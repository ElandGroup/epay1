package wx

import (
	"epaygo/helper"
	"epaygo/helper/cryptoHelper"
	"epaygo/wx/wxConst"
	"errors"
	"net/http"
	"strings"
	"time"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/smallnest/goreq"
)

type WxPayService struct {
}

func (a *WxPayService) DirectPay(params map[string]string) (result string, err error) {

	wxPayData := a.BuildCommonparam(params)

	a.SetValue(wxPayData, wxConst.Body, params[wxConst.BodyMap])
	a.SetValue(wxPayData, wxConst.OutTradeNo, params[wxConst.OutTradeNoMap])
	a.SetValue(wxPayData, wxConst.TotalFee, params[wxConst.TotalFeeMap])
	a.SetValue(wxPayData, wxConst.AuthCode, params[wxConst.AuthCodeMap])
	a.SetValue(wxPayData, wxConst.DeviceInfo, params[wxConst.DeviceInfoMap])

	a.SetValue(wxPayData, wxConst.Detail, params[wxConst.DetailMap])
	a.SetValue(wxPayData, wxConst.Attach, params[wxConst.AttachMap])
	a.SetValue(wxPayData, wxConst.FeeType, params[wxConst.FeeTypeMap])
	a.SetValue(wxPayData, wxConst.GoodsTag, params[wxConst.GoodsTagMap])
	a.SetValue(wxPayData, wxConst.LimitPay, params[wxConst.LimitPayMap])

	a.SetValue(wxPayData, wxConst.Sign, wxPayData.MakeSign(params[wxConst.KeyMap]))

	xmlParam := wxPayData.ToXml()
	req, body, reqErr := goreq.New().Post(wxConst.MicroPay_Url).ContentType("xml").SendRawString(xmlParam).End()

	return a.ParseResult(req, body, reqErr, params[wxConst.KeyMap])

}

func (a *WxPayService) Refund(params map[string]string) (result string, err error) {
	wxPayData := a.BuildCommonparam(params)

	wxPayData.RemoveKey(wxConst.SpbillCreateIp)
	a.SetValue(wxPayData, wxConst.DeviceInfo, params[wxConst.DeviceInfoMap])
	a.SetValue(wxPayData, wxConst.TransactionId, params[wxConst.TransactionIdMap])
	a.SetValue(wxPayData, wxConst.OutRefundNo, params[wxConst.OutRefundNoMap])
	a.SetValue(wxPayData, wxConst.OutTradeNo, params[wxConst.OutTradeNoMap])
	a.SetValue(wxPayData, wxConst.RefundId, params[wxConst.RefundIdMap])

	a.SetValue(wxPayData, wxConst.TotalFee, params[wxConst.TotalFeeMap])
	a.SetValue(wxPayData, wxConst.RefundFee, params[wxConst.RefundFeeMap])
	a.SetValue(wxPayData, wxConst.RefundFeeType, params[wxConst.RefundFeeTypeMap])
	a.SetValue(wxPayData, wxConst.OpUserId, params[wxConst.OpUserIdMap])

	a.SetValue(wxPayData, wxConst.Sign, wxPayData.MakeSign(params[wxConst.KeyMap]))

	xmlParam := wxPayData.ToXml()
	reqNew := goreq.New()

	certName := params[wxConst.CertNameMap]
	certKey := params[wxConst.CertKeyMap]
	rootCa := params[wxConst.RootCaMap]
	if transport, e := cryptoHelper.CertTransport(&certName, &certKey, &rootCa); e == nil {

		reqNew.Transport = transport
		reqNew.Client = &http.Client{Transport: transport}
	} else {
		return "", errors.New("cert error:" + e.Error())

	}

	req, body, reqErr := reqNew.Post(wxConst.Refund_Url).ContentType("xml").SendRawString(xmlParam).End()

	return a.ParseResult(req, body, reqErr, params[wxConst.KeyMap])

}

func (a *WxPayService) OrderQuery(params map[string]string) (result string, err error) {

	wxPayData := a.BuildCommonparam(params)

	a.SetValue(wxPayData, wxConst.TransactionId, params[wxConst.TransactionIdMap])
	a.SetValue(wxPayData, wxConst.OutTradeNo, params[wxConst.OutTradeNoMap])

	a.SetValue(wxPayData, wxConst.Sign, wxPayData.MakeSign(params[wxConst.KeyMap]))

	xmlParam := wxPayData.ToXml()
	req, body, reqErr := goreq.New().Post(wxConst.OrderQuery_Url).ContentType("xml").SendRawString(xmlParam).End()

	return a.ParseResult(req, body, reqErr, params[wxConst.KeyMap])

}

