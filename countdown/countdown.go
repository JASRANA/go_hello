package main

import (
	"fmt"
	"io"
	"os"
)

// 依赖注入，运行时打印到命令行，测试时输出到 buffer
func Countdown(writer io.Writer)  {
	fmt.Fprint(writer, 3)
}

func main() {
	Countdown(os.Stdout)
}