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
	// Check for help
	if len(os.Args) != 2 || 
	os.Args[1] == "help" || 
	os.Args[1] == "-h"   || 
	os.Args[1] == "--help" {
		fmt.Printf("%s", HELP_MESSAGE)
		return
	}

	// Parse args
	dX := os.Args[1]

	num, sides, err := parseArgs(dX)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < num; i++ {
		outcome, err := rollDie(sides)
		if err != nil {
			fmt.Println(sides)
			return
		}

		fmt.Printf("Rolled: %d\n", outcome)
	}
}

func rollDie(sides int) (int, error) {
	// Check for valid number of sides
	if sides < 1 {
		return 0, fmt.Errorf("Invalid number of sides: %d", sides)
	}

	// Roll the die
	return rand.Intn(sides) + 1, nil
}

func parseArgs(dX string) (int, int, error) {
	// Check for version flag
	if dX == "-v" || dX == "--version" {
		printVersion()
		os.Exit(0)
	}

	// Check for valid number of dice
	if dX[0] == 'd' {
		dX = dX[1:]
		sides, err := strconv.Atoi(dX)
		if err != nil {
			return 0, 0, err
		}
		return 1, sides, nil
	}

	// Parse the number of dice and sides
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

func printVersion() {
	fmt.Printf("Version: %s\nBuild Date: %s\nCommit: %s\n", Version, BuildDate, Commit)
}
