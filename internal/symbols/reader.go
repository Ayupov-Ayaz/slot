package symbols

import (
	"Ayupov-Ayaz/slot/internal/reader"
	"embed"
	"fmt"
)

//go:embed symbols.txt
var symbols embed.FS

func parseReels(data [][]string) ([]Symbols, error) {
	return nil, nil
}

// ReadReels - read symbols from file
func ReadReels() (*Reels, error) {
	// обрати внимание, что в файле symbols.txt символы разделены через \t
	// и что в конце каждой строки есть \n
	// символ -1 нужен только для выравнивания таблицы
	data, err := reader.Read(symbols, "symbols.txt")
	if err != nil {
		return nil, fmt.Errorf("reader.Read(): %w", err)
	}

	reelSymbols, err := parseReels(data)
	if err != nil {
		return nil, fmt.Errorf("parseReels(): %w", err)
	}

	return NewReels(reelSymbols), nil
}
