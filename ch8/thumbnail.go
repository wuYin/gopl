package main

import (
	"fmt"
	"log"
	"os"
	"sync"
)

var imgs = []string{"twitter.jpg", "facebook.png", "google.png", "ford.svg"}

// 模拟 thumbnail 操作
func ImageFile(fileName string) (string, error) {
	// 进行文件处理
	// ...
	return fileName + " done", nil
}

// 一般顺序操作
func handle() {
	for _, img := range imgs {
		if _, err := ImageFile(img); err != nil {
			log.Println(err)
		}
	}
}

// 直接 go
func handle2() {
	for _, img := range imgs {
		go ImageFile(img) // 不正确的。启动了所有子 goroutine 但主 goroutine 没有等待它们执行结束就返回了
	}
}

// 通过无缓冲（同步）channel 来报告每个子 goroutine 执行结束
func handle3() {
	event := make(chan struct{})
	for _, img := range imgs {
		go func(f string) {
			ImageFile(f)        // 忽略错误处理
			event <- struct{}{} // 报告执行结束
		}(img) // 这里显式地将参数传递给内部匿名函数。否则迭代完毕后所有内部匿名函数都将共享一个外部变量：f，即最后一个 img
	}

	// 因为知道确切子 goroutine 的数量，完全可以使用一个同步 channel 来报告子 goroutine 的执行结束
	for range imgs {
		<-event
	}
}

// 错误处理
func handle4(files []string) error {
	errors := make(chan error)
	for _, f := range files {
		go func(f string) {
			_, err := ImageFile(f)
			// 统统发送。这里是有问题的，当任何一个 goroutine 发生错误时候直接 return
			// 那些还在执行的 goroutine 将阻塞在此处，数据量大时可能导致内存泄漏
			errors <- err
		}(f)
	}

	for range files {
		if err := <-errors; err != nil {
			return err // 错误的。可能会导致一些启动了的子 goroutine 未执行完毕
		}
	}
	return nil
}

// 将同步 channel 修改为缓冲 channel 来避免 goroutine 泄漏
func handle5(files []string) (thumbs []string, err error) {
	type result struct {
		thumb string
		err   error
	}

	ch := make(chan result, len(files))
	for _, f := range files {
		go func(f string) {
			var res result
			res.thumb, res.err = ImageFile(f)
			ch <- res
		}(f)
	}

	for range files {
		res := <-ch
		if res.err != nil {
			return nil, res.err // 就算有一个 goroutine 出错提前返回，其他的也能将数据发送到缓冲 channel 中，不至于导致资源泄漏
		}
		thumbs = append(thumbs, res.thumb)
	}

	return thumbs, nil
}

// 更为优雅的 goroutine 计数器
// 通用的并发循环处理模式
// worker & sync.WaitGroup & closer() && for range
func handle6(files []string) (int64, error) {
	size := make(chan int64)
	var wg sync.WaitGroup

	for _, f := range files {
		wg.Add(1)

		// worker
		go func(f string) {
			defer wg.Done()
			thumb, err := ImageFile(f)
			if err != nil {
				fmt.Println(err)
				return // 自动结束执行
			}
			info, _ := os.Stat(thumb)
			size <- info.Size()
		}(f)
	}

	// closer
	// 一直等待所有 worker 执行完毕，再关闭 size channel
	// 主 goroutine 才能执行结束
	go func() {
		wg.Wait()
		close(size)
	}()

	var total int64
	for n := range size {
		total += n
	}

	return total, nil
}
