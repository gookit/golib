package objbox_test

import (
	"fmt"
	"testing"

	"github.com/gookit/golib/objbox"
	"github.com/gookit/goutil/testutil/assert"
)

type testObj struct {
	Name string
}

func Example_std() {
	objbox.Add("obj1", &testObj{Name: "inhere"})

	obj1 := objbox.Get[*testObj]("obj1")
	fmt.Println(obj1.Name) // "inhere"
}

func TestStd_Get(t *testing.T) {
	assert.Eq(t, 0, objbox.Len())

	objbox.Add("obj1", &testObj{Name: "inhere"})
	assert.Eq(t, 1, objbox.Len())

	assert.True(t, objbox.Has("obj1"))

	obj1 := objbox.Get[*testObj]("obj1")
	assert.Eq(t, "inhere", obj1.Name)

	obj1, ok := objbox.Lookup[*testObj]("obj1")
	assert.True(t, ok)
	assert.Eq(t, "inhere", obj1.Name)

	obj2, ok := objbox.Lookup[*testObj]("obj2")
	assert.False(t, ok)
	assert.Nil(t, obj2)

	objbox.Reset()
	assert.Eq(t, 0, objbox.Len())
}
