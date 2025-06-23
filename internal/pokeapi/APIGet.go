package pokeapi

import (
	"encoding/json"
	"net/http"
)

func GetMapData(url string) (Locations, error) {
	res, err := http.Get(url)
	if err != nil {
		return Locations{}, err
	}
	defer res.Body.Close()

	var locations Locations
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&locations); err != nil {
		return Locations{}, err
	}

	return locations, nil
}
