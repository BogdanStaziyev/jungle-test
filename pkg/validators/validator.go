package validators

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator"
)

type Validator interface {
	ValidateRequest(r *http.Request, i interface{}) error
}

type validate struct {
	validator *validator.Validate
}

func NewValidator() *validate {
	return &validate{
		validator: validator.New(),
	}
}

// ValidateRequest validates the request to a data structure depending on the tags
func (v *validate) ValidateRequest(r *http.Request, i interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
		return err
	}
	if err := v.validator.Struct(i); err != nil {
		return err
	}
	return nil
}
