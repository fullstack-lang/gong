package probe

import (
	"fmt"
	"go/types"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	golang_models "github.com/fullstack-lang/gong/go/golang/models"
	"github.com/fullstack-lang/gong/go/models"
)

type structInfo struct {
	pkgField   *models.StageSetField
	pkgName    string
	pkgPath    string
	structName string
	typeQual   string
	mPkg       *models.ModelPkg
	fields     []golang_models.StageSetStructField
}

func CodeGeneratorStageSetProbe(
	modelPkg *models.ModelPkg,
	pkgName string,
	pkgPath string,
	pkgPathRoot string,
	useSplitlite bool,
) {
	stageSet := modelPkg.StageSet
	if stageSet == nil || len(stageSet.Fields) == 0 {
		return
	}

	if before, _, ok := strings.Cut(pkgPathRoot, "/go/models"); ok {
		pkgPathRoot = before + "/go"
	} else {
		pkgPathRoot = strings.ReplaceAll(pkgPathRoot, "/models", "")
	}

	goModDir, modPath := models.FindGoMod(pkgPath)
	depPkgPaths, _ := models.DiscoverModelDependencies(pkgPath)
	_ = depPkgPaths

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

	// Collect all structs across all packages in StageSet
	var allStructs []structInfo
	for _, f := range stageSet.Fields {
		mPkg := fieldToModelPkg[f]
		if mPkg == nil {
			continue
		}
		var sNames []string
		for sName, gs := range mPkg.GongStructs {
			if gs.ModelPkg.PkgPath == mPkg.PkgPath && gs.HasNameField() && !gs.IsOmittedForMarshalling {
				sNames = append(sNames, sName[len(mPkg.PkgPath)+1:])
			}
		}
		sort.Strings(sNames)
		for _, sName := range sNames {
			typeQual := sName
			if !f.IsLocal {
				typeQual = f.PackageName + "." + sName
			} else {
				typeQual = "models." + sName
			}
			flds := golang_models.ExtractStructFields(mPkg, sName)
			allStructs = append(allStructs, structInfo{
				pkgField:   f,
				pkgName:    f.PackageName,
				pkgPath:    f.PackagePath,
				structName: sName,
				typeQual:   typeQual,
				mPkg:       mPkg,
				fields:     flds,
			})
		}
	}

	// Generate probe/stageset_form_helpers.go
	generateStageSetFormHelpers(pkgPath)

	// Generate probe/stageset_probe.go
	generateStageSetProbeMain(pkgPath, pkgPathRoot, modelPkg, stageSet, useSplitlite)

	// Generate probe/stageset_probe_ux_tree.go
	generateStageSetProbeUxTree(pkgPath, pkgPathRoot, modelPkg, stageSet, fieldToModelPkg)

	// Generate probe/stageset_probe_ux_table.go
	generateStageSetProbeUxTable(pkgPath, pkgPathRoot, modelPkg, stageSet, allStructs, pkgPathToField)

	// Generate probe/stageset_fill_up_form.go
	generateStageSetFillUpForm(pkgPath, pkgPathRoot, modelPkg, stageSet, allStructs, pkgPathToField)

	// Generate probe/stageset_form_callback.go
	generateStageSetFormCallback(pkgPath, pkgPathRoot, modelPkg, stageSet, allStructs, pkgPathToField)
}

