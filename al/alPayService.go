package al

import (
	"encoding/json"
	"epaygo/al/alConst"
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
	payData := *a.BuildCommonparam(params, alConst.PayCustom)

	bizContent := make(map[string]string)
	a.SetValue(&bizContent, alConst.Scence, alConst.BarCodeScenceCustom)
	a.SetValue(&bizContent, alConst.OutTradeNo, params[alConst.OutTradeNoMap])
	a.SetValue(&bizContent, alConst.AuthCode, params[alConst.AuthCodeMap])
	a.SetValue(&bizContent, alConst.TotalAmount, params[alConst.TotalAmountMap])
	a.SetValue(&bizContent, alConst.Subject, params[alConst.SubjectMap])

	a.SetValue(&bizContent, alConst.StoreId, params[alConst.StoreIdMap])
	a.SetValue(&bizContent, alConst.SellerId, params[alConst.SellerIdMap])
	a.SetValue(&bizContent, alConst.TimeExpire, params[alConst.TimeExpireMap])
	a.SetValue(&bizContent, alConst.ExtendParams, params[alConst.ExtendParamsMap])

	b, _ := json.Marshal(bizContent)
	payData[alConst.BizContent] = string(b)

	payData[alConst.Sign], _ = cryptoHelper.GetSha1Hash(mapHelper.SortedUrl(&payData), params[alConst.SellerPrivateKey])

	p, _ := json.Marshal(payData)

	if _, body, reqErr := goreq.New().Get(alConst.OpenApi).Query(string(p)).End(); len(reqErr) != 0 {
		return "", fmt.Errorf("Network error", reqErr[0])
	} else {
		return a.ParseResponse(body, params[alConst.AliPublicKey], alConst.PayNode)
	}

}

func (a *AlPayService) OrderQuery(params map[string]string) (result string, err error) {
	payData := *a.BuildCommonparam(params, alConst.QueryCustom)

	bizContent := make(map[string]string)
	a.SetValue(&bizContent, alConst.OutTradeNo, params[alConst.OutTradeNoMap])
	a.SetValue(&bizContent, alConst.TradeNo, params[alConst.TradeNoMap])

	b, _ := json.Marshal(bizContent)
	payData[alConst.BizContent] = string(b)
	payData[alConst.Sign], _ = cryptoHelper.GetSha1Hash(mapHelper.SortedUrl(&payData), params[alConst.SellerPrivateKeyMap])

	p, _ := json.Marshal(payData)

	if _, body, reqErr := goreq.New().Get(alConst.OpenApi).Query(string(p)).End(); len(reqErr) != 0 {
		return "", fmt.Errorf("Network error", reqErr[0])
	} else {
		return a.ParseResponse(body, params[alConst.AliPublicKeyMap], alConst.QueryNode)
	}
}

func (a *AlPayService) Refund(params map[string]string) (result string, err error) {
	payData := *a.BuildCommonparam(params, alConst.RefundCustom)

	bizContent := make(map[string]string)
	//a.SetValue(&bizContent, alConst.TradeNo, dto.TradeNo)
	a.SetValue(&bizContent, alConst.OutTradeNo, params[alConst.OutTradeNoMap])
	a.SetValue(&bizContent, alConst.RefundAmount, params[alConst.RefundAmountMap])
	a.SetValue(&bizContent, alConst.OutRequestNo, params[alConst.OutRequestNoMap])
	a.SetValue(&bizContent, alConst.RefundReason, params[alConst.RefundReasonMap])
	a.SetValue(&bizContent, alConst.StoreId, params[alConst.StoreIdMap])

	b, _ := json.Marshal(bizContent)
	payData[alConst.BizContent] = string(b)

	payData[alConst.Sign], _ = cryptoHelper.GetSha1Hash(mapHelper.SortedUrl(&payData), params[alConst.SellerPrivateKeyMap])

	p, _ := json.Marshal(payData)
	if _, body, reqErr := goreq.New().Get(alConst.OpenApi).Query(string(p)).End(); len(reqErr) != 0 {
		return "", fmt.Errorf("Network error", reqErr[0])
	} else {
		return a.ParseResponse(body, params[alConst.AliPublicKeyMap], alConst.RefundNode)
	}
}

func (a *AlPayService) OrderReverse(params map[string]string, count int) (result string, err error) {
	if count <= 0 {
		return "", errors.New("Numbers cannot be less than zero")
	}

	payData := *a.BuildCommonparam(params, alConst.ReverseCustom)

	bizContent := make(map[string]string)
	a.SetValue(&bizContent, alConst.OutTradeNo, params[alConst.OutTradeNoMap])
	a.SetValue(&bizContent, alConst.TradeNo, params[alConst.TradeNoMap])
	b, _ := json.Marshal(bizContent)
	payData[alConst.BizContent] = string(b)

	payData[alConst.Sign], _ = cryptoHelper.GetSha1Hash(mapHelper.SortedUrl(&payData), params[alConst.SellerPrivateKeyMap])

	p, _ := json.Marshal(payData)

	if _, body, reqErr := goreq.New().Get(alConst.OpenApi).Query(string(p)).End(); len(reqErr) != 0 {
		return "", fmt.Errorf("Network error", reqErr[0])
	} else {

		if result, e := a.ParseResponse(body, params[alConst.AliPublicKeyMap], alConst.ReverseNode); len(result) != 0 && e == nil {
			return result, nil
		} else {
			if len(result) == 0 {
				return "", e
			}
			rJson, _ := simplejson.NewJson([]byte(result))

			if recall, _ := rJson.Get(alConst.RetryFlag).String(); recall == "Y" {
				time.Sleep(10000 * time.Millisecond) //10s
				count = count - 1
				return a.OrderReverse(params, count)
			} else {
				if v, e := rJson.Get(alConst.Code).String(); e != nil {
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
	a.SetValue(&payData, alConst.AppId, commonParams[alConst.AppIdMap])
	a.SetValue(&payData, alConst.Charset, alConst.CharsetCustom)
	a.SetValue(&payData, alConst.Method, method)
	a.SetValue(&payData, alConst.SignType, alConst.SignTypeCustom)
	t := time.Now()
	a.SetValue(&payData, alConst.TimeStamp, fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second()))

	a.SetValue(&payData, alConst.Version, alConst.VersionCustom)
	a.SetValue(&payData, alConst.SignType, alConst.SignTypeCustom)

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
		sign, err := js.Get(alConst.Sign).String()
		if err != nil {
			return "", fmt.Errorf("parse response error", err.Error())
		}
		if isValid := cryptoHelper.CheckPubKey(bodyJs, sign, pubKey); isValid {
			if code, _ := jsDetail.Get(alConst.Code).String(); code == "10000" {
				return bodyJs, nil
			} else if code == "10003" {
				return "", fmt.Errorf("result is unknown", err.Error())
			} else {
				subCode, _ := jsDetail.Get(alConst.SubCode).String()
				return "", errors.New(subCode)
			}

		} else {
			return "", errors.New("response signature failed")
		}

	}
}
