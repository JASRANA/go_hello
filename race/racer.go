package race

import (
	"net/http"
	"time"
)

func Racer(urlA, urlB string) (winner string) {
	startA := time.Now()
	http.Get(urlA)
	durationUrlA := time.Since(startA)

	startB := time.Now()
	http.Get(urlB)
	durationUrlB := time.Since(startB)

	if durationUrlA < durationUrlB {
		return urlA
	}
	return urlB
}