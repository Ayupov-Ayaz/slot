package slot

import (
	"Ayupov-Ayaz/slot/internal/symbols"
	"Ayupov-Ayaz/slot/internal/win"
)

type RNG interface {
	Random(min uint32, max uint32) uint32
}

type Calculator interface {
	Calculate(symbols symbols.Symbols) ([]win.Win, error)
}

type Generator interface {
	Generate(rng RNG) (symbols.Symbols, error)
}

type Slot struct {
	generator  Generator
	calculator Calculator
}

func NewSlot() *Slot {
	return &Slot{}
}

func (s *Slot) Spin(rng RNG) ([]win.Win, error) {
	panic("implement me")
}
