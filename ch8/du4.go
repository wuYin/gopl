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

var v = flag.Bool("v", false, "-v: show progress")

var done = make(chan struct{})

func canceled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

// du4 相比 du3 引入了 done channel 来中断正在执行的 goroutine
// 一般运行时的协程是不固定的，而且彼此也不能中断彼此的执行
// 取消执行可以添加广播机制：当广播 channel 中有数据时，不再创建新的 goroutine 并且保证现有的 goroutine 能执行完毕
func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done) // 广播机制 // 在 done channel 上进行 2 路 select 来判断
	}()

	sizeCh := make(chan int64)
	var wg sync.WaitGroup
	for _, root := range roots {
		wg.Add(1)
		go walkDir(root, &wg, sizeCh)
	}

	go func() {
		wg.Wait()
		close(sizeCh)
	}()

	ticker := make(<-chan time.Time)
	if *v {
		ticker = time.Tick(100 * time.Millisecond)
	}
	var count, total int64

label:
	for {
		select {
		case <-done:
			for range sizeCh {
				// do nothing // 将通道内的值全部耗尽
			}
		case size, ok := <-sizeCh:
			if !ok {
				break label
			}
			count++
			total += size
		case <-ticker:
			printUshage(count, total)
		}
	}

	printUshage(count, total)
}

func printUshage(files, size int64) {
	fmt.Printf("%d files: %.2fGB\n", files, float64(size)/1e9)
}

func walkDir(dir string, wg *sync.WaitGroup, sizeCh chan<- int64) {
	defer wg.Done()
	if canceled() {
		return
	}
	for _, f := range listDir(dir) {
		if f.IsDir() {
			wg.Add(1)
			subdir := path.Join(dir, f.Name())
			walkDir(subdir, wg, sizeCh)
		} else {
			sizeCh <- f.Size()
		}
	}
}

var sema = make(chan struct{}, 20)

func listDir(dir string) []os.FileInfo {

	select {
	case <-done:
		return nil // 剩下的协程不必再执行
	case sema <- struct{}{}: // 获得信号量
	}
	defer func() {
		<-sema
	}()

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return files
}
