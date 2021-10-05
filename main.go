package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/koseburak/marvel/config"
	"github.com/koseburak/marvel/marvel"
)

func main() {

	// get entered value from terminal
	character := ReadCharacterFromTerminal()

	conf, err := config.Config()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	defaultHTTPClient := &http.Client{}

	marvelClient := marvel.NewMarvelClient(conf, defaultHTTPClient)
	// fetch the Characters from Marvel API
	characters, err := marvelClient.GetCharacters(character)
	if err != nil {
		log.Println("Got error while fetching characters from Marvel API", err)
	}

	// marshal the http response data
	data, err := json.MarshalIndent(characters, "", "  ")
	if err != nil {
		log.Println("Got error while pretty print", err)
	}

	// print searched character information to terminal
	fmt.Println(string(data))
}
