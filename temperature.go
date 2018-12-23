package temperature

// Temperature represents measurement of temperature in any system - Celsius, Fahnrenheit, ...
type Temperature float64

// CtoF converts temperature from Celsius to Fahrenheit
func CtoF(c Temperature) Temperature {
	return (c * (9 / 5)) + 32
}

// FtoC converts temperature from Fahrenheit to Celsius
func FtoC(f Temperature) Temperature {
	return (f - 32) * 5 / 9
}
