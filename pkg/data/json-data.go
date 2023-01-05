package data

import "fmt"

// Наш сервис который работает только с JSON форматом
// Нам нужно применить паттерн проектирования Adapter

type JsonDocument struct {
}

func (document JsonDocument) ConvertToXml() string {
	return "<xml></xml>"
}

type JsonDocumentAdapter struct {
	//Инжектим все поля в наш адаптер
	JsonDocument *JsonDocument
}

// Мы создаем реализацию интерфейса не в самом JsonDocument а в его адаптере

func (adapter JsonDocumentAdapter) SendXmlData() {
	adapter.JsonDocument.ConvertToXml()
	fmt.Println("Отправка XML данных в")
}
