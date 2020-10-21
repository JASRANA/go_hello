package walk

import "reflect"

// 通过反射，获取结构体的属性
func walk(x interface{}, fn func(input string))  {
	val := reflect.ValueOf(x)
	field := val.Field(0)
	// 如果不是 string，会报错的
	fn(field.String())
}
