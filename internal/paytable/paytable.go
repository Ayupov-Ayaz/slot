package paytable

import "embed"

//go:embed pay_table.txt
var paytable embed.FS

type PayTable struct{}

func ReadPayTable() (*PayTable, error) {
	panic("implement me")
	// обрати внимание на то, что в файле pay_table.txt символы разделены через \t
	// и что в конце каждой строки есть \n
	// первая строка относится к таблице выплат символа - 1
	// последняя строка относится к таблице выплат символа - 7
}
