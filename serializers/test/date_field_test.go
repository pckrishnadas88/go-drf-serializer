package serializers_test

import (
	"testing"

	"github.com/pckrishnadas88/go-drf-serializer/serializers"
)

func TestDateField_Public(t *testing.T) {
	field := serializers.DateFieldField(true)

	if err := field.Validate("2025-11-29"); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if err := field.Validate("29-11-2025"); err == nil {
		t.Error("Expected error for invalid date format")
	}

	if err := field.Validate(nil); err == nil {
		t.Error("Expected error for nil value on required field")
	}
}
