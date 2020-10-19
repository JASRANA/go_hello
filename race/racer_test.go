package race

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	//t.Run("测试谁先胜出", func(t *testing.T) {
	//	slowServer := makeDelayedServer(10 * time.Second)
	//	fastServer := makeDelayedServer(0)
	//
	//	// 包裹 defer 的函数结束时调用后边的语句
	//	// https://studygolang.gitbook.io/learn-go-with-tests/go-ji-chu/select#defer
	//	defer slowServer.Close()
	//	defer fastServer.Close()
	//
	//	got, _ := Racer(slowURL, fastURL)
	//	want := fastURL
	//
	//	slowServer.Close()
	//	fastServer.Close()
	//
	//	if got != want {
	//		t.Errorf("got '%s' want '%s'", got, want)
	//	}
	//})

	t.Run("超过 10s，报错", func(t *testing.T) {
		slowServer := makeDelayedServer(11 * time.Second)
		fastServer := makeDelayedServer(12 * time.Second)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		_, err := Racer(slowURL, fastURL)

		if err == nil {
			t.Fatal("expected an error but got nil.")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		time.Sleep(delay)
		writer.WriteHeader(http.StatusOK)
	}))
}
