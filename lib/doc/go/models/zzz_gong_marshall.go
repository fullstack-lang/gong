// generated code - do not edit
package models

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
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
}
`

const IdentifiersDecls = `
	{{Identifier}} := (&models.{{GeneratedStructName}}{Name: ` + "`" + `{{GeneratedFieldNameValue}}` + "`" + `}).Stage(stage)`

// previous version does not hanldle embedded structs (https://github.com/golang/go/issues/9859)
// simpler version but the name of the instance cannot be human read before the fields initialization
const IdentifiersDeclsWithoutNameInit = `
	{{Identifier}} := (&models.{{GeneratedStructName}}{}).Stage(stage)`/* */

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
		log.Println(name + " is not a go filename")
	}

	log.Printf("Marshalling %s", name)

	res, err := stage.MarshallToString(modelsPackageName, packageName)
	if err != nil {
		log.Fatalln("Error marshalling to string:", err)
	}

	fmt.Fprintln(file, res)
}

// MarshallToString marshall the stage content into a string
func (stage *Stage) MarshallToString(modelsPackageName, packageName string) (res string, err error) {

	res = marshallRes
	res = strings.ReplaceAll(res, "{{PackageName}}", packageName)
	res = strings.ReplaceAll(res, "{{ModelsPackageName}}", modelsPackageName)

	// map of identifiers
	// var StageMapDstructIds map[*Dstruct]string
	identifiersDecl := ""
	initializerStatements := ""
	pointersInitializesStatements := ""

	decl := ""
	_ = decl
	setValueField := ""
	_ = setValueField

	// insertion initialization of objects to stage
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
	for _, attributeshape := range attributeshapeOrdered {
	
		identifiersDecl += attributeshape.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", attributeshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attributeshape.Name))
		initializerStatements += setValueField

		if str, ok := attributeshape.IdentifierMeta.(string); ok {
			setValueField = MetaFieldStructInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", attributeshape.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IdentifierMeta")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", str)
			initializerStatements += setValueField
		}

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", attributeshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FieldTypeAsString")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attributeshape.FieldTypeAsString))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", attributeshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Structname")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attributeshape.Structname))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", attributeshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Fieldtypename")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attributeshape.Fieldtypename))
		initializerStatements += setValueField

	}

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
	for _, classdiagram := range classdiagramOrdered {
	
		identifiersDecl += classdiagram.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", classdiagram.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(classdiagram.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", classdiagram.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Description")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(classdiagram.Description))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", classdiagram.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsIncludedInStaticWebSite")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", classdiagram.IsIncludedInStaticWebSite))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", classdiagram.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ShowNbInstances")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", classdiagram.ShowNbInstances))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", classdiagram.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ShowMultiplicity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", classdiagram.ShowMultiplicity))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", classdiagram.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ShowLinkNames")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", classdiagram.ShowLinkNames))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", classdiagram.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsInRenameMode")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", classdiagram.IsInRenameMode))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", classdiagram.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", classdiagram.IsExpanded))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", classdiagram.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "NodeGongStructsIsExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", classdiagram.NodeGongStructsIsExpanded))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", classdiagram.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "NodeGongStructNodeExpansion")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(classdiagram.NodeGongStructNodeExpansion))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", classdiagram.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "NodeGongEnumsIsExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", classdiagram.NodeGongEnumsIsExpanded))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", classdiagram.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "NodeGongEnumNodeExpansion")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(classdiagram.NodeGongEnumNodeExpansion))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", classdiagram.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "NodeGongNotesIsExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", classdiagram.NodeGongNotesIsExpanded))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", classdiagram.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "NodeGongNoteNodeExpansion")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(classdiagram.NodeGongNoteNodeExpansion))
		initializerStatements += setValueField

	}

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
	for _, diagrampackage := range diagrampackageOrdered {
	
		identifiersDecl += diagrampackage.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", diagrampackage.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagrampackage.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", diagrampackage.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Path")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagrampackage.Path))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", diagrampackage.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "GongModelPath")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagrampackage.GongModelPath))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", diagrampackage.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "AbsolutePathToDiagramPackage")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagrampackage.AbsolutePathToDiagramPackage))
		initializerStatements += setValueField

	}

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
	for _, gongenumshape := range gongenumshapeOrdered {
	
		identifiersDecl += gongenumshape.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongenumshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongenumshape.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongenumshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "X")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gongenumshape.X))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongenumshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Y")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gongenumshape.Y))
		initializerStatements += setValueField

		if str, ok := gongenumshape.IdentifierMeta.(string); ok {
			setValueField = MetaFieldStructInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongenumshape.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IdentifierMeta")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", str)
			initializerStatements += setValueField
		}

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongenumshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Width")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gongenumshape.Width))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongenumshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Height")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gongenumshape.Height))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongenumshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", gongenumshape.IsExpanded))
		initializerStatements += setValueField

	}

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
	for _, gongenumvalueshape := range gongenumvalueshapeOrdered {
	
		identifiersDecl += gongenumvalueshape.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongenumvalueshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongenumvalueshape.Name))
		initializerStatements += setValueField

		if str, ok := gongenumvalueshape.IdentifierMeta.(string); ok {
			setValueField = MetaFieldStructInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongenumvalueshape.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IdentifierMeta")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", str)
			initializerStatements += setValueField
		}

	}

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
	for _, gongnotelinkshape := range gongnotelinkshapeOrdered {
	
		identifiersDecl += gongnotelinkshape.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongnotelinkshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongnotelinkshape.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongnotelinkshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Identifier")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongnotelinkshape.Identifier))
		initializerStatements += setValueField

		if gongnotelinkshape.Type != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongnotelinkshape.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Type")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+gongnotelinkshape.Type.ToCodeString())
			initializerStatements += setValueField
		}

	}

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
	for _, gongnoteshape := range gongnoteshapeOrdered {
	
		identifiersDecl += gongnoteshape.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongnoteshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongnoteshape.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongnoteshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Identifier")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongnoteshape.Identifier))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongnoteshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Body")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongnoteshape.Body))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongnoteshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "BodyHTML")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongnoteshape.BodyHTML))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongnoteshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "X")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gongnoteshape.X))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongnoteshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Y")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gongnoteshape.Y))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongnoteshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Width")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gongnoteshape.Width))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongnoteshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Height")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gongnoteshape.Height))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongnoteshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Matched")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", gongnoteshape.Matched))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongnoteshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", gongnoteshape.IsExpanded))
		initializerStatements += setValueField

	}

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
	for _, gongstructshape := range gongstructshapeOrdered {
	
		identifiersDecl += gongstructshape.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongstructshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongstructshape.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongstructshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "X")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gongstructshape.X))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongstructshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Y")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gongstructshape.Y))
		initializerStatements += setValueField

		if str, ok := gongstructshape.IdentifierMeta.(string); ok {
			setValueField = MetaFieldStructInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongstructshape.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IdentifierMeta")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", str)
			initializerStatements += setValueField
		}

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongstructshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Width")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gongstructshape.Width))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongstructshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Height")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gongstructshape.Height))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongstructshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsSelected")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", gongstructshape.IsSelected))
		initializerStatements += setValueField

	}

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
	for _, linkshape := range linkshapeOrdered {
	
		identifiersDecl += linkshape.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", linkshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(linkshape.Name))
		initializerStatements += setValueField

		if str, ok := linkshape.IdentifierMeta.(string); ok {
			setValueField = MetaFieldStructInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", linkshape.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IdentifierMeta")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", str)
			initializerStatements += setValueField
		}

		if str, ok := linkshape.FieldTypeIdentifierMeta.(string); ok {
			setValueField = MetaFieldStructInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", linkshape.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FieldTypeIdentifierMeta")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", str)
			initializerStatements += setValueField
		}

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", linkshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FieldOffsetX")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", linkshape.FieldOffsetX))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", linkshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FieldOffsetY")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", linkshape.FieldOffsetY))
		initializerStatements += setValueField

		if linkshape.TargetMultiplicity != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", linkshape.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "TargetMultiplicity")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+linkshape.TargetMultiplicity.ToCodeString())
			initializerStatements += setValueField
		}

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", linkshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "TargetMultiplicityOffsetX")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", linkshape.TargetMultiplicityOffsetX))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", linkshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "TargetMultiplicityOffsetY")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", linkshape.TargetMultiplicityOffsetY))
		initializerStatements += setValueField

		if linkshape.SourceMultiplicity != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", linkshape.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "SourceMultiplicity")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+linkshape.SourceMultiplicity.ToCodeString())
			initializerStatements += setValueField
		}

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", linkshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "SourceMultiplicityOffsetX")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", linkshape.SourceMultiplicityOffsetX))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", linkshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "SourceMultiplicityOffsetY")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", linkshape.SourceMultiplicityOffsetY))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", linkshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "X")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", linkshape.X))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", linkshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Y")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", linkshape.Y))
		initializerStatements += setValueField

		if linkshape.StartOrientation != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", linkshape.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StartOrientation")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+linkshape.StartOrientation.ToCodeString())
			initializerStatements += setValueField
		}

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", linkshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StartRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", linkshape.StartRatio))
		initializerStatements += setValueField

		if linkshape.EndOrientation != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", linkshape.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EndOrientation")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+linkshape.EndOrientation.ToCodeString())
			initializerStatements += setValueField
		}

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", linkshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EndRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", linkshape.EndRatio))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", linkshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CornerOffsetRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", linkshape.CornerOffsetRatio))
		initializerStatements += setValueField

	}

	// insertion initialization of objects to stage
	if len(attributeshapeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of AttributeShape instances pointers"
	}
	for _, attributeshape := range attributeshapeOrdered {
		_ = attributeshape
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
	}

	if len(classdiagramOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Classdiagram instances pointers"
	}
	for _, classdiagram := range classdiagramOrdered {
		_ = classdiagram
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
		for _, _gongstructshape := range classdiagram.GongStructShapes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", classdiagram.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "GongStructShapes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _gongstructshape.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		for _, _gongenumshape := range classdiagram.GongEnumShapes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", classdiagram.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "GongEnumShapes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _gongenumshape.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		for _, _gongnoteshape := range classdiagram.GongNoteShapes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", classdiagram.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "GongNoteShapes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _gongnoteshape.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	}

	if len(diagrampackageOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of DiagramPackage instances pointers"
	}
	for _, diagrampackage := range diagrampackageOrdered {
		_ = diagrampackage
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
		for _, _classdiagram := range diagrampackage.Classdiagrams {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", diagrampackage.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Classdiagrams")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _classdiagram.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		if diagrampackage.SelectedClassdiagram != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", diagrampackage.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SelectedClassdiagram")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", diagrampackage.SelectedClassdiagram.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	}

	if len(gongenumshapeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of GongEnumShape instances pointers"
	}
	for _, gongenumshape := range gongenumshapeOrdered {
		_ = gongenumshape
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
		for _, _gongenumvalueshape := range gongenumshape.GongEnumValueShapes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", gongenumshape.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "GongEnumValueShapes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _gongenumvalueshape.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	}

	if len(gongenumvalueshapeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of GongEnumValueShape instances pointers"
	}
	for _, gongenumvalueshape := range gongenumvalueshapeOrdered {
		_ = gongenumvalueshape
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
	}

	if len(gongnotelinkshapeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of GongNoteLinkShape instances pointers"
	}
	for _, gongnotelinkshape := range gongnotelinkshapeOrdered {
		_ = gongnotelinkshape
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
	}

	if len(gongnoteshapeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of GongNoteShape instances pointers"
	}
	for _, gongnoteshape := range gongnoteshapeOrdered {
		_ = gongnoteshape
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
		for _, _gongnotelinkshape := range gongnoteshape.GongNoteLinkShapes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", gongnoteshape.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "GongNoteLinkShapes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _gongnotelinkshape.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	}

	if len(gongstructshapeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of GongStructShape instances pointers"
	}
	for _, gongstructshape := range gongstructshapeOrdered {
		_ = gongstructshape
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
		for _, _attributeshape := range gongstructshape.AttributeShapes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", gongstructshape.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "AttributeShapes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _attributeshape.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		for _, _linkshape := range gongstructshape.LinkShapes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", gongstructshape.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "LinkShapes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _linkshape.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	}

	if len(linkshapeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of LinkShape instances pointers"
	}
	for _, linkshape := range linkshapeOrdered {
		_ = linkshape
		var setPointerField string
		_ = setPointerField

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
	return
}

// insertion initialization of objects to stage
func (attributeshape *AttributeShape) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", attributeshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attributeshape.Name))
		initializerStatements += setValueField
	case "IdentifierMeta":
		if str, ok := attributeshape.IdentifierMeta.(string); ok {
			setValueField = MetaFieldStructInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", attributeshape.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IdentifierMeta")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", str)
			initializerStatements += setValueField
		}
	case "FieldTypeAsString":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", attributeshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FieldTypeAsString")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attributeshape.FieldTypeAsString))
		initializerStatements += setValueField
	case "Structname":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", attributeshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Structname")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attributeshape.Structname))
		initializerStatements += setValueField
	case "Fieldtypename":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", attributeshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Fieldtypename")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attributeshape.Fieldtypename))
		initializerStatements += setValueField


	default:
		log.Panicf("Unknown field %s for Gongstruct AttributeShape", fieldName)
	}
	return
}

func (classdiagram *Classdiagram) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", classdiagram.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(classdiagram.Name))
		initializerStatements += setValueField
	case "Description":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", classdiagram.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Description")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(classdiagram.Description))
		initializerStatements += setValueField
	case "IsIncludedInStaticWebSite":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", classdiagram.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsIncludedInStaticWebSite")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", classdiagram.IsIncludedInStaticWebSite))
		initializerStatements += setValueField
	case "ShowNbInstances":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", classdiagram.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ShowNbInstances")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", classdiagram.ShowNbInstances))
		initializerStatements += setValueField
	case "ShowMultiplicity":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", classdiagram.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ShowMultiplicity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", classdiagram.ShowMultiplicity))
		initializerStatements += setValueField
	case "ShowLinkNames":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", classdiagram.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ShowLinkNames")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", classdiagram.ShowLinkNames))
		initializerStatements += setValueField
	case "IsInRenameMode":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", classdiagram.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsInRenameMode")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", classdiagram.IsInRenameMode))
		initializerStatements += setValueField
	case "IsExpanded":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", classdiagram.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", classdiagram.IsExpanded))
		initializerStatements += setValueField
	case "NodeGongStructsIsExpanded":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", classdiagram.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "NodeGongStructsIsExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", classdiagram.NodeGongStructsIsExpanded))
		initializerStatements += setValueField
	case "NodeGongStructNodeExpansion":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", classdiagram.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "NodeGongStructNodeExpansion")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(classdiagram.NodeGongStructNodeExpansion))
		initializerStatements += setValueField
	case "NodeGongEnumsIsExpanded":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", classdiagram.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "NodeGongEnumsIsExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", classdiagram.NodeGongEnumsIsExpanded))
		initializerStatements += setValueField
	case "NodeGongEnumNodeExpansion":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", classdiagram.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "NodeGongEnumNodeExpansion")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(classdiagram.NodeGongEnumNodeExpansion))
		initializerStatements += setValueField
	case "NodeGongNotesIsExpanded":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", classdiagram.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "NodeGongNotesIsExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", classdiagram.NodeGongNotesIsExpanded))
		initializerStatements += setValueField
	case "NodeGongNoteNodeExpansion":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", classdiagram.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "NodeGongNoteNodeExpansion")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(classdiagram.NodeGongNoteNodeExpansion))
		initializerStatements += setValueField

	case "GongStructShapes":
		for _, _gongstructshape := range classdiagram.GongStructShapes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", classdiagram.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "GongStructShapes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _gongstructshape.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "GongEnumShapes":
		for _, _gongenumshape := range classdiagram.GongEnumShapes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", classdiagram.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "GongEnumShapes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _gongenumshape.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "GongNoteShapes":
		for _, _gongnoteshape := range classdiagram.GongNoteShapes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", classdiagram.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "GongNoteShapes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _gongnoteshape.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct Classdiagram", fieldName)
	}
	return
}

func (diagrampackage *DiagramPackage) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", diagrampackage.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagrampackage.Name))
		initializerStatements += setValueField
	case "Path":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", diagrampackage.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Path")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagrampackage.Path))
		initializerStatements += setValueField
	case "GongModelPath":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", diagrampackage.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "GongModelPath")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagrampackage.GongModelPath))
		initializerStatements += setValueField
	case "AbsolutePathToDiagramPackage":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", diagrampackage.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "AbsolutePathToDiagramPackage")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagrampackage.AbsolutePathToDiagramPackage))
		initializerStatements += setValueField

	case "Classdiagrams":
		for _, _classdiagram := range diagrampackage.Classdiagrams {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", diagrampackage.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Classdiagrams")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _classdiagram.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "SelectedClassdiagram":
		if diagrampackage.SelectedClassdiagram != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", diagrampackage.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SelectedClassdiagram")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", diagrampackage.SelectedClassdiagram.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct DiagramPackage", fieldName)
	}
	return
}

func (gongenumshape *GongEnumShape) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongenumshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongenumshape.Name))
		initializerStatements += setValueField
	case "X":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongenumshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "X")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gongenumshape.X))
		initializerStatements += setValueField
	case "Y":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongenumshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Y")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gongenumshape.Y))
		initializerStatements += setValueField
	case "IdentifierMeta":
		if str, ok := gongenumshape.IdentifierMeta.(string); ok {
			setValueField = MetaFieldStructInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongenumshape.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IdentifierMeta")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", str)
			initializerStatements += setValueField
		}
	case "Width":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongenumshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Width")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gongenumshape.Width))
		initializerStatements += setValueField
	case "Height":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongenumshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Height")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gongenumshape.Height))
		initializerStatements += setValueField
	case "IsExpanded":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongenumshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", gongenumshape.IsExpanded))
		initializerStatements += setValueField

	case "GongEnumValueShapes":
		for _, _gongenumvalueshape := range gongenumshape.GongEnumValueShapes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", gongenumshape.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "GongEnumValueShapes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _gongenumvalueshape.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct GongEnumShape", fieldName)
	}
	return
}

func (gongenumvalueshape *GongEnumValueShape) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongenumvalueshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongenumvalueshape.Name))
		initializerStatements += setValueField
	case "IdentifierMeta":
		if str, ok := gongenumvalueshape.IdentifierMeta.(string); ok {
			setValueField = MetaFieldStructInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongenumvalueshape.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IdentifierMeta")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", str)
			initializerStatements += setValueField
		}


	default:
		log.Panicf("Unknown field %s for Gongstruct GongEnumValueShape", fieldName)
	}
	return
}

func (gongnotelinkshape *GongNoteLinkShape) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongnotelinkshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongnotelinkshape.Name))
		initializerStatements += setValueField
	case "Identifier":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongnotelinkshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Identifier")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongnotelinkshape.Identifier))
		initializerStatements += setValueField
	case "Type":
		if gongnotelinkshape.Type != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongnotelinkshape.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Type")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+gongnotelinkshape.Type.ToCodeString())
			initializerStatements += setValueField
		}


	default:
		log.Panicf("Unknown field %s for Gongstruct GongNoteLinkShape", fieldName)
	}
	return
}

func (gongnoteshape *GongNoteShape) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongnoteshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongnoteshape.Name))
		initializerStatements += setValueField
	case "Identifier":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongnoteshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Identifier")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongnoteshape.Identifier))
		initializerStatements += setValueField
	case "Body":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongnoteshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Body")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongnoteshape.Body))
		initializerStatements += setValueField
	case "BodyHTML":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongnoteshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "BodyHTML")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongnoteshape.BodyHTML))
		initializerStatements += setValueField
	case "X":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongnoteshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "X")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gongnoteshape.X))
		initializerStatements += setValueField
	case "Y":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongnoteshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Y")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gongnoteshape.Y))
		initializerStatements += setValueField
	case "Width":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongnoteshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Width")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gongnoteshape.Width))
		initializerStatements += setValueField
	case "Height":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongnoteshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Height")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gongnoteshape.Height))
		initializerStatements += setValueField
	case "Matched":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongnoteshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Matched")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", gongnoteshape.Matched))
		initializerStatements += setValueField
	case "IsExpanded":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongnoteshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", gongnoteshape.IsExpanded))
		initializerStatements += setValueField

	case "GongNoteLinkShapes":
		for _, _gongnotelinkshape := range gongnoteshape.GongNoteLinkShapes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", gongnoteshape.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "GongNoteLinkShapes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _gongnotelinkshape.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct GongNoteShape", fieldName)
	}
	return
}

func (gongstructshape *GongStructShape) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongstructshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongstructshape.Name))
		initializerStatements += setValueField
	case "X":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongstructshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "X")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gongstructshape.X))
		initializerStatements += setValueField
	case "Y":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongstructshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Y")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gongstructshape.Y))
		initializerStatements += setValueField
	case "IdentifierMeta":
		if str, ok := gongstructshape.IdentifierMeta.(string); ok {
			setValueField = MetaFieldStructInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongstructshape.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IdentifierMeta")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", str)
			initializerStatements += setValueField
		}
	case "Width":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongstructshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Width")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gongstructshape.Width))
		initializerStatements += setValueField
	case "Height":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongstructshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Height")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gongstructshape.Height))
		initializerStatements += setValueField
	case "IsSelected":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongstructshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsSelected")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", gongstructshape.IsSelected))
		initializerStatements += setValueField

	case "AttributeShapes":
		for _, _attributeshape := range gongstructshape.AttributeShapes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", gongstructshape.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "AttributeShapes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _attributeshape.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "LinkShapes":
		for _, _linkshape := range gongstructshape.LinkShapes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", gongstructshape.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "LinkShapes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _linkshape.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct GongStructShape", fieldName)
	}
	return
}

func (linkshape *LinkShape) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", linkshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(linkshape.Name))
		initializerStatements += setValueField
	case "IdentifierMeta":
		if str, ok := linkshape.IdentifierMeta.(string); ok {
			setValueField = MetaFieldStructInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", linkshape.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IdentifierMeta")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", str)
			initializerStatements += setValueField
		}
	case "FieldTypeIdentifierMeta":
		if str, ok := linkshape.FieldTypeIdentifierMeta.(string); ok {
			setValueField = MetaFieldStructInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", linkshape.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FieldTypeIdentifierMeta")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", str)
			initializerStatements += setValueField
		}
	case "FieldOffsetX":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", linkshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FieldOffsetX")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", linkshape.FieldOffsetX))
		initializerStatements += setValueField
	case "FieldOffsetY":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", linkshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FieldOffsetY")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", linkshape.FieldOffsetY))
		initializerStatements += setValueField
	case "TargetMultiplicity":
		if linkshape.TargetMultiplicity != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", linkshape.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "TargetMultiplicity")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+linkshape.TargetMultiplicity.ToCodeString())
			initializerStatements += setValueField
		}
	case "TargetMultiplicityOffsetX":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", linkshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "TargetMultiplicityOffsetX")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", linkshape.TargetMultiplicityOffsetX))
		initializerStatements += setValueField
	case "TargetMultiplicityOffsetY":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", linkshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "TargetMultiplicityOffsetY")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", linkshape.TargetMultiplicityOffsetY))
		initializerStatements += setValueField
	case "SourceMultiplicity":
		if linkshape.SourceMultiplicity != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", linkshape.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "SourceMultiplicity")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+linkshape.SourceMultiplicity.ToCodeString())
			initializerStatements += setValueField
		}
	case "SourceMultiplicityOffsetX":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", linkshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "SourceMultiplicityOffsetX")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", linkshape.SourceMultiplicityOffsetX))
		initializerStatements += setValueField
	case "SourceMultiplicityOffsetY":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", linkshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "SourceMultiplicityOffsetY")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", linkshape.SourceMultiplicityOffsetY))
		initializerStatements += setValueField
	case "X":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", linkshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "X")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", linkshape.X))
		initializerStatements += setValueField
	case "Y":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", linkshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Y")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", linkshape.Y))
		initializerStatements += setValueField
	case "StartOrientation":
		if linkshape.StartOrientation != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", linkshape.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StartOrientation")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+linkshape.StartOrientation.ToCodeString())
			initializerStatements += setValueField
		}
	case "StartRatio":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", linkshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StartRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", linkshape.StartRatio))
		initializerStatements += setValueField
	case "EndOrientation":
		if linkshape.EndOrientation != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", linkshape.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EndOrientation")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+linkshape.EndOrientation.ToCodeString())
			initializerStatements += setValueField
		}
	case "EndRatio":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", linkshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EndRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", linkshape.EndRatio))
		initializerStatements += setValueField
	case "CornerOffsetRatio":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", linkshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CornerOffsetRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", linkshape.CornerOffsetRatio))
		initializerStatements += setValueField


	default:
		log.Panicf("Unknown field %s for Gongstruct LinkShape", fieldName)
	}
	return
}
