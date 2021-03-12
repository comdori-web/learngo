package main

import (
	"fmt"

	"github.com/comdori-web/learngo/banking"
)

func main() {
	account := banking.Account{Owner: "kyuha", Balance: 100}
	fmt.Println(account)
}
