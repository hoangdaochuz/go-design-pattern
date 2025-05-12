package main

import (
	"fmt"
	templatemethodpattern "git/go-design-pattern/template-method-pattern"
)

func main() {
	otp := &templatemethodpattern.Otp{}

	// Send SMS Otp
	smsOtp := &templatemethodpattern.SmsOtp{}
	otp.IOtp = smsOtp
	otp.HandleSendOTP(4)
	fmt.Println("--------------------------------")
	// Send Mail Otp
	mailOtp := &templatemethodpattern.MailOtp{}
	otp.IOtp = mailOtp
	otp.HandleSendOTP(5)
}
