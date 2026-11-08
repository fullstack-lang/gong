// generated code - do not edit
package models

import (
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"sort"
	"strings"
	"time"
)

var (
	_ = sort.Slice
	_ = time.Hour
)

const GongIdentifiersDecls = `
	{{Identifier}} := (&models.{{GeneratedStructName}}{Name: {{GeneratedFieldNameValue}}}).Stage(stage)`

const GongUnstageStmt = `
	{{Identifier}}.Unstage(stage)`


const GongStringInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const GongMetaFieldStructInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = ` + `{{GeneratedFieldNameValue}}`

const GongStringEnumInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const GongNumberInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const GongPointerFieldInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const GongSliceOfPointersFieldInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = append({{Identifier}}.{{GeneratedFieldName}}, {{GeneratedFieldNameValue}})`

const GongTimeInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}}, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "{{GeneratedFieldNameValue}}")`

// __gong__toRawStringLiteral formats a string into safe Go source code,
// using backticks to preserve newlines and readability.
func __gong__toRawStringLiteral(s string) string {
	// Step 1: Replace every backtick with a closing backtick,
	// a double-quoted backtick, and an opening backtick.
	escaped := strings.ReplaceAll(s, "`", "` + \"`\" + `")

	// Step 2: Wrap the entire resulting string in backticks.
	result := "`" + escaped + "`"

	// Step 3: Clean up any empty raw strings (``) at the boundaries
	// just in case your original string started or ended with a backtick.
	result = strings.ReplaceAll(result, "`` + ", "")
	result = strings.ReplaceAll(result, " + ``", "")

	return result
}

func __gong__marshallString[T ~string](ident, fieldName string, val T) string {
	res := strings.ReplaceAll(GongStringInitStatement, "{{Identifier}}", ident)
	res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", fieldName)
	return strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(string(val)))
}

func __gong__marshallInt[T ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64](ident, fieldName string, val T) string {
	res := strings.ReplaceAll(GongNumberInitStatement, "{{Identifier}}", ident)
	res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", fieldName)
	return strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", val))
}

func __gong__marshallBool[T ~bool](ident, fieldName string, val T) string {
	res := strings.ReplaceAll(GongNumberInitStatement, "{{Identifier}}", ident)
	res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", fieldName)
	return strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", val))
}

func __gong__marshallFloat[T ~float32 | ~float64](ident, fieldName string, val T) string {
	res := strings.ReplaceAll(GongNumberInitStatement, "{{Identifier}}", ident)
	res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", fieldName)
	return strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", val))
}

func __gong__marshallTime(ident, fieldName, valStr string) string {
	res := strings.ReplaceAll(GongTimeInitStatement, "{{Identifier}}", ident)
	res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", fieldName)
	return strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", valStr)
}

func __gong__marshallPointer(ident, fieldName, targetIdent string) string {
	res := strings.ReplaceAll(GongPointerFieldInitStatement, "{{Identifier}}", ident)
	res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", fieldName)
	return strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", targetIdent)
}

func __gong__marshallSliceOfPointers(ident, fieldName, targetIdent string) string {
	res := strings.ReplaceAll(GongSliceOfPointersFieldInitStatement, "{{Identifier}}", ident)
	res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", fieldName)
	return strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", targetIdent)
}

func __gong__marshallEnumString(ident, fieldName, codeStr string) string {
	val := "\"\""
	if codeStr != "" {
		val = "models." + codeStr
	}
	res := strings.ReplaceAll(GongStringEnumInitStatement, "{{Identifier}}", ident)
	res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", fieldName)
	return strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", val)
}

func __gong__marshallEnumInt(ident, fieldName, codeStr string) string {
	val := "0"
	if codeStr != "" {
		val = "models." + codeStr
	}
	res := strings.ReplaceAll(GongNumberInitStatement, "{{Identifier}}", ident)
	res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", fieldName)
	return strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", val)
}

func __gong__marshallMeta(ident, fieldName, val string) string {
	res := strings.ReplaceAll(GongMetaFieldStructInitStatement, "{{Identifier}}", ident)
	res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", fieldName)
	return strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", val)
}

