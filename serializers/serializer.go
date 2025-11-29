package serializers

type Serializer struct {
	Fields     map[string]Field
	Validators []func(data map[string]any) error
}

func New(fields map[string]Field) *Serializer {
	return &Serializer{
		Fields:     fields,
		Validators: []func(data map[string]any) error{},
	}
}

func (s *Serializer) Validate(data map[string]any) ValidationErrors {
	errors := ValidationErrors{}

	// Field-level validation
	for key, field := range s.Fields {
		value, exists := data[key]
		if !exists {
			value = nil
		}

		if err := field.Validate(value); err != nil {
			errors.AddFieldError(key, err.Error())
		}
	}

	// Serializer-level validation
	for _, v := range s.Validators {
		if err := v(data); err != nil {
			switch e := err.(type) {

			case FieldError:
				errors.AddFieldError(e.Field, e.Msg)

			default:
				errors.AddGlobalError(err.Error())
			}
		}
	}

	if len(errors) == 0 {
		return nil
	}
	return errors
}
