// generated code - do not edit
package models

import (
	"embed"
	"errors"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"path/filepath"
	"slices"
	"strings"
	"time"

	gong_runtime "github.com/fullstack-lang/gong/pkg/runtime"
	"github.com/fullstack-lang/gong/app/xsd/tests/musicxml/go/models/xlink"
	"github.com/fullstack-lang/gong/app/xsd/tests/musicxml/go/models/xmlxsd"
)

var (
	_ = errors.New
	_ = parser.ParseFile
	_ = filepath.Base
	_ = slices.Index[[]int, int]
	_ = strings.Trim
	_ = time.Hour
)

// swagger:ignore
type GONG__ExpressionType = gong_runtime.GONG__ExpressionType

const (
	GONG__STRUCT_INSTANCE      = gong_runtime.GONG__STRUCT_INSTANCE
	GONG__FIELD_OR_CONST_VALUE = gong_runtime.GONG__FIELD_OR_CONST_VALUE
	GONG__FIELD_VALUE          = gong_runtime.GONG__FIELD_VALUE
	GONG__ENUM_CAST_INT        = gong_runtime.GONG__ENUM_CAST_INT
	GONG__ENUM_CAST_STRING     = gong_runtime.GONG__ENUM_CAST_STRING
	GONG__IDENTIFIER_CONST     = gong_runtime.GONG__IDENTIFIER_CONST
)

// ------------------------------------------------------------------------------------------------
// STATIC AST PARSING LOGIC
// ------------------------------------------------------------------------------------------------

// GongModelUnmarshaller abstracts the logic for setting fields on a staged instance
type GongModelUnmarshaller interface {
	// Initialize creates the struct, stages it, and returns the pointer as 'any'
	Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error)

	// UnmarshallField sets a field's value based on the AST expression
	UnmarshallField(stage *Stage, instance GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error
}

type ModelUnmarshaller = GongModelUnmarshaller

// ParseAstFile Parse pathToFile and stages all instances declared in the file
func (stage *Stage) ParseAstFile(pathToFile string, preserveOrder bool) error {
	inFile, fset, err := gong_runtime.ParseAstFile(pathToFile)
	if err != nil {
		return err
	}
	return stage.ParseAstFileFromAst(inFile, fset, preserveOrder)
}

// ParseAstEmbeddedFile parses the Go source code from an embedded file
func (stage *Stage) ParseAstEmbeddedFile(directory embed.FS, pathToFile string) error {
	inFile, fset, err := gong_runtime.ParseAstEmbeddedFile(directory, pathToFile, stage.GetName())
	if err != nil {
		return err
	}
	return stage.ParseAstFileFromAst(inFile, fset, false)
}

// ParseAstString parses the Go source code from a string
func (stage *Stage) ParseAstString(blob string, preserveOrder bool) error {
	var inFile *ast.File
	var fset *token.FileSet
	var err error
	if strings.HasPrefix(strings.TrimSpace(blob), "package ") || strings.Contains(blob, "\npackage ") {
		fset = token.NewFileSet()
		inFile, err = parser.ParseFile(fset, "", blob, parser.ParseComments)
		if err != nil {
			return errors.New("Unable to parse " + err.Error())
		}
	} else {
		inFile, fset, err = gong_runtime.ParseAstString(blob)
		if err != nil {
			return err
		}
	}
	return stage.ParseAstFileFromAst(inFile, fset, preserveOrder)
}

