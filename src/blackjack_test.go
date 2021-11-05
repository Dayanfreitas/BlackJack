package blackjack

import (
	"testing"

	"github.com/adamclerk/deck"
	"github.com/stretchr/testify/assert"
)

func TestBlackJack(t *testing.T) {
	black_jack := New(OptionsBlackJack{})
	d, _ := deck.New(deck.WithCards(deck.NewCard(deck.ACE, deck.CLUB), deck.NewCard(deck.JACK, deck.DIAMOND)), deck.Unshuffled)
	assert.Equal(t, black_jack.Dealer.IsBlackjack(&d.Cards), true, "Is BlackJack")
}

func TestBlackJackHowTwoCardsOfTeenAndOneCardOfAce(t *testing.T) {
	// black_jack := New(OptionsBlackJack{})
	// d, _ := deck.New(
	// 	deck.WithCards(
	// 		deck.NewCard(deck.ACE, deck.CLUB),
	// 		deck.NewCard(deck.JACK, deck.DIAMOND),
	// 		deck.NewCard(deck.QUEEN, deck.SPADE),
	// 	), deck.Unshuffled)
	// black_jack.Players[0] = Player{Hand: d.Cards}
	// var p player
	// p = black_jack.Players[0]

	assert.Equal(t, false, true, "Two Cards Of Teen And One Card Of Ace")
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
	// d, _ := deck.New(deck.WithCards(deck.NewCard(deck.ACE, deck.CLUB), deck.NewCard(deck.JACK, deck.DIAMOND), deck.NewCard(deck.JACK, deck.DIAMOND)), deck.Unshuffled)
	// assert.Equal(t, black_jack.Dealer.IsHandBurst(&d.Cards), true, "Hand is Burst")
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
