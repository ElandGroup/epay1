package wx

import (
	"epaygo/core/common"
	. "epaygo/core/commonDto"
	"epaygo/core/wxConst"
	"epaygo/helper"
	"epaygo/helper/cryptoHelper"
	"net/http"
	"strconv"
	"strings"
	"time"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/smallnest/goreq"
)

type WxPayService struct {
}

func (a *WxPayService) DirectPay(params map[string]string) (result string, apiError *APIError) {

	wxPayData := a.BuildCommonparam(params)

	a.SetValue(wxPayData, wxConst.RawBody, params[wxConst.Body])
	a.SetValue(wxPayData, wxConst.RawOutTradeNo, params[wxConst.OutTradeNo])
	a.SetValue(wxPayData, wxConst.RawTotalFee, params[wxConst.TotalFee])
	a.SetValue(wxPayData, wxConst.RawAuthCode, params[wxConst.AuthCode])
	a.SetValue(wxPayData, wxConst.RawDeviceInfo, params[wxConst.DeviceInfo])

	a.SetValue(wxPayData, wxConst.RawDetail, params[wxConst.Detail])
	a.SetValue(wxPayData, wxConst.RawAttach, params[wxConst.Attach])
	a.SetValue(wxPayData, wxConst.RawFeeType, params[wxConst.FeeType])
	a.SetValue(wxPayData, wxConst.RawGoodsTag, params[wxConst.GoodsTag])
	a.SetValue(wxPayData, wxConst.RawLimitPay, params[wxConst.LimitPay])

	a.SetValue(wxPayData, wxConst.RawSign, wxPayData.MakeSign(params[wxConst.Key]))

	xmlParam := wxPayData.ToXml()
	req, body, reqErr := goreq.New().Post(wxConst.MicroPay_Url).ContentType("xml").SendRawString(xmlParam).End()

	return a.ParseResult(req, body, reqErr, params[wxConst.Key], common.Pay)

}

func (a *WxPayService) Refund(params map[string]string) (result string, apiError *APIError) {
	wxPayData := a.BuildCommonparam(params)

	wxPayData.RemoveKey(wxConst.RawSpbillCreateIp)
	a.SetValue(wxPayData, wxConst.RawDeviceInfo, params[wxConst.DeviceInfo])
	a.SetValue(wxPayData, wxConst.RawTransactionId, params[wxConst.TransactionId])
	a.SetValue(wxPayData, wxConst.RawOutRefundNo, params[wxConst.OutRefundNo])
	a.SetValue(wxPayData, wxConst.RawOutTradeNo, params[wxConst.OutTradeNo])
	a.SetValue(wxPayData, wxConst.RawRefundId, params[wxConst.RefundId])

	a.SetValue(wxPayData, wxConst.RawTotalFee, params[wxConst.TotalFee])
	a.SetValue(wxPayData, wxConst.RawRefundFee, params[wxConst.RefundFee])
	a.SetValue(wxPayData, wxConst.RawRefundFeeType, params[wxConst.RefundFeeType])
	a.SetValue(wxPayData, wxConst.RawOpUserId, params[wxConst.OpUserId])

	a.SetValue(wxPayData, wxConst.RawSign, wxPayData.MakeSign(params[wxConst.Key]))

	xmlParam := wxPayData.ToXml()
	reqNew := goreq.New()

	certName := params[wxConst.CertName]
	certKey := params[wxConst.CertKey]
	rootCa := params[wxConst.RootCa]
	if transport, e := cryptoHelper.CertTransport(&certName, &certKey, &rootCa); e == nil {

		reqNew.Transport = transport
		reqNew.Client = &http.Client{Transport: transport}
	} else {
		commonError := "payType:WX,method:" + common.Refund
		result = ""
		apiError = &APIError{Code: 10014, Message: common.CertificateError, Details: common.ResourceMessage(e.Error(), commonError)}
		return
	}

	req, body, reqErr := reqNew.Post(wxConst.Refund_Url).ContentType("xml").SendRawString(xmlParam).End()

	return a.ParseResult(req, body, reqErr, params[wxConst.Key], common.Refund)

}

