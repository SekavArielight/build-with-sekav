package main

// import "fmt"

// import "piscine"

func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	// panic("Please implement the CanSpy() function")
	// return knightIsAwake && archerIsAwake
	if knightIsAwake == true && archerIsAwake == true && prisonerIsAwake == true{
		return false
	}
	if knightIsAwake == false && archerIsAwake == true || knightIsAwake == true && archerIsAwake == false {
		return false
	}
	return true
}

// func main() {
// 	PrintComb()
// 	var knightIsAwake = false
// 	var archerIsAwake = true
// 	var prisonerIsAwake = false
// 	fmt.Println(CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake))
// }
