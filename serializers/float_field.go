package serializers

import "fmt"

type FloatField struct {
	Required bool
}

func FloatFieldField(required bool) *FloatField {
	return &FloatField{Required: required}
}

func (f *FloatField) Validate(value any) error {
	_, ok := value.(float64)
	if !ok && f.Required {
		return fmt.Errorf("value must be float")
	}
	return nil
}
