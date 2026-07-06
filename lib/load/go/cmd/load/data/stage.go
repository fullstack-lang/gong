package main

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/lib/load/go/models"
	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var (
	_ time.Time
	_ = slices.Index[[]int, int]
)

// function will stage objects
func _(stage *models.Stage) {

	// insertion point for declaration of instances to stage

	__FileToUpload__00000000_ := (&models.FileToUpload{Name: `erze.txt`}).Stage(stage)

	// insertion point for initialization of values

	__FileToUpload__00000000_.Name = `erze.txt`
	__FileToUpload__00000000_.Base64EncodedContent = `ZXJ6`

	// insertion point for setup of pointers
}
