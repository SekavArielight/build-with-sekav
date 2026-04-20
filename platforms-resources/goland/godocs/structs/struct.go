package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func FirstStruct() {
	p1 := Person{"Jude", 12}
	
	fmt.Println(p1.Age)
}