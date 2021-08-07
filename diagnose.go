package inquiry_abstract

import (
	"context"
	"fmt"
)

type diagnose struct {
	Inquiry
}

func (c *diagnose) lockRegister(ctx context.Context) (err error) {
	fmt.Println("diagnose lockRegister")

	return
}
