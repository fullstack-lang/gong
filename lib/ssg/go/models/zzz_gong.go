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

	ssg_go "github.com/fullstack-lang/gong/lib/ssg/go"
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
	Chapters                map[*Chapter]struct{}
	Chapters_reference      map[*Chapter]*Chapter
	Chapters_referenceOrder map[*Chapter]uint // diff Unstage needs the reference order
	Chapters_mapString      map[string]*Chapter

	// insertion point for slice of pointers maps
	Chapter_Pages_reverseMap map[*Page]*Chapter

	OnAfterChapterCreateCallback OnAfterCreateInterface[Chapter]
	OnAfterChapterUpdateCallback OnAfterUpdateInterface[Chapter]
	OnAfterChapterDeleteCallback OnAfterDeleteInterface[Chapter]
	OnAfterChapterReadCallback   OnAfterReadInterface[Chapter]

	Contents                map[*Content]struct{}
	Contents_reference      map[*Content]*Content
	Contents_referenceOrder map[*Content]uint // diff Unstage needs the reference order
	Contents_mapString      map[string]*Content

	// insertion point for slice of pointers maps
	Content_Chapters_reverseMap map[*Chapter]*Content

	OnAfterContentCreateCallback OnAfterCreateInterface[Content]
	OnAfterContentUpdateCallback OnAfterUpdateInterface[Content]
	OnAfterContentDeleteCallback OnAfterDeleteInterface[Content]
	OnAfterContentReadCallback   OnAfterReadInterface[Content]

	Pages                map[*Page]struct{}
	Pages_reference      map[*Page]*Page
	Pages_referenceOrder map[*Page]uint // diff Unstage needs the reference order
	Pages_mapString      map[string]*Page

	// insertion point for slice of pointers maps
	OnAfterPageCreateCallback OnAfterCreateInterface[Page]
	OnAfterPageUpdateCallback OnAfterUpdateInterface[Page]
	OnAfterPageDeleteCallback OnAfterDeleteInterface[Page]
	OnAfterPageReadCallback   OnAfterReadInterface[Page]

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
	ChapterOrder            uint
	ChapterMap_Staged_Order map[*Chapter]uint

	ContentOrder            uint
	ContentMap_Staged_Order map[*Content]uint

	PageOrder            uint
	PageMap_Staged_Order map[*Page]uint

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
	err := GongParseAstString(stage, commitToApply, true)
	if err != nil {
		log.Println("error during ApplyBackwardCommit: ", err)
		return err
	}

	stage.ComputeReferenceAndOrders()

	stage.commitsBehind++

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
	err := GongParseAstString(stage, commitToApply, true)
	if err != nil {
		log.Println("error during ApplyForwardCommit: ", err)
		return err
	}
	stage.ComputeReferenceAndOrders()

	stage.commitsBehind--
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
}

