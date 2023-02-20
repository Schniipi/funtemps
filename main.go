package main

import (
	"flag"
	"fmt"
	"math"

	"github.com/Schniipi/funtemps/conv"
)

var fromUnit, toUnit string
var out string

var fahr float64
var cels float64
var kelv float64
var result float64
var value float64

func init() {

	flag.Float64Var(&fahr, "F", 0.0, "temperatur i grader fahrenheit")
	flag.Float64Var(&cels, "C", 0.0, "temperatur i grader celsius")
	flag.Float64Var(&kelv, "K", 0.0, "temperatur i grader kelvin")

	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")

}

func main() {

	flag.Parse()

	// Her skriver man inn input
	fmt.Print("Enter a temperature value: ")
	fmt.Scan(&value)

	fmt.Print("Enter the temperature unit (C, F or K): ")
	fmt.Scan(&fromUnit)

	fmt.Print("Enter the desired temperature unit (C, F or K): ")
	fmt.Scan(&toUnit)

	switch fromUnit {
	case "C":
		if toUnit == "F" {
			result = conv.CelsiusToFahrenheit(value)
		} else if toUnit == "K" {
			result = conv.CelsiusToKelvin(value)
		} else {
			result = value
		}

	case "F":
		if toUnit == "C" {
			result = conv.FahrenheitToCelsius(value)
		} else if toUnit == "K" {
			result = conv.FahrenheitToKelvin(value)
		} else {
			result = value
		}

	case "K":
		if toUnit == "C" {
			result = conv.KelvinToCelsius(value)
		} else if toUnit == "F" {
			result = conv.KelvinToFahrenheit(value)
		} else {
			result = value
		}
	}

	// Fjerner desimal hvis tallet er heltall
	if math.Mod(result, 1) == 0 {
		fmt.Printf("%.0f %s\n", result, toUnit)
	} else {
		fmt.Printf("%.2f %s\n", result, toUnit)
	}
}
