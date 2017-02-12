package al

import (
	"encoding/json"
	"epaygo/core/alConst"
	"epaygo/core/common"
	. "epaygo/core/commonDto"
	"epaygo/helper/cryptoHelper"
	"epaygo/helper/mapHelper"
	"fmt"
	"strconv"
	"strings"
	"time"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/smallnest/goreq"
)

type AlPayService struct {
}

func (a *AlPayService) DirectPay(params map[string]string) (result string, apiError *APIError) {
	payData := *a.BuildCommonparam(params, alConst.ReqPay)

	bizContent := make(map[string]string)
	a.SetValue(&bizContent, alConst.RawScence, alConst.ReqScenceBarCode)
	a.SetValue(&bizContent, alConst.RawOutTradeNo, params[alConst.OutTradeNo])
	a.SetValue(&bizContent, alConst.RawAuthCode, params[alConst.AuthCode])
	a.SetValue(&bizContent, alConst.RawTotalAmount, params[alConst.TotalAmount])
	a.SetValue(&bizContent, alConst.RawSubject, params[alConst.Subject])

	a.SetValue(&bizContent, alConst.RawStoreId, params[alConst.StoreId])
	a.SetValue(&bizContent, alConst.RawSellerId, params[alConst.SellerId])
	a.SetValue(&bizContent, alConst.RawTimeExpire, params[alConst.TimeExpire])
	a.SetValue(&bizContent, alConst.RawExtendParams, params[alConst.ExtendParams])
	a.SetValue(&bizContent, alConst.RawALAuthToken, params[alConst.ALAuthToken])

	b, _ := json.Marshal(bizContent)
	payData[alConst.RawBizContent] = string(b)

	payData[alConst.RawSign], _ = cryptoHelper.GetSha1Hash(mapHelper.SortedUrl(&payData), params[alConst.SellerPrivateKey])

	p, _ := json.Marshal(payData)

	if _, body, reqErr := goreq.New().Get(alConst.OpenApi).Query(string(p)).End(); len(reqErr) != 0 {
		result = ""
		commonError := "payType:AL,method:" + alConst.ReqPay
		apiError = &APIError{Code: 20001, Message: common.RequestError, Details: common.ResourceMessage(reqErr[0].Error(), commonError)}
		return
	} else {
		return a.ParseResponse(body, params[alConst.AliPublicKey], alConst.RespPay)
	}

}

func (a *AlPayService) OrderQuery(params map[string]string) (result string, apiError *APIError) {
	payData := *a.BuildCommonparam(params, alConst.ReqQuery)

	bizContent := make(map[string]string)
	a.SetValue(&bizContent, alConst.RawOutTradeNo, params[alConst.OutTradeNo])
	a.SetValue(&bizContent, alConst.RawTradeNo, params[alConst.TradeNo])
	a.SetValue(&bizContent, alConst.RawExtendParams, params[alConst.ExtendParams])

	b, _ := json.Marshal(bizContent)
	payData[alConst.RawBizContent] = string(b)
	payData[alConst.RawSign], _ = cryptoHelper.GetSha1Hash(mapHelper.SortedUrl(&payData), params[alConst.SellerPrivateKey])

	p, _ := json.Marshal(payData)
	if _, body, reqErr := goreq.New().Get(alConst.OpenApi).Query(string(p)).End(); len(reqErr) != 0 {
		result = ""
		commonError := "payType:AL,method:" + alConst.ReqQuery
		apiError = &APIError{Code: 20001, Message: common.RequestError, Details: common.ResourceMessage(reqErr[0].Error(), commonError)}
		return
	} else {
		return a.ParseResponse(body, params[alConst.AliPublicKey], alConst.RespQuery)
	}
}

func (a *AlPayService) Refund(params map[string]string) (result string, apiError *APIError) {
	payData := *a.BuildCommonparam(params, alConst.ReqRefund)

	bizContent := make(map[string]string)
	//a.SetValue(&bizContent, alConst.TradeNo, dto.TradeNo)
	a.SetValue(&bizContent, alConst.RawOutTradeNo, params[alConst.OutTradeNo])
	a.SetValue(&bizContent, alConst.RawRefundAmount, params[alConst.RefundAmount])
	a.SetValue(&bizContent, alConst.RawOutRequestNo, params[alConst.OutRequestNo])
	a.SetValue(&bizContent, alConst.RawRefundReason, params[alConst.RefundReason])
	a.SetValue(&bizContent, alConst.RawStoreId, params[alConst.StoreId])
	a.SetValue(&bizContent, alConst.RawAuthCode, params[alConst.AuthCode])
	a.SetValue(&bizContent, alConst.RawExtendParams, params[alConst.ExtendParams])
	b, _ := json.Marshal(bizContent)
	payData[alConst.RawBizContent] = string(b)

	payData[alConst.RawSign], _ = cryptoHelper.GetSha1Hash(mapHelper.SortedUrl(&payData), params[alConst.SellerPrivateKey])

	p, _ := json.Marshal(payData)
	if _, body, reqErr := goreq.New().Get(alConst.OpenApi).Query(string(p)).End(); len(reqErr) != 0 {
		result = ""
		commonError := "payType:AL,method:" + alConst.ReqRefund
		apiError = &APIError{Code: 20001, Message: common.RequestError, Details: common.ResourceMessage(reqErr[0].Error(), commonError)}
		return
	} else {
		return a.ParseResponse(body, params[alConst.AliPublicKey], alConst.RespRefund)
	}
}

