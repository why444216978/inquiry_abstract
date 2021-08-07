package inquiry_abstract

import (
	"context"
	"fmt"
)

// freeConsultation 免费咨询业务实现
type freeConsultation struct {
	Inquiry
}

func (c *freeConsultation) lockRegister(ctx context.Context) (err error) {
	fmt.Println("freeConsultation lockRegister")

	return
}
