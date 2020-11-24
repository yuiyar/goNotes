package main

import (
	"fmt"
	"log"
)

func main() {
	z, err := div(3, 1)

	if err != nil {
		switch e := err.(type) {
		case DivError:
			fmt.Println(e, e.x, e.y)
		default:
			fmt.Println(e)
		}
		log.Fatalln(err)
	}

	fmt.Println(z)
}

type DivError struct {
	x, y int
}

func (DivError) Error() string {
	return "Division by zero"
}

func div(x, y int) (int, error) {
	if y == 0 {
		return 0, DivError{x, y}
	}
	return x / y, nil
}
