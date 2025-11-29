package serializers

type Field interface {
	Validate(value any) error
}
