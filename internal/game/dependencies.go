package game

import (
	"Ayupov-Ayaz/slot/internal/configs/symbols"
	"Ayupov-Ayaz/slot/internal/game/win"
	"Ayupov-Ayaz/slot/internal/rng"
)

type Calculator interface {
	Calculate(symbols symbols.Reels) ([]win.Win, error)
}

type Generator interface {
	Generate(rng rng.RNG) (symbols.Reels, error)
}
