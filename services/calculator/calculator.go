package calculator

import (
	"Ayupov-Ayaz/slot/internal/lines"
	"Ayupov-Ayaz/slot/internal/paytable"
	"Ayupov-Ayaz/slot/internal/symbols"
	"Ayupov-Ayaz/slot/internal/win"
)

// WILD - специальный символ, который может заменить любой другой символ
// он не имеет своего выигрыша, но может увеличить выигрыш за счет замены другого символа
const WILD = symbols.Symbol(0)

type Calculator struct {
	lines    *lines.Lines
	payTable *paytable.PayTable
}

func (c *Calculator) Calculate(spinSymbols symbols.Symbols) ([]win.Win, error) {
	panic("implement me")
}
