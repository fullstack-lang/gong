// generated code - do not edit
package models

import (
	"cmp"
	"embed"
	"errors"
	"fmt"
	"log"
	"math"
	"slices"
	"sort"
	"strings"
	"time"

	docx_go "github.com/fullstack-lang/gong/pkg/docx/go"
)

// can be used for
//
//	days := __Gong__Abs(int(int(inferedInstance.ComputedDuration.Hours()) / 24))
func __Gong__Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

var _ = __Gong__Abs
var _ = strings.Clone("")

const ProbeTreeSidebarSuffix = ":sidebar of the probe"
const ProbeTableSuffix = ":table of the probe"
const ProbeNotificationTableSuffix = ":notification table of the probe"
const ProbeFormSuffix = ":form of the probe"
const ProbeSplitSuffix = ":probe of the probe"

func (stage *Stage) GetProbeTreeSidebarStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeTreeSidebarSuffix
}

func (stage *Stage) GetProbeFormStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeFormSuffix
}

func (stage *Stage) GetProbeTableStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeTableSuffix
}

func (stage *Stage) GetProbeNotificationTableStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeNotificationTableSuffix
}

func (stage *Stage) GetProbeSplitStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeSplitSuffix
}

// errUnkownEnum is returns when a value cannot match enum values
var errUnkownEnum = errors.New("unkown enum")
var _ = errUnkownEnum

// needed to avoid when fmt package is not needed by generated code
var __dummy__fmt_variable fmt.Scanner

var _ = __dummy__fmt_variable

// idem for math package when not need by generated code
var __dummy_math_variable = math.E

var _ = __dummy_math_variable

// swagger:ignore
type __void any

// needed for creating set of instances in the stage
var __member __void
var _ = __member

// GongStructInterface is the interface met by GongStructs
// It allows runtime reflexion of instances (without the hassle of the "reflect" package)
type GongStructInterface interface {
	GetName() (res string)
	// GetID() (res int)
	// GetFields() (res []string)
	// GetFieldStringValue(fieldName string) (res string)
	GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error
	GongGetGongstructName() string
}

// Stage enables storage of staged instances
type Stage struct {
	name string

	// isInDeltaMode is true when the stage is used to compute difference between
	// succesive commit
	isInDeltaMode bool

	// insertion point for definition of arrays registering instances
	Bodys                map[*Body]struct{}
	Bodys_reference      map[*Body]*Body
	Bodys_referenceOrder map[*Body]uint // diff Unstage needs the reference order
	Bodys_mapString      map[string]*Body

	// insertion point for slice of pointers maps
	Body_Paragraphs_reverseMap map[*Paragraph]*Body

	Body_Tables_reverseMap map[*Table]*Body

	OnAfterBodyCreateCallback OnAfterCreateInterface[Body]
	OnAfterBodyUpdateCallback OnAfterUpdateInterface[Body]
	OnAfterBodyDeleteCallback OnAfterDeleteInterface[Body]
	OnAfterBodyReadCallback   OnAfterReadInterface[Body]

	Documents                map[*Document]struct{}
	Documents_reference      map[*Document]*Document
	Documents_referenceOrder map[*Document]uint // diff Unstage needs the reference order
	Documents_mapString      map[string]*Document

	// insertion point for slice of pointers maps
	OnAfterDocumentCreateCallback OnAfterCreateInterface[Document]
	OnAfterDocumentUpdateCallback OnAfterUpdateInterface[Document]
	OnAfterDocumentDeleteCallback OnAfterDeleteInterface[Document]
	OnAfterDocumentReadCallback   OnAfterReadInterface[Document]

	Docxs                map[*Docx]struct{}
	Docxs_reference      map[*Docx]*Docx
	Docxs_referenceOrder map[*Docx]uint // diff Unstage needs the reference order
	Docxs_mapString      map[string]*Docx

	// insertion point for slice of pointers maps
	Docx_Files_reverseMap map[*File]*Docx

	OnAfterDocxCreateCallback OnAfterCreateInterface[Docx]
	OnAfterDocxUpdateCallback OnAfterUpdateInterface[Docx]
	OnAfterDocxDeleteCallback OnAfterDeleteInterface[Docx]
	OnAfterDocxReadCallback   OnAfterReadInterface[Docx]

	Files                map[*File]struct{}
	Files_reference      map[*File]*File
	Files_referenceOrder map[*File]uint // diff Unstage needs the reference order
	Files_mapString      map[string]*File

	// insertion point for slice of pointers maps
	OnAfterFileCreateCallback OnAfterCreateInterface[File]
	OnAfterFileUpdateCallback OnAfterUpdateInterface[File]
	OnAfterFileDeleteCallback OnAfterDeleteInterface[File]
	OnAfterFileReadCallback   OnAfterReadInterface[File]

	Nodes                map[*Node]struct{}
	Nodes_reference      map[*Node]*Node
	Nodes_referenceOrder map[*Node]uint // diff Unstage needs the reference order
	Nodes_mapString      map[string]*Node

	// insertion point for slice of pointers maps
	Node_Nodes_reverseMap map[*Node]*Node

	OnAfterNodeCreateCallback OnAfterCreateInterface[Node]
	OnAfterNodeUpdateCallback OnAfterUpdateInterface[Node]
	OnAfterNodeDeleteCallback OnAfterDeleteInterface[Node]
	OnAfterNodeReadCallback   OnAfterReadInterface[Node]

	Paragraphs                map[*Paragraph]struct{}
	Paragraphs_reference      map[*Paragraph]*Paragraph
	Paragraphs_referenceOrder map[*Paragraph]uint // diff Unstage needs the reference order
	Paragraphs_mapString      map[string]*Paragraph

	// insertion point for slice of pointers maps
	Paragraph_Runes_reverseMap map[*Rune]*Paragraph

	OnAfterParagraphCreateCallback OnAfterCreateInterface[Paragraph]
	OnAfterParagraphUpdateCallback OnAfterUpdateInterface[Paragraph]
	OnAfterParagraphDeleteCallback OnAfterDeleteInterface[Paragraph]
	OnAfterParagraphReadCallback   OnAfterReadInterface[Paragraph]

	ParagraphPropertiess                map[*ParagraphProperties]struct{}
	ParagraphPropertiess_reference      map[*ParagraphProperties]*ParagraphProperties
	ParagraphPropertiess_referenceOrder map[*ParagraphProperties]uint // diff Unstage needs the reference order
	ParagraphPropertiess_mapString      map[string]*ParagraphProperties

	// insertion point for slice of pointers maps
	OnAfterParagraphPropertiesCreateCallback OnAfterCreateInterface[ParagraphProperties]
	OnAfterParagraphPropertiesUpdateCallback OnAfterUpdateInterface[ParagraphProperties]
	OnAfterParagraphPropertiesDeleteCallback OnAfterDeleteInterface[ParagraphProperties]
	OnAfterParagraphPropertiesReadCallback   OnAfterReadInterface[ParagraphProperties]

	ParagraphStyles                map[*ParagraphStyle]struct{}
	ParagraphStyles_reference      map[*ParagraphStyle]*ParagraphStyle
	ParagraphStyles_referenceOrder map[*ParagraphStyle]uint // diff Unstage needs the reference order
	ParagraphStyles_mapString      map[string]*ParagraphStyle

	// insertion point for slice of pointers maps
	OnAfterParagraphStyleCreateCallback OnAfterCreateInterface[ParagraphStyle]
	OnAfterParagraphStyleUpdateCallback OnAfterUpdateInterface[ParagraphStyle]
	OnAfterParagraphStyleDeleteCallback OnAfterDeleteInterface[ParagraphStyle]
	OnAfterParagraphStyleReadCallback   OnAfterReadInterface[ParagraphStyle]

	Runes                map[*Rune]struct{}
	Runes_reference      map[*Rune]*Rune
	Runes_referenceOrder map[*Rune]uint // diff Unstage needs the reference order
	Runes_mapString      map[string]*Rune

	// insertion point for slice of pointers maps
	OnAfterRuneCreateCallback OnAfterCreateInterface[Rune]
	OnAfterRuneUpdateCallback OnAfterUpdateInterface[Rune]
	OnAfterRuneDeleteCallback OnAfterDeleteInterface[Rune]
	OnAfterRuneReadCallback   OnAfterReadInterface[Rune]

	RunePropertiess                map[*RuneProperties]struct{}
	RunePropertiess_reference      map[*RuneProperties]*RuneProperties
	RunePropertiess_referenceOrder map[*RuneProperties]uint // diff Unstage needs the reference order
	RunePropertiess_mapString      map[string]*RuneProperties

	// insertion point for slice of pointers maps
	OnAfterRunePropertiesCreateCallback OnAfterCreateInterface[RuneProperties]
	OnAfterRunePropertiesUpdateCallback OnAfterUpdateInterface[RuneProperties]
	OnAfterRunePropertiesDeleteCallback OnAfterDeleteInterface[RuneProperties]
	OnAfterRunePropertiesReadCallback   OnAfterReadInterface[RuneProperties]

	Tables                map[*Table]struct{}
	Tables_reference      map[*Table]*Table
	Tables_referenceOrder map[*Table]uint // diff Unstage needs the reference order
	Tables_mapString      map[string]*Table

	// insertion point for slice of pointers maps
	Table_TableRows_reverseMap map[*TableRow]*Table

	OnAfterTableCreateCallback OnAfterCreateInterface[Table]
	OnAfterTableUpdateCallback OnAfterUpdateInterface[Table]
	OnAfterTableDeleteCallback OnAfterDeleteInterface[Table]
	OnAfterTableReadCallback   OnAfterReadInterface[Table]

	TableColumns                map[*TableColumn]struct{}
	TableColumns_reference      map[*TableColumn]*TableColumn
	TableColumns_referenceOrder map[*TableColumn]uint // diff Unstage needs the reference order
	TableColumns_mapString      map[string]*TableColumn

	// insertion point for slice of pointers maps
	TableColumn_Paragraphs_reverseMap map[*Paragraph]*TableColumn

	OnAfterTableColumnCreateCallback OnAfterCreateInterface[TableColumn]
	OnAfterTableColumnUpdateCallback OnAfterUpdateInterface[TableColumn]
	OnAfterTableColumnDeleteCallback OnAfterDeleteInterface[TableColumn]
	OnAfterTableColumnReadCallback   OnAfterReadInterface[TableColumn]

	TablePropertiess                map[*TableProperties]struct{}
	TablePropertiess_reference      map[*TableProperties]*TableProperties
	TablePropertiess_referenceOrder map[*TableProperties]uint // diff Unstage needs the reference order
	TablePropertiess_mapString      map[string]*TableProperties

	// insertion point for slice of pointers maps
	OnAfterTablePropertiesCreateCallback OnAfterCreateInterface[TableProperties]
	OnAfterTablePropertiesUpdateCallback OnAfterUpdateInterface[TableProperties]
	OnAfterTablePropertiesDeleteCallback OnAfterDeleteInterface[TableProperties]
	OnAfterTablePropertiesReadCallback   OnAfterReadInterface[TableProperties]

	TableRows                map[*TableRow]struct{}
	TableRows_reference      map[*TableRow]*TableRow
	TableRows_referenceOrder map[*TableRow]uint // diff Unstage needs the reference order
	TableRows_mapString      map[string]*TableRow

	// insertion point for slice of pointers maps
	TableRow_TableColumns_reverseMap map[*TableColumn]*TableRow

	OnAfterTableRowCreateCallback OnAfterCreateInterface[TableRow]
	OnAfterTableRowUpdateCallback OnAfterUpdateInterface[TableRow]
	OnAfterTableRowDeleteCallback OnAfterDeleteInterface[TableRow]
	OnAfterTableRowReadCallback   OnAfterReadInterface[TableRow]

	TableStyles                map[*TableStyle]struct{}
	TableStyles_reference      map[*TableStyle]*TableStyle
	TableStyles_referenceOrder map[*TableStyle]uint // diff Unstage needs the reference order
	TableStyles_mapString      map[string]*TableStyle

	// insertion point for slice of pointers maps
	OnAfterTableStyleCreateCallback OnAfterCreateInterface[TableStyle]
	OnAfterTableStyleUpdateCallback OnAfterUpdateInterface[TableStyle]
	OnAfterTableStyleDeleteCallback OnAfterDeleteInterface[TableStyle]
	OnAfterTableStyleReadCallback   OnAfterReadInterface[TableStyle]

	Texts                map[*Text]struct{}
	Texts_reference      map[*Text]*Text
	Texts_referenceOrder map[*Text]uint // diff Unstage needs the reference order
	Texts_mapString      map[string]*Text

	// insertion point for slice of pointers maps
	OnAfterTextCreateCallback OnAfterCreateInterface[Text]
	OnAfterTextUpdateCallback OnAfterUpdateInterface[Text]
	OnAfterTextDeleteCallback OnAfterDeleteInterface[Text]
	OnAfterTextReadCallback   OnAfterReadInterface[Text]

	AllModelsStructCreateCallback AllModelsStructCreateInterface

	AllModelsStructDeleteCallback AllModelsStructDeleteInterface

	BackRepo BackRepoInterface

	// if set will be called before each commit to the back repo
	OnInitCommitCallback          OnInitCommitInterface
	OnInitCommitFromFrontCallback OnInitCommitInterface
	OnInitCommitFromBackCallback  OnInitCommitInterface

	// store the number of instance per gongstruct
	Map_GongStructName_InstancesNb map[string]int

	// store meta package import
	MetaPackageImportPath  string
	MetaPackageImportAlias string

	// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
	// map to enable docLink renaming when an identifier is renamed
	Map_DocLink_Renaming map[string]GONG__Identifier
	// the to be removed stops here

	// store the stage order of each instance in order to
	// preserve this order when serializing them
	// insertion point for order fields declaration
	BodyOrder            uint
	BodyMap_Staged_Order map[*Body]uint

	DocumentOrder            uint
	DocumentMap_Staged_Order map[*Document]uint

	DocxOrder            uint
	DocxMap_Staged_Order map[*Docx]uint

	FileOrder            uint
	FileMap_Staged_Order map[*File]uint

	NodeOrder            uint
	NodeMap_Staged_Order map[*Node]uint

	ParagraphOrder            uint
	ParagraphMap_Staged_Order map[*Paragraph]uint

	ParagraphPropertiesOrder            uint
	ParagraphPropertiesMap_Staged_Order map[*ParagraphProperties]uint

	ParagraphStyleOrder            uint
	ParagraphStyleMap_Staged_Order map[*ParagraphStyle]uint

	RuneOrder            uint
	RuneMap_Staged_Order map[*Rune]uint

	RunePropertiesOrder            uint
	RunePropertiesMap_Staged_Order map[*RuneProperties]uint

	TableOrder            uint
	TableMap_Staged_Order map[*Table]uint

	TableColumnOrder            uint
	TableColumnMap_Staged_Order map[*TableColumn]uint

	TablePropertiesOrder            uint
	TablePropertiesMap_Staged_Order map[*TableProperties]uint

	TableRowOrder            uint
	TableRowMap_Staged_Order map[*TableRow]uint

	TableStyleOrder            uint
	TableStyleMap_Staged_Order map[*TableStyle]uint

	TextOrder            uint
	TextMap_Staged_Order map[*Text]uint

	// end of insertion point

	NamedStructs []*NamedStruct

	// GongUnmarshallers is the registry of all model unmarshallers
	GongUnmarshallers map[string]ModelUnmarshaller

	// probeIF is the interface to the probe that allows log
	// commit event to the probe
	probeIF ProbeIF

	forwardCommits  []string
	backwardCommits []string

	// when navigating the commit history
	// navigationMode is set to Navigating
	navigationMode gongStageNavigationMode
	commitsBehind  int // the number of commits the stage is behind the front of the history
}

type gongStageNavigationMode string

const (
	GongNavigationModeNormal gongStageNavigationMode = "Normal"
	// when the mode is navigating, each commit backward and forward
	// it is possible to go apply the nbCommitsBackward forward commits
	GongNavigationModeNavigating gongStageNavigationMode = "Navigating"
)

// ApplyBackwardCommit applies the commit before the current one
func (stage *Stage) ApplyBackwardCommit() error {

	if len(stage.backwardCommits) == 0 {
		return errors.New("no backward commit to apply")
	}

	if stage.navigationMode == GongNavigationModeNormal && stage.commitsBehind != 0 {
		return errors.New("in navigation mode normal, cannot have commitsBehind != 0")
	}

	if stage.navigationMode == GongNavigationModeNormal {
		stage.navigationMode = GongNavigationModeNavigating
	}

	if stage.commitsBehind >= len(stage.backwardCommits) {
		return errors.New("no more backward commit to apply")
	}

	commitToApply := stage.backwardCommits[len(stage.backwardCommits)-1-stage.commitsBehind]

	// umarshall the backward commit to the stage

	// the parsing of the commit will call the UX update
	// therefore, it is important to stage.commitsBehind before because it is used in the
	// UX
	stage.commitsBehind++
	err := GongParseAstString(stage, commitToApply, true)
	if err != nil {
		log.Println("error during ApplyBackwardCommit: ", err)
		return err
	}

	stage.ComputeReferenceAndOrders()

	return nil
}

func (stage *Stage) GetForwardCommits() []string {
	return stage.forwardCommits
}

func (stage *Stage) GetBackwardCommits() []string {
	return stage.backwardCommits
}

func (stage *Stage) ApplyForwardCommit() error {
	if stage.navigationMode == GongNavigationModeNormal && stage.commitsBehind != 0 {
		return errors.New("in navigation mode normal, cannot have commitsBehind != 0")
	}

	if stage.commitsBehind == 0 {
		return errors.New("no more forward commit to apply")
	}

	if stage.navigationMode == GongNavigationModeNormal {
		stage.navigationMode = GongNavigationModeNavigating
	}

	commitToApply := stage.forwardCommits[len(stage.forwardCommits)-1-stage.commitsBehind+1]

	// the parsing of the commit will call the UX update
	// therefore, it is important to stage.commitsBehind before because it is used in the
	// UX
	stage.commitsBehind--
	err := GongParseAstString(stage, commitToApply, true)
	if err != nil {
		log.Println("error during ApplyForwardCommit: ", err)
		return err
	}
	stage.ComputeReferenceAndOrders()

	return nil
}

func (stage *Stage) GetCommitsBehind() int {
	return stage.commitsBehind
}

// ResetHard removes the more recent
// commitsBehind forward/backward Commits from the
// stage
func (stage *Stage) ResetHard() {

	newCommitsLen := len(stage.forwardCommits) - stage.GetCommitsBehind()

	stage.forwardCommits = stage.forwardCommits[:newCommitsLen]
	stage.backwardCommits = stage.backwardCommits[:newCommitsLen]
	stage.commitsBehind = 0
	stage.navigationMode = GongNavigationModeNormal

	stage.ComputeInstancesNb()
	stage.ComputeReferenceAndOrders()
	if stage.OnInitCommitCallback != nil {
		stage.OnInitCommitCallback.BeforeCommit(stage)
	}
	if stage.OnInitCommitFromBackCallback != nil {
		stage.OnInitCommitFromBackCallback.BeforeCommit(stage)
	}
}

