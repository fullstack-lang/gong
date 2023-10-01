package x

import (
	"errors"
	"fmt"

	test2_models "github.com/fullstack-lang/gong/test2/go/models"
)

// errUnkownEnum is returns when a value cannot match enum values
var errUnkownEnum = errors.New("unkown enum")

// needed to avoid when fmt package is not needed by generated code
var __dummy__fmt_variable fmt.Scanner

// swagger:ignore
type __void any

// needed for creating set of instances in the stage
var __member __void

// GongStructInterface is the interface met by GongStructs
// It allows runtime reflexion of instances (without the hassle of the "reflect" package)
type GongStructInterface interface {
	GetName() (res string)
	// GetID() (res int)
	// GetFields() (res []string)
	// GetFieldStringValue(fieldName string) (res string)
}

type StageStruct struct {
	X_As           map[*A]any
	X_As_mapString map[string]*A

	// each imported packages to the x package is compiled to
	// a field to this stage struct of the imported package
	// name deconflixion relies on the aliasing in the import
	//
	// packages imports within the module construe a directed acyclic graph (DAG)
	// where each package is represented only once
	// when a commit is performed, the first stage of the commit is performed
	// recursively on all depdendant stages in the DAG. One stage can be present
	// reached through multiple path but it performs the local first stage commit
	// only once
	//
	GONG_Imported_Package_test2_models *test2_models.StageStruct
}

func NewStage(path string) (stage *StageStruct) {

	stage.GONG_Imported_Package_test2_models = test2_models.NewStage(path)

	return new(StageStruct)
}
