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
	modelsPackageName := "github.com/fullstack-lang/gong/test/test1/go/models"
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
		for _, astruct := range __gong__sortStageSetInstances(stage.Astructs, stage.Astruct_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			astructIdent := astruct.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Astruct{Name: %s}).Stage(stage)", astructIdent, __gong__toRawStringLiteral(astruct.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", astructIdent, __gong__toRawStringLiteral(astruct.Name)))
			if str, ok := astruct.Field.(string); ok {
				values.WriteString(fmt.Sprintf("\n\t%s.Field = %s", astructIdent, str))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Date, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", astructIdent, astruct.Date.String()))
			values.WriteString(fmt.Sprintf("\n\t%s.Date2, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", astructIdent, astruct.Date2.String()))
			values.WriteString(fmt.Sprintf("\n\t%s.Booleanfield = %t", astructIdent, astruct.Booleanfield))
			if astruct.Aenum.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.Aenum = models.%s", astructIdent, astruct.Aenum.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.Aenum = \"\"", astructIdent))
			}
			if astruct.Aenum_2.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.Aenum_2 = models.%s", astructIdent, astruct.Aenum_2.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.Aenum_2 = \"\"", astructIdent))
			}
			if astruct.Benum.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.Benum = models.%s", astructIdent, astruct.Benum.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.Benum = \"\"", astructIdent))
			}
			if astruct.CEnum.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.CEnum = models.%s", astructIdent, astruct.CEnum.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.CEnum = 0", astructIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.CName = %s", astructIdent, __gong__toRawStringLiteral(astruct.CName)))
			values.WriteString(fmt.Sprintf("\n\t%s.CFloatfield = %f", astructIdent, astruct.CFloatfield))
			values.WriteString(fmt.Sprintf("\n\t%s.Floatfield = %f", astructIdent, astruct.Floatfield))
			values.WriteString(fmt.Sprintf("\n\t%s.Intfield = %d", astructIdent, astruct.Intfield))
			values.WriteString(fmt.Sprintf("\n\t%s.Anotherbooleanfield = %t", astructIdent, astruct.Anotherbooleanfield))
			values.WriteString(fmt.Sprintf("\n\t%s.Duration1 = time.Duration(%d)", astructIdent, int64(astruct.Duration1)))
			values.WriteString(fmt.Sprintf("\n\t%s.TextFieldBespokeSize = %s", astructIdent, __gong__toRawStringLiteral(astruct.TextFieldBespokeSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.TextArea = %s", astructIdent, __gong__toRawStringLiteral(astruct.TextArea)))
			if astruct.Associationtob != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := astruct.Associationtob.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Associationtob = %s", astructIdent, targetIdent))
			}
			for _, elem := range astruct.Anarrayofb {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Anarrayofb = append(%s.Anarrayofb, %s)", astructIdent, astructIdent, targetIdent))
			}
			if astruct.Anotherassociationtob_2 != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := astruct.Anotherassociationtob_2.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Anotherassociationtob_2 = %s", astructIdent, targetIdent))
			}
			if astruct.Bstruct != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := astruct.Bstruct.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Bstruct = %s", astructIdent, targetIdent))
			}
			if astruct.Bstruct2 != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := astruct.Bstruct2.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Bstruct2 = %s", astructIdent, targetIdent))
			}
			if astruct.Dstruct != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := astruct.Dstruct.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Dstruct = %s", astructIdent, targetIdent))
			}
			if astruct.Dstruct2 != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := astruct.Dstruct2.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Dstruct2 = %s", astructIdent, targetIdent))
			}
			if astruct.Dstruct3 != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := astruct.Dstruct3.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Dstruct3 = %s", astructIdent, targetIdent))
			}
			if astruct.Dstruct4 != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := astruct.Dstruct4.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Dstruct4 = %s", astructIdent, targetIdent))
			}
			for _, elem := range astruct.Dstruct4s {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Dstruct4s = append(%s.Dstruct4s, %s)", astructIdent, astructIdent, targetIdent))
			}
			for _, elem := range astruct.Anarrayofa {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Anarrayofa = append(%s.Anarrayofa, %s)", astructIdent, astructIdent, targetIdent))
			}
			for _, elem := range astruct.Anotherarrayofb {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Anotherarrayofb = append(%s.Anotherarrayofb, %s)", astructIdent, astructIdent, targetIdent))
			}
			for _, elem := range astruct.AnarrayofbUse {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AnarrayofbUse = append(%s.AnarrayofbUse, %s)", astructIdent, astructIdent, targetIdent))
			}
			for _, elem := range astruct.Anarrayofb2Use {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Anarrayofb2Use = append(%s.Anarrayofb2Use, %s)", astructIdent, astructIdent, targetIdent))
			}
			if astruct.AnAstruct != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := astruct.AnAstruct.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AnAstruct = %s", astructIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, astructbstruct2use := range __gong__sortStageSetInstances(stage.AstructBstruct2Uses, stage.AstructBstruct2Use_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			astructbstruct2useIdent := astructbstruct2use.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.AstructBstruct2Use{Name: %s}).Stage(stage)", astructbstruct2useIdent, __gong__toRawStringLiteral(astructbstruct2use.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", astructbstruct2useIdent, __gong__toRawStringLiteral(astructbstruct2use.Name)))
			if astructbstruct2use.Bstrcut2 != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := astructbstruct2use.Bstrcut2.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Bstrcut2 = %s", astructbstruct2useIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, astructbstructuse := range __gong__sortStageSetInstances(stage.AstructBstructUses, stage.AstructBstructUse_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			astructbstructuseIdent := astructbstructuse.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.AstructBstructUse{Name: %s}).Stage(stage)", astructbstructuseIdent, __gong__toRawStringLiteral(astructbstructuse.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", astructbstructuseIdent, __gong__toRawStringLiteral(astructbstructuse.Name)))
			if astructbstructuse.Bstruct2 != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := astructbstructuse.Bstruct2.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Bstruct2 = %s", astructbstructuseIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, bstruct := range __gong__sortStageSetInstances(stage.Bstructs, stage.Bstruct_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			bstructIdent := bstruct.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Bstruct{Name: %s}).Stage(stage)", bstructIdent, __gong__toRawStringLiteral(bstruct.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", bstructIdent, __gong__toRawStringLiteral(bstruct.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Floatfield = %f", bstructIdent, bstruct.Floatfield))
			values.WriteString(fmt.Sprintf("\n\t%s.Floatfield2 = %f", bstructIdent, bstruct.Floatfield2))
			values.WriteString(fmt.Sprintf("\n\t%s.Intfield = %d", bstructIdent, bstruct.Intfield))
			if str, ok := bstruct.Struct.(string); ok {
				values.WriteString(fmt.Sprintf("\n\t%s.Struct = %s", bstructIdent, str))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.ToBeIgnored = %d", bstructIdent, bstruct.ToBeIgnored))
		}
	}
	if stage != nil {
		for _, dstruct := range __gong__sortStageSetInstances(stage.Dstructs, stage.Dstruct_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			dstructIdent := dstruct.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Dstruct{Name: %s}).Stage(stage)", dstructIdent, __gong__toRawStringLiteral(dstruct.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", dstructIdent, __gong__toRawStringLiteral(dstruct.Name)))
			for _, elem := range dstruct.Anarrayofb {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Anarrayofb = append(%s.Anarrayofb, %s)", dstructIdent, dstructIdent, targetIdent))
			}
			if dstruct.Gstruct != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := dstruct.Gstruct.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Gstruct = %s", dstructIdent, targetIdent))
			}
			for _, elem := range dstruct.Gstructs {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Gstructs = append(%s.Gstructs, %s)", dstructIdent, dstructIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, f0123456789012345678901234567890 := range __gong__sortStageSetInstances(stage.F0123456789012345678901234567890s, stage.F0123456789012345678901234567890_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			f0123456789012345678901234567890Ident := f0123456789012345678901234567890.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.F0123456789012345678901234567890{Name: %s}).Stage(stage)", f0123456789012345678901234567890Ident, __gong__toRawStringLiteral(f0123456789012345678901234567890.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", f0123456789012345678901234567890Ident, __gong__toRawStringLiteral(f0123456789012345678901234567890.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Date, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", f0123456789012345678901234567890Ident, f0123456789012345678901234567890.Date.String()))
		}
	}
	if stage != nil {
		for _, gstruct := range __gong__sortStageSetInstances(stage.Gstructs, stage.Gstruct_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			gstructIdent := gstruct.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Gstruct{Name: %s}).Stage(stage)", gstructIdent, __gong__toRawStringLiteral(gstruct.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", gstructIdent, __gong__toRawStringLiteral(gstruct.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Floatfield = %f", gstructIdent, gstruct.Floatfield))
			values.WriteString(fmt.Sprintf("\n\t%s.Floatfield2 = %f", gstructIdent, gstruct.Floatfield2))
			values.WriteString(fmt.Sprintf("\n\t%s.Intfield = %d", gstructIdent, gstruct.Intfield))
			if str, ok := gstruct.Struct.(string); ok {
				values.WriteString(fmt.Sprintf("\n\t%s.Struct = %s", gstructIdent, str))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.ToBeIgnored = %d", gstructIdent, gstruct.ToBeIgnored))
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
go run github.com/fullstack-lang/gong/test/test1/go/models/probe/cmd/probe@%s stage.go
*/
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/test/test1/go/models"%s
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

	if modelsPackageName != "" && modelsPackageName != "github.com/fullstack-lang/gong/test/test1/go/models" {
		res = strings.ReplaceAll(res, "\"github.com/fullstack-lang/gong/test/test1/go/models\"", fmt.Sprintf("%q", modelsPackageName))
	}

	return res, nil
}

// insertion point for marshall field methods
func (astruct *Astruct) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := astruct.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", astruct.Name)
	case "Field":
		if str, ok := astruct.Field.(string); ok {
			res = __gong__marshallMeta(ident, "Field", str)
		}
	case "Date":
		res = __gong__marshallTime(ident, "Date", astruct.Date.String())
	case "Date2":
		res = __gong__marshallTime(ident, "Date2", astruct.Date2.String())
	case "Booleanfield":
		res = __gong__marshallBool(ident, "Booleanfield", astruct.Booleanfield)
	case "Aenum":
		res = __gong__marshallEnumString(ident, "Aenum", astruct.Aenum.ToCodeString())
	case "Aenum_2":
		res = __gong__marshallEnumString(ident, "Aenum_2", astruct.Aenum_2.ToCodeString())
	case "Benum":
		res = __gong__marshallEnumString(ident, "Benum", astruct.Benum.ToCodeString())
	case "CEnum":
		res = __gong__marshallEnumInt(ident, "CEnum", astruct.CEnum.ToCodeString())
	case "CName":
		res = __gong__marshallString(ident, "CName", astruct.CName)
	case "CFloatfield":
		res = __gong__marshallFloat(ident, "CFloatfield", astruct.CFloatfield)
	case "Floatfield":
		res = __gong__marshallFloat(ident, "Floatfield", astruct.Floatfield)
	case "Intfield":
		res = __gong__marshallInt(ident, "Intfield", astruct.Intfield)
	case "Anotherbooleanfield":
		res = __gong__marshallBool(ident, "Anotherbooleanfield", astruct.Anotherbooleanfield)
	case "Duration1":
		res = __gong__marshallInt(ident, "Duration1", astruct.Duration1)
	case "TextFieldBespokeSize":
		res = __gong__marshallString(ident, "TextFieldBespokeSize", astruct.TextFieldBespokeSize)
	case "TextArea":
		res = __gong__marshallString(ident, "TextArea", astruct.TextArea)

	case "Associationtob":
		if astruct.Associationtob != nil {
			res = __gong__marshallPointer(ident, "Associationtob", astruct.Associationtob.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Associationtob", "nil")
		}
	case "Anarrayofb":
		var sb strings.Builder
		for _, _bstruct := range astruct.Anarrayofb {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Anarrayofb", _bstruct.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Anotherassociationtob_2":
		if astruct.Anotherassociationtob_2 != nil {
			res = __gong__marshallPointer(ident, "Anotherassociationtob_2", astruct.Anotherassociationtob_2.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Anotherassociationtob_2", "nil")
		}
	case "Bstruct":
		if astruct.Bstruct != nil {
			res = __gong__marshallPointer(ident, "Bstruct", astruct.Bstruct.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Bstruct", "nil")
		}
	case "Bstruct2":
		if astruct.Bstruct2 != nil {
			res = __gong__marshallPointer(ident, "Bstruct2", astruct.Bstruct2.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Bstruct2", "nil")
		}
	case "Dstruct":
		if astruct.Dstruct != nil {
			res = __gong__marshallPointer(ident, "Dstruct", astruct.Dstruct.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Dstruct", "nil")
		}
	case "Dstruct2":
		if astruct.Dstruct2 != nil {
			res = __gong__marshallPointer(ident, "Dstruct2", astruct.Dstruct2.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Dstruct2", "nil")
		}
	case "Dstruct3":
		if astruct.Dstruct3 != nil {
			res = __gong__marshallPointer(ident, "Dstruct3", astruct.Dstruct3.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Dstruct3", "nil")
		}
	case "Dstruct4":
		if astruct.Dstruct4 != nil {
			res = __gong__marshallPointer(ident, "Dstruct4", astruct.Dstruct4.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Dstruct4", "nil")
		}
	case "Dstruct4s":
		var sb strings.Builder
		for _, _dstruct := range astruct.Dstruct4s {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Dstruct4s", _dstruct.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Anarrayofa":
		var sb strings.Builder
		for _, _astruct := range astruct.Anarrayofa {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Anarrayofa", _astruct.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Anotherarrayofb":
		var sb strings.Builder
		for _, _bstruct := range astruct.Anotherarrayofb {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Anotherarrayofb", _bstruct.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "AnarrayofbUse":
		var sb strings.Builder
		for _, _astructbstructuse := range astruct.AnarrayofbUse {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "AnarrayofbUse", _astructbstructuse.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Anarrayofb2Use":
		var sb strings.Builder
		for _, _astructbstruct2use := range astruct.Anarrayofb2Use {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Anarrayofb2Use", _astructbstruct2use.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "AnAstruct":
		if astruct.AnAstruct != nil {
			res = __gong__marshallPointer(ident, "AnAstruct", astruct.AnAstruct.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "AnAstruct", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Astruct", fieldName)
	}
	return
}

func (astructbstruct2use *AstructBstruct2Use) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := astructbstruct2use.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", astructbstruct2use.Name)

	case "Bstrcut2":
		if astructbstruct2use.Bstrcut2 != nil {
			res = __gong__marshallPointer(ident, "Bstrcut2", astructbstruct2use.Bstrcut2.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Bstrcut2", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct AstructBstruct2Use", fieldName)
	}
	return
}

func (astructbstructuse *AstructBstructUse) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := astructbstructuse.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", astructbstructuse.Name)

	case "Bstruct2":
		if astructbstructuse.Bstruct2 != nil {
			res = __gong__marshallPointer(ident, "Bstruct2", astructbstructuse.Bstruct2.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Bstruct2", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct AstructBstructUse", fieldName)
	}
	return
}

func (bstruct *Bstruct) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := bstruct.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", bstruct.Name)
	case "Floatfield":
		res = __gong__marshallFloat(ident, "Floatfield", bstruct.Floatfield)
	case "Floatfield2":
		res = __gong__marshallFloat(ident, "Floatfield2", bstruct.Floatfield2)
	case "Intfield":
		res = __gong__marshallInt(ident, "Intfield", bstruct.Intfield)

	default:
		log.Panicf("Unknown field %s for Gongstruct Bstruct", fieldName)
	}
	return
}

func (dstruct *Dstruct) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := dstruct.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", dstruct.Name)

	case "Anarrayofb":
		var sb strings.Builder
		for _, _bstruct := range dstruct.Anarrayofb {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Anarrayofb", _bstruct.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Gstruct":
		if dstruct.Gstruct != nil {
			res = __gong__marshallPointer(ident, "Gstruct", dstruct.Gstruct.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Gstruct", "nil")
		}
	case "Gstructs":
		var sb strings.Builder
		for _, _gstruct := range dstruct.Gstructs {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Gstructs", _gstruct.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Dstruct", fieldName)
	}
	return
}

func (f0123456789012345678901234567890 *F0123456789012345678901234567890) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := f0123456789012345678901234567890.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", f0123456789012345678901234567890.Name)
	case "Date":
		res = __gong__marshallTime(ident, "Date", f0123456789012345678901234567890.Date.String())

	default:
		log.Panicf("Unknown field %s for Gongstruct F0123456789012345678901234567890", fieldName)
	}
	return
}

func (gstruct *Gstruct) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := gstruct.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", gstruct.Name)
	case "Floatfield":
		res = __gong__marshallFloat(ident, "Floatfield", gstruct.Floatfield)
	case "Floatfield2":
		res = __gong__marshallFloat(ident, "Floatfield2", gstruct.Floatfield2)
	case "Intfield":
		res = __gong__marshallInt(ident, "Intfield", gstruct.Intfield)

	default:
		log.Panicf("Unknown field %s for Gongstruct Gstruct", fieldName)
	}
	return
}

// insertion point for marshall all fields methods
func (astruct *Astruct) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(astruct.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(astruct.GongMarshallField(stage, "Field"))
		pointersInitializesStatements.WriteString(astruct.GongMarshallField(stage, "Associationtob"))
		pointersInitializesStatements.WriteString(astruct.GongMarshallField(stage, "Anarrayofb"))
		pointersInitializesStatements.WriteString(astruct.GongMarshallField(stage, "Anotherassociationtob_2"))
		initializerStatements.WriteString(astruct.GongMarshallField(stage, "Date"))
		initializerStatements.WriteString(astruct.GongMarshallField(stage, "Date2"))
		initializerStatements.WriteString(astruct.GongMarshallField(stage, "Booleanfield"))
		initializerStatements.WriteString(astruct.GongMarshallField(stage, "Aenum"))
		initializerStatements.WriteString(astruct.GongMarshallField(stage, "Aenum_2"))
		initializerStatements.WriteString(astruct.GongMarshallField(stage, "Benum"))
		initializerStatements.WriteString(astruct.GongMarshallField(stage, "CEnum"))
		initializerStatements.WriteString(astruct.GongMarshallField(stage, "CName"))
		initializerStatements.WriteString(astruct.GongMarshallField(stage, "CFloatfield"))
		pointersInitializesStatements.WriteString(astruct.GongMarshallField(stage, "Bstruct"))
		pointersInitializesStatements.WriteString(astruct.GongMarshallField(stage, "Bstruct2"))
		pointersInitializesStatements.WriteString(astruct.GongMarshallField(stage, "Dstruct"))
		pointersInitializesStatements.WriteString(astruct.GongMarshallField(stage, "Dstruct2"))
		pointersInitializesStatements.WriteString(astruct.GongMarshallField(stage, "Dstruct3"))
		pointersInitializesStatements.WriteString(astruct.GongMarshallField(stage, "Dstruct4"))
		pointersInitializesStatements.WriteString(astruct.GongMarshallField(stage, "Dstruct4s"))
		initializerStatements.WriteString(astruct.GongMarshallField(stage, "Floatfield"))
		initializerStatements.WriteString(astruct.GongMarshallField(stage, "Intfield"))
		initializerStatements.WriteString(astruct.GongMarshallField(stage, "Anotherbooleanfield"))
		initializerStatements.WriteString(astruct.GongMarshallField(stage, "Duration1"))
		pointersInitializesStatements.WriteString(astruct.GongMarshallField(stage, "Anarrayofa"))
		pointersInitializesStatements.WriteString(astruct.GongMarshallField(stage, "Anotherarrayofb"))
		pointersInitializesStatements.WriteString(astruct.GongMarshallField(stage, "AnarrayofbUse"))
		pointersInitializesStatements.WriteString(astruct.GongMarshallField(stage, "Anarrayofb2Use"))
		pointersInitializesStatements.WriteString(astruct.GongMarshallField(stage, "AnAstruct"))
		initializerStatements.WriteString(astruct.GongMarshallField(stage, "TextFieldBespokeSize"))
		initializerStatements.WriteString(astruct.GongMarshallField(stage, "TextArea"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (astructbstruct2use *AstructBstruct2Use) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(astructbstruct2use.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(astructbstruct2use.GongMarshallField(stage, "Bstrcut2"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (astructbstructuse *AstructBstructUse) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(astructbstructuse.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(astructbstructuse.GongMarshallField(stage, "Bstruct2"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (bstruct *Bstruct) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(bstruct.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(bstruct.GongMarshallField(stage, "Floatfield"))
		initializerStatements.WriteString(bstruct.GongMarshallField(stage, "Floatfield2"))
		initializerStatements.WriteString(bstruct.GongMarshallField(stage, "Intfield"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (dstruct *Dstruct) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(dstruct.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(dstruct.GongMarshallField(stage, "Anarrayofb"))
		pointersInitializesStatements.WriteString(dstruct.GongMarshallField(stage, "Gstruct"))
		pointersInitializesStatements.WriteString(dstruct.GongMarshallField(stage, "Gstructs"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (f0123456789012345678901234567890 *F0123456789012345678901234567890) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(f0123456789012345678901234567890.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(f0123456789012345678901234567890.GongMarshallField(stage, "Date"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (gstruct *Gstruct) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(gstruct.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(gstruct.GongMarshallField(stage, "Floatfield"))
		initializerStatements.WriteString(gstruct.GongMarshallField(stage, "Floatfield2"))
		initializerStatements.WriteString(gstruct.GongMarshallField(stage, "Intfield"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
