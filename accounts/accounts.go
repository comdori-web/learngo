package accounts

// Account struct
type Account struct {
	owner   string
	balance int
}

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	// 오브젝트 복사를 막기 위해 포인터로 리턴해줌
	return &account
}

// Deposit x amount on your account
func (a Account) Deposit(amount int) {
	// a Account : receiver
	// receiver로 method를 만들 수 있음
	// a는 다른 이름이어도 된다.
	// 대신 struct의 첫 글자를 따서 소문자로 지어야 한다!
	a.balance += amount
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}
