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
	AttrNames          []string
	AttrTypes          []string
	IsArray            bool
	ElementType        string
	HasTextUnmarshaler bool
	HasParser          bool
}

func (m *Message) ProtoAttributes() string {
	var s strings.Builder
	for i, name := range m.AttrNames {
		s.WriteString(fmt.Sprintf("    %s %s = %d;\n", toProtoType(m.AttrTypes[i]), name, i+1))
	}
	return s.String()
}

func (m *Message) AttributeTypeByName(attrName string) string {
	for i, n := range m.AttrNames {
		if n == attrName {
			return m.AttrTypes[i]
		}
	}
	return ""
}

func (m *Message) AdjustType(messages map[string]*Message) {
	for i, t := range m.AttrTypes {
		m.AttrTypes[i] = adjustType(t, messages)
	}
}

func (m *Message) importTimestamp() bool {
	for _, typ := range m.AttrTypes {
		if typ == "time.Time" || strings.HasSuffix(typ, ".NullTime") {
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

func parseMessages(pkg *ast.Package) map[string]*Message {
	messages := make(map[string]*Message)
	for _, file := range pkg.Files {
		if file.Scope != nil {
			for name, obj := range file.Scope.Objects {
				if isErrType(name) {
					continue
				}
				if typ, ok := obj.Decl.(*ast.TypeSpec); ok {
					switch t := typ.Type.(type) {
					case *ast.Ident:
						messages[name] = createAliasMessage(name, t)
					case *ast.StructType:
						messages[name] = createStructMessage(name, t)
					case *ast.ArrayType:
						messages[name] = createArrayMessage(name, t)
					}
				}
			}
		}
	}
	for _, file := range pkg.Files {
		for _, n := range file.Decls {
			if fun, ok := n.(*ast.FuncDecl); ok {
				if r, ok := isTextUnmarshaler(fun); ok {
					r = strings.TrimPrefix(r, "*")
					if m, ok := messages[r]; ok {
						m.HasTextUnmarshaler = true
					}
				} else if r, ok := isParseFromString(fun); ok {
					r = strings.TrimPrefix(r, "*")
					if m, ok := messages[r]; ok {
						m.HasParser = true
					}
				}
			}
		}
	}
	for _, m := range messages {
		m.AdjustType(messages)
	}
	return messages
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

var errTypes = []string{"ErrDecodeFailed", "ErrInsertFailed", "ErrUpdateFailed", "ErrUpsertFailed"}

func isErrType(name string) bool {
	for _, n := range errTypes {
		if n == name {
			return true
		}
	}
	return false
}
