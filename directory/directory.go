package directory

type Directory map[string]string

const (
	ErrNotFound = DictionaryErr("could not find the word you were looking for")
	ErrKeyExists = DictionaryErr("key exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

// 使错误更具有可读性
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

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

func (d Directory) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrKeyExists
	default:
		return err
	}

	return nil
}

func (d Directory) Update(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case nil:
		d[key] = value
	case ErrNotFound:
		return ErrWordDoesNotExist
	default:
		return err
	}

	return nil
}