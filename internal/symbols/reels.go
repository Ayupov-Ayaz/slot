package symbols

// Reels - коллекция барабанов
// Выступает в качестве обертки над Symbols,
// чтобы ограничить доступ на редактирование барабанов
// Используется при генерации roundSymbols в services/generator/generator.go
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

// todo: добавить функцию которая будет использоваться в services/generator/generator.go
