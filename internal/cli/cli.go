package cli

import (
	"fmt"
	"bufio"
	"os"
	"time"
	"strings"
	"encoding/json"

	"github.com/dubbersthehoser/pokedex/internal/api"
	"github.com/dubbersthehoser/pokedex/internal/pokecache"
	"github.com/dubbersthehoser/pokedex/internal/pokedata"
)

var playerData pokedata.PlayerData

const prompt string = "Pokedex > "

func Run() {
	if len(os.Args) == 2 {
		openPlayerFile()
		activeGame()
		return
	} else {
		newPlayerFile()
		activeGame()
		return
	}
}

func openPlayerFile() {
	playerfile := os.Args[1]

	if !strings.HasSuffix(playerfile, ".json") {
		fmt.Printf("%s: player file is not json\n", playerfile)
		return
	}

	file, err := os.Open(playerfile)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&playerData)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func newPlayerFile() {
	inputScanner := bufio.NewScanner(os.Stdin)
	var playerName string
	for {
		fmt.Printf("Character name? ")
		ok := inputScanner.Scan()
		if !ok {
			os.Exit(0)
		}
		playerName = inputScanner.Text()

		if playerName != "" {
			playerData = *pokedata.NewPlayerData(playerName)
			return
		}
	}
}

func activeGame() {
	inputScanner := bufio.NewScanner(os.Stdin)
	api.Cache = pokecache.NewCache(time.Minute)
	config := api.Config{}
	for {
		fmt.Print(prompt)
		ok := inputScanner.Scan()
		var line string
		if !ok {
			line = "exit"
		} else {
			line = inputScanner.Text()
		}

		words := cleanInput(line)
		if len(words) == 0 {
			continue
		}

		fword := words[0]
		proc, err := commandLookUp(fword)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		err = proc(&config, words[1:]...)
		if err != nil{
			fmt.Println(err.Error())
		}
	}
}

