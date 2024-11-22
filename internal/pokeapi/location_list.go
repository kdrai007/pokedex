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
	return locationApiResp, nil
}
