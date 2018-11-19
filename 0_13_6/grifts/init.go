package grifts

import (
	"github.com/gobuffalo/0_13_6/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
