package values

// NewBuilder creates a new value builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a value builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithNumber(number byte) Builder
	Now() (Value, error)
}

// Value represents a value
type Value interface {
	Name() string
	Number() byte
}
