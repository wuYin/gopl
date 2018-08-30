package tempconv

// CToF 将摄氏温度转为华氏温度
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC 将华氏温度转为摄氏温度
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// KToC 将开尔文温度转为摄氏温度
func KToC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}
