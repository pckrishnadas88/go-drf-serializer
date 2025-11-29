package serializers

import (
	"fmt"
	"reflect"
)

type ListField struct {
	Required bool
	MinLen   int
	MaxLen   int
}

func ListFieldField(required bool, minLen, maxLen int) *ListField {
	return &ListField{
		Required: required,
		MinLen:   minLen,
		MaxLen:   maxLen,
	}
}

func (f *ListField) Validate(value any) error {
	if value == nil {
		if f.Required {
			return fmt.Errorf("value is required")
		}
		return nil
	}

	v := reflect.ValueOf(value)
	if v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
		return fmt.Errorf("value must be a list/array")
	}

	l := v.Len()
	if f.MinLen > 0 && l < f.MinLen {
		return fmt.Errorf("list must have at least %d items", f.MinLen)
	}
	if f.MaxLen > 0 && l > f.MaxLen {
		return fmt.Errorf("list must have at most %d items", f.MaxLen)
	}

	return nil

}
