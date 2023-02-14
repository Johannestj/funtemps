package main

import (
	"flag"
	"fmt"
	"github.com/Johannestj/funtemps/conv"
	//"github.com/Johannestj/funtemps/conv_test"
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

func main() {

	flag.Parse()

	if out == "C"&& isFlagPassed("F"){
		fahr := conv.CelsiusToFahrenheit(cel)
		fmt.Printf("%v°F is %.2f°C\n", fahr, cel)
	}

	if out == "F" && isFlagPassed("C") {
        fahr := conv.CelsiusToFahrenheit(cel)
        fmt.Printf("%v°C is %.2f°F\n", cel, fahr)
    }

	if out == "K" && isFlagPassed("C") {
        kel := conv.CelsiusToKelvin(cel)
        fmt.Printf("%v°C is %.2f°K\n", cel, kel)
    }

	if out == "C" && isFlagPassed("K") {
        cel := conv.KelvinToCelsius(kel)
        fmt.Printf("%v°K is %.2f°C\n", kel, cel)
    }

	if out == "F" && isFlagPassed("K") {
        fahr := conv.KelvinToFahrenheit(kel)
        fmt.Printf("%v°K is %.2f°F\n", kel, fahr)
    }

	if out == "K" && isFlagPassed("F") {
        kel := conv.FahrenheitToKelvin(fahr)
        fmt.Printf("%v°F is %.2f°K\n", fahr, kel)
    }


	/**
	    Her må logikken for flaggene og kall til funksjoner fra conv og funfacts
	    pakkene implementeres.

	    Det er anbefalt å sette opp en tabell med alle mulige kombinasjoner
	    av flagg. flag-pakken har funksjoner som man kan bruke for å teste
	    hvor mange flagg og argumenter er spesifisert på kommandolinje.

	        fmt.Println("len(flag.Args())", len(flag.Args()))
			    fmt.Println("flag.NFlag()", flag.NFlag())

	    Enkelte kombinasjoner skal ikke være gyldige og da må kontrollstrukturer
	    brukes for å utelukke ugyldige kombinasjoner:
	    -F, -C, -K kan ikke brukes samtidig
	    disse tre kan brukes med -out, men ikke med -funfacts
	    -funfacts kan brukes kun med -t
	    ...
	    Jobb deg gjennom alle tilfellene. Vær obs på at det er en del sjekk
	    implementert i flag-pakken og at den vil skrive ut "Usage" med
	    beskrivelsene av flagg-variablene, som angitt i parameter fire til
	    funksjonene Float64Var og StringVar
	*/

	// Her er noen eksempler du kan bruke i den manuelle testingen
	//fmt.Println(fahr, out, funfacts)

	//fmt.Println("len(flag.Args())", len(flag.Args()))
	//fmt.Println("flag.NFlag()", flag.NFlag())

	//fmt.Println(isFlagPassed("out"))

	// Eksempel på enkel logikk
	//if out == "C" && isFlagPassed("F") {
		// Kalle opp funksjonen FahrenheitToCelsius(fahr), som da
		// skal returnere °C
	//		fmt.Println("0°F er -17.78°C")
	//}

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
