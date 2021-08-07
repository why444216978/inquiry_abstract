package inquiry_abstract

import (
	"context"
	"fmt"
)

type consultation struct {
	Inquiry
}

func (c *consultation) lockRegister(ctx context.Context) (err error) {
	fmt.Println("consultation lockRegister")

	return
}
