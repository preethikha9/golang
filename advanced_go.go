package main

import (
	"fmt"
	"time"
)

func go_thread() {
	go thread()
	fmt.Println("main")
	time.Sleep(1 * time.Second)

}
func thread() {
	fmt.Println("go thread")
}

func channel() {
	ch := make(chan int)
	go send(ch)
	go receive(ch)
	time.Sleep(1 * time.Second)
}

func send(ch chan int) {
	ch <- 1
}

func receive(ch chan int) {
	v := <-ch
	fmt.Println(v)
}

type error struct {
	message string
}

func (e *error) error_fun() string {
	return e.message
}

func error_msg() {
	err := &error{
		message: "encounter error",
	}

	fmt.Println("error", err.error_fun())

}
