package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBadBank(t *testing.T) {
	transactions := []Transaction{
		{
			From: "Chris",
			To:   "Riya",
			Sum:  100,
		},
		{
			From: "Adil",
			To:   "Chris",
			Sum:  25,
		},
	}
	assert.EqualValues(t, 100, BalanceFor(transactions, "Riya"))
	assert.EqualValues(t, -75, BalanceFor(transactions, "Chris"))
	assert.EqualValues(t, -25, BalanceFor(transactions, "Adil"))
}
