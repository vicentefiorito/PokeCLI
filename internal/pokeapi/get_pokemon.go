package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// gets the data from an specific pokemon
func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	// calls the api endpoint
	endpoint := "/pokemon/" + pokemonName
	fullUrl := baseURL + endpoint

	// check the cache here
	data, ok := c.cache.Get(fullUrl)

	if ok {
		pokemonResp := Pokemon{}
		// unmarshall the json here and check for erorr
		err := json.Unmarshal(data, &pokemonResp)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemonResp, nil
	}

	// making the request to the API'
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return Pokemon{}, err
	}

	// execute the request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, nil
	}

	// ensures we close the response object
	// as to not lose resources
	defer resp.Body.Close()

	// reading the response body
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, nil
	}
	// unmarshalling the json response
	pokemonResp := Pokemon{}
	err = json.Unmarshal(data, &pokemonResp)
	if err != nil {
		return Pokemon{}, nil
	}

	// adding it to the cache
	// for quick lookup
	c.cache.Add(fullUrl, data)
	return pokemonResp, nil
}
