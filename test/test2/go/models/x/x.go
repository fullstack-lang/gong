package x

import (
	"github.com/fullstack-lang/gong/test/test2/go/models/y"
)

// A is the same identifier as in the models
// It allows for the same gong_... file generation
type A struct {
	Name string

	Y *y.Y
}
