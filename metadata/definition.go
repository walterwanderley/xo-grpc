package metadata

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"sort"
	"strings"
)

type Definition struct {
	GoModule             string
	DatabaseDriverModule string
	DatabaseDriverName   string
	ModelsPath           string
	Packages             []*Package
}

func (d *Definition) ProtoImports() []string {
	r := make([]string, 0)
	if d.importTimestamp() {
		r = append(r, `import "google/protobuf/timestamp.proto";`)
	}
	if d.importWrappers() {
		r = append(r, `import "google/protobuf/wrappers.proto";`)
	}
	return r
}

func (d *Definition) Messages() map[string]*Message {
	if len(d.Packages) > 0 {
		return d.Packages[0].Messages
	}
	return nil
}

func (d *Definition) importTimestamp() bool {
	for _, m := range d.Messages() {
		if m.importTimestamp() {
			return true
		}
	}

	return false
}

func (d *Definition) importWrappers() bool {
	for _, m := range d.Messages() {
		if m.importWrappers() {
			return true
		}
	}

	return false
}

type Package struct {
	Package    string
	GoModule   string
	SrcPath    string
	SrcPackage string
	Services   []*Service
	Messages   map[string]*Message
}

func (p *Package) ProtoImports() []string {
	r := make([]string, 0)
	if p.importEmpty() {
		r = append(r, `import "google/protobuf/empty.proto";`)
	}
	if p.importTimestamp() {
		r = append(r, `import "google/protobuf/timestamp.proto";`)
	}
	if p.importWrappers() {
		r = append(r, `import "google/protobuf/wrappers.proto";`)
	}
	return r
}

func (p *Package) importEmpty() bool {
	for _, s := range p.Services {
		if s.EmptyInput() || s.EmptyOutput() {
			return true
		}
	}
	return false
}

func (p *Package) importTimestamp() bool {
	for _, s := range p.Services {
		for _, n := range s.InputTypes {
			if n == "time.Time" || strings.HasSuffix(n, ".NullTime") {
				return true
			}
		}
		for _, n := range s.InputMethodTypes {
			if n == "time.Time" || strings.HasSuffix(n, ".NullTime") {
				return true
			}
		}
		for _, n := range s.Output {
			if n == "time.Time" || strings.HasSuffix(n, ".NullTime") {
				return true
			}
		}
	}
	return false
}

func (p *Package) importWrappers() bool {
	for _, s := range p.Services {
		for _, n := range s.InputTypes {
			if strings.HasPrefix(n, "sql.Null") && !strings.HasSuffix(n, ".NullTime") {
				return true
			}
		}
		for _, n := range s.InputMethodTypes {
			if strings.HasPrefix(n, "sql.Null") && !strings.HasSuffix(n, ".NullTime") {
				return true
			}
		}
		for _, n := range s.Output {
			if strings.HasPrefix(n, "sql.Null") && !strings.HasSuffix(n, ".NullTime") {
				return true
			}
		}
	}
	return false
}

func ParsePackages(src, module string) ([]*Package, error) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, src, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}
	if total := len(pkgs); total != 1 {
		return nil, fmt.Errorf("too many packages: %d", total)
	}

	var pkgName string
	var pkg *ast.Package
	for pkgName, pkg = range pkgs {
		break
	}

	messages := parseMessages(pkg)

	owners := make(map[string][]*Service)
	for _, file := range pkg.Files {
		for _, n := range file.Decls {
			if fun, ok := n.(*ast.FuncDecl); ok {
				owner, srv := analyseFunc(fun, messages)
				if srv != nil {
					srv.Messages = messages
					if _, ok := owners[owner]; !ok {
						owners[owner] = make([]*Service, 0)
					}
					owners[owner] = append(owners[owner], srv)
				}
			}
		}
	}
	result := make([]*Package, 0)
	for owner, services := range owners {
		sort.SliceStable(services, func(i, j int) bool {
			return strings.Compare(services[i].Name, services[j].Name) < 0
		})
		p := Package{
			Package:    owner,
			SrcPath:    src,
			SrcPackage: pkgName,
			GoModule:   module,
			Messages:   messages,
			Services:   services,
		}
		configReadEntity(&p)
		result = append(result, &p)
	}
	sort.SliceStable(result, func(i, j int) bool {
		return strings.Compare(result[i].Package, result[j].Package) < 0
	})
	return result, nil
}

func configReadEntity(p *Package) {
	m, ok := p.Messages[p.Package]
	if !ok {
		return
	}

	for _, s := range p.Services {
		if len(s.Output) != 1 {
			continue
		}
		if m.Name != strings.TrimPrefix(s.Output[0], "*") {
			continue
		}
		if !strings.HasPrefix(s.Name, p.Package+"By") {
			continue
		}
		if len(m.PkNames) != len(s.InputNames) {
			continue
		}
		var incompatibleInterface bool
		for _, pk := range m.PkNames {
			pkLower := strings.ToLower(pk)
			var found bool
			for _, in := range s.InputNames {
				if strings.ToLower(in) == pkLower {
					found = true
					break
				}
			}
			if !found {
				incompatibleInterface = true
				break
			}
		}
		if incompatibleInterface {
			continue
		}
		m.ReaderService = s
		return
	}
}
