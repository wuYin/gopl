// dup3 统计命令行指定文件中出现超过 2 次的行
package main

import (
	"os"
	"io/ioutil"
	"fmt"
	"strings"
)

func main() {
	counts := make(map[string]int)
	// 读取各文件的全部内容
	for _, fName := range os.Args[1:] {
		data, err := ioutil.ReadFile(fName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "readfile error: %v\n", err)
			return
		}
		// 统计重复行
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	// 输出重复行
	for line, count := range counts {
		if count >= 2 {
			fmt.Printf("%q\t%d\n", line, count)
		}
	}
}
