package main

import (
	"fmt"

	"github.com/jkyoung37/learngo/accounts"
)

func main() {

	account := accounts.NewAccount("nico")
	fmt.Println(account)

}
