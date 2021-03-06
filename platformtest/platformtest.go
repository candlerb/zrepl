package platformtest

import (
	"context"
	"fmt"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type Context struct {
	context.Context
	RootDataset string
}

var FailNowSentinel = fmt.Errorf("platformtest: FailNow called on context")

var SkipNowSentinel = fmt.Errorf("platformtest: SkipNow called on context")

var _ assert.TestingT = (*Context)(nil)
var _ require.TestingT = (*Context)(nil)

func (c *Context) Logf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	GetLog(c).Info(msg)
}

func (c *Context) Errorf(format string, args ...interface{}) {
	GetLog(c).Printf(format, args...)
	c.FailNow()
}

func (c *Context) FailNow() {
	panic(FailNowSentinel)
}

func (c *Context) SkipNow() {
	panic(SkipNowSentinel)
}
