package main

import (
	"fmt"

	facadepattern "git/go-design-pattern/facade-pattern"
)

// strategypattern "git/go-design-pattern/strategy-pattern"
// "git/go-design-pattern/factory-method-pattern/notifier"
// templatemethodpattern "git/go-design-pattern/template-method-pattern"

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
	// paymentProcessor := &strategypattern.PaymentProcessor{}
	// // Choose payment method
	// // Credit-card-method:
	// creditCardPaymentMethod := strategypattern.CreditCardPayment{}
	// paymentProcessor.ChoosePaymentMethod(&creditCardPaymentMethod)
	// paymentProcessor.ProcessPayment(100)

	// // Paypal-method:
	// paypalPaymentMethod := strategypattern.PaypalPayment{}
	// paymentProcessor.ChoosePaymentMethod(&paypalPaymentMethod)
	// paymentProcessor.ProcessPayment(200)

	// // Crypto-method:
	// cryptoPaymentMethod := strategypattern.CryptoPayment{}
	// paymentProcessor.ChoosePaymentMethod(&cryptoPaymentMethod)
	// paymentProcessor.ProcessPayment(300)

	// ==================== PROXY PATTERN ====================
	// realService := &proxypattern.Service{}
	// nginxProxy := proxypattern.NewNginxProxy(realService, make(map[string]int), 2)
	// status, response := nginxProxy.HandleRequest("https://google.com", "GET")
	// fmt.Println(status, response)

	// status, response = nginxProxy.HandleRequest("https://google.com", "GET")
	// fmt.Println(status, response)

	// status, response = nginxProxy.HandleRequest("https://google.com", "GET")
	// fmt.Println(status, response)

	// ==================== FACADE PATTERN ====================
	// this helps hide the complexity of the subsystem, so the client only need to know the facade interface.
	// And client doesn't need to care to other dependencies from the subsystem (decoupling).
	// The facade pattern is a good choice when you want to simplify a complex system and hide the complexity from the client.
	// The facade pattern collect all the dependencies from the subsystem and provide a simple interface to the client.
	walletFacade := facadepattern.NewMyWalletFacade("0271001108850", "1234", 100)
	if err := walletFacade.AddMoneyToWallet("0271001108850", "1234", 100); err != nil {
		fmt.Println(err)
	}
}
