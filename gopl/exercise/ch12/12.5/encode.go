/*
 * File: \src\ch12\sexpr\encode.go                                             *
 * Project: Golang_Starting                                                    *
 * Created At: Monday, 2022/05/23 , 20:38:35                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Tuesday, 2022/05/24 , 13:03:06                               *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package json

import (
	"bytes"
	"fmt"
	"reflect"
)

func Marshal(v any) ([]byte, error) {
	var buf bytes.Buffer

	if err := encode(&buf, reflect.ValueOf(v)); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func MarshalIndent(v any) ([]byte, error) {
	var buf bytes.Buffer

	if err := encode_indent(&buf, reflect.ValueOf(v), 0); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func encode(buf *bytes.Buffer, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("")

	case reflect.Bool:
		if v.Bool() {
			buf.WriteString("true")
		} else {
			buf.WriteString("false")
		}

	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64:
		fmt.Fprintf(buf, "%d", v.Uint())

	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(buf, "%v", v.Float())

	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())

	case reflect.Ptr:
		return encode(buf, v.Elem())

	case reflect.Array, reflect.Slice:
		buf.WriteByte('[')
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				buf.WriteString(",")
			}
			if err := encode(buf, v.Index(i)); err != nil {
				return err
			}
		}
		buf.WriteByte(']')

	case reflect.Interface:
		v = v.Elem()
		// fmt.Fprintf(buf, "%q", v.Type().Name())
		encode(buf, v)

	case reflect.Struct:
		buf.WriteByte('{')
		for i, l := 0, v.NumField(); i < l; i++ {
			fmt.Fprintf(buf, "%q:", v.Type().Field(i).Name)
			if err := encode(buf, v.Field(i)); err != nil {
				return err
			}
			if i < l-1 {
				buf.WriteString(", ")
			}
		}
		buf.WriteByte('}')

	case reflect.Map:
		buf.WriteByte('{')
		// keys := v.MapKeys()
		// // sort.Slice(keys)
		// for i, l := 0, v.Len(); i < l; i++ {
		// 	key := keys[i]

		// 	if err := encode(buf, key); err != nil {
		// 		return err
		// 	}

		// 	buf.WriteString(": ")
		// 	if err := encode(buf, v.MapIndex(key)); err != nil {
		// 		return err
		// 	}
		// 	if i < l-1 {
		// 		buf.WriteString(", ")
		// 	}
		// }
		iter := v.MapRange()
		i, l := 0, v.Len()-1
		for iter.Next() {
			key := iter.Key()
			value := iter.Value()
			if err := encode(buf, key); err != nil {
				return err
			}

			buf.WriteString(":")

			if err := encode(buf, value); err != nil {
				return err
			}

			if i < l {
				buf.WriteString(", ")
			}
			i++
		}
		buf.WriteByte('}')
	default:
		return fmt.Errorf("json: unsupported type: %s", v.Type())
	}

	return nil
}

func writeSpace(buf *bytes.Buffer, size uint) {
	for size > 0 {
		buf.WriteByte(' ')
		size--
	}
}

func encode_indent(buf *bytes.Buffer, v reflect.Value, indent uint) error {

	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("nil")

	case reflect.Bool:
		if v.Bool() {
			buf.WriteByte('t')
		} else {
			buf.WriteString("nil")
		}

	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64:
		fmt.Fprintf(buf, "%d", v.Uint())

	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(buf, "%v", v.Float())

	case reflect.Complex64, reflect.Complex128:
		c := v.Complex()
		fmt.Fprintf(buf, "#(%v,%v)", real(c), imag(c))

	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())

	case reflect.Ptr:
		return encode_indent(buf, v.Elem(), indent)

	case reflect.Array, reflect.Slice:
		buf.WriteByte('(')
		indent++

		for i, l := 0, v.Len(); i < l; i++ {
			if i > 0 {
				writeSpace(buf, indent)
			}
			if err := encode_indent(buf, v.Index(i), indent); err != nil {
				return err
			}
			if i < l-1 {
				buf.WriteByte('\n')
			}

		}
		buf.WriteByte(')')

	case reflect.Interface:
		v = v.Elem()
		fmt.Fprintf(buf, "%q", v.Type().Name())
		encode(buf, v)

	case reflect.Struct:
		buf.WriteByte('(')
		indent++
		for i, l := 0, v.NumField(); i < l; i++ {
			if i > 0 {
				writeSpace(buf, indent)
			}
			size, _ := fmt.Fprintf(buf, "(%s ", v.Type().Field(i).Name)

			if err := encode_indent(buf, v.Field(i), indent+uint(size)); err != nil {
				return err
			}

			buf.WriteByte(')')
			if i < l-1 {
				buf.WriteByte('\n')
			}
		}
		buf.WriteByte(')')

	case reflect.Map:
		buf.WriteByte('(')
		indent++
		keys := v.MapKeys()
		for i, l := 0, v.Len(); i < l; i++ {
			if i > 0 {
				writeSpace(buf, indent)
			}
			buf.WriteByte('(')
			key := keys[i]
			if err := encode_indent(buf, key, indent); err != nil {
				return err
			}
			buf.WriteByte(' ')
			if err := encode_indent(buf, v.MapIndex(key), indent); err != nil {
				return err
			}
			buf.WriteByte(')')
			if i < l-1 {
				// buf.WriteByte(' ')
				buf.WriteByte('\n')
			}
		}
		buf.WriteByte(')')
	default:
		fmt.Errorf("unsupported type:%s", v.Type())
	}

	return nil
}
