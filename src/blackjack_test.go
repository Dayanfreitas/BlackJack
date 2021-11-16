package blackjack

import (
	"testing"

	"github.com/adamclerk/deck"
	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/assert"
)

const (
	INCREMENTS = 1
	NEUTRO     = 0
	DECREMENT  = -1
)

func TestProcessPointsInHand(t *testing.T) {
	// Player{}
	// fmt.Print(Player{})
	black_jack := New(OptionsBlackJack{})

	assert.Equal(t, 1, 2, "process points hand player")
}

func TestGiveCards(t *testing.T) {
	black_jack := New(OptionsBlackJack{})

	lastCardForBlackJack := black_jack.Deck.Cards[len(black_jack.Deck.Cards)-1]
	player := Player{}
	_ = faker.FakeData(&player)
	player.SetDownTable(black_jack)

	black_jack.Dealer.GiveCards(&player)

	assert.Equal(t, lastCardForBlackJack, player.Hand[0], "The last card of deck blackjack should be the first card of player")
}

func TestSetDownTable(t *testing.T) {
	black_jack := New(OptionsBlackJack{})
	player := Player{}
	_ = faker.FakeData(&player)
	player.SetDownTable(black_jack)
	assert.Equal(t, black_jack.HashPlayers[player.Id].Id, player.Id, "Player set down")
}

func TestCanotSetDownTable(t *testing.T) {
	black_jack := New(OptionsBlackJack{})

	for i := 0; i < black_jack.MaxPlayers+1; i++ {
		player := Player{}
		_ = faker.FakeData(&player)
		player.SetDownTable(black_jack)
	}

	assert.Equal(t, len(black_jack.HashPlayers), 7, "Can't sit at the table")
}

func TestGiveUpTable(t *testing.T) {
	black_jack := New(OptionsBlackJack{})
	player := Player{}
	_ = faker.FakeData(&player)
	player.SetDownTable(black_jack)
	player.GiveUp()
	assert.Equal(t, len(black_jack.HashPlayers), 0, "Player Give Up")
}

func TestCountCard(t *testing.T) {
	d, _ := deck.New()
	player := Player{}
	for _, v := range d.Cards {
		player.CountCards(&v)
	}

	assert.Equal(t, NEUTRO, player.CountCard, "Count at the end must be zero")
}

func TestCountCardIncrement(t *testing.T) {
	player := Player{}
	tests := []deck.Card{
		deck.NewCard(deck.TWO, deck.CLUB),
		deck.NewCard(deck.THREE, deck.CLUB),
		deck.NewCard(deck.FOUR, deck.CLUB),
		deck.NewCard(deck.FIVE, deck.CLUB),
		deck.NewCard(deck.SIX, deck.CLUB),
	}

	for _, card := range tests {
		assert.Equal(t, INCREMENTS, player.IncrementsDecrementsOrNeutral(&card), "Increment count card ", card)
	}
}

func TestCountCardNeutro(t *testing.T) {
	player := Player{}
	tests := []deck.Card{
		deck.NewCard(deck.SEVEN, deck.CLUB),
		deck.NewCard(deck.EIGHT, deck.CLUB),
		deck.NewCard(deck.NINE, deck.CLUB),
	}

	for _, card := range tests {
		assert.Equal(t, Neutral, player.IncrementsDecrementsOrNeutral(&card), "Neutro count card ", card)
	}
}

func TestCountCardDecrement(t *testing.T) {
	player := Player{}

	tests := []deck.Card{
		deck.NewCard(deck.TEN, deck.CLUB),
		deck.NewCard(deck.JACK, deck.CLUB),
		deck.NewCard(deck.QUEEN, deck.CLUB),
		deck.NewCard(deck.KING, deck.CLUB),
		deck.NewCard(deck.ACE, deck.CLUB),
	}

	for _, card := range tests {
		assert.Equal(t, DECREMENT, player.IncrementsDecrementsOrNeutral(&card), "Decrement count card ", card)
	}
}

func TestCanSplit(t *testing.T) {
	black_jack := New(OptionsBlackJack{})
	d, _ := deck.New(
		deck.WithCards(
			deck.NewCard(deck.QUEEN, deck.DIAMOND),
			deck.NewCard(deck.QUEEN, deck.SPADE),
		), deck.Unshuffled)

	p := Player{Hand: d.Cards}

	assert.Equal(t, true, black_jack.Dealer.CanSplit(&p), "Can split")
}

func TestCanSplitCase2(t *testing.T) {
	black_jack := New(OptionsBlackJack{})
	d, _ := deck.New(
		deck.WithCards(
			deck.NewCard(deck.TEN, deck.DIAMOND),
			deck.NewCard(deck.QUEEN, deck.SPADE),
		), deck.Unshuffled)

	p := Player{Hand: d.Cards}
	// black_jack.Players = append(black_jack.Players, )

	assert.Equal(t, true, black_jack.Dealer.CanSplit(&p), "Can split 10 and dama")
}

func TestCanotSplit(t *testing.T) {
	black_jack := New(OptionsBlackJack{})
	d, _ := deck.New(
		deck.WithCards(
			deck.NewCard(deck.ACE, deck.DIAMOND),
			deck.NewCard(deck.QUEEN, deck.SPADE),
		), deck.Unshuffled)

	p := Player{Hand: d.Cards}
	// black_jack.Players = append(black_jack.Players, )

	assert.Equal(t, false, black_jack.Dealer.CanSplit(&p), "Canot split")
}

