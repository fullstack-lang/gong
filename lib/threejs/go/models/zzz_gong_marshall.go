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
	modelsPackageName := "github.com/fullstack-lang/gong/lib/threejs/go/models"
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
	localAliasPrefix := "__models"
	if stage.IsInDeltaMode() || stage.GetGongMarshallingMode() == GongMarshallingAppendCommit {
		localAliasPrefix = ""
	}
	_ = localAliasPrefix

	if stage != nil {
		for _, ambiantlight := range __gong__sortStageSetInstances(stage.AmbiantLights, stage.AmbiantLight_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			ambiantlightIdent := localAliasPrefix + ambiantlight.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.AmbiantLight{Name: %s}).Stage(stage)", ambiantlightIdent, __gong__toRawStringLiteral(ambiantlight.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", ambiantlightIdent, __gong__toRawStringLiteral(ambiantlight.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Intensity = %f", ambiantlightIdent, ambiantlight.Intensity))
		}
	}
	if stage != nil {
		for _, boxgeometry := range __gong__sortStageSetInstances(stage.BoxGeometrys, stage.BoxGeometry_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			boxgeometryIdent := localAliasPrefix + boxgeometry.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.BoxGeometry{Name: %s}).Stage(stage)", boxgeometryIdent, __gong__toRawStringLiteral(boxgeometry.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", boxgeometryIdent, __gong__toRawStringLiteral(boxgeometry.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", boxgeometryIdent, boxgeometry.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", boxgeometryIdent, boxgeometry.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.Depth = %f", boxgeometryIdent, boxgeometry.Depth))
			values.WriteString(fmt.Sprintf("\n\t%s.WidthSegments = %d", boxgeometryIdent, boxgeometry.WidthSegments))
			values.WriteString(fmt.Sprintf("\n\t%s.HeightSegments = %d", boxgeometryIdent, boxgeometry.HeightSegments))
			values.WriteString(fmt.Sprintf("\n\t%s.DepthSegments = %d", boxgeometryIdent, boxgeometry.DepthSegments))
		}
	}
	if stage != nil {
		for _, buffergeometry := range __gong__sortStageSetInstances(stage.BufferGeometrys, stage.BufferGeometry_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			buffergeometryIdent := localAliasPrefix + buffergeometry.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.BufferGeometry{Name: %s}).Stage(stage)", buffergeometryIdent, __gong__toRawStringLiteral(buffergeometry.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", buffergeometryIdent, __gong__toRawStringLiteral(buffergeometry.Name)))
			for _, elem := range buffergeometry.Vertices {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Vertices = append(%s.Vertices, %s)", buffergeometryIdent, buffergeometryIdent, targetIdent))
			}
			for _, elem := range buffergeometry.Faces {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Faces = append(%s.Faces, %s)", buffergeometryIdent, buffergeometryIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, camera := range __gong__sortStageSetInstances(stage.Cameras, stage.Camera_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			cameraIdent := localAliasPrefix + camera.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Camera{Name: %s}).Stage(stage)", cameraIdent, __gong__toRawStringLiteral(camera.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", cameraIdent, __gong__toRawStringLiteral(camera.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", cameraIdent, camera.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", cameraIdent, camera.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Z = %f", cameraIdent, camera.Z))
			values.WriteString(fmt.Sprintf("\n\t%s.TargetX = %f", cameraIdent, camera.TargetX))
			values.WriteString(fmt.Sprintf("\n\t%s.TargetY = %f", cameraIdent, camera.TargetY))
			values.WriteString(fmt.Sprintf("\n\t%s.TargetZ = %f", cameraIdent, camera.TargetZ))
			values.WriteString(fmt.Sprintf("\n\t%s.Fov = %f", cameraIdent, camera.Fov))
		}
	}
	if stage != nil {
		for _, canvas := range __gong__sortStageSetInstances(stage.Canvass, stage.Canvas_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			canvasIdent := localAliasPrefix + canvas.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Canvas{Name: %s}).Stage(stage)", canvasIdent, __gong__toRawStringLiteral(canvas.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", canvasIdent, __gong__toRawStringLiteral(canvas.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsWithLastRenderingUpdate = %t", canvasIdent, canvas.IsWithLastRenderingUpdate))
			values.WriteString(fmt.Sprintf("\n\t%s.LastRendering, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", canvasIdent, canvas.LastRendering.String()))
			values.WriteString(fmt.Sprintf("\n\t%s.Frame64BitsEncoded = %s", canvasIdent, __gong__toRawStringLiteral(canvas.Frame64BitsEncoded)))
			for _, elem := range canvas.DirectionalLights {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DirectionalLights = append(%s.DirectionalLights, %s)", canvasIdent, canvasIdent, targetIdent))
			}
			if canvas.AmbiantLight != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + canvas.AmbiantLight.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AmbiantLight = %s", canvasIdent, targetIdent))
			}
			for _, elem := range canvas.Meshs {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Meshs = append(%s.Meshs, %s)", canvasIdent, canvasIdent, targetIdent))
			}
			if canvas.Camera != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + canvas.Camera.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Camera = %s", canvasIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, curve := range __gong__sortStageSetInstances(stage.Curves, stage.Curve_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			curveIdent := localAliasPrefix + curve.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Curve{Name: %s}).Stage(stage)", curveIdent, __gong__toRawStringLiteral(curve.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", curveIdent, __gong__toRawStringLiteral(curve.Name)))
			for _, elem := range curve.Points {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Points = append(%s.Points, %s)", curveIdent, curveIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, cylindergeometry := range __gong__sortStageSetInstances(stage.CylinderGeometrys, stage.CylinderGeometry_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			cylindergeometryIdent := localAliasPrefix + cylindergeometry.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.CylinderGeometry{Name: %s}).Stage(stage)", cylindergeometryIdent, __gong__toRawStringLiteral(cylindergeometry.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", cylindergeometryIdent, __gong__toRawStringLiteral(cylindergeometry.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.RadiusTop = %f", cylindergeometryIdent, cylindergeometry.RadiusTop))
			values.WriteString(fmt.Sprintf("\n\t%s.RadiusBottom = %f", cylindergeometryIdent, cylindergeometry.RadiusBottom))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", cylindergeometryIdent, cylindergeometry.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.RadialSegments = %d", cylindergeometryIdent, cylindergeometry.RadialSegments))
			values.WriteString(fmt.Sprintf("\n\t%s.HeightSegments = %d", cylindergeometryIdent, cylindergeometry.HeightSegments))
			values.WriteString(fmt.Sprintf("\n\t%s.OpenEnded = %t", cylindergeometryIdent, cylindergeometry.OpenEnded))
			values.WriteString(fmt.Sprintf("\n\t%s.ThetaStart = %f", cylindergeometryIdent, cylindergeometry.ThetaStart))
			values.WriteString(fmt.Sprintf("\n\t%s.ThetaLength = %f", cylindergeometryIdent, cylindergeometry.ThetaLength))
		}
	}
	if stage != nil {
		for _, directionallight := range __gong__sortStageSetInstances(stage.DirectionalLights, stage.DirectionalLight_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			directionallightIdent := localAliasPrefix + directionallight.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.DirectionalLight{Name: %s}).Stage(stage)", directionallightIdent, __gong__toRawStringLiteral(directionallight.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", directionallightIdent, __gong__toRawStringLiteral(directionallight.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", directionallightIdent, directionallight.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", directionallightIdent, directionallight.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Z = %f", directionallightIdent, directionallight.Z))
			values.WriteString(fmt.Sprintf("\n\t%s.Intensity = %f", directionallightIdent, directionallight.Intensity))
			values.WriteString(fmt.Sprintf("\n\t%s.IsWithCastShadow = %t", directionallightIdent, directionallight.IsWithCastShadow))
		}
	}
	if stage != nil {
		for _, extrudegeometry := range __gong__sortStageSetInstances(stage.ExtrudeGeometrys, stage.ExtrudeGeometry_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			extrudegeometryIdent := localAliasPrefix + extrudegeometry.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ExtrudeGeometry{Name: %s}).Stage(stage)", extrudegeometryIdent, __gong__toRawStringLiteral(extrudegeometry.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", extrudegeometryIdent, __gong__toRawStringLiteral(extrudegeometry.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Steps = %d", extrudegeometryIdent, extrudegeometry.Steps))
			if extrudegeometry.Shape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + extrudegeometry.Shape.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Shape = %s", extrudegeometryIdent, targetIdent))
			}
			if extrudegeometry.ExtrudePath != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + extrudegeometry.ExtrudePath.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ExtrudePath = %s", extrudegeometryIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, mesh := range __gong__sortStageSetInstances(stage.Meshs, stage.Mesh_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			meshIdent := localAliasPrefix + mesh.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Mesh{Name: %s}).Stage(stage)", meshIdent, __gong__toRawStringLiteral(mesh.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", meshIdent, __gong__toRawStringLiteral(mesh.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", meshIdent, mesh.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", meshIdent, mesh.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Z = %f", meshIdent, mesh.Z))
			if mesh.MeshMaterialBasic != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + mesh.MeshMaterialBasic.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MeshMaterialBasic = %s", meshIdent, targetIdent))
			}
			if mesh.MeshPhysicalMaterial != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + mesh.MeshPhysicalMaterial.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MeshPhysicalMaterial = %s", meshIdent, targetIdent))
			}
			if mesh.CylinderGeometry != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + mesh.CylinderGeometry.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CylinderGeometry = %s", meshIdent, targetIdent))
			}
			if mesh.BoxGeometry != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + mesh.BoxGeometry.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BoxGeometry = %s", meshIdent, targetIdent))
			}
			if mesh.SphereGeometry != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + mesh.SphereGeometry.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SphereGeometry = %s", meshIdent, targetIdent))
			}
			if mesh.TorusGeometry != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + mesh.TorusGeometry.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TorusGeometry = %s", meshIdent, targetIdent))
			}
			if mesh.PlaneGeometry != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + mesh.PlaneGeometry.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PlaneGeometry = %s", meshIdent, targetIdent))
			}
			if mesh.TubeGeometry != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + mesh.TubeGeometry.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TubeGeometry = %s", meshIdent, targetIdent))
			}
			if mesh.ExtrudeGeometry != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + mesh.ExtrudeGeometry.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ExtrudeGeometry = %s", meshIdent, targetIdent))
			}
			if mesh.BufferGeometry != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + mesh.BufferGeometry.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BufferGeometry = %s", meshIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, meshmaterialbasic := range __gong__sortStageSetInstances(stage.MeshMaterialBasics, stage.MeshMaterialBasic_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			meshmaterialbasicIdent := localAliasPrefix + meshmaterialbasic.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.MeshMaterialBasic{Name: %s}).Stage(stage)", meshmaterialbasicIdent, __gong__toRawStringLiteral(meshmaterialbasic.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", meshmaterialbasicIdent, __gong__toRawStringLiteral(meshmaterialbasic.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", meshmaterialbasicIdent, __gong__toRawStringLiteral(meshmaterialbasic.Color)))
		}
	}
	if stage != nil {
		for _, meshphysicalmaterial := range __gong__sortStageSetInstances(stage.MeshPhysicalMaterials, stage.MeshPhysicalMaterial_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			meshphysicalmaterialIdent := localAliasPrefix + meshphysicalmaterial.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.MeshPhysicalMaterial{Name: %s}).Stage(stage)", meshphysicalmaterialIdent, __gong__toRawStringLiteral(meshphysicalmaterial.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", meshphysicalmaterialIdent, __gong__toRawStringLiteral(meshphysicalmaterial.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", meshphysicalmaterialIdent, __gong__toRawStringLiteral(meshphysicalmaterial.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Wireframe = %t", meshphysicalmaterialIdent, meshphysicalmaterial.Wireframe))
			values.WriteString(fmt.Sprintf("\n\t%s.Opacity = %f", meshphysicalmaterialIdent, meshphysicalmaterial.Opacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Transparent = %t", meshphysicalmaterialIdent, meshphysicalmaterial.Transparent))
			values.WriteString(fmt.Sprintf("\n\t%s.Visible = %t", meshphysicalmaterialIdent, meshphysicalmaterial.Visible))
		}
	}
	if stage != nil {
		for _, planegeometry := range __gong__sortStageSetInstances(stage.PlaneGeometrys, stage.PlaneGeometry_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			planegeometryIdent := localAliasPrefix + planegeometry.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.PlaneGeometry{Name: %s}).Stage(stage)", planegeometryIdent, __gong__toRawStringLiteral(planegeometry.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", planegeometryIdent, __gong__toRawStringLiteral(planegeometry.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", planegeometryIdent, planegeometry.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", planegeometryIdent, planegeometry.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.WidthSegments = %d", planegeometryIdent, planegeometry.WidthSegments))
			values.WriteString(fmt.Sprintf("\n\t%s.HeightSegments = %d", planegeometryIdent, planegeometry.HeightSegments))
		}
	}
	if stage != nil {
		for _, shape := range __gong__sortStageSetInstances(stage.Shapes, stage.Shape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			shapeIdent := localAliasPrefix + shape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Shape{Name: %s}).Stage(stage)", shapeIdent, __gong__toRawStringLiteral(shape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", shapeIdent, __gong__toRawStringLiteral(shape.Name)))
			for _, elem := range shape.Points {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Points = append(%s.Points, %s)", shapeIdent, shapeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, spheregeometry := range __gong__sortStageSetInstances(stage.SphereGeometrys, stage.SphereGeometry_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			spheregeometryIdent := localAliasPrefix + spheregeometry.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.SphereGeometry{Name: %s}).Stage(stage)", spheregeometryIdent, __gong__toRawStringLiteral(spheregeometry.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", spheregeometryIdent, __gong__toRawStringLiteral(spheregeometry.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Radius = %f", spheregeometryIdent, spheregeometry.Radius))
			values.WriteString(fmt.Sprintf("\n\t%s.WidthSegments = %d", spheregeometryIdent, spheregeometry.WidthSegments))
			values.WriteString(fmt.Sprintf("\n\t%s.HeightSegments = %d", spheregeometryIdent, spheregeometry.HeightSegments))
			values.WriteString(fmt.Sprintf("\n\t%s.PhiStart = %f", spheregeometryIdent, spheregeometry.PhiStart))
			values.WriteString(fmt.Sprintf("\n\t%s.PhiLength = %f", spheregeometryIdent, spheregeometry.PhiLength))
			values.WriteString(fmt.Sprintf("\n\t%s.ThetaStart = %f", spheregeometryIdent, spheregeometry.ThetaStart))
			values.WriteString(fmt.Sprintf("\n\t%s.ThetaLength = %f", spheregeometryIdent, spheregeometry.ThetaLength))
		}
	}
	if stage != nil {
		for _, torusgeometry := range __gong__sortStageSetInstances(stage.TorusGeometrys, stage.TorusGeometry_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			torusgeometryIdent := localAliasPrefix + torusgeometry.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.TorusGeometry{Name: %s}).Stage(stage)", torusgeometryIdent, __gong__toRawStringLiteral(torusgeometry.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", torusgeometryIdent, __gong__toRawStringLiteral(torusgeometry.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Radius = %f", torusgeometryIdent, torusgeometry.Radius))
			values.WriteString(fmt.Sprintf("\n\t%s.Tube = %f", torusgeometryIdent, torusgeometry.Tube))
			values.WriteString(fmt.Sprintf("\n\t%s.RadialSegments = %d", torusgeometryIdent, torusgeometry.RadialSegments))
			values.WriteString(fmt.Sprintf("\n\t%s.TubularSegments = %d", torusgeometryIdent, torusgeometry.TubularSegments))
			values.WriteString(fmt.Sprintf("\n\t%s.Arc = %f", torusgeometryIdent, torusgeometry.Arc))
		}
	}
	if stage != nil {
		for _, triangle := range __gong__sortStageSetInstances(stage.Triangles, stage.Triangle_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			triangleIdent := localAliasPrefix + triangle.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Triangle{Name: %s}).Stage(stage)", triangleIdent, __gong__toRawStringLiteral(triangle.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", triangleIdent, __gong__toRawStringLiteral(triangle.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.V1 = %d", triangleIdent, triangle.V1))
			values.WriteString(fmt.Sprintf("\n\t%s.V2 = %d", triangleIdent, triangle.V2))
			values.WriteString(fmt.Sprintf("\n\t%s.V3 = %d", triangleIdent, triangle.V3))
		}
	}
	if stage != nil {
		for _, tubegeometry := range __gong__sortStageSetInstances(stage.TubeGeometrys, stage.TubeGeometry_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			tubegeometryIdent := localAliasPrefix + tubegeometry.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.TubeGeometry{Name: %s}).Stage(stage)", tubegeometryIdent, __gong__toRawStringLiteral(tubegeometry.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", tubegeometryIdent, __gong__toRawStringLiteral(tubegeometry.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.TubularSegments = %d", tubegeometryIdent, tubegeometry.TubularSegments))
			values.WriteString(fmt.Sprintf("\n\t%s.Radius = %f", tubegeometryIdent, tubegeometry.Radius))
			values.WriteString(fmt.Sprintf("\n\t%s.RadialSegments = %d", tubegeometryIdent, tubegeometry.RadialSegments))
			values.WriteString(fmt.Sprintf("\n\t%s.Closed = %t", tubegeometryIdent, tubegeometry.Closed))
			if tubegeometry.Path != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + tubegeometry.Path.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Path = %s", tubegeometryIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, vector2 := range __gong__sortStageSetInstances(stage.Vector2s, stage.Vector2_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			vector2Ident := localAliasPrefix + vector2.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Vector2{Name: %s}).Stage(stage)", vector2Ident, __gong__toRawStringLiteral(vector2.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", vector2Ident, __gong__toRawStringLiteral(vector2.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", vector2Ident, vector2.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", vector2Ident, vector2.Y))
		}
	}
	if stage != nil {
		for _, vector3 := range __gong__sortStageSetInstances(stage.Vector3s, stage.Vector3_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			vector3Ident := localAliasPrefix + vector3.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Vector3{Name: %s}).Stage(stage)", vector3Ident, __gong__toRawStringLiteral(vector3.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", vector3Ident, __gong__toRawStringLiteral(vector3.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", vector3Ident, vector3.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", vector3Ident, vector3.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Z = %f", vector3Ident, vector3.Z))
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
go run github.com/fullstack-lang/gong/lib/threejs/go/models/probe/cmd/probe@%s stage.go
*/
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/lib/threejs/go/models"%s
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

	if modelsPackageName != "" && modelsPackageName != "github.com/fullstack-lang/gong/lib/threejs/go/models" {
		res = strings.ReplaceAll(res, "\"github.com/fullstack-lang/gong/lib/threejs/go/models\"", fmt.Sprintf("%q", modelsPackageName))
	}

	return res, nil
}

// insertion point for marshall field methods
func (ambiantlight *AmbiantLight) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := ambiantlight.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", ambiantlight.Name)
	case "Intensity":
		res = __gong__marshallFloat(ident, "Intensity", ambiantlight.Intensity)

	default:
		log.Panicf("Unknown field %s for Gongstruct AmbiantLight", fieldName)
	}
	return
}

