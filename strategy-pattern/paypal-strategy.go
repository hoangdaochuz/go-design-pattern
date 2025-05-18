package strategypattern

import "fmt"

type PaypalPayment struct{}

func (p *PaypalPayment) Pay(amount float64) error {
	fmt.Printf("Processing PAYPAL payment of $%.2f\n", amount)
	return nil
}