// ParseAstFileFromAst traverses the AST and stages instances across stage and its sub-stages
func (stage *Stage) ParseAstFileFromAst(inFile *ast.File, fset *token.FileSet, preserveOrder bool) error {
	gong_runtime.CheckModuleVersion(stage.GetProbeIF(), inFile)

	identifierMap := make(map[string]any)

	if stage.XmlxsdStage != nil {
		for _, instance := range stage.XmlxsdStage.GetInstances() {
			identifierMap[instance.GongGetIdentifier(stage.XmlxsdStage)] = instance
		}
	}
	if stage.XlinkStage != nil {
		for _, instance := range stage.XlinkStage.GetInstances() {
			identifierMap[instance.GongGetIdentifier(stage.XlinkStage)] = instance
		}
	}
	if stage != nil {
		for _, instance := range stage.GetInstances() {
			identifierMap[instance.GongGetIdentifier(stage)] = instance
		}
	}

	aliasToCanonical := make(map[string]string)
	for _, imp := range inFile.Imports {
		p := strings.Trim(imp.Path.Value, "\"`")
		alias := filepath.Base(p)
		if imp.Name != nil {
			alias = imp.Name.Name
		}
		switch p {
		case "github.com/fullstack-lang/gong/app/xsd/tests/musicxml/go/models/xmlxsd":
			aliasToCanonical[alias] = "xmlxsd"
		case "github.com/fullstack-lang/gong/app/xsd/tests/musicxml/go/models/xlink":
			aliasToCanonical[alias] = "xlink"
		case "github.com/fullstack-lang/gong/app/xsd/tests/musicxml/go/models":
			aliasToCanonical[alias] = "models"
		}
	}

	ast.Inspect(inFile, func(n ast.Node) bool {
		switch node := n.(type) {
		case *ast.AssignStmt:
			if len(node.Lhs) < 1 || len(node.Rhs) < 1 {
				return true
			}

			// CASE 1: Initialization ( := )
			if node.Tok == token.DEFINE {
				if ident, ok := node.Lhs[0].(*ast.Ident); ok {
					var pkgAlias string
					var typeName string
					var instanceName string

					ast.Inspect(node.Rhs[0], func(expr ast.Node) bool {
						if compLit, ok := expr.(*ast.CompositeLit); ok {
							if selExpr, ok := compLit.Type.(*ast.SelectorExpr); ok {
								if pkgId, ok := selExpr.X.(*ast.Ident); ok {
									pkgAlias = pkgId.Name
								}
								typeName = selExpr.Sel.Name
								for _, elt := range compLit.Elts {
									if kv, ok := elt.(*ast.KeyValueExpr); ok {
										if k, ok := kv.Key.(*ast.Ident); ok && k.Name == "Name" {
											if v, ok := kv.Value.(*ast.BasicLit); ok {
												instanceName = strings.Trim(v.Value, "\"`")
											}
										}
									}
								}
								return false
							}
						}
						return true
					})

					if canonical, ok := aliasToCanonical[pkgAlias]; ok {
						pkgAlias = canonical
					} else if pkgAlias == "" || pkgAlias == "models" || len(inFile.Imports) == 0 {
						pkgAlias = "models"
					}

					switch pkgAlias {
			case "xmlxsd":
				switch typeName {
				case "Xmlxsd_SchemaElement":
					identifierMap[ident.Name] = __gong__stageSetInit(new(xmlxsd.Xmlxsd_SchemaElement), stage.XmlxsdStage, ident.Name, instanceName, preserveOrder)
				}
			case "xlink":
				switch typeName {
				case "ArcType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(xlink.ArcType), stage.XlinkStage, ident.Name, instanceName, preserveOrder)
				case "Extended":
					identifierMap[ident.Name] = __gong__stageSetInit(new(xlink.Extended), stage.XlinkStage, ident.Name, instanceName, preserveOrder)
				case "LocatorType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(xlink.LocatorType), stage.XlinkStage, ident.Name, instanceName, preserveOrder)
				case "ResourceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(xlink.ResourceType), stage.XlinkStage, ident.Name, instanceName, preserveOrder)
				case "Simple":
					identifierMap[ident.Name] = __gong__stageSetInit(new(xlink.Simple), stage.XlinkStage, ident.Name, instanceName, preserveOrder)
				case "TitleEltType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(xlink.TitleEltType), stage.XlinkStage, ident.Name, instanceName, preserveOrder)
				}
			case "models":
				switch typeName {
				case "A_directive":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_directive), stage, ident.Name, instanceName, preserveOrder)
				case "A_measure":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_measure), stage, ident.Name, instanceName, preserveOrder)
				case "A_measure_1":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_measure_1), stage, ident.Name, instanceName, preserveOrder)
				case "A_part":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_part), stage, ident.Name, instanceName, preserveOrder)
				case "A_part_1":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_part_1), stage, ident.Name, instanceName, preserveOrder)
				case "Accidental":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Accidental), stage, ident.Name, instanceName, preserveOrder)
				case "Accidental_mark":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Accidental_mark), stage, ident.Name, instanceName, preserveOrder)
				case "Accidental_text":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Accidental_text), stage, ident.Name, instanceName, preserveOrder)
				case "Accord":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Accord), stage, ident.Name, instanceName, preserveOrder)
				case "Accordion_registration":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Accordion_registration), stage, ident.Name, instanceName, preserveOrder)
				case "Appearance":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Appearance), stage, ident.Name, instanceName, preserveOrder)
				case "Arpeggiate":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Arpeggiate), stage, ident.Name, instanceName, preserveOrder)
				case "Arrow":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Arrow), stage, ident.Name, instanceName, preserveOrder)
				case "Articulations":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Articulations), stage, ident.Name, instanceName, preserveOrder)
				case "Assess":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Assess), stage, ident.Name, instanceName, preserveOrder)
				case "Attributes":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Attributes), stage, ident.Name, instanceName, preserveOrder)
				case "Backup":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Backup), stage, ident.Name, instanceName, preserveOrder)
				case "Bar_style_color":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Bar_style_color), stage, ident.Name, instanceName, preserveOrder)
				case "Barline":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Barline), stage, ident.Name, instanceName, preserveOrder)
				case "Barre":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Barre), stage, ident.Name, instanceName, preserveOrder)
				case "Bass":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Bass), stage, ident.Name, instanceName, preserveOrder)
				case "Bass_step":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Bass_step), stage, ident.Name, instanceName, preserveOrder)
				case "Beam":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Beam), stage, ident.Name, instanceName, preserveOrder)
				case "Beat_repeat":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Beat_repeat), stage, ident.Name, instanceName, preserveOrder)
				case "Beat_unit_tied":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Beat_unit_tied), stage, ident.Name, instanceName, preserveOrder)
				case "Beater":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Beater), stage, ident.Name, instanceName, preserveOrder)
				case "Bend":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Bend), stage, ident.Name, instanceName, preserveOrder)
				case "Bookmark":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Bookmark), stage, ident.Name, instanceName, preserveOrder)
				case "Bracket":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Bracket), stage, ident.Name, instanceName, preserveOrder)
				case "Breath_mark":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Breath_mark), stage, ident.Name, instanceName, preserveOrder)
				case "Caesura":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Caesura), stage, ident.Name, instanceName, preserveOrder)
				case "Cancel":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Cancel), stage, ident.Name, instanceName, preserveOrder)
				case "Clef":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Clef), stage, ident.Name, instanceName, preserveOrder)
				case "Coda":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Coda), stage, ident.Name, instanceName, preserveOrder)
				case "Credit":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Credit), stage, ident.Name, instanceName, preserveOrder)
				case "Dashes":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Dashes), stage, ident.Name, instanceName, preserveOrder)
				case "Defaults":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Defaults), stage, ident.Name, instanceName, preserveOrder)
				case "Degree":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Degree), stage, ident.Name, instanceName, preserveOrder)
				case "Degree_alter":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Degree_alter), stage, ident.Name, instanceName, preserveOrder)
				case "Degree_type":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Degree_type), stage, ident.Name, instanceName, preserveOrder)
				case "Degree_value":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Degree_value), stage, ident.Name, instanceName, preserveOrder)
				case "Direction":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Direction), stage, ident.Name, instanceName, preserveOrder)
				case "Direction_type":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Direction_type), stage, ident.Name, instanceName, preserveOrder)
				case "Distance":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Distance), stage, ident.Name, instanceName, preserveOrder)
				case "Double":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Double), stage, ident.Name, instanceName, preserveOrder)
				case "Dynamics":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Dynamics), stage, ident.Name, instanceName, preserveOrder)
				case "Effect":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Effect), stage, ident.Name, instanceName, preserveOrder)
				case "Elision":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Elision), stage, ident.Name, instanceName, preserveOrder)
				case "Empty":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Empty), stage, ident.Name, instanceName, preserveOrder)
				case "Empty_font":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Empty_font), stage, ident.Name, instanceName, preserveOrder)
				case "Empty_line":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Empty_line), stage, ident.Name, instanceName, preserveOrder)
				case "Empty_placement":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Empty_placement), stage, ident.Name, instanceName, preserveOrder)
				case "Empty_placement_smufl":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Empty_placement_smufl), stage, ident.Name, instanceName, preserveOrder)
				case "Empty_print_object_style_align":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Empty_print_object_style_align), stage, ident.Name, instanceName, preserveOrder)
				case "Empty_print_style":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Empty_print_style), stage, ident.Name, instanceName, preserveOrder)
				case "Empty_print_style_align":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Empty_print_style_align), stage, ident.Name, instanceName, preserveOrder)
				case "Empty_print_style_align_id":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Empty_print_style_align_id), stage, ident.Name, instanceName, preserveOrder)
				case "Empty_trill_sound":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Empty_trill_sound), stage, ident.Name, instanceName, preserveOrder)
				case "Encoding":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Encoding), stage, ident.Name, instanceName, preserveOrder)
				case "Ending":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Ending), stage, ident.Name, instanceName, preserveOrder)
				case "Extend":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Extend), stage, ident.Name, instanceName, preserveOrder)
				case "Feature":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Feature), stage, ident.Name, instanceName, preserveOrder)
				case "Fermata":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Fermata), stage, ident.Name, instanceName, preserveOrder)
				case "Figure":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Figure), stage, ident.Name, instanceName, preserveOrder)
				case "Figured_bass":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Figured_bass), stage, ident.Name, instanceName, preserveOrder)
				case "Fingering":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Fingering), stage, ident.Name, instanceName, preserveOrder)
				case "First_fret":
					identifierMap[ident.Name] = __gong__stageSetInit(new(First_fret), stage, ident.Name, instanceName, preserveOrder)
				case "For_part":
					identifierMap[ident.Name] = __gong__stageSetInit(new(For_part), stage, ident.Name, instanceName, preserveOrder)
				case "Formatted_symbol":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Formatted_symbol), stage, ident.Name, instanceName, preserveOrder)
				case "Formatted_symbol_id":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Formatted_symbol_id), stage, ident.Name, instanceName, preserveOrder)
				case "Formatted_text":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Formatted_text), stage, ident.Name, instanceName, preserveOrder)
				case "Formatted_text_id":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Formatted_text_id), stage, ident.Name, instanceName, preserveOrder)
				case "Forward":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Forward), stage, ident.Name, instanceName, preserveOrder)
				case "Frame":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Frame), stage, ident.Name, instanceName, preserveOrder)
				case "Frame_note":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Frame_note), stage, ident.Name, instanceName, preserveOrder)
				case "Fret":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Fret), stage, ident.Name, instanceName, preserveOrder)
				case "Glass":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Glass), stage, ident.Name, instanceName, preserveOrder)
				case "Glissando":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Glissando), stage, ident.Name, instanceName, preserveOrder)
				case "Glyph":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Glyph), stage, ident.Name, instanceName, preserveOrder)
				case "Grace":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Grace), stage, ident.Name, instanceName, preserveOrder)
				case "Group_barline":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Group_barline), stage, ident.Name, instanceName, preserveOrder)
				case "Group_name":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Group_name), stage, ident.Name, instanceName, preserveOrder)
				case "Group_symbol":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Group_symbol), stage, ident.Name, instanceName, preserveOrder)
				case "Grouping":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Grouping), stage, ident.Name, instanceName, preserveOrder)
				case "Hammer_on_pull_off":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Hammer_on_pull_off), stage, ident.Name, instanceName, preserveOrder)
				case "Handbell":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Handbell), stage, ident.Name, instanceName, preserveOrder)
				case "Harmon_closed":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Harmon_closed), stage, ident.Name, instanceName, preserveOrder)
				case "Harmon_mute":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Harmon_mute), stage, ident.Name, instanceName, preserveOrder)
				case "Harmonic":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Harmonic), stage, ident.Name, instanceName, preserveOrder)
				case "Harmony":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Harmony), stage, ident.Name, instanceName, preserveOrder)
				case "Harmony_alter":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Harmony_alter), stage, ident.Name, instanceName, preserveOrder)
				case "Harp_pedals":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Harp_pedals), stage, ident.Name, instanceName, preserveOrder)
				case "Heel_toe":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Heel_toe), stage, ident.Name, instanceName, preserveOrder)
				case "Hole":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Hole), stage, ident.Name, instanceName, preserveOrder)
				case "Hole_closed":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Hole_closed), stage, ident.Name, instanceName, preserveOrder)
				case "Horizontal_turn":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Horizontal_turn), stage, ident.Name, instanceName, preserveOrder)
				case "Identification":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Identification), stage, ident.Name, instanceName, preserveOrder)
				case "Image":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Image), stage, ident.Name, instanceName, preserveOrder)
				case "Instrument":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Instrument), stage, ident.Name, instanceName, preserveOrder)
				case "Instrument_change":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Instrument_change), stage, ident.Name, instanceName, preserveOrder)
				case "Instrument_link":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Instrument_link), stage, ident.Name, instanceName, preserveOrder)
				case "Interchangeable":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Interchangeable), stage, ident.Name, instanceName, preserveOrder)
				case "Inversion":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Inversion), stage, ident.Name, instanceName, preserveOrder)
				case "Key":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Key), stage, ident.Name, instanceName, preserveOrder)
				case "Key_accidental":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Key_accidental), stage, ident.Name, instanceName, preserveOrder)
				case "Key_octave":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Key_octave), stage, ident.Name, instanceName, preserveOrder)
				case "Kind":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Kind), stage, ident.Name, instanceName, preserveOrder)
				case "Level":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Level), stage, ident.Name, instanceName, preserveOrder)
				case "Line_detail":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Line_detail), stage, ident.Name, instanceName, preserveOrder)
				case "Line_width":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Line_width), stage, ident.Name, instanceName, preserveOrder)
				case "Link":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Link), stage, ident.Name, instanceName, preserveOrder)
				case "Listen":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Listen), stage, ident.Name, instanceName, preserveOrder)
				case "Listening":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Listening), stage, ident.Name, instanceName, preserveOrder)
				case "Lyric":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Lyric), stage, ident.Name, instanceName, preserveOrder)
				case "Lyric_font":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Lyric_font), stage, ident.Name, instanceName, preserveOrder)
				case "Lyric_language":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Lyric_language), stage, ident.Name, instanceName, preserveOrder)
				case "Measure_layout":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Measure_layout), stage, ident.Name, instanceName, preserveOrder)
				case "Measure_numbering":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Measure_numbering), stage, ident.Name, instanceName, preserveOrder)
				case "Measure_repeat":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Measure_repeat), stage, ident.Name, instanceName, preserveOrder)
				case "Measure_style":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Measure_style), stage, ident.Name, instanceName, preserveOrder)
				case "Membrane":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Membrane), stage, ident.Name, instanceName, preserveOrder)
				case "Metal":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Metal), stage, ident.Name, instanceName, preserveOrder)
				case "Metronome":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Metronome), stage, ident.Name, instanceName, preserveOrder)
				case "Metronome_beam":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Metronome_beam), stage, ident.Name, instanceName, preserveOrder)
				case "Metronome_note":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Metronome_note), stage, ident.Name, instanceName, preserveOrder)
				case "Metronome_tied":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Metronome_tied), stage, ident.Name, instanceName, preserveOrder)
				case "Metronome_tuplet":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Metronome_tuplet), stage, ident.Name, instanceName, preserveOrder)
				case "Midi_device":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Midi_device), stage, ident.Name, instanceName, preserveOrder)
				case "Midi_instrument":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Midi_instrument), stage, ident.Name, instanceName, preserveOrder)
				case "Miscellaneous":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Miscellaneous), stage, ident.Name, instanceName, preserveOrder)
				case "Miscellaneous_field":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Miscellaneous_field), stage, ident.Name, instanceName, preserveOrder)
				case "Mordent":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Mordent), stage, ident.Name, instanceName, preserveOrder)
				case "Multiple_rest":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Multiple_rest), stage, ident.Name, instanceName, preserveOrder)
				case "Name_display":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Name_display), stage, ident.Name, instanceName, preserveOrder)
				case "Non_arpeggiate":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Non_arpeggiate), stage, ident.Name, instanceName, preserveOrder)
				case "Notations":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Notations), stage, ident.Name, instanceName, preserveOrder)
				case "Note":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Note), stage, ident.Name, instanceName, preserveOrder)
				case "Note_size":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Note_size), stage, ident.Name, instanceName, preserveOrder)
				case "Note_type":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Note_type), stage, ident.Name, instanceName, preserveOrder)
				case "Notehead":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Notehead), stage, ident.Name, instanceName, preserveOrder)
				case "Notehead_text":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Notehead_text), stage, ident.Name, instanceName, preserveOrder)
				case "Numeral":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Numeral), stage, ident.Name, instanceName, preserveOrder)
				case "Numeral_key":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Numeral_key), stage, ident.Name, instanceName, preserveOrder)
				case "Numeral_root":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Numeral_root), stage, ident.Name, instanceName, preserveOrder)
				case "Octave_shift":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Octave_shift), stage, ident.Name, instanceName, preserveOrder)
				case "Offset":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Offset), stage, ident.Name, instanceName, preserveOrder)
				case "Opus":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Opus), stage, ident.Name, instanceName, preserveOrder)
				case "Ornaments":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Ornaments), stage, ident.Name, instanceName, preserveOrder)
				case "Other_appearance":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Other_appearance), stage, ident.Name, instanceName, preserveOrder)
				case "Other_direction":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Other_direction), stage, ident.Name, instanceName, preserveOrder)
				case "Other_listening":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Other_listening), stage, ident.Name, instanceName, preserveOrder)
				case "Other_notation":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Other_notation), stage, ident.Name, instanceName, preserveOrder)
				case "Other_placement_text":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Other_placement_text), stage, ident.Name, instanceName, preserveOrder)
				case "Other_play":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Other_play), stage, ident.Name, instanceName, preserveOrder)
				case "Other_text":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Other_text), stage, ident.Name, instanceName, preserveOrder)
				case "Page_layout":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Page_layout), stage, ident.Name, instanceName, preserveOrder)
				case "Page_margins":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Page_margins), stage, ident.Name, instanceName, preserveOrder)
				case "Part_clef":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Part_clef), stage, ident.Name, instanceName, preserveOrder)
				case "Part_group":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Part_group), stage, ident.Name, instanceName, preserveOrder)
				case "Part_link":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Part_link), stage, ident.Name, instanceName, preserveOrder)
				case "Part_list":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Part_list), stage, ident.Name, instanceName, preserveOrder)
				case "Part_name":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Part_name), stage, ident.Name, instanceName, preserveOrder)
				case "Part_symbol":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Part_symbol), stage, ident.Name, instanceName, preserveOrder)
				case "Part_transpose":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Part_transpose), stage, ident.Name, instanceName, preserveOrder)
				case "Pedal":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Pedal), stage, ident.Name, instanceName, preserveOrder)
				case "Pedal_tuning":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Pedal_tuning), stage, ident.Name, instanceName, preserveOrder)
				case "Per_minute":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Per_minute), stage, ident.Name, instanceName, preserveOrder)
				case "Percussion":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Percussion), stage, ident.Name, instanceName, preserveOrder)
				case "Pitch":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Pitch), stage, ident.Name, instanceName, preserveOrder)
				case "Pitched":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Pitched), stage, ident.Name, instanceName, preserveOrder)
				case "Placement_text":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Placement_text), stage, ident.Name, instanceName, preserveOrder)
				case "Play":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Play), stage, ident.Name, instanceName, preserveOrder)
				case "Player":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Player), stage, ident.Name, instanceName, preserveOrder)
				case "Principal_voice":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Principal_voice), stage, ident.Name, instanceName, preserveOrder)
				case "Print_XSD":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Print_XSD), stage, ident.Name, instanceName, preserveOrder)
				case "Release":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Release), stage, ident.Name, instanceName, preserveOrder)
				case "Repeat":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Repeat), stage, ident.Name, instanceName, preserveOrder)
				case "Rest":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Rest), stage, ident.Name, instanceName, preserveOrder)
				case "Root":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Root), stage, ident.Name, instanceName, preserveOrder)
				case "Root_step":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Root_step), stage, ident.Name, instanceName, preserveOrder)
				case "Scaling":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Scaling), stage, ident.Name, instanceName, preserveOrder)
				case "Scordatura":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Scordatura), stage, ident.Name, instanceName, preserveOrder)
				case "Score_instrument":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Score_instrument), stage, ident.Name, instanceName, preserveOrder)
				case "Score_part":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Score_part), stage, ident.Name, instanceName, preserveOrder)
				case "Score_partwise":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Score_partwise), stage, ident.Name, instanceName, preserveOrder)
				case "Score_timewise":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Score_timewise), stage, ident.Name, instanceName, preserveOrder)
				case "Segno":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Segno), stage, ident.Name, instanceName, preserveOrder)
				case "Slash":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Slash), stage, ident.Name, instanceName, preserveOrder)
				case "Slide":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Slide), stage, ident.Name, instanceName, preserveOrder)
				case "Slur":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Slur), stage, ident.Name, instanceName, preserveOrder)
				case "Sound":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Sound), stage, ident.Name, instanceName, preserveOrder)
				case "Staff_details":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Staff_details), stage, ident.Name, instanceName, preserveOrder)
				case "Staff_divide":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Staff_divide), stage, ident.Name, instanceName, preserveOrder)
				case "Staff_layout":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Staff_layout), stage, ident.Name, instanceName, preserveOrder)
				case "Staff_size":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Staff_size), stage, ident.Name, instanceName, preserveOrder)
				case "Staff_tuning":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Staff_tuning), stage, ident.Name, instanceName, preserveOrder)
				case "Stem":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Stem), stage, ident.Name, instanceName, preserveOrder)
				case "Stick":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Stick), stage, ident.Name, instanceName, preserveOrder)
				case "String_mute":
					identifierMap[ident.Name] = __gong__stageSetInit(new(String_mute), stage, ident.Name, instanceName, preserveOrder)
				case "String_type":
					identifierMap[ident.Name] = __gong__stageSetInit(new(String_type), stage, ident.Name, instanceName, preserveOrder)
				case "Strong_accent":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Strong_accent), stage, ident.Name, instanceName, preserveOrder)
				case "Style_text":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Style_text), stage, ident.Name, instanceName, preserveOrder)
				case "Supports":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Supports), stage, ident.Name, instanceName, preserveOrder)
				case "Swing":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Swing), stage, ident.Name, instanceName, preserveOrder)
				case "Sync":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Sync), stage, ident.Name, instanceName, preserveOrder)
				case "System_dividers":
					identifierMap[ident.Name] = __gong__stageSetInit(new(System_dividers), stage, ident.Name, instanceName, preserveOrder)
				case "System_layout":
					identifierMap[ident.Name] = __gong__stageSetInit(new(System_layout), stage, ident.Name, instanceName, preserveOrder)
				case "System_margins":
					identifierMap[ident.Name] = __gong__stageSetInit(new(System_margins), stage, ident.Name, instanceName, preserveOrder)
				case "Tap":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Tap), stage, ident.Name, instanceName, preserveOrder)
				case "Technical":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Technical), stage, ident.Name, instanceName, preserveOrder)
				case "Text_element_data":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Text_element_data), stage, ident.Name, instanceName, preserveOrder)
				case "Tie":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Tie), stage, ident.Name, instanceName, preserveOrder)
				case "Tied":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Tied), stage, ident.Name, instanceName, preserveOrder)
				case "Time":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Time), stage, ident.Name, instanceName, preserveOrder)
				case "Time_modification":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Time_modification), stage, ident.Name, instanceName, preserveOrder)
				case "Timpani":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Timpani), stage, ident.Name, instanceName, preserveOrder)
				case "Transpose":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Transpose), stage, ident.Name, instanceName, preserveOrder)
				case "Tremolo":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Tremolo), stage, ident.Name, instanceName, preserveOrder)
				case "Tuplet":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Tuplet), stage, ident.Name, instanceName, preserveOrder)
				case "Tuplet_dot":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Tuplet_dot), stage, ident.Name, instanceName, preserveOrder)
				case "Tuplet_number":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Tuplet_number), stage, ident.Name, instanceName, preserveOrder)
				case "Tuplet_portion":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Tuplet_portion), stage, ident.Name, instanceName, preserveOrder)
				case "Tuplet_type":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Tuplet_type), stage, ident.Name, instanceName, preserveOrder)
				case "Typed_text":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Typed_text), stage, ident.Name, instanceName, preserveOrder)
				case "Unpitched":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Unpitched), stage, ident.Name, instanceName, preserveOrder)
				case "Virtual_instrument":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Virtual_instrument), stage, ident.Name, instanceName, preserveOrder)
				case "Wait":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Wait), stage, ident.Name, instanceName, preserveOrder)
				case "Wavy_line":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Wavy_line), stage, ident.Name, instanceName, preserveOrder)
				case "Wedge":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Wedge), stage, ident.Name, instanceName, preserveOrder)
				case "Wood":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Wood), stage, ident.Name, instanceName, preserveOrder)
				case "Work":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Work), stage, ident.Name, instanceName, preserveOrder)
				}
					}
				}
				return false
			}

			// CASE 2: Assignment ( = )
			if node.Tok == token.ASSIGN {
				if selExpr, ok := node.Lhs[0].(*ast.SelectorExpr); ok {
					if ident, ok := selExpr.X.(*ast.Ident); ok {
						if instance, exists := identifierMap[ident.Name]; exists {
							fieldName := selExpr.Sel.Name
							rhs := node.Rhs[0]
							switch inst := instance.(type) {
				case *xmlxsd.Xmlxsd_SchemaElement:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					}
				case *xlink.ArcType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type_XSD":
						inst.Type_XSD = GongExtractString(rhs)
					case "Arcrole":
						inst.Arcrole = GongExtractString(rhs)
					case "Title":
						inst.Title = GongExtractString(rhs)
					case "Show":
						inst.Show = GongExtractString(rhs)
					case "Actuate":
						inst.Actuate = GongExtractString(rhs)
					case "From":
						inst.From = GongExtractString(rhs)
					case "To":
						inst.To = GongExtractString(rhs)
					case "Title_1":
						__gong__assignSliceOfPointers(&inst.Title_1, rhs, identifierMap)
					}
				case *xlink.Extended:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type_XSD":
						inst.Type_XSD = GongExtractString(rhs)
					case "Role":
						inst.Role = GongExtractString(rhs)
					case "Title":
						inst.Title = GongExtractString(rhs)
					case "Title_1":
						__gong__assignPointer(&inst.Title_1, rhs, identifierMap)
					case "Resource":
						__gong__assignPointer(&inst.Resource, rhs, identifierMap)
					case "Locator":
						__gong__assignPointer(&inst.Locator, rhs, identifierMap)
					case "Arc":
						__gong__assignPointer(&inst.Arc, rhs, identifierMap)
					}
				case *xlink.LocatorType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type_XSD":
						inst.Type_XSD = GongExtractString(rhs)
					case "Href":
						inst.Href = GongExtractString(rhs)
					case "Role":
						inst.Role = GongExtractString(rhs)
					case "Title":
						inst.Title = GongExtractString(rhs)
					case "Label":
						inst.Label = GongExtractString(rhs)
					case "Title_1":
						__gong__assignSliceOfPointers(&inst.Title_1, rhs, identifierMap)
					}
				case *xlink.ResourceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type_XSD":
						inst.Type_XSD = GongExtractString(rhs)
					case "Role":
						inst.Role = GongExtractString(rhs)
					case "Title":
						inst.Title = GongExtractString(rhs)
					case "Label":
						inst.Label = GongExtractString(rhs)
					}
				case *xlink.Simple:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type_XSD":
						inst.Type_XSD = GongExtractString(rhs)
					case "Href":
						inst.Href = GongExtractString(rhs)
					case "Role":
						inst.Role = GongExtractString(rhs)
					case "Arcrole":
						inst.Arcrole = GongExtractString(rhs)
					case "Title":
						inst.Title = GongExtractString(rhs)
					case "Show":
						inst.Show = GongExtractString(rhs)
					case "Actuate":
						inst.Actuate = GongExtractString(rhs)
					}
				case *xlink.TitleEltType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type_XSD":
						inst.Type_XSD = GongExtractString(rhs)
					case "Lang":
						inst.Lang = GongExtractString(rhs)
					}
				case *A_directive:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Lang":
						inst.Lang = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *A_measure:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractString(rhs)
					case "Text":
						inst.Text = GongExtractString(rhs)
					case "Implicit":
						GongUnmarshallEnum(&inst.Implicit, rhs)
					case "Non_controlling":
						GongUnmarshallEnum(&inst.Non_controlling, rhs)
					case "Width":
						inst.Width = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Note":
						__gong__assignSliceOfPointers(&inst.Note, rhs, identifierMap)
					case "Backup":
						__gong__assignSliceOfPointers(&inst.Backup, rhs, identifierMap)
					case "Forward":
						__gong__assignSliceOfPointers(&inst.Forward, rhs, identifierMap)
					case "Direction":
						__gong__assignSliceOfPointers(&inst.Direction, rhs, identifierMap)
					case "Attributes":
						__gong__assignSliceOfPointers(&inst.Attributes, rhs, identifierMap)
					case "Harmony":
						__gong__assignSliceOfPointers(&inst.Harmony, rhs, identifierMap)
					case "Figured_bass":
						__gong__assignSliceOfPointers(&inst.Figured_bass, rhs, identifierMap)
					case "Print_XSD":
						__gong__assignSliceOfPointers(&inst.Print_XSD, rhs, identifierMap)
					case "Sound":
						__gong__assignSliceOfPointers(&inst.Sound, rhs, identifierMap)
					case "Listening":
						__gong__assignSliceOfPointers(&inst.Listening, rhs, identifierMap)
					case "Barline":
						__gong__assignSliceOfPointers(&inst.Barline, rhs, identifierMap)
					case "Grouping":
						__gong__assignSliceOfPointers(&inst.Grouping, rhs, identifierMap)
					case "Link":
						__gong__assignSliceOfPointers(&inst.Link, rhs, identifierMap)
					case "Bookmark":
						__gong__assignSliceOfPointers(&inst.Bookmark, rhs, identifierMap)
					}
				case *A_measure_1:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractString(rhs)
					case "Text":
						inst.Text = GongExtractString(rhs)
					case "Implicit":
						GongUnmarshallEnum(&inst.Implicit, rhs)
					case "Non_controlling":
						GongUnmarshallEnum(&inst.Non_controlling, rhs)
					case "Width":
						inst.Width = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Part":
						__gong__assignSliceOfPointers(&inst.Part, rhs, identifierMap)
					}
				case *A_part:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Measure":
						__gong__assignSliceOfPointers(&inst.Measure, rhs, identifierMap)
					}
				case *A_part_1:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Note":
						__gong__assignSliceOfPointers(&inst.Note, rhs, identifierMap)
					case "Backup":
						__gong__assignSliceOfPointers(&inst.Backup, rhs, identifierMap)
					case "Forward":
						__gong__assignSliceOfPointers(&inst.Forward, rhs, identifierMap)
					case "Direction":
						__gong__assignSliceOfPointers(&inst.Direction, rhs, identifierMap)
					case "Attributes":
						__gong__assignSliceOfPointers(&inst.Attributes, rhs, identifierMap)
					case "Harmony":
						__gong__assignSliceOfPointers(&inst.Harmony, rhs, identifierMap)
					case "Figured_bass":
						__gong__assignSliceOfPointers(&inst.Figured_bass, rhs, identifierMap)
					case "Print_XSD":
						__gong__assignSliceOfPointers(&inst.Print_XSD, rhs, identifierMap)
					case "Sound":
						__gong__assignSliceOfPointers(&inst.Sound, rhs, identifierMap)
					case "Listening":
						__gong__assignSliceOfPointers(&inst.Listening, rhs, identifierMap)
					case "Barline":
						__gong__assignSliceOfPointers(&inst.Barline, rhs, identifierMap)
					case "Grouping":
						__gong__assignSliceOfPointers(&inst.Grouping, rhs, identifierMap)
					case "Link":
						__gong__assignSliceOfPointers(&inst.Link, rhs, identifierMap)
					case "Bookmark":
						__gong__assignSliceOfPointers(&inst.Bookmark, rhs, identifierMap)
					}
				case *Accidental:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Cautionary":
						GongUnmarshallEnum(&inst.Cautionary, rhs)
					case "Editorial":
						GongUnmarshallEnum(&inst.Editorial, rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "Parentheses":
						GongUnmarshallEnum(&inst.Parentheses, rhs)
					case "Bracket":
						GongUnmarshallEnum(&inst.Bracket, rhs)
					case "Size":
						GongUnmarshallEnum(&inst.Size, rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "EnclosedText":
						GongUnmarshallEnum(&inst.EnclosedText, rhs)
					}
				case *Accidental_mark:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "Parentheses":
						GongUnmarshallEnum(&inst.Parentheses, rhs)
					case "Bracket":
						GongUnmarshallEnum(&inst.Bracket, rhs)
					case "Size":
						GongUnmarshallEnum(&inst.Size, rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						GongUnmarshallEnum(&inst.Placement, rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "EnclosedText":
						GongUnmarshallEnum(&inst.EnclosedText, rhs)
					}
				case *Accidental_text:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "Lang":
						inst.Lang = GongExtractString(rhs)
					case "Space":
						inst.Space = GongExtractString(rhs)
					case "Justify":
						GongUnmarshallEnum(&inst.Justify, rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Halign":
						GongUnmarshallEnum(&inst.Halign, rhs)
					case "Valign":
						GongUnmarshallEnum(&inst.Valign, rhs)
					case "Underline":
						inst.Underline = GongExtractInt(rhs)
					case "Overline":
						inst.Overline = GongExtractInt(rhs)
					case "Line_through":
						inst.Line_through = GongExtractInt(rhs)
					case "Rotation":
						inst.Rotation = GongExtractString(rhs)
					case "Letter_spacing":
						inst.Letter_spacing = GongExtractString(rhs)
					case "Line_height":
						inst.Line_height = GongExtractString(rhs)
					case "Dir":
						GongUnmarshallEnum(&inst.Dir, rhs)
					case "Enclosure":
						GongUnmarshallEnum(&inst.Enclosure, rhs)
					case "EnclosedText":
						GongUnmarshallEnum(&inst.EnclosedText, rhs)
					}
				case *Accord:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "String_XSD":
						inst.String_XSD = GongExtractInt(rhs)
					case "Tuning_step":
						GongUnmarshallEnum(&inst.Tuning_step, rhs)
					case "Tuning_alter":
						inst.Tuning_alter = GongExtractString(rhs)
					case "Tuning_octave":
						inst.Tuning_octave = GongExtractInt(rhs)
					}
				case *Accordion_registration:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Halign":
						GongUnmarshallEnum(&inst.Halign, rhs)
					case "Valign":
						GongUnmarshallEnum(&inst.Valign, rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Accordion_high":
						inst.Accordion_high = GongExtractString(rhs)
					case "Accordion_middle":
						inst.Accordion_middle = GongExtractInt(rhs)
					case "Accordion_low":
						inst.Accordion_low = GongExtractString(rhs)
					}
				case *Appearance:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Line_width":
						__gong__assignSliceOfPointers(&inst.Line_width, rhs, identifierMap)
					case "Note_size":
						__gong__assignSliceOfPointers(&inst.Note_size, rhs, identifierMap)
					case "Distance":
						__gong__assignSliceOfPointers(&inst.Distance, rhs, identifierMap)
					case "Glyph":
						__gong__assignSliceOfPointers(&inst.Glyph, rhs, identifierMap)
					case "Other_appearance":
						__gong__assignSliceOfPointers(&inst.Other_appearance, rhs, identifierMap)
					}
				case *Arpeggiate:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Direction":
						GongUnmarshallEnum(&inst.Direction, rhs)
					case "Unbroken":
						GongUnmarshallEnum(&inst.Unbroken, rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Placement":
						GongUnmarshallEnum(&inst.Placement, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					}
				case *Arrow:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						GongUnmarshallEnum(&inst.Placement, rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "Arrow_direction":
						GongUnmarshallEnum(&inst.Arrow_direction, rhs)
					case "Arrow_style":
						GongUnmarshallEnum(&inst.Arrow_style, rhs)
					case "Arrowhead":
						inst.Arrowhead = GongExtractString(rhs)
					case "Circular_arrow":
						GongUnmarshallEnum(&inst.Circular_arrow, rhs)
					}
				case *Articulations:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Accent":
						__gong__assignSliceOfPointers(&inst.Accent, rhs, identifierMap)
					case "Strong_accent":
						__gong__assignSliceOfPointers(&inst.Strong_accent, rhs, identifierMap)
					case "Staccato":
						__gong__assignSliceOfPointers(&inst.Staccato, rhs, identifierMap)
					case "Tenuto":
						__gong__assignSliceOfPointers(&inst.Tenuto, rhs, identifierMap)
					case "Detached_legato":
						__gong__assignSliceOfPointers(&inst.Detached_legato, rhs, identifierMap)
					case "Staccatissimo":
						__gong__assignSliceOfPointers(&inst.Staccatissimo, rhs, identifierMap)
					case "Spiccato":
						__gong__assignSliceOfPointers(&inst.Spiccato, rhs, identifierMap)
					case "Scoop":
						__gong__assignSliceOfPointers(&inst.Scoop, rhs, identifierMap)
					case "Plop":
						__gong__assignSliceOfPointers(&inst.Plop, rhs, identifierMap)
					case "Doit":
						__gong__assignSliceOfPointers(&inst.Doit, rhs, identifierMap)
					case "Falloff":
						__gong__assignSliceOfPointers(&inst.Falloff, rhs, identifierMap)
					case "Breath_mark":
						__gong__assignSliceOfPointers(&inst.Breath_mark, rhs, identifierMap)
					case "Caesura":
						__gong__assignSliceOfPointers(&inst.Caesura, rhs, identifierMap)
					case "Stress":
						__gong__assignSliceOfPointers(&inst.Stress, rhs, identifierMap)
					case "Unstress":
						__gong__assignSliceOfPointers(&inst.Unstress, rhs, identifierMap)
					case "Soft_accent":
						__gong__assignSliceOfPointers(&inst.Soft_accent, rhs, identifierMap)
					case "Other_articulation":
						__gong__assignSliceOfPointers(&inst.Other_articulation, rhs, identifierMap)
					}
				case *Assess:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type_XSD":
						GongUnmarshallEnum(&inst.Type_XSD, rhs)
					case "Player":
						inst.Player = GongExtractString(rhs)
					case "Time_only":
						inst.Time_only = GongExtractString(rhs)
					}
				case *Attributes:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Footnote":
						__gong__assignPointer(&inst.Footnote, rhs, identifierMap)
					case "Level":
						__gong__assignPointer(&inst.Level, rhs, identifierMap)
					case "Divisions":
						inst.Divisions = GongExtractString(rhs)
					case "Key":
						__gong__assignSliceOfPointers(&inst.Key, rhs, identifierMap)
					case "Time":
						inst.Time = GongExtractString(rhs)
					case "Staves":
						inst.Staves = GongExtractInt(rhs)
					case "Part_symbol":
						__gong__assignPointer(&inst.Part_symbol, rhs, identifierMap)
					case "Instruments":
						inst.Instruments = GongExtractInt(rhs)
					case "Clef":
						__gong__assignSliceOfPointers(&inst.Clef, rhs, identifierMap)
					case "Staff_details":
						__gong__assignSliceOfPointers(&inst.Staff_details, rhs, identifierMap)
					case "Transpose":
						__gong__assignSliceOfPointers(&inst.Transpose, rhs, identifierMap)
					case "For_part":
						__gong__assignSliceOfPointers(&inst.For_part, rhs, identifierMap)
					case "Directive":
						__gong__assignSliceOfPointers(&inst.Directive, rhs, identifierMap)
					case "Measure_style":
						__gong__assignSliceOfPointers(&inst.Measure_style, rhs, identifierMap)
					}
				case *Backup:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Footnote":
						__gong__assignPointer(&inst.Footnote, rhs, identifierMap)
					case "Level":
						__gong__assignPointer(&inst.Level, rhs, identifierMap)
					}
				case *Bar_style_color:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "EnclosedText":
						GongUnmarshallEnum(&inst.EnclosedText, rhs)
					}
				case *Barline:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Location":
						GongUnmarshallEnum(&inst.Location, rhs)
					case "Segno":
						inst.Segno = GongExtractString(rhs)
					case "Coda":
						inst.Coda = GongExtractString(rhs)
					case "Divisions":
						inst.Divisions = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Bar_style":
						__gong__assignPointer(&inst.Bar_style, rhs, identifierMap)
					case "Footnote":
						__gong__assignPointer(&inst.Footnote, rhs, identifierMap)
					case "Level":
						__gong__assignPointer(&inst.Level, rhs, identifierMap)
					case "Wavy_line":
						__gong__assignPointer(&inst.Wavy_line, rhs, identifierMap)
					case "Segno_1":
						__gong__assignPointer(&inst.Segno_1, rhs, identifierMap)
					case "Coda_1":
						__gong__assignPointer(&inst.Coda_1, rhs, identifierMap)
					case "Fermata":
						__gong__assignPointer(&inst.Fermata, rhs, identifierMap)
					case "Ending":
						__gong__assignPointer(&inst.Ending, rhs, identifierMap)
					case "Repeat":
						__gong__assignPointer(&inst.Repeat, rhs, identifierMap)
					}
				case *Barre:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type_XSD":
						GongUnmarshallEnum(&inst.Type_XSD, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					}
				case *Bass:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Arrangement":
						GongUnmarshallEnum(&inst.Arrangement, rhs)
					case "Bass_separator":
						__gong__assignPointer(&inst.Bass_separator, rhs, identifierMap)
					case "Bass_step":
						__gong__assignPointer(&inst.Bass_step, rhs, identifierMap)
					case "Bass_alter":
						__gong__assignPointer(&inst.Bass_alter, rhs, identifierMap)
					}
				case *Bass_step:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Text":
						inst.Text = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "EnclosedText":
						GongUnmarshallEnum(&inst.EnclosedText, rhs)
					}
				case *Beam:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Repeater":
						GongUnmarshallEnum(&inst.Repeater, rhs)
					case "Fan":
						GongUnmarshallEnum(&inst.Fan, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "EnclosedText":
						GongUnmarshallEnum(&inst.EnclosedText, rhs)
					}
				case *Beat_repeat:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type_XSD":
						GongUnmarshallEnum(&inst.Type_XSD, rhs)
					case "Slashes":
						inst.Slashes = GongExtractInt(rhs)
					case "Use_dots":
						GongUnmarshallEnum(&inst.Use_dots, rhs)
					case "Slash_type":
						GongUnmarshallEnum(&inst.Slash_type, rhs)
					case "Slash_dot":
						inst.Slash_dot = GongExtractString(rhs)
					case "Except_voice":
						inst.Except_voice = GongExtractString(rhs)
					}
				case *Beat_unit_tied:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Beat_unit":
						GongUnmarshallEnum(&inst.Beat_unit, rhs)
					case "Beat_unit_dot":
						inst.Beat_unit_dot = GongExtractString(rhs)
					}
				case *Beater:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Tip":
						GongUnmarshallEnum(&inst.Tip, rhs)
					case "EnclosedText":
						GongUnmarshallEnum(&inst.EnclosedText, rhs)
					}
				case *Bend:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Shape":
						GongUnmarshallEnum(&inst.Shape, rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Accelerate":
						GongUnmarshallEnum(&inst.Accelerate, rhs)
					case "Beats":
						inst.Beats = GongExtractString(rhs)
					case "First_beat":
						inst.First_beat = GongExtractString(rhs)
					case "Last_beat":
						inst.Last_beat = GongExtractString(rhs)
					case "Bend_alter":
						inst.Bend_alter = GongExtractString(rhs)
					case "Pre_bend":
						inst.Pre_bend = GongExtractString(rhs)
					case "Release":
						__gong__assignPointer(&inst.Release, rhs, identifierMap)
					case "With_bar":
						__gong__assignPointer(&inst.With_bar, rhs, identifierMap)
					}
				case *Bookmark:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "NameXSD":
						inst.NameXSD = GongExtractString(rhs)
					case "Element":
						inst.Element = GongExtractString(rhs)
					case "Position":
						inst.Position = GongExtractInt(rhs)
					}
				case *Bracket:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type_XSD":
						GongUnmarshallEnum(&inst.Type_XSD, rhs)
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Line_end":
						GongUnmarshallEnum(&inst.Line_end, rhs)
					case "End_length":
						inst.End_length = GongExtractString(rhs)
					case "Line_type":
						GongUnmarshallEnum(&inst.Line_type, rhs)
					case "Dash_length":
						inst.Dash_length = GongExtractString(rhs)
					case "Space_length":
						inst.Space_length = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					}
				case *Breath_mark:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						GongUnmarshallEnum(&inst.Placement, rhs)
					case "EnclosedText":
						GongUnmarshallEnum(&inst.EnclosedText, rhs)
					}
				case *Caesura:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						GongUnmarshallEnum(&inst.Placement, rhs)
					case "EnclosedText":
						GongUnmarshallEnum(&inst.EnclosedText, rhs)
					}
				case *Cancel:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Location":
						GongUnmarshallEnum(&inst.Location, rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractInt(rhs)
					}
				case *Clef:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Additional":
						GongUnmarshallEnum(&inst.Additional, rhs)
					case "Size":
						GongUnmarshallEnum(&inst.Size, rhs)
					case "After_barline":
						GongUnmarshallEnum(&inst.After_barline, rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Print_object":
						GongUnmarshallEnum(&inst.Print_object, rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Sign":
						GongUnmarshallEnum(&inst.Sign, rhs)
					case "Line":
						inst.Line = GongExtractInt(rhs)
					case "Clef_octave_change":
						inst.Clef_octave_change = GongExtractInt(rhs)
					}
				case *Coda:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Halign":
						GongUnmarshallEnum(&inst.Halign, rhs)
					case "Valign":
						GongUnmarshallEnum(&inst.Valign, rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					}
				case *Credit:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Page":
						inst.Page = GongExtractInt(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Credit_type":
						inst.Credit_type = GongExtractString(rhs)
					case "Link":
						__gong__assignSliceOfPointers(&inst.Link, rhs, identifierMap)
					case "Bookmark":
						__gong__assignSliceOfPointers(&inst.Bookmark, rhs, identifierMap)
					case "Credit_image":
						__gong__assignPointer(&inst.Credit_image, rhs, identifierMap)
					case "Credit_words":
						__gong__assignPointer(&inst.Credit_words, rhs, identifierMap)
					case "Credit_symbol":
						__gong__assignPointer(&inst.Credit_symbol, rhs, identifierMap)
					}
				case *Dashes:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type_XSD":
						GongUnmarshallEnum(&inst.Type_XSD, rhs)
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Dash_length":
						inst.Dash_length = GongExtractString(rhs)
					case "Space_length":
						inst.Space_length = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					}
				case *Defaults:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Scaling":
						__gong__assignPointer(&inst.Scaling, rhs, identifierMap)
					case "Concert_score":
						inst.Concert_score = GongExtractString(rhs)
					case "Page_layout":
						__gong__assignPointer(&inst.Page_layout, rhs, identifierMap)
					case "System_layout":
						__gong__assignPointer(&inst.System_layout, rhs, identifierMap)
					case "Staff_layout":
						__gong__assignSliceOfPointers(&inst.Staff_layout, rhs, identifierMap)
					case "Appearance":
						__gong__assignPointer(&inst.Appearance, rhs, identifierMap)
					case "Music_font":
						__gong__assignPointer(&inst.Music_font, rhs, identifierMap)
					case "Word_font":
						__gong__assignPointer(&inst.Word_font, rhs, identifierMap)
					case "Lyric_font":
						__gong__assignSliceOfPointers(&inst.Lyric_font, rhs, identifierMap)
					case "Lyric_language":
						__gong__assignSliceOfPointers(&inst.Lyric_language, rhs, identifierMap)
					}
				case *Degree:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Print_object":
						GongUnmarshallEnum(&inst.Print_object, rhs)
					case "Degree_value":
						__gong__assignPointer(&inst.Degree_value, rhs, identifierMap)
					case "Degree_alter":
						__gong__assignPointer(&inst.Degree_alter, rhs, identifierMap)
					case "Degree_type":
						__gong__assignPointer(&inst.Degree_type, rhs, identifierMap)
					}
				case *Degree_alter:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Plus_minus":
						GongUnmarshallEnum(&inst.Plus_minus, rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Degree_type:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Text":
						inst.Text = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "EnclosedText":
						GongUnmarshallEnum(&inst.EnclosedText, rhs)
					}
				case *Degree_value:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Symbol":
						GongUnmarshallEnum(&inst.Symbol, rhs)
					case "Text":
						inst.Text = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractInt(rhs)
					}
				case *Direction:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Placement":
						GongUnmarshallEnum(&inst.Placement, rhs)
					case "Directive":
						GongUnmarshallEnum(&inst.Directive, rhs)
					case "System":
						inst.System = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Direction_type":
						__gong__assignSliceOfPointers(&inst.Direction_type, rhs, identifierMap)
					case "Offset":
						__gong__assignPointer(&inst.Offset, rhs, identifierMap)
					case "Footnote":
						__gong__assignPointer(&inst.Footnote, rhs, identifierMap)
					case "Level":
						__gong__assignPointer(&inst.Level, rhs, identifierMap)
					case "Voice":
						inst.Voice = GongExtractString(rhs)
					case "Staff":
						inst.Staff = GongExtractInt(rhs)
					case "Sound":
						__gong__assignPointer(&inst.Sound, rhs, identifierMap)
					case "Listening":
						__gong__assignPointer(&inst.Listening, rhs, identifierMap)
					}
				case *Direction_type:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Rehearsal":
						__gong__assignSliceOfPointers(&inst.Rehearsal, rhs, identifierMap)
					case "Segno":
						__gong__assignSliceOfPointers(&inst.Segno, rhs, identifierMap)
					case "Coda":
						__gong__assignSliceOfPointers(&inst.Coda, rhs, identifierMap)
					case "Words":
						__gong__assignSliceOfPointers(&inst.Words, rhs, identifierMap)
					case "Symbol":
						__gong__assignSliceOfPointers(&inst.Symbol, rhs, identifierMap)
					case "Wedge":
						__gong__assignPointer(&inst.Wedge, rhs, identifierMap)
					case "Dynamics":
						__gong__assignSliceOfPointers(&inst.Dynamics, rhs, identifierMap)
					case "Dashes":
						__gong__assignPointer(&inst.Dashes, rhs, identifierMap)
					case "Bracket":
						__gong__assignPointer(&inst.Bracket, rhs, identifierMap)
					case "Pedal":
						__gong__assignPointer(&inst.Pedal, rhs, identifierMap)
					case "Metronome":
						__gong__assignPointer(&inst.Metronome, rhs, identifierMap)
					case "Octave_shift":
						__gong__assignPointer(&inst.Octave_shift, rhs, identifierMap)
					case "Harp_pedals":
						__gong__assignPointer(&inst.Harp_pedals, rhs, identifierMap)
					case "Damp":
						__gong__assignPointer(&inst.Damp, rhs, identifierMap)
					case "Damp_all":
						__gong__assignPointer(&inst.Damp_all, rhs, identifierMap)
					case "Eyeglasses":
						__gong__assignPointer(&inst.Eyeglasses, rhs, identifierMap)
					case "String_mute":
						__gong__assignPointer(&inst.String_mute, rhs, identifierMap)
					case "Scordatura":
						__gong__assignPointer(&inst.Scordatura, rhs, identifierMap)
					case "Image":
						__gong__assignPointer(&inst.Image, rhs, identifierMap)
					case "Principal_voice":
						__gong__assignPointer(&inst.Principal_voice, rhs, identifierMap)
					case "Percussion":
						__gong__assignSliceOfPointers(&inst.Percussion, rhs, identifierMap)
					case "Accordion_registration":
						__gong__assignPointer(&inst.Accordion_registration, rhs, identifierMap)
					case "Staff_divide":
						__gong__assignPointer(&inst.Staff_divide, rhs, identifierMap)
					case "Other_direction":
						__gong__assignPointer(&inst.Other_direction, rhs, identifierMap)
					}
				case *Distance:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type_XSD":
						inst.Type_XSD = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Double:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Above":
						GongUnmarshallEnum(&inst.Above, rhs)
					}
				case *Dynamics:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Halign":
						GongUnmarshallEnum(&inst.Halign, rhs)
					case "Valign":
						GongUnmarshallEnum(&inst.Valign, rhs)
					case "Placement":
						GongUnmarshallEnum(&inst.Placement, rhs)
					case "Underline":
						inst.Underline = GongExtractInt(rhs)
					case "Overline":
						inst.Overline = GongExtractInt(rhs)
					case "Line_through":
						inst.Line_through = GongExtractInt(rhs)
					case "Enclosure":
						GongUnmarshallEnum(&inst.Enclosure, rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "P":
						inst.P = GongExtractString(rhs)
					case "Pp":
						inst.Pp = GongExtractString(rhs)
					case "Ppp":
						inst.Ppp = GongExtractString(rhs)
					case "Pppp":
						inst.Pppp = GongExtractString(rhs)
					case "Ppppp":
						inst.Ppppp = GongExtractString(rhs)
					case "Pppppp":
						inst.Pppppp = GongExtractString(rhs)
					case "F":
						inst.F = GongExtractString(rhs)
					case "Ff":
						inst.Ff = GongExtractString(rhs)
					case "Fff":
						inst.Fff = GongExtractString(rhs)
					case "Ffff":
						inst.Ffff = GongExtractString(rhs)
					case "Fffff":
						inst.Fffff = GongExtractString(rhs)
					case "Ffffff":
						inst.Ffffff = GongExtractString(rhs)
					case "Mp":
						inst.Mp = GongExtractString(rhs)
					case "Mf":
						inst.Mf = GongExtractString(rhs)
					case "Sf":
						inst.Sf = GongExtractString(rhs)
					case "Sfp":
						inst.Sfp = GongExtractString(rhs)
					case "Sfpp":
						inst.Sfpp = GongExtractString(rhs)
					case "Fp":
						inst.Fp = GongExtractString(rhs)
					case "Rf":
						inst.Rf = GongExtractString(rhs)
					case "Rfz":
						inst.Rfz = GongExtractString(rhs)
					case "Sfz":
						inst.Sfz = GongExtractString(rhs)
					case "Sffz":
						inst.Sffz = GongExtractString(rhs)
					case "Fz":
						inst.Fz = GongExtractString(rhs)
					case "N":
						inst.N = GongExtractString(rhs)
					case "Pf":
						inst.Pf = GongExtractString(rhs)
					case "Sfzp":
						inst.Sfzp = GongExtractString(rhs)
					case "Other_dynamics":
						__gong__assignSliceOfPointers(&inst.Other_dynamics, rhs, identifierMap)
					}
				case *Effect:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "EnclosedText":
						GongUnmarshallEnum(&inst.EnclosedText, rhs)
					}
				case *Elision:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Empty:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					}
				case *Empty_font:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					}
				case *Empty_line:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Line_shape":
						GongUnmarshallEnum(&inst.Line_shape, rhs)
					case "Line_type":
						GongUnmarshallEnum(&inst.Line_type, rhs)
					case "Line_length":
						GongUnmarshallEnum(&inst.Line_length, rhs)
					case "Dash_length":
						inst.Dash_length = GongExtractString(rhs)
					case "Space_length":
						inst.Space_length = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						GongUnmarshallEnum(&inst.Placement, rhs)
					}
				case *Empty_placement:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						GongUnmarshallEnum(&inst.Placement, rhs)
					}
				case *Empty_placement_smufl:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						GongUnmarshallEnum(&inst.Placement, rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					}
				case *Empty_print_object_style_align:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Print_object":
						GongUnmarshallEnum(&inst.Print_object, rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Halign":
						GongUnmarshallEnum(&inst.Halign, rhs)
					case "Valign":
						GongUnmarshallEnum(&inst.Valign, rhs)
					}
				case *Empty_print_style:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					}
				case *Empty_print_style_align:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Halign":
						GongUnmarshallEnum(&inst.Halign, rhs)
					case "Valign":
						GongUnmarshallEnum(&inst.Valign, rhs)
					}
				case *Empty_print_style_align_id:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Halign":
						GongUnmarshallEnum(&inst.Halign, rhs)
					case "Valign":
						GongUnmarshallEnum(&inst.Valign, rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					}
				case *Empty_trill_sound:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						GongUnmarshallEnum(&inst.Placement, rhs)
					case "Start_note":
						GongUnmarshallEnum(&inst.Start_note, rhs)
					case "Trill_step":
						GongUnmarshallEnum(&inst.Trill_step, rhs)
					case "Two_note_turn":
						GongUnmarshallEnum(&inst.Two_note_turn, rhs)
					case "Accelerate":
						GongUnmarshallEnum(&inst.Accelerate, rhs)
					case "Beats":
						inst.Beats = GongExtractString(rhs)
					case "Second_beat":
						inst.Second_beat = GongExtractString(rhs)
					case "Last_beat":
						inst.Last_beat = GongExtractString(rhs)
					}
				case *Encoding:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Encoding_date":
						inst.Encoding_date = GongExtractString(rhs)
					case "Encoder":
						__gong__assignSliceOfPointers(&inst.Encoder, rhs, identifierMap)
					case "Software":
						inst.Software = GongExtractString(rhs)
					case "Encoding_description":
						inst.Encoding_description = GongExtractString(rhs)
					case "Supports":
						__gong__assignSliceOfPointers(&inst.Supports, rhs, identifierMap)
					}
				case *Ending:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractString(rhs)
					case "Type_XSD":
						GongUnmarshallEnum(&inst.Type_XSD, rhs)
					case "End_length":
						inst.End_length = GongExtractString(rhs)
					case "Text_x":
						inst.Text_x = GongExtractString(rhs)
					case "Text_y":
						inst.Text_y = GongExtractString(rhs)
					case "Print_object":
						GongUnmarshallEnum(&inst.Print_object, rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "System":
						inst.System = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Extend:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type_XSD":
						GongUnmarshallEnum(&inst.Type_XSD, rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					}
				case *Feature:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type_XSD":
						inst.Type_XSD = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Fermata:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type_XSD":
						GongUnmarshallEnum(&inst.Type_XSD, rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "EnclosedText":
						GongUnmarshallEnum(&inst.EnclosedText, rhs)
					}
				case *Figure:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Prefix":
						__gong__assignPointer(&inst.Prefix, rhs, identifierMap)
					case "Figure_number":
						__gong__assignPointer(&inst.Figure_number, rhs, identifierMap)
					case "Suffix":
						__gong__assignPointer(&inst.Suffix, rhs, identifierMap)
					case "Extend":
						__gong__assignPointer(&inst.Extend, rhs, identifierMap)
					case "Footnote":
						__gong__assignPointer(&inst.Footnote, rhs, identifierMap)
					case "Level":
						__gong__assignPointer(&inst.Level, rhs, identifierMap)
					}
				case *Figured_bass:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Parentheses":
						GongUnmarshallEnum(&inst.Parentheses, rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Halign":
						GongUnmarshallEnum(&inst.Halign, rhs)
					case "Valign":
						GongUnmarshallEnum(&inst.Valign, rhs)
					case "Placement":
						GongUnmarshallEnum(&inst.Placement, rhs)
					case "Print_dot":
						GongUnmarshallEnum(&inst.Print_dot, rhs)
					case "Print_lyric":
						GongUnmarshallEnum(&inst.Print_lyric, rhs)
					case "Print_object":
						GongUnmarshallEnum(&inst.Print_object, rhs)
					case "Print_spacing":
						GongUnmarshallEnum(&inst.Print_spacing, rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Figure":
						__gong__assignSliceOfPointers(&inst.Figure, rhs, identifierMap)
					case "Footnote":
						__gong__assignPointer(&inst.Footnote, rhs, identifierMap)
					case "Level":
						__gong__assignPointer(&inst.Level, rhs, identifierMap)
					}
				case *Fingering:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Substitution":
						GongUnmarshallEnum(&inst.Substitution, rhs)
					case "Alternate":
						GongUnmarshallEnum(&inst.Alternate, rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						GongUnmarshallEnum(&inst.Placement, rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *First_fret:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Text":
						inst.Text = GongExtractString(rhs)
					case "Location":
						GongUnmarshallEnum(&inst.Location, rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractInt(rhs)
					}
				case *For_part:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Part_clef":
						__gong__assignPointer(&inst.Part_clef, rhs, identifierMap)
					case "Part_transpose":
						__gong__assignPointer(&inst.Part_transpose, rhs, identifierMap)
					}
				case *Formatted_symbol:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Justify":
						GongUnmarshallEnum(&inst.Justify, rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Halign":
						GongUnmarshallEnum(&inst.Halign, rhs)
					case "Valign":
						GongUnmarshallEnum(&inst.Valign, rhs)
					case "Underline":
						inst.Underline = GongExtractInt(rhs)
					case "Overline":
						inst.Overline = GongExtractInt(rhs)
					case "Line_through":
						inst.Line_through = GongExtractInt(rhs)
					case "Rotation":
						inst.Rotation = GongExtractString(rhs)
					case "Letter_spacing":
						inst.Letter_spacing = GongExtractString(rhs)
					case "Line_height":
						inst.Line_height = GongExtractString(rhs)
					case "Dir":
						GongUnmarshallEnum(&inst.Dir, rhs)
					case "Enclosure":
						GongUnmarshallEnum(&inst.Enclosure, rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Formatted_symbol_id:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Justify":
						GongUnmarshallEnum(&inst.Justify, rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Halign":
						GongUnmarshallEnum(&inst.Halign, rhs)
					case "Valign":
						GongUnmarshallEnum(&inst.Valign, rhs)
					case "Underline":
						inst.Underline = GongExtractInt(rhs)
					case "Overline":
						inst.Overline = GongExtractInt(rhs)
					case "Line_through":
						inst.Line_through = GongExtractInt(rhs)
					case "Rotation":
						inst.Rotation = GongExtractString(rhs)
					case "Letter_spacing":
						inst.Letter_spacing = GongExtractString(rhs)
					case "Line_height":
						inst.Line_height = GongExtractString(rhs)
					case "Dir":
						GongUnmarshallEnum(&inst.Dir, rhs)
					case "Enclosure":
						GongUnmarshallEnum(&inst.Enclosure, rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Formatted_text:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Lang":
						inst.Lang = GongExtractString(rhs)
					case "Space":
						inst.Space = GongExtractString(rhs)
					case "Justify":
						GongUnmarshallEnum(&inst.Justify, rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Halign":
						GongUnmarshallEnum(&inst.Halign, rhs)
					case "Valign":
						GongUnmarshallEnum(&inst.Valign, rhs)
					case "Underline":
						inst.Underline = GongExtractInt(rhs)
					case "Overline":
						inst.Overline = GongExtractInt(rhs)
					case "Line_through":
						inst.Line_through = GongExtractInt(rhs)
					case "Rotation":
						inst.Rotation = GongExtractString(rhs)
					case "Letter_spacing":
						inst.Letter_spacing = GongExtractString(rhs)
					case "Line_height":
						inst.Line_height = GongExtractString(rhs)
					case "Dir":
						GongUnmarshallEnum(&inst.Dir, rhs)
					case "Enclosure":
						GongUnmarshallEnum(&inst.Enclosure, rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Formatted_text_id:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Lang":
						inst.Lang = GongExtractString(rhs)
					case "Space":
						inst.Space = GongExtractString(rhs)
					case "Justify":
						GongUnmarshallEnum(&inst.Justify, rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Halign":
						GongUnmarshallEnum(&inst.Halign, rhs)
					case "Valign":
						GongUnmarshallEnum(&inst.Valign, rhs)
					case "Underline":
						inst.Underline = GongExtractInt(rhs)
					case "Overline":
						inst.Overline = GongExtractInt(rhs)
					case "Line_through":
						inst.Line_through = GongExtractInt(rhs)
					case "Rotation":
						inst.Rotation = GongExtractString(rhs)
					case "Letter_spacing":
						inst.Letter_spacing = GongExtractString(rhs)
					case "Line_height":
						inst.Line_height = GongExtractString(rhs)
					case "Dir":
						GongUnmarshallEnum(&inst.Dir, rhs)
					case "Enclosure":
						GongUnmarshallEnum(&inst.Enclosure, rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Forward:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Footnote":
						__gong__assignPointer(&inst.Footnote, rhs, identifierMap)
					case "Level":
						__gong__assignPointer(&inst.Level, rhs, identifierMap)
					case "Voice":
						inst.Voice = GongExtractString(rhs)
					case "Staff":
						inst.Staff = GongExtractInt(rhs)
					}
				case *Frame:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Height":
						inst.Height = GongExtractString(rhs)
					case "Width":
						inst.Width = GongExtractString(rhs)
					case "Unplayed":
						inst.Unplayed = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Halign":
						GongUnmarshallEnum(&inst.Halign, rhs)
					case "Valign":
						GongUnmarshallEnum(&inst.Valign, rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Frame_strings":
						inst.Frame_strings = GongExtractInt(rhs)
					case "Frame_frets":
						inst.Frame_frets = GongExtractInt(rhs)
					case "First_fret":
						__gong__assignPointer(&inst.First_fret, rhs, identifierMap)
					case "Frame_note":
						__gong__assignSliceOfPointers(&inst.Frame_note, rhs, identifierMap)
					}
				case *Frame_note:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "String_XSD":
						__gong__assignPointer(&inst.String_XSD, rhs, identifierMap)
					case "Fret":
						__gong__assignPointer(&inst.Fret, rhs, identifierMap)
					case "Fingering":
						__gong__assignPointer(&inst.Fingering, rhs, identifierMap)
					case "Barre":
						__gong__assignPointer(&inst.Barre, rhs, identifierMap)
					}
				case *Fret:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractInt(rhs)
					}
				case *Glass:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "EnclosedText":
						GongUnmarshallEnum(&inst.EnclosedText, rhs)
					}
				case *Glissando:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type_XSD":
						GongUnmarshallEnum(&inst.Type_XSD, rhs)
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Line_type":
						GongUnmarshallEnum(&inst.Line_type, rhs)
					case "Dash_length":
						inst.Dash_length = GongExtractString(rhs)
					case "Space_length":
						inst.Space_length = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Glyph:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type_XSD":
						inst.Type_XSD = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Grace:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Steal_time_previous":
						inst.Steal_time_previous = GongExtractString(rhs)
					case "Steal_time_following":
						inst.Steal_time_following = GongExtractString(rhs)
					case "Make_time":
						inst.Make_time = GongExtractString(rhs)
					case "Slash":
						GongUnmarshallEnum(&inst.Slash, rhs)
					}
				case *Group_barline:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "EnclosedText":
						GongUnmarshallEnum(&inst.EnclosedText, rhs)
					}
				case *Group_name:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Justify":
						GongUnmarshallEnum(&inst.Justify, rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Group_symbol:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "EnclosedText":
						GongUnmarshallEnum(&inst.EnclosedText, rhs)
					}
				case *Grouping:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type_XSD":
						GongUnmarshallEnum(&inst.Type_XSD, rhs)
					case "Number":
						inst.Number = GongExtractString(rhs)
					case "Member_of":
						inst.Member_of = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Feature":
						__gong__assignSliceOfPointers(&inst.Feature, rhs, identifierMap)
					}
				case *Hammer_on_pull_off:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type_XSD":
						GongUnmarshallEnum(&inst.Type_XSD, rhs)
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						GongUnmarshallEnum(&inst.Placement, rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Handbell:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						GongUnmarshallEnum(&inst.Placement, rhs)
					case "EnclosedText":
						GongUnmarshallEnum(&inst.EnclosedText, rhs)
					}
				case *Harmon_closed:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Location":
						GongUnmarshallEnum(&inst.Location, rhs)
					case "EnclosedText":
						GongUnmarshallEnum(&inst.EnclosedText, rhs)
					}
				case *Harmon_mute:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						GongUnmarshallEnum(&inst.Placement, rhs)
					case "Harmon_closed":
						__gong__assignPointer(&inst.Harmon_closed, rhs, identifierMap)
					}
				case *Harmonic:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Print_object":
						GongUnmarshallEnum(&inst.Print_object, rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						GongUnmarshallEnum(&inst.Placement, rhs)
					case "Natural":
						inst.Natural = GongExtractString(rhs)
					case "Artificial":
						inst.Artificial = GongExtractString(rhs)
					case "Base_pitch":
						inst.Base_pitch = GongExtractString(rhs)
					case "Touching_pitch":
						inst.Touching_pitch = GongExtractString(rhs)
					case "Sounding_pitch":
						inst.Sounding_pitch = GongExtractString(rhs)
					}
				case *Harmony:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type_XSD":
						GongUnmarshallEnum(&inst.Type_XSD, rhs)
					case "Print_frame":
						GongUnmarshallEnum(&inst.Print_frame, rhs)
					case "Arrangement":
						GongUnmarshallEnum(&inst.Arrangement, rhs)
					case "Print_object":
						GongUnmarshallEnum(&inst.Print_object, rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						GongUnmarshallEnum(&inst.Placement, rhs)
					case "System":
						inst.System = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Root":
						__gong__assignPointer(&inst.Root, rhs, identifierMap)
					case "Numeral":
						__gong__assignPointer(&inst.Numeral, rhs, identifierMap)
					case "Function":
						__gong__assignPointer(&inst.Function, rhs, identifierMap)
					case "Kind":
						__gong__assignPointer(&inst.Kind, rhs, identifierMap)
					case "Inversion":
						__gong__assignPointer(&inst.Inversion, rhs, identifierMap)
					case "Bass":
						__gong__assignPointer(&inst.Bass, rhs, identifierMap)
					case "Degree":
						__gong__assignSliceOfPointers(&inst.Degree, rhs, identifierMap)
					case "Frame":
						__gong__assignPointer(&inst.Frame, rhs, identifierMap)
					case "Offset":
						__gong__assignPointer(&inst.Offset, rhs, identifierMap)
					case "Footnote":
						__gong__assignPointer(&inst.Footnote, rhs, identifierMap)
					case "Level":
						__gong__assignPointer(&inst.Level, rhs, identifierMap)
					case "Staff":
						inst.Staff = GongExtractInt(rhs)
					}
				case *Harmony_alter:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Location":
						GongUnmarshallEnum(&inst.Location, rhs)
					case "Print_object":
						GongUnmarshallEnum(&inst.Print_object, rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Harp_pedals:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Halign":
						GongUnmarshallEnum(&inst.Halign, rhs)
					case "Valign":
						GongUnmarshallEnum(&inst.Valign, rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Pedal_tuning":
						__gong__assignSliceOfPointers(&inst.Pedal_tuning, rhs, identifierMap)
					}
				case *Heel_toe:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						GongUnmarshallEnum(&inst.Placement, rhs)
					case "Substitution":
						GongUnmarshallEnum(&inst.Substitution, rhs)
					}
				case *Hole:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						GongUnmarshallEnum(&inst.Placement, rhs)
					case "Hole_type":
						inst.Hole_type = GongExtractString(rhs)
					case "Hole_closed":
						__gong__assignPointer(&inst.Hole_closed, rhs, identifierMap)
					case "Hole_shape":
						inst.Hole_shape = GongExtractString(rhs)
					}
				case *Hole_closed:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Location":
						GongUnmarshallEnum(&inst.Location, rhs)
					case "EnclosedText":
						GongUnmarshallEnum(&inst.EnclosedText, rhs)
					}
				case *Horizontal_turn:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Slash":
						GongUnmarshallEnum(&inst.Slash, rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						GongUnmarshallEnum(&inst.Placement, rhs)
					case "Start_note":
						GongUnmarshallEnum(&inst.Start_note, rhs)
					case "Trill_step":
						GongUnmarshallEnum(&inst.Trill_step, rhs)
					case "Two_note_turn":
						GongUnmarshallEnum(&inst.Two_note_turn, rhs)
					case "Accelerate":
						GongUnmarshallEnum(&inst.Accelerate, rhs)
					case "Beats":
						inst.Beats = GongExtractString(rhs)
					case "Second_beat":
						inst.Second_beat = GongExtractString(rhs)
					case "Last_beat":
						inst.Last_beat = GongExtractString(rhs)
					}
				case *Identification:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Creator":
						__gong__assignSliceOfPointers(&inst.Creator, rhs, identifierMap)
					case "Rights":
						__gong__assignSliceOfPointers(&inst.Rights, rhs, identifierMap)
					case "Encoding":
						__gong__assignPointer(&inst.Encoding, rhs, identifierMap)
					case "Source":
						inst.Source = GongExtractString(rhs)
					case "Relation":
						__gong__assignSliceOfPointers(&inst.Relation, rhs, identifierMap)
					case "Miscellaneous":
						__gong__assignPointer(&inst.Miscellaneous, rhs, identifierMap)
					}
				case *Image:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Source":
						inst.Source = GongExtractString(rhs)
					case "Type_XSD":
						inst.Type_XSD = GongExtractString(rhs)
					case "Height":
						inst.Height = GongExtractString(rhs)
					case "Width":
						inst.Width = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Halign":
						GongUnmarshallEnum(&inst.Halign, rhs)
					case "Valign":
						GongUnmarshallEnum(&inst.Valign, rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					}
				case *Instrument:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					}
				case *Instrument_change:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Instrument_sound":
						inst.Instrument_sound = GongExtractString(rhs)
					case "Solo":
						inst.Solo = GongExtractString(rhs)
					case "Ensemble":
						inst.Ensemble = GongExtractString(rhs)
					case "Virtual_instrument":
						__gong__assignPointer(&inst.Virtual_instrument, rhs, identifierMap)
					}
				case *Instrument_link:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					}
				case *Interchangeable:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Symbol":
						GongUnmarshallEnum(&inst.Symbol, rhs)
					case "Separator":
						GongUnmarshallEnum(&inst.Separator, rhs)
					case "Time_relation":
						GongUnmarshallEnum(&inst.Time_relation, rhs)
					case "Beats":
						inst.Beats = GongExtractString(rhs)
					case "Beat_type":
						inst.Beat_type = GongExtractString(rhs)
					}
				case *Inversion:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Text":
						inst.Text = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractInt(rhs)
					}
				case *Key:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Print_object":
						GongUnmarshallEnum(&inst.Print_object, rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Cancel":
						__gong__assignPointer(&inst.Cancel, rhs, identifierMap)
					case "Fifths":
						inst.Fifths = GongExtractInt(rhs)
					case "Mode":
						inst.Mode = GongExtractString(rhs)
					case "Key_step":
						GongUnmarshallEnum(&inst.Key_step, rhs)
					case "Key_alter":
						inst.Key_alter = GongExtractString(rhs)
					case "Key_accidental":
						__gong__assignPointer(&inst.Key_accidental, rhs, identifierMap)
					case "Key_octave":
						__gong__assignSliceOfPointers(&inst.Key_octave, rhs, identifierMap)
					}
				case *Key_accidental:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "EnclosedText":
						GongUnmarshallEnum(&inst.EnclosedText, rhs)
					}
				case *Key_octave:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Cancel":
						GongUnmarshallEnum(&inst.Cancel, rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractInt(rhs)
					}
				case *Kind:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Use_symbols":
						GongUnmarshallEnum(&inst.Use_symbols, rhs)
					case "Text":
						inst.Text = GongExtractString(rhs)
					case "Stack_degrees":
						GongUnmarshallEnum(&inst.Stack_degrees, rhs)
					case "Parentheses_degrees":
						GongUnmarshallEnum(&inst.Parentheses_degrees, rhs)
					case "Bracket_degrees":
						GongUnmarshallEnum(&inst.Bracket_degrees, rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Halign":
						GongUnmarshallEnum(&inst.Halign, rhs)
					case "Valign":
						GongUnmarshallEnum(&inst.Valign, rhs)
					case "EnclosedText":
						GongUnmarshallEnum(&inst.EnclosedText, rhs)
					}
				case *Level:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Reference_XSD":
						GongUnmarshallEnum(&inst.Reference_XSD, rhs)
					case "Type_XSD":
						GongUnmarshallEnum(&inst.Type_XSD, rhs)
					case "Parentheses":
						GongUnmarshallEnum(&inst.Parentheses, rhs)
					case "Bracket":
						GongUnmarshallEnum(&inst.Bracket, rhs)
					case "Size":
						GongUnmarshallEnum(&inst.Size, rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Line_detail:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Line":
						inst.Line = GongExtractInt(rhs)
					case "Width":
						inst.Width = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Line_type":
						GongUnmarshallEnum(&inst.Line_type, rhs)
					case "Print_object":
						GongUnmarshallEnum(&inst.Print_object, rhs)
					}
				case *Line_width:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type_XSD":
						inst.Type_XSD = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Link:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NameXSD":
						inst.NameXSD = GongExtractString(rhs)
					case "Href":
						inst.Href = GongExtractString(rhs)
					case "Type_XSD":
						inst.Type_XSD = GongExtractString(rhs)
					case "Role":
						inst.Role = GongExtractString(rhs)
					case "Title":
						inst.Title = GongExtractString(rhs)
					case "Show":
						inst.Show = GongExtractString(rhs)
					case "Actuate":
						inst.Actuate = GongExtractString(rhs)
					case "Element":
						inst.Element = GongExtractString(rhs)
					case "Position":
						inst.Position = GongExtractInt(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					}
				case *Listen:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Assess":
						__gong__assignSliceOfPointers(&inst.Assess, rhs, identifierMap)
					case "Wait":
						__gong__assignSliceOfPointers(&inst.Wait, rhs, identifierMap)
					case "Other_listen":
						__gong__assignSliceOfPointers(&inst.Other_listen, rhs, identifierMap)
					}
				case *Listening:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Sync":
						__gong__assignSliceOfPointers(&inst.Sync, rhs, identifierMap)
					case "Other_listening":
						__gong__assignSliceOfPointers(&inst.Other_listening, rhs, identifierMap)
					case "Offset":
						__gong__assignPointer(&inst.Offset, rhs, identifierMap)
					}
				case *Lyric:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractString(rhs)
					case "NameXSD":
						inst.NameXSD = GongExtractString(rhs)
					case "Time_only":
						inst.Time_only = GongExtractString(rhs)
					case "Justify":
						GongUnmarshallEnum(&inst.Justify, rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Placement":
						GongUnmarshallEnum(&inst.Placement, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Print_object":
						GongUnmarshallEnum(&inst.Print_object, rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Syllabic":
						GongUnmarshallEnum(&inst.Syllabic, rhs)
					case "Text":
						__gong__assignPointer(&inst.Text, rhs, identifierMap)
					case "Elision":
						__gong__assignSliceOfPointers(&inst.Elision, rhs, identifierMap)
					case "Extend":
						__gong__assignPointer(&inst.Extend, rhs, identifierMap)
					case "Laughing":
						inst.Laughing = GongExtractString(rhs)
					case "Humming":
						inst.Humming = GongExtractString(rhs)
					case "End_line":
						inst.End_line = GongExtractString(rhs)
					case "End_paragraph":
						inst.End_paragraph = GongExtractString(rhs)
					case "Footnote":
						__gong__assignPointer(&inst.Footnote, rhs, identifierMap)
					case "Level":
						__gong__assignPointer(&inst.Level, rhs, identifierMap)
					}
				case *Lyric_font:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractString(rhs)
					case "NameXSD":
						inst.NameXSD = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					}
				case *Lyric_language:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractString(rhs)
					case "NameXSD":
						inst.NameXSD = GongExtractString(rhs)
					case "Lang":
						inst.Lang = GongExtractString(rhs)
					}
				case *Measure_layout:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Measure_distance":
						inst.Measure_distance = GongExtractString(rhs)
					}
				case *Measure_numbering:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "System":
						GongUnmarshallEnum(&inst.System, rhs)
					case "Staff":
						inst.Staff = GongExtractInt(rhs)
					case "Multiple_rest_always":
						GongUnmarshallEnum(&inst.Multiple_rest_always, rhs)
					case "Multiple_rest_range":
						GongUnmarshallEnum(&inst.Multiple_rest_range, rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Halign":
						GongUnmarshallEnum(&inst.Halign, rhs)
					case "Valign":
						GongUnmarshallEnum(&inst.Valign, rhs)
					case "EnclosedText":
						GongUnmarshallEnum(&inst.EnclosedText, rhs)
					}
				case *Measure_repeat:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type_XSD":
						GongUnmarshallEnum(&inst.Type_XSD, rhs)
					case "Slashes":
						inst.Slashes = GongExtractInt(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Measure_style:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Multiple_rest":
						__gong__assignPointer(&inst.Multiple_rest, rhs, identifierMap)
					case "Measure_repeat":
						__gong__assignPointer(&inst.Measure_repeat, rhs, identifierMap)
					case "Beat_repeat":
						__gong__assignPointer(&inst.Beat_repeat, rhs, identifierMap)
					case "Slash":
						__gong__assignPointer(&inst.Slash, rhs, identifierMap)
					}
				case *Membrane:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "EnclosedText":
						GongUnmarshallEnum(&inst.EnclosedText, rhs)
					}
				case *Metal:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "EnclosedText":
						GongUnmarshallEnum(&inst.EnclosedText, rhs)
					}
				case *Metronome:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Parentheses":
						GongUnmarshallEnum(&inst.Parentheses, rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Halign":
						GongUnmarshallEnum(&inst.Halign, rhs)
					case "Valign":
						GongUnmarshallEnum(&inst.Valign, rhs)
					case "Print_object":
						GongUnmarshallEnum(&inst.Print_object, rhs)
					case "Justify":
						GongUnmarshallEnum(&inst.Justify, rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Beat_unit":
						GongUnmarshallEnum(&inst.Beat_unit, rhs)
					case "Beat_unit_dot":
						inst.Beat_unit_dot = GongExtractString(rhs)
					case "Beat_unit_tied":
						__gong__assignSliceOfPointers(&inst.Beat_unit_tied, rhs, identifierMap)
					case "Per_minute":
						__gong__assignPointer(&inst.Per_minute, rhs, identifierMap)
					case "Metronome_arrows":
						inst.Metronome_arrows = GongExtractString(rhs)
					case "Metronome_note":
						__gong__assignSliceOfPointers(&inst.Metronome_note, rhs, identifierMap)
					case "Metronome_relation":
						inst.Metronome_relation = GongExtractString(rhs)
					}
				case *Metronome_beam:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "EnclosedText":
						GongUnmarshallEnum(&inst.EnclosedText, rhs)
					}
				case *Metronome_note:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Metronome_type":
						GongUnmarshallEnum(&inst.Metronome_type, rhs)
					case "Metronome_dot":
						inst.Metronome_dot = GongExtractString(rhs)
					case "Metronome_beam":
						__gong__assignSliceOfPointers(&inst.Metronome_beam, rhs, identifierMap)
					case "Metronome_tied":
						__gong__assignPointer(&inst.Metronome_tied, rhs, identifierMap)
					case "Metronome_tuplet":
						__gong__assignPointer(&inst.Metronome_tuplet, rhs, identifierMap)
					}
				case *Metronome_tied:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type_XSD":
						GongUnmarshallEnum(&inst.Type_XSD, rhs)
					}
				case *Metronome_tuplet:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Actual_notes":
						inst.Actual_notes = GongExtractInt(rhs)
					case "Normal_notes":
						inst.Normal_notes = GongExtractInt(rhs)
					case "Normal_type":
						GongUnmarshallEnum(&inst.Normal_type, rhs)
					case "Normal_dot":
						inst.Normal_dot = GongExtractString(rhs)
					case "Type_XSD":
						GongUnmarshallEnum(&inst.Type_XSD, rhs)
					case "Bracket":
						GongUnmarshallEnum(&inst.Bracket, rhs)
					case "Show_number":
						GongUnmarshallEnum(&inst.Show_number, rhs)
					}
				case *Midi_device:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Port":
						inst.Port = GongExtractInt(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Midi_instrument:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Midi_channel":
						inst.Midi_channel = GongExtractInt(rhs)
					case "Midi_name":
						inst.Midi_name = GongExtractString(rhs)
					case "Midi_bank":
						inst.Midi_bank = GongExtractInt(rhs)
					case "Midi_program":
						inst.Midi_program = GongExtractInt(rhs)
					case "Midi_unpitched":
						inst.Midi_unpitched = GongExtractInt(rhs)
					case "Volume":
						inst.Volume = GongExtractString(rhs)
					case "Pan":
						inst.Pan = GongExtractString(rhs)
					case "Elevation":
						inst.Elevation = GongExtractString(rhs)
					}
				case *Miscellaneous:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Miscellaneous_field":
						__gong__assignSliceOfPointers(&inst.Miscellaneous_field, rhs, identifierMap)
					}
				case *Miscellaneous_field:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NameXSD":
						inst.NameXSD = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Mordent:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						GongUnmarshallEnum(&inst.Placement, rhs)
					case "Start_note":
						GongUnmarshallEnum(&inst.Start_note, rhs)
					case "Trill_step":
						GongUnmarshallEnum(&inst.Trill_step, rhs)
					case "Two_note_turn":
						GongUnmarshallEnum(&inst.Two_note_turn, rhs)
					case "Accelerate":
						GongUnmarshallEnum(&inst.Accelerate, rhs)
					case "Beats":
						inst.Beats = GongExtractString(rhs)
					case "Second_beat":
						inst.Second_beat = GongExtractString(rhs)
					case "Last_beat":
						inst.Last_beat = GongExtractString(rhs)
					case "Long":
						GongUnmarshallEnum(&inst.Long, rhs)
					case "Approach":
						GongUnmarshallEnum(&inst.Approach, rhs)
					case "Departure":
						GongUnmarshallEnum(&inst.Departure, rhs)
					}
				case *Multiple_rest:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Use_symbols":
						GongUnmarshallEnum(&inst.Use_symbols, rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractInt(rhs)
					}
				case *Name_display:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Print_object":
						GongUnmarshallEnum(&inst.Print_object, rhs)
					case "Display_text":
						__gong__assignSliceOfPointers(&inst.Display_text, rhs, identifierMap)
					case "Accidental_text":
						__gong__assignSliceOfPointers(&inst.Accidental_text, rhs, identifierMap)
					}
				case *Non_arpeggiate:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type_XSD":
						GongUnmarshallEnum(&inst.Type_XSD, rhs)
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Placement":
						GongUnmarshallEnum(&inst.Placement, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					}
				case *Notations:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Print_object":
						GongUnmarshallEnum(&inst.Print_object, rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Footnote":
						__gong__assignPointer(&inst.Footnote, rhs, identifierMap)
					case "Level":
						__gong__assignPointer(&inst.Level, rhs, identifierMap)
					case "Tied":
						__gong__assignSliceOfPointers(&inst.Tied, rhs, identifierMap)
					case "Slur":
						__gong__assignSliceOfPointers(&inst.Slur, rhs, identifierMap)
					case "Tuplet":
						__gong__assignSliceOfPointers(&inst.Tuplet, rhs, identifierMap)
					case "Glissando":
						__gong__assignSliceOfPointers(&inst.Glissando, rhs, identifierMap)
					case "Slide":
						__gong__assignSliceOfPointers(&inst.Slide, rhs, identifierMap)
					case "Ornaments":
						__gong__assignSliceOfPointers(&inst.Ornaments, rhs, identifierMap)
					case "Technical":
						__gong__assignSliceOfPointers(&inst.Technical, rhs, identifierMap)
					case "Articulations":
						__gong__assignSliceOfPointers(&inst.Articulations, rhs, identifierMap)
					case "Dynamics":
						__gong__assignSliceOfPointers(&inst.Dynamics, rhs, identifierMap)
					case "Fermata":
						__gong__assignSliceOfPointers(&inst.Fermata, rhs, identifierMap)
					case "Arpeggiate":
						__gong__assignSliceOfPointers(&inst.Arpeggiate, rhs, identifierMap)
					case "Non_arpeggiate":
						__gong__assignSliceOfPointers(&inst.Non_arpeggiate, rhs, identifierMap)
					case "Accidental_mark":
						__gong__assignSliceOfPointers(&inst.Accidental_mark, rhs, identifierMap)
					case "Other_notation":
						__gong__assignSliceOfPointers(&inst.Other_notation, rhs, identifierMap)
					}
				case *Note:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Print_leger":
						GongUnmarshallEnum(&inst.Print_leger, rhs)
					case "Dynamics":
						inst.Dynamics = GongExtractString(rhs)
					case "End_dynamics":
						inst.End_dynamics = GongExtractString(rhs)
					case "Attack":
						inst.Attack = GongExtractString(rhs)
					case "Release":
						inst.Release = GongExtractString(rhs)
					case "Time_only":
						inst.Time_only = GongExtractString(rhs)
					case "Pizzicato":
						GongUnmarshallEnum(&inst.Pizzicato, rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Print_dot":
						GongUnmarshallEnum(&inst.Print_dot, rhs)
					case "Print_lyric":
						GongUnmarshallEnum(&inst.Print_lyric, rhs)
					case "Print_object":
						GongUnmarshallEnum(&inst.Print_object, rhs)
					case "Print_spacing":
						GongUnmarshallEnum(&inst.Print_spacing, rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Grace":
						__gong__assignPointer(&inst.Grace, rhs, identifierMap)
					case "Chord":
						inst.Chord = GongExtractString(rhs)
					case "Pitch":
						__gong__assignPointer(&inst.Pitch, rhs, identifierMap)
					case "Unpitched":
						__gong__assignPointer(&inst.Unpitched, rhs, identifierMap)
					case "Rest":
						__gong__assignPointer(&inst.Rest, rhs, identifierMap)
					case "Tie":
						__gong__assignPointer(&inst.Tie, rhs, identifierMap)
					case "Cue":
						inst.Cue = GongExtractString(rhs)
					case "Instrument":
						__gong__assignSliceOfPointers(&inst.Instrument, rhs, identifierMap)
					case "Footnote":
						__gong__assignPointer(&inst.Footnote, rhs, identifierMap)
					case "Level":
						__gong__assignPointer(&inst.Level, rhs, identifierMap)
					case "Voice":
						inst.Voice = GongExtractString(rhs)
					case "Type_XSD":
						__gong__assignPointer(&inst.Type_XSD, rhs, identifierMap)
					case "Dot":
						__gong__assignSliceOfPointers(&inst.Dot, rhs, identifierMap)
					case "Accidental":
						__gong__assignPointer(&inst.Accidental, rhs, identifierMap)
					case "Time_modification":
						__gong__assignPointer(&inst.Time_modification, rhs, identifierMap)
					case "Stem":
						__gong__assignPointer(&inst.Stem, rhs, identifierMap)
					case "Notehead":
						__gong__assignPointer(&inst.Notehead, rhs, identifierMap)
					case "Notehead_text":
						__gong__assignPointer(&inst.Notehead_text, rhs, identifierMap)
					case "Staff":
						inst.Staff = GongExtractInt(rhs)
					case "Beam":
						__gong__assignPointer(&inst.Beam, rhs, identifierMap)
					case "Notations":
						__gong__assignSliceOfPointers(&inst.Notations, rhs, identifierMap)
					case "Lyric":
						__gong__assignSliceOfPointers(&inst.Lyric, rhs, identifierMap)
					case "Play":
						__gong__assignPointer(&inst.Play, rhs, identifierMap)
					case "Listen":
						__gong__assignPointer(&inst.Listen, rhs, identifierMap)
					}
				case *Note_size:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type_XSD":
						GongUnmarshallEnum(&inst.Type_XSD, rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Note_type:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Size":
						GongUnmarshallEnum(&inst.Size, rhs)
					case "EnclosedText":
						GongUnmarshallEnum(&inst.EnclosedText, rhs)
					}
				case *Notehead:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Filled":
						GongUnmarshallEnum(&inst.Filled, rhs)
					case "Parentheses":
						GongUnmarshallEnum(&inst.Parentheses, rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "EnclosedText":
						GongUnmarshallEnum(&inst.EnclosedText, rhs)
					}
				case *Notehead_text:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Display_text":
						__gong__assignSliceOfPointers(&inst.Display_text, rhs, identifierMap)
					case "Accidental_text":
						__gong__assignSliceOfPointers(&inst.Accidental_text, rhs, identifierMap)
					}
				case *Numeral:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Numeral_root":
						__gong__assignPointer(&inst.Numeral_root, rhs, identifierMap)
					case "Numeral_alter":
						__gong__assignPointer(&inst.Numeral_alter, rhs, identifierMap)
					case "Numeral_key":
						__gong__assignPointer(&inst.Numeral_key, rhs, identifierMap)
					}
				case *Numeral_key:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Print_object":
						GongUnmarshallEnum(&inst.Print_object, rhs)
					case "Numeral_fifths":
						inst.Numeral_fifths = GongExtractInt(rhs)
					case "Numeral_mode":
						GongUnmarshallEnum(&inst.Numeral_mode, rhs)
					}
				case *Numeral_root:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Text":
						inst.Text = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractInt(rhs)
					}
				case *Octave_shift:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type_XSD":
						GongUnmarshallEnum(&inst.Type_XSD, rhs)
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Size":
						inst.Size = GongExtractInt(rhs)
					case "Dash_length":
						inst.Dash_length = GongExtractString(rhs)
					case "Space_length":
						inst.Space_length = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					}
				case *Offset:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Sound":
						GongUnmarshallEnum(&inst.Sound, rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Opus:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Href":
						inst.Href = GongExtractString(rhs)
					case "Type_XSD":
						inst.Type_XSD = GongExtractString(rhs)
					case "Role":
						inst.Role = GongExtractString(rhs)
					case "Title":
						inst.Title = GongExtractString(rhs)
					case "Show":
						inst.Show = GongExtractString(rhs)
					case "Actuate":
						inst.Actuate = GongExtractString(rhs)
					}
				case *Ornaments:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Trill_mark":
						__gong__assignSliceOfPointers(&inst.Trill_mark, rhs, identifierMap)
					case "Turn":
						__gong__assignSliceOfPointers(&inst.Turn, rhs, identifierMap)
					case "Delayed_turn":
						__gong__assignSliceOfPointers(&inst.Delayed_turn, rhs, identifierMap)
					case "Inverted_turn":
						__gong__assignSliceOfPointers(&inst.Inverted_turn, rhs, identifierMap)
					case "Delayed_inverted_turn":
						__gong__assignSliceOfPointers(&inst.Delayed_inverted_turn, rhs, identifierMap)
					case "Vertical_turn":
						__gong__assignSliceOfPointers(&inst.Vertical_turn, rhs, identifierMap)
					case "Inverted_vertical_turn":
						__gong__assignSliceOfPointers(&inst.Inverted_vertical_turn, rhs, identifierMap)
					case "Shake":
						__gong__assignSliceOfPointers(&inst.Shake, rhs, identifierMap)
					case "Wavy_line":
						__gong__assignSliceOfPointers(&inst.Wavy_line, rhs, identifierMap)
					case "Mordent":
						__gong__assignSliceOfPointers(&inst.Mordent, rhs, identifierMap)
					case "Inverted_mordent":
						__gong__assignSliceOfPointers(&inst.Inverted_mordent, rhs, identifierMap)
					case "Schleifer":
						__gong__assignSliceOfPointers(&inst.Schleifer, rhs, identifierMap)
					case "Tremolo":
						__gong__assignSliceOfPointers(&inst.Tremolo, rhs, identifierMap)
					case "Haydn":
						__gong__assignSliceOfPointers(&inst.Haydn, rhs, identifierMap)
					case "Other_ornament":
						__gong__assignSliceOfPointers(&inst.Other_ornament, rhs, identifierMap)
					case "Accidental_mark":
						__gong__assignSliceOfPointers(&inst.Accidental_mark, rhs, identifierMap)
					}
				case *Other_appearance:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type_XSD":
						inst.Type_XSD = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Other_direction:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Print_object":
						GongUnmarshallEnum(&inst.Print_object, rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Halign":
						GongUnmarshallEnum(&inst.Halign, rhs)
					case "Valign":
						GongUnmarshallEnum(&inst.Valign, rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Other_listening:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type_XSD":
						inst.Type_XSD = GongExtractString(rhs)
					case "Player":
						inst.Player = GongExtractString(rhs)
					case "Time_only":
						inst.Time_only = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Other_notation:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type_XSD":
						GongUnmarshallEnum(&inst.Type_XSD, rhs)
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Print_object":
						GongUnmarshallEnum(&inst.Print_object, rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						GongUnmarshallEnum(&inst.Placement, rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Other_placement_text:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						GongUnmarshallEnum(&inst.Placement, rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Other_play:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type_XSD":
						inst.Type_XSD = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Other_text:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Page_layout:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Page_height":
						inst.Page_height = GongExtractString(rhs)
					case "Page_width":
						inst.Page_width = GongExtractString(rhs)
					case "Page_margins":
						__gong__assignPointer(&inst.Page_margins, rhs, identifierMap)
					}
				case *Page_margins:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type_XSD":
						GongUnmarshallEnum(&inst.Type_XSD, rhs)
					case "Left_margin":
						inst.Left_margin = GongExtractString(rhs)
					case "Right_margin":
						inst.Right_margin = GongExtractString(rhs)
					case "Top_margin":
						inst.Top_margin = GongExtractString(rhs)
					case "Bottom_margin":
						inst.Bottom_margin = GongExtractString(rhs)
					}
				case *Part_clef:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Sign":
						GongUnmarshallEnum(&inst.Sign, rhs)
					case "Line":
						inst.Line = GongExtractInt(rhs)
					case "Clef_octave_change":
						inst.Clef_octave_change = GongExtractInt(rhs)
					}
				case *Part_group:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type_XSD":
						GongUnmarshallEnum(&inst.Type_XSD, rhs)
					case "Number":
						inst.Number = GongExtractString(rhs)
					case "Group_name":
						__gong__assignPointer(&inst.Group_name, rhs, identifierMap)
					case "Group_name_display":
						__gong__assignPointer(&inst.Group_name_display, rhs, identifierMap)
					case "Group_abbreviation":
						__gong__assignPointer(&inst.Group_abbreviation, rhs, identifierMap)
					case "Group_abbreviation_display":
						__gong__assignPointer(&inst.Group_abbreviation_display, rhs, identifierMap)
					case "Group_symbol":
						__gong__assignPointer(&inst.Group_symbol, rhs, identifierMap)
					case "Group_barline":
						__gong__assignPointer(&inst.Group_barline, rhs, identifierMap)
					case "Group_time":
						inst.Group_time = GongExtractString(rhs)
					case "Footnote":
						__gong__assignPointer(&inst.Footnote, rhs, identifierMap)
					case "Level":
						__gong__assignPointer(&inst.Level, rhs, identifierMap)
					}
				case *Part_link:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Href":
						inst.Href = GongExtractString(rhs)
					case "Type_XSD":
						inst.Type_XSD = GongExtractString(rhs)
					case "Role":
						inst.Role = GongExtractString(rhs)
					case "Title":
						inst.Title = GongExtractString(rhs)
					case "Show":
						inst.Show = GongExtractString(rhs)
					case "Actuate":
						inst.Actuate = GongExtractString(rhs)
					case "Instrument_link":
						__gong__assignSliceOfPointers(&inst.Instrument_link, rhs, identifierMap)
					case "Group_link":
						inst.Group_link = GongExtractString(rhs)
					}
				case *Part_list:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Part_group":
						__gong__assignPointer(&inst.Part_group, rhs, identifierMap)
					case "Score_part":
						__gong__assignPointer(&inst.Score_part, rhs, identifierMap)
					}
				case *Part_name:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Print_object":
						GongUnmarshallEnum(&inst.Print_object, rhs)
					case "Justify":
						GongUnmarshallEnum(&inst.Justify, rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Part_symbol:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Top_staff":
						inst.Top_staff = GongExtractInt(rhs)
					case "Bottom_staff":
						inst.Bottom_staff = GongExtractInt(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "EnclosedText":
						GongUnmarshallEnum(&inst.EnclosedText, rhs)
					}
				case *Part_transpose:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Diatonic":
						inst.Diatonic = GongExtractInt(rhs)
					case "Chromatic":
						inst.Chromatic = GongExtractString(rhs)
					case "Octave_change":
						inst.Octave_change = GongExtractInt(rhs)
					case "Double":
						inst.Double = GongExtractFloat(rhs)
					}
				case *Pedal:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type_XSD":
						GongUnmarshallEnum(&inst.Type_XSD, rhs)
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Line":
						GongUnmarshallEnum(&inst.Line, rhs)
					case "Sign":
						GongUnmarshallEnum(&inst.Sign, rhs)
					case "Abbreviated":
						GongUnmarshallEnum(&inst.Abbreviated, rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Halign":
						GongUnmarshallEnum(&inst.Halign, rhs)
					case "Valign":
						GongUnmarshallEnum(&inst.Valign, rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					}
				case *Pedal_tuning:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Pedal_step":
						GongUnmarshallEnum(&inst.Pedal_step, rhs)
					case "Pedal_alter":
						inst.Pedal_alter = GongExtractString(rhs)
					}
				case *Per_minute:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Percussion:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Halign":
						GongUnmarshallEnum(&inst.Halign, rhs)
					case "Valign":
						GongUnmarshallEnum(&inst.Valign, rhs)
					case "Enclosure":
						GongUnmarshallEnum(&inst.Enclosure, rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Glass":
						__gong__assignPointer(&inst.Glass, rhs, identifierMap)
					case "Metal":
						__gong__assignPointer(&inst.Metal, rhs, identifierMap)
					case "Wood":
						__gong__assignPointer(&inst.Wood, rhs, identifierMap)
					case "Pitched":
						__gong__assignPointer(&inst.Pitched, rhs, identifierMap)
					case "Membrane":
						__gong__assignPointer(&inst.Membrane, rhs, identifierMap)
					case "Effect":
						__gong__assignPointer(&inst.Effect, rhs, identifierMap)
					case "Timpani":
						__gong__assignPointer(&inst.Timpani, rhs, identifierMap)
					case "Beater":
						__gong__assignPointer(&inst.Beater, rhs, identifierMap)
					case "Stick":
						__gong__assignPointer(&inst.Stick, rhs, identifierMap)
					case "Stick_location":
						GongUnmarshallEnum(&inst.Stick_location, rhs)
					case "Other_percussion":
						__gong__assignPointer(&inst.Other_percussion, rhs, identifierMap)
					}
				case *Pitch:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Step":
						GongUnmarshallEnum(&inst.Step, rhs)
					case "Alter":
						inst.Alter = GongExtractString(rhs)
					case "Octave":
						inst.Octave = GongExtractInt(rhs)
					}
				case *Pitched:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "EnclosedText":
						GongUnmarshallEnum(&inst.EnclosedText, rhs)
					}
				case *Placement_text:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						GongUnmarshallEnum(&inst.Placement, rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Play:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Ipa":
						inst.Ipa = GongExtractString(rhs)
					case "Mute":
						GongUnmarshallEnum(&inst.Mute, rhs)
					case "Semi_pitched":
						GongUnmarshallEnum(&inst.Semi_pitched, rhs)
					case "Other_play":
						__gong__assignSliceOfPointers(&inst.Other_play, rhs, identifierMap)
					}
				case *Player:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Player_name":
						inst.Player_name = GongExtractString(rhs)
					}
				case *Principal_voice:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type_XSD":
						GongUnmarshallEnum(&inst.Type_XSD, rhs)
					case "Symbol":
						GongUnmarshallEnum(&inst.Symbol, rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Halign":
						GongUnmarshallEnum(&inst.Halign, rhs)
					case "Valign":
						GongUnmarshallEnum(&inst.Valign, rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Print_XSD:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Staff_spacing":
						inst.Staff_spacing = GongExtractString(rhs)
					case "New_system":
						GongUnmarshallEnum(&inst.New_system, rhs)
					case "New_page":
						GongUnmarshallEnum(&inst.New_page, rhs)
					case "Blank_page":
						inst.Blank_page = GongExtractInt(rhs)
					case "Page_number":
						inst.Page_number = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Page_layout":
						__gong__assignPointer(&inst.Page_layout, rhs, identifierMap)
					case "System_layout":
						__gong__assignPointer(&inst.System_layout, rhs, identifierMap)
					case "Staff_layout":
						__gong__assignSliceOfPointers(&inst.Staff_layout, rhs, identifierMap)
					case "Measure_layout":
						__gong__assignPointer(&inst.Measure_layout, rhs, identifierMap)
					case "Measure_numbering":
						__gong__assignPointer(&inst.Measure_numbering, rhs, identifierMap)
					case "Part_name_display":
						__gong__assignPointer(&inst.Part_name_display, rhs, identifierMap)
					case "Part_abbreviation_display":
						__gong__assignPointer(&inst.Part_abbreviation_display, rhs, identifierMap)
					}
				case *Release:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Offset":
						inst.Offset = GongExtractString(rhs)
					}
				case *Repeat:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Direction":
						GongUnmarshallEnum(&inst.Direction, rhs)
					case "Times":
						inst.Times = GongExtractInt(rhs)
					case "After_jump":
						GongUnmarshallEnum(&inst.After_jump, rhs)
					case "Winged":
						GongUnmarshallEnum(&inst.Winged, rhs)
					}
				case *Rest:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Measure":
						GongUnmarshallEnum(&inst.Measure, rhs)
					case "Display_step":
						GongUnmarshallEnum(&inst.Display_step, rhs)
					case "Display_octave":
						inst.Display_octave = GongExtractInt(rhs)
					}
				case *Root:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Root_step":
						__gong__assignPointer(&inst.Root_step, rhs, identifierMap)
					case "Root_alter":
						__gong__assignPointer(&inst.Root_alter, rhs, identifierMap)
					}
				case *Root_step:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Text":
						inst.Text = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "EnclosedText":
						GongUnmarshallEnum(&inst.EnclosedText, rhs)
					}
				case *Scaling:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Millimeters":
						inst.Millimeters = GongExtractString(rhs)
					case "Tenths":
						inst.Tenths = GongExtractString(rhs)
					}
				case *Scordatura:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Accord":
						__gong__assignSliceOfPointers(&inst.Accord, rhs, identifierMap)
					}
				case *Score_instrument:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Instrument_name":
						inst.Instrument_name = GongExtractString(rhs)
					case "Instrument_abbreviation":
						inst.Instrument_abbreviation = GongExtractString(rhs)
					case "Instrument_sound":
						inst.Instrument_sound = GongExtractString(rhs)
					case "Solo":
						inst.Solo = GongExtractString(rhs)
					case "Ensemble":
						inst.Ensemble = GongExtractString(rhs)
					case "Virtual_instrument":
						__gong__assignPointer(&inst.Virtual_instrument, rhs, identifierMap)
					}
				case *Score_part:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Identification":
						__gong__assignPointer(&inst.Identification, rhs, identifierMap)
					case "Part_link":
						__gong__assignSliceOfPointers(&inst.Part_link, rhs, identifierMap)
					case "Part_name":
						__gong__assignPointer(&inst.Part_name, rhs, identifierMap)
					case "Part_name_display":
						__gong__assignPointer(&inst.Part_name_display, rhs, identifierMap)
					case "Part_abbreviation":
						__gong__assignPointer(&inst.Part_abbreviation, rhs, identifierMap)
					case "Part_abbreviation_display":
						__gong__assignPointer(&inst.Part_abbreviation_display, rhs, identifierMap)
					case "Group":
						inst.Group = GongExtractString(rhs)
					case "Score_instrument":
						__gong__assignSliceOfPointers(&inst.Score_instrument, rhs, identifierMap)
					case "Player":
						__gong__assignSliceOfPointers(&inst.Player, rhs, identifierMap)
					case "Midi_device":
						__gong__assignSliceOfPointers(&inst.Midi_device, rhs, identifierMap)
					case "Midi_instrument":
						__gong__assignSliceOfPointers(&inst.Midi_instrument, rhs, identifierMap)
					}
				case *Score_partwise:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Version":
						inst.Version = GongExtractString(rhs)
					case "Work":
						__gong__assignPointer(&inst.Work, rhs, identifierMap)
					case "Movement_number":
						inst.Movement_number = GongExtractString(rhs)
					case "Movement_title":
						inst.Movement_title = GongExtractString(rhs)
					case "Identification":
						__gong__assignPointer(&inst.Identification, rhs, identifierMap)
					case "Defaults":
						__gong__assignPointer(&inst.Defaults, rhs, identifierMap)
					case "Credit":
						__gong__assignSliceOfPointers(&inst.Credit, rhs, identifierMap)
					case "Part_list":
						__gong__assignPointer(&inst.Part_list, rhs, identifierMap)
					case "Part":
						__gong__assignSliceOfPointers(&inst.Part, rhs, identifierMap)
					}
				case *Score_timewise:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Version":
						inst.Version = GongExtractString(rhs)
					case "Work":
						__gong__assignPointer(&inst.Work, rhs, identifierMap)
					case "Movement_number":
						inst.Movement_number = GongExtractString(rhs)
					case "Movement_title":
						inst.Movement_title = GongExtractString(rhs)
					case "Identification":
						__gong__assignPointer(&inst.Identification, rhs, identifierMap)
					case "Defaults":
						__gong__assignPointer(&inst.Defaults, rhs, identifierMap)
					case "Credit":
						__gong__assignSliceOfPointers(&inst.Credit, rhs, identifierMap)
					case "Part_list":
						__gong__assignPointer(&inst.Part_list, rhs, identifierMap)
					case "Measure":
						__gong__assignSliceOfPointers(&inst.Measure, rhs, identifierMap)
					}
				case *Segno:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Halign":
						GongUnmarshallEnum(&inst.Halign, rhs)
					case "Valign":
						GongUnmarshallEnum(&inst.Valign, rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					}
				case *Slash:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type_XSD":
						GongUnmarshallEnum(&inst.Type_XSD, rhs)
					case "Use_dots":
						GongUnmarshallEnum(&inst.Use_dots, rhs)
					case "Use_stems":
						GongUnmarshallEnum(&inst.Use_stems, rhs)
					case "Slash_type":
						GongUnmarshallEnum(&inst.Slash_type, rhs)
					case "Slash_dot":
						inst.Slash_dot = GongExtractString(rhs)
					case "Except_voice":
						inst.Except_voice = GongExtractString(rhs)
					}
				case *Slide:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type_XSD":
						GongUnmarshallEnum(&inst.Type_XSD, rhs)
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Line_type":
						GongUnmarshallEnum(&inst.Line_type, rhs)
					case "Dash_length":
						inst.Dash_length = GongExtractString(rhs)
					case "Space_length":
						inst.Space_length = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Accelerate":
						GongUnmarshallEnum(&inst.Accelerate, rhs)
					case "Beats":
						inst.Beats = GongExtractString(rhs)
					case "First_beat":
						inst.First_beat = GongExtractString(rhs)
					case "Last_beat":
						inst.Last_beat = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Slur:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type_XSD":
						GongUnmarshallEnum(&inst.Type_XSD, rhs)
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Line_type":
						GongUnmarshallEnum(&inst.Line_type, rhs)
					case "Dash_length":
						inst.Dash_length = GongExtractString(rhs)
					case "Space_length":
						inst.Space_length = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Placement":
						GongUnmarshallEnum(&inst.Placement, rhs)
					case "Orientation":
						GongUnmarshallEnum(&inst.Orientation, rhs)
					case "Bezier_x":
						inst.Bezier_x = GongExtractString(rhs)
					case "Bezier_y":
						inst.Bezier_y = GongExtractString(rhs)
					case "Bezier_x2":
						inst.Bezier_x2 = GongExtractString(rhs)
					case "Bezier_y2":
						inst.Bezier_y2 = GongExtractString(rhs)
					case "Bezier_offset":
						inst.Bezier_offset = GongExtractString(rhs)
					case "Bezier_offset2":
						inst.Bezier_offset2 = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					}
				case *Sound:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Tempo":
						inst.Tempo = GongExtractString(rhs)
					case "Dynamics":
						inst.Dynamics = GongExtractString(rhs)
					case "Dacapo":
						GongUnmarshallEnum(&inst.Dacapo, rhs)
					case "Segno":
						inst.Segno = GongExtractString(rhs)
					case "Dalsegno":
						inst.Dalsegno = GongExtractString(rhs)
					case "Coda":
						inst.Coda = GongExtractString(rhs)
					case "Tocoda":
						inst.Tocoda = GongExtractString(rhs)
					case "Divisions":
						inst.Divisions = GongExtractString(rhs)
					case "Forward_repeat":
						GongUnmarshallEnum(&inst.Forward_repeat, rhs)
					case "Fine":
						inst.Fine = GongExtractString(rhs)
					case "Time_only":
						inst.Time_only = GongExtractString(rhs)
					case "Pizzicato":
						GongUnmarshallEnum(&inst.Pizzicato, rhs)
					case "Pan":
						inst.Pan = GongExtractString(rhs)
					case "Elevation":
						inst.Elevation = GongExtractString(rhs)
					case "Damper_pedal":
						inst.Damper_pedal = GongExtractString(rhs)
					case "Soft_pedal":
						inst.Soft_pedal = GongExtractString(rhs)
					case "Sostenuto_pedal":
						inst.Sostenuto_pedal = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Instrument_change":
						__gong__assignSliceOfPointers(&inst.Instrument_change, rhs, identifierMap)
					case "Midi_device":
						__gong__assignSliceOfPointers(&inst.Midi_device, rhs, identifierMap)
					case "Midi_instrument":
						__gong__assignSliceOfPointers(&inst.Midi_instrument, rhs, identifierMap)
					case "Play":
						__gong__assignSliceOfPointers(&inst.Play, rhs, identifierMap)
					case "Swing":
						__gong__assignPointer(&inst.Swing, rhs, identifierMap)
					case "Offset":
						__gong__assignPointer(&inst.Offset, rhs, identifierMap)
					}
				case *Staff_details:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Show_frets":
						GongUnmarshallEnum(&inst.Show_frets, rhs)
					case "Print_object":
						GongUnmarshallEnum(&inst.Print_object, rhs)
					case "Print_spacing":
						GongUnmarshallEnum(&inst.Print_spacing, rhs)
					case "Staff_type":
						GongUnmarshallEnum(&inst.Staff_type, rhs)
					case "Staff_lines":
						inst.Staff_lines = GongExtractInt(rhs)
					case "Line_detail":
						__gong__assignSliceOfPointers(&inst.Line_detail, rhs, identifierMap)
					case "Staff_tuning":
						__gong__assignSliceOfPointers(&inst.Staff_tuning, rhs, identifierMap)
					case "Capo":
						inst.Capo = GongExtractInt(rhs)
					case "Staff_size":
						__gong__assignPointer(&inst.Staff_size, rhs, identifierMap)
					}
				case *Staff_divide:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type_XSD":
						GongUnmarshallEnum(&inst.Type_XSD, rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Halign":
						GongUnmarshallEnum(&inst.Halign, rhs)
					case "Valign":
						GongUnmarshallEnum(&inst.Valign, rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					}
				case *Staff_layout:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Staff_distance":
						inst.Staff_distance = GongExtractString(rhs)
					}
				case *Staff_size:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Scaling":
						inst.Scaling = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Staff_tuning:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Line":
						inst.Line = GongExtractInt(rhs)
					case "Tuning_step":
						GongUnmarshallEnum(&inst.Tuning_step, rhs)
					case "Tuning_alter":
						inst.Tuning_alter = GongExtractString(rhs)
					case "Tuning_octave":
						inst.Tuning_octave = GongExtractInt(rhs)
					}
				case *Stem:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "EnclosedText":
						GongUnmarshallEnum(&inst.EnclosedText, rhs)
					}
				case *Stick:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Tip":
						GongUnmarshallEnum(&inst.Tip, rhs)
					case "Parentheses":
						GongUnmarshallEnum(&inst.Parentheses, rhs)
					case "Dashed_circle":
						GongUnmarshallEnum(&inst.Dashed_circle, rhs)
					case "Stick_type":
						GongUnmarshallEnum(&inst.Stick_type, rhs)
					case "Stick_material":
						GongUnmarshallEnum(&inst.Stick_material, rhs)
					}
				case *String_mute:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type_XSD":
						GongUnmarshallEnum(&inst.Type_XSD, rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Halign":
						GongUnmarshallEnum(&inst.Halign, rhs)
					case "Valign":
						GongUnmarshallEnum(&inst.Valign, rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					}
				case *String_type:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						GongUnmarshallEnum(&inst.Placement, rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractInt(rhs)
					}
				case *Strong_accent:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						GongUnmarshallEnum(&inst.Placement, rhs)
					case "Type_XSD":
						GongUnmarshallEnum(&inst.Type_XSD, rhs)
					}
				case *Style_text:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Supports:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type_XSD":
						GongUnmarshallEnum(&inst.Type_XSD, rhs)
					case "Element":
						inst.Element = GongExtractString(rhs)
					case "Attribute":
						inst.Attribute = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *Swing:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Straight":
						inst.Straight = GongExtractString(rhs)
					case "First":
						inst.First = GongExtractInt(rhs)
					case "Second":
						inst.Second = GongExtractInt(rhs)
					case "Swing_type":
						inst.Swing_type = GongExtractString(rhs)
					case "Swing_style":
						inst.Swing_style = GongExtractString(rhs)
					}
				case *Sync:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type_XSD":
						GongUnmarshallEnum(&inst.Type_XSD, rhs)
					case "Latency":
						inst.Latency = GongExtractInt(rhs)
					case "Player":
						inst.Player = GongExtractString(rhs)
					case "Time_only":
						inst.Time_only = GongExtractString(rhs)
					}
				case *System_dividers:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Left_divider":
						__gong__assignPointer(&inst.Left_divider, rhs, identifierMap)
					case "Right_divider":
						__gong__assignPointer(&inst.Right_divider, rhs, identifierMap)
					}
				case *System_layout:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "System_margins":
						__gong__assignPointer(&inst.System_margins, rhs, identifierMap)
					case "System_distance":
						inst.System_distance = GongExtractString(rhs)
					case "Top_system_distance":
						inst.Top_system_distance = GongExtractString(rhs)
					case "System_dividers":
						__gong__assignPointer(&inst.System_dividers, rhs, identifierMap)
					}
				case *System_margins:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Left_margin":
						inst.Left_margin = GongExtractString(rhs)
					case "Right_margin":
						inst.Right_margin = GongExtractString(rhs)
					}
				case *Tap:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Hand":
						GongUnmarshallEnum(&inst.Hand, rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						GongUnmarshallEnum(&inst.Placement, rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Technical:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Up_bow":
						__gong__assignSliceOfPointers(&inst.Up_bow, rhs, identifierMap)
					case "Down_bow":
						__gong__assignSliceOfPointers(&inst.Down_bow, rhs, identifierMap)
					case "Harmonic":
						__gong__assignSliceOfPointers(&inst.Harmonic, rhs, identifierMap)
					case "Open_string":
						__gong__assignSliceOfPointers(&inst.Open_string, rhs, identifierMap)
					case "Thumb_position":
						__gong__assignSliceOfPointers(&inst.Thumb_position, rhs, identifierMap)
					case "Fingering":
						__gong__assignSliceOfPointers(&inst.Fingering, rhs, identifierMap)
					case "Pluck":
						__gong__assignSliceOfPointers(&inst.Pluck, rhs, identifierMap)
					case "Double_tongue":
						__gong__assignSliceOfPointers(&inst.Double_tongue, rhs, identifierMap)
					case "Triple_tongue":
						__gong__assignSliceOfPointers(&inst.Triple_tongue, rhs, identifierMap)
					case "Stopped":
						__gong__assignSliceOfPointers(&inst.Stopped, rhs, identifierMap)
					case "Snap_pizzicato":
						__gong__assignSliceOfPointers(&inst.Snap_pizzicato, rhs, identifierMap)
					case "Fret":
						__gong__assignSliceOfPointers(&inst.Fret, rhs, identifierMap)
					case "String_XSD":
						__gong__assignSliceOfPointers(&inst.String_XSD, rhs, identifierMap)
					case "Hammer_on":
						__gong__assignSliceOfPointers(&inst.Hammer_on, rhs, identifierMap)
					case "Pull_off":
						__gong__assignSliceOfPointers(&inst.Pull_off, rhs, identifierMap)
					case "Bend":
						__gong__assignSliceOfPointers(&inst.Bend, rhs, identifierMap)
					case "Tap":
						__gong__assignSliceOfPointers(&inst.Tap, rhs, identifierMap)
					case "Heel":
						__gong__assignSliceOfPointers(&inst.Heel, rhs, identifierMap)
					case "Toe":
						__gong__assignSliceOfPointers(&inst.Toe, rhs, identifierMap)
					case "Fingernails":
						__gong__assignSliceOfPointers(&inst.Fingernails, rhs, identifierMap)
					case "Hole":
						__gong__assignSliceOfPointers(&inst.Hole, rhs, identifierMap)
					case "Arrow":
						__gong__assignSliceOfPointers(&inst.Arrow, rhs, identifierMap)
					case "Handbell":
						__gong__assignSliceOfPointers(&inst.Handbell, rhs, identifierMap)
					case "Brass_bend":
						__gong__assignSliceOfPointers(&inst.Brass_bend, rhs, identifierMap)
					case "Flip":
						__gong__assignSliceOfPointers(&inst.Flip, rhs, identifierMap)
					case "Smear":
						__gong__assignSliceOfPointers(&inst.Smear, rhs, identifierMap)
					case "Open":
						__gong__assignSliceOfPointers(&inst.Open, rhs, identifierMap)
					case "Half_muted":
						__gong__assignSliceOfPointers(&inst.Half_muted, rhs, identifierMap)
					case "Harmon_mute":
						__gong__assignSliceOfPointers(&inst.Harmon_mute, rhs, identifierMap)
					case "Golpe":
						__gong__assignSliceOfPointers(&inst.Golpe, rhs, identifierMap)
					case "Other_technical":
						__gong__assignSliceOfPointers(&inst.Other_technical, rhs, identifierMap)
					}
				case *Text_element_data:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Lang":
						inst.Lang = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Underline":
						inst.Underline = GongExtractInt(rhs)
					case "Overline":
						inst.Overline = GongExtractInt(rhs)
					case "Line_through":
						inst.Line_through = GongExtractInt(rhs)
					case "Rotation":
						inst.Rotation = GongExtractString(rhs)
					case "Letter_spacing":
						inst.Letter_spacing = GongExtractString(rhs)
					case "Dir":
						GongUnmarshallEnum(&inst.Dir, rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Tie:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type_XSD":
						GongUnmarshallEnum(&inst.Type_XSD, rhs)
					case "Time_only":
						inst.Time_only = GongExtractString(rhs)
					}
				case *Tied:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type_XSD":
						GongUnmarshallEnum(&inst.Type_XSD, rhs)
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Line_type":
						GongUnmarshallEnum(&inst.Line_type, rhs)
					case "Dash_length":
						inst.Dash_length = GongExtractString(rhs)
					case "Space_length":
						inst.Space_length = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Placement":
						GongUnmarshallEnum(&inst.Placement, rhs)
					case "Orientation":
						GongUnmarshallEnum(&inst.Orientation, rhs)
					case "Bezier_x":
						inst.Bezier_x = GongExtractString(rhs)
					case "Bezier_y":
						inst.Bezier_y = GongExtractString(rhs)
					case "Bezier_x2":
						inst.Bezier_x2 = GongExtractString(rhs)
					case "Bezier_y2":
						inst.Bezier_y2 = GongExtractString(rhs)
					case "Bezier_offset":
						inst.Bezier_offset = GongExtractString(rhs)
					case "Bezier_offset2":
						inst.Bezier_offset2 = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					}
				case *Time:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Symbol":
						GongUnmarshallEnum(&inst.Symbol, rhs)
					case "Separator":
						GongUnmarshallEnum(&inst.Separator, rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Halign":
						GongUnmarshallEnum(&inst.Halign, rhs)
					case "Valign":
						GongUnmarshallEnum(&inst.Valign, rhs)
					case "Print_object":
						GongUnmarshallEnum(&inst.Print_object, rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Beats":
						inst.Beats = GongExtractString(rhs)
					case "Beat_type":
						inst.Beat_type = GongExtractString(rhs)
					case "Interchangeable":
						__gong__assignPointer(&inst.Interchangeable, rhs, identifierMap)
					case "Senza_misura":
						inst.Senza_misura = GongExtractString(rhs)
					}
				case *Time_modification:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Actual_notes":
						inst.Actual_notes = GongExtractInt(rhs)
					case "Normal_notes":
						inst.Normal_notes = GongExtractInt(rhs)
					case "Normal_type":
						GongUnmarshallEnum(&inst.Normal_type, rhs)
					case "Normal_dot":
						inst.Normal_dot = GongExtractString(rhs)
					}
				case *Timpani:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					}
				case *Transpose:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Diatonic":
						inst.Diatonic = GongExtractInt(rhs)
					case "Chromatic":
						inst.Chromatic = GongExtractString(rhs)
					case "Octave_change":
						inst.Octave_change = GongExtractInt(rhs)
					case "Double":
						inst.Double = GongExtractFloat(rhs)
					}
				case *Tremolo:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type_XSD":
						GongUnmarshallEnum(&inst.Type_XSD, rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						GongUnmarshallEnum(&inst.Placement, rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractInt(rhs)
					}
				case *Tuplet:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type_XSD":
						GongUnmarshallEnum(&inst.Type_XSD, rhs)
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Bracket":
						GongUnmarshallEnum(&inst.Bracket, rhs)
					case "Show_number":
						GongUnmarshallEnum(&inst.Show_number, rhs)
					case "Show_type":
						GongUnmarshallEnum(&inst.Show_type, rhs)
					case "Line_shape":
						GongUnmarshallEnum(&inst.Line_shape, rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Placement":
						GongUnmarshallEnum(&inst.Placement, rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Tuplet_actual":
						__gong__assignPointer(&inst.Tuplet_actual, rhs, identifierMap)
					case "Tuplet_normal":
						__gong__assignPointer(&inst.Tuplet_normal, rhs, identifierMap)
					}
				case *Tuplet_dot:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					}
				case *Tuplet_number:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractInt(rhs)
					}
				case *Tuplet_portion:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Tuplet_number":
						__gong__assignPointer(&inst.Tuplet_number, rhs, identifierMap)
					case "Tuplet_type":
						__gong__assignPointer(&inst.Tuplet_type, rhs, identifierMap)
					case "Tuplet_dot":
						__gong__assignSliceOfPointers(&inst.Tuplet_dot, rhs, identifierMap)
					}
				case *Tuplet_type:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						GongUnmarshallEnum(&inst.Font_style, rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						GongUnmarshallEnum(&inst.Font_weight, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "EnclosedText":
						GongUnmarshallEnum(&inst.EnclosedText, rhs)
					}
				case *Typed_text:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type_XSD":
						inst.Type_XSD = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Unpitched:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Display_step":
						GongUnmarshallEnum(&inst.Display_step, rhs)
					case "Display_octave":
						inst.Display_octave = GongExtractInt(rhs)
					}
				case *Virtual_instrument:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Virtual_library":
						inst.Virtual_library = GongExtractString(rhs)
					case "Virtual_name":
						inst.Virtual_name = GongExtractString(rhs)
					}
				case *Wait:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Player":
						inst.Player = GongExtractString(rhs)
					case "Time_only":
						inst.Time_only = GongExtractString(rhs)
					}
				case *Wavy_line:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type_XSD":
						GongUnmarshallEnum(&inst.Type_XSD, rhs)
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Placement":
						GongUnmarshallEnum(&inst.Placement, rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Start_note":
						GongUnmarshallEnum(&inst.Start_note, rhs)
					case "Trill_step":
						GongUnmarshallEnum(&inst.Trill_step, rhs)
					case "Two_note_turn":
						GongUnmarshallEnum(&inst.Two_note_turn, rhs)
					case "Accelerate":
						GongUnmarshallEnum(&inst.Accelerate, rhs)
					case "Beats":
						inst.Beats = GongExtractString(rhs)
					case "Second_beat":
						inst.Second_beat = GongExtractString(rhs)
					case "Last_beat":
						inst.Last_beat = GongExtractString(rhs)
					}
				case *Wedge:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type_XSD":
						GongUnmarshallEnum(&inst.Type_XSD, rhs)
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Spread":
						inst.Spread = GongExtractString(rhs)
					case "Niente":
						GongUnmarshallEnum(&inst.Niente, rhs)
					case "Line_type":
						GongUnmarshallEnum(&inst.Line_type, rhs)
					case "Dash_length":
						inst.Dash_length = GongExtractString(rhs)
					case "Space_length":
						inst.Space_length = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					}
				case *Wood:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "EnclosedText":
						GongUnmarshallEnum(&inst.EnclosedText, rhs)
					}
				case *Work:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Work_number":
						inst.Work_number = GongExtractString(rhs)
					case "Work_title":
						inst.Work_title = GongExtractString(rhs)
					case "Opus":
						__gong__assignPointer(&inst.Opus, rhs, identifierMap)
					}
							}
						}
					}
				}
			}
		case *ast.ExprStmt:
			if call, ok := node.X.(*ast.CallExpr); ok {
				if sel, ok := call.Fun.(*ast.SelectorExpr); ok {
					if sel.Sel.Name == "Unstage" {
						if ident, ok := sel.X.(*ast.Ident); ok {
							if instance, ok := identifierMap[ident.Name]; ok {
								if inst, ok := instance.(xmlxsd.GongstructIF); ok && stage.XmlxsdStage != nil {
									inst.UnstageVoid(stage.XmlxsdStage)
								}
								if inst, ok := instance.(xlink.GongstructIF); ok && stage.XlinkStage != nil {
									inst.UnstageVoid(stage.XlinkStage)
								}
								if inst, ok := instance.(GongstructIF); ok {
									inst.UnstageVoid(stage)
								}
							}
						}
					}
					if sel.Sel.Name == "Commit" {
						if ident, ok := sel.X.(*ast.Ident); ok && ident.Name == "stage" {
							if stage.IsInDeltaMode() && stage.GetNavigationMode() != GongNavigationModeNavigating {
								stage.Commit()
							} else {
								stage.ComputeInstancesNb()
								stage.ComputeReferenceAndOrders()
								if stage.OnInitCommitCallback != nil {
									stage.OnInitCommitCallback.BeforeCommit(stage)
								}
								if stage.OnInitCommitFromBackCallback != nil {
									stage.OnInitCommitFromBackCallback.BeforeCommit(stage)
								}
								stage.RunBeforeCommitHooks()
								stage.RunAfterCommitHooks()
							}
						}
					}
				}
			}
		}
		return true
	})

	return nil
}

