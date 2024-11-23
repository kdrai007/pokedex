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
	if respData, ok := c.cache.Get(fullURL); ok {
		locationResp := shallowApiResponse{}
		err := json.Unmarshal(respData, &locationResp)
		if err == nil {
			return locationResp, nil
		}
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

	c.cache.Add(fullURL, data)
	return locationApiResp, nil
}