// recomputeOrders recomputes the next order for each struct
// this is necessary because the order might have been incremented
// during the commits that have been discarded
// insertion point for max order recomputation
func (stage *Stage) recomputeOrders() {
	// insertion point for max order recomputation 
	var maxChapterOrder uint
	var foundChapter bool
	for _, order := range stage.ChapterMap_Staged_Order {
		if !foundChapter || order > maxChapterOrder {
			maxChapterOrder = order
			foundChapter = true
		}
	}
	if foundChapter {
		stage.ChapterOrder = maxChapterOrder + 1
	} else {
		stage.ChapterOrder = 0
	}

	var maxContentOrder uint
	var foundContent bool
	for _, order := range stage.ContentMap_Staged_Order {
		if !foundContent || order > maxContentOrder {
			maxContentOrder = order
			foundContent = true
		}
	}
	if foundContent {
		stage.ContentOrder = maxContentOrder + 1
	} else {
		stage.ContentOrder = 0
	}

	var maxPageOrder uint
	var foundPage bool
	for _, order := range stage.PageMap_Staged_Order {
		if !foundPage || order > maxPageOrder {
			maxPageOrder = order
			foundPage = true
		}
	}
	if foundPage {
		stage.PageOrder = maxPageOrder + 1
	} else {
		stage.PageOrder = 0
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
	case *Chapter:
		tmp := GetStructInstancesByOrder(stage.Chapters, stage.ChapterMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Chapter implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Content:
		tmp := GetStructInstancesByOrder(stage.Contents, stage.ContentMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Content implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Page:
		tmp := GetStructInstancesByOrder(stage.Pages, stage.PageMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Page implements.
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
	case "Chapter":
		res = GetNamedStructInstances(stage.Chapters, stage.ChapterMap_Staged_Order)
	case "Content":
		res = GetNamedStructInstances(stage.Contents, stage.ContentMap_Staged_Order)
	case "Page":
		res = GetNamedStructInstances(stage.Pages, stage.PageMap_Staged_Order)
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
	return "github.com/fullstack-lang/gong/lib/ssg/go/models"
}

func (stage *Stage) GetMap_GongStructName_InstancesNb() map[string]int {
	return stage.Map_GongStructName_InstancesNb
}

func (stage *Stage) GetModelsEmbededDir() embed.FS {
	return ssg_go.GoModelsDir
}

func (stage *Stage) GetDigramsEmbededDir() embed.FS {
	return ssg_go.GoDiagramsDir
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
	CommitChapter(chapter *Chapter)
	CheckoutChapter(chapter *Chapter)
	CommitContent(content *Content)
	CheckoutContent(content *Content)
	CommitPage(page *Page)
	CheckoutPage(page *Page)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(name string) (stage *Stage) {

	stage = &Stage{ // insertion point for array initiatialisation
		Chapters:           make(map[*Chapter]struct{}),
		Chapters_mapString: make(map[string]*Chapter),

		Contents:           make(map[*Content]struct{}),
		Contents_mapString: make(map[string]*Content),

		Pages:           make(map[*Page]struct{}),
		Pages_mapString: make(map[string]*Page),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		name: name,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		ChapterMap_Staged_Order: make(map[*Chapter]uint),

		ContentMap_Staged_Order: make(map[*Content]uint),

		PageMap_Staged_Order: make(map[*Page]uint),

		// end of insertion point
		GongUnmarshallers: map[string]ModelUnmarshaller{ // insertion point for unmarshallers
			"Chapter": &ChapterUnmarshaller{},

			"Content": &ContentUnmarshaller{},

			"Page": &PageUnmarshaller{},

			// end of insertion point
		},

		NamedStructs: []*NamedStruct{ // insertion point for order map initialisations
			{name: "Chapter"},
			{name: "Content"},
			{name: "Page"},
		}, // end of insertion point

		navigationMode: GongNavigationModeNormal,
	}

	return
}

func GetOrder[Type Gongstruct](stage *Stage, instance *Type) uint {

	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *Chapter:
		return stage.ChapterMap_Staged_Order[instance]
	case *Content:
		return stage.ContentMap_Staged_Order[instance]
	case *Page:
		return stage.PageMap_Staged_Order[instance]
	default:
		return 0 // should not happen
	}
}

func GetOrderPointerGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) uint {

	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *Chapter:
		return stage.ChapterMap_Staged_Order[instance]
	case *Content:
		return stage.ContentMap_Staged_Order[instance]
	case *Page:
		return stage.PageMap_Staged_Order[instance]
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
	stage.Map_GongStructName_InstancesNb["Chapter"] = len(stage.Chapters)
	stage.Map_GongStructName_InstancesNb["Content"] = len(stage.Contents)
	stage.Map_GongStructName_InstancesNb["Page"] = len(stage.Pages)
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
// Stage puts chapter to the model stage
func (chapter *Chapter) Stage(stage *Stage) *Chapter {

	if _, ok := stage.Chapters[chapter]; !ok {
		stage.Chapters[chapter] = struct{}{}
		stage.ChapterMap_Staged_Order[chapter] = stage.ChapterOrder
		stage.ChapterOrder++
	}
	stage.Chapters_mapString[chapter.Name] = chapter

	return chapter
}

// StagePreserveOrder puts chapter to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ChapterOrder
// - update stage.ChapterOrder accordingly
func (chapter *Chapter) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.Chapters[chapter]; !ok {
		stage.Chapters[chapter] = struct{}{}

		if order > stage.ChapterOrder {
			stage.ChapterOrder = order
		}
		stage.ChapterMap_Staged_Order[chapter] = order
		stage.ChapterOrder++
	}
	stage.Chapters_mapString[chapter.Name] = chapter
}

// Unstage removes chapter off the model stage
func (chapter *Chapter) Unstage(stage *Stage) *Chapter {
	delete(stage.Chapters, chapter)
	delete(stage.ChapterMap_Staged_Order, chapter)
	delete(stage.Chapters_mapString, chapter.Name)

	return chapter
}

// UnstageVoid removes chapter off the model stage
func (chapter *Chapter) UnstageVoid(stage *Stage) {
	delete(stage.Chapters, chapter)
	delete(stage.ChapterMap_Staged_Order, chapter)
	delete(stage.Chapters_mapString, chapter.Name)
}

// commit chapter to the back repo (if it is already staged)
func (chapter *Chapter) Commit(stage *Stage) *Chapter {
	if _, ok := stage.Chapters[chapter]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitChapter(chapter)
		}
	}
	return chapter
}

func (chapter *Chapter) CommitVoid(stage *Stage) {
	chapter.Commit(stage)
}

func (chapter *Chapter) StageVoid(stage *Stage) {
	chapter.Stage(stage)
}

// Checkout chapter to the back repo (if it is already staged)
func (chapter *Chapter) Checkout(stage *Stage) *Chapter {
	if _, ok := stage.Chapters[chapter]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutChapter(chapter)
		}
	}
	return chapter
}

