package main

import (
	strategypattern "git/go-design-pattern/strategy-pattern"
	// "git/go-design-pattern/factory-method-pattern/notifier"
	// templatemethodpattern "git/go-design-pattern/template-method-pattern"
)

func main() {

	// =================== TEMPLATE METHOD PATTERN ==========================
	//otp := &templatemethodpattern.Otp{}

	//// Send SMS Otp
	//smsOtp := &templatemethodpattern.SmsOtp{}
	//otp.IOtp = smsOtp
	//otp.HandleSendOTP(4)
	//fmt.Println("--------------------------------")
	//// Send Mail Otp
	//mailOtp := &templatemethodpattern.MailOtp{}
	//otp.IOtp = mailOtp
	//otp.HandleSendOTP(5)

	// ==================== FACTORY METHOD PATTERN ============================
	// slackFactory := &notifactory.SlackNotifierFactory{}
	// notifier := slackFactory.CreateNotifier()
	// notifier.SendNotification("Hello")

	// mailFactory := &notifactory.MailNotifierFactory{}
	// notifier = mailFactory.CreateNotifier()
	// notifier.SendNotification("slack")

	// ==================== STRATEGY PATTERN ============================
	paymentProcessor := &strategypattern.PaymentProcessor{}
	// Choose payment method
	// Credit-card-method:
	creditCardPaymentMethod := strategypattern.CreditCardPayment{}
	paymentProcessor.ChoosePaymentMethod(&creditCardPaymentMethod)
	paymentProcessor.ProcessPayment(100)

	// Paypal-method:
	paypalPaymentMethod := strategypattern.PaypalPayment{}
	paymentProcessor.ChoosePaymentMethod(&paypalPaymentMethod)
	paymentProcessor.ProcessPayment(200)

	// Crypto-method:
	cryptoPaymentMethod := strategypattern.CryptoPayment{}
	paymentProcessor.ChoosePaymentMethod(&cryptoPaymentMethod)
	paymentProcessor.ProcessPayment(300)
}
