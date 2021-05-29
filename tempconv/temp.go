/*
 * @Descripttion:
 * @version: v0.1
 * @Author: Elon C
 * @Date: 2021-05-28 21:58:40
 * @LastEditors: Elon C
 * @LastEditTime: 2021-05-29 14:54:48
 * @FilePath: \Golang_Starting\tempconv\temp.go
 */
//包tempconv进行摄氏温度和华氏温度的转换
package tempconv

import "fmt"

type Celsius float64    //摄氏温标
type Fahrenheit float64 //华氏温标
type Kelvin float64     //热力学温标

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100

	AbsoluteZeroF Fahrenheit = -459.67
	FreezingF     Fahrenheit = 32
	BoilingF      Fahrenheit = 212

	AbsoluteZeroK Kelvin = 0
	FreezingK     Kelvin = 273.15
	BoilingK      Kelvin = 373.15
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%gK", k)
}
