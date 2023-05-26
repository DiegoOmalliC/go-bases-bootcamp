package main

import (
	"errors"
	"fmt"
	"time"
)

// MyError is an error implementation that includes a time and message.
type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("%v: %v", e.When, e.What)
}

func oops() error {
	return &MyError{
		time.Date(1989, 3, 15, 22, 30, 0, 0, time.UTC),
		"the file system has gone away",
	}
}

func main() {
	if err := oops(); err != nil {
		fmt.Println(err)
	}
	err := oops()
	var error *MyError

	isMyError := errors.As(err, &error)
	fmt.Println(isMyError)

	isMyError = errors.Is(err, fmt.Errorf("the file system has gone away"))
	fmt.Println(isMyError)

}
