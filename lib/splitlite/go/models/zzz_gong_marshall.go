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
	modelsPackageName := "github.com/fullstack-lang/gong/lib/splitlite/go/models"
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
		for _, assplit := range __gong__sortStageSetInstances(stage.AsSplits, stage.AsSplit_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			assplitIdent := assplit.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.AsSplit{Name: %s}).Stage(stage)", assplitIdent, __gong__toRawStringLiteral(assplit.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", assplitIdent, __gong__toRawStringLiteral(assplit.Name)))
			if assplit.Direction.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.Direction = models.%s", assplitIdent, assplit.Direction.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.Direction = \"\"", assplitIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.IsSizeInPixel = %t", assplitIdent, assplit.IsSizeInPixel))
			values.WriteString(fmt.Sprintf("\n\t%s.IsWithCustomGutterSize = %t", assplitIdent, assplit.IsWithCustomGutterSize))
			values.WriteString(fmt.Sprintf("\n\t%s.GutterSize = %f", assplitIdent, assplit.GutterSize))
			for _, elem := range assplit.AsSplitAreas {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AsSplitAreas = append(%s.AsSplitAreas, %s)", assplitIdent, assplitIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, assplitarea := range __gong__sortStageSetInstances(stage.AsSplitAreas, stage.AsSplitArea_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			assplitareaIdent := assplitarea.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.AsSplitArea{Name: %s}).Stage(stage)", assplitareaIdent, __gong__toRawStringLiteral(assplitarea.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", assplitareaIdent, __gong__toRawStringLiteral(assplitarea.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowNameInHeader = %t", assplitareaIdent, assplitarea.ShowNameInHeader))
			values.WriteString(fmt.Sprintf("\n\t%s.Size = %f", assplitareaIdent, assplitarea.Size))
			values.WriteString(fmt.Sprintf("\n\t%s.IsAny = %t", assplitareaIdent, assplitarea.IsAny))
			values.WriteString(fmt.Sprintf("\n\t%s.HasDiv = %t", assplitareaIdent, assplitarea.HasDiv))
			values.WriteString(fmt.Sprintf("\n\t%s.DivStyle = %s", assplitareaIdent, __gong__toRawStringLiteral(assplitarea.DivStyle)))
			if assplitarea.AsSplit != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := assplitarea.AsSplit.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AsSplit = %s", assplitareaIdent, targetIdent))
			}
			if assplitarea.Button != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := assplitarea.Button.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Button = %s", assplitareaIdent, targetIdent))
			}
			if assplitarea.Form != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := assplitarea.Form.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Form = %s", assplitareaIdent, targetIdent))
			}
			if assplitarea.Load != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := assplitarea.Load.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Load = %s", assplitareaIdent, targetIdent))
			}
			if assplitarea.Split != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := assplitarea.Split.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Split = %s", assplitareaIdent, targetIdent))
			}
			if assplitarea.Svg != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := assplitarea.Svg.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Svg = %s", assplitareaIdent, targetIdent))
			}
			if assplitarea.Table != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := assplitarea.Table.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Table = %s", assplitareaIdent, targetIdent))
			}
			if assplitarea.Tree != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := assplitarea.Tree.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Tree = %s", assplitareaIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, button := range __gong__sortStageSetInstances(stage.Buttons, stage.Button_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			buttonIdent := button.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Button{Name: %s}).Stage(stage)", buttonIdent, __gong__toRawStringLiteral(button.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", buttonIdent, __gong__toRawStringLiteral(button.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StackName = %s", buttonIdent, __gong__toRawStringLiteral(button.StackName)))
		}
	}
	if stage != nil {
		for _, favicon := range __gong__sortStageSetInstances(stage.FavIcons, stage.FavIcon_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			faviconIdent := favicon.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.FavIcon{Name: %s}).Stage(stage)", faviconIdent, __gong__toRawStringLiteral(favicon.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", faviconIdent, __gong__toRawStringLiteral(favicon.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.SVG = %s", faviconIdent, __gong__toRawStringLiteral(favicon.SVG)))
		}
	}
	if stage != nil {
		for _, form := range __gong__sortStageSetInstances(stage.Forms, stage.Form_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			formIdent := form.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Form{Name: %s}).Stage(stage)", formIdent, __gong__toRawStringLiteral(form.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", formIdent, __gong__toRawStringLiteral(form.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StackName = %s", formIdent, __gong__toRawStringLiteral(form.StackName)))
		}
	}
	if stage != nil {
		for _, load := range __gong__sortStageSetInstances(stage.Loads, stage.Load_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			loadIdent := load.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Load{Name: %s}).Stage(stage)", loadIdent, __gong__toRawStringLiteral(load.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", loadIdent, __gong__toRawStringLiteral(load.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StackName = %s", loadIdent, __gong__toRawStringLiteral(load.StackName)))
		}
	}
	if stage != nil {
		for _, logoontheleft := range __gong__sortStageSetInstances(stage.LogoOnTheLefts, stage.LogoOnTheLeft_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			logoontheleftIdent := logoontheleft.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.LogoOnTheLeft{Name: %s}).Stage(stage)", logoontheleftIdent, __gong__toRawStringLiteral(logoontheleft.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", logoontheleftIdent, __gong__toRawStringLiteral(logoontheleft.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %d", logoontheleftIdent, logoontheleft.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %d", logoontheleftIdent, logoontheleft.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.SVG = %s", logoontheleftIdent, __gong__toRawStringLiteral(logoontheleft.SVG)))
		}
	}
	if stage != nil {
		for _, logoontheright := range __gong__sortStageSetInstances(stage.LogoOnTheRights, stage.LogoOnTheRight_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			logoontherightIdent := logoontheright.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.LogoOnTheRight{Name: %s}).Stage(stage)", logoontherightIdent, __gong__toRawStringLiteral(logoontheright.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", logoontherightIdent, __gong__toRawStringLiteral(logoontheright.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %d", logoontherightIdent, logoontheright.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %d", logoontherightIdent, logoontheright.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.SVG = %s", logoontherightIdent, __gong__toRawStringLiteral(logoontheright.SVG)))
		}
	}
	if stage != nil {
		for _, split := range __gong__sortStageSetInstances(stage.Splits, stage.Split_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			splitIdent := split.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Split{Name: %s}).Stage(stage)", splitIdent, __gong__toRawStringLiteral(split.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", splitIdent, __gong__toRawStringLiteral(split.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StackName = %s", splitIdent, __gong__toRawStringLiteral(split.StackName)))
		}
	}
	if stage != nil {
		for _, svg := range __gong__sortStageSetInstances(stage.Svgs, stage.Svg_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			svgIdent := svg.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Svg{Name: %s}).Stage(stage)", svgIdent, __gong__toRawStringLiteral(svg.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", svgIdent, __gong__toRawStringLiteral(svg.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StackName = %s", svgIdent, __gong__toRawStringLiteral(svg.StackName)))
			values.WriteString(fmt.Sprintf("\n\t%s.Style = %s", svgIdent, __gong__toRawStringLiteral(svg.Style)))
		}
	}
	if stage != nil {
		for _, table := range __gong__sortStageSetInstances(stage.Tables, stage.Table_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			tableIdent := table.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Table{Name: %s}).Stage(stage)", tableIdent, __gong__toRawStringLiteral(table.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", tableIdent, __gong__toRawStringLiteral(table.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StackName = %s", tableIdent, __gong__toRawStringLiteral(table.StackName)))
		}
	}
	if stage != nil {
		for _, title := range __gong__sortStageSetInstances(stage.Titles, stage.Title_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			titleIdent := title.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Title{Name: %s}).Stage(stage)", titleIdent, __gong__toRawStringLiteral(title.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", titleIdent, __gong__toRawStringLiteral(title.Name)))
		}
	}
	if stage != nil {
		for _, tree := range __gong__sortStageSetInstances(stage.Trees, stage.Tree_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			treeIdent := tree.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Tree{Name: %s}).Stage(stage)", treeIdent, __gong__toRawStringLiteral(tree.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", treeIdent, __gong__toRawStringLiteral(tree.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StackName = %s", treeIdent, __gong__toRawStringLiteral(tree.StackName)))
		}
	}
	if stage != nil {
		for _, view := range __gong__sortStageSetInstances(stage.Views, stage.View_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			viewIdent := view.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.View{Name: %s}).Stage(stage)", viewIdent, __gong__toRawStringLiteral(view.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", viewIdent, __gong__toRawStringLiteral(view.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowViewName = %t", viewIdent, view.ShowViewName))
			values.WriteString(fmt.Sprintf("\n\t%s.IsSelectedView = %t", viewIdent, view.IsSelectedView))
			if view.Direction.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.Direction = models.%s", viewIdent, view.Direction.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.Direction = \"\"", viewIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.IsSecondaryView = %t", viewIdent, view.IsSecondaryView))
			values.WriteString(fmt.Sprintf("\n\t%s.IsSizeInPixel = %t", viewIdent, view.IsSizeInPixel))
			values.WriteString(fmt.Sprintf("\n\t%s.IsWithCustomGutterSize = %t", viewIdent, view.IsWithCustomGutterSize))
			values.WriteString(fmt.Sprintf("\n\t%s.GutterSize = %f", viewIdent, view.GutterSize))
			for _, elem := range view.RootAsSplitAreas {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootAsSplitAreas = append(%s.RootAsSplitAreas, %s)", viewIdent, viewIdent, targetIdent))
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
go run github.com/fullstack-lang/gong/lib/splitlite/go/models/probe/cmd/probe@%s stage.go
*/
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/lib/splitlite/go/models"%s
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

	if modelsPackageName != "" && modelsPackageName != "github.com/fullstack-lang/gong/lib/splitlite/go/models" {
		res = strings.ReplaceAll(res, "\"github.com/fullstack-lang/gong/lib/splitlite/go/models\"", fmt.Sprintf("%q", modelsPackageName))
	}

	return res, nil
}

// insertion point for marshall field methods
func (assplit *AsSplit) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := assplit.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", assplit.Name)
	case "Direction":
		res = __gong__marshallEnumString(ident, "Direction", assplit.Direction.ToCodeString())
	case "IsSizeInPixel":
		res = __gong__marshallBool(ident, "IsSizeInPixel", assplit.IsSizeInPixel)
	case "IsWithCustomGutterSize":
		res = __gong__marshallBool(ident, "IsWithCustomGutterSize", assplit.IsWithCustomGutterSize)
	case "GutterSize":
		res = __gong__marshallFloat(ident, "GutterSize", assplit.GutterSize)

	case "AsSplitAreas":
		var sb strings.Builder
		for _, _assplitarea := range assplit.AsSplitAreas {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "AsSplitAreas", _assplitarea.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct AsSplit", fieldName)
	}
	return
}

func (assplitarea *AsSplitArea) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := assplitarea.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", assplitarea.Name)
	case "ShowNameInHeader":
		res = __gong__marshallBool(ident, "ShowNameInHeader", assplitarea.ShowNameInHeader)
	case "Size":
		res = __gong__marshallFloat(ident, "Size", assplitarea.Size)
	case "IsAny":
		res = __gong__marshallBool(ident, "IsAny", assplitarea.IsAny)
	case "HasDiv":
		res = __gong__marshallBool(ident, "HasDiv", assplitarea.HasDiv)
	case "DivStyle":
		res = __gong__marshallString(ident, "DivStyle", assplitarea.DivStyle)

	case "AsSplit":
		if assplitarea.AsSplit != nil {
			res = __gong__marshallPointer(ident, "AsSplit", assplitarea.AsSplit.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "AsSplit", "nil")
		}
	case "Button":
		if assplitarea.Button != nil {
			res = __gong__marshallPointer(ident, "Button", assplitarea.Button.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Button", "nil")
		}
	case "Form":
		if assplitarea.Form != nil {
			res = __gong__marshallPointer(ident, "Form", assplitarea.Form.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Form", "nil")
		}
	case "Load":
		if assplitarea.Load != nil {
			res = __gong__marshallPointer(ident, "Load", assplitarea.Load.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Load", "nil")
		}
	case "Split":
		if assplitarea.Split != nil {
			res = __gong__marshallPointer(ident, "Split", assplitarea.Split.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Split", "nil")
		}
	case "Svg":
		if assplitarea.Svg != nil {
			res = __gong__marshallPointer(ident, "Svg", assplitarea.Svg.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Svg", "nil")
		}
	case "Table":
		if assplitarea.Table != nil {
			res = __gong__marshallPointer(ident, "Table", assplitarea.Table.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Table", "nil")
		}
	case "Tree":
		if assplitarea.Tree != nil {
			res = __gong__marshallPointer(ident, "Tree", assplitarea.Tree.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Tree", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct AsSplitArea", fieldName)
	}
	return
}

func (button *Button) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := button.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", button.Name)
	case "StackName":
		res = __gong__marshallString(ident, "StackName", button.StackName)

	default:
		log.Panicf("Unknown field %s for Gongstruct Button", fieldName)
	}
	return
}

func (favicon *FavIcon) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := favicon.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", favicon.Name)
	case "SVG":
		res = __gong__marshallString(ident, "SVG", favicon.SVG)

	default:
		log.Panicf("Unknown field %s for Gongstruct FavIcon", fieldName)
	}
	return
}

func (form *Form) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := form.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", form.Name)
	case "StackName":
		res = __gong__marshallString(ident, "StackName", form.StackName)

	default:
		log.Panicf("Unknown field %s for Gongstruct Form", fieldName)
	}
	return
}

