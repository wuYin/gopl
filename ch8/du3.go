package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"sync"
	"time"
)

var v = flag.Bool("v", false, "-v: show more detail")

// 在 du2 的基础上套上循环并发模式，并发地统计占用量
// gr du2.go -v /Users/wuyin 	// 18s
// gr du3.go -v /Users/wuyin 	// 5s
func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	now := time.Now()

	fileSizeCh := make(chan int64)
	var wg sync.WaitGroup
	for _, root := range roots {
		wg.Add(1)
		go walkDir(root, &wg, fileSizeCh) // 光这里开协程不够啊
	}

	go func() {
		wg.Wait()
		close(fileSizeCh)
	}()

	var fileCount, totalSize int64
	ticker := make(<-chan time.Time) // ticker 本质上还是一个时间值接收 channel
	if *v {
		ticker = time.Tick(200 * time.Millisecond)
	}

loop:
	for {
		select {
		case size, ok := <-fileSizeCh: // 在 channel 上不进行 range 后必须判断当前 channel 是否已关闭。否则可能 panic
			if !ok {
				break loop // 需要使用 label 跳出 select 和 for
			}
			fileCount++
			totalSize += size
		case <-ticker:
			printUshage(fileCount, totalSize)
		}
	}

	printUshage(fileCount, totalSize) // ok
	fmt.Printf("cost: %.2fs\n", time.Now().Sub(now).Seconds())
}

func printUshage(files, size int64) {
	fmt.Printf("%d files: %.2fGB\n", files, float64(size)/1e9)
}

func walkDir(dir string, wg *sync.WaitGroup, fileSize chan<- int64) {
	defer wg.Done()
	for _, f := range read(dir) {
		if f.IsDir() {
			wg.Add(1)
			subDir := path.Join(dir, f.Name())
			go walkDir(subDir, wg, fileSize) // 去遍历处理子目录的时候也是要开协程去处理滴。幸好 goroutine 的递归调用栈很大
		} else {
			fileSize <- f.Size()
		}
	}
}

// 信号量控制 goroutine 数量
var sema = make(chan struct{}, 20)

func read(dir string) []os.FileInfo {
	sema <- struct{}{} // 信号量+1
	defer func() {
		<-sema // 信号量-1
	}()

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return files
}
