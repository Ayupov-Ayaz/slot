package paytable

import (
	"Ayupov-Ayaz/slot/internal/configs/reader"
	"Ayupov-Ayaz/slot/internal/configs/symbols"
	"embed"
	"fmt"
)

//go:embed pay_table.txt
var payTable embed.FS

func parsePayouts(data [][]string) (map[symbols.Symbol]Payout, error) {
	// todo: implement me
	return nil, nil
}

func ReadPayTable() (*PayTable, error) {
	data, err := reader.Read(payTable, "pay_table.txt")
	if err != nil {
		return nil, fmt.Errorf("reader.Read(): %w", err)
	}

	payouts, err := parsePayouts(data)
	if err != nil {
		return nil, fmt.Errorf("parsePayouts(): %w", err)
	}

	return NewPayTable(payouts), nil
}
