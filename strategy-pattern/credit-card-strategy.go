package strategypattern

import "fmt"

type CreditCardPayment struct{}

func (c *CreditCardPayment) Pay(amount float64) error {
	fmt.Printf("Processing CREDIT CARD payment of $%.2f\n", amount)
	return nil
}
