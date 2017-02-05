package epaygo

import (
	"epaygo/al"
	"epaygo/bp"
	"epaygo/wx"
	"errors"
	"strings"
)

func CreatePayment(payType string) (IPayService, error) {

	switch strings.ToUpper(payType) {
	case "WX":
		return new(wx.WxPayService), nil
	case "AL":
		return new(al.AlPayService), nil
	case "BP":
		return new(bp.BpPayService), nil
	default:
		return nil, errors.New("Invalid Payment Type")
	}
}
