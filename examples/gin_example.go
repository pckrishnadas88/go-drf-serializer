package main

import (
	"fmt"
	"unicode"

	"github.com/gin-gonic/gin"
	"github.com/pckrishnadas88/go-drf-serializer/serializers"
)

// Custom validator: only letters
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

	// Field-level
	name := serializers.CharFieldField(true, 20)
	name.Validators = append(name.Validators, OnlyLetters)

	age := serializers.IntegerFieldField(true)

	// Serializer
	userSerializer := serializers.New(map[string]serializers.Field{
		"name": name,
		"age":  age,
	})

	// Serializer-level validator: only range check, no type check
	userSerializer.Validators = append(userSerializer.Validators, func(data map[string]any) error {
		if ageValue, ok := data["age"].(int); ok {
			if ageValue < 18 {
				return serializers.FieldError{Field: "age", Msg: "age must be >= 18"}
			}
		}
		return nil
	})

	// Gin endpoint
	r.POST("/users", func(c *gin.Context) {
		data, ok := serializers.GinBindAndValidate(c, userSerializer)
		if !ok {
			return // GinBindAndValidate already returned DRF-style errors
		}
		c.JSON(200, gin.H{
			"message": "Validation passed!",
			"data":    data,
		})
	})

	r.Run(":8080")
}
