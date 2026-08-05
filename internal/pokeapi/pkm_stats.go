package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Pokemon struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Exp    int    `json:"base_experience"`
	Height int    `json:"height"`
	Stats  []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
}

func (c *Client) CatchPokemon(name string) (Pokemon, error) {
	baseURL := "https://pokeapi.co/api/v2/pokemon/"
	baseURL += name
	if val, ok := c.pokeCache.Get(baseURL); ok {
		pkmStatResp := Pokemon{}
		err := json.Unmarshal(val, &pkmStatResp)
		if err != nil {
			return Pokemon{}, err
		}
		return pkmStatResp, nil
	}
	resp, err := http.Get(baseURL)
	if err != nil {
		return Pokemon{}, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}
	resp.Body.Close()
	if resp.StatusCode > 299 {
		return Pokemon{}, fmt.Errorf("response failed with status code %d and\nbody: %s\n", resp.StatusCode, body)
	}
	pokemon := Pokemon{}
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}
	c.pokeCache.Add(baseURL, body)
	return pokemon, nil
}

func PrintStats(pkm Pokemon) {
	fmt.Printf("Name: %s\n", pkm.Name)
	fmt.Printf("Height: %d\n", pkm.Height)
	fmt.Printf("Weight: %d\n", pkm.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range pkm.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, t := range pkm.Types {
		fmt.Printf("  -%s\n", t.Type.Name)
	}
}
