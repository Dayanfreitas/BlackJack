package blackjack

import (
	"fmt"

	"github.com/adamclerk/deck"
	"github.com/bxcodec/faker/v3"
)

const (
	OddsDefault       = float64(3.0 / 2.0)
	DeckNumberDefault = 1
	DealerCanHit      = 16
	BlackJackPoints   = 21
)

// const pointCartOfBalck
// card values ​​in black ja
type Play interface {
	teste() string
	// GetHandPoint() int
}

func teste(p Play) string {
	return p.teste()
}

func (d *Dealer) teste() string {
	return "oi"
}

func (p *Player) teste() string {
	return "oi player"
}

// func CalcPoint(p player) int {
// 	return p.GetHandPoint()
// }

type Player struct {
	Id   string      `faker:"uuid_digit"`
	Name string      `faker:"first_name"`
	Hand []deck.Card `faker:"-"`
}

type Dealer struct {
	Player
	odds  float64 `faker:"-"`
	Point int     `faker:"-"`
}

type BlackJack struct {
	// numberPlayers int
	DeckNumber int
	Odds       float64
	Deck       deck.Deck
	Cards      []deck.Card
	Players    []Player
	Dealer     Dealer
}

func Hello() {

	// d1, _ := deck.New()
	// d2, _ := deck.New()

	// fmt.Print(d2)

	player := Player{}
	err := faker.FakeData(&player)
	if err != nil {
		fmt.Println(err)
	}

	dealer := Dealer{}
	err = faker.FakeData(&dealer)
	// err := faker.FakeData(&dealer.Player)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("player")
	fmt.Println(player)

	fmt.Println("Delaer")
	fmt.Println(dealer)

	// black_jack := BlackJack{Deck: *d1}
	black_jack := New(OptionsBlackJack{})
	// black_jack.Dealer = Dealer{}
	black_jack.Players = append(black_jack.Players, Player{Id: "2", Name: "Dayan"})

	// fmt.Println(black_jack.Dealer.teste())
	// fmt.Println(black_jack.Players[0].teste())
	// fmt.Printf("%s\n", teste(&black_jack.Dealer))
	// fmt.Printf("%s\n", teste(&black_jack.Players[0]))
	// fmt.Println(black_jack.Dealer.odds)
	// fmt.Println(teste(black_jack.Players[0]))
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
	return d.GetHandPoint(&d.Hand) <= DealerCanHit
}

func (d *Dealer) IsBlackjack(hand *[]deck.Card) bool {
	return d.GetHandPoint(hand) == BlackJackPoints
}

func (d *Dealer) IsHandBurst(hand *[]deck.Card) bool {
	return d.GetHandPoint(hand) > BlackJackPoints
}

func (d *Dealer) GetHandPoint(hand *[]deck.Card) int {

	point := 0
	for _, v := range *hand {

		if v.Face() == 0 {
			point += 11
		}

		if v.Face() >= 10 {
			point += 10
		}
	}

	return point

}

func (p *Player) GetHandPoint(hand *[]deck.Card) int {

	point := 0
	for _, v := range *hand {

		if v.Face() == 0 {
			point += 11
		}

		if v.Face() >= 10 {
			point += 10
		}
	}

	// ace convenience
	for _, v := range *hand {
		// ace
		if point > BlackJackPoints {
			if v.Face() == 0 {
				point -= 10
			}
		}
	}

	return point

}

// ace convenience

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