func (load *Load) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := load.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", load.Name)
	case "StackName":
		res = __gong__marshallString(ident, "StackName", load.StackName)

	default:
		log.Panicf("Unknown field %s for Gongstruct Load", fieldName)
	}
	return
}

func (logoontheleft *LogoOnTheLeft) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := logoontheleft.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", logoontheleft.Name)
	case "Width":
		res = __gong__marshallInt(ident, "Width", logoontheleft.Width)
	case "Height":
		res = __gong__marshallInt(ident, "Height", logoontheleft.Height)
	case "SVG":
		res = __gong__marshallString(ident, "SVG", logoontheleft.SVG)

	default:
		log.Panicf("Unknown field %s for Gongstruct LogoOnTheLeft", fieldName)
	}
	return
}

func (logoontheright *LogoOnTheRight) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := logoontheright.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", logoontheright.Name)
	case "Width":
		res = __gong__marshallInt(ident, "Width", logoontheright.Width)
	case "Height":
		res = __gong__marshallInt(ident, "Height", logoontheright.Height)
	case "SVG":
		res = __gong__marshallString(ident, "SVG", logoontheright.SVG)

	default:
		log.Panicf("Unknown field %s for Gongstruct LogoOnTheRight", fieldName)
	}
	return
}

func (split *Split) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := split.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", split.Name)
	case "StackName":
		res = __gong__marshallString(ident, "StackName", split.StackName)

	default:
		log.Panicf("Unknown field %s for Gongstruct Split", fieldName)
	}
	return
}

