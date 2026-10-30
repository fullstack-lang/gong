// generated code - do not edit
package aixmfeatures

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
		case "github.com/fullstack-lang/gong/app/xsd/tests/aixm/airport/go/models/aixmfeatures":
			aliasToCanonical[alias] = "aixmfeatures"
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
					} else if pkgAlias == "" || pkgAlias == "aixmfeatures" || len(inFile.Imports) == 0 {
						pkgAlias = "aixmfeatures"
					}

					switch pkgAlias {
			case "aixmfeatures":
				switch typeName {
				case "A_extension":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_1":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_1), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_10":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_10), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_11":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_11), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_12":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_12), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_13":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_13), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_14":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_14), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_15":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_15), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_16":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_16), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_17":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_17), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_18":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_18), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_19":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_19), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_2":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_2), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_20":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_20), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_21":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_21), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_22":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_22), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_23":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_23), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_24":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_24), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_25":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_25), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_26":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_26), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_27":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_27), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_28":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_28), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_29":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_29), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_3":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_3), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_30":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_30), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_31":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_31), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_32":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_32), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_33":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_33), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_34":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_34), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_35":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_35), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_36":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_36), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_37":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_37), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_38":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_38), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_39":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_39), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_4":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_4), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_40":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_40), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_41":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_41), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_42":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_42), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_43":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_43), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_44":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_44), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_45":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_45), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_46":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_46), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_47":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_47), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_48":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_48), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_49":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_49), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_5":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_5), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_50":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_50), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_51":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_51), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_52":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_52), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_53":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_53), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_54":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_54), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_55":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_55), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_56":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_56), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_57":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_57), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_58":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_58), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_59":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_59), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_6":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_6), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_60":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_60), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_61":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_61), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_62":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_62), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_63":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_63), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_64":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_64), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_65":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_65), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_66":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_66), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_67":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_67), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_68":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_68), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_69":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_69), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_7":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_7), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_70":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_70), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_71":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_71), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_72":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_72), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_73":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_73), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_74":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_74), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_75":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_75), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_76":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_76), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_77":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_77), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_78":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_78), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_79":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_79), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_8":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_8), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_80":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_80), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_81":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_81), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_82":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_82), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_83":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_83), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_84":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_84), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_85":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_85), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_86":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_86), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_87":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_87), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_88":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_88), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_89":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_89), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_9":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_9), stage, ident.Name, instanceName, preserveOrder)
				case "AbstractAirportHeliportProtectionAreaType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AbstractAirportHeliportProtectionAreaType), stage, ident.Name, instanceName, preserveOrder)
				case "AbstractGroundLightSystemType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AbstractGroundLightSystemType), stage, ident.Name, instanceName, preserveOrder)
				case "AbstractMarkingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AbstractMarkingType), stage, ident.Name, instanceName, preserveOrder)
				case "AbstractPropertiesWithScheduleType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AbstractPropertiesWithScheduleType), stage, ident.Name, instanceName, preserveOrder)
				case "AbstractSurfaceContaminationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AbstractSurfaceContaminationType), stage, ident.Name, instanceName, preserveOrder)
				case "AbstractUsageConditionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AbstractUsageConditionType), stage, ident.Name, instanceName, preserveOrder)
				case "AircraftCharacteristicPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AircraftCharacteristicPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "AircraftCharacteristicType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AircraftCharacteristicType), stage, ident.Name, instanceName, preserveOrder)
				case "AircraftStandContaminationPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AircraftStandContaminationPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "AircraftStandContaminationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AircraftStandContaminationType), stage, ident.Name, instanceName, preserveOrder)
				case "AircraftStandPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AircraftStandPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "AircraftStandTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AircraftStandTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "AircraftStandTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AircraftStandTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "AircraftStandType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AircraftStandType), stage, ident.Name, instanceName, preserveOrder)
				case "AirportHeliportAvailabilityPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AirportHeliportAvailabilityPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "AirportHeliportAvailabilityType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AirportHeliportAvailabilityType), stage, ident.Name, instanceName, preserveOrder)
				case "AirportHeliportCollocationTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AirportHeliportCollocationTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "AirportHeliportCollocationTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AirportHeliportCollocationTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "AirportHeliportCollocationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AirportHeliportCollocationType), stage, ident.Name, instanceName, preserveOrder)
				case "AirportHeliportContaminationPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AirportHeliportContaminationPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "AirportHeliportContaminationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AirportHeliportContaminationType), stage, ident.Name, instanceName, preserveOrder)
				case "AirportHeliportPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AirportHeliportPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "AirportHeliportProtectionAreaPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AirportHeliportProtectionAreaPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "AirportHeliportResponsibilityOrganisationPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AirportHeliportResponsibilityOrganisationPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "AirportHeliportResponsibilityOrganisationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AirportHeliportResponsibilityOrganisationType), stage, ident.Name, instanceName, preserveOrder)
				case "AirportHeliportTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AirportHeliportTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "AirportHeliportTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AirportHeliportTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "AirportHeliportType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AirportHeliportType), stage, ident.Name, instanceName, preserveOrder)
				case "AirportHeliportUsagePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AirportHeliportUsagePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "AirportHeliportUsageType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AirportHeliportUsageType), stage, ident.Name, instanceName, preserveOrder)
				case "AirportHotSpotTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AirportHotSpotTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "AirportHotSpotTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AirportHotSpotTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "AirportHotSpotType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AirportHotSpotType), stage, ident.Name, instanceName, preserveOrder)
				case "AirportProtectionAreaMarkingTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AirportProtectionAreaMarkingTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "AirportProtectionAreaMarkingTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AirportProtectionAreaMarkingTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "AirportProtectionAreaMarkingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AirportProtectionAreaMarkingType), stage, ident.Name, instanceName, preserveOrder)
				case "AirportSuppliesServicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AirportSuppliesServicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "AltimeterSourcePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AltimeterSourcePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "AltimeterSourceStatusPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AltimeterSourceStatusPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "AltimeterSourceStatusType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AltimeterSourceStatusType), stage, ident.Name, instanceName, preserveOrder)
				case "AltimeterSourceTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AltimeterSourceTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "AltimeterSourceTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AltimeterSourceTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "AltimeterSourceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AltimeterSourceType), stage, ident.Name, instanceName, preserveOrder)
				case "ApproachLightingSystemTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ApproachLightingSystemTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "ApproachLightingSystemTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ApproachLightingSystemTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "ApproachLightingSystemType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ApproachLightingSystemType), stage, ident.Name, instanceName, preserveOrder)
				case "ApronAreaAvailabilityPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ApronAreaAvailabilityPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "ApronAreaAvailabilityType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ApronAreaAvailabilityType), stage, ident.Name, instanceName, preserveOrder)
				case "ApronAreaUsagePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ApronAreaUsagePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "ApronAreaUsageType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ApronAreaUsageType), stage, ident.Name, instanceName, preserveOrder)
				case "ApronContaminationPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ApronContaminationPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "ApronContaminationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ApronContaminationType), stage, ident.Name, instanceName, preserveOrder)
				case "ApronElementPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ApronElementPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "ApronElementTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ApronElementTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "ApronElementTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ApronElementTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "ApronElementType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ApronElementType), stage, ident.Name, instanceName, preserveOrder)
				case "ApronLightSystemTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ApronLightSystemTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "ApronLightSystemTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ApronLightSystemTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "ApronLightSystemType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ApronLightSystemType), stage, ident.Name, instanceName, preserveOrder)
				case "ApronMarkingTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ApronMarkingTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "ApronMarkingTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ApronMarkingTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "ApronMarkingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ApronMarkingType), stage, ident.Name, instanceName, preserveOrder)
				case "ApronPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ApronPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "ApronTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ApronTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "ApronTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ApronTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "ApronType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ApronType), stage, ident.Name, instanceName, preserveOrder)
				case "ArrestingGearTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ArrestingGearTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "ArrestingGearTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ArrestingGearTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "ArrestingGearType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ArrestingGearType), stage, ident.Name, instanceName, preserveOrder)
				case "CityPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CityPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "CityType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CityType), stage, ident.Name, instanceName, preserveOrder)
				case "ConditionCombinationPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ConditionCombinationPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "ConditionCombinationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ConditionCombinationType), stage, ident.Name, instanceName, preserveOrder)
				case "ContactInformationPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ContactInformationPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "ContactInformationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ContactInformationType), stage, ident.Name, instanceName, preserveOrder)
				case "CurveType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CurveType), stage, ident.Name, instanceName, preserveOrder)
				case "DeicingAreaMarkingTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DeicingAreaMarkingTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "DeicingAreaMarkingTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DeicingAreaMarkingTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "DeicingAreaMarkingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DeicingAreaMarkingType), stage, ident.Name, instanceName, preserveOrder)
				case "DeicingAreaPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DeicingAreaPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "DeicingAreaTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DeicingAreaTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "DeicingAreaTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DeicingAreaTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "DeicingAreaType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DeicingAreaType), stage, ident.Name, instanceName, preserveOrder)
				case "ElevatedCurvePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ElevatedCurvePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "ElevatedCurveType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ElevatedCurveType), stage, ident.Name, instanceName, preserveOrder)
				case "ElevatedPointPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ElevatedPointPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "ElevatedPointType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ElevatedPointType), stage, ident.Name, instanceName, preserveOrder)
				case "ElevatedSurfacePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ElevatedSurfacePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "ElevatedSurfaceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ElevatedSurfaceType), stage, ident.Name, instanceName, preserveOrder)
				case "FlightCharacteristicPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(FlightCharacteristicPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "FlightCharacteristicType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(FlightCharacteristicType), stage, ident.Name, instanceName, preserveOrder)
				case "FloatingDockSitePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(FloatingDockSitePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "FloatingDockSiteTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(FloatingDockSiteTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "FloatingDockSiteTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(FloatingDockSiteTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "FloatingDockSiteType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(FloatingDockSiteType), stage, ident.Name, instanceName, preserveOrder)
				case "GroundLightingAvailabilityPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(GroundLightingAvailabilityPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "GroundLightingAvailabilityType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(GroundLightingAvailabilityType), stage, ident.Name, instanceName, preserveOrder)
				case "GuidanceLineLightSystemTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(GuidanceLineLightSystemTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "GuidanceLineLightSystemTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(GuidanceLineLightSystemTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "GuidanceLineLightSystemType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(GuidanceLineLightSystemType), stage, ident.Name, instanceName, preserveOrder)
				case "GuidanceLineMarkingTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(GuidanceLineMarkingTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "GuidanceLineMarkingTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(GuidanceLineMarkingTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "GuidanceLineMarkingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(GuidanceLineMarkingType), stage, ident.Name, instanceName, preserveOrder)
				case "GuidanceLinePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(GuidanceLinePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "GuidanceLineTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(GuidanceLineTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "GuidanceLineTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(GuidanceLineTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "GuidanceLineType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(GuidanceLineType), stage, ident.Name, instanceName, preserveOrder)
				case "LightElementPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(LightElementPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "LightElementStatusPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(LightElementStatusPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "LightElementStatusType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(LightElementStatusType), stage, ident.Name, instanceName, preserveOrder)
				case "LightElementType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(LightElementType), stage, ident.Name, instanceName, preserveOrder)
				case "LinguisticNotePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(LinguisticNotePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "LinguisticNoteType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(LinguisticNoteType), stage, ident.Name, instanceName, preserveOrder)
				case "ManoeuvringAreaAvailabilityPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ManoeuvringAreaAvailabilityPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "ManoeuvringAreaAvailabilityType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ManoeuvringAreaAvailabilityType), stage, ident.Name, instanceName, preserveOrder)
				case "ManoeuvringAreaUsagePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ManoeuvringAreaUsagePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "ManoeuvringAreaUsageType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ManoeuvringAreaUsageType), stage, ident.Name, instanceName, preserveOrder)
				case "MarkingBuoyTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(MarkingBuoyTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "MarkingBuoyTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(MarkingBuoyTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "MarkingBuoyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(MarkingBuoyType), stage, ident.Name, instanceName, preserveOrder)
				case "MarkingElementPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(MarkingElementPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "MarkingElementType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(MarkingElementType), stage, ident.Name, instanceName, preserveOrder)
				case "MeteorologyPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(MeteorologyPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "MeteorologyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(MeteorologyType), stage, ident.Name, instanceName, preserveOrder)
				case "NavaidEquipmentDistancePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NavaidEquipmentDistancePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "NavaidEquipmentDistanceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NavaidEquipmentDistanceType), stage, ident.Name, instanceName, preserveOrder)
				case "NavaidEquipmentPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NavaidEquipmentPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "NonMovementAreaTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NonMovementAreaTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "NonMovementAreaTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NonMovementAreaTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "NonMovementAreaType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NonMovementAreaType), stage, ident.Name, instanceName, preserveOrder)
				case "NotePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NotePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "NoteType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NoteType), stage, ident.Name, instanceName, preserveOrder)
				case "OnlineContactPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(OnlineContactPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "OnlineContactType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(OnlineContactType), stage, ident.Name, instanceName, preserveOrder)
				case "OrganisationAuthorityPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(OrganisationAuthorityPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "PassengerLoadingBridgeTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(PassengerLoadingBridgeTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "PassengerLoadingBridgeTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(PassengerLoadingBridgeTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "PassengerLoadingBridgeType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(PassengerLoadingBridgeType), stage, ident.Name, instanceName, preserveOrder)
				case "PointType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(PointType), stage, ident.Name, instanceName, preserveOrder)
				case "PostalAddressPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(PostalAddressPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "PostalAddressType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(PostalAddressType), stage, ident.Name, instanceName, preserveOrder)
				case "RidgePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RidgePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "RidgeType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RidgeType), stage, ident.Name, instanceName, preserveOrder)
				case "RoadTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RoadTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "RoadTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RoadTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "RoadType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RoadType), stage, ident.Name, instanceName, preserveOrder)
				case "RunwayBlastPadTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RunwayBlastPadTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "RunwayBlastPadTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RunwayBlastPadTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "RunwayBlastPadType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RunwayBlastPadType), stage, ident.Name, instanceName, preserveOrder)
				case "RunwayCentrelinePointPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RunwayCentrelinePointPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "RunwayCentrelinePointTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RunwayCentrelinePointTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "RunwayCentrelinePointTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RunwayCentrelinePointTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "RunwayCentrelinePointType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RunwayCentrelinePointType), stage, ident.Name, instanceName, preserveOrder)
				case "RunwayContaminationPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RunwayContaminationPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "RunwayContaminationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RunwayContaminationType), stage, ident.Name, instanceName, preserveOrder)
				case "RunwayDeclaredDistancePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RunwayDeclaredDistancePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "RunwayDeclaredDistanceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RunwayDeclaredDistanceType), stage, ident.Name, instanceName, preserveOrder)
				case "RunwayDeclaredDistanceValuePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RunwayDeclaredDistanceValuePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "RunwayDeclaredDistanceValueType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RunwayDeclaredDistanceValueType), stage, ident.Name, instanceName, preserveOrder)
				case "RunwayDirectionLightSystemTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RunwayDirectionLightSystemTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "RunwayDirectionLightSystemTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RunwayDirectionLightSystemTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "RunwayDirectionLightSystemType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RunwayDirectionLightSystemType), stage, ident.Name, instanceName, preserveOrder)
				case "RunwayDirectionPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RunwayDirectionPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "RunwayDirectionTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RunwayDirectionTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "RunwayDirectionTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RunwayDirectionTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "RunwayDirectionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RunwayDirectionType), stage, ident.Name, instanceName, preserveOrder)
				case "RunwayElementPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RunwayElementPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "RunwayElementTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RunwayElementTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "RunwayElementTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RunwayElementTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "RunwayElementType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RunwayElementType), stage, ident.Name, instanceName, preserveOrder)
				case "RunwayMarkingTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RunwayMarkingTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "RunwayMarkingTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RunwayMarkingTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "RunwayMarkingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RunwayMarkingType), stage, ident.Name, instanceName, preserveOrder)
				case "RunwayPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RunwayPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "RunwayProtectAreaLightSystemTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RunwayProtectAreaLightSystemTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "RunwayProtectAreaLightSystemTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RunwayProtectAreaLightSystemTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "RunwayProtectAreaLightSystemType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RunwayProtectAreaLightSystemType), stage, ident.Name, instanceName, preserveOrder)
				case "RunwayProtectAreaPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RunwayProtectAreaPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "RunwayProtectAreaTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RunwayProtectAreaTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "RunwayProtectAreaTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RunwayProtectAreaTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "RunwayProtectAreaType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RunwayProtectAreaType), stage, ident.Name, instanceName, preserveOrder)
				case "RunwaySectionContaminationPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RunwaySectionContaminationPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "RunwaySectionContaminationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RunwaySectionContaminationType), stage, ident.Name, instanceName, preserveOrder)
				case "RunwayTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RunwayTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "RunwayTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RunwayTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "RunwayType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RunwayType), stage, ident.Name, instanceName, preserveOrder)
				case "RunwayVisualRangeTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RunwayVisualRangeTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "RunwayVisualRangeTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RunwayVisualRangeTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "RunwayVisualRangeType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RunwayVisualRangeType), stage, ident.Name, instanceName, preserveOrder)
				case "SeaplaneLandingAreaPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SeaplaneLandingAreaPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "SeaplaneLandingAreaTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SeaplaneLandingAreaTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "SeaplaneLandingAreaTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SeaplaneLandingAreaTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "SeaplaneLandingAreaType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SeaplaneLandingAreaType), stage, ident.Name, instanceName, preserveOrder)
				case "SeaplaneRampSitePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SeaplaneRampSitePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "SeaplaneRampSiteTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SeaplaneRampSiteTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "SeaplaneRampSiteTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SeaplaneRampSiteTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "SeaplaneRampSiteType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SeaplaneRampSiteType), stage, ident.Name, instanceName, preserveOrder)
				case "StandMarkingTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(StandMarkingTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "StandMarkingTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(StandMarkingTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "StandMarkingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(StandMarkingType), stage, ident.Name, instanceName, preserveOrder)
				case "SurfaceCharacteristicsPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SurfaceCharacteristicsPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "SurfaceCharacteristicsType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SurfaceCharacteristicsType), stage, ident.Name, instanceName, preserveOrder)
				case "SurfaceContaminationLayerPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SurfaceContaminationLayerPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "SurfaceContaminationLayerType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SurfaceContaminationLayerType), stage, ident.Name, instanceName, preserveOrder)
				case "SurfaceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SurfaceType), stage, ident.Name, instanceName, preserveOrder)
				case "SurveyControlPointTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SurveyControlPointTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "SurveyControlPointTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SurveyControlPointTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "SurveyControlPointType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SurveyControlPointType), stage, ident.Name, instanceName, preserveOrder)
				case "TaxiHoldingPositionLightSystemTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TaxiHoldingPositionLightSystemTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "TaxiHoldingPositionLightSystemTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TaxiHoldingPositionLightSystemTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "TaxiHoldingPositionLightSystemType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TaxiHoldingPositionLightSystemType), stage, ident.Name, instanceName, preserveOrder)
				case "TaxiHoldingPositionMarkingTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TaxiHoldingPositionMarkingTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "TaxiHoldingPositionMarkingTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TaxiHoldingPositionMarkingTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "TaxiHoldingPositionMarkingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TaxiHoldingPositionMarkingType), stage, ident.Name, instanceName, preserveOrder)
				case "TaxiHoldingPositionPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TaxiHoldingPositionPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "TaxiHoldingPositionTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TaxiHoldingPositionTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "TaxiHoldingPositionTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TaxiHoldingPositionTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "TaxiHoldingPositionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TaxiHoldingPositionType), stage, ident.Name, instanceName, preserveOrder)
				case "TaxiwayContaminationPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TaxiwayContaminationPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "TaxiwayContaminationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TaxiwayContaminationType), stage, ident.Name, instanceName, preserveOrder)
				case "TaxiwayElementPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TaxiwayElementPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "TaxiwayElementTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TaxiwayElementTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "TaxiwayElementTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TaxiwayElementTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "TaxiwayElementType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TaxiwayElementType), stage, ident.Name, instanceName, preserveOrder)
				case "TaxiwayLightSystemTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TaxiwayLightSystemTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "TaxiwayLightSystemTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TaxiwayLightSystemTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "TaxiwayLightSystemType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TaxiwayLightSystemType), stage, ident.Name, instanceName, preserveOrder)
				case "TaxiwayMarkingTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TaxiwayMarkingTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "TaxiwayMarkingTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TaxiwayMarkingTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "TaxiwayMarkingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TaxiwayMarkingType), stage, ident.Name, instanceName, preserveOrder)
				case "TaxiwayPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TaxiwayPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "TaxiwayTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TaxiwayTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "TaxiwayTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TaxiwayTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "TaxiwayType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TaxiwayType), stage, ident.Name, instanceName, preserveOrder)
				case "TelephoneContactPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TelephoneContactPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "TelephoneContactType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TelephoneContactType), stage, ident.Name, instanceName, preserveOrder)
				case "TimesheetPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TimesheetPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "TimesheetType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TimesheetType), stage, ident.Name, instanceName, preserveOrder)
				case "TouchDownLiftOffContaminationPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TouchDownLiftOffContaminationPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "TouchDownLiftOffContaminationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TouchDownLiftOffContaminationType), stage, ident.Name, instanceName, preserveOrder)
				case "TouchDownLiftOffLightSystemTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TouchDownLiftOffLightSystemTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "TouchDownLiftOffLightSystemTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TouchDownLiftOffLightSystemTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "TouchDownLiftOffLightSystemType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TouchDownLiftOffLightSystemType), stage, ident.Name, instanceName, preserveOrder)
				case "TouchDownLiftOffMarkingTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TouchDownLiftOffMarkingTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "TouchDownLiftOffMarkingTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TouchDownLiftOffMarkingTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "TouchDownLiftOffMarkingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TouchDownLiftOffMarkingType), stage, ident.Name, instanceName, preserveOrder)
				case "TouchDownLiftOffPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TouchDownLiftOffPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "TouchDownLiftOffSafeAreaTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TouchDownLiftOffSafeAreaTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "TouchDownLiftOffSafeAreaTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TouchDownLiftOffSafeAreaTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "TouchDownLiftOffSafeAreaType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TouchDownLiftOffSafeAreaType), stage, ident.Name, instanceName, preserveOrder)
				case "TouchDownLiftOffTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TouchDownLiftOffTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "TouchDownLiftOffTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TouchDownLiftOffTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "TouchDownLiftOffType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TouchDownLiftOffType), stage, ident.Name, instanceName, preserveOrder)
				case "VisualGlideSlopeIndicatorTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(VisualGlideSlopeIndicatorTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "VisualGlideSlopeIndicatorTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(VisualGlideSlopeIndicatorTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "VisualGlideSlopeIndicatorType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(VisualGlideSlopeIndicatorType), stage, ident.Name, instanceName, preserveOrder)
				case "WorkAreaTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(WorkAreaTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "WorkAreaTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(WorkAreaTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "WorkAreaType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(WorkAreaType), stage, ident.Name, instanceName, preserveOrder)
				case "WorkareaActivityPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(WorkareaActivityPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "WorkareaActivityType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(WorkareaActivityType), stage, ident.Name, instanceName, preserveOrder)
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
				case *A_extension:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_1:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_10:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_11:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_12:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_13:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_14:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_15:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_16:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_17:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_18:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_19:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_2:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_20:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_21:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_22:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_23:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_24:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_25:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_26:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_27:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_28:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_29:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_3:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_30:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_31:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_32:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_33:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_34:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_35:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_36:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_37:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_38:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_39:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_4:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_40:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_41:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_42:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_43:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_44:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_45:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_46:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_47:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_48:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_49:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_5:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_50:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_51:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_52:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_53:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_54:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_55:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_56:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_57:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_58:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_59:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_6:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_60:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_61:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_62:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_63:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_64:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_65:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_66:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_67:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_68:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_69:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_7:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_70:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_71:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_72:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_73:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_74:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_75:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_76:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_77:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_78:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_79:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_8:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_80:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_81:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_82:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_83:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_84:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_85:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_86:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_87:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_88:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_89:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *A_extension_9:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *AbstractAirportHeliportProtectionAreaType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					}
				case *AbstractGroundLightSystemType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					}
				case *AbstractMarkingType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					}
				case *AbstractPropertiesWithScheduleType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					}
				case *AbstractSurfaceContaminationType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					}
				case *AbstractUsageConditionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					}
				case *AircraftCharacteristicPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "AircraftCharacteristic":
						__gong__assignPointer(&inst.AircraftCharacteristic, rhs, identifierMap)
					}
				case *AircraftCharacteristicType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *AircraftStandContaminationPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "AircraftStandContamination":
						__gong__assignPointer(&inst.AircraftStandContamination, rhs, identifierMap)
					}
				case *AircraftStandContaminationType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "CriticalRidge":
						__gong__assignSliceOfPointers(&inst.CriticalRidge, rhs, identifierMap)
					case "Layer":
						__gong__assignSliceOfPointers(&inst.Layer, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *AircraftStandPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
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
				case *AircraftStandTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AircraftStandTimeSlice":
						__gong__assignPointer(&inst.AircraftStandTimeSlice, rhs, identifierMap)
					}
				case *AircraftStandTimeSliceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Interpretation":
						inst.Interpretation = GongExtractString(rhs)
					case "SequenceNumber":
						inst.SequenceNumber = GongExtractInt(rhs)
					case "CorrectionNumber":
						inst.CorrectionNumber = GongExtractInt(rhs)
					case "SurfaceProperties":
						__gong__assignPointer(&inst.SurfaceProperties, rhs, identifierMap)
					case "Location":
						__gong__assignPointer(&inst.Location, rhs, identifierMap)
					case "ApronLocation":
						__gong__assignPointer(&inst.ApronLocation, rhs, identifierMap)
					case "Extent":
						__gong__assignPointer(&inst.Extent, rhs, identifierMap)
					case "Contaminant":
						__gong__assignSliceOfPointers(&inst.Contaminant, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *AircraftStandType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeSlice":
						__gong__assignSliceOfPointers(&inst.TimeSlice, rhs, identifierMap)
					}
				case *AirportHeliportAvailabilityPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "AirportHeliportAvailability":
						__gong__assignPointer(&inst.AirportHeliportAvailability, rhs, identifierMap)
					}
				case *AirportHeliportAvailabilityType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeInterval":
						__gong__assignSliceOfPointers(&inst.TimeInterval, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "SpecialDateAuthority":
						__gong__assignSliceOfPointers(&inst.SpecialDateAuthority, rhs, identifierMap)
					case "Usage":
						__gong__assignSliceOfPointers(&inst.Usage, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *AirportHeliportCollocationTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AirportHeliportCollocationTimeSlice":
						__gong__assignPointer(&inst.AirportHeliportCollocationTimeSlice, rhs, identifierMap)
					}
				case *AirportHeliportCollocationTimeSliceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Interpretation":
						inst.Interpretation = GongExtractString(rhs)
					case "SequenceNumber":
						inst.SequenceNumber = GongExtractInt(rhs)
					case "CorrectionNumber":
						inst.CorrectionNumber = GongExtractInt(rhs)
					case "HostAirport":
						__gong__assignPointer(&inst.HostAirport, rhs, identifierMap)
					case "DependentAirport":
						__gong__assignPointer(&inst.DependentAirport, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *AirportHeliportCollocationType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeSlice":
						__gong__assignSliceOfPointers(&inst.TimeSlice, rhs, identifierMap)
					}
				case *AirportHeliportContaminationPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "AirportHeliportContamination":
						__gong__assignPointer(&inst.AirportHeliportContamination, rhs, identifierMap)
					}
				case *AirportHeliportContaminationType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "CriticalRidge":
						__gong__assignSliceOfPointers(&inst.CriticalRidge, rhs, identifierMap)
					case "Layer":
						__gong__assignSliceOfPointers(&inst.Layer, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *AirportHeliportPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
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
				case *AirportHeliportProtectionAreaPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
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
				case *AirportHeliportResponsibilityOrganisationPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "AirportHeliportResponsibilityOrganisation":
						__gong__assignPointer(&inst.AirportHeliportResponsibilityOrganisation, rhs, identifierMap)
					}
				case *AirportHeliportResponsibilityOrganisationType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeInterval":
						__gong__assignSliceOfPointers(&inst.TimeInterval, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "SpecialDateAuthority":
						__gong__assignSliceOfPointers(&inst.SpecialDateAuthority, rhs, identifierMap)
					case "TheOrganisationAuthority":
						__gong__assignPointer(&inst.TheOrganisationAuthority, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *AirportHeliportTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AirportHeliportTimeSlice":
						__gong__assignPointer(&inst.AirportHeliportTimeSlice, rhs, identifierMap)
					}
				case *AirportHeliportTimeSliceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Interpretation":
						inst.Interpretation = GongExtractString(rhs)
					case "SequenceNumber":
						inst.SequenceNumber = GongExtractInt(rhs)
					case "CorrectionNumber":
						inst.CorrectionNumber = GongExtractInt(rhs)
					case "Contaminant":
						__gong__assignSliceOfPointers(&inst.Contaminant, rhs, identifierMap)
					case "ServedCity":
						__gong__assignSliceOfPointers(&inst.ServedCity, rhs, identifierMap)
					case "ResponsibleOrganisation":
						__gong__assignPointer(&inst.ResponsibleOrganisation, rhs, identifierMap)
					case "ARP":
						__gong__assignPointer(&inst.ARP, rhs, identifierMap)
					case "AviationBoundary":
						__gong__assignPointer(&inst.AviationBoundary, rhs, identifierMap)
					case "AltimeterSource":
						__gong__assignSliceOfPointers(&inst.AltimeterSource, rhs, identifierMap)
					case "Contact":
						__gong__assignSliceOfPointers(&inst.Contact, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *AirportHeliportType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeSlice":
						__gong__assignSliceOfPointers(&inst.TimeSlice, rhs, identifierMap)
					}
				case *AirportHeliportUsagePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "AirportHeliportUsage":
						__gong__assignPointer(&inst.AirportHeliportUsage, rhs, identifierMap)
					}
				case *AirportHeliportUsageType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Contact":
						__gong__assignSliceOfPointers(&inst.Contact, rhs, identifierMap)
					case "Selection":
						__gong__assignPointer(&inst.Selection, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *AirportHotSpotTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AirportHotSpotTimeSlice":
						__gong__assignPointer(&inst.AirportHotSpotTimeSlice, rhs, identifierMap)
					}
				case *AirportHotSpotTimeSliceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Interpretation":
						inst.Interpretation = GongExtractString(rhs)
					case "SequenceNumber":
						inst.SequenceNumber = GongExtractInt(rhs)
					case "CorrectionNumber":
						inst.CorrectionNumber = GongExtractInt(rhs)
					case "Area":
						__gong__assignPointer(&inst.Area, rhs, identifierMap)
					case "AffectedAirport":
						__gong__assignPointer(&inst.AffectedAirport, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *AirportHotSpotType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeSlice":
						__gong__assignSliceOfPointers(&inst.TimeSlice, rhs, identifierMap)
					}
				case *AirportProtectionAreaMarkingTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AirportProtectionAreaMarkingTimeSlice":
						__gong__assignPointer(&inst.AirportProtectionAreaMarkingTimeSlice, rhs, identifierMap)
					}
				case *AirportProtectionAreaMarkingTimeSliceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Interpretation":
						inst.Interpretation = GongExtractString(rhs)
					case "SequenceNumber":
						inst.SequenceNumber = GongExtractInt(rhs)
					case "CorrectionNumber":
						inst.CorrectionNumber = GongExtractInt(rhs)
					case "Element":
						__gong__assignSliceOfPointers(&inst.Element, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "MarkedProtectionArea":
						__gong__assignPointer(&inst.MarkedProtectionArea, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *AirportProtectionAreaMarkingType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeSlice":
						__gong__assignSliceOfPointers(&inst.TimeSlice, rhs, identifierMap)
					}
				case *AirportSuppliesServicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
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
				case *AltimeterSourcePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
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
				case *AltimeterSourceStatusPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "AltimeterSourceStatus":
						__gong__assignPointer(&inst.AltimeterSourceStatus, rhs, identifierMap)
					}
				case *AltimeterSourceStatusType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeInterval":
						__gong__assignSliceOfPointers(&inst.TimeInterval, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "SpecialDateAuthority":
						__gong__assignSliceOfPointers(&inst.SpecialDateAuthority, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *AltimeterSourceTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AltimeterSourceTimeSlice":
						__gong__assignPointer(&inst.AltimeterSourceTimeSlice, rhs, identifierMap)
					}
				case *AltimeterSourceTimeSliceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Interpretation":
						inst.Interpretation = GongExtractString(rhs)
					case "SequenceNumber":
						inst.SequenceNumber = GongExtractInt(rhs)
					case "CorrectionNumber":
						inst.CorrectionNumber = GongExtractInt(rhs)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *AltimeterSourceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeSlice":
						__gong__assignSliceOfPointers(&inst.TimeSlice, rhs, identifierMap)
					}
				case *ApproachLightingSystemTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "ApproachLightingSystemTimeSlice":
						__gong__assignPointer(&inst.ApproachLightingSystemTimeSlice, rhs, identifierMap)
					}
				case *ApproachLightingSystemTimeSliceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Interpretation":
						inst.Interpretation = GongExtractString(rhs)
					case "SequenceNumber":
						inst.SequenceNumber = GongExtractInt(rhs)
					case "CorrectionNumber":
						inst.CorrectionNumber = GongExtractInt(rhs)
					case "Element":
						__gong__assignSliceOfPointers(&inst.Element, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "ServedRunwayDirection":
						__gong__assignPointer(&inst.ServedRunwayDirection, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *ApproachLightingSystemType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeSlice":
						__gong__assignSliceOfPointers(&inst.TimeSlice, rhs, identifierMap)
					}
				case *ApronAreaAvailabilityPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "ApronAreaAvailability":
						__gong__assignPointer(&inst.ApronAreaAvailability, rhs, identifierMap)
					}
				case *ApronAreaAvailabilityType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeInterval":
						__gong__assignSliceOfPointers(&inst.TimeInterval, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "SpecialDateAuthority":
						__gong__assignSliceOfPointers(&inst.SpecialDateAuthority, rhs, identifierMap)
					case "Usage":
						__gong__assignSliceOfPointers(&inst.Usage, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *ApronAreaUsagePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "ApronAreaUsage":
						__gong__assignPointer(&inst.ApronAreaUsage, rhs, identifierMap)
					}
				case *ApronAreaUsageType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Contact":
						__gong__assignSliceOfPointers(&inst.Contact, rhs, identifierMap)
					case "Selection":
						__gong__assignPointer(&inst.Selection, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *ApronContaminationPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "ApronContamination":
						__gong__assignPointer(&inst.ApronContamination, rhs, identifierMap)
					}
				case *ApronContaminationType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "CriticalRidge":
						__gong__assignSliceOfPointers(&inst.CriticalRidge, rhs, identifierMap)
					case "Layer":
						__gong__assignSliceOfPointers(&inst.Layer, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *ApronElementPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
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
				case *ApronElementTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "ApronElementTimeSlice":
						__gong__assignPointer(&inst.ApronElementTimeSlice, rhs, identifierMap)
					}
				case *ApronElementTimeSliceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Interpretation":
						inst.Interpretation = GongExtractString(rhs)
					case "SequenceNumber":
						inst.SequenceNumber = GongExtractInt(rhs)
					case "CorrectionNumber":
						inst.CorrectionNumber = GongExtractInt(rhs)
					case "AssociatedApron":
						__gong__assignPointer(&inst.AssociatedApron, rhs, identifierMap)
					case "SurfaceProperties":
						__gong__assignPointer(&inst.SurfaceProperties, rhs, identifierMap)
					case "Extent":
						__gong__assignPointer(&inst.Extent, rhs, identifierMap)
					case "SupplyService":
						__gong__assignSliceOfPointers(&inst.SupplyService, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *ApronElementType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeSlice":
						__gong__assignSliceOfPointers(&inst.TimeSlice, rhs, identifierMap)
					}
				case *ApronLightSystemTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "ApronLightSystemTimeSlice":
						__gong__assignPointer(&inst.ApronLightSystemTimeSlice, rhs, identifierMap)
					}
				case *ApronLightSystemTimeSliceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Interpretation":
						inst.Interpretation = GongExtractString(rhs)
					case "SequenceNumber":
						inst.SequenceNumber = GongExtractInt(rhs)
					case "CorrectionNumber":
						inst.CorrectionNumber = GongExtractInt(rhs)
					case "Element":
						__gong__assignSliceOfPointers(&inst.Element, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "LightedApron":
						__gong__assignPointer(&inst.LightedApron, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *ApronLightSystemType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeSlice":
						__gong__assignSliceOfPointers(&inst.TimeSlice, rhs, identifierMap)
					}
				case *ApronMarkingTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "ApronMarkingTimeSlice":
						__gong__assignPointer(&inst.ApronMarkingTimeSlice, rhs, identifierMap)
					}
				case *ApronMarkingTimeSliceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Interpretation":
						inst.Interpretation = GongExtractString(rhs)
					case "SequenceNumber":
						inst.SequenceNumber = GongExtractInt(rhs)
					case "CorrectionNumber":
						inst.CorrectionNumber = GongExtractInt(rhs)
					case "Element":
						__gong__assignSliceOfPointers(&inst.Element, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "MarkedApron":
						__gong__assignPointer(&inst.MarkedApron, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *ApronMarkingType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeSlice":
						__gong__assignSliceOfPointers(&inst.TimeSlice, rhs, identifierMap)
					}
				case *ApronPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
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
				case *ApronTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "ApronTimeSlice":
						__gong__assignPointer(&inst.ApronTimeSlice, rhs, identifierMap)
					}
				case *ApronTimeSliceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Interpretation":
						inst.Interpretation = GongExtractString(rhs)
					case "SequenceNumber":
						inst.SequenceNumber = GongExtractInt(rhs)
					case "CorrectionNumber":
						inst.CorrectionNumber = GongExtractInt(rhs)
					case "SurfaceProperties":
						__gong__assignPointer(&inst.SurfaceProperties, rhs, identifierMap)
					case "AssociatedAirportHeliport":
						__gong__assignPointer(&inst.AssociatedAirportHeliport, rhs, identifierMap)
					case "Contaminant":
						__gong__assignSliceOfPointers(&inst.Contaminant, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *ApronType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeSlice":
						__gong__assignSliceOfPointers(&inst.TimeSlice, rhs, identifierMap)
					}
				case *ArrestingGearTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "ArrestingGearTimeSlice":
						__gong__assignPointer(&inst.ArrestingGearTimeSlice, rhs, identifierMap)
					}
				case *ArrestingGearTimeSliceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Interpretation":
						inst.Interpretation = GongExtractString(rhs)
					case "SequenceNumber":
						inst.SequenceNumber = GongExtractInt(rhs)
					case "CorrectionNumber":
						inst.CorrectionNumber = GongExtractInt(rhs)
					case "RunwayDirection":
						__gong__assignSliceOfPointers(&inst.RunwayDirection, rhs, identifierMap)
					case "SurfaceProperties":
						__gong__assignPointer(&inst.SurfaceProperties, rhs, identifierMap)
					case "Extent_surfaceExtent":
						__gong__assignPointer(&inst.Extent_surfaceExtent, rhs, identifierMap)
					case "Extent_curveExtent":
						__gong__assignPointer(&inst.Extent_curveExtent, rhs, identifierMap)
					case "Extent_pointExtent":
						__gong__assignPointer(&inst.Extent_pointExtent, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *ArrestingGearType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeSlice":
						__gong__assignSliceOfPointers(&inst.TimeSlice, rhs, identifierMap)
					}
				case *CityPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "City":
						__gong__assignPointer(&inst.City, rhs, identifierMap)
					}
				case *CityType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *ConditionCombinationPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "ConditionCombination":
						__gong__assignPointer(&inst.ConditionCombination, rhs, identifierMap)
					}
				case *ConditionCombinationType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeInterval":
						__gong__assignSliceOfPointers(&inst.TimeInterval, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "SpecialDateAuthority":
						__gong__assignSliceOfPointers(&inst.SpecialDateAuthority, rhs, identifierMap)
					case "Weather":
						__gong__assignSliceOfPointers(&inst.Weather, rhs, identifierMap)
					case "Aircraft":
						__gong__assignSliceOfPointers(&inst.Aircraft, rhs, identifierMap)
					case "Flight":
						__gong__assignSliceOfPointers(&inst.Flight, rhs, identifierMap)
					case "SubCondition":
						__gong__assignSliceOfPointers(&inst.SubCondition, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *ContactInformationPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "ContactInformation":
						__gong__assignPointer(&inst.ContactInformation, rhs, identifierMap)
					}
				case *ContactInformationType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Address":
						__gong__assignSliceOfPointers(&inst.Address, rhs, identifierMap)
					case "NetworkNode":
						__gong__assignSliceOfPointers(&inst.NetworkNode, rhs, identifierMap)
					case "PhoneFax":
						__gong__assignSliceOfPointers(&inst.PhoneFax, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *CurveType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "SrsName":
						inst.SrsName = GongExtractString(rhs)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					}
				case *DeicingAreaMarkingTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "DeicingAreaMarkingTimeSlice":
						__gong__assignPointer(&inst.DeicingAreaMarkingTimeSlice, rhs, identifierMap)
					}
				case *DeicingAreaMarkingTimeSliceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Interpretation":
						inst.Interpretation = GongExtractString(rhs)
					case "SequenceNumber":
						inst.SequenceNumber = GongExtractInt(rhs)
					case "CorrectionNumber":
						inst.CorrectionNumber = GongExtractInt(rhs)
					case "Element":
						__gong__assignSliceOfPointers(&inst.Element, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "MarkedDeicingArea":
						__gong__assignPointer(&inst.MarkedDeicingArea, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *DeicingAreaMarkingType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeSlice":
						__gong__assignSliceOfPointers(&inst.TimeSlice, rhs, identifierMap)
					}
				case *DeicingAreaPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
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
				case *DeicingAreaTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "DeicingAreaTimeSlice":
						__gong__assignPointer(&inst.DeicingAreaTimeSlice, rhs, identifierMap)
					}
				case *DeicingAreaTimeSliceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Interpretation":
						inst.Interpretation = GongExtractString(rhs)
					case "SequenceNumber":
						inst.SequenceNumber = GongExtractInt(rhs)
					case "CorrectionNumber":
						inst.CorrectionNumber = GongExtractInt(rhs)
					case "AssociatedApron":
						__gong__assignPointer(&inst.AssociatedApron, rhs, identifierMap)
					case "TaxiwayLocation":
						__gong__assignPointer(&inst.TaxiwayLocation, rhs, identifierMap)
					case "StandLocation":
						__gong__assignPointer(&inst.StandLocation, rhs, identifierMap)
					case "SurfaceProperties":
						__gong__assignPointer(&inst.SurfaceProperties, rhs, identifierMap)
					case "Extent":
						__gong__assignPointer(&inst.Extent, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *DeicingAreaType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeSlice":
						__gong__assignSliceOfPointers(&inst.TimeSlice, rhs, identifierMap)
					}
				case *ElevatedCurvePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "ElevatedCurve":
						__gong__assignPointer(&inst.ElevatedCurve, rhs, identifierMap)
					}
				case *ElevatedCurveType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "SrsName":
						inst.SrsName = GongExtractString(rhs)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *ElevatedPointPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "ElevatedPoint":
						__gong__assignPointer(&inst.ElevatedPoint, rhs, identifierMap)
					}
				case *ElevatedPointType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "SrsName":
						inst.SrsName = GongExtractString(rhs)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *ElevatedSurfacePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "ElevatedSurface":
						__gong__assignPointer(&inst.ElevatedSurface, rhs, identifierMap)
					}
				case *ElevatedSurfaceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "SrsName":
						inst.SrsName = GongExtractString(rhs)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *FlightCharacteristicPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "FlightCharacteristic":
						__gong__assignPointer(&inst.FlightCharacteristic, rhs, identifierMap)
					}
				case *FlightCharacteristicType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *FloatingDockSitePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
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
				case *FloatingDockSiteTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "FloatingDockSiteTimeSlice":
						__gong__assignPointer(&inst.FloatingDockSiteTimeSlice, rhs, identifierMap)
					}
				case *FloatingDockSiteTimeSliceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Interpretation":
						inst.Interpretation = GongExtractString(rhs)
					case "SequenceNumber":
						inst.SequenceNumber = GongExtractInt(rhs)
					case "CorrectionNumber":
						inst.CorrectionNumber = GongExtractInt(rhs)
					case "Extent":
						__gong__assignPointer(&inst.Extent, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *FloatingDockSiteType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeSlice":
						__gong__assignSliceOfPointers(&inst.TimeSlice, rhs, identifierMap)
					}
				case *GroundLightingAvailabilityPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "GroundLightingAvailability":
						__gong__assignPointer(&inst.GroundLightingAvailability, rhs, identifierMap)
					}
				case *GroundLightingAvailabilityType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeInterval":
						__gong__assignSliceOfPointers(&inst.TimeInterval, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "SpecialDateAuthority":
						__gong__assignSliceOfPointers(&inst.SpecialDateAuthority, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *GuidanceLineLightSystemTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "GuidanceLineLightSystemTimeSlice":
						__gong__assignPointer(&inst.GuidanceLineLightSystemTimeSlice, rhs, identifierMap)
					}
				case *GuidanceLineLightSystemTimeSliceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Interpretation":
						inst.Interpretation = GongExtractString(rhs)
					case "SequenceNumber":
						inst.SequenceNumber = GongExtractInt(rhs)
					case "CorrectionNumber":
						inst.CorrectionNumber = GongExtractInt(rhs)
					case "Element":
						__gong__assignSliceOfPointers(&inst.Element, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "LightedGuidanceLine":
						__gong__assignPointer(&inst.LightedGuidanceLine, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *GuidanceLineLightSystemType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeSlice":
						__gong__assignSliceOfPointers(&inst.TimeSlice, rhs, identifierMap)
					}
				case *GuidanceLineMarkingTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "GuidanceLineMarkingTimeSlice":
						__gong__assignPointer(&inst.GuidanceLineMarkingTimeSlice, rhs, identifierMap)
					}
				case *GuidanceLineMarkingTimeSliceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Interpretation":
						inst.Interpretation = GongExtractString(rhs)
					case "SequenceNumber":
						inst.SequenceNumber = GongExtractInt(rhs)
					case "CorrectionNumber":
						inst.CorrectionNumber = GongExtractInt(rhs)
					case "Element":
						__gong__assignSliceOfPointers(&inst.Element, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "MarkedGuidanceLine":
						__gong__assignPointer(&inst.MarkedGuidanceLine, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *GuidanceLineMarkingType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeSlice":
						__gong__assignSliceOfPointers(&inst.TimeSlice, rhs, identifierMap)
					}
				case *GuidanceLinePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
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
				case *GuidanceLineTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "GuidanceLineTimeSlice":
						__gong__assignPointer(&inst.GuidanceLineTimeSlice, rhs, identifierMap)
					}
				case *GuidanceLineTimeSliceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Interpretation":
						inst.Interpretation = GongExtractString(rhs)
					case "SequenceNumber":
						inst.SequenceNumber = GongExtractInt(rhs)
					case "CorrectionNumber":
						inst.CorrectionNumber = GongExtractInt(rhs)
					case "ConnectedTouchDownLiftOff":
						__gong__assignSliceOfPointers(&inst.ConnectedTouchDownLiftOff, rhs, identifierMap)
					case "ConnectedRunwayCentrelinePoint":
						__gong__assignSliceOfPointers(&inst.ConnectedRunwayCentrelinePoint, rhs, identifierMap)
					case "ConnectedApron":
						__gong__assignSliceOfPointers(&inst.ConnectedApron, rhs, identifierMap)
					case "ConnectedStand":
						__gong__assignSliceOfPointers(&inst.ConnectedStand, rhs, identifierMap)
					case "Extent":
						__gong__assignPointer(&inst.Extent, rhs, identifierMap)
					case "ConnectedTaxiway":
						__gong__assignSliceOfPointers(&inst.ConnectedTaxiway, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *GuidanceLineType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeSlice":
						__gong__assignSliceOfPointers(&inst.TimeSlice, rhs, identifierMap)
					}
				case *LightElementPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "LightElement":
						__gong__assignPointer(&inst.LightElement, rhs, identifierMap)
					}
				case *LightElementStatusPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "LightElementStatus":
						__gong__assignPointer(&inst.LightElementStatus, rhs, identifierMap)
					}
				case *LightElementStatusType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeInterval":
						__gong__assignSliceOfPointers(&inst.TimeInterval, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "SpecialDateAuthority":
						__gong__assignSliceOfPointers(&inst.SpecialDateAuthority, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *LightElementType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Location":
						__gong__assignPointer(&inst.Location, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *LinguisticNotePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "LinguisticNote":
						__gong__assignPointer(&inst.LinguisticNote, rhs, identifierMap)
					}
				case *LinguisticNoteType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *ManoeuvringAreaAvailabilityPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "ManoeuvringAreaAvailability":
						__gong__assignPointer(&inst.ManoeuvringAreaAvailability, rhs, identifierMap)
					}
				case *ManoeuvringAreaAvailabilityType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeInterval":
						__gong__assignSliceOfPointers(&inst.TimeInterval, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "SpecialDateAuthority":
						__gong__assignSliceOfPointers(&inst.SpecialDateAuthority, rhs, identifierMap)
					case "Usage":
						__gong__assignSliceOfPointers(&inst.Usage, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *ManoeuvringAreaUsagePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "ManoeuvringAreaUsage":
						__gong__assignPointer(&inst.ManoeuvringAreaUsage, rhs, identifierMap)
					}
				case *ManoeuvringAreaUsageType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Contact":
						__gong__assignSliceOfPointers(&inst.Contact, rhs, identifierMap)
					case "Selection":
						__gong__assignPointer(&inst.Selection, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *MarkingBuoyTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "MarkingBuoyTimeSlice":
						__gong__assignPointer(&inst.MarkingBuoyTimeSlice, rhs, identifierMap)
					}
				case *MarkingBuoyTimeSliceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Interpretation":
						inst.Interpretation = GongExtractString(rhs)
					case "SequenceNumber":
						inst.SequenceNumber = GongExtractInt(rhs)
					case "CorrectionNumber":
						inst.CorrectionNumber = GongExtractInt(rhs)
					case "TheSeaplaneLandingArea":
						__gong__assignPointer(&inst.TheSeaplaneLandingArea, rhs, identifierMap)
					case "Location":
						__gong__assignPointer(&inst.Location, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *MarkingBuoyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeSlice":
						__gong__assignSliceOfPointers(&inst.TimeSlice, rhs, identifierMap)
					}
				case *MarkingElementPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "MarkingElement":
						__gong__assignPointer(&inst.MarkingElement, rhs, identifierMap)
					}
				case *MarkingElementType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Extent_curveExtent":
						__gong__assignPointer(&inst.Extent_curveExtent, rhs, identifierMap)
					case "Extent_surfaceExtent":
						__gong__assignPointer(&inst.Extent_surfaceExtent, rhs, identifierMap)
					case "Extent_location":
						__gong__assignPointer(&inst.Extent_location, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *MeteorologyPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "Meteorology":
						__gong__assignPointer(&inst.Meteorology, rhs, identifierMap)
					}
				case *MeteorologyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *NavaidEquipmentDistancePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "NavaidEquipmentDistance":
						__gong__assignPointer(&inst.NavaidEquipmentDistance, rhs, identifierMap)
					}
				case *NavaidEquipmentDistanceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "TheNavaidEquipment":
						__gong__assignPointer(&inst.TheNavaidEquipment, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *NavaidEquipmentPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
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
				case *NonMovementAreaTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "NonMovementAreaTimeSlice":
						__gong__assignPointer(&inst.NonMovementAreaTimeSlice, rhs, identifierMap)
					}
				case *NonMovementAreaTimeSliceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Interpretation":
						inst.Interpretation = GongExtractString(rhs)
					case "SequenceNumber":
						inst.SequenceNumber = GongExtractInt(rhs)
					case "CorrectionNumber":
						inst.CorrectionNumber = GongExtractInt(rhs)
					case "AssociatedAirportHeliport":
						__gong__assignPointer(&inst.AssociatedAirportHeliport, rhs, identifierMap)
					case "Extent":
						__gong__assignPointer(&inst.Extent, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *NonMovementAreaType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeSlice":
						__gong__assignSliceOfPointers(&inst.TimeSlice, rhs, identifierMap)
					}
				case *NotePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "Note":
						__gong__assignPointer(&inst.Note, rhs, identifierMap)
					}
				case *NoteType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TranslatedNote":
						__gong__assignSliceOfPointers(&inst.TranslatedNote, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *OnlineContactPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "OnlineContact":
						__gong__assignPointer(&inst.OnlineContact, rhs, identifierMap)
					}
				case *OnlineContactType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeInterval":
						__gong__assignSliceOfPointers(&inst.TimeInterval, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "SpecialDateAuthority":
						__gong__assignSliceOfPointers(&inst.SpecialDateAuthority, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *OrganisationAuthorityPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
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
				case *PassengerLoadingBridgeTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "PassengerLoadingBridgeTimeSlice":
						__gong__assignPointer(&inst.PassengerLoadingBridgeTimeSlice, rhs, identifierMap)
					}
				case *PassengerLoadingBridgeTimeSliceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Interpretation":
						inst.Interpretation = GongExtractString(rhs)
					case "SequenceNumber":
						inst.SequenceNumber = GongExtractInt(rhs)
					case "CorrectionNumber":
						inst.CorrectionNumber = GongExtractInt(rhs)
					case "Extent":
						__gong__assignPointer(&inst.Extent, rhs, identifierMap)
					case "AssociatedStand":
						__gong__assignSliceOfPointers(&inst.AssociatedStand, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *PassengerLoadingBridgeType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeSlice":
						__gong__assignSliceOfPointers(&inst.TimeSlice, rhs, identifierMap)
					}
				case *PointType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "SrsName":
						inst.SrsName = GongExtractString(rhs)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					}
				case *PostalAddressPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "PostalAddress":
						__gong__assignPointer(&inst.PostalAddress, rhs, identifierMap)
					}
				case *PostalAddressType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeInterval":
						__gong__assignSliceOfPointers(&inst.TimeInterval, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "SpecialDateAuthority":
						__gong__assignSliceOfPointers(&inst.SpecialDateAuthority, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *RidgePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "Ridge":
						__gong__assignPointer(&inst.Ridge, rhs, identifierMap)
					}
				case *RidgeType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *RoadTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "RoadTimeSlice":
						__gong__assignPointer(&inst.RoadTimeSlice, rhs, identifierMap)
					}
				case *RoadTimeSliceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Interpretation":
						inst.Interpretation = GongExtractString(rhs)
					case "SequenceNumber":
						inst.SequenceNumber = GongExtractInt(rhs)
					case "CorrectionNumber":
						inst.CorrectionNumber = GongExtractInt(rhs)
					case "AssociatedAirport":
						__gong__assignPointer(&inst.AssociatedAirport, rhs, identifierMap)
					case "SurfaceProperties":
						__gong__assignPointer(&inst.SurfaceProperties, rhs, identifierMap)
					case "AccessibleStand":
						__gong__assignSliceOfPointers(&inst.AccessibleStand, rhs, identifierMap)
					case "SurfaceExtent":
						__gong__assignPointer(&inst.SurfaceExtent, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *RoadType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeSlice":
						__gong__assignSliceOfPointers(&inst.TimeSlice, rhs, identifierMap)
					}
				case *RunwayBlastPadTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "RunwayBlastPadTimeSlice":
						__gong__assignPointer(&inst.RunwayBlastPadTimeSlice, rhs, identifierMap)
					}
				case *RunwayBlastPadTimeSliceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Interpretation":
						inst.Interpretation = GongExtractString(rhs)
					case "SequenceNumber":
						inst.SequenceNumber = GongExtractInt(rhs)
					case "CorrectionNumber":
						inst.CorrectionNumber = GongExtractInt(rhs)
					case "UsedRunwayDirection":
						__gong__assignPointer(&inst.UsedRunwayDirection, rhs, identifierMap)
					case "Extent":
						__gong__assignPointer(&inst.Extent, rhs, identifierMap)
					case "SurfaceProperties":
						__gong__assignPointer(&inst.SurfaceProperties, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *RunwayBlastPadType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeSlice":
						__gong__assignSliceOfPointers(&inst.TimeSlice, rhs, identifierMap)
					}
				case *RunwayCentrelinePointPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
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
				case *RunwayCentrelinePointTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "RunwayCentrelinePointTimeSlice":
						__gong__assignPointer(&inst.RunwayCentrelinePointTimeSlice, rhs, identifierMap)
					}
				case *RunwayCentrelinePointTimeSliceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Interpretation":
						inst.Interpretation = GongExtractString(rhs)
					case "SequenceNumber":
						inst.SequenceNumber = GongExtractInt(rhs)
					case "CorrectionNumber":
						inst.CorrectionNumber = GongExtractInt(rhs)
					case "Location":
						__gong__assignPointer(&inst.Location, rhs, identifierMap)
					case "OnRunway":
						__gong__assignPointer(&inst.OnRunway, rhs, identifierMap)
					case "AssociatedDeclaredDistance":
						__gong__assignSliceOfPointers(&inst.AssociatedDeclaredDistance, rhs, identifierMap)
					case "NavaidEquipment":
						__gong__assignSliceOfPointers(&inst.NavaidEquipment, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *RunwayCentrelinePointType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeSlice":
						__gong__assignSliceOfPointers(&inst.TimeSlice, rhs, identifierMap)
					}
				case *RunwayContaminationPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "RunwayContamination":
						__gong__assignPointer(&inst.RunwayContamination, rhs, identifierMap)
					}
				case *RunwayContaminationType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "CriticalRidge":
						__gong__assignSliceOfPointers(&inst.CriticalRidge, rhs, identifierMap)
					case "Layer":
						__gong__assignSliceOfPointers(&inst.Layer, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *RunwayDeclaredDistancePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "RunwayDeclaredDistance":
						__gong__assignPointer(&inst.RunwayDeclaredDistance, rhs, identifierMap)
					}
				case *RunwayDeclaredDistanceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "DeclaredValue":
						__gong__assignSliceOfPointers(&inst.DeclaredValue, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *RunwayDeclaredDistanceValuePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "RunwayDeclaredDistanceValue":
						__gong__assignPointer(&inst.RunwayDeclaredDistanceValue, rhs, identifierMap)
					}
				case *RunwayDeclaredDistanceValueType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeInterval":
						__gong__assignSliceOfPointers(&inst.TimeInterval, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "SpecialDateAuthority":
						__gong__assignSliceOfPointers(&inst.SpecialDateAuthority, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *RunwayDirectionLightSystemTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "RunwayDirectionLightSystemTimeSlice":
						__gong__assignPointer(&inst.RunwayDirectionLightSystemTimeSlice, rhs, identifierMap)
					}
				case *RunwayDirectionLightSystemTimeSliceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Interpretation":
						inst.Interpretation = GongExtractString(rhs)
					case "SequenceNumber":
						inst.SequenceNumber = GongExtractInt(rhs)
					case "CorrectionNumber":
						inst.CorrectionNumber = GongExtractInt(rhs)
					case "Element":
						__gong__assignSliceOfPointers(&inst.Element, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "AssociatedRunwayDirection":
						__gong__assignPointer(&inst.AssociatedRunwayDirection, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *RunwayDirectionLightSystemType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeSlice":
						__gong__assignSliceOfPointers(&inst.TimeSlice, rhs, identifierMap)
					}
				case *RunwayDirectionPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
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
				case *RunwayDirectionTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "RunwayDirectionTimeSlice":
						__gong__assignPointer(&inst.RunwayDirectionTimeSlice, rhs, identifierMap)
					}
				case *RunwayDirectionTimeSliceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Interpretation":
						inst.Interpretation = GongExtractString(rhs)
					case "SequenceNumber":
						inst.SequenceNumber = GongExtractInt(rhs)
					case "CorrectionNumber":
						inst.CorrectionNumber = GongExtractInt(rhs)
					case "UsedRunway":
						__gong__assignPointer(&inst.UsedRunway, rhs, identifierMap)
					case "StartingElement":
						__gong__assignPointer(&inst.StartingElement, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *RunwayDirectionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeSlice":
						__gong__assignSliceOfPointers(&inst.TimeSlice, rhs, identifierMap)
					}
				case *RunwayElementPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
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
				case *RunwayElementTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "RunwayElementTimeSlice":
						__gong__assignPointer(&inst.RunwayElementTimeSlice, rhs, identifierMap)
					}
				case *RunwayElementTimeSliceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Interpretation":
						inst.Interpretation = GongExtractString(rhs)
					case "SequenceNumber":
						inst.SequenceNumber = GongExtractInt(rhs)
					case "CorrectionNumber":
						inst.CorrectionNumber = GongExtractInt(rhs)
					case "SurfaceProperties":
						__gong__assignPointer(&inst.SurfaceProperties, rhs, identifierMap)
					case "AssociatedRunway":
						__gong__assignSliceOfPointers(&inst.AssociatedRunway, rhs, identifierMap)
					case "Extent":
						__gong__assignPointer(&inst.Extent, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *RunwayElementType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeSlice":
						__gong__assignSliceOfPointers(&inst.TimeSlice, rhs, identifierMap)
					}
				case *RunwayMarkingTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "RunwayMarkingTimeSlice":
						__gong__assignPointer(&inst.RunwayMarkingTimeSlice, rhs, identifierMap)
					}
				case *RunwayMarkingTimeSliceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Interpretation":
						inst.Interpretation = GongExtractString(rhs)
					case "SequenceNumber":
						inst.SequenceNumber = GongExtractInt(rhs)
					case "CorrectionNumber":
						inst.CorrectionNumber = GongExtractInt(rhs)
					case "Element":
						__gong__assignSliceOfPointers(&inst.Element, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "MarkedRunway":
						__gong__assignPointer(&inst.MarkedRunway, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *RunwayMarkingType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeSlice":
						__gong__assignSliceOfPointers(&inst.TimeSlice, rhs, identifierMap)
					}
				case *RunwayPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
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
				case *RunwayProtectAreaLightSystemTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "RunwayProtectAreaLightSystemTimeSlice":
						__gong__assignPointer(&inst.RunwayProtectAreaLightSystemTimeSlice, rhs, identifierMap)
					}
				case *RunwayProtectAreaLightSystemTimeSliceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Interpretation":
						inst.Interpretation = GongExtractString(rhs)
					case "SequenceNumber":
						inst.SequenceNumber = GongExtractInt(rhs)
					case "CorrectionNumber":
						inst.CorrectionNumber = GongExtractInt(rhs)
					case "Element":
						__gong__assignSliceOfPointers(&inst.Element, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "LightedArea":
						__gong__assignPointer(&inst.LightedArea, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *RunwayProtectAreaLightSystemType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeSlice":
						__gong__assignSliceOfPointers(&inst.TimeSlice, rhs, identifierMap)
					}
				case *RunwayProtectAreaPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
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
				case *RunwayProtectAreaTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "RunwayProtectAreaTimeSlice":
						__gong__assignPointer(&inst.RunwayProtectAreaTimeSlice, rhs, identifierMap)
					}
				case *RunwayProtectAreaTimeSliceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Interpretation":
						inst.Interpretation = GongExtractString(rhs)
					case "SequenceNumber":
						inst.SequenceNumber = GongExtractInt(rhs)
					case "CorrectionNumber":
						inst.CorrectionNumber = GongExtractInt(rhs)
					case "SurfaceProperties":
						__gong__assignPointer(&inst.SurfaceProperties, rhs, identifierMap)
					case "Extent":
						__gong__assignPointer(&inst.Extent, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "ProtectedRunwayDirection":
						__gong__assignPointer(&inst.ProtectedRunwayDirection, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *RunwayProtectAreaType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeSlice":
						__gong__assignSliceOfPointers(&inst.TimeSlice, rhs, identifierMap)
					}
				case *RunwaySectionContaminationPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "RunwaySectionContamination":
						__gong__assignPointer(&inst.RunwaySectionContamination, rhs, identifierMap)
					}
				case *RunwaySectionContaminationType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "CriticalRidge":
						__gong__assignSliceOfPointers(&inst.CriticalRidge, rhs, identifierMap)
					case "Layer":
						__gong__assignSliceOfPointers(&inst.Layer, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *RunwayTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "RunwayTimeSlice":
						__gong__assignPointer(&inst.RunwayTimeSlice, rhs, identifierMap)
					}
				case *RunwayTimeSliceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Interpretation":
						inst.Interpretation = GongExtractString(rhs)
					case "SequenceNumber":
						inst.SequenceNumber = GongExtractInt(rhs)
					case "CorrectionNumber":
						inst.CorrectionNumber = GongExtractInt(rhs)
					case "SurfaceProperties":
						__gong__assignPointer(&inst.SurfaceProperties, rhs, identifierMap)
					case "AssociatedAirportHeliport":
						__gong__assignPointer(&inst.AssociatedAirportHeliport, rhs, identifierMap)
					case "OverallContaminant":
						__gong__assignSliceOfPointers(&inst.OverallContaminant, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "AreaContaminant":
						__gong__assignSliceOfPointers(&inst.AreaContaminant, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *RunwayType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeSlice":
						__gong__assignSliceOfPointers(&inst.TimeSlice, rhs, identifierMap)
					}
				case *RunwayVisualRangeTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "RunwayVisualRangeTimeSlice":
						__gong__assignPointer(&inst.RunwayVisualRangeTimeSlice, rhs, identifierMap)
					}
				case *RunwayVisualRangeTimeSliceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Interpretation":
						inst.Interpretation = GongExtractString(rhs)
					case "SequenceNumber":
						inst.SequenceNumber = GongExtractInt(rhs)
					case "CorrectionNumber":
						inst.CorrectionNumber = GongExtractInt(rhs)
					case "AssociatedRunwayDirection":
						__gong__assignSliceOfPointers(&inst.AssociatedRunwayDirection, rhs, identifierMap)
					case "Location":
						__gong__assignPointer(&inst.Location, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *RunwayVisualRangeType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeSlice":
						__gong__assignSliceOfPointers(&inst.TimeSlice, rhs, identifierMap)
					}
				case *SeaplaneLandingAreaPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
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
				case *SeaplaneLandingAreaTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "SeaplaneLandingAreaTimeSlice":
						__gong__assignPointer(&inst.SeaplaneLandingAreaTimeSlice, rhs, identifierMap)
					}
				case *SeaplaneLandingAreaTimeSliceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Interpretation":
						inst.Interpretation = GongExtractString(rhs)
					case "SequenceNumber":
						inst.SequenceNumber = GongExtractInt(rhs)
					case "CorrectionNumber":
						inst.CorrectionNumber = GongExtractInt(rhs)
					case "RampSite":
						__gong__assignSliceOfPointers(&inst.RampSite, rhs, identifierMap)
					case "DockSite":
						__gong__assignSliceOfPointers(&inst.DockSite, rhs, identifierMap)
					case "Extent":
						__gong__assignPointer(&inst.Extent, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *SeaplaneLandingAreaType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeSlice":
						__gong__assignSliceOfPointers(&inst.TimeSlice, rhs, identifierMap)
					}
				case *SeaplaneRampSitePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
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
				case *SeaplaneRampSiteTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "SeaplaneRampSiteTimeSlice":
						__gong__assignPointer(&inst.SeaplaneRampSiteTimeSlice, rhs, identifierMap)
					}
				case *SeaplaneRampSiteTimeSliceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Interpretation":
						inst.Interpretation = GongExtractString(rhs)
					case "SequenceNumber":
						inst.SequenceNumber = GongExtractInt(rhs)
					case "CorrectionNumber":
						inst.CorrectionNumber = GongExtractInt(rhs)
					case "Extent":
						__gong__assignPointer(&inst.Extent, rhs, identifierMap)
					case "Centreline":
						__gong__assignPointer(&inst.Centreline, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *SeaplaneRampSiteType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeSlice":
						__gong__assignSliceOfPointers(&inst.TimeSlice, rhs, identifierMap)
					}
				case *StandMarkingTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "StandMarkingTimeSlice":
						__gong__assignPointer(&inst.StandMarkingTimeSlice, rhs, identifierMap)
					}
				case *StandMarkingTimeSliceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Interpretation":
						inst.Interpretation = GongExtractString(rhs)
					case "SequenceNumber":
						inst.SequenceNumber = GongExtractInt(rhs)
					case "CorrectionNumber":
						inst.CorrectionNumber = GongExtractInt(rhs)
					case "Element":
						__gong__assignSliceOfPointers(&inst.Element, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "MarkedStand":
						__gong__assignPointer(&inst.MarkedStand, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *StandMarkingType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeSlice":
						__gong__assignSliceOfPointers(&inst.TimeSlice, rhs, identifierMap)
					}
				case *SurfaceCharacteristicsPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "SurfaceCharacteristics":
						__gong__assignPointer(&inst.SurfaceCharacteristics, rhs, identifierMap)
					}
				case *SurfaceCharacteristicsType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *SurfaceContaminationLayerPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "SurfaceContaminationLayer":
						__gong__assignPointer(&inst.SurfaceContaminationLayer, rhs, identifierMap)
					}
				case *SurfaceContaminationLayerType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Extent":
						__gong__assignSliceOfPointers(&inst.Extent, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *SurfaceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "SrsName":
						inst.SrsName = GongExtractString(rhs)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					}
				case *SurveyControlPointTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "SurveyControlPointTimeSlice":
						__gong__assignPointer(&inst.SurveyControlPointTimeSlice, rhs, identifierMap)
					}
				case *SurveyControlPointTimeSliceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Interpretation":
						inst.Interpretation = GongExtractString(rhs)
					case "SequenceNumber":
						inst.SequenceNumber = GongExtractInt(rhs)
					case "CorrectionNumber":
						inst.CorrectionNumber = GongExtractInt(rhs)
					case "AssociatedAirportHeliport":
						__gong__assignPointer(&inst.AssociatedAirportHeliport, rhs, identifierMap)
					case "Location":
						__gong__assignPointer(&inst.Location, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *SurveyControlPointType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeSlice":
						__gong__assignSliceOfPointers(&inst.TimeSlice, rhs, identifierMap)
					}
				case *TaxiHoldingPositionLightSystemTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "TaxiHoldingPositionLightSystemTimeSlice":
						__gong__assignPointer(&inst.TaxiHoldingPositionLightSystemTimeSlice, rhs, identifierMap)
					}
				case *TaxiHoldingPositionLightSystemTimeSliceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Interpretation":
						inst.Interpretation = GongExtractString(rhs)
					case "SequenceNumber":
						inst.SequenceNumber = GongExtractInt(rhs)
					case "CorrectionNumber":
						inst.CorrectionNumber = GongExtractInt(rhs)
					case "Element":
						__gong__assignSliceOfPointers(&inst.Element, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "TaxiHolding":
						__gong__assignPointer(&inst.TaxiHolding, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *TaxiHoldingPositionLightSystemType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeSlice":
						__gong__assignSliceOfPointers(&inst.TimeSlice, rhs, identifierMap)
					}
				case *TaxiHoldingPositionMarkingTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "TaxiHoldingPositionMarkingTimeSlice":
						__gong__assignPointer(&inst.TaxiHoldingPositionMarkingTimeSlice, rhs, identifierMap)
					}
				case *TaxiHoldingPositionMarkingTimeSliceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Interpretation":
						inst.Interpretation = GongExtractString(rhs)
					case "SequenceNumber":
						inst.SequenceNumber = GongExtractInt(rhs)
					case "CorrectionNumber":
						inst.CorrectionNumber = GongExtractInt(rhs)
					case "Element":
						__gong__assignSliceOfPointers(&inst.Element, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "MarkedTaxiHold":
						__gong__assignPointer(&inst.MarkedTaxiHold, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *TaxiHoldingPositionMarkingType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeSlice":
						__gong__assignSliceOfPointers(&inst.TimeSlice, rhs, identifierMap)
					}
				case *TaxiHoldingPositionPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
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
				case *TaxiHoldingPositionTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "TaxiHoldingPositionTimeSlice":
						__gong__assignPointer(&inst.TaxiHoldingPositionTimeSlice, rhs, identifierMap)
					}
				case *TaxiHoldingPositionTimeSliceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Interpretation":
						inst.Interpretation = GongExtractString(rhs)
					case "SequenceNumber":
						inst.SequenceNumber = GongExtractInt(rhs)
					case "CorrectionNumber":
						inst.CorrectionNumber = GongExtractInt(rhs)
					case "AssociatedGuidanceLine":
						__gong__assignPointer(&inst.AssociatedGuidanceLine, rhs, identifierMap)
					case "ProtectedRunway":
						__gong__assignSliceOfPointers(&inst.ProtectedRunway, rhs, identifierMap)
					case "Location":
						__gong__assignPointer(&inst.Location, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *TaxiHoldingPositionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeSlice":
						__gong__assignSliceOfPointers(&inst.TimeSlice, rhs, identifierMap)
					}
				case *TaxiwayContaminationPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "TaxiwayContamination":
						__gong__assignPointer(&inst.TaxiwayContamination, rhs, identifierMap)
					}
				case *TaxiwayContaminationType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "CriticalRidge":
						__gong__assignSliceOfPointers(&inst.CriticalRidge, rhs, identifierMap)
					case "Layer":
						__gong__assignSliceOfPointers(&inst.Layer, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *TaxiwayElementPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
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
				case *TaxiwayElementTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "TaxiwayElementTimeSlice":
						__gong__assignPointer(&inst.TaxiwayElementTimeSlice, rhs, identifierMap)
					}
				case *TaxiwayElementTimeSliceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Interpretation":
						inst.Interpretation = GongExtractString(rhs)
					case "SequenceNumber":
						inst.SequenceNumber = GongExtractInt(rhs)
					case "CorrectionNumber":
						inst.CorrectionNumber = GongExtractInt(rhs)
					case "SurfaceProperties":
						__gong__assignPointer(&inst.SurfaceProperties, rhs, identifierMap)
					case "AssociatedTaxiway":
						__gong__assignPointer(&inst.AssociatedTaxiway, rhs, identifierMap)
					case "Extent":
						__gong__assignPointer(&inst.Extent, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *TaxiwayElementType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeSlice":
						__gong__assignSliceOfPointers(&inst.TimeSlice, rhs, identifierMap)
					}
				case *TaxiwayLightSystemTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "TaxiwayLightSystemTimeSlice":
						__gong__assignPointer(&inst.TaxiwayLightSystemTimeSlice, rhs, identifierMap)
					}
				case *TaxiwayLightSystemTimeSliceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Interpretation":
						inst.Interpretation = GongExtractString(rhs)
					case "SequenceNumber":
						inst.SequenceNumber = GongExtractInt(rhs)
					case "CorrectionNumber":
						inst.CorrectionNumber = GongExtractInt(rhs)
					case "Element":
						__gong__assignSliceOfPointers(&inst.Element, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "LightedTaxiway":
						__gong__assignPointer(&inst.LightedTaxiway, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *TaxiwayLightSystemType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeSlice":
						__gong__assignSliceOfPointers(&inst.TimeSlice, rhs, identifierMap)
					}
				case *TaxiwayMarkingTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "TaxiwayMarkingTimeSlice":
						__gong__assignPointer(&inst.TaxiwayMarkingTimeSlice, rhs, identifierMap)
					}
				case *TaxiwayMarkingTimeSliceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Interpretation":
						inst.Interpretation = GongExtractString(rhs)
					case "SequenceNumber":
						inst.SequenceNumber = GongExtractInt(rhs)
					case "CorrectionNumber":
						inst.CorrectionNumber = GongExtractInt(rhs)
					case "Element":
						__gong__assignSliceOfPointers(&inst.Element, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "MarkedTaxiway":
						__gong__assignPointer(&inst.MarkedTaxiway, rhs, identifierMap)
					case "MarkedElement":
						__gong__assignPointer(&inst.MarkedElement, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *TaxiwayMarkingType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeSlice":
						__gong__assignSliceOfPointers(&inst.TimeSlice, rhs, identifierMap)
					}
				case *TaxiwayPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
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
				case *TaxiwayTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "TaxiwayTimeSlice":
						__gong__assignPointer(&inst.TaxiwayTimeSlice, rhs, identifierMap)
					}
				case *TaxiwayTimeSliceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Interpretation":
						inst.Interpretation = GongExtractString(rhs)
					case "SequenceNumber":
						inst.SequenceNumber = GongExtractInt(rhs)
					case "CorrectionNumber":
						inst.CorrectionNumber = GongExtractInt(rhs)
					case "SurfaceProperties":
						__gong__assignPointer(&inst.SurfaceProperties, rhs, identifierMap)
					case "AssociatedAirportHeliport":
						__gong__assignPointer(&inst.AssociatedAirportHeliport, rhs, identifierMap)
					case "Contaminant":
						__gong__assignSliceOfPointers(&inst.Contaminant, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *TaxiwayType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeSlice":
						__gong__assignSliceOfPointers(&inst.TimeSlice, rhs, identifierMap)
					}
				case *TelephoneContactPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "TelephoneContact":
						__gong__assignPointer(&inst.TelephoneContact, rhs, identifierMap)
					}
				case *TelephoneContactType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeInterval":
						__gong__assignSliceOfPointers(&inst.TimeInterval, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "SpecialDateAuthority":
						__gong__assignSliceOfPointers(&inst.SpecialDateAuthority, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *TimesheetPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "Timesheet":
						__gong__assignPointer(&inst.Timesheet, rhs, identifierMap)
					}
				case *TimesheetType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *TouchDownLiftOffContaminationPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "TouchDownLiftOffContamination":
						__gong__assignPointer(&inst.TouchDownLiftOffContamination, rhs, identifierMap)
					}
				case *TouchDownLiftOffContaminationType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "CriticalRidge":
						__gong__assignSliceOfPointers(&inst.CriticalRidge, rhs, identifierMap)
					case "Layer":
						__gong__assignSliceOfPointers(&inst.Layer, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *TouchDownLiftOffLightSystemTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "TouchDownLiftOffLightSystemTimeSlice":
						__gong__assignPointer(&inst.TouchDownLiftOffLightSystemTimeSlice, rhs, identifierMap)
					}
				case *TouchDownLiftOffLightSystemTimeSliceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Interpretation":
						inst.Interpretation = GongExtractString(rhs)
					case "SequenceNumber":
						inst.SequenceNumber = GongExtractInt(rhs)
					case "CorrectionNumber":
						inst.CorrectionNumber = GongExtractInt(rhs)
					case "Element":
						__gong__assignSliceOfPointers(&inst.Element, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "LightedTouchDownLiftOff":
						__gong__assignPointer(&inst.LightedTouchDownLiftOff, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *TouchDownLiftOffLightSystemType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeSlice":
						__gong__assignSliceOfPointers(&inst.TimeSlice, rhs, identifierMap)
					}
				case *TouchDownLiftOffMarkingTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "TouchDownLiftOffMarkingTimeSlice":
						__gong__assignPointer(&inst.TouchDownLiftOffMarkingTimeSlice, rhs, identifierMap)
					}
				case *TouchDownLiftOffMarkingTimeSliceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Interpretation":
						inst.Interpretation = GongExtractString(rhs)
					case "SequenceNumber":
						inst.SequenceNumber = GongExtractInt(rhs)
					case "CorrectionNumber":
						inst.CorrectionNumber = GongExtractInt(rhs)
					case "Element":
						__gong__assignSliceOfPointers(&inst.Element, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "MarkedTouchDownLiftOff":
						__gong__assignPointer(&inst.MarkedTouchDownLiftOff, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *TouchDownLiftOffMarkingType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeSlice":
						__gong__assignSliceOfPointers(&inst.TimeSlice, rhs, identifierMap)
					}
				case *TouchDownLiftOffPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
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
				case *TouchDownLiftOffSafeAreaTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "TouchDownLiftOffSafeAreaTimeSlice":
						__gong__assignPointer(&inst.TouchDownLiftOffSafeAreaTimeSlice, rhs, identifierMap)
					}
				case *TouchDownLiftOffSafeAreaTimeSliceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Interpretation":
						inst.Interpretation = GongExtractString(rhs)
					case "SequenceNumber":
						inst.SequenceNumber = GongExtractInt(rhs)
					case "CorrectionNumber":
						inst.CorrectionNumber = GongExtractInt(rhs)
					case "SurfaceProperties":
						__gong__assignPointer(&inst.SurfaceProperties, rhs, identifierMap)
					case "Extent":
						__gong__assignPointer(&inst.Extent, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "ProtectedTouchDownLiftOff":
						__gong__assignPointer(&inst.ProtectedTouchDownLiftOff, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *TouchDownLiftOffSafeAreaType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeSlice":
						__gong__assignSliceOfPointers(&inst.TimeSlice, rhs, identifierMap)
					}
				case *TouchDownLiftOffTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "TouchDownLiftOffTimeSlice":
						__gong__assignPointer(&inst.TouchDownLiftOffTimeSlice, rhs, identifierMap)
					}
				case *TouchDownLiftOffTimeSliceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Interpretation":
						inst.Interpretation = GongExtractString(rhs)
					case "SequenceNumber":
						inst.SequenceNumber = GongExtractInt(rhs)
					case "CorrectionNumber":
						inst.CorrectionNumber = GongExtractInt(rhs)
					case "AimingPoint":
						__gong__assignPointer(&inst.AimingPoint, rhs, identifierMap)
					case "Extent":
						__gong__assignPointer(&inst.Extent, rhs, identifierMap)
					case "SurfaceProperties":
						__gong__assignPointer(&inst.SurfaceProperties, rhs, identifierMap)
					case "AssociatedAirportHeliport":
						__gong__assignPointer(&inst.AssociatedAirportHeliport, rhs, identifierMap)
					case "ApproachTakeOffArea":
						__gong__assignPointer(&inst.ApproachTakeOffArea, rhs, identifierMap)
					case "Contaminant":
						__gong__assignSliceOfPointers(&inst.Contaminant, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *TouchDownLiftOffType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeSlice":
						__gong__assignSliceOfPointers(&inst.TimeSlice, rhs, identifierMap)
					}
				case *VisualGlideSlopeIndicatorTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "VisualGlideSlopeIndicatorTimeSlice":
						__gong__assignPointer(&inst.VisualGlideSlopeIndicatorTimeSlice, rhs, identifierMap)
					}
				case *VisualGlideSlopeIndicatorTimeSliceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Interpretation":
						inst.Interpretation = GongExtractString(rhs)
					case "SequenceNumber":
						inst.SequenceNumber = GongExtractInt(rhs)
					case "CorrectionNumber":
						inst.CorrectionNumber = GongExtractInt(rhs)
					case "Element":
						__gong__assignSliceOfPointers(&inst.Element, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "RunwayDirection":
						__gong__assignPointer(&inst.RunwayDirection, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *VisualGlideSlopeIndicatorType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeSlice":
						__gong__assignSliceOfPointers(&inst.TimeSlice, rhs, identifierMap)
					}
				case *WorkAreaTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "WorkAreaTimeSlice":
						__gong__assignPointer(&inst.WorkAreaTimeSlice, rhs, identifierMap)
					}
				case *WorkAreaTimeSliceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Interpretation":
						inst.Interpretation = GongExtractString(rhs)
					case "SequenceNumber":
						inst.SequenceNumber = GongExtractInt(rhs)
					case "CorrectionNumber":
						inst.CorrectionNumber = GongExtractInt(rhs)
					case "AssociatedAirportHeliport":
						__gong__assignPointer(&inst.AssociatedAirportHeliport, rhs, identifierMap)
					case "Extent":
						__gong__assignPointer(&inst.Extent, rhs, identifierMap)
					case "Activation":
						__gong__assignSliceOfPointers(&inst.Activation, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *WorkAreaType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeSlice":
						__gong__assignSliceOfPointers(&inst.TimeSlice, rhs, identifierMap)
					}
				case *WorkareaActivityPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "WorkareaActivity":
						__gong__assignPointer(&inst.WorkareaActivity, rhs, identifierMap)
					}
				case *WorkareaActivityType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TimeInterval":
						__gong__assignSliceOfPointers(&inst.TimeInterval, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "SpecialDateAuthority":
						__gong__assignSliceOfPointers(&inst.SpecialDateAuthority, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
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
type A_extensionUnmarshaller struct{}

func (u *A_extensionUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extensionUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_1Unmarshaller struct{}

func (u *A_extension_1Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_1), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_1Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_1)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_10Unmarshaller struct{}

func (u *A_extension_10Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_10), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_10Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_10)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_11Unmarshaller struct{}

func (u *A_extension_11Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_11), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_11Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_11)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_12Unmarshaller struct{}

