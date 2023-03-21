package validators

import (
	"errors"
	"strings"

	"github.com/go-playground/validator"
)

type validators struct {
	validator *validator.Validate
}

func NewValidator() *validators {
	return &validators{
		validator: validator.New(),
	}
}

// Validate validates the request to a data structure depending on the tags
func (v *validators) Validate(i interface{}) error {
	err := v.validator.Struct(i)
	if err == nil {
		return nil
	}
	err = errors.New(strings.Replace(err.Error(), "\n", ", ", -1))
	return err
}
