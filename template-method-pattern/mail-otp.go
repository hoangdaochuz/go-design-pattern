package templatemethodpattern

import "fmt"

type MailOtp struct {
	Otp
}

func (m *MailOtp) GenerateOtpCode(otpLength int) string {
	fmt.Println("generated mail otp code")
	otpCode := "alidwed"
	return otpCode
}

func (m *MailOtp) SaveOtpCache(otp string) error {
	fmt.Println("Saved MAIL OTP to cache")
	return nil
}

func (m *MailOtp) GetMessage(message string) string {
	fmt.Println("Get Mail Otp message successfully")

	return message
}

func (m *MailOtp) SendNotification(message string) error {
	fmt.Println("Send Mail Otp notification successfully")
	return nil
}
