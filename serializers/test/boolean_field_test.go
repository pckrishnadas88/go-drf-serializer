package serializers_test

import (
	"testing"

	"github.com/pckrishnadas88/go-drf-serializer/serializers"
)

func TestBooleanField_Public(t *testing.T) {
	field := serializers.BooleanFieldField(true)

	if err := field.Validate(true); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if err := field.Validate(nil); err == nil {
		t.Error("Expected error for nil value on required field")
	}

	if err := field.Validate("yes"); err == nil {
		t.Error("Expected error for non-boolean value")
	}
}
