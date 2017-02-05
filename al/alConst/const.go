package alConst

const (
	OpenApi          = "https://openapi.alipay.com/gateway.do"
	SellerPrivateKey = "SellerPrivateKey"
	AliPublicKey     = "AliPublicKey"
)

type Person struct {
	Name string
}

//const For request alipay
const (

	//customer
	PayCustom           = "alipay.trade.pay"
	ReverseCustom       = "alipay.trade.cancel"
	RefundCustom        = "alipay.trade.refund"
	QueryCustom         = "alipay.trade.query"
	BarCodeScenceCustom = "bar_code"

	CharsetCustom  = "utf-8"
	VersionCustom  = "1.0"
	SignTypeCustom = "RSA"
	//common part
	AppId      = "app_id"
	Method     = "method"
	TimeStamp  = "timestamp"
	Charset    = "charset"
	Version    = "version"
	SignType   = "sign_type"
	OutTradeNo = "out_trade_no"

	BizContent = "biz_content"
	Sign       = "sign"

	//direct pay
	AuthCode    = "auth_code"
	TotalAmount = "total_amount"
	Subject     = "subject"
	StoreId     = "store_id"

	SellerId     = "seller_id"
	TimeExpire   = "time_expire"
	ExtendParams = "extend_params"
	Scence       = "scene"

	//order query
	TradeNo = "trade_no"

	//refund
	RefundAmount = "refund_amount"
	OutRequestNo = "out_request_no"
	RefundReason = "refund_reason"
	//service provider
	SysServiceProviderId = "sys_service_provider_id"
)

//const For response alipay
const (
	//Sign      = "sign"
	Code      = "code"
	SubCode   = "sub_code"
	Msg       = "msg"
	SubMsg    = "sub_msg"
	RetryFlag = "retry_flag"

	PayNode     = "alipay_trade_pay_response"
	ReverseNode = "alipay_trade_cancel_response"
	RefundNode  = "alipay_trade_refund_response"
	QueryNode   = "alipay_trade_query_response"
	//direct pay
	//TotalAmount = "total_amount"

	//order Query
	TradeStatus = "trade_status"

	//refund
	RefundFee   = "refund_fee"
	OutRefundNo = "out_refund_no"
	//OutTradeNo  = "out_trade_no"
	//TradeNo     = "trade_no"

	//Reverse

)

const (
	AppIdMap            = "AppId"
	SellerPrivateKeyMap = "SellerPrivateKey"
	AliPublicKeyMap     = "AliPublicKey"
	OutTradeNoMap       = "OutTradeNo"
	ALAuthTokenMap      = "ALAuthToken"

	AuthCodeMap    = "AuthCode"
	TotalAmountMap = "TotalAmount"
	SubjectMap     = "Subject"
	StoreIdMap     = "StoreId"

	SellerIdMap     = "SellerId"
	TimeExpireMap   = "TimeExpire"
	ExtendParamsMap = "ExtendParams"
	//query
	TradeNoMap = "TradeNo"

	//refund
	OutRequestNoMap = "OutRequestNo"
	RefundReasonMap = "RefundReason"
	RefundAmountMap = "RefundAmount"
)
