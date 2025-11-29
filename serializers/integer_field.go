package serializers

import "fmt"

type IntegerField struct {
	Required bool
}

func IntegerFieldField(required bool) *IntegerField {
	return &IntegerField{Required: required}
}

func (f *IntegerField) Validate(value any) error {
	_, ok := value.(int)
	if !ok && f.Required {
		return fmt.Errorf("value must be integer")
	}
	return nil
}
