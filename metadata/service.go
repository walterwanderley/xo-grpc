package metadata

import (
	"fmt"
	"go/ast"
	"sort"
	"strings"
)

const textUnmarshalerTypePrefix = "encoding.TextUnmarshaler."

type Service struct {
	Name             string
	InputNames       []string
	InputTypes       []string
	Output           []string
	IsMethod         bool
	InputMethodNames []string
	InputMethodTypes []string

	Messages map[string]*Message
}

func (s *Service) MethodInputType() string {
	switch {
	case s.EmptyInput():
		return "emptypb.Empty"
	default:
		return fmt.Sprintf("pb.%sRequest", s.Name)
	}
}

func (s *Service) MethodOutputType() string {
	switch {
	case s.EmptyOutput():
		return "emptypb.Empty"
	case s.HasCustomOutput():
		return fmt.Sprintf("typespb.%s", strings.TrimPrefix(s.Output[0], "*"))
	default:
		return fmt.Sprintf("pb.%sResponse", s.Name)
	}
}

func (s *Service) ReturnCallDatabase() string {
	if !s.EmptyOutput() {
		return "result,"
	}
	return ""
}

func (s *Service) ParamsCallDatabase() string {
	if s.EmptyInput() {
		return ""
	}
	return ", " + strings.Join(s.InputNames, ", ")
}

func (s *Service) InputGrpc() []string {
	res := make([]string, 0)
	if s.EmptyInput() {
		return res
	}

	for i, attr := range s.InputMethodNames {
		res = append(res, bindToGo("req", fmt.Sprintf("m.%s", attr), attr, s.InputMethodTypes[i], false)...)
	}

	if s.HasCustomParams() {
		typ := s.InputTypes[0]
		in := s.InputNames[0]
		res = append(res, fmt.Sprintf("var %s %s", in, typ))
		m := s.Messages[typ]
		for i, name := range m.AttrNames {
			attrName := UpperFirstCharacter(name)
			res = append(res, bindToGo("req", fmt.Sprintf("%s.%s", in, attrName), attrName, m.AttrTypes[i], false)...)
		}
	} else {
		for i, n := range s.InputNames {
			res = append(res, bindToGo("req", n, UpperFirstCharacter(n), s.InputTypes[i], true)...)
		}
	}

	return res
}

func (s *Service) OutputGrpc() []string {
	res := make([]string, 0)

	if s.EmptyOutput() {
		return res
	}

	if s.HasArrayOutput() {
		res = append(res, "for _, r := range result {")
		typ := strings.TrimPrefix(strings.TrimPrefix(s.Output[0], "[]"), "*")
		res = append(res, fmt.Sprintf("var item typespb.%s", typ))
		m := s.Messages[typ]
		for i, attr := range m.AttrNames {
			res = append(res, bindToProto("r", "item", UpperFirstCharacter(attr), m.AttrTypes[i])...)
		}
		res = append(res, "res.Value = append(res.Value, &item)")
		res = append(res, "}")
		return res
	}

	if s.HasCustomOutput() {
		for _, n := range s.Output {
			m := s.Messages[strings.TrimPrefix(n, "*")]
			for i, attr := range m.AttrNames {
				res = append(res, bindToProto("result", "res", UpperFirstCharacter(attr), m.AttrTypes[i])...)
			}
		}
		return res
	}
	if !s.EmptyOutput() {
		res = append(res, "res.Value = result")
		return res
	}

	return res
}

