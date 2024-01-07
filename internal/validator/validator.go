package validator

import (
	"github.com/go-playground/validator"
)

type ValidationErrorResponse struct {
	Error       bool
	FailedField string
	Tag         string
	Value       any
}

type Validator struct {
	valdtr *validator.Validate
}

func New() *Validator {
	v := validator.New()
	return &Validator{
		valdtr: v,
	}
}

func (v *Validator) Validate(data any) []ValidationErrorResponse {
	validationErrors := []ValidationErrorResponse{}
	errs := v.valdtr.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			var elem ValidationErrorResponse
			elem.FailedField = err.Field()
			elem.Tag = err.Tag()
			elem.Value = err.Value()
			elem.Error = true
			validationErrors = append(validationErrors, elem)
		}
	}
	return validationErrors
}
