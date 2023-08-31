package calculator

import (
	"Ayupov-Ayaz/slot/internal/lines"
	"Ayupov-Ayaz/slot/internal/paytable"
	"Ayupov-Ayaz/slot/internal/symbols"
	"Ayupov-Ayaz/slot/internal/win"
)

type Calculator struct {
	lines    *lines.Lines
	payTable *paytable.PayTable
}

func (c *Calculator) Calculate(spinSymbols symbols.Symbols) ([]win.Win, error) {
	panic("implement me")
}
