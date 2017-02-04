package epaygo

import (
	"encoding/json"
	"epaygo/al"
	"epaygo/helper/cryptoHelper"
	"epaygo/helper/mapHelper"
	"errors"
	"fmt"
	"strings"
	"time"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/smallnest/goreq"
)

type AlPayService struct {
}

func (a *AlPayService) DirectPay(params map[string]string) (result string, err error) {
	payData := *a.BuildCommonparam(params, al.PayCustom)

	bizContent := make(map[string]string)
	a.SetValue(&bizContent, al.Scence, al.BarCodeScenceCustom)
	a.SetValue(&bizContent, al.OutTradeNo, params[al.OutTradeNoMap])
	a.SetValue(&bizContent, al.AuthCode, params[al.AuthCodeMap])
	a.SetValue(&bizContent, al.TotalAmount, params[al.TotalAmountMap])
	a.SetValue(&bizContent, al.Subject, params[al.SubjectMap])

	a.SetValue(&bizContent, al.StoreId, params[al.StoreIdMap])
	a.SetValue(&bizContent, al.SellerId, params[al.SellerIdMap])
	a.SetValue(&bizContent, al.TimeExpire, params[al.TimeExpireMap])
	a.SetValue(&bizContent, al.ExtendParams, params[al.ExtendParamsMap])

	b, _ := json.Marshal(bizContent)
	payData[al.BizContent] = string(b)

	payData[al.Sign], _ = cryptoHelper.GetSha1Hash(mapHelper.SortedUrl(&payData), params[al.SellerPrivateKey])

	p, _ := json.Marshal(payData)

	if _, body, reqErr := goreq.New().Get(al.OpenApi).Query(string(p)).End(); len(reqErr) != 0 {
		return "", fmt.Errorf("Network error", reqErr[0])
	} else {
		return a.ParseResponse(body, params[al.AliPublicKey], al.PayNode)
	}

}

func (a *AlPayService) OrderQuery(params map[string]string) (result string, err error) {
	payData := *a.BuildCommonparam(params, al.QueryCustom)

	bizContent := make(map[string]string)
	a.SetValue(&bizContent, al.OutTradeNo, params[al.OutTradeNoMap])
	a.SetValue(&bizContent, al.TradeNo, params[al.TradeNoMap])

	b, _ := json.Marshal(bizContent)
	payData[al.BizContent] = string(b)
	payData[al.Sign], _ = cryptoHelper.GetSha1Hash(mapHelper.SortedUrl(&payData), params[al.SellerPrivateKeyMap])

	p, _ := json.Marshal(payData)

	if _, body, reqErr := goreq.New().Get(al.OpenApi).Query(string(p)).End(); len(reqErr) != 0 {
		return "", fmt.Errorf("Network error", reqErr[0])
	} else {
		return a.ParseResponse(body, params[al.AliPublicKeyMap], al.QueryNode)
	}
}

func (a *AlPayService) Refund(params map[string]string) (result string, err error) {
	payData := *a.BuildCommonparam(params, al.RefundCustom)

	bizContent := make(map[string]string)
	//a.SetValue(&bizContent, alConst.TradeNo, dto.TradeNo)
	a.SetValue(&bizContent, al.OutTradeNo, params[al.OutTradeNoMap])
	a.SetValue(&bizContent, al.RefundAmount, params[al.RefundAmountMap])
	a.SetValue(&bizContent, al.OutRequestNo, params[al.OutRequestNoMap])
	a.SetValue(&bizContent, al.RefundReason, params[al.RefundReasonMap])
	a.SetValue(&bizContent, al.StoreId, params[al.StoreIdMap])

	b, _ := json.Marshal(bizContent)
	payData[al.BizContent] = string(b)

	payData[al.Sign], _ = cryptoHelper.GetSha1Hash(mapHelper.SortedUrl(&payData), params[al.SellerPrivateKeyMap])

	p, _ := json.Marshal(payData)
	if _, body, reqErr := goreq.New().Get(al.OpenApi).Query(string(p)).End(); len(reqErr) != 0 {
		return "", fmt.Errorf("Network error", reqErr[0])
	} else {
		return a.ParseResponse(body, params[al.AliPublicKeyMap], al.RefundNode)
	}
}

