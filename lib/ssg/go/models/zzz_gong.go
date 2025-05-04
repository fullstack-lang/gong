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
	"time"

	ssg_go "github.com/fullstack-lang/gong/lib/ssg/go"
)

// can be used for
//     days := __Gong__Abs(int(int(inferedInstance.ComputedDuration.Hours()) / 24))
func __Gong__Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

var _ = __Gong__Abs

const ProbeTreeSidebarSuffix = "-sidebar"
const ProbeTableSuffix = "-table"
const ProbeFormSuffix = "-form"
const ProbeSplitSuffix = "-probe"

func (stage *Stage) GetProbeTreeSidebarStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeTreeSidebarSuffix
}

func (stage *Stage) GetProbeFormStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeFormSuffix
}

func (stage *Stage) GetProbeTableStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeTableSuffix
}

func (stage *Stage) GetProbeSplitStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeSplitSuffix
}

// errUnkownEnum is returns when a value cannot match enum values
var errUnkownEnum = errors.New("unkown enum")

// needed to avoid when fmt package is not needed by generated code
var __dummy__fmt_variable fmt.Scanner

// idem for math package when not need by generated code
var __dummy_math_variable = math.E

// swagger:ignore
type __void any

// needed for creating set of instances in the stage
var __member __void

// GongStructInterface is the interface met by GongStructs
// It allows runtime reflexion of instances (without the hassle of the "reflect" package)
type GongStructInterface interface {
	GetName() (res string)
	// GetID() (res int)
	// GetFields() (res []string)
	// GetFieldStringValue(fieldName string) (res string)
}

// Stage enables storage of staged instances
// swagger:ignore
type Stage struct {
	name string

	// insertion point for definition of arrays registering instances
	Chapters           map[*Chapter]any
	Chapters_mapString map[string]*Chapter

	// insertion point for slice of pointers maps
	Chapter_Pages_reverseMap map[*Page]*Chapter

	OnAfterChapterCreateCallback OnAfterCreateInterface[Chapter]
	OnAfterChapterUpdateCallback OnAfterUpdateInterface[Chapter]
	OnAfterChapterDeleteCallback OnAfterDeleteInterface[Chapter]
	OnAfterChapterReadCallback   OnAfterReadInterface[Chapter]

	Contents           map[*Content]any
	Contents_mapString map[string]*Content

	// insertion point for slice of pointers maps
	Content_Chapters_reverseMap map[*Chapter]*Content

	OnAfterContentCreateCallback OnAfterCreateInterface[Content]
	OnAfterContentUpdateCallback OnAfterUpdateInterface[Content]
	OnAfterContentDeleteCallback OnAfterDeleteInterface[Content]
	OnAfterContentReadCallback   OnAfterReadInterface[Content]

	Pages           map[*Page]any
	Pages_mapString map[string]*Page

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
}

// GetNamedStructs implements models.ProbebStage.
func (stage *Stage) GetNamedStructsNames() (res []string) {

	for _, namedStruct := range stage.NamedStructs {
		res = append(res, namedStruct.name)
	}

	return
}

func GetNamedStructInstances[T PointerToGongstruct](set map[T]any, order map[T]uint) (res []string) {

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
		Chapters:           make(map[*Chapter]any),
		Chapters_mapString: make(map[string]*Chapter),

		Contents:           make(map[*Content]any),
		Contents_mapString: make(map[string]*Content),

		Pages:           make(map[*Page]any),
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

		NamedStructs: []*NamedStruct{ // insertion point for order map initialisations
			{name: "Chapter"},
			{name: "Content"},
			{name: "Page"},
		}, // end of insertion point
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

	if stage.BackRepo != nil {
		stage.BackRepo.Commit(stage)
	}

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
	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["Chapter"] = len(stage.Chapters)
	stage.Map_GongStructName_InstancesNb["Content"] = len(stage.Contents)
	stage.Map_GongStructName_InstancesNb["Page"] = len(stage.Pages)

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
		stage.Chapters[chapter] = __member
		stage.ChapterMap_Staged_Order[chapter] = stage.ChapterOrder
		stage.ChapterOrder++
	}
	stage.Chapters_mapString[chapter.Name] = chapter

	return chapter
}

// Unstage removes chapter off the model stage
func (chapter *Chapter) Unstage(stage *Stage) *Chapter {
	delete(stage.Chapters, chapter)
	delete(stage.Chapters_mapString, chapter.Name)
	return chapter
}

