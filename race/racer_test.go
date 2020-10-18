package race

import "testing"

func TestRacer(t *testing.T) {
	slowURl := "http://www.github.com"
	fastURL := "http://www.baidu.com"

	got := Racer(slowURl, fastURL)
	want := fastURL

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