func (a *AlPayService) OrderReverse(params map[string]string, count int) (result string, apiError *APIError) {
	commonError := "payType:AL,method:" + alConst.ReqReverse
	if count <= 0 {
		result = ""
		apiError = &APIError{Code: 20001, Message: common.RequestError, Details: common.ResourceMessage("request count:"+strconv.Itoa(count), commonError)}
		return
	}

	payData := *a.BuildCommonparam(params, alConst.ReqReverse)

	bizContent := make(map[string]string)
	a.SetValue(&bizContent, alConst.RawOutTradeNo, params[alConst.OutTradeNo])
	a.SetValue(&bizContent, alConst.RawTradeNo, params[alConst.TradeNo])
	a.SetValue(&bizContent, alConst.RawAuthCode, params[alConst.AuthCode])
	a.SetValue(&bizContent, alConst.RawExtendParams, params[alConst.ExtendParams])
	b, _ := json.Marshal(bizContent)
	payData[alConst.RawBizContent] = string(b)

	payData[alConst.RawSign], _ = cryptoHelper.GetSha1Hash(mapHelper.SortedUrl(&payData), params[alConst.SellerPrivateKey])

	p, _ := json.Marshal(payData)

	if _, body, reqErr := goreq.New().Get(alConst.OpenApi).Query(string(p)).End(); len(reqErr) != 0 {
		result = ""
		apiError = &APIError{Code: 20001, Message: common.RequestError, Details: common.ResourceMessage(reqErr[0].Error(), commonError)}
		return
	} else {

		if responseResult, e := a.ParseResponse(body, params[alConst.AliPublicKey], alConst.RespReverse); e == nil {
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
			if recall, err = messgeJson.Get(alConst.RawRetryFlag).String(); err != nil {
				result = ""
				apiError = &APIError{Code: 20001, Message: common.ResponseParseError, Details: common.ResourceMessage(err.Error(), commonError)}
				return
			} else if recall == "Y" {
				time.Sleep(10000 * time.Millisecond) //10s
				count = count - 1
				return a.OrderReverse(params, count)
			} else {
				if v, e := messgeJson.Get(alConst.RawCode).String(); e != nil {
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

func (a *AlPayService) BuildCommonparam(commonParams map[string]string, method string) *map[string]string {
	payData := make(map[string]string)
	a.SetValue(&payData, alConst.RawAppId, commonParams[alConst.AppId])
	a.SetValue(&payData, alConst.RawCharset, alConst.ReqCharset)
	a.SetValue(&payData, alConst.RawMethod, method)
	a.SetValue(&payData, alConst.RawSignType, alConst.ReqSignType)
	t := time.Now()
	a.SetValue(&payData, alConst.RawTimeStamp, fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second()))

	a.SetValue(&payData, alConst.RawVersion, alConst.ReqVersion)

	return &payData
}

func (a *AlPayService) SetValue(mapData *map[string]string, key string, value string) {
	if len(strings.TrimSpace(value)) != 0 {
		(*mapData)[key] = value
	}
}

func (a *AlPayService) ParseResponse(body string, pubKey string, repType string) (result string, apiError *APIError) {

	commonError := "payType:AL,method:" + repType
	if js, err := simplejson.NewJson([]byte(body)); err != nil {
		result = ""
		apiError = &APIError{Code: 20001, Message: common.ResponseParseError, Details: common.ResourceMessage(err.Error(), commonError)}
		return
	} else {
		jsDetail := js.Get(repType)
		body, err := jsDetail.Map()
		if err != nil {
			result = ""
			apiError = &APIError{Code: 20001, Message: common.ResponseParseError, Details: common.ResourceMessage(err.Error(), commonError)}
			return
		}
		bodyArray, err := json.Marshal(body)
		if err != nil {
			return "", nil
		}
		bodyJs := string(bodyArray)
		sign, err := js.Get(alConst.RawSign).String()
		if err != nil {
			result = ""
			apiError = &APIError{Code: 20001, Message: common.ResponseParseError, Details: common.ResourceMessage(err.Error(), commonError)}
			return
		}
		if isValid := cryptoHelper.CheckPubKey(bodyJs, sign, pubKey); isValid {
			if code, err := jsDetail.Get(alConst.RawCode).String(); err != nil {
				result = ""
				apiError = &APIError{Code: 20001, Message: common.ResponseParseError, Details: common.ResourceMessage(err.Error(), commonError)}
				return
			} else if code == "10000" {
				return bodyJs, nil
			} else if code == "10003" {
				result = ""
				apiError = &APIError{Code: 10001, Message: common.SystemError, Details: common.ResourceMessage(err.Error(), commonError)}
				return
			} else {
				if subCode, err := jsDetail.Get(alConst.RawSubCode).String(); err != nil {
					result = ""
					apiError = &APIError{Code: 20001, Message: common.ResponseParseError, Details: common.ResourceMessage(err.Error(), commonError)}
				} else {
					result = ""
					apiError = &APIError{Code: 20005, Message: common.ResponseError, Details: common.ResourceMessage(subCode, commonError)}
				}
				return
			}

		} else {
			result = ""
			apiError = &APIError{Code: 20003, Message: common.ResponseSignError, Details: common.ResourceMessage(err.Error(), commonError)}
			return
		}

	}
}
