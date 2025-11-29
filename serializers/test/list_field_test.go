package serializers_test

import (
	"testing"

	"github.com/pckrishnadas88/go-drf-serializer/serializers"
)

func TestListField_Public(t *testing.T) {
	field := serializers.ListFieldField(true, 1, 3)

	if err := field.Validate([]any{1}); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if err := field.Validate([]any{}); err == nil {
		t.Error("Expected error for below min length")
	}

	if err := field.Validate([]any{1, 2, 3, 4}); err == nil {
		t.Error("Expected error for exceeding max length")
	}

	if err := field.Validate(nil); err == nil {
		t.Error("Expected error for nil value on required field")
	}
}
