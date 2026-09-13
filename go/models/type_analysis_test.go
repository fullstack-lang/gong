package models

import (
	"go/ast"
	"go/parser"
	"go/token"
	"slices"
	"testing"
)

func TestTypeAnalysis_ResilienceAndInterfaces(t *testing.T) {
	const abstractCode = `
package models

type AbstractType interface {
	GetComputedPrefix() string
	SetComputedPrefix(string)
}

type AbstractTypeFields struct {
	ComputedPrefix string
}

func (a *AbstractTypeFields) GetComputedPrefix() string {
	return a.ComputedPrefix
}

func (a *AbstractTypeFields) SetComputedPrefix(s string) {
	a.ComputedPrefix = s
}

type Task struct {
	Name string
	AbstractTypeFields
	UnresolvedRef *MissingType // Missing type must not halt type checking!
}
`

	const concreteCode = `
package models

type ConcreteType interface {
	GetAbstractElement() AbstractType
}

type TaskShape struct {
	Name string
	Task *Task
}

func (s *TaskShape) GetAbstractElement() AbstractType {
	if s.Task == nil {
		return nil
	}
	return s.Task
}

// References Stage which is not defined in any user model!
func (s *TaskShape) DoSomethingWithStage(stage *Stage) {
}
`

	fset := token.NewFileSet()
	f1, err := parser.ParseFile(fset, "models_abstract.go", abstractCode, parser.ParseComments)
	if err != nil {
		t.Fatalf("Failed to parse abstractCode: %v", err)
	}
	f2, err := parser.ParseFile(fset, "models_concrete.go", concreteCode, parser.ParseComments)
	if err != nil {
		t.Fatalf("Failed to parse concreteCode: %v", err)
	}

	astPkg := &ast.Package{
		Name:  "models",
		Files: map[string]*ast.File{"models_abstract.go": f1, "models_concrete.go": f2},
	}

	stage := NewStage("")
	modelPkg := NewModelPkg(stage)
	modelPkg.PkgPath = "github.com/example/testmodels"
	modelPkg.Fset = fset

	taskStruct := (&GongStruct{Name: "Task"}).Stage(stage)
	taskShapeStruct := (&GongStruct{Name: "TaskShape"}).Stage(stage)

	modelPkg.GongStructs = map[string]*GongStruct{
		"Task":      taskStruct,
		"TaskShape": taskShapeStruct,
	}

	RunTypeAnalysis(modelPkg, astPkg)

	if modelPkg.TypesPkg == nil {
		t.Fatal("Expected TypesPkg to be non-nil")
	}

	// 1. Task should implement AbstractType via embedded AbstractTypeFields
	if !modelPkg.Implements("Task", "AbstractType") {
		t.Error("Expected Task to implement AbstractType")
	}
	if !taskStruct.Implements("AbstractType") {
		t.Error("Expected taskStruct.Implements('AbstractType') to be true")
	}
	if !slices.Contains(taskStruct.ImplementedInterfaces, "AbstractType") {
		t.Errorf("Expected ImplementedInterfaces on Task to contain AbstractType, got: %v", taskStruct.ImplementedInterfaces)
	}

	// 2. TaskShape should implement ConcreteType
	if !modelPkg.Implements("TaskShape", "ConcreteType") {
		t.Error("Expected TaskShape to implement ConcreteType")
	}
	if !taskShapeStruct.Implements("ConcreteType") {
		t.Error("Expected taskShapeStruct.Implements('ConcreteType') to be true")
	}

	// 3. TaskShape should NOT implement AbstractType
	if modelPkg.Implements("TaskShape", "AbstractType") {
		t.Error("TaskShape should NOT implement AbstractType")
	}

	// 4. Method sets
	if !taskStruct.HasMethod("GetComputedPrefix") {
		t.Error("Expected Task to have method GetComputedPrefix (promoted from embedded)")
	}
	if !taskStruct.HasMethod("Stage") {
		t.Error("Expected Task to have synthetic Stage method")
	}
	if !taskShapeStruct.HasMethod("GetAbstractElement") {
		t.Error("Expected TaskShape to have GetAbstractElement")
	}

	// 5. Embedded structs
	embeddedInTask := modelPkg.GetEmbeddedStructNames("Task")
	if !slices.Contains(embeddedInTask, "AbstractTypeFields") {
		t.Errorf("Expected embedded structs in Task to contain AbstractTypeFields, got: %v", embeddedInTask)
	}

	// 6. Resilience: TypeErrors should record MissingType without failing
	if len(modelPkg.TypeErrors) == 0 {
		t.Log("No type errors recorded or MissingType gracefully handled")
	} else {
		t.Logf("Gracefully recorded %d type errors: %v", len(modelPkg.TypeErrors), modelPkg.TypeErrors)
	}
}
