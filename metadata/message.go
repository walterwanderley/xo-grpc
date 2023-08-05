package metadata

import (
	"fmt"
	"go/ast"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Message struct {
	Name               string
	PkNames            []string
	IndexNames         map[string]struct{}
	AttrNames          []string
	AttrTypes          []string
	IsArray            bool
	ElementType        string
	HasTextUnmarshaler bool
	HasParser          bool
	ReaderService      *Service
}

func (m *Message) ProtoAttributes() string {
	var s strings.Builder
	for i, name := range m.AttrNames {
		s.WriteString(fmt.Sprintf("    %s %s = %d;\n", toProtoType(m.AttrTypes[i]), ToSnakeCase(name), i+1))
	}
	return s.String()
}

func (m *Message) entityParentPK() (entity []string, parent []string) {
	for _, pk := range m.PkNames {
		found := false
		for idx := range m.IndexNames {
			if idx == pk {
				found = true
				break
			}
		}
		if found {
			parent = append(parent, pk)
			continue
		}
		entity = append(entity, pk)
	}
	return
}

func (m *Message) attributeTypeByName(attrName string) string {
	for i, n := range m.AttrNames {
		if n == attrName {
			return m.AttrTypes[i]
		}
	}
	return ""
}

func (m *Message) adjustType(messages map[string]*Message) {
	for i, t := range m.AttrTypes {
		m.AttrTypes[i] = adjustType(t, messages)
	}
}

func (m *Message) importTimestamp() bool {
	for _, typ := range m.AttrTypes {
		if typ == "time.Time" || strings.HasSuffix(typ, ".NullTime") || typ == "*Time" {
			return true
		}
	}

	return false
}

func (m *Message) importWrappers() bool {
	for _, typ := range m.AttrTypes {
		if strings.HasPrefix(typ, "sql.Null") && !strings.HasSuffix(typ, ".NullTime") {
			return true
		}
	}

	return false
}

func createStructMessage(name string, s *ast.StructType) *Message {
	names := make([]string, 0)
	types := make([]string, 0)
	for _, f := range s.Fields.List {
		if len(f.Names) == 0 || !firstIsUpper(f.Names[0].Name) {
			continue
		}
		types = append(types, exprToStr(f.Type))
		names = append(names, f.Names[0].Name)
	}
	return &Message{
		Name:      name,
		AttrNames: names,
		AttrTypes: types,
	}
}

func createArrayMessage(name string, s *ast.ArrayType) *Message {
	return &Message{
		Name:        name,
		IsArray:     true,
		ElementType: exprToStr(s.Elt),
	}
}

func createAliasMessage(name string, s *ast.Ident) *Message {
	return &Message{
		Name:        name,
		ElementType: exprToStr(s),
	}
}

func customType(typ string) bool {
	typ = strings.TrimPrefix(typ, "*")
	return firstIsUpper(typ)
}

func firstIsUpper(s string) bool {
	ru, _ := utf8.DecodeRuneInString(s[0:1])
	return unicode.IsUpper(ru)
}

func adjustType(typ string, messages map[string]*Message) string {
	if m, ok := messages[typ]; ok {
		var prefix string
		if m.IsArray {
			prefix = "[]"
		}
		switch {
		case m.HasTextUnmarshaler:
			return fmt.Sprintf("%s%s.%s", prefix, textUnmarshalerTypePrefix, typ)
		case m.HasParser:
			return fmt.Sprintf("%s%s.%s", parserTypePrefix, prefix, typ)
		case m.ElementType != "":
			return prefix + m.ElementType
		}
	}

	return typ
}

var errTypes = []string{"ErrDecodeFailed", "ErrInsertFailed", "ErrUpdateFailed", "ErrUpsertFailed"}

func isErrType(name string) bool {
	for _, n := range errTypes {
		if n == name {
			return true
		}
	}
	return false
}
