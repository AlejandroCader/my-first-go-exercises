package main

import (
	"fmt"
	"time"
)

func miGoroutine2(c chan int) {
	for i := 0; i < 5; i++ {
		c <- i
		time.Sleep(time.Millisecond * 5000)
	}
	close(c)
}

func main() {
	c := make(chan int)

	go miGoroutine2(c)
	for val := range c {
		fmt.Println("Valor recibido de la GoRoutine:", val)
	}
	fmt.Println("Programa principal: Listo!")
}
