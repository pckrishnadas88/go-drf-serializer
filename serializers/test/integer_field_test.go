package serializers_test

import (
	"testing"

	"github.com/pckrishnadas88/go-drf-serializer/serializers"
)

func TestIntegerField_Public(t *testing.T) {
	field := serializers.IntegerFieldField(true)

	if err := field.Validate(10); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if err := field.Validate("abc"); err == nil {
		t.Error("Expected error for non-integer value")
	}

	if err := field.Validate(nil); err == nil {
		t.Error("Expected error for nil value on required field")
	}
}
