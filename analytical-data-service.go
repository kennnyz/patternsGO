package data_service

// Внешний пакет, который неизменяем и работает только с данными формата XML

import "fmt"

type AnalyticalDataService interface {
	SendXmlData()
}

type XmlDocument struct {
}

func (doc XmlDocument) SendXmlData() {
	fmt.Println("Отправка XML документа")
}
