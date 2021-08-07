package inquiry_abstract

import (
	"context"
	"fmt"
)

type hisOneAdapter struct {
	his hisOne
}

func (h *hisOneAdapter) register(ctx context.Context, orderCode string) (err error) {
	h.his.register(ctx, orderCode)

	return
}

type hisOne struct{}

func (*hisOne) register(ctx context.Context, orderCode string) (err error) {
	fmt.Println("hisA register")

	return
}
