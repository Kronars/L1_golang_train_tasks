package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Запись в канал
func Writer(cancel chan bool) chan int {
	main_thread := make(chan int)

	go func() {
		i := 0
		defer close(main_thread)
		for {
			i++
			select {
			case main_thread <- i:
			case <-cancel:
				return
			}
		}
	}()

	return main_thread
}

// Чтение канала
func Worker(inp chan int, work func(v int)) {
	for v := range inp {
		work(v)
	}
}

// Типа полезная работа
func PretendHelpfull(value int) {
	time.Sleep(500 * time.Millisecond)
	fmt.Print(value, ", ")
}

func main() {
	sigs := make(chan os.Signal)  // Канал с сигналами операционной системы
	cancel := make(chan bool)     // Канал с уведомлением о отмене по Ctrl+C
	main_thread := Writer(cancel) // Главный поток

	const amount = 4

	// Запуск воркеров
	for i := 0; i < amount; i++ {
		go Worker(main_thread, PretendHelpfull)
	}

	// Канал оиждания прерывания
	done := make(chan bool)
	// Подписка на уведомление о прерывании
	signal.Notify(sigs, syscall.SIGINT)

	go func() {
		<-sigs         // Получение прерывания
		cancel <- true // Завершение горутин воркеров
		done <- true   // Завершение этой горутины
	}()

	<-done // Ожидание команды завершения
}