func (svg *Svg) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := svg.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", svg.Name)
	case "StackName":
		res = __gong__marshallString(ident, "StackName", svg.StackName)
	case "Style":
		res = __gong__marshallString(ident, "Style", svg.Style)

	default:
		log.Panicf("Unknown field %s for Gongstruct Svg", fieldName)
	}
	return
}

func (table *Table) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := table.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", table.Name)
	case "StackName":
		res = __gong__marshallString(ident, "StackName", table.StackName)

	default:
		log.Panicf("Unknown field %s for Gongstruct Table", fieldName)
	}
	return
}

func (title *Title) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := title.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", title.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct Title", fieldName)
	}
	return
}

func (tree *Tree) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := tree.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", tree.Name)
	case "StackName":
		res = __gong__marshallString(ident, "StackName", tree.StackName)

	default:
		log.Panicf("Unknown field %s for Gongstruct Tree", fieldName)
	}
	return
}

func (view *View) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := view.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", view.Name)
	case "ShowViewName":
		res = __gong__marshallBool(ident, "ShowViewName", view.ShowViewName)
	case "IsSelectedView":
		res = __gong__marshallBool(ident, "IsSelectedView", view.IsSelectedView)
	case "Direction":
		res = __gong__marshallEnumString(ident, "Direction", view.Direction.ToCodeString())
	case "IsSecondaryView":
		res = __gong__marshallBool(ident, "IsSecondaryView", view.IsSecondaryView)
	case "IsSizeInPixel":
		res = __gong__marshallBool(ident, "IsSizeInPixel", view.IsSizeInPixel)
	case "IsWithCustomGutterSize":
		res = __gong__marshallBool(ident, "IsWithCustomGutterSize", view.IsWithCustomGutterSize)
	case "GutterSize":
		res = __gong__marshallFloat(ident, "GutterSize", view.GutterSize)

	case "RootAsSplitAreas":
		var sb strings.Builder
		for _, _assplitarea := range view.RootAsSplitAreas {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "RootAsSplitAreas", _assplitarea.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct View", fieldName)
	}
	return
}