// recomputeOrders recomputes the next order for each struct
// this is necessary because the order might have been incremented
// during the commits that have been discarded
// insertion point for max order recomputation
func (stage *Stage) recomputeOrders() {
	// insertion point for max order recomputation 
	var maxBodyOrder uint
	var foundBody bool
	for _, order := range stage.BodyMap_Staged_Order {
		if !foundBody || order > maxBodyOrder {
			maxBodyOrder = order
			foundBody = true
		}
	}
	if foundBody {
		stage.BodyOrder = maxBodyOrder + 1
	} else {
		stage.BodyOrder = 0
	}

	var maxDocumentOrder uint
	var foundDocument bool
	for _, order := range stage.DocumentMap_Staged_Order {
		if !foundDocument || order > maxDocumentOrder {
			maxDocumentOrder = order
			foundDocument = true
		}
	}
	if foundDocument {
		stage.DocumentOrder = maxDocumentOrder + 1
	} else {
		stage.DocumentOrder = 0
	}

	var maxDocxOrder uint
	var foundDocx bool
	for _, order := range stage.DocxMap_Staged_Order {
		if !foundDocx || order > maxDocxOrder {
			maxDocxOrder = order
			foundDocx = true
		}
	}
	if foundDocx {
		stage.DocxOrder = maxDocxOrder + 1
	} else {
		stage.DocxOrder = 0
	}

	var maxFileOrder uint
	var foundFile bool
	for _, order := range stage.FileMap_Staged_Order {
		if !foundFile || order > maxFileOrder {
			maxFileOrder = order
			foundFile = true
		}
	}
	if foundFile {
		stage.FileOrder = maxFileOrder + 1
	} else {
		stage.FileOrder = 0
	}

	var maxNodeOrder uint
	var foundNode bool
	for _, order := range stage.NodeMap_Staged_Order {
		if !foundNode || order > maxNodeOrder {
			maxNodeOrder = order
			foundNode = true
		}
	}
	if foundNode {
		stage.NodeOrder = maxNodeOrder + 1
	} else {
		stage.NodeOrder = 0
	}

	var maxParagraphOrder uint
	var foundParagraph bool
	for _, order := range stage.ParagraphMap_Staged_Order {
		if !foundParagraph || order > maxParagraphOrder {
			maxParagraphOrder = order
			foundParagraph = true
		}
	}
	if foundParagraph {
		stage.ParagraphOrder = maxParagraphOrder + 1
	} else {
		stage.ParagraphOrder = 0
	}

	var maxParagraphPropertiesOrder uint
	var foundParagraphProperties bool
	for _, order := range stage.ParagraphPropertiesMap_Staged_Order {
		if !foundParagraphProperties || order > maxParagraphPropertiesOrder {
			maxParagraphPropertiesOrder = order
			foundParagraphProperties = true
		}
	}
	if foundParagraphProperties {
		stage.ParagraphPropertiesOrder = maxParagraphPropertiesOrder + 1
	} else {
		stage.ParagraphPropertiesOrder = 0
	}

	var maxParagraphStyleOrder uint
	var foundParagraphStyle bool
	for _, order := range stage.ParagraphStyleMap_Staged_Order {
		if !foundParagraphStyle || order > maxParagraphStyleOrder {
			maxParagraphStyleOrder = order
			foundParagraphStyle = true
		}
	}
	if foundParagraphStyle {
		stage.ParagraphStyleOrder = maxParagraphStyleOrder + 1
	} else {
		stage.ParagraphStyleOrder = 0
	}

	var maxRuneOrder uint
	var foundRune bool
	for _, order := range stage.RuneMap_Staged_Order {
		if !foundRune || order > maxRuneOrder {
			maxRuneOrder = order
			foundRune = true
		}
	}
	if foundRune {
		stage.RuneOrder = maxRuneOrder + 1
	} else {
		stage.RuneOrder = 0
	}

	var maxRunePropertiesOrder uint
	var foundRuneProperties bool
	for _, order := range stage.RunePropertiesMap_Staged_Order {
		if !foundRuneProperties || order > maxRunePropertiesOrder {
			maxRunePropertiesOrder = order
			foundRuneProperties = true
		}
	}
	if foundRuneProperties {
		stage.RunePropertiesOrder = maxRunePropertiesOrder + 1
	} else {
		stage.RunePropertiesOrder = 0
	}

	var maxTableOrder uint
	var foundTable bool
	for _, order := range stage.TableMap_Staged_Order {
		if !foundTable || order > maxTableOrder {
			maxTableOrder = order
			foundTable = true
		}
	}
	if foundTable {
		stage.TableOrder = maxTableOrder + 1
	} else {
		stage.TableOrder = 0
	}

	var maxTableColumnOrder uint
	var foundTableColumn bool
	for _, order := range stage.TableColumnMap_Staged_Order {
		if !foundTableColumn || order > maxTableColumnOrder {
			maxTableColumnOrder = order
			foundTableColumn = true
		}
	}
	if foundTableColumn {
		stage.TableColumnOrder = maxTableColumnOrder + 1
	} else {
		stage.TableColumnOrder = 0
	}

	var maxTablePropertiesOrder uint
	var foundTableProperties bool
	for _, order := range stage.TablePropertiesMap_Staged_Order {
		if !foundTableProperties || order > maxTablePropertiesOrder {
			maxTablePropertiesOrder = order
			foundTableProperties = true
		}
	}
	if foundTableProperties {
		stage.TablePropertiesOrder = maxTablePropertiesOrder + 1
	} else {
		stage.TablePropertiesOrder = 0
	}

	var maxTableRowOrder uint
	var foundTableRow bool
	for _, order := range stage.TableRowMap_Staged_Order {
		if !foundTableRow || order > maxTableRowOrder {
			maxTableRowOrder = order
			foundTableRow = true
		}
	}
	if foundTableRow {
		stage.TableRowOrder = maxTableRowOrder + 1
	} else {
		stage.TableRowOrder = 0
	}

	var maxTableStyleOrder uint
	var foundTableStyle bool
	for _, order := range stage.TableStyleMap_Staged_Order {
		if !foundTableStyle || order > maxTableStyleOrder {
			maxTableStyleOrder = order
			foundTableStyle = true
		}
	}
	if foundTableStyle {
		stage.TableStyleOrder = maxTableStyleOrder + 1
	} else {
		stage.TableStyleOrder = 0
	}

	var maxTextOrder uint
	var foundText bool
	for _, order := range stage.TextMap_Staged_Order {
		if !foundText || order > maxTextOrder {
			maxTextOrder = order
			foundText = true
		}
	}
	if foundText {
		stage.TextOrder = maxTextOrder + 1
	} else {
		stage.TextOrder = 0
	}

}

func (stage *Stage) SetDeltaMode(inDeltaMode bool) {
	stage.isInDeltaMode = inDeltaMode
}

func (stage *Stage) IsInDeltaMode() bool {
	return stage.isInDeltaMode
}

func (stage *Stage) SetProbeIF(probeIF ProbeIF) {
	stage.probeIF = probeIF
}

func (stage *Stage) GetProbeIF() ProbeIF {
	if stage.probeIF == nil {
		return nil
	}

	return stage.probeIF
}

// GetNamedStructs implements models.ProbebStage.
func (stage *Stage) GetNamedStructsNames() (res []string) {

	for _, namedStruct := range stage.NamedStructs {
		res = append(res, namedStruct.name)
	}

	return
}

func GetNamedStructInstances[T PointerToGongstruct](set map[T]struct{}, order map[T]uint) (res []string) {

	orderedSet := []T{}
	for instance := range set {
		orderedSet = append(orderedSet, instance)
	}
	sort.Slice(orderedSet[:], func(i, j int) bool {
		instancei := orderedSet[i]
		instancej := orderedSet[j]
		i_order, oki := order[instancei]
		j_order, okj := order[instancej]
		if !oki || !okj {
			log.Fatalf("GetNamedStructInstances: pointer not found")
		}
		return i_order < j_order
	})

	for _, instance := range orderedSet {
		res = append(res, instance.GetName())
	}

	return
}

func GetStructInstancesByOrderAuto[T PointerToGongstruct](stage *Stage) (res []T) {
	var t T
	switch any(t).(type) {
	// insertion point for case
	case *Body:
		tmp := GetStructInstancesByOrder(stage.Bodys, stage.BodyMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Body implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Document:
		tmp := GetStructInstancesByOrder(stage.Documents, stage.DocumentMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Document implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Docx:
		tmp := GetStructInstancesByOrder(stage.Docxs, stage.DocxMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Docx implements.
			res = append(res, any(v).(T))
		}
		return res
	case *File:
		tmp := GetStructInstancesByOrder(stage.Files, stage.FileMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *File implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Node:
		tmp := GetStructInstancesByOrder(stage.Nodes, stage.NodeMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Node implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Paragraph:
		tmp := GetStructInstancesByOrder(stage.Paragraphs, stage.ParagraphMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Paragraph implements.
			res = append(res, any(v).(T))
		}
		return res
	case *ParagraphProperties:
		tmp := GetStructInstancesByOrder(stage.ParagraphPropertiess, stage.ParagraphPropertiesMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *ParagraphProperties implements.
			res = append(res, any(v).(T))
		}
		return res
	case *ParagraphStyle:
		tmp := GetStructInstancesByOrder(stage.ParagraphStyles, stage.ParagraphStyleMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *ParagraphStyle implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Rune:
		tmp := GetStructInstancesByOrder(stage.Runes, stage.RuneMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Rune implements.
			res = append(res, any(v).(T))
		}
		return res
	case *RuneProperties:
		tmp := GetStructInstancesByOrder(stage.RunePropertiess, stage.RunePropertiesMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *RuneProperties implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Table:
		tmp := GetStructInstancesByOrder(stage.Tables, stage.TableMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Table implements.
			res = append(res, any(v).(T))
		}
		return res
	case *TableColumn:
		tmp := GetStructInstancesByOrder(stage.TableColumns, stage.TableColumnMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *TableColumn implements.
			res = append(res, any(v).(T))
		}
		return res
	case *TableProperties:
		tmp := GetStructInstancesByOrder(stage.TablePropertiess, stage.TablePropertiesMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *TableProperties implements.
			res = append(res, any(v).(T))
		}
		return res
	case *TableRow:
		tmp := GetStructInstancesByOrder(stage.TableRows, stage.TableRowMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *TableRow implements.
			res = append(res, any(v).(T))
		}
		return res
	case *TableStyle:
		tmp := GetStructInstancesByOrder(stage.TableStyles, stage.TableStyleMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *TableStyle implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Text:
		tmp := GetStructInstancesByOrder(stage.Texts, stage.TextMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Text implements.
			res = append(res, any(v).(T))
		}
		return res

	}
	return
}

func GetStructInstancesByOrder[T PointerToGongstruct](set map[T]struct{}, order map[T]uint) (res []T) {

	orderedSet := []T{}
	for instance := range set {
		orderedSet = append(orderedSet, instance)
	}
	sort.Slice(orderedSet[:], func(i, j int) bool {
		instancei := orderedSet[i]
		instancej := orderedSet[j]
		i_order, oki := order[instancei]
		j_order, okj := order[instancej]
		if !oki || !okj {
			log.Fatalf("GetNamedStructInstances: pointer not found")
		}
		return i_order < j_order
	})

	res = append(res, orderedSet...)

	return
}

func (stage *Stage) GetNamedStructNamesByOrder(namedStructName string) (res []string) {

	switch namedStructName {
	// insertion point for case
	case "Body":
		res = GetNamedStructInstances(stage.Bodys, stage.BodyMap_Staged_Order)
	case "Document":
		res = GetNamedStructInstances(stage.Documents, stage.DocumentMap_Staged_Order)
	case "Docx":
		res = GetNamedStructInstances(stage.Docxs, stage.DocxMap_Staged_Order)
	case "File":
		res = GetNamedStructInstances(stage.Files, stage.FileMap_Staged_Order)
	case "Node":
		res = GetNamedStructInstances(stage.Nodes, stage.NodeMap_Staged_Order)
	case "Paragraph":
		res = GetNamedStructInstances(stage.Paragraphs, stage.ParagraphMap_Staged_Order)
	case "ParagraphProperties":
		res = GetNamedStructInstances(stage.ParagraphPropertiess, stage.ParagraphPropertiesMap_Staged_Order)
	case "ParagraphStyle":
		res = GetNamedStructInstances(stage.ParagraphStyles, stage.ParagraphStyleMap_Staged_Order)
	case "Rune":
		res = GetNamedStructInstances(stage.Runes, stage.RuneMap_Staged_Order)
	case "RuneProperties":
		res = GetNamedStructInstances(stage.RunePropertiess, stage.RunePropertiesMap_Staged_Order)
	case "Table":
		res = GetNamedStructInstances(stage.Tables, stage.TableMap_Staged_Order)
	case "TableColumn":
		res = GetNamedStructInstances(stage.TableColumns, stage.TableColumnMap_Staged_Order)
	case "TableProperties":
		res = GetNamedStructInstances(stage.TablePropertiess, stage.TablePropertiesMap_Staged_Order)
	case "TableRow":
		res = GetNamedStructInstances(stage.TableRows, stage.TableRowMap_Staged_Order)
	case "TableStyle":
		res = GetNamedStructInstances(stage.TableStyles, stage.TableStyleMap_Staged_Order)
	case "Text":
		res = GetNamedStructInstances(stage.Texts, stage.TextMap_Staged_Order)
	}

	return
}

type NamedStruct struct {
	name string
}

func (namedStruct *NamedStruct) GetName() string {
	return namedStruct.name
}

func (stage *Stage) GetType() string {
	return "github.com/fullstack-lang/gong/pkg/docx/go/models"
}

func (stage *Stage) GetMap_GongStructName_InstancesNb() map[string]int {
	return stage.Map_GongStructName_InstancesNb
}

func (stage *Stage) GetModelsEmbededDir() embed.FS {
	return docx_go.GoModelsDir
}

func (stage *Stage) GetDigramsEmbededDir() embed.FS {
	return docx_go.GoDiagramsDir
}

type GONG__Identifier struct {
	Ident string
	Type  GONG__ExpressionType
}

type OnInitCommitInterface interface {
	BeforeCommit(stage *Stage)
}

// OnAfterCreateInterface callback when an instance is updated from the front
type OnAfterCreateInterface[Type Gongstruct] interface {
	OnAfterCreate(stage *Stage,
		instance *Type)
}

// OnAfterReadInterface callback when an instance is updated from the front
type OnAfterReadInterface[Type Gongstruct] interface {
	OnAfterRead(stage *Stage,
		instance *Type)
}

// OnAfterUpdateInterface callback when an instance is updated from the front
type OnAfterUpdateInterface[Type Gongstruct] interface {
	OnAfterUpdate(stage *Stage, old, new *Type)
}

// OnAfterDeleteInterface callback when an instance is updated from the front
type OnAfterDeleteInterface[Type Gongstruct] interface {
	OnAfterDelete(stage *Stage,
		staged, front *Type)
}

type BackRepoInterface interface {
	Commit(stage *Stage)
	Checkout(stage *Stage)
	Backup(stage *Stage, dirPath string)
	Restore(stage *Stage, dirPath string)
	BackupXL(stage *Stage, dirPath string)
	RestoreXL(stage *Stage, dirPath string)
	// insertion point for Commit and Checkout signatures
	CommitBody(body *Body)
	CheckoutBody(body *Body)
	CommitDocument(document *Document)
	CheckoutDocument(document *Document)
	CommitDocx(docx *Docx)
	CheckoutDocx(docx *Docx)
	CommitFile(file *File)
	CheckoutFile(file *File)
	CommitNode(node *Node)
	CheckoutNode(node *Node)
	CommitParagraph(paragraph *Paragraph)
	CheckoutParagraph(paragraph *Paragraph)
	CommitParagraphProperties(paragraphproperties *ParagraphProperties)
	CheckoutParagraphProperties(paragraphproperties *ParagraphProperties)
	CommitParagraphStyle(paragraphstyle *ParagraphStyle)
	CheckoutParagraphStyle(paragraphstyle *ParagraphStyle)
	CommitRune(rune *Rune)
	CheckoutRune(rune *Rune)
	CommitRuneProperties(runeproperties *RuneProperties)
	CheckoutRuneProperties(runeproperties *RuneProperties)
	CommitTable(table *Table)
	CheckoutTable(table *Table)
	CommitTableColumn(tablecolumn *TableColumn)
	CheckoutTableColumn(tablecolumn *TableColumn)
	CommitTableProperties(tableproperties *TableProperties)
	CheckoutTableProperties(tableproperties *TableProperties)
	CommitTableRow(tablerow *TableRow)
	CheckoutTableRow(tablerow *TableRow)
	CommitTableStyle(tablestyle *TableStyle)
	CheckoutTableStyle(tablestyle *TableStyle)
	CommitText(text *Text)
	CheckoutText(text *Text)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(name string) (stage *Stage) {

	stage = &Stage{ // insertion point for array initiatialisation
		Bodys:           make(map[*Body]struct{}),
		Bodys_mapString: make(map[string]*Body),

		Documents:           make(map[*Document]struct{}),
		Documents_mapString: make(map[string]*Document),

		Docxs:           make(map[*Docx]struct{}),
		Docxs_mapString: make(map[string]*Docx),

		Files:           make(map[*File]struct{}),
		Files_mapString: make(map[string]*File),

		Nodes:           make(map[*Node]struct{}),
		Nodes_mapString: make(map[string]*Node),

		Paragraphs:           make(map[*Paragraph]struct{}),
		Paragraphs_mapString: make(map[string]*Paragraph),

		ParagraphPropertiess:           make(map[*ParagraphProperties]struct{}),
		ParagraphPropertiess_mapString: make(map[string]*ParagraphProperties),

		ParagraphStyles:           make(map[*ParagraphStyle]struct{}),
		ParagraphStyles_mapString: make(map[string]*ParagraphStyle),

		Runes:           make(map[*Rune]struct{}),
		Runes_mapString: make(map[string]*Rune),

		RunePropertiess:           make(map[*RuneProperties]struct{}),
		RunePropertiess_mapString: make(map[string]*RuneProperties),

		Tables:           make(map[*Table]struct{}),
		Tables_mapString: make(map[string]*Table),

		TableColumns:           make(map[*TableColumn]struct{}),
		TableColumns_mapString: make(map[string]*TableColumn),

		TablePropertiess:           make(map[*TableProperties]struct{}),
		TablePropertiess_mapString: make(map[string]*TableProperties),

		TableRows:           make(map[*TableRow]struct{}),
		TableRows_mapString: make(map[string]*TableRow),

		TableStyles:           make(map[*TableStyle]struct{}),
		TableStyles_mapString: make(map[string]*TableStyle),

		Texts:           make(map[*Text]struct{}),
		Texts_mapString: make(map[string]*Text),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		name: name,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		BodyMap_Staged_Order: make(map[*Body]uint),

		DocumentMap_Staged_Order: make(map[*Document]uint),

		DocxMap_Staged_Order: make(map[*Docx]uint),

		FileMap_Staged_Order: make(map[*File]uint),

		NodeMap_Staged_Order: make(map[*Node]uint),

		ParagraphMap_Staged_Order: make(map[*Paragraph]uint),

		ParagraphPropertiesMap_Staged_Order: make(map[*ParagraphProperties]uint),

		ParagraphStyleMap_Staged_Order: make(map[*ParagraphStyle]uint),

		RuneMap_Staged_Order: make(map[*Rune]uint),

		RunePropertiesMap_Staged_Order: make(map[*RuneProperties]uint),

		TableMap_Staged_Order: make(map[*Table]uint),

		TableColumnMap_Staged_Order: make(map[*TableColumn]uint),

		TablePropertiesMap_Staged_Order: make(map[*TableProperties]uint),

		TableRowMap_Staged_Order: make(map[*TableRow]uint),

		TableStyleMap_Staged_Order: make(map[*TableStyle]uint),

		TextMap_Staged_Order: make(map[*Text]uint),

		// end of insertion point
		GongUnmarshallers: map[string]ModelUnmarshaller{ // insertion point for unmarshallers
			"Body": &BodyUnmarshaller{},

			"Document": &DocumentUnmarshaller{},

			"Docx": &DocxUnmarshaller{},

			"File": &FileUnmarshaller{},

			"Node": &NodeUnmarshaller{},

			"Paragraph": &ParagraphUnmarshaller{},

			"ParagraphProperties": &ParagraphPropertiesUnmarshaller{},

			"ParagraphStyle": &ParagraphStyleUnmarshaller{},

			"Rune": &RuneUnmarshaller{},

			"RuneProperties": &RunePropertiesUnmarshaller{},

			"Table": &TableUnmarshaller{},

			"TableColumn": &TableColumnUnmarshaller{},

			"TableProperties": &TablePropertiesUnmarshaller{},

			"TableRow": &TableRowUnmarshaller{},

			"TableStyle": &TableStyleUnmarshaller{},

			"Text": &TextUnmarshaller{},

			// end of insertion point
		},

		NamedStructs: []*NamedStruct{ // insertion point for order map initialisations
			{name: "Body"},
			{name: "Document"},
			{name: "Docx"},
			{name: "File"},
			{name: "Node"},
			{name: "Paragraph"},
			{name: "ParagraphProperties"},
			{name: "ParagraphStyle"},
			{name: "Rune"},
			{name: "RuneProperties"},
			{name: "Table"},
			{name: "TableColumn"},
			{name: "TableProperties"},
			{name: "TableRow"},
			{name: "TableStyle"},
			{name: "Text"},
		}, // end of insertion point

		navigationMode: GongNavigationModeNormal,
	}

	return
}

func GetOrder[Type Gongstruct](stage *Stage, instance *Type) uint {

	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *Body:
		return stage.BodyMap_Staged_Order[instance]
	case *Document:
		return stage.DocumentMap_Staged_Order[instance]
	case *Docx:
		return stage.DocxMap_Staged_Order[instance]
	case *File:
		return stage.FileMap_Staged_Order[instance]
	case *Node:
		return stage.NodeMap_Staged_Order[instance]
	case *Paragraph:
		return stage.ParagraphMap_Staged_Order[instance]
	case *ParagraphProperties:
		return stage.ParagraphPropertiesMap_Staged_Order[instance]
	case *ParagraphStyle:
		return stage.ParagraphStyleMap_Staged_Order[instance]
	case *Rune:
		return stage.RuneMap_Staged_Order[instance]
	case *RuneProperties:
		return stage.RunePropertiesMap_Staged_Order[instance]
	case *Table:
		return stage.TableMap_Staged_Order[instance]
	case *TableColumn:
		return stage.TableColumnMap_Staged_Order[instance]
	case *TableProperties:
		return stage.TablePropertiesMap_Staged_Order[instance]
	case *TableRow:
		return stage.TableRowMap_Staged_Order[instance]
	case *TableStyle:
		return stage.TableStyleMap_Staged_Order[instance]
	case *Text:
		return stage.TextMap_Staged_Order[instance]
	default:
		return 0 // should not happen
	}
}

func GetOrderPointerGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) uint {

	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *Body:
		return stage.BodyMap_Staged_Order[instance]
	case *Document:
		return stage.DocumentMap_Staged_Order[instance]
	case *Docx:
		return stage.DocxMap_Staged_Order[instance]
	case *File:
		return stage.FileMap_Staged_Order[instance]
	case *Node:
		return stage.NodeMap_Staged_Order[instance]
	case *Paragraph:
		return stage.ParagraphMap_Staged_Order[instance]
	case *ParagraphProperties:
		return stage.ParagraphPropertiesMap_Staged_Order[instance]
	case *ParagraphStyle:
		return stage.ParagraphStyleMap_Staged_Order[instance]
	case *Rune:
		return stage.RuneMap_Staged_Order[instance]
	case *RuneProperties:
		return stage.RunePropertiesMap_Staged_Order[instance]
	case *Table:
		return stage.TableMap_Staged_Order[instance]
	case *TableColumn:
		return stage.TableColumnMap_Staged_Order[instance]
	case *TableProperties:
		return stage.TablePropertiesMap_Staged_Order[instance]
	case *TableRow:
		return stage.TableRowMap_Staged_Order[instance]
	case *TableStyle:
		return stage.TableStyleMap_Staged_Order[instance]
	case *Text:
		return stage.TextMap_Staged_Order[instance]
	default:
		return 0 // should not happen
	}
}

func (stage *Stage) GetName() string {
	return stage.name
}

func (stage *Stage) CommitWithSuspendedCallbacks() {

	tmp := stage.OnInitCommitFromBackCallback
	stage.OnInitCommitFromBackCallback = nil
	stage.Commit()
	stage.OnInitCommitFromBackCallback = tmp
}

func (stage *Stage) Commit() {
	stage.ComputeReverseMaps()

	if stage.OnInitCommitCallback != nil {
		stage.OnInitCommitCallback.BeforeCommit(stage)
	}
	if stage.OnInitCommitFromBackCallback != nil {
		stage.OnInitCommitFromBackCallback.BeforeCommit(stage)
	}

	if stage.BackRepo != nil {
		stage.BackRepo.Commit(stage)
	}
	stage.ComputeInstancesNb()

	// if a commit is applied when in navigation mode
	// this will reset the commits behind and swith the
	// naviagation
	if stage.isInDeltaMode && stage.navigationMode == GongNavigationModeNavigating && stage.GetCommitsBehind() > 0 {
		stage.ResetHard()
	}

	if stage.IsInDeltaMode() {
		stage.ComputeForwardAndBackwardCommits()
		stage.ComputeReferenceAndOrders()
		if stage.GetProbeIF() != nil {
			stage.GetProbeIF().Refresh()
		}
	}
}

func (stage *Stage) ComputeInstancesNb() {
	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["Body"] = len(stage.Bodys)
	stage.Map_GongStructName_InstancesNb["Document"] = len(stage.Documents)
	stage.Map_GongStructName_InstancesNb["Docx"] = len(stage.Docxs)
	stage.Map_GongStructName_InstancesNb["File"] = len(stage.Files)
	stage.Map_GongStructName_InstancesNb["Node"] = len(stage.Nodes)
	stage.Map_GongStructName_InstancesNb["Paragraph"] = len(stage.Paragraphs)
	stage.Map_GongStructName_InstancesNb["ParagraphProperties"] = len(stage.ParagraphPropertiess)
	stage.Map_GongStructName_InstancesNb["ParagraphStyle"] = len(stage.ParagraphStyles)
	stage.Map_GongStructName_InstancesNb["Rune"] = len(stage.Runes)
	stage.Map_GongStructName_InstancesNb["RuneProperties"] = len(stage.RunePropertiess)
	stage.Map_GongStructName_InstancesNb["Table"] = len(stage.Tables)
	stage.Map_GongStructName_InstancesNb["TableColumn"] = len(stage.TableColumns)
	stage.Map_GongStructName_InstancesNb["TableProperties"] = len(stage.TablePropertiess)
	stage.Map_GongStructName_InstancesNb["TableRow"] = len(stage.TableRows)
	stage.Map_GongStructName_InstancesNb["TableStyle"] = len(stage.TableStyles)
	stage.Map_GongStructName_InstancesNb["Text"] = len(stage.Texts)
}

func (stage *Stage) Checkout() {
	if stage.BackRepo != nil {
		stage.BackRepo.Checkout(stage)
	}

	stage.ComputeReverseMaps()
	stage.ComputeInstancesNb()
}

// backup generates backup files in the dirPath
func (stage *Stage) Backup(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Backup(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *Stage) Restore(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Restore(stage, dirPath)
	}
}

// backup generates backup files in the dirPath
func (stage *Stage) BackupXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.BackupXL(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *Stage) RestoreXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.RestoreXL(stage, dirPath)
	}
}

// insertion point for cumulative sub template with model space calls
// Stage puts body to the model stage
func (body *Body) Stage(stage *Stage) *Body {

	if _, ok := stage.Bodys[body]; !ok {
		stage.Bodys[body] = struct{}{}
		stage.BodyMap_Staged_Order[body] = stage.BodyOrder
		stage.BodyOrder++
	}
	stage.Bodys_mapString[body.Name] = body

	return body
}

// StagePreserveOrder puts body to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.BodyOrder
// - update stage.BodyOrder accordingly
func (body *Body) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.Bodys[body]; !ok {
		stage.Bodys[body] = struct{}{}

		if order > stage.BodyOrder {
			stage.BodyOrder = order
		}
		stage.BodyMap_Staged_Order[body] = order
		stage.BodyOrder++
	}
	stage.Bodys_mapString[body.Name] = body
}

// Unstage removes body off the model stage
func (body *Body) Unstage(stage *Stage) *Body {
	delete(stage.Bodys, body)
	delete(stage.BodyMap_Staged_Order, body)
	delete(stage.Bodys_mapString, body.Name)

	return body
}

// UnstageVoid removes body off the model stage
func (body *Body) UnstageVoid(stage *Stage) {
	delete(stage.Bodys, body)
	delete(stage.BodyMap_Staged_Order, body)
	delete(stage.Bodys_mapString, body.Name)
}

// commit body to the back repo (if it is already staged)
func (body *Body) Commit(stage *Stage) *Body {
	if _, ok := stage.Bodys[body]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitBody(body)
		}
	}
	return body
}

