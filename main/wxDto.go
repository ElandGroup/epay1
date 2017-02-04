package main

type WxPayBaseDto struct {
	AppId          string
	MchId          string
	SubMchId       string
	SubAppId       string
	SpbillCreateIp string
	Key            string
	OutTradeNo     string
}

type WxDirectPayDto struct {
	//WxPayBaseDto
	AppId          string
	MchId          string
	SubMchId       string
	SubAppId       string
	SpbillCreateIp string
	Key            string
	OutTradeNo     string

	DeviceInfo string
	Body       string
	Detail     string
	Attach     string
	TotalFee   string
	FeeType    string
	GoodsTag   string
	LimitPay   string
	AuthCode   string
}

type WxRefundDto struct {
	//WxPayBaseDto

	AppId          string
	MchId          string
	SubMchId       string
	SubAppId       string
	SpbillCreateIp string
	Key            string
	OutTradeNo     string

	DeviceInfo    string
	TransactionId string
	RefundId      string
	TotalFee      string
	OutRefundNo   string
	RefundFee     string
	RefundFeeType string
	OpUserId      string
	CertName      string
	CertKey       string
	RootCa        string
}

type WxOrderQueryDto struct {
	//WxPayBaseDto
	AppId          string
	MchId          string
	SubMchId       string
	SubAppId       string
	SpbillCreateIp string
	Key            string
	OutTradeNo     string

	TransactionId string
}

type WxReverseDto struct {
	//WxPayBaseDto
	AppId          string
	MchId          string
	SubMchId       string
	SubAppId       string
	SpbillCreateIp string
	Key            string
	OutTradeNo     string

	TransactionId string
	CertName      string
	CertKey       string
	RootCa        string
}
