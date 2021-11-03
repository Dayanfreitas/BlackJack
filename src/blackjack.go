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
	Cards []deck.Card
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

func PaymentOdds(odds float64, bet float64) float64 {
	return bet * odds
}
