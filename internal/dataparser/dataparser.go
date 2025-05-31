package dataparser

// DataParser определяет интерфейс для парсинга данных и получения информации
type DataParser interface {
	// Parse парсит строку с данными и заполняет структуру
	Parse(datastring string) error

	// ActionInfo возвращает строку с информацией о данных
	ActionInfo() (string, error)
}

// Info принимает объект, реализующий интерфейс DataParser,
// и возвращает информацию о данных в виде строки
func Info(dp DataParser) (string, error) {
	return dp.ActionInfo()
}
