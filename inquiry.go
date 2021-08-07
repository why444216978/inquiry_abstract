package inquiry_abstract

import (
	"context"
	"fmt"
)

// InquiryInterface 问诊接口抽象
type InquiryInterface interface {
	// lockRegister 锁号
	lockRegister(ctx context.Context) (err error)
	// createOrder 创建支付订单
	createOrder(ctx context.Context) (err error)
	// createInquiryOrder 创建问诊订单
	createInquiryOrder(ctx context.Context) (err error)
	// createDialog 创建聊天会话
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
