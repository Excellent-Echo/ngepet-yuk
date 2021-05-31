package helper

import "github.com/go-playground/validator/v10"

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func APIResponse(msg string, code int, status string, data interface{}) Response {
	var meta = Meta{
		Message: msg,
		Code:    code,
		Status:  status,
	}

	var response = Response{
		Meta: meta,
		Data: data,
	}

	return response
}

func SplitErrorInformation(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}
