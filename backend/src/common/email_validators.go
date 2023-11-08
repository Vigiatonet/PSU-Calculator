package common

import (
	"regexp"

	"github.com/Vigiatonet/PSU-Calculator/config"
	"github.com/Vigiatonet/PSU-Calculator/pkg/logging"
)

var EmailRegex = `^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`
var log = logging.NewLogger(config.GetConfig())

func ValidateEmail(email string) bool {
	matched, err := regexp.MatchString(EmailRegex, email)
	if err != nil {
		log.Error(err, logging.General, logging.SubCategory(logging.Validation), "regex failed to match", nil)
		return false
	}
	return matched
}
