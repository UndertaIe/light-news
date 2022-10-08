package main

import "fmt"

type Job func()

func NewJob(r *Rule) Job {
	return func() {
		fmt.Println("job running")
	}
}
