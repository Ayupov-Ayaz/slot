# Техническое задание на должность Golang Developer [GameBeat.studio](https://gamebeat.studio/)

Подразделение разработка игр

## Цели
1) Проверить уровень знаний языка программирования Golang
2) Проверить уровень знаний алгоритмов и структур данных
3) Проверить знания и умение применять систему контроля версий Git
4) Показать кандидату задачи, которые он будет решать в рамках работы.  
Чтобы кандидат мог оценить свои силы и понять, сможет ли он справиться с ними.


## Задача
Необходимо реализовать механику слотовой игры

### Описание игры
Игра представляет собой слот с 5 барабанами и 3 рядами (5x3).
В игре есть 10 символов, которые могут выпадать на барабанах.
В игре есть 5 линий выплат, которые определяются комбинацией символов на барабанах.
Игрок получает выигрыш если выпадает комбинация символов,
количество которых соответствует одной из [линий выплат](./internal/lines/lines.txt). 
Величину выигрыша определяет [таблица выплат](./internal/paytable/pay_table.txt).
Символы которые участвуют в игре прописаны в файле [символов](./internal/symbols/symbols.txt).

> Условием для получения выигрыша является подряд идущие символы на одной из линий выплат.
> Например, если на линии выпадает комбинация символов `1 1 1 1 1`,
> то игрок получает выигрыш в размере 500 [если смотреть по таблице выплат](./internal/paytable/pay_table.txt).
> Важно отметить, что в игре присуствует WILD символ, индекс которого равен 0.
> Этот символ заменяет любой другой символ. То есть комбинация `1 1 0 1 1`
является идентичной комбинации `1 1 1 1 1`.

### Механика игры
Игру можно разделить условно на несколько этапов:

1) Игрок делает ставку
> Это классическая slot игра, по этому ставка(TotalBet) будет равна количеству линий

2) Идет генерация барабанов, для игрового поля 5x3
> Необходимо пробежать с 1 по 5 [барабан](./internal/symbols/symbols.txt) используя
> [Генератор случайных чисел](./internal/rng/mercer.go) 
> определиться с выпавшими символами 
> (для этого необходимо получить случайное число от 0 до длины каждого
> барабана и взять первые 3 символа начиная с этого индекса)
>
> Важное замечание:  
> Если выпадет индекс в конце барабана (к примеру 17, когда барабан имеет длину 18 индексов),
> то необходимо взять символы с 17, 18, 1 индекса. То есть зациклить барабан. 

3) Необходимо определить выигрыш игрока
> Для этого необходимо пробежать по всем [линиям выплат](./internal/lines/lines.txt)
> и посчитать количество символов на каждой линии выплат, идущих подряд.
> Если количество символов соответствует какому-то выигрышу из [таблицы выплат](./internal/paytable/pay_table.txt),
> то игрок получает выигрыш в размере, указанном в таблице выплат.
> 

4) Необходимо вернуть игроку результат игры
> В результате игры игрок должен получить:
> - Ставку которую он сделал (TotalBet uint64)
> - Выигрыш, который он получил (TotalWin uint64)
> - Символы, которые выпали на барабанах (Symbols [][]uint16)
> - Seed - это число, которое использовалось для генерации барабанов (Seed int64)
> - Массив выигрышных линий (WinLines []LineWin)
> 
> LineWin - это структура, которая содержит в себе:
> - Номер линии выплаты (Line uint8)
> - Количество символов на линии выплаты (Count uint8)
> - Выигрышный символ (Symbol uint16)
> - Выигрыш (Win uint64)


5) Реализация симулятора
> Что бы проверит, что все отрабатывает корректно, необходимо реализовать симулятор игры.
> Это такая сущность которая делает ставку 20 млн раз и считает сумму выигрышей и проигрышей.
> Каждая ставка должна быть выполнена с новым сидом!
>(достаточно, если в качестве Seed будет использоваться TimeStamp в nano секундах)
> 
> В конце симуляции необходимо вывести в консоль в читабельном виде:
> - Сумму ставок (TotalBet * 20млн)
> - Сумму выигрышей (TotalWin)
> - RTP (Return to player) - это отношение суммы выигрышей к сумме ставок

### Требования к реализации
1) Игра должна быть реализована на языке программирования Golang
2) Игра должна быть реализована в виде пакета, который можно будет подключить к любому проекту
3) Обязательно наличие тестов с покрытием не менее 80%. Для тестов можете использовать любой удобный для вас мокинг фреймворк:
* [gomock](https://github.com/golang/mock)
* [testify/mock](https://github.com/stretchr/testify)
* [golang-mock](https://pkg.go.dev/github.com/stretchr/testify/mock)
* [mockery](https://github.com/vektra/mockery)
* и т.д.
4) Работу необходимо выложить на github и прислать ссылку на репозиторий на почту указанную в пункте 3.
5) Оповестить о готовности задания сотрудника который проводил первичное интервью
6) На тестовое задание дается 1 неделя, может быть продлено, если будет согласовано с сотрудником, который проводил первичное интервью
7) Конфиги игры должны оставаться в *.txt файлах, чтобы иметь возможность динамически их настраивать не влезая в код игры
8) Чтение конфигов происходит только 1 раз, при создании игры
9) Конфиги игры должны быть только для чтения, их нельзя изменять во время игры
10) Результаты игры должны быть только для чтения


### Полезные ссылки
* [go embed](https://pkg.go.dev/embed) - для чтения файла
* [fortune-three](https://avocado.prod.gamebeat.cloud/gamebeat/?game=fortune_three&home=https://gamebeat.studio/&clientType=mobile) 
похожая игра, можно посмотреть визуально как работает
*  [semantic commit](https://gist.github.com/joshbuchea/6f47e86d2510bce28f8e7f42ae84c716)
будет плюсом, если вы будете комментировать свой код в этом стиле
* Какие-то уточняющие вопросы по реализации можно задавать на почту
ayaz.ayupov@releaseband.com (просматриваю в день по 2-3 раза)