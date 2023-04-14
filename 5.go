package main

import (
	"fmt"
	"time"
)

func main() {
	duration := 3

	ch := Writer(duration)
	go Reader(ch)
	time.Sleep(time.Second * 4)
}

func Writer(duration int) chan int {
	ch := make(chan int)
	count := 0

	alarm := time.After(time.Duration(duration * int(time.Second)))
	go func() {
		defer close(ch)
		for {
			count++
			time.Sleep(time.Millisecond * 250)
			select {
			case ch <- count:
			case t := <-alarm:
				fmt.Println(t)
				return
			}
		}
	}()

	return ch
}

func Reader(inp chan int) {
	for val := range inp {
		fmt.Printf("Считал - %d\n", val)
	}
}
