package serializers

import (
	"fmt"
)

type IntegerField struct {
	Required bool
}

func IntegerFieldField(required bool) *IntegerField {
	return &IntegerField{Required: required}
}

// Validate accepts JSON numbers (float64) and converts to int internally
func (f *IntegerField) Validate(value any) error {
	if value == nil {
		if f.Required {
			return fmt.Errorf("value is required")
		}
		return nil
	}

	switch v := value.(type) {
	case int:
		return nil
	case float64:
		// JSON numbers come as float64, check if integer
		if v != float64(int(v)) {
			return fmt.Errorf("value must be integer")
		}
		return nil
	default:
		return fmt.Errorf("value must be integer")
	}
}
