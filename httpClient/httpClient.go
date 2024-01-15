package httpClient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const apiKey = "<API_KEY_HERE>"

type RequestObject struct {
	Lat float64
	Lon float64
}

type CustomHTTPClient struct {
	Client *http.Client
	ApiKey string
	URL    url.URL
}

type ResponseData struct {
	Status string `json:"status"`
	Data   []struct {
		Coordinates []struct {
			Lat   float64 `json:"lat"`
			Lon   float64 `json:"lon"`
			Dates []struct {
				Date  string  `json:"dates"`
				Value float64 `json:"value"`
			}
		} `json:"coordinates"`
	} `json:"data"`
}

func NewCustomHttpClient(reqObj RequestObject) *CustomHTTPClient {
	return &CustomHTTPClient{
		Client: &http.Client{},
		ApiKey: apiKey,
		URL: url.URL{
			Scheme: "https",
			Host:   "api.meteomatics.com",
			Path:   fmt.Sprintf("%s/t_2m:F/%f,%f/json", time.Now().Format(time.RFC3339), reqObj.Lat, reqObj.Lon),
		},
	}
}

func (chc CustomHTTPClient) MakeWeatherRequest() (*ResponseData, error) {
	endpoint := chc.URL.ResolveReference(&chc.URL)

	req, err := http.NewRequest(http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", chc.ApiKey))

	res, err := chc.Client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	response := &ResponseData{}
	if err := json.NewDecoder(res.Body).Decode(response); err != nil {
		fmt.Printf("Error decoding weather data: %s", err)

		return nil, err
	}

	return response, nil
}
