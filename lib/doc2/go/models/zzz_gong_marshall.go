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

	// Declaration of instances to stage{{Identifiers}}

	// Setup of values{{ValueInitializers}}

	// Setup of pointers{{PointersInitializers}}
}`

const IdentifiersDecls = `
	{{Identifier}} := (&models.{{GeneratedStructName}}{}).Stage(stage)`

const StringInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = ` + "`" + `{{GeneratedFieldNameValue}}` + "`"

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
	map_AttributeShape_Identifiers := make(map[*AttributeShape]string)
	_ = map_AttributeShape_Identifiers

	attributeshapeOrdered := []*AttributeShape{}
	for attributeshape := range stage.AttributeShapes {
		attributeshapeOrdered = append(attributeshapeOrdered, attributeshape)
	}
	sort.Slice(attributeshapeOrdered[:], func(i, j int) bool {
		attributeshapei := attributeshapeOrdered[i]
		attributeshapej := attributeshapeOrdered[j]
		attributeshapei_order, oki := stage.AttributeShapeMap_Staged_Order[attributeshapei]
		attributeshapej_order, okj := stage.AttributeShapeMap_Staged_Order[attributeshapej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return attributeshapei_order < attributeshapej_order
	})
	if len(attributeshapeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, attributeshape := range attributeshapeOrdered {

		id = generatesIdentifier("AttributeShape", idx, attributeshape.Name)
		map_AttributeShape_Identifiers[attributeshape] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "AttributeShape")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attributeshape.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attributeshape.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "\n\t{{Identifier}}",
			fmt.Sprintf("\n\n\t//gong:ident [%s] comment added to overcome the problem with the comment map association\n\t{{Identifier}}",
				string(attributeshape.Identifier)))
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Identifier")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attributeshape.Identifier))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FieldTypeAsString")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attributeshape.FieldTypeAsString))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Structname")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attributeshape.Structname))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Fieldtypename")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attributeshape.Fieldtypename))
		initializerStatements += setValueField

	}

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
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsInRenameMode")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", classdiagram.IsInRenameMode))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", classdiagram.IsExpanded))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "NodeGongStructsIsExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", classdiagram.NodeGongStructsIsExpanded))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "NodeGongStructNodeExpansionBinaryEncoding")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", classdiagram.NodeGongStructNodeExpansionBinaryEncoding))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "NodeGongEnumsIsExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", classdiagram.NodeGongEnumsIsExpanded))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "NodeGongEnumNodeExpansionBinaryEncoding")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", classdiagram.NodeGongEnumNodeExpansionBinaryEncoding))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "NodeGongNotesIsExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", classdiagram.NodeGongNotesIsExpanded))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "NodeGongNoteNodeExpansionBinaryEncoding")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", classdiagram.NodeGongNoteNodeExpansionBinaryEncoding))
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

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "AbsolutePathToDiagramPackage")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagrampackage.AbsolutePathToDiagramPackage))
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

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "X")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gongenumshape.X))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Y")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gongenumshape.Y))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "\n\t{{Identifier}}",
			fmt.Sprintf("\n\n\t//gong:ident [%s] comment added to overcome the problem with the comment map association\n\t{{Identifier}}",
				string(gongenumshape.Identifier)))
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

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", gongenumshape.IsExpanded))
		initializerStatements += setValueField

	}

	map_GongEnumValueShape_Identifiers := make(map[*GongEnumValueShape]string)
	_ = map_GongEnumValueShape_Identifiers

	gongenumvalueshapeOrdered := []*GongEnumValueShape{}
	for gongenumvalueshape := range stage.GongEnumValueShapes {
		gongenumvalueshapeOrdered = append(gongenumvalueshapeOrdered, gongenumvalueshape)
	}
	sort.Slice(gongenumvalueshapeOrdered[:], func(i, j int) bool {
		gongenumvalueshapei := gongenumvalueshapeOrdered[i]
		gongenumvalueshapej := gongenumvalueshapeOrdered[j]
		gongenumvalueshapei_order, oki := stage.GongEnumValueShapeMap_Staged_Order[gongenumvalueshapei]
		gongenumvalueshapej_order, okj := stage.GongEnumValueShapeMap_Staged_Order[gongenumvalueshapej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return gongenumvalueshapei_order < gongenumvalueshapej_order
	})
	if len(gongenumvalueshapeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, gongenumvalueshape := range gongenumvalueshapeOrdered {

		id = generatesIdentifier("GongEnumValueShape", idx, gongenumvalueshape.Name)
		map_GongEnumValueShape_Identifiers[gongenumvalueshape] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "GongEnumValueShape")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", gongenumvalueshape.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongenumvalueshape.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "\n\t{{Identifier}}",
			fmt.Sprintf("\n\n\t//gong:ident [%s] comment added to overcome the problem with the comment map association\n\t{{Identifier}}",
				string(gongenumvalueshape.Identifier)))
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Identifier")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongenumvalueshape.Identifier))
		initializerStatements += setValueField

	}

	map_GongNoteLinkShape_Identifiers := make(map[*GongNoteLinkShape]string)
	_ = map_GongNoteLinkShape_Identifiers

	gongnotelinkshapeOrdered := []*GongNoteLinkShape{}
	for gongnotelinkshape := range stage.GongNoteLinkShapes {
		gongnotelinkshapeOrdered = append(gongnotelinkshapeOrdered, gongnotelinkshape)
	}
	sort.Slice(gongnotelinkshapeOrdered[:], func(i, j int) bool {
		gongnotelinkshapei := gongnotelinkshapeOrdered[i]
		gongnotelinkshapej := gongnotelinkshapeOrdered[j]
		gongnotelinkshapei_order, oki := stage.GongNoteLinkShapeMap_Staged_Order[gongnotelinkshapei]
		gongnotelinkshapej_order, okj := stage.GongNoteLinkShapeMap_Staged_Order[gongnotelinkshapej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return gongnotelinkshapei_order < gongnotelinkshapej_order
	})
	if len(gongnotelinkshapeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, gongnotelinkshape := range gongnotelinkshapeOrdered {

		id = generatesIdentifier("GongNoteLinkShape", idx, gongnotelinkshape.Name)
		map_GongNoteLinkShape_Identifiers[gongnotelinkshape] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "GongNoteLinkShape")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", gongnotelinkshape.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongnotelinkshape.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "\n\t{{Identifier}}",
			fmt.Sprintf("\n\n\t//gong:ident [%s] comment added to overcome the problem with the comment map association\n\t{{Identifier}}",
				string(gongnotelinkshape.Identifier)))
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Identifier")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongnotelinkshape.Identifier))
		initializerStatements += setValueField

		if gongnotelinkshape.Type != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Type")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+gongnotelinkshape.Type.ToCodeString())
			initializerStatements += setValueField
		}

	}

	map_GongNoteShape_Identifiers := make(map[*GongNoteShape]string)
	_ = map_GongNoteShape_Identifiers

	gongnoteshapeOrdered := []*GongNoteShape{}
	for gongnoteshape := range stage.GongNoteShapes {
		gongnoteshapeOrdered = append(gongnoteshapeOrdered, gongnoteshape)
	}
	sort.Slice(gongnoteshapeOrdered[:], func(i, j int) bool {
		gongnoteshapei := gongnoteshapeOrdered[i]
		gongnoteshapej := gongnoteshapeOrdered[j]
		gongnoteshapei_order, oki := stage.GongNoteShapeMap_Staged_Order[gongnoteshapei]
		gongnoteshapej_order, okj := stage.GongNoteShapeMap_Staged_Order[gongnoteshapej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return gongnoteshapei_order < gongnoteshapej_order
	})
	if len(gongnoteshapeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, gongnoteshape := range gongnoteshapeOrdered {

		id = generatesIdentifier("GongNoteShape", idx, gongnoteshape.Name)
		map_GongNoteShape_Identifiers[gongnoteshape] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "GongNoteShape")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", gongnoteshape.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongnoteshape.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "\n\t{{Identifier}}",
			fmt.Sprintf("\n\n\t//gong:ident [%s] comment added to overcome the problem with the comment map association\n\t{{Identifier}}",
				string(gongnoteshape.Identifier)))
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Identifier")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongnoteshape.Identifier))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Body")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongnoteshape.Body))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "BodyHTML")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongnoteshape.BodyHTML))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "X")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gongnoteshape.X))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Y")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gongnoteshape.Y))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Width")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gongnoteshape.Width))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Height")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gongnoteshape.Height))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Matched")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", gongnoteshape.Matched))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", gongnoteshape.IsExpanded))
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

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "X")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gongstructshape.X))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Y")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gongstructshape.Y))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "\n\t{{Identifier}}",
			fmt.Sprintf("\n\n\t//gong:ident [%s] comment added to overcome the problem with the comment map association\n\t{{Identifier}}",
				string(gongstructshape.Identifier)))
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

	map_LinkShape_Identifiers := make(map[*LinkShape]string)
	_ = map_LinkShape_Identifiers

	linkshapeOrdered := []*LinkShape{}
	for linkshape := range stage.LinkShapes {
		linkshapeOrdered = append(linkshapeOrdered, linkshape)
	}
	sort.Slice(linkshapeOrdered[:], func(i, j int) bool {
		linkshapei := linkshapeOrdered[i]
		linkshapej := linkshapeOrdered[j]
		linkshapei_order, oki := stage.LinkShapeMap_Staged_Order[linkshapei]
		linkshapej_order, okj := stage.LinkShapeMap_Staged_Order[linkshapej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return linkshapei_order < linkshapej_order
	})
	if len(linkshapeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, linkshape := range linkshapeOrdered {

		id = generatesIdentifier("LinkShape", idx, linkshape.Name)
		map_LinkShape_Identifiers[linkshape] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "LinkShape")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", linkshape.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(linkshape.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "\n\t{{Identifier}}",
			fmt.Sprintf("\n\n\t//gong:ident [%s] comment added to overcome the problem with the comment map association\n\t{{Identifier}}",
				string(linkshape.Identifier)))
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Identifier")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(linkshape.Identifier))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "\n\t{{Identifier}}",
			fmt.Sprintf("\n\n\t//gong:ident [%s] comment added to overcome the problem with the comment map association\n\t{{Identifier}}",
				string(linkshape.Fieldtypename)))
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Fieldtypename")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(linkshape.Fieldtypename))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FieldOffsetX")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", linkshape.FieldOffsetX))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FieldOffsetY")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", linkshape.FieldOffsetY))
		initializerStatements += setValueField

		if linkshape.TargetMultiplicity != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "TargetMultiplicity")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+linkshape.TargetMultiplicity.ToCodeString())
			initializerStatements += setValueField
		}

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "TargetMultiplicityOffsetX")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", linkshape.TargetMultiplicityOffsetX))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "TargetMultiplicityOffsetY")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", linkshape.TargetMultiplicityOffsetY))
		initializerStatements += setValueField

		if linkshape.SourceMultiplicity != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "SourceMultiplicity")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+linkshape.SourceMultiplicity.ToCodeString())
			initializerStatements += setValueField
		}

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "SourceMultiplicityOffsetX")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", linkshape.SourceMultiplicityOffsetX))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "SourceMultiplicityOffsetY")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", linkshape.SourceMultiplicityOffsetY))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "X")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", linkshape.X))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Y")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", linkshape.Y))
		initializerStatements += setValueField

		if linkshape.StartOrientation != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StartOrientation")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+linkshape.StartOrientation.ToCodeString())
			initializerStatements += setValueField
		}

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StartRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", linkshape.StartRatio))
		initializerStatements += setValueField

		if linkshape.EndOrientation != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EndOrientation")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+linkshape.EndOrientation.ToCodeString())
			initializerStatements += setValueField
		}

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EndRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", linkshape.EndRatio))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CornerOffsetRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", linkshape.CornerOffsetRatio))
		initializerStatements += setValueField

	}

	// insertion initialization of objects to stage
	if len(attributeshapeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of AttributeShape instances pointers"
	}
	for idx, attributeshape := range attributeshapeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("AttributeShape", idx, attributeshape.Name)
		map_AttributeShape_Identifiers[attributeshape] = id

		// Initialisation of values
	}

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

		for _, _gongnoteshape := range classdiagram.GongNoteShapes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "GongNoteShapes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_GongNoteShape_Identifiers[_gongnoteshape])
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
		for _, _gongenumvalueshape := range gongenumshape.GongEnumValueShapes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "GongEnumValueShapes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_GongEnumValueShape_Identifiers[_gongenumvalueshape])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(gongenumvalueshapeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of GongEnumValueShape instances pointers"
	}
	for idx, gongenumvalueshape := range gongenumvalueshapeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("GongEnumValueShape", idx, gongenumvalueshape.Name)
		map_GongEnumValueShape_Identifiers[gongenumvalueshape] = id

		// Initialisation of values
	}

	if len(gongnotelinkshapeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of GongNoteLinkShape instances pointers"
	}
	for idx, gongnotelinkshape := range gongnotelinkshapeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("GongNoteLinkShape", idx, gongnotelinkshape.Name)
		map_GongNoteLinkShape_Identifiers[gongnotelinkshape] = id

		// Initialisation of values
	}

	if len(gongnoteshapeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of GongNoteShape instances pointers"
	}
	for idx, gongnoteshape := range gongnoteshapeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("GongNoteShape", idx, gongnoteshape.Name)
		map_GongNoteShape_Identifiers[gongnoteshape] = id

		// Initialisation of values
		for _, _gongnotelinkshape := range gongnoteshape.GongNoteLinkShapes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "GongNoteLinkShapes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_GongNoteLinkShape_Identifiers[_gongnotelinkshape])
			pointersInitializesStatements += setPointerField
		}

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
		for _, _attributeshape := range gongstructshape.AttributeShapes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "AttributeShapes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_AttributeShape_Identifiers[_attributeshape])
			pointersInitializesStatements += setPointerField
		}

		for _, _linkshape := range gongstructshape.LinkShapes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "LinkShapes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_LinkShape_Identifiers[_linkshape])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(linkshapeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of LinkShape instances pointers"
	}
	for idx, linkshape := range linkshapeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("LinkShape", idx, linkshape.Name)
		map_LinkShape_Identifiers[linkshape] = id

		// Initialisation of values
	}

	res = strings.ReplaceAll(res, "{{Identifiers}}", identifiersDecl)
	res = strings.ReplaceAll(res, "{{ValueInitializers}}", initializerStatements)
	res = strings.ReplaceAll(res, "{{PointersInitializers}}", pointersInitializesStatements)

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

	fmt.Fprintln(file, res)
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
