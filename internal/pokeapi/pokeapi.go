package pokeapi

import(
	"fmt"
	"net/http"
	"encoding/json"
	"io"
	"github.com/molkobahn/pokedexcli/internal/pokecache"
)

type PokeStruct struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetRequest(url string, cache *pokecache.Cache) (PokeStruct, error) {
	fmt.Println("The funtion was entered")
	var data []byte

	cacheData, status := cache.Get(url)
	fmt.Println("cache was checked")
	if status{
		data = cacheData
	} else {
	
		res, err := http.Get(url)
		fmt.Println("http Get request was made")
		if err != nil {
			return PokeStruct{}, err
		}
		defer res.Body.Close()

		newData, err := io.ReadAll(res.Body)
		fmt.Println("response was read")
		if err != nil {
			return PokeStruct{}, err
		}

		data = newData

		cache.Add(url, data)
		fmt.Println("added to cache")
	}

	var pokeStruct PokeStruct
	if err := json.Unmarshal(data, &pokeStruct); err != nil {
		return PokeStruct{}, err
	}

	return pokeStruct, nil
	
}