// --- Generic Helpers for Unmarshallers (delegating to gong_runtime) ---

var GongExtractString = gong_runtime.ExtractString
var GongExtractInt = gong_runtime.ExtractInt
var GongExtractFloat = gong_runtime.ExtractFloat
var GongExtractBool = gong_runtime.ExtractBool
var GongExtractExpr = gong_runtime.ExtractExpr
var GongExtractDate = gong_runtime.ExtractDate

func __gong__extractMiddleUint(input string) (uint, error) {
	return gong_runtime.ExtractMiddleUint(input)
}

func __gong__stageSetInit[P interface {
	SetName(string)
	StageVoid(S)
	StagePreserveOrder(S, uint)
}, S any](instance P, stage S, identifier string, instanceName string, preserveOrder bool) any {
	instance.SetName(instanceName)
	if !preserveOrder {
		instance.StageVoid(stage)
	} else {
		if order, err := __gong__extractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifier", identifier)
			instance.StageVoid(stage)
		} else {
			instance.StagePreserveOrder(stage, order)
		}
	}
	return instance
}

func __gong__assignPointer[T any](targetPtr **T, rhs ast.Expr, identifierMap map[string]any) {
	if rIdent, ok := rhs.(*ast.Ident); ok {
		if rIdent.Name == "nil" {
			*targetPtr = nil
			return
		}
		if target, ok := identifierMap[rIdent.Name]; ok {
			if typedTarget, ok := target.(*T); ok {
				*targetPtr = typedTarget
			}
		}
	}
}

