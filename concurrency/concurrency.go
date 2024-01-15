package concurrency

import (
	"concurrency_weather/customHttpClient"
	"fmt"
	"sync"
	"time"
)

type City struct {
	Name      string
	Latitude  float64
	Longitude float64
}

type Weather struct {
	City City
	Temp float64
}

func (c City) FetchWeatherSingleThread() (*Weather, error) {
	reqObj := customHttpClient.RequestObject{Lat: c.Latitude, Lon: c.Longitude}
	client := customHttpClient.NewCustomHttpClient(reqObj)

	response, err := client.MakeWeatherRequest()
	if err != nil {
		return nil, err
	}

	weather := &Weather{
		City: c,
		Temp: response.Data[0].Coordinates[0].Dates[0].Value,
	}
	return weather, nil
}

func (c City) FetchWeatherConcurrently(ch chan<- Weather, wg *sync.WaitGroup) (*Weather, error) {
	defer wg.Done()
	reqObj := customHttpClient.RequestObject{Lat: c.Latitude, Lon: c.Longitude}
	client := customHttpClient.NewCustomHttpClient(reqObj)
	response, err := client.MakeWeatherRequest()
	if err != nil {
		return nil, err
	}

	weather := &Weather{
		City: c,
		Temp: response.Data[0].Coordinates[0].Dates[0].Value,
	}
	ch <- *weather
	return weather, nil

}

func RunConcurrentTest(cities []City) {
	// Concurrent code
	concurrentTimeStartNow := time.Now()
	ch := make(chan Weather)
	var wg sync.WaitGroup
	for _, city := range cities {
		wg.Add(1)
		go city.FetchWeatherConcurrently(ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		fmt.Printf("%s temperature is %f degrees farenheit.\n", result.City.Name, result.Temp)
	}

	fmt.Println("This concurrent operation took: ", time.Since(concurrentTimeStartNow))
}

func RunSingleThreadTest(cities []City) {
	// Single threaded code
	singleThreadTimeStartNow := time.Now()
	for _, city := range cities {
		if result, err := city.FetchWeatherSingleThread(); err == nil {
			fmt.Printf("%s temperature is %f degrees farenheit.\n", result.City.Name, result.Temp)
		}
	}
	fmt.Println("This single threaded operation took: ", time.Since(singleThreadTimeStartNow))
}
