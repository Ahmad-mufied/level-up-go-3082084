package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// Use pointers and errors to handle edge cases

// The TASK
// Given a simple calculator implement edge case and input verification that ouputs
// error messages and gracefully shuts down the program.

// operators is the map of legal operators and their functions
var operators = map[string]func(x, y float64) float64{
	"+": func(x, y float64) float64 { return x + y },
	"-": func(x, y float64) float64 { return x - y },
	"*": func(x, y float64) float64 { return x * y },
	"/": func(x, y float64) float64 { return x / y },
}

// parseOperand parses a string to a float64
func parseOperand(op string) (*float64, error) {
	parsedOp, err := strconv.ParseFloat(op, 64)
	if err != nil {
		return nil, fmt.Errorf("Cannot parse: %v", err)
	}
	return &parsedOp, nil
}

// calculate returns the result of a 2 operand mathematical expression
func calculate(expr string) (*float64, error) {
	ops := strings.Fields(expr)
	nops := len(ops)
	if nops != 3 {
		return nil, fmt.Errorf("cannot calculate: need 3 operands, got %d", nops)
	}
	left, err := parseOperand(ops[0])
	if err != nil {
		return nil, err
	}
	right, err := parseOperand(ops[2])
	if err != nil {
		return nil, err

	}
	f, ok := operators[ops[1]]
	if !ok {
		return nil, fmt.Errorf("cannot calculate: %s is unknown", ops[1])
	}
	result := f(*left, *right)
	return &result, nil
}

func main() {
	expr := flag.String("expr", "",
		"The expression to calculate on, separated by spaces.")
	flag.Parse()
	result, err := calculate(*expr)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	log.Printf("%s = %.2f\n", *expr, *result)
}
