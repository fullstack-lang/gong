package x

import (
	"github.com/fullstack-lang/gong/test/test2/go/models"
)

// A is the same identifier as in the models
// It allows for the same gong_... file generation
type A struct {
	Name string

	A *models.A

	// Y_A *y.A
}
