package main

import (
	"fmt"
	"os"
)

func main() {
	TestNew()
}

func TestCommand() {
	fmt.Println(os.Getwd())
	err := os.Chdir("..")
	if err != nil {
		fmt.Println("error")
	} else {
		fmt.Println(os.Getwd())
	}
}

func TestNew() {
	a := float64(0.2)
	b := float64(0.4)
	fmt.Println(a + b)
}
