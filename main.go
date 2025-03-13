package main

import (
	"fmt"
	"time"
)

func main() {

	// Получает настоящее время
	t0 := time.Now()

	//Созаёт канал для объения между го рутинами
	channel := make(chan string)
	//Создаётся другой канал для общения между горутинами
	anotherChannel := make(chan string)

	//Запускаем функцию асинхронно и заносим в его канал данные
	go func() {
		channel <- "data" //Запись данных в канал
	}()
	//Запускаем другую функцию асинхронно и заносим туда другие данные
	go func() {
		anotherChannel <- "another-data" //Запись данных в канал
	}()

	//Ожидаем когда в канале будут достигнуты условия
	select {
	// Условие что в канала что-то есть
	case msg1 := <-channel:
		fmt.Println(msg1)
	//Условие что в другом канале что-то есть
	case msg2 := <-anotherChannel:
		fmt.Println(msg2)
	}

	printComplete(t0)
}

func printComplete(t time.Time) {
	// Форматирвоание текста в целеный цвет
	green := "\033[32m"
	reset := "\033[0m"

	//Печать сколько времени запрос выполяется
	fmt.Printf("\n%sзавершено в: %v%s", green, time.Since(t), reset)
}
