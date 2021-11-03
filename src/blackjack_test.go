package blackjack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAmountOfDeckDefault(t *testing.T) {
	black_jack := New(OptionsBlackJack{})
	assert.Equal(t, black_jack.DeckNumber, 1, "Number of deck default")
}

func TestAmountOfDeckOne(t *testing.T) {
	odds := float64(3.0 / 2.0)
	black_jack := New(OptionsBlackJack{Odds: odds, DeckNumber: 1})
	assert.Equal(t, black_jack.DeckNumber, 1, "BlackJack com apenas um baralho")
}

func TestAmountOfDeckTwo(t *testing.T) {
	odds := float64(3.0 / 2.0)
	black_jack := New(OptionsBlackJack{Odds: odds, DeckNumber: 2})
	assert.Equal(t, black_jack.DeckNumber, 2, "BlackJack com dos baralhos")
}

func TestPaymentOddsDefault(t *testing.T) {
	black_jack := New(OptionsBlackJack{})
	assert.Equal(t, black_jack.Odds, float64(3.0/2.0), "Odds payment default")
}

func TestPaymentOddsThreeForTwo(t *testing.T) {
	odds := float64(3.0 / 2.0)
	black_jack := New(OptionsBlackJack{Odds: odds})
	betValue := 100.0
	assert.Equal(t, 150.0, black_jack.Dealer.PaymentOdds(betValue), "Pagamento do cassino 3 para 2")
}

func TestPaymentOddsSixForFive(t *testing.T) {
	odds := float64(6.0 / 5.0)
	black_jack := New(OptionsBlackJack{Odds: odds})
	betValue := 100.0
	assert.Equal(t, 120.0, black_jack.Dealer.PaymentOdds(betValue), "Pagamento do cassino 6 para 5")
}
