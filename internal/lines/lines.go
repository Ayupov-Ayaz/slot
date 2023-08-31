package lines

import "embed"

//go:embed lines.txt
var lines embed.FS

type Line struct{}

type Lines []Line

func ReadLines() (Lines, error) {
	panic("implement me")
	// обрати внимание на то, что в файле lines.txt символы разделены через \t
	// и что в конце каждой строки есть \n
	// первая строка относится к линии - 1
	// последняя строка относится к линии - 5
}
