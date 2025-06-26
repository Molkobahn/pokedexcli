package pokeapi

import(
	"fmt"
	"net/http"
	"encoding/json"
	"io"
	"github.com/molkobahn/pokedexcli/internal/pokecache"
)

type RespLocations struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}


type RespLocationDetails struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	GameIndex            int    `json:"game_index"`
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
			MaxChance        int `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel        int   `json:"min_level"`
				MaxLevel        int   `json:"max_level"`
				ConditionValues []any `json:"condition_values"`
				Chance          int   `json:"chance"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}


func ListLocations(url string, cache *pokecache.Cache) (RespLocations, error) {
	var data []byte

	cacheData, status := cache.Get(url)
	
	if status{
		data = cacheData
	} else {
		res, err := http.Get(url)
		if err != nil {
			return RespLocations{}, err
		}
		defer res.Body.Close()

		newData, err := io.ReadAll(res.Body)
		if err != nil {
			return RespLocations{}, err
		}

		data = newData

		cache.Add(url, data)
	}

	var pokeStruct RespLocations
	if err := json.Unmarshal(data, &pokeStruct); err != nil {
		return RespLocations{}, err
	}

	return pokeStruct, nil
}


func LocationDetails(url string, cache *pokecache.Cache) (RespLocationDetails, error) {
	var data []byte

	cacheData, status := cache.Get(url)
	
	if status{
		data = cacheData
	} else {
		res, err := http.Get(url)
		if err != nil {
			return RespLocationDetails{}, err
		}
		defer res.Body.Close()

		if res.StatusCode != 200 {
			return RespLocationDetails{}, fmt.Errorf("API request failed with status %d: location '%s' not found", res.StatusCode, url)
		}

		newData, err := io.ReadAll(res.Body)
		if err != nil {
			return RespLocationDetails{}, err
		}

		data = newData

		cache.Add(url, data)
	}

	var pokeStruct RespLocationDetails
	if err := json.Unmarshal(data, &pokeStruct); err != nil {
		return RespLocationDetails{}, err
	}

	return pokeStruct, nil

}