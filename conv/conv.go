package conv

/*
  I denne pakken skal alle konverteringfunksjonene
  implementeres. Bruk engelsk.
    FahrenheitToCelsius
    CelsiusToFahrenheit
    KelvinToFahrenheit
    ...
*/

// Konverterer Fahrenheit til Celsius
func FahrenheitToCelsius(fahr float64) float64 {
	cel := (fahr - 32) * 5/9
	return cel
}

func CelsiusToFahrenheit(cel float64) float64 {
	fahr := cel * 9/5 + 32
	return fahr
}

func FahrenheitToKelvin(fahr float64) float64 {
	kel := (fahr - 32) * 5/9 + 273.15
	return kel
}

func KelvinToFahrenheit(kel float64) float64 {
	fahr := (kel - 273.15) * 9/5 + 32
	return fahr
}

func KelvinToCelsius(kel float64) float64 {
	cel := kel - 273.15
	return cel
}

func CelsiusToKelvin(cel float64) float64 {
	kel := cel + 273.15
	return kel
}

// De andre konverteringsfunksjonene implementere her
// ...
