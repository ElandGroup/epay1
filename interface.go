package epaygo

import (
	. "epaygo/core/common"
)

type IPayService interface {
	DirectPay(params map[string]string) (result string, apiError *APIError)
	OrderQuery(params map[string]string) (result string, apiError *APIError)
	Refund(params map[string]string) (result string, apiError *APIError)
	OrderReverse(params map[string]string, count int) (result string, apiError *APIError)
}
