package serializers

import (
	"fmt"
	"reflect"
)

type ChoiceField struct {
	Required bool
	Choices  []any
}

func ChoiceFieldField(required bool, choices []any) *ChoiceField {
	return &ChoiceField{
		Required: required,
		Choices:  choices,
	}
}

func (f *ChoiceField) Validate(value any) error {
	if value == nil {
		if f.Required {
			return fmt.Errorf("value is required")
		}
		return nil
	}
	for _, c := range f.Choices {
		if reflect.DeepEqual(value, c) {
			return nil
		}
	}
	return fmt.Errorf("value must be one of %v", f.Choices)
}
