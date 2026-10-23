// Package rpsはじゃんけんの手や勝敗判定の関数を提供する
package rps

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

type Hand int

const (
	Rock Hand = iota
	Paper
	Scissors
)

type Result int

const (
	Win Result = iota
	Lose
	Draw
)

func (h Hand) String() string {
	switch h {
	case Rock:
		return "Rock"
	case Paper:
		return "Paper"
	case Scissors:
		return "Scissors"
	default:
		return ""
	}
}

func (r Result) String() string {
	switch r {
	case Win:
		return "Win"
	case Lose:
		return "Lose"
	case Draw:
		return "Draw"
	default:
		return ""
	}
}

// ParseHand は受け取った文字列に応じたHand型を返す
// rock, paper, scissorsのみ受け付ける(大文字，小文字は問わない)
func ParseHand(s string) (Hand, error) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "rock":
		return Rock, nil
	case "paper":
		return Paper, nil
	case "scissors":
		return Scissors, nil
	default:
		return 0, fmt.Errorf("invalid hand: %q", s)
	}
}

// New はランダムなHand型の値を返す
func New() Hand {
	return Hand(rand.IntN(3))
}

// Judge はじゃんけんの勝敗を返す
// 渡した手が不正な場合は-1を返す
func Judge(hand, opponentHand Hand) Result {
	if hand == opponentHand {
		return Draw
	}
	switch hand {
	case Rock:
		if opponentHand == Paper {
			return Lose
		}
		return Win
	case Scissors:
		if opponentHand == Rock {
			return Lose
		}
		return Win
	case Paper:
		if opponentHand == Scissors {
			return Lose
		}
		return Win
	}
	return -1
}