func bindToProto(src, dst, attrName, attrType string) []string {
	res := make([]string, 0)
	switch attrType {
	case "sql.NullBool":
		res = append(res, fmt.Sprintf("if %s.%s.Valid {", src, attrName))
		res = append(res, fmt.Sprintf("%s.%s = wrapperspb.Bool(%s.%s.Bool) }", dst, attrName, src, attrName))
	case "sql.NullInt32":
		res = append(res, fmt.Sprintf("if %s.%s.Valid {", src, attrName))
		res = append(res, fmt.Sprintf("%s.%s = wrapperspb.Int32(%s.%s.Int32) }", dst, attrName, src, attrName))
	case "sql.NullInt64":
		res = append(res, fmt.Sprintf("if %s.%s.Valid {", src, attrName))
		res = append(res, fmt.Sprintf("%s.%s = wrapperspb.Int64(%s.%s.Int64) }", dst, attrName, src, attrName))
	case "sql.NullFloat64":
		res = append(res, fmt.Sprintf("if %s.%s.Valid {", src, attrName))
		res = append(res, fmt.Sprintf("%s.%s = wrapperspb.Float64(%s.%s.Float64) }", dst, attrName, src, attrName))
	case "sql.NullString":
		res = append(res, fmt.Sprintf("if %s.%s.Valid {", src, attrName))
		res = append(res, fmt.Sprintf("%s.%s = wrapperspb.String(%s.%s.String) }", dst, attrName, src, attrName))
	case "sql.NullTime", "pq.NullTime", "mysql.NullTime":
		res = append(res, fmt.Sprintf("if %s.%s.Valid {", src, attrName))
		res = append(res, fmt.Sprintf("%s.%s = timestamppb.New(%s.%s.Time) }", dst, attrName, src, attrName))
	case "time.Time":
		res = append(res, fmt.Sprintf("%s.%s = timestamppb.New(%s.%s)", dst, attrName, src, attrName))
	case "xoutil.SqTime":
		res = append(res, fmt.Sprintf("%s.%s = timestamppb.New(%s.%s.Timr)", dst, attrName, src, attrName))
	case "uuid.UUID", "net.HardwareAddr", "net.IP":
		res = append(res, fmt.Sprintf("%s.%s = %s.%s.String()", dst, attrName, src, attrName))
	case "int16":
		res = append(res, fmt.Sprintf("%s.%s = int32(%s.%s)", dst, attrName, src, attrName))
	case "int":
		res = append(res, fmt.Sprintf("%s.%s = int64(%s.%s)", dst, attrName, src, attrName))
	case "uint16":
		res = append(res, fmt.Sprintf("%s.%s = uint32(%s.%s)", dst, attrName, src, attrName))
	default:
		if strings.Contains(attrType, textUnmarshalerTypePrefix) {
			res = append(res, fmt.Sprintf("%s.%s = %s.%s.String()", dst, attrName, src, attrName))
		} else {
			res = append(res, fmt.Sprintf("%s.%s = %s.%s", dst, attrName, src, attrName))
		}
	}
	return res
}

