package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemonInfo(name string) (CatchPokemonInfo, error) {
	url := baseURL + "/pokemon/" + name + "/"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return CatchPokemonInfo{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return CatchPokemonInfo{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return CatchPokemonInfo{}, err
	}

	pokeInfo := CatchPokemonInfo{}
	err = json.Unmarshal(data, &pokeInfo)
	if err != nil {
		return CatchPokemonInfo{}, err
	}

	return pokeInfo, nil
}
