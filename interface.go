package epaygo

type IPayService interface {
	DirectPay(params map[string]string) (result string, err error)
	OrderQuery(params map[string]string) (result string, err error)
	Refund(params map[string]string) (result string, err error)
	OrderReverse(params map[string]string, count int) (result string, err error)
}