func (u *A_extension_12Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_12), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_12Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_12)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_13Unmarshaller struct{}

func (u *A_extension_13Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_13), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_13Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_13)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_14Unmarshaller struct{}

func (u *A_extension_14Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_14), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_14Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_14)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_15Unmarshaller struct{}

func (u *A_extension_15Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_15), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_15Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_15)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_16Unmarshaller struct{}

func (u *A_extension_16Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_16), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_16Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_16)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_17Unmarshaller struct{}

func (u *A_extension_17Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_17), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_17Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_17)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_18Unmarshaller struct{}

func (u *A_extension_18Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_18), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_18Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_18)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_19Unmarshaller struct{}

func (u *A_extension_19Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_19), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_19Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_19)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_2Unmarshaller struct{}

func (u *A_extension_2Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_2), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_2Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_2)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_20Unmarshaller struct{}

func (u *A_extension_20Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_20), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_20Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_20)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_21Unmarshaller struct{}

func (u *A_extension_21Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_21), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_21Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_21)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_22Unmarshaller struct{}

func (u *A_extension_22Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_22), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_22Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_22)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_23Unmarshaller struct{}

func (u *A_extension_23Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_23), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_23Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_23)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_24Unmarshaller struct{}

func (u *A_extension_24Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_24), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_24Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_24)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_25Unmarshaller struct{}

func (u *A_extension_25Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_25), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_25Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_25)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_26Unmarshaller struct{}

