package validator

import (
	"strings"
	"unicode/utf8"
)

type UserValidator struct{}

func (uv *UserValidator) Validate(name string) map[string]string {
	errors := make(map[string]string)

	if strings.TrimSpace(name) == "" {
		errors["name"] = "The name field cannot be blank"
	}

	if utf8.RuneCountInString(name) > 255 {
		errors["name"] = "The name field cannot be mmore than 255 characters long"
	}

	return errors
}
