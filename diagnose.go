package inquiry_abstract

import (
	"context"
	"fmt"
)

// diagnose 收费复诊业务实现
type diagnose struct {
	Inquiry
}

func (c *diagnose) lockRegister(ctx context.Context) (err error) {
	fmt.Println("diagnose lockRegister")

	return
}
