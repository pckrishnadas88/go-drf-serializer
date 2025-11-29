package serializers

import "fmt"

type BooleanField struct {
	Required bool
}

func BooleanFieldField(required bool) *BooleanField {
	return &BooleanField{Required: required}
}

func (f *BooleanField) Validate(value any) error {
	_, ok := value.(bool)
	if !ok && f.Required {
		return fmt.Errorf("value must be boolean")
	}
	return nil
}
