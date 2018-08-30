package main

import (
	"gopl/ch2/tempconv"
	"fmt"
)

func main() {
	var c tempconv.Celsius = 233
	var f tempconv.Fahrenheit = 233
	var k tempconv.Kelvin = 233

	fmt.Printf("%s = %s\n", c, tempconv.CToF(c))
	fmt.Printf("%s = %s\n", f, tempconv.FToC(f))
	fmt.Printf("%s = %s\n", k, tempconv.KToC(k))
}
