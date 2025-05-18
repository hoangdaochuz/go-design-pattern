package strategypattern

type PaymentStrategy interface {
	Pay(amount float64) error
}

type PaymentProcessor struct {
	PaymentMethod PaymentStrategy
}

func (p *PaymentProcessor) ChoosePaymentMethod(method PaymentStrategy) {
	p.PaymentMethod = method
}

func (p *PaymentProcessor) ProcessPayment(amount float64) error {
	if err := p.PaymentMethod.Pay(amount); err != nil {
		return err
	}
	return nil
}