func (u *A_extension_26Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_26), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_26Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_26)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_27Unmarshaller struct{}

func (u *A_extension_27Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_27), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_27Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_27)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_28Unmarshaller struct{}

func (u *A_extension_28Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_28), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_28Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_28)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_29Unmarshaller struct{}

func (u *A_extension_29Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_29), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_29Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_29)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_3Unmarshaller struct{}

func (u *A_extension_3Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_3), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_3Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_3)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_30Unmarshaller struct{}

func (u *A_extension_30Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_30), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_30Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_30)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_31Unmarshaller struct{}

func (u *A_extension_31Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_31), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_31Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_31)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_32Unmarshaller struct{}

func (u *A_extension_32Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_32), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_32Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_32)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_33Unmarshaller struct{}

func (u *A_extension_33Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_33), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_33Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_33)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_34Unmarshaller struct{}

func (u *A_extension_34Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_34), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_34Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_34)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_35Unmarshaller struct{}

func (u *A_extension_35Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_35), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_35Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_35)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_36Unmarshaller struct{}

func (u *A_extension_36Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_36), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_36Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_36)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_37Unmarshaller struct{}

func (u *A_extension_37Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_37), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_37Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_37)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_38Unmarshaller struct{}

func (u *A_extension_38Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_38), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_38Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_38)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_39Unmarshaller struct{}

func (u *A_extension_39Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_39), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_39Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_39)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_4Unmarshaller struct{}

