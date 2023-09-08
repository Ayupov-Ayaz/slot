package symbols

// Reels - коллекция барабанов
// Выступает в качестве обертки над Symbols,
// чтобы ограничить доступ на редактирование барабанов
// Используется при генерации roundSymbols в services/generator/generator.go
type Reels struct {
	// read only
	symbols []Symbols
}

func NewReels(symbols []Symbols) *Reels {
	return &Reels{symbols: symbols}
}

func (r *Reels) GetSymbols(reelIndex int, startIndex int, count int) (Symbols, error) {
	panic("implement me")
}
