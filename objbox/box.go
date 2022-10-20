package objbox

import "github.com/gookit/goutil"

// Box struct
type Box struct {
	objMap map[string]any
}

// New instance
func New() *Box {
	return &Box{
		objMap: map[string]any{},
	}
}

// Len of all objects
func (b *Box) Len() int {
	return len(b.objMap)
}

// Set object by name
func (b *Box) Set(name string, obj any) {
	b.objMap[name] = obj
}

// Add object by name
func (b *Box) Add(name string, obj any) {
	b.objMap[name] = obj
}

// Has object by name
func (b *Box) Has(name string) bool {
	_, ok := b.objMap[name]
	return ok
}

// Get object by name, if not exists will panic
func (b *Box) Get(name string) any {
	obj, ok := b.objMap[name]
	if !ok {
		goutil.Panicf("object %q not exists in box", name)
	}
	return obj
}

// Lookup object by name
func (b *Box) Lookup(name string) (any, bool) {
	obj, ok := b.objMap[name]
	return obj, ok
}

// Delete object by name
func (b *Box) Delete(name string) bool {
	_, ok := b.objMap[name]
	if ok {
		delete(b.objMap, name)
	}
	return ok
}

// Reset all objects
func (b *Box) Reset() {
	b.objMap = map[string]any{}
}
