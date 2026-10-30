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
		case "github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go/models/aixmfeatures":
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
				case "A_extension_5":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_5), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_6":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_6), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_7":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_7), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_8":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_8), stage, ident.Name, instanceName, preserveOrder)
				case "A_extension_9":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_extension_9), stage, ident.Name, instanceName, preserveOrder)
				case "AbstractNavaidEquipmentType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AbstractNavaidEquipmentType), stage, ident.Name, instanceName, preserveOrder)
				case "AbstractNavigationSystemCheckpointType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AbstractNavigationSystemCheckpointType), stage, ident.Name, instanceName, preserveOrder)
				case "AbstractPropertiesWithScheduleType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AbstractPropertiesWithScheduleType), stage, ident.Name, instanceName, preserveOrder)
				case "AbstractRadarEquipmentType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AbstractRadarEquipmentType), stage, ident.Name, instanceName, preserveOrder)
				case "AbstractSurveillanceRadarType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AbstractSurveillanceRadarType), stage, ident.Name, instanceName, preserveOrder)
				case "AirportHeliportPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AirportHeliportPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "AirspacePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AirspacePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "AngleIndicationTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AngleIndicationTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "AngleIndicationTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AngleIndicationTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "AngleIndicationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AngleIndicationType), stage, ident.Name, instanceName, preserveOrder)
				case "AuthorityForNavaidEquipmentPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AuthorityForNavaidEquipmentPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "AuthorityForNavaidEquipmentType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AuthorityForNavaidEquipmentType), stage, ident.Name, instanceName, preserveOrder)
				case "AuthorityForSpecialNavigationStationPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AuthorityForSpecialNavigationStationPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "AuthorityForSpecialNavigationStationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AuthorityForSpecialNavigationStationType), stage, ident.Name, instanceName, preserveOrder)
				case "AuthorityForSpecialNavigationSystemPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AuthorityForSpecialNavigationSystemPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "AuthorityForSpecialNavigationSystemType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AuthorityForSpecialNavigationSystemType), stage, ident.Name, instanceName, preserveOrder)
				case "AzimuthTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AzimuthTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "AzimuthTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AzimuthTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "AzimuthType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AzimuthType), stage, ident.Name, instanceName, preserveOrder)
				case "CheckpointINSTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CheckpointINSTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "CheckpointINSTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CheckpointINSTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "CheckpointINSType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CheckpointINSType), stage, ident.Name, instanceName, preserveOrder)
				case "CheckpointVORTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CheckpointVORTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "CheckpointVORTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CheckpointVORTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "CheckpointVORType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CheckpointVORType), stage, ident.Name, instanceName, preserveOrder)
				case "ContactInformationPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ContactInformationPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "ContactInformationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ContactInformationType), stage, ident.Name, instanceName, preserveOrder)
				case "DMETimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DMETimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "DMETimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DMETimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "DMEType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DMEType), stage, ident.Name, instanceName, preserveOrder)
				case "DesignatedPointPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DesignatedPointPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "DesignatedPointTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DesignatedPointTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "DesignatedPointTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DesignatedPointTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "DesignatedPointType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DesignatedPointType), stage, ident.Name, instanceName, preserveOrder)
				case "DirectionFinderTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DirectionFinderTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "DirectionFinderTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DirectionFinderTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "DirectionFinderType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DirectionFinderType), stage, ident.Name, instanceName, preserveOrder)
				case "DistanceIndicationTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DistanceIndicationTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "DistanceIndicationTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DistanceIndicationTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "DistanceIndicationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DistanceIndicationType), stage, ident.Name, instanceName, preserveOrder)
				case "ElevatedPointPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ElevatedPointPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "ElevatedPointType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ElevatedPointType), stage, ident.Name, instanceName, preserveOrder)
				case "ElevationTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ElevationTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "ElevationTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ElevationTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "ElevationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ElevationType), stage, ident.Name, instanceName, preserveOrder)
				case "GlidepathTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(GlidepathTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "GlidepathTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(GlidepathTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "GlidepathType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(GlidepathType), stage, ident.Name, instanceName, preserveOrder)
				case "InformationServicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(InformationServicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "LinguisticNotePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(LinguisticNotePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "LinguisticNoteType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(LinguisticNoteType), stage, ident.Name, instanceName, preserveOrder)
				case "LocalizerTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(LocalizerTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "LocalizerTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(LocalizerTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "LocalizerType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(LocalizerType), stage, ident.Name, instanceName, preserveOrder)
				case "MarkerBeaconTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(MarkerBeaconTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "MarkerBeaconTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(MarkerBeaconTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "MarkerBeaconType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(MarkerBeaconType), stage, ident.Name, instanceName, preserveOrder)
				case "NDBTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NDBTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "NDBTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NDBTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "NDBType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NDBType), stage, ident.Name, instanceName, preserveOrder)
				case "NavaidComponentPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NavaidComponentPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "NavaidComponentType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NavaidComponentType), stage, ident.Name, instanceName, preserveOrder)
				case "NavaidEquipmentMonitoringPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NavaidEquipmentMonitoringPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "NavaidEquipmentMonitoringType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NavaidEquipmentMonitoringType), stage, ident.Name, instanceName, preserveOrder)
				case "NavaidEquipmentPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NavaidEquipmentPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "NavaidOperationalStatusPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NavaidOperationalStatusPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "NavaidOperationalStatusType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NavaidOperationalStatusType), stage, ident.Name, instanceName, preserveOrder)
				case "NavaidPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NavaidPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "NavaidTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NavaidTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "NavaidTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NavaidTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "NavaidType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NavaidType), stage, ident.Name, instanceName, preserveOrder)
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
				case "PointPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(PointPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "PointType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(PointType), stage, ident.Name, instanceName, preserveOrder)
				case "PostalAddressPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(PostalAddressPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "PostalAddressType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(PostalAddressType), stage, ident.Name, instanceName, preserveOrder)
				case "PrecisionApproachRadarTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(PrecisionApproachRadarTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "PrecisionApproachRadarTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(PrecisionApproachRadarTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "PrecisionApproachRadarType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(PrecisionApproachRadarType), stage, ident.Name, instanceName, preserveOrder)
				case "PrimarySurveillanceRadarTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(PrimarySurveillanceRadarTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "PrimarySurveillanceRadarTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(PrimarySurveillanceRadarTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "PrimarySurveillanceRadarType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(PrimarySurveillanceRadarType), stage, ident.Name, instanceName, preserveOrder)
				case "RadarComponentPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RadarComponentPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "RadarComponentType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RadarComponentType), stage, ident.Name, instanceName, preserveOrder)
				case "RadarEquipmentPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RadarEquipmentPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "RadarSystemTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RadarSystemTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "RadarSystemTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RadarSystemTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "RadarSystemType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RadarSystemType), stage, ident.Name, instanceName, preserveOrder)
				case "ReflectorPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ReflectorPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "ReflectorType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ReflectorType), stage, ident.Name, instanceName, preserveOrder)
				case "RunwayCentrelinePointPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RunwayCentrelinePointPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "RunwayDirectionPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RunwayDirectionPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "RunwayPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RunwayPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "SDFTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SDFTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "SDFTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SDFTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "SDFType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SDFType), stage, ident.Name, instanceName, preserveOrder)
				case "SecondarySurveillanceRadarTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SecondarySurveillanceRadarTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "SecondarySurveillanceRadarTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SecondarySurveillanceRadarTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "SecondarySurveillanceRadarType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SecondarySurveillanceRadarType), stage, ident.Name, instanceName, preserveOrder)
				case "SignificantPointInAirspaceTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SignificantPointInAirspaceTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "SignificantPointInAirspaceTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SignificantPointInAirspaceTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "SignificantPointInAirspaceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SignificantPointInAirspaceType), stage, ident.Name, instanceName, preserveOrder)
				case "SpecialNavigationStationStatusPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SpecialNavigationStationStatusPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "SpecialNavigationStationStatusType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SpecialNavigationStationStatusType), stage, ident.Name, instanceName, preserveOrder)
				case "SpecialNavigationStationTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SpecialNavigationStationTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "SpecialNavigationStationTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SpecialNavigationStationTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "SpecialNavigationStationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SpecialNavigationStationType), stage, ident.Name, instanceName, preserveOrder)
				case "SpecialNavigationSystemPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SpecialNavigationSystemPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "SpecialNavigationSystemTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SpecialNavigationSystemTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "SpecialNavigationSystemTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SpecialNavigationSystemTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "SpecialNavigationSystemType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SpecialNavigationSystemType), stage, ident.Name, instanceName, preserveOrder)
				case "SurveillanceGroundStationPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SurveillanceGroundStationPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "SurveillanceGroundStationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SurveillanceGroundStationType), stage, ident.Name, instanceName, preserveOrder)
				case "TACANTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TACANTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "TACANTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TACANTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "TACANType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TACANType), stage, ident.Name, instanceName, preserveOrder)
				case "TelephoneContactPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TelephoneContactPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "TelephoneContactType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TelephoneContactType), stage, ident.Name, instanceName, preserveOrder)
				case "TimesheetPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TimesheetPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "TimesheetType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TimesheetType), stage, ident.Name, instanceName, preserveOrder)
				case "TouchDownLiftOffPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TouchDownLiftOffPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "UnitPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(UnitPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "VORPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(VORPropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "VORTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(VORTimeSlicePropertyType), stage, ident.Name, instanceName, preserveOrder)
				case "VORTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(VORTimeSliceType), stage, ident.Name, instanceName, preserveOrder)
				case "VORType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(VORType), stage, ident.Name, instanceName, preserveOrder)
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
				case *A_extension_5:
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
				case *A_extension_7:
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
				case *A_extension_9:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *AbstractNavaidEquipmentType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					}
				case *AbstractNavigationSystemCheckpointType:
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
				case *AbstractRadarEquipmentType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					}
				case *AbstractSurveillanceRadarType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
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
				case *AirspacePropertyType:
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
				case *AngleIndicationTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AngleIndicationTimeSlice":
						__gong__assignPointer(&inst.AngleIndicationTimeSlice, rhs, identifierMap)
					}
				case *AngleIndicationTimeSliceType:
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
					case "Fix":
						__gong__assignPointer(&inst.Fix, rhs, identifierMap)
					case "PointChoice_fixDesignatedPoint":
						__gong__assignPointer(&inst.PointChoice_fixDesignatedPoint, rhs, identifierMap)
					case "PointChoice_navaidSystem":
						__gong__assignPointer(&inst.PointChoice_navaidSystem, rhs, identifierMap)
					case "PointChoice_position":
						__gong__assignPointer(&inst.PointChoice_position, rhs, identifierMap)
					case "PointChoice_runwayPoint":
						__gong__assignPointer(&inst.PointChoice_runwayPoint, rhs, identifierMap)
					case "PointChoice_aimingPoint":
						__gong__assignPointer(&inst.PointChoice_aimingPoint, rhs, identifierMap)
					case "PointChoice_airportReferencePoint":
						__gong__assignPointer(&inst.PointChoice_airportReferencePoint, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *AngleIndicationType:
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
				case *AuthorityForNavaidEquipmentPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "AuthorityForNavaidEquipment":
						__gong__assignPointer(&inst.AuthorityForNavaidEquipment, rhs, identifierMap)
					}
				case *AuthorityForNavaidEquipmentType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "TheOrganisationAuthority":
						__gong__assignPointer(&inst.TheOrganisationAuthority, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *AuthorityForSpecialNavigationStationPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "AuthorityForSpecialNavigationStation":
						__gong__assignPointer(&inst.AuthorityForSpecialNavigationStation, rhs, identifierMap)
					}
				case *AuthorityForSpecialNavigationStationType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "TheOrganisationAuthority":
						__gong__assignPointer(&inst.TheOrganisationAuthority, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *AuthorityForSpecialNavigationSystemPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "AuthorityForSpecialNavigationSystem":
						__gong__assignPointer(&inst.AuthorityForSpecialNavigationSystem, rhs, identifierMap)
					}
				case *AuthorityForSpecialNavigationSystemType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "TheOrganisationAuthority":
						__gong__assignPointer(&inst.TheOrganisationAuthority, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *AzimuthTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AzimuthTimeSlice":
						__gong__assignPointer(&inst.AzimuthTimeSlice, rhs, identifierMap)
					}
				case *AzimuthTimeSliceType:
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
					case "Authority":
						__gong__assignSliceOfPointers(&inst.Authority, rhs, identifierMap)
					case "Monitoring":
						__gong__assignSliceOfPointers(&inst.Monitoring, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *AzimuthType:
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
				case *CheckpointINSTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "CheckpointINSTimeSlice":
						__gong__assignPointer(&inst.CheckpointINSTimeSlice, rhs, identifierMap)
					}
				case *CheckpointINSTimeSliceType:
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
					case "Position":
						__gong__assignPointer(&inst.Position, rhs, identifierMap)
					case "AirportHeliport":
						__gong__assignPointer(&inst.AirportHeliport, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *CheckpointINSType:
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
				case *CheckpointVORTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "CheckpointVORTimeSlice":
						__gong__assignPointer(&inst.CheckpointVORTimeSlice, rhs, identifierMap)
					}
				case *CheckpointVORTimeSliceType:
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
					case "Position":
						__gong__assignPointer(&inst.Position, rhs, identifierMap)
					case "AirportHeliport":
						__gong__assignPointer(&inst.AirportHeliport, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "CheckPointFacility":
						__gong__assignPointer(&inst.CheckPointFacility, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *CheckpointVORType:
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
				case *DMETimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "DMETimeSlice":
						__gong__assignPointer(&inst.DMETimeSlice, rhs, identifierMap)
					}
				case *DMETimeSliceType:
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
					case "Authority":
						__gong__assignSliceOfPointers(&inst.Authority, rhs, identifierMap)
					case "Monitoring":
						__gong__assignSliceOfPointers(&inst.Monitoring, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *DMEType:
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
				case *DesignatedPointPropertyType:
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
				case *DesignatedPointTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "DesignatedPointTimeSlice":
						__gong__assignPointer(&inst.DesignatedPointTimeSlice, rhs, identifierMap)
					}
				case *DesignatedPointTimeSliceType:
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
					case "AimingPoint":
						__gong__assignPointer(&inst.AimingPoint, rhs, identifierMap)
					case "AirportHeliport":
						__gong__assignPointer(&inst.AirportHeliport, rhs, identifierMap)
					case "RunwayPoint":
						__gong__assignPointer(&inst.RunwayPoint, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *DesignatedPointType:
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
				case *DirectionFinderTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "DirectionFinderTimeSlice":
						__gong__assignPointer(&inst.DirectionFinderTimeSlice, rhs, identifierMap)
					}
				case *DirectionFinderTimeSliceType:
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
					case "Authority":
						__gong__assignSliceOfPointers(&inst.Authority, rhs, identifierMap)
					case "Monitoring":
						__gong__assignSliceOfPointers(&inst.Monitoring, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "InformationProvision":
						__gong__assignSliceOfPointers(&inst.InformationProvision, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *DirectionFinderType:
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
				case *DistanceIndicationTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "DistanceIndicationTimeSlice":
						__gong__assignPointer(&inst.DistanceIndicationTimeSlice, rhs, identifierMap)
					}
				case *DistanceIndicationTimeSliceType:
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
					case "Fix":
						__gong__assignPointer(&inst.Fix, rhs, identifierMap)
					case "PointChoice_fixDesignatedPoint":
						__gong__assignPointer(&inst.PointChoice_fixDesignatedPoint, rhs, identifierMap)
					case "PointChoice_navaidSystem":
						__gong__assignPointer(&inst.PointChoice_navaidSystem, rhs, identifierMap)
					case "PointChoice_position":
						__gong__assignPointer(&inst.PointChoice_position, rhs, identifierMap)
					case "PointChoice_runwayPoint":
						__gong__assignPointer(&inst.PointChoice_runwayPoint, rhs, identifierMap)
					case "PointChoice_aimingPoint":
						__gong__assignPointer(&inst.PointChoice_aimingPoint, rhs, identifierMap)
					case "PointChoice_airportReferencePoint":
						__gong__assignPointer(&inst.PointChoice_airportReferencePoint, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *DistanceIndicationType:
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
				case *ElevationTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "ElevationTimeSlice":
						__gong__assignPointer(&inst.ElevationTimeSlice, rhs, identifierMap)
					}
				case *ElevationTimeSliceType:
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
					case "Authority":
						__gong__assignSliceOfPointers(&inst.Authority, rhs, identifierMap)
					case "Monitoring":
						__gong__assignSliceOfPointers(&inst.Monitoring, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *ElevationType:
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
				case *GlidepathTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "GlidepathTimeSlice":
						__gong__assignPointer(&inst.GlidepathTimeSlice, rhs, identifierMap)
					}
				case *GlidepathTimeSliceType:
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
					case "Authority":
						__gong__assignSliceOfPointers(&inst.Authority, rhs, identifierMap)
					case "Monitoring":
						__gong__assignSliceOfPointers(&inst.Monitoring, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *GlidepathType:
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
				case *InformationServicePropertyType:
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
				case *LocalizerTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "LocalizerTimeSlice":
						__gong__assignPointer(&inst.LocalizerTimeSlice, rhs, identifierMap)
					}
				case *LocalizerTimeSliceType:
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
					case "Authority":
						__gong__assignSliceOfPointers(&inst.Authority, rhs, identifierMap)
					case "Monitoring":
						__gong__assignSliceOfPointers(&inst.Monitoring, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *LocalizerType:
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
				case *MarkerBeaconTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "MarkerBeaconTimeSlice":
						__gong__assignPointer(&inst.MarkerBeaconTimeSlice, rhs, identifierMap)
					}
				case *MarkerBeaconTimeSliceType:
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
					case "Authority":
						__gong__assignSliceOfPointers(&inst.Authority, rhs, identifierMap)
					case "Monitoring":
						__gong__assignSliceOfPointers(&inst.Monitoring, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *MarkerBeaconType:
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
				case *NDBTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "NDBTimeSlice":
						__gong__assignPointer(&inst.NDBTimeSlice, rhs, identifierMap)
					}
				case *NDBTimeSliceType:
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
					case "Authority":
						__gong__assignSliceOfPointers(&inst.Authority, rhs, identifierMap)
					case "Monitoring":
						__gong__assignSliceOfPointers(&inst.Monitoring, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *NDBType:
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
				case *NavaidComponentPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "NavaidComponent":
						__gong__assignPointer(&inst.NavaidComponent, rhs, identifierMap)
					}
				case *NavaidComponentType:
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
				case *NavaidEquipmentMonitoringPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "NavaidEquipmentMonitoring":
						__gong__assignPointer(&inst.NavaidEquipmentMonitoring, rhs, identifierMap)
					}
				case *NavaidEquipmentMonitoringType:
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
				case *NavaidOperationalStatusPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "NavaidOperationalStatus":
						__gong__assignPointer(&inst.NavaidOperationalStatus, rhs, identifierMap)
					}
				case *NavaidOperationalStatusType:
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
				case *NavaidPropertyType:
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
				case *NavaidTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "NavaidTimeSlice":
						__gong__assignPointer(&inst.NavaidTimeSlice, rhs, identifierMap)
					}
				case *NavaidTimeSliceType:
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
					case "TouchDownLiftOff":
						__gong__assignSliceOfPointers(&inst.TouchDownLiftOff, rhs, identifierMap)
					case "NavaidEquipment":
						__gong__assignSliceOfPointers(&inst.NavaidEquipment, rhs, identifierMap)
					case "Location":
						__gong__assignPointer(&inst.Location, rhs, identifierMap)
					case "RunwayDirection":
						__gong__assignSliceOfPointers(&inst.RunwayDirection, rhs, identifierMap)
					case "ServedAirport":
						__gong__assignSliceOfPointers(&inst.ServedAirport, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *NavaidType:
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
				case *PointPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "Point":
						__gong__assignPointer(&inst.Point, rhs, identifierMap)
					case "ElevatedPoint":
						__gong__assignPointer(&inst.ElevatedPoint, rhs, identifierMap)
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
				case *PrecisionApproachRadarTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "PrecisionApproachRadarTimeSlice":
						__gong__assignPointer(&inst.PrecisionApproachRadarTimeSlice, rhs, identifierMap)
					}
				case *PrecisionApproachRadarTimeSliceType:
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
					case "Contact":
						__gong__assignSliceOfPointers(&inst.Contact, rhs, identifierMap)
					case "Location":
						__gong__assignPointer(&inst.Location, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Reflector":
						__gong__assignSliceOfPointers(&inst.Reflector, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *PrecisionApproachRadarType:
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
				case *PrimarySurveillanceRadarTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "PrimarySurveillanceRadarTimeSlice":
						__gong__assignPointer(&inst.PrimarySurveillanceRadarTimeSlice, rhs, identifierMap)
					}
				case *PrimarySurveillanceRadarTimeSliceType:
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
					case "Contact":
						__gong__assignSliceOfPointers(&inst.Contact, rhs, identifierMap)
					case "Location":
						__gong__assignPointer(&inst.Location, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "GroundStation":
						__gong__assignSliceOfPointers(&inst.GroundStation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *PrimarySurveillanceRadarType:
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
				case *RadarComponentPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "RadarComponent":
						__gong__assignPointer(&inst.RadarComponent, rhs, identifierMap)
					}
				case *RadarComponentType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "TheRadarEquipment":
						__gong__assignPointer(&inst.TheRadarEquipment, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *RadarEquipmentPropertyType:
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
				case *RadarSystemTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "RadarSystemTimeSlice":
						__gong__assignPointer(&inst.RadarSystemTimeSlice, rhs, identifierMap)
					}
				case *RadarSystemTimeSliceType:
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
					case "RadarEquipment":
						__gong__assignSliceOfPointers(&inst.RadarEquipment, rhs, identifierMap)
					case "Office":
						__gong__assignSliceOfPointers(&inst.Office, rhs, identifierMap)
					case "AirportHeliport":
						__gong__assignPointer(&inst.AirportHeliport, rhs, identifierMap)
					case "PARRunway":
						__gong__assignSliceOfPointers(&inst.PARRunway, rhs, identifierMap)
					case "Location":
						__gong__assignPointer(&inst.Location, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *RadarSystemType:
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
				case *ReflectorPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "Reflector":
						__gong__assignPointer(&inst.Reflector, rhs, identifierMap)
					}
				case *ReflectorType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "TouchdownReflector":
						__gong__assignPointer(&inst.TouchdownReflector, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
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
				case *SDFTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "SDFTimeSlice":
						__gong__assignPointer(&inst.SDFTimeSlice, rhs, identifierMap)
					}
				case *SDFTimeSliceType:
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
					case "Authority":
						__gong__assignSliceOfPointers(&inst.Authority, rhs, identifierMap)
					case "Monitoring":
						__gong__assignSliceOfPointers(&inst.Monitoring, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *SDFType:
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
				case *SecondarySurveillanceRadarTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "SecondarySurveillanceRadarTimeSlice":
						__gong__assignPointer(&inst.SecondarySurveillanceRadarTimeSlice, rhs, identifierMap)
					}
				case *SecondarySurveillanceRadarTimeSliceType:
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
					case "Contact":
						__gong__assignSliceOfPointers(&inst.Contact, rhs, identifierMap)
					case "Location":
						__gong__assignPointer(&inst.Location, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "GroundStation":
						__gong__assignSliceOfPointers(&inst.GroundStation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *SecondarySurveillanceRadarType:
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
				case *SignificantPointInAirspaceTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "SignificantPointInAirspaceTimeSlice":
						__gong__assignPointer(&inst.SignificantPointInAirspaceTimeSlice, rhs, identifierMap)
					}
				case *SignificantPointInAirspaceTimeSliceType:
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
					case "ContainingAirspace":
						__gong__assignPointer(&inst.ContainingAirspace, rhs, identifierMap)
					case "Location_fixDesignatedPoint":
						__gong__assignPointer(&inst.Location_fixDesignatedPoint, rhs, identifierMap)
					case "Location_navaidSystem":
						__gong__assignPointer(&inst.Location_navaidSystem, rhs, identifierMap)
					case "Location_position":
						__gong__assignPointer(&inst.Location_position, rhs, identifierMap)
					case "Location_runwayPoint":
						__gong__assignPointer(&inst.Location_runwayPoint, rhs, identifierMap)
					case "Location_aimingPoint":
						__gong__assignPointer(&inst.Location_aimingPoint, rhs, identifierMap)
					case "Location_airportReferencePoint":
						__gong__assignPointer(&inst.Location_airportReferencePoint, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *SignificantPointInAirspaceType:
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
				case *SpecialNavigationStationStatusPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "SpecialNavigationStationStatus":
						__gong__assignPointer(&inst.SpecialNavigationStationStatus, rhs, identifierMap)
					}
				case *SpecialNavigationStationStatusType:
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
				case *SpecialNavigationStationTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "SpecialNavigationStationTimeSlice":
						__gong__assignPointer(&inst.SpecialNavigationStationTimeSlice, rhs, identifierMap)
					}
				case *SpecialNavigationStationTimeSliceType:
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
					case "SystemChain":
						__gong__assignPointer(&inst.SystemChain, rhs, identifierMap)
					case "ResponsibleOrganisation":
						__gong__assignPointer(&inst.ResponsibleOrganisation, rhs, identifierMap)
					case "Position":
						__gong__assignPointer(&inst.Position, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *SpecialNavigationStationType:
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
				case *SpecialNavigationSystemPropertyType:
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
				case *SpecialNavigationSystemTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "SpecialNavigationSystemTimeSlice":
						__gong__assignPointer(&inst.SpecialNavigationSystemTimeSlice, rhs, identifierMap)
					}
				case *SpecialNavigationSystemTimeSliceType:
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
					case "ResponsibleOrganisation":
						__gong__assignPointer(&inst.ResponsibleOrganisation, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *SpecialNavigationSystemType:
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
				case *SurveillanceGroundStationPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "SurveillanceGroundStation":
						__gong__assignPointer(&inst.SurveillanceGroundStation, rhs, identifierMap)
					}
				case *SurveillanceGroundStationType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "TheUnit":
						__gong__assignPointer(&inst.TheUnit, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *TACANTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "TACANTimeSlice":
						__gong__assignPointer(&inst.TACANTimeSlice, rhs, identifierMap)
					}
				case *TACANTimeSliceType:
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
					case "Authority":
						__gong__assignSliceOfPointers(&inst.Authority, rhs, identifierMap)
					case "Monitoring":
						__gong__assignSliceOfPointers(&inst.Monitoring, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *TACANType:
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
				case *UnitPropertyType:
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
				case *VORPropertyType:
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
				case *VORTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "VORTimeSlice":
						__gong__assignPointer(&inst.VORTimeSlice, rhs, identifierMap)
					}
				case *VORTimeSliceType:
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
					case "Authority":
						__gong__assignSliceOfPointers(&inst.Authority, rhs, identifierMap)
					case "Monitoring":
						__gong__assignSliceOfPointers(&inst.Monitoring, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *VORType:
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

type AbstractNavaidEquipmentTypeUnmarshaller struct{}

func (u *AbstractNavaidEquipmentTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AbstractNavaidEquipmentType), stage, identifier, instanceName, preserveOrder)
}

func (u *AbstractNavaidEquipmentTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AbstractNavaidEquipmentType)
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

type AbstractNavigationSystemCheckpointTypeUnmarshaller struct{}

func (u *AbstractNavigationSystemCheckpointTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AbstractNavigationSystemCheckpointType), stage, identifier, instanceName, preserveOrder)
}

func (u *AbstractNavigationSystemCheckpointTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AbstractNavigationSystemCheckpointType)
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

type AbstractRadarEquipmentTypeUnmarshaller struct{}

func (u *AbstractRadarEquipmentTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AbstractRadarEquipmentType), stage, identifier, instanceName, preserveOrder)
}

func (u *AbstractRadarEquipmentTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AbstractRadarEquipmentType)
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

type AbstractSurveillanceRadarTypeUnmarshaller struct{}

func (u *AbstractSurveillanceRadarTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AbstractSurveillanceRadarType), stage, identifier, instanceName, preserveOrder)
}

func (u *AbstractSurveillanceRadarTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AbstractSurveillanceRadarType)
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

type AirspacePropertyTypeUnmarshaller struct{}

func (u *AirspacePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AirspacePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *AirspacePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AirspacePropertyType)
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

type AngleIndicationTimeSlicePropertyTypeUnmarshaller struct{}

func (u *AngleIndicationTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AngleIndicationTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *AngleIndicationTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AngleIndicationTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "AngleIndicationTimeSlice":
		GongUnmarshallPointer(&instance.AngleIndicationTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type AngleIndicationTimeSliceTypeUnmarshaller struct{}

func (u *AngleIndicationTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AngleIndicationTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *AngleIndicationTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AngleIndicationTimeSliceType)
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
	case "Fix":
		GongUnmarshallPointer(&instance.Fix, valueExpr, identifierMap)
	case "PointChoice_fixDesignatedPoint":
		GongUnmarshallPointer(&instance.PointChoice_fixDesignatedPoint, valueExpr, identifierMap)
	case "PointChoice_navaidSystem":
		GongUnmarshallPointer(&instance.PointChoice_navaidSystem, valueExpr, identifierMap)
	case "PointChoice_position":
		GongUnmarshallPointer(&instance.PointChoice_position, valueExpr, identifierMap)
	case "PointChoice_runwayPoint":
		GongUnmarshallPointer(&instance.PointChoice_runwayPoint, valueExpr, identifierMap)
	case "PointChoice_aimingPoint":
		GongUnmarshallPointer(&instance.PointChoice_aimingPoint, valueExpr, identifierMap)
	case "PointChoice_airportReferencePoint":
		GongUnmarshallPointer(&instance.PointChoice_airportReferencePoint, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type AngleIndicationTypeUnmarshaller struct{}

func (u *AngleIndicationTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AngleIndicationType), stage, identifier, instanceName, preserveOrder)
}

func (u *AngleIndicationTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AngleIndicationType)
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

type AuthorityForNavaidEquipmentPropertyTypeUnmarshaller struct{}

func (u *AuthorityForNavaidEquipmentPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AuthorityForNavaidEquipmentPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *AuthorityForNavaidEquipmentPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AuthorityForNavaidEquipmentPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "AuthorityForNavaidEquipment":
		GongUnmarshallPointer(&instance.AuthorityForNavaidEquipment, valueExpr, identifierMap)
	}
	return nil
}

