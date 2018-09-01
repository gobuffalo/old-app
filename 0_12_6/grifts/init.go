package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/oldapp/0_12_6/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
