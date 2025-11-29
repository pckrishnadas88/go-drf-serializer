# go-drf-serializer

**DRF-style serializers and validation for Go** — minimal, dependency-free, and Gin-friendly.

This package provides:

* Field-level validation (`CharField`, `IntegerField`, `EmailField`, `BooleanField`, `ChoiceField`, `DateField`, `ListField`, `URLField`)
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
	name := serializers.CharFieldField(true, 20)
	name.Validators = append(name.Validators, OnlyLetters)

	age := serializers.IntegerFieldField(true)
	email := serializers.EmailFieldField(true)
	isActive := serializers.BooleanFieldField(false)
	status := serializers.ChoiceFieldField(true, []any{"active", "inactive", "pending"})
	birthdate := serializers.DateFieldField(true)
	website := serializers.URLFieldField(true)
	tags := serializers.ListFieldField(false, 1, 5)

	user := serializers.New(map[string]serializers.Field{
		"name":      name,
		"age":       age,
		"email":     email,
		"isActive":  isActive,
		"status":    status,
		"birthdate": birthdate,
		"website":   website,
		"tags":      tags,
	})

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

	data := map[string]any{
		"name":      "Krishna123",
		"age":       17,
		"email":     "krish@example.com",
		"isActive":  true,
		"status":    "blocked",
		"birthdate": "2025-11-29",
		"website":   "https:://example.com",
		"tags":      []string{"go", "api"},
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

## Gin Example

```go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pckrishnadas88/go-drf-serializer/serializers"
	"unicode"
)

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
	r := gin.Default()

	name := serializers.CharFieldField(true, 20)
	name.Validators = append(name.Validators, OnlyLetters)
	age := serializers.IntegerFieldField(true)

	userSerializer := serializers.New(map[string]serializers.Field{
		"name": name,
		"age":  age,
	})

	userSerializer.Validators = append(userSerializer.Validators, func(data map[string]any) error {
		if ageValue, ok := data["age"].(int); ok && ageValue < 18 {
			return serializers.FieldError{Field: "age", Msg: "age must be >= 18"}
		}
		return nil
	})

	r.POST("/users", func(c *gin.Context) {
		data, ok := serializers.GinBindAndValidate(c, userSerializer)
		if !ok {
			return
		}
		c.JSON(200, gin.H{
			"message": "Validation passed!",
			"data":    data,
		})
	})

	r.Run(":8080")
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
| `ChoiceFieldField(choices)`  | true/false | Only allows values in the provided slice         |
| `DateFieldField()`           | true/false | Validates `YYYY-MM-DD` format                    |
| `ListFieldField(min,max)`    | true/false | Validates slice/array length                     |
| `URLFieldField()`            | true/false | Validates URL scheme and host                    |

---

## 🔧 Custom Validators

```go
name.Validators = append(name.Validators, func(value any) error {
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

## Running Tests

Run unit tests to ensure field and serializer validation works correctly:

```bash
cd serializers
go test -v
```

---

## 🏗 Roadmap

* Nested serializers
* Extended DRF-like options (MinLength, MaxValue, Regex)
* Additional field types
* More Gin helpers

---

## License

This project is licensed under the **MIT License**.
See [LICENSE](LICENSE) file for details.

---

## Contact

Email: [pckrishnadas88@gmail.com](mailto:pckrishnadas88@gmail.com)
Twitter / X: [@pckrishnadas88](https://twitter.com/pckrishnadas88)