func __gong__assignSliceOfPointers[T any](slice *[]*T, rhs ast.Expr, identifierMap map[string]any) {
	if call, ok := rhs.(*ast.CallExpr); ok {
		funcName := ""
		var isSlices bool
		if sel, ok := call.Fun.(*ast.SelectorExpr); ok {
			if ident, ok := sel.X.(*ast.Ident); ok && ident.Name == "slices" {
				isSlices = true
				funcName = sel.Sel.Name
			}
		} else if ident, ok := call.Fun.(*ast.Ident); ok {
			funcName = ident.Name
		}

		if isSlices {
			if funcName == "Delete" && len(call.Args) == 3 {
				start := GongExtractInt(call.Args[1])
				end := GongExtractInt(call.Args[2])
				if end <= len(*slice) {
					*slice = slices.Delete(*slice, start, end)
				}
			} else if funcName == "Insert" && len(call.Args) == 3 {
				index := GongExtractInt(call.Args[1])
				if index <= len(*slice) {
					if rIdent, ok := call.Args[2].(*ast.Ident); ok {
						if target, ok := identifierMap[rIdent.Name]; ok {
							if typedTarget, ok := target.(*T); ok {
								*slice = slices.Insert(*slice, index, typedTarget)
							}
						}
					}
				}
			}
		} else if funcName == "append" && len(call.Args) >= 2 {
			for i := 1; i < len(call.Args); i++ {
				if rIdent, ok := call.Args[i].(*ast.Ident); ok {
					if target, ok := identifierMap[rIdent.Name]; ok {
						if typedTarget, ok := target.(*T); ok {
							*slice = append(*slice, typedTarget)
						}
					}
				}
			}
		}
	}
}

// GongUnmarshallSliceOfPointers handles append, slices.Delete, and slices.Insert for slice fields
func GongUnmarshallSliceOfPointers[T GongstructPtr](
	slice *[]T,
	valueExpr ast.Expr,
	identifierMap map[string]GongstructIF) (err error) {
	return gong_runtime.UnmarshallSliceOfPointers(slice, valueExpr, identifierMap)
}

// GongUnmarshallPointer handles assignment of a single pointer field
func GongUnmarshallPointer[T GongstructPtr](
	ptr *T,
	valueExpr ast.Expr,
	identifierMap map[string]GongstructIF) {
	gong_runtime.UnmarshallPointer(ptr, valueExpr, identifierMap)
}

// GongUnmarshallEnum handles assignment of enum fields (via SelectorExpr or String fallback)
func GongUnmarshallEnum[T interface{ FromCodeString(string) error }](
	ptr T,
	valueExpr ast.Expr) {
	gong_runtime.UnmarshallEnum(ptr, valueExpr)
}

// GongInitialize initializes a staged instance, sets its name, and stages it
func GongInitialize[P interface {
	GongstructIF
	StagePreserveOrder(stage *Stage, order uint)
}](instance P, stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance.SetName(instanceName)
	if !preserveOrder {
		instance.StageVoid(stage)
	} else {
		if newOrder, err := __gong__extractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.StageVoid(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

// insertion point per named struct
type A_directiveUnmarshaller struct{}

func (u *A_directiveUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_directive), stage, identifier, instanceName, preserveOrder)
}

func (u *A_directiveUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_directive)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Lang":
		instance.Lang = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type A_measureUnmarshaller struct{}

func (u *A_measureUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_measure), stage, identifier, instanceName, preserveOrder)
}

func (u *A_measureUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_measure)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Number":
		instance.Number = GongExtractString(valueExpr)
	case "Text":
		instance.Text = GongExtractString(valueExpr)
	case "Implicit":
		GongUnmarshallEnum(&instance.Implicit, valueExpr)
	case "Non_controlling":
		GongUnmarshallEnum(&instance.Non_controlling, valueExpr)
	case "Width":
		instance.Width = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Note":
		GongUnmarshallSliceOfPointers(&instance.Note, valueExpr, identifierMap)
	case "Backup":
		GongUnmarshallSliceOfPointers(&instance.Backup, valueExpr, identifierMap)
	case "Forward":
		GongUnmarshallSliceOfPointers(&instance.Forward, valueExpr, identifierMap)
	case "Direction":
		GongUnmarshallSliceOfPointers(&instance.Direction, valueExpr, identifierMap)
	case "Attributes":
		GongUnmarshallSliceOfPointers(&instance.Attributes, valueExpr, identifierMap)
	case "Harmony":
		GongUnmarshallSliceOfPointers(&instance.Harmony, valueExpr, identifierMap)
	case "Figured_bass":
		GongUnmarshallSliceOfPointers(&instance.Figured_bass, valueExpr, identifierMap)
	case "Print_XSD":
		GongUnmarshallSliceOfPointers(&instance.Print_XSD, valueExpr, identifierMap)
	case "Sound":
		GongUnmarshallSliceOfPointers(&instance.Sound, valueExpr, identifierMap)
	case "Listening":
		GongUnmarshallSliceOfPointers(&instance.Listening, valueExpr, identifierMap)
	case "Barline":
		GongUnmarshallSliceOfPointers(&instance.Barline, valueExpr, identifierMap)
	case "Grouping":
		GongUnmarshallSliceOfPointers(&instance.Grouping, valueExpr, identifierMap)
	case "Link":
		GongUnmarshallSliceOfPointers(&instance.Link, valueExpr, identifierMap)
	case "Bookmark":
		GongUnmarshallSliceOfPointers(&instance.Bookmark, valueExpr, identifierMap)
	}
	return nil
}

type A_measure_1Unmarshaller struct{}

func (u *A_measure_1Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_measure_1), stage, identifier, instanceName, preserveOrder)
}

func (u *A_measure_1Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_measure_1)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Number":
		instance.Number = GongExtractString(valueExpr)
	case "Text":
		instance.Text = GongExtractString(valueExpr)
	case "Implicit":
		GongUnmarshallEnum(&instance.Implicit, valueExpr)
	case "Non_controlling":
		GongUnmarshallEnum(&instance.Non_controlling, valueExpr)
	case "Width":
		instance.Width = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Part":
		GongUnmarshallSliceOfPointers(&instance.Part, valueExpr, identifierMap)
	}
	return nil
}

type A_partUnmarshaller struct{}

func (u *A_partUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_part), stage, identifier, instanceName, preserveOrder)
}

func (u *A_partUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_part)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Measure":
		GongUnmarshallSliceOfPointers(&instance.Measure, valueExpr, identifierMap)
	}
	return nil
}

type A_part_1Unmarshaller struct{}

func (u *A_part_1Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_part_1), stage, identifier, instanceName, preserveOrder)
}

func (u *A_part_1Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_part_1)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Note":
		GongUnmarshallSliceOfPointers(&instance.Note, valueExpr, identifierMap)
	case "Backup":
		GongUnmarshallSliceOfPointers(&instance.Backup, valueExpr, identifierMap)
	case "Forward":
		GongUnmarshallSliceOfPointers(&instance.Forward, valueExpr, identifierMap)
	case "Direction":
		GongUnmarshallSliceOfPointers(&instance.Direction, valueExpr, identifierMap)
	case "Attributes":
		GongUnmarshallSliceOfPointers(&instance.Attributes, valueExpr, identifierMap)
	case "Harmony":
		GongUnmarshallSliceOfPointers(&instance.Harmony, valueExpr, identifierMap)
	case "Figured_bass":
		GongUnmarshallSliceOfPointers(&instance.Figured_bass, valueExpr, identifierMap)
	case "Print_XSD":
		GongUnmarshallSliceOfPointers(&instance.Print_XSD, valueExpr, identifierMap)
	case "Sound":
		GongUnmarshallSliceOfPointers(&instance.Sound, valueExpr, identifierMap)
	case "Listening":
		GongUnmarshallSliceOfPointers(&instance.Listening, valueExpr, identifierMap)
	case "Barline":
		GongUnmarshallSliceOfPointers(&instance.Barline, valueExpr, identifierMap)
	case "Grouping":
		GongUnmarshallSliceOfPointers(&instance.Grouping, valueExpr, identifierMap)
	case "Link":
		GongUnmarshallSliceOfPointers(&instance.Link, valueExpr, identifierMap)
	case "Bookmark":
		GongUnmarshallSliceOfPointers(&instance.Bookmark, valueExpr, identifierMap)
	}
	return nil
}

type AccidentalUnmarshaller struct{}

func (u *AccidentalUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Accidental), stage, identifier, instanceName, preserveOrder)
}

