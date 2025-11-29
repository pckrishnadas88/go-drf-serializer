Here’s a **ready-to-use README.md** based on your full example code, including installation, features, and usage examples:

---

# go-drf-serializer

**DRF-style serializers and validation for Go** — minimal, dependency-free, and Gin-friendly.

This package provides:

* Field-level validation (`CharField`, `IntegerField`, `EmailField`, `BooleanField`)
* Custom validators
* Serializer-level validation
* Django REST Framework–style error output

---

## 📦 Installation

```bash
go get github.com/pckrishnadas88/go-drf-serializer
```

Import in your code:

```go
import "github.com/pckrishnadas88/go-drf-serializer/serializers"
```

---

## ⚡ Features

* DRF-like field declarations
* Required and optional fields
* Custom validation per field
* Serializer-level validators
* DRF-style error responses:

```json
{
  "name": ["must contain letters only"],
  "age": ["age must be >= 18"]
}
```

---

## 📄 Example Usage

### main.go

```go
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
```

---

## ✅ Sample Output

```json
Validation error:
{
  "name": [
    "must contain letters only"
  ],
  "age": [
    "age must be >= 18"
  ]
}
```

---

## 🧱 Supported Fields

| Field                        | Required   | Notes                                            |
| ---------------------------- | ---------- | ------------------------------------------------ |
| `CharFieldField(maxLen int)` | true/false | Supports `Validators` slice                      |
| `IntegerFieldField()`        | true/false | Supports min/max validation via custom validator |
| `EmailFieldField()`          | true/false | Validates email format                           |
| `BooleanFieldField()`        | true/false | Simple boolean check                             |

---

## 🔧 Adding Custom Validators

```go
name.Validators = append(name.Validators, func(value any) error {
    // custom logic
    return nil
})
```

Serializer-level validators:

```go
user.Validators = append(user.Validators, func(data map[string]any) error {
    if data["age"].(int) < 18 {
        return serializers.FieldError{Field: "age", Msg: "age must be >= 18"}
    }
    return nil
})
```

---

## 🏗 Roadmap

* Nested serializers
* ListField support
* Automatic JSON binding from HTTP requests
* Extended DRF-like field options (MinLength, MaxValue, Regex)