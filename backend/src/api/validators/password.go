package validators

import (
	"github.com/Vigiatonet/PSU-Calculator/common"
	"github.com/go-playground/validator/v10"
)

func ValidatePassword(fld validator.FieldLevel) bool {
	value, ok := fld.Field().Interface().(string)
	if !ok {
		fld.Param()
		return false
	}
	return common.ValidatePassword(value)
}
