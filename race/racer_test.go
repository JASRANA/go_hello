package race

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

const slowDelayTime = 20
const fastDelayTime = 0

func TestRacer(t *testing.T) {
	slowServer := makeDelayedServer(slowDelayTime)
	fastServer := makeDelayedServer(fastDelayTime)

	// 包裹 defer 的函数结束时调用后边的语句
	// https://studygolang.gitbook.io/learn-go-with-tests/go-ji-chu/select#defer
	defer slowServer.Close()
	defer fastServer.Close()

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	got := Racer(slowURL, fastURL)
	want := fastURL

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}

}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		time.Sleep(delay * time.Millisecond)
		writer.WriteHeader(http.StatusOK)
	}))
}
