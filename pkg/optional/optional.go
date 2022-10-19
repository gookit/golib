//go:build go1.18
// +build go1.18

package optional

var empty = &optional{v: nil}

// optional struct
type optional struct {
	v any
}

// Of data
func Of(data any) *optional {
	return &optional{v: data}
}

// OfNillable data
func OfNillable(data any) *optional {
	if data == nil {
		return empty
	}

	return &optional{v: data}
}

func (o *optional) Map(fn func(v any) any) *optional {
	if o.v == nil {
		return empty
	}

	return OfNillable(fn(o.v))
}

// Get value, will panic on value is nil
func (o *optional) Get() any {
	if o.v == nil {
		panic("nil value")
	}

	return o.v
}

// OrElse get value.
func (o *optional) OrElse(v any) any {
	if o.v == nil {
		return v
	}

	return o.v
}

// OrElseGet value.
func (o *optional) OrElseGet(v any) any {
	if o.v == nil {
		return v
	}
	return o.v
}
