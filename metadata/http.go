package metadata

import (
	"fmt"
	"strings"
)

func (s *Service) HttpMethod() string {
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

func (s *Service) HttpPath() string {
	path := "/v1/" + toKebabCase(s.Owner)
	switch s.Name {
	case "Insert":
		return path
	case "Delete", "Update":
		if s.IsMethod {
			return path + s.pkURLParams()
		}
	default:
		if s.RelationshipMethod() {
			if pk := s.SimplePK(); pk != "" {
				return fmt.Sprintf("%s/{%s}/%s", path, pk, toKebabCase(s.Name))
			}
			return path + s.pkURLParams() + "/" + toKebabCase(s.Name)

		}
		if s.IsReadEntity() {
			return path + s.pkURLParams()
		}

		name := strings.TrimPrefix(s.Name, s.Owner+"sBy")
		name = strings.TrimPrefix(name, s.Owner+"By")
		path = path + "/" + toKebabCase(name)
	}
	method := s.HttpMethod()

	if method == "get" && !s.HasCustomParams() && !s.HasArrayParams() {
		if len(s.InputNames) == 1 {
			path = fmt.Sprintf("%s/{%s}", strings.TrimSuffix(path, "/"), UpperFirstCharacter(s.InputNames[0]))
		} else if len(s.InputMethodNames) == 1 {
			path = fmt.Sprintf("%s/{%s}", strings.TrimSuffix(path, "/"), UpperFirstCharacter(s.InputMethodNames[0]))

		}
	}
	return path
}

func (s *Service) pkURLParams() string {
	if pk := s.SimplePK(); pk != "" {
		return fmt.Sprintf("/{%s}", pk)
	}
	var buf strings.Builder
	for _, attr := range s.PK() {
		buf.WriteString(fmt.Sprintf("/%s/{%s}", strings.TrimSuffix(toKebabCase(attr), "-id"), attr))
	}
	return buf.String()

}

func (s *Service) HttpBody() string {
	switch s.HttpMethod() {
	case "get", "delete":
		return ""
	default:
		return "*"
	}
}

func (s *Service) HttpResponseBody() string {
	if s.HasArrayOutput() {
		return "value"
	}
	return ""
}

func (s *Service) HttpOptions() []string {
	res := make([]string, 0)
	res = append(res, fmt.Sprintf("%s: \"%s\"", s.HttpMethod(), s.HttpPath()))
	body := s.HttpBody()
	if body != "" {
		res = append(res, fmt.Sprintf("body: \"%s\"", body))
	}
	responseBody := s.HttpResponseBody()
	if responseBody != "" {
		res = append(res, fmt.Sprintf("response_body: \"%s\"", responseBody))
	}
	return res
}
