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
	AppId          = "appid"
	MchId          = "mch_id"
	NonceStr       = "nonce_str"
	SubMchId       = "sub_mch_id"
	SubAppId       = "sub_appid"
	SpbillCreateIp = "spbill_create_ip" //此参数可手动配置也可在程序中自动获取
	DeviceInfo     = "device_info"

	Sign       = "sign"
	Body       = "body"
	Detail     = "detail"
	Attach     = "attach"
	OutTradeNo = "out_trade_no"
	FeeType    = "fee_type"

	TotalFee = "total_fee"
	GoodsTag = "goods_tag"
	LimitPay = "limit_pay"

	//For WxDirectPayDto
	AuthCode = "auth_code"
	//For Sign
	Key = "key"

	//For WxPayPrecreateDto
	TimeStart  = "time_start"
	TimeExpire = "time_expire"
	NotifyUrl  = "notify_url"
	TradeType  = "trade_type"
	ProductId  = "product_id"
	OpenId     = "openid"

	//For OrderQuery
	TransactionId = "transaction_id"

	//For Refund
	/// <summary>
	/// 商户系统内部的退款单号，商户系统内部唯一，同一退款单号多次请求只退一笔
	/// </summary>
	OutRefundNo = "out_refund_no"
	/// <summary>
	/// 退款总金额，订单总金额，单位为分，只能为整数，详见支付金额
	/// </summary>
	RefundFee = "refund_fee"
	/// <summary>
	/// 货币类型，符合ISO 4217标准的三位字母代码，默认人民币：CNY，其他值列表详见货币类型
	/// </summary>
	RefundFeeType = "refund_fee_type"
	/// <summary>
	/// 操作员帐号, 默认为商户号
	/// </summary>
	OpUserId = "op_user_id"
	/// <summary>
	/// 微信生成的退款单号，在申请退款接口有返回
	/// </summary>
	RefundId = "refund_id"

	//For DownloadBill
	/// <summary>
	/// 下载对账单的日期，格式：20140603
	/// </summary>
	BillDate = "bill_date"
	/// <summary>
	/// ALL，返回当日所有订单信息，默认值,
	/// SUCCESS，返回当日成功支付的订单,
	/// REFUND，返回当日退款订单,
	/// REVOKED，已撤销的订单
	/// </summary>
	BillType = "bill_type"

	//For QRCode
	/// <summary>
	/// 时间戳
	/// </summary>
	TimeStamp = "time_stamp"

	//For ShortUrl
	/// <summary>
	/// 需要转换的URL，签名用原串，传输需URLencode
	/// </summary>
	LongUrl = "long_url"

	SpbillCreateIpValue = "8.8.8.8"
)

//const For response wechat
const (
	ReturnCode = "return_code"
	ReturnMsg  = "return_msg"
	ResultCode = "result_code"
	ErrCode    = "err_code"
	ErrCodeDes = "err_code_des"

	TradeState = "trade_state"

	TradeStateDesc = "trade_state_desc"

	Recall = "recall"

	TimeEnd = "time_end"

	CodeUrl = "code_url"

	PrepayId = "prepay_id"

	//TradeType = "trade_type"

	BankType = "bank_type"

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

	RefundCount = "refund_count"

	outRefundNoLike = "out_refund_no_"

	RefundChannelLike = "refund_channel_"

	RefundFeeLike = "refund_fee_"

	SettlementRefundFeeLike = "settlement_refund_fee_"

	RefundStatusLike = "refund_status_"

	RefundRecvAccount = "refund_recv_accout_"

	//ErrCode For wx
	SystemError   = "SYSTEMERROR"
	OrderNotExist = "ORDERNOTEXIST"
	BankError     = "BANKERROR"
	UserPaying    = "USERPAYING"

	//refund
	//RefundFee   = "refund_fee"
	//OutRefundNo = "out_refund_no"
)

const (
	//common param
	AppIdMap          = "AppId"
	SpbillCreateIpMap = "SpbillCreateIp"
	MchIdMap          = "MchId"
	SubAppIdMap       = "SubAppId"
	SubMchIdMap       = "SubMchId"
	KeyMap            = "Key"
	OutTradeNoMap     = "OutTradeNo"

	//direct pay
	BodyMap       = "Body"
	TotalFeeMap   = "TotalFee"
	AuthCodeMap   = "AuthCode"
	DeviceInfoMap = "DeviceInfo"

	DetailMap   = "Detail"
	AttachMap   = "Attach"
	FeeTypeMap  = "FeeType"
	GoodsTagMap = "GoodsTag"

	LimitPayMap = "LimitPay"

	OpUserIdMap = "OpUserId"

	//refund
	TransactionIdMap = "TransactionId"
	OutRefundNoMap   = "OutRefundNo"
	RefundIdMap      = "RefundId"
	RefundFeeMap     = "RefundFee"
	RefundFeeTypeMap = "RefundFeeType"

	CertNameMap = "CertName"
	CertKeyMap  = "CertKey"
	RootCaMap   = "RootCa"
)
