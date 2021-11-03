package blackjack

import (
	"fmt"

	"github.com/adamclerk/deck"
)

type Hand struct {
}

type Player struct {
	Id   int
	Name string
	Hand []deck.Card
}

type Dealer struct {
	Player
	odds float64
}

type BlackJack struct {
	// numberPlayers int
	DeckNumber int
	Odds       float64
	Cards      []deck.Card
	Dealer     Dealer
}

func Hello() {
	d, _ := deck.New()

	black_jack := BlackJack{Cards: d.Cards}

	fmt.Printf("%v", black_jack)

	p := Player{}
	dealer := Dealer{}

	fmt.Println()
	dealer.Name = "Dealer"
	dealer.Hand = *&d.Cards
	// dealer.odds = float64() 3/5

	p.Name = "Dayan"
	p.Hand = *&d.Cards

	// fmt.Printf("player -> %T\n", p)
	// fmt.Printf("dealer -> %T\n", dealer)

	fmt.Println("JOGADOR")
	fmt.Println(dealer.Name)
	fmt.Println(dealer.Hand)

	fmt.Println("JOGADOR")
	fmt.Println(p.Name)
	fmt.Println(p.Hand)

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

func New(op OptionsBlackJack) *BlackJack {
	// Contrui black jack
	oddsDefault := float64(3.0 / 2.0)
	black_jack := &BlackJack{}
	black_jack.Odds = oddsDefault

	if op.Odds != oddsDefault && op.Odds > 0 {
		black_jack.Odds = op.Odds
	}

	black_jack.Dealer = Dealer{odds: black_jack.Odds}

	black_jack.DeckNumber = 1
	if op.DeckNumber != 0 {
		black_jack.DeckNumber = op.DeckNumber
	}

	return black_jack
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
