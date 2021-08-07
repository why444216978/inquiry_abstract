package inquiry_abstract

import (
	"context"
)

type HisInterface interface {
	register(ctx context.Context, orderCode string) (err error)
}
