package serializers

import (
	"fmt"
	"net/url"
)

type URLField struct {
	Required bool
}

func URLFieldField(required bool) *URLField {
	return &URLField{Required: required}
}

func (f *URLField) Validate(value any) error {
	if value == nil {
		if f.Required {
			return fmt.Errorf("value is required")
		}
		return nil
	}
	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("value must be a string")
	}
	u, err := url.ParseRequestURI(str)
	if err != nil {
		return fmt.Errorf("value must be a valid URL")
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return fmt.Errorf("URL scheme must be http or https")
	}
	if u.Host == "" {
		return fmt.Errorf("URL must have a valid host")
	}
	return nil
}
