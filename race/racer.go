package race

import "net/http"

// 问题引入：无法同时测试，浪费时间
func Racer(urlA, urlB string) (winner string) {
	// 谁先返回，谁就胜出
	// select 允许同时在多个 channel 等待
	select {
	case <-ping(urlA):
		return urlA
	case <-ping(urlB):
		return urlB
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