// fetch 输出从 URL 获取的内容
package main

import (
	"os"
	"net/http"
	"fmt"
	"io/ioutil"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "get error: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close() // 手动关闭 body 数据，避免资源泄漏
		if err != nil {
			fmt.Fprintf(os.Stderr, "readall error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", b) // []byte 可使用 %s 直接输出
	}
}
