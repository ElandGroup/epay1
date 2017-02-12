package alDto

type AlResponseDto struct {
	AlCommonDto
	TradeNo    string `json:"trade_no"`
	OutTradeNo string `json:"out_trade_no"`
	RetryFlag  string `json:"retry_flag"`
	Action     string `json:"action"`
}

type AlCommonDto struct {
	Code   string `json:code`
	Msg    string `json:msg`
	SubMsg string `json:sub_msg`
	Sign   string `json:sign`
}
