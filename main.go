package main

import "fmt"

func main() {

	names := []string{"ni", "ci", "ki"}
	names = append(names, "ka")

	fmt.Println(names)
}
