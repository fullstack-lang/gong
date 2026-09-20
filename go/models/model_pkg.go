package models

import (
	"embed"
	"go/token"
	"go/types"
	"path/filepath"
	"strings"
)

// ModelPkg is the go package where the gong source code is located
//
// It contains the list of GongStructs & GongEnum
type ModelPkg struct {
	Name      string // should be "models"
	PkgGoName string // Go package name (e.g. "models", "x", "y")
	PkgPath   string // for instance "github.com/.../models"

	// Stage_ is where the ModelPkg lives
	Stage_ *Stage

	GongStructs map[string]*GongStruct `gorm:"-"` // sql3Lite does not support maps
	GongEnums   map[string]*GongEnum   `gorm:"-"`
	GongNotes   map[string]*GongNote   `gorm:"-"`

	// Fset is the token.FileSet for AST files
	// swagger:ignore
	Fset *token.FileSet `gorm:"-"`

	// TypesPkg is the type-checked Go package
	// swagger:ignore
	TypesPkg *types.Package `gorm:"-"`

	// TypesInfo holds type resolution info (Defs, Uses, Selections, Types)
	// swagger:ignore
	TypesInfo *types.Info `gorm:"-"`

	// TypeErrors holds any non-fatal type errors encountered during type analysis
	// swagger:ignore
	TypeErrors []error `gorm:"-"`

	// StageSet holds metadata if a StageSet struct is defined in the package
	// swagger:ignore
	StageSet *StageSetModel `gorm:"-"`

	// PathToGoSubDirectory for instance "/tmp"
	PathToGoSubDirectory string

	// OrmPkgGenPath is target path for orm package, for instance "/tmp/libraryorm"
	OrmPkgGenPath string

	// DbOrmPkgGenPath is target path for orm package, for instance "/tmp/libraryorm/db"
	DbOrmPkgGenPath string

	// DbOrmPkgGenPath is target path for orm package, for instance "/tmp/libraryorm/db"
	DbLiteOrmPkgGenPath string

	// DbPkgGenPath is target path for orm package, for instance "/tmp/libraryorm/db"
	DbPkgGenPath string

	// ControllersPkgGenPath is target path for controllers package
	ControllersPkgGenPath string

	// FullstackPkgGenPath is target path for Fullstack package
	FullstackPkgGenPath string

	// StackPkgGenPath is target path for Stack package
	StackPkgGenPath string

	// StackPkgGenPath is target path for the level 1 Stack package
	Level1StackPkgGenPath string

	// StaticPkgGenPath is target path for Static package
	StaticPkgGenPath string

	// ProbePkgGenPath is target path for Data package
	ProbePkgGenPath string

	// NgWorkspacePath is the path to the Ng Workspace
	NgWorkspacePath string

	// NgWorkspaceName is the name of the angular workspace
	//
	// note : initialy the name was "ng" for all angular workspace but "npm workspaces"
	// does not support workspaces that have the same name (it does not know
	// about aliasing)
	NgWorkspaceName string

	// NgDataLibrarySourceCodeDirectory is the "<pkgName>/src/lib" directory where,
	// by angular CLI convention,
	// the source code for the library's components,
	// services, modules, and other features are located.
	//
	// gong generate regenerates at each compilation a material angular library "<pkgName>" for
	// having the code related to data manipulation of objects of the
	// stack
	NgDataLibrarySourceCodeDirectory string

	// NgSpecificLibrarySourceCodeDirectory is the "<pkgName>specific/src/lib"
	// where the developper stores its specific code for the front component
	//
	// This library is generated once at the stack creation
	NgSpecificLibrarySourceCodeDirectory string

	// MaterialLibDatamodelTargetPath is the "<pkgName>datamodel/src/lib"
	MaterialLibDatamodelTargetPath string
}

func (modelPkg *ModelPkg) GetStage() (stage *Stage) {
	stage = modelPkg.Stage_
	return stage
}

