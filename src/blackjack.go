package blackjack

import (
	"fmt"

	"github.com/adamclerk/deck"
	"github.com/bxcodec/faker"
)

const (
	OddsDefault       = float64(3.0 / 2.0)
	DeckNumberDefault = 1
	DealerCanHit      = 16
	BlackJackPoints   = 21
	Increments        = 1
	Decrements        = -1
	Neutral           = 0
	MAX_PLAYERS       = 7
)

var TablePointCards = map[int]int{
	int(deck.ACE):   0,
	int(deck.TWO):   2,
	int(deck.THREE): 3,
	int(deck.FOUR):  4,
	int(deck.FIVE):  5,
	int(deck.SIX):   6,
	int(deck.SEVEN): 7,
	int(deck.EIGHT): 8,
	int(deck.NINE):  9,
	int(deck.TEN):   10,
	int(deck.JACK):  10,
	int(deck.QUEEN): 10,
	int(deck.KING):  10,
}

type player interface {
	HandCount() int
}

func HandCount(p player) int {
	return p.HandCount()
}

func (d *Dealer) HandCount() int {
	return d.GetHandPoint(&d.Hand)
}

func (p *Player) HandCount() int {
	return p.GetHandPoint(&p.Hand)
}

type Player struct {
	Id        string      `faker:"uuid_digit"`
	Name      string      `faker:"first_name"`
	Hand      []deck.Card `faker:"-"`
	CountCard int         `faker:"-"`
	BlackJack *BlackJack  `faker:"-"`
}

type Dealer struct {
	Player
	odds  float64 `faker:"-"`
	Point int     `faker:"-"`
}

type BlackJack struct {
	DeckNumber  int
	Odds        float64
	Deck        deck.Deck
	Cards       []deck.Card
	Players     []Player
	HashPlayers map[string]*Player
	Dealer      Dealer
	MaxPlayers  int
	Surrender   bool
}

func Play() {
	black_jack := New(OptionsBlackJack{})

	player := Player{}
	err := faker.FakeData(&player)
	if err != nil {
		fmt.Println(err)
	}

	player.SetDownTable(black_jack)
	black_jack.showPlayer()

	for i := 0; i < 2; i++ {
		for _, player := range black_jack.HashPlayers {
			// fmt.Println("Hand")
			black_jack.Dealer.GiveCards(player)
			// fmt.Println(player.Hand)
		}
		// fmt.Println("Delaer Hand")
		// fmt.Println(black_jack.Dealer.Player)
		// fmt.Scan()
		black_jack.Dealer.GiveCards(&black_jack.Dealer.Player)
		// fmt.Println(black_jack.Dealer.Player)
	}

	// SHOW HANDS
	fmt.Println("Delaer Hand")
	fmt.Println(black_jack.Dealer.Player.Hand[0])
	// fmt.Println(black_jack.Dealer.Player.Hand[1])

	for _, player := range black_jack.HashPlayers {
		fmt.Println("player.Hand")
		fmt.Println(player.Hand)
	}

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

	dealer := Player{}
	faker.FakeData(&dealer)

	black_jack := BlackJack{
		DeckNumber:  options.DeckNumber,
		Deck:        *d,
		Odds:        options.Odds,
		Dealer:      Dealer{Player: dealer, odds: options.Odds},
		MaxPlayers:  MAX_PLAYERS,
		HashPlayers: make(map[string]*Player),
	}

	black_jack.Dealer.BlackJack = &black_jack
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
	return HandCount(p)
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

func (d *Dealer) CanSplit(p *Player) bool {
	card1 := int(p.Hand[0].Face())
	card2 := int(p.Hand[1].Face())

	return (len(p.Hand) == 2) && (TablePointCards[card1] == TablePointCards[card2])
}

func (b *BlackJack) AddPlayer(p *Player) {
	p.BlackJack = b
	b.HashPlayers[p.Id] = p
	// fmt.Printf("Mesa tal %d/%d jogadores\n", len(b.HashPlayers), b.MaxPlayers)
}

func (b *BlackJack) RemovePlayer(p *Player) {
	_, exist := b.HashPlayers[p.Id]
	if exist {
		delete(b.HashPlayers, p.Id)
	}
	_, exist = b.HashPlayers[p.Id]

	// fmt.Println()
	// fmt.Printf("Player [%s] give up\n", p.Name)
	// fmt.Printf("Mesa tal %d/%d jogadores\n", len(b.HashPlayers), b.MaxPlayers)
}

func (p *Player) SetDownTable(b *BlackJack) {
	if len(b.HashPlayers) >= b.MaxPlayers {
		return
	}

	b.AddPlayer(p)
}

func (p *Player) GiveUp() {
	p.BlackJack.RemovePlayer(p)
}

func (b *BlackJack) showPlayer() {
	fmt.Printf("___________Players_____________\n")
	for _, v := range b.HashPlayers {
		fmt.Printf("[%s]-> %s\n", v.Id, v.Name)
	}
	fmt.Printf("___________Dealer_____________\n")
	fmt.Printf("[%s]-> %s\n", b.Dealer.Id, b.Dealer.Name)

}

func (d *Dealer) GiveCards(p *Player) {
	card := Pop(&d.BlackJack.Deck)
	p.Hand = append(p.Hand, card)
}

func Pop(d *deck.Deck) deck.Card {
	// fmt.Println(d.Cards)
	last := len(d.Cards) - 1

	card := d.Cards[last]
	d.Cards = d.Cards[:last]

	return card
}

// func (d *Dealer) DrawCard(deck *deck.Card)  {

// }

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
