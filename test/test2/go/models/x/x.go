package x

import (
	"github.com/fullstack-lang/gong/test/test2/go/models/y"
)

// X is the same identifier as in the models
// It allows for the same gong_... file generation
type X struct {
	Name string

	Y *y.Y
}
