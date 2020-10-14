package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const countdownStart = 3
const finalWord = "Go!"

// 依赖注入，运行时打印到命令行，测试时输出到 buffer
func Countdown(writer io.Writer)  {
	for i := countdownStart; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Fprintln(writer, i)
	}
	time.Sleep(1 * time.Second)
	fmt.Fprint(writer, finalWord)
}

func main() {
	Countdown(os.Stdout)
}