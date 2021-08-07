package inquiry_abstract

import (
	"context"
	"fmt"
)

// consultation 收费咨询业务实现
type consultation struct {
	Inquiry
}

func (c *consultation) lockRegister(ctx context.Context) (err error) {
	fmt.Println("consultation lockRegister")

	return
}