func (u *A_extension_4Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_4), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_4Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_4)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_40Unmarshaller struct{}

func (u *A_extension_40Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_40), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_40Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_40)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_41Unmarshaller struct{}

func (u *A_extension_41Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_41), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_41Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_41)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_42Unmarshaller struct{}

func (u *A_extension_42Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_42), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_42Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_42)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_43Unmarshaller struct{}

func (u *A_extension_43Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_43), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_43Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_43)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_44Unmarshaller struct{}

func (u *A_extension_44Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_44), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_44Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_44)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_45Unmarshaller struct{}

func (u *A_extension_45Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_45), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_45Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_45)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_46Unmarshaller struct{}

func (u *A_extension_46Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_46), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_46Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_46)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_47Unmarshaller struct{}

func (u *A_extension_47Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_47), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_47Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_47)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_48Unmarshaller struct{}

func (u *A_extension_48Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_48), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_48Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_48)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_49Unmarshaller struct{}

func (u *A_extension_49Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_49), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_49Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_49)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_5Unmarshaller struct{}

func (u *A_extension_5Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_5), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_5Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_5)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_50Unmarshaller struct{}

func (u *A_extension_50Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_50), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_50Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_50)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_51Unmarshaller struct{}

func (u *A_extension_51Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_51), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_51Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_51)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_52Unmarshaller struct{}

func (u *A_extension_52Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_52), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_52Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_52)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_53Unmarshaller struct{}

func (u *A_extension_53Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_53), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_53Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_53)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_54Unmarshaller struct{}

func (u *A_extension_54Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_54), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_54Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_54)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_55Unmarshaller struct{}

func (u *A_extension_55Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_55), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_55Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_55)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_56Unmarshaller struct{}

func (u *A_extension_56Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_56), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_56Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_56)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_57Unmarshaller struct{}

func (u *A_extension_57Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_57), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_57Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_57)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_58Unmarshaller struct{}

func (u *A_extension_58Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_58), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_58Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_58)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_59Unmarshaller struct{}

func (u *A_extension_59Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_59), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_59Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_59)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_6Unmarshaller struct{}

func (u *A_extension_6Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_6), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_6Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_6)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_60Unmarshaller struct{}

func (u *A_extension_60Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_60), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_60Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_60)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_61Unmarshaller struct{}

func (u *A_extension_61Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_61), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_61Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_61)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_62Unmarshaller struct{}

func (u *A_extension_62Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_62), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_62Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_62)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_63Unmarshaller struct{}

func (u *A_extension_63Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_63), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_63Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_63)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_64Unmarshaller struct{}

func (u *A_extension_64Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_64), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_64Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_64)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_65Unmarshaller struct{}

func (u *A_extension_65Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_65), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_65Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_65)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_66Unmarshaller struct{}

func (u *A_extension_66Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_66), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_66Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_66)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_67Unmarshaller struct{}

func (u *A_extension_67Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_67), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_67Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_67)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_68Unmarshaller struct{}

func (u *A_extension_68Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_68), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_68Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_68)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_69Unmarshaller struct{}

func (u *A_extension_69Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_69), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_69Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_69)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_7Unmarshaller struct{}

func (u *A_extension_7Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_7), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_7Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_7)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_70Unmarshaller struct{}

func (u *A_extension_70Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_70), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_70Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_70)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_71Unmarshaller struct{}

func (u *A_extension_71Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_71), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_71Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_71)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_72Unmarshaller struct{}

func (u *A_extension_72Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_72), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_72Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_72)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_73Unmarshaller struct{}

func (u *A_extension_73Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_73), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_73Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_73)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_74Unmarshaller struct{}

func (u *A_extension_74Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_74), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_74Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_74)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_75Unmarshaller struct{}

func (u *A_extension_75Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_75), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_75Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_75)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_76Unmarshaller struct{}

func (u *A_extension_76Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_76), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_76Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_76)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_77Unmarshaller struct{}

func (u *A_extension_77Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_77), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_77Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_77)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_78Unmarshaller struct{}

func (u *A_extension_78Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_78), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_78Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_78)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_79Unmarshaller struct{}

func (u *A_extension_79Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_79), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_79Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_79)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_8Unmarshaller struct{}

func (u *A_extension_8Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_8), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_8Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_8)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_80Unmarshaller struct{}

func (u *A_extension_80Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_80), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_80Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_80)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_81Unmarshaller struct{}

func (u *A_extension_81Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_81), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_81Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_81)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_82Unmarshaller struct{}

func (u *A_extension_82Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_82), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_82Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_82)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_83Unmarshaller struct{}

func (u *A_extension_83Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_83), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_83Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_83)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_84Unmarshaller struct{}

func (u *A_extension_84Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_84), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_84Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_84)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_85Unmarshaller struct{}

func (u *A_extension_85Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_85), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_85Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_85)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_86Unmarshaller struct{}

func (u *A_extension_86Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_86), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_86Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_86)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_87Unmarshaller struct{}

func (u *A_extension_87Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_87), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_87Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_87)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_88Unmarshaller struct{}

func (u *A_extension_88Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_88), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_88Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_88)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_89Unmarshaller struct{}

func (u *A_extension_89Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_89), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_89Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_89)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type A_extension_9Unmarshaller struct{}

func (u *A_extension_9Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(A_extension_9), stage, identifier, instanceName, preserveOrder)
}

func (u *A_extension_9Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*A_extension_9)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	}
	return nil
}

type AbstractAirportHeliportProtectionAreaTypeUnmarshaller struct{}

func (u *AbstractAirportHeliportProtectionAreaTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AbstractAirportHeliportProtectionAreaType), stage, identifier, instanceName, preserveOrder)
}

func (u *AbstractAirportHeliportProtectionAreaTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AbstractAirportHeliportProtectionAreaType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	}
	return nil
}

type AbstractGroundLightSystemTypeUnmarshaller struct{}

func (u *AbstractGroundLightSystemTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AbstractGroundLightSystemType), stage, identifier, instanceName, preserveOrder)
}

func (u *AbstractGroundLightSystemTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AbstractGroundLightSystemType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	}
	return nil
}

type AbstractMarkingTypeUnmarshaller struct{}

func (u *AbstractMarkingTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AbstractMarkingType), stage, identifier, instanceName, preserveOrder)
}

func (u *AbstractMarkingTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AbstractMarkingType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	}
	return nil
}

type AbstractPropertiesWithScheduleTypeUnmarshaller struct{}

func (u *AbstractPropertiesWithScheduleTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AbstractPropertiesWithScheduleType), stage, identifier, instanceName, preserveOrder)
}

func (u *AbstractPropertiesWithScheduleTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AbstractPropertiesWithScheduleType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	}
	return nil
}

type AbstractSurfaceContaminationTypeUnmarshaller struct{}

func (u *AbstractSurfaceContaminationTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AbstractSurfaceContaminationType), stage, identifier, instanceName, preserveOrder)
}

func (u *AbstractSurfaceContaminationTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AbstractSurfaceContaminationType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	}
	return nil
}

type AbstractUsageConditionTypeUnmarshaller struct{}

func (u *AbstractUsageConditionTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AbstractUsageConditionType), stage, identifier, instanceName, preserveOrder)
}

func (u *AbstractUsageConditionTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AbstractUsageConditionType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	}
	return nil
}

type AircraftCharacteristicPropertyTypeUnmarshaller struct{}

func (u *AircraftCharacteristicPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AircraftCharacteristicPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *AircraftCharacteristicPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AircraftCharacteristicPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "AircraftCharacteristic":
		GongUnmarshallPointer(&instance.AircraftCharacteristic, valueExpr, identifierMap)
	}
	return nil
}

type AircraftCharacteristicTypeUnmarshaller struct{}

func (u *AircraftCharacteristicTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AircraftCharacteristicType), stage, identifier, instanceName, preserveOrder)
}

func (u *AircraftCharacteristicTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AircraftCharacteristicType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type AircraftStandContaminationPropertyTypeUnmarshaller struct{}

func (u *AircraftStandContaminationPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AircraftStandContaminationPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *AircraftStandContaminationPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AircraftStandContaminationPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "AircraftStandContamination":
		GongUnmarshallPointer(&instance.AircraftStandContamination, valueExpr, identifierMap)
	}
	return nil
}

type AircraftStandContaminationTypeUnmarshaller struct{}

func (u *AircraftStandContaminationTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AircraftStandContaminationType), stage, identifier, instanceName, preserveOrder)
}

