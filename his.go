package inquiry_abstract

import (
	"context"
)

// HisInterface HIS交互接口抽象
type HisInterface interface {
	register(ctx context.Context, orderCode string) (err error)
}