func (boxgeometry *BoxGeometry) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := boxgeometry.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", boxgeometry.Name)
	case "Width":
		res = __gong__marshallFloat(ident, "Width", boxgeometry.Width)
	case "Height":
		res = __gong__marshallFloat(ident, "Height", boxgeometry.Height)
	case "Depth":
		res = __gong__marshallFloat(ident, "Depth", boxgeometry.Depth)
	case "WidthSegments":
		res = __gong__marshallInt(ident, "WidthSegments", boxgeometry.WidthSegments)
	case "HeightSegments":
		res = __gong__marshallInt(ident, "HeightSegments", boxgeometry.HeightSegments)
	case "DepthSegments":
		res = __gong__marshallInt(ident, "DepthSegments", boxgeometry.DepthSegments)

	default:
		log.Panicf("Unknown field %s for Gongstruct BoxGeometry", fieldName)
	}
	return
}

func (buffergeometry *BufferGeometry) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := buffergeometry.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", buffergeometry.Name)

	case "Vertices":
		var sb strings.Builder
		for _, _vector3 := range buffergeometry.Vertices {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Vertices", _vector3.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Faces":
		var sb strings.Builder
		for _, _triangle := range buffergeometry.Faces {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Faces", _triangle.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct BufferGeometry", fieldName)
	}
	return
}

