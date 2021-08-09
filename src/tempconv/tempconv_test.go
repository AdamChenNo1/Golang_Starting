/*
 * @Descripttion:
 * @version: v0.1
 * @Author: Elon C
 * @Date: 2021-05-28 22:17:00
 * @LastEditors: Elon C
 * @LastEditTime: 2021-05-29 15:02:31
 * @FilePath: \Golang_Starting\tempconv\tempconv_test.go
 */
package tempconv

import (
	"fmt"
	"testing"
)

func ExampleCelsius() {
	fmt.Printf("Brrrr! %v\n", AbsoluteZeroC)
	//输出
	//"Brrrr! -273.15°C"
}
func ExampleFahrenheit() {
	var f Fahrenheit = -459.67
	fmt.Printf("Brrrr! %v\n", f)
	//输出
	//"Brrrr! -273.15°C"
}

func ExampleKelvin() {
	fmt.Printf("Brrrr! %v\n", AbsoluteZeroK)
	//输出
	//"Brrrr! 0K"
}

func TestCToF(t *testing.T) {
	tests := []struct {
		c    Celsius
		want string
	}{
		{AbsoluteZeroC, "-459.67°F"},
		{FreezingC, "32.00°F"},
		{BoilingC, "212.00°F"},
	}
	for _, test := range tests {
		f := float64(CToF(test.c))
		s := fmt.Sprintf("%.2f°F", f)

		if s != test.want {
			t.Errorf("CToF(%g) = %s,want %s", test.c, s, test.want)
		}
	}
}

func TestFToC(t *testing.T) {
	tests := []struct {
		f    Fahrenheit
		want string
	}{
		{-459.67, "-273.15°C"},
		{32.00, "0.00°C"},
		{212.00, "100.00°C"},
	}
	for _, test := range tests {
		c := float64(FToC(test.f))
		s := fmt.Sprintf("%.2f°C", c)

		if s != test.want {
			t.Errorf("CToF(%g) = %s,want %s", test.f, s, test.want)
		}
	}
}

func TestKToF(t *testing.T) {
	tests := []struct {
		k    Kelvin
		want string
	}{
		{AbsoluteZeroK, "-459.67°F"},
		{FreezingK, "32.00°F"},
		{BoilingK, "212.00°F"},
	}
	for _, test := range tests {
		f := float64(KToF(test.k))
		s := fmt.Sprintf("%.2f°F", f)

		if s != test.want {
			t.Errorf("CToF(%g) = %s,want %s", test.k, s, test.want)
		}
	}
}

func TestKToC(t *testing.T) {
	tests := []struct {
		k    Kelvin
		want string
	}{
		{AbsoluteZeroK, "-273.15°C"},
		{FreezingK, "0.00°C"},
		{BoilingK, "100.00°C"},
	}
	for _, test := range tests {
		c := float64(KToC(test.k))
		s := fmt.Sprintf("%.2f°C", c)

		if s != test.want {
			t.Errorf("CToF(%g) = %s,want %s", test.k, s, test.want)
		}
	}
}

func TestFToK(t *testing.T) {
	tests := []struct {
		f    Fahrenheit
		want string
	}{
		{AbsoluteZeroF, "0.00K"},
		{FreezingF, "273.15K"},
		{BoilingF, "373.15K"},
	}
	for _, test := range tests {
		k := float64(FToK(test.f))
		s := fmt.Sprintf("%.2fK", k)

		if s != test.want {
			t.Errorf("CToF(%g) = %s,want %s", test.f, s, test.want)
		}
	}
}

func TestCToK(t *testing.T) {
	tests := []struct {
		c    Celsius
		want string
	}{
		{AbsoluteZeroC, "0.00K"},
		{FreezingC, "273.15K"},
		{BoilingC, "373.15K"},
	}
	for _, test := range tests {
		k := float64(CToK(test.c))
		s := fmt.Sprintf("%.2fK", k)

		if s != test.want {
			t.Errorf("CToF(%g) = %s,want %s", test.c, s, test.want)
		}
	}
}
