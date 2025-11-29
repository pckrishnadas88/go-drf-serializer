package serializers_test

import (
	"testing"

	"github.com/pckrishnadas88/go-drf-serializer/serializers"
)

func TestCharField_Public(t *testing.T) {
	field := serializers.CharFieldField(true, 5)
	if err := field.Validate(""); err == nil {
		t.Error("Expected error for empty string")
	}
}