// UnstageVoid removes chapter off the model stage
func (chapter *Chapter) UnstageVoid(stage *Stage) {
	delete(stage.Chapters, chapter)
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

// Stage puts content to the model stage
func (content *Content) Stage(stage *Stage) *Content {

	if _, ok := stage.Contents[content]; !ok {
		stage.Contents[content] = __member
		stage.ContentMap_Staged_Order[content] = stage.ContentOrder
		stage.ContentOrder++
	}
	stage.Contents_mapString[content.Name] = content

	return content
}

// Unstage removes content off the model stage
func (content *Content) Unstage(stage *Stage) *Content {
	delete(stage.Contents, content)
	delete(stage.Contents_mapString, content.Name)
	return content
}

// UnstageVoid removes content off the model stage
func (content *Content) UnstageVoid(stage *Stage) {
	delete(stage.Contents, content)
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

// Stage puts page to the model stage
func (page *Page) Stage(stage *Stage) *Page {

	if _, ok := stage.Pages[page]; !ok {
		stage.Pages[page] = __member
		stage.PageMap_Staged_Order[page] = stage.PageOrder
		stage.PageOrder++
	}
	stage.Pages_mapString[page.Name] = page

	return page
}

// Unstage removes page off the model stage
func (page *Page) Unstage(stage *Stage) *Page {
	delete(stage.Pages, page)
	delete(stage.Pages_mapString, page.Name)
	return page
}

// UnstageVoid removes page off the model stage
func (page *Page) UnstageVoid(stage *Stage) {
	delete(stage.Pages, page)
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
	stage.Chapters = make(map[*Chapter]any)
	stage.Chapters_mapString = make(map[string]*Chapter)
	stage.ChapterMap_Staged_Order = make(map[*Chapter]uint)
	stage.ChapterOrder = 0

	stage.Contents = make(map[*Content]any)
	stage.Contents_mapString = make(map[string]*Content)
	stage.ContentMap_Staged_Order = make(map[*Content]uint)
	stage.ContentOrder = 0

	stage.Pages = make(map[*Page]any)
	stage.Pages_mapString = make(map[string]*Page)
	stage.PageMap_Staged_Order = make(map[*Page]uint)
	stage.PageOrder = 0

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
type PointerToGongstruct interface {
	GetName() string
	CommitVoid(*Stage)
	UnstageVoid(stage *Stage)
	comparable
}

func CompareGongstructByName[T PointerToGongstruct](a, b T) int {
	return cmp.Compare(a.GetName(), b.GetName())
}

func SortGongstructSetByName[T PointerToGongstruct](set map[T]any) (sortedSlice []T) {

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

// GongGetMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GongGetMap[Type GongstructMapString](stage *Stage) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case map[string]*Chapter:
		return any(&stage.Chapters_mapString).(*Type)
	case map[string]*Content:
		return any(&stage.Contents_mapString).(*Type)
	case map[string]*Page:
		return any(&stage.Pages_mapString).(*Type)
	default:
		return nil
	}
}

// GetGongstructInstancesSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSet[Type Gongstruct](stage *Stage) *map[*Type]any {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case Chapter:
		return any(&stage.Chapters).(*map[*Type]any)
	case Content:
		return any(&stage.Contents).(*map[*Type]any)
	case Page:
		return any(&stage.Pages).(*map[*Type]any)
	default:
		return nil
	}
}

// GetGongstructInstancesSetFromPointerType returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSetFromPointerType[Type PointerToGongstruct](stage *Stage) *map[Type]any {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case *Chapter:
		return any(&stage.Chapters).(*map[Type]any)
	case *Content:
		return any(&stage.Contents).(*map[Type]any)
	case *Page:
		return any(&stage.Pages).(*map[Type]any)
	default:
		return nil
	}
}

// GetGongstructInstancesMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
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
func GetSliceOfPointersReverseMap[Start, End Gongstruct](fieldname string, stage *Stage) map[*End]*Start {

	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of Chapter
	case Chapter:
		switch fieldname {
		// insertion point for per direct association field
		case "Pages":
			res := make(map[*Page]*Chapter)
			for chapter := range stage.Chapters {
				for _, page_ := range chapter.Pages {
					res[page_] = chapter
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of Content
	case Content:
		switch fieldname {
		// insertion point for per direct association field
		case "Chapters":
			res := make(map[*Chapter]*Content)
			for content := range stage.Contents {
				for _, chapter_ := range content.Chapters {
					res[chapter_] = content
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of Page
	case Page:
		switch fieldname {
		// insertion point for per direct association field
		}
	}
	return nil
}

// GetGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetGongstructName[Type Gongstruct]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case Chapter:
		res = "Chapter"
	case Content:
		res = "Content"
	case Page:
		res = "Page"
	}
	return res
}

// GetPointerToGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetPointerToGongstructName[Type PointerToGongstruct]() (res string) {

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

// GetFields return the array of the fields
func GetFields[Type Gongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case Chapter:
		res = []string{"Name", "MardownContent", "Pages"}
	case Content:
		res = []string{"Name", "MardownContent", "ContentPath", "OutputPath", "LayoutPath", "StaticPath", "Target", "Chapters"}
	case Page:
		res = []string{"Name", "MardownContent"}
	}
	return
}

type ReverseField struct {
	GongstructName string
	Fieldname      string
}

func GetReverseFields[Type Gongstruct]() (res []ReverseField) {

	res = make([]ReverseField, 0)

	var ret Type

	switch any(ret).(type) {

	// insertion point for generic get gongstruct name
	case Chapter:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Content"
		rf.Fieldname = "Chapters"
		res = append(res, rf)
	case Content:
		var rf ReverseField
		_ = rf
	case Page:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Chapter"
		rf.Fieldname = "Pages"
		res = append(res, rf)
	}
	return
}

// GetFieldsFromPointer return the array of the fields
func GetFieldsFromPointer[Type PointerToGongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *Chapter:
		res = []string{"Name", "MardownContent", "Pages"}
	case *Content:
		res = []string{"Name", "MardownContent", "ContentPath", "OutputPath", "LayoutPath", "StaticPath", "Target", "Chapters"}
	case *Page:
		res = []string{"Name", "MardownContent"}
	}
	return
}

type GongFieldValueType string

const (
	GongFieldValueTypeInt    GongFieldValueType = "GongFieldValueTypeInt"
	GongFieldValueTypeFloat  GongFieldValueType = "GongFieldValueTypeFloat"
	GongFieldValueTypeBool   GongFieldValueType = "GongFieldValueTypeBool"
	GongFieldValueTypeOthers GongFieldValueType = "GongFieldValueTypeOthers"
)

type GongFieldValue struct {
	valueString string
	GongFieldValueType
	valueInt   int
	valueFloat float64
	valueBool  bool
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

func GetFieldStringValueFromPointer(instance any, fieldName string) (res GongFieldValue) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case *Chapter:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "MardownContent":
			res.valueString = inferedInstance.MardownContent
		case "Pages":
			for idx, __instance__ := range inferedInstance.Pages {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case *Content:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "MardownContent":
			res.valueString = inferedInstance.MardownContent
		case "ContentPath":
			res.valueString = inferedInstance.ContentPath
		case "OutputPath":
			res.valueString = inferedInstance.OutputPath
		case "LayoutPath":
			res.valueString = inferedInstance.LayoutPath
		case "StaticPath":
			res.valueString = inferedInstance.StaticPath
		case "Target":
			enum := inferedInstance.Target
			res.valueString = enum.ToCodeString()
		case "Chapters":
			for idx, __instance__ := range inferedInstance.Chapters {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case *Page:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "MardownContent":
			res.valueString = inferedInstance.MardownContent
		}
	default:
		_ = inferedInstance
	}
	return
}

func GetFieldStringValue(instance any, fieldName string) (res GongFieldValue) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case Chapter:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "MardownContent":
			res.valueString = inferedInstance.MardownContent
		case "Pages":
			for idx, __instance__ := range inferedInstance.Pages {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case Content:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "MardownContent":
			res.valueString = inferedInstance.MardownContent
		case "ContentPath":
			res.valueString = inferedInstance.ContentPath
		case "OutputPath":
			res.valueString = inferedInstance.OutputPath
		case "LayoutPath":
			res.valueString = inferedInstance.LayoutPath
		case "StaticPath":
			res.valueString = inferedInstance.StaticPath
		case "Target":
			enum := inferedInstance.Target
			res.valueString = enum.ToCodeString()
		case "Chapters":
			for idx, __instance__ := range inferedInstance.Chapters {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case Page:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "MardownContent":
			res.valueString = inferedInstance.MardownContent
		}
	default:
		_ = inferedInstance
	}
	return
}

// Last line of the template