func (u *AccidentalUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Accidental)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Cautionary":
		GongUnmarshallEnum(&instance.Cautionary, valueExpr)
	case "Editorial":
		GongUnmarshallEnum(&instance.Editorial, valueExpr)
	case "Smufl":
		instance.Smufl = GongExtractString(valueExpr)
	case "Parentheses":
		GongUnmarshallEnum(&instance.Parentheses, valueExpr)
	case "Bracket":
		GongUnmarshallEnum(&instance.Bracket, valueExpr)
	case "Size":
		GongUnmarshallEnum(&instance.Size, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "EnclosedText":
		GongUnmarshallEnum(&instance.EnclosedText, valueExpr)
	}
	return nil
}

type Accidental_markUnmarshaller struct{}

func (u *Accidental_markUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Accidental_mark), stage, identifier, instanceName, preserveOrder)
}

func (u *Accidental_markUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Accidental_mark)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Smufl":
		instance.Smufl = GongExtractString(valueExpr)
	case "Parentheses":
		GongUnmarshallEnum(&instance.Parentheses, valueExpr)
	case "Bracket":
		GongUnmarshallEnum(&instance.Bracket, valueExpr)
	case "Size":
		GongUnmarshallEnum(&instance.Size, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Placement":
		GongUnmarshallEnum(&instance.Placement, valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "EnclosedText":
		GongUnmarshallEnum(&instance.EnclosedText, valueExpr)
	}
	return nil
}

type Accidental_textUnmarshaller struct{}

func (u *Accidental_textUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Accidental_text), stage, identifier, instanceName, preserveOrder)
}

func (u *Accidental_textUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Accidental_text)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Smufl":
		instance.Smufl = GongExtractString(valueExpr)
	case "Lang":
		instance.Lang = GongExtractString(valueExpr)
	case "Space":
		instance.Space = GongExtractString(valueExpr)
	case "Justify":
		GongUnmarshallEnum(&instance.Justify, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Halign":
		GongUnmarshallEnum(&instance.Halign, valueExpr)
	case "Valign":
		GongUnmarshallEnum(&instance.Valign, valueExpr)
	case "Underline":
		instance.Underline = GongExtractInt(valueExpr)
	case "Overline":
		instance.Overline = GongExtractInt(valueExpr)
	case "Line_through":
		instance.Line_through = GongExtractInt(valueExpr)
	case "Rotation":
		instance.Rotation = GongExtractString(valueExpr)
	case "Letter_spacing":
		instance.Letter_spacing = GongExtractString(valueExpr)
	case "Line_height":
		instance.Line_height = GongExtractString(valueExpr)
	case "Dir":
		GongUnmarshallEnum(&instance.Dir, valueExpr)
	case "Enclosure":
		GongUnmarshallEnum(&instance.Enclosure, valueExpr)
	case "EnclosedText":
		GongUnmarshallEnum(&instance.EnclosedText, valueExpr)
	}
	return nil
}

type AccordUnmarshaller struct{}

func (u *AccordUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Accord), stage, identifier, instanceName, preserveOrder)
}

func (u *AccordUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Accord)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "String_XSD":
		instance.String_XSD = GongExtractInt(valueExpr)
	case "Tuning_step":
		GongUnmarshallEnum(&instance.Tuning_step, valueExpr)
	case "Tuning_alter":
		instance.Tuning_alter = GongExtractString(valueExpr)
	case "Tuning_octave":
		instance.Tuning_octave = GongExtractInt(valueExpr)
	}
	return nil
}

type Accordion_registrationUnmarshaller struct{}

func (u *Accordion_registrationUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Accordion_registration), stage, identifier, instanceName, preserveOrder)
}

func (u *Accordion_registrationUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Accordion_registration)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Halign":
		GongUnmarshallEnum(&instance.Halign, valueExpr)
	case "Valign":
		GongUnmarshallEnum(&instance.Valign, valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Accordion_high":
		instance.Accordion_high = GongExtractString(valueExpr)
	case "Accordion_middle":
		instance.Accordion_middle = GongExtractInt(valueExpr)
	case "Accordion_low":
		instance.Accordion_low = GongExtractString(valueExpr)
	}
	return nil
}

type AppearanceUnmarshaller struct{}

func (u *AppearanceUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Appearance), stage, identifier, instanceName, preserveOrder)
}

func (u *AppearanceUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Appearance)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Line_width":
		GongUnmarshallSliceOfPointers(&instance.Line_width, valueExpr, identifierMap)
	case "Note_size":
		GongUnmarshallSliceOfPointers(&instance.Note_size, valueExpr, identifierMap)
	case "Distance":
		GongUnmarshallSliceOfPointers(&instance.Distance, valueExpr, identifierMap)
	case "Glyph":
		GongUnmarshallSliceOfPointers(&instance.Glyph, valueExpr, identifierMap)
	case "Other_appearance":
		GongUnmarshallSliceOfPointers(&instance.Other_appearance, valueExpr, identifierMap)
	}
	return nil
}

type ArpeggiateUnmarshaller struct{}

func (u *ArpeggiateUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Arpeggiate), stage, identifier, instanceName, preserveOrder)
}

func (u *ArpeggiateUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Arpeggiate)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "Direction":
		GongUnmarshallEnum(&instance.Direction, valueExpr)
	case "Unbroken":
		GongUnmarshallEnum(&instance.Unbroken, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Placement":
		GongUnmarshallEnum(&instance.Placement, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	}
	return nil
}

type ArrowUnmarshaller struct{}

func (u *ArrowUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Arrow), stage, identifier, instanceName, preserveOrder)
}

func (u *ArrowUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Arrow)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Placement":
		GongUnmarshallEnum(&instance.Placement, valueExpr)
	case "Smufl":
		instance.Smufl = GongExtractString(valueExpr)
	case "Arrow_direction":
		GongUnmarshallEnum(&instance.Arrow_direction, valueExpr)
	case "Arrow_style":
		GongUnmarshallEnum(&instance.Arrow_style, valueExpr)
	case "Arrowhead":
		instance.Arrowhead = GongExtractString(valueExpr)
	case "Circular_arrow":
		GongUnmarshallEnum(&instance.Circular_arrow, valueExpr)
	}
	return nil
}

type ArticulationsUnmarshaller struct{}

func (u *ArticulationsUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Articulations), stage, identifier, instanceName, preserveOrder)
}

func (u *ArticulationsUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Articulations)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Accent":
		GongUnmarshallSliceOfPointers(&instance.Accent, valueExpr, identifierMap)
	case "Strong_accent":
		GongUnmarshallSliceOfPointers(&instance.Strong_accent, valueExpr, identifierMap)
	case "Staccato":
		GongUnmarshallSliceOfPointers(&instance.Staccato, valueExpr, identifierMap)
	case "Tenuto":
		GongUnmarshallSliceOfPointers(&instance.Tenuto, valueExpr, identifierMap)
	case "Detached_legato":
		GongUnmarshallSliceOfPointers(&instance.Detached_legato, valueExpr, identifierMap)
	case "Staccatissimo":
		GongUnmarshallSliceOfPointers(&instance.Staccatissimo, valueExpr, identifierMap)
	case "Spiccato":
		GongUnmarshallSliceOfPointers(&instance.Spiccato, valueExpr, identifierMap)
	case "Scoop":
		GongUnmarshallSliceOfPointers(&instance.Scoop, valueExpr, identifierMap)
	case "Plop":
		GongUnmarshallSliceOfPointers(&instance.Plop, valueExpr, identifierMap)
	case "Doit":
		GongUnmarshallSliceOfPointers(&instance.Doit, valueExpr, identifierMap)
	case "Falloff":
		GongUnmarshallSliceOfPointers(&instance.Falloff, valueExpr, identifierMap)
	case "Breath_mark":
		GongUnmarshallSliceOfPointers(&instance.Breath_mark, valueExpr, identifierMap)
	case "Caesura":
		GongUnmarshallSliceOfPointers(&instance.Caesura, valueExpr, identifierMap)
	case "Stress":
		GongUnmarshallSliceOfPointers(&instance.Stress, valueExpr, identifierMap)
	case "Unstress":
		GongUnmarshallSliceOfPointers(&instance.Unstress, valueExpr, identifierMap)
	case "Soft_accent":
		GongUnmarshallSliceOfPointers(&instance.Soft_accent, valueExpr, identifierMap)
	case "Other_articulation":
		GongUnmarshallSliceOfPointers(&instance.Other_articulation, valueExpr, identifierMap)
	}
	return nil
}

type AssessUnmarshaller struct{}

func (u *AssessUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Assess), stage, identifier, instanceName, preserveOrder)
}

func (u *AssessUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Assess)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type_XSD":
		GongUnmarshallEnum(&instance.Type_XSD, valueExpr)
	case "Player":
		instance.Player = GongExtractString(valueExpr)
	case "Time_only":
		instance.Time_only = GongExtractString(valueExpr)
	}
	return nil
}

type AttributesUnmarshaller struct{}

func (u *AttributesUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Attributes), stage, identifier, instanceName, preserveOrder)
}

func (u *AttributesUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Attributes)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Footnote":
		GongUnmarshallPointer(&instance.Footnote, valueExpr, identifierMap)
	case "Level":
		GongUnmarshallPointer(&instance.Level, valueExpr, identifierMap)
	case "Divisions":
		instance.Divisions = GongExtractString(valueExpr)
	case "Key":
		GongUnmarshallSliceOfPointers(&instance.Key, valueExpr, identifierMap)
	case "Time":
		instance.Time = GongExtractString(valueExpr)
	case "Staves":
		instance.Staves = GongExtractInt(valueExpr)
	case "Part_symbol":
		GongUnmarshallPointer(&instance.Part_symbol, valueExpr, identifierMap)
	case "Instruments":
		instance.Instruments = GongExtractInt(valueExpr)
	case "Clef":
		GongUnmarshallSliceOfPointers(&instance.Clef, valueExpr, identifierMap)
	case "Staff_details":
		GongUnmarshallSliceOfPointers(&instance.Staff_details, valueExpr, identifierMap)
	case "Transpose":
		GongUnmarshallSliceOfPointers(&instance.Transpose, valueExpr, identifierMap)
	case "For_part":
		GongUnmarshallSliceOfPointers(&instance.For_part, valueExpr, identifierMap)
	case "Directive":
		GongUnmarshallSliceOfPointers(&instance.Directive, valueExpr, identifierMap)
	case "Measure_style":
		GongUnmarshallSliceOfPointers(&instance.Measure_style, valueExpr, identifierMap)
	}
	return nil
}

type BackupUnmarshaller struct{}

func (u *BackupUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Backup), stage, identifier, instanceName, preserveOrder)
}

func (u *BackupUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Backup)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Footnote":
		GongUnmarshallPointer(&instance.Footnote, valueExpr, identifierMap)
	case "Level":
		GongUnmarshallPointer(&instance.Level, valueExpr, identifierMap)
	}
	return nil
}

type Bar_style_colorUnmarshaller struct{}

func (u *Bar_style_colorUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Bar_style_color), stage, identifier, instanceName, preserveOrder)
}

func (u *Bar_style_colorUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Bar_style_color)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "EnclosedText":
		GongUnmarshallEnum(&instance.EnclosedText, valueExpr)
	}
	return nil
}

type BarlineUnmarshaller struct{}

func (u *BarlineUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Barline), stage, identifier, instanceName, preserveOrder)
}

func (u *BarlineUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Barline)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Location":
		GongUnmarshallEnum(&instance.Location, valueExpr)
	case "Segno":
		instance.Segno = GongExtractString(valueExpr)
	case "Coda":
		instance.Coda = GongExtractString(valueExpr)
	case "Divisions":
		instance.Divisions = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Bar_style":
		GongUnmarshallPointer(&instance.Bar_style, valueExpr, identifierMap)
	case "Footnote":
		GongUnmarshallPointer(&instance.Footnote, valueExpr, identifierMap)
	case "Level":
		GongUnmarshallPointer(&instance.Level, valueExpr, identifierMap)
	case "Wavy_line":
		GongUnmarshallPointer(&instance.Wavy_line, valueExpr, identifierMap)
	case "Segno_1":
		GongUnmarshallPointer(&instance.Segno_1, valueExpr, identifierMap)
	case "Coda_1":
		GongUnmarshallPointer(&instance.Coda_1, valueExpr, identifierMap)
	case "Fermata":
		GongUnmarshallPointer(&instance.Fermata, valueExpr, identifierMap)
	case "Ending":
		GongUnmarshallPointer(&instance.Ending, valueExpr, identifierMap)
	case "Repeat":
		GongUnmarshallPointer(&instance.Repeat, valueExpr, identifierMap)
	}
	return nil
}

type BarreUnmarshaller struct{}

func (u *BarreUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Barre), stage, identifier, instanceName, preserveOrder)
}

func (u *BarreUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Barre)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type_XSD":
		GongUnmarshallEnum(&instance.Type_XSD, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	}
	return nil
}

type BassUnmarshaller struct{}

func (u *BassUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Bass), stage, identifier, instanceName, preserveOrder)
}

func (u *BassUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Bass)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Arrangement":
		GongUnmarshallEnum(&instance.Arrangement, valueExpr)
	case "Bass_separator":
		GongUnmarshallPointer(&instance.Bass_separator, valueExpr, identifierMap)
	case "Bass_step":
		GongUnmarshallPointer(&instance.Bass_step, valueExpr, identifierMap)
	case "Bass_alter":
		GongUnmarshallPointer(&instance.Bass_alter, valueExpr, identifierMap)
	}
	return nil
}

type Bass_stepUnmarshaller struct{}

func (u *Bass_stepUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Bass_step), stage, identifier, instanceName, preserveOrder)
}

func (u *Bass_stepUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Bass_step)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Text":
		instance.Text = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "EnclosedText":
		GongUnmarshallEnum(&instance.EnclosedText, valueExpr)
	}
	return nil
}

type BeamUnmarshaller struct{}

func (u *BeamUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Beam), stage, identifier, instanceName, preserveOrder)
}

func (u *BeamUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Beam)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "Repeater":
		GongUnmarshallEnum(&instance.Repeater, valueExpr)
	case "Fan":
		GongUnmarshallEnum(&instance.Fan, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "EnclosedText":
		GongUnmarshallEnum(&instance.EnclosedText, valueExpr)
	}
	return nil
}

type Beat_repeatUnmarshaller struct{}

func (u *Beat_repeatUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Beat_repeat), stage, identifier, instanceName, preserveOrder)
}

func (u *Beat_repeatUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Beat_repeat)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type_XSD":
		GongUnmarshallEnum(&instance.Type_XSD, valueExpr)
	case "Slashes":
		instance.Slashes = GongExtractInt(valueExpr)
	case "Use_dots":
		GongUnmarshallEnum(&instance.Use_dots, valueExpr)
	case "Slash_type":
		GongUnmarshallEnum(&instance.Slash_type, valueExpr)
	case "Slash_dot":
		instance.Slash_dot = GongExtractString(valueExpr)
	case "Except_voice":
		instance.Except_voice = GongExtractString(valueExpr)
	}
	return nil
}

type Beat_unit_tiedUnmarshaller struct{}

func (u *Beat_unit_tiedUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Beat_unit_tied), stage, identifier, instanceName, preserveOrder)
}

func (u *Beat_unit_tiedUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Beat_unit_tied)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Beat_unit":
		GongUnmarshallEnum(&instance.Beat_unit, valueExpr)
	case "Beat_unit_dot":
		instance.Beat_unit_dot = GongExtractString(valueExpr)
	}
	return nil
}

type BeaterUnmarshaller struct{}

func (u *BeaterUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Beater), stage, identifier, instanceName, preserveOrder)
}

func (u *BeaterUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Beater)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Tip":
		GongUnmarshallEnum(&instance.Tip, valueExpr)
	case "EnclosedText":
		GongUnmarshallEnum(&instance.EnclosedText, valueExpr)
	}
	return nil
}

type BendUnmarshaller struct{}

func (u *BendUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Bend), stage, identifier, instanceName, preserveOrder)
}

func (u *BendUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Bend)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Shape":
		GongUnmarshallEnum(&instance.Shape, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Accelerate":
		GongUnmarshallEnum(&instance.Accelerate, valueExpr)
	case "Beats":
		instance.Beats = GongExtractString(valueExpr)
	case "First_beat":
		instance.First_beat = GongExtractString(valueExpr)
	case "Last_beat":
		instance.Last_beat = GongExtractString(valueExpr)
	case "Bend_alter":
		instance.Bend_alter = GongExtractString(valueExpr)
	case "Pre_bend":
		instance.Pre_bend = GongExtractString(valueExpr)
	case "Release":
		GongUnmarshallPointer(&instance.Release, valueExpr, identifierMap)
	case "With_bar":
		GongUnmarshallPointer(&instance.With_bar, valueExpr, identifierMap)
	}
	return nil
}

type BookmarkUnmarshaller struct{}

func (u *BookmarkUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Bookmark), stage, identifier, instanceName, preserveOrder)
}

func (u *BookmarkUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Bookmark)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "NameXSD":
		instance.NameXSD = GongExtractString(valueExpr)
	case "Element":
		instance.Element = GongExtractString(valueExpr)
	case "Position":
		instance.Position = GongExtractInt(valueExpr)
	}
	return nil
}

type BracketUnmarshaller struct{}

func (u *BracketUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Bracket), stage, identifier, instanceName, preserveOrder)
}

func (u *BracketUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Bracket)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type_XSD":
		GongUnmarshallEnum(&instance.Type_XSD, valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "Line_end":
		GongUnmarshallEnum(&instance.Line_end, valueExpr)
	case "End_length":
		instance.End_length = GongExtractString(valueExpr)
	case "Line_type":
		GongUnmarshallEnum(&instance.Line_type, valueExpr)
	case "Dash_length":
		instance.Dash_length = GongExtractString(valueExpr)
	case "Space_length":
		instance.Space_length = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	}
	return nil
}

type Breath_markUnmarshaller struct{}

func (u *Breath_markUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Breath_mark), stage, identifier, instanceName, preserveOrder)
}

func (u *Breath_markUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Breath_mark)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Placement":
		GongUnmarshallEnum(&instance.Placement, valueExpr)
	case "EnclosedText":
		GongUnmarshallEnum(&instance.EnclosedText, valueExpr)
	}
	return nil
}

type CaesuraUnmarshaller struct{}

func (u *CaesuraUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Caesura), stage, identifier, instanceName, preserveOrder)
}

func (u *CaesuraUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Caesura)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Placement":
		GongUnmarshallEnum(&instance.Placement, valueExpr)
	case "EnclosedText":
		GongUnmarshallEnum(&instance.EnclosedText, valueExpr)
	}
	return nil
}

type CancelUnmarshaller struct{}

func (u *CancelUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Cancel), stage, identifier, instanceName, preserveOrder)
}

func (u *CancelUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Cancel)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Location":
		GongUnmarshallEnum(&instance.Location, valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractInt(valueExpr)
	}
	return nil
}

type ClefUnmarshaller struct{}

func (u *ClefUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Clef), stage, identifier, instanceName, preserveOrder)
}

func (u *ClefUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Clef)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "Additional":
		GongUnmarshallEnum(&instance.Additional, valueExpr)
	case "Size":
		GongUnmarshallEnum(&instance.Size, valueExpr)
	case "After_barline":
		GongUnmarshallEnum(&instance.After_barline, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Print_object":
		GongUnmarshallEnum(&instance.Print_object, valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Sign":
		GongUnmarshallEnum(&instance.Sign, valueExpr)
	case "Line":
		instance.Line = GongExtractInt(valueExpr)
	case "Clef_octave_change":
		instance.Clef_octave_change = GongExtractInt(valueExpr)
	}
	return nil
}

type CodaUnmarshaller struct{}

func (u *CodaUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Coda), stage, identifier, instanceName, preserveOrder)
}

func (u *CodaUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Coda)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Smufl":
		instance.Smufl = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Halign":
		GongUnmarshallEnum(&instance.Halign, valueExpr)
	case "Valign":
		GongUnmarshallEnum(&instance.Valign, valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	}
	return nil
}

type CreditUnmarshaller struct{}

func (u *CreditUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Credit), stage, identifier, instanceName, preserveOrder)
}

func (u *CreditUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Credit)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Page":
		instance.Page = GongExtractInt(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Credit_type":
		instance.Credit_type = GongExtractString(valueExpr)
	case "Link":
		GongUnmarshallSliceOfPointers(&instance.Link, valueExpr, identifierMap)
	case "Bookmark":
		GongUnmarshallSliceOfPointers(&instance.Bookmark, valueExpr, identifierMap)
	case "Credit_image":
		GongUnmarshallPointer(&instance.Credit_image, valueExpr, identifierMap)
	case "Credit_words":
		GongUnmarshallPointer(&instance.Credit_words, valueExpr, identifierMap)
	case "Credit_symbol":
		GongUnmarshallPointer(&instance.Credit_symbol, valueExpr, identifierMap)
	}
	return nil
}

type DashesUnmarshaller struct{}

func (u *DashesUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Dashes), stage, identifier, instanceName, preserveOrder)
}

func (u *DashesUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Dashes)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type_XSD":
		GongUnmarshallEnum(&instance.Type_XSD, valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "Dash_length":
		instance.Dash_length = GongExtractString(valueExpr)
	case "Space_length":
		instance.Space_length = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	}
	return nil
}

type DefaultsUnmarshaller struct{}

func (u *DefaultsUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Defaults), stage, identifier, instanceName, preserveOrder)
}

func (u *DefaultsUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Defaults)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Scaling":
		GongUnmarshallPointer(&instance.Scaling, valueExpr, identifierMap)
	case "Concert_score":
		instance.Concert_score = GongExtractString(valueExpr)
	case "Page_layout":
		GongUnmarshallPointer(&instance.Page_layout, valueExpr, identifierMap)
	case "System_layout":
		GongUnmarshallPointer(&instance.System_layout, valueExpr, identifierMap)
	case "Staff_layout":
		GongUnmarshallSliceOfPointers(&instance.Staff_layout, valueExpr, identifierMap)
	case "Appearance":
		GongUnmarshallPointer(&instance.Appearance, valueExpr, identifierMap)
	case "Music_font":
		GongUnmarshallPointer(&instance.Music_font, valueExpr, identifierMap)
	case "Word_font":
		GongUnmarshallPointer(&instance.Word_font, valueExpr, identifierMap)
	case "Lyric_font":
		GongUnmarshallSliceOfPointers(&instance.Lyric_font, valueExpr, identifierMap)
	case "Lyric_language":
		GongUnmarshallSliceOfPointers(&instance.Lyric_language, valueExpr, identifierMap)
	}
	return nil
}

type DegreeUnmarshaller struct{}

func (u *DegreeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Degree), stage, identifier, instanceName, preserveOrder)
}

func (u *DegreeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Degree)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Print_object":
		GongUnmarshallEnum(&instance.Print_object, valueExpr)
	case "Degree_value":
		GongUnmarshallPointer(&instance.Degree_value, valueExpr, identifierMap)
	case "Degree_alter":
		GongUnmarshallPointer(&instance.Degree_alter, valueExpr, identifierMap)
	case "Degree_type":
		GongUnmarshallPointer(&instance.Degree_type, valueExpr, identifierMap)
	}
	return nil
}

type Degree_alterUnmarshaller struct{}

func (u *Degree_alterUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Degree_alter), stage, identifier, instanceName, preserveOrder)
}

func (u *Degree_alterUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Degree_alter)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Plus_minus":
		GongUnmarshallEnum(&instance.Plus_minus, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type Degree_typeUnmarshaller struct{}

func (u *Degree_typeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Degree_type), stage, identifier, instanceName, preserveOrder)
}

func (u *Degree_typeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Degree_type)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Text":
		instance.Text = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "EnclosedText":
		GongUnmarshallEnum(&instance.EnclosedText, valueExpr)
	}
	return nil
}

type Degree_valueUnmarshaller struct{}

func (u *Degree_valueUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Degree_value), stage, identifier, instanceName, preserveOrder)
}

func (u *Degree_valueUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Degree_value)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Symbol":
		GongUnmarshallEnum(&instance.Symbol, valueExpr)
	case "Text":
		instance.Text = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractInt(valueExpr)
	}
	return nil
}

type DirectionUnmarshaller struct{}

func (u *DirectionUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Direction), stage, identifier, instanceName, preserveOrder)
}

func (u *DirectionUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Direction)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Placement":
		GongUnmarshallEnum(&instance.Placement, valueExpr)
	case "Directive":
		GongUnmarshallEnum(&instance.Directive, valueExpr)
	case "System":
		instance.System = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Direction_type":
		GongUnmarshallSliceOfPointers(&instance.Direction_type, valueExpr, identifierMap)
	case "Offset":
		GongUnmarshallPointer(&instance.Offset, valueExpr, identifierMap)
	case "Footnote":
		GongUnmarshallPointer(&instance.Footnote, valueExpr, identifierMap)
	case "Level":
		GongUnmarshallPointer(&instance.Level, valueExpr, identifierMap)
	case "Voice":
		instance.Voice = GongExtractString(valueExpr)
	case "Staff":
		instance.Staff = GongExtractInt(valueExpr)
	case "Sound":
		GongUnmarshallPointer(&instance.Sound, valueExpr, identifierMap)
	case "Listening":
		GongUnmarshallPointer(&instance.Listening, valueExpr, identifierMap)
	}
	return nil
}

type Direction_typeUnmarshaller struct{}

func (u *Direction_typeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Direction_type), stage, identifier, instanceName, preserveOrder)
}

func (u *Direction_typeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Direction_type)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Rehearsal":
		GongUnmarshallSliceOfPointers(&instance.Rehearsal, valueExpr, identifierMap)
	case "Segno":
		GongUnmarshallSliceOfPointers(&instance.Segno, valueExpr, identifierMap)
	case "Coda":
		GongUnmarshallSliceOfPointers(&instance.Coda, valueExpr, identifierMap)
	case "Words":
		GongUnmarshallSliceOfPointers(&instance.Words, valueExpr, identifierMap)
	case "Symbol":
		GongUnmarshallSliceOfPointers(&instance.Symbol, valueExpr, identifierMap)
	case "Wedge":
		GongUnmarshallPointer(&instance.Wedge, valueExpr, identifierMap)
	case "Dynamics":
		GongUnmarshallSliceOfPointers(&instance.Dynamics, valueExpr, identifierMap)
	case "Dashes":
		GongUnmarshallPointer(&instance.Dashes, valueExpr, identifierMap)
	case "Bracket":
		GongUnmarshallPointer(&instance.Bracket, valueExpr, identifierMap)
	case "Pedal":
		GongUnmarshallPointer(&instance.Pedal, valueExpr, identifierMap)
	case "Metronome":
		GongUnmarshallPointer(&instance.Metronome, valueExpr, identifierMap)
	case "Octave_shift":
		GongUnmarshallPointer(&instance.Octave_shift, valueExpr, identifierMap)
	case "Harp_pedals":
		GongUnmarshallPointer(&instance.Harp_pedals, valueExpr, identifierMap)
	case "Damp":
		GongUnmarshallPointer(&instance.Damp, valueExpr, identifierMap)
	case "Damp_all":
		GongUnmarshallPointer(&instance.Damp_all, valueExpr, identifierMap)
	case "Eyeglasses":
		GongUnmarshallPointer(&instance.Eyeglasses, valueExpr, identifierMap)
	case "String_mute":
		GongUnmarshallPointer(&instance.String_mute, valueExpr, identifierMap)
	case "Scordatura":
		GongUnmarshallPointer(&instance.Scordatura, valueExpr, identifierMap)
	case "Image":
		GongUnmarshallPointer(&instance.Image, valueExpr, identifierMap)
	case "Principal_voice":
		GongUnmarshallPointer(&instance.Principal_voice, valueExpr, identifierMap)
	case "Percussion":
		GongUnmarshallSliceOfPointers(&instance.Percussion, valueExpr, identifierMap)
	case "Accordion_registration":
		GongUnmarshallPointer(&instance.Accordion_registration, valueExpr, identifierMap)
	case "Staff_divide":
		GongUnmarshallPointer(&instance.Staff_divide, valueExpr, identifierMap)
	case "Other_direction":
		GongUnmarshallPointer(&instance.Other_direction, valueExpr, identifierMap)
	}
	return nil
}

type DistanceUnmarshaller struct{}

func (u *DistanceUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Distance), stage, identifier, instanceName, preserveOrder)
}

func (u *DistanceUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Distance)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type_XSD":
		instance.Type_XSD = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type DoubleUnmarshaller struct{}

func (u *DoubleUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Double), stage, identifier, instanceName, preserveOrder)
}

func (u *DoubleUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Double)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Above":
		GongUnmarshallEnum(&instance.Above, valueExpr)
	}
	return nil
}

type DynamicsUnmarshaller struct{}

func (u *DynamicsUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Dynamics), stage, identifier, instanceName, preserveOrder)
}

func (u *DynamicsUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Dynamics)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Halign":
		GongUnmarshallEnum(&instance.Halign, valueExpr)
	case "Valign":
		GongUnmarshallEnum(&instance.Valign, valueExpr)
	case "Placement":
		GongUnmarshallEnum(&instance.Placement, valueExpr)
	case "Underline":
		instance.Underline = GongExtractInt(valueExpr)
	case "Overline":
		instance.Overline = GongExtractInt(valueExpr)
	case "Line_through":
		instance.Line_through = GongExtractInt(valueExpr)
	case "Enclosure":
		GongUnmarshallEnum(&instance.Enclosure, valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "P":
		instance.P = GongExtractString(valueExpr)
	case "Pp":
		instance.Pp = GongExtractString(valueExpr)
	case "Ppp":
		instance.Ppp = GongExtractString(valueExpr)
	case "Pppp":
		instance.Pppp = GongExtractString(valueExpr)
	case "Ppppp":
		instance.Ppppp = GongExtractString(valueExpr)
	case "Pppppp":
		instance.Pppppp = GongExtractString(valueExpr)
	case "F":
		instance.F = GongExtractString(valueExpr)
	case "Ff":
		instance.Ff = GongExtractString(valueExpr)
	case "Fff":
		instance.Fff = GongExtractString(valueExpr)
	case "Ffff":
		instance.Ffff = GongExtractString(valueExpr)
	case "Fffff":
		instance.Fffff = GongExtractString(valueExpr)
	case "Ffffff":
		instance.Ffffff = GongExtractString(valueExpr)
	case "Mp":
		instance.Mp = GongExtractString(valueExpr)
	case "Mf":
		instance.Mf = GongExtractString(valueExpr)
	case "Sf":
		instance.Sf = GongExtractString(valueExpr)
	case "Sfp":
		instance.Sfp = GongExtractString(valueExpr)
	case "Sfpp":
		instance.Sfpp = GongExtractString(valueExpr)
	case "Fp":
		instance.Fp = GongExtractString(valueExpr)
	case "Rf":
		instance.Rf = GongExtractString(valueExpr)
	case "Rfz":
		instance.Rfz = GongExtractString(valueExpr)
	case "Sfz":
		instance.Sfz = GongExtractString(valueExpr)
	case "Sffz":
		instance.Sffz = GongExtractString(valueExpr)
	case "Fz":
		instance.Fz = GongExtractString(valueExpr)
	case "N":
		instance.N = GongExtractString(valueExpr)
	case "Pf":
		instance.Pf = GongExtractString(valueExpr)
	case "Sfzp":
		instance.Sfzp = GongExtractString(valueExpr)
	case "Other_dynamics":
		GongUnmarshallSliceOfPointers(&instance.Other_dynamics, valueExpr, identifierMap)
	}
	return nil
}

type EffectUnmarshaller struct{}

func (u *EffectUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Effect), stage, identifier, instanceName, preserveOrder)
}

func (u *EffectUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Effect)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Smufl":
		instance.Smufl = GongExtractString(valueExpr)
	case "EnclosedText":
		GongUnmarshallEnum(&instance.EnclosedText, valueExpr)
	}
	return nil
}

type ElisionUnmarshaller struct{}

func (u *ElisionUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Elision), stage, identifier, instanceName, preserveOrder)
}

func (u *ElisionUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Elision)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Smufl":
		instance.Smufl = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type EmptyUnmarshaller struct{}

func (u *EmptyUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Empty), stage, identifier, instanceName, preserveOrder)
}

func (u *EmptyUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Empty)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type Empty_fontUnmarshaller struct{}

func (u *Empty_fontUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Empty_font), stage, identifier, instanceName, preserveOrder)
}

func (u *Empty_fontUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Empty_font)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	}
	return nil
}

type Empty_lineUnmarshaller struct{}

func (u *Empty_lineUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Empty_line), stage, identifier, instanceName, preserveOrder)
}

func (u *Empty_lineUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Empty_line)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Line_shape":
		GongUnmarshallEnum(&instance.Line_shape, valueExpr)
	case "Line_type":
		GongUnmarshallEnum(&instance.Line_type, valueExpr)
	case "Line_length":
		GongUnmarshallEnum(&instance.Line_length, valueExpr)
	case "Dash_length":
		instance.Dash_length = GongExtractString(valueExpr)
	case "Space_length":
		instance.Space_length = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Placement":
		GongUnmarshallEnum(&instance.Placement, valueExpr)
	}
	return nil
}

type Empty_placementUnmarshaller struct{}

func (u *Empty_placementUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Empty_placement), stage, identifier, instanceName, preserveOrder)
}

func (u *Empty_placementUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Empty_placement)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Placement":
		GongUnmarshallEnum(&instance.Placement, valueExpr)
	}
	return nil
}

type Empty_placement_smuflUnmarshaller struct{}

func (u *Empty_placement_smuflUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Empty_placement_smufl), stage, identifier, instanceName, preserveOrder)
}

func (u *Empty_placement_smuflUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Empty_placement_smufl)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Placement":
		GongUnmarshallEnum(&instance.Placement, valueExpr)
	case "Smufl":
		instance.Smufl = GongExtractString(valueExpr)
	}
	return nil
}

type Empty_print_object_style_alignUnmarshaller struct{}

func (u *Empty_print_object_style_alignUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Empty_print_object_style_align), stage, identifier, instanceName, preserveOrder)
}

func (u *Empty_print_object_style_alignUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Empty_print_object_style_align)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Print_object":
		GongUnmarshallEnum(&instance.Print_object, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Halign":
		GongUnmarshallEnum(&instance.Halign, valueExpr)
	case "Valign":
		GongUnmarshallEnum(&instance.Valign, valueExpr)
	}
	return nil
}

type Empty_print_styleUnmarshaller struct{}

func (u *Empty_print_styleUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Empty_print_style), stage, identifier, instanceName, preserveOrder)
}

func (u *Empty_print_styleUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Empty_print_style)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	}
	return nil
}

type Empty_print_style_alignUnmarshaller struct{}

func (u *Empty_print_style_alignUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Empty_print_style_align), stage, identifier, instanceName, preserveOrder)
}

func (u *Empty_print_style_alignUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Empty_print_style_align)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Halign":
		GongUnmarshallEnum(&instance.Halign, valueExpr)
	case "Valign":
		GongUnmarshallEnum(&instance.Valign, valueExpr)
	}
	return nil
}

type Empty_print_style_align_idUnmarshaller struct{}

func (u *Empty_print_style_align_idUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Empty_print_style_align_id), stage, identifier, instanceName, preserveOrder)
}

func (u *Empty_print_style_align_idUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Empty_print_style_align_id)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Halign":
		GongUnmarshallEnum(&instance.Halign, valueExpr)
	case "Valign":
		GongUnmarshallEnum(&instance.Valign, valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	}
	return nil
}

type Empty_trill_soundUnmarshaller struct{}

func (u *Empty_trill_soundUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Empty_trill_sound), stage, identifier, instanceName, preserveOrder)
}

func (u *Empty_trill_soundUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Empty_trill_sound)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Placement":
		GongUnmarshallEnum(&instance.Placement, valueExpr)
	case "Start_note":
		GongUnmarshallEnum(&instance.Start_note, valueExpr)
	case "Trill_step":
		GongUnmarshallEnum(&instance.Trill_step, valueExpr)
	case "Two_note_turn":
		GongUnmarshallEnum(&instance.Two_note_turn, valueExpr)
	case "Accelerate":
		GongUnmarshallEnum(&instance.Accelerate, valueExpr)
	case "Beats":
		instance.Beats = GongExtractString(valueExpr)
	case "Second_beat":
		instance.Second_beat = GongExtractString(valueExpr)
	case "Last_beat":
		instance.Last_beat = GongExtractString(valueExpr)
	}
	return nil
}

