// 修改 dup2 的程序，输出出现重复行的文件的文件名称
package main

import (
	"os"
	"bufio"
	"fmt"
)

func main() {
	counts := make(map[string]int)
	l2Files := make(map[string][]string) // duplicateLine => files
	files := os.Args[1:]
	// 从 stdin 读取
	if len(files) == 0 {
		countFiles(os.Stdin, counts, l2Files)
		return
	}
	// 从文件列表读取
	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "open file %s error: %v\n", file, err)
			continue
		}
		countFiles(f, counts, l2Files)
		f.Close()
	}
	// 输出有重复行的文件
	for line, n := range counts {
		if n >= 2 {
			fmt.Printf("%d\t%q\t%q\n", n, line, l2Files[line])
		}
	}
}

func countFiles(f *os.File, counts map[string]int, l2Files map[string][]string) {
	input := bufio.NewScanner(f)
	fName := f.Name()
	for input.Scan() {
		if input.Err() != nil {
			fmt.Fprintf(os.Stderr, "%v\n", input.Err())
			break
		}
		line := input.Text()
		if !isIn(fName, l2Files[line]) {
			l2Files[line] = append(l2Files[line], fName)
		}
		counts[line]++
	}
}

func isIn(needle string, haystack []string) (in bool) {
	for _, s := range haystack {
		if needle == s {
			in = true
		}
	}
	return
}
