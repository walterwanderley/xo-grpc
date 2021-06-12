package metadata

import (
	"fmt"
	"go/ast"
	"sort"
	"strings"
)

func analyseFunc(fun *ast.FuncDecl, messages map[string]*Message) (owner string, srv *Service) {
	if !isMethodValid(fun) {
		return
	}

	srv = &Service{
		Name:       fun.Name.String(),
		InputNames: make([]string, 0),
		InputTypes: make([]string, 0),
		Output:     make([]string, 0),
		IsMethod:   fun.Recv != nil && len(fun.Recv.List) > 0,
	}
	// context is the first parameter and DB is the second parameter
	for i := 0; i < len(fun.Type.Params.List); i++ {
		p := fun.Type.Params.List[i]
		if exprToStr(p.Type) == "context.Context" {
			srv.HasContext = true
			continue
		}
		if exprToStr(p.Type) == "DB" {
			continue
		}

		for _, n := range p.Names {
			srv.InputNames = append(srv.InputNames, n.Name)
			srv.InputTypes = append(srv.InputTypes, adjustType(exprToStr(p.Type), messages))
		}
	}

	// error is the last result
	for i := 0; i < len(fun.Type.Results.List)-1; i++ {
		p := fun.Type.Results.List[i]
		srv.Output = append(srv.Output, adjustType(exprToStr(p.Type), messages))
	}

	owner = strings.TrimPrefix(strings.TrimPrefix(getOwner(fun), "[]"), "*")
	if srv.IsMethod {
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

		if srv.Name == "Delete" {
			receiver.PkNames = make([]string, 0)
			receiver.PkNames = append(receiver.PkNames, methodAttributes...)

		}
		srv.InputMethodNames = methodAttributes
		srv.InputMethodTypes = methodTypes
	}
	srv.Owner = owner

	return
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

	if exprToStr(fun.Type.Results.List[len(fun.Type.Results.List)-1].Type) != "error" {
		return false
	}

	for _, param := range fun.Type.Params.List {
		if t, ok := param.Type.(*ast.Ident); ok && t.Name == "DB" {
			return true
		}
	}

	return false
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

func isParseFromString(fun *ast.FuncDecl) (receiver string, ok bool) {
	if fun.Name.Name != "Parse" {
		return
	}

	if fun.Recv == nil || len(fun.Recv.List) != 1 ||
		fun.Type.Params == nil || len(fun.Type.Params.List) != 1 ||
		fun.Type.Results == nil || len(fun.Type.Results.List) != 1 {
		return
	}

	if exprToStr(fun.Type.Params.List[0].Type) != "string" {
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
