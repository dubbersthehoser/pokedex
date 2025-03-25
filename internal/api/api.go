package api

import (
	"net/http"
	"encoding/json"
	"io"
	"net/url"
)

const (
	EPBase string = "https://pokeapi.co/api/v2"
	EPLocationArea = "https://pokeapi.co/api/v2/location-area"
	EPPokemon = "https://pokeapi.co/api/v2/pokemon"
)

type Cacher interface {
	Add(string, []byte)
	Get(string) ([]byte, bool)
} 

type nilCache struct{}
func (n nilCache) Get(key string) ([]byte, bool) {
	return nil, false
}
func (n nilCache) Add(key string, data []byte) {
	return
}

var Cache Cacher = nilCache{}

type Config struct {
	Resource string
	Next     string
	Previous string
}

func getData(url string) ([]byte, error){
	gotCache := false
	bytes, gotCache := Cache.Get(url)
	if !gotCache {
		req, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer req.Body.Close()
		bytes, err = io.ReadAll(req.Body)
		if err != nil {
			return nil, err
		}
	}
	return bytes, nil
}

func GetLocationAreas(config *Config) ([]NamedAPIResource, error) {
	var epUrl string

	if config.Next != "" && config.Resource == EPLocationArea {
		epUrl = config.Next
	} else {
		epUrl = EPLocationArea
	}

	data, err := getData(epUrl)
	if err != nil {
		return nil, err
	}

	namedResList := NamedAPIResourceList{}
	err = json.Unmarshal(data, &namedResList)
	if err != nil {
		return nil, err
	}

	config.Resource = EPLocationArea
	if namedResList.Previous != nil {
		config.Previous = *namedResList.Previous
	} 
	if namedResList.Next != nil {
		config.Next = *namedResList.Next
	}
	return namedResList.Results, nil
}
func GetLocationArea(config *Config, id string) (LocationArea, error) {
	 epUrl, err := url.JoinPath(EPLocationArea, id)
	 if err != nil {
	 	return LocationArea{}, err
	 }
	data, err := getData(epUrl)
	if err != nil {
	 	return LocationArea{}, err
	}
	locationArea := LocationArea{}
	err = json.Unmarshal(data, &locationArea)
	if err != nil {
	 	return LocationArea{}, err
	}
	return locationArea, nil
}

func GetPokemon(config *Config, id string) (Pokemon, error) {
	epUrl, err := url.JoinPath(EPPokemon, id)
	if err != nil {
		return Pokemon{}, err
	}
	data, err := getData(epUrl)
	if err != nil {
		return Pokemon{}, err
	}
	pokemon := Pokemon{}
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}
	return pokemon, nil
}
