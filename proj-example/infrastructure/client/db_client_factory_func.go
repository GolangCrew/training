package client

type ClientFactoryFunc func() (Client, error) //сигнатура функции, которую принимает конструктор репозитория


// функция, которая возвращает настоящий клиент базы
func NewDBClient() (Client, error) {
	return nil, nil
}
