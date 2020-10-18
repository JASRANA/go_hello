package race

import (
	"net/http"
	"time"
)

func Racer(urlA, urlB string) (winner string) {
	durationUrlA := measureResponseTime(urlA)
	durationUrlB := measureResponseTime(urlB)

	if durationUrlA < durationUrlB {
		return urlA
	}
	return urlB
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}