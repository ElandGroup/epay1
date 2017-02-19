package common

//error
const (
	SystemError        = "system error"
	ResponseSignError  = "response signature failed"
	RequestError       = "request error"
	ResponseParseError = "response parse error"
	CertificateError   = "certificate is missing or doesn't right"
	ResponseMessage    = "response messge"
)

const (
	Pay     = "Pay"
	Reverse = "Reverse"
	Refund  = "Refund"
	Query   = "Query"
)

//resource
const ()

var MessageMap = map[int]string{
	10001: "System error:%v",
	10012: "A required parameter is missing or doesn't have the right format:%v",
	20001: "%v is required",
	20004: "%v: Not Exist",
	20007: "No row is affected",
	20008: "The parameter format should be %v",

	20012: "The value of %v is incorrect",
	20013: "%v has a wrong format",
	20014: "Request error, please try again ...",
	20015: "The response format is incorrect",
	20016: "No fields found in response:%v",

	20017: "Message from WeChat:%v",
	20018: "Message from Alipay:%v",
	20019: "pay failure, please try again ...",
	20020: "Signature error"}
