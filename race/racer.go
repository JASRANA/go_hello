package race

import (
	"fmt"
	"net/http"
	"time"
)

// 问题引入：无法同时测试，浪费时间
func Racer(urlA, urlB string) (winner string, err error) {
	// 谁先返回，谁就胜出
	// select 允许同时在多个 channel 等待
	select {
	case <-ping(urlA):
		return urlA, nil
	case <-ping(urlB):
		return urlB, nil
		// time.After() 会在定义的时间过后发送一个 chan
	case <-time.After(10 * time.Second):
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