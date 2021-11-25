package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Celsius float64
type Fahrenheit float64
type Kelvin float64

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		input := bufio.NewScanner(os.Stdin)
		input.Split(bufio.ScanWords)
		for input.Scan() {
			t, err := strconv.ParseFloat(input.Text(), 64)
			if err != nil {
				fmt.Printf("Error converting input to a float\n")
				os.Exit(1)
			}
			var c Celsius = Celsius(t)
			fmt.Println("Kelvin: ", CToK(c).String())
			fmt.Println("Fahrenheit: ", CToF(c).String())
		}
	} else {
		for _, arg := range args {
			t, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Printf("Error converting input to a float\n")
				os.Exit(1)
			}
			var c Celsius = Celsius(t)
			fmt.Println("Kelvin: ", CToK(c).String())
			fmt.Println("Fahrenheit: ", CToF(c).String())
		}
	}
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
