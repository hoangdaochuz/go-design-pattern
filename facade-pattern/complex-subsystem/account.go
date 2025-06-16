package complexsubsystem

type Account struct {
	accountNumber string
}

func NewAccount(accountNumber string) *Account {
	return &Account{
		accountNumber: accountNumber,
	}
}

func (acc *Account) CheckValidAccount(accountNumber string) bool {
	return acc.accountNumber == accountNumber
}