func TestBlackJack(t *testing.T) {
	black_jack := New(OptionsBlackJack{})
	d, _ := deck.New(deck.WithCards(deck.NewCard(deck.ACE, deck.CLUB), deck.NewCard(deck.JACK, deck.DIAMOND)), deck.Unshuffled)
	assert.Equal(t, black_jack.Dealer.IsBlackjack(&d.Cards), true, "Is BlackJack")
}

func TestBlackJackHowTwoCardsOfTeenAndOneCardOfAce(t *testing.T) {
	black_jack := New(OptionsBlackJack{})
	d, _ := deck.New(
		deck.WithCards(
			deck.NewCard(deck.ACE, deck.CLUB),
			deck.NewCard(deck.JACK, deck.DIAMOND),
			deck.NewCard(deck.QUEEN, deck.SPADE),
		), deck.Unshuffled)

	black_jack.Players = append(black_jack.Players, Player{Hand: d.Cards})

	assert.Equal(t, 21, black_jack.Dealer.CountPoints(&black_jack.Players[0]), "Two Cards Of Teen And One Card Of Ace")
}

func TestHandIsBurst(t *testing.T) {
	black_jack := New(OptionsBlackJack{})
	d, _ := deck.New(deck.WithCards(deck.NewCard(deck.ACE, deck.CLUB), deck.NewCard(deck.JACK, deck.DIAMOND), deck.NewCard(deck.JACK, deck.DIAMOND)), deck.Unshuffled)
	assert.Equal(t, black_jack.Dealer.IsHandBurst(&d.Cards), true, "Hand is Burst")
}

func TestRandomHand(t *testing.T) {
	type test struct {
		Deck   deck.Deck
		Answer int
	}

	black_jack := New(OptionsBlackJack{})
	d1, _ := deck.New(deck.WithCards(deck.NewCard(deck.ACE, deck.CLUB), deck.NewCard(deck.JACK, deck.DIAMOND)), deck.Unshuffled)
	d2, _ := deck.New(deck.WithCards(deck.NewCard(deck.ACE, deck.CLUB), deck.NewCard(deck.JACK, deck.DIAMOND)), deck.Unshuffled)
	d3, _ := deck.New(deck.WithCards(deck.NewCard(deck.ACE, deck.CLUB), deck.NewCard(deck.JACK, deck.DIAMOND)), deck.Unshuffled)

	tests := []test{
		{Deck: *d1, Answer: 21},
		{Deck: *d2, Answer: 21},
		{Deck: *d3, Answer: 21},
	}

	for _, v := range tests {
		assert.Equal(t, black_jack.Dealer.GetHandPoint(&v.Deck.Cards), v.Answer, "Random Hand")
	}
}

func TestAmountOfDeckDefault(t *testing.T) {
	black_jack := New(OptionsBlackJack{})
	assert.Equal(t, black_jack.Deck.NumberOfDecks, 1, "Number of deck default")
}

func TestAmountOfCart(t *testing.T) {
	black_jack := New(OptionsBlackJack{})
	assert.Equal(t, len(black_jack.Deck.Cards), 52, "Number of default cards 52")
}

func TestAmountOfDeckOne(t *testing.T) {
	black_jack := New(OptionsBlackJack{DeckNumber: 1})
	assert.Equal(t, black_jack.Deck.NumberOfDecks, 1, "BlackJack com apenas um baralho")
}

func TestAmountOfDeckTwo(t *testing.T) {
	black_jack := New(OptionsBlackJack{DeckNumber: 2})
	assert.Equal(t, black_jack.DeckNumber, 2, "BlackJack com dos baralhos")
}

func TestPaymentOddsDefault(t *testing.T) {
	black_jack := New(OptionsBlackJack{})
	assert.Equal(t, black_jack.Odds, float64(3.0/2.0), "Odds payment default")
}

func TestPaymentOddsThreeForTwo(t *testing.T) {
	black_jack := New(OptionsBlackJack{Odds: float64(3.0 / 2.0)})
	betValue := 100.0
	assert.Equal(t, 150.0, black_jack.Dealer.PaymentOdds(betValue), "Pagamento do cassino 3 para 2")
}

func TestPaymentOddsSixForFive(t *testing.T) {
	black_jack := New(OptionsBlackJack{Odds: float64(6.0 / 5.0)})
	betValue := 100.0
	assert.Equal(t, 120.0, black_jack.Dealer.PaymentOdds(betValue), "Pagamento do cassino 6 para 5")
}

//Dealer
func TestDealerCanHit(t *testing.T) {
	black_jack := New(OptionsBlackJack{})
	assert.Equal(t, black_jack.Dealer.CanHit(), true, "Dealer can hit")
}

func TestDealerCanotHit(t *testing.T) {
	black_jack := New(OptionsBlackJack{})
	d, _ := deck.New(deck.WithCards(deck.NewCard(deck.ACE, deck.CLUB), deck.NewCard(deck.JACK, deck.DIAMOND)), deck.Unshuffled)
	black_jack.Dealer.Hand = d.Cards
	assert.Equal(t, black_jack.Dealer.CanHit(), false, "Dealer can´t hit")
}
