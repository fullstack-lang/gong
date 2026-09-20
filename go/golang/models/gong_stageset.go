package models

import (
	"fmt"
	"go/types"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/fullstack-lang/gong/go/models"
)

type StageSetStructField struct {
	Name              string
	IsPointer         bool
	IsSliceOfPointer  bool
	TargetPackagePath string
	TargetStructName  string
	BasicKind         types.BasicKind
	IsEnum            bool
	EnumTypeQual      string
	IsDuration        bool
	IsTime            bool
}

func ExtractStructFields(mPkg *models.ModelPkg, structName string) []StageSetStructField {
	var result []StageSetStructField
	if mPkg.TypesPkg == nil {
		return result
	}
	obj := mPkg.TypesPkg.Scope().Lookup(structName)
	if obj == nil {
		return result
	}
	named, ok := obj.Type().(*types.Named)
	if !ok {
		return result
	}
	st, ok := named.Underlying().(*types.Struct)
	if !ok {
		return result
	}

	var inspectFields func(st *types.Struct)
	inspectFields = func(st *types.Struct) {
		for i := 0; i < st.NumFields(); i++ {
			fld := st.Field(i)
			if fld.Anonymous() {
				if embNamed, ok := fld.Type().(*types.Named); ok {
					if embSt, ok := embNamed.Underlying().(*types.Struct); ok {
						inspectFields(embSt)
					}
				}
				continue
			}
			// check pointer
			if ptr, ok := fld.Type().(*types.Pointer); ok {
				if targetNamed, ok := ptr.Elem().(*types.Named); ok {
					if targetPkg := targetNamed.Obj().Pkg(); targetPkg != nil {
						result = append(result, StageSetStructField{
							Name:              fld.Name(),
							IsPointer:         true,
							TargetPackagePath: targetPkg.Path(),
							TargetStructName:  targetNamed.Obj().Name(),
						})
						continue
					}
				}
			}
			// check slice of pointer
			if sl, ok := fld.Type().(*types.Slice); ok {
				if ptr, ok := sl.Elem().(*types.Pointer); ok {
					if targetNamed, ok := ptr.Elem().(*types.Named); ok {
						if targetPkg := targetNamed.Obj().Pkg(); targetPkg != nil {
							result = append(result, StageSetStructField{
								Name:              fld.Name(),
								IsSliceOfPointer:  true,
								TargetPackagePath: targetPkg.Path(),
								TargetStructName:  targetNamed.Obj().Name(),
							})
							continue
						}
					}
				}
			}
			// check basic type
			if basic, ok := fld.Type().(*types.Basic); ok {
				result = append(result, StageSetStructField{
					Name:      fld.Name(),
					BasicKind: basic.Kind(),
				})
				continue
			}
			// check named basic (enum or time.Duration / time.Time)
			if namedType, ok := fld.Type().(*types.Named); ok {
				if namedType.Obj().Pkg() != nil && namedType.Obj().Pkg().Path() == "time" && namedType.Obj().Name() == "Duration" {
					result = append(result, StageSetStructField{
						Name:       fld.Name(),
						IsDuration: true,
					})
					continue
				}
				if namedType.Obj().Pkg() != nil && namedType.Obj().Pkg().Path() == "time" && namedType.Obj().Name() == "Time" {
					result = append(result, StageSetStructField{
						Name:   fld.Name(),
						IsTime: true,
					})
					continue
				}
				if basic, ok := namedType.Underlying().(*types.Basic); ok {
					qual := namedType.Obj().Name()
					if namedType.Obj().Pkg() != nil && namedType.Obj().Pkg().Path() != mPkg.PkgPath {
						qual = namedType.Obj().Pkg().Name() + "." + qual
					}
					result = append(result, StageSetStructField{
						Name:         fld.Name(),
						IsEnum:       true,
						BasicKind:    basic.Kind(),
						EnumTypeQual: qual,
					})
					continue
				}
			}
		}
	}

	inspectFields(st)
	return result
}