// __gong__sortStageSetInstances sorts instances by their staged order
func __gong__sortStageSetInstances[T comparable](instances map[T]struct{}, orderMap map[T]uint) []T {
	ordered := make([]T, 0, len(instances))
	for inst := range instances {
		ordered = append(ordered, inst)
	}
	sort.Slice(ordered, func(i, j int) bool {
		return orderMap[ordered[i]] < orderMap[ordered[j]]
	})
	return ordered
}

// MarshallFile marshall the stage content into a file as an instanciation into a stage
// according to the marshalling policy of the stage.
//
// In GongMarshallingAppendCommit mode, it will append the last commit to the file.
// In other modes, it will rewrite the entire file.
func (stage *Stage) MarshallFile(filename string, args ...string) {

	if stage.GetGongMarshallingMode() == GongMarshallingAppendCommit {
		contentBytes, err := os.ReadFile(filename)

		// if the file does not exist, marshall the full stage
		if os.IsNotExist(err) {
			file, createErr := os.Create(filename)
			if createErr != nil {
				log.Fatal(createErr.Error())
			}
			defer file.Close()
			stage.Marshall(file, args...)
			return
		}
		if err != nil {
			log.Fatal(err.Error())
		}

		content := string(contentBytes)

		if stage.IsSquashing() {
			// we squash: we want to clear the current function body
			// and let the append logic write the squashed commit
			const fnSig = "func _(stage *models.Stage) {"
			if before, _, ok := strings.Cut(content, fnSig); ok {
				content = before + fnSig + "\n}\n"
			}
		}

		if stage.IsApplyingBackwardCommit() {
			// we are going backward, we need to remove the last forward commit from the file

			// because commitsBehind has been incremented before the call to this function
			// the index of the commit to remove is len(forwardCommits) - commitsBehind
			forwardCommits := stage.GetForwardCommits()
			commitIndexToRemove := len(forwardCommits) - stage.GetCommitsBehind()

			if commitIndexToRemove < 0 || commitIndexToRemove >= len(forwardCommits) {
				return // Should not happen if history is consistent
			}

			commitToRemove := forwardCommits[commitIndexToRemove]

			lastIndex := strings.LastIndex(content, commitToRemove+"\n")
			if lastIndex != -1 {
				newContent := content[:lastIndex] + content[lastIndex+len(commitToRemove)+1:]
				err = os.WriteFile(filename, []byte(newContent), 0644)
				if err != nil {
					log.Fatal(err.Error())
				}
			} else {
				lastIndex = strings.LastIndex(content, commitToRemove)
				if lastIndex != -1 {
					newContent := content[:lastIndex] + content[lastIndex+len(commitToRemove):]
					err = os.WriteFile(filename, []byte(newContent), 0644)
					if err != nil {
						log.Fatal(err.Error())
					}
				} else {
					// The commit block was not found. This typically happens for the initial
					// commit which is formatted differently (the lines after func _(stage *models.Stage) {).
					// We rewrite the entire file with the current (rewound) stage state to safely remove it.
					file, createErr := os.Create(filename)
					if createErr != nil {
						log.Fatal(createErr.Error())
					}
					defer file.Close()
					stage.Marshall(file, args...)
				}
			}
			return // we are done for the backward case
		}

		if stage.IsApplyingForwardCommit() {
			// bypass the modified check
		} else if !stage.IsModified() {
			return
		}

		forwardCommits := stage.GetForwardCommits()
		if len(forwardCommits) == 0 {
			return // nothing to do
		}

		activeCommits := len(forwardCommits) - stage.GetCommitsBehind()
		if activeCommits <= 0 {
			return
		}
		forwardCommit := forwardCommits[activeCommits-1]

		// append before the ending brace of the func
		lastBrace := strings.LastIndex(content, "}")
		if lastBrace == -1 {
			// if no brace, something is wrong with the file, so we rewrite it
			file, createErr := os.Create(filename)
			if createErr != nil {
				log.Fatal(createErr.Error())
			}
			defer file.Close()
			stage.Marshall(file, args...)
			return
		}

		contentBeforeBrace := content[:lastBrace]
		trimmedContentBeforeBrace := strings.TrimSpace(contentBeforeBrace)
		emptyBody := stage.IsSquashing() ||
			strings.HasSuffix(trimmedContentBeforeBrace, "func _(stage *models.Stage) {") ||
			strings.HasSuffix(trimmedContentBeforeBrace, "// insertion point for setup of pointers") ||
			strings.HasSuffix(trimmedContentBeforeBrace, "// Phase 3: Pointer Setups (Intra-stage and Cross-stage pointers)\n\t// ------------------------------------------------------------------------")

		// check if the file ends with stage.Commit() before the brace
		if !emptyBody && !strings.HasSuffix(trimmedContentBeforeBrace, "stage.Commit()") {
			contentBeforeBrace = contentBeforeBrace + "\n\tstage.Commit()\n"
		}

		// insert the commit statements before the last brace
		newContent := contentBeforeBrace + forwardCommit + "\n" + content[lastBrace:]

		err = os.WriteFile(filename, []byte(newContent), 0644)
		if err != nil {
			log.Fatal(err.Error())
		}

	} else {
		file, err := os.Create(filename)
		if err != nil {
			log.Fatal(err.Error())
		}
		defer file.Close()

		stage.Marshall(file, args...)
	}
}

