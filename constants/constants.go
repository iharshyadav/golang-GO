package main

import "fmt"

const age int = 20

func main() {
	const name string = "harsh"

	// group constants

	const (
		host = "localhost"
		port = 3000
	)

	fmt.Println(age,name,host,port)
}