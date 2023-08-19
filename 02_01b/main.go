package main

import (
	"flag"
	"log"
)

// Goroutines and channels are the Go concurrency mechanism.

// The TASK
// Given a list of messages and a number N implement a function that outputs
// the same message N times concurrently.

var messages = []string{
	"Hello!",
	"How are you?",
	"Are you just going to repeat what I say?",
	"So immature",
	"Stop copying me!",
}

// repeat concurrently prints out the given message n times
func repeat(n int, message string) {
	ch := make(chan struct{})
	for i := 0; i < n; i++ {
		go func(i int) {
			log.Printf("[G%d] %s\n", i, message)
			ch <- struct{}{}
		}(i)
	}
	for i := 0; i < n; i++ {
		<-ch
	}
}

func main() {
	factor := flag.Int64("factor", 0, "The fan-out factor to repeat by")
	flag.Parse()
	for _, m := range messages {
		log.Println(m)
		repeat(int(*factor), m)
	}
}
