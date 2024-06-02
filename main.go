package main

import (
	"errors"
	"fmt"

	"github.com/siddhant-k-code/experiment-stacktrace/stacktrace"
)

func main() {
	fmt.Println("Hello")
	fmt.Println(f())
	fmt.Println(stacktrace.WithStack(f()))
}

func f() error {
	return errors.New("f func error")
}
