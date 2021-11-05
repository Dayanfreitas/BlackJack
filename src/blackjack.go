package blackjack

import (
	"fmt"

	"github.com/adamclerk/deck"
)

const (
	OddsDefault       = float64(3.0 / 2.0)
	DeckNumberDefault = 1
	DealerCanHit      = 16
	BlackJackPoints   = 21
	Increments        = 1
	Decrements        = -1
	Neutral           = 0
)

// const pointCartOfBalck
// card values ​​in black ja
type player interface {
	teste() int
	// GetHandPoint() int
}

func teste(p player) int {
	return p.teste()
}

func (d *Dealer) teste() int {
	fmt.Println("p.Hand")
	fmt.Println(d.Hand)

	return 21
}

func (p *Player) teste() int {
	fmt.Println("p.Hand")
	fmt.Println(p.Hand)

	return 20
}

// func CalcPoint(p player) int {
// 	return p.GetHandPoint()
// }

type Player struct {
	Id        string      `faker:"uuid_digit"`
	Name      string      `faker:"first_name"`
	Hand      []deck.Card `faker:"-"`
	CountCard int         `faker:"-"`
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

	d1, _ := deck.New()
	d2, _ := deck.New()

	// fmt.Print(d2)

	// player := Player{}
	// err := faker.FakeData(&player)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// dealer := Dealer{}
	// err = faker.FakeData(&dealer)
	// // err := faker.FakeData(&dealer.Player)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println("player")
	// fmt.Println(player)

	// fmt.Println("Delaer")
	// fmt.Println(dealer)

	// black_jack := BlackJack{Deck: *d1}
	black_jack := New(OptionsBlackJack{})
	// black_jack.Dealer = Dealer{}
	black_jack.Players = append(black_jack.Players, Player{Id: "2", Name: "Dayan"})

	// fmt.Println(black_jack.Dealer.teste())
	// fmt.Println(black_jack.Players[0].teste())
	// fmt.Printf("%s\n", teste(&black_jack.Dealer))
	// fmt.Printf("%s\n", teste(&black_jack.Players[0]))
	// fmt.Println(black_jack.Dealer.odds)

	black_jack.Players[0].Hand = d1.Cards
	black_jack.Dealer.Hand = d2.Cards

	fmt.Printf("%d\n", black_jack.Dealer.CountPoints(&black_jack.Players[0]))
	fmt.Printf("%d\n", black_jack.Dealer.CountPoints(&black_jack.Dealer))
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
		} else {
			point += v.Face()
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

func (d *Dealer) CountPoints(p player) int {
	return teste(p)
}

func InBetween(value, start, end int) bool {

	if (value >= start) && (value <= end) {
		return true
	}
	return false
}

func (p *Player) IncrementsDecrementsOrNeutral(c *deck.Card) int {
	//<2>----<6> +1
	//<7>----<9>  0
	//<10>---<A> -1

	v := c.Face()

	if InBetween(v, int(deck.TWO), int(deck.SIX)) {
		return Increments
	}

	if InBetween(v, int(deck.TEN), int(deck.KING)) || (v == int(deck.ACE)) {
		return Decrements
	}

	return Neutral
}

func (p *Player) CountCards(c *deck.Card) int {
	p.CountCard += p.IncrementsDecrementsOrNeutral(c)
	return p.CountCard
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
