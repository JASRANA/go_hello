package race

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	// 制作测试用 HTTP 服务器，通过控制请求返回时间达到测试目的
	slowServer := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		// 先休眠 20 毫秒
		time.Sleep(20 * time.Millisecond)
		writer.WriteHeader(http.StatusOK)
	}))

	fastServer := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
	}))

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	got := Racer(slowURL, fastURL)
	want := fastURL

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}

	slowServer.Close()
	fastServer.Close()
}
