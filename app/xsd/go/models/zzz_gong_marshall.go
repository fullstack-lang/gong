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
	modelsPackageName := "github.com/fullstack-lang/gong/app/xsd/go/models"
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
		for _, all := range __gong__sortStageSetInstances(stage.Alls, stage.All_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			allIdent := localAliasPrefix + all.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.All{Name: %s}).Stage(stage)", allIdent, __gong__toRawStringLiteral(all.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", allIdent, __gong__toRawStringLiteral(all.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.OuterElementName = %s", allIdent, __gong__toRawStringLiteral(all.OuterElementName)))
			values.WriteString(fmt.Sprintf("\n\t%s.Order = %d", allIdent, all.Order))
			values.WriteString(fmt.Sprintf("\n\t%s.Depth = %d", allIdent, all.Depth))
			values.WriteString(fmt.Sprintf("\n\t%s.MinOccurs = %s", allIdent, __gong__toRawStringLiteral(all.MinOccurs)))
			values.WriteString(fmt.Sprintf("\n\t%s.MaxOccurs = %s", allIdent, __gong__toRawStringLiteral(all.MaxOccurs)))
			if all.Annotation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + all.Annotation.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = %s", allIdent, targetIdent))
			}
			for _, elem := range all.Sequences {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Sequences = append(%s.Sequences, %s)", allIdent, allIdent, targetIdent))
			}
			for _, elem := range all.Alls {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Alls = append(%s.Alls, %s)", allIdent, allIdent, targetIdent))
			}
			for _, elem := range all.Choices {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Choices = append(%s.Choices, %s)", allIdent, allIdent, targetIdent))
			}
			for _, elem := range all.Groups {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Groups = append(%s.Groups, %s)", allIdent, allIdent, targetIdent))
			}
			for _, elem := range all.Elements {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Elements = append(%s.Elements, %s)", allIdent, allIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, annotation := range __gong__sortStageSetInstances(stage.Annotations, stage.Annotation_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			annotationIdent := localAliasPrefix + annotation.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Annotation{Name: %s}).Stage(stage)", annotationIdent, __gong__toRawStringLiteral(annotation.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", annotationIdent, __gong__toRawStringLiteral(annotation.Name)))
			for _, elem := range annotation.Documentations {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Documentations = append(%s.Documentations, %s)", annotationIdent, annotationIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, attribute := range __gong__sortStageSetInstances(stage.Attributes, stage.Attribute_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			attributeIdent := localAliasPrefix + attribute.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Attribute{Name: %s}).Stage(stage)", attributeIdent, __gong__toRawStringLiteral(attribute.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", attributeIdent, __gong__toRawStringLiteral(attribute.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NameXSD = %s", attributeIdent, __gong__toRawStringLiteral(attribute.NameXSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", attributeIdent, __gong__toRawStringLiteral(attribute.Type)))
			values.WriteString(fmt.Sprintf("\n\t%s.HasNameConflict = %t", attributeIdent, attribute.HasNameConflict))
			values.WriteString(fmt.Sprintf("\n\t%s.GoIdentifier = %s", attributeIdent, __gong__toRawStringLiteral(attribute.GoIdentifier)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default = %s", attributeIdent, __gong__toRawStringLiteral(attribute.Default)))
			values.WriteString(fmt.Sprintf("\n\t%s.Use = %s", attributeIdent, __gong__toRawStringLiteral(attribute.Use)))
			values.WriteString(fmt.Sprintf("\n\t%s.Form = %s", attributeIdent, __gong__toRawStringLiteral(attribute.Form)))
			values.WriteString(fmt.Sprintf("\n\t%s.Fixed = %s", attributeIdent, __gong__toRawStringLiteral(attribute.Fixed)))
			values.WriteString(fmt.Sprintf("\n\t%s.Ref = %s", attributeIdent, __gong__toRawStringLiteral(attribute.Ref)))
			values.WriteString(fmt.Sprintf("\n\t%s.TargetNamespace = %s", attributeIdent, __gong__toRawStringLiteral(attribute.TargetNamespace)))
			values.WriteString(fmt.Sprintf("\n\t%s.SimpleType = %s", attributeIdent, __gong__toRawStringLiteral(attribute.SimpleType)))
			values.WriteString(fmt.Sprintf("\n\t%s.IDXSD = %s", attributeIdent, __gong__toRawStringLiteral(attribute.IDXSD)))
			if attribute.Annotation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + attribute.Annotation.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = %s", attributeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, attributegroup := range __gong__sortStageSetInstances(stage.AttributeGroups, stage.AttributeGroup_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			attributegroupIdent := localAliasPrefix + attributegroup.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.AttributeGroup{Name: %s}).Stage(stage)", attributegroupIdent, __gong__toRawStringLiteral(attributegroup.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", attributegroupIdent, __gong__toRawStringLiteral(attributegroup.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NameXSD = %s", attributegroupIdent, __gong__toRawStringLiteral(attributegroup.NameXSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.HasNameConflict = %t", attributegroupIdent, attributegroup.HasNameConflict))
			values.WriteString(fmt.Sprintf("\n\t%s.GoIdentifier = %s", attributegroupIdent, __gong__toRawStringLiteral(attributegroup.GoIdentifier)))
			values.WriteString(fmt.Sprintf("\n\t%s.Ref = %s", attributegroupIdent, __gong__toRawStringLiteral(attributegroup.Ref)))
			values.WriteString(fmt.Sprintf("\n\t%s.Order = %d", attributegroupIdent, attributegroup.Order))
			values.WriteString(fmt.Sprintf("\n\t%s.Depth = %d", attributegroupIdent, attributegroup.Depth))
			if attributegroup.Annotation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + attributegroup.Annotation.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = %s", attributegroupIdent, targetIdent))
			}
			for _, elem := range attributegroup.AttributeGroups {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AttributeGroups = append(%s.AttributeGroups, %s)", attributegroupIdent, attributegroupIdent, targetIdent))
			}
			for _, elem := range attributegroup.Attributes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Attributes = append(%s.Attributes, %s)", attributegroupIdent, attributegroupIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, choice := range __gong__sortStageSetInstances(stage.Choices, stage.Choice_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			choiceIdent := localAliasPrefix + choice.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Choice{Name: %s}).Stage(stage)", choiceIdent, __gong__toRawStringLiteral(choice.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", choiceIdent, __gong__toRawStringLiteral(choice.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.OuterElementName = %s", choiceIdent, __gong__toRawStringLiteral(choice.OuterElementName)))
			values.WriteString(fmt.Sprintf("\n\t%s.Order = %d", choiceIdent, choice.Order))
			values.WriteString(fmt.Sprintf("\n\t%s.Depth = %d", choiceIdent, choice.Depth))
			values.WriteString(fmt.Sprintf("\n\t%s.MinOccurs = %s", choiceIdent, __gong__toRawStringLiteral(choice.MinOccurs)))
			values.WriteString(fmt.Sprintf("\n\t%s.MaxOccurs = %s", choiceIdent, __gong__toRawStringLiteral(choice.MaxOccurs)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsDuplicatedInXSD = %t", choiceIdent, choice.IsDuplicatedInXSD))
			if choice.Annotation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + choice.Annotation.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = %s", choiceIdent, targetIdent))
			}
			for _, elem := range choice.Sequences {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Sequences = append(%s.Sequences, %s)", choiceIdent, choiceIdent, targetIdent))
			}
			for _, elem := range choice.Alls {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Alls = append(%s.Alls, %s)", choiceIdent, choiceIdent, targetIdent))
			}
			for _, elem := range choice.Choices {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Choices = append(%s.Choices, %s)", choiceIdent, choiceIdent, targetIdent))
			}
			for _, elem := range choice.Groups {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Groups = append(%s.Groups, %s)", choiceIdent, choiceIdent, targetIdent))
			}
			for _, elem := range choice.Elements {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Elements = append(%s.Elements, %s)", choiceIdent, choiceIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, complexcontent := range __gong__sortStageSetInstances(stage.ComplexContents, stage.ComplexContent_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			complexcontentIdent := localAliasPrefix + complexcontent.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ComplexContent{Name: %s}).Stage(stage)", complexcontentIdent, __gong__toRawStringLiteral(complexcontent.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", complexcontentIdent, __gong__toRawStringLiteral(complexcontent.Name)))
		}
	}
	if stage != nil {
		for _, complextype := range __gong__sortStageSetInstances(stage.ComplexTypes, stage.ComplexType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			complextypeIdent := localAliasPrefix + complextype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ComplexType{Name: %s}).Stage(stage)", complextypeIdent, __gong__toRawStringLiteral(complextype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", complextypeIdent, __gong__toRawStringLiteral(complextype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.HasNameConflict = %t", complextypeIdent, complextype.HasNameConflict))
			values.WriteString(fmt.Sprintf("\n\t%s.GoIdentifier = %s", complextypeIdent, __gong__toRawStringLiteral(complextype.GoIdentifier)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsAnonymous = %t", complextypeIdent, complextype.IsAnonymous))
			values.WriteString(fmt.Sprintf("\n\t%s.NameXSD = %s", complextypeIdent, __gong__toRawStringLiteral(complextype.NameXSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.OuterElementName = %s", complextypeIdent, __gong__toRawStringLiteral(complextype.OuterElementName)))
			values.WriteString(fmt.Sprintf("\n\t%s.Order = %d", complextypeIdent, complextype.Order))
			values.WriteString(fmt.Sprintf("\n\t%s.Depth = %d", complextypeIdent, complextype.Depth))
			values.WriteString(fmt.Sprintf("\n\t%s.MinOccurs = %s", complextypeIdent, __gong__toRawStringLiteral(complextype.MinOccurs)))
			values.WriteString(fmt.Sprintf("\n\t%s.MaxOccurs = %s", complextypeIdent, __gong__toRawStringLiteral(complextype.MaxOccurs)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsDuplicatedInXSD = %t", complextypeIdent, complextype.IsDuplicatedInXSD))
			if complextype.OuterElement != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + complextype.OuterElement.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OuterElement = %s", complextypeIdent, targetIdent))
			}
			if complextype.Annotation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + complextype.Annotation.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = %s", complextypeIdent, targetIdent))
			}
			for _, elem := range complextype.Sequences {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Sequences = append(%s.Sequences, %s)", complextypeIdent, complextypeIdent, targetIdent))
			}
			for _, elem := range complextype.Alls {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Alls = append(%s.Alls, %s)", complextypeIdent, complextypeIdent, targetIdent))
			}
			for _, elem := range complextype.Choices {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Choices = append(%s.Choices, %s)", complextypeIdent, complextypeIdent, targetIdent))
			}
			for _, elem := range complextype.Groups {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Groups = append(%s.Groups, %s)", complextypeIdent, complextypeIdent, targetIdent))
			}
			for _, elem := range complextype.Elements {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Elements = append(%s.Elements, %s)", complextypeIdent, complextypeIdent, targetIdent))
			}
			if complextype.Extension != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + complextype.Extension.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = %s", complextypeIdent, targetIdent))
			}
			if complextype.SimpleContent != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + complextype.SimpleContent.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SimpleContent = %s", complextypeIdent, targetIdent))
			}
			if complextype.ComplexContent != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + complextype.ComplexContent.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ComplexContent = %s", complextypeIdent, targetIdent))
			}
			for _, elem := range complextype.Attributes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Attributes = append(%s.Attributes, %s)", complextypeIdent, complextypeIdent, targetIdent))
			}
			for _, elem := range complextype.AttributeGroups {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AttributeGroups = append(%s.AttributeGroups, %s)", complextypeIdent, complextypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, documentation := range __gong__sortStageSetInstances(stage.Documentations, stage.Documentation_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			documentationIdent := localAliasPrefix + documentation.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Documentation{Name: %s}).Stage(stage)", documentationIdent, __gong__toRawStringLiteral(documentation.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", documentationIdent, __gong__toRawStringLiteral(documentation.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Text = %s", documentationIdent, __gong__toRawStringLiteral(documentation.Text)))
			values.WriteString(fmt.Sprintf("\n\t%s.Source = %s", documentationIdent, __gong__toRawStringLiteral(documentation.Source)))
			values.WriteString(fmt.Sprintf("\n\t%s.Lang = %s", documentationIdent, __gong__toRawStringLiteral(documentation.Lang)))
		}
	}
	if stage != nil {
		for _, element := range __gong__sortStageSetInstances(stage.Elements, stage.Element_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			elementIdent := localAliasPrefix + element.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Element{Name: %s}).Stage(stage)", elementIdent, __gong__toRawStringLiteral(element.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", elementIdent, __gong__toRawStringLiteral(element.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Order = %d", elementIdent, element.Order))
			values.WriteString(fmt.Sprintf("\n\t%s.Depth = %d", elementIdent, element.Depth))
			values.WriteString(fmt.Sprintf("\n\t%s.HasNameConflict = %t", elementIdent, element.HasNameConflict))
			values.WriteString(fmt.Sprintf("\n\t%s.GoIdentifier = %s", elementIdent, __gong__toRawStringLiteral(element.GoIdentifier)))
			values.WriteString(fmt.Sprintf("\n\t%s.NameXSD = %s", elementIdent, __gong__toRawStringLiteral(element.NameXSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", elementIdent, __gong__toRawStringLiteral(element.Type)))
			values.WriteString(fmt.Sprintf("\n\t%s.MinOccurs = %s", elementIdent, __gong__toRawStringLiteral(element.MinOccurs)))
			values.WriteString(fmt.Sprintf("\n\t%s.MaxOccurs = %s", elementIdent, __gong__toRawStringLiteral(element.MaxOccurs)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default = %s", elementIdent, __gong__toRawStringLiteral(element.Default)))
			values.WriteString(fmt.Sprintf("\n\t%s.Fixed = %s", elementIdent, __gong__toRawStringLiteral(element.Fixed)))
			values.WriteString(fmt.Sprintf("\n\t%s.Nillable = %s", elementIdent, __gong__toRawStringLiteral(element.Nillable)))
			values.WriteString(fmt.Sprintf("\n\t%s.Ref = %s", elementIdent, __gong__toRawStringLiteral(element.Ref)))
			values.WriteString(fmt.Sprintf("\n\t%s.Abstract = %s", elementIdent, __gong__toRawStringLiteral(element.Abstract)))
			values.WriteString(fmt.Sprintf("\n\t%s.Form = %s", elementIdent, __gong__toRawStringLiteral(element.Form)))
			values.WriteString(fmt.Sprintf("\n\t%s.Block = %s", elementIdent, __gong__toRawStringLiteral(element.Block)))
			values.WriteString(fmt.Sprintf("\n\t%s.Final = %s", elementIdent, __gong__toRawStringLiteral(element.Final)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsDuplicatedInXSD = %t", elementIdent, element.IsDuplicatedInXSD))
			if element.Annotation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + element.Annotation.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = %s", elementIdent, targetIdent))
			}
			if element.SimpleType != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + element.SimpleType.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SimpleType = %s", elementIdent, targetIdent))
			}
			if element.ComplexType != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + element.ComplexType.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ComplexType = %s", elementIdent, targetIdent))
			}
			for _, elem := range element.Groups {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Groups = append(%s.Groups, %s)", elementIdent, elementIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, enumeration := range __gong__sortStageSetInstances(stage.Enumerations, stage.Enumeration_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			enumerationIdent := localAliasPrefix + enumeration.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Enumeration{Name: %s}).Stage(stage)", enumerationIdent, __gong__toRawStringLiteral(enumeration.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", enumerationIdent, __gong__toRawStringLiteral(enumeration.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", enumerationIdent, __gong__toRawStringLiteral(enumeration.Value)))
			if enumeration.Annotation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + enumeration.Annotation.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = %s", enumerationIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, extension := range __gong__sortStageSetInstances(stage.Extensions, stage.Extension_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			extensionIdent := localAliasPrefix + extension.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Extension{Name: %s}).Stage(stage)", extensionIdent, __gong__toRawStringLiteral(extension.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", extensionIdent, __gong__toRawStringLiteral(extension.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.OuterElementName = %s", extensionIdent, __gong__toRawStringLiteral(extension.OuterElementName)))
			values.WriteString(fmt.Sprintf("\n\t%s.Order = %d", extensionIdent, extension.Order))
			values.WriteString(fmt.Sprintf("\n\t%s.Depth = %d", extensionIdent, extension.Depth))
			values.WriteString(fmt.Sprintf("\n\t%s.MinOccurs = %s", extensionIdent, __gong__toRawStringLiteral(extension.MinOccurs)))
			values.WriteString(fmt.Sprintf("\n\t%s.MaxOccurs = %s", extensionIdent, __gong__toRawStringLiteral(extension.MaxOccurs)))
			values.WriteString(fmt.Sprintf("\n\t%s.Base = %s", extensionIdent, __gong__toRawStringLiteral(extension.Base)))
			values.WriteString(fmt.Sprintf("\n\t%s.Ref = %s", extensionIdent, __gong__toRawStringLiteral(extension.Ref)))
			for _, elem := range extension.Sequences {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Sequences = append(%s.Sequences, %s)", extensionIdent, extensionIdent, targetIdent))
			}
			for _, elem := range extension.Alls {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Alls = append(%s.Alls, %s)", extensionIdent, extensionIdent, targetIdent))
			}
			for _, elem := range extension.Choices {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Choices = append(%s.Choices, %s)", extensionIdent, extensionIdent, targetIdent))
			}
			for _, elem := range extension.Groups {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Groups = append(%s.Groups, %s)", extensionIdent, extensionIdent, targetIdent))
			}
			for _, elem := range extension.Elements {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Elements = append(%s.Elements, %s)", extensionIdent, extensionIdent, targetIdent))
			}
			for _, elem := range extension.Attributes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Attributes = append(%s.Attributes, %s)", extensionIdent, extensionIdent, targetIdent))
			}
			for _, elem := range extension.AttributeGroups {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AttributeGroups = append(%s.AttributeGroups, %s)", extensionIdent, extensionIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, group := range __gong__sortStageSetInstances(stage.Groups, stage.Group_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			groupIdent := localAliasPrefix + group.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Group{Name: %s}).Stage(stage)", groupIdent, __gong__toRawStringLiteral(group.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", groupIdent, __gong__toRawStringLiteral(group.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NameXSD = %s", groupIdent, __gong__toRawStringLiteral(group.NameXSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Ref = %s", groupIdent, __gong__toRawStringLiteral(group.Ref)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsAnonymous = %t", groupIdent, group.IsAnonymous))
			values.WriteString(fmt.Sprintf("\n\t%s.HasNameConflict = %t", groupIdent, group.HasNameConflict))
			values.WriteString(fmt.Sprintf("\n\t%s.GoIdentifier = %s", groupIdent, __gong__toRawStringLiteral(group.GoIdentifier)))
			values.WriteString(fmt.Sprintf("\n\t%s.OuterElementName = %s", groupIdent, __gong__toRawStringLiteral(group.OuterElementName)))
			values.WriteString(fmt.Sprintf("\n\t%s.Order = %d", groupIdent, group.Order))
			values.WriteString(fmt.Sprintf("\n\t%s.Depth = %d", groupIdent, group.Depth))
			values.WriteString(fmt.Sprintf("\n\t%s.MinOccurs = %s", groupIdent, __gong__toRawStringLiteral(group.MinOccurs)))
			values.WriteString(fmt.Sprintf("\n\t%s.MaxOccurs = %s", groupIdent, __gong__toRawStringLiteral(group.MaxOccurs)))
			if group.Annotation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + group.Annotation.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = %s", groupIdent, targetIdent))
			}
			if group.OuterElement != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + group.OuterElement.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OuterElement = %s", groupIdent, targetIdent))
			}
			for _, elem := range group.Sequences {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Sequences = append(%s.Sequences, %s)", groupIdent, groupIdent, targetIdent))
			}
			for _, elem := range group.Alls {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Alls = append(%s.Alls, %s)", groupIdent, groupIdent, targetIdent))
			}
			for _, elem := range group.Choices {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Choices = append(%s.Choices, %s)", groupIdent, groupIdent, targetIdent))
			}
			for _, elem := range group.Groups {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Groups = append(%s.Groups, %s)", groupIdent, groupIdent, targetIdent))
			}
			for _, elem := range group.Elements {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Elements = append(%s.Elements, %s)", groupIdent, groupIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, length := range __gong__sortStageSetInstances(stage.Lengths, stage.Length_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			lengthIdent := localAliasPrefix + length.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Length{Name: %s}).Stage(stage)", lengthIdent, __gong__toRawStringLiteral(length.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", lengthIdent, __gong__toRawStringLiteral(length.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", lengthIdent, __gong__toRawStringLiteral(length.Value)))
			if length.Annotation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + length.Annotation.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = %s", lengthIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, maxinclusive := range __gong__sortStageSetInstances(stage.MaxInclusives, stage.MaxInclusive_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			maxinclusiveIdent := localAliasPrefix + maxinclusive.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.MaxInclusive{Name: %s}).Stage(stage)", maxinclusiveIdent, __gong__toRawStringLiteral(maxinclusive.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", maxinclusiveIdent, __gong__toRawStringLiteral(maxinclusive.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", maxinclusiveIdent, __gong__toRawStringLiteral(maxinclusive.Value)))
			if maxinclusive.Annotation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + maxinclusive.Annotation.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = %s", maxinclusiveIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, maxlength := range __gong__sortStageSetInstances(stage.MaxLengths, stage.MaxLength_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			maxlengthIdent := localAliasPrefix + maxlength.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.MaxLength{Name: %s}).Stage(stage)", maxlengthIdent, __gong__toRawStringLiteral(maxlength.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", maxlengthIdent, __gong__toRawStringLiteral(maxlength.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", maxlengthIdent, __gong__toRawStringLiteral(maxlength.Value)))
			if maxlength.Annotation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + maxlength.Annotation.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = %s", maxlengthIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, mininclusive := range __gong__sortStageSetInstances(stage.MinInclusives, stage.MinInclusive_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			mininclusiveIdent := localAliasPrefix + mininclusive.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.MinInclusive{Name: %s}).Stage(stage)", mininclusiveIdent, __gong__toRawStringLiteral(mininclusive.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", mininclusiveIdent, __gong__toRawStringLiteral(mininclusive.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", mininclusiveIdent, __gong__toRawStringLiteral(mininclusive.Value)))
			if mininclusive.Annotation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + mininclusive.Annotation.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = %s", mininclusiveIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, minlength := range __gong__sortStageSetInstances(stage.MinLengths, stage.MinLength_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			minlengthIdent := localAliasPrefix + minlength.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.MinLength{Name: %s}).Stage(stage)", minlengthIdent, __gong__toRawStringLiteral(minlength.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", minlengthIdent, __gong__toRawStringLiteral(minlength.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", minlengthIdent, __gong__toRawStringLiteral(minlength.Value)))
			if minlength.Annotation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + minlength.Annotation.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = %s", minlengthIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, pattern := range __gong__sortStageSetInstances(stage.Patterns, stage.Pattern_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			patternIdent := localAliasPrefix + pattern.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Pattern{Name: %s}).Stage(stage)", patternIdent, __gong__toRawStringLiteral(pattern.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", patternIdent, __gong__toRawStringLiteral(pattern.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", patternIdent, __gong__toRawStringLiteral(pattern.Value)))
			if pattern.Annotation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + pattern.Annotation.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = %s", patternIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, restriction := range __gong__sortStageSetInstances(stage.Restrictions, stage.Restriction_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			restrictionIdent := localAliasPrefix + restriction.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Restriction{Name: %s}).Stage(stage)", restrictionIdent, __gong__toRawStringLiteral(restriction.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", restrictionIdent, __gong__toRawStringLiteral(restriction.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Base = %s", restrictionIdent, __gong__toRawStringLiteral(restriction.Base)))
			if restriction.Annotation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + restriction.Annotation.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = %s", restrictionIdent, targetIdent))
			}
			for _, elem := range restriction.Enumerations {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Enumerations = append(%s.Enumerations, %s)", restrictionIdent, restrictionIdent, targetIdent))
			}
			if restriction.MinInclusive != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + restriction.MinInclusive.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MinInclusive = %s", restrictionIdent, targetIdent))
			}
			if restriction.MaxInclusive != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + restriction.MaxInclusive.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MaxInclusive = %s", restrictionIdent, targetIdent))
			}
			if restriction.Pattern != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + restriction.Pattern.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Pattern = %s", restrictionIdent, targetIdent))
			}
			if restriction.WhiteSpace != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + restriction.WhiteSpace.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.WhiteSpace = %s", restrictionIdent, targetIdent))
			}
			if restriction.MinLength != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + restriction.MinLength.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MinLength = %s", restrictionIdent, targetIdent))
			}
			if restriction.MaxLength != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + restriction.MaxLength.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MaxLength = %s", restrictionIdent, targetIdent))
			}
			if restriction.Length != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + restriction.Length.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Length = %s", restrictionIdent, targetIdent))
			}
			if restriction.TotalDigit != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + restriction.TotalDigit.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TotalDigit = %s", restrictionIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, schema := range __gong__sortStageSetInstances(stage.Schemas, stage.Schema_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			schemaIdent := localAliasPrefix + schema.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Schema{Name: %s}).Stage(stage)", schemaIdent, __gong__toRawStringLiteral(schema.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", schemaIdent, __gong__toRawStringLiteral(schema.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Xs = %s", schemaIdent, __gong__toRawStringLiteral(schema.Xs)))
			values.WriteString(fmt.Sprintf("\n\t%s.Order = %d", schemaIdent, schema.Order))
			values.WriteString(fmt.Sprintf("\n\t%s.Depth = %d", schemaIdent, schema.Depth))
			if schema.Annotation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + schema.Annotation.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = %s", schemaIdent, targetIdent))
			}
			for _, elem := range schema.Elements {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Elements = append(%s.Elements, %s)", schemaIdent, schemaIdent, targetIdent))
			}
			for _, elem := range schema.SimpleTypes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SimpleTypes = append(%s.SimpleTypes, %s)", schemaIdent, schemaIdent, targetIdent))
			}
			for _, elem := range schema.ComplexTypes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ComplexTypes = append(%s.ComplexTypes, %s)", schemaIdent, schemaIdent, targetIdent))
			}
			for _, elem := range schema.AttributeGroups {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AttributeGroups = append(%s.AttributeGroups, %s)", schemaIdent, schemaIdent, targetIdent))
			}
			for _, elem := range schema.Groups {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Groups = append(%s.Groups, %s)", schemaIdent, schemaIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, sequence := range __gong__sortStageSetInstances(stage.Sequences, stage.Sequence_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			sequenceIdent := localAliasPrefix + sequence.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Sequence{Name: %s}).Stage(stage)", sequenceIdent, __gong__toRawStringLiteral(sequence.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", sequenceIdent, __gong__toRawStringLiteral(sequence.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.OuterElementName = %s", sequenceIdent, __gong__toRawStringLiteral(sequence.OuterElementName)))
			values.WriteString(fmt.Sprintf("\n\t%s.Order = %d", sequenceIdent, sequence.Order))
			values.WriteString(fmt.Sprintf("\n\t%s.Depth = %d", sequenceIdent, sequence.Depth))
			values.WriteString(fmt.Sprintf("\n\t%s.MinOccurs = %s", sequenceIdent, __gong__toRawStringLiteral(sequence.MinOccurs)))
			values.WriteString(fmt.Sprintf("\n\t%s.MaxOccurs = %s", sequenceIdent, __gong__toRawStringLiteral(sequence.MaxOccurs)))
			if sequence.Annotation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + sequence.Annotation.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = %s", sequenceIdent, targetIdent))
			}
			for _, elem := range sequence.Sequences {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Sequences = append(%s.Sequences, %s)", sequenceIdent, sequenceIdent, targetIdent))
			}
			for _, elem := range sequence.Alls {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Alls = append(%s.Alls, %s)", sequenceIdent, sequenceIdent, targetIdent))
			}
			for _, elem := range sequence.Choices {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Choices = append(%s.Choices, %s)", sequenceIdent, sequenceIdent, targetIdent))
			}
			for _, elem := range sequence.Groups {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Groups = append(%s.Groups, %s)", sequenceIdent, sequenceIdent, targetIdent))
			}
			for _, elem := range sequence.Elements {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Elements = append(%s.Elements, %s)", sequenceIdent, sequenceIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, simplecontent := range __gong__sortStageSetInstances(stage.SimpleContents, stage.SimpleContent_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			simplecontentIdent := localAliasPrefix + simplecontent.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.SimpleContent{Name: %s}).Stage(stage)", simplecontentIdent, __gong__toRawStringLiteral(simplecontent.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", simplecontentIdent, __gong__toRawStringLiteral(simplecontent.Name)))
			if simplecontent.Extension != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + simplecontent.Extension.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = %s", simplecontentIdent, targetIdent))
			}
			if simplecontent.Restriction != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + simplecontent.Restriction.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Restriction = %s", simplecontentIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, simpletype := range __gong__sortStageSetInstances(stage.SimpleTypes, stage.SimpleType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			simpletypeIdent := localAliasPrefix + simpletype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.SimpleType{Name: %s}).Stage(stage)", simpletypeIdent, __gong__toRawStringLiteral(simpletype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", simpletypeIdent, __gong__toRawStringLiteral(simpletype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NameXSD = %s", simpletypeIdent, __gong__toRawStringLiteral(simpletype.NameXSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Order = %d", simpletypeIdent, simpletype.Order))
			values.WriteString(fmt.Sprintf("\n\t%s.Depth = %d", simpletypeIdent, simpletype.Depth))
			if simpletype.Annotation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + simpletype.Annotation.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = %s", simpletypeIdent, targetIdent))
			}
			if simpletype.Restriction != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + simpletype.Restriction.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Restriction = %s", simpletypeIdent, targetIdent))
			}
			if simpletype.Union != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + simpletype.Union.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Union = %s", simpletypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, totaldigit := range __gong__sortStageSetInstances(stage.TotalDigits, stage.TotalDigit_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			totaldigitIdent := localAliasPrefix + totaldigit.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.TotalDigit{Name: %s}).Stage(stage)", totaldigitIdent, __gong__toRawStringLiteral(totaldigit.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", totaldigitIdent, __gong__toRawStringLiteral(totaldigit.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", totaldigitIdent, __gong__toRawStringLiteral(totaldigit.Value)))
			if totaldigit.Annotation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + totaldigit.Annotation.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = %s", totaldigitIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, union := range __gong__sortStageSetInstances(stage.Unions, stage.Union_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			unionIdent := localAliasPrefix + union.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Union{Name: %s}).Stage(stage)", unionIdent, __gong__toRawStringLiteral(union.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", unionIdent, __gong__toRawStringLiteral(union.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.MemberTypes = %s", unionIdent, __gong__toRawStringLiteral(union.MemberTypes)))
			if union.Annotation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + union.Annotation.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = %s", unionIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, whitespace := range __gong__sortStageSetInstances(stage.WhiteSpaces, stage.WhiteSpace_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			whitespaceIdent := localAliasPrefix + whitespace.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.WhiteSpace{Name: %s}).Stage(stage)", whitespaceIdent, __gong__toRawStringLiteral(whitespace.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", whitespaceIdent, __gong__toRawStringLiteral(whitespace.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", whitespaceIdent, __gong__toRawStringLiteral(whitespace.Value)))
			if whitespace.Annotation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + whitespace.Annotation.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = %s", whitespaceIdent, targetIdent))
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
go run github.com/fullstack-lang/gong/app/xsd/go/models/probe/cmd/probe@%s stage.go
*/
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/app/xsd/go/models"%s
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

	if modelsPackageName != "" && modelsPackageName != "github.com/fullstack-lang/gong/app/xsd/go/models" {
		res = strings.ReplaceAll(res, "\"github.com/fullstack-lang/gong/app/xsd/go/models\"", fmt.Sprintf("%q", modelsPackageName))
	}

	return res, nil
}