func bindToGo(src, dst, attrName, attrType string, newVar bool) []string {
	res := make([]string, 0)
	switch attrType {
	case "sql.NullBool":
		if newVar {
			res = append(res, fmt.Sprintf("var %s %s", dst, attrType))
		}
		res = append(res, fmt.Sprintf("if v := %s.Get%s(); v != nil {", src, attrName))
		res = append(res, fmt.Sprintf("%s = sql.NullBool{Valid: true, Bool: v.Value}", dst))
		res = append(res, "}")
	case "sql.NullInt32":
		if newVar {
			res = append(res, fmt.Sprintf("var %s %s", dst, attrType))
		}
		res = append(res, fmt.Sprintf("if v := %s.Get%s(); v != nil {", src, attrName))
		res = append(res, fmt.Sprintf("%s = sql.NullInt32{Valid: true, Int32: v.Value}", dst))
		res = append(res, "}")
	case "sql.NullInt64":
		if newVar {
			res = append(res, fmt.Sprintf("var %s %s", dst, attrType))
		}
		res = append(res, fmt.Sprintf("if v := %s.Get%s(); v != nil {", src, attrName))
		res = append(res, fmt.Sprintf("%s = sql.NullInt64{Valid: true, Int64: v.Value}", dst))
		res = append(res, "}")
	case "sql.NullFloat64":
		if newVar {
			res = append(res, fmt.Sprintf("var %s %s", dst, attrType))
		}
		res = append(res, fmt.Sprintf("if v := %s.Get%s(); v != nil {", src, attrName))
		res = append(res, fmt.Sprintf("%s = sql.NullFloat64{Valid: true, Float64: v.Value}", dst))
		res = append(res, "}")
	case "sql.NullString":
		if newVar {
			res = append(res, fmt.Sprintf("var %s %s", dst, attrType))
		}
		res = append(res, fmt.Sprintf("if v := %s.Get%s(); v != nil {", src, attrName))
		res = append(res, fmt.Sprintf("%s = sql.NullString{Valid: true, String: v.Value}", dst))
		res = append(res, "}")
	case "sql.NullTime", "pq.NullTime", "mysql.NullTime":
		if newVar {
			res = append(res, fmt.Sprintf("var %s %s", dst, attrType))
		}
		res = append(res, fmt.Sprintf("if v := %s.Get%s(); v != nil {", src, attrName))
		res = append(res, fmt.Sprintf("if err = v.CheckValid(); err != nil { err = fmt.Errorf(\"invalid %s: %%s%%w\", err.Error(), validation.ErrUserInput)", attrName))
		res = append(res, "return }")
		res = append(res, "if t := v.AsTime(); !t.IsZero() {")
		res = append(res, fmt.Sprintf("%s.Valid = true", dst))
		res = append(res, fmt.Sprintf("%s.Time = t } }", dst))
	case "time.Time":
		if newVar {
			res = append(res, fmt.Sprintf("var %s %s", dst, attrType))
		}
		res = append(res, fmt.Sprintf("if v := %s.Get%s(); v != nil {", src, attrName))
		res = append(res, fmt.Sprintf("if err = v.CheckValid(); err != nil { err = fmt.Errorf(\"invalid %s: %%s%%w\", err.Error(), validation.ErrUserInput)", attrName))
		res = append(res, "return }")
		res = append(res, fmt.Sprintf("%s = v.AsTime()", dst))
		res = append(res, fmt.Sprintf("} else { err = fmt.Errorf(\"the %s attribute is required%%w\", validation.ErrUserInput)", attrName))
		res = append(res, "return }")
	case "xoutil.SqTime":
		if newVar {
			res = append(res, fmt.Sprintf("var %s %s", dst, attrType))
		}
		res = append(res, fmt.Sprintf("if v := %s.Get%s(); v != nil {", src, attrName))
		res = append(res, fmt.Sprintf("if err = v.CheckValid(); err != nil { err = fmt.Errorf(\"invalid %s: %%s%%w\", err.Error(), validation.ErrUserInput)", attrName))
		res = append(res, "return }")
		res = append(res, fmt.Sprintf("%s.Time = v.AsTime()", dst))
		res = append(res, fmt.Sprintf("}"))
	case "uuid.UUID":
		if newVar {
			res = append(res, fmt.Sprintf("var %s %s", dst, attrType))
		}
		res = append(res, fmt.Sprintf("if %s, err = uuid.Parse(%s.Get%s()); err != nil {", dst, src, attrName))
		res = append(res, fmt.Sprintf("err = fmt.Errorf(\"invalid %s: %%s%%w\", err.Error(), validation.ErrUserInput)", attrName))
		res = append(res, "return }")
	case "net.HardwareAddr":
		if newVar {
			res = append(res, fmt.Sprintf("var %s %s", dst, attrType))
		}
		res = append(res, fmt.Sprintf("if %s, err = net.ParseMAC(%s.Get%s()); err != nil {", dst, src, attrName))
		res = append(res, fmt.Sprintf("err = fmt.Errorf(\"invalid %s: %%s%%w\", err.Error(), validation.ErrUserInput)", attrName))
		res = append(res, "return }")
	case "net.IP":
		if newVar {
			res = append(res, fmt.Sprintf("%s := net.ParseIP(%s.Get%s())", dst, src, attrName))
		} else {
			res = append(res, fmt.Sprintf("%s = net.ParseIP(%s.Get%s())", dst, src, attrName))
		}
	case "int16":
		if newVar {
			res = append(res, fmt.Sprintf("%s := int16(%s.Get%s())", dst, src, attrName))
		} else {
			res = append(res, fmt.Sprintf("%s = int16(%s.Get%s())", dst, src, attrName))
		}
	case "int":
		if newVar {
			res = append(res, fmt.Sprintf("%s := int(%s.Get%s())", dst, src, attrName))
		} else {
			res = append(res, fmt.Sprintf("%s = int(%s.Get%s())", dst, src, attrName))
		}
	case "uint16":
		if newVar {
			res = append(res, fmt.Sprintf("%s := uint16(%s.Get%s())", dst, src, attrName))
		} else {
			res = append(res, fmt.Sprintf("%s = uint16(%s.Get%s())", dst, src, attrName))
		}
	default:

		if strings.Contains(attrType, textUnmarshalerTypePrefix) {
			attrType = strings.ReplaceAll(attrType, textUnmarshalerTypePrefix, "")
			if newVar {
				res = append(res, fmt.Sprintf("%s := new(%s)", dst, attrType))
			}
			res = append(res, fmt.Sprintf("if err = %s.UnmarshalText([]byte(%s.Get%s())); err != nil {", dst, src, attrName))
			res = append(res, fmt.Sprintf("err = fmt.Errorf(\"invalid %s: %%s%%w\", err.Error(), validation.ErrUserInput)", attrName))
			res = append(res, "return }")
		} else {
			if newVar {
				res = append(res, fmt.Sprintf("%s := %s.Get%s()", dst, src, attrName))
			} else {
				res = append(res, fmt.Sprintf("%s = %s.Get%s()", dst, src, attrName))
			}
		}
	}
	return res
}

