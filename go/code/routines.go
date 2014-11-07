package main

import "fmt"
import "time"

func ping(c chan string) {
	for i := 0; i < 10; i++ {
		c <- fmt.Sprintf("hello %v", i)
	}
}

func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string)
	go ping(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