// insertion point for marshall field methods
func (all *All) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := all.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", all.Name)
	case "OuterElementName":
		res = __gong__marshallString(ident, "OuterElementName", all.OuterElementName)
	case "Order":
		res = __gong__marshallInt(ident, "Order", all.Order)
	case "Depth":
		res = __gong__marshallInt(ident, "Depth", all.Depth)
	case "MinOccurs":
		res = __gong__marshallString(ident, "MinOccurs", all.MinOccurs)
	case "MaxOccurs":
		res = __gong__marshallString(ident, "MaxOccurs", all.MaxOccurs)

	case "Annotation":
		if all.Annotation != nil {
			res = __gong__marshallPointer(ident, "Annotation", all.Annotation.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Annotation", "nil")
		}
	case "Sequences":
		var sb strings.Builder
		for _, _sequence := range all.Sequences {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Sequences", _sequence.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Alls":
		var sb strings.Builder
		for _, _all := range all.Alls {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Alls", _all.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Choices":
		var sb strings.Builder
		for _, _choice := range all.Choices {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Choices", _choice.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Groups":
		var sb strings.Builder
		for _, _group := range all.Groups {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Groups", _group.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Elements":
		var sb strings.Builder
		for _, _element := range all.Elements {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Elements", _element.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct All", fieldName)
	}
	return
}

func (annotation *Annotation) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := annotation.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", annotation.Name)

	case "Documentations":
		var sb strings.Builder
		for _, _documentation := range annotation.Documentations {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Documentations", _documentation.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Annotation", fieldName)
	}
	return
}

func (attribute *Attribute) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := attribute.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", attribute.Name)
	case "NameXSD":
		res = __gong__marshallString(ident, "NameXSD", attribute.NameXSD)
	case "Type":
		res = __gong__marshallString(ident, "Type", attribute.Type)
	case "HasNameConflict":
		res = __gong__marshallBool(ident, "HasNameConflict", attribute.HasNameConflict)
	case "GoIdentifier":
		res = __gong__marshallString(ident, "GoIdentifier", attribute.GoIdentifier)
	case "Default":
		res = __gong__marshallString(ident, "Default", attribute.Default)
	case "Use":
		res = __gong__marshallString(ident, "Use", attribute.Use)
	case "Form":
		res = __gong__marshallString(ident, "Form", attribute.Form)
	case "Fixed":
		res = __gong__marshallString(ident, "Fixed", attribute.Fixed)
	case "Ref":
		res = __gong__marshallString(ident, "Ref", attribute.Ref)
	case "TargetNamespace":
		res = __gong__marshallString(ident, "TargetNamespace", attribute.TargetNamespace)
	case "SimpleType":
		res = __gong__marshallString(ident, "SimpleType", attribute.SimpleType)
	case "IDXSD":
		res = __gong__marshallString(ident, "IDXSD", attribute.IDXSD)

	case "Annotation":
		if attribute.Annotation != nil {
			res = __gong__marshallPointer(ident, "Annotation", attribute.Annotation.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Annotation", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Attribute", fieldName)
	}
	return
}

