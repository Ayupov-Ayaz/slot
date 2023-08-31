package result

import (
	"Ayupov-Ayaz/slot/internal/symbols"
	"Ayupov-Ayaz/slot/internal/win"
)

type Round struct {
	seed        int64
	spinSymbols symbols.Symbols
	wins        []win.Win
	totalBet    uint64
	totalWin    uint64
}
