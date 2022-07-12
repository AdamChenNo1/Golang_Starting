/*
 * File: /get_fields_values_of_struct.go/get_fields.go                         *
 * Project: go-demo                                                            *
 * Created At: Tuesday, 2022/06/28 , 06:16:06                                  *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Tuesday, 2022/06/28 , 12:05:08                               *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package main

import (
	"log"
	"reflect"
)

func get_fields(s any) []any {
	var res []any

	// value := reflect.ValueOf(&s).Elem()

	// log.Printf("%T", value)
	var toSlice func(v any)
	toSlice = func(v any) {
		rv := reflect.Indirect(reflect.ValueOf(&v)).Elem()
		if k := rv.Kind(); k != reflect.Struct {
			res = append(res, v)
		} else {
			for i, n := 0, rv.NumField(); i < n; i++ {
				field := rv.Field(i)
				// if field := rv.Field(i); field.Kind() == reflect.Struct {
				// toSlice(field.Interface())
				toSlice(field.Interface())
				// }
			}
		}
		// switch k {
		// case reflect.Int:
		// 	log.Printf("%T", v)

		// 	fields = append(fields, int(v.Int()))
		// case reflect.Int8:
		// 	fields = append(fields, int8(v.Int()))
		// case reflect.Int16:
		// 	fields = append(fields, int16(v.Int()))
		// case reflect.Int32:
		// 	fields = append(fields, int32(v.Int()))
		// case reflect.Int64:
		// 	fields = append(fields, int64(v.Int()))
		// case reflect.Bool:
		// 	fields = append(fields, v.Bool())
		// case reflect.Float32:
		// 	fields = append(fields, float32(v.Float()))
		// case reflect.Float64:
		// 	fields = append(fields, float64(v.Float()))
		// case reflect.Array:
		// 	log.Println(v.Type().Elem())
		// 	var arr = toSlice(v.Index(0))
		// 	for i := 1; i < v.Len(); i++ {
		// 		arr = append(arr, toSlice(v.Index(i))...)
		// 	}
		// 	fields = append(fields, arr)
		// case reflect.String:
		// 	fields = append(fields, v.String())
		// }

		// for i := 0; i < value.NumField(); i++ {
		// 	log.Println(value.Interface())
		// 	fi := value.Field(i).Elem()
		// 	fields = append(fields, toSlice(fi)...)
		// }
		// return fields
	}

	log.Printf("%T", &s)
	toSlice(s)
	return res
}

func get_fields2(s any) []any {
	var res []any

	var toSlice func(v reflect.Value)
	toSlice = func(v reflect.Value) {
		if k := v.Kind(); k != reflect.Struct {
			res = append(res, v.Interface())
		} else {
			for i, n := 0, v.NumField(); i < n; i++ {
				field := v.Field(i)
				toSlice(field)
			}
		}
	}

	log.Printf("%T", &s)
	toSlice(reflect.ValueOf(s))
	return res
}
