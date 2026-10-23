package main

import (
	"flag"
	"os"
	"strings"

	rps "github.com/Suiren91/rps_cli"
)

func main() {
	flag.Parse()
	arg := strings.ToLower(flag.Arg(0))
	hand, _ := rps.ParseHand(arg)
	oppo := rps.New()

	res := rps.Judge(hand, oppo)
	switch res {
	case rps.Win:
		os.Exit(0)
	case rps.Lose:
		os.Exit(1)
	case rps.Draw:
		os.Exit(3)
	}
}
