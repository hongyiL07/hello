package main

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func main() {
	fmt.Println("摄氏度和华氏度的转换")
}

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit((c*9/5 + 32))
}
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

//类型转换不改变类型的值，仅改变类型