func generateStageSetFormHelpers(pkgPath string) {
	code := `// generated code - do not edit
package probe

import (
	"fmt"
	"sort"
	"strings"
	"time"

	form "github.com/fullstack-lang/gong/lib/form/go/models"
)

// StageSetBasicFieldtoForm appends a basic field to the form
func StageSetBasicFieldtoForm(
	fieldName string,
	field any,
	formStage *form.Stage,
	formGroup *form.FormGroup,
) {
	formDiv := (&form.FormDiv{
		Name: fieldName,
	}).Stage(formStage)
	formGroup.FormDivs = append(formGroup.FormDivs, formDiv)

	if boolVal, ok := field.(bool); ok {
		checkBox := (&form.CheckBox{
			Name:  fieldName,
			Value: boolVal,
		}).Stage(formStage)
		formDiv.CheckBoxs = append(formDiv.CheckBoxs, checkBox)
		return
	}

	formField := (&form.FormField{
		Name:        fieldName,
		Label:       fieldName,
		Placeholder: "",
	}).Stage(formStage)
	formDiv.FormFields = append(formDiv.FormFields, formField)

	switch fieldVal := field.(type) {
	case string:
		formField.FormFieldString = (&form.FormFieldString{
			Name:  "string",
			Value: fieldVal,
		}).Stage(formStage)
	case int:
		formField.FormFieldInt = (&form.FormFieldInt{
			Name:  "int",
			Value: fieldVal,
		}).Stage(formStage)
	case int8:
		formField.FormFieldInt = (&form.FormFieldInt{
			Name:  "int",
			Value: int(fieldVal),
		}).Stage(formStage)
	case int16:
		formField.FormFieldInt = (&form.FormFieldInt{
			Name:  "int",
			Value: int(fieldVal),
		}).Stage(formStage)
	case int32:
		formField.FormFieldInt = (&form.FormFieldInt{
			Name:  "int",
			Value: int(fieldVal),
		}).Stage(formStage)
	case int64:
		formField.FormFieldInt = (&form.FormFieldInt{
			Name:  "int",
			Value: int(fieldVal),
		}).Stage(formStage)
	case uint:
		formField.FormFieldInt = (&form.FormFieldInt{
			Name:  "int",
			Value: int(fieldVal),
		}).Stage(formStage)
	case uint8:
		formField.FormFieldInt = (&form.FormFieldInt{
			Name:  "int",
			Value: int(fieldVal),
		}).Stage(formStage)
	case uint16:
		formField.FormFieldInt = (&form.FormFieldInt{
			Name:  "int",
			Value: int(fieldVal),
		}).Stage(formStage)
	case uint32:
		formField.FormFieldInt = (&form.FormFieldInt{
			Name:  "int",
			Value: int(fieldVal),
		}).Stage(formStage)
	case uint64:
		formField.FormFieldInt = (&form.FormFieldInt{
			Name:  "int",
			Value: int(fieldVal),
		}).Stage(formStage)
	case float32:
		formField.FormFieldFloat64 = (&form.FormFieldFloat64{
			Name:  "float",
			Value: float64(fieldVal),
		}).Stage(formStage)
	case float64:
		formField.FormFieldFloat64 = (&form.FormFieldFloat64{
			Name:  "float",
			Value: fieldVal,
		}).Stage(formStage)
	case time.Time:
		formField.FormFieldDate = (&form.FormFieldDate{
			Name:  fieldName + "Date",
			Value: fieldVal,
		}).Stage(formStage)
	case time.Duration:
		formField.FormFieldString = (&form.FormFieldString{
			Name:  "duration",
			Value: fieldVal.String(),
		}).Stage(formStage)
	default:
		formField.FormFieldString = (&form.FormFieldString{
			Name:  "string",
			Value: fmt.Sprintf("%v", fieldVal),
		}).Stage(formStage)
	}
}

// StageSetAssociationFieldToForm appends a FormFieldSelect dropdown for a pointer field
func StageSetAssociationFieldToForm[T interface {
	GetName() string
	comparable
}](
	fieldName string,
	field T,
	formGroup *form.FormGroup,
	instancesSet *map[T]struct{},
	formStage *form.Stage,
) {
	formDiv := (&form.FormDiv{
		Name: fieldName,
	}).Stage(formStage)
	formGroup.FormDivs = append(formGroup.FormDivs, formDiv)

	formField := (&form.FormField{
		Name:        fieldName,
		Label:       fieldName,
		Placeholder: "",
	}).Stage(formStage)
	formDiv.FormFields = append(formDiv.FormFields, formField)

	formFieldSelect := (&form.FormFieldSelect{
		Name:       "association",
		CanBeEmpty: true,
	}).Stage(formStage)
	formField.FormFieldSelect = formFieldSelect

	formFieldSelect.Options = make([]*form.Option, 0)
	if instancesSet != nil {
		instancesSorted := make([]T, 0, len(*instancesSet))
		for inst := range *instancesSet {
			instancesSorted = append(instancesSorted, inst)
		}
		sort.Slice(instancesSorted, func(i, j int) bool {
			return instancesSorted[i].GetName() < instancesSorted[j].GetName()
		})
		for _, instance := range instancesSorted {
			option := (&form.Option{
				Name: instance.GetName(),
			}).Stage(formStage)

			if instance == field {
				formFieldSelect.Value = option
			}
			formFieldSelect.Options = append(formFieldSelect.Options, option)
		}
	}
}

// StageSetFormDivSelectFieldToField updates a pointer field from the selected dropdown option
func StageSetFormDivSelectFieldToField[T interface {
	GetName() string
	comparable
}](
	field *T,
	instancesSet *map[T]struct{},
	formDiv *form.FormDiv,
) {
	if len(formDiv.FormFields) == 0 || formDiv.FormFields[0].FormFieldSelect == nil || formDiv.FormFields[0].FormFieldSelect.Value == nil {
		var zero T
		*field = zero
		return
	}
	selectedName := formDiv.FormFields[0].FormFieldSelect.Value.GetName()
	var zero T
	*field = zero
	if instancesSet != nil {
		for inst := range *instancesSet {
			if inst.GetName() == selectedName {
				*field = inst
				return
			}
		}
	}
}

// StageSetAssociationReverseFieldToForm shows reverse references pointing to this instance
func StageSetAssociationReverseFieldToForm(
	ownerStructName string,
	fieldName string,
	referencingNames []string,
	formGroup *form.FormGroup,
	formStage *form.Stage,
) {
	formDiv := (&form.FormDiv{
		Name: ownerStructName + ":" + fieldName,
	}).Stage(formStage)
	formGroup.FormDivs = append(formGroup.FormDivs, formDiv)

	val := strings.Join(referencingNames, ", ")
	formField := (&form.FormField{
		Name:  ownerStructName + ":" + fieldName,
		Label: "(" + ownerStructName + ") -> " + fieldName,
		FormFieldString: &form.FormFieldString{
			Value: val,
		},
	}).Stage(formStage)
	formDiv.FormFields = append(formDiv.FormFields, formField)
}
`
	writeFile(filepath.Join(pkgPath, "probe/stageset_form_helpers.go"), code)
}

