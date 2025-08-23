// Package weather is ....
package weather

// CurrentCondition is a string ...
var CurrentCondition string

// CurrentLocation is a string to store location.
var CurrentLocation string

// Forecast takes the city and condition and returns a string.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
