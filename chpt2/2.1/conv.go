package main

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

func main() {

	const (
		AbsoluteZeroC Celsius = -273.15
		FreezingC     Celsius = 0
		BoilingC      Celsius = 100
	)

	//var c Celsius = 0
	//var f Fahrenheit = 0
	var k Kelvin = 0
	var f Fahrenheit = 0
	var c Celsius = 0
	fmt.Println(KToF(k).String())
	fmt.Println(FToK(f).String())
	fmt.Println(CToK(c).String())
	//tmp := CToF(AbsoluteZeroC)
	//fmt.Println(AbsoluteZeroC.String())
}

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// CToK converts a Celsius temperature to Kelvin.
func CToK(c Celsius) Kelvin { return Kelvin(c + 273.15) }

// FToK converts a Fahrenheit temperature to Kelvin.
func FToK(f Fahrenheit) Kelvin { return Kelvin(((f - 32) * 5 / 9) + 273.15) }

// KToC converts a Kelvin temperature to Celsius.
func KToC(k Kelvin) Celsius { return Celsius(k - 273.15) }

// KToF converts a Kelvin temperature to Fahrenheit.
func KToF(k Kelvin) Fahrenheit { return Fahrenheit((k-273.15)*9/5 + 32) }

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g°K", k) }