func (body *Body) CommitVoid(stage *Stage) {
	body.Commit(stage)
}

func (body *Body) StageVoid(stage *Stage) {
	body.Stage(stage)
}

// Checkout body to the back repo (if it is already staged)
func (body *Body) Checkout(stage *Stage) *Body {
	if _, ok := stage.Bodys[body]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutBody(body)
		}
	}
	return body
}

// for satisfaction of GongStruct interface
func (body *Body) GetName() (res string) {
	return body.Name
}

// for satisfaction of GongStruct interface
func (body *Body) SetName(name string) {
	body.Name = name
}

// Stage puts document to the model stage
func (document *Document) Stage(stage *Stage) *Document {

	if _, ok := stage.Documents[document]; !ok {
		stage.Documents[document] = struct{}{}
		stage.DocumentMap_Staged_Order[document] = stage.DocumentOrder
		stage.DocumentOrder++
	}
	stage.Documents_mapString[document.Name] = document

	return document
}

// StagePreserveOrder puts document to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DocumentOrder
// - update stage.DocumentOrder accordingly
func (document *Document) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.Documents[document]; !ok {
		stage.Documents[document] = struct{}{}

		if order > stage.DocumentOrder {
			stage.DocumentOrder = order
		}
		stage.DocumentMap_Staged_Order[document] = order
		stage.DocumentOrder++
	}
	stage.Documents_mapString[document.Name] = document
}

// Unstage removes document off the model stage
func (document *Document) Unstage(stage *Stage) *Document {
	delete(stage.Documents, document)
	delete(stage.DocumentMap_Staged_Order, document)
	delete(stage.Documents_mapString, document.Name)

	return document
}

// UnstageVoid removes document off the model stage
func (document *Document) UnstageVoid(stage *Stage) {
	delete(stage.Documents, document)
	delete(stage.DocumentMap_Staged_Order, document)
	delete(stage.Documents_mapString, document.Name)
}

// commit document to the back repo (if it is already staged)
func (document *Document) Commit(stage *Stage) *Document {
	if _, ok := stage.Documents[document]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDocument(document)
		}
	}
	return document
}

func (document *Document) CommitVoid(stage *Stage) {
	document.Commit(stage)
}

func (document *Document) StageVoid(stage *Stage) {
	document.Stage(stage)
}

// Checkout document to the back repo (if it is already staged)
func (document *Document) Checkout(stage *Stage) *Document {
	if _, ok := stage.Documents[document]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutDocument(document)
		}
	}
	return document
}

// for satisfaction of GongStruct interface
func (document *Document) GetName() (res string) {
	return document.Name
}

// for satisfaction of GongStruct interface
func (document *Document) SetName(name string) {
	document.Name = name
}

// Stage puts docx to the model stage
func (docx *Docx) Stage(stage *Stage) *Docx {

	if _, ok := stage.Docxs[docx]; !ok {
		stage.Docxs[docx] = struct{}{}
		stage.DocxMap_Staged_Order[docx] = stage.DocxOrder
		stage.DocxOrder++
	}
	stage.Docxs_mapString[docx.Name] = docx

	return docx
}

// StagePreserveOrder puts docx to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DocxOrder
// - update stage.DocxOrder accordingly
func (docx *Docx) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.Docxs[docx]; !ok {
		stage.Docxs[docx] = struct{}{}

		if order > stage.DocxOrder {
			stage.DocxOrder = order
		}
		stage.DocxMap_Staged_Order[docx] = order
		stage.DocxOrder++
	}
	stage.Docxs_mapString[docx.Name] = docx
}

// Unstage removes docx off the model stage
func (docx *Docx) Unstage(stage *Stage) *Docx {
	delete(stage.Docxs, docx)
	delete(stage.DocxMap_Staged_Order, docx)
	delete(stage.Docxs_mapString, docx.Name)

	return docx
}

// UnstageVoid removes docx off the model stage
func (docx *Docx) UnstageVoid(stage *Stage) {
	delete(stage.Docxs, docx)
	delete(stage.DocxMap_Staged_Order, docx)
	delete(stage.Docxs_mapString, docx.Name)
}

// commit docx to the back repo (if it is already staged)
func (docx *Docx) Commit(stage *Stage) *Docx {
	if _, ok := stage.Docxs[docx]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDocx(docx)
		}
	}
	return docx
}

func (docx *Docx) CommitVoid(stage *Stage) {
	docx.Commit(stage)
}

func (docx *Docx) StageVoid(stage *Stage) {
	docx.Stage(stage)
}

// Checkout docx to the back repo (if it is already staged)
func (docx *Docx) Checkout(stage *Stage) *Docx {
	if _, ok := stage.Docxs[docx]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutDocx(docx)
		}
	}
	return docx
}

// for satisfaction of GongStruct interface
func (docx *Docx) GetName() (res string) {
	return docx.Name
}

// for satisfaction of GongStruct interface
func (docx *Docx) SetName(name string) {
	docx.Name = name
}

// Stage puts file to the model stage
func (file *File) Stage(stage *Stage) *File {

	if _, ok := stage.Files[file]; !ok {
		stage.Files[file] = struct{}{}
		stage.FileMap_Staged_Order[file] = stage.FileOrder
		stage.FileOrder++
	}
	stage.Files_mapString[file.Name] = file

	return file
}

// StagePreserveOrder puts file to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.FileOrder
// - update stage.FileOrder accordingly
func (file *File) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.Files[file]; !ok {
		stage.Files[file] = struct{}{}

		if order > stage.FileOrder {
			stage.FileOrder = order
		}
		stage.FileMap_Staged_Order[file] = order
		stage.FileOrder++
	}
	stage.Files_mapString[file.Name] = file
}

// Unstage removes file off the model stage
func (file *File) Unstage(stage *Stage) *File {
	delete(stage.Files, file)
	delete(stage.FileMap_Staged_Order, file)
	delete(stage.Files_mapString, file.Name)

	return file
}

// UnstageVoid removes file off the model stage
func (file *File) UnstageVoid(stage *Stage) {
	delete(stage.Files, file)
	delete(stage.FileMap_Staged_Order, file)
	delete(stage.Files_mapString, file.Name)
}

// commit file to the back repo (if it is already staged)
func (file *File) Commit(stage *Stage) *File {
	if _, ok := stage.Files[file]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitFile(file)
		}
	}
	return file
}

func (file *File) CommitVoid(stage *Stage) {
	file.Commit(stage)
}

func (file *File) StageVoid(stage *Stage) {
	file.Stage(stage)
}

// Checkout file to the back repo (if it is already staged)
func (file *File) Checkout(stage *Stage) *File {
	if _, ok := stage.Files[file]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutFile(file)
		}
	}
	return file
}

// for satisfaction of GongStruct interface
func (file *File) GetName() (res string) {
	return file.Name
}

// for satisfaction of GongStruct interface
func (file *File) SetName(name string) {
	file.Name = name
}

// Stage puts node to the model stage
func (node *Node) Stage(stage *Stage) *Node {

	if _, ok := stage.Nodes[node]; !ok {
		stage.Nodes[node] = struct{}{}
		stage.NodeMap_Staged_Order[node] = stage.NodeOrder
		stage.NodeOrder++
	}
	stage.Nodes_mapString[node.Name] = node

	return node
}

// StagePreserveOrder puts node to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.NodeOrder
// - update stage.NodeOrder accordingly
func (node *Node) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.Nodes[node]; !ok {
		stage.Nodes[node] = struct{}{}

		if order > stage.NodeOrder {
			stage.NodeOrder = order
		}
		stage.NodeMap_Staged_Order[node] = order
		stage.NodeOrder++
	}
	stage.Nodes_mapString[node.Name] = node
}

// Unstage removes node off the model stage
func (node *Node) Unstage(stage *Stage) *Node {
	delete(stage.Nodes, node)
	delete(stage.NodeMap_Staged_Order, node)
	delete(stage.Nodes_mapString, node.Name)

	return node
}

// UnstageVoid removes node off the model stage
func (node *Node) UnstageVoid(stage *Stage) {
	delete(stage.Nodes, node)
	delete(stage.NodeMap_Staged_Order, node)
	delete(stage.Nodes_mapString, node.Name)
}

// commit node to the back repo (if it is already staged)
func (node *Node) Commit(stage *Stage) *Node {
	if _, ok := stage.Nodes[node]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitNode(node)
		}
	}
	return node
}

func (node *Node) CommitVoid(stage *Stage) {
	node.Commit(stage)
}

func (node *Node) StageVoid(stage *Stage) {
	node.Stage(stage)
}

// Checkout node to the back repo (if it is already staged)
func (node *Node) Checkout(stage *Stage) *Node {
	if _, ok := stage.Nodes[node]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutNode(node)
		}
	}
	return node
}

// for satisfaction of GongStruct interface
func (node *Node) GetName() (res string) {
	return node.Name
}

// for satisfaction of GongStruct interface
func (node *Node) SetName(name string) {
	node.Name = name
}

// Stage puts paragraph to the model stage
func (paragraph *Paragraph) Stage(stage *Stage) *Paragraph {

	if _, ok := stage.Paragraphs[paragraph]; !ok {
		stage.Paragraphs[paragraph] = struct{}{}
		stage.ParagraphMap_Staged_Order[paragraph] = stage.ParagraphOrder
		stage.ParagraphOrder++
	}
	stage.Paragraphs_mapString[paragraph.Name] = paragraph

	return paragraph
}

// StagePreserveOrder puts paragraph to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ParagraphOrder
// - update stage.ParagraphOrder accordingly
func (paragraph *Paragraph) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.Paragraphs[paragraph]; !ok {
		stage.Paragraphs[paragraph] = struct{}{}

		if order > stage.ParagraphOrder {
			stage.ParagraphOrder = order
		}
		stage.ParagraphMap_Staged_Order[paragraph] = order
		stage.ParagraphOrder++
	}
	stage.Paragraphs_mapString[paragraph.Name] = paragraph
}

// Unstage removes paragraph off the model stage
func (paragraph *Paragraph) Unstage(stage *Stage) *Paragraph {
	delete(stage.Paragraphs, paragraph)
	delete(stage.ParagraphMap_Staged_Order, paragraph)
	delete(stage.Paragraphs_mapString, paragraph.Name)

	return paragraph
}

// UnstageVoid removes paragraph off the model stage
func (paragraph *Paragraph) UnstageVoid(stage *Stage) {
	delete(stage.Paragraphs, paragraph)
	delete(stage.ParagraphMap_Staged_Order, paragraph)
	delete(stage.Paragraphs_mapString, paragraph.Name)
}

// commit paragraph to the back repo (if it is already staged)
func (paragraph *Paragraph) Commit(stage *Stage) *Paragraph {
	if _, ok := stage.Paragraphs[paragraph]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitParagraph(paragraph)
		}
	}
	return paragraph
}

func (paragraph *Paragraph) CommitVoid(stage *Stage) {
	paragraph.Commit(stage)
}

func (paragraph *Paragraph) StageVoid(stage *Stage) {
	paragraph.Stage(stage)
}

// Checkout paragraph to the back repo (if it is already staged)
func (paragraph *Paragraph) Checkout(stage *Stage) *Paragraph {
	if _, ok := stage.Paragraphs[paragraph]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutParagraph(paragraph)
		}
	}
	return paragraph
}

// for satisfaction of GongStruct interface
func (paragraph *Paragraph) GetName() (res string) {
	return paragraph.Name
}