// for satisfaction of GongStruct interface
func (chapter *Chapter) GetName() (res string) {
	return chapter.Name
}

// for satisfaction of GongStruct interface
func (chapter *Chapter) SetName(name string) {
	chapter.Name = name
}

// Stage puts content to the model stage
func (content *Content) Stage(stage *Stage) *Content {

	if _, ok := stage.Contents[content]; !ok {
		stage.Contents[content] = struct{}{}
		stage.ContentMap_Staged_Order[content] = stage.ContentOrder
		stage.ContentOrder++
	}
	stage.Contents_mapString[content.Name] = content

	return content
}

// StagePreserveOrder puts content to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ContentOrder
// - update stage.ContentOrder accordingly
func (content *Content) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.Contents[content]; !ok {
		stage.Contents[content] = struct{}{}

		if order > stage.ContentOrder {
			stage.ContentOrder = order
		}
		stage.ContentMap_Staged_Order[content] = order
		stage.ContentOrder++
	}
	stage.Contents_mapString[content.Name] = content
}

// Unstage removes content off the model stage
func (content *Content) Unstage(stage *Stage) *Content {
	delete(stage.Contents, content)
	delete(stage.ContentMap_Staged_Order, content)
	delete(stage.Contents_mapString, content.Name)

	return content
}

// UnstageVoid removes content off the model stage
func (content *Content) UnstageVoid(stage *Stage) {
	delete(stage.Contents, content)
	delete(stage.ContentMap_Staged_Order, content)
	delete(stage.Contents_mapString, content.Name)
}

// commit content to the back repo (if it is already staged)
func (content *Content) Commit(stage *Stage) *Content {
	if _, ok := stage.Contents[content]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitContent(content)
		}
	}
	return content
}

func (content *Content) CommitVoid(stage *Stage) {
	content.Commit(stage)
}

func (content *Content) StageVoid(stage *Stage) {
	content.Stage(stage)
}

// Checkout content to the back repo (if it is already staged)
func (content *Content) Checkout(stage *Stage) *Content {
	if _, ok := stage.Contents[content]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutContent(content)
		}
	}
	return content
}

// for satisfaction of GongStruct interface
func (content *Content) GetName() (res string) {
	return content.Name
}

// for satisfaction of GongStruct interface
func (content *Content) SetName(name string) {
	content.Name = name
}

