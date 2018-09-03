package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	// Ctrl + D = EOF
	for scanner.Scan() {
		counts[scanner.Text()]++
	}
	if scanner.Err() != nil {
		fmt.Println(scanner.Err())
		return
	}
	for w, c := range counts {
		fmt.Printf("%q: %d\n", w, c)
	}
}
