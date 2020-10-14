package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const countdownStart = 3
const finalWord = "Go!"

type Sleeper interface {
	Sleep()
}

// mock spies 监视器，监视执行的次数
type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type ConfigurableSleeper struct {
	duration time.Duration
}

func (o *ConfigurableSleeper) Sleep() {
	time.Sleep(o.duration)
}

// 记录操作次数与操作类型
type CountdownOperationSpy struct {
	Calls []string
}

func (s *CountdownOperationSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

// 实现了 io.writer
func (s *CountdownOperationSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const sleep = "sleep"
const write = "write"

// 依赖注入，运行时打印到命令行，测试时输出到 buffer
// 问题引入：每次需要花费 4s 的时间执行 Countdown 测试
// 问题引入：代码的执行顺序无法简单使用 SpySleeper 监视
// mock 的原则
// https://studygolang.gitbook.io/learn-go-with-tests/go-ji-chu/mocking
func Countdown(writer io.Writer, sleeper Sleeper)  {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(writer, i)
	}
	sleeper.Sleep()
	fmt.Fprint(writer, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second}
	Countdown(os.Stdout, sleeper)
}