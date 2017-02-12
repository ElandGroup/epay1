package alConst

const (
	OpenApi = "https://openapi.alipay.com/gateway.do"
)

type Person struct {
	Name string
}

//const For request alipay
const (

	//customer
	ReqPay           = "alipay.trade.pay"
	ReqReverse       = "alipay.trade.cancel"
	ReqRefund        = "alipay.trade.refund"
	ReqQuery         = "alipay.trade.query"
	ReqScenceBarCode = "bar_code"

	ReqCharset  = "utf-8"
	ReqVersion  = "1.0"
	ReqSignType = "RSA"
	//common part
	RawAppId      = "app_id"
	RawMethod     = "method"
	RawTimeStamp  = "timestamp"
	RawCharset    = "charset"
	RawVersion    = "version"
	RawSignType   = "sign_type"
	RawOutTradeNo = "out_trade_no"

	RawBizContent = "biz_content"
	RawSign       = "sign"

	//direct pay
	RawAuthCode    = "auth_code"
	RawTotalAmount = "total_amount"
	RawSubject     = "subject"
	RawStoreId     = "store_id"

	RawSellerId     = "seller_id"
	RawTimeExpire   = "time_expire"
	RawExtendParams = "extend_params"
	RawScence       = "scene"

	//order query
	RawTradeNo = "trade_no"

	//refund
	RawRefundAmount = "refund_amount"
	RawOutRequestNo = "out_request_no"
	RawRefundReason = "refund_reason"
	//service provider
	RawSysServiceProviderId = "sys_service_provider_id"
)

//const For response alipay
const (
	RespPay     = "alipay_trade_pay_response"
	RespReverse = "alipay_trade_cancel_response"
	RespRefund  = "alipay_trade_refund_response"
	RespQuery   = "alipay_trade_query_response"

	//Sign      = "sign"
	RawCode      = "code"
	RawSubCode   = "sub_code"
	RawMsg       = "msg"
	RawSubMsg    = "sub_msg"
	RawRetryFlag = "retry_flag"

	//direct pay
	//TotalAmount = "total_amount"

	//order Query
	RawTradeStatus = "trade_status"

	//refund
	RawRefundFee   = "refund_fee"
	RawOutRefundNo = "out_refund_no"
	//OutTradeNo  = "out_trade_no"
	//TradeNo     = "trade_no"

	//Reverse

)

const (
	AppId            = "AppId"
	SellerPrivateKey = "SellerPrivateKey"
	AliPublicKey     = "AliPublicKey"
	OutTradeNo       = "OutTradeNo"
	ALAuthToken      = "ALAuthToken"

	AuthCode    = "AuthCode"
	TotalAmount = "TotalAmount"
	Subject     = "Subject"
	StoreId     = "StoreId"

	SellerId     = "SellerId"
	TimeExpire   = "TimeExpire"
	ExtendParams = "ExtendParams"
	//query
	TradeNo = "TradeNo"

	//refund
	OutRequestNo         = "OutRequestNo"
	RefundReason         = "RefundReason"
	RefundAmount         = "RefundAmount"
	SysServiceProviderId = "SysServiceProviderId"
)
