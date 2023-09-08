package paytable

import "Ayupov-Ayaz/slot/internal/symbols"

type Payout []uint64

type PayTable struct {
	symbolPayouts map[symbols.Symbol]Payout
}

func NewPayTable(symbolPayouts map[symbols.Symbol]Payout) *PayTable {
	return &PayTable{symbolPayouts: symbolPayouts}
}
