package serializers_test

import (
	"testing"

	"github.com/pckrishnadas88/go-drf-serializer/serializers"
)

func TestChoiceField_Public(t *testing.T) {
	field := serializers.ChoiceFieldField(true, []any{"active", "inactive"})

	if err := field.Validate("active"); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if err := field.Validate("blocked"); err == nil {
		t.Error("Expected error for invalid choice")
	}

	if err := field.Validate(nil); err == nil {
		t.Error("Expected error for nil value on required field")
	}
}
