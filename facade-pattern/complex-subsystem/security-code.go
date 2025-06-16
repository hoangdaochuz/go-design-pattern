package complexsubsystem

type SecurityCode struct {
	code string
}

func NewSecurityCode(code string) *SecurityCode {
	return &SecurityCode{
		code: code,
	}
}

func (sc *SecurityCode) CheckValidSecurityCode(code string) bool {
	return sc.code == code
}
