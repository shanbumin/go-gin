package main


import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"golang.org/x/net/context"
	"golang.org/x/sync/errgroup"
)

var (
	Version   = "N/A"
	BuildTime = "N/A"
)

//https://blog.csdn.net/tianlongtc/article/details/80038235

//go run main.go  -timeout 1000ms . fmt
//为了测试 sync.ErrGroup 的所有特性， 我已经写了一个递归搜索指定目录的Go程序。并且我还加了一个超时的机制 。当程序超时了，所有的 goroutines 将被取消，程序退出。

func main() {

	//第一步:命令行传进来的参数做了解析
	duration := flag.Duration("timeout", 500*time.Millisecond, "timeout in milliseconds")
	//一般用来提醒的，我们可以覆盖掉包中默认的，改成我们自定义的
	flag.Usage = func() {
		fmt.Printf("%s by Brian Ketelsen\n", os.Args[0])
		fmt.Printf("Version %s, Built: %s \n", Version, BuildTime)
		fmt.Println("Usage:")
		fmt.Printf("	gogrep [flags] path pattern \n")
		fmt.Println("Flags:")
		flag.PrintDefaults()
	}
	flag.Parse()
	if flag.NArg() != 2 {
		flag.Usage()
		os.Exit(-1)
	}
	path := flag.Arg(0)
	pattern := flag.Arg(1)

	//设置超时时间
	ctx, cf := context.WithTimeout(context.Background(), *duration)
	defer cf()



	m, err := search(ctx, path, pattern)
	if err != nil {
		log.Fatal(err)
	}
	for _, name := range m {
		fmt.Println(name)
	}
	fmt.Println(len(m), "hits")
}

func search(ctx context.Context, root string, pattern string) ([]string, error) {
	g, ctx := errgroup.WithContext(ctx)
	paths := make(chan string, 100)
	// get all the paths

	g.Go(func() error {
		defer close(paths)

		return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.Mode().IsRegular() {
				return nil
			}
			if !info.IsDir() && !strings.HasSuffix(info.Name(), ".go") {
				return nil
			}

			select {
			case paths <- path:
			case <-ctx.Done():
				return ctx.Err()
			}
			return nil
		})

	})

	c := make(chan string, 100)
	for path := range paths {
		p := path
		g.Go(func() error {
			data, err := ioutil.ReadFile(p)
			if err != nil {
				return err
			}
			if !bytes.Contains(data, []byte(pattern)) {
				return nil
			}
			select {
			case c <- p:
			case <-ctx.Done():
				return ctx.Err()
			}
			return nil
		})
	}
	go func() {
		g.Wait()
		close(c)
	}()

	var m []string
	for r := range c {
		m = append(m, r)
	}
	return m, g.Wait()
}