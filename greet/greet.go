package main

import (
	"fmt"
	"io"
	"net/http"
)

// 依赖注入的原理
// https://studygolang.gitbook.io/learn-go-with-tests/go-ji-chu/dependency-injection
// 想要重用，注入自己的 io.Writer 即可
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

// go run greet.go
// localhost:5000
func main() {
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}