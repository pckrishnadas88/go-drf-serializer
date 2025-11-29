package serializers_test

import (
	"testing"

	"github.com/pckrishnadas88/go-drf-serializer/serializers"
)

func TestSerializer_FieldLevelValidation_Public(t *testing.T) {
	name := serializers.CharFieldField(true, 5)
	age := serializers.IntegerFieldField(true)

	s := serializers.New(map[string]serializers.Field{
		"name": name,
		"age":  age,
	})

	data := map[string]any{
		"name": "abcdef", // too long
		"age":  "xyz",    // invalid type
	}

	errs := s.Validate(data)
	if errs == nil {
		t.Error("Expected validation errors")
	}

	if len(errs["name"]) == 0 || len(errs["age"]) == 0 {
		t.Error("Expected errors for name and age fields")
	}
}

func TestSerializer_SerializerLevelValidation_Public(t *testing.T) {
	age := serializers.IntegerFieldField(true)

	s := serializers.New(map[string]serializers.Field{
		"age": age,
	})

	s.Validators = append(s.Validators, func(data map[string]any) error {
		v, ok := data["age"].(int)
		if !ok {
			return serializers.FieldError{Field: "age", Msg: "must be integer"}
		}
		if v < 18 {
			return serializers.FieldError{Field: "age", Msg: "must be >= 18"}
		}
		return nil
	})

	data := map[string]any{"age": 16}

	errs := s.Validate(data)
	if errs == nil || len(errs["age"]) == 0 {
		t.Error("Expected serializer-level validation error for age")
	}
}