func (modelPkg *ModelPkg) SetStage(stage *Stage) {
	modelPkg.Stage_ = stage
}

func NewModelPkg(stage *Stage) (modelPkg *ModelPkg) {

	modelPkg = new(ModelPkg)
	modelPkg.SetStage(stage)
	return
}

// dir, initialized with a //go:embed directive, is the root
// the embedded source code
// usualy, it embeds go/models go/diagrams
func LoadEmbedded(stage *Stage, goModelsDir embed.FS) (modelPkg *ModelPkg, err error) {

	modelPkg = NewModelPkg(stage)
	modelPkg.Fset = token.NewFileSet()

	// since the source is embedded, one needs to
	// compute the Abstract syntax tree in a special manner
	pkgs := ParseEmbedModelWithFset(goModelsDir, "models", modelPkg.Fset)

	WalkParser(pkgs, modelPkg, nil)

	modelPkg.SerializeToStage()
	stage.Commit()

	return modelPkg, nil
}

func LoadSource(stage *Stage, pkgPath string) (modelPkg *ModelPkg, err error) {

	// check existance of go.mod file in the path to the 'models' package
	//
	// if no go.mod file is found above the 'models' package, gong generate fails
	//
	// if go.mod exists, it means the package path has been defined
	// for instance "github.com/fullstack-lang/gongsvg"
	//
	// if go.mod does not exist, gong generate can only infer the package name
	// from the name of directory that is two levels above "go/models"
	// it is up to the developper to change the module name after the first gong generation
	pkgName, fullPkgPath := ComputePkgPathFromGoModFile(pkgPath)

	// initiate model package
	modelPkg = (&ModelPkg{
		Name:    pkgName,
		PkgPath: fullPkgPath,
		Stage_:  stage,
	})

	Walk(pkgPath, modelPkg)

	if modelPkg.StageSet == nil {
		_ = modelPkg.SynthesizeStageSetFromDependencies(pkgPath)
	}

	modelPkg.SerializeToStage()

	return modelPkg, nil
}

// SerializeToStage stages modelPkg and
// recursively stage all structs and all fields of all structs
func (modelPkg *ModelPkg) SerializeToStage() {
	modelPkg.Stage(modelPkg.GetStage())
	for _, gongStruct := range modelPkg.GongStructs {
		gongStruct.Stage(modelPkg.GetStage())

		for _, field := range gongStruct.Fields {
			switch field := field.(type) {
			case *GongBasicField:

				field.Stage(modelPkg.GetStage())
				gongStruct.GongBasicFields = append(gongStruct.GongBasicFields, field)

			case *GongTimeField:
				field.Stage(modelPkg.GetStage())
				gongStruct.GongTimeFields = append(gongStruct.GongTimeFields, field)

			case *PointerToGongStructField:
				field.Stage(modelPkg.GetStage())
				gongStruct.PointerToGongStructFields = append(gongStruct.PointerToGongStructFields, field)

			case *SliceOfPointerToGongStructField:
				field.Stage(modelPkg.GetStage())
				gongStruct.SliceOfPointerToGongStructFields = append(gongStruct.SliceOfPointerToGongStructFields,
					field)
			}
		}

	}
	for _, gongEnum := range modelPkg.GongEnums {
		gongEnum.Stage(modelPkg.GetStage())

		for _, gongEnumValue := range gongEnum.GongEnumValues {
			gongEnumValue.Stage(modelPkg.GetStage())
		}
	}
	for _, gongNote := range modelPkg.GongNotes {
		gongNote.Stage(modelPkg.GetStage())
	}
	modelPkg.GetStage().Commit()
}

