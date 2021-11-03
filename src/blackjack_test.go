package blackjack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPaymentOddsThreeForTwo(t *testing.T) {
	odds := float64(3.0 / 2.0)
	betValue := 100.0

	assert.Equal(t, 150.0, PaymentOdds(odds, betValue), "Pagamento do cassino 3 para 2")
}

func TestPaymentOddsSixForFive(t *testing.T) {
	odds := (6.0 / 5.0)
	betValue := 100.0

	assert.Equal(t, 120.0, PaymentOdds(odds, betValue), "Pagamento do cassino 6 para 5")
}
