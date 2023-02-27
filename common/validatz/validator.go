package validatz

import (
	"github.com/go-playground/validator/v10"
)

var (
	xValidator *validator.Validate = validator.New()
)

func Validate(data any, lang string) string {
	err := xValidator.Struct(data)
	if err == nil {
		return ""
	}

	// todo: add translation

	return err.Error()
}
