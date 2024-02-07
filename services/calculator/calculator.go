package calculator

import (
	"Ayupov-Ayaz/slot/internal/configs/lines"
	"Ayupov-Ayaz/slot/internal/configs/paytable"
	"Ayupov-Ayaz/slot/internal/configs/symbols"
	"Ayupov-Ayaz/slot/internal/game/win"
)

// WILD - специальный символ, который может заменить любой другой символ
// он не имеет своего выигрыша, но может увеличить выигрыш за счет замены другого символа
const WILD = symbols.Symbol(0)

type Calculator struct {
	lines    lines.Lines
	payTable *paytable.PayTable
}

func NewCalculator(lines lines.Lines, payTable *paytable.PayTable) *Calculator {
	return &Calculator{lines: lines, payTable: payTable}
}

func (c *Calculator) Calculate(spinSymbols symbols.Reels) ([]win.Win, error) {
	// todo: implement me
	return nil, nil
}
