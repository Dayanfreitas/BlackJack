package blackjack

import (
	"fmt"

	"github.com/adamclerk/deck"
)

const (
	OddsDefault       = float64(3.0 / 2.0)
	DeckNumberDefault = 1
)

type Player struct {
	Id   int
	Name string
	Hand []deck.Card
}

type Dealer struct {
	Player
	odds  float64
	Point int
}

type BlackJack struct {
	// numberPlayers int
	DeckNumber int
	Odds       float64
	Deck       deck.Deck
	Cards      []deck.Card
	Dealer     Dealer
}

func Hello() {
	d1, _ := deck.New()
	d2, _ := deck.New()

	black_jack := BlackJack{Deck: *d1}

	fmt.Printf("BLACK JACK -> %v", black_jack)
	index := 0
	for _, v := range d2.Cards {
		index++
		fmt.Println(v.Face())
	}

	fmt.Println("Count ->", index)

	// p := Player{}
	// dealer := Dealer{}

	// fmt.Println()
	// dealer.Name = "Dealer"
	// dealer.Hand = *&d1.Cards

	// p.Name = "Dayan"
	// p.Hand = *&d2.Cards

	// fmt.Println("JOGADOR")
	// fmt.Println(dealer.Name)
	// fmt.Println(dealer.Hand)

	// fmt.Println("JOGADOR")
	// fmt.Println(p.Name)
	// fmt.Println(p.Hand)

	// player1Hand, _ := deck.New(deck.Empty)
	// player2Hand, _ := deck.New(deck.Empty)

	// // deck.Debugf(false, "INIT")
	// card1 := deck.NewCard(deck.TWO, deck.SPADE)ls
	// deck.Debugf(true, "DECK_>\n%s", d)

	// fmt.Println(len(d.Cards))
	// fmt.Print("player1Hand->\n")
	// fmt.Print(player1Hand)

	// fmt.Print("player2Hand->\n")
	// fmt.Print(player2Hand)
	// fmt.Print()

	// fmt.Print(deck.TWO)
	// fmt.Printf("type -> %T\n", d.Cards)

	// for _, c := range black_jack.Cards {
	// 	fmt.Printf("type -> %T\n", c)
	// }

}

func (d *Dealer) PaymentOdds(bet float64) float64 {
	return bet * d.odds
}

type OptionsBlackJack struct {
	Odds       float64
	DeckNumber int
}

func New(optionsBlackJack OptionsBlackJack) *BlackJack {
	options := OptionsBlackJack{
		Odds:       OddsDefault,
		DeckNumber: DeckNumberDefault,
	}

	if optionsBlackJack.DeckNumber > 1 {
		options.DeckNumber = optionsBlackJack.DeckNumber
	}

	if optionsBlackJack.Odds != options.Odds && optionsBlackJack.Odds > 0 {
		options.Odds = optionsBlackJack.Odds
	}
	d, _ := deck.New()

	black_jack := BlackJack{
		DeckNumber: options.DeckNumber,
		Deck:       *d,
		Odds:       options.Odds,
		Dealer:     Dealer{odds: options.Odds},
	}
	return &black_jack
}

func (d *Dealer) CanHit() bool {
	if d.Point <= 16 {
		return true
	}

	return false
}

//REFERENCIA
// type Person struct {
// 	firstName string
// 	lastName  string
// }
// func printIfPerson(object interface{}) {
// 	person, ok := object.(Person)
// 	if ok {
// 		fmt.Printf("Hello %s!\n", person.firstName)
// 	}
// }
