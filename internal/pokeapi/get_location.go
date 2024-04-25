package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// this gets the data for one specific location
func (c *Client) GetLocation(location string) (Location, error) {
	endpoint := "/location-area/" + location
	fullURL := baseURL + endpoint

	// check the cache here
	data, ok := c.cache.Get(fullURL)

	if ok {
		// unmarshalling the json
		locationResp := Location{}
		err := json.Unmarshal(data, &locationResp)
		if err != nil {
			fmt.Println(err)
		}

		return locationResp, nil
	}
	// cache miss
	fmt.Println("cache miss")

	//making the request to the API
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return Location{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}

	// closes the resp object before the function returns
	defer resp.Body.Close()

	// checking status code of the response
	if resp.StatusCode > 399 {
		return Location{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	// reading the response body
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, err
	}

	// unmarshalling the json
	locationResp := Location{}
	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return Location{}, nil
	}

	c.cache.Add(fullURL, data)

	return locationResp, nil
}
