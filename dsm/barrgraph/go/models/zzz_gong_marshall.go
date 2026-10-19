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
	modelsPackageName := "github.com/fullstack-lang/gong/dsm/barrgraph/go/models"
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
		for _, artefacttype := range __gong__sortStageSetInstances(stage.ArtefactTypes, stage.ArtefactType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			artefacttypeIdent := artefacttype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ArtefactType{Name: %s}).Stage(stage)", artefacttypeIdent, __gong__toRawStringLiteral(artefacttype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", artefacttypeIdent, __gong__toRawStringLiteral(artefacttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", artefacttypeIdent, __gong__toRawStringLiteral(artefacttype.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", artefacttypeIdent, artefacttype.IsExpanded))
		}
	}
	if stage != nil {
		for _, artefacttypeshape := range __gong__sortStageSetInstances(stage.ArtefactTypeShapes, stage.ArtefactTypeShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			artefacttypeshapeIdent := artefacttypeshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ArtefactTypeShape{Name: %s}).Stage(stage)", artefacttypeshapeIdent, __gong__toRawStringLiteral(artefacttypeshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", artefacttypeshapeIdent, __gong__toRawStringLiteral(artefacttypeshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", artefacttypeshapeIdent, artefacttypeshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", artefacttypeshapeIdent, artefacttypeshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", artefacttypeshapeIdent, artefacttypeshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", artefacttypeshapeIdent, artefacttypeshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", artefacttypeshapeIdent, artefacttypeshape.IsHidden))
			if artefacttypeshape.ArtefactType != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := artefacttypeshape.ArtefactType.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ArtefactType = %s", artefacttypeshapeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, artist := range __gong__sortStageSetInstances(stage.Artists, stage.Artist_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			artistIdent := artist.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Artist{Name: %s}).Stage(stage)", artistIdent, __gong__toRawStringLiteral(artist.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", artistIdent, __gong__toRawStringLiteral(artist.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", artistIdent, __gong__toRawStringLiteral(artist.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", artistIdent, artist.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsDead = %t", artistIdent, artist.IsDead))
			values.WriteString(fmt.Sprintf("\n\t%s.DateOfDeath, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", artistIdent, artist.DateOfDeath.String()))
			if artist.Place != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := artist.Place.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Place = %s", artistIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, artistshape := range __gong__sortStageSetInstances(stage.ArtistShapes, stage.ArtistShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			artistshapeIdent := artistshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ArtistShape{Name: %s}).Stage(stage)", artistshapeIdent, __gong__toRawStringLiteral(artistshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", artistshapeIdent, __gong__toRawStringLiteral(artistshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", artistshapeIdent, artistshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", artistshapeIdent, artistshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", artistshapeIdent, artistshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", artistshapeIdent, artistshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", artistshapeIdent, artistshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.ImagePng_X = %f", artistshapeIdent, artistshape.ImagePng_X))
			values.WriteString(fmt.Sprintf("\n\t%s.ImagePng_Y = %f", artistshapeIdent, artistshape.ImagePng_Y))
			values.WriteString(fmt.Sprintf("\n\t%s.ImagePng_Width = %f", artistshapeIdent, artistshape.ImagePng_Width))
			values.WriteString(fmt.Sprintf("\n\t%s.ImagePng_Height = %f", artistshapeIdent, artistshape.ImagePng_Height))
			values.WriteString(fmt.Sprintf("\n\t%s.ImagePng_X_Offset = %f", artistshapeIdent, artistshape.ImagePng_X_Offset))
			values.WriteString(fmt.Sprintf("\n\t%s.ImagePng_Y_Offset = %f", artistshapeIdent, artistshape.ImagePng_Y_Offset))
			if artistshape.ImagePng_RectAnchorType.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.ImagePng_RectAnchorType = models.%s", artistshapeIdent, artistshape.ImagePng_RectAnchorType.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.ImagePng_RectAnchorType = \"\"", artistshapeIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.ImagePngBase64Content = %s", artistshapeIdent, __gong__toRawStringLiteral(artistshape.ImagePngBase64Content)))
			if artistshape.Artist != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := artistshape.Artist.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Artist = %s", artistshapeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, controlpointshape := range __gong__sortStageSetInstances(stage.ControlPointShapes, stage.ControlPointShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			controlpointshapeIdent := controlpointshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ControlPointShape{Name: %s}).Stage(stage)", controlpointshapeIdent, __gong__toRawStringLiteral(controlpointshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", controlpointshapeIdent, __gong__toRawStringLiteral(controlpointshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.X_Relative = %f", controlpointshapeIdent, controlpointshape.X_Relative))
			values.WriteString(fmt.Sprintf("\n\t%s.Y_Relative = %f", controlpointshapeIdent, controlpointshape.Y_Relative))
			values.WriteString(fmt.Sprintf("\n\t%s.IsStartShapeTheClosestShape = %t", controlpointshapeIdent, controlpointshape.IsStartShapeTheClosestShape))
		}
	}
	if stage != nil {
		for _, desk := range __gong__sortStageSetInstances(stage.Desks, stage.Desk_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			deskIdent := desk.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Desk{Name: %s}).Stage(stage)", deskIdent, __gong__toRawStringLiteral(desk.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", deskIdent, __gong__toRawStringLiteral(desk.Name)))
			if desk.SelectedDiagram != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := desk.SelectedDiagram.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SelectedDiagram = %s", deskIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, diagram := range __gong__sortStageSetInstances(stage.Diagrams, stage.Diagram_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			diagramIdent := diagram.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Diagram{Name: %s}).Stage(stage)", diagramIdent, __gong__toRawStringLiteral(diagram.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", diagramIdent, __gong__toRawStringLiteral(diagram.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", diagramIdent, __gong__toRawStringLiteral(diagram.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", diagramIdent, diagram.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsChecked = %t", diagramIdent, diagram.IsChecked))
			values.WriteString(fmt.Sprintf("\n\t%s.IsEditable = %t", diagramIdent, diagram.IsEditable))
			values.WriteString(fmt.Sprintf("\n\t%s.IsNodeExpanded = %t", diagramIdent, diagram.IsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsMovementCategoryNodeExpanded = %t", diagramIdent, diagram.IsMovementCategoryNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsArtefactTypeCategoryNodeExpanded = %t", diagramIdent, diagram.IsArtefactTypeCategoryNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsArtistCategoryNodeExpanded = %t", diagramIdent, diagram.IsArtistCategoryNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsInfluenceCategoryNodeExpanded = %t", diagramIdent, diagram.IsInfluenceCategoryNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsMovementCategoryHidden = %t", diagramIdent, diagram.IsMovementCategoryHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsArtefactTypeCategoryHidden = %t", diagramIdent, diagram.IsArtefactTypeCategoryHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsArtistCategoryHidden = %t", diagramIdent, diagram.IsArtistCategoryHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsInfluenceCategoryHidden = %t", diagramIdent, diagram.IsInfluenceCategoryHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.StartDate, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", diagramIdent, diagram.StartDate.String()))
			values.WriteString(fmt.Sprintf("\n\t%s.EndDate, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", diagramIdent, diagram.EndDate.String()))
			values.WriteString(fmt.Sprintf("\n\t%s.NbYearsForIntervals = %d", diagramIdent, diagram.NbYearsForIntervals))
			values.WriteString(fmt.Sprintf("\n\t%s.XMargin = %f", diagramIdent, diagram.XMargin))
			values.WriteString(fmt.Sprintf("\n\t%s.YMargin = %f", diagramIdent, diagram.YMargin))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", diagramIdent, diagram.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.NextVerticalDateXMargin = %f", diagramIdent, diagram.NextVerticalDateXMargin))
			values.WriteString(fmt.Sprintf("\n\t%s.RedColorCode = %s", diagramIdent, __gong__toRawStringLiteral(diagram.RedColorCode)))
			values.WriteString(fmt.Sprintf("\n\t%s.BackgroundGreyColorCode = %s", diagramIdent, __gong__toRawStringLiteral(diagram.BackgroundGreyColorCode)))
			values.WriteString(fmt.Sprintf("\n\t%s.GrayColorCode = %s", diagramIdent, __gong__toRawStringLiteral(diagram.GrayColorCode)))
			values.WriteString(fmt.Sprintf("\n\t%s.BottomBoxYOffset = %f", diagramIdent, diagram.BottomBoxYOffset))
			values.WriteString(fmt.Sprintf("\n\t%s.BottomBoxWidth = %f", diagramIdent, diagram.BottomBoxWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.BottomBoxHeigth = %f", diagramIdent, diagram.BottomBoxHeigth))
			values.WriteString(fmt.Sprintf("\n\t%s.BottomBoxFontSize = %s", diagramIdent, __gong__toRawStringLiteral(diagram.BottomBoxFontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.BottomBoxFontWeigth = %s", diagramIdent, __gong__toRawStringLiteral(diagram.BottomBoxFontWeigth)))
			values.WriteString(fmt.Sprintf("\n\t%s.BottomBoxFontFamily = %s", diagramIdent, __gong__toRawStringLiteral(diagram.BottomBoxFontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.BottomBoxLetterSpacing = %s", diagramIdent, __gong__toRawStringLiteral(diagram.BottomBoxLetterSpacing)))
			values.WriteString(fmt.Sprintf("\n\t%s.BottomBoxLetterColorCode = %s", diagramIdent, __gong__toRawStringLiteral(diagram.BottomBoxLetterColorCode)))
			if diagram.MovementRectAnchorType.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.MovementRectAnchorType = models.%s", diagramIdent, diagram.MovementRectAnchorType.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.MovementRectAnchorType = \"\"", diagramIdent))
			}
			if diagram.MovementTextAnchorType.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.MovementTextAnchorType = models.%s", diagramIdent, diagram.MovementTextAnchorType.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.MovementTextAnchorType = \"\"", diagramIdent))
			}
			if diagram.MovementDominantBaselineType.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.MovementDominantBaselineType = models.%s", diagramIdent, diagram.MovementDominantBaselineType.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.MovementDominantBaselineType = \"\"", diagramIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.MovementFontSize = %s", diagramIdent, __gong__toRawStringLiteral(diagram.MovementFontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.MajorMovementFontSize = %s", diagramIdent, __gong__toRawStringLiteral(diagram.MajorMovementFontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.MinorMovementFontSize = %s", diagramIdent, __gong__toRawStringLiteral(diagram.MinorMovementFontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.MovementFontWeigth = %s", diagramIdent, __gong__toRawStringLiteral(diagram.MovementFontWeigth)))
			values.WriteString(fmt.Sprintf("\n\t%s.MovementFontFamily = %s", diagramIdent, __gong__toRawStringLiteral(diagram.MovementFontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.MovementLetterSpacing = %s", diagramIdent, __gong__toRawStringLiteral(diagram.MovementLetterSpacing)))
			values.WriteString(fmt.Sprintf("\n\t%s.AbstractMovementFontSize = %s", diagramIdent, __gong__toRawStringLiteral(diagram.AbstractMovementFontSize)))
			if diagram.AbstractMovementRectAnchorType.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.AbstractMovementRectAnchorType = models.%s", diagramIdent, diagram.AbstractMovementRectAnchorType.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.AbstractMovementRectAnchorType = \"\"", diagramIdent))
			}
			if diagram.AbstractMovementTextAnchorType.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.AbstractMovementTextAnchorType = models.%s", diagramIdent, diagram.AbstractMovementTextAnchorType.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.AbstractMovementTextAnchorType = \"\"", diagramIdent))
			}
			if diagram.AbstractDominantBaselineType.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.AbstractDominantBaselineType = models.%s", diagramIdent, diagram.AbstractDominantBaselineType.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.AbstractDominantBaselineType = \"\"", diagramIdent))
			}
			if diagram.MovementDateRectAnchorType.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.MovementDateRectAnchorType = models.%s", diagramIdent, diagram.MovementDateRectAnchorType.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.MovementDateRectAnchorType = \"\"", diagramIdent))
			}
			if diagram.MovementDateTextAnchorType.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.MovementDateTextAnchorType = models.%s", diagramIdent, diagram.MovementDateTextAnchorType.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.MovementDateTextAnchorType = \"\"", diagramIdent))
			}
			if diagram.MovementDateTextDominantBaselineType.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.MovementDateTextDominantBaselineType = models.%s", diagramIdent, diagram.MovementDateTextDominantBaselineType.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.MovementDateTextDominantBaselineType = \"\"", diagramIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.MovementDateAndPlacesFontSize = %s", diagramIdent, __gong__toRawStringLiteral(diagram.MovementDateAndPlacesFontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.MovementDateAndPlacesFontWeigth = %s", diagramIdent, __gong__toRawStringLiteral(diagram.MovementDateAndPlacesFontWeigth)))
			values.WriteString(fmt.Sprintf("\n\t%s.MovementDateAndPlacesFontFamily = %s", diagramIdent, __gong__toRawStringLiteral(diagram.MovementDateAndPlacesFontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.MovementDateAndPlacesLetterSpacing = %s", diagramIdent, __gong__toRawStringLiteral(diagram.MovementDateAndPlacesLetterSpacing)))
			values.WriteString(fmt.Sprintf("\n\t%s.MovementBelowArcY_Offset = %f", diagramIdent, diagram.MovementBelowArcY_Offset))
			values.WriteString(fmt.Sprintf("\n\t%s.MovementBelowArcY_OffsetPerPlace = %f", diagramIdent, diagram.MovementBelowArcY_OffsetPerPlace))
			if diagram.MovementPlacesRectAnchorType.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.MovementPlacesRectAnchorType = models.%s", diagramIdent, diagram.MovementPlacesRectAnchorType.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.MovementPlacesRectAnchorType = \"\"", diagramIdent))
			}
			if diagram.MovementPlacesTextAnchorType.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.MovementPlacesTextAnchorType = models.%s", diagramIdent, diagram.MovementPlacesTextAnchorType.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.MovementPlacesTextAnchorType = \"\"", diagramIdent))
			}
			if diagram.MovementPlacesDominantBaselineType.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.MovementPlacesDominantBaselineType = models.%s", diagramIdent, diagram.MovementPlacesDominantBaselineType.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.MovementPlacesDominantBaselineType = \"\"", diagramIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.ArtefactTypeFontSize = %s", diagramIdent, __gong__toRawStringLiteral(diagram.ArtefactTypeFontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.ArtefactTypeFontWeigth = %s", diagramIdent, __gong__toRawStringLiteral(diagram.ArtefactTypeFontWeigth)))
			values.WriteString(fmt.Sprintf("\n\t%s.ArtefactTypeFontFamily = %s", diagramIdent, __gong__toRawStringLiteral(diagram.ArtefactTypeFontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.ArtefactTypeLetterSpacing = %s", diagramIdent, __gong__toRawStringLiteral(diagram.ArtefactTypeLetterSpacing)))
			if diagram.ArtefactTypeRectAnchorType.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.ArtefactTypeRectAnchorType = models.%s", diagramIdent, diagram.ArtefactTypeRectAnchorType.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.ArtefactTypeRectAnchorType = \"\"", diagramIdent))
			}
			if diagram.ArtefactDominantBaselineType.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.ArtefactDominantBaselineType = models.%s", diagramIdent, diagram.ArtefactDominantBaselineType.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.ArtefactDominantBaselineType = \"\"", diagramIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.ArtefactTypeStrokeWidth = %f", diagramIdent, diagram.ArtefactTypeStrokeWidth))
			if diagram.ArtistRectAnchorType.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.ArtistRectAnchorType = models.%s", diagramIdent, diagram.ArtistRectAnchorType.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.ArtistRectAnchorType = \"\"", diagramIdent))
			}
			if diagram.ArtistTextAnchorType.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.ArtistTextAnchorType = models.%s", diagramIdent, diagram.ArtistTextAnchorType.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.ArtistTextAnchorType = \"\"", diagramIdent))
			}
			if diagram.ArtistDominantBaselineType.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.ArtistDominantBaselineType = models.%s", diagramIdent, diagram.ArtistDominantBaselineType.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.ArtistDominantBaselineType = \"\"", diagramIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.ArtistFontSize = %s", diagramIdent, __gong__toRawStringLiteral(diagram.ArtistFontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.MajorArtistFontSize = %s", diagramIdent, __gong__toRawStringLiteral(diagram.MajorArtistFontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.MinorArtistFontSize = %s", diagramIdent, __gong__toRawStringLiteral(diagram.MinorArtistFontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.ArtistFontWeigth = %s", diagramIdent, __gong__toRawStringLiteral(diagram.ArtistFontWeigth)))
			values.WriteString(fmt.Sprintf("\n\t%s.ArtistFontFamily = %s", diagramIdent, __gong__toRawStringLiteral(diagram.ArtistFontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.ArtistLetterSpacing = %s", diagramIdent, __gong__toRawStringLiteral(diagram.ArtistLetterSpacing)))
			if diagram.ArtistDateRectAnchorType.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.ArtistDateRectAnchorType = models.%s", diagramIdent, diagram.ArtistDateRectAnchorType.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.ArtistDateRectAnchorType = \"\"", diagramIdent))
			}
			if diagram.ArtistDateTextAnchorType.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.ArtistDateTextAnchorType = models.%s", diagramIdent, diagram.ArtistDateTextAnchorType.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.ArtistDateTextAnchorType = \"\"", diagramIdent))
			}
			if diagram.ArtistDateDominantBaselineType.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.ArtistDateDominantBaselineType = models.%s", diagramIdent, diagram.ArtistDateDominantBaselineType.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.ArtistDateDominantBaselineType = \"\"", diagramIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.ArtistDateAndPlacesFontSize = %s", diagramIdent, __gong__toRawStringLiteral(diagram.ArtistDateAndPlacesFontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.ArtistDateAndPlacesFontWeigth = %s", diagramIdent, __gong__toRawStringLiteral(diagram.ArtistDateAndPlacesFontWeigth)))
			values.WriteString(fmt.Sprintf("\n\t%s.ArtistDateAndPlacesFontFamily = %s", diagramIdent, __gong__toRawStringLiteral(diagram.ArtistDateAndPlacesFontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.ArtistDateAndPlacesLetterSpacing = %s", diagramIdent, __gong__toRawStringLiteral(diagram.ArtistDateAndPlacesLetterSpacing)))
			if diagram.ArtistPlacesRectAnchorType.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.ArtistPlacesRectAnchorType = models.%s", diagramIdent, diagram.ArtistPlacesRectAnchorType.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.ArtistPlacesRectAnchorType = \"\"", diagramIdent))
			}
			if diagram.ArtistPlacesTextAnchorType.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.ArtistPlacesTextAnchorType = models.%s", diagramIdent, diagram.ArtistPlacesTextAnchorType.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.ArtistPlacesTextAnchorType = \"\"", diagramIdent))
			}
			if diagram.ArtistPlacesDominantBaselineType.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.ArtistPlacesDominantBaselineType = models.%s", diagramIdent, diagram.ArtistPlacesDominantBaselineType.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.ArtistPlacesDominantBaselineType = \"\"", diagramIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.InfluenceArrowSize = %f", diagramIdent, diagram.InfluenceArrowSize))
			values.WriteString(fmt.Sprintf("\n\t%s.InfluenceArrowStartOffset = %f", diagramIdent, diagram.InfluenceArrowStartOffset))
			values.WriteString(fmt.Sprintf("\n\t%s.InfluenceArrowEndOffset = %f", diagramIdent, diagram.InfluenceArrowEndOffset))
			values.WriteString(fmt.Sprintf("\n\t%s.InfluenceCornerRadius = %f", diagramIdent, diagram.InfluenceCornerRadius))
			values.WriteString(fmt.Sprintf("\n\t%s.InfluenceDashedLinePattern = %s", diagramIdent, __gong__toRawStringLiteral(diagram.InfluenceDashedLinePattern)))
			for _, elem := range diagram.MovementShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MovementShapes = append(%s.MovementShapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ArtefactTypeShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ArtefactTypeShapes = append(%s.ArtefactTypeShapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ArtistShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ArtistShapes = append(%s.ArtistShapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.InfluenceShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.InfluenceShapes = append(%s.InfluenceShapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, influence := range __gong__sortStageSetInstances(stage.Influences, stage.Influence_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			influenceIdent := influence.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Influence{Name: %s}).Stage(stage)", influenceIdent, __gong__toRawStringLiteral(influence.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", influenceIdent, __gong__toRawStringLiteral(influence.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", influenceIdent, __gong__toRawStringLiteral(influence.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", influenceIdent, influence.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHypothtical = %t", influenceIdent, influence.IsHypothtical))
			if influence.SourceMovement != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := influence.SourceMovement.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SourceMovement = %s", influenceIdent, targetIdent))
			}
			if influence.SourceArtefactType != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := influence.SourceArtefactType.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SourceArtefactType = %s", influenceIdent, targetIdent))
			}
			if influence.SourceArtist != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := influence.SourceArtist.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SourceArtist = %s", influenceIdent, targetIdent))
			}
			if influence.TargetMovement != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := influence.TargetMovement.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TargetMovement = %s", influenceIdent, targetIdent))
			}
			if influence.TargetArtefactType != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := influence.TargetArtefactType.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TargetArtefactType = %s", influenceIdent, targetIdent))
			}
			if influence.TargetArtist != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := influence.TargetArtist.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TargetArtist = %s", influenceIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, influenceshape := range __gong__sortStageSetInstances(stage.InfluenceShapes, stage.InfluenceShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			influenceshapeIdent := influenceshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.InfluenceShape{Name: %s}).Stage(stage)", influenceshapeIdent, __gong__toRawStringLiteral(influenceshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", influenceshapeIdent, __gong__toRawStringLiteral(influenceshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", influenceshapeIdent, influenceshape.IsHidden))
			if influenceshape.Influence != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := influenceshape.Influence.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Influence = %s", influenceshapeIdent, targetIdent))
			}
			for _, elem := range influenceshape.ControlPointShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ControlPointShapes = append(%s.ControlPointShapes, %s)", influenceshapeIdent, influenceshapeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, library := range __gong__sortStageSetInstances(stage.Librarys, stage.Library_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			libraryIdent := library.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Library{Name: %s}).Stage(stage)", libraryIdent, __gong__toRawStringLiteral(library.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", libraryIdent, __gong__toRawStringLiteral(library.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", libraryIdent, __gong__toRawStringLiteral(library.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", libraryIdent, __gong__toRawStringLiteral(library.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", libraryIdent, library.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsRootLibrary = %t", libraryIdent, library.IsRootLibrary))
			values.WriteString(fmt.Sprintf("\n\t%s.IsSubLibrariesNodeExpanded = %t", libraryIdent, library.IsSubLibrariesNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.NbPixPerCharacter = %f", libraryIdent, library.NbPixPerCharacter))
			values.WriteString(fmt.Sprintf("\n\t%s.LogoSVGFile = %s", libraryIdent, __gong__toRawStringLiteral(library.LogoSVGFile)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpandedTmp = %t", libraryIdent, library.IsExpandedTmp))
			for _, elem := range library.SubLibraries {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SubLibraries = append(%s.SubLibraries, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.SubLibrariesWhoseNodeIsExpanded {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SubLibrariesWhoseNodeIsExpanded = append(%s.SubLibrariesWhoseNodeIsExpanded, %s)", libraryIdent, libraryIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, movement := range __gong__sortStageSetInstances(stage.Movements, stage.Movement_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			movementIdent := movement.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Movement{Name: %s}).Stage(stage)", movementIdent, __gong__toRawStringLiteral(movement.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", movementIdent, __gong__toRawStringLiteral(movement.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", movementIdent, __gong__toRawStringLiteral(movement.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", movementIdent, movement.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.Date, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", movementIdent, movement.Date.String()))
			values.WriteString(fmt.Sprintf("\n\t%s.HideDate = %t", movementIdent, movement.HideDate))
			values.WriteString(fmt.Sprintf("\n\t%s.HasTaxonomicFilter = %t", movementIdent, movement.HasTaxonomicFilter))
			values.WriteString(fmt.Sprintf("\n\t%s.TaxonomicFilter = %s", movementIdent, __gong__toRawStringLiteral(movement.TaxonomicFilter)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsFeatured = %t", movementIdent, movement.IsFeatured))
			values.WriteString(fmt.Sprintf("\n\t%s.FeaturePrefix = %s", movementIdent, __gong__toRawStringLiteral(movement.FeaturePrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsMajor = %t", movementIdent, movement.IsMajor))
			values.WriteString(fmt.Sprintf("\n\t%s.IsMinor = %t", movementIdent, movement.IsMinor))
			values.WriteString(fmt.Sprintf("\n\t%s.AdditionnalName = %s", movementIdent, __gong__toRawStringLiteral(movement.AdditionnalName)))
			for _, elem := range movement.Places {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Places = append(%s.Places, %s)", movementIdent, movementIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, movementshape := range __gong__sortStageSetInstances(stage.MovementShapes, stage.MovementShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			movementshapeIdent := movementshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.MovementShape{Name: %s}).Stage(stage)", movementshapeIdent, __gong__toRawStringLiteral(movementshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", movementshapeIdent, __gong__toRawStringLiteral(movementshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", movementshapeIdent, movementshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", movementshapeIdent, movementshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", movementshapeIdent, movementshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", movementshapeIdent, movementshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", movementshapeIdent, movementshape.IsHidden))
			if movementshape.Movement != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := movementshape.Movement.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Movement = %s", movementshapeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, place := range __gong__sortStageSetInstances(stage.Places, stage.Place_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			placeIdent := place.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Place{Name: %s}).Stage(stage)", placeIdent, __gong__toRawStringLiteral(place.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", placeIdent, __gong__toRawStringLiteral(place.Name)))
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
go run github.com/fullstack-lang/gong/dsm/barrgraph/go/models/probe/cmd/probe@%s stage.go
*/
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/dsm/barrgraph/go/models"%s
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

	if modelsPackageName != "" && modelsPackageName != "github.com/fullstack-lang/gong/dsm/barrgraph/go/models" {
		res = strings.ReplaceAll(res, "\"github.com/fullstack-lang/gong/dsm/barrgraph/go/models\"", fmt.Sprintf("%q", modelsPackageName))
	}

	return res, nil
}

// insertion point for marshall field methods
func (artefacttype *ArtefactType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := artefacttype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", artefacttype.Name)
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", artefacttype.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", artefacttype.IsExpanded)

	default:
		log.Panicf("Unknown field %s for Gongstruct ArtefactType", fieldName)
	}
	return
}

func (artefacttypeshape *ArtefactTypeShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := artefacttypeshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", artefacttypeshape.Name)
	case "X":
		res = __gong__marshallFloat(ident, "X", artefacttypeshape.X)
	case "Y":
		res = __gong__marshallFloat(ident, "Y", artefacttypeshape.Y)
	case "Width":
		res = __gong__marshallFloat(ident, "Width", artefacttypeshape.Width)
	case "Height":
		res = __gong__marshallFloat(ident, "Height", artefacttypeshape.Height)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", artefacttypeshape.IsHidden)

	case "ArtefactType":
		if artefacttypeshape.ArtefactType != nil {
			res = __gong__marshallPointer(ident, "ArtefactType", artefacttypeshape.ArtefactType.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ArtefactType", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ArtefactTypeShape", fieldName)
	}
	return
}

func (artist *Artist) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := artist.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", artist.Name)
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", artist.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", artist.IsExpanded)
	case "IsDead":
		res = __gong__marshallBool(ident, "IsDead", artist.IsDead)
	case "DateOfDeath":
		res = __gong__marshallTime(ident, "DateOfDeath", artist.DateOfDeath.String())

	case "Place":
		if artist.Place != nil {
			res = __gong__marshallPointer(ident, "Place", artist.Place.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Place", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Artist", fieldName)
	}
	return
}

func (artistshape *ArtistShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := artistshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", artistshape.Name)
	case "X":
		res = __gong__marshallFloat(ident, "X", artistshape.X)
	case "Y":
		res = __gong__marshallFloat(ident, "Y", artistshape.Y)
	case "Width":
		res = __gong__marshallFloat(ident, "Width", artistshape.Width)
	case "Height":
		res = __gong__marshallFloat(ident, "Height", artistshape.Height)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", artistshape.IsHidden)
	case "ImagePng_X":
		res = __gong__marshallFloat(ident, "ImagePng_X", artistshape.ImagePng_X)
	case "ImagePng_Y":
		res = __gong__marshallFloat(ident, "ImagePng_Y", artistshape.ImagePng_Y)
	case "ImagePng_Width":
		res = __gong__marshallFloat(ident, "ImagePng_Width", artistshape.ImagePng_Width)
	case "ImagePng_Height":
		res = __gong__marshallFloat(ident, "ImagePng_Height", artistshape.ImagePng_Height)
	case "ImagePng_X_Offset":
		res = __gong__marshallFloat(ident, "ImagePng_X_Offset", artistshape.ImagePng_X_Offset)
	case "ImagePng_Y_Offset":
		res = __gong__marshallFloat(ident, "ImagePng_Y_Offset", artistshape.ImagePng_Y_Offset)
	case "ImagePng_RectAnchorType":
		res = __gong__marshallEnumString(ident, "ImagePng_RectAnchorType", artistshape.ImagePng_RectAnchorType.ToCodeString())
	case "ImagePngBase64Content":
		res = __gong__marshallString(ident, "ImagePngBase64Content", artistshape.ImagePngBase64Content)

	case "Artist":
		if artistshape.Artist != nil {
			res = __gong__marshallPointer(ident, "Artist", artistshape.Artist.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Artist", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ArtistShape", fieldName)
	}
	return
}

func (controlpointshape *ControlPointShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := controlpointshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", controlpointshape.Name)
	case "X_Relative":
		res = __gong__marshallFloat(ident, "X_Relative", controlpointshape.X_Relative)
	case "Y_Relative":
		res = __gong__marshallFloat(ident, "Y_Relative", controlpointshape.Y_Relative)
	case "IsStartShapeTheClosestShape":
		res = __gong__marshallBool(ident, "IsStartShapeTheClosestShape", controlpointshape.IsStartShapeTheClosestShape)

	default:
		log.Panicf("Unknown field %s for Gongstruct ControlPointShape", fieldName)
	}
	return
}

func (desk *Desk) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := desk.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", desk.Name)

	case "SelectedDiagram":
		if desk.SelectedDiagram != nil {
			res = __gong__marshallPointer(ident, "SelectedDiagram", desk.SelectedDiagram.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "SelectedDiagram", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Desk", fieldName)
	}
	return
}

func (diagram *Diagram) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := diagram.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", diagram.Name)
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", diagram.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", diagram.IsExpanded)
	case "IsChecked":
		res = __gong__marshallBool(ident, "IsChecked", diagram.IsChecked)
	case "IsEditable":
		res = __gong__marshallBool(ident, "IsEditable", diagram.IsEditable)
	case "IsNodeExpanded":
		res = __gong__marshallBool(ident, "IsNodeExpanded", diagram.IsNodeExpanded)
	case "IsMovementCategoryNodeExpanded":
		res = __gong__marshallBool(ident, "IsMovementCategoryNodeExpanded", diagram.IsMovementCategoryNodeExpanded)
	case "IsArtefactTypeCategoryNodeExpanded":
		res = __gong__marshallBool(ident, "IsArtefactTypeCategoryNodeExpanded", diagram.IsArtefactTypeCategoryNodeExpanded)
	case "IsArtistCategoryNodeExpanded":
		res = __gong__marshallBool(ident, "IsArtistCategoryNodeExpanded", diagram.IsArtistCategoryNodeExpanded)
	case "IsInfluenceCategoryNodeExpanded":
		res = __gong__marshallBool(ident, "IsInfluenceCategoryNodeExpanded", diagram.IsInfluenceCategoryNodeExpanded)
	case "IsMovementCategoryHidden":
		res = __gong__marshallBool(ident, "IsMovementCategoryHidden", diagram.IsMovementCategoryHidden)
	case "IsArtefactTypeCategoryHidden":
		res = __gong__marshallBool(ident, "IsArtefactTypeCategoryHidden", diagram.IsArtefactTypeCategoryHidden)
	case "IsArtistCategoryHidden":
		res = __gong__marshallBool(ident, "IsArtistCategoryHidden", diagram.IsArtistCategoryHidden)
	case "IsInfluenceCategoryHidden":
		res = __gong__marshallBool(ident, "IsInfluenceCategoryHidden", diagram.IsInfluenceCategoryHidden)
	case "StartDate":
		res = __gong__marshallTime(ident, "StartDate", diagram.StartDate.String())
	case "EndDate":
		res = __gong__marshallTime(ident, "EndDate", diagram.EndDate.String())
	case "NbYearsForIntervals":
		res = __gong__marshallInt(ident, "NbYearsForIntervals", diagram.NbYearsForIntervals)
	case "XMargin":
		res = __gong__marshallFloat(ident, "XMargin", diagram.XMargin)
	case "YMargin":
		res = __gong__marshallFloat(ident, "YMargin", diagram.YMargin)
	case "Height":
		res = __gong__marshallFloat(ident, "Height", diagram.Height)
	case "NextVerticalDateXMargin":
		res = __gong__marshallFloat(ident, "NextVerticalDateXMargin", diagram.NextVerticalDateXMargin)
	case "RedColorCode":
		res = __gong__marshallString(ident, "RedColorCode", diagram.RedColorCode)
	case "BackgroundGreyColorCode":
		res = __gong__marshallString(ident, "BackgroundGreyColorCode", diagram.BackgroundGreyColorCode)
	case "GrayColorCode":
		res = __gong__marshallString(ident, "GrayColorCode", diagram.GrayColorCode)
	case "BottomBoxYOffset":
		res = __gong__marshallFloat(ident, "BottomBoxYOffset", diagram.BottomBoxYOffset)
	case "BottomBoxWidth":
		res = __gong__marshallFloat(ident, "BottomBoxWidth", diagram.BottomBoxWidth)
	case "BottomBoxHeigth":
		res = __gong__marshallFloat(ident, "BottomBoxHeigth", diagram.BottomBoxHeigth)
	case "BottomBoxFontSize":
		res = __gong__marshallString(ident, "BottomBoxFontSize", diagram.BottomBoxFontSize)
	case "BottomBoxFontWeigth":
		res = __gong__marshallString(ident, "BottomBoxFontWeigth", diagram.BottomBoxFontWeigth)
	case "BottomBoxFontFamily":
		res = __gong__marshallString(ident, "BottomBoxFontFamily", diagram.BottomBoxFontFamily)
	case "BottomBoxLetterSpacing":
		res = __gong__marshallString(ident, "BottomBoxLetterSpacing", diagram.BottomBoxLetterSpacing)
	case "BottomBoxLetterColorCode":
		res = __gong__marshallString(ident, "BottomBoxLetterColorCode", diagram.BottomBoxLetterColorCode)
	case "MovementRectAnchorType":
		res = __gong__marshallEnumString(ident, "MovementRectAnchorType", diagram.MovementRectAnchorType.ToCodeString())
	case "MovementTextAnchorType":
		res = __gong__marshallEnumString(ident, "MovementTextAnchorType", diagram.MovementTextAnchorType.ToCodeString())
	case "MovementDominantBaselineType":
		res = __gong__marshallEnumString(ident, "MovementDominantBaselineType", diagram.MovementDominantBaselineType.ToCodeString())
	case "MovementFontSize":
		res = __gong__marshallString(ident, "MovementFontSize", diagram.MovementFontSize)
	case "MajorMovementFontSize":
		res = __gong__marshallString(ident, "MajorMovementFontSize", diagram.MajorMovementFontSize)
	case "MinorMovementFontSize":
		res = __gong__marshallString(ident, "MinorMovementFontSize", diagram.MinorMovementFontSize)
	case "MovementFontWeigth":
		res = __gong__marshallString(ident, "MovementFontWeigth", diagram.MovementFontWeigth)
	case "MovementFontFamily":
		res = __gong__marshallString(ident, "MovementFontFamily", diagram.MovementFontFamily)
	case "MovementLetterSpacing":
		res = __gong__marshallString(ident, "MovementLetterSpacing", diagram.MovementLetterSpacing)
	case "AbstractMovementFontSize":
		res = __gong__marshallString(ident, "AbstractMovementFontSize", diagram.AbstractMovementFontSize)
	case "AbstractMovementRectAnchorType":
		res = __gong__marshallEnumString(ident, "AbstractMovementRectAnchorType", diagram.AbstractMovementRectAnchorType.ToCodeString())
	case "AbstractMovementTextAnchorType":
		res = __gong__marshallEnumString(ident, "AbstractMovementTextAnchorType", diagram.AbstractMovementTextAnchorType.ToCodeString())
	case "AbstractDominantBaselineType":
		res = __gong__marshallEnumString(ident, "AbstractDominantBaselineType", diagram.AbstractDominantBaselineType.ToCodeString())
	case "MovementDateRectAnchorType":
		res = __gong__marshallEnumString(ident, "MovementDateRectAnchorType", diagram.MovementDateRectAnchorType.ToCodeString())
	case "MovementDateTextAnchorType":
		res = __gong__marshallEnumString(ident, "MovementDateTextAnchorType", diagram.MovementDateTextAnchorType.ToCodeString())
	case "MovementDateTextDominantBaselineType":
		res = __gong__marshallEnumString(ident, "MovementDateTextDominantBaselineType", diagram.MovementDateTextDominantBaselineType.ToCodeString())
	case "MovementDateAndPlacesFontSize":
		res = __gong__marshallString(ident, "MovementDateAndPlacesFontSize", diagram.MovementDateAndPlacesFontSize)
	case "MovementDateAndPlacesFontWeigth":
		res = __gong__marshallString(ident, "MovementDateAndPlacesFontWeigth", diagram.MovementDateAndPlacesFontWeigth)
	case "MovementDateAndPlacesFontFamily":
		res = __gong__marshallString(ident, "MovementDateAndPlacesFontFamily", diagram.MovementDateAndPlacesFontFamily)
	case "MovementDateAndPlacesLetterSpacing":
		res = __gong__marshallString(ident, "MovementDateAndPlacesLetterSpacing", diagram.MovementDateAndPlacesLetterSpacing)
	case "MovementBelowArcY_Offset":
		res = __gong__marshallFloat(ident, "MovementBelowArcY_Offset", diagram.MovementBelowArcY_Offset)
	case "MovementBelowArcY_OffsetPerPlace":
		res = __gong__marshallFloat(ident, "MovementBelowArcY_OffsetPerPlace", diagram.MovementBelowArcY_OffsetPerPlace)
	case "MovementPlacesRectAnchorType":
		res = __gong__marshallEnumString(ident, "MovementPlacesRectAnchorType", diagram.MovementPlacesRectAnchorType.ToCodeString())
	case "MovementPlacesTextAnchorType":
		res = __gong__marshallEnumString(ident, "MovementPlacesTextAnchorType", diagram.MovementPlacesTextAnchorType.ToCodeString())
	case "MovementPlacesDominantBaselineType":
		res = __gong__marshallEnumString(ident, "MovementPlacesDominantBaselineType", diagram.MovementPlacesDominantBaselineType.ToCodeString())
	case "ArtefactTypeFontSize":
		res = __gong__marshallString(ident, "ArtefactTypeFontSize", diagram.ArtefactTypeFontSize)
	case "ArtefactTypeFontWeigth":
		res = __gong__marshallString(ident, "ArtefactTypeFontWeigth", diagram.ArtefactTypeFontWeigth)
	case "ArtefactTypeFontFamily":
		res = __gong__marshallString(ident, "ArtefactTypeFontFamily", diagram.ArtefactTypeFontFamily)
	case "ArtefactTypeLetterSpacing":
		res = __gong__marshallString(ident, "ArtefactTypeLetterSpacing", diagram.ArtefactTypeLetterSpacing)
	case "ArtefactTypeRectAnchorType":
		res = __gong__marshallEnumString(ident, "ArtefactTypeRectAnchorType", diagram.ArtefactTypeRectAnchorType.ToCodeString())
	case "ArtefactDominantBaselineType":
		res = __gong__marshallEnumString(ident, "ArtefactDominantBaselineType", diagram.ArtefactDominantBaselineType.ToCodeString())
	case "ArtefactTypeStrokeWidth":
		res = __gong__marshallFloat(ident, "ArtefactTypeStrokeWidth", diagram.ArtefactTypeStrokeWidth)
	case "ArtistRectAnchorType":
		res = __gong__marshallEnumString(ident, "ArtistRectAnchorType", diagram.ArtistRectAnchorType.ToCodeString())
	case "ArtistTextAnchorType":
		res = __gong__marshallEnumString(ident, "ArtistTextAnchorType", diagram.ArtistTextAnchorType.ToCodeString())
	case "ArtistDominantBaselineType":
		res = __gong__marshallEnumString(ident, "ArtistDominantBaselineType", diagram.ArtistDominantBaselineType.ToCodeString())
	case "ArtistFontSize":
		res = __gong__marshallString(ident, "ArtistFontSize", diagram.ArtistFontSize)
	case "MajorArtistFontSize":
		res = __gong__marshallString(ident, "MajorArtistFontSize", diagram.MajorArtistFontSize)
	case "MinorArtistFontSize":
		res = __gong__marshallString(ident, "MinorArtistFontSize", diagram.MinorArtistFontSize)
	case "ArtistFontWeigth":
		res = __gong__marshallString(ident, "ArtistFontWeigth", diagram.ArtistFontWeigth)
	case "ArtistFontFamily":
		res = __gong__marshallString(ident, "ArtistFontFamily", diagram.ArtistFontFamily)
	case "ArtistLetterSpacing":
		res = __gong__marshallString(ident, "ArtistLetterSpacing", diagram.ArtistLetterSpacing)
	case "ArtistDateRectAnchorType":
		res = __gong__marshallEnumString(ident, "ArtistDateRectAnchorType", diagram.ArtistDateRectAnchorType.ToCodeString())
	case "ArtistDateTextAnchorType":
		res = __gong__marshallEnumString(ident, "ArtistDateTextAnchorType", diagram.ArtistDateTextAnchorType.ToCodeString())
	case "ArtistDateDominantBaselineType":
		res = __gong__marshallEnumString(ident, "ArtistDateDominantBaselineType", diagram.ArtistDateDominantBaselineType.ToCodeString())
	case "ArtistDateAndPlacesFontSize":
		res = __gong__marshallString(ident, "ArtistDateAndPlacesFontSize", diagram.ArtistDateAndPlacesFontSize)
	case "ArtistDateAndPlacesFontWeigth":
		res = __gong__marshallString(ident, "ArtistDateAndPlacesFontWeigth", diagram.ArtistDateAndPlacesFontWeigth)
	case "ArtistDateAndPlacesFontFamily":
		res = __gong__marshallString(ident, "ArtistDateAndPlacesFontFamily", diagram.ArtistDateAndPlacesFontFamily)
	case "ArtistDateAndPlacesLetterSpacing":
		res = __gong__marshallString(ident, "ArtistDateAndPlacesLetterSpacing", diagram.ArtistDateAndPlacesLetterSpacing)
	case "ArtistPlacesRectAnchorType":
		res = __gong__marshallEnumString(ident, "ArtistPlacesRectAnchorType", diagram.ArtistPlacesRectAnchorType.ToCodeString())
	case "ArtistPlacesTextAnchorType":
		res = __gong__marshallEnumString(ident, "ArtistPlacesTextAnchorType", diagram.ArtistPlacesTextAnchorType.ToCodeString())
	case "ArtistPlacesDominantBaselineType":
		res = __gong__marshallEnumString(ident, "ArtistPlacesDominantBaselineType", diagram.ArtistPlacesDominantBaselineType.ToCodeString())
	case "InfluenceArrowSize":
		res = __gong__marshallFloat(ident, "InfluenceArrowSize", diagram.InfluenceArrowSize)
	case "InfluenceArrowStartOffset":
		res = __gong__marshallFloat(ident, "InfluenceArrowStartOffset", diagram.InfluenceArrowStartOffset)
	case "InfluenceArrowEndOffset":
		res = __gong__marshallFloat(ident, "InfluenceArrowEndOffset", diagram.InfluenceArrowEndOffset)
	case "InfluenceCornerRadius":
		res = __gong__marshallFloat(ident, "InfluenceCornerRadius", diagram.InfluenceCornerRadius)
	case "InfluenceDashedLinePattern":
		res = __gong__marshallString(ident, "InfluenceDashedLinePattern", diagram.InfluenceDashedLinePattern)

	case "MovementShapes":
		var sb strings.Builder
		for _, _movementshape := range diagram.MovementShapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "MovementShapes", _movementshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ArtefactTypeShapes":
		var sb strings.Builder
		for _, _artefacttypeshape := range diagram.ArtefactTypeShapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ArtefactTypeShapes", _artefacttypeshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ArtistShapes":
		var sb strings.Builder
		for _, _artistshape := range diagram.ArtistShapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ArtistShapes", _artistshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "InfluenceShapes":
		var sb strings.Builder
		for _, _influenceshape := range diagram.InfluenceShapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "InfluenceShapes", _influenceshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Diagram", fieldName)
	}
	return
}

func (influence *Influence) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := influence.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", influence.Name)
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", influence.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", influence.IsExpanded)
	case "IsHypothtical":
		res = __gong__marshallBool(ident, "IsHypothtical", influence.IsHypothtical)

	case "SourceMovement":
		if influence.SourceMovement != nil {
			res = __gong__marshallPointer(ident, "SourceMovement", influence.SourceMovement.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "SourceMovement", "nil")
		}
	case "SourceArtefactType":
		if influence.SourceArtefactType != nil {
			res = __gong__marshallPointer(ident, "SourceArtefactType", influence.SourceArtefactType.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "SourceArtefactType", "nil")
		}
	case "SourceArtist":
		if influence.SourceArtist != nil {
			res = __gong__marshallPointer(ident, "SourceArtist", influence.SourceArtist.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "SourceArtist", "nil")
		}
	case "TargetMovement":
		if influence.TargetMovement != nil {
			res = __gong__marshallPointer(ident, "TargetMovement", influence.TargetMovement.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "TargetMovement", "nil")
		}
	case "TargetArtefactType":
		if influence.TargetArtefactType != nil {
			res = __gong__marshallPointer(ident, "TargetArtefactType", influence.TargetArtefactType.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "TargetArtefactType", "nil")
		}
	case "TargetArtist":
		if influence.TargetArtist != nil {
			res = __gong__marshallPointer(ident, "TargetArtist", influence.TargetArtist.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "TargetArtist", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Influence", fieldName)
	}
	return
}

func (influenceshape *InfluenceShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := influenceshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", influenceshape.Name)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", influenceshape.IsHidden)

	case "Influence":
		if influenceshape.Influence != nil {
			res = __gong__marshallPointer(ident, "Influence", influenceshape.Influence.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Influence", "nil")
		}
	case "ControlPointShapes":
		var sb strings.Builder
		for _, _controlpointshape := range influenceshape.ControlPointShapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ControlPointShapes", _controlpointshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct InfluenceShape", fieldName)
	}
	return
}

func (library *Library) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := library.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", library.Name)
	case "Description":
		res = __gong__marshallString(ident, "Description", library.Description)
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", library.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", library.IsExpanded)
	case "IsRootLibrary":
		res = __gong__marshallBool(ident, "IsRootLibrary", library.IsRootLibrary)
	case "IsSubLibrariesNodeExpanded":
		res = __gong__marshallBool(ident, "IsSubLibrariesNodeExpanded", library.IsSubLibrariesNodeExpanded)
	case "NbPixPerCharacter":
		res = __gong__marshallFloat(ident, "NbPixPerCharacter", library.NbPixPerCharacter)
	case "LogoSVGFile":
		res = __gong__marshallString(ident, "LogoSVGFile", library.LogoSVGFile)
	case "IsExpandedTmp":
		res = __gong__marshallBool(ident, "IsExpandedTmp", library.IsExpandedTmp)

	case "SubLibraries":
		var sb strings.Builder
		for _, _library := range library.SubLibraries {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SubLibraries", _library.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "SubLibrariesWhoseNodeIsExpanded":
		var sb strings.Builder
		for _, _library := range library.SubLibrariesWhoseNodeIsExpanded {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SubLibrariesWhoseNodeIsExpanded", _library.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Library", fieldName)
	}
	return
}

func (movement *Movement) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := movement.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", movement.Name)
	case "ComputedPrefix":
		res = __gong__marshallString(ident, "ComputedPrefix", movement.ComputedPrefix)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", movement.IsExpanded)
	case "Date":
		res = __gong__marshallTime(ident, "Date", movement.Date.String())
	case "HideDate":
		res = __gong__marshallBool(ident, "HideDate", movement.HideDate)
	case "HasTaxonomicFilter":
		res = __gong__marshallBool(ident, "HasTaxonomicFilter", movement.HasTaxonomicFilter)
	case "TaxonomicFilter":
		res = __gong__marshallString(ident, "TaxonomicFilter", movement.TaxonomicFilter)
	case "IsFeatured":
		res = __gong__marshallBool(ident, "IsFeatured", movement.IsFeatured)
	case "FeaturePrefix":
		res = __gong__marshallString(ident, "FeaturePrefix", movement.FeaturePrefix)
	case "IsMajor":
		res = __gong__marshallBool(ident, "IsMajor", movement.IsMajor)
	case "IsMinor":
		res = __gong__marshallBool(ident, "IsMinor", movement.IsMinor)
	case "AdditionnalName":
		res = __gong__marshallString(ident, "AdditionnalName", movement.AdditionnalName)

	case "Places":
		var sb strings.Builder
		for _, _place := range movement.Places {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Places", _place.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Movement", fieldName)
	}
	return
}

