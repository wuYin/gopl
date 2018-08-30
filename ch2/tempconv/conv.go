// 包 tempconv 负责摄氏温度与华氏温度的转换计算
package tempconv

import "fmt"

type Celsius float64    // 摄氏温度
type Fahrenheit float64 // 华氏温度
type Kelvin float64     // 开尔文温度

const (
	AbsoluteZeroC Celsius = 273.15
	FrezzingC     Celsius = 0
	BoilingC      Celsius = 100
)

// 自定义输出函数
func (c Celsius) String() string {
	return fmt.Sprintf("%.3g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%.3g°F", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%.3g°K", k)
}