// for satisfaction of GongStruct interface
func (paragraph *Paragraph) SetName(name string) {
	paragraph.Name = name
}

// Stage puts paragraphproperties to the model stage
func (paragraphproperties *ParagraphProperties) Stage(stage *Stage) *ParagraphProperties {

	if _, ok := stage.ParagraphPropertiess[paragraphproperties]; !ok {
		stage.ParagraphPropertiess[paragraphproperties] = struct{}{}
		stage.ParagraphPropertiesMap_Staged_Order[paragraphproperties] = stage.ParagraphPropertiesOrder
		stage.ParagraphPropertiesOrder++
	}
	stage.ParagraphPropertiess_mapString[paragraphproperties.Name] = paragraphproperties

	return paragraphproperties
}

// StagePreserveOrder puts paragraphproperties to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ParagraphPropertiesOrder
// - update stage.ParagraphPropertiesOrder accordingly
func (paragraphproperties *ParagraphProperties) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.ParagraphPropertiess[paragraphproperties]; !ok {
		stage.ParagraphPropertiess[paragraphproperties] = struct{}{}

		if order > stage.ParagraphPropertiesOrder {
			stage.ParagraphPropertiesOrder = order
		}
		stage.ParagraphPropertiesMap_Staged_Order[paragraphproperties] = order
		stage.ParagraphPropertiesOrder++
	}
	stage.ParagraphPropertiess_mapString[paragraphproperties.Name] = paragraphproperties
}

// Unstage removes paragraphproperties off the model stage
func (paragraphproperties *ParagraphProperties) Unstage(stage *Stage) *ParagraphProperties {
	delete(stage.ParagraphPropertiess, paragraphproperties)
	delete(stage.ParagraphPropertiesMap_Staged_Order, paragraphproperties)
	delete(stage.ParagraphPropertiess_mapString, paragraphproperties.Name)

	return paragraphproperties
}

// UnstageVoid removes paragraphproperties off the model stage
func (paragraphproperties *ParagraphProperties) UnstageVoid(stage *Stage) {
	delete(stage.ParagraphPropertiess, paragraphproperties)
	delete(stage.ParagraphPropertiesMap_Staged_Order, paragraphproperties)
	delete(stage.ParagraphPropertiess_mapString, paragraphproperties.Name)
}

// commit paragraphproperties to the back repo (if it is already staged)
func (paragraphproperties *ParagraphProperties) Commit(stage *Stage) *ParagraphProperties {
	if _, ok := stage.ParagraphPropertiess[paragraphproperties]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitParagraphProperties(paragraphproperties)
		}
	}
	return paragraphproperties
}

func (paragraphproperties *ParagraphProperties) CommitVoid(stage *Stage) {
	paragraphproperties.Commit(stage)
}

func (paragraphproperties *ParagraphProperties) StageVoid(stage *Stage) {
	paragraphproperties.Stage(stage)
}

// Checkout paragraphproperties to the back repo (if it is already staged)
func (paragraphproperties *ParagraphProperties) Checkout(stage *Stage) *ParagraphProperties {
	if _, ok := stage.ParagraphPropertiess[paragraphproperties]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutParagraphProperties(paragraphproperties)
		}
	}
	return paragraphproperties
}

// for satisfaction of GongStruct interface
func (paragraphproperties *ParagraphProperties) GetName() (res string) {
	return paragraphproperties.Name
}

// for satisfaction of GongStruct interface
func (paragraphproperties *ParagraphProperties) SetName(name string) {
	paragraphproperties.Name = name
}

// Stage puts paragraphstyle to the model stage
func (paragraphstyle *ParagraphStyle) Stage(stage *Stage) *ParagraphStyle {

	if _, ok := stage.ParagraphStyles[paragraphstyle]; !ok {
		stage.ParagraphStyles[paragraphstyle] = struct{}{}
		stage.ParagraphStyleMap_Staged_Order[paragraphstyle] = stage.ParagraphStyleOrder
		stage.ParagraphStyleOrder++
	}
	stage.ParagraphStyles_mapString[paragraphstyle.Name] = paragraphstyle

	return paragraphstyle
}

// StagePreserveOrder puts paragraphstyle to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ParagraphStyleOrder
// - update stage.ParagraphStyleOrder accordingly
func (paragraphstyle *ParagraphStyle) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.ParagraphStyles[paragraphstyle]; !ok {
		stage.ParagraphStyles[paragraphstyle] = struct{}{}

		if order > stage.ParagraphStyleOrder {
			stage.ParagraphStyleOrder = order
		}
		stage.ParagraphStyleMap_Staged_Order[paragraphstyle] = order
		stage.ParagraphStyleOrder++
	}
	stage.ParagraphStyles_mapString[paragraphstyle.Name] = paragraphstyle
}

// Unstage removes paragraphstyle off the model stage
func (paragraphstyle *ParagraphStyle) Unstage(stage *Stage) *ParagraphStyle {
	delete(stage.ParagraphStyles, paragraphstyle)
	delete(stage.ParagraphStyleMap_Staged_Order, paragraphstyle)
	delete(stage.ParagraphStyles_mapString, paragraphstyle.Name)

	return paragraphstyle
}

// UnstageVoid removes paragraphstyle off the model stage
func (paragraphstyle *ParagraphStyle) UnstageVoid(stage *Stage) {
	delete(stage.ParagraphStyles, paragraphstyle)
	delete(stage.ParagraphStyleMap_Staged_Order, paragraphstyle)
	delete(stage.ParagraphStyles_mapString, paragraphstyle.Name)
}

// commit paragraphstyle to the back repo (if it is already staged)
func (paragraphstyle *ParagraphStyle) Commit(stage *Stage) *ParagraphStyle {
	if _, ok := stage.ParagraphStyles[paragraphstyle]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitParagraphStyle(paragraphstyle)
		}
	}
	return paragraphstyle
}

func (paragraphstyle *ParagraphStyle) CommitVoid(stage *Stage) {
	paragraphstyle.Commit(stage)
}

func (paragraphstyle *ParagraphStyle) StageVoid(stage *Stage) {
	paragraphstyle.Stage(stage)
}

// Checkout paragraphstyle to the back repo (if it is already staged)
func (paragraphstyle *ParagraphStyle) Checkout(stage *Stage) *ParagraphStyle {
	if _, ok := stage.ParagraphStyles[paragraphstyle]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutParagraphStyle(paragraphstyle)
		}
	}
	return paragraphstyle
}

// for satisfaction of GongStruct interface
func (paragraphstyle *ParagraphStyle) GetName() (res string) {
	return paragraphstyle.Name
}

// for satisfaction of GongStruct interface
func (paragraphstyle *ParagraphStyle) SetName(name string) {
	paragraphstyle.Name = name
}

// Stage puts rune to the model stage
func (rune *Rune) Stage(stage *Stage) *Rune {

	if _, ok := stage.Runes[rune]; !ok {
		stage.Runes[rune] = struct{}{}
		stage.RuneMap_Staged_Order[rune] = stage.RuneOrder
		stage.RuneOrder++
	}
	stage.Runes_mapString[rune.Name] = rune

	return rune
}

// StagePreserveOrder puts rune to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.RuneOrder
// - update stage.RuneOrder accordingly
func (rune *Rune) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.Runes[rune]; !ok {
		stage.Runes[rune] = struct{}{}

		if order > stage.RuneOrder {
			stage.RuneOrder = order
		}
		stage.RuneMap_Staged_Order[rune] = order
		stage.RuneOrder++
	}
	stage.Runes_mapString[rune.Name] = rune
}

// Unstage removes rune off the model stage
func (rune *Rune) Unstage(stage *Stage) *Rune {
	delete(stage.Runes, rune)
	delete(stage.RuneMap_Staged_Order, rune)
	delete(stage.Runes_mapString, rune.Name)

	return rune
}

// UnstageVoid removes rune off the model stage
func (rune *Rune) UnstageVoid(stage *Stage) {
	delete(stage.Runes, rune)
	delete(stage.RuneMap_Staged_Order, rune)
	delete(stage.Runes_mapString, rune.Name)
}

// commit rune to the back repo (if it is already staged)
func (rune *Rune) Commit(stage *Stage) *Rune {
	if _, ok := stage.Runes[rune]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitRune(rune)
		}
	}
	return rune
}

func (rune *Rune) CommitVoid(stage *Stage) {
	rune.Commit(stage)
}

func (rune *Rune) StageVoid(stage *Stage) {
	rune.Stage(stage)
}

// Checkout rune to the back repo (if it is already staged)
func (rune *Rune) Checkout(stage *Stage) *Rune {
	if _, ok := stage.Runes[rune]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutRune(rune)
		}
	}
	return rune
}

// for satisfaction of GongStruct interface
func (rune *Rune) GetName() (res string) {
	return rune.Name
}

// for satisfaction of GongStruct interface
func (rune *Rune) SetName(name string) {
	rune.Name = name
}

// Stage puts runeproperties to the model stage
func (runeproperties *RuneProperties) Stage(stage *Stage) *RuneProperties {

	if _, ok := stage.RunePropertiess[runeproperties]; !ok {
		stage.RunePropertiess[runeproperties] = struct{}{}
		stage.RunePropertiesMap_Staged_Order[runeproperties] = stage.RunePropertiesOrder
		stage.RunePropertiesOrder++
	}
	stage.RunePropertiess_mapString[runeproperties.Name] = runeproperties

	return runeproperties
}

// StagePreserveOrder puts runeproperties to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.RunePropertiesOrder
// - update stage.RunePropertiesOrder accordingly
func (runeproperties *RuneProperties) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.RunePropertiess[runeproperties]; !ok {
		stage.RunePropertiess[runeproperties] = struct{}{}

		if order > stage.RunePropertiesOrder {
			stage.RunePropertiesOrder = order
		}
		stage.RunePropertiesMap_Staged_Order[runeproperties] = order
		stage.RunePropertiesOrder++
	}
	stage.RunePropertiess_mapString[runeproperties.Name] = runeproperties
}

// Unstage removes runeproperties off the model stage
func (runeproperties *RuneProperties) Unstage(stage *Stage) *RuneProperties {
	delete(stage.RunePropertiess, runeproperties)
	delete(stage.RunePropertiesMap_Staged_Order, runeproperties)
	delete(stage.RunePropertiess_mapString, runeproperties.Name)

	return runeproperties
}

// UnstageVoid removes runeproperties off the model stage
func (runeproperties *RuneProperties) UnstageVoid(stage *Stage) {
	delete(stage.RunePropertiess, runeproperties)
	delete(stage.RunePropertiesMap_Staged_Order, runeproperties)
	delete(stage.RunePropertiess_mapString, runeproperties.Name)
}

// commit runeproperties to the back repo (if it is already staged)
func (runeproperties *RuneProperties) Commit(stage *Stage) *RuneProperties {
	if _, ok := stage.RunePropertiess[runeproperties]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitRuneProperties(runeproperties)
		}
	}
	return runeproperties
}

func (runeproperties *RuneProperties) CommitVoid(stage *Stage) {
	runeproperties.Commit(stage)
}

func (runeproperties *RuneProperties) StageVoid(stage *Stage) {
	runeproperties.Stage(stage)
}

// Checkout runeproperties to the back repo (if it is already staged)
func (runeproperties *RuneProperties) Checkout(stage *Stage) *RuneProperties {
	if _, ok := stage.RunePropertiess[runeproperties]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutRuneProperties(runeproperties)
		}
	}
	return runeproperties
}

// for satisfaction of GongStruct interface
func (runeproperties *RuneProperties) GetName() (res string) {
	return runeproperties.Name
}

// for satisfaction of GongStruct interface
func (runeproperties *RuneProperties) SetName(name string) {
	runeproperties.Name = name
}

// Stage puts table to the model stage
func (table *Table) Stage(stage *Stage) *Table {

	if _, ok := stage.Tables[table]; !ok {
		stage.Tables[table] = struct{}{}
		stage.TableMap_Staged_Order[table] = stage.TableOrder
		stage.TableOrder++
	}
	stage.Tables_mapString[table.Name] = table

	return table
}

// StagePreserveOrder puts table to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TableOrder
// - update stage.TableOrder accordingly
func (table *Table) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.Tables[table]; !ok {
		stage.Tables[table] = struct{}{}

		if order > stage.TableOrder {
			stage.TableOrder = order
		}
		stage.TableMap_Staged_Order[table] = order
		stage.TableOrder++
	}
	stage.Tables_mapString[table.Name] = table
}

// Unstage removes table off the model stage
func (table *Table) Unstage(stage *Stage) *Table {
	delete(stage.Tables, table)
	delete(stage.TableMap_Staged_Order, table)
	delete(stage.Tables_mapString, table.Name)

	return table
}

// UnstageVoid removes table off the model stage
func (table *Table) UnstageVoid(stage *Stage) {
	delete(stage.Tables, table)
	delete(stage.TableMap_Staged_Order, table)
	delete(stage.Tables_mapString, table.Name)
}

// commit table to the back repo (if it is already staged)
func (table *Table) Commit(stage *Stage) *Table {
	if _, ok := stage.Tables[table]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTable(table)
		}
	}
	return table
}

func (table *Table) CommitVoid(stage *Stage) {
	table.Commit(stage)
}

func (table *Table) StageVoid(stage *Stage) {
	table.Stage(stage)
}

// Checkout table to the back repo (if it is already staged)
func (table *Table) Checkout(stage *Stage) *Table {
	if _, ok := stage.Tables[table]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutTable(table)
		}
	}
	return table
}

// for satisfaction of GongStruct interface
func (table *Table) GetName() (res string) {
	return table.Name
}

// for satisfaction of GongStruct interface
func (table *Table) SetName(name string) {
	table.Name = name
}

// Stage puts tablecolumn to the model stage
func (tablecolumn *TableColumn) Stage(stage *Stage) *TableColumn {

	if _, ok := stage.TableColumns[tablecolumn]; !ok {
		stage.TableColumns[tablecolumn] = struct{}{}
		stage.TableColumnMap_Staged_Order[tablecolumn] = stage.TableColumnOrder
		stage.TableColumnOrder++
	}
	stage.TableColumns_mapString[tablecolumn.Name] = tablecolumn

	return tablecolumn
}

// StagePreserveOrder puts tablecolumn to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TableColumnOrder
// - update stage.TableColumnOrder accordingly
func (tablecolumn *TableColumn) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.TableColumns[tablecolumn]; !ok {
		stage.TableColumns[tablecolumn] = struct{}{}

		if order > stage.TableColumnOrder {
			stage.TableColumnOrder = order
		}
		stage.TableColumnMap_Staged_Order[tablecolumn] = order
		stage.TableColumnOrder++
	}
	stage.TableColumns_mapString[tablecolumn.Name] = tablecolumn
}

// Unstage removes tablecolumn off the model stage
func (tablecolumn *TableColumn) Unstage(stage *Stage) *TableColumn {
	delete(stage.TableColumns, tablecolumn)
	delete(stage.TableColumnMap_Staged_Order, tablecolumn)
	delete(stage.TableColumns_mapString, tablecolumn.Name)

	return tablecolumn
}

// UnstageVoid removes tablecolumn off the model stage
func (tablecolumn *TableColumn) UnstageVoid(stage *Stage) {
	delete(stage.TableColumns, tablecolumn)
	delete(stage.TableColumnMap_Staged_Order, tablecolumn)
	delete(stage.TableColumns_mapString, tablecolumn.Name)
}

// commit tablecolumn to the back repo (if it is already staged)
func (tablecolumn *TableColumn) Commit(stage *Stage) *TableColumn {
	if _, ok := stage.TableColumns[tablecolumn]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTableColumn(tablecolumn)
		}
	}
	return tablecolumn
}

func (tablecolumn *TableColumn) CommitVoid(stage *Stage) {
	tablecolumn.Commit(stage)
}

func (tablecolumn *TableColumn) StageVoid(stage *Stage) {
	tablecolumn.Stage(stage)
}

// Checkout tablecolumn to the back repo (if it is already staged)
func (tablecolumn *TableColumn) Checkout(stage *Stage) *TableColumn {
	if _, ok := stage.TableColumns[tablecolumn]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutTableColumn(tablecolumn)
		}
	}
	return tablecolumn
}

// for satisfaction of GongStruct interface
func (tablecolumn *TableColumn) GetName() (res string) {
	return tablecolumn.Name
}

// for satisfaction of GongStruct interface
func (tablecolumn *TableColumn) SetName(name string) {
	tablecolumn.Name = name
}

// Stage puts tableproperties to the model stage
func (tableproperties *TableProperties) Stage(stage *Stage) *TableProperties {

	if _, ok := stage.TablePropertiess[tableproperties]; !ok {
		stage.TablePropertiess[tableproperties] = struct{}{}
		stage.TablePropertiesMap_Staged_Order[tableproperties] = stage.TablePropertiesOrder
		stage.TablePropertiesOrder++
	}
	stage.TablePropertiess_mapString[tableproperties.Name] = tableproperties

	return tableproperties
}

// StagePreserveOrder puts tableproperties to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TablePropertiesOrder
// - update stage.TablePropertiesOrder accordingly
func (tableproperties *TableProperties) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.TablePropertiess[tableproperties]; !ok {
		stage.TablePropertiess[tableproperties] = struct{}{}

		if order > stage.TablePropertiesOrder {
			stage.TablePropertiesOrder = order
		}
		stage.TablePropertiesMap_Staged_Order[tableproperties] = order
		stage.TablePropertiesOrder++
	}
	stage.TablePropertiess_mapString[tableproperties.Name] = tableproperties
}

// Unstage removes tableproperties off the model stage
func (tableproperties *TableProperties) Unstage(stage *Stage) *TableProperties {
	delete(stage.TablePropertiess, tableproperties)
	delete(stage.TablePropertiesMap_Staged_Order, tableproperties)
	delete(stage.TablePropertiess_mapString, tableproperties.Name)

	return tableproperties
}

// UnstageVoid removes tableproperties off the model stage
func (tableproperties *TableProperties) UnstageVoid(stage *Stage) {
	delete(stage.TablePropertiess, tableproperties)
	delete(stage.TablePropertiesMap_Staged_Order, tableproperties)
	delete(stage.TablePropertiess_mapString, tableproperties.Name)
}

// commit tableproperties to the back repo (if it is already staged)
func (tableproperties *TableProperties) Commit(stage *Stage) *TableProperties {
	if _, ok := stage.TablePropertiess[tableproperties]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTableProperties(tableproperties)
		}
	}
	return tableproperties
}

func (tableproperties *TableProperties) CommitVoid(stage *Stage) {
	tableproperties.Commit(stage)
}

func (tableproperties *TableProperties) StageVoid(stage *Stage) {
	tableproperties.Stage(stage)
}

// Checkout tableproperties to the back repo (if it is already staged)
func (tableproperties *TableProperties) Checkout(stage *Stage) *TableProperties {
	if _, ok := stage.TablePropertiess[tableproperties]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutTableProperties(tableproperties)
		}
	}
	return tableproperties
}

// for satisfaction of GongStruct interface
func (tableproperties *TableProperties) GetName() (res string) {
	return tableproperties.Name
}

// for satisfaction of GongStruct interface
func (tableproperties *TableProperties) SetName(name string) {
	tableproperties.Name = name
}

// Stage puts tablerow to the model stage
func (tablerow *TableRow) Stage(stage *Stage) *TableRow {

	if _, ok := stage.TableRows[tablerow]; !ok {
		stage.TableRows[tablerow] = struct{}{}
		stage.TableRowMap_Staged_Order[tablerow] = stage.TableRowOrder
		stage.TableRowOrder++
	}
	stage.TableRows_mapString[tablerow.Name] = tablerow

	return tablerow
}

// StagePreserveOrder puts tablerow to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TableRowOrder
// - update stage.TableRowOrder accordingly
func (tablerow *TableRow) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.TableRows[tablerow]; !ok {
		stage.TableRows[tablerow] = struct{}{}

		if order > stage.TableRowOrder {
			stage.TableRowOrder = order
		}
		stage.TableRowMap_Staged_Order[tablerow] = order
		stage.TableRowOrder++
	}
	stage.TableRows_mapString[tablerow.Name] = tablerow
}

