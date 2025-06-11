// generated code - do not edit
package models

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/sergi/go-diff/diffmatchpatch"
)

const marshallRes = `package {{PackageName}}

import (
	"time"

	"{{ModelsPackageName}}"
	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// _ point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.Stage) {

	const __write__local_time = "{{LocalTimeStamp}}"
	const __write__utc_time__ = "{{UTCTimeStamp}}"

	const __commitId__ = "{{CommitId}}"

	// Declaration of instances to stage{{Identifiers}}

	// Setup of values{{ValueInitializers}}

	// Setup of pointers{{PointersInitializers}}
}
`

const IdentifiersDecls = `
	{{Identifier}} := (&models.{{GeneratedStructName}}{}).Stage(stage)`

const StringInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = ` + "`" + `{{GeneratedFieldNameValue}}` + "`"

const MetaFieldStructInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = ` + `{{GeneratedFieldNameValue}}`

const StringEnumInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const NumberInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const PointerFieldInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const SliceOfPointersFieldInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = append({{Identifier}}.{{GeneratedFieldName}}, {{GeneratedFieldNameValue}})`

const TimeInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}}, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "{{GeneratedFieldNameValue}}")`

// Marshall marshall the stage content into the file as an instanciation into a stage
func (stage *Stage) Marshall(file *os.File, modelsPackageName, packageName string) {

	name := file.Name()

	if !strings.HasSuffix(name, ".go") {
		log.Fatalln(name + " is not a go filename")
	}

	log.Printf("%s Marshalling %s", time.Now().Format("2006-01-02 15:04:05.000000"), name)
	newBase := filepath.Base(file.Name())

	res := marshallRes
	res = strings.ReplaceAll(res, "{{databaseName}}", strings.ReplaceAll(newBase, ".go", ""))
	res = strings.ReplaceAll(res, "{{PackageName}}", packageName)
	res = strings.ReplaceAll(res, "{{ModelsPackageName}}", modelsPackageName)

	// map of identifiers
	// var StageMapDstructIds map[*Dstruct]string
	identifiersDecl := ""
	initializerStatements := ""
	pointersInitializesStatements := ""

	id := ""
	_ = id
	decl := ""
	_ = decl
	setValueField := ""
	_ = setValueField

	// insertion initialization of objects to stage
	map_Classdiagram_Identifiers := make(map[*Classdiagram]string)
	_ = map_Classdiagram_Identifiers

	classdiagramOrdered := []*Classdiagram{}
	for classdiagram := range stage.Classdiagrams {
		classdiagramOrdered = append(classdiagramOrdered, classdiagram)
	}
	sort.Slice(classdiagramOrdered[:], func(i, j int) bool {
		classdiagrami := classdiagramOrdered[i]
		classdiagramj := classdiagramOrdered[j]
		classdiagrami_order, oki := stage.ClassdiagramMap_Staged_Order[classdiagrami]
		classdiagramj_order, okj := stage.ClassdiagramMap_Staged_Order[classdiagramj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return classdiagrami_order < classdiagramj_order
	})
	if len(classdiagramOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, classdiagram := range classdiagramOrdered {

		id = generatesIdentifier("Classdiagram", idx, classdiagram.Name)
		map_Classdiagram_Identifiers[classdiagram] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Classdiagram")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", classdiagram.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(classdiagram.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsInDrawMode")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", classdiagram.IsInDrawMode))
		initializerStatements += setValueField

	}

	map_DiagramPackage_Identifiers := make(map[*DiagramPackage]string)
	_ = map_DiagramPackage_Identifiers

	diagrampackageOrdered := []*DiagramPackage{}
	for diagrampackage := range stage.DiagramPackages {
		diagrampackageOrdered = append(diagrampackageOrdered, diagrampackage)
	}
	sort.Slice(diagrampackageOrdered[:], func(i, j int) bool {
		diagrampackagei := diagrampackageOrdered[i]
		diagrampackagej := diagrampackageOrdered[j]
		diagrampackagei_order, oki := stage.DiagramPackageMap_Staged_Order[diagrampackagei]
		diagrampackagej_order, okj := stage.DiagramPackageMap_Staged_Order[diagrampackagej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return diagrampackagei_order < diagrampackagej_order
	})
	if len(diagrampackageOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, diagrampackage := range diagrampackageOrdered {

		id = generatesIdentifier("DiagramPackage", idx, diagrampackage.Name)
		map_DiagramPackage_Identifiers[diagrampackage] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DiagramPackage")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", diagrampackage.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagrampackage.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Path")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagrampackage.Path))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "GongModelPath")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagrampackage.GongModelPath))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsEditable")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagrampackage.IsEditable))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsReloaded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagrampackage.IsReloaded))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "AbsolutePathToDiagramPackage")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagrampackage.AbsolutePathToDiagramPackage))
		initializerStatements += setValueField

	}

	map_Field_Identifiers := make(map[*Field]string)
	_ = map_Field_Identifiers

	fieldOrdered := []*Field{}
	for field := range stage.Fields {
		fieldOrdered = append(fieldOrdered, field)
	}
	sort.Slice(fieldOrdered[:], func(i, j int) bool {
		fieldi := fieldOrdered[i]
		fieldj := fieldOrdered[j]
		fieldi_order, oki := stage.FieldMap_Staged_Order[fieldi]
		fieldj_order, okj := stage.FieldMap_Staged_Order[fieldj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return fieldi_order < fieldj_order
	})
	if len(fieldOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, field := range fieldOrdered {

		id = generatesIdentifier("Field", idx, field.Name)
		map_Field_Identifiers[field] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Field")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", field.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(field.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Identifier")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(field.Identifier))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FieldTypeAsString")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(field.FieldTypeAsString))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Structname")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(field.Structname))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Fieldtypename")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(field.Fieldtypename))
		initializerStatements += setValueField

	}

	map_GongEnumShape_Identifiers := make(map[*GongEnumShape]string)
	_ = map_GongEnumShape_Identifiers

	gongenumshapeOrdered := []*GongEnumShape{}
	for gongenumshape := range stage.GongEnumShapes {
		gongenumshapeOrdered = append(gongenumshapeOrdered, gongenumshape)
	}
	sort.Slice(gongenumshapeOrdered[:], func(i, j int) bool {
		gongenumshapei := gongenumshapeOrdered[i]
		gongenumshapej := gongenumshapeOrdered[j]
		gongenumshapei_order, oki := stage.GongEnumShapeMap_Staged_Order[gongenumshapei]
		gongenumshapej_order, okj := stage.GongEnumShapeMap_Staged_Order[gongenumshapej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return gongenumshapei_order < gongenumshapej_order
	})
	if len(gongenumshapeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, gongenumshape := range gongenumshapeOrdered {

		id = generatesIdentifier("GongEnumShape", idx, gongenumshape.Name)
		map_GongEnumShape_Identifiers[gongenumshape] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "GongEnumShape")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", gongenumshape.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongenumshape.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Identifier")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongenumshape.Identifier))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Width")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gongenumshape.Width))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Height")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gongenumshape.Height))
		initializerStatements += setValueField

	}

	map_GongEnumValueEntry_Identifiers := make(map[*GongEnumValueEntry]string)
	_ = map_GongEnumValueEntry_Identifiers

	gongenumvalueentryOrdered := []*GongEnumValueEntry{}
	for gongenumvalueentry := range stage.GongEnumValueEntrys {
		gongenumvalueentryOrdered = append(gongenumvalueentryOrdered, gongenumvalueentry)
	}
	sort.Slice(gongenumvalueentryOrdered[:], func(i, j int) bool {
		gongenumvalueentryi := gongenumvalueentryOrdered[i]
		gongenumvalueentryj := gongenumvalueentryOrdered[j]
		gongenumvalueentryi_order, oki := stage.GongEnumValueEntryMap_Staged_Order[gongenumvalueentryi]
		gongenumvalueentryj_order, okj := stage.GongEnumValueEntryMap_Staged_Order[gongenumvalueentryj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return gongenumvalueentryi_order < gongenumvalueentryj_order
	})
	if len(gongenumvalueentryOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, gongenumvalueentry := range gongenumvalueentryOrdered {

		id = generatesIdentifier("GongEnumValueEntry", idx, gongenumvalueentry.Name)
		map_GongEnumValueEntry_Identifiers[gongenumvalueentry] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "GongEnumValueEntry")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", gongenumvalueentry.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongenumvalueentry.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Identifier")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongenumvalueentry.Identifier))
		initializerStatements += setValueField

	}

	map_GongStructShape_Identifiers := make(map[*GongStructShape]string)
	_ = map_GongStructShape_Identifiers

	gongstructshapeOrdered := []*GongStructShape{}
	for gongstructshape := range stage.GongStructShapes {
		gongstructshapeOrdered = append(gongstructshapeOrdered, gongstructshape)
	}
	sort.Slice(gongstructshapeOrdered[:], func(i, j int) bool {
		gongstructshapei := gongstructshapeOrdered[i]
		gongstructshapej := gongstructshapeOrdered[j]
		gongstructshapei_order, oki := stage.GongStructShapeMap_Staged_Order[gongstructshapei]
		gongstructshapej_order, okj := stage.GongStructShapeMap_Staged_Order[gongstructshapej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return gongstructshapei_order < gongstructshapej_order
	})
	if len(gongstructshapeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, gongstructshape := range gongstructshapeOrdered {

		id = generatesIdentifier("GongStructShape", idx, gongstructshape.Name)
		map_GongStructShape_Identifiers[gongstructshape] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "GongStructShape")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", gongstructshape.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongstructshape.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Identifier")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongstructshape.Identifier))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ShowNbInstances")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", gongstructshape.ShowNbInstances))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "NbInstances")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", gongstructshape.NbInstances))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Width")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gongstructshape.Width))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Height")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gongstructshape.Height))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsSelected")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", gongstructshape.IsSelected))
		initializerStatements += setValueField

	}

	map_Link_Identifiers := make(map[*Link]string)
	_ = map_Link_Identifiers

	linkOrdered := []*Link{}
	for link := range stage.Links {
		linkOrdered = append(linkOrdered, link)
	}
	sort.Slice(linkOrdered[:], func(i, j int) bool {
		linki := linkOrdered[i]
		linkj := linkOrdered[j]
		linki_order, oki := stage.LinkMap_Staged_Order[linki]
		linkj_order, okj := stage.LinkMap_Staged_Order[linkj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return linki_order < linkj_order
	})
	if len(linkOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, link := range linkOrdered {

		id = generatesIdentifier("Link", idx, link.Name)
		map_Link_Identifiers[link] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Link")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", link.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(link.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Identifier")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(link.Identifier))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Fieldtypename")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(link.Fieldtypename))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FieldOffsetX")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", link.FieldOffsetX))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FieldOffsetY")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", link.FieldOffsetY))
		initializerStatements += setValueField

		if link.TargetMultiplicity != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "TargetMultiplicity")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+link.TargetMultiplicity.ToCodeString())
			initializerStatements += setValueField
		}

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "TargetMultiplicityOffsetX")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", link.TargetMultiplicityOffsetX))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "TargetMultiplicityOffsetY")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", link.TargetMultiplicityOffsetY))
		initializerStatements += setValueField

		if link.SourceMultiplicity != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "SourceMultiplicity")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+link.SourceMultiplicity.ToCodeString())
			initializerStatements += setValueField
		}

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "SourceMultiplicityOffsetX")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", link.SourceMultiplicityOffsetX))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "SourceMultiplicityOffsetY")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", link.SourceMultiplicityOffsetY))
		initializerStatements += setValueField

		if link.StartOrientation != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StartOrientation")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+link.StartOrientation.ToCodeString())
			initializerStatements += setValueField
		}

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StartRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", link.StartRatio))
		initializerStatements += setValueField

		if link.EndOrientation != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EndOrientation")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+link.EndOrientation.ToCodeString())
			initializerStatements += setValueField
		}

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EndRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", link.EndRatio))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CornerOffsetRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", link.CornerOffsetRatio))
		initializerStatements += setValueField

	}

	map_NoteShape_Identifiers := make(map[*NoteShape]string)
	_ = map_NoteShape_Identifiers

	noteshapeOrdered := []*NoteShape{}
	for noteshape := range stage.NoteShapes {
		noteshapeOrdered = append(noteshapeOrdered, noteshape)
	}
	sort.Slice(noteshapeOrdered[:], func(i, j int) bool {
		noteshapei := noteshapeOrdered[i]
		noteshapej := noteshapeOrdered[j]
		noteshapei_order, oki := stage.NoteShapeMap_Staged_Order[noteshapei]
		noteshapej_order, okj := stage.NoteShapeMap_Staged_Order[noteshapej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return noteshapei_order < noteshapej_order
	})
	if len(noteshapeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, noteshape := range noteshapeOrdered {

		id = generatesIdentifier("NoteShape", idx, noteshape.Name)
		map_NoteShape_Identifiers[noteshape] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "NoteShape")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", noteshape.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(noteshape.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Identifier")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(noteshape.Identifier))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Body")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(noteshape.Body))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "BodyHTML")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(noteshape.BodyHTML))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "X")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", noteshape.X))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Y")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", noteshape.Y))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Width")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", noteshape.Width))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Height")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", noteshape.Height))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Matched")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", noteshape.Matched))
		initializerStatements += setValueField

	}

	map_NoteShapeLink_Identifiers := make(map[*NoteShapeLink]string)
	_ = map_NoteShapeLink_Identifiers

	noteshapelinkOrdered := []*NoteShapeLink{}
	for noteshapelink := range stage.NoteShapeLinks {
		noteshapelinkOrdered = append(noteshapelinkOrdered, noteshapelink)
	}
	sort.Slice(noteshapelinkOrdered[:], func(i, j int) bool {
		noteshapelinki := noteshapelinkOrdered[i]
		noteshapelinkj := noteshapelinkOrdered[j]
		noteshapelinki_order, oki := stage.NoteShapeLinkMap_Staged_Order[noteshapelinki]
		noteshapelinkj_order, okj := stage.NoteShapeLinkMap_Staged_Order[noteshapelinkj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return noteshapelinki_order < noteshapelinkj_order
	})
	if len(noteshapelinkOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, noteshapelink := range noteshapelinkOrdered {

		id = generatesIdentifier("NoteShapeLink", idx, noteshapelink.Name)
		map_NoteShapeLink_Identifiers[noteshapelink] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "NoteShapeLink")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", noteshapelink.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(noteshapelink.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Identifier")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(noteshapelink.Identifier))
		initializerStatements += setValueField

		if noteshapelink.Type != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Type")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+noteshapelink.Type.ToCodeString())
			initializerStatements += setValueField
		}

	}

	map_Position_Identifiers := make(map[*Position]string)
	_ = map_Position_Identifiers

	positionOrdered := []*Position{}
	for position := range stage.Positions {
		positionOrdered = append(positionOrdered, position)
	}
	sort.Slice(positionOrdered[:], func(i, j int) bool {
		positioni := positionOrdered[i]
		positionj := positionOrdered[j]
		positioni_order, oki := stage.PositionMap_Staged_Order[positioni]
		positionj_order, okj := stage.PositionMap_Staged_Order[positionj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return positioni_order < positionj_order
	})
	if len(positionOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, position := range positionOrdered {

		id = generatesIdentifier("Position", idx, position.Name)
		map_Position_Identifiers[position] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Position")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", position.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "X")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", position.X))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Y")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", position.Y))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(position.Name))
		initializerStatements += setValueField

	}

	map_UmlState_Identifiers := make(map[*UmlState]string)
	_ = map_UmlState_Identifiers

	umlstateOrdered := []*UmlState{}
	for umlstate := range stage.UmlStates {
		umlstateOrdered = append(umlstateOrdered, umlstate)
	}
	sort.Slice(umlstateOrdered[:], func(i, j int) bool {
		umlstatei := umlstateOrdered[i]
		umlstatej := umlstateOrdered[j]
		umlstatei_order, oki := stage.UmlStateMap_Staged_Order[umlstatei]
		umlstatej_order, okj := stage.UmlStateMap_Staged_Order[umlstatej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return umlstatei_order < umlstatej_order
	})
	if len(umlstateOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, umlstate := range umlstateOrdered {

		id = generatesIdentifier("UmlState", idx, umlstate.Name)
		map_UmlState_Identifiers[umlstate] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "UmlState")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", umlstate.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(umlstate.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "X")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", umlstate.X))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Y")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", umlstate.Y))
		initializerStatements += setValueField

	}

	map_Umlsc_Identifiers := make(map[*Umlsc]string)
	_ = map_Umlsc_Identifiers

	umlscOrdered := []*Umlsc{}
	for umlsc := range stage.Umlscs {
		umlscOrdered = append(umlscOrdered, umlsc)
	}
	sort.Slice(umlscOrdered[:], func(i, j int) bool {
		umlsci := umlscOrdered[i]
		umlscj := umlscOrdered[j]
		umlsci_order, oki := stage.UmlscMap_Staged_Order[umlsci]
		umlscj_order, okj := stage.UmlscMap_Staged_Order[umlscj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return umlsci_order < umlscj_order
	})
	if len(umlscOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, umlsc := range umlscOrdered {

		id = generatesIdentifier("Umlsc", idx, umlsc.Name)
		map_Umlsc_Identifiers[umlsc] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Umlsc")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", umlsc.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(umlsc.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Activestate")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(umlsc.Activestate))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsInDrawMode")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", umlsc.IsInDrawMode))
		initializerStatements += setValueField

	}

	map_Vertice_Identifiers := make(map[*Vertice]string)
	_ = map_Vertice_Identifiers

	verticeOrdered := []*Vertice{}
	for vertice := range stage.Vertices {
		verticeOrdered = append(verticeOrdered, vertice)
	}
	sort.Slice(verticeOrdered[:], func(i, j int) bool {
		verticei := verticeOrdered[i]
		verticej := verticeOrdered[j]
		verticei_order, oki := stage.VerticeMap_Staged_Order[verticei]
		verticej_order, okj := stage.VerticeMap_Staged_Order[verticej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return verticei_order < verticej_order
	})
	if len(verticeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, vertice := range verticeOrdered {

		id = generatesIdentifier("Vertice", idx, vertice.Name)
		map_Vertice_Identifiers[vertice] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Vertice")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", vertice.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "X")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", vertice.X))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Y")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", vertice.Y))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(vertice.Name))
		initializerStatements += setValueField

	}

	// insertion initialization of objects to stage
	if len(classdiagramOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Classdiagram instances pointers"
	}
	for idx, classdiagram := range classdiagramOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Classdiagram", idx, classdiagram.Name)
		map_Classdiagram_Identifiers[classdiagram] = id

		// Initialisation of values
		for _, _gongstructshape := range classdiagram.GongStructShapes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "GongStructShapes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_GongStructShape_Identifiers[_gongstructshape])
			pointersInitializesStatements += setPointerField
		}

		for _, _gongenumshape := range classdiagram.GongEnumShapes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "GongEnumShapes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_GongEnumShape_Identifiers[_gongenumshape])
			pointersInitializesStatements += setPointerField
		}

		for _, _noteshape := range classdiagram.NoteShapes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "NoteShapes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_NoteShape_Identifiers[_noteshape])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(diagrampackageOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of DiagramPackage instances pointers"
	}
	for idx, diagrampackage := range diagrampackageOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("DiagramPackage", idx, diagrampackage.Name)
		map_DiagramPackage_Identifiers[diagrampackage] = id

		// Initialisation of values
		for _, _classdiagram := range diagrampackage.Classdiagrams {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Classdiagrams")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Classdiagram_Identifiers[_classdiagram])
			pointersInitializesStatements += setPointerField
		}

		if diagrampackage.SelectedClassdiagram != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SelectedClassdiagram")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Classdiagram_Identifiers[diagrampackage.SelectedClassdiagram])
			pointersInitializesStatements += setPointerField
		}

		for _, _umlsc := range diagrampackage.Umlscs {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Umlscs")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Umlsc_Identifiers[_umlsc])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(fieldOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Field instances pointers"
	}
	for idx, field := range fieldOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Field", idx, field.Name)
		map_Field_Identifiers[field] = id

		// Initialisation of values
	}

	if len(gongenumshapeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of GongEnumShape instances pointers"
	}
	for idx, gongenumshape := range gongenumshapeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("GongEnumShape", idx, gongenumshape.Name)
		map_GongEnumShape_Identifiers[gongenumshape] = id

		// Initialisation of values
		if gongenumshape.Position != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Position")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Position_Identifiers[gongenumshape.Position])
			pointersInitializesStatements += setPointerField
		}

		for _, _gongenumvalueentry := range gongenumshape.GongEnumValueEntrys {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "GongEnumValueEntrys")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_GongEnumValueEntry_Identifiers[_gongenumvalueentry])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(gongenumvalueentryOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of GongEnumValueEntry instances pointers"
	}
	for idx, gongenumvalueentry := range gongenumvalueentryOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("GongEnumValueEntry", idx, gongenumvalueentry.Name)
		map_GongEnumValueEntry_Identifiers[gongenumvalueentry] = id

		// Initialisation of values
	}

	if len(gongstructshapeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of GongStructShape instances pointers"
	}
	for idx, gongstructshape := range gongstructshapeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("GongStructShape", idx, gongstructshape.Name)
		map_GongStructShape_Identifiers[gongstructshape] = id

		// Initialisation of values
		if gongstructshape.Position != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Position")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Position_Identifiers[gongstructshape.Position])
			pointersInitializesStatements += setPointerField
		}

		for _, _field := range gongstructshape.Fields {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Fields")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Field_Identifiers[_field])
			pointersInitializesStatements += setPointerField
		}

		for _, _link := range gongstructshape.Links {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Links")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Link_Identifiers[_link])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(linkOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Link instances pointers"
	}
	for idx, link := range linkOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Link", idx, link.Name)
		map_Link_Identifiers[link] = id

		// Initialisation of values
		if link.Middlevertice != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Middlevertice")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Vertice_Identifiers[link.Middlevertice])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(noteshapeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of NoteShape instances pointers"
	}
	for idx, noteshape := range noteshapeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("NoteShape", idx, noteshape.Name)
		map_NoteShape_Identifiers[noteshape] = id

		// Initialisation of values
		for _, _noteshapelink := range noteshape.NoteShapeLinks {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "NoteShapeLinks")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_NoteShapeLink_Identifiers[_noteshapelink])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(noteshapelinkOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of NoteShapeLink instances pointers"
	}
	for idx, noteshapelink := range noteshapelinkOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("NoteShapeLink", idx, noteshapelink.Name)
		map_NoteShapeLink_Identifiers[noteshapelink] = id

		// Initialisation of values
	}

	if len(positionOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Position instances pointers"
	}
	for idx, position := range positionOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Position", idx, position.Name)
		map_Position_Identifiers[position] = id

		// Initialisation of values
	}

	if len(umlstateOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of UmlState instances pointers"
	}
	for idx, umlstate := range umlstateOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("UmlState", idx, umlstate.Name)
		map_UmlState_Identifiers[umlstate] = id

		// Initialisation of values
	}

	if len(umlscOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Umlsc instances pointers"
	}
	for idx, umlsc := range umlscOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Umlsc", idx, umlsc.Name)
		map_Umlsc_Identifiers[umlsc] = id

		// Initialisation of values
		for _, _umlstate := range umlsc.States {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "States")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_UmlState_Identifiers[_umlstate])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(verticeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Vertice instances pointers"
	}
	for idx, vertice := range verticeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Vertice", idx, vertice.Name)
		map_Vertice_Identifiers[vertice] = id

		// Initialisation of values
	}

	res = strings.ReplaceAll(res, "{{Identifiers}}", identifiersDecl)
	res = strings.ReplaceAll(res, "{{ValueInitializers}}", initializerStatements)
	res = strings.ReplaceAll(res, "{{PointersInitializers}}", pointersInitializesStatements)

	// Local time with timezone
	localTimestamp := stage.commitTimeStamp.Format("2006-01-02 15:04:05.000000 MST")

	// UTC time
	utcTimestamp := stage.commitTimeStamp.UTC().Format("2006-01-02 15:04:05.000000 UTC")
	res = strings.ReplaceAll(res, "{{LocalTimeStamp}}", localTimestamp)
	res = strings.ReplaceAll(res, "{{UTCTimeStamp}}", utcTimestamp)
	res = strings.ReplaceAll(res, "{{CommitId}}", fmt.Sprintf("%.10d", stage.commitId))

	if stage.MetaPackageImportAlias != "" {
		res = strings.ReplaceAll(res, "{{ImportPackageDeclaration}}",
			fmt.Sprintf("\n\t%s %s", stage.MetaPackageImportAlias, stage.MetaPackageImportPath))

		res = strings.ReplaceAll(res, "{{ImportPackageDummyDeclaration}}",
			fmt.Sprintf("\nvar _ %s.Stage",
				stage.MetaPackageImportAlias))

		var entries string

		// regenerate the map of doc link renaming
		// the key and value are set to the value because
		// if it has been renamed, this is the new value that matters
		valuesOrdered := make([]GONG__Identifier, 0)
		for _, value := range stage.Map_DocLink_Renaming {
			valuesOrdered = append(valuesOrdered, value)
		}
		sort.Slice(valuesOrdered[:], func(i, j int) bool {
			return valuesOrdered[i].Ident < valuesOrdered[j].Ident
		})
		for _, value := range valuesOrdered {

			// get the number of points in the value to find if it is a field
			// or a struct

			switch value.Type {
			case GONG__ENUM_CAST_INT:
				entries += fmt.Sprintf("\n\n\t\"%s\": %s(0),", value.Ident, value.Ident)
			case GONG__ENUM_CAST_STRING:
				entries += fmt.Sprintf("\n\n\t\"%s\": %s(\"\"),", value.Ident, value.Ident)
			case GONG__FIELD_VALUE:
				// substitute the second point with "{})."
				joker := "__substitute_for_first_point__"
				valueIdentifier := strings.Replace(value.Ident, ".", joker, 1)
				valueIdentifier = strings.Replace(valueIdentifier, ".", "{}).", 1)
				valueIdentifier = strings.Replace(valueIdentifier, joker, ".", 1)
				entries += fmt.Sprintf("\n\n\t\"%s\": (%s,", value.Ident, valueIdentifier)
			case GONG__IDENTIFIER_CONST:
				entries += fmt.Sprintf("\n\n\t\"%s\": %s,", value.Ident, value.Ident)
			case GONG__STRUCT_INSTANCE:
				entries += fmt.Sprintf("\n\n\t\"%s\": &(%s{}),", value.Ident, value.Ident)
			}
		}

		// res = strings.ReplaceAll(res, "{{EntriesDocLinkStringDocLinkIdentifier}}", entries)
	}

	if stage.generatesDiff {
		diff := computeDiff(stage.contentWhenParsed, res)
		os.WriteFile(fmt.Sprintf("%s-%.10d-%.10d.delta", name, stage.commitIdWhenParsed, stage.commitId), []byte(diff), os.FileMode(0666))
		diff = ComputeDiff(stage.contentWhenParsed, res)
		os.WriteFile(fmt.Sprintf("%s-%.10d-%.10d.diff", name, stage.commitIdWhenParsed, stage.commitId), []byte(diff), os.FileMode(0666))
	}
	stage.contentWhenParsed = res
	stage.commitIdWhenParsed = stage.commitId

	fmt.Fprintln(file, res)
}

// computeDiff calculates the git-style unified diff between two strings.
func computeDiff(a, b string) string {
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(a, b, false)
	return dmp.DiffToDelta(diffs)
}

// computePrettyDiff calculates the git-style unified diff between two strings.
func computePrettyDiff(a, b string) string {
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(a, b, false)
	return dmp.DiffPrettyHtml(diffs)
}

// applyDiff reconstructs the original string 'a' from the new string 'b' and the diff string 'c'.
func applyDiff(b, c string) (string, error) {
	dmp := diffmatchpatch.New()
	diffs, err := dmp.DiffFromDelta(b, c)
	if err != nil {
		return "", err
	}
	patches := dmp.PatchMake(b, diffs)
	// We are applying the patch in reverse to get from 'b' to 'a'.
	// The library's PatchApply function returns the new string and a slice of booleans indicating the success of each patch application.
	result, _ := dmp.PatchApply(patches, b)
	return result, nil
}

// unique identifier per struct
func generatesIdentifier(gongStructName string, idx int, instanceName string) (identifier string) {

	identifier = instanceName
	// Make a Regex to say we only want letters and numbers
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(instanceName, "_")

	identifier = fmt.Sprintf("__%s__%06d_%s", gongStructName, idx, processedString)

	return
}
