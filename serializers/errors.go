package serializers

type ValidationErrors map[string][]string

func (ve ValidationErrors) AddFieldError(field, msg string) {
	ve[field] = append(ve[field], msg)
}

const GlobalErrorKey = "non_field_errors"

func (ve ValidationErrors) AddGlobalError(msg string) {
	ve[GlobalErrorKey] = append(ve[GlobalErrorKey], msg)
}

// Returned by validators to attach error to a field
type FieldError struct {
	Field string
	Msg   string
}

func (fe FieldError) Error() string {
	return fe.Msg
}