// Unstage removes tablerow off the model stage
func (tablerow *TableRow) Unstage(stage *Stage) *TableRow {
	delete(stage.TableRows, tablerow)
	delete(stage.TableRowMap_Staged_Order, tablerow)
	delete(stage.TableRows_mapString, tablerow.Name)

	return tablerow
}

// UnstageVoid removes tablerow off the model stage
func (tablerow *TableRow) UnstageVoid(stage *Stage) {
	delete(stage.TableRows, tablerow)
	delete(stage.TableRowMap_Staged_Order, tablerow)
	delete(stage.TableRows_mapString, tablerow.Name)
}

// commit tablerow to the back repo (if it is already staged)
func (tablerow *TableRow) Commit(stage *Stage) *TableRow {
	if _, ok := stage.TableRows[tablerow]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTableRow(tablerow)
		}
	}
	return tablerow
}

func (tablerow *TableRow) CommitVoid(stage *Stage) {
	tablerow.Commit(stage)
}

func (tablerow *TableRow) StageVoid(stage *Stage) {
	tablerow.Stage(stage)
}

// Checkout tablerow to the back repo (if it is already staged)
func (tablerow *TableRow) Checkout(stage *Stage) *TableRow {
	if _, ok := stage.TableRows[tablerow]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutTableRow(tablerow)
		}
	}
	return tablerow
}

// for satisfaction of GongStruct interface
func (tablerow *TableRow) GetName() (res string) {
	return tablerow.Name
}

// for satisfaction of GongStruct interface
func (tablerow *TableRow) SetName(name string) {
	tablerow.Name = name
}

// Stage puts tablestyle to the model stage
func (tablestyle *TableStyle) Stage(stage *Stage) *TableStyle {

	if _, ok := stage.TableStyles[tablestyle]; !ok {
		stage.TableStyles[tablestyle] = struct{}{}
		stage.TableStyleMap_Staged_Order[tablestyle] = stage.TableStyleOrder
		stage.TableStyleOrder++
	}
	stage.TableStyles_mapString[tablestyle.Name] = tablestyle

	return tablestyle
}

// StagePreserveOrder puts tablestyle to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TableStyleOrder
// - update stage.TableStyleOrder accordingly
func (tablestyle *TableStyle) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.TableStyles[tablestyle]; !ok {
		stage.TableStyles[tablestyle] = struct{}{}

		if order > stage.TableStyleOrder {
			stage.TableStyleOrder = order
		}
		stage.TableStyleMap_Staged_Order[tablestyle] = order
		stage.TableStyleOrder++
	}
	stage.TableStyles_mapString[tablestyle.Name] = tablestyle
}

// Unstage removes tablestyle off the model stage
func (tablestyle *TableStyle) Unstage(stage *Stage) *TableStyle {
	delete(stage.TableStyles, tablestyle)
	delete(stage.TableStyleMap_Staged_Order, tablestyle)
	delete(stage.TableStyles_mapString, tablestyle.Name)

	return tablestyle
}

// UnstageVoid removes tablestyle off the model stage
func (tablestyle *TableStyle) UnstageVoid(stage *Stage) {
	delete(stage.TableStyles, tablestyle)
	delete(stage.TableStyleMap_Staged_Order, tablestyle)
	delete(stage.TableStyles_mapString, tablestyle.Name)
}

// commit tablestyle to the back repo (if it is already staged)
func (tablestyle *TableStyle) Commit(stage *Stage) *TableStyle {
	if _, ok := stage.TableStyles[tablestyle]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTableStyle(tablestyle)
		}
	}
	return tablestyle
}

func (tablestyle *TableStyle) CommitVoid(stage *Stage) {
	tablestyle.Commit(stage)
}

func (tablestyle *TableStyle) StageVoid(stage *Stage) {
	tablestyle.Stage(stage)
}

// Checkout tablestyle to the back repo (if it is already staged)
func (tablestyle *TableStyle) Checkout(stage *Stage) *TableStyle {
	if _, ok := stage.TableStyles[tablestyle]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutTableStyle(tablestyle)
		}
	}
	return tablestyle
}

// for satisfaction of GongStruct interface
func (tablestyle *TableStyle) GetName() (res string) {
	return tablestyle.Name
}

// for satisfaction of GongStruct interface
func (tablestyle *TableStyle) SetName(name string) {
	tablestyle.Name = name
}

// Stage puts text to the model stage
func (text *Text) Stage(stage *Stage) *Text {

	if _, ok := stage.Texts[text]; !ok {
		stage.Texts[text] = struct{}{}
		stage.TextMap_Staged_Order[text] = stage.TextOrder
		stage.TextOrder++
	}
	stage.Texts_mapString[text.Name] = text

	return text
}

// StagePreserveOrder puts text to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TextOrder
// - update stage.TextOrder accordingly
func (text *Text) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.Texts[text]; !ok {
		stage.Texts[text] = struct{}{}

		if order > stage.TextOrder {
			stage.TextOrder = order
		}
		stage.TextMap_Staged_Order[text] = order
		stage.TextOrder++
	}
	stage.Texts_mapString[text.Name] = text
}

// Unstage removes text off the model stage
func (text *Text) Unstage(stage *Stage) *Text {
	delete(stage.Texts, text)
	delete(stage.TextMap_Staged_Order, text)
	delete(stage.Texts_mapString, text.Name)

	return text
}

// UnstageVoid removes text off the model stage
func (text *Text) UnstageVoid(stage *Stage) {
	delete(stage.Texts, text)
	delete(stage.TextMap_Staged_Order, text)
	delete(stage.Texts_mapString, text.Name)
}

// commit text to the back repo (if it is already staged)
func (text *Text) Commit(stage *Stage) *Text {
	if _, ok := stage.Texts[text]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitText(text)
		}
	}
	return text
}

func (text *Text) CommitVoid(stage *Stage) {
	text.Commit(stage)
}

func (text *Text) StageVoid(stage *Stage) {
	text.Stage(stage)
}

// Checkout text to the back repo (if it is already staged)
func (text *Text) Checkout(stage *Stage) *Text {
	if _, ok := stage.Texts[text]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutText(text)
		}
	}
	return text
}

// for satisfaction of GongStruct interface
func (text *Text) GetName() (res string) {
	return text.Name
}

// for satisfaction of GongStruct interface
func (text *Text) SetName(name string) {
	text.Name = name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMBody(Body *Body)
	CreateORMDocument(Document *Document)
	CreateORMDocx(Docx *Docx)
	CreateORMFile(File *File)
	CreateORMNode(Node *Node)
	CreateORMParagraph(Paragraph *Paragraph)
	CreateORMParagraphProperties(ParagraphProperties *ParagraphProperties)
	CreateORMParagraphStyle(ParagraphStyle *ParagraphStyle)
	CreateORMRune(Rune *Rune)
	CreateORMRuneProperties(RuneProperties *RuneProperties)
	CreateORMTable(Table *Table)
	CreateORMTableColumn(TableColumn *TableColumn)
	CreateORMTableProperties(TableProperties *TableProperties)
	CreateORMTableRow(TableRow *TableRow)
	CreateORMTableStyle(TableStyle *TableStyle)
	CreateORMText(Text *Text)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMBody(Body *Body)
	DeleteORMDocument(Document *Document)
	DeleteORMDocx(Docx *Docx)
	DeleteORMFile(File *File)
	DeleteORMNode(Node *Node)
	DeleteORMParagraph(Paragraph *Paragraph)
	DeleteORMParagraphProperties(ParagraphProperties *ParagraphProperties)
	DeleteORMParagraphStyle(ParagraphStyle *ParagraphStyle)
	DeleteORMRune(Rune *Rune)
	DeleteORMRuneProperties(RuneProperties *RuneProperties)
	DeleteORMTable(Table *Table)
	DeleteORMTableColumn(TableColumn *TableColumn)
	DeleteORMTableProperties(TableProperties *TableProperties)
	DeleteORMTableRow(TableRow *TableRow)
	DeleteORMTableStyle(TableStyle *TableStyle)
	DeleteORMText(Text *Text)
}

func (stage *Stage) Reset() { // insertion point for array reset
	stage.Bodys = make(map[*Body]struct{})
	stage.Bodys_mapString = make(map[string]*Body)
	stage.BodyMap_Staged_Order = make(map[*Body]uint)
	stage.BodyOrder = 0

	stage.Documents = make(map[*Document]struct{})
	stage.Documents_mapString = make(map[string]*Document)
	stage.DocumentMap_Staged_Order = make(map[*Document]uint)
	stage.DocumentOrder = 0

	stage.Docxs = make(map[*Docx]struct{})
	stage.Docxs_mapString = make(map[string]*Docx)
	stage.DocxMap_Staged_Order = make(map[*Docx]uint)
	stage.DocxOrder = 0

	stage.Files = make(map[*File]struct{})
	stage.Files_mapString = make(map[string]*File)
	stage.FileMap_Staged_Order = make(map[*File]uint)
	stage.FileOrder = 0

	stage.Nodes = make(map[*Node]struct{})
	stage.Nodes_mapString = make(map[string]*Node)
	stage.NodeMap_Staged_Order = make(map[*Node]uint)
	stage.NodeOrder = 0

	stage.Paragraphs = make(map[*Paragraph]struct{})
	stage.Paragraphs_mapString = make(map[string]*Paragraph)
	stage.ParagraphMap_Staged_Order = make(map[*Paragraph]uint)
	stage.ParagraphOrder = 0

	stage.ParagraphPropertiess = make(map[*ParagraphProperties]struct{})
	stage.ParagraphPropertiess_mapString = make(map[string]*ParagraphProperties)
	stage.ParagraphPropertiesMap_Staged_Order = make(map[*ParagraphProperties]uint)
	stage.ParagraphPropertiesOrder = 0

	stage.ParagraphStyles = make(map[*ParagraphStyle]struct{})
	stage.ParagraphStyles_mapString = make(map[string]*ParagraphStyle)
	stage.ParagraphStyleMap_Staged_Order = make(map[*ParagraphStyle]uint)
	stage.ParagraphStyleOrder = 0

	stage.Runes = make(map[*Rune]struct{})
	stage.Runes_mapString = make(map[string]*Rune)
	stage.RuneMap_Staged_Order = make(map[*Rune]uint)
	stage.RuneOrder = 0

	stage.RunePropertiess = make(map[*RuneProperties]struct{})
	stage.RunePropertiess_mapString = make(map[string]*RuneProperties)
	stage.RunePropertiesMap_Staged_Order = make(map[*RuneProperties]uint)
	stage.RunePropertiesOrder = 0

	stage.Tables = make(map[*Table]struct{})
	stage.Tables_mapString = make(map[string]*Table)
	stage.TableMap_Staged_Order = make(map[*Table]uint)
	stage.TableOrder = 0

	stage.TableColumns = make(map[*TableColumn]struct{})
	stage.TableColumns_mapString = make(map[string]*TableColumn)
	stage.TableColumnMap_Staged_Order = make(map[*TableColumn]uint)
	stage.TableColumnOrder = 0

	stage.TablePropertiess = make(map[*TableProperties]struct{})
	stage.TablePropertiess_mapString = make(map[string]*TableProperties)
	stage.TablePropertiesMap_Staged_Order = make(map[*TableProperties]uint)
	stage.TablePropertiesOrder = 0

	stage.TableRows = make(map[*TableRow]struct{})
	stage.TableRows_mapString = make(map[string]*TableRow)
	stage.TableRowMap_Staged_Order = make(map[*TableRow]uint)
	stage.TableRowOrder = 0

	stage.TableStyles = make(map[*TableStyle]struct{})
	stage.TableStyles_mapString = make(map[string]*TableStyle)
	stage.TableStyleMap_Staged_Order = make(map[*TableStyle]uint)
	stage.TableStyleOrder = 0

	stage.Texts = make(map[*Text]struct{})
	stage.Texts_mapString = make(map[string]*Text)
	stage.TextMap_Staged_Order = make(map[*Text]uint)
	stage.TextOrder = 0

	if stage.GetProbeIF() != nil {
		stage.GetProbeIF().ResetNotifications()
	}
	if stage.IsInDeltaMode() {
		stage.ComputeReferenceAndOrders()
	}
}

func (stage *Stage) Nil() { // insertion point for array nil
	stage.Bodys = nil
	stage.Bodys_mapString = nil

	stage.Documents = nil
	stage.Documents_mapString = nil

	stage.Docxs = nil
	stage.Docxs_mapString = nil

	stage.Files = nil
	stage.Files_mapString = nil

	stage.Nodes = nil
	stage.Nodes_mapString = nil

	stage.Paragraphs = nil
	stage.Paragraphs_mapString = nil

	stage.ParagraphPropertiess = nil
	stage.ParagraphPropertiess_mapString = nil

	stage.ParagraphStyles = nil
	stage.ParagraphStyles_mapString = nil

	stage.Runes = nil
	stage.Runes_mapString = nil

	stage.RunePropertiess = nil
	stage.RunePropertiess_mapString = nil

	stage.Tables = nil
	stage.Tables_mapString = nil

	stage.TableColumns = nil
	stage.TableColumns_mapString = nil

	stage.TablePropertiess = nil
	stage.TablePropertiess_mapString = nil

	stage.TableRows = nil
	stage.TableRows_mapString = nil

	stage.TableStyles = nil
	stage.TableStyles_mapString = nil

	stage.Texts = nil
	stage.Texts_mapString = nil

}

func (stage *Stage) Unstage() { // insertion point for array nil
	for body := range stage.Bodys {
		body.Unstage(stage)
	}

	for document := range stage.Documents {
		document.Unstage(stage)
	}

	for docx := range stage.Docxs {
		docx.Unstage(stage)
	}

	for file := range stage.Files {
		file.Unstage(stage)
	}

	for node := range stage.Nodes {
		node.Unstage(stage)
	}

	for paragraph := range stage.Paragraphs {
		paragraph.Unstage(stage)
	}

	for paragraphproperties := range stage.ParagraphPropertiess {
		paragraphproperties.Unstage(stage)
	}

	for paragraphstyle := range stage.ParagraphStyles {
		paragraphstyle.Unstage(stage)
	}

	for rune := range stage.Runes {
		rune.Unstage(stage)
	}

	for runeproperties := range stage.RunePropertiess {
		runeproperties.Unstage(stage)
	}

	for table := range stage.Tables {
		table.Unstage(stage)
	}

	for tablecolumn := range stage.TableColumns {
		tablecolumn.Unstage(stage)
	}

	for tableproperties := range stage.TablePropertiess {
		tableproperties.Unstage(stage)
	}

	for tablerow := range stage.TableRows {
		tablerow.Unstage(stage)
	}

	for tablestyle := range stage.TableStyles {
		tablestyle.Unstage(stage)
	}

	for text := range stage.Texts {
		text.Unstage(stage)
	}

}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type Gongstruct interface {
}

type GongtructBasicField interface {
	int | float64 | bool | string | time.Time | time.Duration
}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type GongstructIF interface {
	GetName() string
	SetName(string)
	CommitVoid(*Stage)
	StageVoid(*Stage)
	UnstageVoid(stage *Stage)
	GongGetFieldHeaders() []GongFieldHeader
	GongClean(stage *Stage) (modified bool)
	GongGetFieldValue(fieldName string, stage *Stage) GongFieldValue
	GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error
	GongGetGongstructName() string
	GongGetOrder(stage *Stage) uint
	GongGetReferenceIdentifier(stage *Stage) string
	GongGetIdentifier(stage *Stage) string
	GongCopy() GongstructIF
	GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) string
	GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) GongstructIF
}
type PointerToGongstruct interface {
	GongstructIF
	comparable
}

func CompareGongstructByName[T PointerToGongstruct](a, b T) int {
	return cmp.Compare(a.GetName(), b.GetName())
}

func SortGongstructSetByName[T PointerToGongstruct](set map[T]struct{}) (sortedSlice []T) {

	for key := range set {
		sortedSlice = append(sortedSlice, key)
	}
	slices.SortFunc(sortedSlice, CompareGongstructByName)

	return
}

func GetGongstrucsSorted[T PointerToGongstruct](stage *Stage) (sortedSlice []T) {

	set := GetGongstructInstancesSetFromPointerType[T](stage)
	sortedSlice = SortGongstructSetByName(*set)

	return
}

type GongstructSet interface {
	map[any]any
}

type GongstructMapString interface {
	map[any]any
}

// GongGetSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GongGetSet[Type GongstructSet](stage *Stage) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case map[*Body]any:
		return any(&stage.Bodys).(*Type)
	case map[*Document]any:
		return any(&stage.Documents).(*Type)
	case map[*Docx]any:
		return any(&stage.Docxs).(*Type)
	case map[*File]any:
		return any(&stage.Files).(*Type)
	case map[*Node]any:
		return any(&stage.Nodes).(*Type)
	case map[*Paragraph]any:
		return any(&stage.Paragraphs).(*Type)
	case map[*ParagraphProperties]any:
		return any(&stage.ParagraphPropertiess).(*Type)
	case map[*ParagraphStyle]any:
		return any(&stage.ParagraphStyles).(*Type)
	case map[*Rune]any:
		return any(&stage.Runes).(*Type)
	case map[*RuneProperties]any:
		return any(&stage.RunePropertiess).(*Type)
	case map[*Table]any:
		return any(&stage.Tables).(*Type)
	case map[*TableColumn]any:
		return any(&stage.TableColumns).(*Type)
	case map[*TableProperties]any:
		return any(&stage.TablePropertiess).(*Type)
	case map[*TableRow]any:
		return any(&stage.TableRows).(*Type)
	case map[*TableStyle]any:
		return any(&stage.TableStyles).(*Type)
	case map[*Text]any:
		return any(&stage.Texts).(*Type)
	default:
		return nil
	}
}

// GongGetMap returns the map of staged Gonstruct instance by their name
// Can be usefull if names are unique
func GongGetMap[Type GongstructIF](stage *Stage) map[string]Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case *Body:
		return any(stage.Bodys_mapString).(map[string]Type)
	case *Document:
		return any(stage.Documents_mapString).(map[string]Type)
	case *Docx:
		return any(stage.Docxs_mapString).(map[string]Type)
	case *File:
		return any(stage.Files_mapString).(map[string]Type)
	case *Node:
		return any(stage.Nodes_mapString).(map[string]Type)
	case *Paragraph:
		return any(stage.Paragraphs_mapString).(map[string]Type)
	case *ParagraphProperties:
		return any(stage.ParagraphPropertiess_mapString).(map[string]Type)
	case *ParagraphStyle:
		return any(stage.ParagraphStyles_mapString).(map[string]Type)
	case *Rune:
		return any(stage.Runes_mapString).(map[string]Type)
	case *RuneProperties:
		return any(stage.RunePropertiess_mapString).(map[string]Type)
	case *Table:
		return any(stage.Tables_mapString).(map[string]Type)
	case *TableColumn:
		return any(stage.TableColumns_mapString).(map[string]Type)
	case *TableProperties:
		return any(stage.TablePropertiess_mapString).(map[string]Type)
	case *TableRow:
		return any(stage.TableRows_mapString).(map[string]Type)
	case *TableStyle:
		return any(stage.TableStyles_mapString).(map[string]Type)
	case *Text:
		return any(stage.Texts_mapString).(map[string]Type)
	default:
		return nil
	}
}

// GetGongstructInstancesSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSet[Type Gongstruct](stage *Stage) *map[*Type]struct{} {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case Body:
		return any(&stage.Bodys).(*map[*Type]struct{})
	case Document:
		return any(&stage.Documents).(*map[*Type]struct{})
	case Docx:
		return any(&stage.Docxs).(*map[*Type]struct{})
	case File:
		return any(&stage.Files).(*map[*Type]struct{})
	case Node:
		return any(&stage.Nodes).(*map[*Type]struct{})
	case Paragraph:
		return any(&stage.Paragraphs).(*map[*Type]struct{})
	case ParagraphProperties:
		return any(&stage.ParagraphPropertiess).(*map[*Type]struct{})
	case ParagraphStyle:
		return any(&stage.ParagraphStyles).(*map[*Type]struct{})
	case Rune:
		return any(&stage.Runes).(*map[*Type]struct{})
	case RuneProperties:
		return any(&stage.RunePropertiess).(*map[*Type]struct{})
	case Table:
		return any(&stage.Tables).(*map[*Type]struct{})
	case TableColumn:
		return any(&stage.TableColumns).(*map[*Type]struct{})
	case TableProperties:
		return any(&stage.TablePropertiess).(*map[*Type]struct{})
	case TableRow:
		return any(&stage.TableRows).(*map[*Type]struct{})
	case TableStyle:
		return any(&stage.TableStyles).(*map[*Type]struct{})
	case Text:
		return any(&stage.Texts).(*map[*Type]struct{})
	default:
		return nil
	}
}

// GetGongstructInstancesSetFromPointerType returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSetFromPointerType[Type PointerToGongstruct](stage *Stage) *map[Type]struct{} {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case *Body:
		return any(&stage.Bodys).(*map[Type]struct{})
	case *Document:
		return any(&stage.Documents).(*map[Type]struct{})
	case *Docx:
		return any(&stage.Docxs).(*map[Type]struct{})
	case *File:
		return any(&stage.Files).(*map[Type]struct{})
	case *Node:
		return any(&stage.Nodes).(*map[Type]struct{})
	case *Paragraph:
		return any(&stage.Paragraphs).(*map[Type]struct{})
	case *ParagraphProperties:
		return any(&stage.ParagraphPropertiess).(*map[Type]struct{})
	case *ParagraphStyle:
		return any(&stage.ParagraphStyles).(*map[Type]struct{})
	case *Rune:
		return any(&stage.Runes).(*map[Type]struct{})
	case *RuneProperties:
		return any(&stage.RunePropertiess).(*map[Type]struct{})
	case *Table:
		return any(&stage.Tables).(*map[Type]struct{})
	case *TableColumn:
		return any(&stage.TableColumns).(*map[Type]struct{})
	case *TableProperties:
		return any(&stage.TablePropertiess).(*map[Type]struct{})
	case *TableRow:
		return any(&stage.TableRows).(*map[Type]struct{})
	case *TableStyle:
		return any(&stage.TableStyles).(*map[Type]struct{})
	case *Text:
		return any(&stage.Texts).(*map[Type]struct{})
	default:
		return nil
	}
}

// GetGongstructInstancesMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesMap[Type Gongstruct](stage *Stage) *map[string]*Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case Body:
		return any(&stage.Bodys_mapString).(*map[string]*Type)
	case Document:
		return any(&stage.Documents_mapString).(*map[string]*Type)
	case Docx:
		return any(&stage.Docxs_mapString).(*map[string]*Type)
	case File:
		return any(&stage.Files_mapString).(*map[string]*Type)
	case Node:
		return any(&stage.Nodes_mapString).(*map[string]*Type)
	case Paragraph:
		return any(&stage.Paragraphs_mapString).(*map[string]*Type)
	case ParagraphProperties:
		return any(&stage.ParagraphPropertiess_mapString).(*map[string]*Type)
	case ParagraphStyle:
		return any(&stage.ParagraphStyles_mapString).(*map[string]*Type)
	case Rune:
		return any(&stage.Runes_mapString).(*map[string]*Type)
	case RuneProperties:
		return any(&stage.RunePropertiess_mapString).(*map[string]*Type)
	case Table:
		return any(&stage.Tables_mapString).(*map[string]*Type)
	case TableColumn:
		return any(&stage.TableColumns_mapString).(*map[string]*Type)
	case TableProperties:
		return any(&stage.TablePropertiess_mapString).(*map[string]*Type)
	case TableRow:
		return any(&stage.TableRows_mapString).(*map[string]*Type)
	case TableStyle:
		return any(&stage.TableStyles_mapString).(*map[string]*Type)
	case Text:
		return any(&stage.Texts_mapString).(*map[string]*Type)
	default:
		return nil
	}
}

// GetAssociationName is a generic function that returns an instance of Type
// where each association is filled with an instance whose name is the name of the association
//
// This function can be handy for generating navigation function that are refactorable
func GetAssociationName[Type Gongstruct]() *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for instance with special fields
	case Body:
		return any(&Body{
			// Initialisation of associations
			// field is initialized with an instance of Paragraph with the name of the field
			Paragraphs: []*Paragraph{{Name: "Paragraphs"}},
			// field is initialized with an instance of Table with the name of the field
			Tables: []*Table{{Name: "Tables"}},
			// field is initialized with an instance of Paragraph with the name of the field
			LastParagraph: &Paragraph{Name: "LastParagraph"},
		}).(*Type)
	case Document:
		return any(&Document{
			// Initialisation of associations
			// field is initialized with an instance of File with the name of the field
			File: &File{Name: "File"},
			// field is initialized with an instance of Node with the name of the field
			Root: &Node{Name: "Root"},
			// field is initialized with an instance of Body with the name of the field
			Body: &Body{Name: "Body"},
		}).(*Type)
	case Docx:
		return any(&Docx{
			// Initialisation of associations
			// field is initialized with an instance of File with the name of the field
			Files: []*File{{Name: "Files"}},
			// field is initialized with an instance of Document with the name of the field
			Document: &Document{Name: "Document"},
		}).(*Type)
	case File:
		return any(&File{
			// Initialisation of associations
		}).(*Type)
	case Node:
		return any(&Node{
			// Initialisation of associations
			// field is initialized with an instance of Node with the name of the field
			Nodes: []*Node{{Name: "Nodes"}},
		}).(*Type)
	case Paragraph:
		return any(&Paragraph{
			// Initialisation of associations
			// field is initialized with an instance of Node with the name of the field
			Node: &Node{Name: "Node"},
			// field is initialized with an instance of ParagraphProperties with the name of the field
			ParagraphProperties: &ParagraphProperties{Name: "ParagraphProperties"},
			// field is initialized with an instance of Rune with the name of the field
			Runes: []*Rune{{Name: "Runes"}},
			// field is initialized with an instance of Paragraph with the name of the field
			Next: &Paragraph{Name: "Next"},
			// field is initialized with an instance of Paragraph with the name of the field
			Previous: &Paragraph{Name: "Previous"},
			// field is initialized with an instance of Body with the name of the field
			EnclosingBody: &Body{Name: "EnclosingBody"},
			// field is initialized with an instance of TableColumn with the name of the field
			EnclosingTableColumn: &TableColumn{Name: "EnclosingTableColumn"},
		}).(*Type)
	case ParagraphProperties:
		return any(&ParagraphProperties{
			// Initialisation of associations
			// field is initialized with an instance of ParagraphStyle with the name of the field
			ParagraphStyle: &ParagraphStyle{Name: "ParagraphStyle"},
			// field is initialized with an instance of Node with the name of the field
			Node: &Node{Name: "Node"},
		}).(*Type)
	case ParagraphStyle:
		return any(&ParagraphStyle{
			// Initialisation of associations
			// field is initialized with an instance of Node with the name of the field
			Node: &Node{Name: "Node"},
		}).(*Type)
	case Rune:
		return any(&Rune{
			// Initialisation of associations
			// field is initialized with an instance of Node with the name of the field
			Node: &Node{Name: "Node"},
			// field is initialized with an instance of Text with the name of the field
			Text: &Text{Name: "Text"},
			// field is initialized with an instance of RuneProperties with the name of the field
			RuneProperties: &RuneProperties{Name: "RuneProperties"},
			// field is initialized with an instance of Paragraph with the name of the field
			EnclosingParagraph: &Paragraph{Name: "EnclosingParagraph"},
		}).(*Type)
	case RuneProperties:
		return any(&RuneProperties{
			// Initialisation of associations
			// field is initialized with an instance of Node with the name of the field
			Node: &Node{Name: "Node"},
		}).(*Type)
	case Table:
		return any(&Table{
			// Initialisation of associations
			// field is initialized with an instance of Node with the name of the field
			Node: &Node{Name: "Node"},
			// field is initialized with an instance of TableProperties with the name of the field
			TableProperties: &TableProperties{Name: "TableProperties"},
			// field is initialized with an instance of TableRow with the name of the field
			TableRows: []*TableRow{{Name: "TableRows"}},
		}).(*Type)
	case TableColumn:
		return any(&TableColumn{
			// Initialisation of associations
			// field is initialized with an instance of Node with the name of the field
			Node: &Node{Name: "Node"},
			// field is initialized with an instance of Paragraph with the name of the field
			Paragraphs: []*Paragraph{{Name: "Paragraphs"}},
		}).(*Type)
	case TableProperties:
		return any(&TableProperties{
			// Initialisation of associations
			// field is initialized with an instance of Node with the name of the field
			Node: &Node{Name: "Node"},
			// field is initialized with an instance of TableStyle with the name of the field
			TableStyle: &TableStyle{Name: "TableStyle"},
		}).(*Type)
	case TableRow:
		return any(&TableRow{
			// Initialisation of associations
			// field is initialized with an instance of Node with the name of the field
			Node: &Node{Name: "Node"},
			// field is initialized with an instance of TableColumn with the name of the field
			TableColumns: []*TableColumn{{Name: "TableColumns"}},
		}).(*Type)
	case TableStyle:
		return any(&TableStyle{
			// Initialisation of associations
			// field is initialized with an instance of Node with the name of the field
			Node: &Node{Name: "Node"},
		}).(*Type)
	case Text:
		return any(&Text{
			// Initialisation of associations
			// field is initialized with an instance of Node with the name of the field
			Node: &Node{Name: "Node"},
			// field is initialized with an instance of Rune with the name of the field
			EnclosingRune: &Rune{Name: "EnclosingRune"},
		}).(*Type)
	default:
		return nil
	}
}

