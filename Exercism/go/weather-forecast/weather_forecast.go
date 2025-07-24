// Package weather provides functionality to manage weather forecasts.
package weather

// CurrentCondition holds the current weather condition.
var CurrentCondition string
// CurrentLocation holds the name of the current location.
var CurrentLocation string

// Forecast sets the current weather condition for a given city and returns a formatted string.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