func (camera *Camera) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := camera.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", camera.Name)
	case "X":
		res = __gong__marshallFloat(ident, "X", camera.X)
	case "Y":
		res = __gong__marshallFloat(ident, "Y", camera.Y)
	case "Z":
		res = __gong__marshallFloat(ident, "Z", camera.Z)
	case "TargetX":
		res = __gong__marshallFloat(ident, "TargetX", camera.TargetX)
	case "TargetY":
		res = __gong__marshallFloat(ident, "TargetY", camera.TargetY)
	case "TargetZ":
		res = __gong__marshallFloat(ident, "TargetZ", camera.TargetZ)
	case "Fov":
		res = __gong__marshallFloat(ident, "Fov", camera.Fov)

	default:
		log.Panicf("Unknown field %s for Gongstruct Camera", fieldName)
	}
	return
}

func (canvas *Canvas) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := canvas.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", canvas.Name)
	case "IsWithLastRenderingUpdate":
		res = __gong__marshallBool(ident, "IsWithLastRenderingUpdate", canvas.IsWithLastRenderingUpdate)
	case "LastRendering":
		res = __gong__marshallTime(ident, "LastRendering", canvas.LastRendering.String())
	case "Frame64BitsEncoded":
		res = __gong__marshallString(ident, "Frame64BitsEncoded", canvas.Frame64BitsEncoded)

	case "DirectionalLights":
		var sb strings.Builder
		for _, _directionallight := range canvas.DirectionalLights {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "DirectionalLights", _directionallight.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "AmbiantLight":
		if canvas.AmbiantLight != nil {
			res = __gong__marshallPointer(ident, "AmbiantLight", canvas.AmbiantLight.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "AmbiantLight", "nil")
		}
	case "Meshs":
		var sb strings.Builder
		for _, _mesh := range canvas.Meshs {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Meshs", _mesh.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Camera":
		if canvas.Camera != nil {
			res = __gong__marshallPointer(ident, "Camera", canvas.Camera.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Camera", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Canvas", fieldName)
	}
	return
}

func (curve *Curve) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := curve.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", curve.Name)

	case "Points":
		var sb strings.Builder
		for _, _vector3 := range curve.Points {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Points", _vector3.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Curve", fieldName)
	}
	return
}

