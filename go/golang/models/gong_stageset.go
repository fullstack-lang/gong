package models

import (
	"fmt"
	"go/types"
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
	IsGongEnum        bool
	EnumTypeQual      string
	EnumPkgAlias      string
	IsDuration        bool
	IsTime            bool
	IsAny             bool
}

func ExtractStructFields(mPkg *models.ModelPkg, structName string, pkgPathToField ...map[string]*models.StageSetField) []StageSetStructField {
	var p2f map[string]*models.StageSetField
	if len(pkgPathToField) > 0 {
		p2f = pkgPathToField[0]
	}
	var result []StageSetStructField
	if mPkg.TypesPkg != nil {
		if obj := mPkg.TypesPkg.Scope().Lookup(structName); obj != nil {
			if named, ok := obj.Type().(*types.Named); ok {
				if st, ok := named.Underlying().(*types.Struct); ok {
					var inspectFields func(st *types.Struct)
					inspectFields = func(st *types.Struct) {
						for i := 0; i < st.NumFields(); i++ {
							fld := st.Field(i)
							if fld.Anonymous() {
								embType := fld.Type()
								if ptr, ok := embType.(*types.Pointer); ok {
									embType = ptr.Elem()
								}
								if embNamed, ok := embType.(*types.Named); ok {
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
										targetName := targetNamed.Obj().Name()
										if targetPkg.Path() == mPkg.PkgPath {
											if _, ok := mPkg.GongStructs[mPkg.PkgPath+"."+targetName]; !ok {
												continue
											}
										} else {
											if p2f == nil || p2f[targetPkg.Path()] == nil {
												continue
											}
										}
										result = append(result, StageSetStructField{
											Name:              fld.Name(),
											IsPointer:         true,
											TargetPackagePath: targetPkg.Path(),
											TargetStructName:  targetName,
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
											targetName := targetNamed.Obj().Name()
											if targetPkg.Path() == mPkg.PkgPath {
												if _, ok := mPkg.GongStructs[mPkg.PkgPath+"."+targetName]; !ok {
													continue
												}
											} else {
												if p2f == nil || p2f[targetPkg.Path()] == nil {
													continue
												}
											}
											result = append(result, StageSetStructField{
												Name:              fld.Name(),
												IsSliceOfPointer:  true,
												TargetPackagePath: targetPkg.Path(),
												TargetStructName:  targetName,
											})
											continue
										}
									}
								}
							}
							// check empty interface (any)
							if iface, ok := fld.Type().Underlying().(*types.Interface); ok && iface.Empty() {
								result = append(result, StageSetStructField{
									Name:  fld.Name(),
									IsAny: true,
								})
								continue
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
									enumPkgAlias := mPkg.PkgGoName
									if enumPkgAlias == "" {
										enumPkgAlias = "models"
									}
									pkgPath := mPkg.PkgPath
									if namedType.Obj().Pkg() != nil {
										pkgPath = namedType.Obj().Pkg().Path()
										if p2f != nil && p2f[pkgPath] != nil {
											enumPkgAlias = p2f[pkgPath].ImportAlias
										} else {
											enumPkgAlias = namedType.Obj().Pkg().Name()
										}
										if pkgPath != mPkg.PkgPath {
											qual = enumPkgAlias + "." + qual
										}
									} else if p2f != nil && p2f[mPkg.PkgPath] != nil {
										enumPkgAlias = p2f[mPkg.PkgPath].ImportAlias
									}
									_, isGongEnum := mPkg.GongEnums[pkgPath+"."+namedType.Obj().Name()]
									if !isGongEnum {
										_, isGongEnum = mPkg.GongEnums[mPkg.PkgPath+"."+namedType.Obj().Name()]
									}
									if !isGongEnum {
										if types.NewMethodSet(types.NewPointer(namedType)).Lookup(nil, "FromCodeString") != nil {
											isGongEnum = true
										}
									}
									result = append(result, StageSetStructField{
										Name:         fld.Name(),
										IsEnum:       true,
										IsGongEnum:   isGongEnum,
										BasicKind:    basic.Kind(),
										EnumTypeQual: qual,
										EnumPkgAlias: enumPkgAlias,
									})
									continue
								}
							}
						}
					}

					inspectFields(st)
				}
			}
		}
	}

	if len(result) == 0 {
		if gs, ok := mPkg.GongStructs[mPkg.PkgPath+"."+structName]; ok {
			enumPkgAlias := mPkg.PkgGoName
			if enumPkgAlias == "" {
				enumPkgAlias = "models"
			}
			if p2f != nil && p2f[mPkg.PkgPath] != nil {
				enumPkgAlias = p2f[mPkg.PkgPath].ImportAlias
			}
			for _, field := range gs.Fields {
				switch f := field.(type) {
				case *models.GongBasicField:
					if f.GetBasicKind() == types.UntypedNil {
						result = append(result, StageSetStructField{
							Name:  f.Name,
							IsAny: true,
						})
					} else if f.DeclaredType == "time.Duration" {
						result = append(result, StageSetStructField{
							Name:       f.Name,
							IsDuration: true,
						})
					} else if f.GongEnum != nil {
						result = append(result, StageSetStructField{
							Name:         f.Name,
							IsEnum:       true,
							IsGongEnum:   true,
							BasicKind:    f.GetBasicKind(),
							EnumTypeQual: f.GongEnum.Name,
							EnumPkgAlias: enumPkgAlias,
						})
					} else {
						result = append(result, StageSetStructField{
							Name:      f.Name,
							BasicKind: f.GetBasicKind(),
						})
					}
				case *models.GongTimeField:
					result = append(result, StageSetStructField{
						Name:   f.Name,
						IsTime: true,
					})
				case *models.PointerToGongStructField:
					if f.GongStruct != nil {
						result = append(result, StageSetStructField{
							Name:              f.Name,
							IsPointer:         true,
							TargetPackagePath: mPkg.PkgPath,
							TargetStructName:  f.GongStruct.Name,
						})
					}
				case *models.SliceOfPointerToGongStructField:
					if f.GongStruct != nil {
						result = append(result, StageSetStructField{
							Name:              f.Name,
							IsSliceOfPointer:  true,
							TargetPackagePath: mPkg.PkgPath,
							TargetStructName:  f.GongStruct.Name,
						})
					}
				}
			}
		}
	}

	return result
}

