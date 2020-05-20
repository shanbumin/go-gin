package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
)

//sync.WaitGroup 的加强版
//最近 Go 的团队在试验的代码仓库中增加了一个包叫 sync.ErrGroup。sync.ErrGroup 相当于为 sync.WaitGroup 增加了错误返回的功能。下面我们来看一个同样的例子：



func main(){
	var g errgroup.Group
	var urls = []string{
		"http://www.baidu.com/",
		"http://www.zuzuche.com/",
		"http://www.mokasz.com/",
		"http://www.mokasz.com/",
		"http://www.mokasz.com/",
		"http://www.mokasz.com/",
		"http://www.mokasz.com/",
		"http://www.mokasz.com/",
	}
	for _, url := range urls {
		url := url
		g.Go(func() error {
			resp, err := http.Get(url)
			if err == nil {
				resp.Body.Close()
			}
			return err
		})
	}
	if err := g.Wait(); err == nil {
		fmt.Println("Successfully fetched all URLs.")
	}else{
		fmt.Println(err)
	}
}

//g.Go() 这个方法不仅允许你传一个匿名的函数，而且还能捕获错误信息，你只要像这样返回一个错误 return err。
//这对使用 goroutines 的开发者来说在功能上是一个很大的提升。
