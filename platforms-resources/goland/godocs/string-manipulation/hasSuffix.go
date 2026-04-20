package main

import (
	"fmt"
	"strings"
)

func main() {
	hasSuffix()
}

func hasSuffix() {
	variable := strings.HasSuffix("jude", "e")
	fmt.Println(variable)
}
