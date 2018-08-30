// 米与英寸转换程序
package main

import (
	"fmt"
	"os"
	"strconv"
	"bufio"
)

type Meter float64
type Inch float64

func (m Meter) String() string {
	return fmt.Sprintf("%.4g m", m)
}

func (i Inch) String() string {
	return fmt.Sprintf("%.4g inch", i)
}

func i2m(i Inch) Meter {
	return Meter(i * 0.0254)
}

func m2i(m Meter) Inch {
	return Inch(m * 39.37)
}

func main() {
	if len(os.Args) >= 2 {
		for _, v := range os.Args[:1] {
			printConv(v)
		}
		return
	}
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		printConv(scan.Text())
	}
}

func printConv(s string) {
	l, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}
	fmt.Printf("%f = %s\n", l, i2m(Inch(l)))
	fmt.Printf("%f = %s\n", l, m2i(Meter(l)))
}
