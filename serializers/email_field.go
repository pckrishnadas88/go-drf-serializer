package serializers

import (
	"fmt"
	"regexp"
)

type EmailField struct {
	Required bool
}

func EmailFieldField(required bool) *EmailField {
	return &EmailField{Required: required}
}

func (f *EmailField) Validate(value any) error {
	str, ok := value.(string)
	if !ok || str == "" {
		if f.Required {
			return fmt.Errorf("field is required")
		}
		return nil
	}
	regex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	if !regex.MatchString(str) {
		return fmt.Errorf("invalid email")
	}
	return nil
}
