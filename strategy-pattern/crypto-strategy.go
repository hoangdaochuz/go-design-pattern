package strategypattern

import "fmt"

type CryptoPayment struct{}

func (c *CryptoPayment) Pay(amount float64) error {
	fmt.Printf("Processing CRYPTO payment of $%.2f\n", amount)
	return nil
}
