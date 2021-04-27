package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatME(words ...string) {

	fmt.Println(words)
}

func main() {
	repeatME("2", "3", "4", "5")

}
