package accounts

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("Can't withdraw")

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	// 오브젝트 복사를 막기 위해 포인터로 리턴해줌
	return &account
}

// Deposit x amount on your account
func (a *Account) Deposit(amount int) {
	// a Account : receiver
	// receiver로 method를 만들 수 있음
	// a는 다른 이름이어도 된다.
	// 대신 struct의 첫 글자를 따서 소문자로 지어야 한다!
	a.balance += amount
	// 신기한 것은 포인터로 변경해도 호출부분과 접근 부분의 차이가 없네...
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

// Withdraw x amount from your account
func (a *Account) Withdraw(amount int) error {
	// go에는 exception 같은 것이 없다.
	// err를 return을 직접 해줘야 한다.
	if a.balance < amount {
		//return error.Error()
		return errNoMoney
	}

	a.balance -= amount

	return nil
}

// ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner of the account
func (a Account) Owner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account.\nHas : ", a.Balance())
}
