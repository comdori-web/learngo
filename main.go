package main

import (
	"fmt"

	"github.com/comdori-web/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("kyuha")
	fmt.Println(account)
}
