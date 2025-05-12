package templatemethodpattern

type IOtp interface {
	GenerateOtpCode(int) string
	SaveOtpCache(string) error
	GetMessage(string) string
	SendNotification(string) error
}

type Otp struct {
	IOtp IOtp // declare IOtp field to inject specific subclass that implement baseclass
}

// This is a common algorithm to send OTP notification

func (o *Otp) HandleSendOTP(otpLength int) error {
	otpString := o.IOtp.GenerateOtpCode(otpLength)
	if err := o.IOtp.SaveOtpCache(otpString); err != nil {
		return err
	}

	message := o.IOtp.GetMessage(otpString)

	if err := o.IOtp.SendNotification(message); err != nil {
		return err
	}
	return nil
}
