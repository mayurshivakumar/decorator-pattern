package main

import (
	"fmt"
	"time"
)

func decorator(f func(s string)) func(s string) {
	return func(s string) {
		start := time.Now()
		f(s)
		elapsed := time.Since(start)
		fmt.Println("job took: ", elapsed)
	}
}

func quickJob(s string) {
	fmt.Println(s)
}

func longRunningJob(s string) {
	time.Sleep(time.Second * 2)
	fmt.Println(s)
}

func main() {
	decorator(quickJob)("finished quick job")
	decorator(longRunningJob)("finished long job")
}
