package inquiry_abstract

import "context"

// HandlerInterface 标准流程抽象Handler
type HandlerInterface interface {
	// CreateOrder 创建订单
	CreateOrder(ctx context.Context, inquiry InquiryInterface, his HisInterface, orderCode string) (err error)
}

type Handler struct{}

func (*Handler) CreateOrder(ctx context.Context, inquiry InquiryInterface, his HisInterface, orderCode string) (err error) {
	inquiry.lockRegister(ctx)
	his.register(ctx, orderCode)
	inquiry.createOrder(ctx)
	inquiry.createInquiryOrder(ctx)
	inquiry.createDialog(ctx)

	return
}