func (cylindergeometry *CylinderGeometry) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := cylindergeometry.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", cylindergeometry.Name)
	case "RadiusTop":
		res = __gong__marshallFloat(ident, "RadiusTop", cylindergeometry.RadiusTop)
	case "RadiusBottom":
		res = __gong__marshallFloat(ident, "RadiusBottom", cylindergeometry.RadiusBottom)
	case "Height":
		res = __gong__marshallFloat(ident, "Height", cylindergeometry.Height)
	case "RadialSegments":
		res = __gong__marshallInt(ident, "RadialSegments", cylindergeometry.RadialSegments)
	case "HeightSegments":
		res = __gong__marshallInt(ident, "HeightSegments", cylindergeometry.HeightSegments)
	case "OpenEnded":
		res = __gong__marshallBool(ident, "OpenEnded", cylindergeometry.OpenEnded)
	case "ThetaStart":
		res = __gong__marshallFloat(ident, "ThetaStart", cylindergeometry.ThetaStart)
	case "ThetaLength":
		res = __gong__marshallFloat(ident, "ThetaLength", cylindergeometry.ThetaLength)

	default:
		log.Panicf("Unknown field %s for Gongstruct CylinderGeometry", fieldName)
	}
	return
}

func (directionallight *DirectionalLight) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := directionallight.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", directionallight.Name)
	case "X":
		res = __gong__marshallFloat(ident, "X", directionallight.X)
	case "Y":
		res = __gong__marshallFloat(ident, "Y", directionallight.Y)
	case "Z":
		res = __gong__marshallFloat(ident, "Z", directionallight.Z)
	case "Intensity":
		res = __gong__marshallFloat(ident, "Intensity", directionallight.Intensity)
	case "IsWithCastShadow":
		res = __gong__marshallBool(ident, "IsWithCastShadow", directionallight.IsWithCastShadow)

	default:
		log.Panicf("Unknown field %s for Gongstruct DirectionalLight", fieldName)
	}
	return
}