func (a *WxPayService) OrderQuery(params map[string]string) (result string, apiError *APIError) {

	wxPayData := a.BuildCommonparam(params)

	a.SetValue(wxPayData, wxConst.RawTransactionId, params[wxConst.TransactionId])
	a.SetValue(wxPayData, wxConst.RawOutTradeNo, params[wxConst.OutTradeNo])

	a.SetValue(wxPayData, wxConst.RawSign, wxPayData.MakeSign(params[wxConst.Key]))

	xmlParam := wxPayData.ToXml()
	req, body, reqErr := goreq.New().Post(wxConst.OrderQuery_Url).ContentType("xml").SendRawString(xmlParam).End()

	return a.ParseResult(req, body, reqErr, params[wxConst.Key], common.Query)

}

func (a *WxPayService) OrderReverse(params map[string]string, count int) (result string, apiError *APIError) {
	commonError := "payType:WX,method:" + common.Reverse
	if count <= 0 {
		result = ""
		apiError = &APIError{Code: 20001, Message: common.RequestError, Details: common.ResourceMessage("request count:"+strconv.Itoa(count), commonError)}
		return
	}
	wxPayData := a.BuildCommonparam(params)
	wxPayData.RemoveKey(wxConst.RawSpbillCreateIp)
	a.SetValue(wxPayData, wxConst.RawTransactionId, params[wxConst.TransactionId])
	a.SetValue(wxPayData, wxConst.RawOutTradeNo, params[wxConst.OutTradeNo])

	a.SetValue(wxPayData, wxConst.RawSign, wxPayData.MakeSign(params[wxConst.Key]))

	xmlParam := wxPayData.ToXml()
	reqNew := goreq.New()
	certName := params[wxConst.CertName]
	certKey := params[wxConst.CertKey]
	rootCa := params[wxConst.RootCa]
	if transport, e := cryptoHelper.CertTransport(&certName, &certKey, &rootCa); e == nil {
		reqNew.Transport = transport
		reqNew.Client = &http.Client{Transport: transport}
	} else {
		result = ""
		apiError = &APIError{Code: 10014, Message: common.CertificateError, Details: common.ResourceMessage(e.Error(), commonError)}
		return
	}

	if req, body, reqErr := reqNew.Post(wxConst.Reverse_Url).ContentType("xml").SendRawString(xmlParam).End(); reqErr != nil {
		result = ""
		apiError = &APIError{Code: 20001, Message: common.RequestError, Details: common.ResourceMessage(reqErr[0].Error(), commonError)}
		return
	} else {

		if responseResult, e := a.ParseResult(req, body, reqErr, params[wxConst.Key], common.Reverse); e == nil {
			result = responseResult
			apiError = nil
			return
		} else {
			var messgeJson *simplejson.Json
			var err error
			if messgeJson, err = simplejson.NewJson([]byte(responseResult)); err != nil {
				result = ""
				apiError = &APIError{Code: 20001, Message: common.ResponseParseError, Details: common.ResourceMessage(err.Error(), commonError)}
				return
			}
			var recall string
			if recall, err = messgeJson.Get(wxConst.RawRecall).String(); err != nil {
				result = ""
				apiError = &APIError{Code: 20001, Message: common.ResponseParseError, Details: common.ResourceMessage(err.Error(), commonError)}
				return
			} else if recall == "Y" {
				time.Sleep(10000 * time.Millisecond) //10s
				count = count - 1
				return a.OrderReverse(params, count)
			} else {
				if v, e := messgeJson.Get(wxConst.RawErrCode).String(); e != nil {
					result = ""
					apiError = &APIError{Code: 20001, Message: common.ResponseParseError, Details: common.ResourceMessage(err.Error(), commonError)}
					return
				} else {
					result = ""
					apiError = &APIError{Code: 20001, Message: common.ResponseParseError, Details: common.ResourceMessage(v, commonError)}
					return
				}
			}

		}

	}

}

