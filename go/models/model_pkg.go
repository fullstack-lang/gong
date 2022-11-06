package models

import "embed"

// ModelPkg is the go package where the gong source code is located
//
// It contains the list of GongStructs & GongEnum
type ModelPkg struct {
	Name    string // should be "models"
	PkgPath string // for instance "github.com/.../models"

	GongStructs map[string]*GongStruct `gorm:"-"` // sql3Lite does not support maps
	GongEnums   map[string]*GongEnum   `gorm:"-"`
	GongNotes   map[string]*GongNote   `gorm:"-"`
}

// dir, initialized with a //go:embed directive, is the root
// the embedded source code
// usualy, it embeds go/models go/diagrams
func LoadEmbedded(dir embed.FS) (modelPkg *ModelPkg, err error) {

	modelPkg = &ModelPkg{}

	// since the source is embedded, one needs to
	// compute the Abstract syntax tree in a special manner
	pkgs := ParseEmbedModel(dir, "go/models")

	WalkParser(pkgs, modelPkg)
	modelPkg.SerializeToStage()
	Stage.Commit()

	return modelPkg, nil
}

func LoadSource(pkgPath string) (modelPkg *ModelPkg, err error) {

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
	})

	Walk(pkgPath, modelPkg)

	modelPkg.SerializeToStage()

	return modelPkg, nil
}

// SerializeToStage stages modelPkg and
// recursively stage all structs and all fields of all structs
func (modelPkg *ModelPkg) SerializeToStage() {
	modelPkg.Stage()
	for _, gongStruct := range modelPkg.GongStructs {
		gongStruct.Stage()

		for _, field := range gongStruct.Fields {
			switch field := field.(type) {
			case *GongBasicField:

				field.Stage()
				gongStruct.GongBasicFields = append(gongStruct.GongBasicFields, field)

			case *GongTimeField:
				field.Stage()
				gongStruct.GongTimeFields = append(gongStruct.GongTimeFields, field)

			case *PointerToGongStructField:
				field.Stage()
				gongStruct.PointerToGongStructFields = append(gongStruct.PointerToGongStructFields, field)

			case *SliceOfPointerToGongStructField:
				field.Stage()
				gongStruct.SliceOfPointerToGongStructFields = append(gongStruct.SliceOfPointerToGongStructFields,
					field)
			}
		}

	}
	for _, gongEnum := range modelPkg.GongEnums {
		gongEnum.Stage()

		for _, gongEnumValue := range gongEnum.GongEnumValues {
			gongEnumValue.Stage()
		}
	}
	for _, gongNote := range modelPkg.GongNotes {
		gongNote.Stage()
	}
	Stage.Commit()
}