func (extrudegeometry *ExtrudeGeometry) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := extrudegeometry.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", extrudegeometry.Name)
	case "Steps":
		res = __gong__marshallInt(ident, "Steps", extrudegeometry.Steps)

	case "Shape":
		if extrudegeometry.Shape != nil {
			res = __gong__marshallPointer(ident, "Shape", extrudegeometry.Shape.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Shape", "nil")
		}
	case "ExtrudePath":
		if extrudegeometry.ExtrudePath != nil {
			res = __gong__marshallPointer(ident, "ExtrudePath", extrudegeometry.ExtrudePath.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ExtrudePath", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ExtrudeGeometry", fieldName)
	}
	return
}

func (mesh *Mesh) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := mesh.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", mesh.Name)
	case "X":
		res = __gong__marshallFloat(ident, "X", mesh.X)
	case "Y":
		res = __gong__marshallFloat(ident, "Y", mesh.Y)
	case "Z":
		res = __gong__marshallFloat(ident, "Z", mesh.Z)

	case "MeshMaterialBasic":
		if mesh.MeshMaterialBasic != nil {
			res = __gong__marshallPointer(ident, "MeshMaterialBasic", mesh.MeshMaterialBasic.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "MeshMaterialBasic", "nil")
		}
	case "MeshPhysicalMaterial":
		if mesh.MeshPhysicalMaterial != nil {
			res = __gong__marshallPointer(ident, "MeshPhysicalMaterial", mesh.MeshPhysicalMaterial.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "MeshPhysicalMaterial", "nil")
		}
	case "CylinderGeometry":
		if mesh.CylinderGeometry != nil {
			res = __gong__marshallPointer(ident, "CylinderGeometry", mesh.CylinderGeometry.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "CylinderGeometry", "nil")
		}
	case "BoxGeometry":
		if mesh.BoxGeometry != nil {
			res = __gong__marshallPointer(ident, "BoxGeometry", mesh.BoxGeometry.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "BoxGeometry", "nil")
		}
	case "SphereGeometry":
		if mesh.SphereGeometry != nil {
			res = __gong__marshallPointer(ident, "SphereGeometry", mesh.SphereGeometry.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "SphereGeometry", "nil")
		}
	case "TorusGeometry":
		if mesh.TorusGeometry != nil {
			res = __gong__marshallPointer(ident, "TorusGeometry", mesh.TorusGeometry.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "TorusGeometry", "nil")
		}
	case "PlaneGeometry":
		if mesh.PlaneGeometry != nil {
			res = __gong__marshallPointer(ident, "PlaneGeometry", mesh.PlaneGeometry.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "PlaneGeometry", "nil")
		}
	case "TubeGeometry":
		if mesh.TubeGeometry != nil {
			res = __gong__marshallPointer(ident, "TubeGeometry", mesh.TubeGeometry.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "TubeGeometry", "nil")
		}
	case "ExtrudeGeometry":
		if mesh.ExtrudeGeometry != nil {
			res = __gong__marshallPointer(ident, "ExtrudeGeometry", mesh.ExtrudeGeometry.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ExtrudeGeometry", "nil")
		}
	case "BufferGeometry":
		if mesh.BufferGeometry != nil {
			res = __gong__marshallPointer(ident, "BufferGeometry", mesh.BufferGeometry.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "BufferGeometry", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Mesh", fieldName)
	}
	return
}

func (meshmaterialbasic *MeshMaterialBasic) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := meshmaterialbasic.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", meshmaterialbasic.Name)
	case "Color":
		res = __gong__marshallString(ident, "Color", meshmaterialbasic.Color)

	default:
		log.Panicf("Unknown field %s for Gongstruct MeshMaterialBasic", fieldName)
	}
	return
}