type ResolvedSubStages struct {
	ExternalImports                      string
	SubStageFields                       string
	NewSubStageStatements                string
	CommitStatements                     string
	CheckoutStatements                   string
	ResetStatements                      string
	CleanStatements                      string
	ComputeReverseMapsStatements         string
	ComputeInstancesNbStatements         string
	ComputeReferenceAndOrdersStatements  string
	SyntheticImports                     string
	DummyDeclarations                    string
	MainPkgImportAlias                   string
	ProbePackageName                     string
	MarshallBody                         string
	ImportPathCases                      string
	PrepopulateIdentifierMapStatements   string
	DefineCases                          string
	AssignCases                          string
	UnstageCases                         string
}

func ResolveSubStages(
	modelPkg *models.ModelPkg,
	pkgPath string,
) *ResolvedSubStages {
	stageSet := modelPkg.StageSet
	rootPkgGoName := modelPkg.PkgGoName
	if rootPkgGoName == "" {
		rootPkgGoName = "models"
	}
	if stageSet == nil || len(stageSet.Fields) == 0 {
		stageSet = &models.StageSetModel{
			Name: "StageSet",
			Fields: []*models.StageSetField{
				{
					Name:        "Stage",
					PackageName: rootPkgGoName,
					PackagePath: modelPkg.PkgPath,
					IsLocal:     true,
					ImportAlias: rootPkgGoName,
				},
			},
		}
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
	var subStageFields strings.Builder
	var newSubStageStatements strings.Builder
	for _, f := range stageSet.Fields {
		if !f.IsLocal {
			if f.ImportAlias != f.PackageName {
				externalImports.WriteString(fmt.Sprintf("\n\t%s \"%s\"", f.ImportAlias, f.PackagePath))
			} else {
				externalImports.WriteString(fmt.Sprintf("\n\t\"%s\"", f.PackagePath))
			}
			subStageFields.WriteString(fmt.Sprintf("\t%s *%s.Stage\n", f.Name, f.ImportAlias))
			newSubStageStatements.WriteString(fmt.Sprintf("\tsubPath_%s := \"%s\"\n\tif name != \"\" {\n\t\tsubPath_%s = name + \"_%s\"\n\t}\n\tstage.%s = %s.NewStage(subPath_%s)\n", f.Name, f.ImportAlias, f.Name, f.ImportAlias, f.Name, f.ImportAlias, f.Name))
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
		if !f.IsLocal {
			commitStatements.WriteString(fmt.Sprintf("\tif stage.%s != nil {\n\t\tstage.%s.Commit()\n\t}\n", f.Name, f.Name))
			checkoutStatements.WriteString(fmt.Sprintf("\tif stage.%s != nil {\n\t\tstage.%s.Checkout()\n\t}\n", f.Name, f.Name))
			resetStatements.WriteString(fmt.Sprintf("\tif stage.%s != nil {\n\t\tstage.%s.Reset()\n\t}\n", f.Name, f.Name))
			cleanStatements.WriteString(fmt.Sprintf("\tif stage.%s != nil {\n\t\tmodified = stage.%s.Clean() || modified\n\t}\n", f.Name, f.Name))
			computeReverseMapsStatements.WriteString(fmt.Sprintf("\tif stage.%s != nil {\n\t\tstage.%s.ComputeReverseMaps()\n\t}\n", f.Name, f.Name))
			computeInstancesNbStatements.WriteString(fmt.Sprintf("\tif stage.%s != nil {\n\t\tstage.%s.ComputeInstancesNb()\n\t}\n", f.Name, f.Name))
			computeReferenceAndOrdersStatements.WriteString(fmt.Sprintf("\tif stage.%s != nil {\n\t\tstage.%s.ComputeReferenceAndOrders()\n\t}\n", f.Name, f.Name))
		}
	}

	var syntheticImports strings.Builder
	var dummyDeclarations strings.Builder
	for _, f := range stageSet.Fields {
		if f.ImportAlias != f.PackageName {
			syntheticImports.WriteString(fmt.Sprintf("\n\t%s \"%s\"", f.ImportAlias, f.PackagePath))
		} else {
			syntheticImports.WriteString(fmt.Sprintf("\n\t\"%s\"", f.PackagePath))
		}
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
	if mainPkgImportAlias == "" {
		mainPkgImportAlias = rootPkgGoName
	}

	var mainPkgPath string
	for _, f := range stageSet.Fields {
		if f.IsLocal {
			mainPkgPath = f.PackagePath
			break
		}
	}
	if mainPkgPath == "" && len(stageSet.Fields) > 0 {
		mainPkgPath = stageSet.Fields[0].PackagePath
	}
	if mainPkgPath == "" {
		mainPkgPath = modelPkg.PkgPath
	}
	probePackageName := mainPkgPath + "/probe/cmd/probe"

	// Build code inside MarshallToString
	var marshallBody strings.Builder
	marshallBody.WriteString("\tvar declarations strings.Builder\n")
	marshallBody.WriteString("\tvar values strings.Builder\n")
	marshallBody.WriteString("\tvar pointers strings.Builder\n")
	marshallBody.WriteString("\tvar lastStageDecl string\n")
	marshallBody.WriteString("\tvar lastStageVal string\n")
	marshallBody.WriteString("\tvar lastStagePtr string\n")
	marshallBody.WriteString("\t_ = lastStageDecl\n")
	marshallBody.WriteString("\t_ = lastStageVal\n")
	marshallBody.WriteString("\t_ = lastStagePtr\n")
	marshallBody.WriteString(fmt.Sprintf("\tlocalAliasPrefix := \"__%s\"\n", mainPkgImportAlias))
	marshallBody.WriteString("\tif stage.IsInDeltaMode() || stage.GetGongMarshallingMode() == GongMarshallingAppendCommit {\n")
	marshallBody.WriteString("\t\tlocalAliasPrefix = \"\"\n")
	marshallBody.WriteString("\t}\n")
	marshallBody.WriteString("\t_ = localAliasPrefix\n\n")

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

		stageExpr := "stage"
		aliasPrefixExpr := "localAliasPrefix"
		if !f.IsLocal {
			stageExpr = "stage." + f.Name
			aliasPrefixExpr = fmt.Sprintf("\"__%s\"", f.ImportAlias)
		}

		for _, sName := range structNames {
			sVar := strings.ToLower(sName)
			sPlural := sName + "s"

			fields := ExtractStructFields(mPkg, sName, pkgPathToField)

			marshallBody.WriteString(fmt.Sprintf("\tif %s != nil {\n", stageExpr))
			marshallBody.WriteString(fmt.Sprintf("\t\tfor _, %s := range __gong__sortStageSetInstances(%s.%s, %s.%s_stagedOrder) {\n", sVar, stageExpr, sPlural, stageExpr, sName))
			marshallBody.WriteString(fmt.Sprintf("\t\t\tif lastStageDecl != %q {\n", f.Name))
			marshallBody.WriteString("\t\t\t\tif declarations.Len() > 0 {\n")
			marshallBody.WriteString("\t\t\t\t\tdeclarations.WriteString(\"\\n\")\n")
			marshallBody.WriteString("\t\t\t\t}\n")
			marshallBody.WriteString(fmt.Sprintf("\t\t\t\tlastStageDecl = %q\n", f.Name))
			marshallBody.WriteString("\t\t\t}\n")
			marshallBody.WriteString(fmt.Sprintf("\t\t\t%sIdent := %s + %s.GongGetIdentifier(%s)\n", sVar, aliasPrefixExpr, sVar, stageExpr))
			// Phase 1 declaration
			marshallBody.WriteString(fmt.Sprintf("\t\t\tdeclarations.WriteString(fmt.Sprintf(\"\\n\\t%%s := (&%s.%s{Name: %%s}).Stage(%s)\", %sIdent, __gong__toRawStringLiteral(%s.Name)))\n", f.ImportAlias, sName, stageExpr, sVar, sVar))

			// Phase 2 values
			marshallBody.WriteString(fmt.Sprintf("\t\t\tif lastStageVal != %q {\n", f.Name))
			marshallBody.WriteString("\t\t\t\tif values.Len() > 0 {\n")
			marshallBody.WriteString("\t\t\t\t\tvalues.WriteString(\"\\n\")\n")
			marshallBody.WriteString("\t\t\t\t}\n")
			marshallBody.WriteString(fmt.Sprintf("\t\t\t\tlastStageVal = %q\n", f.Name))
			marshallBody.WriteString("\t\t\t}\n")
			for _, fld := range fields {
				if fld.IsPointer || fld.IsSliceOfPointer {
					continue
				}
				if fld.IsAny {
					marshallBody.WriteString(fmt.Sprintf("\t\t\tif str, ok := %s.%s.(string); ok {\n\t\t\t\tvalues.WriteString(fmt.Sprintf(\"\\n\\t%%s.%s = %%s\", %sIdent, str))\n\t\t\t}\n", sVar, fld.Name, fld.Name, sVar))
					continue
				}
				if fld.IsGongEnum {
					switch fld.BasicKind {
					case types.Int, types.Int64, types.Int32, types.Int16, types.Int8, types.Uint, types.Uint64, types.Uint32, types.Uint16, types.Uint8:
						marshallBody.WriteString(fmt.Sprintf("\t\t\tif %s.%s.ToCodeString() != \"\" {\n\t\t\t\tvalues.WriteString(fmt.Sprintf(\"\\n\\t%%s.%s = %s.%%s\", %sIdent, %s.%s.ToCodeString()))\n\t\t\t} else {\n\t\t\t\tvalues.WriteString(fmt.Sprintf(\"\\n\\t%%s.%s = 0\", %sIdent))\n\t\t\t}\n", sVar, fld.Name, fld.Name, fld.EnumPkgAlias, sVar, sVar, fld.Name, fld.Name, sVar))
					default:
						marshallBody.WriteString(fmt.Sprintf("\t\t\tif %s.%s.ToCodeString() != \"\" {\n\t\t\t\tvalues.WriteString(fmt.Sprintf(\"\\n\\t%%s.%s = %s.%%s\", %sIdent, %s.%s.ToCodeString()))\n\t\t\t} else {\n\t\t\t\tvalues.WriteString(fmt.Sprintf(\"\\n\\t%%s.%s = \\\"\\\"\", %sIdent))\n\t\t\t}\n", sVar, fld.Name, fld.Name, fld.EnumPkgAlias, sVar, sVar, fld.Name, fld.Name, sVar))
					}
					continue
				}
				switch fld.BasicKind {
				case types.String:
					if fld.IsEnum {
						marshallBody.WriteString(fmt.Sprintf("\t\t\tvalues.WriteString(fmt.Sprintf(\"\\n\\t%%s.%s = %%s\", %sIdent, __gong__toRawStringLiteral(string(%s.%s))))\n", fld.Name, sVar, sVar, fld.Name))
					} else {
						marshallBody.WriteString(fmt.Sprintf("\t\t\tvalues.WriteString(fmt.Sprintf(\"\\n\\t%%s.%s = %%s\", %sIdent, __gong__toRawStringLiteral(%s.%s)))\n", fld.Name, sVar, sVar, fld.Name))
					}
				case types.Int, types.Int64, types.Int32, types.Int16, types.Int8, types.Uint, types.Uint64, types.Uint32, types.Uint16, types.Uint8:
					if fld.IsEnum {
						marshallBody.WriteString(fmt.Sprintf("\t\t\tvalues.WriteString(fmt.Sprintf(\"\\n\\t%%s.%s = %%d\", %sIdent, int(%s.%s)))\n", fld.Name, sVar, sVar, fld.Name))
					} else {
						marshallBody.WriteString(fmt.Sprintf("\t\t\tvalues.WriteString(fmt.Sprintf(\"\\n\\t%%s.%s = %%d\", %sIdent, %s.%s))\n", fld.Name, sVar, sVar, fld.Name))
					}
				case types.Float64, types.Float32:
					if fld.IsEnum {
						marshallBody.WriteString(fmt.Sprintf("\t\t\tvalues.WriteString(fmt.Sprintf(\"\\n\\t%%s.%s = %%f\", %sIdent, float64(%s.%s)))\n", fld.Name, sVar, sVar, fld.Name))
					} else {
						marshallBody.WriteString(fmt.Sprintf("\t\t\tvalues.WriteString(fmt.Sprintf(\"\\n\\t%%s.%s = %%f\", %sIdent, %s.%s))\n", fld.Name, sVar, sVar, fld.Name))
					}
				case types.Bool:
					if fld.IsEnum {
						marshallBody.WriteString(fmt.Sprintf("\t\t\tvalues.WriteString(fmt.Sprintf(\"\\n\\t%%s.%s = %%t\", %sIdent, bool(%s.%s)))\n", fld.Name, sVar, sVar, fld.Name))
					} else {
						marshallBody.WriteString(fmt.Sprintf("\t\t\tvalues.WriteString(fmt.Sprintf(\"\\n\\t%%s.%s = %%t\", %sIdent, %s.%s))\n", fld.Name, sVar, sVar, fld.Name))
					}
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
						if targetPkg := fieldToModelPkg[targetSSF]; targetPkg != nil {
							targetGS, isGS := targetPkg.GongStructs[targetPkg.PkgPath+"."+fld.TargetStructName]
							if !isGS || targetGS.IsOmittedForMarshalling {
								continue
							}
						}
						targetStageExpr := "stage"
						targetAliasPrefixExpr := "localAliasPrefix"
						if !targetSSF.IsLocal {
							targetStageExpr = "stage." + targetSSF.Name
							targetAliasPrefixExpr = fmt.Sprintf("\"__%s\"", targetSSF.ImportAlias)
						}
						marshallBody.WriteString(fmt.Sprintf("\t\t\tif %s.%s != nil {\n", sVar, fld.Name))
						marshallBody.WriteString(fmt.Sprintf("\t\t\t\tif lastStagePtr != %q {\n", f.Name))
						marshallBody.WriteString("\t\t\t\t\tif pointers.Len() > 0 {\n")
						marshallBody.WriteString("\t\t\t\t\t\tpointers.WriteString(\"\\n\")\n")
						marshallBody.WriteString("\t\t\t\t\t}\n")
						marshallBody.WriteString(fmt.Sprintf("\t\t\t\t\tlastStagePtr = %q\n", f.Name))
						marshallBody.WriteString("\t\t\t\t}\n")
						marshallBody.WriteString(fmt.Sprintf("\t\t\t\ttargetIdent := %s + %s.%s.GongGetIdentifier(%s)\n", targetAliasPrefixExpr, sVar, fld.Name, targetStageExpr))
						marshallBody.WriteString(fmt.Sprintf("\t\t\t\tpointers.WriteString(fmt.Sprintf(\"\\n\\t%%s.%s = %%s\", %sIdent, targetIdent))\n", fld.Name, sVar))
						marshallBody.WriteString("\t\t\t}\n")
					}
				} else if fld.IsSliceOfPointer {
					targetSSF := pkgPathToField[fld.TargetPackagePath]
					if targetSSF != nil {
						if targetPkg := fieldToModelPkg[targetSSF]; targetPkg != nil {
							targetGS, isGS := targetPkg.GongStructs[targetPkg.PkgPath+"."+fld.TargetStructName]
							if !isGS || targetGS.IsOmittedForMarshalling {
								continue
							}
						}
						targetStageExpr := "stage"
						targetAliasPrefixExpr := "localAliasPrefix"
						if !targetSSF.IsLocal {
							targetStageExpr = "stage." + targetSSF.Name
							targetAliasPrefixExpr = fmt.Sprintf("\"__%s\"", targetSSF.ImportAlias)
						}
						marshallBody.WriteString(fmt.Sprintf("\t\t\tfor _, elem := range %s.%s {\n", sVar, fld.Name))
						marshallBody.WriteString(fmt.Sprintf("\t\t\t\tif lastStagePtr != %q {\n", f.Name))
						marshallBody.WriteString("\t\t\t\t\tif pointers.Len() > 0 {\n")
						marshallBody.WriteString("\t\t\t\t\t\tpointers.WriteString(\"\\n\")\n")
						marshallBody.WriteString("\t\t\t\t\t}\n")
						marshallBody.WriteString(fmt.Sprintf("\t\t\t\t\tlastStagePtr = %q\n", f.Name))
						marshallBody.WriteString("\t\t\t\t}\n")
						marshallBody.WriteString(fmt.Sprintf("\t\t\t\ttargetIdent := %s + elem.GongGetIdentifier(%s)\n", targetAliasPrefixExpr, targetStageExpr))
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
	var prepopulateIdentifierMap strings.Builder
	var defineCases strings.Builder
	var assignCases strings.Builder
	var unstageCases strings.Builder

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

		stageExpr := "stage"
		if !f.IsLocal {
			stageExpr = "stage." + f.Name
		}

		if len(structNames) > 0 {
			prepopulateIdentifierMap.WriteString(fmt.Sprintf("\tif %s != nil {\n", stageExpr))
			prepopulateIdentifierMap.WriteString(fmt.Sprintf("\t\tfor _, instance := range %s.GetInstances() {\n", stageExpr))
			prepopulateIdentifierMap.WriteString(fmt.Sprintf("\t\t\tidentifierMap[instance.GongGetIdentifier(%s)] = instance\n", stageExpr))
			prepopulateIdentifierMap.WriteString(fmt.Sprintf("\t\t\tidentifierMap[\"__%s\"+instance.GongGetIdentifier(%s)] = instance\n", f.ImportAlias, stageExpr))
			prepopulateIdentifierMap.WriteString("\t\t}\n")
			prepopulateIdentifierMap.WriteString("\t}\n")

			if f.IsLocal {
				unstageCases.WriteString("\t\t\t\t\t\t\t\tif inst, ok := instance.(GongstructIF); ok {\n")
				unstageCases.WriteString("\t\t\t\t\t\t\t\t\tinst.UnstageVoid(stage)\n")
				unstageCases.WriteString("\t\t\t\t\t\t\t\t}\n")
			} else {
				unstageCases.WriteString(fmt.Sprintf("\t\t\t\t\t\t\t\tif inst, ok := instance.(%s.GongstructIF); ok && stage.%s != nil {\n", f.ImportAlias, f.Name))
				unstageCases.WriteString(fmt.Sprintf("\t\t\t\t\t\t\t\t\tinst.UnstageVoid(stage.%s)\n", f.Name))
				unstageCases.WriteString("\t\t\t\t\t\t\t\t}\n")
			}
		}

		defineCases.WriteString(fmt.Sprintf("\t\t\tcase \"%s\":\n", f.ImportAlias))
		defineCases.WriteString("\t\t\t\tswitch typeName {\n")
		for _, sName := range structNames {
			typeQual := sName
			if !f.IsLocal {
				typeQual = f.ImportAlias + "." + sName
			}
			defineCases.WriteString(fmt.Sprintf("\t\t\t\tcase \"%s\":\n", sName))
			defineCases.WriteString(fmt.Sprintf("\t\t\t\t\tidentifierMap[ident.Name] = __gong__stageSetInit(new(%s), %s, ident.Name, instanceName, preserveOrder)\n", typeQual, stageExpr))
		}
		defineCases.WriteString("\t\t\t\t}\n")

		for _, sName := range structNames {
			typeQual := sName
			if !f.IsLocal {
				typeQual = f.ImportAlias + "." + sName
			}
			fields := ExtractStructFields(mPkg, sName, pkgPathToField)

			assignCases.WriteString(fmt.Sprintf("\t\t\t\tcase *%s:\n", typeQual))
			assignCases.WriteString("\t\t\t\t\tswitch fieldName {\n")
			for _, fld := range fields {
				if fld.IsPointer {
					targetSSF := pkgPathToField[fld.TargetPackagePath]
					if targetSSF != nil {
						if targetPkg := fieldToModelPkg[targetSSF]; targetPkg != nil {
							targetGS, isGS := targetPkg.GongStructs[targetPkg.PkgPath+"."+fld.TargetStructName]
							if !isGS || targetGS.IsOmittedForMarshalling {
								continue
							}
						}
						assignCases.WriteString(fmt.Sprintf("\t\t\t\t\tcase \"%s\":\n\t\t\t\t\t\t__gong__assignPointer(&inst.%s, rhs, identifierMap)\n", fld.Name, fld.Name))
					}
				} else if fld.IsSliceOfPointer {
					targetSSF := pkgPathToField[fld.TargetPackagePath]
					if targetSSF != nil {
						if targetPkg := fieldToModelPkg[targetSSF]; targetPkg != nil {
							targetGS, isGS := targetPkg.GongStructs[targetPkg.PkgPath+"."+fld.TargetStructName]
							if !isGS || targetGS.IsOmittedForMarshalling {
								continue
							}
						}
						assignCases.WriteString(fmt.Sprintf("\t\t\t\t\tcase \"%s\":\n\t\t\t\t\t\t__gong__assignSliceOfPointers(&inst.%s, rhs, identifierMap)\n", fld.Name, fld.Name))
					}
				} else if fld.IsAny {
					assignCases.WriteString(fmt.Sprintf("\t\t\t\t\tcase \"%s\":\n\t\t\t\t\t\tinst.%s = GongExtractExpr(rhs)\n", fld.Name, fld.Name))
				} else if fld.IsGongEnum {
					assignCases.WriteString(fmt.Sprintf("\t\t\t\t\tcase \"%s\":\n\t\t\t\t\t\tGongUnmarshallEnum(&inst.%s, rhs)\n", fld.Name, fld.Name))
				} else {
					switch fld.BasicKind {
					case types.String:
						if fld.IsEnum {
							assignCases.WriteString(fmt.Sprintf("\t\t\t\t\tcase \"%s\":\n\t\t\t\t\t\tinst.%s = %s(GongExtractString(rhs))\n", fld.Name, fld.Name, fld.EnumTypeQual))
						} else {
							assignCases.WriteString(fmt.Sprintf("\t\t\t\t\tcase \"%s\":\n\t\t\t\t\t\tinst.%s = GongExtractString(rhs)\n", fld.Name, fld.Name))
						}
					case types.Int:
						if fld.IsEnum {
							assignCases.WriteString(fmt.Sprintf("\t\t\t\t\tcase \"%s\":\n\t\t\t\t\t\tinst.%s = %s(GongExtractInt(rhs))\n", fld.Name, fld.Name, fld.EnumTypeQual))
						} else {
							assignCases.WriteString(fmt.Sprintf("\t\t\t\t\tcase \"%s\":\n\t\t\t\t\t\tinst.%s = GongExtractInt(rhs)\n", fld.Name, fld.Name))
						}
					case types.Int8:
						if fld.IsEnum {
							assignCases.WriteString(fmt.Sprintf("\t\t\t\t\tcase \"%s\":\n\t\t\t\t\t\tinst.%s = %s(GongExtractInt(rhs))\n", fld.Name, fld.Name, fld.EnumTypeQual))
						} else {
							assignCases.WriteString(fmt.Sprintf("\t\t\t\t\tcase \"%s\":\n\t\t\t\t\t\tinst.%s = int8(GongExtractInt(rhs))\n", fld.Name, fld.Name))
						}
					case types.Int16:
						if fld.IsEnum {
							assignCases.WriteString(fmt.Sprintf("\t\t\t\t\tcase \"%s\":\n\t\t\t\t\t\tinst.%s = %s(GongExtractInt(rhs))\n", fld.Name, fld.Name, fld.EnumTypeQual))
						} else {
							assignCases.WriteString(fmt.Sprintf("\t\t\t\t\tcase \"%s\":\n\t\t\t\t\t\tinst.%s = int16(GongExtractInt(rhs))\n", fld.Name, fld.Name))
						}
					case types.Int32:
						if fld.IsEnum {
							assignCases.WriteString(fmt.Sprintf("\t\t\t\t\tcase \"%s\":\n\t\t\t\t\t\tinst.%s = %s(GongExtractInt(rhs))\n", fld.Name, fld.Name, fld.EnumTypeQual))
						} else {
							assignCases.WriteString(fmt.Sprintf("\t\t\t\t\tcase \"%s\":\n\t\t\t\t\t\tinst.%s = int32(GongExtractInt(rhs))\n", fld.Name, fld.Name))
						}
					case types.Int64:
						if fld.IsEnum {
							assignCases.WriteString(fmt.Sprintf("\t\t\t\t\tcase \"%s\":\n\t\t\t\t\t\tinst.%s = %s(GongExtractInt(rhs))\n", fld.Name, fld.Name, fld.EnumTypeQual))
						} else {
							assignCases.WriteString(fmt.Sprintf("\t\t\t\t\tcase \"%s\":\n\t\t\t\t\t\tinst.%s = int64(GongExtractInt(rhs))\n", fld.Name, fld.Name))
						}
					case types.Uint:
						if fld.IsEnum {
							assignCases.WriteString(fmt.Sprintf("\t\t\t\t\tcase \"%s\":\n\t\t\t\t\t\tinst.%s = %s(GongExtractInt(rhs))\n", fld.Name, fld.Name, fld.EnumTypeQual))
						} else {
							assignCases.WriteString(fmt.Sprintf("\t\t\t\t\tcase \"%s\":\n\t\t\t\t\t\tinst.%s = uint(GongExtractInt(rhs))\n", fld.Name, fld.Name))
						}
					case types.Uint8:
						if fld.IsEnum {
							assignCases.WriteString(fmt.Sprintf("\t\t\t\t\tcase \"%s\":\n\t\t\t\t\t\tinst.%s = %s(GongExtractInt(rhs))\n", fld.Name, fld.Name, fld.EnumTypeQual))
						} else {
							assignCases.WriteString(fmt.Sprintf("\t\t\t\t\tcase \"%s\":\n\t\t\t\t\t\tinst.%s = uint8(GongExtractInt(rhs))\n", fld.Name, fld.Name))
						}
					case types.Uint16:
						if fld.IsEnum {
							assignCases.WriteString(fmt.Sprintf("\t\t\t\t\tcase \"%s\":\n\t\t\t\t\t\tinst.%s = %s(GongExtractInt(rhs))\n", fld.Name, fld.Name, fld.EnumTypeQual))
						} else {
							assignCases.WriteString(fmt.Sprintf("\t\t\t\t\tcase \"%s\":\n\t\t\t\t\t\tinst.%s = uint16(GongExtractInt(rhs))\n", fld.Name, fld.Name))
						}
					case types.Uint32:
						if fld.IsEnum {
							assignCases.WriteString(fmt.Sprintf("\t\t\t\t\tcase \"%s\":\n\t\t\t\t\t\tinst.%s = %s(GongExtractInt(rhs))\n", fld.Name, fld.Name, fld.EnumTypeQual))
						} else {
							assignCases.WriteString(fmt.Sprintf("\t\t\t\t\tcase \"%s\":\n\t\t\t\t\t\tinst.%s = uint32(GongExtractInt(rhs))\n", fld.Name, fld.Name))
						}
					case types.Uint64:
						if fld.IsEnum {
							assignCases.WriteString(fmt.Sprintf("\t\t\t\t\tcase \"%s\":\n\t\t\t\t\t\tinst.%s = %s(GongExtractInt(rhs))\n", fld.Name, fld.Name, fld.EnumTypeQual))
						} else {
							assignCases.WriteString(fmt.Sprintf("\t\t\t\t\tcase \"%s\":\n\t\t\t\t\t\tinst.%s = uint64(GongExtractInt(rhs))\n", fld.Name, fld.Name))
						}
					case types.Float32:
						if fld.IsEnum {
							assignCases.WriteString(fmt.Sprintf("\t\t\t\t\tcase \"%s\":\n\t\t\t\t\t\tinst.%s = %s(GongExtractFloat(rhs))\n", fld.Name, fld.Name, fld.EnumTypeQual))
						} else {
							assignCases.WriteString(fmt.Sprintf("\t\t\t\t\tcase \"%s\":\n\t\t\t\t\t\tinst.%s = float32(GongExtractFloat(rhs))\n", fld.Name, fld.Name))
						}
					case types.Float64:
						if fld.IsEnum {
							assignCases.WriteString(fmt.Sprintf("\t\t\t\t\tcase \"%s\":\n\t\t\t\t\t\tinst.%s = %s(GongExtractFloat(rhs))\n", fld.Name, fld.Name, fld.EnumTypeQual))
						} else {
							assignCases.WriteString(fmt.Sprintf("\t\t\t\t\tcase \"%s\":\n\t\t\t\t\t\tinst.%s = GongExtractFloat(rhs)\n", fld.Name, fld.Name))
						}
					case types.Bool:
						if fld.IsEnum {
							assignCases.WriteString(fmt.Sprintf("\t\t\t\t\tcase \"%s\":\n\t\t\t\t\t\tinst.%s = %s(GongExtractBool(rhs))\n", fld.Name, fld.Name, fld.EnumTypeQual))
						} else {
							assignCases.WriteString(fmt.Sprintf("\t\t\t\t\tcase \"%s\":\n\t\t\t\t\t\tinst.%s = GongExtractBool(rhs)\n", fld.Name, fld.Name))
						}
					default:
						if fld.IsTime {
							assignCases.WriteString(fmt.Sprintf("\t\t\t\t\tcase \"%s\":\n\t\t\t\t\t\tinst.%s = GongExtractDate(rhs)\n", fld.Name, fld.Name))
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

	var importPathCases strings.Builder
	for _, f := range orderedFields {
		if f.PackagePath != "" {
			importPathCases.WriteString(fmt.Sprintf("\t\tcase \"%s\":\n\t\t\taliasToCanonical[alias] = \"%s\"\n", f.PackagePath, f.ImportAlias))
		}
	}

	return &ResolvedSubStages{
		ExternalImports:                     externalImports.String(),
		SubStageFields:                      subStageFields.String(),
		NewSubStageStatements:               newSubStageStatements.String(),
		CommitStatements:                    commitStatements.String(),
		CheckoutStatements:                  checkoutStatements.String(),
		ResetStatements:                     resetStatements.String(),
		CleanStatements:                     cleanStatements.String(),
		ComputeReverseMapsStatements:        computeReverseMapsStatements.String(),
		ComputeInstancesNbStatements:        computeInstancesNbStatements.String(),
		ComputeReferenceAndOrdersStatements: computeReferenceAndOrdersStatements.String(),
		SyntheticImports:                    syntheticImports.String(),
		DummyDeclarations:                   dummyDeclarations.String(),
		MainPkgImportAlias:                  mainPkgImportAlias,
		ProbePackageName:                    probePackageName,
		MarshallBody:                        marshallBody.String(),
		ImportPathCases:                     importPathCases.String(),
		PrepopulateIdentifierMapStatements:  prepopulateIdentifierMap.String(),
		DefineCases:                         defineCases.String(),
		AssignCases:                         assignCases.String(),
		UnstageCases:                        unstageCases.String(),
	}
}

