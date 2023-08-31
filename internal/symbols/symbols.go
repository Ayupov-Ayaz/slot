package symbols

import "embed"

//go:embed symbols.txt
var symbols embed.FS

type Symbol = uint16

type Symbols []Symbol