func (u *AircraftStandContaminationTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AircraftStandContaminationType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "CriticalRidge":
		GongUnmarshallSliceOfPointers(&instance.CriticalRidge, valueExpr, identifierMap)
	case "Layer":
		GongUnmarshallSliceOfPointers(&instance.Layer, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type AircraftStandPropertyTypeUnmarshaller struct{}

func (u *AircraftStandPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AircraftStandPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *AircraftStandPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AircraftStandPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "Type_XSD":
		instance.Type_XSD = GongExtractString(valueExpr)
	case "Href":
		instance.Href = GongExtractString(valueExpr)
	case "Role":
		instance.Role = GongExtractString(valueExpr)
	case "Arcrole":
		instance.Arcrole = GongExtractString(valueExpr)
	case "Title":
		instance.Title = GongExtractString(valueExpr)
	case "Show":
		instance.Show = GongExtractString(valueExpr)
	case "Actuate":
		instance.Actuate = GongExtractString(valueExpr)
	}
	return nil
}

type AircraftStandTimeSlicePropertyTypeUnmarshaller struct{}

func (u *AircraftStandTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AircraftStandTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *AircraftStandTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AircraftStandTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "AircraftStandTimeSlice":
		GongUnmarshallPointer(&instance.AircraftStandTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type AircraftStandTimeSliceTypeUnmarshaller struct{}

func (u *AircraftStandTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AircraftStandTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *AircraftStandTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AircraftStandTimeSliceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Interpretation":
		instance.Interpretation = GongExtractString(valueExpr)
	case "SequenceNumber":
		instance.SequenceNumber = GongExtractInt(valueExpr)
	case "CorrectionNumber":
		instance.CorrectionNumber = GongExtractInt(valueExpr)
	case "SurfaceProperties":
		GongUnmarshallPointer(&instance.SurfaceProperties, valueExpr, identifierMap)
	case "Location":
		GongUnmarshallPointer(&instance.Location, valueExpr, identifierMap)
	case "ApronLocation":
		GongUnmarshallPointer(&instance.ApronLocation, valueExpr, identifierMap)
	case "Extent":
		GongUnmarshallPointer(&instance.Extent, valueExpr, identifierMap)
	case "Contaminant":
		GongUnmarshallSliceOfPointers(&instance.Contaminant, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Availability":
		GongUnmarshallSliceOfPointers(&instance.Availability, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type AircraftStandTypeUnmarshaller struct{}

func (u *AircraftStandTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AircraftStandType), stage, identifier, instanceName, preserveOrder)
}

func (u *AircraftStandTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AircraftStandType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeSlice":
		GongUnmarshallSliceOfPointers(&instance.TimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type AirportHeliportAvailabilityPropertyTypeUnmarshaller struct{}

func (u *AirportHeliportAvailabilityPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AirportHeliportAvailabilityPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *AirportHeliportAvailabilityPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AirportHeliportAvailabilityPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "AirportHeliportAvailability":
		GongUnmarshallPointer(&instance.AirportHeliportAvailability, valueExpr, identifierMap)
	}
	return nil
}

type AirportHeliportAvailabilityTypeUnmarshaller struct{}

func (u *AirportHeliportAvailabilityTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AirportHeliportAvailabilityType), stage, identifier, instanceName, preserveOrder)
}

func (u *AirportHeliportAvailabilityTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AirportHeliportAvailabilityType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeInterval":
		GongUnmarshallSliceOfPointers(&instance.TimeInterval, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "SpecialDateAuthority":
		GongUnmarshallSliceOfPointers(&instance.SpecialDateAuthority, valueExpr, identifierMap)
	case "Usage":
		GongUnmarshallSliceOfPointers(&instance.Usage, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type AirportHeliportCollocationTimeSlicePropertyTypeUnmarshaller struct{}

func (u *AirportHeliportCollocationTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AirportHeliportCollocationTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *AirportHeliportCollocationTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AirportHeliportCollocationTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "AirportHeliportCollocationTimeSlice":
		GongUnmarshallPointer(&instance.AirportHeliportCollocationTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type AirportHeliportCollocationTimeSliceTypeUnmarshaller struct{}

func (u *AirportHeliportCollocationTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AirportHeliportCollocationTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *AirportHeliportCollocationTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AirportHeliportCollocationTimeSliceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Interpretation":
		instance.Interpretation = GongExtractString(valueExpr)
	case "SequenceNumber":
		instance.SequenceNumber = GongExtractInt(valueExpr)
	case "CorrectionNumber":
		instance.CorrectionNumber = GongExtractInt(valueExpr)
	case "HostAirport":
		GongUnmarshallPointer(&instance.HostAirport, valueExpr, identifierMap)
	case "DependentAirport":
		GongUnmarshallPointer(&instance.DependentAirport, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type AirportHeliportCollocationTypeUnmarshaller struct{}

func (u *AirportHeliportCollocationTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AirportHeliportCollocationType), stage, identifier, instanceName, preserveOrder)
}

func (u *AirportHeliportCollocationTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AirportHeliportCollocationType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeSlice":
		GongUnmarshallSliceOfPointers(&instance.TimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type AirportHeliportContaminationPropertyTypeUnmarshaller struct{}

func (u *AirportHeliportContaminationPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AirportHeliportContaminationPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *AirportHeliportContaminationPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AirportHeliportContaminationPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "AirportHeliportContamination":
		GongUnmarshallPointer(&instance.AirportHeliportContamination, valueExpr, identifierMap)
	}
	return nil
}

type AirportHeliportContaminationTypeUnmarshaller struct{}

func (u *AirportHeliportContaminationTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AirportHeliportContaminationType), stage, identifier, instanceName, preserveOrder)
}

func (u *AirportHeliportContaminationTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AirportHeliportContaminationType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "CriticalRidge":
		GongUnmarshallSliceOfPointers(&instance.CriticalRidge, valueExpr, identifierMap)
	case "Layer":
		GongUnmarshallSliceOfPointers(&instance.Layer, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type AirportHeliportPropertyTypeUnmarshaller struct{}

func (u *AirportHeliportPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AirportHeliportPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *AirportHeliportPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AirportHeliportPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "Type_XSD":
		instance.Type_XSD = GongExtractString(valueExpr)
	case "Href":
		instance.Href = GongExtractString(valueExpr)
	case "Role":
		instance.Role = GongExtractString(valueExpr)
	case "Arcrole":
		instance.Arcrole = GongExtractString(valueExpr)
	case "Title":
		instance.Title = GongExtractString(valueExpr)
	case "Show":
		instance.Show = GongExtractString(valueExpr)
	case "Actuate":
		instance.Actuate = GongExtractString(valueExpr)
	}
	return nil
}

type AirportHeliportProtectionAreaPropertyTypeUnmarshaller struct{}

func (u *AirportHeliportProtectionAreaPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AirportHeliportProtectionAreaPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *AirportHeliportProtectionAreaPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AirportHeliportProtectionAreaPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "Type_XSD":
		instance.Type_XSD = GongExtractString(valueExpr)
	case "Href":
		instance.Href = GongExtractString(valueExpr)
	case "Role":
		instance.Role = GongExtractString(valueExpr)
	case "Arcrole":
		instance.Arcrole = GongExtractString(valueExpr)
	case "Title":
		instance.Title = GongExtractString(valueExpr)
	case "Show":
		instance.Show = GongExtractString(valueExpr)
	case "Actuate":
		instance.Actuate = GongExtractString(valueExpr)
	}
	return nil
}

type AirportHeliportResponsibilityOrganisationPropertyTypeUnmarshaller struct{}

func (u *AirportHeliportResponsibilityOrganisationPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AirportHeliportResponsibilityOrganisationPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *AirportHeliportResponsibilityOrganisationPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AirportHeliportResponsibilityOrganisationPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "AirportHeliportResponsibilityOrganisation":
		GongUnmarshallPointer(&instance.AirportHeliportResponsibilityOrganisation, valueExpr, identifierMap)
	}
	return nil
}

type AirportHeliportResponsibilityOrganisationTypeUnmarshaller struct{}

func (u *AirportHeliportResponsibilityOrganisationTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AirportHeliportResponsibilityOrganisationType), stage, identifier, instanceName, preserveOrder)
}

func (u *AirportHeliportResponsibilityOrganisationTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AirportHeliportResponsibilityOrganisationType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeInterval":
		GongUnmarshallSliceOfPointers(&instance.TimeInterval, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "SpecialDateAuthority":
		GongUnmarshallSliceOfPointers(&instance.SpecialDateAuthority, valueExpr, identifierMap)
	case "TheOrganisationAuthority":
		GongUnmarshallPointer(&instance.TheOrganisationAuthority, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type AirportHeliportTimeSlicePropertyTypeUnmarshaller struct{}

func (u *AirportHeliportTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AirportHeliportTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *AirportHeliportTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AirportHeliportTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "AirportHeliportTimeSlice":
		GongUnmarshallPointer(&instance.AirportHeliportTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type AirportHeliportTimeSliceTypeUnmarshaller struct{}

func (u *AirportHeliportTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AirportHeliportTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *AirportHeliportTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AirportHeliportTimeSliceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Interpretation":
		instance.Interpretation = GongExtractString(valueExpr)
	case "SequenceNumber":
		instance.SequenceNumber = GongExtractInt(valueExpr)
	case "CorrectionNumber":
		instance.CorrectionNumber = GongExtractInt(valueExpr)
	case "Contaminant":
		GongUnmarshallSliceOfPointers(&instance.Contaminant, valueExpr, identifierMap)
	case "ServedCity":
		GongUnmarshallSliceOfPointers(&instance.ServedCity, valueExpr, identifierMap)
	case "ResponsibleOrganisation":
		GongUnmarshallPointer(&instance.ResponsibleOrganisation, valueExpr, identifierMap)
	case "ARP":
		GongUnmarshallPointer(&instance.ARP, valueExpr, identifierMap)
	case "AviationBoundary":
		GongUnmarshallPointer(&instance.AviationBoundary, valueExpr, identifierMap)
	case "AltimeterSource":
		GongUnmarshallSliceOfPointers(&instance.AltimeterSource, valueExpr, identifierMap)
	case "Contact":
		GongUnmarshallSliceOfPointers(&instance.Contact, valueExpr, identifierMap)
	case "Availability":
		GongUnmarshallSliceOfPointers(&instance.Availability, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type AirportHeliportTypeUnmarshaller struct{}

func (u *AirportHeliportTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AirportHeliportType), stage, identifier, instanceName, preserveOrder)
}

func (u *AirportHeliportTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AirportHeliportType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeSlice":
		GongUnmarshallSliceOfPointers(&instance.TimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type AirportHeliportUsagePropertyTypeUnmarshaller struct{}

func (u *AirportHeliportUsagePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AirportHeliportUsagePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *AirportHeliportUsagePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AirportHeliportUsagePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "AirportHeliportUsage":
		GongUnmarshallPointer(&instance.AirportHeliportUsage, valueExpr, identifierMap)
	}
	return nil
}

type AirportHeliportUsageTypeUnmarshaller struct{}

func (u *AirportHeliportUsageTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AirportHeliportUsageType), stage, identifier, instanceName, preserveOrder)
}

func (u *AirportHeliportUsageTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AirportHeliportUsageType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Contact":
		GongUnmarshallSliceOfPointers(&instance.Contact, valueExpr, identifierMap)
	case "Selection":
		GongUnmarshallPointer(&instance.Selection, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type AirportHotSpotTimeSlicePropertyTypeUnmarshaller struct{}

func (u *AirportHotSpotTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AirportHotSpotTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *AirportHotSpotTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AirportHotSpotTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "AirportHotSpotTimeSlice":
		GongUnmarshallPointer(&instance.AirportHotSpotTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type AirportHotSpotTimeSliceTypeUnmarshaller struct{}

func (u *AirportHotSpotTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AirportHotSpotTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *AirportHotSpotTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AirportHotSpotTimeSliceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Interpretation":
		instance.Interpretation = GongExtractString(valueExpr)
	case "SequenceNumber":
		instance.SequenceNumber = GongExtractInt(valueExpr)
	case "CorrectionNumber":
		instance.CorrectionNumber = GongExtractInt(valueExpr)
	case "Area":
		GongUnmarshallPointer(&instance.Area, valueExpr, identifierMap)
	case "AffectedAirport":
		GongUnmarshallPointer(&instance.AffectedAirport, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type AirportHotSpotTypeUnmarshaller struct{}

func (u *AirportHotSpotTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AirportHotSpotType), stage, identifier, instanceName, preserveOrder)
}

func (u *AirportHotSpotTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AirportHotSpotType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeSlice":
		GongUnmarshallSliceOfPointers(&instance.TimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type AirportProtectionAreaMarkingTimeSlicePropertyTypeUnmarshaller struct{}

func (u *AirportProtectionAreaMarkingTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AirportProtectionAreaMarkingTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *AirportProtectionAreaMarkingTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AirportProtectionAreaMarkingTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "AirportProtectionAreaMarkingTimeSlice":
		GongUnmarshallPointer(&instance.AirportProtectionAreaMarkingTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type AirportProtectionAreaMarkingTimeSliceTypeUnmarshaller struct{}

func (u *AirportProtectionAreaMarkingTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AirportProtectionAreaMarkingTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *AirportProtectionAreaMarkingTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AirportProtectionAreaMarkingTimeSliceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Interpretation":
		instance.Interpretation = GongExtractString(valueExpr)
	case "SequenceNumber":
		instance.SequenceNumber = GongExtractInt(valueExpr)
	case "CorrectionNumber":
		instance.CorrectionNumber = GongExtractInt(valueExpr)
	case "Element":
		GongUnmarshallSliceOfPointers(&instance.Element, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "MarkedProtectionArea":
		GongUnmarshallPointer(&instance.MarkedProtectionArea, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type AirportProtectionAreaMarkingTypeUnmarshaller struct{}

func (u *AirportProtectionAreaMarkingTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AirportProtectionAreaMarkingType), stage, identifier, instanceName, preserveOrder)
}

func (u *AirportProtectionAreaMarkingTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AirportProtectionAreaMarkingType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeSlice":
		GongUnmarshallSliceOfPointers(&instance.TimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type AirportSuppliesServicePropertyTypeUnmarshaller struct{}

func (u *AirportSuppliesServicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AirportSuppliesServicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *AirportSuppliesServicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AirportSuppliesServicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "Type_XSD":
		instance.Type_XSD = GongExtractString(valueExpr)
	case "Href":
		instance.Href = GongExtractString(valueExpr)
	case "Role":
		instance.Role = GongExtractString(valueExpr)
	case "Arcrole":
		instance.Arcrole = GongExtractString(valueExpr)
	case "Title":
		instance.Title = GongExtractString(valueExpr)
	case "Show":
		instance.Show = GongExtractString(valueExpr)
	case "Actuate":
		instance.Actuate = GongExtractString(valueExpr)
	}
	return nil
}

type AltimeterSourcePropertyTypeUnmarshaller struct{}

func (u *AltimeterSourcePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AltimeterSourcePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *AltimeterSourcePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AltimeterSourcePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "Type_XSD":
		instance.Type_XSD = GongExtractString(valueExpr)
	case "Href":
		instance.Href = GongExtractString(valueExpr)
	case "Role":
		instance.Role = GongExtractString(valueExpr)
	case "Arcrole":
		instance.Arcrole = GongExtractString(valueExpr)
	case "Title":
		instance.Title = GongExtractString(valueExpr)
	case "Show":
		instance.Show = GongExtractString(valueExpr)
	case "Actuate":
		instance.Actuate = GongExtractString(valueExpr)
	}
	return nil
}

type AltimeterSourceStatusPropertyTypeUnmarshaller struct{}

func (u *AltimeterSourceStatusPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AltimeterSourceStatusPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *AltimeterSourceStatusPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AltimeterSourceStatusPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "AltimeterSourceStatus":
		GongUnmarshallPointer(&instance.AltimeterSourceStatus, valueExpr, identifierMap)
	}
	return nil
}

type AltimeterSourceStatusTypeUnmarshaller struct{}

func (u *AltimeterSourceStatusTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AltimeterSourceStatusType), stage, identifier, instanceName, preserveOrder)
}

func (u *AltimeterSourceStatusTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AltimeterSourceStatusType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeInterval":
		GongUnmarshallSliceOfPointers(&instance.TimeInterval, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "SpecialDateAuthority":
		GongUnmarshallSliceOfPointers(&instance.SpecialDateAuthority, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type AltimeterSourceTimeSlicePropertyTypeUnmarshaller struct{}

func (u *AltimeterSourceTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AltimeterSourceTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *AltimeterSourceTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AltimeterSourceTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "AltimeterSourceTimeSlice":
		GongUnmarshallPointer(&instance.AltimeterSourceTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type AltimeterSourceTimeSliceTypeUnmarshaller struct{}

func (u *AltimeterSourceTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AltimeterSourceTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *AltimeterSourceTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AltimeterSourceTimeSliceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Interpretation":
		instance.Interpretation = GongExtractString(valueExpr)
	case "SequenceNumber":
		instance.SequenceNumber = GongExtractInt(valueExpr)
	case "CorrectionNumber":
		instance.CorrectionNumber = GongExtractInt(valueExpr)
	case "Availability":
		GongUnmarshallSliceOfPointers(&instance.Availability, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type AltimeterSourceTypeUnmarshaller struct{}

func (u *AltimeterSourceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AltimeterSourceType), stage, identifier, instanceName, preserveOrder)
}

func (u *AltimeterSourceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AltimeterSourceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeSlice":
		GongUnmarshallSliceOfPointers(&instance.TimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type ApproachLightingSystemTimeSlicePropertyTypeUnmarshaller struct{}

func (u *ApproachLightingSystemTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ApproachLightingSystemTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *ApproachLightingSystemTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ApproachLightingSystemTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "ApproachLightingSystemTimeSlice":
		GongUnmarshallPointer(&instance.ApproachLightingSystemTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type ApproachLightingSystemTimeSliceTypeUnmarshaller struct{}

func (u *ApproachLightingSystemTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ApproachLightingSystemTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *ApproachLightingSystemTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ApproachLightingSystemTimeSliceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Interpretation":
		instance.Interpretation = GongExtractString(valueExpr)
	case "SequenceNumber":
		instance.SequenceNumber = GongExtractInt(valueExpr)
	case "CorrectionNumber":
		instance.CorrectionNumber = GongExtractInt(valueExpr)
	case "Element":
		GongUnmarshallSliceOfPointers(&instance.Element, valueExpr, identifierMap)
	case "Availability":
		GongUnmarshallSliceOfPointers(&instance.Availability, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "ServedRunwayDirection":
		GongUnmarshallPointer(&instance.ServedRunwayDirection, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type ApproachLightingSystemTypeUnmarshaller struct{}

func (u *ApproachLightingSystemTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ApproachLightingSystemType), stage, identifier, instanceName, preserveOrder)
}

func (u *ApproachLightingSystemTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ApproachLightingSystemType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeSlice":
		GongUnmarshallSliceOfPointers(&instance.TimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type ApronAreaAvailabilityPropertyTypeUnmarshaller struct{}

func (u *ApronAreaAvailabilityPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ApronAreaAvailabilityPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *ApronAreaAvailabilityPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ApronAreaAvailabilityPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "ApronAreaAvailability":
		GongUnmarshallPointer(&instance.ApronAreaAvailability, valueExpr, identifierMap)
	}
	return nil
}

type ApronAreaAvailabilityTypeUnmarshaller struct{}

func (u *ApronAreaAvailabilityTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ApronAreaAvailabilityType), stage, identifier, instanceName, preserveOrder)
}

func (u *ApronAreaAvailabilityTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ApronAreaAvailabilityType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeInterval":
		GongUnmarshallSliceOfPointers(&instance.TimeInterval, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "SpecialDateAuthority":
		GongUnmarshallSliceOfPointers(&instance.SpecialDateAuthority, valueExpr, identifierMap)
	case "Usage":
		GongUnmarshallSliceOfPointers(&instance.Usage, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type ApronAreaUsagePropertyTypeUnmarshaller struct{}

func (u *ApronAreaUsagePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ApronAreaUsagePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *ApronAreaUsagePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ApronAreaUsagePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "ApronAreaUsage":
		GongUnmarshallPointer(&instance.ApronAreaUsage, valueExpr, identifierMap)
	}
	return nil
}

type ApronAreaUsageTypeUnmarshaller struct{}

func (u *ApronAreaUsageTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ApronAreaUsageType), stage, identifier, instanceName, preserveOrder)
}

func (u *ApronAreaUsageTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ApronAreaUsageType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Contact":
		GongUnmarshallSliceOfPointers(&instance.Contact, valueExpr, identifierMap)
	case "Selection":
		GongUnmarshallPointer(&instance.Selection, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type ApronContaminationPropertyTypeUnmarshaller struct{}

func (u *ApronContaminationPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ApronContaminationPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *ApronContaminationPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ApronContaminationPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "ApronContamination":
		GongUnmarshallPointer(&instance.ApronContamination, valueExpr, identifierMap)
	}
	return nil
}

type ApronContaminationTypeUnmarshaller struct{}

func (u *ApronContaminationTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ApronContaminationType), stage, identifier, instanceName, preserveOrder)
}

func (u *ApronContaminationTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ApronContaminationType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "CriticalRidge":
		GongUnmarshallSliceOfPointers(&instance.CriticalRidge, valueExpr, identifierMap)
	case "Layer":
		GongUnmarshallSliceOfPointers(&instance.Layer, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type ApronElementPropertyTypeUnmarshaller struct{}

func (u *ApronElementPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ApronElementPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *ApronElementPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ApronElementPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "Type_XSD":
		instance.Type_XSD = GongExtractString(valueExpr)
	case "Href":
		instance.Href = GongExtractString(valueExpr)
	case "Role":
		instance.Role = GongExtractString(valueExpr)
	case "Arcrole":
		instance.Arcrole = GongExtractString(valueExpr)
	case "Title":
		instance.Title = GongExtractString(valueExpr)
	case "Show":
		instance.Show = GongExtractString(valueExpr)
	case "Actuate":
		instance.Actuate = GongExtractString(valueExpr)
	}
	return nil
}

type ApronElementTimeSlicePropertyTypeUnmarshaller struct{}

func (u *ApronElementTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ApronElementTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *ApronElementTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ApronElementTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "ApronElementTimeSlice":
		GongUnmarshallPointer(&instance.ApronElementTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type ApronElementTimeSliceTypeUnmarshaller struct{}

func (u *ApronElementTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ApronElementTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *ApronElementTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ApronElementTimeSliceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Interpretation":
		instance.Interpretation = GongExtractString(valueExpr)
	case "SequenceNumber":
		instance.SequenceNumber = GongExtractInt(valueExpr)
	case "CorrectionNumber":
		instance.CorrectionNumber = GongExtractInt(valueExpr)
	case "AssociatedApron":
		GongUnmarshallPointer(&instance.AssociatedApron, valueExpr, identifierMap)
	case "SurfaceProperties":
		GongUnmarshallPointer(&instance.SurfaceProperties, valueExpr, identifierMap)
	case "Extent":
		GongUnmarshallPointer(&instance.Extent, valueExpr, identifierMap)
	case "SupplyService":
		GongUnmarshallSliceOfPointers(&instance.SupplyService, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Availability":
		GongUnmarshallSliceOfPointers(&instance.Availability, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type ApronElementTypeUnmarshaller struct{}

func (u *ApronElementTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ApronElementType), stage, identifier, instanceName, preserveOrder)
}

func (u *ApronElementTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ApronElementType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeSlice":
		GongUnmarshallSliceOfPointers(&instance.TimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type ApronLightSystemTimeSlicePropertyTypeUnmarshaller struct{}

func (u *ApronLightSystemTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ApronLightSystemTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *ApronLightSystemTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ApronLightSystemTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "ApronLightSystemTimeSlice":
		GongUnmarshallPointer(&instance.ApronLightSystemTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type ApronLightSystemTimeSliceTypeUnmarshaller struct{}

func (u *ApronLightSystemTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ApronLightSystemTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *ApronLightSystemTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ApronLightSystemTimeSliceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Interpretation":
		instance.Interpretation = GongExtractString(valueExpr)
	case "SequenceNumber":
		instance.SequenceNumber = GongExtractInt(valueExpr)
	case "CorrectionNumber":
		instance.CorrectionNumber = GongExtractInt(valueExpr)
	case "Element":
		GongUnmarshallSliceOfPointers(&instance.Element, valueExpr, identifierMap)
	case "Availability":
		GongUnmarshallSliceOfPointers(&instance.Availability, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "LightedApron":
		GongUnmarshallPointer(&instance.LightedApron, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type ApronLightSystemTypeUnmarshaller struct{}

func (u *ApronLightSystemTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ApronLightSystemType), stage, identifier, instanceName, preserveOrder)
}

func (u *ApronLightSystemTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ApronLightSystemType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeSlice":
		GongUnmarshallSliceOfPointers(&instance.TimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type ApronMarkingTimeSlicePropertyTypeUnmarshaller struct{}

func (u *ApronMarkingTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ApronMarkingTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *ApronMarkingTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ApronMarkingTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "ApronMarkingTimeSlice":
		GongUnmarshallPointer(&instance.ApronMarkingTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type ApronMarkingTimeSliceTypeUnmarshaller struct{}

func (u *ApronMarkingTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ApronMarkingTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *ApronMarkingTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ApronMarkingTimeSliceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Interpretation":
		instance.Interpretation = GongExtractString(valueExpr)
	case "SequenceNumber":
		instance.SequenceNumber = GongExtractInt(valueExpr)
	case "CorrectionNumber":
		instance.CorrectionNumber = GongExtractInt(valueExpr)
	case "Element":
		GongUnmarshallSliceOfPointers(&instance.Element, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "MarkedApron":
		GongUnmarshallPointer(&instance.MarkedApron, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type ApronMarkingTypeUnmarshaller struct{}

func (u *ApronMarkingTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ApronMarkingType), stage, identifier, instanceName, preserveOrder)
}

func (u *ApronMarkingTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ApronMarkingType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeSlice":
		GongUnmarshallSliceOfPointers(&instance.TimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type ApronPropertyTypeUnmarshaller struct{}

func (u *ApronPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ApronPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *ApronPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ApronPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "Type_XSD":
		instance.Type_XSD = GongExtractString(valueExpr)
	case "Href":
		instance.Href = GongExtractString(valueExpr)
	case "Role":
		instance.Role = GongExtractString(valueExpr)
	case "Arcrole":
		instance.Arcrole = GongExtractString(valueExpr)
	case "Title":
		instance.Title = GongExtractString(valueExpr)
	case "Show":
		instance.Show = GongExtractString(valueExpr)
	case "Actuate":
		instance.Actuate = GongExtractString(valueExpr)
	}
	return nil
}

type ApronTimeSlicePropertyTypeUnmarshaller struct{}

func (u *ApronTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ApronTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *ApronTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ApronTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "ApronTimeSlice":
		GongUnmarshallPointer(&instance.ApronTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type ApronTimeSliceTypeUnmarshaller struct{}

func (u *ApronTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ApronTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *ApronTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ApronTimeSliceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Interpretation":
		instance.Interpretation = GongExtractString(valueExpr)
	case "SequenceNumber":
		instance.SequenceNumber = GongExtractInt(valueExpr)
	case "CorrectionNumber":
		instance.CorrectionNumber = GongExtractInt(valueExpr)
	case "SurfaceProperties":
		GongUnmarshallPointer(&instance.SurfaceProperties, valueExpr, identifierMap)
	case "AssociatedAirportHeliport":
		GongUnmarshallPointer(&instance.AssociatedAirportHeliport, valueExpr, identifierMap)
	case "Contaminant":
		GongUnmarshallSliceOfPointers(&instance.Contaminant, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Availability":
		GongUnmarshallSliceOfPointers(&instance.Availability, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type ApronTypeUnmarshaller struct{}

func (u *ApronTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ApronType), stage, identifier, instanceName, preserveOrder)
}

func (u *ApronTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ApronType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeSlice":
		GongUnmarshallSliceOfPointers(&instance.TimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type ArrestingGearTimeSlicePropertyTypeUnmarshaller struct{}

func (u *ArrestingGearTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ArrestingGearTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *ArrestingGearTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ArrestingGearTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "ArrestingGearTimeSlice":
		GongUnmarshallPointer(&instance.ArrestingGearTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type ArrestingGearTimeSliceTypeUnmarshaller struct{}

func (u *ArrestingGearTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ArrestingGearTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *ArrestingGearTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ArrestingGearTimeSliceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Interpretation":
		instance.Interpretation = GongExtractString(valueExpr)
	case "SequenceNumber":
		instance.SequenceNumber = GongExtractInt(valueExpr)
	case "CorrectionNumber":
		instance.CorrectionNumber = GongExtractInt(valueExpr)
	case "RunwayDirection":
		GongUnmarshallSliceOfPointers(&instance.RunwayDirection, valueExpr, identifierMap)
	case "SurfaceProperties":
		GongUnmarshallPointer(&instance.SurfaceProperties, valueExpr, identifierMap)
	case "Extent_surfaceExtent":
		GongUnmarshallPointer(&instance.Extent_surfaceExtent, valueExpr, identifierMap)
	case "Extent_curveExtent":
		GongUnmarshallPointer(&instance.Extent_curveExtent, valueExpr, identifierMap)
	case "Extent_pointExtent":
		GongUnmarshallPointer(&instance.Extent_pointExtent, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type ArrestingGearTypeUnmarshaller struct{}

func (u *ArrestingGearTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ArrestingGearType), stage, identifier, instanceName, preserveOrder)
}

func (u *ArrestingGearTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ArrestingGearType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeSlice":
		GongUnmarshallSliceOfPointers(&instance.TimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type CityPropertyTypeUnmarshaller struct{}

func (u *CityPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CityPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *CityPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CityPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "City":
		GongUnmarshallPointer(&instance.City, valueExpr, identifierMap)
	}
	return nil
}

type CityTypeUnmarshaller struct{}

func (u *CityTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CityType), stage, identifier, instanceName, preserveOrder)
}

func (u *CityTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CityType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type ConditionCombinationPropertyTypeUnmarshaller struct{}

func (u *ConditionCombinationPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ConditionCombinationPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *ConditionCombinationPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ConditionCombinationPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "ConditionCombination":
		GongUnmarshallPointer(&instance.ConditionCombination, valueExpr, identifierMap)
	}
	return nil
}

type ConditionCombinationTypeUnmarshaller struct{}

func (u *ConditionCombinationTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ConditionCombinationType), stage, identifier, instanceName, preserveOrder)
}

func (u *ConditionCombinationTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ConditionCombinationType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeInterval":
		GongUnmarshallSliceOfPointers(&instance.TimeInterval, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "SpecialDateAuthority":
		GongUnmarshallSliceOfPointers(&instance.SpecialDateAuthority, valueExpr, identifierMap)
	case "Weather":
		GongUnmarshallSliceOfPointers(&instance.Weather, valueExpr, identifierMap)
	case "Aircraft":
		GongUnmarshallSliceOfPointers(&instance.Aircraft, valueExpr, identifierMap)
	case "Flight":
		GongUnmarshallSliceOfPointers(&instance.Flight, valueExpr, identifierMap)
	case "SubCondition":
		GongUnmarshallSliceOfPointers(&instance.SubCondition, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type ContactInformationPropertyTypeUnmarshaller struct{}

func (u *ContactInformationPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ContactInformationPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *ContactInformationPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ContactInformationPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "ContactInformation":
		GongUnmarshallPointer(&instance.ContactInformation, valueExpr, identifierMap)
	}
	return nil
}

type ContactInformationTypeUnmarshaller struct{}

func (u *ContactInformationTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ContactInformationType), stage, identifier, instanceName, preserveOrder)
}

func (u *ContactInformationTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ContactInformationType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Address":
		GongUnmarshallSliceOfPointers(&instance.Address, valueExpr, identifierMap)
	case "NetworkNode":
		GongUnmarshallSliceOfPointers(&instance.NetworkNode, valueExpr, identifierMap)
	case "PhoneFax":
		GongUnmarshallSliceOfPointers(&instance.PhoneFax, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type CurveTypeUnmarshaller struct{}

func (u *CurveTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CurveType), stage, identifier, instanceName, preserveOrder)
}

func (u *CurveTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CurveType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "SrsName":
		instance.SrsName = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	}
	return nil
}

type DeicingAreaMarkingTimeSlicePropertyTypeUnmarshaller struct{}

func (u *DeicingAreaMarkingTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DeicingAreaMarkingTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *DeicingAreaMarkingTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DeicingAreaMarkingTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "DeicingAreaMarkingTimeSlice":
		GongUnmarshallPointer(&instance.DeicingAreaMarkingTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type DeicingAreaMarkingTimeSliceTypeUnmarshaller struct{}

func (u *DeicingAreaMarkingTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DeicingAreaMarkingTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *DeicingAreaMarkingTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DeicingAreaMarkingTimeSliceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Interpretation":
		instance.Interpretation = GongExtractString(valueExpr)
	case "SequenceNumber":
		instance.SequenceNumber = GongExtractInt(valueExpr)
	case "CorrectionNumber":
		instance.CorrectionNumber = GongExtractInt(valueExpr)
	case "Element":
		GongUnmarshallSliceOfPointers(&instance.Element, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "MarkedDeicingArea":
		GongUnmarshallPointer(&instance.MarkedDeicingArea, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type DeicingAreaMarkingTypeUnmarshaller struct{}

func (u *DeicingAreaMarkingTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DeicingAreaMarkingType), stage, identifier, instanceName, preserveOrder)
}

func (u *DeicingAreaMarkingTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DeicingAreaMarkingType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeSlice":
		GongUnmarshallSliceOfPointers(&instance.TimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type DeicingAreaPropertyTypeUnmarshaller struct{}

func (u *DeicingAreaPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DeicingAreaPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *DeicingAreaPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DeicingAreaPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "Type_XSD":
		instance.Type_XSD = GongExtractString(valueExpr)
	case "Href":
		instance.Href = GongExtractString(valueExpr)
	case "Role":
		instance.Role = GongExtractString(valueExpr)
	case "Arcrole":
		instance.Arcrole = GongExtractString(valueExpr)
	case "Title":
		instance.Title = GongExtractString(valueExpr)
	case "Show":
		instance.Show = GongExtractString(valueExpr)
	case "Actuate":
		instance.Actuate = GongExtractString(valueExpr)
	}
	return nil
}

type DeicingAreaTimeSlicePropertyTypeUnmarshaller struct{}

func (u *DeicingAreaTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DeicingAreaTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *DeicingAreaTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DeicingAreaTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "DeicingAreaTimeSlice":
		GongUnmarshallPointer(&instance.DeicingAreaTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type DeicingAreaTimeSliceTypeUnmarshaller struct{}

func (u *DeicingAreaTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DeicingAreaTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *DeicingAreaTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DeicingAreaTimeSliceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Interpretation":
		instance.Interpretation = GongExtractString(valueExpr)
	case "SequenceNumber":
		instance.SequenceNumber = GongExtractInt(valueExpr)
	case "CorrectionNumber":
		instance.CorrectionNumber = GongExtractInt(valueExpr)
	case "AssociatedApron":
		GongUnmarshallPointer(&instance.AssociatedApron, valueExpr, identifierMap)
	case "TaxiwayLocation":
		GongUnmarshallPointer(&instance.TaxiwayLocation, valueExpr, identifierMap)
	case "StandLocation":
		GongUnmarshallPointer(&instance.StandLocation, valueExpr, identifierMap)
	case "SurfaceProperties":
		GongUnmarshallPointer(&instance.SurfaceProperties, valueExpr, identifierMap)
	case "Extent":
		GongUnmarshallPointer(&instance.Extent, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Availability":
		GongUnmarshallSliceOfPointers(&instance.Availability, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type DeicingAreaTypeUnmarshaller struct{}

func (u *DeicingAreaTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DeicingAreaType), stage, identifier, instanceName, preserveOrder)
}

func (u *DeicingAreaTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DeicingAreaType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeSlice":
		GongUnmarshallSliceOfPointers(&instance.TimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type ElevatedCurvePropertyTypeUnmarshaller struct{}

func (u *ElevatedCurvePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ElevatedCurvePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *ElevatedCurvePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ElevatedCurvePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "ElevatedCurve":
		GongUnmarshallPointer(&instance.ElevatedCurve, valueExpr, identifierMap)
	}
	return nil
}

type ElevatedCurveTypeUnmarshaller struct{}

func (u *ElevatedCurveTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ElevatedCurveType), stage, identifier, instanceName, preserveOrder)
}

func (u *ElevatedCurveTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ElevatedCurveType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "SrsName":
		instance.SrsName = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type ElevatedPointPropertyTypeUnmarshaller struct{}

func (u *ElevatedPointPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ElevatedPointPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *ElevatedPointPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ElevatedPointPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "ElevatedPoint":
		GongUnmarshallPointer(&instance.ElevatedPoint, valueExpr, identifierMap)
	}
	return nil
}

type ElevatedPointTypeUnmarshaller struct{}

func (u *ElevatedPointTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ElevatedPointType), stage, identifier, instanceName, preserveOrder)
}

func (u *ElevatedPointTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ElevatedPointType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "SrsName":
		instance.SrsName = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type ElevatedSurfacePropertyTypeUnmarshaller struct{}

func (u *ElevatedSurfacePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ElevatedSurfacePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *ElevatedSurfacePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ElevatedSurfacePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "ElevatedSurface":
		GongUnmarshallPointer(&instance.ElevatedSurface, valueExpr, identifierMap)
	}
	return nil
}

type ElevatedSurfaceTypeUnmarshaller struct{}

func (u *ElevatedSurfaceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ElevatedSurfaceType), stage, identifier, instanceName, preserveOrder)
}

func (u *ElevatedSurfaceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ElevatedSurfaceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "SrsName":
		instance.SrsName = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type FlightCharacteristicPropertyTypeUnmarshaller struct{}

func (u *FlightCharacteristicPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FlightCharacteristicPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *FlightCharacteristicPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FlightCharacteristicPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "FlightCharacteristic":
		GongUnmarshallPointer(&instance.FlightCharacteristic, valueExpr, identifierMap)
	}
	return nil
}

type FlightCharacteristicTypeUnmarshaller struct{}

func (u *FlightCharacteristicTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FlightCharacteristicType), stage, identifier, instanceName, preserveOrder)
}

func (u *FlightCharacteristicTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FlightCharacteristicType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type FloatingDockSitePropertyTypeUnmarshaller struct{}

func (u *FloatingDockSitePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FloatingDockSitePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *FloatingDockSitePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FloatingDockSitePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "Type_XSD":
		instance.Type_XSD = GongExtractString(valueExpr)
	case "Href":
		instance.Href = GongExtractString(valueExpr)
	case "Role":
		instance.Role = GongExtractString(valueExpr)
	case "Arcrole":
		instance.Arcrole = GongExtractString(valueExpr)
	case "Title":
		instance.Title = GongExtractString(valueExpr)
	case "Show":
		instance.Show = GongExtractString(valueExpr)
	case "Actuate":
		instance.Actuate = GongExtractString(valueExpr)
	}
	return nil
}

type FloatingDockSiteTimeSlicePropertyTypeUnmarshaller struct{}

func (u *FloatingDockSiteTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FloatingDockSiteTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *FloatingDockSiteTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FloatingDockSiteTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "FloatingDockSiteTimeSlice":
		GongUnmarshallPointer(&instance.FloatingDockSiteTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type FloatingDockSiteTimeSliceTypeUnmarshaller struct{}

func (u *FloatingDockSiteTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FloatingDockSiteTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *FloatingDockSiteTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FloatingDockSiteTimeSliceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Interpretation":
		instance.Interpretation = GongExtractString(valueExpr)
	case "SequenceNumber":
		instance.SequenceNumber = GongExtractInt(valueExpr)
	case "CorrectionNumber":
		instance.CorrectionNumber = GongExtractInt(valueExpr)
	case "Extent":
		GongUnmarshallPointer(&instance.Extent, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type FloatingDockSiteTypeUnmarshaller struct{}

func (u *FloatingDockSiteTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(FloatingDockSiteType), stage, identifier, instanceName, preserveOrder)
}

func (u *FloatingDockSiteTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*FloatingDockSiteType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeSlice":
		GongUnmarshallSliceOfPointers(&instance.TimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type GroundLightingAvailabilityPropertyTypeUnmarshaller struct{}

func (u *GroundLightingAvailabilityPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(GroundLightingAvailabilityPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *GroundLightingAvailabilityPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GroundLightingAvailabilityPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "GroundLightingAvailability":
		GongUnmarshallPointer(&instance.GroundLightingAvailability, valueExpr, identifierMap)
	}
	return nil
}

type GroundLightingAvailabilityTypeUnmarshaller struct{}

func (u *GroundLightingAvailabilityTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(GroundLightingAvailabilityType), stage, identifier, instanceName, preserveOrder)
}

func (u *GroundLightingAvailabilityTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GroundLightingAvailabilityType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeInterval":
		GongUnmarshallSliceOfPointers(&instance.TimeInterval, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "SpecialDateAuthority":
		GongUnmarshallSliceOfPointers(&instance.SpecialDateAuthority, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type GuidanceLineLightSystemTimeSlicePropertyTypeUnmarshaller struct{}

func (u *GuidanceLineLightSystemTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(GuidanceLineLightSystemTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *GuidanceLineLightSystemTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GuidanceLineLightSystemTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "GuidanceLineLightSystemTimeSlice":
		GongUnmarshallPointer(&instance.GuidanceLineLightSystemTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type GuidanceLineLightSystemTimeSliceTypeUnmarshaller struct{}

func (u *GuidanceLineLightSystemTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(GuidanceLineLightSystemTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *GuidanceLineLightSystemTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GuidanceLineLightSystemTimeSliceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Interpretation":
		instance.Interpretation = GongExtractString(valueExpr)
	case "SequenceNumber":
		instance.SequenceNumber = GongExtractInt(valueExpr)
	case "CorrectionNumber":
		instance.CorrectionNumber = GongExtractInt(valueExpr)
	case "Element":
		GongUnmarshallSliceOfPointers(&instance.Element, valueExpr, identifierMap)
	case "Availability":
		GongUnmarshallSliceOfPointers(&instance.Availability, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "LightedGuidanceLine":
		GongUnmarshallPointer(&instance.LightedGuidanceLine, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type GuidanceLineLightSystemTypeUnmarshaller struct{}

func (u *GuidanceLineLightSystemTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(GuidanceLineLightSystemType), stage, identifier, instanceName, preserveOrder)
}

func (u *GuidanceLineLightSystemTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GuidanceLineLightSystemType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeSlice":
		GongUnmarshallSliceOfPointers(&instance.TimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type GuidanceLineMarkingTimeSlicePropertyTypeUnmarshaller struct{}

func (u *GuidanceLineMarkingTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(GuidanceLineMarkingTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *GuidanceLineMarkingTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GuidanceLineMarkingTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "GuidanceLineMarkingTimeSlice":
		GongUnmarshallPointer(&instance.GuidanceLineMarkingTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type GuidanceLineMarkingTimeSliceTypeUnmarshaller struct{}

func (u *GuidanceLineMarkingTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(GuidanceLineMarkingTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *GuidanceLineMarkingTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GuidanceLineMarkingTimeSliceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Interpretation":
		instance.Interpretation = GongExtractString(valueExpr)
	case "SequenceNumber":
		instance.SequenceNumber = GongExtractInt(valueExpr)
	case "CorrectionNumber":
		instance.CorrectionNumber = GongExtractInt(valueExpr)
	case "Element":
		GongUnmarshallSliceOfPointers(&instance.Element, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "MarkedGuidanceLine":
		GongUnmarshallPointer(&instance.MarkedGuidanceLine, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type GuidanceLineMarkingTypeUnmarshaller struct{}

func (u *GuidanceLineMarkingTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(GuidanceLineMarkingType), stage, identifier, instanceName, preserveOrder)
}

func (u *GuidanceLineMarkingTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GuidanceLineMarkingType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeSlice":
		GongUnmarshallSliceOfPointers(&instance.TimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type GuidanceLinePropertyTypeUnmarshaller struct{}

func (u *GuidanceLinePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(GuidanceLinePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *GuidanceLinePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GuidanceLinePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "Type_XSD":
		instance.Type_XSD = GongExtractString(valueExpr)
	case "Href":
		instance.Href = GongExtractString(valueExpr)
	case "Role":
		instance.Role = GongExtractString(valueExpr)
	case "Arcrole":
		instance.Arcrole = GongExtractString(valueExpr)
	case "Title":
		instance.Title = GongExtractString(valueExpr)
	case "Show":
		instance.Show = GongExtractString(valueExpr)
	case "Actuate":
		instance.Actuate = GongExtractString(valueExpr)
	}
	return nil
}

type GuidanceLineTimeSlicePropertyTypeUnmarshaller struct{}

func (u *GuidanceLineTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(GuidanceLineTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *GuidanceLineTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GuidanceLineTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "GuidanceLineTimeSlice":
		GongUnmarshallPointer(&instance.GuidanceLineTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type GuidanceLineTimeSliceTypeUnmarshaller struct{}

func (u *GuidanceLineTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(GuidanceLineTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *GuidanceLineTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GuidanceLineTimeSliceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Interpretation":
		instance.Interpretation = GongExtractString(valueExpr)
	case "SequenceNumber":
		instance.SequenceNumber = GongExtractInt(valueExpr)
	case "CorrectionNumber":
		instance.CorrectionNumber = GongExtractInt(valueExpr)
	case "ConnectedTouchDownLiftOff":
		GongUnmarshallSliceOfPointers(&instance.ConnectedTouchDownLiftOff, valueExpr, identifierMap)
	case "ConnectedRunwayCentrelinePoint":
		GongUnmarshallSliceOfPointers(&instance.ConnectedRunwayCentrelinePoint, valueExpr, identifierMap)
	case "ConnectedApron":
		GongUnmarshallSliceOfPointers(&instance.ConnectedApron, valueExpr, identifierMap)
	case "ConnectedStand":
		GongUnmarshallSliceOfPointers(&instance.ConnectedStand, valueExpr, identifierMap)
	case "Extent":
		GongUnmarshallPointer(&instance.Extent, valueExpr, identifierMap)
	case "ConnectedTaxiway":
		GongUnmarshallSliceOfPointers(&instance.ConnectedTaxiway, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type GuidanceLineTypeUnmarshaller struct{}

func (u *GuidanceLineTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(GuidanceLineType), stage, identifier, instanceName, preserveOrder)
}

func (u *GuidanceLineTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GuidanceLineType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeSlice":
		GongUnmarshallSliceOfPointers(&instance.TimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type LightElementPropertyTypeUnmarshaller struct{}

func (u *LightElementPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(LightElementPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *LightElementPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*LightElementPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "LightElement":
		GongUnmarshallPointer(&instance.LightElement, valueExpr, identifierMap)
	}
	return nil
}

type LightElementStatusPropertyTypeUnmarshaller struct{}

func (u *LightElementStatusPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(LightElementStatusPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *LightElementStatusPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*LightElementStatusPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "LightElementStatus":
		GongUnmarshallPointer(&instance.LightElementStatus, valueExpr, identifierMap)
	}
	return nil
}

type LightElementStatusTypeUnmarshaller struct{}

func (u *LightElementStatusTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(LightElementStatusType), stage, identifier, instanceName, preserveOrder)
}

func (u *LightElementStatusTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*LightElementStatusType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeInterval":
		GongUnmarshallSliceOfPointers(&instance.TimeInterval, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "SpecialDateAuthority":
		GongUnmarshallSliceOfPointers(&instance.SpecialDateAuthority, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type LightElementTypeUnmarshaller struct{}

func (u *LightElementTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(LightElementType), stage, identifier, instanceName, preserveOrder)
}

func (u *LightElementTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*LightElementType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Location":
		GongUnmarshallPointer(&instance.Location, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Availability":
		GongUnmarshallSliceOfPointers(&instance.Availability, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type LinguisticNotePropertyTypeUnmarshaller struct{}

func (u *LinguisticNotePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(LinguisticNotePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *LinguisticNotePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*LinguisticNotePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "LinguisticNote":
		GongUnmarshallPointer(&instance.LinguisticNote, valueExpr, identifierMap)
	}
	return nil
}

type LinguisticNoteTypeUnmarshaller struct{}

func (u *LinguisticNoteTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(LinguisticNoteType), stage, identifier, instanceName, preserveOrder)
}

func (u *LinguisticNoteTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*LinguisticNoteType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type ManoeuvringAreaAvailabilityPropertyTypeUnmarshaller struct{}

func (u *ManoeuvringAreaAvailabilityPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ManoeuvringAreaAvailabilityPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *ManoeuvringAreaAvailabilityPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ManoeuvringAreaAvailabilityPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "ManoeuvringAreaAvailability":
		GongUnmarshallPointer(&instance.ManoeuvringAreaAvailability, valueExpr, identifierMap)
	}
	return nil
}

type ManoeuvringAreaAvailabilityTypeUnmarshaller struct{}

func (u *ManoeuvringAreaAvailabilityTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ManoeuvringAreaAvailabilityType), stage, identifier, instanceName, preserveOrder)
}

func (u *ManoeuvringAreaAvailabilityTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ManoeuvringAreaAvailabilityType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeInterval":
		GongUnmarshallSliceOfPointers(&instance.TimeInterval, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "SpecialDateAuthority":
		GongUnmarshallSliceOfPointers(&instance.SpecialDateAuthority, valueExpr, identifierMap)
	case "Usage":
		GongUnmarshallSliceOfPointers(&instance.Usage, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type ManoeuvringAreaUsagePropertyTypeUnmarshaller struct{}

func (u *ManoeuvringAreaUsagePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ManoeuvringAreaUsagePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *ManoeuvringAreaUsagePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ManoeuvringAreaUsagePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "ManoeuvringAreaUsage":
		GongUnmarshallPointer(&instance.ManoeuvringAreaUsage, valueExpr, identifierMap)
	}
	return nil
}

type ManoeuvringAreaUsageTypeUnmarshaller struct{}

func (u *ManoeuvringAreaUsageTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ManoeuvringAreaUsageType), stage, identifier, instanceName, preserveOrder)
}

func (u *ManoeuvringAreaUsageTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ManoeuvringAreaUsageType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Contact":
		GongUnmarshallSliceOfPointers(&instance.Contact, valueExpr, identifierMap)
	case "Selection":
		GongUnmarshallPointer(&instance.Selection, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type MarkingBuoyTimeSlicePropertyTypeUnmarshaller struct{}

func (u *MarkingBuoyTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(MarkingBuoyTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *MarkingBuoyTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MarkingBuoyTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "MarkingBuoyTimeSlice":
		GongUnmarshallPointer(&instance.MarkingBuoyTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type MarkingBuoyTimeSliceTypeUnmarshaller struct{}

func (u *MarkingBuoyTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(MarkingBuoyTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *MarkingBuoyTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MarkingBuoyTimeSliceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Interpretation":
		instance.Interpretation = GongExtractString(valueExpr)
	case "SequenceNumber":
		instance.SequenceNumber = GongExtractInt(valueExpr)
	case "CorrectionNumber":
		instance.CorrectionNumber = GongExtractInt(valueExpr)
	case "TheSeaplaneLandingArea":
		GongUnmarshallPointer(&instance.TheSeaplaneLandingArea, valueExpr, identifierMap)
	case "Location":
		GongUnmarshallPointer(&instance.Location, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type MarkingBuoyTypeUnmarshaller struct{}

func (u *MarkingBuoyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(MarkingBuoyType), stage, identifier, instanceName, preserveOrder)
}

func (u *MarkingBuoyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MarkingBuoyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeSlice":
		GongUnmarshallSliceOfPointers(&instance.TimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type MarkingElementPropertyTypeUnmarshaller struct{}

func (u *MarkingElementPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(MarkingElementPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *MarkingElementPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MarkingElementPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "MarkingElement":
		GongUnmarshallPointer(&instance.MarkingElement, valueExpr, identifierMap)
	}
	return nil
}

type MarkingElementTypeUnmarshaller struct{}

func (u *MarkingElementTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(MarkingElementType), stage, identifier, instanceName, preserveOrder)
}

func (u *MarkingElementTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MarkingElementType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Extent_curveExtent":
		GongUnmarshallPointer(&instance.Extent_curveExtent, valueExpr, identifierMap)
	case "Extent_surfaceExtent":
		GongUnmarshallPointer(&instance.Extent_surfaceExtent, valueExpr, identifierMap)
	case "Extent_location":
		GongUnmarshallPointer(&instance.Extent_location, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type MeteorologyPropertyTypeUnmarshaller struct{}

func (u *MeteorologyPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(MeteorologyPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *MeteorologyPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MeteorologyPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "Meteorology":
		GongUnmarshallPointer(&instance.Meteorology, valueExpr, identifierMap)
	}
	return nil
}

type MeteorologyTypeUnmarshaller struct{}

func (u *MeteorologyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(MeteorologyType), stage, identifier, instanceName, preserveOrder)
}

func (u *MeteorologyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MeteorologyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type NavaidEquipmentDistancePropertyTypeUnmarshaller struct{}

func (u *NavaidEquipmentDistancePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(NavaidEquipmentDistancePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *NavaidEquipmentDistancePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*NavaidEquipmentDistancePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "NavaidEquipmentDistance":
		GongUnmarshallPointer(&instance.NavaidEquipmentDistance, valueExpr, identifierMap)
	}
	return nil
}

type NavaidEquipmentDistanceTypeUnmarshaller struct{}

func (u *NavaidEquipmentDistanceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(NavaidEquipmentDistanceType), stage, identifier, instanceName, preserveOrder)
}

func (u *NavaidEquipmentDistanceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*NavaidEquipmentDistanceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "TheNavaidEquipment":
		GongUnmarshallPointer(&instance.TheNavaidEquipment, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type NavaidEquipmentPropertyTypeUnmarshaller struct{}

func (u *NavaidEquipmentPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(NavaidEquipmentPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *NavaidEquipmentPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*NavaidEquipmentPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "Type_XSD":
		instance.Type_XSD = GongExtractString(valueExpr)
	case "Href":
		instance.Href = GongExtractString(valueExpr)
	case "Role":
		instance.Role = GongExtractString(valueExpr)
	case "Arcrole":
		instance.Arcrole = GongExtractString(valueExpr)
	case "Title":
		instance.Title = GongExtractString(valueExpr)
	case "Show":
		instance.Show = GongExtractString(valueExpr)
	case "Actuate":
		instance.Actuate = GongExtractString(valueExpr)
	}
	return nil
}

type NonMovementAreaTimeSlicePropertyTypeUnmarshaller struct{}

func (u *NonMovementAreaTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(NonMovementAreaTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *NonMovementAreaTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*NonMovementAreaTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "NonMovementAreaTimeSlice":
		GongUnmarshallPointer(&instance.NonMovementAreaTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type NonMovementAreaTimeSliceTypeUnmarshaller struct{}

func (u *NonMovementAreaTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(NonMovementAreaTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *NonMovementAreaTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*NonMovementAreaTimeSliceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Interpretation":
		instance.Interpretation = GongExtractString(valueExpr)
	case "SequenceNumber":
		instance.SequenceNumber = GongExtractInt(valueExpr)
	case "CorrectionNumber":
		instance.CorrectionNumber = GongExtractInt(valueExpr)
	case "AssociatedAirportHeliport":
		GongUnmarshallPointer(&instance.AssociatedAirportHeliport, valueExpr, identifierMap)
	case "Extent":
		GongUnmarshallPointer(&instance.Extent, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type NonMovementAreaTypeUnmarshaller struct{}

func (u *NonMovementAreaTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(NonMovementAreaType), stage, identifier, instanceName, preserveOrder)
}

func (u *NonMovementAreaTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*NonMovementAreaType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeSlice":
		GongUnmarshallSliceOfPointers(&instance.TimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type NotePropertyTypeUnmarshaller struct{}

func (u *NotePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(NotePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *NotePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*NotePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "Note":
		GongUnmarshallPointer(&instance.Note, valueExpr, identifierMap)
	}
	return nil
}

type NoteTypeUnmarshaller struct{}

func (u *NoteTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(NoteType), stage, identifier, instanceName, preserveOrder)
}

func (u *NoteTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*NoteType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TranslatedNote":
		GongUnmarshallSliceOfPointers(&instance.TranslatedNote, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type OnlineContactPropertyTypeUnmarshaller struct{}

func (u *OnlineContactPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(OnlineContactPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *OnlineContactPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*OnlineContactPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "OnlineContact":
		GongUnmarshallPointer(&instance.OnlineContact, valueExpr, identifierMap)
	}
	return nil
}

type OnlineContactTypeUnmarshaller struct{}

func (u *OnlineContactTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(OnlineContactType), stage, identifier, instanceName, preserveOrder)
}

func (u *OnlineContactTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*OnlineContactType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeInterval":
		GongUnmarshallSliceOfPointers(&instance.TimeInterval, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "SpecialDateAuthority":
		GongUnmarshallSliceOfPointers(&instance.SpecialDateAuthority, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type OrganisationAuthorityPropertyTypeUnmarshaller struct{}

func (u *OrganisationAuthorityPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(OrganisationAuthorityPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *OrganisationAuthorityPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*OrganisationAuthorityPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "Type_XSD":
		instance.Type_XSD = GongExtractString(valueExpr)
	case "Href":
		instance.Href = GongExtractString(valueExpr)
	case "Role":
		instance.Role = GongExtractString(valueExpr)
	case "Arcrole":
		instance.Arcrole = GongExtractString(valueExpr)
	case "Title":
		instance.Title = GongExtractString(valueExpr)
	case "Show":
		instance.Show = GongExtractString(valueExpr)
	case "Actuate":
		instance.Actuate = GongExtractString(valueExpr)
	}
	return nil
}

type PassengerLoadingBridgeTimeSlicePropertyTypeUnmarshaller struct{}

func (u *PassengerLoadingBridgeTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PassengerLoadingBridgeTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *PassengerLoadingBridgeTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PassengerLoadingBridgeTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "PassengerLoadingBridgeTimeSlice":
		GongUnmarshallPointer(&instance.PassengerLoadingBridgeTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type PassengerLoadingBridgeTimeSliceTypeUnmarshaller struct{}

func (u *PassengerLoadingBridgeTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PassengerLoadingBridgeTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *PassengerLoadingBridgeTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PassengerLoadingBridgeTimeSliceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Interpretation":
		instance.Interpretation = GongExtractString(valueExpr)
	case "SequenceNumber":
		instance.SequenceNumber = GongExtractInt(valueExpr)
	case "CorrectionNumber":
		instance.CorrectionNumber = GongExtractInt(valueExpr)
	case "Extent":
		GongUnmarshallPointer(&instance.Extent, valueExpr, identifierMap)
	case "AssociatedStand":
		GongUnmarshallSliceOfPointers(&instance.AssociatedStand, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type PassengerLoadingBridgeTypeUnmarshaller struct{}

func (u *PassengerLoadingBridgeTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PassengerLoadingBridgeType), stage, identifier, instanceName, preserveOrder)
}

func (u *PassengerLoadingBridgeTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PassengerLoadingBridgeType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeSlice":
		GongUnmarshallSliceOfPointers(&instance.TimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type PointTypeUnmarshaller struct{}

func (u *PointTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PointType), stage, identifier, instanceName, preserveOrder)
}

func (u *PointTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PointType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "SrsName":
		instance.SrsName = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	}
	return nil
}

type PostalAddressPropertyTypeUnmarshaller struct{}

func (u *PostalAddressPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PostalAddressPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *PostalAddressPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PostalAddressPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "PostalAddress":
		GongUnmarshallPointer(&instance.PostalAddress, valueExpr, identifierMap)
	}
	return nil
}

type PostalAddressTypeUnmarshaller struct{}

func (u *PostalAddressTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PostalAddressType), stage, identifier, instanceName, preserveOrder)
}

func (u *PostalAddressTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PostalAddressType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeInterval":
		GongUnmarshallSliceOfPointers(&instance.TimeInterval, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "SpecialDateAuthority":
		GongUnmarshallSliceOfPointers(&instance.SpecialDateAuthority, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type RidgePropertyTypeUnmarshaller struct{}

func (u *RidgePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RidgePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *RidgePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RidgePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "Ridge":
		GongUnmarshallPointer(&instance.Ridge, valueExpr, identifierMap)
	}
	return nil
}

type RidgeTypeUnmarshaller struct{}

func (u *RidgeTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RidgeType), stage, identifier, instanceName, preserveOrder)
}

func (u *RidgeTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RidgeType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type RoadTimeSlicePropertyTypeUnmarshaller struct{}

func (u *RoadTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RoadTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *RoadTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RoadTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "RoadTimeSlice":
		GongUnmarshallPointer(&instance.RoadTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type RoadTimeSliceTypeUnmarshaller struct{}

func (u *RoadTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RoadTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *RoadTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RoadTimeSliceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Interpretation":
		instance.Interpretation = GongExtractString(valueExpr)
	case "SequenceNumber":
		instance.SequenceNumber = GongExtractInt(valueExpr)
	case "CorrectionNumber":
		instance.CorrectionNumber = GongExtractInt(valueExpr)
	case "AssociatedAirport":
		GongUnmarshallPointer(&instance.AssociatedAirport, valueExpr, identifierMap)
	case "SurfaceProperties":
		GongUnmarshallPointer(&instance.SurfaceProperties, valueExpr, identifierMap)
	case "AccessibleStand":
		GongUnmarshallSliceOfPointers(&instance.AccessibleStand, valueExpr, identifierMap)
	case "SurfaceExtent":
		GongUnmarshallPointer(&instance.SurfaceExtent, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type RoadTypeUnmarshaller struct{}

func (u *RoadTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RoadType), stage, identifier, instanceName, preserveOrder)
}

func (u *RoadTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RoadType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeSlice":
		GongUnmarshallSliceOfPointers(&instance.TimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type RunwayBlastPadTimeSlicePropertyTypeUnmarshaller struct{}

func (u *RunwayBlastPadTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RunwayBlastPadTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *RunwayBlastPadTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RunwayBlastPadTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "RunwayBlastPadTimeSlice":
		GongUnmarshallPointer(&instance.RunwayBlastPadTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type RunwayBlastPadTimeSliceTypeUnmarshaller struct{}

func (u *RunwayBlastPadTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RunwayBlastPadTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *RunwayBlastPadTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RunwayBlastPadTimeSliceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Interpretation":
		instance.Interpretation = GongExtractString(valueExpr)
	case "SequenceNumber":
		instance.SequenceNumber = GongExtractInt(valueExpr)
	case "CorrectionNumber":
		instance.CorrectionNumber = GongExtractInt(valueExpr)
	case "UsedRunwayDirection":
		GongUnmarshallPointer(&instance.UsedRunwayDirection, valueExpr, identifierMap)
	case "Extent":
		GongUnmarshallPointer(&instance.Extent, valueExpr, identifierMap)
	case "SurfaceProperties":
		GongUnmarshallPointer(&instance.SurfaceProperties, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type RunwayBlastPadTypeUnmarshaller struct{}

func (u *RunwayBlastPadTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RunwayBlastPadType), stage, identifier, instanceName, preserveOrder)
}

func (u *RunwayBlastPadTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RunwayBlastPadType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeSlice":
		GongUnmarshallSliceOfPointers(&instance.TimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type RunwayCentrelinePointPropertyTypeUnmarshaller struct{}

func (u *RunwayCentrelinePointPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RunwayCentrelinePointPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *RunwayCentrelinePointPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RunwayCentrelinePointPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "Type_XSD":
		instance.Type_XSD = GongExtractString(valueExpr)
	case "Href":
		instance.Href = GongExtractString(valueExpr)
	case "Role":
		instance.Role = GongExtractString(valueExpr)
	case "Arcrole":
		instance.Arcrole = GongExtractString(valueExpr)
	case "Title":
		instance.Title = GongExtractString(valueExpr)
	case "Show":
		instance.Show = GongExtractString(valueExpr)
	case "Actuate":
		instance.Actuate = GongExtractString(valueExpr)
	}
	return nil
}

type RunwayCentrelinePointTimeSlicePropertyTypeUnmarshaller struct{}

func (u *RunwayCentrelinePointTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RunwayCentrelinePointTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *RunwayCentrelinePointTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RunwayCentrelinePointTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "RunwayCentrelinePointTimeSlice":
		GongUnmarshallPointer(&instance.RunwayCentrelinePointTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type RunwayCentrelinePointTimeSliceTypeUnmarshaller struct{}

func (u *RunwayCentrelinePointTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RunwayCentrelinePointTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *RunwayCentrelinePointTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RunwayCentrelinePointTimeSliceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Interpretation":
		instance.Interpretation = GongExtractString(valueExpr)
	case "SequenceNumber":
		instance.SequenceNumber = GongExtractInt(valueExpr)
	case "CorrectionNumber":
		instance.CorrectionNumber = GongExtractInt(valueExpr)
	case "Location":
		GongUnmarshallPointer(&instance.Location, valueExpr, identifierMap)
	case "OnRunway":
		GongUnmarshallPointer(&instance.OnRunway, valueExpr, identifierMap)
	case "AssociatedDeclaredDistance":
		GongUnmarshallSliceOfPointers(&instance.AssociatedDeclaredDistance, valueExpr, identifierMap)
	case "NavaidEquipment":
		GongUnmarshallSliceOfPointers(&instance.NavaidEquipment, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type RunwayCentrelinePointTypeUnmarshaller struct{}

func (u *RunwayCentrelinePointTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RunwayCentrelinePointType), stage, identifier, instanceName, preserveOrder)
}

func (u *RunwayCentrelinePointTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RunwayCentrelinePointType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeSlice":
		GongUnmarshallSliceOfPointers(&instance.TimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type RunwayContaminationPropertyTypeUnmarshaller struct{}

func (u *RunwayContaminationPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RunwayContaminationPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *RunwayContaminationPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RunwayContaminationPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "RunwayContamination":
		GongUnmarshallPointer(&instance.RunwayContamination, valueExpr, identifierMap)
	}
	return nil
}

type RunwayContaminationTypeUnmarshaller struct{}

func (u *RunwayContaminationTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RunwayContaminationType), stage, identifier, instanceName, preserveOrder)
}

func (u *RunwayContaminationTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RunwayContaminationType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "CriticalRidge":
		GongUnmarshallSliceOfPointers(&instance.CriticalRidge, valueExpr, identifierMap)
	case "Layer":
		GongUnmarshallSliceOfPointers(&instance.Layer, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type RunwayDeclaredDistancePropertyTypeUnmarshaller struct{}

func (u *RunwayDeclaredDistancePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RunwayDeclaredDistancePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *RunwayDeclaredDistancePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RunwayDeclaredDistancePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "RunwayDeclaredDistance":
		GongUnmarshallPointer(&instance.RunwayDeclaredDistance, valueExpr, identifierMap)
	}
	return nil
}

type RunwayDeclaredDistanceTypeUnmarshaller struct{}

func (u *RunwayDeclaredDistanceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RunwayDeclaredDistanceType), stage, identifier, instanceName, preserveOrder)
}

func (u *RunwayDeclaredDistanceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RunwayDeclaredDistanceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "DeclaredValue":
		GongUnmarshallSliceOfPointers(&instance.DeclaredValue, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type RunwayDeclaredDistanceValuePropertyTypeUnmarshaller struct{}

func (u *RunwayDeclaredDistanceValuePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RunwayDeclaredDistanceValuePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *RunwayDeclaredDistanceValuePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RunwayDeclaredDistanceValuePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "RunwayDeclaredDistanceValue":
		GongUnmarshallPointer(&instance.RunwayDeclaredDistanceValue, valueExpr, identifierMap)
	}
	return nil
}

type RunwayDeclaredDistanceValueTypeUnmarshaller struct{}

func (u *RunwayDeclaredDistanceValueTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RunwayDeclaredDistanceValueType), stage, identifier, instanceName, preserveOrder)
}

func (u *RunwayDeclaredDistanceValueTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RunwayDeclaredDistanceValueType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeInterval":
		GongUnmarshallSliceOfPointers(&instance.TimeInterval, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "SpecialDateAuthority":
		GongUnmarshallSliceOfPointers(&instance.SpecialDateAuthority, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type RunwayDirectionLightSystemTimeSlicePropertyTypeUnmarshaller struct{}

func (u *RunwayDirectionLightSystemTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RunwayDirectionLightSystemTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *RunwayDirectionLightSystemTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RunwayDirectionLightSystemTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "RunwayDirectionLightSystemTimeSlice":
		GongUnmarshallPointer(&instance.RunwayDirectionLightSystemTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type RunwayDirectionLightSystemTimeSliceTypeUnmarshaller struct{}

func (u *RunwayDirectionLightSystemTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RunwayDirectionLightSystemTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *RunwayDirectionLightSystemTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RunwayDirectionLightSystemTimeSliceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Interpretation":
		instance.Interpretation = GongExtractString(valueExpr)
	case "SequenceNumber":
		instance.SequenceNumber = GongExtractInt(valueExpr)
	case "CorrectionNumber":
		instance.CorrectionNumber = GongExtractInt(valueExpr)
	case "Element":
		GongUnmarshallSliceOfPointers(&instance.Element, valueExpr, identifierMap)
	case "Availability":
		GongUnmarshallSliceOfPointers(&instance.Availability, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "AssociatedRunwayDirection":
		GongUnmarshallPointer(&instance.AssociatedRunwayDirection, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type RunwayDirectionLightSystemTypeUnmarshaller struct{}

func (u *RunwayDirectionLightSystemTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RunwayDirectionLightSystemType), stage, identifier, instanceName, preserveOrder)
}

func (u *RunwayDirectionLightSystemTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RunwayDirectionLightSystemType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeSlice":
		GongUnmarshallSliceOfPointers(&instance.TimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type RunwayDirectionPropertyTypeUnmarshaller struct{}

func (u *RunwayDirectionPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RunwayDirectionPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *RunwayDirectionPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RunwayDirectionPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "Type_XSD":
		instance.Type_XSD = GongExtractString(valueExpr)
	case "Href":
		instance.Href = GongExtractString(valueExpr)
	case "Role":
		instance.Role = GongExtractString(valueExpr)
	case "Arcrole":
		instance.Arcrole = GongExtractString(valueExpr)
	case "Title":
		instance.Title = GongExtractString(valueExpr)
	case "Show":
		instance.Show = GongExtractString(valueExpr)
	case "Actuate":
		instance.Actuate = GongExtractString(valueExpr)
	}
	return nil
}

type RunwayDirectionTimeSlicePropertyTypeUnmarshaller struct{}

func (u *RunwayDirectionTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RunwayDirectionTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *RunwayDirectionTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RunwayDirectionTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "RunwayDirectionTimeSlice":
		GongUnmarshallPointer(&instance.RunwayDirectionTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type RunwayDirectionTimeSliceTypeUnmarshaller struct{}

func (u *RunwayDirectionTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RunwayDirectionTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *RunwayDirectionTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RunwayDirectionTimeSliceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Interpretation":
		instance.Interpretation = GongExtractString(valueExpr)
	case "SequenceNumber":
		instance.SequenceNumber = GongExtractInt(valueExpr)
	case "CorrectionNumber":
		instance.CorrectionNumber = GongExtractInt(valueExpr)
	case "UsedRunway":
		GongUnmarshallPointer(&instance.UsedRunway, valueExpr, identifierMap)
	case "StartingElement":
		GongUnmarshallPointer(&instance.StartingElement, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Availability":
		GongUnmarshallSliceOfPointers(&instance.Availability, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type RunwayDirectionTypeUnmarshaller struct{}

func (u *RunwayDirectionTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RunwayDirectionType), stage, identifier, instanceName, preserveOrder)
}

func (u *RunwayDirectionTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RunwayDirectionType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeSlice":
		GongUnmarshallSliceOfPointers(&instance.TimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type RunwayElementPropertyTypeUnmarshaller struct{}

func (u *RunwayElementPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RunwayElementPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *RunwayElementPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RunwayElementPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "Type_XSD":
		instance.Type_XSD = GongExtractString(valueExpr)
	case "Href":
		instance.Href = GongExtractString(valueExpr)
	case "Role":
		instance.Role = GongExtractString(valueExpr)
	case "Arcrole":
		instance.Arcrole = GongExtractString(valueExpr)
	case "Title":
		instance.Title = GongExtractString(valueExpr)
	case "Show":
		instance.Show = GongExtractString(valueExpr)
	case "Actuate":
		instance.Actuate = GongExtractString(valueExpr)
	}
	return nil
}

type RunwayElementTimeSlicePropertyTypeUnmarshaller struct{}

func (u *RunwayElementTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RunwayElementTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *RunwayElementTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RunwayElementTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "RunwayElementTimeSlice":
		GongUnmarshallPointer(&instance.RunwayElementTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type RunwayElementTimeSliceTypeUnmarshaller struct{}

func (u *RunwayElementTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RunwayElementTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *RunwayElementTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RunwayElementTimeSliceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Interpretation":
		instance.Interpretation = GongExtractString(valueExpr)
	case "SequenceNumber":
		instance.SequenceNumber = GongExtractInt(valueExpr)
	case "CorrectionNumber":
		instance.CorrectionNumber = GongExtractInt(valueExpr)
	case "SurfaceProperties":
		GongUnmarshallPointer(&instance.SurfaceProperties, valueExpr, identifierMap)
	case "AssociatedRunway":
		GongUnmarshallSliceOfPointers(&instance.AssociatedRunway, valueExpr, identifierMap)
	case "Extent":
		GongUnmarshallPointer(&instance.Extent, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Availability":
		GongUnmarshallSliceOfPointers(&instance.Availability, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type RunwayElementTypeUnmarshaller struct{}

func (u *RunwayElementTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RunwayElementType), stage, identifier, instanceName, preserveOrder)
}

func (u *RunwayElementTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RunwayElementType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeSlice":
		GongUnmarshallSliceOfPointers(&instance.TimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type RunwayMarkingTimeSlicePropertyTypeUnmarshaller struct{}

func (u *RunwayMarkingTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RunwayMarkingTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *RunwayMarkingTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RunwayMarkingTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "RunwayMarkingTimeSlice":
		GongUnmarshallPointer(&instance.RunwayMarkingTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type RunwayMarkingTimeSliceTypeUnmarshaller struct{}

func (u *RunwayMarkingTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RunwayMarkingTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *RunwayMarkingTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RunwayMarkingTimeSliceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Interpretation":
		instance.Interpretation = GongExtractString(valueExpr)
	case "SequenceNumber":
		instance.SequenceNumber = GongExtractInt(valueExpr)
	case "CorrectionNumber":
		instance.CorrectionNumber = GongExtractInt(valueExpr)
	case "Element":
		GongUnmarshallSliceOfPointers(&instance.Element, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "MarkedRunway":
		GongUnmarshallPointer(&instance.MarkedRunway, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type RunwayMarkingTypeUnmarshaller struct{}

func (u *RunwayMarkingTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RunwayMarkingType), stage, identifier, instanceName, preserveOrder)
}

func (u *RunwayMarkingTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RunwayMarkingType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeSlice":
		GongUnmarshallSliceOfPointers(&instance.TimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type RunwayPropertyTypeUnmarshaller struct{}

func (u *RunwayPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RunwayPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *RunwayPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RunwayPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "Type_XSD":
		instance.Type_XSD = GongExtractString(valueExpr)
	case "Href":
		instance.Href = GongExtractString(valueExpr)
	case "Role":
		instance.Role = GongExtractString(valueExpr)
	case "Arcrole":
		instance.Arcrole = GongExtractString(valueExpr)
	case "Title":
		instance.Title = GongExtractString(valueExpr)
	case "Show":
		instance.Show = GongExtractString(valueExpr)
	case "Actuate":
		instance.Actuate = GongExtractString(valueExpr)
	}
	return nil
}

type RunwayProtectAreaLightSystemTimeSlicePropertyTypeUnmarshaller struct{}

func (u *RunwayProtectAreaLightSystemTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RunwayProtectAreaLightSystemTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *RunwayProtectAreaLightSystemTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RunwayProtectAreaLightSystemTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "RunwayProtectAreaLightSystemTimeSlice":
		GongUnmarshallPointer(&instance.RunwayProtectAreaLightSystemTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type RunwayProtectAreaLightSystemTimeSliceTypeUnmarshaller struct{}

func (u *RunwayProtectAreaLightSystemTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RunwayProtectAreaLightSystemTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *RunwayProtectAreaLightSystemTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RunwayProtectAreaLightSystemTimeSliceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Interpretation":
		instance.Interpretation = GongExtractString(valueExpr)
	case "SequenceNumber":
		instance.SequenceNumber = GongExtractInt(valueExpr)
	case "CorrectionNumber":
		instance.CorrectionNumber = GongExtractInt(valueExpr)
	case "Element":
		GongUnmarshallSliceOfPointers(&instance.Element, valueExpr, identifierMap)
	case "Availability":
		GongUnmarshallSliceOfPointers(&instance.Availability, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "LightedArea":
		GongUnmarshallPointer(&instance.LightedArea, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type RunwayProtectAreaLightSystemTypeUnmarshaller struct{}

func (u *RunwayProtectAreaLightSystemTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RunwayProtectAreaLightSystemType), stage, identifier, instanceName, preserveOrder)
}

func (u *RunwayProtectAreaLightSystemTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RunwayProtectAreaLightSystemType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeSlice":
		GongUnmarshallSliceOfPointers(&instance.TimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type RunwayProtectAreaPropertyTypeUnmarshaller struct{}

func (u *RunwayProtectAreaPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RunwayProtectAreaPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *RunwayProtectAreaPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RunwayProtectAreaPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "Type_XSD":
		instance.Type_XSD = GongExtractString(valueExpr)
	case "Href":
		instance.Href = GongExtractString(valueExpr)
	case "Role":
		instance.Role = GongExtractString(valueExpr)
	case "Arcrole":
		instance.Arcrole = GongExtractString(valueExpr)
	case "Title":
		instance.Title = GongExtractString(valueExpr)
	case "Show":
		instance.Show = GongExtractString(valueExpr)
	case "Actuate":
		instance.Actuate = GongExtractString(valueExpr)
	}
	return nil
}

type RunwayProtectAreaTimeSlicePropertyTypeUnmarshaller struct{}

func (u *RunwayProtectAreaTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RunwayProtectAreaTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *RunwayProtectAreaTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RunwayProtectAreaTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "RunwayProtectAreaTimeSlice":
		GongUnmarshallPointer(&instance.RunwayProtectAreaTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type RunwayProtectAreaTimeSliceTypeUnmarshaller struct{}

func (u *RunwayProtectAreaTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RunwayProtectAreaTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *RunwayProtectAreaTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RunwayProtectAreaTimeSliceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Interpretation":
		instance.Interpretation = GongExtractString(valueExpr)
	case "SequenceNumber":
		instance.SequenceNumber = GongExtractInt(valueExpr)
	case "CorrectionNumber":
		instance.CorrectionNumber = GongExtractInt(valueExpr)
	case "SurfaceProperties":
		GongUnmarshallPointer(&instance.SurfaceProperties, valueExpr, identifierMap)
	case "Extent":
		GongUnmarshallPointer(&instance.Extent, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "ProtectedRunwayDirection":
		GongUnmarshallPointer(&instance.ProtectedRunwayDirection, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type RunwayProtectAreaTypeUnmarshaller struct{}

func (u *RunwayProtectAreaTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RunwayProtectAreaType), stage, identifier, instanceName, preserveOrder)
}

func (u *RunwayProtectAreaTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RunwayProtectAreaType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeSlice":
		GongUnmarshallSliceOfPointers(&instance.TimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type RunwaySectionContaminationPropertyTypeUnmarshaller struct{}

func (u *RunwaySectionContaminationPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RunwaySectionContaminationPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *RunwaySectionContaminationPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RunwaySectionContaminationPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "RunwaySectionContamination":
		GongUnmarshallPointer(&instance.RunwaySectionContamination, valueExpr, identifierMap)
	}
	return nil
}

type RunwaySectionContaminationTypeUnmarshaller struct{}

func (u *RunwaySectionContaminationTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RunwaySectionContaminationType), stage, identifier, instanceName, preserveOrder)
}

func (u *RunwaySectionContaminationTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RunwaySectionContaminationType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "CriticalRidge":
		GongUnmarshallSliceOfPointers(&instance.CriticalRidge, valueExpr, identifierMap)
	case "Layer":
		GongUnmarshallSliceOfPointers(&instance.Layer, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type RunwayTimeSlicePropertyTypeUnmarshaller struct{}

func (u *RunwayTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RunwayTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *RunwayTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RunwayTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "RunwayTimeSlice":
		GongUnmarshallPointer(&instance.RunwayTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type RunwayTimeSliceTypeUnmarshaller struct{}

func (u *RunwayTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RunwayTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *RunwayTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RunwayTimeSliceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Interpretation":
		instance.Interpretation = GongExtractString(valueExpr)
	case "SequenceNumber":
		instance.SequenceNumber = GongExtractInt(valueExpr)
	case "CorrectionNumber":
		instance.CorrectionNumber = GongExtractInt(valueExpr)
	case "SurfaceProperties":
		GongUnmarshallPointer(&instance.SurfaceProperties, valueExpr, identifierMap)
	case "AssociatedAirportHeliport":
		GongUnmarshallPointer(&instance.AssociatedAirportHeliport, valueExpr, identifierMap)
	case "OverallContaminant":
		GongUnmarshallSliceOfPointers(&instance.OverallContaminant, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "AreaContaminant":
		GongUnmarshallSliceOfPointers(&instance.AreaContaminant, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type RunwayTypeUnmarshaller struct{}

func (u *RunwayTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RunwayType), stage, identifier, instanceName, preserveOrder)
}

func (u *RunwayTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RunwayType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeSlice":
		GongUnmarshallSliceOfPointers(&instance.TimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type RunwayVisualRangeTimeSlicePropertyTypeUnmarshaller struct{}

func (u *RunwayVisualRangeTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RunwayVisualRangeTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *RunwayVisualRangeTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RunwayVisualRangeTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "RunwayVisualRangeTimeSlice":
		GongUnmarshallPointer(&instance.RunwayVisualRangeTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type RunwayVisualRangeTimeSliceTypeUnmarshaller struct{}

func (u *RunwayVisualRangeTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RunwayVisualRangeTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *RunwayVisualRangeTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RunwayVisualRangeTimeSliceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Interpretation":
		instance.Interpretation = GongExtractString(valueExpr)
	case "SequenceNumber":
		instance.SequenceNumber = GongExtractInt(valueExpr)
	case "CorrectionNumber":
		instance.CorrectionNumber = GongExtractInt(valueExpr)
	case "AssociatedRunwayDirection":
		GongUnmarshallSliceOfPointers(&instance.AssociatedRunwayDirection, valueExpr, identifierMap)
	case "Location":
		GongUnmarshallPointer(&instance.Location, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type RunwayVisualRangeTypeUnmarshaller struct{}

func (u *RunwayVisualRangeTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RunwayVisualRangeType), stage, identifier, instanceName, preserveOrder)
}

func (u *RunwayVisualRangeTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RunwayVisualRangeType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeSlice":
		GongUnmarshallSliceOfPointers(&instance.TimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type SeaplaneLandingAreaPropertyTypeUnmarshaller struct{}

func (u *SeaplaneLandingAreaPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SeaplaneLandingAreaPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *SeaplaneLandingAreaPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SeaplaneLandingAreaPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "Type_XSD":
		instance.Type_XSD = GongExtractString(valueExpr)
	case "Href":
		instance.Href = GongExtractString(valueExpr)
	case "Role":
		instance.Role = GongExtractString(valueExpr)
	case "Arcrole":
		instance.Arcrole = GongExtractString(valueExpr)
	case "Title":
		instance.Title = GongExtractString(valueExpr)
	case "Show":
		instance.Show = GongExtractString(valueExpr)
	case "Actuate":
		instance.Actuate = GongExtractString(valueExpr)
	}
	return nil
}

type SeaplaneLandingAreaTimeSlicePropertyTypeUnmarshaller struct{}

func (u *SeaplaneLandingAreaTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SeaplaneLandingAreaTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *SeaplaneLandingAreaTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SeaplaneLandingAreaTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "SeaplaneLandingAreaTimeSlice":
		GongUnmarshallPointer(&instance.SeaplaneLandingAreaTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type SeaplaneLandingAreaTimeSliceTypeUnmarshaller struct{}

func (u *SeaplaneLandingAreaTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SeaplaneLandingAreaTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *SeaplaneLandingAreaTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SeaplaneLandingAreaTimeSliceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Interpretation":
		instance.Interpretation = GongExtractString(valueExpr)
	case "SequenceNumber":
		instance.SequenceNumber = GongExtractInt(valueExpr)
	case "CorrectionNumber":
		instance.CorrectionNumber = GongExtractInt(valueExpr)
	case "RampSite":
		GongUnmarshallSliceOfPointers(&instance.RampSite, valueExpr, identifierMap)
	case "DockSite":
		GongUnmarshallSliceOfPointers(&instance.DockSite, valueExpr, identifierMap)
	case "Extent":
		GongUnmarshallPointer(&instance.Extent, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Availability":
		GongUnmarshallSliceOfPointers(&instance.Availability, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type SeaplaneLandingAreaTypeUnmarshaller struct{}

func (u *SeaplaneLandingAreaTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SeaplaneLandingAreaType), stage, identifier, instanceName, preserveOrder)
}

func (u *SeaplaneLandingAreaTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SeaplaneLandingAreaType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeSlice":
		GongUnmarshallSliceOfPointers(&instance.TimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type SeaplaneRampSitePropertyTypeUnmarshaller struct{}

func (u *SeaplaneRampSitePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SeaplaneRampSitePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *SeaplaneRampSitePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SeaplaneRampSitePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "Type_XSD":
		instance.Type_XSD = GongExtractString(valueExpr)
	case "Href":
		instance.Href = GongExtractString(valueExpr)
	case "Role":
		instance.Role = GongExtractString(valueExpr)
	case "Arcrole":
		instance.Arcrole = GongExtractString(valueExpr)
	case "Title":
		instance.Title = GongExtractString(valueExpr)
	case "Show":
		instance.Show = GongExtractString(valueExpr)
	case "Actuate":
		instance.Actuate = GongExtractString(valueExpr)
	}
	return nil
}

type SeaplaneRampSiteTimeSlicePropertyTypeUnmarshaller struct{}

func (u *SeaplaneRampSiteTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SeaplaneRampSiteTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *SeaplaneRampSiteTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SeaplaneRampSiteTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "SeaplaneRampSiteTimeSlice":
		GongUnmarshallPointer(&instance.SeaplaneRampSiteTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type SeaplaneRampSiteTimeSliceTypeUnmarshaller struct{}

func (u *SeaplaneRampSiteTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SeaplaneRampSiteTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *SeaplaneRampSiteTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SeaplaneRampSiteTimeSliceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Interpretation":
		instance.Interpretation = GongExtractString(valueExpr)
	case "SequenceNumber":
		instance.SequenceNumber = GongExtractInt(valueExpr)
	case "CorrectionNumber":
		instance.CorrectionNumber = GongExtractInt(valueExpr)
	case "Extent":
		GongUnmarshallPointer(&instance.Extent, valueExpr, identifierMap)
	case "Centreline":
		GongUnmarshallPointer(&instance.Centreline, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type SeaplaneRampSiteTypeUnmarshaller struct{}

func (u *SeaplaneRampSiteTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SeaplaneRampSiteType), stage, identifier, instanceName, preserveOrder)
}

func (u *SeaplaneRampSiteTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SeaplaneRampSiteType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeSlice":
		GongUnmarshallSliceOfPointers(&instance.TimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type StandMarkingTimeSlicePropertyTypeUnmarshaller struct{}

func (u *StandMarkingTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(StandMarkingTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *StandMarkingTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StandMarkingTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "StandMarkingTimeSlice":
		GongUnmarshallPointer(&instance.StandMarkingTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type StandMarkingTimeSliceTypeUnmarshaller struct{}

func (u *StandMarkingTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(StandMarkingTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *StandMarkingTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StandMarkingTimeSliceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Interpretation":
		instance.Interpretation = GongExtractString(valueExpr)
	case "SequenceNumber":
		instance.SequenceNumber = GongExtractInt(valueExpr)
	case "CorrectionNumber":
		instance.CorrectionNumber = GongExtractInt(valueExpr)
	case "Element":
		GongUnmarshallSliceOfPointers(&instance.Element, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "MarkedStand":
		GongUnmarshallPointer(&instance.MarkedStand, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type StandMarkingTypeUnmarshaller struct{}

func (u *StandMarkingTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(StandMarkingType), stage, identifier, instanceName, preserveOrder)
}

func (u *StandMarkingTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StandMarkingType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeSlice":
		GongUnmarshallSliceOfPointers(&instance.TimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type SurfaceCharacteristicsPropertyTypeUnmarshaller struct{}

func (u *SurfaceCharacteristicsPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SurfaceCharacteristicsPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *SurfaceCharacteristicsPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SurfaceCharacteristicsPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "SurfaceCharacteristics":
		GongUnmarshallPointer(&instance.SurfaceCharacteristics, valueExpr, identifierMap)
	}
	return nil
}

type SurfaceCharacteristicsTypeUnmarshaller struct{}

func (u *SurfaceCharacteristicsTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SurfaceCharacteristicsType), stage, identifier, instanceName, preserveOrder)
}

func (u *SurfaceCharacteristicsTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SurfaceCharacteristicsType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type SurfaceContaminationLayerPropertyTypeUnmarshaller struct{}

func (u *SurfaceContaminationLayerPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SurfaceContaminationLayerPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *SurfaceContaminationLayerPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SurfaceContaminationLayerPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "SurfaceContaminationLayer":
		GongUnmarshallPointer(&instance.SurfaceContaminationLayer, valueExpr, identifierMap)
	}
	return nil
}

type SurfaceContaminationLayerTypeUnmarshaller struct{}

func (u *SurfaceContaminationLayerTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SurfaceContaminationLayerType), stage, identifier, instanceName, preserveOrder)
}

func (u *SurfaceContaminationLayerTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SurfaceContaminationLayerType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Extent":
		GongUnmarshallSliceOfPointers(&instance.Extent, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type SurfaceTypeUnmarshaller struct{}

func (u *SurfaceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SurfaceType), stage, identifier, instanceName, preserveOrder)
}

func (u *SurfaceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SurfaceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "SrsName":
		instance.SrsName = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	}
	return nil
}

type SurveyControlPointTimeSlicePropertyTypeUnmarshaller struct{}

func (u *SurveyControlPointTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SurveyControlPointTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *SurveyControlPointTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SurveyControlPointTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "SurveyControlPointTimeSlice":
		GongUnmarshallPointer(&instance.SurveyControlPointTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type SurveyControlPointTimeSliceTypeUnmarshaller struct{}

func (u *SurveyControlPointTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SurveyControlPointTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *SurveyControlPointTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SurveyControlPointTimeSliceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Interpretation":
		instance.Interpretation = GongExtractString(valueExpr)
	case "SequenceNumber":
		instance.SequenceNumber = GongExtractInt(valueExpr)
	case "CorrectionNumber":
		instance.CorrectionNumber = GongExtractInt(valueExpr)
	case "AssociatedAirportHeliport":
		GongUnmarshallPointer(&instance.AssociatedAirportHeliport, valueExpr, identifierMap)
	case "Location":
		GongUnmarshallPointer(&instance.Location, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type SurveyControlPointTypeUnmarshaller struct{}

func (u *SurveyControlPointTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SurveyControlPointType), stage, identifier, instanceName, preserveOrder)
}

func (u *SurveyControlPointTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SurveyControlPointType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeSlice":
		GongUnmarshallSliceOfPointers(&instance.TimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type TaxiHoldingPositionLightSystemTimeSlicePropertyTypeUnmarshaller struct{}

func (u *TaxiHoldingPositionLightSystemTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TaxiHoldingPositionLightSystemTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *TaxiHoldingPositionLightSystemTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TaxiHoldingPositionLightSystemTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "TaxiHoldingPositionLightSystemTimeSlice":
		GongUnmarshallPointer(&instance.TaxiHoldingPositionLightSystemTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type TaxiHoldingPositionLightSystemTimeSliceTypeUnmarshaller struct{}

func (u *TaxiHoldingPositionLightSystemTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TaxiHoldingPositionLightSystemTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *TaxiHoldingPositionLightSystemTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TaxiHoldingPositionLightSystemTimeSliceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Interpretation":
		instance.Interpretation = GongExtractString(valueExpr)
	case "SequenceNumber":
		instance.SequenceNumber = GongExtractInt(valueExpr)
	case "CorrectionNumber":
		instance.CorrectionNumber = GongExtractInt(valueExpr)
	case "Element":
		GongUnmarshallSliceOfPointers(&instance.Element, valueExpr, identifierMap)
	case "Availability":
		GongUnmarshallSliceOfPointers(&instance.Availability, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "TaxiHolding":
		GongUnmarshallPointer(&instance.TaxiHolding, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type TaxiHoldingPositionLightSystemTypeUnmarshaller struct{}

func (u *TaxiHoldingPositionLightSystemTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TaxiHoldingPositionLightSystemType), stage, identifier, instanceName, preserveOrder)
}

func (u *TaxiHoldingPositionLightSystemTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TaxiHoldingPositionLightSystemType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeSlice":
		GongUnmarshallSliceOfPointers(&instance.TimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type TaxiHoldingPositionMarkingTimeSlicePropertyTypeUnmarshaller struct{}

func (u *TaxiHoldingPositionMarkingTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TaxiHoldingPositionMarkingTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *TaxiHoldingPositionMarkingTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TaxiHoldingPositionMarkingTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "TaxiHoldingPositionMarkingTimeSlice":
		GongUnmarshallPointer(&instance.TaxiHoldingPositionMarkingTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type TaxiHoldingPositionMarkingTimeSliceTypeUnmarshaller struct{}

func (u *TaxiHoldingPositionMarkingTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TaxiHoldingPositionMarkingTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *TaxiHoldingPositionMarkingTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TaxiHoldingPositionMarkingTimeSliceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Interpretation":
		instance.Interpretation = GongExtractString(valueExpr)
	case "SequenceNumber":
		instance.SequenceNumber = GongExtractInt(valueExpr)
	case "CorrectionNumber":
		instance.CorrectionNumber = GongExtractInt(valueExpr)
	case "Element":
		GongUnmarshallSliceOfPointers(&instance.Element, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "MarkedTaxiHold":
		GongUnmarshallPointer(&instance.MarkedTaxiHold, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type TaxiHoldingPositionMarkingTypeUnmarshaller struct{}

func (u *TaxiHoldingPositionMarkingTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TaxiHoldingPositionMarkingType), stage, identifier, instanceName, preserveOrder)
}

func (u *TaxiHoldingPositionMarkingTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TaxiHoldingPositionMarkingType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeSlice":
		GongUnmarshallSliceOfPointers(&instance.TimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type TaxiHoldingPositionPropertyTypeUnmarshaller struct{}

func (u *TaxiHoldingPositionPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TaxiHoldingPositionPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *TaxiHoldingPositionPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TaxiHoldingPositionPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "Type_XSD":
		instance.Type_XSD = GongExtractString(valueExpr)
	case "Href":
		instance.Href = GongExtractString(valueExpr)
	case "Role":
		instance.Role = GongExtractString(valueExpr)
	case "Arcrole":
		instance.Arcrole = GongExtractString(valueExpr)
	case "Title":
		instance.Title = GongExtractString(valueExpr)
	case "Show":
		instance.Show = GongExtractString(valueExpr)
	case "Actuate":
		instance.Actuate = GongExtractString(valueExpr)
	}
	return nil
}

type TaxiHoldingPositionTimeSlicePropertyTypeUnmarshaller struct{}

func (u *TaxiHoldingPositionTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TaxiHoldingPositionTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *TaxiHoldingPositionTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TaxiHoldingPositionTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "TaxiHoldingPositionTimeSlice":
		GongUnmarshallPointer(&instance.TaxiHoldingPositionTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type TaxiHoldingPositionTimeSliceTypeUnmarshaller struct{}

func (u *TaxiHoldingPositionTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TaxiHoldingPositionTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *TaxiHoldingPositionTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TaxiHoldingPositionTimeSliceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Interpretation":
		instance.Interpretation = GongExtractString(valueExpr)
	case "SequenceNumber":
		instance.SequenceNumber = GongExtractInt(valueExpr)
	case "CorrectionNumber":
		instance.CorrectionNumber = GongExtractInt(valueExpr)
	case "AssociatedGuidanceLine":
		GongUnmarshallPointer(&instance.AssociatedGuidanceLine, valueExpr, identifierMap)
	case "ProtectedRunway":
		GongUnmarshallSliceOfPointers(&instance.ProtectedRunway, valueExpr, identifierMap)
	case "Location":
		GongUnmarshallPointer(&instance.Location, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type TaxiHoldingPositionTypeUnmarshaller struct{}

func (u *TaxiHoldingPositionTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TaxiHoldingPositionType), stage, identifier, instanceName, preserveOrder)
}

func (u *TaxiHoldingPositionTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TaxiHoldingPositionType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeSlice":
		GongUnmarshallSliceOfPointers(&instance.TimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type TaxiwayContaminationPropertyTypeUnmarshaller struct{}

func (u *TaxiwayContaminationPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TaxiwayContaminationPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *TaxiwayContaminationPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TaxiwayContaminationPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "TaxiwayContamination":
		GongUnmarshallPointer(&instance.TaxiwayContamination, valueExpr, identifierMap)
	}
	return nil
}

type TaxiwayContaminationTypeUnmarshaller struct{}

func (u *TaxiwayContaminationTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TaxiwayContaminationType), stage, identifier, instanceName, preserveOrder)
}

func (u *TaxiwayContaminationTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TaxiwayContaminationType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "CriticalRidge":
		GongUnmarshallSliceOfPointers(&instance.CriticalRidge, valueExpr, identifierMap)
	case "Layer":
		GongUnmarshallSliceOfPointers(&instance.Layer, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type TaxiwayElementPropertyTypeUnmarshaller struct{}

func (u *TaxiwayElementPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TaxiwayElementPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *TaxiwayElementPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TaxiwayElementPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "Type_XSD":
		instance.Type_XSD = GongExtractString(valueExpr)
	case "Href":
		instance.Href = GongExtractString(valueExpr)
	case "Role":
		instance.Role = GongExtractString(valueExpr)
	case "Arcrole":
		instance.Arcrole = GongExtractString(valueExpr)
	case "Title":
		instance.Title = GongExtractString(valueExpr)
	case "Show":
		instance.Show = GongExtractString(valueExpr)
	case "Actuate":
		instance.Actuate = GongExtractString(valueExpr)
	}
	return nil
}

type TaxiwayElementTimeSlicePropertyTypeUnmarshaller struct{}

func (u *TaxiwayElementTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TaxiwayElementTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *TaxiwayElementTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TaxiwayElementTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "TaxiwayElementTimeSlice":
		GongUnmarshallPointer(&instance.TaxiwayElementTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type TaxiwayElementTimeSliceTypeUnmarshaller struct{}

func (u *TaxiwayElementTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TaxiwayElementTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *TaxiwayElementTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TaxiwayElementTimeSliceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Interpretation":
		instance.Interpretation = GongExtractString(valueExpr)
	case "SequenceNumber":
		instance.SequenceNumber = GongExtractInt(valueExpr)
	case "CorrectionNumber":
		instance.CorrectionNumber = GongExtractInt(valueExpr)
	case "SurfaceProperties":
		GongUnmarshallPointer(&instance.SurfaceProperties, valueExpr, identifierMap)
	case "AssociatedTaxiway":
		GongUnmarshallPointer(&instance.AssociatedTaxiway, valueExpr, identifierMap)
	case "Extent":
		GongUnmarshallPointer(&instance.Extent, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Availability":
		GongUnmarshallSliceOfPointers(&instance.Availability, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type TaxiwayElementTypeUnmarshaller struct{}

func (u *TaxiwayElementTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TaxiwayElementType), stage, identifier, instanceName, preserveOrder)
}

func (u *TaxiwayElementTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TaxiwayElementType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeSlice":
		GongUnmarshallSliceOfPointers(&instance.TimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type TaxiwayLightSystemTimeSlicePropertyTypeUnmarshaller struct{}

func (u *TaxiwayLightSystemTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TaxiwayLightSystemTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *TaxiwayLightSystemTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TaxiwayLightSystemTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "TaxiwayLightSystemTimeSlice":
		GongUnmarshallPointer(&instance.TaxiwayLightSystemTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type TaxiwayLightSystemTimeSliceTypeUnmarshaller struct{}

func (u *TaxiwayLightSystemTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TaxiwayLightSystemTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *TaxiwayLightSystemTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TaxiwayLightSystemTimeSliceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Interpretation":
		instance.Interpretation = GongExtractString(valueExpr)
	case "SequenceNumber":
		instance.SequenceNumber = GongExtractInt(valueExpr)
	case "CorrectionNumber":
		instance.CorrectionNumber = GongExtractInt(valueExpr)
	case "Element":
		GongUnmarshallSliceOfPointers(&instance.Element, valueExpr, identifierMap)
	case "Availability":
		GongUnmarshallSliceOfPointers(&instance.Availability, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "LightedTaxiway":
		GongUnmarshallPointer(&instance.LightedTaxiway, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type TaxiwayLightSystemTypeUnmarshaller struct{}

func (u *TaxiwayLightSystemTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TaxiwayLightSystemType), stage, identifier, instanceName, preserveOrder)
}

func (u *TaxiwayLightSystemTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TaxiwayLightSystemType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeSlice":
		GongUnmarshallSliceOfPointers(&instance.TimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type TaxiwayMarkingTimeSlicePropertyTypeUnmarshaller struct{}

func (u *TaxiwayMarkingTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TaxiwayMarkingTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *TaxiwayMarkingTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TaxiwayMarkingTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "TaxiwayMarkingTimeSlice":
		GongUnmarshallPointer(&instance.TaxiwayMarkingTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type TaxiwayMarkingTimeSliceTypeUnmarshaller struct{}

func (u *TaxiwayMarkingTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TaxiwayMarkingTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *TaxiwayMarkingTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TaxiwayMarkingTimeSliceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Interpretation":
		instance.Interpretation = GongExtractString(valueExpr)
	case "SequenceNumber":
		instance.SequenceNumber = GongExtractInt(valueExpr)
	case "CorrectionNumber":
		instance.CorrectionNumber = GongExtractInt(valueExpr)
	case "Element":
		GongUnmarshallSliceOfPointers(&instance.Element, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "MarkedTaxiway":
		GongUnmarshallPointer(&instance.MarkedTaxiway, valueExpr, identifierMap)
	case "MarkedElement":
		GongUnmarshallPointer(&instance.MarkedElement, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type TaxiwayMarkingTypeUnmarshaller struct{}

func (u *TaxiwayMarkingTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TaxiwayMarkingType), stage, identifier, instanceName, preserveOrder)
}

func (u *TaxiwayMarkingTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TaxiwayMarkingType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeSlice":
		GongUnmarshallSliceOfPointers(&instance.TimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type TaxiwayPropertyTypeUnmarshaller struct{}

func (u *TaxiwayPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TaxiwayPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *TaxiwayPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TaxiwayPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "Type_XSD":
		instance.Type_XSD = GongExtractString(valueExpr)
	case "Href":
		instance.Href = GongExtractString(valueExpr)
	case "Role":
		instance.Role = GongExtractString(valueExpr)
	case "Arcrole":
		instance.Arcrole = GongExtractString(valueExpr)
	case "Title":
		instance.Title = GongExtractString(valueExpr)
	case "Show":
		instance.Show = GongExtractString(valueExpr)
	case "Actuate":
		instance.Actuate = GongExtractString(valueExpr)
	}
	return nil
}

type TaxiwayTimeSlicePropertyTypeUnmarshaller struct{}

func (u *TaxiwayTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TaxiwayTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *TaxiwayTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TaxiwayTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "TaxiwayTimeSlice":
		GongUnmarshallPointer(&instance.TaxiwayTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type TaxiwayTimeSliceTypeUnmarshaller struct{}

func (u *TaxiwayTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TaxiwayTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *TaxiwayTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TaxiwayTimeSliceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Interpretation":
		instance.Interpretation = GongExtractString(valueExpr)
	case "SequenceNumber":
		instance.SequenceNumber = GongExtractInt(valueExpr)
	case "CorrectionNumber":
		instance.CorrectionNumber = GongExtractInt(valueExpr)
	case "SurfaceProperties":
		GongUnmarshallPointer(&instance.SurfaceProperties, valueExpr, identifierMap)
	case "AssociatedAirportHeliport":
		GongUnmarshallPointer(&instance.AssociatedAirportHeliport, valueExpr, identifierMap)
	case "Contaminant":
		GongUnmarshallSliceOfPointers(&instance.Contaminant, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Availability":
		GongUnmarshallSliceOfPointers(&instance.Availability, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type TaxiwayTypeUnmarshaller struct{}

func (u *TaxiwayTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TaxiwayType), stage, identifier, instanceName, preserveOrder)
}

func (u *TaxiwayTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TaxiwayType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeSlice":
		GongUnmarshallSliceOfPointers(&instance.TimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type TelephoneContactPropertyTypeUnmarshaller struct{}

func (u *TelephoneContactPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TelephoneContactPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *TelephoneContactPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TelephoneContactPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "TelephoneContact":
		GongUnmarshallPointer(&instance.TelephoneContact, valueExpr, identifierMap)
	}
	return nil
}

type TelephoneContactTypeUnmarshaller struct{}

func (u *TelephoneContactTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TelephoneContactType), stage, identifier, instanceName, preserveOrder)
}

func (u *TelephoneContactTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TelephoneContactType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeInterval":
		GongUnmarshallSliceOfPointers(&instance.TimeInterval, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "SpecialDateAuthority":
		GongUnmarshallSliceOfPointers(&instance.SpecialDateAuthority, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type TimesheetPropertyTypeUnmarshaller struct{}

func (u *TimesheetPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TimesheetPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *TimesheetPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TimesheetPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "Timesheet":
		GongUnmarshallPointer(&instance.Timesheet, valueExpr, identifierMap)
	}
	return nil
}

type TimesheetTypeUnmarshaller struct{}

func (u *TimesheetTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TimesheetType), stage, identifier, instanceName, preserveOrder)
}

func (u *TimesheetTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TimesheetType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type TouchDownLiftOffContaminationPropertyTypeUnmarshaller struct{}

func (u *TouchDownLiftOffContaminationPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TouchDownLiftOffContaminationPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *TouchDownLiftOffContaminationPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TouchDownLiftOffContaminationPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "TouchDownLiftOffContamination":
		GongUnmarshallPointer(&instance.TouchDownLiftOffContamination, valueExpr, identifierMap)
	}
	return nil
}

type TouchDownLiftOffContaminationTypeUnmarshaller struct{}

func (u *TouchDownLiftOffContaminationTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TouchDownLiftOffContaminationType), stage, identifier, instanceName, preserveOrder)
}

func (u *TouchDownLiftOffContaminationTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TouchDownLiftOffContaminationType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "CriticalRidge":
		GongUnmarshallSliceOfPointers(&instance.CriticalRidge, valueExpr, identifierMap)
	case "Layer":
		GongUnmarshallSliceOfPointers(&instance.Layer, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type TouchDownLiftOffLightSystemTimeSlicePropertyTypeUnmarshaller struct{}

func (u *TouchDownLiftOffLightSystemTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TouchDownLiftOffLightSystemTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *TouchDownLiftOffLightSystemTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TouchDownLiftOffLightSystemTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "TouchDownLiftOffLightSystemTimeSlice":
		GongUnmarshallPointer(&instance.TouchDownLiftOffLightSystemTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type TouchDownLiftOffLightSystemTimeSliceTypeUnmarshaller struct{}

func (u *TouchDownLiftOffLightSystemTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TouchDownLiftOffLightSystemTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *TouchDownLiftOffLightSystemTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TouchDownLiftOffLightSystemTimeSliceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Interpretation":
		instance.Interpretation = GongExtractString(valueExpr)
	case "SequenceNumber":
		instance.SequenceNumber = GongExtractInt(valueExpr)
	case "CorrectionNumber":
		instance.CorrectionNumber = GongExtractInt(valueExpr)
	case "Element":
		GongUnmarshallSliceOfPointers(&instance.Element, valueExpr, identifierMap)
	case "Availability":
		GongUnmarshallSliceOfPointers(&instance.Availability, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "LightedTouchDownLiftOff":
		GongUnmarshallPointer(&instance.LightedTouchDownLiftOff, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type TouchDownLiftOffLightSystemTypeUnmarshaller struct{}

func (u *TouchDownLiftOffLightSystemTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TouchDownLiftOffLightSystemType), stage, identifier, instanceName, preserveOrder)
}

func (u *TouchDownLiftOffLightSystemTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TouchDownLiftOffLightSystemType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeSlice":
		GongUnmarshallSliceOfPointers(&instance.TimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type TouchDownLiftOffMarkingTimeSlicePropertyTypeUnmarshaller struct{}

func (u *TouchDownLiftOffMarkingTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TouchDownLiftOffMarkingTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *TouchDownLiftOffMarkingTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TouchDownLiftOffMarkingTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "TouchDownLiftOffMarkingTimeSlice":
		GongUnmarshallPointer(&instance.TouchDownLiftOffMarkingTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type TouchDownLiftOffMarkingTimeSliceTypeUnmarshaller struct{}

func (u *TouchDownLiftOffMarkingTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TouchDownLiftOffMarkingTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *TouchDownLiftOffMarkingTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TouchDownLiftOffMarkingTimeSliceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Interpretation":
		instance.Interpretation = GongExtractString(valueExpr)
	case "SequenceNumber":
		instance.SequenceNumber = GongExtractInt(valueExpr)
	case "CorrectionNumber":
		instance.CorrectionNumber = GongExtractInt(valueExpr)
	case "Element":
		GongUnmarshallSliceOfPointers(&instance.Element, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "MarkedTouchDownLiftOff":
		GongUnmarshallPointer(&instance.MarkedTouchDownLiftOff, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type TouchDownLiftOffMarkingTypeUnmarshaller struct{}

func (u *TouchDownLiftOffMarkingTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TouchDownLiftOffMarkingType), stage, identifier, instanceName, preserveOrder)
}

func (u *TouchDownLiftOffMarkingTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TouchDownLiftOffMarkingType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeSlice":
		GongUnmarshallSliceOfPointers(&instance.TimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type TouchDownLiftOffPropertyTypeUnmarshaller struct{}

func (u *TouchDownLiftOffPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TouchDownLiftOffPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *TouchDownLiftOffPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TouchDownLiftOffPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "Type_XSD":
		instance.Type_XSD = GongExtractString(valueExpr)
	case "Href":
		instance.Href = GongExtractString(valueExpr)
	case "Role":
		instance.Role = GongExtractString(valueExpr)
	case "Arcrole":
		instance.Arcrole = GongExtractString(valueExpr)
	case "Title":
		instance.Title = GongExtractString(valueExpr)
	case "Show":
		instance.Show = GongExtractString(valueExpr)
	case "Actuate":
		instance.Actuate = GongExtractString(valueExpr)
	}
	return nil
}

type TouchDownLiftOffSafeAreaTimeSlicePropertyTypeUnmarshaller struct{}

func (u *TouchDownLiftOffSafeAreaTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TouchDownLiftOffSafeAreaTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *TouchDownLiftOffSafeAreaTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TouchDownLiftOffSafeAreaTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "TouchDownLiftOffSafeAreaTimeSlice":
		GongUnmarshallPointer(&instance.TouchDownLiftOffSafeAreaTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type TouchDownLiftOffSafeAreaTimeSliceTypeUnmarshaller struct{}

func (u *TouchDownLiftOffSafeAreaTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TouchDownLiftOffSafeAreaTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *TouchDownLiftOffSafeAreaTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TouchDownLiftOffSafeAreaTimeSliceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Interpretation":
		instance.Interpretation = GongExtractString(valueExpr)
	case "SequenceNumber":
		instance.SequenceNumber = GongExtractInt(valueExpr)
	case "CorrectionNumber":
		instance.CorrectionNumber = GongExtractInt(valueExpr)
	case "SurfaceProperties":
		GongUnmarshallPointer(&instance.SurfaceProperties, valueExpr, identifierMap)
	case "Extent":
		GongUnmarshallPointer(&instance.Extent, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "ProtectedTouchDownLiftOff":
		GongUnmarshallPointer(&instance.ProtectedTouchDownLiftOff, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type TouchDownLiftOffSafeAreaTypeUnmarshaller struct{}

func (u *TouchDownLiftOffSafeAreaTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TouchDownLiftOffSafeAreaType), stage, identifier, instanceName, preserveOrder)
}

func (u *TouchDownLiftOffSafeAreaTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TouchDownLiftOffSafeAreaType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeSlice":
		GongUnmarshallSliceOfPointers(&instance.TimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type TouchDownLiftOffTimeSlicePropertyTypeUnmarshaller struct{}

func (u *TouchDownLiftOffTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TouchDownLiftOffTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *TouchDownLiftOffTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TouchDownLiftOffTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "TouchDownLiftOffTimeSlice":
		GongUnmarshallPointer(&instance.TouchDownLiftOffTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type TouchDownLiftOffTimeSliceTypeUnmarshaller struct{}

func (u *TouchDownLiftOffTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TouchDownLiftOffTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *TouchDownLiftOffTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TouchDownLiftOffTimeSliceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Interpretation":
		instance.Interpretation = GongExtractString(valueExpr)
	case "SequenceNumber":
		instance.SequenceNumber = GongExtractInt(valueExpr)
	case "CorrectionNumber":
		instance.CorrectionNumber = GongExtractInt(valueExpr)
	case "AimingPoint":
		GongUnmarshallPointer(&instance.AimingPoint, valueExpr, identifierMap)
	case "Extent":
		GongUnmarshallPointer(&instance.Extent, valueExpr, identifierMap)
	case "SurfaceProperties":
		GongUnmarshallPointer(&instance.SurfaceProperties, valueExpr, identifierMap)
	case "AssociatedAirportHeliport":
		GongUnmarshallPointer(&instance.AssociatedAirportHeliport, valueExpr, identifierMap)
	case "ApproachTakeOffArea":
		GongUnmarshallPointer(&instance.ApproachTakeOffArea, valueExpr, identifierMap)
	case "Contaminant":
		GongUnmarshallSliceOfPointers(&instance.Contaminant, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Availability":
		GongUnmarshallSliceOfPointers(&instance.Availability, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type TouchDownLiftOffTypeUnmarshaller struct{}

func (u *TouchDownLiftOffTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TouchDownLiftOffType), stage, identifier, instanceName, preserveOrder)
}

func (u *TouchDownLiftOffTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TouchDownLiftOffType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeSlice":
		GongUnmarshallSliceOfPointers(&instance.TimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type VisualGlideSlopeIndicatorTimeSlicePropertyTypeUnmarshaller struct{}

func (u *VisualGlideSlopeIndicatorTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(VisualGlideSlopeIndicatorTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *VisualGlideSlopeIndicatorTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*VisualGlideSlopeIndicatorTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "VisualGlideSlopeIndicatorTimeSlice":
		GongUnmarshallPointer(&instance.VisualGlideSlopeIndicatorTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type VisualGlideSlopeIndicatorTimeSliceTypeUnmarshaller struct{}

func (u *VisualGlideSlopeIndicatorTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(VisualGlideSlopeIndicatorTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *VisualGlideSlopeIndicatorTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*VisualGlideSlopeIndicatorTimeSliceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Interpretation":
		instance.Interpretation = GongExtractString(valueExpr)
	case "SequenceNumber":
		instance.SequenceNumber = GongExtractInt(valueExpr)
	case "CorrectionNumber":
		instance.CorrectionNumber = GongExtractInt(valueExpr)
	case "Element":
		GongUnmarshallSliceOfPointers(&instance.Element, valueExpr, identifierMap)
	case "Availability":
		GongUnmarshallSliceOfPointers(&instance.Availability, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "RunwayDirection":
		GongUnmarshallPointer(&instance.RunwayDirection, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type VisualGlideSlopeIndicatorTypeUnmarshaller struct{}

func (u *VisualGlideSlopeIndicatorTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(VisualGlideSlopeIndicatorType), stage, identifier, instanceName, preserveOrder)
}

func (u *VisualGlideSlopeIndicatorTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*VisualGlideSlopeIndicatorType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeSlice":
		GongUnmarshallSliceOfPointers(&instance.TimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type WorkAreaTimeSlicePropertyTypeUnmarshaller struct{}

func (u *WorkAreaTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(WorkAreaTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *WorkAreaTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*WorkAreaTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "WorkAreaTimeSlice":
		GongUnmarshallPointer(&instance.WorkAreaTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type WorkAreaTimeSliceTypeUnmarshaller struct{}

func (u *WorkAreaTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(WorkAreaTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *WorkAreaTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*WorkAreaTimeSliceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "Interpretation":
		instance.Interpretation = GongExtractString(valueExpr)
	case "SequenceNumber":
		instance.SequenceNumber = GongExtractInt(valueExpr)
	case "CorrectionNumber":
		instance.CorrectionNumber = GongExtractInt(valueExpr)
	case "AssociatedAirportHeliport":
		GongUnmarshallPointer(&instance.AssociatedAirportHeliport, valueExpr, identifierMap)
	case "Extent":
		GongUnmarshallPointer(&instance.Extent, valueExpr, identifierMap)
	case "Activation":
		GongUnmarshallSliceOfPointers(&instance.Activation, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type WorkAreaTypeUnmarshaller struct{}

func (u *WorkAreaTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(WorkAreaType), stage, identifier, instanceName, preserveOrder)
}

func (u *WorkAreaTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*WorkAreaType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeSlice":
		GongUnmarshallSliceOfPointers(&instance.TimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type WorkareaActivityPropertyTypeUnmarshaller struct{}

func (u *WorkareaActivityPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(WorkareaActivityPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *WorkareaActivityPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*WorkareaActivityPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "WorkareaActivity":
		GongUnmarshallPointer(&instance.WorkareaActivity, valueExpr, identifierMap)
	}
	return nil
}

type WorkareaActivityTypeUnmarshaller struct{}

func (u *WorkareaActivityTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(WorkareaActivityType), stage, identifier, instanceName, preserveOrder)
}

func (u *WorkareaActivityTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*WorkareaActivityType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TimeInterval":
		GongUnmarshallSliceOfPointers(&instance.TimeInterval, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "SpecialDateAuthority":
		GongUnmarshallSliceOfPointers(&instance.SpecialDateAuthority, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}
