// cf 把它的参数值参数转换为摄氏温度和华氏温度
package main

import (
	"os"
	"strconv"
	"fmt"
	"gopl/ch2/tempconv" // 导入的包
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "conv error: %v\n", err)
			return
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
}