func (meshphysicalmaterial *MeshPhysicalMaterial) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := meshphysicalmaterial.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", meshphysicalmaterial.Name)
	case "Color":
		res = __gong__marshallString(ident, "Color", meshphysicalmaterial.Color)
	case "Wireframe":
		res = __gong__marshallBool(ident, "Wireframe", meshphysicalmaterial.Wireframe)
	case "Opacity":
		res = __gong__marshallFloat(ident, "Opacity", meshphysicalmaterial.Opacity)
	case "Transparent":
		res = __gong__marshallBool(ident, "Transparent", meshphysicalmaterial.Transparent)
	case "Visible":
		res = __gong__marshallBool(ident, "Visible", meshphysicalmaterial.Visible)

	default:
		log.Panicf("Unknown field %s for Gongstruct MeshPhysicalMaterial", fieldName)
	}
	return
}

func (planegeometry *PlaneGeometry) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := planegeometry.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", planegeometry.Name)
	case "Width":
		res = __gong__marshallFloat(ident, "Width", planegeometry.Width)
	case "Height":
		res = __gong__marshallFloat(ident, "Height", planegeometry.Height)
	case "WidthSegments":
		res = __gong__marshallInt(ident, "WidthSegments", planegeometry.WidthSegments)
	case "HeightSegments":
		res = __gong__marshallInt(ident, "HeightSegments", planegeometry.HeightSegments)

	default:
		log.Panicf("Unknown field %s for Gongstruct PlaneGeometry", fieldName)
	}
	return
}

func (shape *Shape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := shape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", shape.Name)

	case "Points":
		var sb strings.Builder
		for _, _vector2 := range shape.Points {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Points", _vector2.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Shape", fieldName)
	}
	return
}

func (spheregeometry *SphereGeometry) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := spheregeometry.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", spheregeometry.Name)
	case "Radius":
		res = __gong__marshallFloat(ident, "Radius", spheregeometry.Radius)
	case "WidthSegments":
		res = __gong__marshallInt(ident, "WidthSegments", spheregeometry.WidthSegments)
	case "HeightSegments":
		res = __gong__marshallInt(ident, "HeightSegments", spheregeometry.HeightSegments)
	case "PhiStart":
		res = __gong__marshallFloat(ident, "PhiStart", spheregeometry.PhiStart)
	case "PhiLength":
		res = __gong__marshallFloat(ident, "PhiLength", spheregeometry.PhiLength)
	case "ThetaStart":
		res = __gong__marshallFloat(ident, "ThetaStart", spheregeometry.ThetaStart)
	case "ThetaLength":
		res = __gong__marshallFloat(ident, "ThetaLength", spheregeometry.ThetaLength)

	default:
		log.Panicf("Unknown field %s for Gongstruct SphereGeometry", fieldName)
	}
	return
}

func (torusgeometry *TorusGeometry) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := torusgeometry.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", torusgeometry.Name)
	case "Radius":
		res = __gong__marshallFloat(ident, "Radius", torusgeometry.Radius)
	case "Tube":
		res = __gong__marshallFloat(ident, "Tube", torusgeometry.Tube)
	case "RadialSegments":
		res = __gong__marshallInt(ident, "RadialSegments", torusgeometry.RadialSegments)
	case "TubularSegments":
		res = __gong__marshallInt(ident, "TubularSegments", torusgeometry.TubularSegments)
	case "Arc":
		res = __gong__marshallFloat(ident, "Arc", torusgeometry.Arc)

	default:
		log.Panicf("Unknown field %s for Gongstruct TorusGeometry", fieldName)
	}
	return
}

