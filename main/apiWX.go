package main

import (
	"epaygo"
	"epaygo/helper"
	"fmt"
	"net/http"

	"epay.api.biz.go/epay.api.go/common"
	"epay.api.biz.go/epay.api.go/common/smt"

	"github.com/labstack/echo"
)

func DirectPayWX(c echo.Context) error {
	directPayDto := new(WxDirectPayDto)
	if err := c.Bind(directPayDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, APIResult{Success: false, Error: APIError{Message: "A required parameter is missing or doesn't have the right format" + "directPayDto"}})
	}
	directPayDto.OutTradeNo = helper.Uuid32()
	fmt.Printf("%#v", directPayDto)

	wxPayService := new(epaygo.WxPayService)

	dtoP := structToMap(directPayDto)

	if result, err := wxPayService.DirectPay(dtoP); err != nil {
		return c.JSON(http.StatusOK, APIResult{Success: false, Error: APIError{Code: err.Error()}})
	} else {
		fmt.Println(result)
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
	smt.Debug.Printf("%#v", dto)

	wxPayService := new(epaygo.WxPayService)
	dtoP := structToMap(dto)

	if result, err := wxPayService.OrderQueryWx(dtoP); err != nil {
		return c.JSON(http.StatusOK, APIResult{Success: false, Error: APIError{Code: err.Error()}})
	} else {
		smt.Debug.Println(result)
		return c.JSON(http.StatusOK, APIResult{Success: true, Result: result})
	}

}

func RefundWX(c echo.Context) error {
	dto := new(WxRefundDto)
	if err := c.Bind(dto); err != nil {
		return c.JSON(http.StatusBadRequest, APIResult{Success: false, Error: APIError{Message: "A required parameter is missing or doesn't have the right format" + "WxRefundDto"}})
	}
	dto.OutRefundNo = common.Uuid32()
	smt.Debug.Printf("%#v", dto)
	wxPayService := new(epaygo.WxPayService)
	dtoP := structToMap(dto)
	if result, err := wxPayService.RefundWx(dtoP); err != nil {
		return c.JSON(http.StatusOK, APIResult{Success: false, Error: APIError{Code: err.Error()}})
	} else {
		smt.Debug.Println(result)
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
	smt.Debug.Printf("%#v", dto)
	wxPayService := new(epaygo.WxPayService)
	dtoP := structToMap(dto)
	if result, err := wxPayService.ReverseWx(dtoP, 10); err != nil {
		return c.JSON(http.StatusOK, APIResult{Success: false, Error: APIError{Code: err.Error()}})
	} else {
		//c.JSON(http.StatusOK, APIResult{Success: true, Result: result})
		return c.JSON(http.StatusOK, APIResult{Success: true, Result: result})
		//c.String(http.StatusOK, result)
	}

}
