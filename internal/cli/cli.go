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

func Run() {
	prompt := "Pokedex > "
	inputScanner := bufio.NewScanner(os.Stdin)
	config := api.Config{}
	api.Cache = pokecache.NewCache(time.Minute)
	pokeDex = make(map[string]api.Pokemon)
	for {
		fmt.Print(prompt)
		ok := inputScanner.Scan()
		if !ok {
			fmt.Println("EXIT")
			return
		}

		words := cleanInput(inputScanner.Text())
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
