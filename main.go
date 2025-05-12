package main

import (
	// "fmt"
	notifactory "git/go-design-pattern/factory-method-pattern/noti-factory"
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
	slackFactory := &notifactory.SlackNotifierFactory{}
	notifier := slackFactory.CreateNotifier()
	notifier.SendNotification("Hello")

	mailFactory := &notifactory.MailNotifierFactory{}
	notifier = mailFactory.CreateNotifier()
	notifier.SendNotification("slack")

}
