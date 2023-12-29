package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

const HELP_MESSAGE = "Usage: roll <num>d<sides>\n"

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("%s", HELP_MESSAGE)
		return
	}

	if os.Args[1] == "help" || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Printf("%s", HELP_MESSAGE)
		return
	}


	dX := os.Args[1]

	num, sides, err := parseArgs(dX)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Rolling %dd%d\n", num, sides)

	for i := 0; i < num; i++ {
		outcome, err := rollDie(sides)
		if err != nil {
			fmt.Println(sides)
			return
		}

		fmt.Printf("Rolled a %d\n", outcome)
	}
}

func rollDie(sides int) (int, error) {
	if sides < 1 {
		return 0, fmt.Errorf("Invalid number of sides: %d", sides)
	}
	return rand.Intn(sides) + 1, nil
}

func parseArgs(dX string) (int, int, error) {
	if dX[0] == 'd' {
		dX = dX[1:]
		sides, err := strconv.Atoi(dX)
		if err != nil {
			return 0, 0, err
		}
		return 1, sides, nil
	}

	parts := strings.Split(dX, "d")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("Invalid die format: %s", dX)
	}

	num, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, err
	}

	sides, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, err
	}

	return num, sides, nil

}
