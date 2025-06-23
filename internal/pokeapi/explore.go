package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ExploreArea(area string) (AreaPokemon, error) {
	url := baseURL + "/location-area/" + area + "/"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return AreaPokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return AreaPokemon{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return AreaPokemon{}, err
	}

	pokemon := AreaPokemon{}
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return AreaPokemon{}, err
	}

	return pokemon, nil
}
