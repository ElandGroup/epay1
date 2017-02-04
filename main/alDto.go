package main

type AlPayBaseDto struct {
	AppId            string
	SellerPrivateKey string
	AliPublicKey     string
	OutTradeNo       string
	ALAuthToken      string
}

type AlDirectPayDto struct {
	AppId            string
	SellerPrivateKey string
	AliPublicKey     string
	OutTradeNo       string
	ALAuthToken      string
	AuthCode         string
	TotalAmount      string
	Subject          string
	StoreId          string

	SellerId     string
	TimeExpire   string
	ExtendParams string
}

type AlOrderQueryDto struct {
	//AlPayBaseDto `mapstructure:",squash"`
	AppId            string
	SellerPrivateKey string
	AliPublicKey     string
	OutTradeNo       string
	ALAuthToken      string
	AuthCode         string
	TotalAmount      string
	Subject          string
	StoreId          string

	TradeNo string
}

type AlRefundDto struct {
	//AlPayBaseDto
	AppId            string
	SellerPrivateKey string
	AliPublicKey     string
	OutTradeNo       string
	ALAuthToken      string
	AuthCode         string
	TotalAmount      string
	Subject          string
	StoreId          string

	TradeNo      string
	RefundAmount string
	OutRequestNo string
	RefundReason string
}

type AlReverseDto struct {
	//AlPayBaseDto
	AppId            string
	SellerPrivateKey string
	AliPublicKey     string
	OutTradeNo       string
	ALAuthToken      string
	AuthCode         string
	TotalAmount      string
	Subject          string
	StoreId          string
	TradeNo          string
}
