package temperature

// CtoF converts temperature from Celsius to Fahrenheit
func CtoF(c float64) float64 {
	return (c * (9 / 5)) + 32
}

// FtoC converts temperature from Fahrenheit to Celsius
func FtoC(f float64) float64 {
	return (f - 32) * 5 / 9
}
