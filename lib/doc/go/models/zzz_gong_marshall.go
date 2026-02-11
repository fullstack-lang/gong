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

	// insertion point for declaration of instances to stage{{Identifiers}}

	// insertion point for initialization of values{{ValueInitializers}}

	// insertion point for setup of pointers{{PointersInitializers}}
}`

const GongIdentifiersDecls = `
	{{Identifier}} := (&models.{{GeneratedStructName}}{Name: ` + "`" + `{{GeneratedFieldNameValue}}` + "`" + `}).Stage(stage)`

const GongUnstageStmt = `
	{{Identifier}}.Unstage(stage)`

// previous version does not hanldle embedded structs (https://github.com/golang/go/issues/9859)
// simpler version but the name of the instance cannot be human read before the fields initialization
const IdentifiersDeclsWithoutNameInit = `
	{{Identifier}} := (&models.{{GeneratedStructName}}{}).Stage(stage)` /* */

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
	var identifiersDecl strings.Builder
	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder

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
		identifiersDecl.WriteString("\n")
	}
	for _, attributeshape := range attributeshapeOrdered {

		identifiersDecl.WriteString(attributeshape.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(attributeshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(attributeshape.GongMarshallField(stage, "IdentifierMeta"))
		initializerStatements.WriteString(attributeshape.GongMarshallField(stage, "FieldTypeAsString"))
		initializerStatements.WriteString(attributeshape.GongMarshallField(stage, "Structname"))
		initializerStatements.WriteString(attributeshape.GongMarshallField(stage, "Fieldtypename"))
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
		identifiersDecl.WriteString("\n")
	}
	for _, classdiagram := range classdiagramOrdered {

		identifiersDecl.WriteString(classdiagram.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(classdiagram.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(classdiagram.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(classdiagram.GongMarshallField(stage, "IsIncludedInStaticWebSite"))
		pointersInitializesStatements.WriteString(classdiagram.GongMarshallField(stage, "GongStructShapes"))
		pointersInitializesStatements.WriteString(classdiagram.GongMarshallField(stage, "GongEnumShapes"))
		pointersInitializesStatements.WriteString(classdiagram.GongMarshallField(stage, "GongNoteShapes"))
		initializerStatements.WriteString(classdiagram.GongMarshallField(stage, "ShowNbInstances"))
		initializerStatements.WriteString(classdiagram.GongMarshallField(stage, "ShowMultiplicity"))
		initializerStatements.WriteString(classdiagram.GongMarshallField(stage, "ShowLinkNames"))
		initializerStatements.WriteString(classdiagram.GongMarshallField(stage, "IsInRenameMode"))
		initializerStatements.WriteString(classdiagram.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(classdiagram.GongMarshallField(stage, "NodeGongStructsIsExpanded"))
		initializerStatements.WriteString(classdiagram.GongMarshallField(stage, "NodeGongStructNodeExpansion"))
		initializerStatements.WriteString(classdiagram.GongMarshallField(stage, "NodeGongEnumsIsExpanded"))
		initializerStatements.WriteString(classdiagram.GongMarshallField(stage, "NodeGongEnumNodeExpansion"))
		initializerStatements.WriteString(classdiagram.GongMarshallField(stage, "NodeGongNotesIsExpanded"))
		initializerStatements.WriteString(classdiagram.GongMarshallField(stage, "NodeGongNoteNodeExpansion"))
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
		identifiersDecl.WriteString("\n")
	}
	for _, diagrampackage := range diagrampackageOrdered {

		identifiersDecl.WriteString(diagrampackage.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(diagrampackage.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(diagrampackage.GongMarshallField(stage, "Path"))
		initializerStatements.WriteString(diagrampackage.GongMarshallField(stage, "GongModelPath"))
		pointersInitializesStatements.WriteString(diagrampackage.GongMarshallField(stage, "Classdiagrams"))
		pointersInitializesStatements.WriteString(diagrampackage.GongMarshallField(stage, "SelectedClassdiagram"))
		initializerStatements.WriteString(diagrampackage.GongMarshallField(stage, "AbsolutePathToDiagramPackage"))
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
		identifiersDecl.WriteString("\n")
	}
	for _, gongenumshape := range gongenumshapeOrdered {

		identifiersDecl.WriteString(gongenumshape.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(gongenumshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(gongenumshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(gongenumshape.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(gongenumshape.GongMarshallField(stage, "IdentifierMeta"))
		pointersInitializesStatements.WriteString(gongenumshape.GongMarshallField(stage, "GongEnumValueShapes"))
		initializerStatements.WriteString(gongenumshape.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(gongenumshape.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(gongenumshape.GongMarshallField(stage, "IsExpanded"))
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
		identifiersDecl.WriteString("\n")
	}
	for _, gongenumvalueshape := range gongenumvalueshapeOrdered {

		identifiersDecl.WriteString(gongenumvalueshape.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(gongenumvalueshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(gongenumvalueshape.GongMarshallField(stage, "IdentifierMeta"))
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
		identifiersDecl.WriteString("\n")
	}
	for _, gongnotelinkshape := range gongnotelinkshapeOrdered {

		identifiersDecl.WriteString(gongnotelinkshape.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(gongnotelinkshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(gongnotelinkshape.GongMarshallField(stage, "Identifier"))
		initializerStatements.WriteString(gongnotelinkshape.GongMarshallField(stage, "Type"))
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
		identifiersDecl.WriteString("\n")
	}
	for _, gongnoteshape := range gongnoteshapeOrdered {

		identifiersDecl.WriteString(gongnoteshape.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(gongnoteshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(gongnoteshape.GongMarshallField(stage, "Identifier"))
		initializerStatements.WriteString(gongnoteshape.GongMarshallField(stage, "Body"))
		initializerStatements.WriteString(gongnoteshape.GongMarshallField(stage, "BodyHTML"))
		initializerStatements.WriteString(gongnoteshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(gongnoteshape.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(gongnoteshape.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(gongnoteshape.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(gongnoteshape.GongMarshallField(stage, "Matched"))
		pointersInitializesStatements.WriteString(gongnoteshape.GongMarshallField(stage, "GongNoteLinkShapes"))
		initializerStatements.WriteString(gongnoteshape.GongMarshallField(stage, "IsExpanded"))
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
		identifiersDecl.WriteString("\n")
	}
	for _, gongstructshape := range gongstructshapeOrdered {

		identifiersDecl.WriteString(gongstructshape.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(gongstructshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(gongstructshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(gongstructshape.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(gongstructshape.GongMarshallField(stage, "IdentifierMeta"))
		pointersInitializesStatements.WriteString(gongstructshape.GongMarshallField(stage, "AttributeShapes"))
		pointersInitializesStatements.WriteString(gongstructshape.GongMarshallField(stage, "LinkShapes"))
		initializerStatements.WriteString(gongstructshape.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(gongstructshape.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(gongstructshape.GongMarshallField(stage, "IsSelected"))
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
		identifiersDecl.WriteString("\n")
	}
	for _, linkshape := range linkshapeOrdered {

		identifiersDecl.WriteString(linkshape.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(linkshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(linkshape.GongMarshallField(stage, "IdentifierMeta"))
		initializerStatements.WriteString(linkshape.GongMarshallField(stage, "FieldTypeIdentifierMeta"))
		initializerStatements.WriteString(linkshape.GongMarshallField(stage, "FieldOffsetX"))
		initializerStatements.WriteString(linkshape.GongMarshallField(stage, "FieldOffsetY"))
		initializerStatements.WriteString(linkshape.GongMarshallField(stage, "TargetMultiplicity"))
		initializerStatements.WriteString(linkshape.GongMarshallField(stage, "TargetMultiplicityOffsetX"))
		initializerStatements.WriteString(linkshape.GongMarshallField(stage, "TargetMultiplicityOffsetY"))
		initializerStatements.WriteString(linkshape.GongMarshallField(stage, "SourceMultiplicity"))
		initializerStatements.WriteString(linkshape.GongMarshallField(stage, "SourceMultiplicityOffsetX"))
		initializerStatements.WriteString(linkshape.GongMarshallField(stage, "SourceMultiplicityOffsetY"))
		initializerStatements.WriteString(linkshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(linkshape.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(linkshape.GongMarshallField(stage, "StartOrientation"))
		initializerStatements.WriteString(linkshape.GongMarshallField(stage, "StartRatio"))
		initializerStatements.WriteString(linkshape.GongMarshallField(stage, "EndOrientation"))
		initializerStatements.WriteString(linkshape.GongMarshallField(stage, "EndRatio"))
		initializerStatements.WriteString(linkshape.GongMarshallField(stage, "CornerOffsetRatio"))
	}

	// insertion initialization of objects to stage
	for _, attributeshape := range attributeshapeOrdered {
		_ = attributeshape
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, classdiagram := range classdiagramOrdered {
		_ = classdiagram
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, diagrampackage := range diagrampackageOrdered {
		_ = diagrampackage
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, gongenumshape := range gongenumshapeOrdered {
		_ = gongenumshape
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, gongenumvalueshape := range gongenumvalueshapeOrdered {
		_ = gongenumvalueshape
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, gongnotelinkshape := range gongnotelinkshapeOrdered {
		_ = gongnotelinkshape
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, gongnoteshape := range gongnoteshapeOrdered {
		_ = gongnoteshape
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, gongstructshape := range gongstructshapeOrdered {
		_ = gongstructshape
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, linkshape := range linkshapeOrdered {
		_ = linkshape
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	res = strings.ReplaceAll(res, "{{Identifiers}}", identifiersDecl.String())
	res = strings.ReplaceAll(res, "{{ValueInitializers}}", initializerStatements.String())
	res = strings.ReplaceAll(res, "{{PointersInitializers}}", pointersInitializesStatements.String())

	if stage.MetaPackageImportAlias != "" {
		res = strings.ReplaceAll(res, "{{ImportPackageDeclaration}}",
			fmt.Sprintf("\n\t%s %s", stage.MetaPackageImportAlias, stage.MetaPackageImportPath))

		res = strings.ReplaceAll(res, "{{ImportPackageDummyDeclaration}}",
			fmt.Sprintf("\nvar _ %s.Stage",
				stage.MetaPackageImportAlias))

		var entries strings.Builder

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
				entries.WriteString(fmt.Sprintf("\n\n\t\"%s\": %s(0),", value.Ident, value.Ident))
			case GONG__ENUM_CAST_STRING:
				entries.WriteString(fmt.Sprintf("\n\n\t\"%s\": %s(\"\"),", value.Ident, value.Ident))
			case GONG__FIELD_VALUE:
				// substitute the second point with "{})."
				joker := "__substitute_for_first_point__"
				valueIdentifier := strings.Replace(value.Ident, ".", joker, 1)
				valueIdentifier = strings.Replace(valueIdentifier, ".", "{}).", 1)
				valueIdentifier = strings.Replace(valueIdentifier, joker, ".", 1)
				entries.WriteString(fmt.Sprintf("\n\n\t\"%s\": (%s,", value.Ident, valueIdentifier))
			case GONG__IDENTIFIER_CONST:
				entries.WriteString(fmt.Sprintf("\n\n\t\"%s\": %s,", value.Ident, value.Ident))
			case GONG__STRUCT_INSTANCE:
				entries.WriteString(fmt.Sprintf("\n\n\t\"%s\": &(%s{}),", value.Ident, value.Ident))
			}
		}

		// res = strings.ReplaceAll(res, "{{EntriesDocLinkStringDocLinkIdentifier}}", entries.String())
	}
	return
}

// insertion point for marshall field methods
func (attributeshape *AttributeShape) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", attributeshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(attributeshape.Name))
	case "IdentifierMeta":
		if str, ok := attributeshape.IdentifierMeta.(string); ok {
			res = MetaFieldStructInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", attributeshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IdentifierMeta")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", str)
		}
	case "FieldTypeAsString":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", attributeshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FieldTypeAsString")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(attributeshape.FieldTypeAsString))
	case "Structname":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", attributeshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Structname")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(attributeshape.Structname))
	case "Fieldtypename":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", attributeshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Fieldtypename")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(attributeshape.Fieldtypename))

	default:
		log.Panicf("Unknown field %s for Gongstruct AttributeShape", fieldName)
	}
	return
}

func (classdiagram *Classdiagram) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", classdiagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(classdiagram.Name))
	case "Description":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", classdiagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Description")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(classdiagram.Description))
	case "IsIncludedInStaticWebSite":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", classdiagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsIncludedInStaticWebSite")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", classdiagram.IsIncludedInStaticWebSite))
	case "ShowNbInstances":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", classdiagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShowNbInstances")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", classdiagram.ShowNbInstances))
	case "ShowMultiplicity":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", classdiagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShowMultiplicity")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", classdiagram.ShowMultiplicity))
	case "ShowLinkNames":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", classdiagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ShowLinkNames")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", classdiagram.ShowLinkNames))
	case "IsInRenameMode":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", classdiagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsInRenameMode")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", classdiagram.IsInRenameMode))
	case "IsExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", classdiagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", classdiagram.IsExpanded))
	case "NodeGongStructsIsExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", classdiagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "NodeGongStructsIsExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", classdiagram.NodeGongStructsIsExpanded))
	case "NodeGongStructNodeExpansion":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", classdiagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "NodeGongStructNodeExpansion")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(classdiagram.NodeGongStructNodeExpansion))
	case "NodeGongEnumsIsExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", classdiagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "NodeGongEnumsIsExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", classdiagram.NodeGongEnumsIsExpanded))
	case "NodeGongEnumNodeExpansion":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", classdiagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "NodeGongEnumNodeExpansion")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(classdiagram.NodeGongEnumNodeExpansion))
	case "NodeGongNotesIsExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", classdiagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "NodeGongNotesIsExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", classdiagram.NodeGongNotesIsExpanded))
	case "NodeGongNoteNodeExpansion":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", classdiagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "NodeGongNoteNodeExpansion")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(classdiagram.NodeGongNoteNodeExpansion))

	case "GongStructShapes":
		var sb strings.Builder
		for _, _gongstructshape := range classdiagram.GongStructShapes {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", classdiagram.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "GongStructShapes")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _gongstructshape.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "GongEnumShapes":
		var sb strings.Builder
		for _, _gongenumshape := range classdiagram.GongEnumShapes {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", classdiagram.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "GongEnumShapes")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _gongenumshape.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "GongNoteShapes":
		var sb strings.Builder
		for _, _gongnoteshape := range classdiagram.GongNoteShapes {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", classdiagram.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "GongNoteShapes")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _gongnoteshape.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Classdiagram", fieldName)
	}
	return
}

func (diagrampackage *DiagramPackage) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagrampackage.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(diagrampackage.Name))
	case "Path":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagrampackage.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Path")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(diagrampackage.Path))
	case "GongModelPath":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagrampackage.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "GongModelPath")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(diagrampackage.GongModelPath))
	case "AbsolutePathToDiagramPackage":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagrampackage.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "AbsolutePathToDiagramPackage")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(diagrampackage.AbsolutePathToDiagramPackage))

	case "Classdiagrams":
		var sb strings.Builder
		for _, _classdiagram := range diagrampackage.Classdiagrams {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", diagrampackage.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Classdiagrams")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _classdiagram.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "SelectedClassdiagram":
		if diagrampackage.SelectedClassdiagram != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", diagrampackage.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SelectedClassdiagram")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", diagrampackage.SelectedClassdiagram.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", diagrampackage.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SelectedClassdiagram")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct DiagramPackage", fieldName)
	}
	return
}

func (gongenumshape *GongEnumShape) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gongenumshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(gongenumshape.Name))
	case "X":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gongenumshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "X")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gongenumshape.X))
	case "Y":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gongenumshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Y")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gongenumshape.Y))
	case "IdentifierMeta":
		if str, ok := gongenumshape.IdentifierMeta.(string); ok {
			res = MetaFieldStructInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", gongenumshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IdentifierMeta")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", str)
		}
	case "Width":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gongenumshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Width")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gongenumshape.Width))
	case "Height":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gongenumshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Height")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gongenumshape.Height))
	case "IsExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gongenumshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", gongenumshape.IsExpanded))

	case "GongEnumValueShapes":
		var sb strings.Builder
		for _, _gongenumvalueshape := range gongenumshape.GongEnumValueShapes {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", gongenumshape.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "GongEnumValueShapes")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _gongenumvalueshape.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct GongEnumShape", fieldName)
	}
	return
}

func (gongenumvalueshape *GongEnumValueShape) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gongenumvalueshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(gongenumvalueshape.Name))
	case "IdentifierMeta":
		if str, ok := gongenumvalueshape.IdentifierMeta.(string); ok {
			res = MetaFieldStructInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", gongenumvalueshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IdentifierMeta")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", str)
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct GongEnumValueShape", fieldName)
	}
	return
}

func (gongnotelinkshape *GongNoteLinkShape) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gongnotelinkshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(gongnotelinkshape.Name))
	case "Identifier":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gongnotelinkshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Identifier")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(gongnotelinkshape.Identifier))
	case "Type":
		if gongnotelinkshape.Type.ToCodeString() != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", gongnotelinkshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Type")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+gongnotelinkshape.Type.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", gongnotelinkshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Type")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "\"\"")
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct GongNoteLinkShape", fieldName)
	}
	return
}

func (gongnoteshape *GongNoteShape) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gongnoteshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(gongnoteshape.Name))
	case "Identifier":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gongnoteshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Identifier")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(gongnoteshape.Identifier))
	case "Body":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gongnoteshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Body")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(gongnoteshape.Body))
	case "BodyHTML":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gongnoteshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "BodyHTML")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(gongnoteshape.BodyHTML))
	case "X":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gongnoteshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "X")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gongnoteshape.X))
	case "Y":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gongnoteshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Y")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gongnoteshape.Y))
	case "Width":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gongnoteshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Width")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gongnoteshape.Width))
	case "Height":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gongnoteshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Height")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gongnoteshape.Height))
	case "Matched":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gongnoteshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Matched")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", gongnoteshape.Matched))
	case "IsExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gongnoteshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", gongnoteshape.IsExpanded))

	case "GongNoteLinkShapes":
		var sb strings.Builder
		for _, _gongnotelinkshape := range gongnoteshape.GongNoteLinkShapes {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", gongnoteshape.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "GongNoteLinkShapes")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _gongnotelinkshape.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct GongNoteShape", fieldName)
	}
	return
}

func (gongstructshape *GongStructShape) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gongstructshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(gongstructshape.Name))
	case "X":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gongstructshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "X")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gongstructshape.X))
	case "Y":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gongstructshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Y")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gongstructshape.Y))
	case "IdentifierMeta":
		if str, ok := gongstructshape.IdentifierMeta.(string); ok {
			res = MetaFieldStructInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", gongstructshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IdentifierMeta")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", str)
		}
	case "Width":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gongstructshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Width")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gongstructshape.Width))
	case "Height":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gongstructshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Height")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gongstructshape.Height))
	case "IsSelected":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gongstructshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsSelected")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", gongstructshape.IsSelected))

	case "AttributeShapes":
		var sb strings.Builder
		for _, _attributeshape := range gongstructshape.AttributeShapes {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", gongstructshape.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "AttributeShapes")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _attributeshape.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "LinkShapes":
		var sb strings.Builder
		for _, _linkshape := range gongstructshape.LinkShapes {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", gongstructshape.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "LinkShapes")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _linkshape.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct GongStructShape", fieldName)
	}
	return
}

func (linkshape *LinkShape) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", linkshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(linkshape.Name))
	case "IdentifierMeta":
		if str, ok := linkshape.IdentifierMeta.(string); ok {
			res = MetaFieldStructInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", linkshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IdentifierMeta")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", str)
		}
	case "FieldTypeIdentifierMeta":
		if str, ok := linkshape.FieldTypeIdentifierMeta.(string); ok {
			res = MetaFieldStructInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", linkshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FieldTypeIdentifierMeta")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", str)
		}
	case "FieldOffsetX":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", linkshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FieldOffsetX")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", linkshape.FieldOffsetX))
	case "FieldOffsetY":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", linkshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FieldOffsetY")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", linkshape.FieldOffsetY))
	case "TargetMultiplicity":
		if linkshape.TargetMultiplicity.ToCodeString() != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", linkshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "TargetMultiplicity")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+linkshape.TargetMultiplicity.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", linkshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "TargetMultiplicity")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "\"\"")
		}
	case "TargetMultiplicityOffsetX":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", linkshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "TargetMultiplicityOffsetX")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", linkshape.TargetMultiplicityOffsetX))
	case "TargetMultiplicityOffsetY":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", linkshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "TargetMultiplicityOffsetY")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", linkshape.TargetMultiplicityOffsetY))
	case "SourceMultiplicity":
		if linkshape.SourceMultiplicity.ToCodeString() != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", linkshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SourceMultiplicity")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+linkshape.SourceMultiplicity.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", linkshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SourceMultiplicity")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "\"\"")
		}
	case "SourceMultiplicityOffsetX":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", linkshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SourceMultiplicityOffsetX")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", linkshape.SourceMultiplicityOffsetX))
	case "SourceMultiplicityOffsetY":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", linkshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SourceMultiplicityOffsetY")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", linkshape.SourceMultiplicityOffsetY))
	case "X":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", linkshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "X")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", linkshape.X))
	case "Y":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", linkshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Y")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", linkshape.Y))
	case "StartOrientation":
		if linkshape.StartOrientation.ToCodeString() != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", linkshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StartOrientation")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+linkshape.StartOrientation.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", linkshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StartOrientation")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "\"\"")
		}
	case "StartRatio":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", linkshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StartRatio")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", linkshape.StartRatio))
	case "EndOrientation":
		if linkshape.EndOrientation.ToCodeString() != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", linkshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EndOrientation")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+linkshape.EndOrientation.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", linkshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EndOrientation")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "\"\"")
		}
	case "EndRatio":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", linkshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EndRatio")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", linkshape.EndRatio))
	case "CornerOffsetRatio":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", linkshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CornerOffsetRatio")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", linkshape.CornerOffsetRatio))

	default:
		log.Panicf("Unknown field %s for Gongstruct LinkShape", fieldName)
	}
	return
}

// insertion point for marshall all fields methods
func (attributeshape *AttributeShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(attributeshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(attributeshape.GongMarshallField(stage, "IdentifierMeta"))
		initializerStatements.WriteString(attributeshape.GongMarshallField(stage, "FieldTypeAsString"))
		initializerStatements.WriteString(attributeshape.GongMarshallField(stage, "Structname"))
		initializerStatements.WriteString(attributeshape.GongMarshallField(stage, "Fieldtypename"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (classdiagram *Classdiagram) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(classdiagram.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(classdiagram.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(classdiagram.GongMarshallField(stage, "IsIncludedInStaticWebSite"))
		pointersInitializesStatements.WriteString(classdiagram.GongMarshallField(stage, "GongStructShapes"))
		pointersInitializesStatements.WriteString(classdiagram.GongMarshallField(stage, "GongEnumShapes"))
		pointersInitializesStatements.WriteString(classdiagram.GongMarshallField(stage, "GongNoteShapes"))
		initializerStatements.WriteString(classdiagram.GongMarshallField(stage, "ShowNbInstances"))
		initializerStatements.WriteString(classdiagram.GongMarshallField(stage, "ShowMultiplicity"))
		initializerStatements.WriteString(classdiagram.GongMarshallField(stage, "ShowLinkNames"))
		initializerStatements.WriteString(classdiagram.GongMarshallField(stage, "IsInRenameMode"))
		initializerStatements.WriteString(classdiagram.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(classdiagram.GongMarshallField(stage, "NodeGongStructsIsExpanded"))
		initializerStatements.WriteString(classdiagram.GongMarshallField(stage, "NodeGongStructNodeExpansion"))
		initializerStatements.WriteString(classdiagram.GongMarshallField(stage, "NodeGongEnumsIsExpanded"))
		initializerStatements.WriteString(classdiagram.GongMarshallField(stage, "NodeGongEnumNodeExpansion"))
		initializerStatements.WriteString(classdiagram.GongMarshallField(stage, "NodeGongNotesIsExpanded"))
		initializerStatements.WriteString(classdiagram.GongMarshallField(stage, "NodeGongNoteNodeExpansion"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (diagrampackage *DiagramPackage) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(diagrampackage.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(diagrampackage.GongMarshallField(stage, "Path"))
		initializerStatements.WriteString(diagrampackage.GongMarshallField(stage, "GongModelPath"))
		pointersInitializesStatements.WriteString(diagrampackage.GongMarshallField(stage, "Classdiagrams"))
		pointersInitializesStatements.WriteString(diagrampackage.GongMarshallField(stage, "SelectedClassdiagram"))
		initializerStatements.WriteString(diagrampackage.GongMarshallField(stage, "AbsolutePathToDiagramPackage"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (gongenumshape *GongEnumShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(gongenumshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(gongenumshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(gongenumshape.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(gongenumshape.GongMarshallField(stage, "IdentifierMeta"))
		pointersInitializesStatements.WriteString(gongenumshape.GongMarshallField(stage, "GongEnumValueShapes"))
		initializerStatements.WriteString(gongenumshape.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(gongenumshape.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(gongenumshape.GongMarshallField(stage, "IsExpanded"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (gongenumvalueshape *GongEnumValueShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(gongenumvalueshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(gongenumvalueshape.GongMarshallField(stage, "IdentifierMeta"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (gongnotelinkshape *GongNoteLinkShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(gongnotelinkshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(gongnotelinkshape.GongMarshallField(stage, "Identifier"))
		initializerStatements.WriteString(gongnotelinkshape.GongMarshallField(stage, "Type"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (gongnoteshape *GongNoteShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(gongnoteshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(gongnoteshape.GongMarshallField(stage, "Identifier"))
		initializerStatements.WriteString(gongnoteshape.GongMarshallField(stage, "Body"))
		initializerStatements.WriteString(gongnoteshape.GongMarshallField(stage, "BodyHTML"))
		initializerStatements.WriteString(gongnoteshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(gongnoteshape.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(gongnoteshape.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(gongnoteshape.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(gongnoteshape.GongMarshallField(stage, "Matched"))
		pointersInitializesStatements.WriteString(gongnoteshape.GongMarshallField(stage, "GongNoteLinkShapes"))
		initializerStatements.WriteString(gongnoteshape.GongMarshallField(stage, "IsExpanded"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (gongstructshape *GongStructShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(gongstructshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(gongstructshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(gongstructshape.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(gongstructshape.GongMarshallField(stage, "IdentifierMeta"))
		pointersInitializesStatements.WriteString(gongstructshape.GongMarshallField(stage, "AttributeShapes"))
		pointersInitializesStatements.WriteString(gongstructshape.GongMarshallField(stage, "LinkShapes"))
		initializerStatements.WriteString(gongstructshape.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(gongstructshape.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(gongstructshape.GongMarshallField(stage, "IsSelected"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (linkshape *LinkShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(linkshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(linkshape.GongMarshallField(stage, "IdentifierMeta"))
		initializerStatements.WriteString(linkshape.GongMarshallField(stage, "FieldTypeIdentifierMeta"))
		initializerStatements.WriteString(linkshape.GongMarshallField(stage, "FieldOffsetX"))
		initializerStatements.WriteString(linkshape.GongMarshallField(stage, "FieldOffsetY"))
		initializerStatements.WriteString(linkshape.GongMarshallField(stage, "TargetMultiplicity"))
		initializerStatements.WriteString(linkshape.GongMarshallField(stage, "TargetMultiplicityOffsetX"))
		initializerStatements.WriteString(linkshape.GongMarshallField(stage, "TargetMultiplicityOffsetY"))
		initializerStatements.WriteString(linkshape.GongMarshallField(stage, "SourceMultiplicity"))
		initializerStatements.WriteString(linkshape.GongMarshallField(stage, "SourceMultiplicityOffsetX"))
		initializerStatements.WriteString(linkshape.GongMarshallField(stage, "SourceMultiplicityOffsetY"))
		initializerStatements.WriteString(linkshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(linkshape.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(linkshape.GongMarshallField(stage, "StartOrientation"))
		initializerStatements.WriteString(linkshape.GongMarshallField(stage, "StartRatio"))
		initializerStatements.WriteString(linkshape.GongMarshallField(stage, "EndOrientation"))
		initializerStatements.WriteString(linkshape.GongMarshallField(stage, "EndRatio"))
		initializerStatements.WriteString(linkshape.GongMarshallField(stage, "CornerOffsetRatio"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
