package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/oldapp/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