func (attributegroup *AttributeGroup) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := attributegroup.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", attributegroup.Name)
	case "NameXSD":
		res = __gong__marshallString(ident, "NameXSD", attributegroup.NameXSD)
	case "HasNameConflict":
		res = __gong__marshallBool(ident, "HasNameConflict", attributegroup.HasNameConflict)
	case "GoIdentifier":
		res = __gong__marshallString(ident, "GoIdentifier", attributegroup.GoIdentifier)
	case "Ref":
		res = __gong__marshallString(ident, "Ref", attributegroup.Ref)
	case "Order":
		res = __gong__marshallInt(ident, "Order", attributegroup.Order)
	case "Depth":
		res = __gong__marshallInt(ident, "Depth", attributegroup.Depth)

	case "Annotation":
		if attributegroup.Annotation != nil {
			res = __gong__marshallPointer(ident, "Annotation", attributegroup.Annotation.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Annotation", "nil")
		}
	case "AttributeGroups":
		var sb strings.Builder
		for _, _attributegroup := range attributegroup.AttributeGroups {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "AttributeGroups", _attributegroup.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Attributes":
		var sb strings.Builder
		for _, _attribute := range attributegroup.Attributes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Attributes", _attribute.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct AttributeGroup", fieldName)
	}
	return
}