type EncodingUnmarshaller struct{}

func (u *EncodingUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Encoding), stage, identifier, instanceName, preserveOrder)
}

func (u *EncodingUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Encoding)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Encoding_date":
		instance.Encoding_date = GongExtractString(valueExpr)
	case "Encoder":
		GongUnmarshallSliceOfPointers(&instance.Encoder, valueExpr, identifierMap)
	case "Software":
		instance.Software = GongExtractString(valueExpr)
	case "Encoding_description":
		instance.Encoding_description = GongExtractString(valueExpr)
	case "Supports":
		GongUnmarshallSliceOfPointers(&instance.Supports, valueExpr, identifierMap)
	}
	return nil
}

type EndingUnmarshaller struct{}

func (u *EndingUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Ending), stage, identifier, instanceName, preserveOrder)
}

func (u *EndingUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Ending)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Number":
		instance.Number = GongExtractString(valueExpr)
	case "Type_XSD":
		GongUnmarshallEnum(&instance.Type_XSD, valueExpr)
	case "End_length":
		instance.End_length = GongExtractString(valueExpr)
	case "Text_x":
		instance.Text_x = GongExtractString(valueExpr)
	case "Text_y":
		instance.Text_y = GongExtractString(valueExpr)
	case "Print_object":
		GongUnmarshallEnum(&instance.Print_object, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "System":
		instance.System = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type ExtendUnmarshaller struct{}

func (u *ExtendUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Extend), stage, identifier, instanceName, preserveOrder)
}

func (u *ExtendUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Extend)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type_XSD":
		GongUnmarshallEnum(&instance.Type_XSD, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	}
	return nil
}

type FeatureUnmarshaller struct{}

func (u *FeatureUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Feature), stage, identifier, instanceName, preserveOrder)
}

func (u *FeatureUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Feature)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type_XSD":
		instance.Type_XSD = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type FermataUnmarshaller struct{}

func (u *FermataUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Fermata), stage, identifier, instanceName, preserveOrder)
}

func (u *FermataUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Fermata)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type_XSD":
		GongUnmarshallEnum(&instance.Type_XSD, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "EnclosedText":
		GongUnmarshallEnum(&instance.EnclosedText, valueExpr)
	}
	return nil
}

type FigureUnmarshaller struct{}

func (u *FigureUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Figure), stage, identifier, instanceName, preserveOrder)
}

func (u *FigureUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Figure)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Prefix":
		GongUnmarshallPointer(&instance.Prefix, valueExpr, identifierMap)
	case "Figure_number":
		GongUnmarshallPointer(&instance.Figure_number, valueExpr, identifierMap)
	case "Suffix":
		GongUnmarshallPointer(&instance.Suffix, valueExpr, identifierMap)
	case "Extend":
		GongUnmarshallPointer(&instance.Extend, valueExpr, identifierMap)
	case "Footnote":
		GongUnmarshallPointer(&instance.Footnote, valueExpr, identifierMap)
	case "Level":
		GongUnmarshallPointer(&instance.Level, valueExpr, identifierMap)
	}
	return nil
}

type Figured_bassUnmarshaller struct{}

func (u *Figured_bassUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Figured_bass), stage, identifier, instanceName, preserveOrder)
}

func (u *Figured_bassUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Figured_bass)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Parentheses":
		GongUnmarshallEnum(&instance.Parentheses, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Halign":
		GongUnmarshallEnum(&instance.Halign, valueExpr)
	case "Valign":
		GongUnmarshallEnum(&instance.Valign, valueExpr)
	case "Placement":
		GongUnmarshallEnum(&instance.Placement, valueExpr)
	case "Print_dot":
		GongUnmarshallEnum(&instance.Print_dot, valueExpr)
	case "Print_lyric":
		GongUnmarshallEnum(&instance.Print_lyric, valueExpr)
	case "Print_object":
		GongUnmarshallEnum(&instance.Print_object, valueExpr)
	case "Print_spacing":
		GongUnmarshallEnum(&instance.Print_spacing, valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Figure":
		GongUnmarshallSliceOfPointers(&instance.Figure, valueExpr, identifierMap)
	case "Footnote":
		GongUnmarshallPointer(&instance.Footnote, valueExpr, identifierMap)
	case "Level":
		GongUnmarshallPointer(&instance.Level, valueExpr, identifierMap)
	}
	return nil
}

type FingeringUnmarshaller struct{}

func (u *FingeringUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Fingering), stage, identifier, instanceName, preserveOrder)
}

func (u *FingeringUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Fingering)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Substitution":
		GongUnmarshallEnum(&instance.Substitution, valueExpr)
	case "Alternate":
		GongUnmarshallEnum(&instance.Alternate, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Placement":
		GongUnmarshallEnum(&instance.Placement, valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type First_fretUnmarshaller struct{}

func (u *First_fretUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(First_fret), stage, identifier, instanceName, preserveOrder)
}

func (u *First_fretUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*First_fret)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Text":
		instance.Text = GongExtractString(valueExpr)
	case "Location":
		GongUnmarshallEnum(&instance.Location, valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractInt(valueExpr)
	}
	return nil
}

type For_partUnmarshaller struct{}

func (u *For_partUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(For_part), stage, identifier, instanceName, preserveOrder)
}

func (u *For_partUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*For_part)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Part_clef":
		GongUnmarshallPointer(&instance.Part_clef, valueExpr, identifierMap)
	case "Part_transpose":
		GongUnmarshallPointer(&instance.Part_transpose, valueExpr, identifierMap)
	}
	return nil
}

type Formatted_symbolUnmarshaller struct{}

func (u *Formatted_symbolUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Formatted_symbol), stage, identifier, instanceName, preserveOrder)
}

func (u *Formatted_symbolUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Formatted_symbol)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Justify":
		GongUnmarshallEnum(&instance.Justify, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Halign":
		GongUnmarshallEnum(&instance.Halign, valueExpr)
	case "Valign":
		GongUnmarshallEnum(&instance.Valign, valueExpr)
	case "Underline":
		instance.Underline = GongExtractInt(valueExpr)
	case "Overline":
		instance.Overline = GongExtractInt(valueExpr)
	case "Line_through":
		instance.Line_through = GongExtractInt(valueExpr)
	case "Rotation":
		instance.Rotation = GongExtractString(valueExpr)
	case "Letter_spacing":
		instance.Letter_spacing = GongExtractString(valueExpr)
	case "Line_height":
		instance.Line_height = GongExtractString(valueExpr)
	case "Dir":
		GongUnmarshallEnum(&instance.Dir, valueExpr)
	case "Enclosure":
		GongUnmarshallEnum(&instance.Enclosure, valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type Formatted_symbol_idUnmarshaller struct{}

func (u *Formatted_symbol_idUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Formatted_symbol_id), stage, identifier, instanceName, preserveOrder)
}

func (u *Formatted_symbol_idUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Formatted_symbol_id)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Justify":
		GongUnmarshallEnum(&instance.Justify, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Halign":
		GongUnmarshallEnum(&instance.Halign, valueExpr)
	case "Valign":
		GongUnmarshallEnum(&instance.Valign, valueExpr)
	case "Underline":
		instance.Underline = GongExtractInt(valueExpr)
	case "Overline":
		instance.Overline = GongExtractInt(valueExpr)
	case "Line_through":
		instance.Line_through = GongExtractInt(valueExpr)
	case "Rotation":
		instance.Rotation = GongExtractString(valueExpr)
	case "Letter_spacing":
		instance.Letter_spacing = GongExtractString(valueExpr)
	case "Line_height":
		instance.Line_height = GongExtractString(valueExpr)
	case "Dir":
		GongUnmarshallEnum(&instance.Dir, valueExpr)
	case "Enclosure":
		GongUnmarshallEnum(&instance.Enclosure, valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type Formatted_textUnmarshaller struct{}

func (u *Formatted_textUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Formatted_text), stage, identifier, instanceName, preserveOrder)
}

func (u *Formatted_textUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Formatted_text)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Lang":
		instance.Lang = GongExtractString(valueExpr)
	case "Space":
		instance.Space = GongExtractString(valueExpr)
	case "Justify":
		GongUnmarshallEnum(&instance.Justify, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Halign":
		GongUnmarshallEnum(&instance.Halign, valueExpr)
	case "Valign":
		GongUnmarshallEnum(&instance.Valign, valueExpr)
	case "Underline":
		instance.Underline = GongExtractInt(valueExpr)
	case "Overline":
		instance.Overline = GongExtractInt(valueExpr)
	case "Line_through":
		instance.Line_through = GongExtractInt(valueExpr)
	case "Rotation":
		instance.Rotation = GongExtractString(valueExpr)
	case "Letter_spacing":
		instance.Letter_spacing = GongExtractString(valueExpr)
	case "Line_height":
		instance.Line_height = GongExtractString(valueExpr)
	case "Dir":
		GongUnmarshallEnum(&instance.Dir, valueExpr)
	case "Enclosure":
		GongUnmarshallEnum(&instance.Enclosure, valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type Formatted_text_idUnmarshaller struct{}

func (u *Formatted_text_idUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Formatted_text_id), stage, identifier, instanceName, preserveOrder)
}

func (u *Formatted_text_idUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Formatted_text_id)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Lang":
		instance.Lang = GongExtractString(valueExpr)
	case "Space":
		instance.Space = GongExtractString(valueExpr)
	case "Justify":
		GongUnmarshallEnum(&instance.Justify, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Halign":
		GongUnmarshallEnum(&instance.Halign, valueExpr)
	case "Valign":
		GongUnmarshallEnum(&instance.Valign, valueExpr)
	case "Underline":
		instance.Underline = GongExtractInt(valueExpr)
	case "Overline":
		instance.Overline = GongExtractInt(valueExpr)
	case "Line_through":
		instance.Line_through = GongExtractInt(valueExpr)
	case "Rotation":
		instance.Rotation = GongExtractString(valueExpr)
	case "Letter_spacing":
		instance.Letter_spacing = GongExtractString(valueExpr)
	case "Line_height":
		instance.Line_height = GongExtractString(valueExpr)
	case "Dir":
		GongUnmarshallEnum(&instance.Dir, valueExpr)
	case "Enclosure":
		GongUnmarshallEnum(&instance.Enclosure, valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type ForwardUnmarshaller struct{}

func (u *ForwardUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Forward), stage, identifier, instanceName, preserveOrder)
}

func (u *ForwardUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Forward)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Footnote":
		GongUnmarshallPointer(&instance.Footnote, valueExpr, identifierMap)
	case "Level":
		GongUnmarshallPointer(&instance.Level, valueExpr, identifierMap)
	case "Voice":
		instance.Voice = GongExtractString(valueExpr)
	case "Staff":
		instance.Staff = GongExtractInt(valueExpr)
	}
	return nil
}

type FrameUnmarshaller struct{}

func (u *FrameUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Frame), stage, identifier, instanceName, preserveOrder)
}

func (u *FrameUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Frame)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Height":
		instance.Height = GongExtractString(valueExpr)
	case "Width":
		instance.Width = GongExtractString(valueExpr)
	case "Unplayed":
		instance.Unplayed = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Halign":
		GongUnmarshallEnum(&instance.Halign, valueExpr)
	case "Valign":
		GongUnmarshallEnum(&instance.Valign, valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Frame_strings":
		instance.Frame_strings = GongExtractInt(valueExpr)
	case "Frame_frets":
		instance.Frame_frets = GongExtractInt(valueExpr)
	case "First_fret":
		GongUnmarshallPointer(&instance.First_fret, valueExpr, identifierMap)
	case "Frame_note":
		GongUnmarshallSliceOfPointers(&instance.Frame_note, valueExpr, identifierMap)
	}
	return nil
}

type Frame_noteUnmarshaller struct{}

func (u *Frame_noteUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Frame_note), stage, identifier, instanceName, preserveOrder)
}

func (u *Frame_noteUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Frame_note)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "String_XSD":
		GongUnmarshallPointer(&instance.String_XSD, valueExpr, identifierMap)
	case "Fret":
		GongUnmarshallPointer(&instance.Fret, valueExpr, identifierMap)
	case "Fingering":
		GongUnmarshallPointer(&instance.Fingering, valueExpr, identifierMap)
	case "Barre":
		GongUnmarshallPointer(&instance.Barre, valueExpr, identifierMap)
	}
	return nil
}

type FretUnmarshaller struct{}

func (u *FretUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Fret), stage, identifier, instanceName, preserveOrder)
}

func (u *FretUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Fret)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractInt(valueExpr)
	}
	return nil
}

type GlassUnmarshaller struct{}

func (u *GlassUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Glass), stage, identifier, instanceName, preserveOrder)
}

func (u *GlassUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Glass)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Smufl":
		instance.Smufl = GongExtractString(valueExpr)
	case "EnclosedText":
		GongUnmarshallEnum(&instance.EnclosedText, valueExpr)
	}
	return nil
}

type GlissandoUnmarshaller struct{}

func (u *GlissandoUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Glissando), stage, identifier, instanceName, preserveOrder)
}

func (u *GlissandoUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Glissando)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type_XSD":
		GongUnmarshallEnum(&instance.Type_XSD, valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "Line_type":
		GongUnmarshallEnum(&instance.Line_type, valueExpr)
	case "Dash_length":
		instance.Dash_length = GongExtractString(valueExpr)
	case "Space_length":
		instance.Space_length = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type GlyphUnmarshaller struct{}

func (u *GlyphUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Glyph), stage, identifier, instanceName, preserveOrder)
}

func (u *GlyphUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Glyph)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type_XSD":
		instance.Type_XSD = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type GraceUnmarshaller struct{}

func (u *GraceUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Grace), stage, identifier, instanceName, preserveOrder)
}

func (u *GraceUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Grace)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Steal_time_previous":
		instance.Steal_time_previous = GongExtractString(valueExpr)
	case "Steal_time_following":
		instance.Steal_time_following = GongExtractString(valueExpr)
	case "Make_time":
		instance.Make_time = GongExtractString(valueExpr)
	case "Slash":
		GongUnmarshallEnum(&instance.Slash, valueExpr)
	}
	return nil
}

type Group_barlineUnmarshaller struct{}

func (u *Group_barlineUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Group_barline), stage, identifier, instanceName, preserveOrder)
}

func (u *Group_barlineUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Group_barline)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "EnclosedText":
		GongUnmarshallEnum(&instance.EnclosedText, valueExpr)
	}
	return nil
}

type Group_nameUnmarshaller struct{}

func (u *Group_nameUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Group_name), stage, identifier, instanceName, preserveOrder)
}

func (u *Group_nameUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Group_name)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Justify":
		GongUnmarshallEnum(&instance.Justify, valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type Group_symbolUnmarshaller struct{}

func (u *Group_symbolUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Group_symbol), stage, identifier, instanceName, preserveOrder)
}

func (u *Group_symbolUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Group_symbol)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "EnclosedText":
		GongUnmarshallEnum(&instance.EnclosedText, valueExpr)
	}
	return nil
}

type GroupingUnmarshaller struct{}

func (u *GroupingUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Grouping), stage, identifier, instanceName, preserveOrder)
}

func (u *GroupingUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Grouping)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type_XSD":
		GongUnmarshallEnum(&instance.Type_XSD, valueExpr)
	case "Number":
		instance.Number = GongExtractString(valueExpr)
	case "Member_of":
		instance.Member_of = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Feature":
		GongUnmarshallSliceOfPointers(&instance.Feature, valueExpr, identifierMap)
	}
	return nil
}

type Hammer_on_pull_offUnmarshaller struct{}

func (u *Hammer_on_pull_offUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Hammer_on_pull_off), stage, identifier, instanceName, preserveOrder)
}

func (u *Hammer_on_pull_offUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Hammer_on_pull_off)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type_XSD":
		GongUnmarshallEnum(&instance.Type_XSD, valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Placement":
		GongUnmarshallEnum(&instance.Placement, valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type HandbellUnmarshaller struct{}

func (u *HandbellUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Handbell), stage, identifier, instanceName, preserveOrder)
}

func (u *HandbellUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Handbell)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Placement":
		GongUnmarshallEnum(&instance.Placement, valueExpr)
	case "EnclosedText":
		GongUnmarshallEnum(&instance.EnclosedText, valueExpr)
	}
	return nil
}

type Harmon_closedUnmarshaller struct{}

func (u *Harmon_closedUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Harmon_closed), stage, identifier, instanceName, preserveOrder)
}

func (u *Harmon_closedUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Harmon_closed)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Location":
		GongUnmarshallEnum(&instance.Location, valueExpr)
	case "EnclosedText":
		GongUnmarshallEnum(&instance.EnclosedText, valueExpr)
	}
	return nil
}

type Harmon_muteUnmarshaller struct{}

func (u *Harmon_muteUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Harmon_mute), stage, identifier, instanceName, preserveOrder)
}

func (u *Harmon_muteUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Harmon_mute)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Placement":
		GongUnmarshallEnum(&instance.Placement, valueExpr)
	case "Harmon_closed":
		GongUnmarshallPointer(&instance.Harmon_closed, valueExpr, identifierMap)
	}
	return nil
}

type HarmonicUnmarshaller struct{}

func (u *HarmonicUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Harmonic), stage, identifier, instanceName, preserveOrder)
}

func (u *HarmonicUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Harmonic)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Print_object":
		GongUnmarshallEnum(&instance.Print_object, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Placement":
		GongUnmarshallEnum(&instance.Placement, valueExpr)
	case "Natural":
		instance.Natural = GongExtractString(valueExpr)
	case "Artificial":
		instance.Artificial = GongExtractString(valueExpr)
	case "Base_pitch":
		instance.Base_pitch = GongExtractString(valueExpr)
	case "Touching_pitch":
		instance.Touching_pitch = GongExtractString(valueExpr)
	case "Sounding_pitch":
		instance.Sounding_pitch = GongExtractString(valueExpr)
	}
	return nil
}

type HarmonyUnmarshaller struct{}

func (u *HarmonyUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Harmony), stage, identifier, instanceName, preserveOrder)
}

func (u *HarmonyUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Harmony)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type_XSD":
		GongUnmarshallEnum(&instance.Type_XSD, valueExpr)
	case "Print_frame":
		GongUnmarshallEnum(&instance.Print_frame, valueExpr)
	case "Arrangement":
		GongUnmarshallEnum(&instance.Arrangement, valueExpr)
	case "Print_object":
		GongUnmarshallEnum(&instance.Print_object, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Placement":
		GongUnmarshallEnum(&instance.Placement, valueExpr)
	case "System":
		instance.System = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Root":
		GongUnmarshallPointer(&instance.Root, valueExpr, identifierMap)
	case "Numeral":
		GongUnmarshallPointer(&instance.Numeral, valueExpr, identifierMap)
	case "Function":
		GongUnmarshallPointer(&instance.Function, valueExpr, identifierMap)
	case "Kind":
		GongUnmarshallPointer(&instance.Kind, valueExpr, identifierMap)
	case "Inversion":
		GongUnmarshallPointer(&instance.Inversion, valueExpr, identifierMap)
	case "Bass":
		GongUnmarshallPointer(&instance.Bass, valueExpr, identifierMap)
	case "Degree":
		GongUnmarshallSliceOfPointers(&instance.Degree, valueExpr, identifierMap)
	case "Frame":
		GongUnmarshallPointer(&instance.Frame, valueExpr, identifierMap)
	case "Offset":
		GongUnmarshallPointer(&instance.Offset, valueExpr, identifierMap)
	case "Footnote":
		GongUnmarshallPointer(&instance.Footnote, valueExpr, identifierMap)
	case "Level":
		GongUnmarshallPointer(&instance.Level, valueExpr, identifierMap)
	case "Staff":
		instance.Staff = GongExtractInt(valueExpr)
	}
	return nil
}

type Harmony_alterUnmarshaller struct{}

func (u *Harmony_alterUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Harmony_alter), stage, identifier, instanceName, preserveOrder)
}

func (u *Harmony_alterUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Harmony_alter)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Location":
		GongUnmarshallEnum(&instance.Location, valueExpr)
	case "Print_object":
		GongUnmarshallEnum(&instance.Print_object, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type Harp_pedalsUnmarshaller struct{}

func (u *Harp_pedalsUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Harp_pedals), stage, identifier, instanceName, preserveOrder)
}

func (u *Harp_pedalsUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Harp_pedals)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Halign":
		GongUnmarshallEnum(&instance.Halign, valueExpr)
	case "Valign":
		GongUnmarshallEnum(&instance.Valign, valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Pedal_tuning":
		GongUnmarshallSliceOfPointers(&instance.Pedal_tuning, valueExpr, identifierMap)
	}
	return nil
}

type Heel_toeUnmarshaller struct{}

func (u *Heel_toeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Heel_toe), stage, identifier, instanceName, preserveOrder)
}

func (u *Heel_toeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Heel_toe)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Placement":
		GongUnmarshallEnum(&instance.Placement, valueExpr)
	case "Substitution":
		GongUnmarshallEnum(&instance.Substitution, valueExpr)
	}
	return nil
}

type HoleUnmarshaller struct{}

func (u *HoleUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Hole), stage, identifier, instanceName, preserveOrder)
}

func (u *HoleUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Hole)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Placement":
		GongUnmarshallEnum(&instance.Placement, valueExpr)
	case "Hole_type":
		instance.Hole_type = GongExtractString(valueExpr)
	case "Hole_closed":
		GongUnmarshallPointer(&instance.Hole_closed, valueExpr, identifierMap)
	case "Hole_shape":
		instance.Hole_shape = GongExtractString(valueExpr)
	}
	return nil
}

type Hole_closedUnmarshaller struct{}

func (u *Hole_closedUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Hole_closed), stage, identifier, instanceName, preserveOrder)
}

func (u *Hole_closedUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Hole_closed)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Location":
		GongUnmarshallEnum(&instance.Location, valueExpr)
	case "EnclosedText":
		GongUnmarshallEnum(&instance.EnclosedText, valueExpr)
	}
	return nil
}

type Horizontal_turnUnmarshaller struct{}

func (u *Horizontal_turnUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Horizontal_turn), stage, identifier, instanceName, preserveOrder)
}

func (u *Horizontal_turnUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Horizontal_turn)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Slash":
		GongUnmarshallEnum(&instance.Slash, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Placement":
		GongUnmarshallEnum(&instance.Placement, valueExpr)
	case "Start_note":
		GongUnmarshallEnum(&instance.Start_note, valueExpr)
	case "Trill_step":
		GongUnmarshallEnum(&instance.Trill_step, valueExpr)
	case "Two_note_turn":
		GongUnmarshallEnum(&instance.Two_note_turn, valueExpr)
	case "Accelerate":
		GongUnmarshallEnum(&instance.Accelerate, valueExpr)
	case "Beats":
		instance.Beats = GongExtractString(valueExpr)
	case "Second_beat":
		instance.Second_beat = GongExtractString(valueExpr)
	case "Last_beat":
		instance.Last_beat = GongExtractString(valueExpr)
	}
	return nil
}

type IdentificationUnmarshaller struct{}

func (u *IdentificationUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Identification), stage, identifier, instanceName, preserveOrder)
}

func (u *IdentificationUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Identification)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Creator":
		GongUnmarshallSliceOfPointers(&instance.Creator, valueExpr, identifierMap)
	case "Rights":
		GongUnmarshallSliceOfPointers(&instance.Rights, valueExpr, identifierMap)
	case "Encoding":
		GongUnmarshallPointer(&instance.Encoding, valueExpr, identifierMap)
	case "Source":
		instance.Source = GongExtractString(valueExpr)
	case "Relation":
		GongUnmarshallSliceOfPointers(&instance.Relation, valueExpr, identifierMap)
	case "Miscellaneous":
		GongUnmarshallPointer(&instance.Miscellaneous, valueExpr, identifierMap)
	}
	return nil
}

type ImageUnmarshaller struct{}

func (u *ImageUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Image), stage, identifier, instanceName, preserveOrder)
}

func (u *ImageUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Image)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Source":
		instance.Source = GongExtractString(valueExpr)
	case "Type_XSD":
		instance.Type_XSD = GongExtractString(valueExpr)
	case "Height":
		instance.Height = GongExtractString(valueExpr)
	case "Width":
		instance.Width = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Halign":
		GongUnmarshallEnum(&instance.Halign, valueExpr)
	case "Valign":
		GongUnmarshallEnum(&instance.Valign, valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	}
	return nil
}

type InstrumentUnmarshaller struct{}

func (u *InstrumentUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Instrument), stage, identifier, instanceName, preserveOrder)
}

func (u *InstrumentUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Instrument)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	}
	return nil
}

type Instrument_changeUnmarshaller struct{}

func (u *Instrument_changeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Instrument_change), stage, identifier, instanceName, preserveOrder)
}

func (u *Instrument_changeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Instrument_change)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Instrument_sound":
		instance.Instrument_sound = GongExtractString(valueExpr)
	case "Solo":
		instance.Solo = GongExtractString(valueExpr)
	case "Ensemble":
		instance.Ensemble = GongExtractString(valueExpr)
	case "Virtual_instrument":
		GongUnmarshallPointer(&instance.Virtual_instrument, valueExpr, identifierMap)
	}
	return nil
}

type Instrument_linkUnmarshaller struct{}

func (u *Instrument_linkUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Instrument_link), stage, identifier, instanceName, preserveOrder)
}

func (u *Instrument_linkUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Instrument_link)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	}
	return nil
}

type InterchangeableUnmarshaller struct{}

func (u *InterchangeableUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Interchangeable), stage, identifier, instanceName, preserveOrder)
}

func (u *InterchangeableUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Interchangeable)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Symbol":
		GongUnmarshallEnum(&instance.Symbol, valueExpr)
	case "Separator":
		GongUnmarshallEnum(&instance.Separator, valueExpr)
	case "Time_relation":
		GongUnmarshallEnum(&instance.Time_relation, valueExpr)
	case "Beats":
		instance.Beats = GongExtractString(valueExpr)
	case "Beat_type":
		instance.Beat_type = GongExtractString(valueExpr)
	}
	return nil
}

type InversionUnmarshaller struct{}

func (u *InversionUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Inversion), stage, identifier, instanceName, preserveOrder)
}

func (u *InversionUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Inversion)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Text":
		instance.Text = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractInt(valueExpr)
	}
	return nil
}

type KeyUnmarshaller struct{}

func (u *KeyUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Key), stage, identifier, instanceName, preserveOrder)
}

func (u *KeyUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Key)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Print_object":
		GongUnmarshallEnum(&instance.Print_object, valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Cancel":
		GongUnmarshallPointer(&instance.Cancel, valueExpr, identifierMap)
	case "Fifths":
		instance.Fifths = GongExtractInt(valueExpr)
	case "Mode":
		instance.Mode = GongExtractString(valueExpr)
	case "Key_step":
		GongUnmarshallEnum(&instance.Key_step, valueExpr)
	case "Key_alter":
		instance.Key_alter = GongExtractString(valueExpr)
	case "Key_accidental":
		GongUnmarshallPointer(&instance.Key_accidental, valueExpr, identifierMap)
	case "Key_octave":
		GongUnmarshallSliceOfPointers(&instance.Key_octave, valueExpr, identifierMap)
	}
	return nil
}

type Key_accidentalUnmarshaller struct{}

func (u *Key_accidentalUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Key_accidental), stage, identifier, instanceName, preserveOrder)
}

func (u *Key_accidentalUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Key_accidental)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Smufl":
		instance.Smufl = GongExtractString(valueExpr)
	case "EnclosedText":
		GongUnmarshallEnum(&instance.EnclosedText, valueExpr)
	}
	return nil
}

type Key_octaveUnmarshaller struct{}

func (u *Key_octaveUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Key_octave), stage, identifier, instanceName, preserveOrder)
}

func (u *Key_octaveUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Key_octave)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "Cancel":
		GongUnmarshallEnum(&instance.Cancel, valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractInt(valueExpr)
	}
	return nil
}

type KindUnmarshaller struct{}

func (u *KindUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Kind), stage, identifier, instanceName, preserveOrder)
}

func (u *KindUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Kind)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Use_symbols":
		GongUnmarshallEnum(&instance.Use_symbols, valueExpr)
	case "Text":
		instance.Text = GongExtractString(valueExpr)
	case "Stack_degrees":
		GongUnmarshallEnum(&instance.Stack_degrees, valueExpr)
	case "Parentheses_degrees":
		GongUnmarshallEnum(&instance.Parentheses_degrees, valueExpr)
	case "Bracket_degrees":
		GongUnmarshallEnum(&instance.Bracket_degrees, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Halign":
		GongUnmarshallEnum(&instance.Halign, valueExpr)
	case "Valign":
		GongUnmarshallEnum(&instance.Valign, valueExpr)
	case "EnclosedText":
		GongUnmarshallEnum(&instance.EnclosedText, valueExpr)
	}
	return nil
}

type LevelUnmarshaller struct{}

func (u *LevelUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Level), stage, identifier, instanceName, preserveOrder)
}

func (u *LevelUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Level)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Reference_XSD":
		GongUnmarshallEnum(&instance.Reference_XSD, valueExpr)
	case "Type_XSD":
		GongUnmarshallEnum(&instance.Type_XSD, valueExpr)
	case "Parentheses":
		GongUnmarshallEnum(&instance.Parentheses, valueExpr)
	case "Bracket":
		GongUnmarshallEnum(&instance.Bracket, valueExpr)
	case "Size":
		GongUnmarshallEnum(&instance.Size, valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type Line_detailUnmarshaller struct{}

func (u *Line_detailUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Line_detail), stage, identifier, instanceName, preserveOrder)
}

func (u *Line_detailUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Line_detail)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Line":
		instance.Line = GongExtractInt(valueExpr)
	case "Width":
		instance.Width = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Line_type":
		GongUnmarshallEnum(&instance.Line_type, valueExpr)
	case "Print_object":
		GongUnmarshallEnum(&instance.Print_object, valueExpr)
	}
	return nil
}

type Line_widthUnmarshaller struct{}

func (u *Line_widthUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Line_width), stage, identifier, instanceName, preserveOrder)
}

func (u *Line_widthUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Line_width)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type_XSD":
		instance.Type_XSD = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type LinkUnmarshaller struct{}

func (u *LinkUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Link), stage, identifier, instanceName, preserveOrder)
}

func (u *LinkUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Link)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NameXSD":
		instance.NameXSD = GongExtractString(valueExpr)
	case "Href":
		instance.Href = GongExtractString(valueExpr)
	case "Type_XSD":
		instance.Type_XSD = GongExtractString(valueExpr)
	case "Role":
		instance.Role = GongExtractString(valueExpr)
	case "Title":
		instance.Title = GongExtractString(valueExpr)
	case "Show":
		instance.Show = GongExtractString(valueExpr)
	case "Actuate":
		instance.Actuate = GongExtractString(valueExpr)
	case "Element":
		instance.Element = GongExtractString(valueExpr)
	case "Position":
		instance.Position = GongExtractInt(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	}
	return nil
}

type ListenUnmarshaller struct{}

func (u *ListenUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Listen), stage, identifier, instanceName, preserveOrder)
}

func (u *ListenUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Listen)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Assess":
		GongUnmarshallSliceOfPointers(&instance.Assess, valueExpr, identifierMap)
	case "Wait":
		GongUnmarshallSliceOfPointers(&instance.Wait, valueExpr, identifierMap)
	case "Other_listen":
		GongUnmarshallSliceOfPointers(&instance.Other_listen, valueExpr, identifierMap)
	}
	return nil
}

type ListeningUnmarshaller struct{}

func (u *ListeningUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Listening), stage, identifier, instanceName, preserveOrder)
}

func (u *ListeningUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Listening)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Sync":
		GongUnmarshallSliceOfPointers(&instance.Sync, valueExpr, identifierMap)
	case "Other_listening":
		GongUnmarshallSliceOfPointers(&instance.Other_listening, valueExpr, identifierMap)
	case "Offset":
		GongUnmarshallPointer(&instance.Offset, valueExpr, identifierMap)
	}
	return nil
}

type LyricUnmarshaller struct{}

func (u *LyricUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Lyric), stage, identifier, instanceName, preserveOrder)
}

func (u *LyricUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Lyric)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Number":
		instance.Number = GongExtractString(valueExpr)
	case "NameXSD":
		instance.NameXSD = GongExtractString(valueExpr)
	case "Time_only":
		instance.Time_only = GongExtractString(valueExpr)
	case "Justify":
		GongUnmarshallEnum(&instance.Justify, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Placement":
		GongUnmarshallEnum(&instance.Placement, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Print_object":
		GongUnmarshallEnum(&instance.Print_object, valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Syllabic":
		GongUnmarshallEnum(&instance.Syllabic, valueExpr)
	case "Text":
		GongUnmarshallPointer(&instance.Text, valueExpr, identifierMap)
	case "Elision":
		GongUnmarshallSliceOfPointers(&instance.Elision, valueExpr, identifierMap)
	case "Extend":
		GongUnmarshallPointer(&instance.Extend, valueExpr, identifierMap)
	case "Laughing":
		instance.Laughing = GongExtractString(valueExpr)
	case "Humming":
		instance.Humming = GongExtractString(valueExpr)
	case "End_line":
		instance.End_line = GongExtractString(valueExpr)
	case "End_paragraph":
		instance.End_paragraph = GongExtractString(valueExpr)
	case "Footnote":
		GongUnmarshallPointer(&instance.Footnote, valueExpr, identifierMap)
	case "Level":
		GongUnmarshallPointer(&instance.Level, valueExpr, identifierMap)
	}
	return nil
}

type Lyric_fontUnmarshaller struct{}

func (u *Lyric_fontUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Lyric_font), stage, identifier, instanceName, preserveOrder)
}

func (u *Lyric_fontUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Lyric_font)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Number":
		instance.Number = GongExtractString(valueExpr)
	case "NameXSD":
		instance.NameXSD = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	}
	return nil
}

type Lyric_languageUnmarshaller struct{}

func (u *Lyric_languageUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Lyric_language), stage, identifier, instanceName, preserveOrder)
}

func (u *Lyric_languageUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Lyric_language)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Number":
		instance.Number = GongExtractString(valueExpr)
	case "NameXSD":
		instance.NameXSD = GongExtractString(valueExpr)
	case "Lang":
		instance.Lang = GongExtractString(valueExpr)
	}
	return nil
}

type Measure_layoutUnmarshaller struct{}

func (u *Measure_layoutUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Measure_layout), stage, identifier, instanceName, preserveOrder)
}

func (u *Measure_layoutUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Measure_layout)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Measure_distance":
		instance.Measure_distance = GongExtractString(valueExpr)
	}
	return nil
}

type Measure_numberingUnmarshaller struct{}

func (u *Measure_numberingUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Measure_numbering), stage, identifier, instanceName, preserveOrder)
}

func (u *Measure_numberingUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Measure_numbering)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "System":
		GongUnmarshallEnum(&instance.System, valueExpr)
	case "Staff":
		instance.Staff = GongExtractInt(valueExpr)
	case "Multiple_rest_always":
		GongUnmarshallEnum(&instance.Multiple_rest_always, valueExpr)
	case "Multiple_rest_range":
		GongUnmarshallEnum(&instance.Multiple_rest_range, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Halign":
		GongUnmarshallEnum(&instance.Halign, valueExpr)
	case "Valign":
		GongUnmarshallEnum(&instance.Valign, valueExpr)
	case "EnclosedText":
		GongUnmarshallEnum(&instance.EnclosedText, valueExpr)
	}
	return nil
}

type Measure_repeatUnmarshaller struct{}

func (u *Measure_repeatUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Measure_repeat), stage, identifier, instanceName, preserveOrder)
}

func (u *Measure_repeatUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Measure_repeat)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type_XSD":
		GongUnmarshallEnum(&instance.Type_XSD, valueExpr)
	case "Slashes":
		instance.Slashes = GongExtractInt(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type Measure_styleUnmarshaller struct{}

func (u *Measure_styleUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Measure_style), stage, identifier, instanceName, preserveOrder)
}

