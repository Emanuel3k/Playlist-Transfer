package request

import (
	"encoding/json"
	"fmt"
	"github.com/emanuel3k/playlist-transfer/pkg/web"
	"github.com/go-playground/validator"
	"net/http"
)

var (
	ErrRequestContentTypeNotJSON = web.UnprocessableEntityError("request content type is not application/json")
	ErrRequestBodyInvalid        = web.BadRequestErrorWithCauses("request body is invalid", nil)
)

func Decode(r *http.Request, body any) *web.AppError {
	if r.Header.Get("Content-Type") != "application/json" {
		return ErrRequestContentTypeNotJSON
	}

	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		return ErrRequestContentTypeNotJSON
	}

	return nil
}

func Validate(body any) *web.AppError {
	validate := validator.New()

	if err := validate.Struct(body); err != nil {
		var causes []web.Cause
		for _, fieldErr := range err.(validator.ValidationErrors) {
			causes = append(causes, web.Cause{
				Field:   fieldErr.StructField(),
				Message: fmt.Sprint(fieldErr.Tag(), fieldErr.Param()),
			})
		}

		ErrRequestBodyInvalid.Causes = causes
		return ErrRequestBodyInvalid
	}

	return nil
}
