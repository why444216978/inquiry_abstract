package inquiry_abstract

import (
	"context"
	"fmt"
	"log"
	"testing"
)

// TestHandler_CreateOrder 创建订单 4*2 八种流程测试
func TestHandler_CreateOrder(t *testing.T) {
	ctx := context.Background()

	consultation, err := NewInquiry(ctx, TypeConsultation, 0)
	if err != nil {
		log.Fatal(err)
		return
	}

	freeConsultation, err := NewInquiry(ctx, TypeConsultation, 1)
	if err != nil {
		log.Fatal(err)
		return
	}

	freeDiagnose, err := NewInquiry(ctx, TypeDiagnose, 0)
	if err != nil {
		log.Fatal(err)
		return
	}

	diagnose, err := NewInquiry(ctx, TypeDiagnose, 1)
	if err != nil {
		log.Fatal(err)
		return
	}

	hisOne, err := NewHis(ctx, HisCodeOne)
	if err != nil {
		log.Fatal(err)
		return
	}

	hisTwo, err := NewHis(ctx, HisCodeTwo)
	if err != nil {
		log.Fatal(err)
		return
	}

	orderCode := "abc"

	handler := &Handler{}

	fmt.Println("\n----------consultation hisOne----------")
	handler.CreateOrder(ctx, consultation, hisOne, orderCode)

	fmt.Println("\n----------freeConsultation hisOne----------")
	handler.CreateOrder(ctx, freeConsultation, hisOne, orderCode)

	fmt.Println("\n----------diagnose hisOne----------")
	handler.CreateOrder(ctx, diagnose, hisOne, orderCode)

	fmt.Println("\n----------freeDiagnose hisOne----------")
	handler.CreateOrder(ctx, freeDiagnose, hisOne, orderCode)

	fmt.Println("\n----------consultation hisTwo----------")
	handler.CreateOrder(ctx, consultation, hisTwo, orderCode)

	fmt.Println("\n----------freeConsultation hisTwo----------")
	handler.CreateOrder(ctx, freeConsultation, hisTwo, orderCode)

	fmt.Println("\n----------diagnose hisTwo----------")
	handler.CreateOrder(ctx, diagnose, hisTwo, orderCode)

	fmt.Println("\n----------freeDiagnose hisTwo----------")
	handler.CreateOrder(ctx, freeDiagnose, hisTwo, orderCode)

}
