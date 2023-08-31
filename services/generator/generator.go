package generator

import "Ayupov-Ayaz/slot/internal/symbols"

type RNG interface {
	Random(min uint32, max uint32) uint32
}

type Symbols struct {
	reels *symbols.Reels
}

func (s *Symbols) Generate(rng RNG) (roundSymbols symbols.Symbols, err error) {
	panic("implement me")
}
