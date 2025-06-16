package facadepattern

import (
	"fmt"
	complexsubsystem "git/go-design-pattern/facade-pattern/complex-subsystem"
)

type MyWalletFacade struct {
	account      *complexsubsystem.Account
	securityCode *complexsubsystem.SecurityCode
	wallet       *complexsubsystem.ThirdPartyWallet
	logger       *complexsubsystem.Logger
	notification *complexsubsystem.Notification
}

func NewMyWalletFacade(accountNumber string, code string, balance int) *MyWalletFacade {
	walletAccount := &MyWalletFacade{
		account:      complexsubsystem.NewAccount(accountNumber),
		securityCode: complexsubsystem.NewSecurityCode(code),
		wallet:       complexsubsystem.NewThirdPartyWallet(balance),
		logger:       complexsubsystem.NewLogger(),
		notification: complexsubsystem.NewNotification(),
	}
	fmt.Println("New account created")
	return walletAccount
}

func (walletFacace *MyWalletFacade) AddMoneyToWallet(accountNumber, securityCode string, amountMoney int) error {
	fmt.Println("Processing Add money to wallet")
	// Check if account is valid
	isValidAccount := walletFacace.account.CheckValidAccount(accountNumber)
	if !isValidAccount {
		return fmt.Errorf("account number is not valid")
	}

	// Check if security code is valid
	isValidSecurityCode := walletFacace.securityCode.CheckValidSecurityCode(securityCode)
	if !isValidSecurityCode {
		return fmt.Errorf("security code is not valid")
	}

	err := walletFacace.wallet.CreditBalance(amountMoney)
	if err != nil {
		return err
	}
	err = walletFacace.logger.Log(accountNumber, amountMoney, "Add money")
	if err != nil {
		return err
	}
	// Send notification to account owner
	err = walletFacace.notification.SendNotification("Add money to wallet successfully", accountNumber)
	if err != nil {
		return err
	}
	fmt.Println("Add money to wallet successfully")
	return nil
}