func CodeGeneratorModelGongStageSet(
	modelPkg *models.ModelPkg,
	pkgName string,
	pkgPath string,
) {
	stageSet := modelPkg.StageSet
	if stageSet == nil || len(stageSet.Fields) == 0 {
		return
	}

	goModDir, modPath := models.FindGoMod(pkgPath)
	depPkgPaths, _ := models.DiscoverModelDependencies(pkgPath)

	fieldToModelPkg := make(map[*models.StageSetField]*models.ModelPkg)
	pkgPathToField := make(map[string]*models.StageSetField)

	for _, f := range stageSet.Fields {
		pkgPathToField[f.PackagePath] = f
		if f.IsLocal {
			fieldToModelPkg[f] = modelPkg
		} else {
			if goModDir != "" && modPath != "" && strings.HasPrefix(f.PackagePath, modPath) {
				rel := strings.TrimPrefix(f.PackagePath, modPath)
				rel = strings.TrimPrefix(rel, "/")
				depDir := filepath.Join(goModDir, filepath.FromSlash(rel))
				depStage := models.NewStage("")
				depPkg, err := models.LoadSource(depStage, depDir)
				if err == nil {
					fieldToModelPkg[f] = depPkg
				}
			}
		}
	}

	// Sort fields in topological order: dependencies (leaves first), then local package
	orderedFields := make([]*models.StageSetField, len(stageSet.Fields))
	copy(orderedFields, stageSet.Fields)
	sort.SliceStable(orderedFields, func(i, j int) bool {
		fi, fj := orderedFields[i], orderedFields[j]
		if fi.IsLocal != fj.IsLocal {
			return !fi.IsLocal
		}
		idxI, idxJ := -1, -1
		for idx, p := range depPkgPaths {
			_, fullPath := models.ComputePkgPathFromGoModFile(p)
			if fullPath == fi.PackagePath {
				idxI = idx
			}
			if fullPath == fj.PackagePath {
				idxJ = idx
			}
		}
		return idxI < idxJ
	})

	var externalImports strings.Builder
	for _, f := range stageSet.Fields {
		if !f.IsLocal {
			externalImports.WriteString(fmt.Sprintf("\n\t\"%s\"", f.PackagePath))
		}
	}

	var commitStatements strings.Builder
	var checkoutStatements strings.Builder
	var resetStatements strings.Builder
	var cleanStatements strings.Builder
	var computeReverseMapsStatements strings.Builder
	var computeInstancesNbStatements strings.Builder
	var computeReferenceAndOrdersStatements strings.Builder
	for _, f := range orderedFields {
		commitStatements.WriteString(fmt.Sprintf("\tif stageSet.%s != nil {\n\t\tstageSet.%s.Commit()\n\t}\n", f.Name, f.Name))
		checkoutStatements.WriteString(fmt.Sprintf("\tif stageSet.%s != nil {\n\t\tstageSet.%s.Checkout()\n\t}\n", f.Name, f.Name))
		resetStatements.WriteString(fmt.Sprintf("\tif stageSet.%s != nil {\n\t\tstageSet.%s.Reset()\n\t}\n", f.Name, f.Name))
		cleanStatements.WriteString(fmt.Sprintf("\tif stageSet.%s != nil {\n\t\tstageSet.%s.Clean()\n\t}\n", f.Name, f.Name))
		computeReverseMapsStatements.WriteString(fmt.Sprintf("\tif stageSet.%s != nil {\n\t\tstageSet.%s.ComputeReverseMaps()\n\t}\n", f.Name, f.Name))
		computeInstancesNbStatements.WriteString(fmt.Sprintf("\tif stageSet.%s != nil {\n\t\tstageSet.%s.ComputeInstancesNb()\n\t}\n", f.Name, f.Name))
		computeReferenceAndOrdersStatements.WriteString(fmt.Sprintf("\tif stageSet.%s != nil {\n\t\tstageSet.%s.ComputeReferenceAndOrders()\n\t}\n", f.Name, f.Name))
	}

	var newStageStatements strings.Builder
	var newStageFromStageStatements strings.Builder
	localFieldName := "Stage"
	for _, f := range stageSet.Fields {
		if f.IsLocal {
			localFieldName = f.Name
			newStageStatements.WriteString(fmt.Sprintf("\tstageSet.%s = NewStage(path)\n", f.Name))
			newStageFromStageStatements.WriteString(fmt.Sprintf("\tstageSet.%s = stage\n", f.Name))
		} else {
			newStageStatements.WriteString(fmt.Sprintf("\tsubPath_%s := \"%s\"\n\tif path != \"\" {\n\tsubPath_%s = path + \"_%s\"\n\t}\n\tstageSet.%s = %s.NewStage(subPath_%s)\n", f.Name, f.PackageName, f.Name, f.PackageName, f.Name, f.PackageName, f.Name))
			newStageFromStageStatements.WriteString(fmt.Sprintf("\tsubPath_%s := \"%s\"\n\tif stage != nil && stage.GetName() != \"\" {\n\tsubPath_%s = stage.GetName() + \"_%s\"\n\t}\n\tstageSet.%s = %s.NewStage(subPath_%s)\n", f.Name, f.PackageName, f.Name, f.PackageName, f.Name, f.PackageName, f.Name))
		}
	}

	var syntheticImports strings.Builder
	var dummyDeclarations strings.Builder
	for _, f := range stageSet.Fields {
		syntheticImports.WriteString(fmt.Sprintf("\n\t%s \"%s\"", f.ImportAlias, f.PackagePath))
		dummyDeclarations.WriteString(fmt.Sprintf("\n\t_ *%s.Stage", f.ImportAlias))
	}

	var mainPkgImportAlias string
	for _, f := range stageSet.Fields {
		if f.IsLocal {
			mainPkgImportAlias = f.ImportAlias
			break
		}
	}
	if mainPkgImportAlias == "" && len(stageSet.Fields) > 0 {
		mainPkgImportAlias = stageSet.Fields[0].ImportAlias
	}

	// Build code inside MarshallToString
	var marshallBody strings.Builder
	marshallBody.WriteString("\tvar declarations strings.Builder\n")
	marshallBody.WriteString("\tvar values strings.Builder\n")
	marshallBody.WriteString("\tvar pointers strings.Builder\n\n")

	for _, f := range orderedFields {
		mPkg := fieldToModelPkg[f]
		if mPkg == nil {
			continue
		}

		// Sort structs
		var structNames []string
		for sName, gs := range mPkg.GongStructs {
			if gs.ModelPkg.PkgPath == mPkg.PkgPath && gs.HasNameField() && !gs.IsOmittedForMarshalling {
				structNames = append(structNames, sName[len(mPkg.PkgPath)+1:])
			}
		}
		sort.Strings(structNames)

		aliasPrefix := f.ImportAlias[:len(f.ImportAlias)-2] // strip trailing __

		for _, sName := range structNames {
			sVar := strings.ToLower(sName)
			sPlural := sName + "s"
			typeQual := sName
			if !f.IsLocal {
				typeQual = f.PackageName + "." + sName
			}

			fields := ExtractStructFields(mPkg, sName)

			marshallBody.WriteString(fmt.Sprintf("\tif stageSet.%s != nil {\n", f.Name))
			marshallBody.WriteString(fmt.Sprintf("\t\t%sOrdered := []*%s{}\n", sVar, typeQual))
			marshallBody.WriteString(fmt.Sprintf("\t\tfor %s := range stageSet.%s.%s {\n", sVar, f.Name, sPlural))
			marshallBody.WriteString(fmt.Sprintf("\t\t\t%sOrdered = append(%sOrdered, %s)\n", sVar, sVar, sVar))
			marshallBody.WriteString("\t\t}\n")
			marshallBody.WriteString(fmt.Sprintf("\t\tsort.Slice(%sOrdered, func(i, j int) bool {\n", sVar))
			marshallBody.WriteString(fmt.Sprintf("\t\t\treturn stageSet.%s.%s_stagedOrder[%sOrdered[i]] < stageSet.%s.%s_stagedOrder[%sOrdered[j]]\n", f.Name, sName, sVar, f.Name, sName, sVar))
			marshallBody.WriteString("\t\t})\n")

			marshallBody.WriteString(fmt.Sprintf("\t\tfor _, %s := range %sOrdered {\n", sVar, sVar))
			marshallBody.WriteString(fmt.Sprintf("\t\t\t%sIdent := \"%s\" + %s.GongGetIdentifier(stageSet.%s)\n", sVar, aliasPrefix, sVar, f.Name))
			// Phase 1 declaration
			marshallBody.WriteString(fmt.Sprintf("\t\t\tdeclarations.WriteString(fmt.Sprintf(\"\\n\\t%%s := (&%s.%s{Name: %%s}).Stage(stageSet.%s)\", %sIdent, __gong__toRawStringLiteral(%s.Name)))\n", f.ImportAlias, sName, f.Name, sVar, sVar))

			// Phase 2 values
			for _, fld := range fields {
				if fld.IsPointer || fld.IsSliceOfPointer {
					continue
				}
				switch fld.BasicKind {
				case types.String:
					marshallBody.WriteString(fmt.Sprintf("\t\t\tvalues.WriteString(fmt.Sprintf(\"\\n\\t%%s.%s = %%s\", %sIdent, __gong__toRawStringLiteral(%s.%s)))\n", fld.Name, sVar, sVar, fld.Name))
				case types.Int, types.Int64, types.Int32, types.Int16, types.Int8, types.Uint, types.Uint64, types.Uint32, types.Uint16, types.Uint8:
					if fld.IsEnum {
						marshallBody.WriteString(fmt.Sprintf("\t\t\tvalues.WriteString(fmt.Sprintf(\"\\n\\t%%s.%s = %%d\", %sIdent, int(%s.%s)))\n", fld.Name, sVar, sVar, fld.Name))
					} else {
						marshallBody.WriteString(fmt.Sprintf("\t\t\tvalues.WriteString(fmt.Sprintf(\"\\n\\t%%s.%s = %%d\", %sIdent, %s.%s))\n", fld.Name, sVar, sVar, fld.Name))
					}
				case types.Float64, types.Float32:
					marshallBody.WriteString(fmt.Sprintf("\t\t\tvalues.WriteString(fmt.Sprintf(\"\\n\\t%%s.%s = %%f\", %sIdent, %s.%s))\n", fld.Name, sVar, sVar, fld.Name))
				case types.Bool:
					marshallBody.WriteString(fmt.Sprintf("\t\t\tvalues.WriteString(fmt.Sprintf(\"\\n\\t%%s.%s = %%t\", %sIdent, %s.%s))\n", fld.Name, sVar, sVar, fld.Name))
				default:
					if fld.IsTime {
						marshallBody.WriteString(fmt.Sprintf("\t\t\tvalues.WriteString(fmt.Sprintf(\"\\n\\t%%s.%s, _ = time.Parse(\\\"2006-01-02 15:04:05.999999999 -0700 MST\\\", \\\"%%s\\\")\", %sIdent, %s.%s.String()))\n", fld.Name, sVar, sVar, fld.Name))
					} else if fld.IsDuration {
						marshallBody.WriteString(fmt.Sprintf("\t\t\tvalues.WriteString(fmt.Sprintf(\"\\n\\t%%s.%s = time.Duration(%%d)\", %sIdent, int64(%s.%s)))\n", fld.Name, sVar, sVar, fld.Name))
					} else if fld.IsEnum {
						marshallBody.WriteString(fmt.Sprintf("\t\t\tvalues.WriteString(fmt.Sprintf(\"\\n\\t%%s.%s = %%s\", %sIdent, __gong__toRawStringLiteral(string(%s.%s))))\n", fld.Name, sVar, sVar, fld.Name))
					}
				}
			}

			// Phase 3 pointers
			for _, fld := range fields {
				if fld.IsPointer {
					targetSSF := pkgPathToField[fld.TargetPackagePath]
					if targetSSF != nil {
						targetAliasPrefix := targetSSF.ImportAlias[:len(targetSSF.ImportAlias)-2]
						marshallBody.WriteString(fmt.Sprintf("\t\t\tif %s.%s != nil {\n", sVar, fld.Name))
						marshallBody.WriteString(fmt.Sprintf("\t\t\t\ttargetIdent := \"%s\" + %s.%s.GongGetIdentifier(stageSet.%s)\n", targetAliasPrefix, sVar, fld.Name, targetSSF.Name))
						marshallBody.WriteString(fmt.Sprintf("\t\t\t\tpointers.WriteString(fmt.Sprintf(\"\\n\\t%%s.%s = %%s\", %sIdent, targetIdent))\n", fld.Name, sVar))
						marshallBody.WriteString("\t\t\t}\n")
					}
				} else if fld.IsSliceOfPointer {
					targetSSF := pkgPathToField[fld.TargetPackagePath]
					if targetSSF != nil {
						targetAliasPrefix := targetSSF.ImportAlias[:len(targetSSF.ImportAlias)-2]
						marshallBody.WriteString(fmt.Sprintf("\t\t\tfor _, elem := range %s.%s {\n", sVar, fld.Name))
						marshallBody.WriteString(fmt.Sprintf("\t\t\t\ttargetIdent := \"%s\" + elem.GongGetIdentifier(stageSet.%s)\n", targetAliasPrefix, targetSSF.Name))
						marshallBody.WriteString(fmt.Sprintf("\t\t\t\tpointers.WriteString(fmt.Sprintf(\"\\n\\t%%s.%s = append(%%s.%s, %%s)\", %sIdent, %sIdent, targetIdent))\n", fld.Name, fld.Name, sVar, sVar))
						marshallBody.WriteString("\t\t\t}\n")
					}
				}
			}

			marshallBody.WriteString("\t\t}\n")
			marshallBody.WriteString("\t}\n")
		}
	}

	// Build AST Parsing (ParseAstFileFromAst)
	var defineCases strings.Builder
	var assignCases strings.Builder

	for _, f := range orderedFields {
		mPkg := fieldToModelPkg[f]
		if mPkg == nil {
			continue
		}
		var structNames []string
		for sName, gs := range mPkg.GongStructs {
			if gs.ModelPkg.PkgPath == mPkg.PkgPath && gs.HasNameField() && !gs.IsOmittedForMarshalling {
				structNames = append(structNames, sName[len(mPkg.PkgPath)+1:])
			}
		}
		sort.Strings(structNames)

		defineCases.WriteString(fmt.Sprintf("\t\t\tcase \"%s\":\n", f.ImportAlias))
		defineCases.WriteString("\t\t\t\tswitch typeName {\n")
		for _, sName := range structNames {
			typeQual := sName
			if !f.IsLocal {
				typeQual = f.PackageName + "." + sName
			}
			defineCases.WriteString(fmt.Sprintf("\t\t\t\tcase \"%s\":\n", sName))
			defineCases.WriteString("\t\t\t\t\tif !preserveOrder {\n")
			defineCases.WriteString(fmt.Sprintf("\t\t\t\t\t\tinst := (&%s{Name: instanceName}).Stage(stageSet.%s)\n", typeQual, f.Name))
			defineCases.WriteString("\t\t\t\t\t\tidentifierMap[ident.Name] = inst\n")
			defineCases.WriteString("\t\t\t\t\t} else {\n")
			defineCases.WriteString(fmt.Sprintf("\t\t\t\t\t\tinst := new(%s)\n", typeQual))
			defineCases.WriteString("\t\t\t\t\t\tinst.Name = instanceName\n")
			defineCases.WriteString("\t\t\t\t\t\torder, _ := __gong__extractMiddleUint(ident.Name)\n")
			defineCases.WriteString(fmt.Sprintf("\t\t\t\t\t\tinst.StagePreserveOrder(stageSet.%s, uint(order))\n", f.Name))
			defineCases.WriteString("\t\t\t\t\t\tidentifierMap[ident.Name] = inst\n")
			defineCases.WriteString("\t\t\t\t\t}\n")
		}
		defineCases.WriteString("\t\t\t\t}\n")

		for _, sName := range structNames {
			typeQual := sName
			if !f.IsLocal {
				typeQual = f.PackageName + "." + sName
			}
			fields := ExtractStructFields(mPkg, sName)

			assignCases.WriteString(fmt.Sprintf("\t\t\t\tcase *%s:\n", typeQual))
			assignCases.WriteString("\t\t\t\t\tswitch fieldName {\n")
			for _, fld := range fields {
				if fld.IsPointer {
					targetSSF := pkgPathToField[fld.TargetPackagePath]
					if targetSSF != nil {
						targetQual := fld.TargetStructName
						if !targetSSF.IsLocal {
							targetQual = targetSSF.PackageName + "." + fld.TargetStructName
						}
						assignCases.WriteString(fmt.Sprintf("\t\t\t\t\tcase \"%s\":\n", fld.Name))
						assignCases.WriteString("\t\t\t\t\t\tif rIdent, ok := rhs.(*ast.Ident); ok {\n")
						assignCases.WriteString("\t\t\t\t\t\t\tif target, ok := identifierMap[rIdent.Name]; ok {\n")
						assignCases.WriteString(fmt.Sprintf("\t\t\t\t\t\t\t\tif typedTarget, ok := target.(*%s); ok {\n", targetQual))
						assignCases.WriteString(fmt.Sprintf("\t\t\t\t\t\t\t\t\tinst.%s = typedTarget\n", fld.Name))
						assignCases.WriteString("\t\t\t\t\t\t\t\t}\n")
						assignCases.WriteString("\t\t\t\t\t\t\t}\n")
						assignCases.WriteString("\t\t\t\t\t\t}\n")
					}
				} else if fld.IsSliceOfPointer {
					targetSSF := pkgPathToField[fld.TargetPackagePath]
					if targetSSF != nil {
						targetQual := fld.TargetStructName
						if !targetSSF.IsLocal {
							targetQual = targetSSF.PackageName + "." + fld.TargetStructName
						}
						assignCases.WriteString(fmt.Sprintf("\t\t\t\t\tcase \"%s\":\n", fld.Name))
						assignCases.WriteString("\t\t\t\t\t\tif call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {\n")
						assignCases.WriteString("\t\t\t\t\t\t\tif rIdent, ok := call.Args[1].(*ast.Ident); ok {\n")
						assignCases.WriteString("\t\t\t\t\t\t\t\tif target, ok := identifierMap[rIdent.Name]; ok {\n")
						assignCases.WriteString(fmt.Sprintf("\t\t\t\t\t\t\t\t\tif typedTarget, ok := target.(*%s); ok {\n", targetQual))
						assignCases.WriteString(fmt.Sprintf("\t\t\t\t\t\t\t\t\t\tinst.%s = append(inst.%s, typedTarget)\n", fld.Name, fld.Name))
						assignCases.WriteString("\t\t\t\t\t\t\t\t\t}\n")
						assignCases.WriteString("\t\t\t\t\t\t\t\t}\n")
						assignCases.WriteString("\t\t\t\t\t\t\t}\n")
						assignCases.WriteString("\t\t\t\t\t\t}\n")
					}
				} else {
					switch fld.BasicKind {
					case types.String:
						assignCases.WriteString(fmt.Sprintf("\t\t\t\t\tcase \"%s\":\n\t\t\t\t\t\tinst.%s = GongExtractString(rhs)\n", fld.Name, fld.Name))
					case types.Int, types.Int64, types.Int32, types.Int16, types.Int8, types.Uint, types.Uint64, types.Uint32, types.Uint16, types.Uint8:
						if fld.IsEnum {
							assignCases.WriteString(fmt.Sprintf("\t\t\t\t\tcase \"%s\":\n\t\t\t\t\t\tinst.%s = %s(GongExtractInt(rhs))\n", fld.Name, fld.Name, fld.EnumTypeQual))
						} else {
							assignCases.WriteString(fmt.Sprintf("\t\t\t\t\tcase \"%s\":\n\t\t\t\t\t\tinst.%s = GongExtractInt(rhs)\n", fld.Name, fld.Name))
						}
					case types.Float64, types.Float32:
						assignCases.WriteString(fmt.Sprintf("\t\t\t\t\tcase \"%s\":\n\t\t\t\t\t\tinst.%s = GongExtractFloat(rhs)\n", fld.Name, fld.Name))
					case types.Bool:
						assignCases.WriteString(fmt.Sprintf("\t\t\t\t\tcase \"%s\":\n\t\t\t\t\t\tinst.%s = GongExtractBool(rhs)\n", fld.Name, fld.Name))
					default:
						if fld.IsTime {
							assignCases.WriteString(fmt.Sprintf("\t\t\t\t\tcase \"%s\":\n\t\t\t\t\t\tif call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {\n\t\t\t\t\t\t\tif bl, ok := call.Args[1].(*ast.BasicLit); ok {\n\t\t\t\t\t\t\t\tinst.%s, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", strings.Trim(bl.Value, \"\\\"`\"))\n\t\t\t\t\t\t\t}\n\t\t\t\t\t\t}\n", fld.Name, fld.Name))
						} else if fld.IsDuration {
							assignCases.WriteString(fmt.Sprintf("\t\t\t\t\tcase \"%s\":\n\t\t\t\t\t\tinst.%s = time.Duration(GongExtractInt(rhs))\n", fld.Name, fld.Name))
						} else if fld.IsEnum {
							assignCases.WriteString(fmt.Sprintf("\t\t\t\t\tcase \"%s\":\n\t\t\t\t\t\tinst.%s = %s(GongExtractString(rhs))\n", fld.Name, fld.Name, fld.EnumTypeQual))
						}
					}
				}
			}
			assignCases.WriteString("\t\t\t\t\t}\n")
		}
	}

	codeGO := ModelGongStageSetTemplate
	codeGO = strings.ReplaceAll(codeGO, "{{PkgGoName}}", modelPkg.PkgGoName)
	codeGO = strings.ReplaceAll(codeGO, "{{ExternalImports}}", externalImports.String())
	codeGO = strings.ReplaceAll(codeGO, "{{CommitStatements}}", commitStatements.String())
	codeGO = strings.ReplaceAll(codeGO, "{{CheckoutStatements}}", checkoutStatements.String())
	codeGO = strings.ReplaceAll(codeGO, "{{ResetStatements}}", resetStatements.String())
	codeGO = strings.ReplaceAll(codeGO, "{{CleanStatements}}", cleanStatements.String())
	codeGO = strings.ReplaceAll(codeGO, "{{ComputeReverseMapsStatements}}", computeReverseMapsStatements.String())
	codeGO = strings.ReplaceAll(codeGO, "{{ComputeInstancesNbStatements}}", computeInstancesNbStatements.String())
	codeGO = strings.ReplaceAll(codeGO, "{{ComputeReferenceAndOrdersStatements}}", computeReferenceAndOrdersStatements.String())
	codeGO = strings.ReplaceAll(codeGO, "{{NewStageStatements}}", newStageStatements.String())
	codeGO = strings.ReplaceAll(codeGO, "{{NewStageFromStageStatements}}", newStageFromStageStatements.String())
	codeGO = strings.ReplaceAll(codeGO, "{{LocalFieldName}}", localFieldName)
	codeGO = strings.ReplaceAll(codeGO, "{{SyntheticImports}}", syntheticImports.String())
	codeGO = strings.ReplaceAll(codeGO, "{{DummyDeclarations}}", dummyDeclarations.String())
	codeGO = strings.ReplaceAll(codeGO, "{{MainPkgImportAlias}}", mainPkgImportAlias)
	codeGO = strings.ReplaceAll(codeGO, "{{MarshallBody}}", marshallBody.String())
	codeGO = strings.ReplaceAll(codeGO, "{{DefineCases}}", defineCases.String())
	codeGO = strings.ReplaceAll(codeGO, "{{AssignCases}}", assignCases.String())

	outFile := filepath.Join(pkgPath, string(models.GeneratedGongStageSetGoFilePath))
	file, err := os.Create(outFile)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	fmt.Fprint(file, codeGO)
}
