package main

import (
	"fmt"
	"log"
	"os"
)

type customer struct {
	name string
}

func (c *customer) getName() string {
	return c.name
}

func loggingDecorator(f func() string, l *log.Logger) func() string {
	return func() string {
		l.Println("start...")
		name := f()
		defer func() { l.Println("end...") }()
		return name
	}
}

func main() {

	myLogger := log.New(os.Stdout, "### ", 11)
	c := customer{name: "bob"}
	decorated := loggingDecorator(c.getName, myLogger)()
	fmt.Println("log decorated    ", decorated)
}
