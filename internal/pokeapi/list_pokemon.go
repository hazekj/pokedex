package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListPokemon(area string) ([]pokemonInArea, error) {
	url := baseURL + "/location-area/" + area

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return []pokemonInArea{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return []pokemonInArea{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return []pokemonInArea{}, err
	}

	pokeRes := pokeArea{}
	err = json.Unmarshal(data, &pokeRes)
	if err != nil {
		return []pokemonInArea{}, err
	}

	return pokeRes.PokemonEncounters, nil

}
