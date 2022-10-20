package objbox_test

import (
	"testing"

	"github.com/gookit/golib/objbox"
	"github.com/gookit/goutil/testutil/assert"
)

func TestBox_Add(t *testing.T) {
	box := objbox.New()
	box.Add("obj1", &testObj{Name: "inhere"})

	obj1 := box.Get("obj1").(*testObj)
	assert.Eq(t, "inhere", obj1.Name)
}
