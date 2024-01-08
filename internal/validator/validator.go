package validator

import (
	"fmt"
	"github.com/go-playground/validator"
)

type ValidationErrorResponse struct {
	Error       bool
	FailedField string
	Tag         string
	Value       any
	Message     string
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

func fmtMessage(field, actualTag string, param, value any) string {
	switch actualTag {
	case "required":
		return fmt.Sprintf("field: %s is %s, and must not be %d", field, actualTag, value)
	case "gte", "lte":
		return fmt.Sprintf("field: %s must be %s than %d", field, actualTag, value)
	default:
		return fmt.Sprint("Error...")
	}

}

func (v *Validator) Validate(data any) []ValidationErrorResponse {
	validationErrors := []ValidationErrorResponse{}
	errs := v.valdtr.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {

			fmt.Println(err.Field())
			fmt.Println(err.Tag())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Type())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			var elem ValidationErrorResponse
			elem.FailedField = err.Field()
			elem.Tag = err.Tag()
			elem.Value = err.Value()
			elem.Error = true
			elem.Message = fmtMessage(err.Field(), err.Tag(), err.Param(), err.Value())
			validationErrors = append(validationErrors, elem)
		}
	}
	return validationErrors
}