func (movementshape *MovementShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := movementshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", movementshape.Name)
	case "X":
		res = __gong__marshallFloat(ident, "X", movementshape.X)
	case "Y":
		res = __gong__marshallFloat(ident, "Y", movementshape.Y)
	case "Width":
		res = __gong__marshallFloat(ident, "Width", movementshape.Width)
	case "Height":
		res = __gong__marshallFloat(ident, "Height", movementshape.Height)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", movementshape.IsHidden)

	case "Movement":
		if movementshape.Movement != nil {
			res = __gong__marshallPointer(ident, "Movement", movementshape.Movement.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Movement", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct MovementShape", fieldName)
	}
	return
}

func (place *Place) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := place.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", place.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct Place", fieldName)
	}
	return
}

// insertion point for marshall all fields methods
func (artefacttype *ArtefactType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(artefacttype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(artefacttype.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(artefacttype.GongMarshallField(stage, "IsExpanded"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (artefacttypeshape *ArtefactTypeShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(artefacttypeshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(artefacttypeshape.GongMarshallField(stage, "ArtefactType"))
		initializerStatements.WriteString(artefacttypeshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(artefacttypeshape.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(artefacttypeshape.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(artefacttypeshape.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(artefacttypeshape.GongMarshallField(stage, "IsHidden"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (artist *Artist) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(artist.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(artist.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(artist.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(artist.GongMarshallField(stage, "IsDead"))
		initializerStatements.WriteString(artist.GongMarshallField(stage, "DateOfDeath"))
		pointersInitializesStatements.WriteString(artist.GongMarshallField(stage, "Place"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (artistshape *ArtistShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(artistshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(artistshape.GongMarshallField(stage, "Artist"))
		initializerStatements.WriteString(artistshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(artistshape.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(artistshape.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(artistshape.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(artistshape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(artistshape.GongMarshallField(stage, "ImagePng_X"))
		initializerStatements.WriteString(artistshape.GongMarshallField(stage, "ImagePng_Y"))
		initializerStatements.WriteString(artistshape.GongMarshallField(stage, "ImagePng_Width"))
		initializerStatements.WriteString(artistshape.GongMarshallField(stage, "ImagePng_Height"))
		initializerStatements.WriteString(artistshape.GongMarshallField(stage, "ImagePng_X_Offset"))
		initializerStatements.WriteString(artistshape.GongMarshallField(stage, "ImagePng_Y_Offset"))
		initializerStatements.WriteString(artistshape.GongMarshallField(stage, "ImagePng_RectAnchorType"))
		initializerStatements.WriteString(artistshape.GongMarshallField(stage, "ImagePngBase64Content"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (controlpointshape *ControlPointShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(controlpointshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(controlpointshape.GongMarshallField(stage, "X_Relative"))
		initializerStatements.WriteString(controlpointshape.GongMarshallField(stage, "Y_Relative"))
		initializerStatements.WriteString(controlpointshape.GongMarshallField(stage, "IsStartShapeTheClosestShape"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (desk *Desk) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(desk.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(desk.GongMarshallField(stage, "SelectedDiagram"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (diagram *Diagram) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsChecked"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "MovementShapes"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "ArtefactTypeShapes"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "ArtistShapes"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "InfluenceShapes"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsEditable"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsNodeExpanded"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsMovementCategoryNodeExpanded"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsArtefactTypeCategoryNodeExpanded"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsArtistCategoryNodeExpanded"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsInfluenceCategoryNodeExpanded"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsMovementCategoryHidden"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsArtefactTypeCategoryHidden"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsArtistCategoryHidden"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsInfluenceCategoryHidden"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "StartDate"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "EndDate"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "NbYearsForIntervals"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "XMargin"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "YMargin"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "NextVerticalDateXMargin"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "RedColorCode"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "BackgroundGreyColorCode"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "GrayColorCode"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "BottomBoxYOffset"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "BottomBoxWidth"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "BottomBoxHeigth"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "BottomBoxFontSize"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "BottomBoxFontWeigth"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "BottomBoxFontFamily"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "BottomBoxLetterSpacing"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "BottomBoxLetterColorCode"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MovementRectAnchorType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MovementTextAnchorType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MovementDominantBaselineType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MovementFontSize"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MajorMovementFontSize"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MinorMovementFontSize"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MovementFontWeigth"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MovementFontFamily"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MovementLetterSpacing"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "AbstractMovementFontSize"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "AbstractMovementRectAnchorType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "AbstractMovementTextAnchorType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "AbstractDominantBaselineType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MovementDateRectAnchorType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MovementDateTextAnchorType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MovementDateTextDominantBaselineType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MovementDateAndPlacesFontSize"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MovementDateAndPlacesFontWeigth"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MovementDateAndPlacesFontFamily"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MovementDateAndPlacesLetterSpacing"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MovementBelowArcY_Offset"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MovementBelowArcY_OffsetPerPlace"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MovementPlacesRectAnchorType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MovementPlacesTextAnchorType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MovementPlacesDominantBaselineType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtefactTypeFontSize"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtefactTypeFontWeigth"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtefactTypeFontFamily"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtefactTypeLetterSpacing"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtefactTypeRectAnchorType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtefactDominantBaselineType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtefactTypeStrokeWidth"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtistRectAnchorType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtistTextAnchorType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtistDominantBaselineType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtistFontSize"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MajorArtistFontSize"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "MinorArtistFontSize"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtistFontWeigth"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtistFontFamily"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtistLetterSpacing"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtistDateRectAnchorType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtistDateTextAnchorType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtistDateDominantBaselineType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtistDateAndPlacesFontSize"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtistDateAndPlacesFontWeigth"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtistDateAndPlacesFontFamily"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtistDateAndPlacesLetterSpacing"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtistPlacesRectAnchorType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtistPlacesTextAnchorType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "ArtistPlacesDominantBaselineType"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "InfluenceArrowSize"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "InfluenceArrowStartOffset"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "InfluenceArrowEndOffset"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "InfluenceCornerRadius"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "InfluenceDashedLinePattern"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (influence *Influence) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(influence.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(influence.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(influence.GongMarshallField(stage, "IsExpanded"))
		pointersInitializesStatements.WriteString(influence.GongMarshallField(stage, "SourceMovement"))
		pointersInitializesStatements.WriteString(influence.GongMarshallField(stage, "SourceArtefactType"))
		pointersInitializesStatements.WriteString(influence.GongMarshallField(stage, "SourceArtist"))
		pointersInitializesStatements.WriteString(influence.GongMarshallField(stage, "TargetMovement"))
		pointersInitializesStatements.WriteString(influence.GongMarshallField(stage, "TargetArtefactType"))
		pointersInitializesStatements.WriteString(influence.GongMarshallField(stage, "TargetArtist"))
		initializerStatements.WriteString(influence.GongMarshallField(stage, "IsHypothtical"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (influenceshape *InfluenceShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(influenceshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(influenceshape.GongMarshallField(stage, "Influence"))
		initializerStatements.WriteString(influenceshape.GongMarshallField(stage, "IsHidden"))
		pointersInitializesStatements.WriteString(influenceshape.GongMarshallField(stage, "ControlPointShapes"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (library *Library) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(library.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsRootLibrary"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "SubLibraries"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsSubLibrariesNodeExpanded"))
		pointersInitializesStatements.WriteString(library.GongMarshallField(stage, "SubLibrariesWhoseNodeIsExpanded"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "NbPixPerCharacter"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "LogoSVGFile"))
		initializerStatements.WriteString(library.GongMarshallField(stage, "IsExpandedTmp"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (movement *Movement) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(movement.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(movement.GongMarshallField(stage, "ComputedPrefix"))
		initializerStatements.WriteString(movement.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(movement.GongMarshallField(stage, "Date"))
		initializerStatements.WriteString(movement.GongMarshallField(stage, "HideDate"))
		pointersInitializesStatements.WriteString(movement.GongMarshallField(stage, "Places"))
		initializerStatements.WriteString(movement.GongMarshallField(stage, "HasTaxonomicFilter"))
		initializerStatements.WriteString(movement.GongMarshallField(stage, "TaxonomicFilter"))
		initializerStatements.WriteString(movement.GongMarshallField(stage, "IsFeatured"))
		initializerStatements.WriteString(movement.GongMarshallField(stage, "FeaturePrefix"))
		initializerStatements.WriteString(movement.GongMarshallField(stage, "IsMajor"))
		initializerStatements.WriteString(movement.GongMarshallField(stage, "IsMinor"))
		initializerStatements.WriteString(movement.GongMarshallField(stage, "AdditionnalName"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (movementshape *MovementShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(movementshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(movementshape.GongMarshallField(stage, "Movement"))
		initializerStatements.WriteString(movementshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(movementshape.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(movementshape.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(movementshape.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(movementshape.GongMarshallField(stage, "IsHidden"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (place *Place) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(place.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