func (a *WxPayService) OrderReverse(params map[string]string, count int) (result string, err error) {
	if count <= 0 {
		return "", errors.New("10005")
	}
	wxPayData := a.BuildCommonparam(params)
	wxPayData.RemoveKey(wxConst.SpbillCreateIp)
	a.SetValue(wxPayData, wxConst.TransactionId, params[wxConst.TransactionIdMap])
	a.SetValue(wxPayData, wxConst.OutTradeNo, params[wxConst.OutTradeNoMap])

	a.SetValue(wxPayData, wxConst.Sign, wxPayData.MakeSign(params[wxConst.KeyMap]))

	xmlParam := wxPayData.ToXml()
	reqNew := goreq.New()
	certName := params[wxConst.CertNameMap]
	certKey := params[wxConst.CertKeyMap]
	rootCa := params[wxConst.RootCaMap]
	if transport, e := cryptoHelper.CertTransport(&certName, &certKey, &rootCa); e == nil {
		reqNew.Transport = transport
		reqNew.Client = &http.Client{Transport: transport}
	} else {
		return "", errors.New("cert error:" + e.Error())
	}

	if req, body, reqErr := reqNew.Post(wxConst.Reverse_Url).ContentType("xml").SendRawString(xmlParam).End(); reqErr != nil {
		return "", reqErr[0]
	} else {

		if result, e := a.ParseResult(req, body, reqErr, params[wxConst.KeyMap]); e == nil {
			return result, nil
		} else {
			if len(result) == 0 {
				return "", e
			}
			rJson, _ := simplejson.NewJson([]byte(result))

			if recall, _ := rJson.Get(wxConst.Recall).String(); recall == "Y" {
				time.Sleep(10000 * time.Millisecond) //10s
				count = count - 1
				return a.OrderReverse(params, count)
			} else {
				if v, e := rJson.Get(wxConst.ErrCode).String(); e != nil {
					return "", errors.New("10007") //no data
				} else {
					return "", errors.New(v)
				}
			}

		}

	}

}

func (a *WxPayService) BuildCommonparam(params map[string]string) WxPayData {
	wxPayData := NewWxPayData()
	a.SetValue(*wxPayData, wxConst.SpbillCreateIp, params[wxConst.SpbillCreateIpMap])
	a.SetValue(*wxPayData, wxConst.AppId, params[wxConst.AppIdMap])
	a.SetValue(*wxPayData, wxConst.MchId, params[wxConst.MchIdMap])
	a.SetValue(*wxPayData, wxConst.SubAppId, params[wxConst.SubAppIdMap])
	a.SetValue(*wxPayData, wxConst.SubMchId, params[wxConst.SubMchIdMap])

	a.SetValue(*wxPayData, wxConst.NonceStr, helper.Uuid32())
	return *wxPayData
}

func (a *WxPayService) SetValue(wxPayData WxPayData, key string, value string) {
	if len(strings.TrimSpace(value)) != 0 {
		wxPayData.SetValue(key, value)
	}
}

func (a *WxPayService) ParseResult(req goreq.Response, body string, reqErrs []error, key string) (result string, err error) {
	//serviceResult := ServiceResult{Result: nil, Success: ResultType.Unknown, Error: APIError{Code: 10004, Message: "", Details: nil}}
	wxResponse := NewWxPayData()
	if err != nil {
		return "", reqErrs[0]
	}
	if req.StatusCode == http.StatusOK {
		if _, err := wxResponse.FromXml(body, key); err != nil {
			return "", errors.New("The request failed, please check whether the network is normal")
		}

		if wxResponse == nil {
			return "", errors.New("The request failed, please check whether the network is normal")
		} else {
			if len(wxResponse.GetValue(wxConst.ReturnCode)) == 0 || strings.ToUpper(wxResponse.GetValue(wxConst.ReturnCode)) != "SUCCESS" {
				return wxResponse.ToJson(), errors.New("The request failed, please check whether the network is normal")
			}
			if len(wxResponse.GetValue(wxConst.ResultCode)) != 0 && strings.ToUpper(wxResponse.GetValue(wxConst.ResultCode)) == "SUCCESS" {
				return wxResponse.ToJson(), nil
			} else {
				errCode := wxResponse.GetValue(wxConst.ErrCode)
				if errCode == wxConst.SystemError || errCode == wxConst.BankError || errCode == wxConst.UserPaying {
					return wxResponse.ToJson(), errors.New("result is unknown")
				} else {
					return wxResponse.ToJson(), errors.New(errCode)
				}
			}
		}
	} else {
		return "", errors.New("The request failed, please check whether the network is normal")
	}
	return "", errors.New("The request failed, please check whether the network is normal")
}
