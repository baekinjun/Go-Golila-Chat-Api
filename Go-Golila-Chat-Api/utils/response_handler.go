package utils

import (
	"net/http"

	"github.com/labstack/echo"
)

type GeneralResponseType struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Payload map[string]interface{} `json:"payload"`
}

// Standard Response Generater
func GeneralResponse(code int, message string, payload map[string]interface{}) *GeneralResponseType {
	return &GeneralResponseType{
		Code:    code,
		Message: message,
		Payload: payload,
	}
}

// Error Handler
func HttpResponseErrorHandler(err error, c echo.Context) error {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		message := he.Message
		return c.JSON(code, GeneralResponse(code, message.(string), map[string]interface{}{}))
	}
	c.Logger().Error(err)
	return c.JSON(code, GeneralResponse(code, err.Error(), map[string]interface{}{}))
}

// General Response Handler
func GeneralResponseHandler(c echo.Context, code int, message string, payload map[string]interface{}) error {
	c.Logger().Debug(code, message, payload)
	return c.JSON(code, GeneralResponse(code, message, payload))
}

type AIServerResponse struct {
	Code    int
	Message string
	Payload map[string]interface{}
}