func (a *WxPayService) BuildCommonparam(params map[string]string) WxPayData {
	wxPayData := NewWxPayData()
	a.SetValue(*wxPayData, wxConst.RawSpbillCreateIp, params[wxConst.SpbillCreateIp])
	a.SetValue(*wxPayData, wxConst.RawAppId, params[wxConst.AppId])
	a.SetValue(*wxPayData, wxConst.RawMchId, params[wxConst.MchId])
	a.SetValue(*wxPayData, wxConst.RawSubAppId, params[wxConst.SubAppId])
	a.SetValue(*wxPayData, wxConst.RawSubMchId, params[wxConst.SubMchId])

	a.SetValue(*wxPayData, wxConst.RawNonceStr, helper.UuIdForPay(""))
	return *wxPayData
}

func (a *WxPayService) SetValue(wxPayData WxPayData, key string, value string) {
	if len(strings.TrimSpace(value)) != 0 {
		wxPayData.SetValue(key, value)
	}
}

func (a *WxPayService) ParseResult(req goreq.Response, body string, reqErrs []error, key string, reqType string) (result string, apiError *APIError) {
	//serviceResult := ServiceResult{Result: nil, Success: ResultType.Unknown, Error: APIError{Code: 10004, Message: "", Details: nil}}
	commonError := "payType:WX,method:" + reqType
	wxResponse := NewWxPayData()
	if len(reqErrs) != 0 {
		result = ""
		apiError = &APIError{Code: 20001, Message: common.ResponseParseError, Details: common.ResourceMessage(reqErrs[0].Error(), commonError)}
		return
	}
	if req.StatusCode == http.StatusOK {
		if err := wxResponse.FromXml(body, key); err != nil {
			result = ""
			apiError = &APIError{Code: 20001, Message: common.ResponseParseError, Details: common.ResourceMessage(err.Error(), commonError)}
			return
		}

		if wxResponse == nil {
			result = ""
			apiError = &APIError{Code: 20005, Message: common.ResponseMessage, Details: common.ResourceMessage(reqErrs[0].Error(), commonError)}
			return
		} else {
			if len(wxResponse.GetValue(wxConst.RawReturnCode)) == 0 || strings.ToUpper(wxResponse.GetValue(wxConst.RawReturnCode)) != "SUCCESS" {
				result = ""
				apiError = &APIError{Code: 20005, Message: common.ResponseMessage, Details: common.ResourceMessage(wxResponse.GetValue(wxConst.RawReturnMsg), commonError)}
				return
			}
			if len(wxResponse.GetValue(wxConst.RawResultCode)) != 0 && strings.ToUpper(wxResponse.GetValue(wxConst.RawResultCode)) == "SUCCESS" {
				result = wxResponse.ToJson()
				apiError = nil
				return
			} else {
				errCode := wxResponse.GetValue(wxConst.RawErrCode)
				if errCode == wxConst.RawSystemError || errCode == wxConst.RawBankError || errCode == wxConst.RawUserPaying {
					result = ""
					apiError = &APIError{Code: 10001, Message: common.SystemError, Details: common.ResourceMessage(errCode, commonError)}
					return
				} else {
					result = ""
					apiError = &APIError{Code: 20005, Message: common.ResponseMessage, Details: common.ResourceMessage(errCode, commonError)}
					return
				}
			}
		}
		return
	} else {
		result = ""
		apiError = &APIError{Code: 20005, Message: common.ResponseMessage, Details: common.ResourceMessage(reqErrs[0].Error(), commonError)}
		return
	}
	result = ""
	apiError = &APIError{Code: 20005, Message: common.ResponseMessage, Details: common.ResourceMessage(reqErrs[0].Error(), commonError)}
	return
}
