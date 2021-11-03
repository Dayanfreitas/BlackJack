package blackjack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAmountOfDeckOne(t *testing.T) {
	black_jack := New(float64(3.0/2.0), 1)

	assert.Equal(t, black_jack.DeckNumber, 1, "BlackJack com apenas um baralho")
}

func TestAmountOfDeckTwo(t *testing.T) {
	black_jack := New(float64(3.0/2.0), 2)

	assert.Equal(t, black_jack.DeckNumber, 2, "BlackJack com dos baralhos")
}

func TestPaymentOddsThreeForTwo(t *testing.T) {
	black_jack := New(float64(3.0 / 2.0))
	// black_jack := BlackJack{Odds: }
	// dealer := Dealer{odds: float64(3.0 / 2.0)}
	// dealer := Dealer{}
	// black_jack.Dealer = dealer

	betValue := 100.0
	assert.Equal(t, 150.0, black_jack.Dealer.PaymentOdds(betValue), "Pagamento do cassino 3 para 2")
}

func TestPaymentOddsSixForFive(t *testing.T) {
	black_jack := New(float64(6.0 / 5.0))
	// black_jack := BlackJack{Odds: float64(6.0 / 5.0)}
	// black_jack.New()
	// dealer := Dealer{odds: float64(6.0 / 5.0)}
	// dealer := Dealer{}
	// black_jack.Dealer = dealer
	betValue := 100.0
	assert.Equal(t, 120.0, black_jack.Dealer.PaymentOdds(betValue), "Pagamento do cassino 6 para 5")
}
