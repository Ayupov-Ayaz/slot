package win

import (
	"Ayupov-Ayaz/slot/internal/symbols"
)

type Win struct {
	amount uint64
	symbol symbols.Symbol
	count  uint8
	line   uint8
}
