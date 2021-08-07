package inquiry_abstract

import (
	"context"
	"fmt"
)

type InquiryInterface interface {
	lockRegister(ctx context.Context) (err error)
	createOrder(ctx context.Context) (err error)
	createInquiryOrder(ctx context.Context) (err error)
	createDialog(ctx context.Context) (err error)
}

type Inquiry struct{}

func (*Inquiry) createOrder(ctx context.Context) (err error) {
	fmt.Println("Inquiry createOrder")

	return
}

func (*Inquiry) createInquiryOrder(ctx context.Context) (err error) {
	fmt.Println("Inquiry createInquiryOrder")

	return
}

func (*Inquiry) createDialog(ctx context.Context) (err error) {
	fmt.Println("Inquiry createDialog")

	return
}
