package main

import (
	"net/http"

	"github.com/labstack/echo"
)

var (
	e = echo.New()
)

func main() {
	RegisterApi()
	e.Start(":1001")
}

func RegisterApi() {

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello epay")
	})

	apiGroup := e.Group("/api")
	v1 := apiGroup.Group("/v1")

	wx := v1.Group("/wx")
	al := v1.Group("/al")

	wx.POST("/pay", DirectPayWX)

	wx.POST("/query", OrderQueryWX)

	wx.POST("/refund", RefundWX)

	wx.POST("/reverse", ReverseWX)

	al.POST("/pay", DirectPayAL)

	al.POST("/query", OrderQueryAL)

	al.POST("/refund", RefundAL)

	al.POST("/reverse", ReverseAL)

}