type AuthorityForNavaidEquipmentTypeUnmarshaller struct{}

func (u *AuthorityForNavaidEquipmentTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AuthorityForNavaidEquipmentType), stage, identifier, instanceName, preserveOrder)
}

func (u *AuthorityForNavaidEquipmentTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AuthorityForNavaidEquipmentType)
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
	case "TheOrganisationAuthority":
		GongUnmarshallPointer(&instance.TheOrganisationAuthority, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type AuthorityForSpecialNavigationStationPropertyTypeUnmarshaller struct{}

func (u *AuthorityForSpecialNavigationStationPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AuthorityForSpecialNavigationStationPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *AuthorityForSpecialNavigationStationPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AuthorityForSpecialNavigationStationPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "AuthorityForSpecialNavigationStation":
		GongUnmarshallPointer(&instance.AuthorityForSpecialNavigationStation, valueExpr, identifierMap)
	}
	return nil
}

type AuthorityForSpecialNavigationStationTypeUnmarshaller struct{}

func (u *AuthorityForSpecialNavigationStationTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AuthorityForSpecialNavigationStationType), stage, identifier, instanceName, preserveOrder)
}

func (u *AuthorityForSpecialNavigationStationTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AuthorityForSpecialNavigationStationType)
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
	case "TheOrganisationAuthority":
		GongUnmarshallPointer(&instance.TheOrganisationAuthority, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type AuthorityForSpecialNavigationSystemPropertyTypeUnmarshaller struct{}

func (u *AuthorityForSpecialNavigationSystemPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AuthorityForSpecialNavigationSystemPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *AuthorityForSpecialNavigationSystemPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AuthorityForSpecialNavigationSystemPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "AuthorityForSpecialNavigationSystem":
		GongUnmarshallPointer(&instance.AuthorityForSpecialNavigationSystem, valueExpr, identifierMap)
	}
	return nil
}

