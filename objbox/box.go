package objbox

import "github.com/gookit/goutil"

// Box struct
type Box[V any] struct {
	objMap map[string]V
}

// New instance
func New[V any]() *Box[V] {
	return &Box[V]{
		objMap: map[string]V{},
	}
}

// Len of all objects
func (b *Box[V]) Len() int {
	return len(b.objMap)
}

// Set object by name
func (b *Box[V]) Set(name string, obj V) {
	b.objMap[name] = obj
}

// Add object by name
func (b *Box[V]) Add(name string, obj V) {
	b.objMap[name] = obj
}

// Has object by name
func (b *Box[V]) Has(name string) bool {
	_, ok := b.objMap[name]
	return ok
}

// Get object by name, if not exists will panic
func (b *Box[V]) Get(name string) V {
	obj, ok := b.objMap[name]
	if !ok {
		goutil.Panicf("object %q not exists in box", name)
	}
	return obj
}

// Lookup object by name
func (b *Box[V]) Lookup(name string) (V, bool) {
	obj, ok := b.objMap[name]
	return obj, ok
}

// Delete object by name
func (b *Box[V]) Delete(name string) bool {
	_, ok := b.objMap[name]
	if ok {
		delete(b.objMap, name)
	}
	return ok
}

// Reset all objects
func (b *Box[V]) Reset() {
	b.objMap = map[string]V{}
}
