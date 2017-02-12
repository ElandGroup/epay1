package wxConst

// api for wechar
const (
	MicroPay_Url   = "https://api.mch.weixin.qq.com/pay/micropay"
	OrderQuery_Url = "https://api.mch.weixin.qq.com/pay/orderquery"
	Refund_Url     = "https://api.mch.weixin.qq.com/secapi/pay/refund"
	Reverse_Url    = "https://api.mch.weixin.qq.com/secapi/pay/reverse"
)

//const For Request wechat
const (

	//For WxPayBaseDto
	RawAppId          = "appid"
	RawMchId          = "mch_id"
	RawNonceStr       = "nonce_str"
	RawSubMchId       = "sub_mch_id"
	RawSubAppId       = "sub_appid"
	RawSpbillCreateIp = "spbill_create_ip" //此参数可手动配置也可在程序中自动获取
	RawDeviceInfo     = "device_info"

	RawSign       = "sign"
	RawBody       = "body"
	RawDetail     = "detail"
	RawAttach     = "attach"
	RawOutTradeNo = "out_trade_no"
	RawFeeType    = "fee_type"

	RawTotalFee = "total_fee"
	RawGoodsTag = "goods_tag"
	RawLimitPay = "limit_pay"

	//For WxDirectPayDto
	RawAuthCode = "auth_code"
	//For Sign
	RawKey = "key"

	//For WxPayPrecreateDto
	RawTimeStart  = "time_start"
	RawTimeExpire = "time_expire"
	RawNotifyUrl  = "notify_url"
	RawTradeType  = "trade_type"
	RawProductId  = "product_id"
	RawOpenId     = "openid"

	//For OrderQuery
	RawTransactionId = "transaction_id"

	//For Refund
	/// <summary>
	/// 商户系统内部的退款单号，商户系统内部唯一，同一退款单号多次请求只退一笔
	/// </summary>
	RawOutRefundNo = "out_refund_no"
	/// <summary>
	/// 退款总金额，订单总金额，单位为分，只能为整数，详见支付金额
	/// </summary>
	RawRefundFee = "refund_fee"
	/// <summary>
	/// 货币类型，符合ISO 4217标准的三位字母代码，默认人民币：CNY，其他值列表详见货币类型
	/// </summary>
	RawRefundFeeType = "refund_fee_type"
	/// <summary>
	/// 操作员帐号, 默认为商户号
	/// </summary>
	RawOpUserId = "op_user_id"
	/// <summary>
	/// 微信生成的退款单号，在申请退款接口有返回
	/// </summary>
	RawRefundId = "refund_id"

	//For DownloadBill
	/// <summary>
	/// 下载对账单的日期，格式：20140603
	/// </summary>
	RawBillDate = "bill_date"
	/// <summary>
	/// ALL，返回当日所有订单信息，默认值,
	/// SUCCESS，返回当日成功支付的订单,
	/// REFUND，返回当日退款订单,
	/// REVOKED，已撤销的订单
	/// </summary>
	RawBillType = "bill_type"

	//For QRCode
	/// <summary>
	/// 时间戳
	/// </summary>
	RawTimeStamp = "time_stamp"

	//For ShortUrl
	/// <summary>
	/// 需要转换的URL，签名用原串，传输需URLencode
	/// </summary>
	RawLongUrl = "long_url"

	RawSpbillCreateIpValue = "8.8.8.8"
)

//const For response wechat
const (
	RawReturnCode = "return_code"
	RawReturnMsg  = "return_msg"
	RawResultCode = "result_code"
	RawErrCode    = "err_code"
	RawErrCodeDes = "err_code_des"

	RawTradeState = "trade_state"

	RawTradeStateDesc = "trade_state_desc"

	RawRecall = "recall"

	RawTimeEnd = "time_end"

	RawCodeUrl = "code_url"

	RawPrepayId = "prepay_id"

	//TradeType = "trade_type"

	RawBankType = "bank_type"

	//TransactionId = "transaction_id"
	//OutTradeNo    = "out_trade_no"

	//TotalFee = "total_fee"
	//Unified Pay
	JSAppId     = "appId"
	JSTimeStamp = "timeStamp"
	JSNonceStr  = "nonceStr"
	JSPackage   = "package"
	JSSignType  = "signType"
	JSPaySign   = "paySign"

	//refund

	RawRefundCount = "refund_count"

	RawOutRefundNoLike = "out_refund_no_"

	RawRefundChannelLike = "refund_channel_"

	RawRefundFeeLike = "refund_fee_"

	RawSettlementRefundFeeLike = "settlement_refund_fee_"

	RawRefundStatusLike = "refund_status_"

	RawRefundRecvAccount = "refund_recv_accout_"

	//ErrCode For wx
	RawSystemError   = "SYSTEMERROR"
	RawOrderNotExist = "ORDERNOTEXIST"
	RawBankError     = "BANKERROR"
	RawUserPaying    = "USERPAYING"

	//refund
	//RefundFee   = "refund_fee"
	//OutRefundNo = "out_refund_no"
)

const (
	//common param
	AppId          = "AppId"
	SpbillCreateIp = "SpbillCreateIp"
	MchId          = "MchId"
	SubAppId       = "SubAppId"
	SubMchId       = "SubMchId"
	Key            = "Key"
	OutTradeNo     = "OutTradeNo"

	//direct pay
	Body       = "Body"
	TotalFee   = "TotalFee"
	AuthCode   = "AuthCode"
	DeviceInfo = "DeviceInfo"

	Detail   = "Detail"
	Attach   = "Attach"
	FeeType  = "FeeType"
	GoodsTag = "GoodsTag"

	LimitPay = "LimitPay"

	OpUserId = "OpUserId"

	//refund
	TransactionId = "TransactionId"
	OutRefundNo   = "OutRefundNo"
	RefundId      = "RefundId"
	RefundFee     = "RefundFee"
	RefundFeeType = "RefundFeeType"

	CertName = "CertName"
	CertKey  = "CertKey"
	RootCa   = "RootCa"
)
