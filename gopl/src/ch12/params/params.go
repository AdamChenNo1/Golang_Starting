/*
 * File: \src\ch12\params\params.go                                            *
 * Project: gopl                                                               *
 * Created At: Thursday, 2022/06/2 , 20:47:22                                  *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Thursday, 2022/06/2 , 21:05:39                               *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package params

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

// Unpack 从HTTP请求req的参数中提取数据填充到ptr所指向的结构体的各个字段
func Unpack(req *http.Request, ptr any) error {
	if err := req.ParseForm(); err != nil {
		return err
	}

	// 创建字段映射表，键为有效字段
	fields := make(map[string]reflect.Value)
	v := reflect.ValueOf(ptr).Elem()

	for i, l := 0, v.NumField(); i < l; i++ {
		fieldInfo := v.Type().Field(i) // a reflect.StructField
		tag := fieldInfo.Tag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		fields[name] = v.Field(i)
	}

	// 对请求中的每个参数更新结构体中对应的字段
	for name, values := range req.Form {
		f := fields[name]
		if !f.IsValid() {
			continue // 忽略不能识别的HTTP参数
		}

		for _, value := range values {
			if f.Kind() == reflect.Slice { // 同一个参数可能会出现多次，
				elem := reflect.New(f.Type().Elem()).Elem()
				if err := populate(elem, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
				f.Set(reflect.Append(f, elem))
			} else {
				if err := populate(f, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
			}
		}
	}
	return nil
}

// populate 从单个HTTP请求参数值填充单个字段v，或者slice字段中的单个元素
func populate(v reflect.Value, value string) error {
	switch v.Kind() {
	case reflect.String:
		v.SetString(value)
	case reflect.Int:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		v.SetInt(i)
	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		v.SetBool(b)
	default:
		return fmt.Errorf("unspported kind %s", v.Type())
	}

	return nil
}