func generateStageSetProbeMain(
	pkgPath string,
	pkgPathRoot string,
	modelPkg *models.ModelPkg,
	stageSet *models.StageSetModel,
	useSplitlite bool,
) {
	splitImportPath := "github.com/fullstack-lang/gong/lib/split"
	splitPkg := "split"
	prepareStageSetSuffix := ""
	if useSplitlite {
		splitImportPath = "github.com/fullstack-lang/gong/lib/splitlite"
		splitPkg = "splitlite"
		prepareStageSetSuffix = "Splitlite"
	}

	var metaPkgImports strings.Builder
	for _, f := range stageSet.Fields {
		alias := "ref_models"
		pkgP := pkgPathRoot + "/models"
		if !f.IsLocal {
			alias = "ref_" + f.PackageName
			pkgP = f.PackagePath
		}
		metaPkgImports.WriteString(fmt.Sprintf("\t\t{Alias: %q, Path: %q},\n", alias, `"`+pkgP+`"`))
	}

	code := `// generated code - do not edit
package probe

import (
	"embed"
	"encoding/base64"
	"fmt"
	"net/http"
	"regexp"
	"time"

	doc_models "github.com/fullstack-lang/gong/lib/doc/go/models"
	"github.com/fullstack-lang/gong/lib/doc/go/prepare"
	form_fullstack "github.com/fullstack-lang/gong/lib/form/go/fullstack"
	load_fullstack "github.com/fullstack-lang/gong/lib/load/go/fullstack"
	{{SplitPkg}}_fullstack "{{SplitImportPath}}/go/fullstack"
	table_fullstack "github.com/fullstack-lang/gong/lib/table/go/fullstack"
	tree_fullstack "github.com/fullstack-lang/gong/lib/tree/go/fullstack"

	form "github.com/fullstack-lang/gong/lib/form/go/models"
	load "github.com/fullstack-lang/gong/lib/load/go/models"
	{{SplitPkg}} "{{SplitImportPath}}/go/models"
	table "github.com/fullstack-lang/gong/lib/table/go/models"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"

	"{{PkgPathRoot}}/models"
)

type StageSetProbe struct {
	r                      *http.ServeMux
	stageSet               *models.StageSet
	treeStage              *tree.Stage
	treeNavigationStage    *tree.Stage
	formStage              *form.Stage
	tableStage             *table.Stage
	notificationTableStage *table.Stage
	splitStage             *{{SplitPkg}}.Stage
	loadStage              *load.Stage

	diagramEditor *{{SplitPkg}}.AsSplitArea
	docStager     *doc_models.Stager

	fileName string

	dataEditor *{{SplitPkg}}.AsSplit

	notification []*Notification

	maxElementsNbPerGongStructNode int
	commitMode                     bool
	bulkDeleteMode                 bool
}

func NewStageSetProbe(
	r *http.ServeMux,
	goModelsDir embed.FS,
	goDiagramsDir embed.FS,
	embeddedDiagrams bool,
	stageSet *models.StageSet,
) (probe *StageSetProbe) {

	splitStage, _ := {{SplitPkg}}_fullstack.NewStackInstance(r, stageSet.GetProbeSplitStageName())
	splitStage.Commit()

	treeStage, _ := tree_fullstack.NewStackInstance(r, stageSet.GetProbeSplitStageName()+"-tree")
	treeNavigationStage, _ := tree_fullstack.NewStackInstance(r, stageSet.GetProbeSplitStageName()+"-tree-nav")

	tableStage, _ := table_fullstack.NewStackInstance(r, stageSet.GetProbeSplitStageName()+"-table")
	tableStage.Commit()

	notificationTableStage, _ := table_fullstack.NewStackInstance(r, stageSet.GetProbeSplitStageName()+"-notif")
	notificationTableStage.Commit()

	formStage, _ := form_fullstack.NewStackInstance(r, stageSet.GetProbeSplitStageName()+"-form")
	formStage.Commit()

	loadStage, _ := load_fullstack.NewStackInstance(r, stageSet.GetProbeSplitStageName()+"-load")

	probe = &StageSetProbe{
		r:                              r,
		stageSet:                       stageSet,
		treeStage:                      treeStage,
		treeNavigationStage:            treeNavigationStage,
		formStage:                      formStage,
		tableStage:                     tableStage,
		notificationTableStage:         notificationTableStage,
		splitStage:                     splitStage,
		loadStage:                      loadStage,
		maxElementsNbPerGongStructNode: 10,
		commitMode:                     true,
	}

	probe.diagramEditor = &{{SplitPkg}}.AsSplitArea{
		Name:             "Bottom",
		ShowNameInHeader: false,
		Size:             50,
	}

	metaPackageImports := []*doc_models.MetaPackageImport{
{{MetaPackageImports}}	}

	probe.docStager = prepare.PrepareStageSet{{PrepareStageSetSuffix}}(
		r,
		embeddedDiagrams,
		"{{PkgPathRoot}}:"+stageSet.GetProbeSplitStageName(),
		metaPackageImports,
		goModelsDir,
		goDiagramsDir,
		probe.diagramEditor,
		nil,
	)

	probe.dataEditor = &{{SplitPkg}}.AsSplit{
		Name:          "StageSet Top, sidebar, table & form",
		Direction:     {{SplitPkg}}.Horizontal,
		IsSizeInPixel: true,
		AsSplitAreas: []*{{SplitPkg}}.AsSplitArea{
			{
				Name: "sidebar",
				Size: 525,
				AsSplit: &{{SplitPkg}}.AsSplit{
					Direction:              {{SplitPkg}}.Vertical,
					IsSizeInPixel:          true,
					IsWithCustomGutterSize: true,
					GutterSize:             1,
					AsSplitAreas: []*{{SplitPkg}}.AsSplitArea{
						{
							Name: "sidebar tree",
							Size: 53,
							Tree: &{{SplitPkg}}.Tree{
								Name:      "Sidebar",
								StackName: probe.treeNavigationStage.GetName(),
							},
						},
						{
							Name:  "sidebar tree",
							IsAny: true,
							Tree: &{{SplitPkg}}.Tree{
								Name:      "Sidebar",
								StackName: probe.treeStage.GetName(),
							},
						},
						{
							Name: "load",
							Size: 70,
							Load: &{{SplitPkg}}.Load{
								Name:      "Table",
								StackName: probe.loadStage.GetName(),
							},
						},
					},
				},
			},
			{
				Name:  "both tables",
				IsAny: true,
				AsSplit: &{{SplitPkg}}.AsSplit{
					Direction: {{SplitPkg}}.Vertical,
					AsSplitAreas: []*{{SplitPkg}}.AsSplitArea{
						{
							Name: "table",
							Size: 50,
							Table: &{{SplitPkg}}.Table{
								Name:      "Table",
								StackName: probe.tableStage.GetName(),
							},
						},
						{
							Name: "notification table",
							Size: 50,
							Table: &{{SplitPkg}}.Table{
								Name:      "Table",
								StackName: probe.notificationTableStage.GetName(),
							},
						},
					},
				},
			},
			{
				Name: "form",
				Size: 525,
				Form: &{{SplitPkg}}.Form{
					Name:      "Form",
					StackName: probe.formStage.GetName(),
				},
			},
		},
	}

	{{SplitPkg}}.StageBranch(probe.splitStage, &{{SplitPkg}}.View{
		Name: "StageSet Main view",
		RootAsSplitAreas: []*{{SplitPkg}}.AsSplitArea{
			{
				Name:    "Top",
				Size:    50,
				AsSplit: probe.dataEditor,
			},
			probe.diagramEditor,
		},
	})
	probe.splitStage.Commit()

	probe.initLoadStage()
	probe.ux_tree()

	return probe
}

func (probe *StageSetProbe) GetDocStager() *doc_models.Stager {
	return probe.docStager
}

func (probe *StageSetProbe) GetStageSet() *models.StageSet {
	return probe.stageSet
}

func (probe *StageSetProbe) GetSplitStage() *{{SplitPkg}}.Stage {
	return probe.splitStage
}

func (probe *StageSetProbe) GetTreeStage() *tree.Stage {
	return probe.treeStage
}

func (probe *StageSetProbe) GetFormStage() *form.Stage {
	return probe.formStage
}

func (probe *StageSetProbe) GetTableStage() *table.Stage {
	return probe.tableStage
}

func (probe *StageSetProbe) SetMaxElementsNbPerGongStructNode(nb int) {
	probe.maxElementsNbPerGongStructNode = nb
}

func (probe *StageSetProbe) GetMaxElementsNbPerGongStructNode() int {
	return probe.maxElementsNbPerGongStructNode
}

func (probe *StageSetProbe) Refresh() {
	probe.ux_tree()
	probe.ux_table()
}

func (probe *StageSetProbe) AddNotification(date time.Time, message string) {
	notification := &Notification{
		Date:    date,
		Message: message,
	}
	probe.notification = append(probe.notification, notification)
}

func (probe *StageSetProbe) CommitNotificationTable() {
	probe.notificationTableStage.Reset()
	tableInstance := new(table.Table)
	tableInstance.Name = "Notifications"
	tableInstance.HasFiltering = true

	column := new(table.DisplayedColumn)
	column.Name = "Date"
	tableInstance.DisplayedColumns = append(tableInstance.DisplayedColumns, column)

	column = new(table.DisplayedColumn)
	column.Name = "Message"
	tableInstance.DisplayedColumns = append(tableInstance.DisplayedColumns, column)

	for _, notification := range probe.notification {
		row := new(table.Row)
		row.Name = notification.Date.Format(time.StampMicro)

		cell := &table.Cell{Name: "Date"}
		cell.CellString = &table.CellString{Value: notification.Date.Format(time.StampMicro)}
		row.Cells = append(row.Cells, cell)

		cell = &table.Cell{Name: "Message"}
		cell.CellString = &table.CellString{Value: notification.Message}
		row.Cells = append(row.Cells, cell)

		tableInstance.Rows = append(tableInstance.Rows, row)
	}

	table.StageBranch(probe.notificationTableStage, tableInstance)
	probe.notificationTableStage.Commit()
}

func (probe *StageSetProbe) ExportStage() {
	probe.loadStage.Reset()

	fileToDownload := new(load.FileToDownload)
	if probe.fileName == "" {
		probe.fileName = "{{ModelPkgName}}-stageset.go"
	}

	prefixRegex := regexp.MustCompile("^\\d{8} \\d{4} ")
	cleanFileName := prefixRegex.ReplaceAllString(probe.fileName, "")

	fileToDownload.Name = time.Now().Format("20060102 1504 ") + cleanFileName

	stageString, err := probe.stageSet.MarshallToString("main")
	if err != nil {
		probe.AddNotification(time.Now(), "Error serializing stage: "+err.Error())
		probe.CommitNotificationTable()
		return
	}

	fileToDownload.Base64EncodedContent = base64.StdEncoding.EncodeToString([]byte(stageString))
	load.StageBranch(probe.loadStage, fileToDownload)
	probe.loadStage.Commit()

	time.Sleep(1 * time.Second)
	probe.initLoadStage()
}

type stageSetLoadProxy struct {
	probe *StageSetProbe
}

func (proxy *stageSetLoadProxy) OnFileUpload(uploadedFile *load.FileToUpload) error {
	proxy.probe.fileName = uploadedFile.GetName()
	decodedBytes, err := base64.StdEncoding.DecodeString(uploadedFile.Base64EncodedContent)
	if err != nil {
		return fmt.Errorf("base64.StdEncoding.DecodeString failed: %w", err)
	}
	proxy.probe.stageSet.Reset()
	err = proxy.probe.stageSet.ParseAstString(string(decodedBytes), true)
	if err != nil {
		return err
	}
	proxy.probe.stageSet.Commit()
	proxy.probe.Refresh()
	return nil
}

func (probe *StageSetProbe) initLoadStage() {
	probe.loadStage.Reset()

	fileToUpload := &load.FileToUpload{
		Name: "Name of file",
		FileToUploadProxy: &stageSetLoadProxy{
			probe: probe,
		},
	}

	load.StageBranch(probe.loadStage,
		fileToUpload,
	)

	message := &load.Message{
		Name: "Drop your stage.go file here or ",
	}

	message.Stage(probe.loadStage)

	probe.loadStage.Commit()
}
`

	code = strings.ReplaceAll(code, "{{SplitPkg}}", splitPkg)
	code = strings.ReplaceAll(code, "{{SplitImportPath}}", splitImportPath)
	code = strings.ReplaceAll(code, "{{PkgPathRoot}}", pkgPathRoot)
	code = strings.ReplaceAll(code, "{{MetaPackageImports}}", metaPkgImports.String())
	code = strings.ReplaceAll(code, "{{PrepareStageSetSuffix}}", prepareStageSetSuffix)
	code = strings.ReplaceAll(code, "{{ModelPkgName}}", modelPkg.Name)

	writeFile(filepath.Join(pkgPath, "probe/stageset_probe.go"), code)
}

