package helper

import "github.com/Vigiatonet/PSU-Calculator/api/validators"

type Response struct {
	Result     interface{}
	StatusCode ResultCode
	Success    bool
	Err        string
}

func GenerateBaseResponse(result interface{}, statusCode ResultCode, success bool) *Response {
	return &Response{
		Result:     result,
		StatusCode: statusCode,
		Success:    success,
		Err:        "",
	}
}

func GenerateBaseResponseWithError(result interface{}, statusCode ResultCode, success bool, err error) *Response {
	return &Response{
		Result:     result,
		StatusCode: statusCode,
		Success:    success,
		Err:        err.Error(),
	}
}

func GenerateBaseResponseWithValidationError(statusCode ResultCode, success bool, err error) *Response {
	ve := validators.GetValidationError(err)
	return &Response{
		Result:     ve,
		StatusCode: statusCode,
		Success:    success,
		Err:        err.Error(),
	}
}
