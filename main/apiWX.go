package main

import (
	"epaygo"
	. "epaygo/core/commonDto"
	"epaygo/core/helper"
	"net/http"

	"github.com/labstack/echo"
)

func DirectPayWX(c echo.Context) error {
	dto := new(WxDirectPayDto)
	if err := c.Bind(dto); err != nil {
		return c.JSON(http.StatusBadRequest, APIResult{Success: false, Error: APIError{Code: 10012, Message: BadRequestMessage(dto)}})
	}
	dto.OutTradeNo = helper.UuIdForPay(UuIdWxOutTradeNo)

	//wxPayService := new(epaygo.WxPayService)
	payService, _ := epaygo.CreatePayment("WX")

	dtoP := structToMap(dto)

	if result, err := payService.DirectPay(dtoP); err != nil {
		return c.JSON(http.StatusOK, APIResult{Success: false, Error: *err})
	} else {
		//c.JSON(http.StatusOK, APIResult{Success: true, Result: result})
		return c.JSON(http.StatusOK, APIResult{Success: true, Result: result})
		//c.String(http.StatusOK, result)
	}

}

func OrderQueryWX(c echo.Context) error {
	dto := new(WxOrderQueryDto)
	if err := c.Bind(dto); err != nil {
		return c.JSON(http.StatusBadRequest, APIResult{Success: false, Error: APIError{Code: 10012, Message: BadRequestMessage(dto)}})
	}

	payService, _ := epaygo.CreatePayment("WX")

	dtoP := structToMap(dto)

	if result, err := payService.OrderQuery(dtoP); err != nil {
		return c.JSON(http.StatusOK, APIResult{Success: false, Error: *err})
	} else {
		return c.JSON(http.StatusOK, APIResult{Success: true, Result: result})
	}

}

func RefundWX(c echo.Context) error {
	dto := new(WxRefundDto)
	if err := c.Bind(dto); err != nil {
		return c.JSON(http.StatusBadRequest, APIResult{Success: false, Error: APIError{Code: 10012, Message: BadRequestMessage(dto)}})
	}
	dto.OutRefundNo = helper.UuIdForPay(UuIdWxRefundNo)
	payService, _ := epaygo.CreatePayment("WX")

	dtoP := structToMap(dto)
	if result, err := payService.Refund(dtoP); err != nil {
		return c.JSON(http.StatusOK, APIResult{Success: false, Error: *err})
	} else {
		//c.JSON(http.StatusOK, APIResult{Success: true, Result: result})
		return c.JSON(http.StatusOK, APIResult{Success: true, Result: result})
		//c.String(http.StatusOK, result)
	}

}

func ReverseWX(c echo.Context) error {

	dto := new(WxReverseDto)
	if err := c.Bind(dto); err != nil {
		return c.JSON(http.StatusBadRequest, APIResult{Success: false, Error: APIError{Code: 10012, Message: BadRequestMessage(dto)}})
	}
	payService, _ := epaygo.CreatePayment("WX")
	dtoP := structToMap(dto)
	if result, err := payService.OrderReverse(dtoP, 10); err != nil {
		return c.JSON(http.StatusOK, APIResult{Success: false, Error: *err})
	} else {
		//c.JSON(http.StatusOK, APIResult{Success: true, Result: result})
		return c.JSON(http.StatusOK, APIResult{Success: true, Result: result})
		//c.String(http.StatusOK, result)
	}

}
