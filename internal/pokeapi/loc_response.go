package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocResponse struct {
	Next      *string    `json:"next"`
	Previous  *string    `json:"previous"`
	Locations []location `json:"results"`
}

type location struct {
	Name string `json:"name"`
}

func (c *Client) ListLocations(url *string) (LocResponse, error) {
	defaultURL := "https://pokeapi.co/api/v2/location-area/"
	if url == nil {
		url = &defaultURL
	}
	if val, ok := c.pokeCache.Get(*url); ok {
		locResponse := LocResponse{}
		err := json.Unmarshal(val, &locResponse)
		if err != nil {
			return LocResponse{}, err
		}
		return locResponse, nil
	}
	resp, err := http.Get(*url)
	if err != nil {
		return LocResponse{}, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocResponse{}, err
	}
	resp.Body.Close()
	if resp.StatusCode > 299 {
		return LocResponse{}, fmt.Errorf("response failed with status code: %d and\nbody: %s\n", resp.StatusCode, body)
	}
	locResponse := LocResponse{}
	err = json.Unmarshal(body, &locResponse)
	if err != nil {
		return LocResponse{}, err
	}
	c.pokeCache.Add(*url, body)
	return locResponse, nil
}