// Marshall marshall the stage content into the file as an instanciation into a stage
func (stage *Stage) Marshall(file *os.File, args ...string) {

	name := file.Name()

	if !strings.HasSuffix(name, ".go") {
		log.Println(name + " is not a go filename")
	}

	log.Printf("Marshalling %s", name)

	res, err := stage.MarshallToString(args...)
	if err != nil {
		log.Fatalln("Error marshalling to string:", err)
	}

	fmt.Fprintln(file, res)
}

// MarshallToString marshall the stage content into a string
func (stage *Stage) MarshallToString(args ...string) (res string, err error) {
	packageName := "main"
	modelsPackageName := "github.com/fullstack-lang/gong/lib/maplibre/go/models"
	if len(args) == 1 {
		packageName = args[0]
	} else if len(args) >= 2 {
		if args[0] != "" {
			modelsPackageName = args[0]
		}
		packageName = args[1]
	}

	var declarations strings.Builder
	var values strings.Builder
	var pointers strings.Builder
	var lastStageDecl string
	var lastStageVal string
	var lastStagePtr string
	_ = lastStageDecl
	_ = lastStageVal
	_ = lastStagePtr

	if stage != nil {
		for _, coordinate := range __gong__sortStageSetInstances(stage.Coordinates, stage.Coordinate_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			coordinateIdent := coordinate.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Coordinate{Name: %s}).Stage(stage)", coordinateIdent, __gong__toRawStringLiteral(coordinate.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", coordinateIdent, __gong__toRawStringLiteral(coordinate.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Longitude = %f", coordinateIdent, coordinate.Longitude))
			values.WriteString(fmt.Sprintf("\n\t%s.Latitude = %f", coordinateIdent, coordinate.Latitude))
			values.WriteString(fmt.Sprintf("\n\t%s.Altitude = %f", coordinateIdent, coordinate.Altitude))
		}
	}
	if stage != nil {
		for _, layer := range __gong__sortStageSetInstances(stage.Layers, stage.Layer_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			layerIdent := layer.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Layer{Name: %s}).Stage(stage)", layerIdent, __gong__toRawStringLiteral(layer.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", layerIdent, __gong__toRawStringLiteral(layer.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsVisible = %t", layerIdent, layer.IsVisible))
			for _, elem := range layer.Polygons {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Polygons = append(%s.Polygons, %s)", layerIdent, layerIdent, targetIdent))
			}
			for _, elem := range layer.Polylines {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Polylines = append(%s.Polylines, %s)", layerIdent, layerIdent, targetIdent))
			}
			for _, elem := range layer.Markers {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Markers = append(%s.Markers, %s)", layerIdent, layerIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, mapview := range __gong__sortStageSetInstances(stage.MapViews, stage.MapView_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			mapviewIdent := mapview.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.MapView{Name: %s}).Stage(stage)", mapviewIdent, __gong__toRawStringLiteral(mapview.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", mapviewIdent, __gong__toRawStringLiteral(mapview.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StyleURL = %s", mapviewIdent, __gong__toRawStringLiteral(mapview.StyleURL)))
			values.WriteString(fmt.Sprintf("\n\t%s.Longitude = %f", mapviewIdent, mapview.Longitude))
			values.WriteString(fmt.Sprintf("\n\t%s.Latitude = %f", mapviewIdent, mapview.Latitude))
			values.WriteString(fmt.Sprintf("\n\t%s.Zoom = %f", mapviewIdent, mapview.Zoom))
			values.WriteString(fmt.Sprintf("\n\t%s.Pitch = %f", mapviewIdent, mapview.Pitch))
			values.WriteString(fmt.Sprintf("\n\t%s.Bearing = %f", mapviewIdent, mapview.Bearing))
			for _, elem := range mapview.Layers {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Layers = append(%s.Layers, %s)", mapviewIdent, mapviewIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, marker := range __gong__sortStageSetInstances(stage.Markers, stage.Marker_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			markerIdent := marker.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Marker{Name: %s}).Stage(stage)", markerIdent, __gong__toRawStringLiteral(marker.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", markerIdent, __gong__toRawStringLiteral(marker.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Label = %s", markerIdent, __gong__toRawStringLiteral(marker.Label)))
			values.WriteString(fmt.Sprintf("\n\t%s.PopupHTML = %s", markerIdent, __gong__toRawStringLiteral(marker.PopupHTML)))
			values.WriteString(fmt.Sprintf("\n\t%s.Longitude = %f", markerIdent, marker.Longitude))
			values.WriteString(fmt.Sprintf("\n\t%s.Latitude = %f", markerIdent, marker.Latitude))
			values.WriteString(fmt.Sprintf("\n\t%s.Altitude = %f", markerIdent, marker.Altitude))
			values.WriteString(fmt.Sprintf("\n\t%s.SVG = %s", markerIdent, __gong__toRawStringLiteral(marker.SVG)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", markerIdent, __gong__toRawStringLiteral(marker.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Scale = %f", markerIdent, marker.Scale))
			values.WriteString(fmt.Sprintf("\n\t%s.Rotation = %f", markerIdent, marker.Rotation))
			if marker.RotationAlignment.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.RotationAlignment = models.%s", markerIdent, marker.RotationAlignment.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.RotationAlignment = \"\"", markerIdent))
			}
			if marker.PitchAlignment.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.PitchAlignment = models.%s", markerIdent, marker.PitchAlignment.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.PitchAlignment = \"\"", markerIdent))
			}
			if marker.Anchor.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.Anchor = models.%s", markerIdent, marker.Anchor.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.Anchor = \"\"", markerIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.IsDraggable = %t", markerIdent, marker.IsDraggable))
		}
	}
	if stage != nil {
		for _, polygon := range __gong__sortStageSetInstances(stage.Polygons, stage.Polygon_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			polygonIdent := polygon.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Polygon{Name: %s}).Stage(stage)", polygonIdent, __gong__toRawStringLiteral(polygon.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", polygonIdent, __gong__toRawStringLiteral(polygon.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Label = %s", polygonIdent, __gong__toRawStringLiteral(polygon.Label)))
			values.WriteString(fmt.Sprintf("\n\t%s.PopupHTML = %s", polygonIdent, __gong__toRawStringLiteral(polygon.PopupHTML)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillColor = %s", polygonIdent, __gong__toRawStringLiteral(polygon.FillColor)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", polygonIdent, polygon.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeColor = %s", polygonIdent, __gong__toRawStringLiteral(polygon.StrokeColor)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", polygonIdent, polygon.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", polygonIdent, polygon.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExtruded = %t", polygonIdent, polygon.IsExtruded))
			values.WriteString(fmt.Sprintf("\n\t%s.BaseAltitude = %f", polygonIdent, polygon.BaseAltitude))
			values.WriteString(fmt.Sprintf("\n\t%s.ExtrusionHeight = %f", polygonIdent, polygon.ExtrusionHeight))
			for _, elem := range polygon.Coordinates {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Coordinates = append(%s.Coordinates, %s)", polygonIdent, polygonIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, polyline := range __gong__sortStageSetInstances(stage.Polylines, stage.Polyline_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			polylineIdent := polyline.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Polyline{Name: %s}).Stage(stage)", polylineIdent, __gong__toRawStringLiteral(polyline.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", polylineIdent, __gong__toRawStringLiteral(polyline.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Label = %s", polylineIdent, __gong__toRawStringLiteral(polyline.Label)))
			values.WriteString(fmt.Sprintf("\n\t%s.PopupHTML = %s", polylineIdent, __gong__toRawStringLiteral(polyline.PopupHTML)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeColor = %s", polylineIdent, __gong__toRawStringLiteral(polyline.StrokeColor)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", polylineIdent, polyline.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", polylineIdent, polyline.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", polylineIdent, __gong__toRawStringLiteral(polyline.StrokeDashArray)))
			for _, elem := range polyline.Coordinates {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Coordinates = append(%s.Coordinates, %s)", polylineIdent, polylineIdent, targetIdent))
			}
		}
	}

	var goModuleVersion string
	if buildInfo, ok := debug.ReadBuildInfo(); ok {
		goModuleVersion = buildInfo.Main.Version
	}

	goModuleVersionWithoutDirty := strings.ReplaceAll(goModuleVersion, "+dirty", "")
	if goModuleVersionWithoutDirty == "" || goModuleVersionWithoutDirty == "(devel)" {
		goModuleVersionWithoutDirty = "latest"
	}

	var metaImports strings.Builder
	var metaDummies strings.Builder
	if len(stage.MetaPackageImports) > 0 {
		for _, imp := range stage.MetaPackageImports {
			metaImports.WriteString(fmt.Sprintf("\n\t%s %s", imp.Alias, imp.Path))
			metaDummies.WriteString(fmt.Sprintf("\nvar _ %s.Stage", imp.Alias))
		}
	} else if stage.MetaPackageImportAlias != "" {
		metaImports.WriteString(fmt.Sprintf("\n\t%s %s", stage.MetaPackageImportAlias, stage.MetaPackageImportPath))
		metaDummies.WriteString(fmt.Sprintf("\nvar _ %s.Stage", stage.MetaPackageImportAlias))
	}

	res = fmt.Sprintf(`// file generated by gong
// go module version: %s
/* to edit this file with the default probe, use the command
go run github.com/fullstack-lang/gong/lib/maplibre/go/models/probe/cmd/probe@%s stage.go
*/
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/lib/maplibre/go/models"%s
)

var (
	_ time.Time
	_ = slices.Index[[]int, int]

	_ *models.Stage
)%s

// function will stage objects across all coordinated stages
func _(stage *models.Stage) {

	// ------------------------------------------------------------------------
	// Phase 1: Declarations (in topological order: leaves first)
	// ------------------------------------------------------------------------%s

	// ------------------------------------------------------------------------
	// Phase 2: Value Initializations
	// ------------------------------------------------------------------------%s

	// ------------------------------------------------------------------------
	// Phase 3: Pointer Setups (Intra-stage and Cross-stage pointers)
	// ------------------------------------------------------------------------%s
}
`, goModuleVersion, goModuleVersionWithoutDirty, packageName, metaImports.String(), metaDummies.String(), declarations.String(), values.String(), pointers.String())

	if modelsPackageName != "" && modelsPackageName != "github.com/fullstack-lang/gong/lib/maplibre/go/models" {
		res = strings.ReplaceAll(res, "\"github.com/fullstack-lang/gong/lib/maplibre/go/models\"", fmt.Sprintf("%q", modelsPackageName))
	}

	return res, nil
}

