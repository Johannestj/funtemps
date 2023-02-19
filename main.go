package main

import (
	"flag"
	"fmt"
	"github.com/Johannestj/funtemps/conv"
	"strings"
)

// Definerer flag-variablene i hoved-"scope"
var fahr float64
var out string
var funfacts string
var kel float64
var cel float64

// Bruker init (som anbefalt i dokumentasjonen) for å sikre at flagvariablene
// er initialisert.
func init() {

	/*
	   Her er eksempler på hvordan man implementerer parsing av flagg.
	   For eksempel, kommando
	       funtemps -F 0 -out C
	   skal returnere output: 0°F er -17.78°C
	*/

	// Definerer og initialiserer flagg-variablene
	flag.Float64Var(&fahr, "F", 0.0, "temperatur i grader fahrenheit")
	// Du må selv definere flag-variablene for "C" og "K"
	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - fahrenheit, K- Kelvin")
	flag.StringVar(&funfacts, "funfacts", "sun", "\"fun-facts\" om sun - Solen, luna - Månen og terra - Jorden")
	flag.Float64Var(&kel, "K", 0.0, "temperatur i grader kelvin")
	flag.Float64Var(&cel, "C", 0.0, "temperatur i grader celcius")
	// Du må selv definere flag-variabelen for -t flagget, som bestemmer
	// hvilken temperaturskala skal brukes når funfacts skal vises

}

func formatNumber(number float64) string {
    numStr := fmt.Sprintf("%.2f", number)
    numParts := strings.Split(numStr, ".")
    intPart := numParts[0]
    intPartLen := len(intPart)

    // Remove trailing zeros after decimal point
    decPart := strings.TrimRight(numParts[1], "0")
    if decPart == "" {
        return intPart
    }

    // Format integer part with thousands separators
    if intPartLen <= 3 {
        return intPart + "." + decPart
    }

    start := intPartLen % 3
    if start == 0 {
        start = 3
    }

    var result string
    for i, digit := range intPart {
        if i == start {
            result += " "
            start += 3
        }
        result += string(digit)
    }

    return result + "." + decPart
}

// removes the decimal of zero for the input value
/*func RemoveDecimal(number float64) string {
    numStr := fmt.Sprintf("%.2f", number)
    numParts := strings.Split(numStr, ".")
    intPart := numParts[0]

    decPart := strings.TrimRight(numParts[1], "0")
    if decPart == "" {
        return intPart
    }
    return intPart + "." + decPart
}*/


func main() {

	flag.Parse()

	if out == "C" && isFlagPassed("F"){
		fahr := conv.FahrenheitToCelsius(fahr)
		fmt.Printf("%.2f°F is %s°C\n", fahr, formatNumber(cel))
	}

	if out == "F" && isFlagPassed("C") {
        fahr := conv.CelsiusToFahrenheit(cel)
        //fmt.Printf("%#v°C is %.2f°F\n", cel, fahr)
		fmt.Printf("%.2f°C is %s°F\n", cel, formatNumber(fahr))
    }

	if out == "K" && isFlagPassed("C") {
        kel := conv.CelsiusToKelvin(cel)
        //fmt.Printf("%#v°C is %.2f°K\n", cel, kel)
		fmt.Printf("%.2f°C is %s°K\n", cel, formatNumber(kel))
    }

	if out == "C" && isFlagPassed("K") {
        cel := conv.KelvinToCelsius(kel)
        //fmt.Printf("%#v°K is %.2f°C\n", kel, cel)
		fmt.Printf("%.2f°K is %s°C\n", kel, formatNumber(cel))
    }

	if out == "F" && isFlagPassed("K") {
        fahr := conv.KelvinToFahrenheit(kel)
        //fmt.Printf("%#v°K is %.2f°F\n", kel, fahr)
		fmt.Printf("%.2f°K is %s°F\n", kel, formatNumber(fahr))
    }

	if out == "K" && isFlagPassed("F") {
        kel := conv.FahrenheitToKelvin(fahr)
       //fmt.Printf("%#v°F is %.2f°K\n", fahr, kel)
	   fmt.Printf("%.2f°F is %s°K\n", fahr, formatNumber(kel))
    }
}

// Funksjonen sjekker om flagget er spesifisert på kommandolinje
// Du trenger ikke å bruke den, men den kan hjelpe med logikken
func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