func (u *Measure_styleUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Measure_style)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Multiple_rest":
		GongUnmarshallPointer(&instance.Multiple_rest, valueExpr, identifierMap)
	case "Measure_repeat":
		GongUnmarshallPointer(&instance.Measure_repeat, valueExpr, identifierMap)
	case "Beat_repeat":
		GongUnmarshallPointer(&instance.Beat_repeat, valueExpr, identifierMap)
	case "Slash":
		GongUnmarshallPointer(&instance.Slash, valueExpr, identifierMap)
	}
	return nil
}

type MembraneUnmarshaller struct{}

func (u *MembraneUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Membrane), stage, identifier, instanceName, preserveOrder)
}

func (u *MembraneUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Membrane)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Smufl":
		instance.Smufl = GongExtractString(valueExpr)
	case "EnclosedText":
		GongUnmarshallEnum(&instance.EnclosedText, valueExpr)
	}
	return nil
}

type MetalUnmarshaller struct{}

func (u *MetalUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Metal), stage, identifier, instanceName, preserveOrder)
}

func (u *MetalUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Metal)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Smufl":
		instance.Smufl = GongExtractString(valueExpr)
	case "EnclosedText":
		GongUnmarshallEnum(&instance.EnclosedText, valueExpr)
	}
	return nil
}

type MetronomeUnmarshaller struct{}

func (u *MetronomeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Metronome), stage, identifier, instanceName, preserveOrder)
}

func (u *MetronomeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Metronome)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Parentheses":
		GongUnmarshallEnum(&instance.Parentheses, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Halign":
		GongUnmarshallEnum(&instance.Halign, valueExpr)
	case "Valign":
		GongUnmarshallEnum(&instance.Valign, valueExpr)
	case "Print_object":
		GongUnmarshallEnum(&instance.Print_object, valueExpr)
	case "Justify":
		GongUnmarshallEnum(&instance.Justify, valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Beat_unit":
		GongUnmarshallEnum(&instance.Beat_unit, valueExpr)
	case "Beat_unit_dot":
		instance.Beat_unit_dot = GongExtractString(valueExpr)
	case "Beat_unit_tied":
		GongUnmarshallSliceOfPointers(&instance.Beat_unit_tied, valueExpr, identifierMap)
	case "Per_minute":
		GongUnmarshallPointer(&instance.Per_minute, valueExpr, identifierMap)
	case "Metronome_arrows":
		instance.Metronome_arrows = GongExtractString(valueExpr)
	case "Metronome_note":
		GongUnmarshallSliceOfPointers(&instance.Metronome_note, valueExpr, identifierMap)
	case "Metronome_relation":
		instance.Metronome_relation = GongExtractString(valueExpr)
	}
	return nil
}

type Metronome_beamUnmarshaller struct{}

func (u *Metronome_beamUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Metronome_beam), stage, identifier, instanceName, preserveOrder)
}

func (u *Metronome_beamUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Metronome_beam)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "EnclosedText":
		GongUnmarshallEnum(&instance.EnclosedText, valueExpr)
	}
	return nil
}

type Metronome_noteUnmarshaller struct{}

func (u *Metronome_noteUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Metronome_note), stage, identifier, instanceName, preserveOrder)
}

func (u *Metronome_noteUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Metronome_note)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Metronome_type":
		GongUnmarshallEnum(&instance.Metronome_type, valueExpr)
	case "Metronome_dot":
		instance.Metronome_dot = GongExtractString(valueExpr)
	case "Metronome_beam":
		GongUnmarshallSliceOfPointers(&instance.Metronome_beam, valueExpr, identifierMap)
	case "Metronome_tied":
		GongUnmarshallPointer(&instance.Metronome_tied, valueExpr, identifierMap)
	case "Metronome_tuplet":
		GongUnmarshallPointer(&instance.Metronome_tuplet, valueExpr, identifierMap)
	}
	return nil
}

type Metronome_tiedUnmarshaller struct{}

func (u *Metronome_tiedUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Metronome_tied), stage, identifier, instanceName, preserveOrder)
}

func (u *Metronome_tiedUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Metronome_tied)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type_XSD":
		GongUnmarshallEnum(&instance.Type_XSD, valueExpr)
	}
	return nil
}

type Metronome_tupletUnmarshaller struct{}

func (u *Metronome_tupletUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Metronome_tuplet), stage, identifier, instanceName, preserveOrder)
}

func (u *Metronome_tupletUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Metronome_tuplet)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Actual_notes":
		instance.Actual_notes = GongExtractInt(valueExpr)
	case "Normal_notes":
		instance.Normal_notes = GongExtractInt(valueExpr)
	case "Normal_type":
		GongUnmarshallEnum(&instance.Normal_type, valueExpr)
	case "Normal_dot":
		instance.Normal_dot = GongExtractString(valueExpr)
	case "Type_XSD":
		GongUnmarshallEnum(&instance.Type_XSD, valueExpr)
	case "Bracket":
		GongUnmarshallEnum(&instance.Bracket, valueExpr)
	case "Show_number":
		GongUnmarshallEnum(&instance.Show_number, valueExpr)
	}
	return nil
}

type Midi_deviceUnmarshaller struct{}

func (u *Midi_deviceUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Midi_device), stage, identifier, instanceName, preserveOrder)
}

func (u *Midi_deviceUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Midi_device)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Port":
		instance.Port = GongExtractInt(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type Midi_instrumentUnmarshaller struct{}

func (u *Midi_instrumentUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Midi_instrument), stage, identifier, instanceName, preserveOrder)
}

func (u *Midi_instrumentUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Midi_instrument)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Midi_channel":
		instance.Midi_channel = GongExtractInt(valueExpr)
	case "Midi_name":
		instance.Midi_name = GongExtractString(valueExpr)
	case "Midi_bank":
		instance.Midi_bank = GongExtractInt(valueExpr)
	case "Midi_program":
		instance.Midi_program = GongExtractInt(valueExpr)
	case "Midi_unpitched":
		instance.Midi_unpitched = GongExtractInt(valueExpr)
	case "Volume":
		instance.Volume = GongExtractString(valueExpr)
	case "Pan":
		instance.Pan = GongExtractString(valueExpr)
	case "Elevation":
		instance.Elevation = GongExtractString(valueExpr)
	}
	return nil
}

type MiscellaneousUnmarshaller struct{}

func (u *MiscellaneousUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Miscellaneous), stage, identifier, instanceName, preserveOrder)
}

func (u *MiscellaneousUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Miscellaneous)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Miscellaneous_field":
		GongUnmarshallSliceOfPointers(&instance.Miscellaneous_field, valueExpr, identifierMap)
	}
	return nil
}

type Miscellaneous_fieldUnmarshaller struct{}

func (u *Miscellaneous_fieldUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Miscellaneous_field), stage, identifier, instanceName, preserveOrder)
}

func (u *Miscellaneous_fieldUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Miscellaneous_field)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NameXSD":
		instance.NameXSD = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type MordentUnmarshaller struct{}

func (u *MordentUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Mordent), stage, identifier, instanceName, preserveOrder)
}

func (u *MordentUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Mordent)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Placement":
		GongUnmarshallEnum(&instance.Placement, valueExpr)
	case "Start_note":
		GongUnmarshallEnum(&instance.Start_note, valueExpr)
	case "Trill_step":
		GongUnmarshallEnum(&instance.Trill_step, valueExpr)
	case "Two_note_turn":
		GongUnmarshallEnum(&instance.Two_note_turn, valueExpr)
	case "Accelerate":
		GongUnmarshallEnum(&instance.Accelerate, valueExpr)
	case "Beats":
		instance.Beats = GongExtractString(valueExpr)
	case "Second_beat":
		instance.Second_beat = GongExtractString(valueExpr)
	case "Last_beat":
		instance.Last_beat = GongExtractString(valueExpr)
	case "Long":
		GongUnmarshallEnum(&instance.Long, valueExpr)
	case "Approach":
		GongUnmarshallEnum(&instance.Approach, valueExpr)
	case "Departure":
		GongUnmarshallEnum(&instance.Departure, valueExpr)
	}
	return nil
}

type Multiple_restUnmarshaller struct{}

func (u *Multiple_restUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Multiple_rest), stage, identifier, instanceName, preserveOrder)
}

func (u *Multiple_restUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Multiple_rest)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Use_symbols":
		GongUnmarshallEnum(&instance.Use_symbols, valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractInt(valueExpr)
	}
	return nil
}

type Name_displayUnmarshaller struct{}

func (u *Name_displayUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Name_display), stage, identifier, instanceName, preserveOrder)
}

func (u *Name_displayUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Name_display)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Print_object":
		GongUnmarshallEnum(&instance.Print_object, valueExpr)
	case "Display_text":
		GongUnmarshallSliceOfPointers(&instance.Display_text, valueExpr, identifierMap)
	case "Accidental_text":
		GongUnmarshallSliceOfPointers(&instance.Accidental_text, valueExpr, identifierMap)
	}
	return nil
}

type Non_arpeggiateUnmarshaller struct{}

func (u *Non_arpeggiateUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Non_arpeggiate), stage, identifier, instanceName, preserveOrder)
}

func (u *Non_arpeggiateUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Non_arpeggiate)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type_XSD":
		GongUnmarshallEnum(&instance.Type_XSD, valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Placement":
		GongUnmarshallEnum(&instance.Placement, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	}
	return nil
}

type NotationsUnmarshaller struct{}

func (u *NotationsUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Notations), stage, identifier, instanceName, preserveOrder)
}

func (u *NotationsUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Notations)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Print_object":
		GongUnmarshallEnum(&instance.Print_object, valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Footnote":
		GongUnmarshallPointer(&instance.Footnote, valueExpr, identifierMap)
	case "Level":
		GongUnmarshallPointer(&instance.Level, valueExpr, identifierMap)
	case "Tied":
		GongUnmarshallSliceOfPointers(&instance.Tied, valueExpr, identifierMap)
	case "Slur":
		GongUnmarshallSliceOfPointers(&instance.Slur, valueExpr, identifierMap)
	case "Tuplet":
		GongUnmarshallSliceOfPointers(&instance.Tuplet, valueExpr, identifierMap)
	case "Glissando":
		GongUnmarshallSliceOfPointers(&instance.Glissando, valueExpr, identifierMap)
	case "Slide":
		GongUnmarshallSliceOfPointers(&instance.Slide, valueExpr, identifierMap)
	case "Ornaments":
		GongUnmarshallSliceOfPointers(&instance.Ornaments, valueExpr, identifierMap)
	case "Technical":
		GongUnmarshallSliceOfPointers(&instance.Technical, valueExpr, identifierMap)
	case "Articulations":
		GongUnmarshallSliceOfPointers(&instance.Articulations, valueExpr, identifierMap)
	case "Dynamics":
		GongUnmarshallSliceOfPointers(&instance.Dynamics, valueExpr, identifierMap)
	case "Fermata":
		GongUnmarshallSliceOfPointers(&instance.Fermata, valueExpr, identifierMap)
	case "Arpeggiate":
		GongUnmarshallSliceOfPointers(&instance.Arpeggiate, valueExpr, identifierMap)
	case "Non_arpeggiate":
		GongUnmarshallSliceOfPointers(&instance.Non_arpeggiate, valueExpr, identifierMap)
	case "Accidental_mark":
		GongUnmarshallSliceOfPointers(&instance.Accidental_mark, valueExpr, identifierMap)
	case "Other_notation":
		GongUnmarshallSliceOfPointers(&instance.Other_notation, valueExpr, identifierMap)
	}
	return nil
}

type NoteUnmarshaller struct{}

func (u *NoteUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Note), stage, identifier, instanceName, preserveOrder)
}

func (u *NoteUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Note)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Print_leger":
		GongUnmarshallEnum(&instance.Print_leger, valueExpr)
	case "Dynamics":
		instance.Dynamics = GongExtractString(valueExpr)
	case "End_dynamics":
		instance.End_dynamics = GongExtractString(valueExpr)
	case "Attack":
		instance.Attack = GongExtractString(valueExpr)
	case "Release":
		instance.Release = GongExtractString(valueExpr)
	case "Time_only":
		instance.Time_only = GongExtractString(valueExpr)
	case "Pizzicato":
		GongUnmarshallEnum(&instance.Pizzicato, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Print_dot":
		GongUnmarshallEnum(&instance.Print_dot, valueExpr)
	case "Print_lyric":
		GongUnmarshallEnum(&instance.Print_lyric, valueExpr)
	case "Print_object":
		GongUnmarshallEnum(&instance.Print_object, valueExpr)
	case "Print_spacing":
		GongUnmarshallEnum(&instance.Print_spacing, valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Grace":
		GongUnmarshallPointer(&instance.Grace, valueExpr, identifierMap)
	case "Chord":
		instance.Chord = GongExtractString(valueExpr)
	case "Pitch":
		GongUnmarshallPointer(&instance.Pitch, valueExpr, identifierMap)
	case "Unpitched":
		GongUnmarshallPointer(&instance.Unpitched, valueExpr, identifierMap)
	case "Rest":
		GongUnmarshallPointer(&instance.Rest, valueExpr, identifierMap)
	case "Tie":
		GongUnmarshallPointer(&instance.Tie, valueExpr, identifierMap)
	case "Cue":
		instance.Cue = GongExtractString(valueExpr)
	case "Instrument":
		GongUnmarshallSliceOfPointers(&instance.Instrument, valueExpr, identifierMap)
	case "Footnote":
		GongUnmarshallPointer(&instance.Footnote, valueExpr, identifierMap)
	case "Level":
		GongUnmarshallPointer(&instance.Level, valueExpr, identifierMap)
	case "Voice":
		instance.Voice = GongExtractString(valueExpr)
	case "Type_XSD":
		GongUnmarshallPointer(&instance.Type_XSD, valueExpr, identifierMap)
	case "Dot":
		GongUnmarshallSliceOfPointers(&instance.Dot, valueExpr, identifierMap)
	case "Accidental":
		GongUnmarshallPointer(&instance.Accidental, valueExpr, identifierMap)
	case "Time_modification":
		GongUnmarshallPointer(&instance.Time_modification, valueExpr, identifierMap)
	case "Stem":
		GongUnmarshallPointer(&instance.Stem, valueExpr, identifierMap)
	case "Notehead":
		GongUnmarshallPointer(&instance.Notehead, valueExpr, identifierMap)
	case "Notehead_text":
		GongUnmarshallPointer(&instance.Notehead_text, valueExpr, identifierMap)
	case "Staff":
		instance.Staff = GongExtractInt(valueExpr)
	case "Beam":
		GongUnmarshallPointer(&instance.Beam, valueExpr, identifierMap)
	case "Notations":
		GongUnmarshallSliceOfPointers(&instance.Notations, valueExpr, identifierMap)
	case "Lyric":
		GongUnmarshallSliceOfPointers(&instance.Lyric, valueExpr, identifierMap)
	case "Play":
		GongUnmarshallPointer(&instance.Play, valueExpr, identifierMap)
	case "Listen":
		GongUnmarshallPointer(&instance.Listen, valueExpr, identifierMap)
	}
	return nil
}

type Note_sizeUnmarshaller struct{}

func (u *Note_sizeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Note_size), stage, identifier, instanceName, preserveOrder)
}

func (u *Note_sizeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Note_size)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type_XSD":
		GongUnmarshallEnum(&instance.Type_XSD, valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type Note_typeUnmarshaller struct{}

func (u *Note_typeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Note_type), stage, identifier, instanceName, preserveOrder)
}

func (u *Note_typeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Note_type)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Size":
		GongUnmarshallEnum(&instance.Size, valueExpr)
	case "EnclosedText":
		GongUnmarshallEnum(&instance.EnclosedText, valueExpr)
	}
	return nil
}

type NoteheadUnmarshaller struct{}

func (u *NoteheadUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Notehead), stage, identifier, instanceName, preserveOrder)
}

func (u *NoteheadUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Notehead)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Filled":
		GongUnmarshallEnum(&instance.Filled, valueExpr)
	case "Parentheses":
		GongUnmarshallEnum(&instance.Parentheses, valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Smufl":
		instance.Smufl = GongExtractString(valueExpr)
	case "EnclosedText":
		GongUnmarshallEnum(&instance.EnclosedText, valueExpr)
	}
	return nil
}

type Notehead_textUnmarshaller struct{}

func (u *Notehead_textUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Notehead_text), stage, identifier, instanceName, preserveOrder)
}

func (u *Notehead_textUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Notehead_text)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Display_text":
		GongUnmarshallSliceOfPointers(&instance.Display_text, valueExpr, identifierMap)
	case "Accidental_text":
		GongUnmarshallSliceOfPointers(&instance.Accidental_text, valueExpr, identifierMap)
	}
	return nil
}

type NumeralUnmarshaller struct{}

func (u *NumeralUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Numeral), stage, identifier, instanceName, preserveOrder)
}

func (u *NumeralUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Numeral)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Numeral_root":
		GongUnmarshallPointer(&instance.Numeral_root, valueExpr, identifierMap)
	case "Numeral_alter":
		GongUnmarshallPointer(&instance.Numeral_alter, valueExpr, identifierMap)
	case "Numeral_key":
		GongUnmarshallPointer(&instance.Numeral_key, valueExpr, identifierMap)
	}
	return nil
}

type Numeral_keyUnmarshaller struct{}

func (u *Numeral_keyUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Numeral_key), stage, identifier, instanceName, preserveOrder)
}

func (u *Numeral_keyUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Numeral_key)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Print_object":
		GongUnmarshallEnum(&instance.Print_object, valueExpr)
	case "Numeral_fifths":
		instance.Numeral_fifths = GongExtractInt(valueExpr)
	case "Numeral_mode":
		GongUnmarshallEnum(&instance.Numeral_mode, valueExpr)
	}
	return nil
}

type Numeral_rootUnmarshaller struct{}

func (u *Numeral_rootUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Numeral_root), stage, identifier, instanceName, preserveOrder)
}

func (u *Numeral_rootUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Numeral_root)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Text":
		instance.Text = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractInt(valueExpr)
	}
	return nil
}

type Octave_shiftUnmarshaller struct{}

func (u *Octave_shiftUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Octave_shift), stage, identifier, instanceName, preserveOrder)
}

func (u *Octave_shiftUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Octave_shift)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type_XSD":
		GongUnmarshallEnum(&instance.Type_XSD, valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "Size":
		instance.Size = GongExtractInt(valueExpr)
	case "Dash_length":
		instance.Dash_length = GongExtractString(valueExpr)
	case "Space_length":
		instance.Space_length = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	}
	return nil
}

type OffsetUnmarshaller struct{}

func (u *OffsetUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Offset), stage, identifier, instanceName, preserveOrder)
}

func (u *OffsetUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Offset)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Sound":
		GongUnmarshallEnum(&instance.Sound, valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type OpusUnmarshaller struct{}

func (u *OpusUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Opus), stage, identifier, instanceName, preserveOrder)
}

func (u *OpusUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Opus)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Href":
		instance.Href = GongExtractString(valueExpr)
	case "Type_XSD":
		instance.Type_XSD = GongExtractString(valueExpr)
	case "Role":
		instance.Role = GongExtractString(valueExpr)
	case "Title":
		instance.Title = GongExtractString(valueExpr)
	case "Show":
		instance.Show = GongExtractString(valueExpr)
	case "Actuate":
		instance.Actuate = GongExtractString(valueExpr)
	}
	return nil
}

type OrnamentsUnmarshaller struct{}

func (u *OrnamentsUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Ornaments), stage, identifier, instanceName, preserveOrder)
}

func (u *OrnamentsUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Ornaments)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Trill_mark":
		GongUnmarshallSliceOfPointers(&instance.Trill_mark, valueExpr, identifierMap)
	case "Turn":
		GongUnmarshallSliceOfPointers(&instance.Turn, valueExpr, identifierMap)
	case "Delayed_turn":
		GongUnmarshallSliceOfPointers(&instance.Delayed_turn, valueExpr, identifierMap)
	case "Inverted_turn":
		GongUnmarshallSliceOfPointers(&instance.Inverted_turn, valueExpr, identifierMap)
	case "Delayed_inverted_turn":
		GongUnmarshallSliceOfPointers(&instance.Delayed_inverted_turn, valueExpr, identifierMap)
	case "Vertical_turn":
		GongUnmarshallSliceOfPointers(&instance.Vertical_turn, valueExpr, identifierMap)
	case "Inverted_vertical_turn":
		GongUnmarshallSliceOfPointers(&instance.Inverted_vertical_turn, valueExpr, identifierMap)
	case "Shake":
		GongUnmarshallSliceOfPointers(&instance.Shake, valueExpr, identifierMap)
	case "Wavy_line":
		GongUnmarshallSliceOfPointers(&instance.Wavy_line, valueExpr, identifierMap)
	case "Mordent":
		GongUnmarshallSliceOfPointers(&instance.Mordent, valueExpr, identifierMap)
	case "Inverted_mordent":
		GongUnmarshallSliceOfPointers(&instance.Inverted_mordent, valueExpr, identifierMap)
	case "Schleifer":
		GongUnmarshallSliceOfPointers(&instance.Schleifer, valueExpr, identifierMap)
	case "Tremolo":
		GongUnmarshallSliceOfPointers(&instance.Tremolo, valueExpr, identifierMap)
	case "Haydn":
		GongUnmarshallSliceOfPointers(&instance.Haydn, valueExpr, identifierMap)
	case "Other_ornament":
		GongUnmarshallSliceOfPointers(&instance.Other_ornament, valueExpr, identifierMap)
	case "Accidental_mark":
		GongUnmarshallSliceOfPointers(&instance.Accidental_mark, valueExpr, identifierMap)
	}
	return nil
}

type Other_appearanceUnmarshaller struct{}

func (u *Other_appearanceUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Other_appearance), stage, identifier, instanceName, preserveOrder)
}

func (u *Other_appearanceUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Other_appearance)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type_XSD":
		instance.Type_XSD = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type Other_directionUnmarshaller struct{}

func (u *Other_directionUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Other_direction), stage, identifier, instanceName, preserveOrder)
}

func (u *Other_directionUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Other_direction)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Print_object":
		GongUnmarshallEnum(&instance.Print_object, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Halign":
		GongUnmarshallEnum(&instance.Halign, valueExpr)
	case "Valign":
		GongUnmarshallEnum(&instance.Valign, valueExpr)
	case "Smufl":
		instance.Smufl = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type Other_listeningUnmarshaller struct{}

func (u *Other_listeningUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Other_listening), stage, identifier, instanceName, preserveOrder)
}

func (u *Other_listeningUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Other_listening)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type_XSD":
		instance.Type_XSD = GongExtractString(valueExpr)
	case "Player":
		instance.Player = GongExtractString(valueExpr)
	case "Time_only":
		instance.Time_only = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type Other_notationUnmarshaller struct{}

func (u *Other_notationUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Other_notation), stage, identifier, instanceName, preserveOrder)
}

func (u *Other_notationUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Other_notation)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type_XSD":
		GongUnmarshallEnum(&instance.Type_XSD, valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "Print_object":
		GongUnmarshallEnum(&instance.Print_object, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Placement":
		GongUnmarshallEnum(&instance.Placement, valueExpr)
	case "Smufl":
		instance.Smufl = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type Other_placement_textUnmarshaller struct{}

func (u *Other_placement_textUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Other_placement_text), stage, identifier, instanceName, preserveOrder)
}

func (u *Other_placement_textUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Other_placement_text)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Placement":
		GongUnmarshallEnum(&instance.Placement, valueExpr)
	case "Smufl":
		instance.Smufl = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type Other_playUnmarshaller struct{}

func (u *Other_playUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Other_play), stage, identifier, instanceName, preserveOrder)
}

func (u *Other_playUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Other_play)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type_XSD":
		instance.Type_XSD = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type Other_textUnmarshaller struct{}

func (u *Other_textUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Other_text), stage, identifier, instanceName, preserveOrder)
}

func (u *Other_textUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Other_text)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Smufl":
		instance.Smufl = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type Page_layoutUnmarshaller struct{}

func (u *Page_layoutUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Page_layout), stage, identifier, instanceName, preserveOrder)
}

func (u *Page_layoutUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Page_layout)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Page_height":
		instance.Page_height = GongExtractString(valueExpr)
	case "Page_width":
		instance.Page_width = GongExtractString(valueExpr)
	case "Page_margins":
		GongUnmarshallPointer(&instance.Page_margins, valueExpr, identifierMap)
	}
	return nil
}

type Page_marginsUnmarshaller struct{}

func (u *Page_marginsUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Page_margins), stage, identifier, instanceName, preserveOrder)
}

func (u *Page_marginsUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Page_margins)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type_XSD":
		GongUnmarshallEnum(&instance.Type_XSD, valueExpr)
	case "Left_margin":
		instance.Left_margin = GongExtractString(valueExpr)
	case "Right_margin":
		instance.Right_margin = GongExtractString(valueExpr)
	case "Top_margin":
		instance.Top_margin = GongExtractString(valueExpr)
	case "Bottom_margin":
		instance.Bottom_margin = GongExtractString(valueExpr)
	}
	return nil
}

type Part_clefUnmarshaller struct{}

func (u *Part_clefUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Part_clef), stage, identifier, instanceName, preserveOrder)
}

func (u *Part_clefUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Part_clef)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Sign":
		GongUnmarshallEnum(&instance.Sign, valueExpr)
	case "Line":
		instance.Line = GongExtractInt(valueExpr)
	case "Clef_octave_change":
		instance.Clef_octave_change = GongExtractInt(valueExpr)
	}
	return nil
}

type Part_groupUnmarshaller struct{}

func (u *Part_groupUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Part_group), stage, identifier, instanceName, preserveOrder)
}

func (u *Part_groupUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Part_group)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type_XSD":
		GongUnmarshallEnum(&instance.Type_XSD, valueExpr)
	case "Number":
		instance.Number = GongExtractString(valueExpr)
	case "Group_name":
		GongUnmarshallPointer(&instance.Group_name, valueExpr, identifierMap)
	case "Group_name_display":
		GongUnmarshallPointer(&instance.Group_name_display, valueExpr, identifierMap)
	case "Group_abbreviation":
		GongUnmarshallPointer(&instance.Group_abbreviation, valueExpr, identifierMap)
	case "Group_abbreviation_display":
		GongUnmarshallPointer(&instance.Group_abbreviation_display, valueExpr, identifierMap)
	case "Group_symbol":
		GongUnmarshallPointer(&instance.Group_symbol, valueExpr, identifierMap)
	case "Group_barline":
		GongUnmarshallPointer(&instance.Group_barline, valueExpr, identifierMap)
	case "Group_time":
		instance.Group_time = GongExtractString(valueExpr)
	case "Footnote":
		GongUnmarshallPointer(&instance.Footnote, valueExpr, identifierMap)
	case "Level":
		GongUnmarshallPointer(&instance.Level, valueExpr, identifierMap)
	}
	return nil
}

type Part_linkUnmarshaller struct{}

func (u *Part_linkUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Part_link), stage, identifier, instanceName, preserveOrder)
}

func (u *Part_linkUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Part_link)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Href":
		instance.Href = GongExtractString(valueExpr)
	case "Type_XSD":
		instance.Type_XSD = GongExtractString(valueExpr)
	case "Role":
		instance.Role = GongExtractString(valueExpr)
	case "Title":
		instance.Title = GongExtractString(valueExpr)
	case "Show":
		instance.Show = GongExtractString(valueExpr)
	case "Actuate":
		instance.Actuate = GongExtractString(valueExpr)
	case "Instrument_link":
		GongUnmarshallSliceOfPointers(&instance.Instrument_link, valueExpr, identifierMap)
	case "Group_link":
		instance.Group_link = GongExtractString(valueExpr)
	}
	return nil
}

type Part_listUnmarshaller struct{}

func (u *Part_listUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Part_list), stage, identifier, instanceName, preserveOrder)
}

func (u *Part_listUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Part_list)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Part_group":
		GongUnmarshallPointer(&instance.Part_group, valueExpr, identifierMap)
	case "Score_part":
		GongUnmarshallPointer(&instance.Score_part, valueExpr, identifierMap)
	}
	return nil
}

type Part_nameUnmarshaller struct{}

func (u *Part_nameUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Part_name), stage, identifier, instanceName, preserveOrder)
}

func (u *Part_nameUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Part_name)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Print_object":
		GongUnmarshallEnum(&instance.Print_object, valueExpr)
	case "Justify":
		GongUnmarshallEnum(&instance.Justify, valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type Part_symbolUnmarshaller struct{}

func (u *Part_symbolUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Part_symbol), stage, identifier, instanceName, preserveOrder)
}

func (u *Part_symbolUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Part_symbol)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Top_staff":
		instance.Top_staff = GongExtractInt(valueExpr)
	case "Bottom_staff":
		instance.Bottom_staff = GongExtractInt(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "EnclosedText":
		GongUnmarshallEnum(&instance.EnclosedText, valueExpr)
	}
	return nil
}

type Part_transposeUnmarshaller struct{}

func (u *Part_transposeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Part_transpose), stage, identifier, instanceName, preserveOrder)
}

func (u *Part_transposeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Part_transpose)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Diatonic":
		instance.Diatonic = GongExtractInt(valueExpr)
	case "Chromatic":
		instance.Chromatic = GongExtractString(valueExpr)
	case "Octave_change":
		instance.Octave_change = GongExtractInt(valueExpr)
	case "Double":
		instance.Double = GongExtractFloat(valueExpr)
	}
	return nil
}

type PedalUnmarshaller struct{}

func (u *PedalUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Pedal), stage, identifier, instanceName, preserveOrder)
}

func (u *PedalUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Pedal)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type_XSD":
		GongUnmarshallEnum(&instance.Type_XSD, valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "Line":
		GongUnmarshallEnum(&instance.Line, valueExpr)
	case "Sign":
		GongUnmarshallEnum(&instance.Sign, valueExpr)
	case "Abbreviated":
		GongUnmarshallEnum(&instance.Abbreviated, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Halign":
		GongUnmarshallEnum(&instance.Halign, valueExpr)
	case "Valign":
		GongUnmarshallEnum(&instance.Valign, valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	}
	return nil
}

type Pedal_tuningUnmarshaller struct{}

func (u *Pedal_tuningUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Pedal_tuning), stage, identifier, instanceName, preserveOrder)
}

func (u *Pedal_tuningUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Pedal_tuning)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Pedal_step":
		GongUnmarshallEnum(&instance.Pedal_step, valueExpr)
	case "Pedal_alter":
		instance.Pedal_alter = GongExtractString(valueExpr)
	}
	return nil
}

type Per_minuteUnmarshaller struct{}

func (u *Per_minuteUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Per_minute), stage, identifier, instanceName, preserveOrder)
}

func (u *Per_minuteUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Per_minute)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type PercussionUnmarshaller struct{}

func (u *PercussionUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Percussion), stage, identifier, instanceName, preserveOrder)
}

func (u *PercussionUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Percussion)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Halign":
		GongUnmarshallEnum(&instance.Halign, valueExpr)
	case "Valign":
		GongUnmarshallEnum(&instance.Valign, valueExpr)
	case "Enclosure":
		GongUnmarshallEnum(&instance.Enclosure, valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Glass":
		GongUnmarshallPointer(&instance.Glass, valueExpr, identifierMap)
	case "Metal":
		GongUnmarshallPointer(&instance.Metal, valueExpr, identifierMap)
	case "Wood":
		GongUnmarshallPointer(&instance.Wood, valueExpr, identifierMap)
	case "Pitched":
		GongUnmarshallPointer(&instance.Pitched, valueExpr, identifierMap)
	case "Membrane":
		GongUnmarshallPointer(&instance.Membrane, valueExpr, identifierMap)
	case "Effect":
		GongUnmarshallPointer(&instance.Effect, valueExpr, identifierMap)
	case "Timpani":
		GongUnmarshallPointer(&instance.Timpani, valueExpr, identifierMap)
	case "Beater":
		GongUnmarshallPointer(&instance.Beater, valueExpr, identifierMap)
	case "Stick":
		GongUnmarshallPointer(&instance.Stick, valueExpr, identifierMap)
	case "Stick_location":
		GongUnmarshallEnum(&instance.Stick_location, valueExpr)
	case "Other_percussion":
		GongUnmarshallPointer(&instance.Other_percussion, valueExpr, identifierMap)
	}
	return nil
}

type PitchUnmarshaller struct{}

func (u *PitchUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Pitch), stage, identifier, instanceName, preserveOrder)
}

func (u *PitchUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Pitch)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Step":
		GongUnmarshallEnum(&instance.Step, valueExpr)
	case "Alter":
		instance.Alter = GongExtractString(valueExpr)
	case "Octave":
		instance.Octave = GongExtractInt(valueExpr)
	}
	return nil
}

type PitchedUnmarshaller struct{}

func (u *PitchedUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Pitched), stage, identifier, instanceName, preserveOrder)
}

func (u *PitchedUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Pitched)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Smufl":
		instance.Smufl = GongExtractString(valueExpr)
	case "EnclosedText":
		GongUnmarshallEnum(&instance.EnclosedText, valueExpr)
	}
	return nil
}

type Placement_textUnmarshaller struct{}

func (u *Placement_textUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Placement_text), stage, identifier, instanceName, preserveOrder)
}

func (u *Placement_textUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Placement_text)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Placement":
		GongUnmarshallEnum(&instance.Placement, valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type PlayUnmarshaller struct{}

func (u *PlayUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Play), stage, identifier, instanceName, preserveOrder)
}

func (u *PlayUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Play)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Ipa":
		instance.Ipa = GongExtractString(valueExpr)
	case "Mute":
		GongUnmarshallEnum(&instance.Mute, valueExpr)
	case "Semi_pitched":
		GongUnmarshallEnum(&instance.Semi_pitched, valueExpr)
	case "Other_play":
		GongUnmarshallSliceOfPointers(&instance.Other_play, valueExpr, identifierMap)
	}
	return nil
}

type PlayerUnmarshaller struct{}

func (u *PlayerUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Player), stage, identifier, instanceName, preserveOrder)
}

func (u *PlayerUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Player)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Player_name":
		instance.Player_name = GongExtractString(valueExpr)
	}
	return nil
}

type Principal_voiceUnmarshaller struct{}

func (u *Principal_voiceUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Principal_voice), stage, identifier, instanceName, preserveOrder)
}

func (u *Principal_voiceUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Principal_voice)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type_XSD":
		GongUnmarshallEnum(&instance.Type_XSD, valueExpr)
	case "Symbol":
		GongUnmarshallEnum(&instance.Symbol, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Halign":
		GongUnmarshallEnum(&instance.Halign, valueExpr)
	case "Valign":
		GongUnmarshallEnum(&instance.Valign, valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type Print_XSDUnmarshaller struct{}

func (u *Print_XSDUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Print_XSD), stage, identifier, instanceName, preserveOrder)
}

func (u *Print_XSDUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Print_XSD)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Staff_spacing":
		instance.Staff_spacing = GongExtractString(valueExpr)
	case "New_system":
		GongUnmarshallEnum(&instance.New_system, valueExpr)
	case "New_page":
		GongUnmarshallEnum(&instance.New_page, valueExpr)
	case "Blank_page":
		instance.Blank_page = GongExtractInt(valueExpr)
	case "Page_number":
		instance.Page_number = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Page_layout":
		GongUnmarshallPointer(&instance.Page_layout, valueExpr, identifierMap)
	case "System_layout":
		GongUnmarshallPointer(&instance.System_layout, valueExpr, identifierMap)
	case "Staff_layout":
		GongUnmarshallSliceOfPointers(&instance.Staff_layout, valueExpr, identifierMap)
	case "Measure_layout":
		GongUnmarshallPointer(&instance.Measure_layout, valueExpr, identifierMap)
	case "Measure_numbering":
		GongUnmarshallPointer(&instance.Measure_numbering, valueExpr, identifierMap)
	case "Part_name_display":
		GongUnmarshallPointer(&instance.Part_name_display, valueExpr, identifierMap)
	case "Part_abbreviation_display":
		GongUnmarshallPointer(&instance.Part_abbreviation_display, valueExpr, identifierMap)
	}
	return nil
}

type ReleaseUnmarshaller struct{}

func (u *ReleaseUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Release), stage, identifier, instanceName, preserveOrder)
}

