package complexsubsystem

import "fmt"

type Logger struct {
}

func NewLogger() *Logger {
	return &Logger{}
}

func (logger *Logger) Log(accountNumber string, ammount int, operation string) error {
	fmt.Printf("Account %s %s %d\n", accountNumber, operation, ammount)
	return nil
}
