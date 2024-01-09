package validator

import (
	"fmt"
	"github.com/go-playground/validator"
)

type ValidationError map[string]string

func (v ValidationError) Error() string {
	var errStr string
	for key, val := range v {
		errStr += fmt.Sprintf("%s:%s", key, val)
	}
	return errStr
}

func (v ValidationError) AddError(key, message string) {
	if _, exists := v[key]; !exists {
		v[key] = message
	}
}

func (v ValidationError) CheckErrors(err validator.FieldError) {
	switch err.ActualTag() {
	case "required":
		v.AddError("required", fmt.Sprintf("field: %s is %s, and must not be blank", err.Field(), err.Tag()))
		break
	case "gte", "lte":
		v.AddError("invalid-value", fmt.Sprintf("field: %s must be %s than %d", err.Field(), err.Tag(), err.Value()))
		break
	case "min":
		v.AddError("invalid-collection", fmt.Sprintf("field: %s must contain at least %s elements", err.Field(), err.Param()))
		break
	}
}

type Validator struct {
	valdtr *validator.Validate
	Errors ValidationError
}

func New() *Validator {
	v := validator.New()
	return &Validator{
		valdtr: v,
		Errors: make(ValidationError),
	}
}

func (v *Validator) Validate(data any) ValidationError {
	errs := v.valdtr.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			v.Errors.CheckErrors(err)
		}
	}
	return v.Errors
}