func generateStageSetProbeUxTree(
	pkgPath string,
	pkgPathRoot string,
	modelPkg *models.ModelPkg,
	stageSet *models.StageSetModel,
	fieldToModelPkg map[*models.StageSetField]*models.ModelPkg,
) {
	var extImports strings.Builder
	for _, f := range stageSet.Fields {
		if !f.IsLocal {
			extImports.WriteString(fmt.Sprintf("\n\t\"%s\"", f.PackagePath))
		}
	}

	var pkgTreeNodes strings.Builder
	for _, f := range stageSet.Fields {
		mPkg := fieldToModelPkg[f]
		if mPkg == nil {
			continue
		}
		var sNames []string
		for sName, gs := range mPkg.GongStructs {
			if gs.ModelPkg.PkgPath == mPkg.PkgPath && gs.HasNameField() && !gs.IsOmittedForMarshalling {
				sNames = append(sNames, sName[len(mPkg.PkgPath)+1:])
			}
		}
		sort.Strings(sNames)

		pkgTreeNodes.WriteString(fmt.Sprintf(`
	// Package node: %s
	pkgNode_%s := &tree_models.Node{
		Name:       "%s",
		IsExpanded: true,
	}
	pkgNode_%s.OnIsExpandedChange = func(isExpanded bool) {
		pkgNode_%s.IsExpanded = isExpanded
	}
	topNode.Children = append(topNode.Children, pkgNode_%s)
`, f.PackageName, f.Name, f.PackageName, f.Name, f.Name, f.Name))

		for _, sName := range sNames {
			sPlural := sName + "s"

			pkgTreeNodes.WriteString(fmt.Sprintf(`
	{
		count := len(probe.stageSet.%s.%s)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("%s (%%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all %s instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		pkgNode_%s.Children = append(pkgNode_%s.Children, nodeGongstruct)

		nodeGongstruct.OnIsExpandedChange = func(isExpanded bool) {
			nodeGongstruct.IsExpanded = isExpanded
		}

		nodeGongstruct.OnClick = func(frontNode *tree_models.Node) {
			updateStageSetTable_%s_%s(probe)
			for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
				node.BackgroundColor = ""
			}
			nodeGongstruct.BackgroundColor = "lightgrey"
			probe.treeStage.Commit()
		}

		instCount := 0
		for _inst := range probe.stageSet.%s.%s {
			if instCount >= probe.GetMaxElementsNbPerGongStructNode() {
				nodeGongstruct.Children = append(nodeGongstruct.Children, &tree_models.Node{Name: "..."})
				break
			}
			instCount++
			_captured := _inst
			nodeInstance := &tree_models.Node{
				Name:            _captured.GetName(),
				IsNodeClickable: true,
				OnClick: func(frontNode *tree_models.Node) {
					StageSetFillUpFormFromGongstruct(_captured, probe)
					for node := range *probe.treeStage.GetInstancesSet[*tree_models.Node]() {
						node.BackgroundColor = ""
					}
					frontNode.BackgroundColor = "lightgrey"
					probe.treeStage.Commit()
				},
			}
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
		}
	}
`, f.Name, sPlural, sName, sName, f.Name, f.Name, sName, f.Name, f.Name, sPlural))
		}
	}

	code := fmt.Sprintf(`// generated code - do not edit
package probe

import (
	"fmt"

	tree_buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree_models "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (probe *StageSetProbe) ux_navigation_tree() {
	probe.treeNavigationStage.Reset()
	sidebar := &tree_models.Tree{Name: "Sidebar"}
	tree_models.StageBranch(probe.treeNavigationStage, sidebar)
	probe.treeNavigationStage.Commit()
}

func (probe *StageSetProbe) ux_tree() {
	probe.ux_navigation_tree()
	probe.treeStage.Reset()

	sidebar := &tree_models.Tree{Name: "Sidebar"}
	topNode := &tree_models.Node{
		Name:       "StageSet",
		IsExpanded: true,
	}
	topNode.OnIsExpandedChange = func(isExpanded bool) {
		topNode.IsExpanded = isExpanded
	}
	sidebar.RootNodes = append(sidebar.RootNodes, topNode)

	refreshButton := &tree_models.Button{
		Name:            "RefreshButton " + string(tree_buttons.BUTTON_refresh),
		Icon:            string(tree_buttons.BUTTON_refresh),
		HasToolTip:      true,
		ToolTipText:     "Refresh probe",
		ToolTipPosition: tree_models.Below,
		OnClick: func() {
			probe.Refresh()
		},
	}
	topNode.Buttons = append(topNode.Buttons, refreshButton)

	resetButton := &tree_models.Button{
		Name:            "ResetButton " + string(tree_buttons.BUTTON_reset_tv),
		Icon:            string(tree_buttons.BUTTON_reset_tv),
		HasToolTip:      true,
		ToolTipText:     "Reset all stages in StageSet",
		ToolTipPosition: tree_models.Below,
		OnClick: func() {
			probe.stageSet.Reset()
			probe.stageSet.Commit()
			probe.Refresh()
		},
	}
	topNode.Buttons = append(topNode.Buttons, resetButton)

	exportGoButton := &tree_models.Button{
		Name:            "ExportGoButton " + string(tree_buttons.BUTTON_file_download),
		Icon:            string(tree_buttons.BUTTON_file_download),
		HasToolTip:      true,
		ToolTipText:     "Export Go code for all stages",
		ToolTipPosition: tree_models.Below,
		OnClick: func() {
			probe.ExportStage()
		},
	}
	topNode.Buttons = append(topNode.Buttons, exportGoButton)
%s
	tree_models.StageBranch(probe.treeStage, sidebar)
	probe.treeStage.Commit()
}
`, pkgTreeNodes.String())

	writeFile(filepath.Join(pkgPath, "probe/stageset_probe_ux_tree.go"), code)
}

