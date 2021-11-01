package main

import (
	"fmt"

	"github.com/adamclerk/deck"
)

func main() {
	d, _ := deck.New()

	player1Hand, _ := deck.New(deck.Empty)
	player2Hand, _ := deck.New(deck.Empty)

	// // deck.Debugf(false, "INIT")
	// card1 := deck.NewCard(deck.TWO, deck.SPADE)
	// deck.Debugf(true, "DECK_>\n%s", d)

	fmt.Println(len(d.Cards))
	fmt.Print("player1Hand->\n")
	fmt.Print(player1Hand)

	fmt.Print("player2Hand->\n")
	fmt.Print(player2Hand)
	fmt.Print()

	// fmt.Printf(""d)
	// fmt.Print(deck.TWO)
}
