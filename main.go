package main

import (
	blackjack "Dayanfreitas/BlackJack/src"
	"fmt"

	"github.com/adamclerk/deck"
)

func main() {

	// blackjack.New(blackjack.OptionsBlackJack{})
	blackjack.Hello()
	// ExampleWithCards()
	// blackjack.BlackJack.Cards
}

func ExampleWithCards() {
	d, _ := deck.New(deck.WithCards(deck.NewCard(deck.TWO, deck.CLUB), deck.NewCard(deck.QUEEN, deck.DIAMOND)), deck.Unshuffled)
	GetPoint(d.Cards)

	// fmt.Printf("%s", d)

}

func GetPoint(d []deck.Card) {
	// pointMap := make(map[string]int)

	// pointMap["2"] = 2

	// pointMap["10"] = 10
	// pointMap["11"] = 10
	// pointMap["13"] = 10

	// fmt.Println("pointMap")
	// fmt.Println(pointMap)
	//
	// var cardv int

	for _, v := range d {
		// fmt.Printf("=> %d\n", pointMap[string(v.Face())])
		fmt.Println(v)
		// fmt.Println(v.Face())
	}
}