func generateStageSetProbeUxTable(
	pkgPath string,
	pkgPathRoot string,
	modelPkg *models.ModelPkg,
	stageSet *models.StageSetModel,
	allStructs []structInfo,
	pkgPathToField map[string]*models.StageSetField,
) {
	var extImports strings.Builder
	for _, f := range stageSet.Fields {
		if !f.IsLocal {
			extImports.WriteString(fmt.Sprintf("\n\t\"%s\"", f.PackagePath))
		}
	}

	var tableSwitchCases strings.Builder
	var tableFunctions strings.Builder

	for _, si := range allStructs {
		tblName := si.structName
		tableSwitchCases.WriteString(fmt.Sprintf("\tcase \"%s\":\n\t\tupdateStageSetTable_%s_%s(probe)\n", tblName, si.structName, si.pkgField.Name))

		// Find reverse references: any struct in allStructs that has a pointer to si
		var reversePointers []struct {
			sourceStruct   structInfo
			fieldName      string
			isSlice        bool
		}
		for _, other := range allStructs {
			for _, fld := range other.fields {
				if (fld.IsPointer || fld.IsSliceOfPointer) && fld.TargetPackagePath == si.pkgPath && fld.TargetStructName == si.structName {
					reversePointers = append(reversePointers, struct {
						sourceStruct   structInfo
						fieldName      string
						isSlice        bool
					}{
						sourceStruct: other,
						fieldName:    fld.Name,
						isSlice:      fld.IsSliceOfPointer,
					})
				}
			}
		}

		var columnDefs strings.Builder
		for _, fld := range si.fields {
			columnDefs.WriteString(fmt.Sprintf("\t{\n\t\tcol := new(table_models.DisplayedColumn)\n\t\tcol.Name = \"%s\"\n\t\ttable.DisplayedColumns = append(table.DisplayedColumns, col)\n\t}\n", fld.Name))
		}
		for _, rev := range reversePointers {
			colName := fmt.Sprintf("(%s) -> %s", rev.sourceStruct.typeQual, rev.fieldName)
			columnDefs.WriteString(fmt.Sprintf("\t{\n\t\tcol := new(table_models.DisplayedColumn)\n\t\tcol.Name = \"%s\"\n\t\ttable.DisplayedColumns = append(table.DisplayedColumns, col)\n\t}\n", colName))
		}

		var cellAssignments strings.Builder
		for _, fld := range si.fields {
			if fld.IsPointer {
				cellAssignments.WriteString(fmt.Sprintf(`
		{
			cell := &table_models.Cell{Name: "%s"}
			val := ""
			if structInstance.%s != nil {
				val = structInstance.%s.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}
`, fld.Name, fld.Name, fld.Name))
			} else if fld.IsSliceOfPointer {
				cellAssignments.WriteString(fmt.Sprintf(`
		{
			cell := &table_models.Cell{Name: "%s"}
			var names []string
			for _, elem := range structInstance.%s {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}
`, fld.Name, fld.Name))
			} else {
				switch fld.BasicKind {
				case types.Int, types.Int8, types.Int16, types.Int32, types.Int64,
					types.Uint, types.Uint8, types.Uint16, types.Uint32, types.Uint64:
					cellAssignments.WriteString(fmt.Sprintf(`
		{
			cell := &table_models.Cell{Name: "%s"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.%s)}
			row.Cells = append(row.Cells, cell)
		}
`, fld.Name, fld.Name))
				case types.Float32, types.Float64:
					cellAssignments.WriteString(fmt.Sprintf(`
		{
			cell := &table_models.Cell{Name: "%s"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.%s)}
			row.Cells = append(row.Cells, cell)
		}
`, fld.Name, fld.Name))
				case types.Bool:
					cellAssignments.WriteString(fmt.Sprintf(`
		{
			cell := &table_models.Cell{Name: "%s"}
			cell.CellBool = &table_models.CellBool{Value: structInstance.%s}
			row.Cells = append(row.Cells, cell)
		}
`, fld.Name, fld.Name))
				default:
					cellAssignments.WriteString(fmt.Sprintf(`
		{
			cell := &table_models.Cell{Name: "%s"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%%v", structInstance.%s)}
			row.Cells = append(row.Cells, cell)
		}
`, fld.Name, fld.Name))
				}
			}
		}

		for _, rev := range reversePointers {
			colName := fmt.Sprintf("(%s) -> %s", rev.sourceStruct.typeQual, rev.fieldName)
			if rev.isSlice {
				cellAssignments.WriteString(fmt.Sprintf(`
		{
			cell := &table_models.Cell{Name: "%s"}
			var refNames []string
			for src := range probe.stageSet.%s.%ss {
				for _, target := range src.%s {
					if target == structInstance {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}
`, colName, rev.sourceStruct.pkgField.Name, rev.sourceStruct.structName, rev.fieldName))
			} else {
				cellAssignments.WriteString(fmt.Sprintf(`
		{
			cell := &table_models.Cell{Name: "%s"}
			var refNames []string
			for src := range probe.stageSet.%s.%ss {
				if src.%s == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}
`, colName, rev.sourceStruct.pkgField.Name, rev.sourceStruct.structName, rev.fieldName))
			}
		}

		tableFunctions.WriteString(fmt.Sprintf(`
func updateStageSetTable_%s_%s(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "%s"
	table.HasColumnSorting = true
	table.HasFiltering = true
	table.HasPaginator = true

	colID := new(table_models.DisplayedColumn)
	colID.Name = "ID"
	table.DisplayedColumns = append(table.DisplayedColumns, colID)

	colDel := new(table_models.DisplayedColumn)
	colDel.Name = "Delete"
	table.DisplayedColumns = append(table.DisplayedColumns, colDel)

%s
	// Sort instances by name
	instances := make([]*%s, 0, len(probe.stageSet.%s.%ss))
	for inst := range probe.stageSet.%s.%ss {
		instances = append(instances, inst)
	}
	sort.Slice(instances, func(i, j int) bool {
		return instances[i].GetName() < instances[j].GetName()
	})

	for idx, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				StageSetFillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: idx}
		row.Cells = append(row.Cells, cellID)

		cellDel := &table_models.Cell{Name: "Delete Icon"}
		cellIcon := &table_models.CellIcon{
			Name:                fmt.Sprintf("Delete %%s", structInstance.GetName()),
			Icon:                string(maticons.BUTTON_delete),
			NeedsConfirmation:   true,
			ConfirmationMessage: "Do you confirm you want to delete this instance?",
		}
		cellIcon.Impl = &table_models.FunctionalCellIconProxy{
			OnUpdated: func(stage *table_models.Stage, ci, uci *table_models.CellIcon) {
				_captured.UnstageVoid(probe.stageSet.%s)
				probe.stageSet.Clean()
				probe.stageSet.Commit()
				updateStageSetTable_%s_%s(probe)
				probe.ux_tree()
			},
		}
		cellDel.CellIcon = cellIcon
		row.Cells = append(row.Cells, cellDel)

%s
		table.Rows = append(table.Rows, row)
	}

	table_models.StageBranch(probe.tableStage, table)
	probe.tableStage.Commit()
}
`, si.structName, si.pkgField.Name, si.structName, columnDefs.String(), si.typeQual, si.pkgField.Name, si.structName, si.pkgField.Name, si.structName, si.pkgField.Name, si.structName, si.pkgField.Name, cellAssignments.String()))
	}

	code := fmt.Sprintf(`// generated code - do not edit
package probe

import (
	"fmt"
	"sort"
	"strings"

	table_models "github.com/fullstack-lang/gong/lib/table/go/models"
	maticons "github.com/fullstack-lang/maticons/maticons"

	"%s/models"%s
)

type tableRowUpdater struct {
	onClick func()
}

func (u *tableRowUpdater) RowUpdated(stage *table_models.Stage, row, updatedRow *table_models.Row) {
	if u.onClick != nil {
		u.onClick()
	}
}

func (probe *StageSetProbe) ux_table() {
	var tableName string
	for tbl := range probe.tableStage.Tables {
		tableName = tbl.Name
	}
	switch tableName {
%s	}
}
%s
`, pkgPathRoot, extImports.String(), tableSwitchCases.String(), tableFunctions.String())

	writeFile(filepath.Join(pkgPath, "probe/stageset_probe_ux_table.go"), code)
}

