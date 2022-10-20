//go:build go1.18

package optional_test

import (
	"testing"

	"github.com/gookit/golib/optional"
	"github.com/gookit/goutil/dump"
)

func TestOptFactory_of(t *testing.T) {
	opt := optional.Of(0)

	var a int
	a = opt.OrElseGet(34)

	dump.P(a, opt.OrElseGet(34))
}
