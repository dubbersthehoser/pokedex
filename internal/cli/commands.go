package cli

import (
	"fmt"
	"os"
	"strings"
	"math/rand"
	"encoding/json"

	"github.com/dubbersthehoser/pokedex/internal/api"
)

type cliProc func(*api.Config, ...string) error

type cliCommand struct {
	name string
	description string
	callback cliProc
}

func commandExit(config *api.Config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *api.Config, args ...string) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for name, cmd := range commandMapping {
		fmt.Printf("%s: %s\n", name, cmd.description)
	}
	return nil
}

func commandMap(config *api.Config, args ...string) error {
	results, err := api.GetLocationAreas(config)
	if err != nil {
		return err
	}
	for _, re := range results {
		fields := strings.Split(*re.Url, "/")
		id := fields[len(fields)-2]
		fmt.Printf("%s: %s\n", id, *re.Name)
	}
	return nil
}
func commandMapB(config *api.Config, args ...string) error {
	
	if config.Previous == "" {
		return fmt.Errorf("no previous page")
	}

	if config.Resource == api.EPLocationArea {
		config.Next = config.Previous // Gets previous page
	}
	results, err := api.GetLocationAreas(config)

	if err != nil {
		return err
	}

	for _, re := range results {
		fields := strings.Split(*re.Url, "/")
		id := fields[len(fields)-2]
		fmt.Printf("%s: %s\n", id, *re.Name)
	}
	return nil
}

func commandExplore(config *api.Config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("location not given")
	}
	location := args[0]
	area, err := api.GetLocationArea(config, location)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s.\n", *area.Name)
	fmt.Println("Found Pokemon:")
	for _, encounter := range area.PokemonEncounters {
		name := *encounter.Pokemon.Name
		fmt.Printf(" - %s\n", name)
	}
	return nil
}

func commandCatch(config *api.Config, args ...string) error {
	if len(args) == 0 {
		fmt.Println("Pokemon was not given")
	} 
	pokemon, ok := playerData.CoughtPokemon[args[0]]
	if ok {
		fmt.Printf("%s already cought\n", *pokemon.Name)
		return nil
	}

	pokemon, err := api.GetPokemon(config, args[0])
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", *pokemon.Name)

	threshold := 100 - (pokemon.BaseExperience / 10)
	if threshold < 5 {
		threshold = 5
	} else if threshold > 95 {
		threshold = 95
	}
	chance := rand.Intn(101)
	isCought := chance <= threshold

	if isCought {
		playerData.CoughtPokemon[*pokemon.Name] = pokemon
		fmt.Printf("%s was caught!\n", *pokemon.Name)
	} else {
		fmt.Printf("%s escaped!\n", *pokemon.Name)
	}
	return nil
}

func commandInspect(config *api.Config, args ...string) error {
	if len(args) == 0 {
		fmt.Println("Pokemon was not given")
	} 
	pokemon, ok := playerData.CoughtPokemon[args[0]]
	if !ok {
		fmt.Printf("%s has not been caught\n", *pokemon.Name)
		return nil
	}

	fmt.Println("Name: ", *pokemon.Name)
	fmt.Println("Height: ", pokemon.Height)
	fmt.Println("Weight: ", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", *stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typ := range pokemon.Types {
		fmt.Printf("  - %s\n", *typ.Type.Name)
	}
	return nil
}

func commandPokedex(config *api.Config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for _, pokemon := range playerData.CoughtPokemon {
		fmt.Printf(" - %s\n", *pokemon.Name)
	}
	return nil
}

func commandSave(config *api.Config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("No save name given")
	}
	filename := fmt.Sprintf("%s.json", args[0])
	file, err := os.Create(filename)
	defer file.Close()
	if err != nil {
		return err
	}
	bytes, err := json.Marshal(playerData)
	if err != nil {
		return err
	}

	_, err = file.Write(bytes)

	if err != nil {
		return err
	}
	fmt.Println("Character saved")
	return nil
}

var commandMapping = map[string]cliCommand{}
func initCommandMapping() {
	commandMapping["exit"] = cliCommand{
		name: "exit",
		description: "Exit the Pokedex",
		callback: commandExit,
	}
	commandMapping["help"] = cliCommand{
		name: "help",
		description: "List commands and their usage",
		callback: commandHelp,
	}
	commandMapping["map"] = cliCommand{
		name: "map",
		description: "Page 20 map locations",
		callback: commandMap,
	}
	commandMapping["mapb"] = cliCommand{
		name: "mapb",
		description: "Back page 20 locations",
		callback: commandMapB,
	}
	commandMapping["explore"] = cliCommand{
		name: "explore",
		description: "List out pokemon from a given area",
		callback: commandExplore,
	}
	commandMapping["catch"] = cliCommand{
		name: "catch",
		description: "Catch a given pokemon",
		callback: commandCatch,
	}
	commandMapping["inspect"] = cliCommand{
		name: "inspect",
		description: "Inspect a cought pokemon",
		callback: commandInspect,
	}
	commandMapping["pokedex"] = cliCommand{
		name: "pokedex",
		description: "List cought pokemon",
		callback: commandPokedex,
	}
	commandMapping["save"] = cliCommand{
		name: "save",
		description: "Save current character",
		callback: commandSave,
	}
}

