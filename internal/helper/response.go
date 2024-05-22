package helper

import "github.com/labstack/echo/v4"

type BaseResponse struct {
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func ResponseData(code int, message string, data interface{}) *BaseResponse {
	var response BaseResponse

	response.Code = code
	response.Message = message

	if data != nil {
		response.Data = data
	}

	return &response
}

func ErrorHandler(c echo.Context, statusCode int, errorMessage string) error {
	response := ResponseData(statusCode, errorMessage, nil)
	return c.JSON(statusCode, response)
}
