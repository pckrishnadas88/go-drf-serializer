package serializers

import (
	"fmt"
)

type CharField struct {
	Required   bool
	MaxLength  int
	Validators []func(value any) error
}

func CharFieldField(required bool, maxLen int) *CharField {
	return &CharField{
		Required:  required,
		MaxLength: maxLen,
	}
}

func (f *CharField) Validate(value any) error {
	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("value must be a string")
	}

	if f.Required && str == "" {
		return fmt.Errorf("field is required")
	}

	if f.MaxLength > 0 && len(str) > f.MaxLength {
		return fmt.Errorf("max length exceeded")
	}

	for _, v := range f.Validators {
		if err := v(str); err != nil {
			return err
		}
	}

	return nil
}
