package concurrency

import "fmt"

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChanel := make(chan result)

	for _, url := range urls {
		// 通过将 go 放在 func() 前，开启 goroutine
		// 匿名函数
		// 如果不将 url 分别传入，那么都会使用同一个 url 的引用
		// 并发写入 Map 将报错！使用 go test -race 发现 *竞争条件*
		// 每个进程使用同一个通道 channel，每次将结果放进通道中
		go func(u string) {
			resultChanel <- result{u, wc(u)}
		}(url)
	}

	fmt.Println("after")
	// 排队将结果输出到 results
	for i := 0; i < len(urls); i++ {
		result := <- resultChanel
		fmt.Println(result.string)
		results[result.string] = result.bool
	}

	// 不需要再进行休眠等待
	// time.Sleep(2 * time.Second)

	return results
}