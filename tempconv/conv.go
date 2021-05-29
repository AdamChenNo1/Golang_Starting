/*
 * @Descripttion:
 * @version: v0.1
 * @Author: Elon C
 * @Date: 2021-05-28 22:15:06
 * @LastEditors: Elon C
 * @LastEditTime: 2021-05-29 15:00:28
 * @FilePath: \Golang_Starting\tempconv\CONV.GO
 */
package tempconv

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func KToF(k Kelvin) Fahrenheit {
	return Fahrenheit((k-273.15)*9/5 + 32)
}

func KToC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

func FToK(f Fahrenheit) Kelvin {
	return Kelvin((f-32)*5/9 + 273.15)
}

func CToK(c Celsius) Kelvin {
	return Kelvin(c + 273.15)
}
