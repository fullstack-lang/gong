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

func stageAccessor(f *models.StageSetField) string {
	if f.IsLocal {
		return "probe.stageOfInterest"
	}
	return "probe.stageOfInterest." + f.Name
}

func CodeGeneratorStageSetProbe(
	modelPkg *models.ModelPkg,
	pkgName string,
	pkgPath string,
	pkgPathRoot string,
	useSplitlite bool,
) {
	stageSet := modelPkg.StageSet
	if stageSet == nil || len(stageSet.Fields) <= 1 {
		return
	}

	if before, _, ok := strings.Cut(pkgPathRoot, "/go/models"); ok {
		pkgPathRoot = before + "/go"
	} else {
		pkgPathRoot = strings.ReplaceAll(pkgPathRoot, "/models", "")
	}

	goModDir, modPath := models.FindGoMod(pkgPath)

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
				typeQual = f.ImportAlias + "." + sName
			} else {
				typeQual = "models." + sName
			}
			flds := golang_models.ExtractStructFields(mPkg, sName, pkgPathToField)
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

	// Generate probe/form_helpers.go
	generateStageSetFormHelpers(pkgPath)

	// Generate probe/probe.go
	generateStageSetProbeMain(pkgPath, pkgPathRoot, modelPkg, stageSet, useSplitlite)

	// Generate probe/probe_ux_tree.go
	generateStageSetProbeUxTree(pkgPath, pkgPathRoot, modelPkg, stageSet, fieldToModelPkg)

	// Generate probe/probe_ux_table.go
	generateStageSetProbeUxTable(pkgPath, pkgPathRoot, modelPkg, stageSet, allStructs, pkgPathToField)

	// Generate probe/fill_up_form.go
	generateStageSetFillUpForm(pkgPath, pkgPathRoot, modelPkg, stageSet, allStructs, pkgPathToField)

	// Generate probe/form_callback.go, probe/fill_up_form_from_gongstruct.go, probe/probe_ux_form.go
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
		Name:                 "association",
		CanBeEmpty:           true,
		PreserveInitialOrder: true,
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
		nameCount := make(map[string]int)
		for _, inst := range instancesSorted {
			nameCount[inst.GetName()]++
		}
		for i, instance := range instancesSorted {
			optName := instance.GetName()
			if optName == "" {
				optName = fmt.Sprintf("[Unnamed] (#%d)", i+1)
			} else if nameCount[optName] > 1 {
				optName = fmt.Sprintf("%s (#%d)", optName, i+1)
			}
			option := (&form.Option{
				Name: optName,
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
		instancesSorted := make([]T, 0, len(*instancesSet))
		for inst := range *instancesSet {
			instancesSorted = append(instancesSorted, inst)
		}
		sort.Slice(instancesSorted, func(i, j int) bool {
			return instancesSorted[i].GetName() < instancesSorted[j].GetName()
		})
		nameCount := make(map[string]int)
		for _, inst := range instancesSorted {
			nameCount[inst.GetName()]++
		}
		for i, inst := range instancesSorted {
			optName := inst.GetName()
			if optName == "" {
				optName = fmt.Sprintf("[Unnamed] (#%d)", i+1)
			} else if nameCount[optName] > 1 {
				optName = fmt.Sprintf("%s (#%d)", optName, i+1)
			}
			if optName == selectedName {
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

// StageSetEnumStringFieldToForm appends a FormFieldSelect dropdown for a string enum
func StageSetEnumStringFieldToForm[TF interface {
	Codes() []string
	CodeValues() []string
	ToString() string
}](
	fieldName string,
	field TF,
	formGroup *form.FormGroup,
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
		Name: "enum",
	}).Stage(formStage)
	formField.FormFieldSelect = formFieldSelect

	formFieldSelect.Options = make([]*form.Option, 0)
	for idx, optionCode := range field.Codes() {
		optionValue := field.CodeValues()[idx]

		option := (&form.Option{
			Name: optionCode,
		}).Stage(formStage)

		if field.ToString() == optionValue {
			formFieldSelect.Value = option
		}

		formFieldSelect.Options =
			append(formFieldSelect.Options, option)
	}
}

// StageSetEnumIntFieldToForm appends a FormFieldSelect dropdown for an int enum
func StageSetEnumIntFieldToForm[TF interface {
	Codes() []string
	CodeValues() []int
}](
	fieldName string,
	field TF,
	formGroup *form.FormGroup,
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
		Name: "enum",
	}).Stage(formStage)
	formField.FormFieldSelect = formFieldSelect

	formFieldSelect.Options = make([]*form.Option, 0)
	for idx, optionCode := range field.Codes() {
		optionValue := field.CodeValues()[idx]

		option := (&form.Option{
			Name: optionCode,
		}).Stage(formStage)

		if any(field) == any(optionValue) {
			formFieldSelect.Value = option
		}

		formFieldSelect.Options =
			append(formFieldSelect.Options, option)
	}
}

// StageSetFormDivEnumStringFieldToField updates a string enum field from the form
func StageSetFormDivEnumStringFieldToField[TF interface {
	FromCodeString(input string) error
}](field TF, formDiv *form.FormDiv) {
	if len(formDiv.FormFields) > 0 && formDiv.FormFields[0].FormFieldSelect != nil && formDiv.FormFields[0].FormFieldSelect.Value != nil {
		_ = field.FromCodeString(formDiv.FormFields[0].FormFieldSelect.Value.GetName())
	}
}

// StageSetFormDivEnumIntFieldToField updates an int enum field from the form
func StageSetFormDivEnumIntFieldToField[TF interface {
	FromCodeString(input string) error
}](field TF, formDiv *form.FormDiv) {
	if len(formDiv.FormFields) > 0 && formDiv.FormFields[0].FormFieldSelect != nil && formDiv.FormFields[0].FormFieldSelect.Value != nil {
		_ = field.FromCodeString(formDiv.FormFields[0].FormFieldSelect.Value.GetName())
	}
}
`
	writeFile(filepath.Join(pkgPath, "probe/form_helpers.go"), code)
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
			alias = "ref_" + f.ImportAlias
			pkgP = f.PackagePath
		}
		metaPkgImports.WriteString(fmt.Sprintf("\t\t{Alias: %q, Path: %q},\n", alias, `"`+pkgP+`"`))
	}

	var computeInstancesNbCode strings.Builder
	for _, f := range stageSet.Fields {
		accessor := stageAccessor(f)
		if f.IsLocal {
			computeInstancesNbCode.WriteString(fmt.Sprintf(`	if %s != nil {
		for k, v := range %s.Map_GongStructName_InstancesNb {
			res[k] = v
			res["models."+k] = v
		}
	}
`, accessor, accessor))
		} else {
			computeInstancesNbCode.WriteString(fmt.Sprintf(`	if %s != nil {
		for k, v := range %s.Map_GongStructName_InstancesNb {
			res[k] = v
			res["%s."+k] = v
		}
	}
`, accessor, accessor, f.ImportAlias))
		}
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

	gong_models "github.com/fullstack-lang/gong/go/models"
	gongprobe "github.com/fullstack-lang/gong/pkg/runtime/probe"

	form "github.com/fullstack-lang/gong/lib/form/go/models"
	load "github.com/fullstack-lang/gong/lib/load/go/models"
	{{SplitPkg}} "{{SplitImportPath}}/go/models"
	table "github.com/fullstack-lang/gong/lib/table/go/models"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"

	"{{PkgPathRoot}}/models"
)

type Notification = gongprobe.Notification

const NbNotificationMax = gongprobe.NbNotificationMax

type Probe struct {
	r                      *http.ServeMux
	stageOfInterest        *models.Stage
	gongStage              *gong_models.Stage
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

	updateSliceOfPointersCallback func(instance any, fieldName string, slicePtr any)
}

func (probe *Probe) UpdateSliceOfPointersCallback(instance any, fieldName string, slicePtr any) {
	if probe.updateSliceOfPointersCallback != nil {
		probe.updateSliceOfPointersCallback(instance, fieldName, slicePtr)
	}
}

func (probe *Probe) SetUpdateSliceOfPointersCallback(cb func(instance any, fieldName string, slicePtr any)) {
	probe.updateSliceOfPointersCallback = cb
}

func (probe *Probe) SetCommitMode(commitMode bool) {
	probe.commitMode = commitMode
}

func (probe *Probe) GetCommitMode() bool {
	return probe.commitMode
}

func (probe *Probe) SetMaxElementsNbPerGongStructNode(nb int) {
	probe.maxElementsNbPerGongStructNode = nb
}

func (probe *Probe) GetMaxElementsNbPerGongStructNode() int {
	return probe.maxElementsNbPerGongStructNode
}

func (probe *Probe) RefreshNavigationTree() {
	probe.ux_navigation_tree()
}

func (probe *Probe) GetProbeLoadStageName() string {
	return probe.stageOfInterest.GetProbeLoadStageName()
}

func NewProbe(
	r *http.ServeMux,
	goModelsDir embed.FS,
	goDiagramsDir embed.FS,
	embeddedDiagrams bool,
	stageOfInterest *models.Stage,
) (probe *Probe) {

	splitStage, _ := {{SplitPkg}}_fullstack.NewStackInstance(r, stageOfInterest.GetProbeSplitStageName())
	splitStage.Commit()

	stageOfInterest.MetaPackageImportPath = "{{PkgPathRoot}}/models"

	stage := gong_models.NewStage(stageOfInterest.GetName())
	gong_models.LoadEmbedded(stage, goModelsDir)

	treeStage, _ := tree_fullstack.NewStackInstance(r, stageOfInterest.GetProbeTreeSidebarStageName())
	treeNavigationStage, _ := tree_fullstack.NewStackInstance(r, stageOfInterest.GetProbeNavigationTreeSidebarStageName())

	tableStage, _ := table_fullstack.NewStackInstance(r, stageOfInterest.GetProbeTableStageName())
	tableStage.Commit()

	notificationTableStage, _ := table_fullstack.NewStackInstance(r, stageOfInterest.GetProbeNotificationTableStageName())
	notificationTableStage.Commit()

	formStage, _ := form_fullstack.NewStackInstance(r, stageOfInterest.GetProbeFormStageName())
	formStage.Commit()

	loadStage, _ := load_fullstack.NewStackInstance(r, stageOfInterest.GetProbeLoadStageName())

	probe = &Probe{
		r:                              r,
		stageOfInterest:                stageOfInterest,
		gongStage:                      stage,
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
		"{{PkgPathRoot}}:"+stageOfInterest.GetName(),
		metaPackageImports,
		goModelsDir,
		goDiagramsDir,
		probe.diagramEditor,
		probe.ComputeInstancesNb(),
	)

	probe.dataEditor = gongprobe.CreateDataEditorLayout{{PrepareStageSetSuffix}}(
		probe.treeNavigationStage.GetName(),
		probe.treeStage.GetName(),
		probe.loadStage.GetName(),
		probe.tableStage.GetName(),
		probe.notificationTableStage.GetName(),
		probe.formStage.GetName(),
	)

	probe.splitStage.StageBranch(&{{SplitPkg}}.View{
		Name: "Main view",
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

func (probe *Probe) GetDocStager() *doc_models.Stager {
	return probe.docStager
}

func (probe *Probe) GetStage() *models.Stage {
	return probe.stageOfInterest
}

func (probe *Probe) GetSplitStage() *{{SplitPkg}}.Stage {
	return probe.splitStage
}

func (probe *Probe) GetTreeStage() *tree.Stage {
	return probe.treeStage
}

func (probe *Probe) GetFormStage() *form.Stage {
	return probe.formStage
}

func (probe *Probe) GetTableStage() *table.Stage {
	return probe.tableStage
}

func (probe *Probe) GetNavigationTreeStage() *tree.Stage {
	return probe.treeNavigationStage
}

func (probe *Probe) GetDataEditor() *{{SplitPkg}}.AsSplit {
	return probe.dataEditor
}

func (probe *Probe) GetDiagramEditor() *{{SplitPkg}}.AsSplitArea {
	return probe.diagramEditor
}

func (probe *Probe) ComputeInstancesNb() map[string]int {
	probe.stageOfInterest.ComputeInstancesNb()
	res := make(map[string]int)
{{ComputeInstancesNbCode}}	return res
}

func (probe *Probe) Refresh() {
	if probe.docStager != nil {
		probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
		probe.docStager.Svg()
	}
	probe.ux_tree()
	probe.ux_table()
	probe.ux_form()
}

func (probe *Probe) AddNotification(date time.Time, message string) {
	gongprobe.AddNotification(&probe.notification, date, message)
}

func (probe *Probe) CommitNotificationTable() {
	probe.UpdateAndCommitNotificationTable()
}

func (probe *Probe) ResetNotifications() {
	probe.notification = make([]*Notification, 0)
	probe.UpdateAndCommitNotificationTable()
}

func (probe *Probe) FillUpFormFromGongstruct(instance any, formName string) {
	FillUpFormFromGongstruct(instance, probe)
}

func (probe *Probe) DownloadNotificationsCSV() {
	gongprobe.DownloadNotificationsCSV(probe.notification, probe.loadStage, probe.initLoadStage)
}

func (probe *Probe) ExportStageExcel() {
	excelBytes, err := probe.stageOfInterest.SerializeStageAsBytes(false)
	gongprobe.ExportStageExcel(
		excelBytes,
		err,
		probe.fileName,
		"{{ModelPkgLower}}",
		probe.stageOfInterest.GetName(),
		probe.loadStage,
		probe.initLoadStage,
		probe.AddNotification,
		probe.CommitNotificationTable,
	)
}

func (probe *Probe) ExportStage() {
	probe.loadStage.Reset()

	fileToDownload := new(load.FileToDownload)
	if probe.fileName == "" {
		probe.fileName = "{{ModelPkgName}}-stage.go"
	}

	prefixRegex := regexp.MustCompile("^\\d{8} \\d{4} ")
	cleanFileName := prefixRegex.ReplaceAllString(probe.fileName, "")

	fileToDownload.Name = time.Now().Format("20060102 1504 ") + cleanFileName

	stageString, err := probe.stageOfInterest.MarshallToString("main")
	if err != nil {
		probe.AddNotification(time.Now(), "Error serializing stage: "+err.Error())
		probe.CommitNotificationTable()
		return
	}

	fileToDownload.Base64EncodedContent = base64.StdEncoding.EncodeToString([]byte(stageString))
	probe.loadStage.StageBranch(fileToDownload)
	probe.loadStage.Commit()

	time.Sleep(1 * time.Second)
	probe.initLoadStage()
}

type stageLoadProxy struct {
	probe *Probe
}

func (proxy *stageLoadProxy) OnFileUpload(uploadedFile *load.FileToUpload) error {
	proxy.probe.fileName = uploadedFile.GetName()
	decodedBytes, err := base64.StdEncoding.DecodeString(uploadedFile.Base64EncodedContent)
	if err != nil {
		return fmt.Errorf("base64.StdEncoding.DecodeString failed: %w", err)
	}
	proxy.probe.stageOfInterest.Reset()
	err = proxy.probe.stageOfInterest.ParseAstString(string(decodedBytes), true)
	if err != nil {
		return err
	}
	proxy.probe.stageOfInterest.Commit()
	proxy.probe.Refresh()
	return nil
}

func (probe *Probe) initLoadStage() {
	probe.loadStage.Reset()

	fileToUpload := &load.FileToUpload{
		Name: "Name of file",
		FileToUploadProxy: &stageLoadProxy{
			probe: probe,
		},
	}

	probe.loadStage.StageBranch(
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
	code = strings.ReplaceAll(code, "{{ComputeInstancesNbCode}}", computeInstancesNbCode.String())
	code = strings.ReplaceAll(code, "{{ModelPkgName}}", modelPkg.Name)
	code = strings.ReplaceAll(code, "{{ModelPkgLower}}", strings.ToLower(modelPkg.Name))

	writeFile(filepath.Join(pkgPath, "probe/probe.go"), code)
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
			if f.ImportAlias != f.PackageName {
				extImports.WriteString(fmt.Sprintf("\n\t%s \"%s\"", f.ImportAlias, f.PackagePath))
			} else {
				extImports.WriteString(fmt.Sprintf("\n\t\"%s\"", f.PackagePath))
			}
		}
	}

	var localField *models.StageSetField
	var nonLocalFields []*models.StageSetField
	for _, f := range stageSet.Fields {
		if f.IsLocal {
			localField = f
		} else {
			nonLocalFields = append(nonLocalFields, f)
		}
	}

	// Sort nonLocalFields so ancestor packages appear before descendant packages
	sort.SliceStable(nonLocalFields, func(i, j int) bool {
		if strings.HasPrefix(nonLocalFields[j].PackagePath, nonLocalFields[i].PackagePath+"/") {
			return true
		}
		if strings.HasPrefix(nonLocalFields[i].PackagePath, nonLocalFields[j].PackagePath+"/") {
			return false
		}
		return false
	})

	var pkgTreeNodes strings.Builder

	generateStructNodes := func(f *models.StageSetField, parentVarName string) {
		mPkg := fieldToModelPkg[f]
		if mPkg == nil {
			return
		}
		var sNames []string
		for sName, gs := range mPkg.GongStructs {
			if gs.ModelPkg.PkgPath == mPkg.PkgPath && gs.HasNameField() && !gs.IsOmittedForMarshalling {
				sNames = append(sNames, sName[len(mPkg.PkgPath)+1:])
			}
		}
		sort.Strings(sNames)

		accessor := stageAccessor(f)
		for _, sName := range sNames {
			sPlural := sName + "s"
			typeQual := sName
			if !f.IsLocal {
				typeQual = f.ImportAlias + "." + sName
			} else {
				typeQual = "models." + sName
			}

			pkgTreeNodes.WriteString(fmt.Sprintf(`
	{
		count := len(%s.%s)
		nodeGongstruct := &tree_models.Node{
			Name:            fmt.Sprintf("%s (%%d)", count),
			HasToolTip:      true,
			ToolTipText:     "Display table of all %s instances",
			ToolTipPosition: tree_models.Right,
			IsExpanded:      true,
			IsNodeClickable: true,
		}
		%s.Children = append(%s.Children, nodeGongstruct)

		addButton := &tree_models.Button{
			Name:            "%s " + string(tree_buttons.BUTTON_add),
			Icon:            string(tree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of %s",
			ToolTipPosition: tree_models.Right,
			OnClick: func() {
				StageSetNewInstance_%s_%s(probe)
			},
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)

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
		for _, _inst := range %s.GetInstancesByOrder[*%s]() {
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
					FillUpFormFromGongstruct(_captured, probe)
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
`, accessor, sPlural, sName, sName, parentVarName, parentVarName, sName, sName, sName, f.Name, sName, f.Name, accessor, typeQual))
		}
	}

	// 1. Root package (localField): its struct nodes attach directly to topNode (no "models" wrapper node)
	if localField != nil {
		generateStructNodes(localField, "topNode")
	}

	// 2. Non-local fields: package nodes are created and nested according to hierarchy
	for _, f := range nonLocalFields {
		var parentField *models.StageSetField
		longestPrefix := ""
		for _, other := range nonLocalFields {
			if other == f {
				continue
			}
			prefix := other.PackagePath + "/"
			if strings.HasPrefix(f.PackagePath, prefix) && len(prefix) > len(longestPrefix) {
				parentField = other
				longestPrefix = prefix
			}
		}

		parentVar := "topNode"
		if parentField != nil {
			parentVar = "pkgNode_" + parentField.Name
		}

		pkgTreeNodes.WriteString(fmt.Sprintf(`
	// Package node: %s
	pkgNode_%s := &tree_models.Node{
		Name:       "%s",
		IsExpanded: true,
	}
	pkgNode_%s.OnIsExpandedChange = func(isExpanded bool) {
		pkgNode_%s.IsExpanded = isExpanded
	}
	%s.Children = append(%s.Children, pkgNode_%s)
`, f.ImportAlias, f.Name, f.ImportAlias, f.Name, f.Name, parentVar, parentVar, f.Name))

		generateStructNodes(f, "pkgNode_"+f.Name)
	}

	code := fmt.Sprintf(`// generated code - do not edit
package probe

import (
	"fmt"

	tree_buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree_models "github.com/fullstack-lang/gong/lib/tree/go/models"
	gongprobe "github.com/fullstack-lang/gong/pkg/runtime/probe"

	"%s/models"%s
)

func (probe *Probe) ux_navigation_tree() {
	stageOfInterest := probe.stageOfInterest
	if !stageOfInterest.IsInDeltaMode() {
		return
	}
	probe.treeNavigationStage.Reset()
	sidebar := &tree_models.Tree{Name: "Sidebar"}
	probe.AddCommitNavigationNode(func(node models.GongNodeIF) {
		sidebar.RootNodes = append(sidebar.RootNodes, node.(*tree_models.Node))
	})
	probe.treeNavigationStage.StageBranch(sidebar)
	probe.treeNavigationStage.Commit()
}

func (probe *Probe) AddCommitNavigationNode(appendChildrenNodeFunc func(models.GongNodeIF)) {
	gongprobe.AddCommitNavigationNode(
		probe.stageOfInterest,
		probe.Refresh,
		probe.AddNotification,
		probe.CommitNotificationTable,
		func(node any) {
			appendChildrenNodeFunc(node.(models.GongNodeIF))
		},
	)
}

func (probe *Probe) ux_tree() {
	probe.ux_navigation_tree()
	probe.treeStage.Reset()

	sidebar := &tree_models.Tree{Name: "Sidebar"}
	topNode := &tree_models.Node{
		Name:       probe.stageOfInterest.GetName(),
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

	exportExcelButton := &tree_models.Button{
		Name:            "ExportExcelButton " + string(tree_buttons.BUTTON_download),
		Icon:            string(tree_buttons.BUTTON_download),
		HasToolTip:      true,
		ToolTipText:     "Export XL",
		ToolTipPosition: tree_models.Below,
		OnClick: func() {
			probe.ExportStageExcel()
		},
	}
	topNode.Buttons = append(topNode.Buttons, exportExcelButton)

	resetButton := &tree_models.Button{
		Name:            "ResetButton " + string(tree_buttons.BUTTON_reset_tv),
		Icon:            string(tree_buttons.BUTTON_reset_tv),
		HasToolTip:      true,
		ToolTipText:     "Reset stage",
		ToolTipPosition: tree_models.Below,
		OnClick: func() {
			probe.stageOfInterest.Reset()
			probe.stageOfInterest.Commit()
			probe.Refresh()
		},
	}
	topNode.Buttons = append(topNode.Buttons, resetButton)

	exportGoButton := &tree_models.Button{
		Name:            "ExportGoButton " + string(tree_buttons.BUTTON_file_download),
		Icon:            string(tree_buttons.BUTTON_file_download),
		HasToolTip:      true,
		ToolTipText:     "Export Go code for stage",
		ToolTipPosition: tree_models.Below,
		OnClick: func() {
			probe.ExportStage()
		},
	}
	topNode.Buttons = append(topNode.Buttons, exportGoButton)
%s
	probe.treeStage.StageBranch(sidebar)
	probe.treeStage.Commit()
}
`, pkgPathRoot, extImports.String(), pkgTreeNodes.String())

	writeFile(filepath.Join(pkgPath, "probe/probe_ux_tree.go"), code)
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
			if f.ImportAlias != f.PackageName {
				extImports.WriteString(fmt.Sprintf("\n\t%s \"%s\"", f.ImportAlias, f.PackagePath))
			} else {
				extImports.WriteString(fmt.Sprintf("\n\t\"%s\"", f.PackagePath))
			}
		}
	}

	var tableSwitchCases strings.Builder
	var tableFunctions strings.Builder

	for _, si := range allStructs {
		tblName := si.typeQual
		if si.pkgField.IsLocal {
			tableSwitchCases.WriteString(fmt.Sprintf("\tcase \"%s\", \"%s\":\n\t\tupdateStageSetTable_%s_%s(probe)\n", tblName, si.structName, si.structName, si.pkgField.Name))
		} else {
			tableSwitchCases.WriteString(fmt.Sprintf("\tcase \"%s\":\n\t\tupdateStageSetTable_%s_%s(probe)\n", tblName, si.structName, si.pkgField.Name))
		}

		// Find reverse references: any struct in allStructs that has a pointer to si
		var reversePointers []struct {
			sourceStruct structInfo
			fieldName    string
			isSlice      bool
		}
		for _, other := range allStructs {
			for _, fld := range other.fields {
				if (fld.IsPointer || fld.IsSliceOfPointer) && fld.TargetPackagePath == si.pkgPath && fld.TargetStructName == si.structName {
					reversePointers = append(reversePointers, struct {
						sourceStruct structInfo
						fieldName    string
						isSlice      bool
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
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.%s}
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
			revAccessor := stageAccessor(rev.sourceStruct.pkgField)
			if rev.isSlice {
				cellAssignments.WriteString(fmt.Sprintf(`
		{
			cell := &table_models.Cell{Name: "%s"}
			var refNames []string
			for src := range %s.%ss {
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
`, colName, revAccessor, rev.sourceStruct.structName, rev.fieldName))
			} else {
				cellAssignments.WriteString(fmt.Sprintf(`
		{
			cell := &table_models.Cell{Name: "%s"}
			var refNames []string
			for src := range %s.%ss {
				if src.%s == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}
`, colName, revAccessor, rev.sourceStruct.structName, rev.fieldName))
			}
		}

		siAccessor := stageAccessor(si.pkgField)
		tableFunctions.WriteString(fmt.Sprintf(`
func updateStageSetTable_%s_%s(probe *Probe) {
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
	instances := %s.GetInstancesByOrder[*%s]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(%s.GetOrder(structInstance))}
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
				_captured.UnstageVoid(%s)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_%s_%s(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		cellDel.CellIcon = cellIcon
		row.Cells = append(row.Cells, cellDel)

%s
		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}
`, si.structName, si.pkgField.Name, si.typeQual, columnDefs.String(), siAccessor, si.typeQual, siAccessor, siAccessor, si.structName, si.pkgField.Name, cellAssignments.String()))
	}

	code := fmt.Sprintf(`// generated code - do not edit
package probe

import (
	"fmt"
	"sort"
	"strings"
	"time"

	table_models "github.com/fullstack-lang/gong/lib/table/go/models"
	tree_buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	maticons "github.com/fullstack-lang/maticons/maticons"

	"%s/models"%s
)

const TableName = "Table"

var (
	_ = fmt.Sprintf
	_ = sort.Strings
	_ = strings.Join
	_ = maticons.BUTTON_delete
)

type tableRowUpdater struct {
	onClick func()
}

func (u *tableRowUpdater) RowUpdated(stage *table_models.Stage, row, updatedRow *table_models.Row) {
	if u.onClick != nil {
		u.onClick()
	}
}

func (probe *Probe) ux_table() {
	var tableName string
	for tbl := range probe.tableStage.Tables {
		tableName = tbl.Name
	}
	switch tableName {
%s	}
}

func (probe *Probe) UpdateAndCommitNotificationTable() {
	probe.notificationTableStage.Reset()

	table := new(table_models.Table)
	table.Name = TableName
	table.HasColumnSorting = true
	table.HasFiltering = true
	table.HasPaginator = true
	table.HasCheckableRows = false
	table.HasSaveButton = false

	notificationsResetButton := &table_models.Button{
		Name:            "NotificationsResetButton",
		Icon:            string(tree_buttons.BUTTON_playlist_remove),
		HasToolTip:      true,
		ToolTipText:     "Reset notification table",
		ToolTipPosition: table_models.Below,
		OnClick: func() {
			probe.ResetNotifications()
		},
	}
	table.Buttons = append(table.Buttons, notificationsResetButton)

	notificationsDownloadCSVButton := &table_models.Button{
		Name:            "NotificationsDownloadCSVButton",
		Icon:            string(tree_buttons.BUTTON_download),
		HasToolTip:      true,
		ToolTipText:     "Download notifications as CSV",
		ToolTipPosition: table_models.Below,
		OnClick: func() {
			probe.DownloadNotificationsCSV()
		},
	}
	table.Buttons = append(table.Buttons, notificationsDownloadCSVButton)

	column := new(table_models.DisplayedColumn)
	column.Name = "Date"
	table.DisplayedColumns = append(table.DisplayedColumns, column)

	column = new(table_models.DisplayedColumn)
	column.Name = "Message"
	table.DisplayedColumns = append(table.DisplayedColumns, column)

	sort.Slice(probe.notification, func(i, j int) bool {
		return probe.notification[i].Date.After(probe.notification[j].Date)
	})

	for _, notification := range probe.notification {
		row := new(table_models.Row)
		table.Rows = append(table.Rows, row)

		{
			cell := new(table_models.Cell)
			cellString := new(table_models.CellString)
			cell.CellString = cellString
			cellString.Value = notification.Date.Format(time.StampMicro)
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := new(table_models.Cell)
			cellString := new(table_models.CellString)
			cell.CellString = cellString
			cellString.Value = notification.Message
			row.Cells = append(row.Cells, cell)
		}
	}

	probe.notificationTableStage.StageBranch(table)
	probe.notificationTableStage.Commit()
}
%s
`, pkgPathRoot, extImports.String(), tableSwitchCases.String(), tableFunctions.String())

	writeFile(filepath.Join(pkgPath, "probe/probe_ux_table.go"), code)
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
			if f.ImportAlias != f.PackageName {
				extImports.WriteString(fmt.Sprintf("\n\t%s \"%s\"", f.ImportAlias, f.PackagePath))
			} else {
				extImports.WriteString(fmt.Sprintf("\n\t\"%s\"", f.PackagePath))
			}
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
						targetTypeQual = targetSSF.ImportAlias + "." + fld.TargetStructName
					} else {
						targetTypeQual = "models." + fld.TargetStructName
					}
					cases.WriteString(fmt.Sprintf("\t\tStageSetAssociationFieldToForm(\"%s\", inst.%s, formGroup, %s.GetInstancesSet[*%s](), probe.formStage)\n", fld.Name, fld.Name, stageAccessor(targetSSF), targetTypeQual))
				}
			} else if fld.IsSliceOfPointer {
				targetSSF := pkgPathToField[fld.TargetPackagePath]
				if targetSSF != nil {
					if si.pkgField.IsLocal && targetSSF.IsLocal {
						cases.WriteString(fmt.Sprintf("\t\tAssociationSliceToForm(\"%s\", inst, &inst.%s, formGroup, probe)\n", fld.Name, fld.Name))
					} else {
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
				}
			} else if fld.IsEnum {
				if fld.BasicKind == types.Int || fld.BasicKind == types.Int8 || fld.BasicKind == types.Int16 || fld.BasicKind == types.Int32 || fld.BasicKind == types.Int64 {
					cases.WriteString(fmt.Sprintf("\t\tStageSetEnumIntFieldToForm(\"%s\", inst.%s, formGroup, probe.formStage)\n", fld.Name, fld.Name))
				} else {
					cases.WriteString(fmt.Sprintf("\t\tStageSetEnumStringFieldToForm(\"%s\", inst.%s, formGroup, probe.formStage)\n", fld.Name, fld.Name))
				}
			} else {
				cases.WriteString(fmt.Sprintf("\t\tStageSetBasicFieldtoForm(\"%s\", inst.%s, probe.formStage, formGroup)\n", fld.Name, fld.Name))
			}
		}

		// Reverse pointers pointing to this struct
		for _, other := range allStructs {
			otherAccessor := stageAccessor(other.pkgField)
			for _, fld := range other.fields {
				if fld.IsPointer && fld.TargetPackagePath == si.pkgPath && fld.TargetStructName == si.structName {
					cases.WriteString(fmt.Sprintf(`
		{
			var refNames []string
			for src := range %s.%ss {
				if src.%s == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("%s", "%s", refNames, formGroup, probe.formStage)
		}
`, otherAccessor, other.structName, fld.Name, other.typeQual, fld.Name))
				}
				if fld.IsSliceOfPointer && fld.TargetPackagePath == si.pkgPath && fld.TargetStructName == si.structName {
					cases.WriteString(fmt.Sprintf(`
		{
			var refNames []string
			for src := range %s.%ss {
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
`, otherAccessor, other.structName, fld.Name, other.typeQual, fld.Name))
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

const FormName = "Form"

var (
	_ = sort.Strings
	_ = strings.Join
)

func FillUpForm(
	instance any,
	formGroup *form.FormGroup,
	probe *Probe,
) {
	switch inst := instance.(type) {
%s	default:
		_ = inst
	}
}
`, pkgPathRoot, extImports.String(), cases.String())

	writeFile(filepath.Join(pkgPath, "probe/fill_up_form.go"), code)
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
			if f.ImportAlias != f.PackageName {
				extImports.WriteString(fmt.Sprintf("\n\t%s \"%s\"", f.ImportAlias, f.PackagePath))
			} else {
				extImports.WriteString(fmt.Sprintf("\n\t\"%s\"", f.PackagePath))
			}
		}
	}

	var fillUpDispatchCases strings.Builder
	var fillUpByNameCases strings.Builder
	var saveFunctions strings.Builder
	var newInstanceFunctions strings.Builder

	for _, si := range allStructs {
		siAccessor := stageAccessor(si.pkgField)
		fillUpDispatchCases.WriteString(fmt.Sprintf(`	case *%s:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     inst.GetName(),
			TypeLabel: "%s",
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		formGroup.OnSave = &functionalStageSetFormCallback{
			onSave: func() {
				%s.Lock()
				defer %s.Unlock()
				probe.formStage.Checkout()
				saveStageSet_%s_%s(inst, probe, formGroup)
				if formGroup.HasSuppressButtonBeenPressed {
					inst.UnstageVoid(%s)
				}
				%s.Commit()
				updateStageSetTable_%s_%s(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		FillUpForm(inst, formGroup, probe)
`, si.typeQual, si.structName, siAccessor, siAccessor, si.structName, si.pkgField.Name, siAccessor, siAccessor, si.structName, si.pkgField.Name))

		if si.pkgField.IsLocal {
			fillUpByNameCases.WriteString(fmt.Sprintf(`	case "%s", "%s":
		StageSetNewInstance_%s_%s(probe)
`, si.structName, si.typeQual, si.structName, si.pkgField.Name))
		} else {
			fillUpByNameCases.WriteString(fmt.Sprintf(`	case "%s":
		StageSetNewInstance_%s_%s(probe)
`, si.typeQual, si.structName, si.pkgField.Name))
		}

		var fieldSaves strings.Builder
		for _, fld := range si.fields {
			if fld.IsPointer {
				targetSSF := pkgPathToField[fld.TargetPackagePath]
				if targetSSF != nil {
					targetTypeQual := fld.TargetStructName
					if !targetSSF.IsLocal {
						targetTypeQual = targetSSF.ImportAlias + "." + fld.TargetStructName
					} else {
						targetTypeQual = "models." + fld.TargetStructName
					}
					fieldSaves.WriteString(fmt.Sprintf("\t\tcase \"%s\":\n\t\t\tStageSetFormDivSelectFieldToField(&inst.%s, %s.GetInstancesSet[*%s](), formDiv)\n", fld.Name, fld.Name, stageAccessor(targetSSF), targetTypeQual))
				}
			} else if fld.IsSliceOfPointer {
				targetSSF := pkgPathToField[fld.TargetPackagePath]
				if targetSSF != nil && si.pkgField.IsLocal && targetSSF.IsLocal {
					fieldSaves.WriteString(fmt.Sprintf("\t\tcase \"%s\":\n\t\t\tFormDivSliceOfPointersToField(inst, \"%s\", &inst.%s, formDiv, probe)\n", fld.Name, fld.Name, fld.Name))
				}
			} else if fld.IsEnum {
				if fld.BasicKind == types.Int || fld.BasicKind == types.Int8 || fld.BasicKind == types.Int16 || fld.BasicKind == types.Int32 || fld.BasicKind == types.Int64 {
					fieldSaves.WriteString(fmt.Sprintf("\t\tcase \"%s\":\n\t\t\tStageSetFormDivEnumIntFieldToField(&inst.%s, formDiv)\n", fld.Name, fld.Name))
				} else {
					fieldSaves.WriteString(fmt.Sprintf("\t\tcase \"%s\":\n\t\t\tStageSetFormDivEnumStringFieldToField(&inst.%s, formDiv)\n", fld.Name, fld.Name))
				}
			} else if fld.IsTime {
				fieldSaves.WriteString(fmt.Sprintf("\t\tcase \"%s\":\n\t\t\tFormDivTimeFieldToField(&inst.%s, formDiv, false)\n", fld.Name, fld.Name))
			} else {
				fieldSaves.WriteString(fmt.Sprintf("\t\tcase \"%s\":\n\t\t\tFormDivBasicFieldToField(&inst.%s, formDiv)\n", fld.Name, fld.Name))
			}
		}

		saveFunctions.WriteString(fmt.Sprintf(`
func saveStageSet_%s_%s(
	inst *%s,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
%s		}
	}
}
`, si.structName, si.pkgField.Name, si.typeQual, fieldSaves.String()))

		newInstanceFunctions.WriteString(fmt.Sprintf(`
func StageSetNewInstance_%s_%s(probe *Probe) {
	probe.formStage.Reset()
	formGroup := (&form.FormGroup{
		Name:  FormName,
		Label: "New %s",
	}).Stage(probe.formStage)
	inst := new(%s)
	formGroup.HasSuppressButton = false
	formGroup.OnSave = &functionalStageSetFormCallback{
		onSave: func() {
			%s.Lock()
			defer %s.Unlock()
			probe.formStage.Checkout()
			inst.Stage(%s)
			saveStageSet_%s_%s(inst, probe, formGroup)
			%s.Commit()
			updateStageSetTable_%s_%s(probe)
			probe.ux_tree()
			if probe.docStager != nil {
				probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
				probe.docStager.Svg()
			}
			FillUpFormFromGongstruct(inst, probe)
		},
	}
	FillUpForm(inst, formGroup, probe)
	probe.formStage.Commit()
}
`, si.structName, si.pkgField.Name, si.structName, si.typeQual, siAccessor, siAccessor, siAccessor, si.structName, si.pkgField.Name, siAccessor, si.structName, si.pkgField.Name))
	}

	fillUpFromGongstructCode := fmt.Sprintf(`// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"%s/models"%s
)

func FillUpFormFromGongstruct(instance any, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()

	FillUpNamedFormFromGongstruct(instance, probe, formStage, FormName)
}

func FillUpNamedFormFromGongstruct(instance any, probe *Probe, formStage *form.Stage, formName string) {
	switch inst := instance.(type) {
%s	default:
		_ = inst
	}

	if probe.GetCommitMode() {
		formStage.Commit()
	}
}
`, pkgPathRoot, extImports.String(), fillUpDispatchCases.String())

	writeFile(filepath.Join(pkgPath, "probe/fill_up_form_from_gongstruct.go"), fillUpFromGongstructCode)

	probeUxFormCode := fmt.Sprintf(`// generated code - do not edit
package probe

func (probe *Probe) ux_form() {
}

func FillUpFormFromGongstructName(
	probe *Probe,
	gongstructName string,
	isNewInstance bool,
) {
	_ = isNewInstance
	switch gongstructName {
%s	}
}
`, fillUpByNameCases.String())

	writeFile(filepath.Join(pkgPath, "probe/probe_ux_form.go"), probeUxFormCode)

	formCallbackCode := fmt.Sprintf(`// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"
	gongprobe "github.com/fullstack-lang/gong/pkg/runtime/probe"

	"%s/models"%s
)

type FormCallbackIF = gongprobe.FormCallbackIF

type functionalStageSetFormCallback struct {
	onSave func()
}

func (cb *functionalStageSetFormCallback) OnSave() {
	if cb.onSave != nil {
		cb.onSave()
	}
}
%s
%s
`, pkgPathRoot, extImports.String(), saveFunctions.String(), newInstanceFunctions.String())

	writeFile(filepath.Join(pkgPath, "probe/form_callback.go"), formCallbackCode)
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