// Implements checks if structName (by pointer or value receiver) implements interfaceName.
// interfaceName can be unqualified ("MyInterface") or qualified ("package.MyInterface").
func (modelPkg *ModelPkg) Implements(structName string, interfaceName string) bool {
	if modelPkg.TypesPkg == nil {
		return false
	}
	obj := modelPkg.TypesPkg.Scope().Lookup(structName)
	if obj == nil {
		return false
	}
	named, ok := obj.Type().(*types.Named)
	if !ok {
		return false
	}

	var iface *types.Interface
	if strings.Contains(interfaceName, ".") {
		parts := strings.Split(interfaceName, ".")
		pkgName := parts[0]
		targetIface := parts[1]
		for _, imp := range modelPkg.TypesPkg.Imports() {
			if imp.Name() == pkgName || imp.Path() == pkgName || filepath.Base(imp.Path()) == pkgName {
				if obj := imp.Scope().Lookup(targetIface); obj != nil {
					if it, ok := obj.Type().Underlying().(*types.Interface); ok {
						iface = it
						break
					}
				}
			}
		}
	} else {
		if obj := modelPkg.TypesPkg.Scope().Lookup(interfaceName); obj != nil {
			if it, ok := obj.Type().Underlying().(*types.Interface); ok {
				iface = it
			}
		}
	}

	if iface == nil {
		return false
	}

	// Check pointer receiver first (most common in Gong structs)
	if types.Implements(types.NewPointer(named), iface) {
		return true
	}
	return types.Implements(named, iface)
}

// GetMethodSet returns the method set of *typeName (pointer receiver, including promoted methods).
func (modelPkg *ModelPkg) GetMethodSet(typeName string) *types.MethodSet {
	if modelPkg.TypesPkg == nil {
		return nil
	}
	obj := modelPkg.TypesPkg.Scope().Lookup(typeName)
	if obj == nil {
		return nil
	}
	named, ok := obj.Type().(*types.Named)
	if !ok {
		return nil
	}
	return types.NewMethodSet(types.NewPointer(named))
}

// GetEmbeddedStructNames returns the names of all structs embedded in structName (both direct and transitive).
func (modelPkg *ModelPkg) GetEmbeddedStructNames(structName string) []string {
	if modelPkg.TypesPkg == nil {
		return nil
	}
	obj := modelPkg.TypesPkg.Scope().Lookup(structName)
	if obj == nil {
		return nil
	}
	named, ok := obj.Type().(*types.Named)
	if !ok {
		return nil
	}
	st, ok := named.Underlying().(*types.Struct)
	if !ok {
		return nil
	}

	var embedded []string
	var collect func(s *types.Struct)
	collect = func(s *types.Struct) {
		for f := range s.Fields() {
			f := f
			if f.Anonymous() {
				embedded = append(embedded, f.Name())
				// Check transitive embedding
				t := f.Type()
				if ptr, ok := t.(*types.Pointer); ok {
					t = ptr.Elem()
				}
				if n, ok := t.(*types.Named); ok {
					if innerSt, ok := n.Underlying().(*types.Struct); ok {
						collect(innerSt)
					}
				}
			}
		}
	}
	collect(st)
	return embedded
}

// GetImplementedInterfaces returns all package-level interfaces satisfied by structName.
func (modelPkg *ModelPkg) GetImplementedInterfaces(structName string) []string {
	if modelPkg.TypesPkg == nil {
		return nil
	}
	var implemented []string
	for _, name := range modelPkg.TypesPkg.Scope().Names() {
		obj := modelPkg.TypesPkg.Scope().Lookup(name)
		if obj == nil {
			continue
		}
		if _, ok := obj.Type().Underlying().(*types.Interface); ok {
			if modelPkg.Implements(structName, name) {
				implemented = append(implemented, name)
			}
		}
	}
	return implemented
}

// LookupType looks up a named type in the model package scope.
func (modelPkg *ModelPkg) LookupType(name string) types.Type {
	if modelPkg.TypesPkg == nil {
		return nil
	}
	obj := modelPkg.TypesPkg.Scope().Lookup(name)
	if obj == nil {
		return nil
	}
	return obj.Type()
}
