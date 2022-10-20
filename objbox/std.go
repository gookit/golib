package objbox

var std = New()

// Set object by name
func Set[T any](name string, obj T) {
	std.Set(name, obj)
}

// Add object by name
func Add[T any](name string, obj T) {
	std.Add(name, obj)
}

// Has object by name
func Has(name string) bool {
	return std.Has(name)
}

// Get object by name, if not exists will panic
func Get[T any](name string) T {
	return std.Get(name).(T)
}

// Lookup object by name
func Lookup[T any](name string) (v T, ok bool) {
	if obj, ok := std.Lookup(name); ok {
		return obj.(T), true
	}
	return
}

// Len of objects
func Len() int {
	return std.Len()
}

// Delete object by name
func Delete(name string) bool {
	return std.Delete(name)
}

// Reset all objects
func Reset() {
	std.Reset()
}
