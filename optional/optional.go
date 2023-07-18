package optional

var empty = &Optional[any]{v: nil}

// Optional struct
type Optional[T any] struct {
	v T
}

// Of value
func Of[T any](v T) *Optional[T] {
	return &Optional[T]{v: v}
}

// OfNillable value
func OfNillable[T any](v T) *Optional[T] {
	if v == nil {
		// return empty
		return &Optional[T]{v: v}
	}

	return &Optional[T]{v: v}
}

func (o *Optional[T]) Map(fn func(v T) T) *Optional[T] {
	if o.v == nil {
		// return empty
		return &Optional[T]{v: o.v}
	}

	return OfNillable(fn(o.v))
}

// Get value, will panic on value is nil
func (o *Optional[T]) Get() T {
	if o.v == nil {
		panic("nil value")
	}
	return o.v
}

// OrElse get value.
func (o *Optional[T]) OrElse(v T) T {
	if o.v == nil {
		return v
	}
	return o.v
}

// OrElseGet value.
func (o *Optional[T]) OrElseGet(v T) T {
	if o.v == nil {
		return v
	}
	return o.v
}

/**
TODO more
- ifPresent()
*/