func (u *ReleaseUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Release)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Offset":
		instance.Offset = GongExtractString(valueExpr)
	}
	return nil
}

type RepeatUnmarshaller struct{}

func (u *RepeatUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Repeat), stage, identifier, instanceName, preserveOrder)
}

func (u *RepeatUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Repeat)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Direction":
		GongUnmarshallEnum(&instance.Direction, valueExpr)
	case "Times":
		instance.Times = GongExtractInt(valueExpr)
	case "After_jump":
		GongUnmarshallEnum(&instance.After_jump, valueExpr)
	case "Winged":
		GongUnmarshallEnum(&instance.Winged, valueExpr)
	}
	return nil
}

type RestUnmarshaller struct{}

func (u *RestUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Rest), stage, identifier, instanceName, preserveOrder)
}

func (u *RestUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Rest)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Measure":
		GongUnmarshallEnum(&instance.Measure, valueExpr)
	case "Display_step":
		GongUnmarshallEnum(&instance.Display_step, valueExpr)
	case "Display_octave":
		instance.Display_octave = GongExtractInt(valueExpr)
	}
	return nil
}

type RootUnmarshaller struct{}

func (u *RootUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Root), stage, identifier, instanceName, preserveOrder)
}

func (u *RootUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Root)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Root_step":
		GongUnmarshallPointer(&instance.Root_step, valueExpr, identifierMap)
	case "Root_alter":
		GongUnmarshallPointer(&instance.Root_alter, valueExpr, identifierMap)
	}
	return nil
}

type Root_stepUnmarshaller struct{}

func (u *Root_stepUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Root_step), stage, identifier, instanceName, preserveOrder)
}

func (u *Root_stepUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Root_step)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Text":
		instance.Text = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "EnclosedText":
		GongUnmarshallEnum(&instance.EnclosedText, valueExpr)
	}
	return nil
}

type ScalingUnmarshaller struct{}

func (u *ScalingUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Scaling), stage, identifier, instanceName, preserveOrder)
}

func (u *ScalingUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Scaling)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Millimeters":
		instance.Millimeters = GongExtractString(valueExpr)
	case "Tenths":
		instance.Tenths = GongExtractString(valueExpr)
	}
	return nil
}

type ScordaturaUnmarshaller struct{}

func (u *ScordaturaUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Scordatura), stage, identifier, instanceName, preserveOrder)
}

func (u *ScordaturaUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Scordatura)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Accord":
		GongUnmarshallSliceOfPointers(&instance.Accord, valueExpr, identifierMap)
	}
	return nil
}

type Score_instrumentUnmarshaller struct{}

func (u *Score_instrumentUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Score_instrument), stage, identifier, instanceName, preserveOrder)
}

func (u *Score_instrumentUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Score_instrument)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Instrument_name":
		instance.Instrument_name = GongExtractString(valueExpr)
	case "Instrument_abbreviation":
		instance.Instrument_abbreviation = GongExtractString(valueExpr)
	case "Instrument_sound":
		instance.Instrument_sound = GongExtractString(valueExpr)
	case "Solo":
		instance.Solo = GongExtractString(valueExpr)
	case "Ensemble":
		instance.Ensemble = GongExtractString(valueExpr)
	case "Virtual_instrument":
		GongUnmarshallPointer(&instance.Virtual_instrument, valueExpr, identifierMap)
	}
	return nil
}

type Score_partUnmarshaller struct{}

func (u *Score_partUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Score_part), stage, identifier, instanceName, preserveOrder)
}

func (u *Score_partUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Score_part)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Identification":
		GongUnmarshallPointer(&instance.Identification, valueExpr, identifierMap)
	case "Part_link":
		GongUnmarshallSliceOfPointers(&instance.Part_link, valueExpr, identifierMap)
	case "Part_name":
		GongUnmarshallPointer(&instance.Part_name, valueExpr, identifierMap)
	case "Part_name_display":
		GongUnmarshallPointer(&instance.Part_name_display, valueExpr, identifierMap)
	case "Part_abbreviation":
		GongUnmarshallPointer(&instance.Part_abbreviation, valueExpr, identifierMap)
	case "Part_abbreviation_display":
		GongUnmarshallPointer(&instance.Part_abbreviation_display, valueExpr, identifierMap)
	case "Group":
		instance.Group = GongExtractString(valueExpr)
	case "Score_instrument":
		GongUnmarshallSliceOfPointers(&instance.Score_instrument, valueExpr, identifierMap)
	case "Player":
		GongUnmarshallSliceOfPointers(&instance.Player, valueExpr, identifierMap)
	case "Midi_device":
		GongUnmarshallSliceOfPointers(&instance.Midi_device, valueExpr, identifierMap)
	case "Midi_instrument":
		GongUnmarshallSliceOfPointers(&instance.Midi_instrument, valueExpr, identifierMap)
	}
	return nil
}

type Score_partwiseUnmarshaller struct{}

func (u *Score_partwiseUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Score_partwise), stage, identifier, instanceName, preserveOrder)
}

func (u *Score_partwiseUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Score_partwise)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Version":
		instance.Version = GongExtractString(valueExpr)
	case "Work":
		GongUnmarshallPointer(&instance.Work, valueExpr, identifierMap)
	case "Movement_number":
		instance.Movement_number = GongExtractString(valueExpr)
	case "Movement_title":
		instance.Movement_title = GongExtractString(valueExpr)
	case "Identification":
		GongUnmarshallPointer(&instance.Identification, valueExpr, identifierMap)
	case "Defaults":
		GongUnmarshallPointer(&instance.Defaults, valueExpr, identifierMap)
	case "Credit":
		GongUnmarshallSliceOfPointers(&instance.Credit, valueExpr, identifierMap)
	case "Part_list":
		GongUnmarshallPointer(&instance.Part_list, valueExpr, identifierMap)
	case "Part":
		GongUnmarshallSliceOfPointers(&instance.Part, valueExpr, identifierMap)
	}
	return nil
}

type Score_timewiseUnmarshaller struct{}

func (u *Score_timewiseUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Score_timewise), stage, identifier, instanceName, preserveOrder)
}

func (u *Score_timewiseUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Score_timewise)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Version":
		instance.Version = GongExtractString(valueExpr)
	case "Work":
		GongUnmarshallPointer(&instance.Work, valueExpr, identifierMap)
	case "Movement_number":
		instance.Movement_number = GongExtractString(valueExpr)
	case "Movement_title":
		instance.Movement_title = GongExtractString(valueExpr)
	case "Identification":
		GongUnmarshallPointer(&instance.Identification, valueExpr, identifierMap)
	case "Defaults":
		GongUnmarshallPointer(&instance.Defaults, valueExpr, identifierMap)
	case "Credit":
		GongUnmarshallSliceOfPointers(&instance.Credit, valueExpr, identifierMap)
	case "Part_list":
		GongUnmarshallPointer(&instance.Part_list, valueExpr, identifierMap)
	case "Measure":
		GongUnmarshallSliceOfPointers(&instance.Measure, valueExpr, identifierMap)
	}
	return nil
}

type SegnoUnmarshaller struct{}

func (u *SegnoUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Segno), stage, identifier, instanceName, preserveOrder)
}

func (u *SegnoUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Segno)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Smufl":
		instance.Smufl = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Halign":
		GongUnmarshallEnum(&instance.Halign, valueExpr)
	case "Valign":
		GongUnmarshallEnum(&instance.Valign, valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	}
	return nil
}

type SlashUnmarshaller struct{}

func (u *SlashUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Slash), stage, identifier, instanceName, preserveOrder)
}

func (u *SlashUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Slash)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type_XSD":
		GongUnmarshallEnum(&instance.Type_XSD, valueExpr)
	case "Use_dots":
		GongUnmarshallEnum(&instance.Use_dots, valueExpr)
	case "Use_stems":
		GongUnmarshallEnum(&instance.Use_stems, valueExpr)
	case "Slash_type":
		GongUnmarshallEnum(&instance.Slash_type, valueExpr)
	case "Slash_dot":
		instance.Slash_dot = GongExtractString(valueExpr)
	case "Except_voice":
		instance.Except_voice = GongExtractString(valueExpr)
	}
	return nil
}

type SlideUnmarshaller struct{}

func (u *SlideUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Slide), stage, identifier, instanceName, preserveOrder)
}

func (u *SlideUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Slide)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type_XSD":
		GongUnmarshallEnum(&instance.Type_XSD, valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "Line_type":
		GongUnmarshallEnum(&instance.Line_type, valueExpr)
	case "Dash_length":
		instance.Dash_length = GongExtractString(valueExpr)
	case "Space_length":
		instance.Space_length = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Accelerate":
		GongUnmarshallEnum(&instance.Accelerate, valueExpr)
	case "Beats":
		instance.Beats = GongExtractString(valueExpr)
	case "First_beat":
		instance.First_beat = GongExtractString(valueExpr)
	case "Last_beat":
		instance.Last_beat = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type SlurUnmarshaller struct{}

func (u *SlurUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Slur), stage, identifier, instanceName, preserveOrder)
}

func (u *SlurUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Slur)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type_XSD":
		GongUnmarshallEnum(&instance.Type_XSD, valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "Line_type":
		GongUnmarshallEnum(&instance.Line_type, valueExpr)
	case "Dash_length":
		instance.Dash_length = GongExtractString(valueExpr)
	case "Space_length":
		instance.Space_length = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Placement":
		GongUnmarshallEnum(&instance.Placement, valueExpr)
	case "Orientation":
		GongUnmarshallEnum(&instance.Orientation, valueExpr)
	case "Bezier_x":
		instance.Bezier_x = GongExtractString(valueExpr)
	case "Bezier_y":
		instance.Bezier_y = GongExtractString(valueExpr)
	case "Bezier_x2":
		instance.Bezier_x2 = GongExtractString(valueExpr)
	case "Bezier_y2":
		instance.Bezier_y2 = GongExtractString(valueExpr)
	case "Bezier_offset":
		instance.Bezier_offset = GongExtractString(valueExpr)
	case "Bezier_offset2":
		instance.Bezier_offset2 = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	}
	return nil
}

type SoundUnmarshaller struct{}

func (u *SoundUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Sound), stage, identifier, instanceName, preserveOrder)
}

func (u *SoundUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Sound)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Tempo":
		instance.Tempo = GongExtractString(valueExpr)
	case "Dynamics":
		instance.Dynamics = GongExtractString(valueExpr)
	case "Dacapo":
		GongUnmarshallEnum(&instance.Dacapo, valueExpr)
	case "Segno":
		instance.Segno = GongExtractString(valueExpr)
	case "Dalsegno":
		instance.Dalsegno = GongExtractString(valueExpr)
	case "Coda":
		instance.Coda = GongExtractString(valueExpr)
	case "Tocoda":
		instance.Tocoda = GongExtractString(valueExpr)
	case "Divisions":
		instance.Divisions = GongExtractString(valueExpr)
	case "Forward_repeat":
		GongUnmarshallEnum(&instance.Forward_repeat, valueExpr)
	case "Fine":
		instance.Fine = GongExtractString(valueExpr)
	case "Time_only":
		instance.Time_only = GongExtractString(valueExpr)
	case "Pizzicato":
		GongUnmarshallEnum(&instance.Pizzicato, valueExpr)
	case "Pan":
		instance.Pan = GongExtractString(valueExpr)
	case "Elevation":
		instance.Elevation = GongExtractString(valueExpr)
	case "Damper_pedal":
		instance.Damper_pedal = GongExtractString(valueExpr)
	case "Soft_pedal":
		instance.Soft_pedal = GongExtractString(valueExpr)
	case "Sostenuto_pedal":
		instance.Sostenuto_pedal = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Instrument_change":
		GongUnmarshallSliceOfPointers(&instance.Instrument_change, valueExpr, identifierMap)
	case "Midi_device":
		GongUnmarshallSliceOfPointers(&instance.Midi_device, valueExpr, identifierMap)
	case "Midi_instrument":
		GongUnmarshallSliceOfPointers(&instance.Midi_instrument, valueExpr, identifierMap)
	case "Play":
		GongUnmarshallSliceOfPointers(&instance.Play, valueExpr, identifierMap)
	case "Swing":
		GongUnmarshallPointer(&instance.Swing, valueExpr, identifierMap)
	case "Offset":
		GongUnmarshallPointer(&instance.Offset, valueExpr, identifierMap)
	}
	return nil
}

type Staff_detailsUnmarshaller struct{}

func (u *Staff_detailsUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Staff_details), stage, identifier, instanceName, preserveOrder)
}

func (u *Staff_detailsUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Staff_details)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "Show_frets":
		GongUnmarshallEnum(&instance.Show_frets, valueExpr)
	case "Print_object":
		GongUnmarshallEnum(&instance.Print_object, valueExpr)
	case "Print_spacing":
		GongUnmarshallEnum(&instance.Print_spacing, valueExpr)
	case "Staff_type":
		GongUnmarshallEnum(&instance.Staff_type, valueExpr)
	case "Staff_lines":
		instance.Staff_lines = GongExtractInt(valueExpr)
	case "Line_detail":
		GongUnmarshallSliceOfPointers(&instance.Line_detail, valueExpr, identifierMap)
	case "Staff_tuning":
		GongUnmarshallSliceOfPointers(&instance.Staff_tuning, valueExpr, identifierMap)
	case "Capo":
		instance.Capo = GongExtractInt(valueExpr)
	case "Staff_size":
		GongUnmarshallPointer(&instance.Staff_size, valueExpr, identifierMap)
	}
	return nil
}

type Staff_divideUnmarshaller struct{}

func (u *Staff_divideUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Staff_divide), stage, identifier, instanceName, preserveOrder)
}

func (u *Staff_divideUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Staff_divide)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type_XSD":
		GongUnmarshallEnum(&instance.Type_XSD, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Halign":
		GongUnmarshallEnum(&instance.Halign, valueExpr)
	case "Valign":
		GongUnmarshallEnum(&instance.Valign, valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	}
	return nil
}

type Staff_layoutUnmarshaller struct{}

func (u *Staff_layoutUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Staff_layout), stage, identifier, instanceName, preserveOrder)
}

func (u *Staff_layoutUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Staff_layout)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "Staff_distance":
		instance.Staff_distance = GongExtractString(valueExpr)
	}
	return nil
}

type Staff_sizeUnmarshaller struct{}

func (u *Staff_sizeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Staff_size), stage, identifier, instanceName, preserveOrder)
}

func (u *Staff_sizeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Staff_size)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Scaling":
		instance.Scaling = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type Staff_tuningUnmarshaller struct{}

func (u *Staff_tuningUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Staff_tuning), stage, identifier, instanceName, preserveOrder)
}

func (u *Staff_tuningUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Staff_tuning)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Line":
		instance.Line = GongExtractInt(valueExpr)
	case "Tuning_step":
		GongUnmarshallEnum(&instance.Tuning_step, valueExpr)
	case "Tuning_alter":
		instance.Tuning_alter = GongExtractString(valueExpr)
	case "Tuning_octave":
		instance.Tuning_octave = GongExtractInt(valueExpr)
	}
	return nil
}

type StemUnmarshaller struct{}

func (u *StemUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Stem), stage, identifier, instanceName, preserveOrder)
}

func (u *StemUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Stem)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "EnclosedText":
		GongUnmarshallEnum(&instance.EnclosedText, valueExpr)
	}
	return nil
}

type StickUnmarshaller struct{}

func (u *StickUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Stick), stage, identifier, instanceName, preserveOrder)
}

func (u *StickUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Stick)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Tip":
		GongUnmarshallEnum(&instance.Tip, valueExpr)
	case "Parentheses":
		GongUnmarshallEnum(&instance.Parentheses, valueExpr)
	case "Dashed_circle":
		GongUnmarshallEnum(&instance.Dashed_circle, valueExpr)
	case "Stick_type":
		GongUnmarshallEnum(&instance.Stick_type, valueExpr)
	case "Stick_material":
		GongUnmarshallEnum(&instance.Stick_material, valueExpr)
	}
	return nil
}

type String_muteUnmarshaller struct{}

func (u *String_muteUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(String_mute), stage, identifier, instanceName, preserveOrder)
}

func (u *String_muteUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*String_mute)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type_XSD":
		GongUnmarshallEnum(&instance.Type_XSD, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Halign":
		GongUnmarshallEnum(&instance.Halign, valueExpr)
	case "Valign":
		GongUnmarshallEnum(&instance.Valign, valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	}
	return nil
}

type String_typeUnmarshaller struct{}

func (u *String_typeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(String_type), stage, identifier, instanceName, preserveOrder)
}

func (u *String_typeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*String_type)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Placement":
		GongUnmarshallEnum(&instance.Placement, valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractInt(valueExpr)
	}
	return nil
}

type Strong_accentUnmarshaller struct{}

func (u *Strong_accentUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Strong_accent), stage, identifier, instanceName, preserveOrder)
}

func (u *Strong_accentUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Strong_accent)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Placement":
		GongUnmarshallEnum(&instance.Placement, valueExpr)
	case "Type_XSD":
		GongUnmarshallEnum(&instance.Type_XSD, valueExpr)
	}
	return nil
}

type Style_textUnmarshaller struct{}

func (u *Style_textUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Style_text), stage, identifier, instanceName, preserveOrder)
}

func (u *Style_textUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Style_text)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type SupportsUnmarshaller struct{}

func (u *SupportsUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Supports), stage, identifier, instanceName, preserveOrder)
}

func (u *SupportsUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Supports)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type_XSD":
		GongUnmarshallEnum(&instance.Type_XSD, valueExpr)
	case "Element":
		instance.Element = GongExtractString(valueExpr)
	case "Attribute":
		instance.Attribute = GongExtractString(valueExpr)
	case "Value":
		instance.Value = GongExtractString(valueExpr)
	}
	return nil
}

type SwingUnmarshaller struct{}

func (u *SwingUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Swing), stage, identifier, instanceName, preserveOrder)
}

func (u *SwingUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Swing)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Straight":
		instance.Straight = GongExtractString(valueExpr)
	case "First":
		instance.First = GongExtractInt(valueExpr)
	case "Second":
		instance.Second = GongExtractInt(valueExpr)
	case "Swing_type":
		instance.Swing_type = GongExtractString(valueExpr)
	case "Swing_style":
		instance.Swing_style = GongExtractString(valueExpr)
	}
	return nil
}

type SyncUnmarshaller struct{}

func (u *SyncUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Sync), stage, identifier, instanceName, preserveOrder)
}

func (u *SyncUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Sync)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type_XSD":
		GongUnmarshallEnum(&instance.Type_XSD, valueExpr)
	case "Latency":
		instance.Latency = GongExtractInt(valueExpr)
	case "Player":
		instance.Player = GongExtractString(valueExpr)
	case "Time_only":
		instance.Time_only = GongExtractString(valueExpr)
	}
	return nil
}

type System_dividersUnmarshaller struct{}

func (u *System_dividersUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(System_dividers), stage, identifier, instanceName, preserveOrder)
}

func (u *System_dividersUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*System_dividers)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Left_divider":
		GongUnmarshallPointer(&instance.Left_divider, valueExpr, identifierMap)
	case "Right_divider":
		GongUnmarshallPointer(&instance.Right_divider, valueExpr, identifierMap)
	}
	return nil
}

type System_layoutUnmarshaller struct{}

func (u *System_layoutUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(System_layout), stage, identifier, instanceName, preserveOrder)
}

func (u *System_layoutUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*System_layout)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "System_margins":
		GongUnmarshallPointer(&instance.System_margins, valueExpr, identifierMap)
	case "System_distance":
		instance.System_distance = GongExtractString(valueExpr)
	case "Top_system_distance":
		instance.Top_system_distance = GongExtractString(valueExpr)
	case "System_dividers":
		GongUnmarshallPointer(&instance.System_dividers, valueExpr, identifierMap)
	}
	return nil
}

type System_marginsUnmarshaller struct{}

func (u *System_marginsUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(System_margins), stage, identifier, instanceName, preserveOrder)
}

func (u *System_marginsUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*System_margins)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Left_margin":
		instance.Left_margin = GongExtractString(valueExpr)
	case "Right_margin":
		instance.Right_margin = GongExtractString(valueExpr)
	}
	return nil
}

type TapUnmarshaller struct{}

func (u *TapUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Tap), stage, identifier, instanceName, preserveOrder)
}

func (u *TapUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Tap)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Hand":
		GongUnmarshallEnum(&instance.Hand, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Placement":
		GongUnmarshallEnum(&instance.Placement, valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type TechnicalUnmarshaller struct{}

func (u *TechnicalUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Technical), stage, identifier, instanceName, preserveOrder)
}

func (u *TechnicalUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Technical)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Up_bow":
		GongUnmarshallSliceOfPointers(&instance.Up_bow, valueExpr, identifierMap)
	case "Down_bow":
		GongUnmarshallSliceOfPointers(&instance.Down_bow, valueExpr, identifierMap)
	case "Harmonic":
		GongUnmarshallSliceOfPointers(&instance.Harmonic, valueExpr, identifierMap)
	case "Open_string":
		GongUnmarshallSliceOfPointers(&instance.Open_string, valueExpr, identifierMap)
	case "Thumb_position":
		GongUnmarshallSliceOfPointers(&instance.Thumb_position, valueExpr, identifierMap)
	case "Fingering":
		GongUnmarshallSliceOfPointers(&instance.Fingering, valueExpr, identifierMap)
	case "Pluck":
		GongUnmarshallSliceOfPointers(&instance.Pluck, valueExpr, identifierMap)
	case "Double_tongue":
		GongUnmarshallSliceOfPointers(&instance.Double_tongue, valueExpr, identifierMap)
	case "Triple_tongue":
		GongUnmarshallSliceOfPointers(&instance.Triple_tongue, valueExpr, identifierMap)
	case "Stopped":
		GongUnmarshallSliceOfPointers(&instance.Stopped, valueExpr, identifierMap)
	case "Snap_pizzicato":
		GongUnmarshallSliceOfPointers(&instance.Snap_pizzicato, valueExpr, identifierMap)
	case "Fret":
		GongUnmarshallSliceOfPointers(&instance.Fret, valueExpr, identifierMap)
	case "String_XSD":
		GongUnmarshallSliceOfPointers(&instance.String_XSD, valueExpr, identifierMap)
	case "Hammer_on":
		GongUnmarshallSliceOfPointers(&instance.Hammer_on, valueExpr, identifierMap)
	case "Pull_off":
		GongUnmarshallSliceOfPointers(&instance.Pull_off, valueExpr, identifierMap)
	case "Bend":
		GongUnmarshallSliceOfPointers(&instance.Bend, valueExpr, identifierMap)
	case "Tap":
		GongUnmarshallSliceOfPointers(&instance.Tap, valueExpr, identifierMap)
	case "Heel":
		GongUnmarshallSliceOfPointers(&instance.Heel, valueExpr, identifierMap)
	case "Toe":
		GongUnmarshallSliceOfPointers(&instance.Toe, valueExpr, identifierMap)
	case "Fingernails":
		GongUnmarshallSliceOfPointers(&instance.Fingernails, valueExpr, identifierMap)
	case "Hole":
		GongUnmarshallSliceOfPointers(&instance.Hole, valueExpr, identifierMap)
	case "Arrow":
		GongUnmarshallSliceOfPointers(&instance.Arrow, valueExpr, identifierMap)
	case "Handbell":
		GongUnmarshallSliceOfPointers(&instance.Handbell, valueExpr, identifierMap)
	case "Brass_bend":
		GongUnmarshallSliceOfPointers(&instance.Brass_bend, valueExpr, identifierMap)
	case "Flip":
		GongUnmarshallSliceOfPointers(&instance.Flip, valueExpr, identifierMap)
	case "Smear":
		GongUnmarshallSliceOfPointers(&instance.Smear, valueExpr, identifierMap)
	case "Open":
		GongUnmarshallSliceOfPointers(&instance.Open, valueExpr, identifierMap)
	case "Half_muted":
		GongUnmarshallSliceOfPointers(&instance.Half_muted, valueExpr, identifierMap)
	case "Harmon_mute":
		GongUnmarshallSliceOfPointers(&instance.Harmon_mute, valueExpr, identifierMap)
	case "Golpe":
		GongUnmarshallSliceOfPointers(&instance.Golpe, valueExpr, identifierMap)
	case "Other_technical":
		GongUnmarshallSliceOfPointers(&instance.Other_technical, valueExpr, identifierMap)
	}
	return nil
}

type Text_element_dataUnmarshaller struct{}

func (u *Text_element_dataUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Text_element_data), stage, identifier, instanceName, preserveOrder)
}

func (u *Text_element_dataUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Text_element_data)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Lang":
		instance.Lang = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Underline":
		instance.Underline = GongExtractInt(valueExpr)
	case "Overline":
		instance.Overline = GongExtractInt(valueExpr)
	case "Line_through":
		instance.Line_through = GongExtractInt(valueExpr)
	case "Rotation":
		instance.Rotation = GongExtractString(valueExpr)
	case "Letter_spacing":
		instance.Letter_spacing = GongExtractString(valueExpr)
	case "Dir":
		GongUnmarshallEnum(&instance.Dir, valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type TieUnmarshaller struct{}

func (u *TieUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Tie), stage, identifier, instanceName, preserveOrder)
}

func (u *TieUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Tie)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type_XSD":
		GongUnmarshallEnum(&instance.Type_XSD, valueExpr)
	case "Time_only":
		instance.Time_only = GongExtractString(valueExpr)
	}
	return nil
}

type TiedUnmarshaller struct{}

func (u *TiedUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Tied), stage, identifier, instanceName, preserveOrder)
}

func (u *TiedUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Tied)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type_XSD":
		GongUnmarshallEnum(&instance.Type_XSD, valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "Line_type":
		GongUnmarshallEnum(&instance.Line_type, valueExpr)
	case "Dash_length":
		instance.Dash_length = GongExtractString(valueExpr)
	case "Space_length":
		instance.Space_length = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Placement":
		GongUnmarshallEnum(&instance.Placement, valueExpr)
	case "Orientation":
		GongUnmarshallEnum(&instance.Orientation, valueExpr)
	case "Bezier_x":
		instance.Bezier_x = GongExtractString(valueExpr)
	case "Bezier_y":
		instance.Bezier_y = GongExtractString(valueExpr)
	case "Bezier_x2":
		instance.Bezier_x2 = GongExtractString(valueExpr)
	case "Bezier_y2":
		instance.Bezier_y2 = GongExtractString(valueExpr)
	case "Bezier_offset":
		instance.Bezier_offset = GongExtractString(valueExpr)
	case "Bezier_offset2":
		instance.Bezier_offset2 = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	}
	return nil
}

type TimeUnmarshaller struct{}

func (u *TimeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Time), stage, identifier, instanceName, preserveOrder)
}

func (u *TimeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Time)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "Symbol":
		GongUnmarshallEnum(&instance.Symbol, valueExpr)
	case "Separator":
		GongUnmarshallEnum(&instance.Separator, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Halign":
		GongUnmarshallEnum(&instance.Halign, valueExpr)
	case "Valign":
		GongUnmarshallEnum(&instance.Valign, valueExpr)
	case "Print_object":
		GongUnmarshallEnum(&instance.Print_object, valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Beats":
		instance.Beats = GongExtractString(valueExpr)
	case "Beat_type":
		instance.Beat_type = GongExtractString(valueExpr)
	case "Interchangeable":
		GongUnmarshallPointer(&instance.Interchangeable, valueExpr, identifierMap)
	case "Senza_misura":
		instance.Senza_misura = GongExtractString(valueExpr)
	}
	return nil
}

type Time_modificationUnmarshaller struct{}

func (u *Time_modificationUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Time_modification), stage, identifier, instanceName, preserveOrder)
}

func (u *Time_modificationUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Time_modification)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Actual_notes":
		instance.Actual_notes = GongExtractInt(valueExpr)
	case "Normal_notes":
		instance.Normal_notes = GongExtractInt(valueExpr)
	case "Normal_type":
		GongUnmarshallEnum(&instance.Normal_type, valueExpr)
	case "Normal_dot":
		instance.Normal_dot = GongExtractString(valueExpr)
	}
	return nil
}

type TimpaniUnmarshaller struct{}

func (u *TimpaniUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Timpani), stage, identifier, instanceName, preserveOrder)
}

func (u *TimpaniUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Timpani)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Smufl":
		instance.Smufl = GongExtractString(valueExpr)
	}
	return nil
}

type TransposeUnmarshaller struct{}

func (u *TransposeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Transpose), stage, identifier, instanceName, preserveOrder)
}

func (u *TransposeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Transpose)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Diatonic":
		instance.Diatonic = GongExtractInt(valueExpr)
	case "Chromatic":
		instance.Chromatic = GongExtractString(valueExpr)
	case "Octave_change":
		instance.Octave_change = GongExtractInt(valueExpr)
	case "Double":
		instance.Double = GongExtractFloat(valueExpr)
	}
	return nil
}

type TremoloUnmarshaller struct{}

func (u *TremoloUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Tremolo), stage, identifier, instanceName, preserveOrder)
}

func (u *TremoloUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Tremolo)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type_XSD":
		GongUnmarshallEnum(&instance.Type_XSD, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Placement":
		GongUnmarshallEnum(&instance.Placement, valueExpr)
	case "Smufl":
		instance.Smufl = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractInt(valueExpr)
	}
	return nil
}

type TupletUnmarshaller struct{}

func (u *TupletUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Tuplet), stage, identifier, instanceName, preserveOrder)
}

func (u *TupletUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Tuplet)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type_XSD":
		GongUnmarshallEnum(&instance.Type_XSD, valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "Bracket":
		GongUnmarshallEnum(&instance.Bracket, valueExpr)
	case "Show_number":
		GongUnmarshallEnum(&instance.Show_number, valueExpr)
	case "Show_type":
		GongUnmarshallEnum(&instance.Show_type, valueExpr)
	case "Line_shape":
		GongUnmarshallEnum(&instance.Line_shape, valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Placement":
		GongUnmarshallEnum(&instance.Placement, valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Tuplet_actual":
		GongUnmarshallPointer(&instance.Tuplet_actual, valueExpr, identifierMap)
	case "Tuplet_normal":
		GongUnmarshallPointer(&instance.Tuplet_normal, valueExpr, identifierMap)
	}
	return nil
}

type Tuplet_dotUnmarshaller struct{}

func (u *Tuplet_dotUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Tuplet_dot), stage, identifier, instanceName, preserveOrder)
}

func (u *Tuplet_dotUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Tuplet_dot)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	}
	return nil
}

type Tuplet_numberUnmarshaller struct{}

func (u *Tuplet_numberUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Tuplet_number), stage, identifier, instanceName, preserveOrder)
}

func (u *Tuplet_numberUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Tuplet_number)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractInt(valueExpr)
	}
	return nil
}

type Tuplet_portionUnmarshaller struct{}

func (u *Tuplet_portionUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Tuplet_portion), stage, identifier, instanceName, preserveOrder)
}

func (u *Tuplet_portionUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Tuplet_portion)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Tuplet_number":
		GongUnmarshallPointer(&instance.Tuplet_number, valueExpr, identifierMap)
	case "Tuplet_type":
		GongUnmarshallPointer(&instance.Tuplet_type, valueExpr, identifierMap)
	case "Tuplet_dot":
		GongUnmarshallSliceOfPointers(&instance.Tuplet_dot, valueExpr, identifierMap)
	}
	return nil
}

type Tuplet_typeUnmarshaller struct{}

func (u *Tuplet_typeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Tuplet_type), stage, identifier, instanceName, preserveOrder)
}

func (u *Tuplet_typeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Tuplet_type)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Font_family":
		instance.Font_family = GongExtractString(valueExpr)
	case "Font_style":
		GongUnmarshallEnum(&instance.Font_style, valueExpr)
	case "Font_size":
		instance.Font_size = GongExtractString(valueExpr)
	case "Font_weight":
		GongUnmarshallEnum(&instance.Font_weight, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "EnclosedText":
		GongUnmarshallEnum(&instance.EnclosedText, valueExpr)
	}
	return nil
}

type Typed_textUnmarshaller struct{}

func (u *Typed_textUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Typed_text), stage, identifier, instanceName, preserveOrder)
}

func (u *Typed_textUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Typed_text)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type_XSD":
		instance.Type_XSD = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type UnpitchedUnmarshaller struct{}

func (u *UnpitchedUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Unpitched), stage, identifier, instanceName, preserveOrder)
}

func (u *UnpitchedUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Unpitched)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Display_step":
		GongUnmarshallEnum(&instance.Display_step, valueExpr)
	case "Display_octave":
		instance.Display_octave = GongExtractInt(valueExpr)
	}
	return nil
}

type Virtual_instrumentUnmarshaller struct{}

func (u *Virtual_instrumentUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Virtual_instrument), stage, identifier, instanceName, preserveOrder)
}

func (u *Virtual_instrumentUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Virtual_instrument)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Virtual_library":
		instance.Virtual_library = GongExtractString(valueExpr)
	case "Virtual_name":
		instance.Virtual_name = GongExtractString(valueExpr)
	}
	return nil
}

type WaitUnmarshaller struct{}

func (u *WaitUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Wait), stage, identifier, instanceName, preserveOrder)
}

func (u *WaitUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Wait)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Player":
		instance.Player = GongExtractString(valueExpr)
	case "Time_only":
		instance.Time_only = GongExtractString(valueExpr)
	}
	return nil
}

type Wavy_lineUnmarshaller struct{}

func (u *Wavy_lineUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Wavy_line), stage, identifier, instanceName, preserveOrder)
}

func (u *Wavy_lineUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Wavy_line)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type_XSD":
		GongUnmarshallEnum(&instance.Type_XSD, valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "Smufl":
		instance.Smufl = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Placement":
		GongUnmarshallEnum(&instance.Placement, valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Start_note":
		GongUnmarshallEnum(&instance.Start_note, valueExpr)
	case "Trill_step":
		GongUnmarshallEnum(&instance.Trill_step, valueExpr)
	case "Two_note_turn":
		GongUnmarshallEnum(&instance.Two_note_turn, valueExpr)
	case "Accelerate":
		GongUnmarshallEnum(&instance.Accelerate, valueExpr)
	case "Beats":
		instance.Beats = GongExtractString(valueExpr)
	case "Second_beat":
		instance.Second_beat = GongExtractString(valueExpr)
	case "Last_beat":
		instance.Last_beat = GongExtractString(valueExpr)
	}
	return nil
}

type WedgeUnmarshaller struct{}

func (u *WedgeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Wedge), stage, identifier, instanceName, preserveOrder)
}

func (u *WedgeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Wedge)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Type_XSD":
		GongUnmarshallEnum(&instance.Type_XSD, valueExpr)
	case "Number":
		instance.Number = GongExtractInt(valueExpr)
	case "Spread":
		instance.Spread = GongExtractString(valueExpr)
	case "Niente":
		GongUnmarshallEnum(&instance.Niente, valueExpr)
	case "Line_type":
		GongUnmarshallEnum(&instance.Line_type, valueExpr)
	case "Dash_length":
		instance.Dash_length = GongExtractString(valueExpr)
	case "Space_length":
		instance.Space_length = GongExtractString(valueExpr)
	case "Default_x":
		instance.Default_x = GongExtractString(valueExpr)
	case "Default_y":
		instance.Default_y = GongExtractString(valueExpr)
	case "Relative_x":
		instance.Relative_x = GongExtractString(valueExpr)
	case "Relative_y":
		instance.Relative_y = GongExtractString(valueExpr)
	case "Color":
		instance.Color = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	}
	return nil
}

type WoodUnmarshaller struct{}

func (u *WoodUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Wood), stage, identifier, instanceName, preserveOrder)
}

func (u *WoodUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Wood)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Smufl":
		instance.Smufl = GongExtractString(valueExpr)
	case "EnclosedText":
		GongUnmarshallEnum(&instance.EnclosedText, valueExpr)
	}
	return nil
}

type WorkUnmarshaller struct{}

func (u *WorkUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Work), stage, identifier, instanceName, preserveOrder)
}

func (u *WorkUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Work)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Work_number":
		instance.Work_number = GongExtractString(valueExpr)
	case "Work_title":
		instance.Work_title = GongExtractString(valueExpr)
	case "Opus":
		GongUnmarshallPointer(&instance.Opus, valueExpr, identifierMap)
	}
	return nil
}
