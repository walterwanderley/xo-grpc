package metadata

import (
	"fmt"
	"go/ast"
	"regexp"
	"strings"
	"unicode"
)

func exprToStr(e ast.Expr) string {
	switch exp := e.(type) {
	case *ast.SelectorExpr:
		return fmt.Sprintf("%s.%s", exprToStr(exp.X), exp.Sel.Name)
	case *ast.Ident:
		return exp.String()
	case *ast.StarExpr:
		return "*" + exprToStr(exp.X)
	case *ast.ArrayType:
		return "[]" + exprToStr(exp.Elt)
	default:
		panic(fmt.Sprintf("invalid type %T", exp))
	}
}

func toProtoType(typ string) string {
	if strings.HasPrefix(typ, "*") {
		return toProtoType(typ[1:])
	}
	if strings.HasPrefix(typ, "[]") && typ != "[]byte" {
		return "repeated " + toProtoType(typ[2:])
	}
	switch typ {
	case "json.RawMessage", "[]byte":
		return "bytes"
	case "sql.NullBool":
		return "google.protobuf.BoolValue"
	case "sql.NullInt32":
		return "google.protobuf.Int32Value"
	case "int":
		return "int64"
	case "int16":
		return "int32"
	case "uint16":
		return "uint32"
	case "sql.NullInt64":
		return "google.protobuf.Int64Value"
	case "float32":
		return "float"
	case "float64":
		return "double"
	case "sql.NullFloat64":
		return "google.protobuf.DoubleValue"
	case "sql.NullString":
		return "google.protobuf.StringValue"
	case "sql.NullTime", "time.Time", "pq.NullTime", "mysql.NullTime", "xoutil.SqTime":
		return "google.protobuf.Timestamp"
	case "uuid.UUID", "net.HardwareAddr", "net.IP":
		return "string"
	default:
		if firstIsUpper(typ) {
			return "typespb." + typ
		}

		if strings.Contains(typ, textUnmarshalerTypePrefix) || strings.Contains(typ, parserTypePrefix) {
			return "string"
		}
		return typ
	}
}

func bindToProto(src, dst, attrName, attrType string) []string {
	isArray := strings.HasPrefix(attrType, "[]")
	attrType = canonicalType(attrType)
	res := make([]string, 0)
	switch attrType {
	case "sql.NullBool":
		if isArray {
			res = append(res, bindToProtoWrappersArray(src, dst, attrName, "Bool")...)
			break
		}
		res = append(res, fmt.Sprintf("if %s.%s.Valid {", src, attrName))
		res = append(res, fmt.Sprintf("%s.%s = wrapperspb.Bool(%s.%s.Bool) }", dst, attrName, src, attrName))
	case "sql.NullInt32":
		if isArray {
			res = append(res, bindToProtoWrappersArray(src, dst, attrName, "Int32")...)
			break
		}
		res = append(res, fmt.Sprintf("if %s.%s.Valid {", src, attrName))
		res = append(res, fmt.Sprintf("%s.%s = wrapperspb.Int32(%s.%s.Int32) }", dst, attrName, src, attrName))
	case "sql.NullInt64":
		if isArray {
			res = append(res, bindToProtoWrappersArray(src, dst, attrName, "Int64")...)
			break
		}
		res = append(res, fmt.Sprintf("if %s.%s.Valid {", src, attrName))
		res = append(res, fmt.Sprintf("%s.%s = wrapperspb.Int64(%s.%s.Int64) }", dst, attrName, src, attrName))
	case "sql.NullFloat64":
		if isArray {
			res = append(res, fmt.Sprintf("%s.%s = make([]*wrapperspb.DoubleValue, 0)", dst, attrName))
			res = append(res, fmt.Sprintf("for _, item := range %s.%s {", src, attrName))
			res = append(res, "if item.Valid {")
			res = append(res, fmt.Sprintf("%s.%s = append(%s.%s, &wrapperspb.DoubleValue{Value: item.Float64})", dst, attrName, dst, attrName))
			res = append(res, "} else {")
			res = append(res, fmt.Sprintf("%s.%s = append(%s.%s, nil)", dst, attrName, dst, attrName))
			res = append(res, "}")
			res = append(res, "}")
			break
		}
		res = append(res, fmt.Sprintf("if %s.%s.Valid {", src, attrName))
		res = append(res, fmt.Sprintf("%s.%s = wrapperspb.Double(%s.%s.Float64) }", dst, attrName, src, attrName))
	case "sql.NullString":
		if isArray {
			res = append(res, bindToProtoWrappersArray(src, dst, attrName, "String")...)
			break
		}
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
		if strings.Contains(attrType, textUnmarshalerTypePrefix) || strings.Contains(attrType, parserTypePrefix) {
			res = append(res, fmt.Sprintf("%s.%s = %s.%s.String()", dst, attrName, src, attrName))
		} else {
			res = append(res, fmt.Sprintf("%s.%s = %s.%s", dst, attrName, src, attrName))
		}
	}
	return res
}

