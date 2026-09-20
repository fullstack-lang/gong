package models

import (
	"github.com/fullstack-lang/gong/test/test2/go/models/x"
	"github.com/fullstack-lang/gong/test/test2/go/models/y"
)

// StageSet coordinates multiple stages across packages
type StageSet struct {
	Stage  *Stage
	XStage *x.Stage
	YStage *y.Stage
}