func (a *AlPayService) Reverse(params map[string]string, count int) (result string, err error) {
	if count <= 0 {
		return "", errors.New("Numbers cannot be less than zero")
	}

	payData := *a.BuildCommonparam(params, al.ReverseCustom)

	bizContent := make(map[string]string)
	a.SetValue(&bizContent, al.OutTradeNo, params[al.OutTradeNoMap])
	a.SetValue(&bizContent, al.TradeNo, params[al.TradeNoMap])
	b, _ := json.Marshal(bizContent)
	payData[al.BizContent] = string(b)

	payData[al.Sign], _ = cryptoHelper.GetSha1Hash(mapHelper.SortedUrl(&payData), params[al.SellerPrivateKeyMap])

	p, _ := json.Marshal(payData)

	if _, body, reqErr := goreq.New().Get(al.OpenApi).Query(string(p)).End(); len(reqErr) != 0 {
		return "", fmt.Errorf("Network error", reqErr[0])
	} else {

		if result, e := a.ParseResponse(body, params[al.AliPublicKeyMap], al.ReverseNode); len(result) != 0 && e == nil {
			return result, nil
		} else {
			if len(result) == 0 {
				return "", e
			}
			rJson, _ := simplejson.NewJson([]byte(result))

			if recall, _ := rJson.Get(al.RetryFlag).String(); recall == "Y" {
				time.Sleep(10000 * time.Millisecond) //10s
				count = count - 1
				return a.Reverse(params, count)
			} else {
				if v, e := rJson.Get(al.Code).String(); e != nil {
					return "", errors.New("No data") //no data
				} else {
					return "", errors.New(v)
				}
			}

		}

	}

}

func (a *AlPayService) BuildCommonparam(commonParams map[string]string, method string) *map[string]string {
	payData := make(map[string]string)
	a.SetValue(&payData, al.AppId, commonParams[al.AppIdMap])
	a.SetValue(&payData, al.Charset, al.CharsetCustom)
	a.SetValue(&payData, al.Method, method)
	a.SetValue(&payData, al.SignType, al.SignTypeCustom)
	t := time.Now()
	a.SetValue(&payData, al.TimeStamp, fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second()))

	a.SetValue(&payData, al.Version, al.VersionCustom)
	a.SetValue(&payData, al.SignType, al.SignTypeCustom)

	return &payData
}

func (a *AlPayService) SetValue(mapData *map[string]string, key string, value string) {
	if len(strings.TrimSpace(value)) != 0 {
		(*mapData)[key] = value
	}
}

func (a *AlPayService) ParseResponse(body string, pubKey string, repType string) (result string, err error) {

	if js, err := simplejson.NewJson([]byte(body)); err != nil {
		return "", fmt.Errorf("parse response error", err.Error())
	} else {
		jsDetail := js.Get(repType)
		bodyMap, err := jsDetail.Map()
		if err != nil {
			return "", fmt.Errorf("parse response error", err.Error())
		}
		bodyArray, err := json.Marshal(bodyMap)
		if err != nil {
			return "", fmt.Errorf("parse response error", err.Error())
		}
		bodyJs := string(bodyArray)
		sign, err := js.Get(al.Sign).String()
		if err != nil {
			return "", fmt.Errorf("parse response error", err.Error())
		}
		if isValid := cryptoHelper.CheckPubKey(bodyJs, sign, pubKey); isValid {
			if code, _ := jsDetail.Get(al.Code).String(); code == "10000" {
				return bodyJs, nil
			} else if code == "10003" {
				return "", fmt.Errorf("result is unknown", err.Error())
			} else {
				subCode, _ := jsDetail.Get(al.SubCode).String()
				return "", errors.New(subCode)
			}

		} else {
			return "", errors.New("response signature failed")
		}

	}
}
