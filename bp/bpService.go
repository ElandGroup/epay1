package bp

import (
	. "epaygo/core/commonDto"
)

type BpPayService struct {
}

func (BpPayService) DirectPay(params map[string]string) (result string, apiError *APIError) {
	if params == nil {
	}
	return "", nil
}
func (BpPayService) OrderQuery(params map[string]string) (result string, apiError *APIError) {
	if params == nil {
	}
	return "", nil
}
func (BpPayService) Refund(params map[string]string) (result string, apiError *APIError) {
	if params == nil {
	}
	return "", nil
}
func (BpPayService) OrderReverse(params map[string]string, count int) (result string, apiError *APIError) {
	if params == nil {
	}
	return "", nil
}
