//go:build go1.18

package optional

var empty = &optional[any]{v: nil}

// optional struct
type optional[T any] struct {
	v T
}

// Of value
func Of[T any](v T) *optional[T] {
	return &optional[T]{v: v}
}

// OfNillable value
func OfNillable[T any](v T) *optional[T] {
	if v == nil {
		// return empty
		return &optional[T]{v: nil}
	}

	return &optional[T]{v: v}
}

func (o *optional[T]) Map(fn func(v T) T) *optional[T] {
	if o.v == nil {
		// return empty
		return &optional[T]{v: nil}
	}

	return OfNillable(fn(o.v))
}

// Get value, will panic on value is nil
func (o *optional[T]) Get() T {
	if o.v == nil {
		panic("nil value")
	}
	return o.v
}

// OrElse get value.
func (o *optional[T]) OrElse(v T) T {
	if o.v == nil {
		return v
	}
	return o.v
}

// OrElseGet value.
func (o *optional[T]) OrElseGet(v T) T {
	if o.v == nil {
		return v
	}
	return o.v
}