// GetPointerReverseMap allows backtrack navigation of any Start.Fieldname
// associations (0..1) that is a pointer from one staged Gongstruct (type Start)
// instances to another (type End)
//
// The function provides a map with keys as instances of End and values to arrays of *Start
// the map is construed by iterating over all Start instances and populationg keys with End instances
// and values with slice of Start instances
func GetPointerReverseMap[Start, End Gongstruct](fieldname string, stage *Stage) map[*End][]*Start {

	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of Body
	case Body:
		switch fieldname {
		// insertion point for per direct association field
		case "LastParagraph":
			res := make(map[*Paragraph][]*Body)
			for body := range stage.Bodys {
				if body.LastParagraph != nil {
					paragraph_ := body.LastParagraph
					var bodys []*Body
					_, ok := res[paragraph_]
					if ok {
						bodys = res[paragraph_]
					} else {
						bodys = make([]*Body, 0)
					}
					bodys = append(bodys, body)
					res[paragraph_] = bodys
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Document
	case Document:
		switch fieldname {
		// insertion point for per direct association field
		case "File":
			res := make(map[*File][]*Document)
			for document := range stage.Documents {
				if document.File != nil {
					file_ := document.File
					var documents []*Document
					_, ok := res[file_]
					if ok {
						documents = res[file_]
					} else {
						documents = make([]*Document, 0)
					}
					documents = append(documents, document)
					res[file_] = documents
				}
			}
			return any(res).(map[*End][]*Start)
		case "Root":
			res := make(map[*Node][]*Document)
			for document := range stage.Documents {
				if document.Root != nil {
					node_ := document.Root
					var documents []*Document
					_, ok := res[node_]
					if ok {
						documents = res[node_]
					} else {
						documents = make([]*Document, 0)
					}
					documents = append(documents, document)
					res[node_] = documents
				}
			}
			return any(res).(map[*End][]*Start)
		case "Body":
			res := make(map[*Body][]*Document)
			for document := range stage.Documents {
				if document.Body != nil {
					body_ := document.Body
					var documents []*Document
					_, ok := res[body_]
					if ok {
						documents = res[body_]
					} else {
						documents = make([]*Document, 0)
					}
					documents = append(documents, document)
					res[body_] = documents
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Docx
	case Docx:
		switch fieldname {
		// insertion point for per direct association field
		case "Document":
			res := make(map[*Document][]*Docx)
			for docx := range stage.Docxs {
				if docx.Document != nil {
					document_ := docx.Document
					var docxs []*Docx
					_, ok := res[document_]
					if ok {
						docxs = res[document_]
					} else {
						docxs = make([]*Docx, 0)
					}
					docxs = append(docxs, docx)
					res[document_] = docxs
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of File
	case File:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Node
	case Node:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Paragraph
	case Paragraph:
		switch fieldname {
		// insertion point for per direct association field
		case "Node":
			res := make(map[*Node][]*Paragraph)
			for paragraph := range stage.Paragraphs {
				if paragraph.Node != nil {
					node_ := paragraph.Node
					var paragraphs []*Paragraph
					_, ok := res[node_]
					if ok {
						paragraphs = res[node_]
					} else {
						paragraphs = make([]*Paragraph, 0)
					}
					paragraphs = append(paragraphs, paragraph)
					res[node_] = paragraphs
				}
			}
			return any(res).(map[*End][]*Start)
		case "ParagraphProperties":
			res := make(map[*ParagraphProperties][]*Paragraph)
			for paragraph := range stage.Paragraphs {
				if paragraph.ParagraphProperties != nil {
					paragraphproperties_ := paragraph.ParagraphProperties
					var paragraphs []*Paragraph
					_, ok := res[paragraphproperties_]
					if ok {
						paragraphs = res[paragraphproperties_]
					} else {
						paragraphs = make([]*Paragraph, 0)
					}
					paragraphs = append(paragraphs, paragraph)
					res[paragraphproperties_] = paragraphs
				}
			}
			return any(res).(map[*End][]*Start)
		case "Next":
			res := make(map[*Paragraph][]*Paragraph)
			for paragraph := range stage.Paragraphs {
				if paragraph.Next != nil {
					paragraph_ := paragraph.Next
					var paragraphs []*Paragraph
					_, ok := res[paragraph_]
					if ok {
						paragraphs = res[paragraph_]
					} else {
						paragraphs = make([]*Paragraph, 0)
					}
					paragraphs = append(paragraphs, paragraph)
					res[paragraph_] = paragraphs
				}
			}
			return any(res).(map[*End][]*Start)
		case "Previous":
			res := make(map[*Paragraph][]*Paragraph)
			for paragraph := range stage.Paragraphs {
				if paragraph.Previous != nil {
					paragraph_ := paragraph.Previous
					var paragraphs []*Paragraph
					_, ok := res[paragraph_]
					if ok {
						paragraphs = res[paragraph_]
					} else {
						paragraphs = make([]*Paragraph, 0)
					}
					paragraphs = append(paragraphs, paragraph)
					res[paragraph_] = paragraphs
				}
			}
			return any(res).(map[*End][]*Start)
		case "EnclosingBody":
			res := make(map[*Body][]*Paragraph)
			for paragraph := range stage.Paragraphs {
				if paragraph.EnclosingBody != nil {
					body_ := paragraph.EnclosingBody
					var paragraphs []*Paragraph
					_, ok := res[body_]
					if ok {
						paragraphs = res[body_]
					} else {
						paragraphs = make([]*Paragraph, 0)
					}
					paragraphs = append(paragraphs, paragraph)
					res[body_] = paragraphs
				}
			}
			return any(res).(map[*End][]*Start)
		case "EnclosingTableColumn":
			res := make(map[*TableColumn][]*Paragraph)
			for paragraph := range stage.Paragraphs {
				if paragraph.EnclosingTableColumn != nil {
					tablecolumn_ := paragraph.EnclosingTableColumn
					var paragraphs []*Paragraph
					_, ok := res[tablecolumn_]
					if ok {
						paragraphs = res[tablecolumn_]
					} else {
						paragraphs = make([]*Paragraph, 0)
					}
					paragraphs = append(paragraphs, paragraph)
					res[tablecolumn_] = paragraphs
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ParagraphProperties
	case ParagraphProperties:
		switch fieldname {
		// insertion point for per direct association field
		case "ParagraphStyle":
			res := make(map[*ParagraphStyle][]*ParagraphProperties)
			for paragraphproperties := range stage.ParagraphPropertiess {
				if paragraphproperties.ParagraphStyle != nil {
					paragraphstyle_ := paragraphproperties.ParagraphStyle
					var paragraphpropertiess []*ParagraphProperties
					_, ok := res[paragraphstyle_]
					if ok {
						paragraphpropertiess = res[paragraphstyle_]
					} else {
						paragraphpropertiess = make([]*ParagraphProperties, 0)
					}
					paragraphpropertiess = append(paragraphpropertiess, paragraphproperties)
					res[paragraphstyle_] = paragraphpropertiess
				}
			}
			return any(res).(map[*End][]*Start)
		case "Node":
			res := make(map[*Node][]*ParagraphProperties)
			for paragraphproperties := range stage.ParagraphPropertiess {
				if paragraphproperties.Node != nil {
					node_ := paragraphproperties.Node
					var paragraphpropertiess []*ParagraphProperties
					_, ok := res[node_]
					if ok {
						paragraphpropertiess = res[node_]
					} else {
						paragraphpropertiess = make([]*ParagraphProperties, 0)
					}
					paragraphpropertiess = append(paragraphpropertiess, paragraphproperties)
					res[node_] = paragraphpropertiess
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ParagraphStyle
	case ParagraphStyle:
		switch fieldname {
		// insertion point for per direct association field
		case "Node":
			res := make(map[*Node][]*ParagraphStyle)
			for paragraphstyle := range stage.ParagraphStyles {
				if paragraphstyle.Node != nil {
					node_ := paragraphstyle.Node
					var paragraphstyles []*ParagraphStyle
					_, ok := res[node_]
					if ok {
						paragraphstyles = res[node_]
					} else {
						paragraphstyles = make([]*ParagraphStyle, 0)
					}
					paragraphstyles = append(paragraphstyles, paragraphstyle)
					res[node_] = paragraphstyles
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Rune
	case Rune:
		switch fieldname {
		// insertion point for per direct association field
		case "Node":
			res := make(map[*Node][]*Rune)
			for rune := range stage.Runes {
				if rune.Node != nil {
					node_ := rune.Node
					var runes []*Rune
					_, ok := res[node_]
					if ok {
						runes = res[node_]
					} else {
						runes = make([]*Rune, 0)
					}
					runes = append(runes, rune)
					res[node_] = runes
				}
			}
			return any(res).(map[*End][]*Start)
		case "Text":
			res := make(map[*Text][]*Rune)
			for rune := range stage.Runes {
				if rune.Text != nil {
					text_ := rune.Text
					var runes []*Rune
					_, ok := res[text_]
					if ok {
						runes = res[text_]
					} else {
						runes = make([]*Rune, 0)
					}
					runes = append(runes, rune)
					res[text_] = runes
				}
			}
			return any(res).(map[*End][]*Start)
		case "RuneProperties":
			res := make(map[*RuneProperties][]*Rune)
			for rune := range stage.Runes {
				if rune.RuneProperties != nil {
					runeproperties_ := rune.RuneProperties
					var runes []*Rune
					_, ok := res[runeproperties_]
					if ok {
						runes = res[runeproperties_]
					} else {
						runes = make([]*Rune, 0)
					}
					runes = append(runes, rune)
					res[runeproperties_] = runes
				}
			}
			return any(res).(map[*End][]*Start)
		case "EnclosingParagraph":
			res := make(map[*Paragraph][]*Rune)
			for rune := range stage.Runes {
				if rune.EnclosingParagraph != nil {
					paragraph_ := rune.EnclosingParagraph
					var runes []*Rune
					_, ok := res[paragraph_]
					if ok {
						runes = res[paragraph_]
					} else {
						runes = make([]*Rune, 0)
					}
					runes = append(runes, rune)
					res[paragraph_] = runes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of RuneProperties
	case RuneProperties:
		switch fieldname {
		// insertion point for per direct association field
		case "Node":
			res := make(map[*Node][]*RuneProperties)
			for runeproperties := range stage.RunePropertiess {
				if runeproperties.Node != nil {
					node_ := runeproperties.Node
					var runepropertiess []*RuneProperties
					_, ok := res[node_]
					if ok {
						runepropertiess = res[node_]
					} else {
						runepropertiess = make([]*RuneProperties, 0)
					}
					runepropertiess = append(runepropertiess, runeproperties)
					res[node_] = runepropertiess
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Table
	case Table:
		switch fieldname {
		// insertion point for per direct association field
		case "Node":
			res := make(map[*Node][]*Table)
			for table := range stage.Tables {
				if table.Node != nil {
					node_ := table.Node
					var tables []*Table
					_, ok := res[node_]
					if ok {
						tables = res[node_]
					} else {
						tables = make([]*Table, 0)
					}
					tables = append(tables, table)
					res[node_] = tables
				}
			}
			return any(res).(map[*End][]*Start)
		case "TableProperties":
			res := make(map[*TableProperties][]*Table)
			for table := range stage.Tables {
				if table.TableProperties != nil {
					tableproperties_ := table.TableProperties
					var tables []*Table
					_, ok := res[tableproperties_]
					if ok {
						tables = res[tableproperties_]
					} else {
						tables = make([]*Table, 0)
					}
					tables = append(tables, table)
					res[tableproperties_] = tables
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of TableColumn
	case TableColumn:
		switch fieldname {
		// insertion point for per direct association field
		case "Node":
			res := make(map[*Node][]*TableColumn)
			for tablecolumn := range stage.TableColumns {
				if tablecolumn.Node != nil {
					node_ := tablecolumn.Node
					var tablecolumns []*TableColumn
					_, ok := res[node_]
					if ok {
						tablecolumns = res[node_]
					} else {
						tablecolumns = make([]*TableColumn, 0)
					}
					tablecolumns = append(tablecolumns, tablecolumn)
					res[node_] = tablecolumns
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of TableProperties
	case TableProperties:
		switch fieldname {
		// insertion point for per direct association field
		case "Node":
			res := make(map[*Node][]*TableProperties)
			for tableproperties := range stage.TablePropertiess {
				if tableproperties.Node != nil {
					node_ := tableproperties.Node
					var tablepropertiess []*TableProperties
					_, ok := res[node_]
					if ok {
						tablepropertiess = res[node_]
					} else {
						tablepropertiess = make([]*TableProperties, 0)
					}
					tablepropertiess = append(tablepropertiess, tableproperties)
					res[node_] = tablepropertiess
				}
			}
			return any(res).(map[*End][]*Start)
		case "TableStyle":
			res := make(map[*TableStyle][]*TableProperties)
			for tableproperties := range stage.TablePropertiess {
				if tableproperties.TableStyle != nil {
					tablestyle_ := tableproperties.TableStyle
					var tablepropertiess []*TableProperties
					_, ok := res[tablestyle_]
					if ok {
						tablepropertiess = res[tablestyle_]
					} else {
						tablepropertiess = make([]*TableProperties, 0)
					}
					tablepropertiess = append(tablepropertiess, tableproperties)
					res[tablestyle_] = tablepropertiess
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of TableRow
	case TableRow:
		switch fieldname {
		// insertion point for per direct association field
		case "Node":
			res := make(map[*Node][]*TableRow)
			for tablerow := range stage.TableRows {
				if tablerow.Node != nil {
					node_ := tablerow.Node
					var tablerows []*TableRow
					_, ok := res[node_]
					if ok {
						tablerows = res[node_]
					} else {
						tablerows = make([]*TableRow, 0)
					}
					tablerows = append(tablerows, tablerow)
					res[node_] = tablerows
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of TableStyle
	case TableStyle:
		switch fieldname {
		// insertion point for per direct association field
		case "Node":
			res := make(map[*Node][]*TableStyle)
			for tablestyle := range stage.TableStyles {
				if tablestyle.Node != nil {
					node_ := tablestyle.Node
					var tablestyles []*TableStyle
					_, ok := res[node_]
					if ok {
						tablestyles = res[node_]
					} else {
						tablestyles = make([]*TableStyle, 0)
					}
					tablestyles = append(tablestyles, tablestyle)
					res[node_] = tablestyles
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Text
	case Text:
		switch fieldname {
		// insertion point for per direct association field
		case "Node":
			res := make(map[*Node][]*Text)
			for text := range stage.Texts {
				if text.Node != nil {
					node_ := text.Node
					var texts []*Text
					_, ok := res[node_]
					if ok {
						texts = res[node_]
					} else {
						texts = make([]*Text, 0)
					}
					texts = append(texts, text)
					res[node_] = texts
				}
			}
			return any(res).(map[*End][]*Start)
		case "EnclosingRune":
			res := make(map[*Rune][]*Text)
			for text := range stage.Texts {
				if text.EnclosingRune != nil {
					rune_ := text.EnclosingRune
					var texts []*Text
					_, ok := res[rune_]
					if ok {
						texts = res[rune_]
					} else {
						texts = make([]*Text, 0)
					}
					texts = append(texts, text)
					res[rune_] = texts
				}
			}
			return any(res).(map[*End][]*Start)
		}
	}
	return nil
}

// GetSliceOfPointersReverseMap allows backtrack navigation of any Start.Fieldname
// associations (0..N) between one staged Gongstruct instances and many others
//
// The function provides a map with keys as instances of End and values to *Start instances
// the map is construed by iterating over all Start instances and populating keys with End instances
// and values with the Start instances
func GetSliceOfPointersReverseMap[Start, End Gongstruct](fieldname string, stage *Stage) map[*End][]*Start {

	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of Body
	case Body:
		switch fieldname {
		// insertion point for per direct association field
		case "Paragraphs":
			res := make(map[*Paragraph][]*Body)
			for body := range stage.Bodys {
				for _, paragraph_ := range body.Paragraphs {
					res[paragraph_] = append(res[paragraph_], body)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Tables":
			res := make(map[*Table][]*Body)
			for body := range stage.Bodys {
				for _, table_ := range body.Tables {
					res[table_] = append(res[table_], body)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Document
	case Document:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Docx
	case Docx:
		switch fieldname {
		// insertion point for per direct association field
		case "Files":
			res := make(map[*File][]*Docx)
			for docx := range stage.Docxs {
				for _, file_ := range docx.Files {
					res[file_] = append(res[file_], docx)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of File
	case File:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Node
	case Node:
		switch fieldname {
		// insertion point for per direct association field
		case "Nodes":
			res := make(map[*Node][]*Node)
			for node := range stage.Nodes {
				for _, node_ := range node.Nodes {
					res[node_] = append(res[node_], node)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Paragraph
	case Paragraph:
		switch fieldname {
		// insertion point for per direct association field
		case "Runes":
			res := make(map[*Rune][]*Paragraph)
			for paragraph := range stage.Paragraphs {
				for _, rune_ := range paragraph.Runes {
					res[rune_] = append(res[rune_], paragraph)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ParagraphProperties
	case ParagraphProperties:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ParagraphStyle
	case ParagraphStyle:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Rune
	case Rune:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of RuneProperties
	case RuneProperties:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Table
	case Table:
		switch fieldname {
		// insertion point for per direct association field
		case "TableRows":
			res := make(map[*TableRow][]*Table)
			for table := range stage.Tables {
				for _, tablerow_ := range table.TableRows {
					res[tablerow_] = append(res[tablerow_], table)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of TableColumn
	case TableColumn:
		switch fieldname {
		// insertion point for per direct association field
		case "Paragraphs":
			res := make(map[*Paragraph][]*TableColumn)
			for tablecolumn := range stage.TableColumns {
				for _, paragraph_ := range tablecolumn.Paragraphs {
					res[paragraph_] = append(res[paragraph_], tablecolumn)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of TableProperties
	case TableProperties:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TableRow
	case TableRow:
		switch fieldname {
		// insertion point for per direct association field
		case "TableColumns":
			res := make(map[*TableColumn][]*TableRow)
			for tablerow := range stage.TableRows {
				for _, tablecolumn_ := range tablerow.TableColumns {
					res[tablecolumn_] = append(res[tablecolumn_], tablerow)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of TableStyle
	case TableStyle:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Text
	case Text:
		switch fieldname {
		// insertion point for per direct association field
		}
	}
	return nil
}

// GetPointerToGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetPointerToGongstructName[Type GongstructIF]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *Body:
		res = "Body"
	case *Document:
		res = "Document"
	case *Docx:
		res = "Docx"
	case *File:
		res = "File"
	case *Node:
		res = "Node"
	case *Paragraph:
		res = "Paragraph"
	case *ParagraphProperties:
		res = "ParagraphProperties"
	case *ParagraphStyle:
		res = "ParagraphStyle"
	case *Rune:
		res = "Rune"
	case *RuneProperties:
		res = "RuneProperties"
	case *Table:
		res = "Table"
	case *TableColumn:
		res = "TableColumn"
	case *TableProperties:
		res = "TableProperties"
	case *TableRow:
		res = "TableRow"
	case *TableStyle:
		res = "TableStyle"
	case *Text:
		res = "Text"
	}
	return res
}

type ReverseField struct {
	GongstructName string
	Fieldname      string
}

func GetReverseFields[Type GongstructIF]() (res []ReverseField) {

	res = make([]ReverseField, 0)

	var ret Type

	switch any(ret).(type) {

	// insertion point for generic get gongstruct name
	case *Body:
		var rf ReverseField
		_ = rf
	case *Document:
		var rf ReverseField
		_ = rf
	case *Docx:
		var rf ReverseField
		_ = rf
	case *File:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Docx"
		rf.Fieldname = "Files"
		res = append(res, rf)
	case *Node:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Node"
		rf.Fieldname = "Nodes"
		res = append(res, rf)
	case *Paragraph:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Body"
		rf.Fieldname = "Paragraphs"
		res = append(res, rf)
		rf.GongstructName = "TableColumn"
		rf.Fieldname = "Paragraphs"
		res = append(res, rf)
	case *ParagraphProperties:
		var rf ReverseField
		_ = rf
	case *ParagraphStyle:
		var rf ReverseField
		_ = rf
	case *Rune:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Paragraph"
		rf.Fieldname = "Runes"
		res = append(res, rf)
	case *RuneProperties:
		var rf ReverseField
		_ = rf
	case *Table:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Body"
		rf.Fieldname = "Tables"
		res = append(res, rf)
	case *TableColumn:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "TableRow"
		rf.Fieldname = "TableColumns"
		res = append(res, rf)
	case *TableProperties:
		var rf ReverseField
		_ = rf
	case *TableRow:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Table"
		rf.Fieldname = "TableRows"
		res = append(res, rf)
	case *TableStyle:
		var rf ReverseField
		_ = rf
	case *Text:
		var rf ReverseField
		_ = rf
	}
	return
}

// insertion point for get fields header method
func (body *Body) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "Paragraphs",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Paragraph",
		},
		{
			Name:                 "Tables",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Table",
		},
		{
			Name:                 "LastParagraph",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Paragraph",
		},
	}
	return
}

func (document *Document) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "File",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "File",
		},
		{
			Name:                 "Root",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Node",
		},
		{
			Name:                 "Body",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Body",
		},
	}
	return
}

func (docx *Docx) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "Files",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "File",
		},
		{
			Name:                 "Document",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Document",
		},
	}
	return
}

func (file *File) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (node *Node) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "Nodes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Node",
		},
	}
	return
}

func (paragraph *Paragraph) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Content",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "Node",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Node",
		},
		{
			Name:                 "ParagraphProperties",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ParagraphProperties",
		},
		{
			Name:                 "Runes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Rune",
		},
		{
			Name:               "CollatedText",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "Next",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Paragraph",
		},
		{
			Name:                 "Previous",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Paragraph",
		},
		{
			Name:                 "EnclosingBody",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Body",
		},
		{
			Name:                 "EnclosingTableColumn",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "TableColumn",
		},
	}
	return
}

func (paragraphproperties *ParagraphProperties) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Content",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "ParagraphStyle",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ParagraphStyle",
		},
		{
			Name:                 "Node",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Node",
		},
	}
	return
}

func (paragraphstyle *ParagraphStyle) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "Node",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Node",
		},
		{
			Name:               "Content",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "ValAttr",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (rune *Rune) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Content",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "Node",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Node",
		},
		{
			Name:                 "Text",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Text",
		},
		{
			Name:                 "RuneProperties",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "RuneProperties",
		},
		{
			Name:                 "EnclosingParagraph",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Paragraph",
		},
	}
	return
}

func (runeproperties *RuneProperties) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "Node",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Node",
		},
		{
			Name:               "IsBold",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "IsStrike",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "IsItalic",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Content",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (table *Table) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "Node",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Node",
		},
		{
			Name:               "Content",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "TableProperties",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "TableProperties",
		},
		{
			Name:                 "TableRows",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "TableRow",
		},
	}
	return
}

func (tablecolumn *TableColumn) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Content",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "Node",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Node",
		},
		{
			Name:                 "Paragraphs",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Paragraph",
		},
	}
	return
}

func (tableproperties *TableProperties) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "Node",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Node",
		},
		{
			Name:               "Content",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "TableStyle",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "TableStyle",
		},
	}
	return
}

func (tablerow *TableRow) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Content",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "Node",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Node",
		},
		{
			Name:                 "TableColumns",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "TableColumn",
		},
	}
	return
}

func (tablestyle *TableStyle) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "Node",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Node",
		},
		{
			Name:               "Content",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Val",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (text *Text) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Content",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "Node",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Node",
		},
		{
			Name:               "PreserveWhiteSpace",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "EnclosingRune",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Rune",
		},
	}
	return
}

// GetFieldsFromPointer return the array of the fields
func GetFieldsFromPointer[Type PointerToGongstruct]() (res []GongFieldHeader) {

	var ret Type
	return ret.GongGetFieldHeaders()
}

type GongFieldValueType string

const (
	GongFieldValueTypeInt             GongFieldValueType = "GongFieldValueTypeInt"
	GongFieldValueTypeFloat           GongFieldValueType = "GongFieldValueTypeFloat"
	GongFieldValueTypeBool            GongFieldValueType = "GongFieldValueTypeBool"
	GongFieldValueTypeString          GongFieldValueType = "GongFieldValueTypeString"
	GongFieldValueTypeBasicKind       GongFieldValueType = "GongFieldValueTypeBasicKind"
	GongFieldValueTypePointer         GongFieldValueType = "GongFieldValueTypePointer"
	GongFieldValueTypeSliceOfPointers GongFieldValueType = "GongFieldValueTypeSliceOfPointers"
)

type GongFieldValue struct {
	GongFieldValueType
	valueString string
	valueInt    int
	valueFloat  float64
	valueBool   bool

	// in case of a pointer, the ID of the pointed element
	// in case of a slice of pointers, the IDs, separated by semi columbs
	ids string
}

type GongFieldHeader struct {
	Name string
	GongFieldValueType
	TargetGongstructName string
}

func (gongValueField *GongFieldValue) GetValueString() string {
	return gongValueField.valueString
}

func (gongValueField *GongFieldValue) GetValueInt() int {
	return gongValueField.valueInt
}

func (gongValueField *GongFieldValue) GetValueFloat() float64 {
	return gongValueField.valueFloat
}

func (gongValueField *GongFieldValue) GetValueBool() bool {
	return gongValueField.valueBool
}

// insertion point for generic get gongstruct field value
func (body *Body) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = body.Name
	case "Paragraphs":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range body.Paragraphs {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "Tables":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range body.Tables {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "LastParagraph":
		res.GongFieldValueType = GongFieldValueTypePointer
		if body.LastParagraph != nil {
			res.valueString = body.LastParagraph.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, body.LastParagraph))
		}
	}
	return
}
func (document *Document) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = document.Name
	case "File":
		res.GongFieldValueType = GongFieldValueTypePointer
		if document.File != nil {
			res.valueString = document.File.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, document.File))
		}
	case "Root":
		res.GongFieldValueType = GongFieldValueTypePointer
		if document.Root != nil {
			res.valueString = document.Root.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, document.Root))
		}
	case "Body":
		res.GongFieldValueType = GongFieldValueTypePointer
		if document.Body != nil {
			res.valueString = document.Body.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, document.Body))
		}
	}
	return
}
func (docx *Docx) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = docx.Name
	case "Files":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range docx.Files {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "Document":
		res.GongFieldValueType = GongFieldValueTypePointer
		if docx.Document != nil {
			res.valueString = docx.Document.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, docx.Document))
		}
	}
	return
}
func (file *File) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = file.Name
	}
	return
}
func (node *Node) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = node.Name
	case "Nodes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range node.Nodes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	}
	return
}
func (paragraph *Paragraph) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = paragraph.Name
	case "Content":
		res.valueString = paragraph.Content
	case "Node":
		res.GongFieldValueType = GongFieldValueTypePointer
		if paragraph.Node != nil {
			res.valueString = paragraph.Node.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, paragraph.Node))
		}
	case "ParagraphProperties":
		res.GongFieldValueType = GongFieldValueTypePointer
		if paragraph.ParagraphProperties != nil {
			res.valueString = paragraph.ParagraphProperties.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, paragraph.ParagraphProperties))
		}
	case "Runes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range paragraph.Runes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "CollatedText":
		res.valueString = paragraph.CollatedText
	case "Next":
		res.GongFieldValueType = GongFieldValueTypePointer
		if paragraph.Next != nil {
			res.valueString = paragraph.Next.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, paragraph.Next))
		}
	case "Previous":
		res.GongFieldValueType = GongFieldValueTypePointer
		if paragraph.Previous != nil {
			res.valueString = paragraph.Previous.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, paragraph.Previous))
		}
	case "EnclosingBody":
		res.GongFieldValueType = GongFieldValueTypePointer
		if paragraph.EnclosingBody != nil {
			res.valueString = paragraph.EnclosingBody.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, paragraph.EnclosingBody))
		}
	case "EnclosingTableColumn":
		res.GongFieldValueType = GongFieldValueTypePointer
		if paragraph.EnclosingTableColumn != nil {
			res.valueString = paragraph.EnclosingTableColumn.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, paragraph.EnclosingTableColumn))
		}
	}
	return
}
func (paragraphproperties *ParagraphProperties) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = paragraphproperties.Name
	case "Content":
		res.valueString = paragraphproperties.Content
	case "ParagraphStyle":
		res.GongFieldValueType = GongFieldValueTypePointer
		if paragraphproperties.ParagraphStyle != nil {
			res.valueString = paragraphproperties.ParagraphStyle.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, paragraphproperties.ParagraphStyle))
		}
	case "Node":
		res.GongFieldValueType = GongFieldValueTypePointer
		if paragraphproperties.Node != nil {
			res.valueString = paragraphproperties.Node.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, paragraphproperties.Node))
		}
	}
	return
}
func (paragraphstyle *ParagraphStyle) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = paragraphstyle.Name
	case "Node":
		res.GongFieldValueType = GongFieldValueTypePointer
		if paragraphstyle.Node != nil {
			res.valueString = paragraphstyle.Node.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, paragraphstyle.Node))
		}
	case "Content":
		res.valueString = paragraphstyle.Content
	case "ValAttr":
		res.valueString = paragraphstyle.ValAttr
	}
	return
}
func (rune *Rune) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = rune.Name
	case "Content":
		res.valueString = rune.Content
	case "Node":
		res.GongFieldValueType = GongFieldValueTypePointer
		if rune.Node != nil {
			res.valueString = rune.Node.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, rune.Node))
		}
	case "Text":
		res.GongFieldValueType = GongFieldValueTypePointer
		if rune.Text != nil {
			res.valueString = rune.Text.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, rune.Text))
		}
	case "RuneProperties":
		res.GongFieldValueType = GongFieldValueTypePointer
		if rune.RuneProperties != nil {
			res.valueString = rune.RuneProperties.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, rune.RuneProperties))
		}
	case "EnclosingParagraph":
		res.GongFieldValueType = GongFieldValueTypePointer
		if rune.EnclosingParagraph != nil {
			res.valueString = rune.EnclosingParagraph.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, rune.EnclosingParagraph))
		}
	}
	return
}
func (runeproperties *RuneProperties) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = runeproperties.Name
	case "Node":
		res.GongFieldValueType = GongFieldValueTypePointer
		if runeproperties.Node != nil {
			res.valueString = runeproperties.Node.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, runeproperties.Node))
		}
	case "IsBold":
		res.valueString = fmt.Sprintf("%t", runeproperties.IsBold)
		res.valueBool = runeproperties.IsBold
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsStrike":
		res.valueString = fmt.Sprintf("%t", runeproperties.IsStrike)
		res.valueBool = runeproperties.IsStrike
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsItalic":
		res.valueString = fmt.Sprintf("%t", runeproperties.IsItalic)
		res.valueBool = runeproperties.IsItalic
		res.GongFieldValueType = GongFieldValueTypeBool
	case "Content":
		res.valueString = runeproperties.Content
	}
	return
}
func (table *Table) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = table.Name
	case "Node":
		res.GongFieldValueType = GongFieldValueTypePointer
		if table.Node != nil {
			res.valueString = table.Node.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, table.Node))
		}
	case "Content":
		res.valueString = table.Content
	case "TableProperties":
		res.GongFieldValueType = GongFieldValueTypePointer
		if table.TableProperties != nil {
			res.valueString = table.TableProperties.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, table.TableProperties))
		}
	case "TableRows":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range table.TableRows {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	}
	return
}
func (tablecolumn *TableColumn) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = tablecolumn.Name
	case "Content":
		res.valueString = tablecolumn.Content
	case "Node":
		res.GongFieldValueType = GongFieldValueTypePointer
		if tablecolumn.Node != nil {
			res.valueString = tablecolumn.Node.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, tablecolumn.Node))
		}
	case "Paragraphs":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range tablecolumn.Paragraphs {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	}
	return
}
func (tableproperties *TableProperties) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = tableproperties.Name
	case "Node":
		res.GongFieldValueType = GongFieldValueTypePointer
		if tableproperties.Node != nil {
			res.valueString = tableproperties.Node.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, tableproperties.Node))
		}
	case "Content":
		res.valueString = tableproperties.Content
	case "TableStyle":
		res.GongFieldValueType = GongFieldValueTypePointer
		if tableproperties.TableStyle != nil {
			res.valueString = tableproperties.TableStyle.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, tableproperties.TableStyle))
		}
	}
	return
}
func (tablerow *TableRow) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = tablerow.Name
	case "Content":
		res.valueString = tablerow.Content
	case "Node":
		res.GongFieldValueType = GongFieldValueTypePointer
		if tablerow.Node != nil {
			res.valueString = tablerow.Node.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, tablerow.Node))
		}
	case "TableColumns":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range tablerow.TableColumns {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	}
	return
}
func (tablestyle *TableStyle) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = tablestyle.Name
	case "Node":
		res.GongFieldValueType = GongFieldValueTypePointer
		if tablestyle.Node != nil {
			res.valueString = tablestyle.Node.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, tablestyle.Node))
		}
	case "Content":
		res.valueString = tablestyle.Content
	case "Val":
		res.valueString = tablestyle.Val
	}
	return
}
func (text *Text) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = text.Name
	case "Content":
		res.valueString = text.Content
	case "Node":
		res.GongFieldValueType = GongFieldValueTypePointer
		if text.Node != nil {
			res.valueString = text.Node.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, text.Node))
		}
	case "PreserveWhiteSpace":
		res.valueString = fmt.Sprintf("%t", text.PreserveWhiteSpace)
		res.valueBool = text.PreserveWhiteSpace
		res.GongFieldValueType = GongFieldValueTypeBool
	case "EnclosingRune":
		res.GongFieldValueType = GongFieldValueTypePointer
		if text.EnclosingRune != nil {
			res.valueString = text.EnclosingRune.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, text.EnclosingRune))
		}
	}
	return
}
func GetFieldStringValueFromPointer(instance GongstructIF, fieldName string, stage *Stage) (res GongFieldValue) {

	res = instance.GongGetFieldValue(fieldName, stage)
	return
}

