package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	// waitForServer("http://wuyin.me")

	log.Printf("test change line1")   // log 会给为换行的输出加上换行符，很贴心啊
	log.Printf("test change line2\n") // 效果一致
	log.Printf("test change line3")
}

// 错误处理的第二种方式
// 在合理的时间范围内重试
func waitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Get(url)
		if err == nil {
			return nil
		}

		fmt.Printf("server response: %v, retry %d...\n", err.Error(), tries) // 出错
		time.Sleep(time.Second << uint(tries))                               // 1 2 4 8... // 超时的指数规避策略
	}

	return fmt.Errorf("server not response: %s after %v", url, deadline)
}
