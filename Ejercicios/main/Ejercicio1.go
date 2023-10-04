package main

import (
	"fmt"
	"time"
)

func miGoroutine() {
	for i := 0; i < 5; i++ {
		fmt.Println("GoRoutine: ", i)
		time.Sleep(time.Millisecond * 5000)
	}
}

func main() {
	go miGoroutine()

	time.Sleep(time.Second * 20)

	fmt.Println("Programa principal: Listo!")
}