// Stage puts page to the model stage
func (page *Page) Stage(stage *Stage) *Page {

	if _, ok := stage.Pages[page]; !ok {
		stage.Pages[page] = struct{}{}
		stage.PageMap_Staged_Order[page] = stage.PageOrder
		stage.PageOrder++
	}
	stage.Pages_mapString[page.Name] = page

	return page
}

// StagePreserveOrder puts page to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.PageOrder
// - update stage.PageOrder accordingly
func (page *Page) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.Pages[page]; !ok {
		stage.Pages[page] = struct{}{}

		if order > stage.PageOrder {
			stage.PageOrder = order
		}
		stage.PageMap_Staged_Order[page] = order
		stage.PageOrder++
	}
	stage.Pages_mapString[page.Name] = page
}

// Unstage removes page off the model stage
func (page *Page) Unstage(stage *Stage) *Page {
	delete(stage.Pages, page)
	delete(stage.PageMap_Staged_Order, page)
	delete(stage.Pages_mapString, page.Name)

	return page
}

// UnstageVoid removes page off the model stage
func (page *Page) UnstageVoid(stage *Stage) {
	delete(stage.Pages, page)
	delete(stage.PageMap_Staged_Order, page)
	delete(stage.Pages_mapString, page.Name)
}

// commit page to the back repo (if it is already staged)
func (page *Page) Commit(stage *Stage) *Page {
	if _, ok := stage.Pages[page]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitPage(page)
		}
	}
	return page
}

func (page *Page) CommitVoid(stage *Stage) {
	page.Commit(stage)
}

func (page *Page) StageVoid(stage *Stage) {
	page.Stage(stage)
}

// Checkout page to the back repo (if it is already staged)
func (page *Page) Checkout(stage *Stage) *Page {
	if _, ok := stage.Pages[page]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutPage(page)
		}
	}
	return page
}

// for satisfaction of GongStruct interface
func (page *Page) GetName() (res string) {
	return page.Name
}

// for satisfaction of GongStruct interface
func (page *Page) SetName(name string) {
	page.Name = name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMChapter(Chapter *Chapter)
	CreateORMContent(Content *Content)
	CreateORMPage(Page *Page)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMChapter(Chapter *Chapter)
	DeleteORMContent(Content *Content)
	DeleteORMPage(Page *Page)
}

func (stage *Stage) Reset() { // insertion point for array reset
	stage.Chapters = make(map[*Chapter]struct{})
	stage.Chapters_mapString = make(map[string]*Chapter)
	stage.ChapterMap_Staged_Order = make(map[*Chapter]uint)
	stage.ChapterOrder = 0

	stage.Contents = make(map[*Content]struct{})
	stage.Contents_mapString = make(map[string]*Content)
	stage.ContentMap_Staged_Order = make(map[*Content]uint)
	stage.ContentOrder = 0

	stage.Pages = make(map[*Page]struct{})
	stage.Pages_mapString = make(map[string]*Page)
	stage.PageMap_Staged_Order = make(map[*Page]uint)
	stage.PageOrder = 0

	if stage.GetProbeIF() != nil {
		stage.GetProbeIF().ResetNotifications()
	}
	if stage.IsInDeltaMode() {
		stage.ComputeReferenceAndOrders()
	}
}

func (stage *Stage) Nil() { // insertion point for array nil
	stage.Chapters = nil
	stage.Chapters_mapString = nil

	stage.Contents = nil
	stage.Contents_mapString = nil

	stage.Pages = nil
	stage.Pages_mapString = nil

}

