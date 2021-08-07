package inquiry_abstract

import (
	"context"
	"fmt"
	"strconv"
)

type hisTwoAdapter struct {
	his hisTwo
}

func (h *hisTwoAdapter) register(ctx context.Context, _orderCode string) (err error) {
	orderCode, _ := strconv.Atoi(_orderCode)
	h.his.register(ctx, orderCode)

	return
}

type hisTwo struct{}

func (*hisTwo) register(ctx context.Context, orderCode int) (err error) {
	fmt.Println("hisB register")

	return
}