func (choice *Choice) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := choice.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", choice.Name)
	case "OuterElementName":
		res = __gong__marshallString(ident, "OuterElementName", choice.OuterElementName)
	case "Order":
		res = __gong__marshallInt(ident, "Order", choice.Order)
	case "Depth":
		res = __gong__marshallInt(ident, "Depth", choice.Depth)
	case "MinOccurs":
		res = __gong__marshallString(ident, "MinOccurs", choice.MinOccurs)
	case "MaxOccurs":
		res = __gong__marshallString(ident, "MaxOccurs", choice.MaxOccurs)
	case "IsDuplicatedInXSD":
		res = __gong__marshallBool(ident, "IsDuplicatedInXSD", choice.IsDuplicatedInXSD)

	case "Annotation":
		if choice.Annotation != nil {
			res = __gong__marshallPointer(ident, "Annotation", choice.Annotation.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Annotation", "nil")
		}
	case "Sequences":
		var sb strings.Builder
		for _, _sequence := range choice.Sequences {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Sequences", _sequence.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Alls":
		var sb strings.Builder
		for _, _all := range choice.Alls {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Alls", _all.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Choices":
		var sb strings.Builder
		for _, _choice := range choice.Choices {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Choices", _choice.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Groups":
		var sb strings.Builder
		for _, _group := range choice.Groups {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Groups", _group.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Elements":
		var sb strings.Builder
		for _, _element := range choice.Elements {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Elements", _element.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Choice", fieldName)
	}
	return
}

func (complexcontent *ComplexContent) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := complexcontent.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", complexcontent.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct ComplexContent", fieldName)
	}
	return
}