func (triangle *Triangle) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := triangle.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", triangle.Name)
	case "V1":
		res = __gong__marshallInt(ident, "V1", triangle.V1)
	case "V2":
		res = __gong__marshallInt(ident, "V2", triangle.V2)
	case "V3":
		res = __gong__marshallInt(ident, "V3", triangle.V3)

	default:
		log.Panicf("Unknown field %s for Gongstruct Triangle", fieldName)
	}
	return
}

func (tubegeometry *TubeGeometry) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := tubegeometry.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", tubegeometry.Name)
	case "TubularSegments":
		res = __gong__marshallInt(ident, "TubularSegments", tubegeometry.TubularSegments)
	case "Radius":
		res = __gong__marshallFloat(ident, "Radius", tubegeometry.Radius)
	case "RadialSegments":
		res = __gong__marshallInt(ident, "RadialSegments", tubegeometry.RadialSegments)
	case "Closed":
		res = __gong__marshallBool(ident, "Closed", tubegeometry.Closed)

	case "Path":
		if tubegeometry.Path != nil {
			res = __gong__marshallPointer(ident, "Path", tubegeometry.Path.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Path", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct TubeGeometry", fieldName)
	}
	return
}

func (vector2 *Vector2) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := vector2.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", vector2.Name)
	case "X":
		res = __gong__marshallFloat(ident, "X", vector2.X)
	case "Y":
		res = __gong__marshallFloat(ident, "Y", vector2.Y)

	default:
		log.Panicf("Unknown field %s for Gongstruct Vector2", fieldName)
	}
	return
}

func (vector3 *Vector3) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := vector3.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", vector3.Name)
	case "X":
		res = __gong__marshallFloat(ident, "X", vector3.X)
	case "Y":
		res = __gong__marshallFloat(ident, "Y", vector3.Y)
	case "Z":
		res = __gong__marshallFloat(ident, "Z", vector3.Z)

	default:
		log.Panicf("Unknown field %s for Gongstruct Vector3", fieldName)
	}
	return
}

