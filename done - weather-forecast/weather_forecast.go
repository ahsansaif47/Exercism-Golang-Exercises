// Package weather provides tools related to weather.
package weather

// CurrentCondition tells us about the current weather condition.
var CurrentCondition string

// CurrentLocations tells about the location we are at.
var CurrentLocation string

// Forecast tells about about the weather forecast.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
