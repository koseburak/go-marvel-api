package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ReadCharacterFromTerminal returns the value entered from the terminal
func ReadCharacterFromTerminal() string {
	colorReset := "\033[0m"
	colorYellow := "\033[33m"
	// colorGreen := "\033[32m"
	// colorBlue := "\033[34m"

	// read and get argument input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(string(colorYellow), "Please enter the Marvel Character name you want to search: ")
	enteredValue, _ := reader.ReadString('\n')

	// convert input argument from CRLF to LF
	enteredValue = strings.Replace(enteredValue, "\n", "", -1)
	enteredValue = strings.Replace(enteredValue, "\r", "", -1)
	fmt.Println(string(colorReset), "")

	return enteredValue
}
