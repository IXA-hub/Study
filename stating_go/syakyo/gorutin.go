package main

import (
	"fmt"
	"time"
)

func receiver(ch <-chan int) {
	for {
		i := <-ch
		fmt.Println(i)
	}
}

func receive(name string, ch <-chan int) {
	for {
		_, ok := <-ch
		if ok == false {
			//受付終了で分岐
			break
		}
		fmt.Println(name + "is done .")
	}
}

func main() {
	ch := make(chan int, 20)

	go receive("1st goroutine", ch)
	go receive("2nd goroutine", ch)
	go receive("3rd goroutine", ch)

	i := 0
	for i < 100 {
		ch <- i
		i++
	}

	close(ch)

	time.Sleep(3 * time.Second)
}
