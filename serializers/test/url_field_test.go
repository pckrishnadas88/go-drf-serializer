package serializers_test

import (
	"testing"

	"github.com/pckrishnadas88/go-drf-serializer/serializers"
)

func TestURLField_Public(t *testing.T) {
	field := serializers.URLFieldField(true)

	if err := field.Validate("https://example.com"); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if err := field.Validate("htp:/invalid"); err == nil {
		t.Error("Expected error for invalid URL")
	}

	if err := field.Validate(nil); err == nil {
		t.Error("Expected error for nil value on required field")
	}
}