// insertion point for marshall all fields methods
func (ambiantlight *AmbiantLight) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(ambiantlight.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(ambiantlight.GongMarshallField(stage, "Intensity"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (boxgeometry *BoxGeometry) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(boxgeometry.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(boxgeometry.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(boxgeometry.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(boxgeometry.GongMarshallField(stage, "Depth"))
		initializerStatements.WriteString(boxgeometry.GongMarshallField(stage, "WidthSegments"))
		initializerStatements.WriteString(boxgeometry.GongMarshallField(stage, "HeightSegments"))
		initializerStatements.WriteString(boxgeometry.GongMarshallField(stage, "DepthSegments"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (buffergeometry *BufferGeometry) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(buffergeometry.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(buffergeometry.GongMarshallField(stage, "Vertices"))
		pointersInitializesStatements.WriteString(buffergeometry.GongMarshallField(stage, "Faces"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (camera *Camera) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(camera.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(camera.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(camera.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(camera.GongMarshallField(stage, "Z"))
		initializerStatements.WriteString(camera.GongMarshallField(stage, "TargetX"))
		initializerStatements.WriteString(camera.GongMarshallField(stage, "TargetY"))
		initializerStatements.WriteString(camera.GongMarshallField(stage, "TargetZ"))
		initializerStatements.WriteString(camera.GongMarshallField(stage, "Fov"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (canvas *Canvas) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(canvas.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(canvas.GongMarshallField(stage, "DirectionalLights"))
		pointersInitializesStatements.WriteString(canvas.GongMarshallField(stage, "AmbiantLight"))
		pointersInitializesStatements.WriteString(canvas.GongMarshallField(stage, "Meshs"))
		pointersInitializesStatements.WriteString(canvas.GongMarshallField(stage, "Camera"))
		initializerStatements.WriteString(canvas.GongMarshallField(stage, "IsWithLastRenderingUpdate"))
		initializerStatements.WriteString(canvas.GongMarshallField(stage, "LastRendering"))
		initializerStatements.WriteString(canvas.GongMarshallField(stage, "Frame64BitsEncoded"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (curve *Curve) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(curve.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(curve.GongMarshallField(stage, "Points"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (cylindergeometry *CylinderGeometry) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(cylindergeometry.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(cylindergeometry.GongMarshallField(stage, "RadiusTop"))
		initializerStatements.WriteString(cylindergeometry.GongMarshallField(stage, "RadiusBottom"))
		initializerStatements.WriteString(cylindergeometry.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(cylindergeometry.GongMarshallField(stage, "RadialSegments"))
		initializerStatements.WriteString(cylindergeometry.GongMarshallField(stage, "HeightSegments"))
		initializerStatements.WriteString(cylindergeometry.GongMarshallField(stage, "OpenEnded"))
		initializerStatements.WriteString(cylindergeometry.GongMarshallField(stage, "ThetaStart"))
		initializerStatements.WriteString(cylindergeometry.GongMarshallField(stage, "ThetaLength"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (directionallight *DirectionalLight) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(directionallight.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(directionallight.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(directionallight.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(directionallight.GongMarshallField(stage, "Z"))
		initializerStatements.WriteString(directionallight.GongMarshallField(stage, "Intensity"))
		initializerStatements.WriteString(directionallight.GongMarshallField(stage, "IsWithCastShadow"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (extrudegeometry *ExtrudeGeometry) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(extrudegeometry.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(extrudegeometry.GongMarshallField(stage, "Shape"))
		pointersInitializesStatements.WriteString(extrudegeometry.GongMarshallField(stage, "ExtrudePath"))
		initializerStatements.WriteString(extrudegeometry.GongMarshallField(stage, "Steps"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (mesh *Mesh) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(mesh.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(mesh.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(mesh.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(mesh.GongMarshallField(stage, "Z"))
		pointersInitializesStatements.WriteString(mesh.GongMarshallField(stage, "MeshMaterialBasic"))
		pointersInitializesStatements.WriteString(mesh.GongMarshallField(stage, "MeshPhysicalMaterial"))
		pointersInitializesStatements.WriteString(mesh.GongMarshallField(stage, "CylinderGeometry"))
		pointersInitializesStatements.WriteString(mesh.GongMarshallField(stage, "BoxGeometry"))
		pointersInitializesStatements.WriteString(mesh.GongMarshallField(stage, "SphereGeometry"))
		pointersInitializesStatements.WriteString(mesh.GongMarshallField(stage, "TorusGeometry"))
		pointersInitializesStatements.WriteString(mesh.GongMarshallField(stage, "PlaneGeometry"))
		pointersInitializesStatements.WriteString(mesh.GongMarshallField(stage, "TubeGeometry"))
		pointersInitializesStatements.WriteString(mesh.GongMarshallField(stage, "ExtrudeGeometry"))
		pointersInitializesStatements.WriteString(mesh.GongMarshallField(stage, "BufferGeometry"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (meshmaterialbasic *MeshMaterialBasic) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(meshmaterialbasic.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(meshmaterialbasic.GongMarshallField(stage, "Color"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (meshphysicalmaterial *MeshPhysicalMaterial) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(meshphysicalmaterial.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(meshphysicalmaterial.GongMarshallField(stage, "Color"))
		initializerStatements.WriteString(meshphysicalmaterial.GongMarshallField(stage, "Wireframe"))
		initializerStatements.WriteString(meshphysicalmaterial.GongMarshallField(stage, "Opacity"))
		initializerStatements.WriteString(meshphysicalmaterial.GongMarshallField(stage, "Transparent"))
		initializerStatements.WriteString(meshphysicalmaterial.GongMarshallField(stage, "Visible"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (planegeometry *PlaneGeometry) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(planegeometry.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(planegeometry.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(planegeometry.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(planegeometry.GongMarshallField(stage, "WidthSegments"))
		initializerStatements.WriteString(planegeometry.GongMarshallField(stage, "HeightSegments"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (shape *Shape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(shape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(shape.GongMarshallField(stage, "Points"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (spheregeometry *SphereGeometry) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(spheregeometry.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(spheregeometry.GongMarshallField(stage, "Radius"))
		initializerStatements.WriteString(spheregeometry.GongMarshallField(stage, "WidthSegments"))
		initializerStatements.WriteString(spheregeometry.GongMarshallField(stage, "HeightSegments"))
		initializerStatements.WriteString(spheregeometry.GongMarshallField(stage, "PhiStart"))
		initializerStatements.WriteString(spheregeometry.GongMarshallField(stage, "PhiLength"))
		initializerStatements.WriteString(spheregeometry.GongMarshallField(stage, "ThetaStart"))
		initializerStatements.WriteString(spheregeometry.GongMarshallField(stage, "ThetaLength"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (torusgeometry *TorusGeometry) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(torusgeometry.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(torusgeometry.GongMarshallField(stage, "Radius"))
		initializerStatements.WriteString(torusgeometry.GongMarshallField(stage, "Tube"))
		initializerStatements.WriteString(torusgeometry.GongMarshallField(stage, "RadialSegments"))
		initializerStatements.WriteString(torusgeometry.GongMarshallField(stage, "TubularSegments"))
		initializerStatements.WriteString(torusgeometry.GongMarshallField(stage, "Arc"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (triangle *Triangle) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(triangle.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(triangle.GongMarshallField(stage, "V1"))
		initializerStatements.WriteString(triangle.GongMarshallField(stage, "V2"))
		initializerStatements.WriteString(triangle.GongMarshallField(stage, "V3"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (tubegeometry *TubeGeometry) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(tubegeometry.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(tubegeometry.GongMarshallField(stage, "Path"))
		initializerStatements.WriteString(tubegeometry.GongMarshallField(stage, "TubularSegments"))
		initializerStatements.WriteString(tubegeometry.GongMarshallField(stage, "Radius"))
		initializerStatements.WriteString(tubegeometry.GongMarshallField(stage, "RadialSegments"))
		initializerStatements.WriteString(tubegeometry.GongMarshallField(stage, "Closed"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (vector2 *Vector2) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(vector2.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(vector2.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(vector2.GongMarshallField(stage, "Y"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (vector3 *Vector3) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(vector3.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(vector3.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(vector3.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(vector3.GongMarshallField(stage, "Z"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
