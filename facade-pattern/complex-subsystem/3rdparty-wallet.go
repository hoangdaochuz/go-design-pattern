package complexsubsystem

import "fmt"

type ThirdPartyWallet struct {
	balance int
}

func NewThirdPartyWallet(balance int) *ThirdPartyWallet {
	return &ThirdPartyWallet{
		balance: balance,
	}
}

func (wallet *ThirdPartyWallet) CreditBalance(amount int) error {
	wallet.balance += amount
	fmt.Println("Credit balance successfully")
	return nil
}

func (wallet *ThirdPartyWallet) DebitBalance(amount int) error {
	if amount > wallet.balance {
		return fmt.Errorf("insufficient balance")
	}
	wallet.balance -= amount
	fmt.Println("Debit balance successfully")
	return nil
}
