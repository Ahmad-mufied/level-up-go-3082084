package main

import (
	"fmt"
	"log"
)

// The TASK
// Given a defined list of resources
// implement a function tha simulates the
// concurrent allocation of resources to consumer goroutines.

// A signal channel, var signal chan struct{}, is a channel whose purpose is to
// synchronize goroutines.

// the number of attendees we need to serve lunch to
const consumerCount = 500

// foodCourses represents the types of resources to pass to the consumers
var foodCourses = []string{
	"Caprese Salad",
	"Spaghetti Carbonara",
	"Vanilla Panna Cotta",
}

// takeLunch is the consumer function for the lunch simulation
// Change the signature of this function as required
func takeLunch(name string, in []chan string, done chan<- struct{}) {
	for _, ch := range in {
		log.Printf("%s eats %s\n", name, <-ch)
	}
	done <- struct{}{}
}

// serveLunch is the producer function for the lunch simulation.
// Change the signature of this function as required
func serveLunch(course string, out chan<- string, done <-chan struct{}) {
	for {
		select {
		case <-done:
			return
		default:
			out <- course
		}
	}
}

func main() {
	log.Printf("Welcome to the conference lunch! Serving %d attendees.\n",
		consumerCount)

	// create a slice of channels to pass to the consumers
	var courses []chan string
	doneEating := make(chan struct{})
	doneServing := make(chan struct{})

	// start the producers
	for _, c := range foodCourses {
		ch := make(chan string)
		courses = append(courses, ch)
		go serveLunch(c, ch, doneServing)
	}

	// start the consumers
	for i := 0; i < consumerCount; i++ {
		name := fmt.Sprintf("Attendee %d", i)
		go takeLunch(name, courses, doneEating)
	}

	// wait for all the consumers to finish
	for i := 0; i < consumerCount; i++ {
		<-doneEating
	}

	// close all channels
	for _, ch := range courses {
		close(ch)
	}

	close(doneServing)
	close(doneEating)

	log.Println("All done!")

}
