package serializers

import (
	"fmt"
	"testing"
)

// Custom validator for letters only
func OnlyLetters(value any) error {
	str, ok := value.(string)
	if !ok {
		return nil // ignore non-string values
	}
	for _, r := range str {
		if r < 'A' || r > 'z' {
			return fmt.Errorf("must contain letters only")
		}
	}
	return nil
}

func TestCharField_Required(t *testing.T) {
	field := CharFieldField(true, 10)
	err := field.Validate(nil)
	if err == nil {
		t.Errorf("Expected error for required field, got nil")
	}

	err = field.Validate("HelloWorld")
	if err != nil {
		t.Errorf("Unexpected error for valid input: %v", err)
	}

	err = field.Validate("HelloWorldTooLong")
	if err == nil {
		t.Errorf("Expected error for string exceeding max length")
	}

}

func TestIntegerField_Required(t *testing.T) {
	field := IntegerFieldField(true)

	// nil should fail
	err := field.Validate(nil)
	if err == nil {
		t.Errorf("Expected error for required integer field, got nil")
	}

	// valid int
	err = field.Validate(42)
	if err != nil {
		t.Errorf("Unexpected error for valid integer: %v", err)
	}

	// JSON-style whole number float
	err = field.Validate(25.0)
	if err != nil {
		t.Errorf("Unexpected error for valid JSON number: %v", err)
	}

	// JSON-style float with decimal part should fail
	err = field.Validate(3.14)
	if err == nil {
		t.Errorf("Expected error for non-integer float")
	}
}

func TestSerializer_FieldLevelValidation(t *testing.T) {
	name := CharFieldField(true, 20)
	name.Validators = append(name.Validators, OnlyLetters)

	age := IntegerFieldField(true)

	s := New(map[string]Field{
		"name": name,
		"age":  age,
	})

	data := map[string]any{
		"name": "Krishna123", // invalid
		"age":  25,
	}

	errs := s.Validate(data)
	if errs == nil {
		t.Errorf("Expected validation errors for name field")
	}
	if _, ok := errs["name"]; !ok {
		t.Errorf("Expected error key 'name' missing")
	}

}

func TestSerializer_SerializerLevelValidation(t *testing.T) {
	age := IntegerFieldField(true)

	s := New(map[string]Field{
		"age": age,
	})

	s.Validators = append(s.Validators, func(data map[string]any) error {
		ageValue, ok := data["age"].(int)
		if !ok {
			return FieldError{Field: "age", Msg: "must be int"}
		}
		if ageValue < 18 {
			return FieldError{Field: "age", Msg: "age must be >= 18"}
		}
		return nil
	})

	data := map[string]any{
		"age": 16,
	}

	errs := s.Validate(data)
	if errs == nil || len(errs["age"]) == 0 {
		t.Errorf("Expected serializer-level validation error for age < 18")
	}

}
