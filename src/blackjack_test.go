package blackjack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPaymentOddsThreeForTwo(t *testing.T) {
	dealer := Dealer{odds: float64(3.0 / 2.0)}
	betValue := 100.0
	assert.Equal(t, 150.0, dealer.PaymentOdds(betValue), "Pagamento do cassino 3 para 2")
}

func TestPaymentOddsSixForFive(t *testing.T) {
	dealer := Dealer{odds: float64(6.0 / 5.0)}
	betValue := 100.0
	assert.Equal(t, 120.0, dealer.PaymentOdds(betValue), "Pagamento do cassino 6 para 5")
}