type AuthorityForSpecialNavigationSystemTypeUnmarshaller struct{}

func (u *AuthorityForSpecialNavigationSystemTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AuthorityForSpecialNavigationSystemType), stage, identifier, instanceName, preserveOrder)
}

func (u *AuthorityForSpecialNavigationSystemTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AuthorityForSpecialNavigationSystemType)
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
	case "TheOrganisationAuthority":
		GongUnmarshallPointer(&instance.TheOrganisationAuthority, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type AzimuthTimeSlicePropertyTypeUnmarshaller struct{}

func (u *AzimuthTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AzimuthTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *AzimuthTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AzimuthTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "AzimuthTimeSlice":
		GongUnmarshallPointer(&instance.AzimuthTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type AzimuthTimeSliceTypeUnmarshaller struct{}

func (u *AzimuthTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AzimuthTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *AzimuthTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AzimuthTimeSliceType)
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
	case "Authority":
		GongUnmarshallSliceOfPointers(&instance.Authority, valueExpr, identifierMap)
	case "Monitoring":
		GongUnmarshallSliceOfPointers(&instance.Monitoring, valueExpr, identifierMap)
	case "Availability":
		GongUnmarshallSliceOfPointers(&instance.Availability, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type AzimuthTypeUnmarshaller struct{}

func (u *AzimuthTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AzimuthType), stage, identifier, instanceName, preserveOrder)
}

func (u *AzimuthTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AzimuthType)
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

type CheckpointINSTimeSlicePropertyTypeUnmarshaller struct{}

func (u *CheckpointINSTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CheckpointINSTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *CheckpointINSTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CheckpointINSTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "CheckpointINSTimeSlice":
		GongUnmarshallPointer(&instance.CheckpointINSTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type CheckpointINSTimeSliceTypeUnmarshaller struct{}

func (u *CheckpointINSTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CheckpointINSTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *CheckpointINSTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CheckpointINSTimeSliceType)
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
	case "Position":
		GongUnmarshallPointer(&instance.Position, valueExpr, identifierMap)
	case "AirportHeliport":
		GongUnmarshallPointer(&instance.AirportHeliport, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type CheckpointINSTypeUnmarshaller struct{}

func (u *CheckpointINSTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CheckpointINSType), stage, identifier, instanceName, preserveOrder)
}

func (u *CheckpointINSTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CheckpointINSType)
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

type CheckpointVORTimeSlicePropertyTypeUnmarshaller struct{}

func (u *CheckpointVORTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CheckpointVORTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *CheckpointVORTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CheckpointVORTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "CheckpointVORTimeSlice":
		GongUnmarshallPointer(&instance.CheckpointVORTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type CheckpointVORTimeSliceTypeUnmarshaller struct{}

func (u *CheckpointVORTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CheckpointVORTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *CheckpointVORTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CheckpointVORTimeSliceType)
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
	case "Position":
		GongUnmarshallPointer(&instance.Position, valueExpr, identifierMap)
	case "AirportHeliport":
		GongUnmarshallPointer(&instance.AirportHeliport, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "CheckPointFacility":
		GongUnmarshallPointer(&instance.CheckPointFacility, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type CheckpointVORTypeUnmarshaller struct{}

func (u *CheckpointVORTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CheckpointVORType), stage, identifier, instanceName, preserveOrder)
}

func (u *CheckpointVORTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CheckpointVORType)
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

type DMETimeSlicePropertyTypeUnmarshaller struct{}

func (u *DMETimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DMETimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *DMETimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DMETimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "DMETimeSlice":
		GongUnmarshallPointer(&instance.DMETimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type DMETimeSliceTypeUnmarshaller struct{}

func (u *DMETimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DMETimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *DMETimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DMETimeSliceType)
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
	case "Authority":
		GongUnmarshallSliceOfPointers(&instance.Authority, valueExpr, identifierMap)
	case "Monitoring":
		GongUnmarshallSliceOfPointers(&instance.Monitoring, valueExpr, identifierMap)
	case "Availability":
		GongUnmarshallSliceOfPointers(&instance.Availability, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type DMETypeUnmarshaller struct{}

func (u *DMETypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DMEType), stage, identifier, instanceName, preserveOrder)
}

func (u *DMETypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DMEType)
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

type DesignatedPointPropertyTypeUnmarshaller struct{}

func (u *DesignatedPointPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DesignatedPointPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *DesignatedPointPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DesignatedPointPropertyType)
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

type DesignatedPointTimeSlicePropertyTypeUnmarshaller struct{}

func (u *DesignatedPointTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DesignatedPointTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *DesignatedPointTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DesignatedPointTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "DesignatedPointTimeSlice":
		GongUnmarshallPointer(&instance.DesignatedPointTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type DesignatedPointTimeSliceTypeUnmarshaller struct{}

func (u *DesignatedPointTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DesignatedPointTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *DesignatedPointTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DesignatedPointTimeSliceType)
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
	case "AimingPoint":
		GongUnmarshallPointer(&instance.AimingPoint, valueExpr, identifierMap)
	case "AirportHeliport":
		GongUnmarshallPointer(&instance.AirportHeliport, valueExpr, identifierMap)
	case "RunwayPoint":
		GongUnmarshallPointer(&instance.RunwayPoint, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type DesignatedPointTypeUnmarshaller struct{}

func (u *DesignatedPointTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DesignatedPointType), stage, identifier, instanceName, preserveOrder)
}

func (u *DesignatedPointTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DesignatedPointType)
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

type DirectionFinderTimeSlicePropertyTypeUnmarshaller struct{}

func (u *DirectionFinderTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DirectionFinderTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *DirectionFinderTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DirectionFinderTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "DirectionFinderTimeSlice":
		GongUnmarshallPointer(&instance.DirectionFinderTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type DirectionFinderTimeSliceTypeUnmarshaller struct{}

func (u *DirectionFinderTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DirectionFinderTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *DirectionFinderTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DirectionFinderTimeSliceType)
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
	case "Authority":
		GongUnmarshallSliceOfPointers(&instance.Authority, valueExpr, identifierMap)
	case "Monitoring":
		GongUnmarshallSliceOfPointers(&instance.Monitoring, valueExpr, identifierMap)
	case "Availability":
		GongUnmarshallSliceOfPointers(&instance.Availability, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "InformationProvision":
		GongUnmarshallSliceOfPointers(&instance.InformationProvision, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type DirectionFinderTypeUnmarshaller struct{}

func (u *DirectionFinderTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DirectionFinderType), stage, identifier, instanceName, preserveOrder)
}

func (u *DirectionFinderTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DirectionFinderType)
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

type DistanceIndicationTimeSlicePropertyTypeUnmarshaller struct{}

func (u *DistanceIndicationTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DistanceIndicationTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *DistanceIndicationTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DistanceIndicationTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "DistanceIndicationTimeSlice":
		GongUnmarshallPointer(&instance.DistanceIndicationTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type DistanceIndicationTimeSliceTypeUnmarshaller struct{}

func (u *DistanceIndicationTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DistanceIndicationTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *DistanceIndicationTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DistanceIndicationTimeSliceType)
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
	case "Fix":
		GongUnmarshallPointer(&instance.Fix, valueExpr, identifierMap)
	case "PointChoice_fixDesignatedPoint":
		GongUnmarshallPointer(&instance.PointChoice_fixDesignatedPoint, valueExpr, identifierMap)
	case "PointChoice_navaidSystem":
		GongUnmarshallPointer(&instance.PointChoice_navaidSystem, valueExpr, identifierMap)
	case "PointChoice_position":
		GongUnmarshallPointer(&instance.PointChoice_position, valueExpr, identifierMap)
	case "PointChoice_runwayPoint":
		GongUnmarshallPointer(&instance.PointChoice_runwayPoint, valueExpr, identifierMap)
	case "PointChoice_aimingPoint":
		GongUnmarshallPointer(&instance.PointChoice_aimingPoint, valueExpr, identifierMap)
	case "PointChoice_airportReferencePoint":
		GongUnmarshallPointer(&instance.PointChoice_airportReferencePoint, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type DistanceIndicationTypeUnmarshaller struct{}

func (u *DistanceIndicationTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DistanceIndicationType), stage, identifier, instanceName, preserveOrder)
}

func (u *DistanceIndicationTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DistanceIndicationType)
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

type ElevationTimeSlicePropertyTypeUnmarshaller struct{}

func (u *ElevationTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ElevationTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *ElevationTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ElevationTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "ElevationTimeSlice":
		GongUnmarshallPointer(&instance.ElevationTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type ElevationTimeSliceTypeUnmarshaller struct{}

func (u *ElevationTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ElevationTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *ElevationTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ElevationTimeSliceType)
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
	case "Authority":
		GongUnmarshallSliceOfPointers(&instance.Authority, valueExpr, identifierMap)
	case "Monitoring":
		GongUnmarshallSliceOfPointers(&instance.Monitoring, valueExpr, identifierMap)
	case "Availability":
		GongUnmarshallSliceOfPointers(&instance.Availability, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type ElevationTypeUnmarshaller struct{}

func (u *ElevationTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ElevationType), stage, identifier, instanceName, preserveOrder)
}

func (u *ElevationTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ElevationType)
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

type GlidepathTimeSlicePropertyTypeUnmarshaller struct{}

func (u *GlidepathTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(GlidepathTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *GlidepathTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GlidepathTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "GlidepathTimeSlice":
		GongUnmarshallPointer(&instance.GlidepathTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type GlidepathTimeSliceTypeUnmarshaller struct{}

func (u *GlidepathTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(GlidepathTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *GlidepathTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GlidepathTimeSliceType)
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
	case "Authority":
		GongUnmarshallSliceOfPointers(&instance.Authority, valueExpr, identifierMap)
	case "Monitoring":
		GongUnmarshallSliceOfPointers(&instance.Monitoring, valueExpr, identifierMap)
	case "Availability":
		GongUnmarshallSliceOfPointers(&instance.Availability, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type GlidepathTypeUnmarshaller struct{}

func (u *GlidepathTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(GlidepathType), stage, identifier, instanceName, preserveOrder)
}

func (u *GlidepathTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GlidepathType)
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

type InformationServicePropertyTypeUnmarshaller struct{}

func (u *InformationServicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(InformationServicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *InformationServicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*InformationServicePropertyType)
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

type LocalizerTimeSlicePropertyTypeUnmarshaller struct{}

func (u *LocalizerTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(LocalizerTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *LocalizerTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*LocalizerTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "LocalizerTimeSlice":
		GongUnmarshallPointer(&instance.LocalizerTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type LocalizerTimeSliceTypeUnmarshaller struct{}

func (u *LocalizerTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(LocalizerTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *LocalizerTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*LocalizerTimeSliceType)
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
	case "Authority":
		GongUnmarshallSliceOfPointers(&instance.Authority, valueExpr, identifierMap)
	case "Monitoring":
		GongUnmarshallSliceOfPointers(&instance.Monitoring, valueExpr, identifierMap)
	case "Availability":
		GongUnmarshallSliceOfPointers(&instance.Availability, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type LocalizerTypeUnmarshaller struct{}

func (u *LocalizerTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(LocalizerType), stage, identifier, instanceName, preserveOrder)
}

func (u *LocalizerTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*LocalizerType)
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

type MarkerBeaconTimeSlicePropertyTypeUnmarshaller struct{}

func (u *MarkerBeaconTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(MarkerBeaconTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *MarkerBeaconTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MarkerBeaconTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "MarkerBeaconTimeSlice":
		GongUnmarshallPointer(&instance.MarkerBeaconTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type MarkerBeaconTimeSliceTypeUnmarshaller struct{}

func (u *MarkerBeaconTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(MarkerBeaconTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *MarkerBeaconTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MarkerBeaconTimeSliceType)
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
	case "Authority":
		GongUnmarshallSliceOfPointers(&instance.Authority, valueExpr, identifierMap)
	case "Monitoring":
		GongUnmarshallSliceOfPointers(&instance.Monitoring, valueExpr, identifierMap)
	case "Availability":
		GongUnmarshallSliceOfPointers(&instance.Availability, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type MarkerBeaconTypeUnmarshaller struct{}

func (u *MarkerBeaconTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(MarkerBeaconType), stage, identifier, instanceName, preserveOrder)
}

func (u *MarkerBeaconTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MarkerBeaconType)
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

type NDBTimeSlicePropertyTypeUnmarshaller struct{}

func (u *NDBTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(NDBTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *NDBTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*NDBTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "NDBTimeSlice":
		GongUnmarshallPointer(&instance.NDBTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type NDBTimeSliceTypeUnmarshaller struct{}

func (u *NDBTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(NDBTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *NDBTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*NDBTimeSliceType)
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
	case "Authority":
		GongUnmarshallSliceOfPointers(&instance.Authority, valueExpr, identifierMap)
	case "Monitoring":
		GongUnmarshallSliceOfPointers(&instance.Monitoring, valueExpr, identifierMap)
	case "Availability":
		GongUnmarshallSliceOfPointers(&instance.Availability, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type NDBTypeUnmarshaller struct{}

func (u *NDBTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(NDBType), stage, identifier, instanceName, preserveOrder)
}

func (u *NDBTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*NDBType)
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

type NavaidComponentPropertyTypeUnmarshaller struct{}

func (u *NavaidComponentPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(NavaidComponentPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *NavaidComponentPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*NavaidComponentPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "NavaidComponent":
		GongUnmarshallPointer(&instance.NavaidComponent, valueExpr, identifierMap)
	}
	return nil
}

type NavaidComponentTypeUnmarshaller struct{}

func (u *NavaidComponentTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(NavaidComponentType), stage, identifier, instanceName, preserveOrder)
}

func (u *NavaidComponentTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*NavaidComponentType)
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

type NavaidEquipmentMonitoringPropertyTypeUnmarshaller struct{}

func (u *NavaidEquipmentMonitoringPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(NavaidEquipmentMonitoringPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *NavaidEquipmentMonitoringPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*NavaidEquipmentMonitoringPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "NavaidEquipmentMonitoring":
		GongUnmarshallPointer(&instance.NavaidEquipmentMonitoring, valueExpr, identifierMap)
	}
	return nil
}

type NavaidEquipmentMonitoringTypeUnmarshaller struct{}

func (u *NavaidEquipmentMonitoringTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(NavaidEquipmentMonitoringType), stage, identifier, instanceName, preserveOrder)
}

func (u *NavaidEquipmentMonitoringTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*NavaidEquipmentMonitoringType)
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

type NavaidOperationalStatusPropertyTypeUnmarshaller struct{}

func (u *NavaidOperationalStatusPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(NavaidOperationalStatusPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *NavaidOperationalStatusPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*NavaidOperationalStatusPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "NavaidOperationalStatus":
		GongUnmarshallPointer(&instance.NavaidOperationalStatus, valueExpr, identifierMap)
	}
	return nil
}

type NavaidOperationalStatusTypeUnmarshaller struct{}

func (u *NavaidOperationalStatusTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(NavaidOperationalStatusType), stage, identifier, instanceName, preserveOrder)
}

func (u *NavaidOperationalStatusTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*NavaidOperationalStatusType)
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

type NavaidPropertyTypeUnmarshaller struct{}

func (u *NavaidPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(NavaidPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *NavaidPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*NavaidPropertyType)
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

type NavaidTimeSlicePropertyTypeUnmarshaller struct{}

func (u *NavaidTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(NavaidTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *NavaidTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*NavaidTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "NavaidTimeSlice":
		GongUnmarshallPointer(&instance.NavaidTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type NavaidTimeSliceTypeUnmarshaller struct{}

func (u *NavaidTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(NavaidTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *NavaidTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*NavaidTimeSliceType)
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
	case "TouchDownLiftOff":
		GongUnmarshallSliceOfPointers(&instance.TouchDownLiftOff, valueExpr, identifierMap)
	case "NavaidEquipment":
		GongUnmarshallSliceOfPointers(&instance.NavaidEquipment, valueExpr, identifierMap)
	case "Location":
		GongUnmarshallPointer(&instance.Location, valueExpr, identifierMap)
	case "RunwayDirection":
		GongUnmarshallSliceOfPointers(&instance.RunwayDirection, valueExpr, identifierMap)
	case "ServedAirport":
		GongUnmarshallSliceOfPointers(&instance.ServedAirport, valueExpr, identifierMap)
	case "Availability":
		GongUnmarshallSliceOfPointers(&instance.Availability, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type NavaidTypeUnmarshaller struct{}

func (u *NavaidTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(NavaidType), stage, identifier, instanceName, preserveOrder)
}

func (u *NavaidTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*NavaidType)
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

type PointPropertyTypeUnmarshaller struct{}

func (u *PointPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PointPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *PointPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PointPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "Point":
		GongUnmarshallPointer(&instance.Point, valueExpr, identifierMap)
	case "ElevatedPoint":
		GongUnmarshallPointer(&instance.ElevatedPoint, valueExpr, identifierMap)
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

type PrecisionApproachRadarTimeSlicePropertyTypeUnmarshaller struct{}

func (u *PrecisionApproachRadarTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PrecisionApproachRadarTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *PrecisionApproachRadarTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PrecisionApproachRadarTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "PrecisionApproachRadarTimeSlice":
		GongUnmarshallPointer(&instance.PrecisionApproachRadarTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type PrecisionApproachRadarTimeSliceTypeUnmarshaller struct{}

func (u *PrecisionApproachRadarTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PrecisionApproachRadarTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *PrecisionApproachRadarTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PrecisionApproachRadarTimeSliceType)
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
	case "Contact":
		GongUnmarshallSliceOfPointers(&instance.Contact, valueExpr, identifierMap)
	case "Location":
		GongUnmarshallPointer(&instance.Location, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Reflector":
		GongUnmarshallSliceOfPointers(&instance.Reflector, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type PrecisionApproachRadarTypeUnmarshaller struct{}

func (u *PrecisionApproachRadarTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PrecisionApproachRadarType), stage, identifier, instanceName, preserveOrder)
}

func (u *PrecisionApproachRadarTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PrecisionApproachRadarType)
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

type PrimarySurveillanceRadarTimeSlicePropertyTypeUnmarshaller struct{}

func (u *PrimarySurveillanceRadarTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PrimarySurveillanceRadarTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *PrimarySurveillanceRadarTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PrimarySurveillanceRadarTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "PrimarySurveillanceRadarTimeSlice":
		GongUnmarshallPointer(&instance.PrimarySurveillanceRadarTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type PrimarySurveillanceRadarTimeSliceTypeUnmarshaller struct{}

func (u *PrimarySurveillanceRadarTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PrimarySurveillanceRadarTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *PrimarySurveillanceRadarTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PrimarySurveillanceRadarTimeSliceType)
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
	case "Contact":
		GongUnmarshallSliceOfPointers(&instance.Contact, valueExpr, identifierMap)
	case "Location":
		GongUnmarshallPointer(&instance.Location, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "GroundStation":
		GongUnmarshallSliceOfPointers(&instance.GroundStation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type PrimarySurveillanceRadarTypeUnmarshaller struct{}

func (u *PrimarySurveillanceRadarTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PrimarySurveillanceRadarType), stage, identifier, instanceName, preserveOrder)
}

func (u *PrimarySurveillanceRadarTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PrimarySurveillanceRadarType)
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

type RadarComponentPropertyTypeUnmarshaller struct{}

func (u *RadarComponentPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RadarComponentPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *RadarComponentPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RadarComponentPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "RadarComponent":
		GongUnmarshallPointer(&instance.RadarComponent, valueExpr, identifierMap)
	}
	return nil
}

type RadarComponentTypeUnmarshaller struct{}

func (u *RadarComponentTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RadarComponentType), stage, identifier, instanceName, preserveOrder)
}

func (u *RadarComponentTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RadarComponentType)
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
	case "TheRadarEquipment":
		GongUnmarshallPointer(&instance.TheRadarEquipment, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type RadarEquipmentPropertyTypeUnmarshaller struct{}

func (u *RadarEquipmentPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RadarEquipmentPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *RadarEquipmentPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RadarEquipmentPropertyType)
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

type RadarSystemTimeSlicePropertyTypeUnmarshaller struct{}

func (u *RadarSystemTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RadarSystemTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *RadarSystemTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RadarSystemTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "RadarSystemTimeSlice":
		GongUnmarshallPointer(&instance.RadarSystemTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type RadarSystemTimeSliceTypeUnmarshaller struct{}

func (u *RadarSystemTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RadarSystemTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *RadarSystemTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RadarSystemTimeSliceType)
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
	case "RadarEquipment":
		GongUnmarshallSliceOfPointers(&instance.RadarEquipment, valueExpr, identifierMap)
	case "Office":
		GongUnmarshallSliceOfPointers(&instance.Office, valueExpr, identifierMap)
	case "AirportHeliport":
		GongUnmarshallPointer(&instance.AirportHeliport, valueExpr, identifierMap)
	case "PARRunway":
		GongUnmarshallSliceOfPointers(&instance.PARRunway, valueExpr, identifierMap)
	case "Location":
		GongUnmarshallPointer(&instance.Location, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type RadarSystemTypeUnmarshaller struct{}

func (u *RadarSystemTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RadarSystemType), stage, identifier, instanceName, preserveOrder)
}

func (u *RadarSystemTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RadarSystemType)
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

type ReflectorPropertyTypeUnmarshaller struct{}

func (u *ReflectorPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ReflectorPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *ReflectorPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ReflectorPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "Reflector":
		GongUnmarshallPointer(&instance.Reflector, valueExpr, identifierMap)
	}
	return nil
}

type ReflectorTypeUnmarshaller struct{}

func (u *ReflectorTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ReflectorType), stage, identifier, instanceName, preserveOrder)
}

func (u *ReflectorTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ReflectorType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "TouchdownReflector":
		GongUnmarshallPointer(&instance.TouchdownReflector, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
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

type SDFTimeSlicePropertyTypeUnmarshaller struct{}

func (u *SDFTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SDFTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *SDFTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SDFTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "SDFTimeSlice":
		GongUnmarshallPointer(&instance.SDFTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type SDFTimeSliceTypeUnmarshaller struct{}

func (u *SDFTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SDFTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *SDFTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SDFTimeSliceType)
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
	case "Authority":
		GongUnmarshallSliceOfPointers(&instance.Authority, valueExpr, identifierMap)
	case "Monitoring":
		GongUnmarshallSliceOfPointers(&instance.Monitoring, valueExpr, identifierMap)
	case "Availability":
		GongUnmarshallSliceOfPointers(&instance.Availability, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type SDFTypeUnmarshaller struct{}

func (u *SDFTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SDFType), stage, identifier, instanceName, preserveOrder)
}

func (u *SDFTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SDFType)
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

type SecondarySurveillanceRadarTimeSlicePropertyTypeUnmarshaller struct{}

func (u *SecondarySurveillanceRadarTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SecondarySurveillanceRadarTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *SecondarySurveillanceRadarTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SecondarySurveillanceRadarTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "SecondarySurveillanceRadarTimeSlice":
		GongUnmarshallPointer(&instance.SecondarySurveillanceRadarTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type SecondarySurveillanceRadarTimeSliceTypeUnmarshaller struct{}

func (u *SecondarySurveillanceRadarTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SecondarySurveillanceRadarTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *SecondarySurveillanceRadarTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SecondarySurveillanceRadarTimeSliceType)
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
	case "Contact":
		GongUnmarshallSliceOfPointers(&instance.Contact, valueExpr, identifierMap)
	case "Location":
		GongUnmarshallPointer(&instance.Location, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "GroundStation":
		GongUnmarshallSliceOfPointers(&instance.GroundStation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type SecondarySurveillanceRadarTypeUnmarshaller struct{}

func (u *SecondarySurveillanceRadarTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SecondarySurveillanceRadarType), stage, identifier, instanceName, preserveOrder)
}

func (u *SecondarySurveillanceRadarTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SecondarySurveillanceRadarType)
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

type SignificantPointInAirspaceTimeSlicePropertyTypeUnmarshaller struct{}

func (u *SignificantPointInAirspaceTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SignificantPointInAirspaceTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *SignificantPointInAirspaceTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SignificantPointInAirspaceTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "SignificantPointInAirspaceTimeSlice":
		GongUnmarshallPointer(&instance.SignificantPointInAirspaceTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type SignificantPointInAirspaceTimeSliceTypeUnmarshaller struct{}

func (u *SignificantPointInAirspaceTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SignificantPointInAirspaceTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *SignificantPointInAirspaceTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SignificantPointInAirspaceTimeSliceType)
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
	case "ContainingAirspace":
		GongUnmarshallPointer(&instance.ContainingAirspace, valueExpr, identifierMap)
	case "Location_fixDesignatedPoint":
		GongUnmarshallPointer(&instance.Location_fixDesignatedPoint, valueExpr, identifierMap)
	case "Location_navaidSystem":
		GongUnmarshallPointer(&instance.Location_navaidSystem, valueExpr, identifierMap)
	case "Location_position":
		GongUnmarshallPointer(&instance.Location_position, valueExpr, identifierMap)
	case "Location_runwayPoint":
		GongUnmarshallPointer(&instance.Location_runwayPoint, valueExpr, identifierMap)
	case "Location_aimingPoint":
		GongUnmarshallPointer(&instance.Location_aimingPoint, valueExpr, identifierMap)
	case "Location_airportReferencePoint":
		GongUnmarshallPointer(&instance.Location_airportReferencePoint, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type SignificantPointInAirspaceTypeUnmarshaller struct{}

func (u *SignificantPointInAirspaceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SignificantPointInAirspaceType), stage, identifier, instanceName, preserveOrder)
}

func (u *SignificantPointInAirspaceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SignificantPointInAirspaceType)
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

type SpecialNavigationStationStatusPropertyTypeUnmarshaller struct{}

func (u *SpecialNavigationStationStatusPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SpecialNavigationStationStatusPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *SpecialNavigationStationStatusPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SpecialNavigationStationStatusPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "SpecialNavigationStationStatus":
		GongUnmarshallPointer(&instance.SpecialNavigationStationStatus, valueExpr, identifierMap)
	}
	return nil
}

type SpecialNavigationStationStatusTypeUnmarshaller struct{}

func (u *SpecialNavigationStationStatusTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SpecialNavigationStationStatusType), stage, identifier, instanceName, preserveOrder)
}

func (u *SpecialNavigationStationStatusTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SpecialNavigationStationStatusType)
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

type SpecialNavigationStationTimeSlicePropertyTypeUnmarshaller struct{}

func (u *SpecialNavigationStationTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SpecialNavigationStationTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *SpecialNavigationStationTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SpecialNavigationStationTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "SpecialNavigationStationTimeSlice":
		GongUnmarshallPointer(&instance.SpecialNavigationStationTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type SpecialNavigationStationTimeSliceTypeUnmarshaller struct{}

func (u *SpecialNavigationStationTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SpecialNavigationStationTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *SpecialNavigationStationTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SpecialNavigationStationTimeSliceType)
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
	case "SystemChain":
		GongUnmarshallPointer(&instance.SystemChain, valueExpr, identifierMap)
	case "ResponsibleOrganisation":
		GongUnmarshallPointer(&instance.ResponsibleOrganisation, valueExpr, identifierMap)
	case "Position":
		GongUnmarshallPointer(&instance.Position, valueExpr, identifierMap)
	case "Availability":
		GongUnmarshallSliceOfPointers(&instance.Availability, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type SpecialNavigationStationTypeUnmarshaller struct{}

func (u *SpecialNavigationStationTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SpecialNavigationStationType), stage, identifier, instanceName, preserveOrder)
}

func (u *SpecialNavigationStationTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SpecialNavigationStationType)
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

type SpecialNavigationSystemPropertyTypeUnmarshaller struct{}

func (u *SpecialNavigationSystemPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SpecialNavigationSystemPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *SpecialNavigationSystemPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SpecialNavigationSystemPropertyType)
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

type SpecialNavigationSystemTimeSlicePropertyTypeUnmarshaller struct{}

func (u *SpecialNavigationSystemTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SpecialNavigationSystemTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *SpecialNavigationSystemTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SpecialNavigationSystemTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "SpecialNavigationSystemTimeSlice":
		GongUnmarshallPointer(&instance.SpecialNavigationSystemTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type SpecialNavigationSystemTimeSliceTypeUnmarshaller struct{}

func (u *SpecialNavigationSystemTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SpecialNavigationSystemTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *SpecialNavigationSystemTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SpecialNavigationSystemTimeSliceType)
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
	case "ResponsibleOrganisation":
		GongUnmarshallPointer(&instance.ResponsibleOrganisation, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type SpecialNavigationSystemTypeUnmarshaller struct{}

func (u *SpecialNavigationSystemTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SpecialNavigationSystemType), stage, identifier, instanceName, preserveOrder)
}

func (u *SpecialNavigationSystemTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SpecialNavigationSystemType)
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

type SurveillanceGroundStationPropertyTypeUnmarshaller struct{}

func (u *SurveillanceGroundStationPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SurveillanceGroundStationPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *SurveillanceGroundStationPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SurveillanceGroundStationPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "SurveillanceGroundStation":
		GongUnmarshallPointer(&instance.SurveillanceGroundStation, valueExpr, identifierMap)
	}
	return nil
}

type SurveillanceGroundStationTypeUnmarshaller struct{}

func (u *SurveillanceGroundStationTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SurveillanceGroundStationType), stage, identifier, instanceName, preserveOrder)
}

func (u *SurveillanceGroundStationTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SurveillanceGroundStationType)
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
	case "TheUnit":
		GongUnmarshallPointer(&instance.TheUnit, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type TACANTimeSlicePropertyTypeUnmarshaller struct{}

func (u *TACANTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TACANTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *TACANTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TACANTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "TACANTimeSlice":
		GongUnmarshallPointer(&instance.TACANTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type TACANTimeSliceTypeUnmarshaller struct{}

func (u *TACANTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TACANTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *TACANTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TACANTimeSliceType)
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
	case "Authority":
		GongUnmarshallSliceOfPointers(&instance.Authority, valueExpr, identifierMap)
	case "Monitoring":
		GongUnmarshallSliceOfPointers(&instance.Monitoring, valueExpr, identifierMap)
	case "Availability":
		GongUnmarshallSliceOfPointers(&instance.Availability, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type TACANTypeUnmarshaller struct{}

func (u *TACANTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TACANType), stage, identifier, instanceName, preserveOrder)
}

func (u *TACANTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TACANType)
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

type UnitPropertyTypeUnmarshaller struct{}

func (u *UnitPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(UnitPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *UnitPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*UnitPropertyType)
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

type VORPropertyTypeUnmarshaller struct{}

func (u *VORPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(VORPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *VORPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*VORPropertyType)
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

type VORTimeSlicePropertyTypeUnmarshaller struct{}

func (u *VORTimeSlicePropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(VORTimeSlicePropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *VORTimeSlicePropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*VORTimeSlicePropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Owns":
		instance.Owns = GongExtractBool(valueExpr)
	case "VORTimeSlice":
		GongUnmarshallPointer(&instance.VORTimeSlice, valueExpr, identifierMap)
	}
	return nil
}

type VORTimeSliceTypeUnmarshaller struct{}

func (u *VORTimeSliceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(VORTimeSliceType), stage, identifier, instanceName, preserveOrder)
}

func (u *VORTimeSliceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*VORTimeSliceType)
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
	case "Authority":
		GongUnmarshallSliceOfPointers(&instance.Authority, valueExpr, identifierMap)
	case "Monitoring":
		GongUnmarshallSliceOfPointers(&instance.Monitoring, valueExpr, identifierMap)
	case "Availability":
		GongUnmarshallSliceOfPointers(&instance.Availability, valueExpr, identifierMap)
	case "Annotation":
		GongUnmarshallSliceOfPointers(&instance.Annotation, valueExpr, identifierMap)
	case "Extension":
		GongUnmarshallSliceOfPointers(&instance.Extension, valueExpr, identifierMap)
	}
	return nil
}

type VORTypeUnmarshaller struct{}

func (u *VORTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(VORType), stage, identifier, instanceName, preserveOrder)
}

func (u *VORTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*VORType)
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
