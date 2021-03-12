package main

import (
	"fmt"

	"github.com/comdori-web/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("kyuha")
	account.Deposit(10)
	fmt.Println(account.Balance())
}
