/*
 * @Descripttion:
 * @version: v0.1
 * @Author: Elon C
 * @Date: 2021-05-28 21:11:56
 * @LastEditors: Elon C
 * @LastEditTime: 2021-05-28 21:22:18
 * @FilePath: \Golang_Starting\ftoc\main.go
 */
// ftoc 输出两个华氏-摄氏温度的转换
package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF)) //"32°F = 0°C"
	fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF))   //"32°F = 0°C"
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
