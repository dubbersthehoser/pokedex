package api

type APIResource struct {
	Url *string `json:"url"`
}

type NamedAPIResource struct {
	Name *string `json:"name"`
	Url  *string `json:"url"`
}

type Name struct {
	Name     *string          `json:"name"`
	Language NamedAPIResource `json:"language"`
}

type APIResourceList struct {
	Count int        `json:"count"`
	Next  *string    `json:"next"`
	Previous *string `json:"previous"`
	Results  []APIResource `json:`
}

type NamedAPIResourceList struct {
	Count    int                `json:"count"`
	Next     *string            `json:"next"`
	Previous *string            `json:"previous"`
	Results  []NamedAPIResource `json:"results"`
}

type Description struct {
	Description *string          `json:"description"`
	Language    NamedAPIResource `json:"language"`
}

type Effect struct {
	Effect   *string          `json:"effect"`
	Language NamedAPIResource `json:"language"`
}

type Encounter struct {
	MinLevel        int                `json:"min_level"`
	MaxLevel        int                `json:"max_level"`
	ConditionValues []NamedAPIResource `json:"condition_values"`
	Chance          int                `json:"chance"`
	Method          NamedAPIResource   `json:"method"`
}

type VersionEncounterDetail struct {
	Version          NamedAPIResource `json:"version"`
	MaxChance        int              `json:"max_chance"`
	EncounterDetails []Encounter      `json:"encounter_details"`
}

type PokemonEncounter struct {
	Pokemon        NamedAPIResource         `json:"pokemon"`
	VersionDetails []VersionEncounterDetail `json:"version_details"`
}


type EncounterMethodRate struct {
	EncounterMethod NamedAPIResource         `json:"encounter_method"`
	VersionDetails  []VersionEncounterDetail `json"version_details"`
}


type LocationArea struct {
	Id                   int                   `json:"id"`
	Name                 *string               `json:"name"`
	GameIndex            int                   `json:"game_index"`
	EncounterMethodRates []EncounterMethodRate `json:"encounter_method_rates"`
	Location             NamedAPIResource      `json:"location"`
	Names                []Name                `json:"names"` 
	PokemonEncounters   []PokemonEncounter    `json:"pokemon_encounters"`
}

type PokemonAbility struct {
	IsHidden bool `json:"is_hidden"`
	Slot int `json:"slot"`
	Ability NamedAPIResource `json:"ability"`
}

type VersionGameIndex struct {
	GameIndex int `json:"game_index"`
	Version NamedAPIResource `json:"version"`
}

type PokemonHeldItem struct {
	Item NamedAPIResource `json:"item"`
	VersionDetails []PokemonHeldItemVersion `json:"version_details"`
}

type PokemonHeldItemVersion struct {
	Version NamedAPIResource `json:"version"`
	Rarity int `json:"rarity"`
}

type PokemonMove struct {
	Move NamedAPIResource `json:"move"`
	VersionGroupDetails []PokemonMoveVersion `json:"version_group_details"`
}

type PokemonMoveVersion struct {
	MoveLearnMethod NamedAPIResource `json:"move_learn_method"`
	VersionGroup NamedAPIResource `json:"version_group"`
	LevelLearnedAt int `json:"level_learned_at"`
}

type PokemonType struct {
	Slot int `json:"slot"`
	Type NamedAPIResource `json:"type"`
}

type PokemonTypePast struct {
	Generation NamedAPIResource `json:"generation"`
	Types []PokemonType `json:"types"`
}

type PokemonSprites struct {
	FrontDefault *string `json:"front_default"`
	FrontShiny *string `json:"front_shiny"`
	FrontFemale *string `json:"front_female"`
	FrontShinyFemale *string `json:"front_shiny_female"`
	BackDefault *string `json:"back_default"`
	BackShiny *string `json:"back_shiny"`
	BackFemale *string `json:"back_female"`
	BackShinyFemale *string `json:"back_shiny_female"`
}

type PokemonCries struct {
	Lastest *string `json:"latest"`
	Legacy *string `json:"legacy"`
}

type PokemonStat struct {
	Stat NamedAPIResource `json:"stat"`
	Effort int  `json:"effort"`
	BaseStat int `json:"base_stat"`
}

type Pokemon struct {
	Id int `json:"id"`
	Name *string `json:"name"`
	BaseExperience int `json:"base_experience"`
	Height int `json:"height"`
	IsDefault bool `json:"is_default"`
	Order int `json:"order"`
	Weight int `json:"weight"`
	Abilities []PokemonAbility `json:"abilities"`
	Forms []NamedAPIResource `json:"forms"`
	GameIndices []VersionGameIndex `json:"game_indices"`
	GameItems []PokemonHeldItem `json:"game_items"`
	LocationAreaEncounters *string `json:"location_area_encounters"`
	Moves []PokemonMove `json:"moves"`
	PastTypes []PokemonTypePast `json:"past_types"`
	Sprites PokemonSprites `json:"sprites"`
	Cries PokemonCries `json:"cries"`
	Species NamedAPIResource `json:"species"`
	Stats []PokemonStat `json:"stats"`
	Types []PokemonType `json:"types"`
}
