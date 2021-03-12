package main

import (
	"fmt"

	"github.com/comdori-web/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("kyuha")
	account.Deposit(10)

	// go 에서 error를 처리하는 방법임.
	err := account.Withdraw(20)
	if err != nil {
		// log.Fatalln(err) // 프로그램을 종료 시켜줌
		fmt.Println(err)
	}
	fmt.Println(account) // go 가 자동으로 String() 을 호출해주고 있는거임. like python
}
