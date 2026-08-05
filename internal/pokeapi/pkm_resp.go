package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type PokemonEncountersResp struct {
	Encounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func (c *Client) ListPokemon(url string) (PokemonEncountersResp, error) {
	if val, ok := c.pokeCache.Get(url); ok {
		pokemonResp := PokemonEncountersResp{}
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return PokemonEncountersResp{}, err
		}
		return pokemonResp, nil
	}
	resp, err := http.Get(url)
	if err != nil {
		return PokemonEncountersResp{}, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonEncountersResp{}, err
	}
	resp.Body.Close()
	if resp.StatusCode > 299 {
		return PokemonEncountersResp{}, fmt.Errorf("response failed with status code %d and\nbody: %s\n", resp.StatusCode, body)
	}
	pokemonResp := PokemonEncountersResp{}
	err = json.Unmarshal(body, &pokemonResp)
	if err != nil {
		return PokemonEncountersResp{}, err
	}
	c.pokeCache.Add(url, body)
	return pokemonResp, nil
}
