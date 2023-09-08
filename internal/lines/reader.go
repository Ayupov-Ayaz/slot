package lines

import (
	"Ayupov-Ayaz/slot/internal/reader"
	"embed"
	"fmt"
)

//go:embed lines.txt
var lines embed.FS

func parseLines(data [][]string) (Lines, error) {
	return nil, nil
}

func ReadLines() (Lines, error) {
	data, err := reader.Read(lines, "lines.txt")
	if err != nil {
		return nil, fmt.Errorf("reader.Read(): %w", err)
	}

	slotLines, err := parseLines(data)
	if err != nil {
		return nil, fmt.Errorf("parseLines(): %w", err)
	}

	return slotLines, nil
}
