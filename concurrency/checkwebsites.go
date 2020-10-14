package concurrency

import "time"

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		// 通过将 go 放在 func() 前，开启 goroutine
		// 匿名函数
		// 如果不将 url 分别传入，那么都会使用同一个 url 的引用
		// 并发写入 Map 将报错！使用 go test -race 发现 *竞争条件*
		go func(u string) {
			results[u] = wc(u)
		}(url)
	}

	time.Sleep(2 * time.Second)

	return results
}