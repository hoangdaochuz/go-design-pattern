package main

import (
	commandpattern "git/go-design-pattern/command-pattern"
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
	// walletFacade := facadepattern.NewMyWalletFacade("0271001108850", "1234", 100)
	// if err := walletFacade.AddMoneyToWallet("0271001108850", "1234", 100); err != nil {
	// 	fmt.Println(err)
	// }

	// ==================== ADAPTER PATTERN ====================
	// current system is using typeC port to charge the phone, but we have a new phone that uses lightning port.
	// we need to create an adapter to make the lightning port compatible with the typeC port.

	// typeCCharger := adapterpattern.NewTypeCCharger()
	// typeCCharger.Charge()

	// lightningCharger := adapterpattern.NewLightningCharger()
	// typeCAdapter := adapterpattern.NewTypeCChargerAdapter(lightningCharger)
	// typeCAdapter.Charge()

	// ==================== DECORATOR PATTERN ====================
	// we want to add a new feature to the existing code without changing the existing code.
	// for example, we want to add a new feature to the existing code to log the request and response.
	// we can create a new decorator to add the new feature to the existing code.

	// // create a base pizza
	// basePizza := decoratorpattern.NewBasePizza("Base pizza ", 100)
	// fmt.Println(basePizza.GetDescription(), basePizza.GetPrice())

	// // add toppings to the base pizza
	// // basePizza + tomato ==> tomato pizza
	// tomatoPizza := decoratorpattern.NewTomatoPizza(basePizza)
	// fmt.Println(tomatoPizza.GetDescription(), tomatoPizza.GetPrice())

	// // tomatoPizza + cheese ==> tomato chees pizza
	// tomatoCheesePizza := decoratorpattern.NewCheesePizza(tomatoPizza)
	// fmt.Println(tomatoCheesePizza.GetDescription(), tomatoCheesePizza.GetPrice())

	// // tomatoCheesePizza + seafood ==> tomato cheese seafood pizza
	// tomatoCheeseSeafoodPizza := decoratorpattern.NewSeafoodPizza(tomatoCheesePizza)
	// fmt.Println(tomatoCheeseSeafoodPizza.GetDescription(), tomatoCheeseSeafoodPizza.GetPrice())

	// // basePizza + seafood ==> seafood pizza
	// seafoodPizza := decoratorpattern.NewSeafoodPizza(basePizza)
	// fmt.Println(seafoodPizza.GetDescription(), seafoodPizza.GetPrice())

	// ==================== SINGLETON PATTERN ====================
	// start 100 goroutine to get the singleton instance
	// for i := 0; i < 100; i++ {
	// 	go func() {
	// 		_ = singletonpattern.GetInstance()
	// 	}()
	// }

	// ==================== ABSTRACT FACTORY PATTERN ====================
	// nikeFactory := &productfactory.NikeFactory{}
	// adidasFactory := &productfactory.AdidasFactory{}
	// pumaFactory := &productfactory.PumaFactory{}
	// suppose user want to buy nike shoes and other products, we can use the abstract factory to create the products
	// we can make sure the products are compatible with each other
	// nikeFactory.CreateShoe()
	// nikeFactory.CreatePant()
	// if user want to buy adidas shoes and other products, we can use the abstract factory to create the products and so on
	// On the real product, it will have config for user to choose the brand and the products they want to buy or sth like that

	// ==================== BUILDER PATTERN ====================
	// we want to build a house with different types of houses
	// regular house, igloo, etc.
	// we can use the builder pattern to build the house
	// we can use the director to build the house
	// we can use the builder to build the house

	// houseStoneBuilder := builderpattern.NewHouseStone()
	// houseWoodBuilder := builderpattern.NewHouseWood()

	// director := builderpattern.NewHouseDirector(houseStoneBuilder)
	// director.BuildHouse("stone", "Blue")

	// director = builderpattern.NewHouseDirector(houseWoodBuilder)
	// director.BuildHouse("wood", "Red")

	// ==================== PROTOTYPE PATTERN ====================
	// we want to create a new object by cloning the existing object
	// we can use the prototype pattern to create a new object by cloning the existing object
	// file1 := prototypepattern.NewFile("File1")
	// file2 := prototypepattern.NewFile("File2")
	// file3 := prototypepattern.NewFile("File3")
	// file4 := prototypepattern.NewFile("File4")
	// subFolder1 := prototypepattern.NewFolder("SubFolder1", []prototypepattern.INode{
	// 	file4,
	// })

	// folder1 := prototypepattern.NewFolder("Folder1", []prototypepattern.INode{
	// 	file1,
	// 	file2,
	// 	file3,
	// 	subFolder1,
	// })

	// folder1.Print()
	// fmt.Println("--------------------------------")
	// folder2 := folder1.Clone()
	// folder2.Print()

	// =============== OBSERVER PATTERN =============
	// user1 := observerpattern.NewUser("khai", "1")
	// user2 := observerpattern.NewUser("Khang", "2")

	// youtubeChannel := observerpattern.NewYoutubeChannel("Mixi Gaming")

	// youtubeChannel.Register(user1)
	// youtubeChannel.Register(user2)

	// youtubeChannel.Notify("New videos have published by " + youtubeChannel.GetChannelName())

	// =============== CHAIN OF RESPONSIBILITY ====================
	// reception := chainofresponsibilitypattern.ReceptionDepartment{}
	// doctor := chainofresponsibilitypattern.DoctorDepartment{}
	// medication := chainofresponsibilitypattern.MedicationDepartment{}
	// payment := chainofresponsibilitypattern.PaymentDepartment{}

	// reception.SetNext(&doctor)
	// doctor.SetNext(&medication)
	// medication.SetNext(&payment)

	// patient := chainofresponsibilitypattern.NewPatient("Khai")

	// reception.Execute(*patient)

	// ================= COMMAND PATTERN ========================
	tv := commandpattern.NewTV(false)
	onCommand := commandpattern.NewOnCommand(tv)
	offCommand := commandpattern.NewOffCommand(tv)
	button := &commandpattern.Button{}
	button.SetCommand(onCommand)
	button.Press()
	button.SetCommand(offCommand)
	button.Press()
}
