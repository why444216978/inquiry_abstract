package inquiry_abstract

import (
	"context"
	"errors"
)

const (
	TypeConsultation uint8 = 1
	TypeDiagnose     uint8 = 1
)

var (
	ErrType = errors.New("type error")
)

var (
	Consultation     InquiryInterface
	FreeConsultation InquiryInterface
	Diagnose         InquiryInterface
	FreeDiagnose     InquiryInterface
)

// init Inquiry单例
func init() {
	Consultation = &consultation{}
	FreeConsultation = &freeConsultation{}
	Diagnose = &diagnose{}
	FreeDiagnose = &freeDiagnose{}
}

// NewInquiry Inquiry工厂
func NewInquiry(ctx context.Context, typ uint8, price int) (obj InquiryInterface, err error) {
	obj = Consultation

	if typ == TypeConsultation && price == 0 {
		obj = FreeConsultation
		return
	}

	if typ == TypeDiagnose && price == 0 {
		obj = FreeDiagnose
		return
	}

	if typ == TypeDiagnose && price > 0 {
		obj = Diagnose
		return
	}

	err = ErrType

	return
}
