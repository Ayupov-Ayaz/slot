package main

const spinsCount = 20_000_00

func run() error {
	panic("implement me")
}

func main() {
	// не нужно использовать многопоточность
	// достаточно если программа будет работать последовательно
	// чтение файлов происходит только один раз
	// в цикле for должен быть вызов только функции Spin
	if err := run(); err != nil {
		panic(err)
	}
}
