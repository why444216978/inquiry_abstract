package inquiry_abstract

import (
	"context"
	"fmt"
)

// freeDiagnose 免费问诊业务实现
type freeDiagnose struct {
	Inquiry
}

func (c *freeDiagnose) lockRegister(ctx context.Context) (err error) {
	fmt.Println("freeDiagnose lockRegister")

	return
}
