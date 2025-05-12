package templatemethodpattern

import "fmt"

type SmsOtp struct {
	Otp // embedded Otp struct to alter all field and method of it
}

func (s *SmsOtp) GenerateOtpCode(otpLength int) string {
	fmt.Println("generated SMS OTP code")
	otpCode := "2323434"
	return otpCode
}
func (s *SmsOtp) SaveOtpCache(otpCode string) error {

	fmt.Println("Saved generated SMS OTP code")
	return nil
}

func (s *SmsOtp) GetMessage(message string) string {
	fmt.Println("Get SMS OTP message successfully")
	return message
}

func (s *SmsOtp) SendNotification(message string) error {
	fmt.Println("Sending SMS OTP notification")
	return nil
}
