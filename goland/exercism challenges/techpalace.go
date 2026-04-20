package main

import (
	"fmt"
	"strings"
)

func Techpalace() {

}

func main() {
	c := "aaa"
	clone := strings.Clone(c)
	fmt.Println(c != clone)
}