func (complextype *ComplexType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := complextype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", complextype.Name)
	case "HasNameConflict":
		res = __gong__marshallBool(ident, "HasNameConflict", complextype.HasNameConflict)
	case "GoIdentifier":
		res = __gong__marshallString(ident, "GoIdentifier", complextype.GoIdentifier)
	case "IsAnonymous":
		res = __gong__marshallBool(ident, "IsAnonymous", complextype.IsAnonymous)
	case "NameXSD":
		res = __gong__marshallString(ident, "NameXSD", complextype.NameXSD)
	case "OuterElementName":
		res = __gong__marshallString(ident, "OuterElementName", complextype.OuterElementName)
	case "Order":
		res = __gong__marshallInt(ident, "Order", complextype.Order)
	case "Depth":
		res = __gong__marshallInt(ident, "Depth", complextype.Depth)
	case "MinOccurs":
		res = __gong__marshallString(ident, "MinOccurs", complextype.MinOccurs)
	case "MaxOccurs":
		res = __gong__marshallString(ident, "MaxOccurs", complextype.MaxOccurs)
	case "IsDuplicatedInXSD":
		res = __gong__marshallBool(ident, "IsDuplicatedInXSD", complextype.IsDuplicatedInXSD)

	case "OuterElement":
		if complextype.OuterElement != nil {
			res = __gong__marshallPointer(ident, "OuterElement", complextype.OuterElement.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "OuterElement", "nil")
		}
	case "Annotation":
		if complextype.Annotation != nil {
			res = __gong__marshallPointer(ident, "Annotation", complextype.Annotation.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Annotation", "nil")
		}
	case "Sequences":
		var sb strings.Builder
		for _, _sequence := range complextype.Sequences {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Sequences", _sequence.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Alls":
		var sb strings.Builder
		for _, _all := range complextype.Alls {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Alls", _all.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Choices":
		var sb strings.Builder
		for _, _choice := range complextype.Choices {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Choices", _choice.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Groups":
		var sb strings.Builder
		for _, _group := range complextype.Groups {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Groups", _group.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Elements":
		var sb strings.Builder
		for _, _element := range complextype.Elements {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Elements", _element.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		if complextype.Extension != nil {
			res = __gong__marshallPointer(ident, "Extension", complextype.Extension.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Extension", "nil")
		}
	case "SimpleContent":
		if complextype.SimpleContent != nil {
			res = __gong__marshallPointer(ident, "SimpleContent", complextype.SimpleContent.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "SimpleContent", "nil")
		}
	case "ComplexContent":
		if complextype.ComplexContent != nil {
			res = __gong__marshallPointer(ident, "ComplexContent", complextype.ComplexContent.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ComplexContent", "nil")
		}
	case "Attributes":
		var sb strings.Builder
		for _, _attribute := range complextype.Attributes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Attributes", _attribute.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "AttributeGroups":
		var sb strings.Builder
		for _, _attributegroup := range complextype.AttributeGroups {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "AttributeGroups", _attributegroup.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct ComplexType", fieldName)
	}
	return
}

func (documentation *Documentation) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := documentation.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", documentation.Name)
	case "Text":
		res = __gong__marshallString(ident, "Text", documentation.Text)
	case "Source":
		res = __gong__marshallString(ident, "Source", documentation.Source)
	case "Lang":
		res = __gong__marshallString(ident, "Lang", documentation.Lang)

	default:
		log.Panicf("Unknown field %s for Gongstruct Documentation", fieldName)
	}
	return
}

func (element *Element) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := element.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", element.Name)
	case "Order":
		res = __gong__marshallInt(ident, "Order", element.Order)
	case "Depth":
		res = __gong__marshallInt(ident, "Depth", element.Depth)
	case "HasNameConflict":
		res = __gong__marshallBool(ident, "HasNameConflict", element.HasNameConflict)
	case "GoIdentifier":
		res = __gong__marshallString(ident, "GoIdentifier", element.GoIdentifier)
	case "NameXSD":
		res = __gong__marshallString(ident, "NameXSD", element.NameXSD)
	case "Type":
		res = __gong__marshallString(ident, "Type", element.Type)
	case "MinOccurs":
		res = __gong__marshallString(ident, "MinOccurs", element.MinOccurs)
	case "MaxOccurs":
		res = __gong__marshallString(ident, "MaxOccurs", element.MaxOccurs)
	case "Default":
		res = __gong__marshallString(ident, "Default", element.Default)
	case "Fixed":
		res = __gong__marshallString(ident, "Fixed", element.Fixed)
	case "Nillable":
		res = __gong__marshallString(ident, "Nillable", element.Nillable)
	case "Ref":
		res = __gong__marshallString(ident, "Ref", element.Ref)
	case "Abstract":
		res = __gong__marshallString(ident, "Abstract", element.Abstract)
	case "Form":
		res = __gong__marshallString(ident, "Form", element.Form)
	case "Block":
		res = __gong__marshallString(ident, "Block", element.Block)
	case "Final":
		res = __gong__marshallString(ident, "Final", element.Final)
	case "IsDuplicatedInXSD":
		res = __gong__marshallBool(ident, "IsDuplicatedInXSD", element.IsDuplicatedInXSD)

	case "Annotation":
		if element.Annotation != nil {
			res = __gong__marshallPointer(ident, "Annotation", element.Annotation.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Annotation", "nil")
		}
	case "SimpleType":
		if element.SimpleType != nil {
			res = __gong__marshallPointer(ident, "SimpleType", element.SimpleType.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "SimpleType", "nil")
		}
	case "ComplexType":
		if element.ComplexType != nil {
			res = __gong__marshallPointer(ident, "ComplexType", element.ComplexType.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ComplexType", "nil")
		}
	case "Groups":
		var sb strings.Builder
		for _, _group := range element.Groups {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Groups", _group.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Element", fieldName)
	}
	return
}

func (enumeration *Enumeration) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := enumeration.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", enumeration.Name)
	case "Value":
		res = __gong__marshallString(ident, "Value", enumeration.Value)

	case "Annotation":
		if enumeration.Annotation != nil {
			res = __gong__marshallPointer(ident, "Annotation", enumeration.Annotation.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Annotation", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Enumeration", fieldName)
	}
	return
}

func (extension *Extension) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := extension.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", extension.Name)
	case "OuterElementName":
		res = __gong__marshallString(ident, "OuterElementName", extension.OuterElementName)
	case "Order":
		res = __gong__marshallInt(ident, "Order", extension.Order)
	case "Depth":
		res = __gong__marshallInt(ident, "Depth", extension.Depth)
	case "MinOccurs":
		res = __gong__marshallString(ident, "MinOccurs", extension.MinOccurs)
	case "MaxOccurs":
		res = __gong__marshallString(ident, "MaxOccurs", extension.MaxOccurs)
	case "Base":
		res = __gong__marshallString(ident, "Base", extension.Base)
	case "Ref":
		res = __gong__marshallString(ident, "Ref", extension.Ref)

	case "Sequences":
		var sb strings.Builder
		for _, _sequence := range extension.Sequences {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Sequences", _sequence.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Alls":
		var sb strings.Builder
		for _, _all := range extension.Alls {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Alls", _all.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Choices":
		var sb strings.Builder
		for _, _choice := range extension.Choices {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Choices", _choice.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Groups":
		var sb strings.Builder
		for _, _group := range extension.Groups {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Groups", _group.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Elements":
		var sb strings.Builder
		for _, _element := range extension.Elements {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Elements", _element.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Attributes":
		var sb strings.Builder
		for _, _attribute := range extension.Attributes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Attributes", _attribute.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "AttributeGroups":
		var sb strings.Builder
		for _, _attributegroup := range extension.AttributeGroups {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "AttributeGroups", _attributegroup.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Extension", fieldName)
	}
	return
}

func (group *Group) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := group.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", group.Name)
	case "NameXSD":
		res = __gong__marshallString(ident, "NameXSD", group.NameXSD)
	case "Ref":
		res = __gong__marshallString(ident, "Ref", group.Ref)
	case "IsAnonymous":
		res = __gong__marshallBool(ident, "IsAnonymous", group.IsAnonymous)
	case "HasNameConflict":
		res = __gong__marshallBool(ident, "HasNameConflict", group.HasNameConflict)
	case "GoIdentifier":
		res = __gong__marshallString(ident, "GoIdentifier", group.GoIdentifier)
	case "OuterElementName":
		res = __gong__marshallString(ident, "OuterElementName", group.OuterElementName)
	case "Order":
		res = __gong__marshallInt(ident, "Order", group.Order)
	case "Depth":
		res = __gong__marshallInt(ident, "Depth", group.Depth)
	case "MinOccurs":
		res = __gong__marshallString(ident, "MinOccurs", group.MinOccurs)
	case "MaxOccurs":
		res = __gong__marshallString(ident, "MaxOccurs", group.MaxOccurs)

	case "Annotation":
		if group.Annotation != nil {
			res = __gong__marshallPointer(ident, "Annotation", group.Annotation.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Annotation", "nil")
		}
	case "OuterElement":
		if group.OuterElement != nil {
			res = __gong__marshallPointer(ident, "OuterElement", group.OuterElement.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "OuterElement", "nil")
		}
	case "Sequences":
		var sb strings.Builder
		for _, _sequence := range group.Sequences {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Sequences", _sequence.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Alls":
		var sb strings.Builder
		for _, _all := range group.Alls {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Alls", _all.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Choices":
		var sb strings.Builder
		for _, _choice := range group.Choices {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Choices", _choice.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Groups":
		var sb strings.Builder
		for _, _group := range group.Groups {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Groups", _group.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Elements":
		var sb strings.Builder
		for _, _element := range group.Elements {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Elements", _element.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Group", fieldName)
	}
	return
}

func (length *Length) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := length.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", length.Name)
	case "Value":
		res = __gong__marshallString(ident, "Value", length.Value)

	case "Annotation":
		if length.Annotation != nil {
			res = __gong__marshallPointer(ident, "Annotation", length.Annotation.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Annotation", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Length", fieldName)
	}
	return
}

func (maxinclusive *MaxInclusive) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := maxinclusive.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", maxinclusive.Name)
	case "Value":
		res = __gong__marshallString(ident, "Value", maxinclusive.Value)

	case "Annotation":
		if maxinclusive.Annotation != nil {
			res = __gong__marshallPointer(ident, "Annotation", maxinclusive.Annotation.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Annotation", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct MaxInclusive", fieldName)
	}
	return
}

func (maxlength *MaxLength) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := maxlength.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", maxlength.Name)
	case "Value":
		res = __gong__marshallString(ident, "Value", maxlength.Value)

	case "Annotation":
		if maxlength.Annotation != nil {
			res = __gong__marshallPointer(ident, "Annotation", maxlength.Annotation.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Annotation", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct MaxLength", fieldName)
	}
	return
}

func (mininclusive *MinInclusive) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := mininclusive.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", mininclusive.Name)
	case "Value":
		res = __gong__marshallString(ident, "Value", mininclusive.Value)

	case "Annotation":
		if mininclusive.Annotation != nil {
			res = __gong__marshallPointer(ident, "Annotation", mininclusive.Annotation.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Annotation", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct MinInclusive", fieldName)
	}
	return
}

func (minlength *MinLength) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := minlength.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", minlength.Name)
	case "Value":
		res = __gong__marshallString(ident, "Value", minlength.Value)

	case "Annotation":
		if minlength.Annotation != nil {
			res = __gong__marshallPointer(ident, "Annotation", minlength.Annotation.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Annotation", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct MinLength", fieldName)
	}
	return
}

func (pattern *Pattern) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := pattern.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", pattern.Name)
	case "Value":
		res = __gong__marshallString(ident, "Value", pattern.Value)

	case "Annotation":
		if pattern.Annotation != nil {
			res = __gong__marshallPointer(ident, "Annotation", pattern.Annotation.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Annotation", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Pattern", fieldName)
	}
	return
}

func (restriction *Restriction) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := restriction.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", restriction.Name)
	case "Base":
		res = __gong__marshallString(ident, "Base", restriction.Base)

	case "Annotation":
		if restriction.Annotation != nil {
			res = __gong__marshallPointer(ident, "Annotation", restriction.Annotation.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Annotation", "nil")
		}
	case "Enumerations":
		var sb strings.Builder
		for _, _enumeration := range restriction.Enumerations {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Enumerations", _enumeration.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "MinInclusive":
		if restriction.MinInclusive != nil {
			res = __gong__marshallPointer(ident, "MinInclusive", restriction.MinInclusive.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "MinInclusive", "nil")
		}
	case "MaxInclusive":
		if restriction.MaxInclusive != nil {
			res = __gong__marshallPointer(ident, "MaxInclusive", restriction.MaxInclusive.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "MaxInclusive", "nil")
		}
	case "Pattern":
		if restriction.Pattern != nil {
			res = __gong__marshallPointer(ident, "Pattern", restriction.Pattern.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Pattern", "nil")
		}
	case "WhiteSpace":
		if restriction.WhiteSpace != nil {
			res = __gong__marshallPointer(ident, "WhiteSpace", restriction.WhiteSpace.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "WhiteSpace", "nil")
		}
	case "MinLength":
		if restriction.MinLength != nil {
			res = __gong__marshallPointer(ident, "MinLength", restriction.MinLength.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "MinLength", "nil")
		}
	case "MaxLength":
		if restriction.MaxLength != nil {
			res = __gong__marshallPointer(ident, "MaxLength", restriction.MaxLength.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "MaxLength", "nil")
		}
	case "Length":
		if restriction.Length != nil {
			res = __gong__marshallPointer(ident, "Length", restriction.Length.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Length", "nil")
		}
	case "TotalDigit":
		if restriction.TotalDigit != nil {
			res = __gong__marshallPointer(ident, "TotalDigit", restriction.TotalDigit.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "TotalDigit", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Restriction", fieldName)
	}
	return
}

func (schema *Schema) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := schema.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", schema.Name)
	case "Xs":
		res = __gong__marshallString(ident, "Xs", schema.Xs)
	case "Order":
		res = __gong__marshallInt(ident, "Order", schema.Order)
	case "Depth":
		res = __gong__marshallInt(ident, "Depth", schema.Depth)

	case "Annotation":
		if schema.Annotation != nil {
			res = __gong__marshallPointer(ident, "Annotation", schema.Annotation.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Annotation", "nil")
		}
	case "Elements":
		var sb strings.Builder
		for _, _element := range schema.Elements {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Elements", _element.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "SimpleTypes":
		var sb strings.Builder
		for _, _simpletype := range schema.SimpleTypes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SimpleTypes", _simpletype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ComplexTypes":
		var sb strings.Builder
		for _, _complextype := range schema.ComplexTypes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ComplexTypes", _complextype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "AttributeGroups":
		var sb strings.Builder
		for _, _attributegroup := range schema.AttributeGroups {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "AttributeGroups", _attributegroup.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Groups":
		var sb strings.Builder
		for _, _group := range schema.Groups {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Groups", _group.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Schema", fieldName)
	}
	return
}

func (sequence *Sequence) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := sequence.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", sequence.Name)
	case "OuterElementName":
		res = __gong__marshallString(ident, "OuterElementName", sequence.OuterElementName)
	case "Order":
		res = __gong__marshallInt(ident, "Order", sequence.Order)
	case "Depth":
		res = __gong__marshallInt(ident, "Depth", sequence.Depth)
	case "MinOccurs":
		res = __gong__marshallString(ident, "MinOccurs", sequence.MinOccurs)
	case "MaxOccurs":
		res = __gong__marshallString(ident, "MaxOccurs", sequence.MaxOccurs)

	case "Annotation":
		if sequence.Annotation != nil {
			res = __gong__marshallPointer(ident, "Annotation", sequence.Annotation.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Annotation", "nil")
		}
	case "Sequences":
		var sb strings.Builder
		for _, _sequence := range sequence.Sequences {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Sequences", _sequence.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Alls":
		var sb strings.Builder
		for _, _all := range sequence.Alls {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Alls", _all.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Choices":
		var sb strings.Builder
		for _, _choice := range sequence.Choices {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Choices", _choice.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Groups":
		var sb strings.Builder
		for _, _group := range sequence.Groups {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Groups", _group.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Elements":
		var sb strings.Builder
		for _, _element := range sequence.Elements {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Elements", _element.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Sequence", fieldName)
	}
	return
}

func (simplecontent *SimpleContent) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := simplecontent.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", simplecontent.Name)

	case "Extension":
		if simplecontent.Extension != nil {
			res = __gong__marshallPointer(ident, "Extension", simplecontent.Extension.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Extension", "nil")
		}
	case "Restriction":
		if simplecontent.Restriction != nil {
			res = __gong__marshallPointer(ident, "Restriction", simplecontent.Restriction.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Restriction", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct SimpleContent", fieldName)
	}
	return
}

func (simpletype *SimpleType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := simpletype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", simpletype.Name)
	case "NameXSD":
		res = __gong__marshallString(ident, "NameXSD", simpletype.NameXSD)
	case "Order":
		res = __gong__marshallInt(ident, "Order", simpletype.Order)
	case "Depth":
		res = __gong__marshallInt(ident, "Depth", simpletype.Depth)

	case "Annotation":
		if simpletype.Annotation != nil {
			res = __gong__marshallPointer(ident, "Annotation", simpletype.Annotation.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Annotation", "nil")
		}
	case "Restriction":
		if simpletype.Restriction != nil {
			res = __gong__marshallPointer(ident, "Restriction", simpletype.Restriction.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Restriction", "nil")
		}
	case "Union":
		if simpletype.Union != nil {
			res = __gong__marshallPointer(ident, "Union", simpletype.Union.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Union", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct SimpleType", fieldName)
	}
	return
}

func (totaldigit *TotalDigit) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := totaldigit.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", totaldigit.Name)
	case "Value":
		res = __gong__marshallString(ident, "Value", totaldigit.Value)

	case "Annotation":
		if totaldigit.Annotation != nil {
			res = __gong__marshallPointer(ident, "Annotation", totaldigit.Annotation.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Annotation", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct TotalDigit", fieldName)
	}
	return
}

func (union *Union) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := union.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", union.Name)
	case "MemberTypes":
		res = __gong__marshallString(ident, "MemberTypes", union.MemberTypes)

	case "Annotation":
		if union.Annotation != nil {
			res = __gong__marshallPointer(ident, "Annotation", union.Annotation.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Annotation", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Union", fieldName)
	}
	return
}

func (whitespace *WhiteSpace) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := whitespace.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", whitespace.Name)
	case "Value":
		res = __gong__marshallString(ident, "Value", whitespace.Value)

	case "Annotation":
		if whitespace.Annotation != nil {
			res = __gong__marshallPointer(ident, "Annotation", whitespace.Annotation.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Annotation", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct WhiteSpace", fieldName)
	}
	return
}

// insertion point for marshall all fields methods
func (all *All) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(all.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(all.GongMarshallField(stage, "Annotation"))
		initializerStatements.WriteString(all.GongMarshallField(stage, "OuterElementName"))
		pointersInitializesStatements.WriteString(all.GongMarshallField(stage, "Sequences"))
		pointersInitializesStatements.WriteString(all.GongMarshallField(stage, "Alls"))
		pointersInitializesStatements.WriteString(all.GongMarshallField(stage, "Choices"))
		pointersInitializesStatements.WriteString(all.GongMarshallField(stage, "Groups"))
		pointersInitializesStatements.WriteString(all.GongMarshallField(stage, "Elements"))
		initializerStatements.WriteString(all.GongMarshallField(stage, "Order"))
		initializerStatements.WriteString(all.GongMarshallField(stage, "Depth"))
		initializerStatements.WriteString(all.GongMarshallField(stage, "MinOccurs"))
		initializerStatements.WriteString(all.GongMarshallField(stage, "MaxOccurs"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (annotation *Annotation) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(annotation.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(annotation.GongMarshallField(stage, "Documentations"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (attribute *Attribute) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(attribute.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(attribute.GongMarshallField(stage, "NameXSD"))
		initializerStatements.WriteString(attribute.GongMarshallField(stage, "Type"))
		pointersInitializesStatements.WriteString(attribute.GongMarshallField(stage, "Annotation"))
		initializerStatements.WriteString(attribute.GongMarshallField(stage, "HasNameConflict"))
		initializerStatements.WriteString(attribute.GongMarshallField(stage, "GoIdentifier"))
		initializerStatements.WriteString(attribute.GongMarshallField(stage, "Default"))
		initializerStatements.WriteString(attribute.GongMarshallField(stage, "Use"))
		initializerStatements.WriteString(attribute.GongMarshallField(stage, "Form"))
		initializerStatements.WriteString(attribute.GongMarshallField(stage, "Fixed"))
		initializerStatements.WriteString(attribute.GongMarshallField(stage, "Ref"))
		initializerStatements.WriteString(attribute.GongMarshallField(stage, "TargetNamespace"))
		initializerStatements.WriteString(attribute.GongMarshallField(stage, "SimpleType"))
		initializerStatements.WriteString(attribute.GongMarshallField(stage, "IDXSD"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (attributegroup *AttributeGroup) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(attributegroup.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(attributegroup.GongMarshallField(stage, "NameXSD"))
		pointersInitializesStatements.WriteString(attributegroup.GongMarshallField(stage, "Annotation"))
		initializerStatements.WriteString(attributegroup.GongMarshallField(stage, "HasNameConflict"))
		initializerStatements.WriteString(attributegroup.GongMarshallField(stage, "GoIdentifier"))
		pointersInitializesStatements.WriteString(attributegroup.GongMarshallField(stage, "AttributeGroups"))
		initializerStatements.WriteString(attributegroup.GongMarshallField(stage, "Ref"))
		pointersInitializesStatements.WriteString(attributegroup.GongMarshallField(stage, "Attributes"))
		initializerStatements.WriteString(attributegroup.GongMarshallField(stage, "Order"))
		initializerStatements.WriteString(attributegroup.GongMarshallField(stage, "Depth"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (choice *Choice) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(choice.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(choice.GongMarshallField(stage, "Annotation"))
		initializerStatements.WriteString(choice.GongMarshallField(stage, "OuterElementName"))
		pointersInitializesStatements.WriteString(choice.GongMarshallField(stage, "Sequences"))
		pointersInitializesStatements.WriteString(choice.GongMarshallField(stage, "Alls"))
		pointersInitializesStatements.WriteString(choice.GongMarshallField(stage, "Choices"))
		pointersInitializesStatements.WriteString(choice.GongMarshallField(stage, "Groups"))
		pointersInitializesStatements.WriteString(choice.GongMarshallField(stage, "Elements"))
		initializerStatements.WriteString(choice.GongMarshallField(stage, "Order"))
		initializerStatements.WriteString(choice.GongMarshallField(stage, "Depth"))
		initializerStatements.WriteString(choice.GongMarshallField(stage, "MinOccurs"))
		initializerStatements.WriteString(choice.GongMarshallField(stage, "MaxOccurs"))
		initializerStatements.WriteString(choice.GongMarshallField(stage, "IsDuplicatedInXSD"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (complexcontent *ComplexContent) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(complexcontent.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (complextype *ComplexType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(complextype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(complextype.GongMarshallField(stage, "HasNameConflict"))
		initializerStatements.WriteString(complextype.GongMarshallField(stage, "GoIdentifier"))
		initializerStatements.WriteString(complextype.GongMarshallField(stage, "IsAnonymous"))
		pointersInitializesStatements.WriteString(complextype.GongMarshallField(stage, "OuterElement"))
		pointersInitializesStatements.WriteString(complextype.GongMarshallField(stage, "Annotation"))
		initializerStatements.WriteString(complextype.GongMarshallField(stage, "NameXSD"))
		initializerStatements.WriteString(complextype.GongMarshallField(stage, "OuterElementName"))
		pointersInitializesStatements.WriteString(complextype.GongMarshallField(stage, "Sequences"))
		pointersInitializesStatements.WriteString(complextype.GongMarshallField(stage, "Alls"))
		pointersInitializesStatements.WriteString(complextype.GongMarshallField(stage, "Choices"))
		pointersInitializesStatements.WriteString(complextype.GongMarshallField(stage, "Groups"))
		pointersInitializesStatements.WriteString(complextype.GongMarshallField(stage, "Elements"))
		initializerStatements.WriteString(complextype.GongMarshallField(stage, "Order"))
		initializerStatements.WriteString(complextype.GongMarshallField(stage, "Depth"))
		initializerStatements.WriteString(complextype.GongMarshallField(stage, "MinOccurs"))
		initializerStatements.WriteString(complextype.GongMarshallField(stage, "MaxOccurs"))
		pointersInitializesStatements.WriteString(complextype.GongMarshallField(stage, "Extension"))
		pointersInitializesStatements.WriteString(complextype.GongMarshallField(stage, "SimpleContent"))
		pointersInitializesStatements.WriteString(complextype.GongMarshallField(stage, "ComplexContent"))
		pointersInitializesStatements.WriteString(complextype.GongMarshallField(stage, "Attributes"))
		pointersInitializesStatements.WriteString(complextype.GongMarshallField(stage, "AttributeGroups"))
		initializerStatements.WriteString(complextype.GongMarshallField(stage, "IsDuplicatedInXSD"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (documentation *Documentation) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(documentation.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(documentation.GongMarshallField(stage, "Text"))
		initializerStatements.WriteString(documentation.GongMarshallField(stage, "Source"))
		initializerStatements.WriteString(documentation.GongMarshallField(stage, "Lang"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (element *Element) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(element.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(element.GongMarshallField(stage, "Order"))
		initializerStatements.WriteString(element.GongMarshallField(stage, "Depth"))
		initializerStatements.WriteString(element.GongMarshallField(stage, "HasNameConflict"))
		initializerStatements.WriteString(element.GongMarshallField(stage, "GoIdentifier"))
		pointersInitializesStatements.WriteString(element.GongMarshallField(stage, "Annotation"))
		initializerStatements.WriteString(element.GongMarshallField(stage, "NameXSD"))
		initializerStatements.WriteString(element.GongMarshallField(stage, "Type"))
		initializerStatements.WriteString(element.GongMarshallField(stage, "MinOccurs"))
		initializerStatements.WriteString(element.GongMarshallField(stage, "MaxOccurs"))
		initializerStatements.WriteString(element.GongMarshallField(stage, "Default"))
		initializerStatements.WriteString(element.GongMarshallField(stage, "Fixed"))
		initializerStatements.WriteString(element.GongMarshallField(stage, "Nillable"))
		initializerStatements.WriteString(element.GongMarshallField(stage, "Ref"))
		initializerStatements.WriteString(element.GongMarshallField(stage, "Abstract"))
		initializerStatements.WriteString(element.GongMarshallField(stage, "Form"))
		initializerStatements.WriteString(element.GongMarshallField(stage, "Block"))
		initializerStatements.WriteString(element.GongMarshallField(stage, "Final"))
		pointersInitializesStatements.WriteString(element.GongMarshallField(stage, "SimpleType"))
		pointersInitializesStatements.WriteString(element.GongMarshallField(stage, "ComplexType"))
		pointersInitializesStatements.WriteString(element.GongMarshallField(stage, "Groups"))
		initializerStatements.WriteString(element.GongMarshallField(stage, "IsDuplicatedInXSD"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (enumeration *Enumeration) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(enumeration.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(enumeration.GongMarshallField(stage, "Annotation"))
		initializerStatements.WriteString(enumeration.GongMarshallField(stage, "Value"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (extension *Extension) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(extension.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(extension.GongMarshallField(stage, "OuterElementName"))
		pointersInitializesStatements.WriteString(extension.GongMarshallField(stage, "Sequences"))
		pointersInitializesStatements.WriteString(extension.GongMarshallField(stage, "Alls"))
		pointersInitializesStatements.WriteString(extension.GongMarshallField(stage, "Choices"))
		pointersInitializesStatements.WriteString(extension.GongMarshallField(stage, "Groups"))
		pointersInitializesStatements.WriteString(extension.GongMarshallField(stage, "Elements"))
		initializerStatements.WriteString(extension.GongMarshallField(stage, "Order"))
		initializerStatements.WriteString(extension.GongMarshallField(stage, "Depth"))
		initializerStatements.WriteString(extension.GongMarshallField(stage, "MinOccurs"))
		initializerStatements.WriteString(extension.GongMarshallField(stage, "MaxOccurs"))
		initializerStatements.WriteString(extension.GongMarshallField(stage, "Base"))
		initializerStatements.WriteString(extension.GongMarshallField(stage, "Ref"))
		pointersInitializesStatements.WriteString(extension.GongMarshallField(stage, "Attributes"))
		pointersInitializesStatements.WriteString(extension.GongMarshallField(stage, "AttributeGroups"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (group *Group) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(group.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(group.GongMarshallField(stage, "Annotation"))
		initializerStatements.WriteString(group.GongMarshallField(stage, "NameXSD"))
		initializerStatements.WriteString(group.GongMarshallField(stage, "Ref"))
		initializerStatements.WriteString(group.GongMarshallField(stage, "IsAnonymous"))
		pointersInitializesStatements.WriteString(group.GongMarshallField(stage, "OuterElement"))
		initializerStatements.WriteString(group.GongMarshallField(stage, "HasNameConflict"))
		initializerStatements.WriteString(group.GongMarshallField(stage, "GoIdentifier"))
		initializerStatements.WriteString(group.GongMarshallField(stage, "OuterElementName"))
		pointersInitializesStatements.WriteString(group.GongMarshallField(stage, "Sequences"))
		pointersInitializesStatements.WriteString(group.GongMarshallField(stage, "Alls"))
		pointersInitializesStatements.WriteString(group.GongMarshallField(stage, "Choices"))
		pointersInitializesStatements.WriteString(group.GongMarshallField(stage, "Groups"))
		pointersInitializesStatements.WriteString(group.GongMarshallField(stage, "Elements"))
		initializerStatements.WriteString(group.GongMarshallField(stage, "Order"))
		initializerStatements.WriteString(group.GongMarshallField(stage, "Depth"))
		initializerStatements.WriteString(group.GongMarshallField(stage, "MinOccurs"))
		initializerStatements.WriteString(group.GongMarshallField(stage, "MaxOccurs"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (length *Length) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(length.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(length.GongMarshallField(stage, "Annotation"))
		initializerStatements.WriteString(length.GongMarshallField(stage, "Value"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (maxinclusive *MaxInclusive) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(maxinclusive.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(maxinclusive.GongMarshallField(stage, "Annotation"))
		initializerStatements.WriteString(maxinclusive.GongMarshallField(stage, "Value"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (maxlength *MaxLength) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(maxlength.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(maxlength.GongMarshallField(stage, "Annotation"))
		initializerStatements.WriteString(maxlength.GongMarshallField(stage, "Value"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (mininclusive *MinInclusive) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(mininclusive.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(mininclusive.GongMarshallField(stage, "Annotation"))
		initializerStatements.WriteString(mininclusive.GongMarshallField(stage, "Value"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (minlength *MinLength) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(minlength.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(minlength.GongMarshallField(stage, "Annotation"))
		initializerStatements.WriteString(minlength.GongMarshallField(stage, "Value"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (pattern *Pattern) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(pattern.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(pattern.GongMarshallField(stage, "Annotation"))
		initializerStatements.WriteString(pattern.GongMarshallField(stage, "Value"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (restriction *Restriction) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(restriction.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(restriction.GongMarshallField(stage, "Annotation"))
		initializerStatements.WriteString(restriction.GongMarshallField(stage, "Base"))
		pointersInitializesStatements.WriteString(restriction.GongMarshallField(stage, "Enumerations"))
		pointersInitializesStatements.WriteString(restriction.GongMarshallField(stage, "MinInclusive"))
		pointersInitializesStatements.WriteString(restriction.GongMarshallField(stage, "MaxInclusive"))
		pointersInitializesStatements.WriteString(restriction.GongMarshallField(stage, "Pattern"))
		pointersInitializesStatements.WriteString(restriction.GongMarshallField(stage, "WhiteSpace"))
		pointersInitializesStatements.WriteString(restriction.GongMarshallField(stage, "MinLength"))
		pointersInitializesStatements.WriteString(restriction.GongMarshallField(stage, "MaxLength"))
		pointersInitializesStatements.WriteString(restriction.GongMarshallField(stage, "Length"))
		pointersInitializesStatements.WriteString(restriction.GongMarshallField(stage, "TotalDigit"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (schema *Schema) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(schema.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(schema.GongMarshallField(stage, "Xs"))
		pointersInitializesStatements.WriteString(schema.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(schema.GongMarshallField(stage, "Elements"))
		pointersInitializesStatements.WriteString(schema.GongMarshallField(stage, "SimpleTypes"))
		pointersInitializesStatements.WriteString(schema.GongMarshallField(stage, "ComplexTypes"))
		pointersInitializesStatements.WriteString(schema.GongMarshallField(stage, "AttributeGroups"))
		pointersInitializesStatements.WriteString(schema.GongMarshallField(stage, "Groups"))
		initializerStatements.WriteString(schema.GongMarshallField(stage, "Order"))
		initializerStatements.WriteString(schema.GongMarshallField(stage, "Depth"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (sequence *Sequence) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(sequence.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(sequence.GongMarshallField(stage, "Annotation"))
		initializerStatements.WriteString(sequence.GongMarshallField(stage, "OuterElementName"))
		pointersInitializesStatements.WriteString(sequence.GongMarshallField(stage, "Sequences"))
		pointersInitializesStatements.WriteString(sequence.GongMarshallField(stage, "Alls"))
		pointersInitializesStatements.WriteString(sequence.GongMarshallField(stage, "Choices"))
		pointersInitializesStatements.WriteString(sequence.GongMarshallField(stage, "Groups"))
		pointersInitializesStatements.WriteString(sequence.GongMarshallField(stage, "Elements"))
		initializerStatements.WriteString(sequence.GongMarshallField(stage, "Order"))
		initializerStatements.WriteString(sequence.GongMarshallField(stage, "Depth"))
		initializerStatements.WriteString(sequence.GongMarshallField(stage, "MinOccurs"))
		initializerStatements.WriteString(sequence.GongMarshallField(stage, "MaxOccurs"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (simplecontent *SimpleContent) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(simplecontent.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(simplecontent.GongMarshallField(stage, "Extension"))
		pointersInitializesStatements.WriteString(simplecontent.GongMarshallField(stage, "Restriction"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (simpletype *SimpleType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(simpletype.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(simpletype.GongMarshallField(stage, "Annotation"))
		initializerStatements.WriteString(simpletype.GongMarshallField(stage, "NameXSD"))
		pointersInitializesStatements.WriteString(simpletype.GongMarshallField(stage, "Restriction"))
		pointersInitializesStatements.WriteString(simpletype.GongMarshallField(stage, "Union"))
		initializerStatements.WriteString(simpletype.GongMarshallField(stage, "Order"))
		initializerStatements.WriteString(simpletype.GongMarshallField(stage, "Depth"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (totaldigit *TotalDigit) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(totaldigit.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(totaldigit.GongMarshallField(stage, "Annotation"))
		initializerStatements.WriteString(totaldigit.GongMarshallField(stage, "Value"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (union *Union) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(union.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(union.GongMarshallField(stage, "Annotation"))
		initializerStatements.WriteString(union.GongMarshallField(stage, "MemberTypes"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (whitespace *WhiteSpace) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(whitespace.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(whitespace.GongMarshallField(stage, "Annotation"))
		initializerStatements.WriteString(whitespace.GongMarshallField(stage, "Value"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
