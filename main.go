package main

import (
	"flag"
	"math/rand/v2"
	"os"
	"strings"
)

func main() {
	flag.Parse()
	hand := strings.ToLower(flag.Arg(0))
	var enemy string
	//MEMO: []stringにすればswitch文はいらない
	switch rand.IntN(3) {
	case 0:
		enemy = "rock"
	case 1:
		enemy = "paper"
	case 2:
		enemy = "scissors"
	default:
		panic("unreachable case")
	}

	switch hand {
	case "rock":
		if enemy == "rock" {
			os.Exit(3)
		}
		if enemy == "paper" {
			os.Exit(1)
		}
		if enemy == "scissors" {
			os.Exit(0)
		}
	case "paper":
		if enemy == "rock" {
			os.Exit(0)
		}
		if enemy == "paper" {
			os.Exit(3)
		}
		if enemy == "scissors" {
			os.Exit(1)
		}
	case "scissors":
		if enemy == "rock" {
			os.Exit(1)
		}
		if enemy == "paper" {
			os.Exit(0)
		}
		if enemy == "scissors" {
			os.Exit(3)
		}
	}
}
