package request

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (req Login) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Username, validation.Required, is.Alphanumeric),
		validation.Field(&req.Password, validation.Required, is.Alphanumeric),
	)
}
