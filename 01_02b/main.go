package main

import (
	"log"
	"strings"
	"time"
)

// The Task
// Implement a function that takes in string and outputs a prolonged version of each word

const delay = 700 * time.Millisecond

// print outputs a message and then sleeps for a pre-determined amount
func print(msg string) {
	log.Println(msg)
	time.Sleep(delay)
}

// slowDown takes the given string and repeats its characters
// according to their index in the string.
func slowDown(msg string) {
	words := strings.Split(msg, " ")
	for _, w := range words {
		var pw []string
		for i, c := range w {
			rb := strings.Repeat(string(c), i+1)
			pw = append(pw, rb)
		}
		print(strings.Join(pw, ""))
	}
}

func main() {
	msg := "Time to learn about Go strings!"
	slowDown(msg)
}