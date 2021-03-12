package main

import (
	"fmt"
	"log"

	"github.com/comdori-web/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("kyuha")
	account.Deposit(10)
	fmt.Println(account.Balance())

	// go 에서 error를 처리하는 방법임.
	err := account.Withdraw(20)
	if err != nil {
		log.Fatalln(err) // 프로그램을 종료 시켜줌
	}
	fmt.Println(account.Balance())
}