// insertion point for marshall field methods
func (coordinate *Coordinate) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := coordinate.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", coordinate.Name)
	case "Longitude":
		res = __gong__marshallFloat(ident, "Longitude", coordinate.Longitude)
	case "Latitude":
		res = __gong__marshallFloat(ident, "Latitude", coordinate.Latitude)
	case "Altitude":
		res = __gong__marshallFloat(ident, "Altitude", coordinate.Altitude)

	default:
		log.Panicf("Unknown field %s for Gongstruct Coordinate", fieldName)
	}
	return
}

func (layer *Layer) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := layer.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", layer.Name)
	case "IsVisible":
		res = __gong__marshallBool(ident, "IsVisible", layer.IsVisible)

	case "Polygons":
		var sb strings.Builder
		for _, _polygon := range layer.Polygons {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Polygons", _polygon.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Polylines":
		var sb strings.Builder
		for _, _polyline := range layer.Polylines {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Polylines", _polyline.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Markers":
		var sb strings.Builder
		for _, _marker := range layer.Markers {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Markers", _marker.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Layer", fieldName)
	}
	return
}

func (mapview *MapView) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := mapview.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", mapview.Name)
	case "StyleURL":
		res = __gong__marshallString(ident, "StyleURL", mapview.StyleURL)
	case "Longitude":
		res = __gong__marshallFloat(ident, "Longitude", mapview.Longitude)
	case "Latitude":
		res = __gong__marshallFloat(ident, "Latitude", mapview.Latitude)
	case "Zoom":
		res = __gong__marshallFloat(ident, "Zoom", mapview.Zoom)
	case "Pitch":
		res = __gong__marshallFloat(ident, "Pitch", mapview.Pitch)
	case "Bearing":
		res = __gong__marshallFloat(ident, "Bearing", mapview.Bearing)

	case "Layers":
		var sb strings.Builder
		for _, _layer := range mapview.Layers {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Layers", _layer.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct MapView", fieldName)
	}
	return
}

