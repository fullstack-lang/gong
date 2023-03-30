package models

import "embed"

// ModelPkg is the go package where the gong source code is located
//
// It contains the list of GongStructs & GongEnum
type ModelPkg struct {
	Name    string // should be "models"
	PkgPath string // for instance "github.com/.../models"

	// Stage_ is where the ModelPkg lives
	Stage_ *StageStruct

	GongStructs map[string]*GongStruct `gorm:"-"` // sql3Lite does not support maps
	GongEnums   map[string]*GongEnum   `gorm:"-"`
	GongNotes   map[string]*GongNote   `gorm:"-"`
}

func (modelPkg *ModelPkg) GetStage() (stage *StageStruct) {
	stage = modelPkg.Stage_
	return stage
}

func (modelPkg *ModelPkg) SetStage(stage *StageStruct) {
	modelPkg.Stage_ = stage
}

func NewModelPkg(stage *StageStruct) (modelPkg *ModelPkg) {

	modelPkg = new(ModelPkg)
	modelPkg.SetStage(stage)
	return
}

// dir, initialized with a //go:embed directive, is the root
// the embedded source code
// usualy, it embeds go/models go/diagrams
func LoadEmbedded(stage *StageStruct, goModelsDir embed.FS) (modelPkg *ModelPkg, err error) {

	modelPkg = NewModelPkg(stage)

	// since the source is embedded, one needs to
	// compute the Abstract syntax tree in a special manner
	pkgs := ParseEmbedModel(goModelsDir, "models")

	WalkParser(pkgs, modelPkg)
	// fetch meta information
	inspectMeta(stage, pkgs["models"])

	modelPkg.SerializeToStage()
	stage.Commit()

	return modelPkg, nil
}

func LoadSource(stage *StageStruct, pkgPath string) (modelPkg *ModelPkg, err error) {

	// check existance of go.mod file in the path to the 'models' package
	//
	// if no go.mod file is found above the 'models' package, gongc fails
	//
	// if go.mod exists, it means the package path has been defined
	// for instance "github.com/fullstack-lang/gongsvg"
	//
	// if go.mod does not exist, gongc can only infer the package name
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
