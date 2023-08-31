package symbols

type Reels struct {
	// read only
	symbols Symbols
}

func NewReels(symbols Symbols) *Reels {
	return &Reels{symbols: symbols}
}

// ReadReels - read symbols from file
func ReadReels() (*Reels, error) {
	// обрати внимание, что в файле symbols.txt символы разделены через \t
	// и что в конце каждой строки есть \n
	// символ -1 нужен только для выравнивания таблицы
	panic("implement me")
}
