package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/pckrishnadas88/go-drf-serializer/serializers"
)

func main() {
	status := serializers.ChoiceFieldField(true, []any{"active", "inactive", "pending"})
	birthdate := serializers.DateFieldField(true)
	website := serializers.URLFieldField(true)
	tags := serializers.ListFieldField(false, 1, 5)

	data := map[string]any{
		"status":    "blocked",
		"birthdate": "2025-11-29",
		"website":   "https://example.com",
		"tags":      []string{"go", "api"},
	}

	serializer := serializers.New(map[string]serializers.Field{
		"status":    status,
		"birthdate": birthdate,
		"website":   website,
		"tags":      tags,
	})

	errs := serializer.Validate(data)
	if errs != nil {
		fmt.Println("Validation error:")
		enc := json.NewEncoder(os.Stdout)
		enc.SetEscapeHTML(false)
		enc.SetIndent("", "  ")
		enc.Encode(errs)
		return
	} else {
		fmt.Println("Validation passed!")
	}
}
