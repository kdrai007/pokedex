package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (shallowApiResponse, error) {
	fullURL := baseURL + "/location-area"
	if pageURL != nil {
		fullURL = *pageURL
	}
	// Fetching if there's any cached location
	if respData, ok := c.cache.Get(fullURL); ok {
		locationResp := shallowApiResponse{}
		err := json.Unmarshal(respData, &locationResp)
		if err != nil {
			return shallowApiResponse{}, err
		}
		return locationResp, nil
	}
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return shallowApiResponse{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return shallowApiResponse{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return shallowApiResponse{}, err
	}

	var locationApiResp shallowApiResponse
	err = json.Unmarshal(data, &locationApiResp)
	if err != nil {
		return shallowApiResponse{}, err
	}

	// caching fetched location
	c.cache.Add(fullURL, data)
	return locationApiResp, nil
}

func (c *Client) ListPokemons(location *string) (LocationArea, error) {
	fullURL := baseURL + "/location-area/" + *location

	if respData, ok := c.cache.Get(fullURL); ok {
		locationResp := LocationArea{}
		err := json.Unmarshal(respData, &locationResp)
		if err != nil {
			return LocationArea{}, err
		}
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationArea{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, err
	}

	var pokemonResp LocationArea
	err = json.Unmarshal(data, &pokemonResp)
	if err != nil {
		return LocationArea{}, err
	}
	c.cache.Add(fullURL, data)
	return pokemonResp, nil
}

func (c *Client) FetchPokemon(name *string) (Pokemon, error) {
	fullURL := baseURL + "/pokemon/" + *name

	if cachePokemonResp, ok := c.cache.Get(*name); ok {
		var pokemonResp Pokemon
		err := json.Unmarshal(cachePokemonResp, &pokemonResp)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)

	if err != nil {
		return Pokemon{}, err
	}

	var pokemonResp Pokemon
	err = json.Unmarshal(data, &pokemonResp)
	if err != nil {
		return Pokemon{}, err
	}
	c.cache.Add(*name, data)
	return pokemonResp, nil
}
