package pokeapi

type locationData struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type shallowApiResponse struct {
	Count   int            `json:"count"`
	Next    string         `json:"next"`
	Prev    string         `json:"previous"`
	Results []locationData `json:"results"`
}