func generateStageSetFillUpForm(
	pkgPath string,
	pkgPathRoot string,
	modelPkg *models.ModelPkg,
	stageSet *models.StageSetModel,
	allStructs []structInfo,
	pkgPathToField map[string]*models.StageSetField,
) {
	var extImports strings.Builder
	for _, f := range stageSet.Fields {
		if !f.IsLocal {
			extImports.WriteString(fmt.Sprintf("\n\t\"%s\"", f.PackagePath))
		}
	}

	var cases strings.Builder
	for _, si := range allStructs {
		cases.WriteString(fmt.Sprintf("\tcase *%s:\n", si.typeQual))

		for _, fld := range si.fields {
			if fld.IsPointer {
				targetSSF := pkgPathToField[fld.TargetPackagePath]
				if targetSSF != nil {
					targetTypeQual := fld.TargetStructName
					if !targetSSF.IsLocal {
						targetTypeQual = targetSSF.PackageName + "." + fld.TargetStructName
					} else {
						targetTypeQual = "models." + fld.TargetStructName
					}
					cases.WriteString(fmt.Sprintf("\t\tStageSetAssociationFieldToForm(\"%s\", inst.%s, formGroup, probe.stageSet.%s.GetInstancesSet[*%s](), probe.formStage)\n", fld.Name, fld.Name, targetSSF.Name, targetTypeQual))
				}
			} else if fld.IsSliceOfPointer {
				targetSSF := pkgPathToField[fld.TargetPackagePath]
				if targetSSF != nil {
					cases.WriteString(fmt.Sprintf(`
		{
			// Slice of pointers: %s
			div := (&form.FormDiv{Name: "%s"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.%s {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "%s",
				Label: "%s",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
`, fld.Name, fld.Name, fld.Name, fld.Name, fld.Name))
				}
			} else {
				cases.WriteString(fmt.Sprintf("\t\tStageSetBasicFieldtoForm(\"%s\", inst.%s, probe.formStage, formGroup)\n", fld.Name, fld.Name))
			}
		}

		// Reverse pointers pointing to this struct
		for _, other := range allStructs {
			for _, fld := range other.fields {
				if fld.IsPointer && fld.TargetPackagePath == si.pkgPath && fld.TargetStructName == si.structName {
					cases.WriteString(fmt.Sprintf(`
		{
			var refNames []string
			for src := range probe.stageSet.%s.%ss {
				if src.%s == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("%s", "%s", refNames, formGroup, probe.formStage)
		}
`, other.pkgField.Name, other.structName, fld.Name, other.typeQual, fld.Name))
				}
				if fld.IsSliceOfPointer && fld.TargetPackagePath == si.pkgPath && fld.TargetStructName == si.structName {
					cases.WriteString(fmt.Sprintf(`
		{
			var refNames []string
			for src := range probe.stageSet.%s.%ss {
				for _, target := range src.%s {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("%s", "%s", refNames, formGroup, probe.formStage)
		}
`, other.pkgField.Name, other.structName, fld.Name, other.typeQual, fld.Name))
				}
			}
		}
	}

	code := fmt.Sprintf(`// generated code - do not edit
package probe

import (
	"sort"
	"strings"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"%s/models"%s
)

var (
	_ = sort.Strings
	_ = strings.Join
)

func StageSetFillUpForm(
	instance any,
	formGroup *form.FormGroup,
	probe *StageSetProbe,
) {
	switch inst := instance.(type) {
%s	default:
		_ = inst
	}
}
`, pkgPathRoot, extImports.String(), cases.String())

	writeFile(filepath.Join(pkgPath, "probe/stageset_fill_up_form.go"), code)
}

