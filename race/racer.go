package race

import (
	"fmt"
	"net/http"
	"time"
)

const timeout = 10 * time.Second

// 真实调用
func Racer(urlA, urlB string) (winner string, err error) {
	return ConfigurableRacer(urlA, urlB, timeout)
}

// 1. 对超时时间有明确的规定
// 2. 我们在测试时并不真正关心超时时间
// 测试调用
// 思考：这不是会引入与业务不相关的代码，造成更多的封装？
func ConfigurableRacer(urlA, urlB string, timeout time.Duration) (winner string, error error) {
	// 问题引入：无法同时测试，浪费时间
	// 谁先返回，谁就胜出
	// select 允许同时在多个 channel 等待
	select {
	case <-ping(urlA):
		return urlA, nil
	case <-ping(urlB):
		return urlB, nil
		// time.After() 会在定义的时间过后发送一个 chan
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", urlA, urlB)
	}
}

func ping(url string) chan bool{
	ch := make(chan bool)
	go func() {
		http.Get(url)
		ch <- true
	}()
	return ch
}