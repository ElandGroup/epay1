package main

import (
	"epaygo"
	"epaygo/helper"
	"net/http"

	"github.com/labstack/echo"
)

func DirectPayWX(c echo.Context) error {
	directPayDto := new(WxDirectPayDto)
	if err := c.Bind(directPayDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, APIResult{Success: false, Error: APIError{Message: "A required parameter is missing or doesn't have the right format" + "directPayDto"}})
	}
	directPayDto.OutTradeNo = helper.Uuid32()

	//wxPayService := new(epaygo.WxPayService)
	payService, _ := epaygo.CreatePayment("WX")

	dtoP := structToMap(directPayDto)

	if result, err := payService.DirectPay(dtoP); err != nil {
		return c.JSON(http.StatusOK, APIResult{Success: false, Error: APIError{Code: err.Error()}})
	} else {
		//c.JSON(http.StatusOK, APIResult{Success: true, Result: result})
		return c.JSON(http.StatusOK, APIResult{Success: true, Result: result})
		//c.String(http.StatusOK, result)
	}

}

func OrderQueryWX(c echo.Context) error {
	dto := new(WxOrderQueryDto)
	if err := c.Bind(dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "A required parameter is missing or doesn't have the right format"+"WxOrderQueryDto")
	}

	payService, _ := epaygo.CreatePayment("WX")

	dtoP := structToMap(dto)

	if result, err := payService.OrderQuery(dtoP); err != nil {
		return c.JSON(http.StatusOK, APIResult{Success: false, Error: APIError{Code: err.Error()}})
	} else {
		return c.JSON(http.StatusOK, APIResult{Success: true, Result: result})
	}

}

func RefundWX(c echo.Context) error {
	dto := new(WxRefundDto)
	if err := c.Bind(dto); err != nil {
		return c.JSON(http.StatusBadRequest, APIResult{Success: false, Error: APIError{Message: "A required parameter is missing or doesn't have the right format" + "WxRefundDto"}})
	}
	dto.OutRefundNo = helper.Uuid32()
	payService, _ := epaygo.CreatePayment("WX")

	dtoP := structToMap(dto)
	if result, err := payService.Refund(dtoP); err != nil {
		return c.JSON(http.StatusOK, APIResult{Success: false, Error: APIError{Code: err.Error()}})
	} else {
		//c.JSON(http.StatusOK, APIResult{Success: true, Result: result})
		return c.JSON(http.StatusOK, APIResult{Success: true, Result: result})
		//c.String(http.StatusOK, result)
	}

}

func ReverseWX(c echo.Context) error {

	dto := new(WxReverseDto)
	if err := c.Bind(dto); err != nil {
		return c.JSON(http.StatusBadRequest, APIResult{Success: false, Error: APIError{Message: "A required parameter is missing or doesn't have the right format" + "WxReverseDto"}})
	}
	payService, _ := epaygo.CreatePayment("WX")
	dtoP := structToMap(dto)
	if result, err := payService.OrderReverse(dtoP, 10); err != nil {
		return c.JSON(http.StatusOK, APIResult{Success: false, Error: APIError{Code: err.Error()}})
	} else {
		//c.JSON(http.StatusOK, APIResult{Success: true, Result: result})
		return c.JSON(http.StatusOK, APIResult{Success: true, Result: result})
		//c.String(http.StatusOK, result)
	}

}
