package common

import (
	"unicode"

	"github.com/Vigiatonet/PSU-Calculator/config"
)

func ValidatePassword(password string) bool {
	cfg := config.GetConfig()

	if len(password) < cfg.Password.MinLength || len(password) > cfg.Password.MaxLength {
		return false
	}

	if cfg.Password.IncludeChars && !HasChars(password) {
		return false
	}
	if cfg.Password.IncludeDigits && !HasDigits(password) {
		return false
	}
	if cfg.Password.IncludeUppercase && !HasUpper(password) {
		return false
	}
	if cfg.Password.IncludeLowercase && !HasLower(password) {
		return false
	}
	return true
}

func HasDigits(s string) bool {
	for _, c := range s {
		if unicode.IsDigit(c) {
			return true
		}
	}
	return false
}

func HasLower(s string) bool {
	for _, c := range s {
		if unicode.IsLower(c) && unicode.IsLetter(c) {
			return true
		}
	}
	return false
}

func HasUpper(s string) bool {
	for _, c := range s {
		if unicode.IsUpper(c) && unicode.IsLetter(c) {
			return true
		}
	}
	return false
}

func HasChars(s string) bool {
	for _, c := range s {
		if unicode.IsLetter(c) {
			return true
		}
	}
	return false
}
