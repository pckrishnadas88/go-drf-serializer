package serializers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GinBindAndValidate binds JSON and validates using the serializer
func GinBindAndValidate(c *gin.Context, s *Serializer) (map[string]any, bool) {
	var jsonData map[string]any
	if err := c.ShouldBindJSON(&jsonData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"non_field_errors": []string{"invalid JSON"},
		})
		return nil, false
	}

	if errs := s.Validate(jsonData); errs != nil {
		c.JSON(http.StatusBadRequest, errs)
		return nil, false
	}

	return jsonData, true
}
