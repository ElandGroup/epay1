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
	10001: "System error",
	10012: "A required parameter is missing or doesn't have the right format:%v",
	20001: "%v is required",
	20002: "PC is under construction...",
	20003: "The correct express number should be:%v",
	20004: "%v: Not Exist",
	20005: "Express has been sent",
	20006: "Order has been grabed",

	20007: "No row is affected",
	20008: "The parameter format should be %v",
	20009: "The order does not exist or has been canceled",
	20010: "Please confirm the parameter name, including uppercase and lowercase letters",
	20011: "Please enter a date of %v digits",
	20012: "The value of %v is incorrect",
	20013: "%v has a wrong format"}
