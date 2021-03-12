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
