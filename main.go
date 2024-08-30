package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type PokemonInfo struct {
	Name  string   `json:"name"`
	Types []string `json:"types"`
	Stats []int    `json:"stats"`
	Image string   `json:"image"`
}

func fetchPokemonData(pokemonID int) (PokemonInfo, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%d", pokemonID)
	resp, err := http.Get(url)
	if err != nil {
		return PokemonInfo{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return PokemonInfo{}, fmt.Errorf("failed to fetch data, status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return PokemonInfo{}, err
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return PokemonInfo{}, err
	}

	name := result["name"].(string)
	typesArray := result["types"].([]interface{})
	var types []string
	for _, t := range typesArray {
		typeInfo := t.(map[string]interface{})
		typeName := typeInfo["type"].(map[string]interface{})["name"].(string)
		types = append(types, typeName)
	}

	statsArray := result["stats"].([]interface{})
	var stats []int
	for _, s := range statsArray {
		baseStat := int(s.(map[string]interface{})["base_stat"].(float64))
		stats = append(stats, baseStat)
	}

	imageURL := result["sprites"].(map[string]interface{})["front_default"].(string)

	return PokemonInfo{
		Name:  name,
		Types: types,
		Stats: stats,
		Image: imageURL,
	}, nil
}

func saveToJSON(filename string, pokemons []PokemonInfo) error {
	data, err := json.MarshalIndent(pokemons, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	var n int = 50

	fmt.Print("Number of pokemons to fetch: ")
	fmt.Scan(&n)

	var pokemons []PokemonInfo
	for i := 1; i <= n; i++ {
		pokemon, err := fetchPokemonData(i)
		if err != nil {
			fmt.Println(err)
			continue
		}
		pokemons = append(pokemons, pokemon)
	}

	err := saveToJSON("pokemons.json", pokemons)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Data saved to pokemons.json")
}
