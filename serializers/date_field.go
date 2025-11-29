package serializers

import (
	"fmt"
	"time"
)

type DateField struct {
	Required bool
	Format   string // e.g. "2006-01-02"
}

func DateFieldField(required bool) *DateField {
	return &DateField{
		Required: required,
		Format:   "2006-01-02",
	}
}

func (f *DateField) Validate(value any) error {
	if value == nil {
		if f.Required {
			return fmt.Errorf("value is required")
		}
		return nil
	}
	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("value must be a string in format %s", f.Format)
	}
	_, err := time.Parse(f.Format, str)
	if err != nil {
		return fmt.Errorf("value must match format %s", f.Format)
	}
	return nil
}
