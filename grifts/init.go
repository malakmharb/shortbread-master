package grifts

import (
	"github.com/daylightdata/shortbread/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
