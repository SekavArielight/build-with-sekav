package main

import (
	"fmt"
	"strings"
)

func Clone() {
	c := "aaa"
	clone := strings.Clone(c)
	fmt.Println(c != clone)

	fmt.Print("Hello, World!")
}