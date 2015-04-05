package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string, c chan string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		c <- fmt.Sprintf("%s %d", msg, i)
		// time.Sleep(time.Second)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)

	}
}

func main() {
	c := make(chan string)
	go boring("Hello", c)
	fmt.Println("I am listening")
	for i := 0; i < 5; i++ {
		fmt.Printf("You say %q\n", <-c)
	}
	fmt.Println(" You are boring. Ditching ya !")
}
