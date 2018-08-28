// dup1 输出标准输入中出现次数大于 1 的行，前面是次数
package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		// # 作结束符
		if strings.TrimSpace(input.Text()) == "#" || input.Err() != nil {
			break
		}
		counts[input.Text()]++
	}
	for text, n := range counts {
		if n >= 2 {
			fmt.Printf("%q\t%d\n", text, n)
		}
	}
}
