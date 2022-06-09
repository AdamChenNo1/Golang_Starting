/*
 * File: \src\ch12\sexpr\encode.go                                             *
 * Project: Golang_Starting                                                    *
 * Created At: Monday, 2022/05/23 , 20:38:35                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Wednesday, 2022/05/25 , 20:01:23                             *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package sexpr

import (
	"bytes"
	"fmt"
	"reflect"
)

func Marshal(v any) ([]byte, error) {
	var buf bytes.Buffer

	if _, err := encode(&buf, reflect.ValueOf(v)); err != nil {
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

func encode(buf *bytes.Buffer, v reflect.Value) (int64, error) {
	if v.IsZero() {
		return 0, nil
	}

	var size int64
	switch v.Kind() {
	case reflect.Invalid:

	case reflect.Bool:
		buf.WriteByte('t')
		size++

	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		nsize, _ := fmt.Fprintf(buf, "%d", v.Int())
		size += int64(nsize)

	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64:
		nsize, _ := fmt.Fprintf(buf, "%d", v.Uint())
		size += int64(nsize)

	case reflect.Float32, reflect.Float64:
		nsize, _ := fmt.Fprintf(buf, "%v", v.Float())
		size += int64(nsize)

	case reflect.Complex64, reflect.Complex128:
		c := v.Complex()
		nsize, _ := fmt.Fprintf(buf, "#(%v,%v)", real(c), imag(c))
		size += int64(nsize)

	case reflect.String:
		nsize, _ := fmt.Fprintf(buf, "%s", v.String())
		size += int64(nsize)

	case reflect.Ptr:
		return encode(buf, v.Elem())

	case reflect.Array, reflect.Slice:
		data := new(bytes.Buffer)

		data.WriteByte('(')

		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				data.WriteByte(' ')
			}
			nsize, err := encode(data, v.Index(i))
			if err != nil {
				return 0, err
			}

			if i > 0 && nsize == 0 {
				data.UnreadByte()
			}
		}
		data.WriteByte(')')
		if data.Len() > 2 {
			n, err := data.WriteTo(buf)
			if err != nil {
				return 0, err
			}
			size = n
		}
	case reflect.Interface:
		data := new(bytes.Buffer)
		v = v.Elem()
		fmt.Fprintf(data, "%q", v.Type().Name())
		encode(data, v)

	case reflect.Struct:
		data := new(bytes.Buffer)

		data.WriteByte('(')
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				data.WriteByte(' ')
			}
			nkey, err := fmt.Fprintf(data, "(%s ", v.Type().Field(i).Name)
			if err != nil {
				return 0, err
			}

			if nkey < 3 {
				data.UnreadByte()
				data.UnreadByte()

				if i > 0 {
					data.UnreadByte()
				}
				continue
			}

			nvalue, err := encode(data, v.Field(i))
			if err != nil {
				return 0, err
			}

			if nvalue == 0 {
				for nkey > 0 {
					data.UnreadByte()
					nkey--
				}
				if i > 0 {
					data.UnreadByte()
				}
				continue
			}
			data.WriteByte(')')
		}

		buf.WriteByte(')')
		if data.Len() > 2 {
			n, err := data.WriteTo(buf)
			if err != nil {
				return 0, err
			}
			size = n
		}

	case reflect.Map:
		data := new(bytes.Buffer)

		data.WriteByte('(')
		keys := v.MapKeys()
		for i, l := 0, v.Len(); i < l; i++ {
			if i > 0 {
				data.WriteByte(' ')
			}
			data.WriteByte('(')
			key := keys[i]
			nkey, err := encode(data, key)
			if err != nil {
				return 0, err
			}
			if nkey == 0 && i > 0 {
				data.UnreadByte()
				continue
			}

			buf.WriteByte(' ')
			nvalue, err := encode(buf, v.MapIndex(key))
			if err != nil {
				return 0, err
			}

			if nvalue == 0 {
				for nkey > -2 {
					data.UnreadByte()
					nkey--
				}

				if i > 0 {
					data.UnreadByte()
				}
				continue
			}

			buf.WriteByte(')')
		}
		buf.WriteByte(')')

		if data.Len() > 2 {
			n, err := data.WriteTo(buf)
			if err != nil {
				return 0, err
			}
			size = n
		}

	default:
		fmt.Errorf("unsupported type:%s", v.Type())
	}

	return size, nil
}

func writeSpace(buf *bytes.Buffer, size uint) {
	for size > 0 {
		buf.WriteByte(' ')
		size--
	}
}

func encode_indent(buf *bytes.Buffer, v reflect.Value, indent uint) (int64, error) {
	var size int64

	if v.IsZero() {
		return 0, nil
	}

	switch v.Kind() {
	case reflect.Bool:
		buf.WriteByte('t')
		size++

	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		n, _ := fmt.Fprintf(buf, "%d", v.Int())
		size += int64(n)

	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64:
		n, _ := fmt.Fprintf(buf, "%d", v.Uint())
		size += int64(n)

	case reflect.Float32, reflect.Float64:
		n, _ := fmt.Fprintf(buf, "%v", v.Float())
		size += int64(n)

	case reflect.Complex64, reflect.Complex128:
		c := v.Complex()
		n, _ := fmt.Fprintf(buf, "#(%v,%v)", real(c), imag(c))
		size += int64(n)

	case reflect.String:
		n, _ := fmt.Fprintf(buf, "%q", v.String())
		size += int64(n)

	case reflect.Ptr:
		return encode_indent(buf, v.Elem(), indent)

	case reflect.Array, reflect.Slice:
		data := new(bytes.Buffer)

		data.WriteByte('(')
		indent++

		for i, l := 0, v.Len(); i < l; i++ {
			if i > 0 {
				writeSpace(data, indent)
			}
			n, err := encode_indent(data, v.Index(i), indent)
			if err != nil {
				return 0, err
			}
			if n == 0 && i > 0 {
				var i uint
				for ; i < indent; i++ {
					data.UnreadByte()
				}
				continue
			}
			if i < l-1 {
				data.WriteByte('\n')
			}
		}
		if data.Bytes()[data.Len()-1] == '\n' {
			data.UnreadByte()
		}
		data.WriteByte(')')
		if data.Len() > 2 {
			n, err := data.WriteTo(buf)
			if err != nil {
				return 0, err
			}

			size = n
		}

	case reflect.Interface:
		v = v.Elem()
		n, err := fmt.Fprintf(buf, "%q", v.Type().Name())
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
