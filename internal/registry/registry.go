package registry

import (
	"github.com/iamgaru/gspotty"
	"github.com/iamgaru/human-presence-sim/internal/core"
)

var All = map[string]func() core.Plugin{
	"gspotty": func() core.Plugin { return gspotty.New() },
}
