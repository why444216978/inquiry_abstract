package inquiry_abstract

import (
	"context"
	"fmt"
)

type freeDiagnose struct {
	Inquiry
}

func (c *freeDiagnose) lockRegister(ctx context.Context) (err error) {
	fmt.Println("freeDiagnose lockRegister")

	return
}
