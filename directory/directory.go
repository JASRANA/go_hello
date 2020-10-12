package directory

import "errors"

type Directory map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")

// 可以同时返回很多值
// map 是引用类型，不需要指针传递就可以修改数据
// 这意味着 maps 可以是 nil 值，因此定义时一定要初始化
// var m map[string]string 会出错，抛出 nil 指针异常
func (d Directory) Search(key string) (string,error) {
	// definition 有值，ok = true
	definition, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Directory) Add(key, value string) {
	d[key] = value
}

func Search(directory map[string]string, key string) string {
	return directory[key]
}
