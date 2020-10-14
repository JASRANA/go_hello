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

// 依赖注入，运行时打印到命令行，测试时输出到 buffer
// 问题引入：每次需要花费 4s 的时间执行 Countdown 测试
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