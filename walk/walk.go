package walk

import "reflect"

// 通过反射，获取结构体的属性
func walk(x interface{}, fn func(input string))  {
	val := reflect.ValueOf(x)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fn(field.String())
	}
}
