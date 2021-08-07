package inquiry_abstract

import (
	"context"
	"errors"
)

const (
	HisCodeOne = "one"
	HisCodeTwo = "two"
)

var (
	ErrHIS = errors.New("his error")
)

var (
	HisOne HisInterface
	HisTwo HisInterface
)

func init() {
	HisOne = &hisOneAdapter{}
	HisTwo = &hisTwoAdapter{}
}

func NewHis(ctx context.Context, hisCode string) (his HisInterface, err error) {
	switch hisCode {
	case HisCodeOne:
		his = HisOne
	case HisCodeTwo:
		his = HisTwo
	default:
		err = ErrHIS
	}

	return
}
