package cli

import (
	"fmt"
	"bufio"
	"os"
	"time"

	"github.com/dubbersthehoser/pokedex/internal/api"
	"github.com/dubbersthehoser/pokedex/internal/pokecache"
)

var pokeDex map[string]api.Pokemon

var histList []string

const prompt string = "Pokedex > "

func Run() {
	inputScanner := bufio.NewScanner(os.Stdin)
	config := api.Config{}
	api.Cache = pokecache.NewCache(time.Minute)

	histList = make([]string, 0)
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

		histList = append(histList, line)
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