func (marker *Marker) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := marker.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", marker.Name)
	case "Label":
		res = __gong__marshallString(ident, "Label", marker.Label)
	case "PopupHTML":
		res = __gong__marshallString(ident, "PopupHTML", marker.PopupHTML)
	case "Longitude":
		res = __gong__marshallFloat(ident, "Longitude", marker.Longitude)
	case "Latitude":
		res = __gong__marshallFloat(ident, "Latitude", marker.Latitude)
	case "Altitude":
		res = __gong__marshallFloat(ident, "Altitude", marker.Altitude)
	case "SVG":
		res = __gong__marshallString(ident, "SVG", marker.SVG)
	case "Color":
		res = __gong__marshallString(ident, "Color", marker.Color)
	case "Scale":
		res = __gong__marshallFloat(ident, "Scale", marker.Scale)
	case "Rotation":
		res = __gong__marshallFloat(ident, "Rotation", marker.Rotation)
	case "RotationAlignment":
		res = __gong__marshallEnumString(ident, "RotationAlignment", marker.RotationAlignment.ToCodeString())
	case "PitchAlignment":
		res = __gong__marshallEnumString(ident, "PitchAlignment", marker.PitchAlignment.ToCodeString())
	case "Anchor":
		res = __gong__marshallEnumString(ident, "Anchor", marker.Anchor.ToCodeString())
	case "IsDraggable":
		res = __gong__marshallBool(ident, "IsDraggable", marker.IsDraggable)

	default:
		log.Panicf("Unknown field %s for Gongstruct Marker", fieldName)
	}
	return
}