func (stage *Stage) Unstage() { // insertion point for array nil
	for chapter := range stage.Chapters {
		chapter.Unstage(stage)
	}

	for content := range stage.Contents {
		content.Unstage(stage)
	}

	for page := range stage.Pages {
		page.Unstage(stage)
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
	case map[*Chapter]any:
		return any(&stage.Chapters).(*Type)
	case map[*Content]any:
		return any(&stage.Contents).(*Type)
	case map[*Page]any:
		return any(&stage.Pages).(*Type)
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
	case *Chapter:
		return any(stage.Chapters_mapString).(map[string]Type)
	case *Content:
		return any(stage.Contents_mapString).(map[string]Type)
	case *Page:
		return any(stage.Pages_mapString).(map[string]Type)
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
	case Chapter:
		return any(&stage.Chapters).(*map[*Type]struct{})
	case Content:
		return any(&stage.Contents).(*map[*Type]struct{})
	case Page:
		return any(&stage.Pages).(*map[*Type]struct{})
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
	case *Chapter:
		return any(&stage.Chapters).(*map[Type]struct{})
	case *Content:
		return any(&stage.Contents).(*map[Type]struct{})
	case *Page:
		return any(&stage.Pages).(*map[Type]struct{})
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
	case Chapter:
		return any(&stage.Chapters_mapString).(*map[string]*Type)
	case Content:
		return any(&stage.Contents_mapString).(*map[string]*Type)
	case Page:
		return any(&stage.Pages_mapString).(*map[string]*Type)
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
	case Chapter:
		return any(&Chapter{
			// Initialisation of associations
			// field is initialized with an instance of Page with the name of the field
			Pages: []*Page{{Name: "Pages"}},
		}).(*Type)
	case Content:
		return any(&Content{
			// Initialisation of associations
			// field is initialized with an instance of Chapter with the name of the field
			Chapters: []*Chapter{{Name: "Chapters"}},
		}).(*Type)
	case Page:
		return any(&Page{
			// Initialisation of associations
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
	// reverse maps of direct associations of Chapter
	case Chapter:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Content
	case Content:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Page
	case Page:
		switch fieldname {
		// insertion point for per direct association field
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
	// reverse maps of direct associations of Chapter
	case Chapter:
		switch fieldname {
		// insertion point for per direct association field
		case "Pages":
			res := make(map[*Page][]*Chapter)
			for chapter := range stage.Chapters {
				for _, page_ := range chapter.Pages {
					res[page_] = append(res[page_], chapter)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Content
	case Content:
		switch fieldname {
		// insertion point for per direct association field
		case "Chapters":
			res := make(map[*Chapter][]*Content)
			for content := range stage.Contents {
				for _, chapter_ := range content.Chapters {
					res[chapter_] = append(res[chapter_], content)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Page
	case Page:
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
	case *Chapter:
		res = "Chapter"
	case *Content:
		res = "Content"
	case *Page:
		res = "Page"
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
	case *Chapter:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Content"
		rf.Fieldname = "Chapters"
		res = append(res, rf)
	case *Content:
		var rf ReverseField
		_ = rf
	case *Page:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Chapter"
		rf.Fieldname = "Pages"
		res = append(res, rf)
	}
	return
}

// insertion point for get fields header method
func (chapter *Chapter) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "MardownContent",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "Pages",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Page",
		},
	}
	return
}

func (content *Content) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "MardownContent",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "ContentPath",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "OutputPath",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "LayoutPath",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StaticPath",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "IsBespokeLogoFileName",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "BespokeLogoFileName",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "IsBespokePageTileLogoFileName",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "BespokePageTileLogoFileName",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Target",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "Chapters",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Chapter",
		},
		{
			Name:               "VersionInfo",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (page *Page) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "MardownContent",
			GongFieldValueType: GongFieldValueTypeBasicKind,
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
func (chapter *Chapter) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = chapter.Name
	case "MardownContent":
		res.valueString = chapter.MardownContent
	case "Pages":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range chapter.Pages {
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
func (content *Content) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = content.Name
	case "MardownContent":
		res.valueString = content.MardownContent
	case "ContentPath":
		res.valueString = content.ContentPath
	case "OutputPath":
		res.valueString = content.OutputPath
	case "LayoutPath":
		res.valueString = content.LayoutPath
	case "StaticPath":
		res.valueString = content.StaticPath
	case "IsBespokeLogoFileName":
		res.valueString = fmt.Sprintf("%t", content.IsBespokeLogoFileName)
		res.valueBool = content.IsBespokeLogoFileName
		res.GongFieldValueType = GongFieldValueTypeBool
	case "BespokeLogoFileName":
		res.valueString = content.BespokeLogoFileName
	case "IsBespokePageTileLogoFileName":
		res.valueString = fmt.Sprintf("%t", content.IsBespokePageTileLogoFileName)
		res.valueBool = content.IsBespokePageTileLogoFileName
		res.GongFieldValueType = GongFieldValueTypeBool
	case "BespokePageTileLogoFileName":
		res.valueString = content.BespokePageTileLogoFileName
	case "Target":
		enum := content.Target
		res.valueString = enum.ToCodeString()
	case "Chapters":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range content.Chapters {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "VersionInfo":
		res.valueString = content.VersionInfo
	}
	return
}
func (page *Page) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = page.Name
	case "MardownContent":
		res.valueString = page.MardownContent
	}
	return
}
func GetFieldStringValueFromPointer(instance GongstructIF, fieldName string, stage *Stage) (res GongFieldValue) {

	res = instance.GongGetFieldValue(fieldName, stage)
	return
}

// insertion point for generic set gongstruct field value
func (chapter *Chapter) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		chapter.Name = value.GetValueString()
	case "MardownContent":
		chapter.MardownContent = value.GetValueString()
	case "Pages":
		chapter.Pages = make([]*Page, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Pages {
					if stage.PageMap_Staged_Order[__instance__] == uint(id) {
						chapter.Pages = append(chapter.Pages, __instance__)
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

func (content *Content) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		content.Name = value.GetValueString()
	case "MardownContent":
		content.MardownContent = value.GetValueString()
	case "ContentPath":
		content.ContentPath = value.GetValueString()
	case "OutputPath":
		content.OutputPath = value.GetValueString()
	case "LayoutPath":
		content.LayoutPath = value.GetValueString()
	case "StaticPath":
		content.StaticPath = value.GetValueString()
	case "IsBespokeLogoFileName":
		content.IsBespokeLogoFileName = value.GetValueBool()
	case "BespokeLogoFileName":
		content.BespokeLogoFileName = value.GetValueString()
	case "IsBespokePageTileLogoFileName":
		content.IsBespokePageTileLogoFileName = value.GetValueBool()
	case "BespokePageTileLogoFileName":
		content.BespokePageTileLogoFileName = value.GetValueString()
	case "Target":
		content.Target.FromCodeString(value.GetValueString())
	case "Chapters":
		content.Chapters = make([]*Chapter, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Chapters {
					if stage.ChapterMap_Staged_Order[__instance__] == uint(id) {
						content.Chapters = append(content.Chapters, __instance__)
						break
					}
				}
			}
		}
	case "VersionInfo":
		content.VersionInfo = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (page *Page) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		page.Name = value.GetValueString()
	case "MardownContent":
		page.MardownContent = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func SetFieldStringValueFromPointer(instance GongstructIF, fieldName string, value GongFieldValue, stage *Stage) error {
	return instance.GongSetFieldValue(fieldName, value, stage)
}

// insertion point for generic get gongstruct name
func (chapter *Chapter) GongGetGongstructName() string {
	return "Chapter"
}

func (content *Content) GongGetGongstructName() string {
	return "Content"
}

func (page *Page) GongGetGongstructName() string {
	return "Page"
}

func GetGongstructNameFromPointer(instance GongstructIF) (res string) {
	res = instance.GongGetGongstructName()
	return
}

func (stage *Stage) ResetMapStrings() {

	// insertion point for generic get gongstruct name
	stage.Chapters_mapString = make(map[string]*Chapter)
	for chapter := range stage.Chapters {
		stage.Chapters_mapString[chapter.Name] = chapter
	}

	stage.Contents_mapString = make(map[string]*Content)
	for content := range stage.Contents {
		stage.Contents_mapString[content.Name] = content
	}

	stage.Pages_mapString = make(map[string]*Page)
	for page := range stage.Pages {
		stage.Pages_mapString[page.Name] = page
	}

}

// Last line of the template
