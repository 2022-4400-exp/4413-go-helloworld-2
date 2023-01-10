package main

import (
	"io"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/hello", Hello())
	e.GET("/api/hello", ApiHelloGet())
	e.GET("/uni", ShiryobukaiUni())

	e.Start(":8080")
}

func Hello() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "4413: 思慮深いウニ「俺が思うに...」")
	}
}

func ApiHelloGet() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{"studentId": "4413", "message": "思慮深いウニ「俺が思うに...」"})
	}
}

func ShiryobukaiUni() echo.HandlerFunc {
	return func(c echo.Context) error {
		url := "https://gist.githubusercontent.com/souhait0614/eaa165f0178985be17aebec6fe55710c/raw/shiryobukai-uni.txt"

		res, _ := http.Get(url)
		defer res.Body.Close()

		body, _ := io.ReadAll(res.Body)

		html := "<p><a href=\"https://gist.github.com/souhait0614/eaa165f0178985be17aebec6fe55710c\" target=\"_blank\">思慮深いウニ</a></p><p>" + strings.Replace(string(body), "\n", "<br>", -1) + "</p>"

		return c.HTML(http.StatusOK, html)
	}
}
