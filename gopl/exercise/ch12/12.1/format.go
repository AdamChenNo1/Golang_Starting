/*
 * File: \exercise\ch12\12.1\format.go                                         *
 * Project: Golang_Starting                                                    *
 * Created At: Monday, 2022/05/23 , 18:57:53                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Monday, 2022/05/23 , 20:24:49                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package display

import (
	"reflect"
	"strconv"
	"strings"
)

func Format(v reflect.Value) string {
	return formatAtom(v)
}

func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'f', -1, 64)
	case reflect.Complex64, reflect.Complex128:
		return strconv.FormatComplex(v.Complex(), 'f', -1, 64)
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr:
		return v.Type().String() + " 0x" + strconv.FormatUint(uint64(v.Pointer()), 16)
	case reflect.Struct:
		var sb strings.Builder
		sb.WriteString(v.Type().String())
		sb.WriteRune('{')
		for i, l := 0, v.NumField(); i < l; i++ {
			sb.WriteString(v.Type().Field(i).Name)
			sb.WriteRune(':')
			sb.WriteString(Format(v.Field(i)))
			if i != l-1 {
				sb.WriteRune(',')
			}
		}
		sb.WriteRune('}')
		return sb.String()
	case reflect.Array, reflect.Slice:
		var sb strings.Builder
		sb.WriteString(v.Type().String())
		sb.WriteRune('{')
		for i, l := 0, v.Len(); i < l; i++ {
			sb.WriteString(Format(v.Index(i)))
			if i != l-1 {
				sb.WriteRune(',')
			}
		}
		sb.WriteRune('}')
		return sb.String()
	case reflect.Map:
		var sb strings.Builder
		sb.WriteString(v.Type().String())
		sb.WriteRune('{')
		keys := v.MapKeys()
		l := v.Len()
		for i, key := range keys {
			sb.WriteString(Format(key))
			sb.WriteRune(':')
			sb.WriteString(Format(v.MapIndex(key)))
			if i != l-1 {
				sb.WriteRune(',')
			}
		}
		sb.WriteRune('}')
		return sb.String()
	default:
		return v.Type().String() + " value"
	}
}
