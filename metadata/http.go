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
			if pk := s.SimplePK(); pk != "" {
				return fmt.Sprintf("%s/{%s}", path, pk)
			}
		}
	default:
		name := strings.TrimPrefix(s.Name, s.Owner+"s")
		name = strings.TrimPrefix(name, s.Owner)
		name = strings.TrimPrefix(name, "By")
		if pk := s.SimplePK(); pk != "" && !s.IsMethod {
			if name == pk {
				name = ""
			}
		}
		path = path + "/" + toKebabCase(name)
	}
	method := s.HttpMethod()

	if method == "get" && !s.HasCustomParams() {
		if len(s.InputNames) == 1 && !s.HasArrayParams() {
			path = fmt.Sprintf("%s/{%s}", strings.TrimSuffix(path, "/"), s.InputNames[0])
		} else if len(s.InputMethodNames) == 1 {
			path = fmt.Sprintf("%s/{%s}", strings.TrimSuffix(path, "/"), s.InputMethodNames[0])

		}
	}
	return path
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
