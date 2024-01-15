package main

import (
	"concurrency_weather/concurrency"
	"fmt"
)

var cities = []concurrency.City{
	{
		Name:      "Chicago",
		Latitude:  41.8781,
		Longitude: 87.6298,
	},
	{
		Name:      "New York",
		Latitude:  40.7128,
		Longitude: 74.0060,
	},
	{
		Name:      "Mexico City",
		Latitude:  19.4326,
		Longitude: 99.1332,
	},
	{
		Name:      "Rio De Janeiro",
		Latitude:  22.9068,
		Longitude: 43.1729,
	},
	{
		Name:      "Tel Aviv",
		Latitude:  32.0853,
		Longitude: 34.7818,
	},
	{
		Name:      "Tokyo",
		Latitude:  35.6764,
		Longitude: 139.6500,
	},
	{
		Name:      "Sydney",
		Latitude:  33.8688,
		Longitude: 151.2093,
	},
	{
		Name:      "London",
		Latitude:  51.5072,
		Longitude: 0.1276,
	},
	{
		Name:      "Nairobi",
		Latitude:  1.2921,
		Longitude: 36.8219,
	},
}

func main() {
	fmt.Println("Starting Tests...")
	concurrency.RunConcurrentTest(cities)
	concurrency.RunSingleThreadTest(cities)
	fmt.Println("Finished Tests...")
}
