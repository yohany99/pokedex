package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type PokemonResponse struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func (c *Client) ListPokemon(url string) (PokemonResponse, error) {
	if val, ok := c.pokeCache.Get(url); ok {
		pokemonResp := PokemonResponse{}
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return PokemonResponse{}, err
		}
		return pokemonResp, nil
	}
	resp, err := http.Get(url)
	if err != nil {
		return PokemonResponse{}, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonResponse{}, err
	}
	resp.Body.Close()
	if resp.StatusCode > 299 {
		return PokemonResponse{}, fmt.Errorf("response failed with status code %d and\nbody: %s\n", resp.StatusCode, body)
	}
	pokemonResp := PokemonResponse{}
	err = json.Unmarshal(body, &pokemonResp)
	if err != nil {
		return PokemonResponse{}, err
	}
	c.pokeCache.Add(url, body)
	return pokemonResp, nil
}