// insertion point for generic set gongstruct field value
func (body *Body) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		body.Name = value.GetValueString()
	case "Paragraphs":
		body.Paragraphs = make([]*Paragraph, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Paragraphs {
					if stage.ParagraphMap_Staged_Order[__instance__] == uint(id) {
						body.Paragraphs = append(body.Paragraphs, __instance__)
						break
					}
				}
			}
		}
	case "Tables":
		body.Tables = make([]*Table, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Tables {
					if stage.TableMap_Staged_Order[__instance__] == uint(id) {
						body.Tables = append(body.Tables, __instance__)
						break
					}
				}
			}
		}
	case "LastParagraph":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			body.LastParagraph = nil
			for __instance__ := range stage.Paragraphs {
				if stage.ParagraphMap_Staged_Order[__instance__] == uint(id) {
					body.LastParagraph = __instance__
					break
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (document *Document) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		document.Name = value.GetValueString()
	case "File":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			document.File = nil
			for __instance__ := range stage.Files {
				if stage.FileMap_Staged_Order[__instance__] == uint(id) {
					document.File = __instance__
					break
				}
			}
		}
	case "Root":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			document.Root = nil
			for __instance__ := range stage.Nodes {
				if stage.NodeMap_Staged_Order[__instance__] == uint(id) {
					document.Root = __instance__
					break
				}
			}
		}
	case "Body":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			document.Body = nil
			for __instance__ := range stage.Bodys {
				if stage.BodyMap_Staged_Order[__instance__] == uint(id) {
					document.Body = __instance__
					break
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (docx *Docx) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		docx.Name = value.GetValueString()
	case "Files":
		docx.Files = make([]*File, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Files {
					if stage.FileMap_Staged_Order[__instance__] == uint(id) {
						docx.Files = append(docx.Files, __instance__)
						break
					}
				}
			}
		}
	case "Document":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			docx.Document = nil
			for __instance__ := range stage.Documents {
				if stage.DocumentMap_Staged_Order[__instance__] == uint(id) {
					docx.Document = __instance__
					break
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (file *File) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		file.Name = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (node *Node) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		node.Name = value.GetValueString()
	case "Nodes":
		node.Nodes = make([]*Node, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Nodes {
					if stage.NodeMap_Staged_Order[__instance__] == uint(id) {
						node.Nodes = append(node.Nodes, __instance__)
						break
					}
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (paragraph *Paragraph) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		paragraph.Name = value.GetValueString()
	case "Content":
		paragraph.Content = value.GetValueString()
	case "Node":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			paragraph.Node = nil
			for __instance__ := range stage.Nodes {
				if stage.NodeMap_Staged_Order[__instance__] == uint(id) {
					paragraph.Node = __instance__
					break
				}
			}
		}
	case "ParagraphProperties":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			paragraph.ParagraphProperties = nil
			for __instance__ := range stage.ParagraphPropertiess {
				if stage.ParagraphPropertiesMap_Staged_Order[__instance__] == uint(id) {
					paragraph.ParagraphProperties = __instance__
					break
				}
			}
		}
	case "Runes":
		paragraph.Runes = make([]*Rune, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Runes {
					if stage.RuneMap_Staged_Order[__instance__] == uint(id) {
						paragraph.Runes = append(paragraph.Runes, __instance__)
						break
					}
				}
			}
		}
	case "CollatedText":
		paragraph.CollatedText = value.GetValueString()
	case "Next":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			paragraph.Next = nil
			for __instance__ := range stage.Paragraphs {
				if stage.ParagraphMap_Staged_Order[__instance__] == uint(id) {
					paragraph.Next = __instance__
					break
				}
			}
		}
	case "Previous":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			paragraph.Previous = nil
			for __instance__ := range stage.Paragraphs {
				if stage.ParagraphMap_Staged_Order[__instance__] == uint(id) {
					paragraph.Previous = __instance__
					break
				}
			}
		}
	case "EnclosingBody":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			paragraph.EnclosingBody = nil
			for __instance__ := range stage.Bodys {
				if stage.BodyMap_Staged_Order[__instance__] == uint(id) {
					paragraph.EnclosingBody = __instance__
					break
				}
			}
		}
	case "EnclosingTableColumn":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			paragraph.EnclosingTableColumn = nil
			for __instance__ := range stage.TableColumns {
				if stage.TableColumnMap_Staged_Order[__instance__] == uint(id) {
					paragraph.EnclosingTableColumn = __instance__
					break
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (paragraphproperties *ParagraphProperties) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		paragraphproperties.Name = value.GetValueString()
	case "Content":
		paragraphproperties.Content = value.GetValueString()
	case "ParagraphStyle":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			paragraphproperties.ParagraphStyle = nil
			for __instance__ := range stage.ParagraphStyles {
				if stage.ParagraphStyleMap_Staged_Order[__instance__] == uint(id) {
					paragraphproperties.ParagraphStyle = __instance__
					break
				}
			}
		}
	case "Node":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			paragraphproperties.Node = nil
			for __instance__ := range stage.Nodes {
				if stage.NodeMap_Staged_Order[__instance__] == uint(id) {
					paragraphproperties.Node = __instance__
					break
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (paragraphstyle *ParagraphStyle) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		paragraphstyle.Name = value.GetValueString()
	case "Node":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			paragraphstyle.Node = nil
			for __instance__ := range stage.Nodes {
				if stage.NodeMap_Staged_Order[__instance__] == uint(id) {
					paragraphstyle.Node = __instance__
					break
				}
			}
		}
	case "Content":
		paragraphstyle.Content = value.GetValueString()
	case "ValAttr":
		paragraphstyle.ValAttr = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (rune *Rune) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		rune.Name = value.GetValueString()
	case "Content":
		rune.Content = value.GetValueString()
	case "Node":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			rune.Node = nil
			for __instance__ := range stage.Nodes {
				if stage.NodeMap_Staged_Order[__instance__] == uint(id) {
					rune.Node = __instance__
					break
				}
			}
		}
	case "Text":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			rune.Text = nil
			for __instance__ := range stage.Texts {
				if stage.TextMap_Staged_Order[__instance__] == uint(id) {
					rune.Text = __instance__
					break
				}
			}
		}
	case "RuneProperties":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			rune.RuneProperties = nil
			for __instance__ := range stage.RunePropertiess {
				if stage.RunePropertiesMap_Staged_Order[__instance__] == uint(id) {
					rune.RuneProperties = __instance__
					break
				}
			}
		}
	case "EnclosingParagraph":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			rune.EnclosingParagraph = nil
			for __instance__ := range stage.Paragraphs {
				if stage.ParagraphMap_Staged_Order[__instance__] == uint(id) {
					rune.EnclosingParagraph = __instance__
					break
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (runeproperties *RuneProperties) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		runeproperties.Name = value.GetValueString()
	case "Node":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			runeproperties.Node = nil
			for __instance__ := range stage.Nodes {
				if stage.NodeMap_Staged_Order[__instance__] == uint(id) {
					runeproperties.Node = __instance__
					break
				}
			}
		}
	case "IsBold":
		runeproperties.IsBold = value.GetValueBool()
	case "IsStrike":
		runeproperties.IsStrike = value.GetValueBool()
	case "IsItalic":
		runeproperties.IsItalic = value.GetValueBool()
	case "Content":
		runeproperties.Content = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (table *Table) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		table.Name = value.GetValueString()
	case "Node":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			table.Node = nil
			for __instance__ := range stage.Nodes {
				if stage.NodeMap_Staged_Order[__instance__] == uint(id) {
					table.Node = __instance__
					break
				}
			}
		}
	case "Content":
		table.Content = value.GetValueString()
	case "TableProperties":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			table.TableProperties = nil
			for __instance__ := range stage.TablePropertiess {
				if stage.TablePropertiesMap_Staged_Order[__instance__] == uint(id) {
					table.TableProperties = __instance__
					break
				}
			}
		}
	case "TableRows":
		table.TableRows = make([]*TableRow, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.TableRows {
					if stage.TableRowMap_Staged_Order[__instance__] == uint(id) {
						table.TableRows = append(table.TableRows, __instance__)
						break
					}
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (tablecolumn *TableColumn) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		tablecolumn.Name = value.GetValueString()
	case "Content":
		tablecolumn.Content = value.GetValueString()
	case "Node":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			tablecolumn.Node = nil
			for __instance__ := range stage.Nodes {
				if stage.NodeMap_Staged_Order[__instance__] == uint(id) {
					tablecolumn.Node = __instance__
					break
				}
			}
		}
	case "Paragraphs":
		tablecolumn.Paragraphs = make([]*Paragraph, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Paragraphs {
					if stage.ParagraphMap_Staged_Order[__instance__] == uint(id) {
						tablecolumn.Paragraphs = append(tablecolumn.Paragraphs, __instance__)
						break
					}
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (tableproperties *TableProperties) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		tableproperties.Name = value.GetValueString()
	case "Node":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			tableproperties.Node = nil
			for __instance__ := range stage.Nodes {
				if stage.NodeMap_Staged_Order[__instance__] == uint(id) {
					tableproperties.Node = __instance__
					break
				}
			}
		}
	case "Content":
		tableproperties.Content = value.GetValueString()
	case "TableStyle":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			tableproperties.TableStyle = nil
			for __instance__ := range stage.TableStyles {
				if stage.TableStyleMap_Staged_Order[__instance__] == uint(id) {
					tableproperties.TableStyle = __instance__
					break
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (tablerow *TableRow) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		tablerow.Name = value.GetValueString()
	case "Content":
		tablerow.Content = value.GetValueString()
	case "Node":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			tablerow.Node = nil
			for __instance__ := range stage.Nodes {
				if stage.NodeMap_Staged_Order[__instance__] == uint(id) {
					tablerow.Node = __instance__
					break
				}
			}
		}
	case "TableColumns":
		tablerow.TableColumns = make([]*TableColumn, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.TableColumns {
					if stage.TableColumnMap_Staged_Order[__instance__] == uint(id) {
						tablerow.TableColumns = append(tablerow.TableColumns, __instance__)
						break
					}
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (tablestyle *TableStyle) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		tablestyle.Name = value.GetValueString()
	case "Node":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			tablestyle.Node = nil
			for __instance__ := range stage.Nodes {
				if stage.NodeMap_Staged_Order[__instance__] == uint(id) {
					tablestyle.Node = __instance__
					break
				}
			}
		}
	case "Content":
		tablestyle.Content = value.GetValueString()
	case "Val":
		tablestyle.Val = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (text *Text) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		text.Name = value.GetValueString()
	case "Content":
		text.Content = value.GetValueString()
	case "Node":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			text.Node = nil
			for __instance__ := range stage.Nodes {
				if stage.NodeMap_Staged_Order[__instance__] == uint(id) {
					text.Node = __instance__
					break
				}
			}
		}
	case "PreserveWhiteSpace":
		text.PreserveWhiteSpace = value.GetValueBool()
	case "EnclosingRune":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			text.EnclosingRune = nil
			for __instance__ := range stage.Runes {
				if stage.RuneMap_Staged_Order[__instance__] == uint(id) {
					text.EnclosingRune = __instance__
					break
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func SetFieldStringValueFromPointer(instance GongstructIF, fieldName string, value GongFieldValue, stage *Stage) error {
	return instance.GongSetFieldValue(fieldName, value, stage)
}

// insertion point for generic get gongstruct name
func (body *Body) GongGetGongstructName() string {
	return "Body"
}

func (document *Document) GongGetGongstructName() string {
	return "Document"
}

func (docx *Docx) GongGetGongstructName() string {
	return "Docx"
}

func (file *File) GongGetGongstructName() string {
	return "File"
}

func (node *Node) GongGetGongstructName() string {
	return "Node"
}

func (paragraph *Paragraph) GongGetGongstructName() string {
	return "Paragraph"
}

func (paragraphproperties *ParagraphProperties) GongGetGongstructName() string {
	return "ParagraphProperties"
}

func (paragraphstyle *ParagraphStyle) GongGetGongstructName() string {
	return "ParagraphStyle"
}

func (rune *Rune) GongGetGongstructName() string {
	return "Rune"
}

func (runeproperties *RuneProperties) GongGetGongstructName() string {
	return "RuneProperties"
}

func (table *Table) GongGetGongstructName() string {
	return "Table"
}

func (tablecolumn *TableColumn) GongGetGongstructName() string {
	return "TableColumn"
}

func (tableproperties *TableProperties) GongGetGongstructName() string {
	return "TableProperties"
}

func (tablerow *TableRow) GongGetGongstructName() string {
	return "TableRow"
}

func (tablestyle *TableStyle) GongGetGongstructName() string {
	return "TableStyle"
}

func (text *Text) GongGetGongstructName() string {
	return "Text"
}

func GetGongstructNameFromPointer(instance GongstructIF) (res string) {
	res = instance.GongGetGongstructName()
	return
}

func (stage *Stage) ResetMapStrings() {

	// insertion point for generic get gongstruct name
	stage.Bodys_mapString = make(map[string]*Body)
	for body := range stage.Bodys {
		stage.Bodys_mapString[body.Name] = body
	}

	stage.Documents_mapString = make(map[string]*Document)
	for document := range stage.Documents {
		stage.Documents_mapString[document.Name] = document
	}

	stage.Docxs_mapString = make(map[string]*Docx)
	for docx := range stage.Docxs {
		stage.Docxs_mapString[docx.Name] = docx
	}

	stage.Files_mapString = make(map[string]*File)
	for file := range stage.Files {
		stage.Files_mapString[file.Name] = file
	}

	stage.Nodes_mapString = make(map[string]*Node)
	for node := range stage.Nodes {
		stage.Nodes_mapString[node.Name] = node
	}

	stage.Paragraphs_mapString = make(map[string]*Paragraph)
	for paragraph := range stage.Paragraphs {
		stage.Paragraphs_mapString[paragraph.Name] = paragraph
	}

	stage.ParagraphPropertiess_mapString = make(map[string]*ParagraphProperties)
	for paragraphproperties := range stage.ParagraphPropertiess {
		stage.ParagraphPropertiess_mapString[paragraphproperties.Name] = paragraphproperties
	}

	stage.ParagraphStyles_mapString = make(map[string]*ParagraphStyle)
	for paragraphstyle := range stage.ParagraphStyles {
		stage.ParagraphStyles_mapString[paragraphstyle.Name] = paragraphstyle
	}

	stage.Runes_mapString = make(map[string]*Rune)
	for rune := range stage.Runes {
		stage.Runes_mapString[rune.Name] = rune
	}

	stage.RunePropertiess_mapString = make(map[string]*RuneProperties)
	for runeproperties := range stage.RunePropertiess {
		stage.RunePropertiess_mapString[runeproperties.Name] = runeproperties
	}

	stage.Tables_mapString = make(map[string]*Table)
	for table := range stage.Tables {
		stage.Tables_mapString[table.Name] = table
	}

	stage.TableColumns_mapString = make(map[string]*TableColumn)
	for tablecolumn := range stage.TableColumns {
		stage.TableColumns_mapString[tablecolumn.Name] = tablecolumn
	}

	stage.TablePropertiess_mapString = make(map[string]*TableProperties)
	for tableproperties := range stage.TablePropertiess {
		stage.TablePropertiess_mapString[tableproperties.Name] = tableproperties
	}

	stage.TableRows_mapString = make(map[string]*TableRow)
	for tablerow := range stage.TableRows {
		stage.TableRows_mapString[tablerow.Name] = tablerow
	}

	stage.TableStyles_mapString = make(map[string]*TableStyle)
	for tablestyle := range stage.TableStyles {
		stage.TableStyles_mapString[tablestyle.Name] = tablestyle
	}

	stage.Texts_mapString = make(map[string]*Text)
	for text := range stage.Texts {
		stage.Texts_mapString[text.Name] = text
	}

}

// Last line of the template
