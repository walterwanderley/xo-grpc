package metadata

import (
	"fmt"
	"strings"
)

func (s *Service) HttpOptions() []string {
	res := make([]string, 0)
	res = append(res, fmt.Sprintf("%s: \"%s\"", s.httpMethod(), s.httpPath()))
	body := s.httpBody()
	if body != "" {
		res = append(res, fmt.Sprintf("body: \"%s\"", body))
	}
	responseBody := s.httpResponseBody()
	if responseBody != "" {
		res = append(res, fmt.Sprintf("response_body: \"%s\"", responseBody))
	}
	return res
}

func (s *Service) httpMethod() string {
	switch s.Name {
	case "Insert", "Upsert":
		return "post"
	case "Delete":
		return "delete"
	case "Update":
		return "put"
	default:
		if s.EmptyOutput() {
			return "post"
		}
		return "get"
	}
}

func (s *Service) httpPath() string {
	path := "/v1"
	switch s.Name {
	case "Insert", "Delete", "Update":
		return path + s.pkURLParams()
	default:
		if s.RelationshipMethod() {
			return path + s.pkURLParams() + "/" + toKebabCase(s.Name)
		}
		if s.isReaderEntity() {
			return path + s.pkURLParams()
		}

		name := strings.TrimPrefix(s.Name, s.Owner+"sBy")
		name = strings.TrimPrefix(name, s.Owner+"By")
		path = path + "/" + toKebabCase(s.Owner) + "/" + trimParentPath(toKebabCase(name))
	}
	method := s.httpMethod()

	if method == "get" && !s.HasCustomParams() && !s.hasArrayParams() {
		if len(s.InputNames) == 1 {
			path = fmt.Sprintf("%s/{%s}", strings.TrimSuffix(path, "/"), ToSnakeCase(s.InputNames[0]))
		} else if len(s.InputMethodNames) == 1 {
			path = fmt.Sprintf("%s/{%s}", strings.TrimSuffix(path, "/"), ToSnakeCase(s.InputMethodNames[0]))

		}
	}
	return path
}

func (s *Service) pkURLParams() string {
	if pk := s.SimplePK(); pk != "" && s.Name != "Insert" {
		return fmt.Sprintf("/%s/{%s}", toKebabCase(s.Owner), ToSnakeCase(s.inputCase(pk)))
	}
	var buf strings.Builder
	entity, parent := s.entityPKParentPK()
	for _, attr := range parent {
		buf.WriteString(fmt.Sprintf("/%s/{%s}", trimParentPath(toKebabCase(attr)), ToSnakeCase(s.inputCase(attr))))
	}
	buf.WriteString("/")
	buf.WriteString(toKebabCase(s.Owner))
	if s.Name != "Insert" {
		if len(entity) == 1 {
			buf.WriteString(fmt.Sprintf("/{%s}", ToSnakeCase(s.inputCase(entity[0]))))
		} else {
			for _, attr := range entity {
				buf.WriteString(fmt.Sprintf("/%s/{%s}", trimParentPath(toKebabCase(attr)), ToSnakeCase(s.inputCase(attr))))
			}
		}
	}
	return buf.String()
}

func (s *Service) inputCase(attr string) string {
	attrLower := strings.ToLower(attr)
	for _, input := range s.InputNames {
		if strings.ToLower(input) == attrLower {
			return input
		}
	}
	for _, input := range s.InputMethodNames {
		if strings.ToLower(input) == attrLower {
			return input
		}
	}
	return attr
}

func (s *Service) LocationURIPattern() string {
	if pk := s.SimplePK(); pk != "" {
		return `fmt.Sprintf("/%v", m.` + pk + ")"
	}
	var buf strings.Builder
	entity, _ := s.entityPKParentPK()
	if len(entity) == 0 {
		return `""`
	}
	buf.WriteString(`fmt.Sprintf("`)
	for _, attr := range entity {
		buf.WriteString(fmt.Sprintf(`/%s/%%v`, trimParentPath(toKebabCase(attr))))
	}
	buf.WriteString("\", ")
	buf.WriteString(s.PKEntityParams("m."))
	buf.WriteString(")")
	return buf.String()
}

func (s *Service) httpBody() string {
	switch s.httpMethod() {
	case "get", "delete":
		return ""
	default:
		if s.Name == "Insert" && len(s.InputMethodNames) == len(s.pk()) {
			return ""
		}
		return "*"
	}
}

func (s *Service) httpResponseBody() string {
	if s.hasArrayOutput() {
		return "value"
	}
	return ""
}

func (s *Service) isReaderEntity() bool {
	reader := s.ReaderEntity()
	if reader == nil {
		return false
	}
	return s.Name == reader.Name
}

func trimParentPath(s string) string {
	return strings.TrimSuffix(strings.TrimSuffix(s, "-id"), "-fk")
}
