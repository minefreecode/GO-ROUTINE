package main

import (
	"fmt"
	"time"
)

// Функция принимающая как параметр Канал запускается асинхронно
func doWork(done <-chan bool) {
	//Бесконечный цикл считывания из канала
	for {
		select {
		//Извлекаем из канала и смотрим есть ли там что-либо
		case <-done:
			return
		default:
			//Информация о том что работа продолжается
			fmt.Println("DOING WORK")
		}
	}
}

func main() {
	//Настоящее время
	t0 := time.Now()
	//Создаём канал для связи с асинхронной функцией
	done := make(chan bool)

	//Запускаем функцию асинхронно
	go doWork(done)

	//Выдерживаем время
	time.Sleep(time.Second * 3)

	//Закрываем канал
	close(done)
	//Печатаем содержимое канала
	fmt.Println(<-done)

	//Печатаем информацию о завершении канала
	printComplete(t0)
}

func printComplete(t time.Time) {
	green := "\033[32m"
	reset := "\033[0m"

	// Print the text in green
	fmt.Printf("\n%scompleted in: %v%s", green, time.Since(t), reset)
}
