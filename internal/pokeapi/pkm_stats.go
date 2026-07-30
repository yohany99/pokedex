package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type PokemonStatResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Exp  int    `json:"base_experience"`
}

type pokemon struct {
	name string
}

type PokemonStorage struct {
	PC map[string]pokemon
}

func (c *Client) CatchPokemon(name string) (PokemonStatResponse, error) {
	baseURL := "https://pokeapi.co/api/v2/pokemon/"
	baseURL += name
	if val, ok := c.pokeCache.Get(baseURL); ok {
		pkmStatResp := PokemonStatResponse{}
		err := json.Unmarshal(val, &pkmStatResp)
		if err != nil {
			return PokemonStatResponse{}, err
		}
		return pkmStatResp, nil
	}
	resp, err := http.Get(baseURL)
	if err != nil {
		return PokemonStatResponse{}, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonStatResponse{}, err
	}
	resp.Body.Close()
	if resp.StatusCode > 299 {
		return PokemonStatResponse{}, fmt.Errorf("response failed with status code %d and\nbody: %s\n", resp.StatusCode, body)
	}
	pkmStatResp := PokemonStatResponse{}
	err = json.Unmarshal(body, &pkmStatResp)
	if err != nil {
		return PokemonStatResponse{}, err
	}
	c.pokeCache.Add(baseURL, body)
	return pkmStatResp, nil
}
