package input

import "fmt"

// Capitalizing names will export them (variables, functions...)

func GetString(prompt string) string {
	var str string
	fmt.Printf("%s: ", prompt)
	fmt.Scan(&str) // pointer for strings (!= C)
	return str
}

func GetInt(prompt string) int {
	var num int
	fmt.Printf("%s: ", prompt)
	fmt.Scan(&num) // pointer for integers (!= C)
	return num
}