func (s *Service) RpcSignature() string {
	var b strings.Builder
	b.WriteString(s.Name)
	b.WriteString("(")
	switch {
	case s.EmptyInput():
		b.WriteString("google.protobuf.Empty")
	default:
		b.WriteString(fmt.Sprintf("%sRequest", s.Name))
	}
	b.WriteString(") returns (")
	switch {
	case s.EmptyOutput():
		b.WriteString("google.protobuf.Empty")
	case s.HasCustomOutput():
		b.WriteString("typespb." + strings.TrimPrefix(s.Output[0], "*"))
	default:
		b.WriteString(fmt.Sprintf("%sResponse", s.Name))
	}
	b.WriteString(")")
	return b.String()
}

func (s *Service) HasCustomParams() bool {
	if s.EmptyInput() || len(s.InputTypes) == 0 {
		return false
	}

	return customType(s.InputTypes[0])
}

func (s *Service) HasCustomOutput() bool {
	if s.EmptyOutput() {
		return false
	}

	return customType(s.Output[0])
}

func (s *Service) HasArrayOutput() bool {
	if s.EmptyOutput() {
		return false
	}
	return strings.HasPrefix(s.Output[0], "[]") && s.Output[0] != "[]byte"
}

func (s *Service) ProtoInputs() string {
	var b strings.Builder
	var count int
	for i, name := range s.InputNames {
		count = count + 1
		fmt.Fprintf(&b, "\n    %s %s = %d;", toProtoType(s.InputTypes[i]), name, count)
	}
	for i, name := range s.InputMethodNames {
		count = count + 1
		fmt.Fprintf(&b, "\n    %s %s = %d;", toProtoType(s.InputMethodTypes[i]), name, count)
	}
	return b.String()
}

func (s *Service) EmptyInput() bool {
	if len(s.InputTypes) > 0 || len(s.InputMethodTypes) > 0 {
		return false
	}
	return true
}

func (s *Service) EmptyOutput() bool {
	return len(s.Output) == 0
}

func (s *Service) ProtoOutputs() string {
	var b strings.Builder
	for i, name := range s.Output {
		fmt.Fprintf(&b, "    %s value = %d;\n", toProtoType(name), i+1)

	}
	return b.String()
}

