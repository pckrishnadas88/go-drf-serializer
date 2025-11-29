package main

import (
	"encoding/json"
	"fmt"
	"os"
	"unicode"

	"github.com/pckrishnadas88/go-drf-serializer/serializers"
)

// Example custom validator: only letters
func OnlyLetters(value any) error {
	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("not a string")
	}
	for _, r := range str {
		if !unicode.IsLetter(r) {
			return fmt.Errorf("must contain letters only")
		}
	}
	return nil
}

func main() {
	// Field-level
	name := serializers.CharFieldField(true, 20)
	name.Validators = append(name.Validators, OnlyLetters)

	age := serializers.IntegerFieldField(true)
	email := serializers.EmailFieldField(true)
	isActive := serializers.BooleanFieldField(false)

	// Serializer
	user := serializers.New(map[string]serializers.Field{
		"name":     name,
		"age":      age,
		"email":    email,
		"isActive": isActive,
	})

	// Serializer-level validator
	user.Validators = append(user.Validators, func(data map[string]any) error {
		age, ok := data["age"].(int)
		if !ok {
			return serializers.FieldError{Field: "age", Msg: "age must be an integer"}
		}
		if age < 18 {
			return serializers.FieldError{Field: "age", Msg: "age must be >= 18"}
		}
		return nil
	})
	// Invalid data
	data := map[string]any{
		"name":     "Krishna123",
		"age":      17,
		"email":    "krish@example.com",
		"isActive": true,
	}

	if err := user.Validate(data); err != nil {
		fmt.Println("Validation error:")
		enc := json.NewEncoder(os.Stdout)
		enc.SetEscapeHTML(false)
		enc.SetIndent("", "  ")
		enc.Encode(err)
		return
	} else {
		fmt.Println("Validation passed!")
	}
}
