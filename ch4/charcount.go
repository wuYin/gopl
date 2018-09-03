// charcount 计算 Unicode 字符的个数
package main

import (
	"unicode/utf8"
	"bufio"
	"os"
	"fmt"
	"io"
	"unicode"
)

func main() {
	counts := make(map[rune]int)
	var utfLen [utf8.UTFMax + 1]int
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "read rune error: %v\n", err)
			os.Exit(1)
		}
		// 不合法字符或 ASCII 字符跳过
		if r == unicode.ReplacementChar || n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utfLen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for r, c := range counts {
		fmt.Printf("%q\t%d\n", r, c)
	}
	fmt.Printf("\nlen\tcount\n")
	for l, c := range utfLen {
		fmt.Printf("%d\t%d\n", l, c)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
