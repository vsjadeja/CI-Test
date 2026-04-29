package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}

	return a / b, nil
}

func modulus(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot modulus by zero")
	}

	return a % b, nil
}