func bindToProtoWrappersArray(src, dst, attrName, typ string) []string {
	res := make([]string, 0)
	res = append(res, fmt.Sprintf("%s.%s = make([]*wrapperspb.%sValue, 0)", dst, attrName, typ))
	res = append(res, fmt.Sprintf("for _, item := range %s.%s {", src, attrName))
	res = append(res, "if item.Valid {")
	res = append(res, fmt.Sprintf("%s.%s = append(%s.%s, &wrapperspb.%sValue{Value: item.%s})", dst, attrName, dst, attrName, typ, typ))
	res = append(res, "} else {")
	res = append(res, fmt.Sprintf("%s.%s = append(%s.%s, nil)", dst, attrName, dst, attrName))
	res = append(res, "}")
	res = append(res, "}")
	return res
}

func bindToGo(src, dst, attrName, attrType string, newVar bool) []string {
	isArray := strings.HasPrefix(attrType, "[]")
	attrType = canonicalType(attrType)
	res := make([]string, 0)
	switch attrType {
	case "sql.NullBool":
		if newVar {
			res = append(res, fmt.Sprintf("var %s %s", dst, attrType))
		}
		if isArray {
			res = append(res, bindToGoWrappersArray(src, dst, attrName, attrType, "Bool")...)
			break
		}
		res = append(res, fmt.Sprintf("if v := %s.Get%s(); v != nil {", src, attrName))
		res = append(res, fmt.Sprintf("%s = sql.NullBool{Valid: true, Bool: v.Value}", dst))
		res = append(res, "}")
	case "sql.NullInt32":
		if newVar {
			res = append(res, fmt.Sprintf("var %s %s", dst, attrType))
		}
		if isArray {
			res = append(res, bindToGoWrappersArray(src, dst, attrName, attrType, "Int32")...)
			break
		}
		res = append(res, fmt.Sprintf("if v := %s.Get%s(); v != nil {", src, attrName))
		res = append(res, fmt.Sprintf("%s = sql.NullInt32{Valid: true, Int32: v.Value}", dst))
		res = append(res, "}")
	case "sql.NullInt64":
		if newVar {
			res = append(res, fmt.Sprintf("var %s %s", dst, attrType))
		}
		if isArray {
			res = append(res, bindToGoWrappersArray(src, dst, attrName, attrType, "Int64")...)
			break
		}
		res = append(res, fmt.Sprintf("if v := %s.Get%s(); v != nil {", src, attrName))
		res = append(res, fmt.Sprintf("%s = sql.NullInt64{Valid: true, Int64: v.Value}", dst))
		res = append(res, "}")
	case "sql.NullFloat64":
		if newVar {
			res = append(res, fmt.Sprintf("var %s %s", dst, attrType))
		}
		if isArray {
			res = append(res, bindToGoWrappersArray(src, dst, attrName, attrType, "Float64")...)
			break
		}
		res = append(res, fmt.Sprintf("if v := %s.Get%s(); v != nil {", src, attrName))
		res = append(res, fmt.Sprintf("%s = sql.NullFloat64{Valid: true, Float64: v.Value}", dst))
		res = append(res, "}")
	case "sql.NullString":
		if newVar {
			res = append(res, fmt.Sprintf("var %s %s", dst, attrType))
		}
		if isArray {
			res = append(res, bindToGoWrappersArray(src, dst, attrName, attrType, "String")...)
			break
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
		res = append(res, "}")
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
		switch {
		case strings.Contains(attrType, textUnmarshalerTypePrefix):
			attrType = strings.ReplaceAll(attrType, textUnmarshalerTypePrefix, "")
			if newVar {
				res = append(res, fmt.Sprintf("%s := new(%s)", dst, attrType))
			}
			res = append(res, fmt.Sprintf("if err = %s.UnmarshalText([]byte(%s.Get%s())); err != nil {", dst, src, attrName))
			res = append(res, fmt.Sprintf("err = fmt.Errorf(\"invalid %s: %%s%%w\", err.Error(), validation.ErrUserInput)", attrName))
			res = append(res, "return }")
		case strings.Contains(attrType, parserTypePrefix):
			attrType = strings.ReplaceAll(attrType, parserTypePrefix, "")
			if newVar {
				res = append(res, fmt.Sprintf("%s := new(%s)", dst, attrType))
			}
			res = append(res, fmt.Sprintf("if err = %s.Parse(%s.Get%s()); err != nil {", dst, src, attrName))
			res = append(res, fmt.Sprintf("err = fmt.Errorf(\"invalid %s: %%s%%w\", err.Error(), validation.ErrUserInput)", attrName))
			res = append(res, "return }")
		default:
			if newVar {
				res = append(res, fmt.Sprintf("%s := %s.Get%s()", dst, src, attrName))
			} else {
				res = append(res, fmt.Sprintf("%s = %s.Get%s()", dst, src, attrName))
			}
		}
	}
	return res
}

func bindToGoWrappersArray(src, dst, attrName, attrType, valueType string) []string {
	res := make([]string, 0)
	res = append(res, fmt.Sprintf("%s = make([]%s, 0)", dst, attrType))
	res = append(res, fmt.Sprintf("for _, item := range %s.Get%s() {", src, attrName))
	res = append(res, fmt.Sprintf("%s = append(%s, %s{Valid: item != nil, %s: item.GetValue()})", dst, dst, attrType, valueType))
	res = append(res, "}")
	return res
}

func UpperFirstCharacter(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return str
}

func lowerFirstCharacter(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return str
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func toKebabCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}-${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}-${2}")
	return strings.ToLower(snake)
}