func (polygon *Polygon) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := polygon.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", polygon.Name)
	case "Label":
		res = __gong__marshallString(ident, "Label", polygon.Label)
	case "PopupHTML":
		res = __gong__marshallString(ident, "PopupHTML", polygon.PopupHTML)
	case "FillColor":
		res = __gong__marshallString(ident, "FillColor", polygon.FillColor)
	case "FillOpacity":
		res = __gong__marshallFloat(ident, "FillOpacity", polygon.FillOpacity)
	case "StrokeColor":
		res = __gong__marshallString(ident, "StrokeColor", polygon.StrokeColor)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", polygon.StrokeOpacity)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", polygon.StrokeWidth)
	case "IsExtruded":
		res = __gong__marshallBool(ident, "IsExtruded", polygon.IsExtruded)
	case "BaseAltitude":
		res = __gong__marshallFloat(ident, "BaseAltitude", polygon.BaseAltitude)
	case "ExtrusionHeight":
		res = __gong__marshallFloat(ident, "ExtrusionHeight", polygon.ExtrusionHeight)

	case "Coordinates":
		var sb strings.Builder
		for _, _coordinate := range polygon.Coordinates {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Coordinates", _coordinate.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Polygon", fieldName)
	}
	return
}

func (polyline *Polyline) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := polyline.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", polyline.Name)
	case "Label":
		res = __gong__marshallString(ident, "Label", polyline.Label)
	case "PopupHTML":
		res = __gong__marshallString(ident, "PopupHTML", polyline.PopupHTML)
	case "StrokeColor":
		res = __gong__marshallString(ident, "StrokeColor", polyline.StrokeColor)
	case "StrokeOpacity":
		res = __gong__marshallFloat(ident, "StrokeOpacity", polyline.StrokeOpacity)
	case "StrokeWidth":
		res = __gong__marshallFloat(ident, "StrokeWidth", polyline.StrokeWidth)
	case "StrokeDashArray":
		res = __gong__marshallString(ident, "StrokeDashArray", polyline.StrokeDashArray)

	case "Coordinates":
		var sb strings.Builder
		for _, _coordinate := range polyline.Coordinates {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Coordinates", _coordinate.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Polyline", fieldName)
	}
	return
}