// insertion point for marshall all fields methods
func (assplit *AsSplit) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(assplit.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(assplit.GongMarshallField(stage, "Direction"))
		pointersInitializesStatements.WriteString(assplit.GongMarshallField(stage, "AsSplitAreas"))
		initializerStatements.WriteString(assplit.GongMarshallField(stage, "IsSizeInPixel"))
		initializerStatements.WriteString(assplit.GongMarshallField(stage, "IsWithCustomGutterSize"))
		initializerStatements.WriteString(assplit.GongMarshallField(stage, "GutterSize"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (assplitarea *AsSplitArea) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(assplitarea.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(assplitarea.GongMarshallField(stage, "ShowNameInHeader"))
		initializerStatements.WriteString(assplitarea.GongMarshallField(stage, "Size"))
		initializerStatements.WriteString(assplitarea.GongMarshallField(stage, "IsAny"))
		pointersInitializesStatements.WriteString(assplitarea.GongMarshallField(stage, "AsSplit"))
		pointersInitializesStatements.WriteString(assplitarea.GongMarshallField(stage, "Button"))
		pointersInitializesStatements.WriteString(assplitarea.GongMarshallField(stage, "Form"))
		pointersInitializesStatements.WriteString(assplitarea.GongMarshallField(stage, "Load"))
		pointersInitializesStatements.WriteString(assplitarea.GongMarshallField(stage, "Split"))
		pointersInitializesStatements.WriteString(assplitarea.GongMarshallField(stage, "Svg"))
		pointersInitializesStatements.WriteString(assplitarea.GongMarshallField(stage, "Table"))
		pointersInitializesStatements.WriteString(assplitarea.GongMarshallField(stage, "Tree"))
		initializerStatements.WriteString(assplitarea.GongMarshallField(stage, "HasDiv"))
		initializerStatements.WriteString(assplitarea.GongMarshallField(stage, "DivStyle"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (button *Button) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(button.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(button.GongMarshallField(stage, "StackName"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (favicon *FavIcon) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(favicon.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(favicon.GongMarshallField(stage, "SVG"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (form *Form) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(form.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(form.GongMarshallField(stage, "StackName"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (load *Load) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(load.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(load.GongMarshallField(stage, "StackName"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (logoontheleft *LogoOnTheLeft) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(logoontheleft.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(logoontheleft.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(logoontheleft.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(logoontheleft.GongMarshallField(stage, "SVG"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (logoontheright *LogoOnTheRight) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(logoontheright.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(logoontheright.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(logoontheright.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(logoontheright.GongMarshallField(stage, "SVG"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (split *Split) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(split.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(split.GongMarshallField(stage, "StackName"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (svg *Svg) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(svg.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(svg.GongMarshallField(stage, "StackName"))
		initializerStatements.WriteString(svg.GongMarshallField(stage, "Style"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (table *Table) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(table.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(table.GongMarshallField(stage, "StackName"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (title *Title) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(title.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (tree *Tree) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(tree.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(tree.GongMarshallField(stage, "StackName"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (view *View) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(view.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(view.GongMarshallField(stage, "ShowViewName"))
		pointersInitializesStatements.WriteString(view.GongMarshallField(stage, "RootAsSplitAreas"))
		initializerStatements.WriteString(view.GongMarshallField(stage, "IsSelectedView"))
		initializerStatements.WriteString(view.GongMarshallField(stage, "Direction"))
		initializerStatements.WriteString(view.GongMarshallField(stage, "IsSecondaryView"))
		initializerStatements.WriteString(view.GongMarshallField(stage, "IsSizeInPixel"))
		initializerStatements.WriteString(view.GongMarshallField(stage, "IsWithCustomGutterSize"))
		initializerStatements.WriteString(view.GongMarshallField(stage, "GutterSize"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
