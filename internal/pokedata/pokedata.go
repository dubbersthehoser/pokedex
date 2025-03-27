package pokedata

import (
	"github.com/dubbersthehoser/pokedex/internal/api"
)

type PlayerData struct {
	Version int `json:"version"`
	CharacterName string                 `json:"character_name"`
	CoughtPokemon map[string]api.Pokemon `json:"cought_pokemon"`
}

func NewPlayerData(playerName string) *PlayerData {
	return &PlayerData{
		Version: 0,
		CoughtPokemon: make(map[string]api.Pokemon),
		CharacterName: playerName,
	}
}