// insertion point for marshall all fields methods
func (coordinate *Coordinate) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(coordinate.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(coordinate.GongMarshallField(stage, "Longitude"))
		initializerStatements.WriteString(coordinate.GongMarshallField(stage, "Latitude"))
		initializerStatements.WriteString(coordinate.GongMarshallField(stage, "Altitude"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (layer *Layer) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(layer.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(layer.GongMarshallField(stage, "IsVisible"))
		pointersInitializesStatements.WriteString(layer.GongMarshallField(stage, "Polygons"))
		pointersInitializesStatements.WriteString(layer.GongMarshallField(stage, "Polylines"))
		pointersInitializesStatements.WriteString(layer.GongMarshallField(stage, "Markers"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (mapview *MapView) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(mapview.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(mapview.GongMarshallField(stage, "StyleURL"))
		initializerStatements.WriteString(mapview.GongMarshallField(stage, "Longitude"))
		initializerStatements.WriteString(mapview.GongMarshallField(stage, "Latitude"))
		initializerStatements.WriteString(mapview.GongMarshallField(stage, "Zoom"))
		initializerStatements.WriteString(mapview.GongMarshallField(stage, "Pitch"))
		initializerStatements.WriteString(mapview.GongMarshallField(stage, "Bearing"))
		pointersInitializesStatements.WriteString(mapview.GongMarshallField(stage, "Layers"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (marker *Marker) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(marker.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(marker.GongMarshallField(stage, "Label"))
		initializerStatements.WriteString(marker.GongMarshallField(stage, "PopupHTML"))
		initializerStatements.WriteString(marker.GongMarshallField(stage, "Longitude"))
		initializerStatements.WriteString(marker.GongMarshallField(stage, "Latitude"))
		initializerStatements.WriteString(marker.GongMarshallField(stage, "Altitude"))
		initializerStatements.WriteString(marker.GongMarshallField(stage, "SVG"))
		initializerStatements.WriteString(marker.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(marker.GongMarshallField(stage, "Scale"))
		initializerStatements.WriteString(marker.GongMarshallField(stage, "Rotation"))
		initializerStatements.WriteString(marker.GongMarshallField(stage, "RotationAlignment"))
		initializerStatements.WriteString(marker.GongMarshallField(stage, "PitchAlignment"))
		initializerStatements.WriteString(marker.GongMarshallField(stage, "Anchor"))
		initializerStatements.WriteString(marker.GongMarshallField(stage, "IsDraggable"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (polygon *Polygon) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(polygon.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(polygon.GongMarshallField(stage, "Label"))
		initializerStatements.WriteString(polygon.GongMarshallField(stage, "PopupHTML"))
		pointersInitializesStatements.WriteString(polygon.GongMarshallField(stage, "Coordinates"))
		initializerStatements.WriteString(polygon.GongMarshallField(stage, "FillColor"))
		initializerStatements.WriteString(polygon.GongMarshallField(stage, "FillOpacity"))
		initializerStatements.WriteString(polygon.GongMarshallField(stage, "StrokeColor"))
		initializerStatements.WriteString(polygon.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(polygon.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(polygon.GongMarshallField(stage, "IsExtruded"))
		initializerStatements.WriteString(polygon.GongMarshallField(stage, "BaseAltitude"))
		initializerStatements.WriteString(polygon.GongMarshallField(stage, "ExtrusionHeight"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (polyline *Polyline) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(polyline.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(polyline.GongMarshallField(stage, "Label"))
		initializerStatements.WriteString(polyline.GongMarshallField(stage, "PopupHTML"))
		pointersInitializesStatements.WriteString(polyline.GongMarshallField(stage, "Coordinates"))
		initializerStatements.WriteString(polyline.GongMarshallField(stage, "StrokeColor"))
		initializerStatements.WriteString(polyline.GongMarshallField(stage, "StrokeOpacity"))
		initializerStatements.WriteString(polyline.GongMarshallField(stage, "StrokeWidth"))
		initializerStatements.WriteString(polyline.GongMarshallField(stage, "StrokeDashArray"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
