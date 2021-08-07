package inquiry_abstract

import (
	"context"
	"fmt"
)

type freeConsultation struct {
	Inquiry
}

func (c *freeConsultation) lockRegister(ctx context.Context) (err error) {
	fmt.Println("freeConsultation lockRegister")

	return
}
