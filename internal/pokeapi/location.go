package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pokeLoc string) (pokeLocation, error) {
	url := baseURL + "/location-area"
	if pokeLoc != "" {
		url = pokeLoc
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return pokeLocation{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return pokeLocation{}, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return pokeLocation{}, err
	}

	pokeRes := pokeLocation{}
	err = json.Unmarshal(data, &pokeRes)
	if err != nil {
		return pokeLocation{}, err
	}

	return pokeRes, nil
}