func generateStageSetFormCallback(
	pkgPath string,
	pkgPathRoot string,
	modelPkg *models.ModelPkg,
	stageSet *models.StageSetModel,
	allStructs []structInfo,
	pkgPathToField map[string]*models.StageSetField,
) {
	var extImports strings.Builder
	for _, f := range stageSet.Fields {
		if !f.IsLocal {
			extImports.WriteString(fmt.Sprintf("\n\t\"%s\"", f.PackagePath))
		}
	}

	var fillUpDispatchCases strings.Builder
	var saveFunctions strings.Builder

	for _, si := range allStructs {
		fillUpDispatchCases.WriteString(fmt.Sprintf(`	case *%s:
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				probe.stageSet.%s.Lock()
				defer probe.stageSet.%s.Unlock()
				probe.formStage.Checkout()
				saveStageSet_%s_%s(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(probe.stageSet.%s)
				}
				probe.stageSet.%s.Commit()
				updateStageSetTable_%s_%s(probe)
				probe.ux_tree()
			},
		}
		StageSetFillUpForm(inst, formGroup, probe)
`, si.typeQual, si.pkgField.Name, si.pkgField.Name, si.structName, si.pkgField.Name, si.pkgField.Name, si.pkgField.Name, si.structName, si.pkgField.Name))

		var fieldSaves strings.Builder
		for _, fld := range si.fields {
			if fld.IsPointer {
				targetSSF := pkgPathToField[fld.TargetPackagePath]
				if targetSSF != nil {
					targetTypeQual := fld.TargetStructName
					if !targetSSF.IsLocal {
						targetTypeQual = targetSSF.PackageName + "." + fld.TargetStructName
					} else {
						targetTypeQual = "models." + fld.TargetStructName
					}
					fieldSaves.WriteString(fmt.Sprintf("\t\tcase \"%s\":\n\t\t\tStageSetFormDivSelectFieldToField(&inst.%s, probe.stageSet.%s.GetInstancesSet[*%s](), formDiv)\n", fld.Name, fld.Name, targetSSF.Name, targetTypeQual))
				}
			} else if !fld.IsSliceOfPointer {
				fieldSaves.WriteString(fmt.Sprintf("\t\tcase \"%s\":\n\t\t\tFormDivBasicFieldToField(&inst.%s, formDiv)\n", fld.Name, fld.Name))
			}
		}

		saveFunctions.WriteString(fmt.Sprintf(`
func saveStageSet_%s_%s(
	inst *%s,
	probe *StageSetProbe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
%s		}
	}
}
`, si.structName, si.pkgField.Name, si.typeQual, fieldSaves.String()))
	}

	code := fmt.Sprintf(`// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"%s/models"%s
)

type functionalStageSetFormCallback struct {
	onSave func()
}

func (cb *functionalStageSetFormCallback) OnSave() {
	if cb.onSave != nil {
		cb.onSave()
	}
}

func StageSetFillUpFormFromGongstruct(
	instance any,
	probe *StageSetProbe,
) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name: "Form",
	}).Stage(probe.formStage)

	switch inst := instance.(type) {
%s	default:
		_ = inst
	}

	probe.formStage.Commit()
}
%s
`, pkgPathRoot, extImports.String(), fillUpDispatchCases.String(), saveFunctions.String())

	writeFile(filepath.Join(pkgPath, "probe/stageset_form_callback.go"), code)
}

func writeFile(filePath string, content string) {
	os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
	file, err := os.Create(filePath)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	fmt.Fprint(file, content)
}
