//go:build go1.18
// +build go1.18

package optional_test

import (
	"testing"

	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/stdutil/optional"
)

func TestOptFactory_of(t *testing.T) {
	opt := optional.Of(nil)

	dump.P(opt.OrElseGet(34))
}