func analyseFunc(fun *ast.FuncDecl, messages map[string]*Message) (owner string, srv *Service) {
	if !isMethodValid(fun) {
		return
	}

	inputNames := make([]string, 0)
	inputTypes := make([]string, 0)
	output := make([]string, 0)

	// XODB is the first parameter
	for i := 1; i < len(fun.Type.Params.List); i++ {
		p := fun.Type.Params.List[i]
		inputNames = append(inputNames, p.Names[0].Name)
		inputTypes = append(inputTypes, checkAliasType(exprToStr(p.Type), messages))
	}

	// error is the last result
	for i := 0; i < len(fun.Type.Results.List)-1; i++ {
		p := fun.Type.Results.List[0]
		output = append(output, checkAliasType(exprToStr(p.Type), messages))
	}

	owner = strings.TrimPrefix(strings.TrimPrefix(getOwner(fun), "[]"), "*")
	srv = &Service{
		Name:       fun.Name.String(),
		InputNames: inputNames,
		InputTypes: inputTypes,
		Output:     output,
		IsMethod:   fun.Recv != nil && len(fun.Recv.List) > 0,
	}
	isMethod := fun.Recv != nil && len(fun.Recv.List) > 0
	if isMethod {
		receiverName := fun.Recv.List[0].Names[0].Name
		methodInputRequirements := make(map[string]struct{})
		for _, stmt := range fun.Body.List {
			ast.Inspect(stmt, func(n ast.Node) bool {
				if selector, ok := n.(*ast.SelectorExpr); ok && fmt.Sprintf("%s", selector.X) == receiverName && firstIsUpper(selector.Sel.Name) {
					methodInputRequirements[selector.Sel.Name] = struct{}{}
					return false
				}
				// Ignore passing by reference
				if _, ok := n.(*ast.UnaryExpr); ok {
					return false
				}
				return true
			})
		}
		methodAttributes := make([]string, 0)
		for name := range methodInputRequirements {
			methodAttributes = append(methodAttributes, name)
		}
		receiverType := strings.TrimPrefix(exprToStr(fun.Recv.List[0].Type), "*")
		receiver := messages[receiverType]
		sort.Strings(methodAttributes)
		methodTypes := make([]string, 0)
		for _, n := range methodAttributes {
			methodTypes = append(methodTypes, receiver.AttributeTypeByName(n))
		}
		srv.InputMethodNames = methodAttributes
		srv.InputMethodTypes = methodTypes
	}

	return
}

func checkAliasType(typ string, messages map[string]*Message) string {
	if m, ok := messages[typ]; ok && m.ElementType != "" {
		var prefix string
		if m.IsArray {
			prefix = "[]"
		}
		if m.IsTextUnmarshaler {
			return fmt.Sprintf("%sencoding.TextUnmarshaler.%s", prefix, typ)
		}
		return prefix + m.ElementType
	}

	return typ
}

func getOwner(fun *ast.FuncDecl) string {
	if fun.Recv != nil && len(fun.Recv.List) > 0 {
		return exprToStr(fun.Recv.List[0].Type)
	}

	if len(fun.Type.Results.List) > 1 {
		return exprToStr(fun.Type.Results.List[0].Type)
	}
	return ""
}

func isMethodValid(fun *ast.FuncDecl) bool {
	if fun.Name == nil {
		return false
	}

	if !fun.Name.IsExported() || fun.Name.Name == "Save" {
		return false
	}

	if fun.Type.Params == nil || len(fun.Type.Params.List) == 0 ||
		fun.Type.Results == nil || len(fun.Type.Results.List) == 0 {
		return false
	}

	if t, ok := fun.Type.Params.List[0].Type.(*ast.Ident); !ok || t.Name != "XODB" {
		return false
	}

	if exprToStr(fun.Type.Results.List[len(fun.Type.Results.List)-1].Type) != "error" {
		return false
	}

	return true
}

func isTextUnmarshaler(fun *ast.FuncDecl) (receiver string, ok bool) {
	if fun.Name.Name != "UnmarshalText" {
		return
	}

	if fun.Recv == nil || len(fun.Recv.List) != 1 ||
		fun.Type.Params == nil || len(fun.Type.Params.List) != 1 ||
		fun.Type.Results == nil || len(fun.Type.Results.List) != 1 {
		return
	}

	if exprToStr(fun.Type.Params.List[0].Type) != "[]byte" {
		return
	}

	if exprToStr(fun.Type.Results.List[0].Type) != "error" {
		return
	}

	receiver = exprToStr(fun.Recv.List[0].Type)

	if !strings.HasPrefix(receiver, "*") {
		return
	}

	return receiver, true
}
