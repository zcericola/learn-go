package main

import "fmt"

//hello takes in a string arg and returns a string
func hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(hello("world"))

}
