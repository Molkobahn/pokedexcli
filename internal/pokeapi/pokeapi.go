package pokeapi

import(
	"fmt"
	"net/http"
	"encoding/json"
	"io"
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

func GetRequest(url string) PokeStruct {
	res, err := http.Get(url)
	if err != nil {
		fmt.Print(err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
	fmt.Print(err)
	}

	var pokeStruct PokeStruct
	if err := json.Unmarshal(data, &pokeStruct); err != nil {
		fmt.Print(err)
	}
	

	return pokeStruct
	
}