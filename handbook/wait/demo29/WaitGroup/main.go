package main

import (
	"fmt"
	"net/http"
	"sync"
)

//Go 的一个很重要的的特性就是其原生的并发，像 channel 和 goroutines 这样的利器。但是对于一个新手来说 goroutines 这个概念可能比较陌生。
//Go 团队发布的第一个 goroutines 的管理工具是 sync.WaitGroup，这个工具允许你创建 WaitGroup 去等待一定数量的 goroutines 执行完成。这里有个例子：

func main() {
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.baidu.com/",
		"http://www.zuzuche.com/",
		"http://www.mokasz.com/",
	}
	for _, url := range urls {
		// Increment the WaitGroup counter.
		wg.Add(1)
		// Launch a goroutine to fetch the URL.
		go func(url string) {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()
			// Fetch the URL.
			resp,_:=http.Get(url)
			fmt.Println(resp.Body)
		}(url)
	}
	// Wait for all HTTP fetches to complete.
	wg.Wait()
}


//WaitGroup 使你在处理并发任务时对 goroutines 的创建和停止的数量控制都变的更加简单。
//每次你创建 goroutine 的时候只要调用 Add() 就可以了。当这个任务结束调用 wg.Done()。
//等待所有的任务完成，调用 wg.Wait()。但是用 WatiGroup 唯一的问题就是当你的 goroutines 出错时，你不能捕获到错误的原因。

