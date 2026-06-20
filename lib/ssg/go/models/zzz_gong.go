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
	"sync"
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

var (
	_ = __Gong__Abs
	_ = strings.Clone("")
)

const (
	ProbeTreeSidebarSuffix           = ":sidebar of the probe"
	ProbeNavigationTreeSidebarSuffix = ":sidebar of the probe, navigation"
	ProbeTableSuffix                 = ":table of the probe"
	ProbeNotificationTableSuffix     = ":notification table of the probe"
	ProbeFormSuffix                  = ":form of the probe"
	ProbeSplitSuffix                 = ":probe of the probe"
	ProbeLoadSuffix                  = ":load of the probe"
)

type GongMarshallingMode string

const (
	// the whole stage is generated at each marshall. This is the default
	GongMarshallingNormal GongMarshallingMode = "GongMarshallingNormal"

	// only the last commit is append to the marshall file
	GongMarshallingAppendCommit GongMarshallingMode = "GongMarshallingAppendCommit"
)

func (stage *Stage) GetProbeTreeSidebarStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeTreeSidebarSuffix
}

func (stage *Stage) GetProbeNavigationTreeSidebarStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeNavigationTreeSidebarSuffix
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

func (stage *Stage) GetProbeLoadStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeLoadSuffix
}

// errUnkownEnum is returns when a value cannot match enum values
var (
	errUnkownEnum = errors.New("unkown enum")
	_             = errUnkownEnum
)

// needed to avoid when fmt package is not needed by generated code
var __dummy__fmt_variable fmt.Scanner

var _ = __dummy__fmt_variable

// idem for math package when not need by generated code
var __dummy_math_variable = math.E

var _ = __dummy_math_variable

// swagger:ignore
type __void any

// needed for creating set of instances in the stage
var (
	__member __void
	_        = __member
)

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

	// gongMarshallingMode set the marshalling mode
	gongMarshallingMode GongMarshallingMode
	// some stages have semantic rules that forbids them to be empty
	// like for git, the commit #0 (genesis commit) cannot be rolled back
	isWithGenesisCommit bool

	// insertion point for definition of arrays registering instances
	Chapters                map[*Chapter]struct{}
	Chapters_instance       map[*Chapter]*Chapter
	Chapters_mapString      map[string]*Chapter
	ChapterOrder            uint
	Chapter_stagedOrder     map[*Chapter]uint
	Chapter_orderStaged     map[uint]*Chapter
	Chapters_reference      map[*Chapter]*Chapter
	Chapters_referenceOrder map[*Chapter]uint

	// insertion point for slice of pointers maps
	Chapter_Sections_reverseMap map[*Section]*Chapter

	Chapter_Pages_reverseMap map[*Page]*Chapter

	Chapter_SubChapters_reverseMap map[*Chapter]*Chapter

	OnAfterChapterCreateCallback OnAfterCreateInterface[Chapter]
	OnAfterChapterUpdateCallback OnAfterUpdateInterface[Chapter]
	OnAfterChapterDeleteCallback OnAfterDeleteInterface[Chapter]
	OnAfterChapterReadCallback   OnAfterReadInterface[Chapter]

	Contents                map[*Content]struct{}
	Contents_instance       map[*Content]*Content
	Contents_mapString      map[string]*Content
	ContentOrder            uint
	Content_stagedOrder     map[*Content]uint
	Content_orderStaged     map[uint]*Content
	Contents_reference      map[*Content]*Content
	Contents_referenceOrder map[*Content]uint

	// insertion point for slice of pointers maps
	Content_Chapters_reverseMap map[*Chapter]*Content

	OnAfterContentCreateCallback OnAfterCreateInterface[Content]
	OnAfterContentUpdateCallback OnAfterUpdateInterface[Content]
	OnAfterContentDeleteCallback OnAfterDeleteInterface[Content]
	OnAfterContentReadCallback   OnAfterReadInterface[Content]

	DownloadableFiles                map[*DownloadableFile]struct{}
	DownloadableFiles_instance       map[*DownloadableFile]*DownloadableFile
	DownloadableFiles_mapString      map[string]*DownloadableFile
	DownloadableFileOrder            uint
	DownloadableFile_stagedOrder     map[*DownloadableFile]uint
	DownloadableFile_orderStaged     map[uint]*DownloadableFile
	DownloadableFiles_reference      map[*DownloadableFile]*DownloadableFile
	DownloadableFiles_referenceOrder map[*DownloadableFile]uint

	// insertion point for slice of pointers maps
	OnAfterDownloadableFileCreateCallback OnAfterCreateInterface[DownloadableFile]
	OnAfterDownloadableFileUpdateCallback OnAfterUpdateInterface[DownloadableFile]
	OnAfterDownloadableFileDeleteCallback OnAfterDeleteInterface[DownloadableFile]
	OnAfterDownloadableFileReadCallback   OnAfterReadInterface[DownloadableFile]

	JpgImages                map[*JpgImage]struct{}
	JpgImages_instance       map[*JpgImage]*JpgImage
	JpgImages_mapString      map[string]*JpgImage
	JpgImageOrder            uint
	JpgImage_stagedOrder     map[*JpgImage]uint
	JpgImage_orderStaged     map[uint]*JpgImage
	JpgImages_reference      map[*JpgImage]*JpgImage
	JpgImages_referenceOrder map[*JpgImage]uint

	// insertion point for slice of pointers maps
	OnAfterJpgImageCreateCallback OnAfterCreateInterface[JpgImage]
	OnAfterJpgImageUpdateCallback OnAfterUpdateInterface[JpgImage]
	OnAfterJpgImageDeleteCallback OnAfterDeleteInterface[JpgImage]
	OnAfterJpgImageReadCallback   OnAfterReadInterface[JpgImage]

	Pages                map[*Page]struct{}
	Pages_instance       map[*Page]*Page
	Pages_mapString      map[string]*Page
	PageOrder            uint
	Page_stagedOrder     map[*Page]uint
	Page_orderStaged     map[uint]*Page
	Pages_reference      map[*Page]*Page
	Pages_referenceOrder map[*Page]uint

	// insertion point for slice of pointers maps
	Page_Sections_reverseMap map[*Section]*Page

	OnAfterPageCreateCallback OnAfterCreateInterface[Page]
	OnAfterPageUpdateCallback OnAfterUpdateInterface[Page]
	OnAfterPageDeleteCallback OnAfterDeleteInterface[Page]
	OnAfterPageReadCallback   OnAfterReadInterface[Page]

	PngImages                map[*PngImage]struct{}
	PngImages_instance       map[*PngImage]*PngImage
	PngImages_mapString      map[string]*PngImage
	PngImageOrder            uint
	PngImage_stagedOrder     map[*PngImage]uint
	PngImage_orderStaged     map[uint]*PngImage
	PngImages_reference      map[*PngImage]*PngImage
	PngImages_referenceOrder map[*PngImage]uint

	// insertion point for slice of pointers maps
	OnAfterPngImageCreateCallback OnAfterCreateInterface[PngImage]
	OnAfterPngImageUpdateCallback OnAfterUpdateInterface[PngImage]
	OnAfterPngImageDeleteCallback OnAfterDeleteInterface[PngImage]
	OnAfterPngImageReadCallback   OnAfterReadInterface[PngImage]

	Sections                map[*Section]struct{}
	Sections_instance       map[*Section]*Section
	Sections_mapString      map[string]*Section
	SectionOrder            uint
	Section_stagedOrder     map[*Section]uint
	Section_orderStaged     map[uint]*Section
	Sections_reference      map[*Section]*Section
	Sections_referenceOrder map[*Section]uint

	// insertion point for slice of pointers maps
	OnAfterSectionCreateCallback OnAfterCreateInterface[Section]
	OnAfterSectionUpdateCallback OnAfterUpdateInterface[Section]
	OnAfterSectionDeleteCallback OnAfterDeleteInterface[Section]
	OnAfterSectionReadCallback   OnAfterReadInterface[Section]

	SvgImages                map[*SvgImage]struct{}
	SvgImages_instance       map[*SvgImage]*SvgImage
	SvgImages_mapString      map[string]*SvgImage
	SvgImageOrder            uint
	SvgImage_stagedOrder     map[*SvgImage]uint
	SvgImage_orderStaged     map[uint]*SvgImage
	SvgImages_reference      map[*SvgImage]*SvgImage
	SvgImages_referenceOrder map[*SvgImage]uint

	// insertion point for slice of pointers maps
	OnAfterSvgImageCreateCallback OnAfterCreateInterface[SvgImage]
	OnAfterSvgImageUpdateCallback OnAfterUpdateInterface[SvgImage]
	OnAfterSvgImageDeleteCallback OnAfterDeleteInterface[SvgImage]
	OnAfterSvgImageReadCallback   OnAfterReadInterface[SvgImage]

	AllModelsStructCreateCallback AllModelsStructCreateInterface

	AllModelsStructDeleteCallback AllModelsStructDeleteInterface

	BackRepo BackRepoInterface

	// if set will be called before each commit to the back repo
	OnInitCommitCallback          OnInitCommitInterface
	OnInitCommitFromFrontCallback OnInitCommitInterface
	OnInitCommitFromBackCallback  OnInitCommitInterface

	// Private slices to hold the registered hooks
	beforeCommitHooks []func(stage *Stage)
	afterCommitHooks  []func(stage *Stage)

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

	isApplyingBackwardCommit bool
	isApplyingForwardCommit  bool
	isSquashing              bool

	modified bool

	lock sync.RWMutex
}

func (s *Stage) SetGongMarshallingMode(mode GongMarshallingMode) {
	s.gongMarshallingMode = mode
}

func (s *Stage) GetGongMarshallingMode() GongMarshallingMode {
	return s.gongMarshallingMode
}

func (s *Stage) SetIsWithGenesisCommit(isWithGenesisCommit bool) {
	s.isWithGenesisCommit = isWithGenesisCommit
}

func (s *Stage) GetIsWithGenesisCommit() bool {
	return s.isWithGenesisCommit
}

// RegisterBeforeCommit adds a hook that runs before the commit happens
func (s *Stage) RegisterBeforeCommit(hook func(stage *Stage)) {
	s.beforeCommitHooks = append(s.beforeCommitHooks, hook)
}

// RegisterAfterCommit adds a hook that runs after the commit succeeds
func (s *Stage) RegisterAfterCommit(hook func(stage *Stage)) {
	s.afterCommitHooks = append(s.afterCommitHooks, hook)
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

	if stage.isWithGenesisCommit && stage.commitsBehind >= len(stage.backwardCommits)-1 {
		return errors.New("cannot rollback genesis commit")
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
	stage.isApplyingBackwardCommit = true
	err := GongParseAstString(stage, commitToApply, true)
	stage.isApplyingBackwardCommit = false
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
	stage.isApplyingForwardCommit = true
	err := GongParseAstString(stage, commitToApply, true)
	stage.isApplyingForwardCommit = false
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

func (stage *Stage) Lock() {
	stage.lock.Lock()
}

func (stage *Stage) Unlock() {
	stage.lock.Unlock()
}

func (stage *Stage) RLock() {
	stage.lock.RLock()
}

func (stage *Stage) RUnlock() {
	stage.lock.RUnlock()
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
	if stage.OnInitCommitCallback != nil {
		stage.OnInitCommitCallback.BeforeCommit(stage)
	}
	if stage.OnInitCommitFromBackCallback != nil {
		stage.OnInitCommitFromBackCallback.BeforeCommit(stage)
	}

	// 1. Run all Before Commit hooks
	for _, hook := range stage.beforeCommitHooks {
		hook(stage)
	}

	// 2. Run all After Commit hooks
	for _, hook := range stage.afterCommitHooks {
		hook(stage)
	}
}

// Squash removes all commits and marshals the stage as a single commit
func (stage *Stage) Squash() {
	stage.forwardCommits = stage.forwardCommits[:0]
	stage.backwardCommits = stage.backwardCommits[:0]
	stage.commitsBehind = 0
	stage.navigationMode = GongNavigationModeNormal

	stage.modified = true
	stage.isSquashing = true

	// insertion point for clear references
	stage.Chapters_reference = make(map[*Chapter]*Chapter)
	stage.Chapters_instance = make(map[*Chapter]*Chapter)
	stage.Chapters_referenceOrder = make(map[*Chapter]uint)

	stage.Contents_reference = make(map[*Content]*Content)
	stage.Contents_instance = make(map[*Content]*Content)
	stage.Contents_referenceOrder = make(map[*Content]uint)

	stage.DownloadableFiles_reference = make(map[*DownloadableFile]*DownloadableFile)
	stage.DownloadableFiles_instance = make(map[*DownloadableFile]*DownloadableFile)
	stage.DownloadableFiles_referenceOrder = make(map[*DownloadableFile]uint)

	stage.JpgImages_reference = make(map[*JpgImage]*JpgImage)
	stage.JpgImages_instance = make(map[*JpgImage]*JpgImage)
	stage.JpgImages_referenceOrder = make(map[*JpgImage]uint)

	stage.Pages_reference = make(map[*Page]*Page)
	stage.Pages_instance = make(map[*Page]*Page)
	stage.Pages_referenceOrder = make(map[*Page]uint)

	stage.PngImages_reference = make(map[*PngImage]*PngImage)
	stage.PngImages_instance = make(map[*PngImage]*PngImage)
	stage.PngImages_referenceOrder = make(map[*PngImage]uint)

	stage.Sections_reference = make(map[*Section]*Section)
	stage.Sections_instance = make(map[*Section]*Section)
	stage.Sections_referenceOrder = make(map[*Section]uint)

	stage.SvgImages_reference = make(map[*SvgImage]*SvgImage)
	stage.SvgImages_instance = make(map[*SvgImage]*SvgImage)
	stage.SvgImages_referenceOrder = make(map[*SvgImage]uint)

	stage.ComputeInstancesNb()
	if stage.OnInitCommitCallback != nil {
		stage.OnInitCommitCallback.BeforeCommit(stage)
	}
	if stage.OnInitCommitFromBackCallback != nil {
		stage.OnInitCommitFromBackCallback.BeforeCommit(stage)
	}

	// 1. Run all Before Commit hooks
	for _, hook := range stage.beforeCommitHooks {
		hook(stage)
	}

	// 2. Run all After Commit hooks
	for _, hook := range stage.afterCommitHooks {
		hook(stage)
	}

	stage.isSquashing = false
}

// recomputeOrders recomputes the next order for each struct
// this is necessary because the order might have been incremented
// during the commits that have been discarded
// insertion point for max order recomputation
func (stage *Stage) recomputeOrders() {
	// insertion point for max order recomputation
	var maxChapterOrder uint
	var foundChapter bool
	for _, order := range stage.Chapter_stagedOrder {
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
	for _, order := range stage.Content_stagedOrder {
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

	var maxDownloadableFileOrder uint
	var foundDownloadableFile bool
	for _, order := range stage.DownloadableFile_stagedOrder {
		if !foundDownloadableFile || order > maxDownloadableFileOrder {
			maxDownloadableFileOrder = order
			foundDownloadableFile = true
		}
	}
	if foundDownloadableFile {
		stage.DownloadableFileOrder = maxDownloadableFileOrder + 1
	} else {
		stage.DownloadableFileOrder = 0
	}

	var maxJpgImageOrder uint
	var foundJpgImage bool
	for _, order := range stage.JpgImage_stagedOrder {
		if !foundJpgImage || order > maxJpgImageOrder {
			maxJpgImageOrder = order
			foundJpgImage = true
		}
	}
	if foundJpgImage {
		stage.JpgImageOrder = maxJpgImageOrder + 1
	} else {
		stage.JpgImageOrder = 0
	}

	var maxPageOrder uint
	var foundPage bool
	for _, order := range stage.Page_stagedOrder {
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

	var maxPngImageOrder uint
	var foundPngImage bool
	for _, order := range stage.PngImage_stagedOrder {
		if !foundPngImage || order > maxPngImageOrder {
			maxPngImageOrder = order
			foundPngImage = true
		}
	}
	if foundPngImage {
		stage.PngImageOrder = maxPngImageOrder + 1
	} else {
		stage.PngImageOrder = 0
	}

	var maxSectionOrder uint
	var foundSection bool
	for _, order := range stage.Section_stagedOrder {
		if !foundSection || order > maxSectionOrder {
			maxSectionOrder = order
			foundSection = true
		}
	}
	if foundSection {
		stage.SectionOrder = maxSectionOrder + 1
	} else {
		stage.SectionOrder = 0
	}

	var maxSvgImageOrder uint
	var foundSvgImage bool
	for _, order := range stage.SvgImage_stagedOrder {
		if !foundSvgImage || order > maxSvgImageOrder {
			maxSvgImageOrder = order
			foundSvgImage = true
		}
	}
	if foundSvgImage {
		stage.SvgImageOrder = maxSvgImageOrder + 1
	} else {
		stage.SvgImageOrder = 0
	}

	// end of insertion point for max order recomputation
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

// GetStructInstancesByOrderAuto returns a slice of generic pointers to gongstructs
// ordered by their order in the stage.
func GetStructInstancesByOrderAuto[T PointerToGongstruct](stage *Stage) (res []T) {
	var t T
	switch any(t).(type) {
	// insertion point for case
	case *Chapter:
		tmp := GetStructInstancesByOrder(stage.Chapters, stage.Chapter_stagedOrder)

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
		tmp := GetStructInstancesByOrder(stage.Contents, stage.Content_stagedOrder)

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
	case *DownloadableFile:
		tmp := GetStructInstancesByOrder(stage.DownloadableFiles, stage.DownloadableFile_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *DownloadableFile implements.
			res = append(res, any(v).(T))
		}
		return res
	case *JpgImage:
		tmp := GetStructInstancesByOrder(stage.JpgImages, stage.JpgImage_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *JpgImage implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Page:
		tmp := GetStructInstancesByOrder(stage.Pages, stage.Page_stagedOrder)

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
	case *PngImage:
		tmp := GetStructInstancesByOrder(stage.PngImages, stage.PngImage_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *PngImage implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Section:
		tmp := GetStructInstancesByOrder(stage.Sections, stage.Section_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Section implements.
			res = append(res, any(v).(T))
		}
		return res
	case *SvgImage:
		tmp := GetStructInstancesByOrder(stage.SvgImages, stage.SvgImage_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *SvgImage implements.
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
		res = GetNamedStructInstances(stage.Chapters, stage.Chapter_stagedOrder)
	case "Content":
		res = GetNamedStructInstances(stage.Contents, stage.Content_stagedOrder)
	case "DownloadableFile":
		res = GetNamedStructInstances(stage.DownloadableFiles, stage.DownloadableFile_stagedOrder)
	case "JpgImage":
		res = GetNamedStructInstances(stage.JpgImages, stage.JpgImage_stagedOrder)
	case "Page":
		res = GetNamedStructInstances(stage.Pages, stage.Page_stagedOrder)
	case "PngImage":
		res = GetNamedStructInstances(stage.PngImages, stage.PngImage_stagedOrder)
	case "Section":
		res = GetNamedStructInstances(stage.Sections, stage.Section_stagedOrder)
	case "SvgImage":
		res = GetNamedStructInstances(stage.SvgImages, stage.SvgImage_stagedOrder)
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
	CommitDownloadableFile(downloadablefile *DownloadableFile)
	CheckoutDownloadableFile(downloadablefile *DownloadableFile)
	CommitJpgImage(jpgimage *JpgImage)
	CheckoutJpgImage(jpgimage *JpgImage)
	CommitPage(page *Page)
	CheckoutPage(page *Page)
	CommitPngImage(pngimage *PngImage)
	CheckoutPngImage(pngimage *PngImage)
	CommitSection(section *Section)
	CheckoutSection(section *Section)
	CommitSvgImage(svgimage *SvgImage)
	CheckoutSvgImage(svgimage *SvgImage)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(name string) (stage *Stage) {
	stage = &Stage{ // insertion point for array initiatialisation
		Chapters:           make(map[*Chapter]struct{}),
		Chapters_mapString: make(map[string]*Chapter),

		Contents:           make(map[*Content]struct{}),
		Contents_mapString: make(map[string]*Content),

		DownloadableFiles:           make(map[*DownloadableFile]struct{}),
		DownloadableFiles_mapString: make(map[string]*DownloadableFile),

		JpgImages:           make(map[*JpgImage]struct{}),
		JpgImages_mapString: make(map[string]*JpgImage),

		Pages:           make(map[*Page]struct{}),
		Pages_mapString: make(map[string]*Page),

		PngImages:           make(map[*PngImage]struct{}),
		PngImages_mapString: make(map[string]*PngImage),

		Sections:           make(map[*Section]struct{}),
		Sections_mapString: make(map[string]*Section),

		SvgImages:           make(map[*SvgImage]struct{}),
		SvgImages_mapString: make(map[string]*SvgImage),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		name: name,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		Chapter_stagedOrder: make(map[*Chapter]uint),
		Chapter_orderStaged: make(map[uint]*Chapter),
		Chapters_reference:  make(map[*Chapter]*Chapter),

		Content_stagedOrder: make(map[*Content]uint),
		Content_orderStaged: make(map[uint]*Content),
		Contents_reference:  make(map[*Content]*Content),

		DownloadableFile_stagedOrder: make(map[*DownloadableFile]uint),
		DownloadableFile_orderStaged: make(map[uint]*DownloadableFile),
		DownloadableFiles_reference:  make(map[*DownloadableFile]*DownloadableFile),

		JpgImage_stagedOrder: make(map[*JpgImage]uint),
		JpgImage_orderStaged: make(map[uint]*JpgImage),
		JpgImages_reference:  make(map[*JpgImage]*JpgImage),

		Page_stagedOrder: make(map[*Page]uint),
		Page_orderStaged: make(map[uint]*Page),
		Pages_reference:  make(map[*Page]*Page),

		PngImage_stagedOrder: make(map[*PngImage]uint),
		PngImage_orderStaged: make(map[uint]*PngImage),
		PngImages_reference:  make(map[*PngImage]*PngImage),

		Section_stagedOrder: make(map[*Section]uint),
		Section_orderStaged: make(map[uint]*Section),
		Sections_reference:  make(map[*Section]*Section),

		SvgImage_stagedOrder: make(map[*SvgImage]uint),
		SvgImage_orderStaged: make(map[uint]*SvgImage),
		SvgImages_reference:  make(map[*SvgImage]*SvgImage),

		// end of insertion point
		GongUnmarshallers: map[string]ModelUnmarshaller{ // insertion point for unmarshallers
			"Chapter": &ChapterUnmarshaller{},

			"Content": &ContentUnmarshaller{},

			"DownloadableFile": &DownloadableFileUnmarshaller{},

			"JpgImage": &JpgImageUnmarshaller{},

			"Page": &PageUnmarshaller{},

			"PngImage": &PngImageUnmarshaller{},

			"Section": &SectionUnmarshaller{},

			"SvgImage": &SvgImageUnmarshaller{},

			// end of insertion point
		},

		NamedStructs: []*NamedStruct{ // insertion point for order map initialisations
			{name: "Chapter"},
			{name: "Content"},
			{name: "DownloadableFile"},
			{name: "JpgImage"},
			{name: "Page"},
			{name: "PngImage"},
			{name: "Section"},
			{name: "SvgImage"},
		}, // end of insertion point

		navigationMode: GongNavigationModeNormal,
	}

	return
}

func GetOrder[Type Gongstruct](stage *Stage, instance *Type) uint {
	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *Chapter:
		return stage.Chapter_stagedOrder[instance]
	case *Content:
		return stage.Content_stagedOrder[instance]
	case *DownloadableFile:
		return stage.DownloadableFile_stagedOrder[instance]
	case *JpgImage:
		return stage.JpgImage_stagedOrder[instance]
	case *Page:
		return stage.Page_stagedOrder[instance]
	case *PngImage:
		return stage.PngImage_stagedOrder[instance]
	case *Section:
		return stage.Section_stagedOrder[instance]
	case *SvgImage:
		return stage.SvgImage_stagedOrder[instance]
	default:
		return 0 // should not happen
	}
}

func GongGetInstanceFromOrder[Type PointerToGongstruct](stage *Stage, order uint) (res Type) {
	var t Type
	switch any(t).(type) {
	// insertion point for order map initialisations
	case *Chapter:
		return any(stage.Chapter_orderStaged[order]).(Type)
	case *Content:
		return any(stage.Content_orderStaged[order]).(Type)
	case *DownloadableFile:
		return any(stage.DownloadableFile_orderStaged[order]).(Type)
	case *JpgImage:
		return any(stage.JpgImage_orderStaged[order]).(Type)
	case *Page:
		return any(stage.Page_orderStaged[order]).(Type)
	case *PngImage:
		return any(stage.PngImage_orderStaged[order]).(Type)
	case *Section:
		return any(stage.Section_orderStaged[order]).(Type)
	case *SvgImage:
		return any(stage.SvgImage_orderStaged[order]).(Type)
	default:
		return // should not happen
	}
}

func GetOrderPointerGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) uint {
	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *Chapter:
		return stage.Chapter_stagedOrder[instance]
	case *Content:
		return stage.Content_stagedOrder[instance]
	case *DownloadableFile:
		return stage.DownloadableFile_stagedOrder[instance]
	case *JpgImage:
		return stage.JpgImage_stagedOrder[instance]
	case *Page:
		return stage.Page_stagedOrder[instance]
	case *PngImage:
		return stage.PngImage_stagedOrder[instance]
	case *Section:
		return stage.Section_stagedOrder[instance]
	case *SvgImage:
		return stage.SvgImage_stagedOrder[instance]
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
	tmp2 := stage.beforeCommitHooks
	stage.beforeCommitHooks = nil
	tmp3 := stage.afterCommitHooks
	stage.afterCommitHooks = nil
	stage.Commit()
	stage.OnInitCommitFromBackCallback = tmp
	stage.beforeCommitHooks = tmp2
	stage.afterCommitHooks = tmp3
}

func (stage *Stage) Commit() {
	stage.ComputeReverseMaps()

	if stage.OnInitCommitCallback != nil {
		stage.OnInitCommitCallback.BeforeCommit(stage)
	}
	if stage.OnInitCommitFromBackCallback != nil {
		stage.OnInitCommitFromBackCallback.BeforeCommit(stage)
	}

	// 1. Run all Before Commit hooks
	for _, hook := range stage.beforeCommitHooks {
		hook(stage)
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
		if stage.probeIF != nil {
			stage.probeIF.RefreshNavigationTree()
		}
	}

	// 2. Run all After Commit hooks
	for _, hook := range stage.afterCommitHooks {
		hook(stage)
	}
}

func (stage *Stage) ComputeInstancesNb() {
	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["Chapter"] = len(stage.Chapters)
	stage.Map_GongStructName_InstancesNb["Content"] = len(stage.Contents)
	stage.Map_GongStructName_InstancesNb["DownloadableFile"] = len(stage.DownloadableFiles)
	stage.Map_GongStructName_InstancesNb["JpgImage"] = len(stage.JpgImages)
	stage.Map_GongStructName_InstancesNb["Page"] = len(stage.Pages)
	stage.Map_GongStructName_InstancesNb["PngImage"] = len(stage.PngImages)
	stage.Map_GongStructName_InstancesNb["Section"] = len(stage.Sections)
	stage.Map_GongStructName_InstancesNb["SvgImage"] = len(stage.SvgImages)
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
		stage.Chapter_stagedOrder[chapter] = stage.ChapterOrder
		stage.Chapter_orderStaged[stage.ChapterOrder] = chapter
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
		stage.Chapter_stagedOrder[chapter] = order
		stage.Chapter_orderStaged[order] = chapter
		stage.ChapterOrder++
	}
	stage.Chapters_mapString[chapter.Name] = chapter
}

// Unstage removes chapter off the model stage
func (chapter *Chapter) Unstage(stage *Stage) *Chapter {
	delete(stage.Chapters, chapter)
	// issue1150
	// delete(stage.Chapter_stagedOrder, chapter)
	delete(stage.Chapters_mapString, chapter.Name)

	return chapter
}

// UnstageVoid removes chapter off the model stage
func (chapter *Chapter) UnstageVoid(stage *Stage) {
	delete(stage.Chapters, chapter)
	// issue1150
	// delete(stage.Chapter_stagedOrder, chapter)
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
		stage.Content_stagedOrder[content] = stage.ContentOrder
		stage.Content_orderStaged[stage.ContentOrder] = content
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
		stage.Content_stagedOrder[content] = order
		stage.Content_orderStaged[order] = content
		stage.ContentOrder++
	}
	stage.Contents_mapString[content.Name] = content
}

// Unstage removes content off the model stage
func (content *Content) Unstage(stage *Stage) *Content {
	delete(stage.Contents, content)
	// issue1150
	// delete(stage.Content_stagedOrder, content)
	delete(stage.Contents_mapString, content.Name)

	return content
}

// UnstageVoid removes content off the model stage
func (content *Content) UnstageVoid(stage *Stage) {
	delete(stage.Contents, content)
	// issue1150
	// delete(stage.Content_stagedOrder, content)
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

// Stage puts downloadablefile to the model stage
func (downloadablefile *DownloadableFile) Stage(stage *Stage) *DownloadableFile {
	if _, ok := stage.DownloadableFiles[downloadablefile]; !ok {
		stage.DownloadableFiles[downloadablefile] = struct{}{}
		stage.DownloadableFile_stagedOrder[downloadablefile] = stage.DownloadableFileOrder
		stage.DownloadableFile_orderStaged[stage.DownloadableFileOrder] = downloadablefile
		stage.DownloadableFileOrder++
	}
	stage.DownloadableFiles_mapString[downloadablefile.Name] = downloadablefile

	return downloadablefile
}

// StagePreserveOrder puts downloadablefile to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DownloadableFileOrder
// - update stage.DownloadableFileOrder accordingly
func (downloadablefile *DownloadableFile) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.DownloadableFiles[downloadablefile]; !ok {
		stage.DownloadableFiles[downloadablefile] = struct{}{}

		if order > stage.DownloadableFileOrder {
			stage.DownloadableFileOrder = order
		}
		stage.DownloadableFile_stagedOrder[downloadablefile] = order
		stage.DownloadableFile_orderStaged[order] = downloadablefile
		stage.DownloadableFileOrder++
	}
	stage.DownloadableFiles_mapString[downloadablefile.Name] = downloadablefile
}

// Unstage removes downloadablefile off the model stage
func (downloadablefile *DownloadableFile) Unstage(stage *Stage) *DownloadableFile {
	delete(stage.DownloadableFiles, downloadablefile)
	// issue1150
	// delete(stage.DownloadableFile_stagedOrder, downloadablefile)
	delete(stage.DownloadableFiles_mapString, downloadablefile.Name)

	return downloadablefile
}

// UnstageVoid removes downloadablefile off the model stage
func (downloadablefile *DownloadableFile) UnstageVoid(stage *Stage) {
	delete(stage.DownloadableFiles, downloadablefile)
	// issue1150
	// delete(stage.DownloadableFile_stagedOrder, downloadablefile)
	delete(stage.DownloadableFiles_mapString, downloadablefile.Name)
}

// commit downloadablefile to the back repo (if it is already staged)
func (downloadablefile *DownloadableFile) Commit(stage *Stage) *DownloadableFile {
	if _, ok := stage.DownloadableFiles[downloadablefile]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDownloadableFile(downloadablefile)
		}
	}
	return downloadablefile
}

func (downloadablefile *DownloadableFile) CommitVoid(stage *Stage) {
	downloadablefile.Commit(stage)
}

func (downloadablefile *DownloadableFile) StageVoid(stage *Stage) {
	downloadablefile.Stage(stage)
}

// Checkout downloadablefile to the back repo (if it is already staged)
func (downloadablefile *DownloadableFile) Checkout(stage *Stage) *DownloadableFile {
	if _, ok := stage.DownloadableFiles[downloadablefile]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutDownloadableFile(downloadablefile)
		}
	}
	return downloadablefile
}

// for satisfaction of GongStruct interface
func (downloadablefile *DownloadableFile) GetName() (res string) {
	return downloadablefile.Name
}

// for satisfaction of GongStruct interface
func (downloadablefile *DownloadableFile) SetName(name string) {
	downloadablefile.Name = name
}

// Stage puts jpgimage to the model stage
func (jpgimage *JpgImage) Stage(stage *Stage) *JpgImage {
	if _, ok := stage.JpgImages[jpgimage]; !ok {
		stage.JpgImages[jpgimage] = struct{}{}
		stage.JpgImage_stagedOrder[jpgimage] = stage.JpgImageOrder
		stage.JpgImage_orderStaged[stage.JpgImageOrder] = jpgimage
		stage.JpgImageOrder++
	}
	stage.JpgImages_mapString[jpgimage.Name] = jpgimage

	return jpgimage
}

// StagePreserveOrder puts jpgimage to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.JpgImageOrder
// - update stage.JpgImageOrder accordingly
func (jpgimage *JpgImage) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.JpgImages[jpgimage]; !ok {
		stage.JpgImages[jpgimage] = struct{}{}

		if order > stage.JpgImageOrder {
			stage.JpgImageOrder = order
		}
		stage.JpgImage_stagedOrder[jpgimage] = order
		stage.JpgImage_orderStaged[order] = jpgimage
		stage.JpgImageOrder++
	}
	stage.JpgImages_mapString[jpgimage.Name] = jpgimage
}

// Unstage removes jpgimage off the model stage
func (jpgimage *JpgImage) Unstage(stage *Stage) *JpgImage {
	delete(stage.JpgImages, jpgimage)
	// issue1150
	// delete(stage.JpgImage_stagedOrder, jpgimage)
	delete(stage.JpgImages_mapString, jpgimage.Name)

	return jpgimage
}

// UnstageVoid removes jpgimage off the model stage
func (jpgimage *JpgImage) UnstageVoid(stage *Stage) {
	delete(stage.JpgImages, jpgimage)
	// issue1150
	// delete(stage.JpgImage_stagedOrder, jpgimage)
	delete(stage.JpgImages_mapString, jpgimage.Name)
}

// commit jpgimage to the back repo (if it is already staged)
func (jpgimage *JpgImage) Commit(stage *Stage) *JpgImage {
	if _, ok := stage.JpgImages[jpgimage]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitJpgImage(jpgimage)
		}
	}
	return jpgimage
}

func (jpgimage *JpgImage) CommitVoid(stage *Stage) {
	jpgimage.Commit(stage)
}

func (jpgimage *JpgImage) StageVoid(stage *Stage) {
	jpgimage.Stage(stage)
}

// Checkout jpgimage to the back repo (if it is already staged)
func (jpgimage *JpgImage) Checkout(stage *Stage) *JpgImage {
	if _, ok := stage.JpgImages[jpgimage]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutJpgImage(jpgimage)
		}
	}
	return jpgimage
}

// for satisfaction of GongStruct interface
func (jpgimage *JpgImage) GetName() (res string) {
	return jpgimage.Name
}

// for satisfaction of GongStruct interface
func (jpgimage *JpgImage) SetName(name string) {
	jpgimage.Name = name
}

// Stage puts page to the model stage
func (page *Page) Stage(stage *Stage) *Page {
	if _, ok := stage.Pages[page]; !ok {
		stage.Pages[page] = struct{}{}
		stage.Page_stagedOrder[page] = stage.PageOrder
		stage.Page_orderStaged[stage.PageOrder] = page
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
		stage.Page_stagedOrder[page] = order
		stage.Page_orderStaged[order] = page
		stage.PageOrder++
	}
	stage.Pages_mapString[page.Name] = page
}

// Unstage removes page off the model stage
func (page *Page) Unstage(stage *Stage) *Page {
	delete(stage.Pages, page)
	// issue1150
	// delete(stage.Page_stagedOrder, page)
	delete(stage.Pages_mapString, page.Name)

	return page
}

// UnstageVoid removes page off the model stage
func (page *Page) UnstageVoid(stage *Stage) {
	delete(stage.Pages, page)
	// issue1150
	// delete(stage.Page_stagedOrder, page)
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

// Stage puts pngimage to the model stage
func (pngimage *PngImage) Stage(stage *Stage) *PngImage {
	if _, ok := stage.PngImages[pngimage]; !ok {
		stage.PngImages[pngimage] = struct{}{}
		stage.PngImage_stagedOrder[pngimage] = stage.PngImageOrder
		stage.PngImage_orderStaged[stage.PngImageOrder] = pngimage
		stage.PngImageOrder++
	}
	stage.PngImages_mapString[pngimage.Name] = pngimage

	return pngimage
}

// StagePreserveOrder puts pngimage to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.PngImageOrder
// - update stage.PngImageOrder accordingly
func (pngimage *PngImage) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.PngImages[pngimage]; !ok {
		stage.PngImages[pngimage] = struct{}{}

		if order > stage.PngImageOrder {
			stage.PngImageOrder = order
		}
		stage.PngImage_stagedOrder[pngimage] = order
		stage.PngImage_orderStaged[order] = pngimage
		stage.PngImageOrder++
	}
	stage.PngImages_mapString[pngimage.Name] = pngimage
}

// Unstage removes pngimage off the model stage
func (pngimage *PngImage) Unstage(stage *Stage) *PngImage {
	delete(stage.PngImages, pngimage)
	// issue1150
	// delete(stage.PngImage_stagedOrder, pngimage)
	delete(stage.PngImages_mapString, pngimage.Name)

	return pngimage
}

// UnstageVoid removes pngimage off the model stage
func (pngimage *PngImage) UnstageVoid(stage *Stage) {
	delete(stage.PngImages, pngimage)
	// issue1150
	// delete(stage.PngImage_stagedOrder, pngimage)
	delete(stage.PngImages_mapString, pngimage.Name)
}

// commit pngimage to the back repo (if it is already staged)
func (pngimage *PngImage) Commit(stage *Stage) *PngImage {
	if _, ok := stage.PngImages[pngimage]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitPngImage(pngimage)
		}
	}
	return pngimage
}

func (pngimage *PngImage) CommitVoid(stage *Stage) {
	pngimage.Commit(stage)
}

func (pngimage *PngImage) StageVoid(stage *Stage) {
	pngimage.Stage(stage)
}

// Checkout pngimage to the back repo (if it is already staged)
func (pngimage *PngImage) Checkout(stage *Stage) *PngImage {
	if _, ok := stage.PngImages[pngimage]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutPngImage(pngimage)
		}
	}
	return pngimage
}

// for satisfaction of GongStruct interface
func (pngimage *PngImage) GetName() (res string) {
	return pngimage.Name
}

// for satisfaction of GongStruct interface
func (pngimage *PngImage) SetName(name string) {
	pngimage.Name = name
}

// Stage puts section to the model stage
func (section *Section) Stage(stage *Stage) *Section {
	if _, ok := stage.Sections[section]; !ok {
		stage.Sections[section] = struct{}{}
		stage.Section_stagedOrder[section] = stage.SectionOrder
		stage.Section_orderStaged[stage.SectionOrder] = section
		stage.SectionOrder++
	}
	stage.Sections_mapString[section.Name] = section

	return section
}

// StagePreserveOrder puts section to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.SectionOrder
// - update stage.SectionOrder accordingly
func (section *Section) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Sections[section]; !ok {
		stage.Sections[section] = struct{}{}

		if order > stage.SectionOrder {
			stage.SectionOrder = order
		}
		stage.Section_stagedOrder[section] = order
		stage.Section_orderStaged[order] = section
		stage.SectionOrder++
	}
	stage.Sections_mapString[section.Name] = section
}

// Unstage removes section off the model stage
func (section *Section) Unstage(stage *Stage) *Section {
	delete(stage.Sections, section)
	// issue1150
	// delete(stage.Section_stagedOrder, section)
	delete(stage.Sections_mapString, section.Name)

	return section
}

// UnstageVoid removes section off the model stage
func (section *Section) UnstageVoid(stage *Stage) {
	delete(stage.Sections, section)
	// issue1150
	// delete(stage.Section_stagedOrder, section)
	delete(stage.Sections_mapString, section.Name)
}

// commit section to the back repo (if it is already staged)
func (section *Section) Commit(stage *Stage) *Section {
	if _, ok := stage.Sections[section]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSection(section)
		}
	}
	return section
}

func (section *Section) CommitVoid(stage *Stage) {
	section.Commit(stage)
}

func (section *Section) StageVoid(stage *Stage) {
	section.Stage(stage)
}

// Checkout section to the back repo (if it is already staged)
func (section *Section) Checkout(stage *Stage) *Section {
	if _, ok := stage.Sections[section]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSection(section)
		}
	}
	return section
}

// for satisfaction of GongStruct interface
func (section *Section) GetName() (res string) {
	return section.Name
}

// for satisfaction of GongStruct interface
func (section *Section) SetName(name string) {
	section.Name = name
}

// Stage puts svgimage to the model stage
func (svgimage *SvgImage) Stage(stage *Stage) *SvgImage {
	if _, ok := stage.SvgImages[svgimage]; !ok {
		stage.SvgImages[svgimage] = struct{}{}
		stage.SvgImage_stagedOrder[svgimage] = stage.SvgImageOrder
		stage.SvgImage_orderStaged[stage.SvgImageOrder] = svgimage
		stage.SvgImageOrder++
	}
	stage.SvgImages_mapString[svgimage.Name] = svgimage

	return svgimage
}

// StagePreserveOrder puts svgimage to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.SvgImageOrder
// - update stage.SvgImageOrder accordingly
func (svgimage *SvgImage) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.SvgImages[svgimage]; !ok {
		stage.SvgImages[svgimage] = struct{}{}

		if order > stage.SvgImageOrder {
			stage.SvgImageOrder = order
		}
		stage.SvgImage_stagedOrder[svgimage] = order
		stage.SvgImage_orderStaged[order] = svgimage
		stage.SvgImageOrder++
	}
	stage.SvgImages_mapString[svgimage.Name] = svgimage
}

// Unstage removes svgimage off the model stage
func (svgimage *SvgImage) Unstage(stage *Stage) *SvgImage {
	delete(stage.SvgImages, svgimage)
	// issue1150
	// delete(stage.SvgImage_stagedOrder, svgimage)
	delete(stage.SvgImages_mapString, svgimage.Name)

	return svgimage
}

// UnstageVoid removes svgimage off the model stage
func (svgimage *SvgImage) UnstageVoid(stage *Stage) {
	delete(stage.SvgImages, svgimage)
	// issue1150
	// delete(stage.SvgImage_stagedOrder, svgimage)
	delete(stage.SvgImages_mapString, svgimage.Name)
}

// commit svgimage to the back repo (if it is already staged)
func (svgimage *SvgImage) Commit(stage *Stage) *SvgImage {
	if _, ok := stage.SvgImages[svgimage]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSvgImage(svgimage)
		}
	}
	return svgimage
}

func (svgimage *SvgImage) CommitVoid(stage *Stage) {
	svgimage.Commit(stage)
}

func (svgimage *SvgImage) StageVoid(stage *Stage) {
	svgimage.Stage(stage)
}

// Checkout svgimage to the back repo (if it is already staged)
func (svgimage *SvgImage) Checkout(stage *Stage) *SvgImage {
	if _, ok := stage.SvgImages[svgimage]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSvgImage(svgimage)
		}
	}
	return svgimage
}

// for satisfaction of GongStruct interface
func (svgimage *SvgImage) GetName() (res string) {
	return svgimage.Name
}

// for satisfaction of GongStruct interface
func (svgimage *SvgImage) SetName(name string) {
	svgimage.Name = name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMChapter(Chapter *Chapter)
	CreateORMContent(Content *Content)
	CreateORMDownloadableFile(DownloadableFile *DownloadableFile)
	CreateORMJpgImage(JpgImage *JpgImage)
	CreateORMPage(Page *Page)
	CreateORMPngImage(PngImage *PngImage)
	CreateORMSection(Section *Section)
	CreateORMSvgImage(SvgImage *SvgImage)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMChapter(Chapter *Chapter)
	DeleteORMContent(Content *Content)
	DeleteORMDownloadableFile(DownloadableFile *DownloadableFile)
	DeleteORMJpgImage(JpgImage *JpgImage)
	DeleteORMPage(Page *Page)
	DeleteORMPngImage(PngImage *PngImage)
	DeleteORMSection(Section *Section)
	DeleteORMSvgImage(SvgImage *SvgImage)
}

func (stage *Stage) Reset() { // insertion point for array reset
	stage.Chapters = make(map[*Chapter]struct{})
	stage.Chapters_mapString = make(map[string]*Chapter)
	stage.Chapter_stagedOrder = make(map[*Chapter]uint)
	stage.ChapterOrder = 0

	stage.Contents = make(map[*Content]struct{})
	stage.Contents_mapString = make(map[string]*Content)
	stage.Content_stagedOrder = make(map[*Content]uint)
	stage.ContentOrder = 0

	stage.DownloadableFiles = make(map[*DownloadableFile]struct{})
	stage.DownloadableFiles_mapString = make(map[string]*DownloadableFile)
	stage.DownloadableFile_stagedOrder = make(map[*DownloadableFile]uint)
	stage.DownloadableFileOrder = 0

	stage.JpgImages = make(map[*JpgImage]struct{})
	stage.JpgImages_mapString = make(map[string]*JpgImage)
	stage.JpgImage_stagedOrder = make(map[*JpgImage]uint)
	stage.JpgImageOrder = 0

	stage.Pages = make(map[*Page]struct{})
	stage.Pages_mapString = make(map[string]*Page)
	stage.Page_stagedOrder = make(map[*Page]uint)
	stage.PageOrder = 0

	stage.PngImages = make(map[*PngImage]struct{})
	stage.PngImages_mapString = make(map[string]*PngImage)
	stage.PngImage_stagedOrder = make(map[*PngImage]uint)
	stage.PngImageOrder = 0

	stage.Sections = make(map[*Section]struct{})
	stage.Sections_mapString = make(map[string]*Section)
	stage.Section_stagedOrder = make(map[*Section]uint)
	stage.SectionOrder = 0

	stage.SvgImages = make(map[*SvgImage]struct{})
	stage.SvgImages_mapString = make(map[string]*SvgImage)
	stage.SvgImage_stagedOrder = make(map[*SvgImage]uint)
	stage.SvgImageOrder = 0

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

	stage.DownloadableFiles = nil
	stage.DownloadableFiles_mapString = nil

	stage.JpgImages = nil
	stage.JpgImages_mapString = nil

	stage.Pages = nil
	stage.Pages_mapString = nil

	stage.PngImages = nil
	stage.PngImages_mapString = nil

	stage.Sections = nil
	stage.Sections_mapString = nil

	stage.SvgImages = nil
	stage.SvgImages_mapString = nil

	// end of insertion point for array nil
}

func (stage *Stage) Unstage() { // insertion point for array nil
	for chapter := range stage.Chapters {
		chapter.Unstage(stage)
	}

	for content := range stage.Contents {
		content.Unstage(stage)
	}

	for downloadablefile := range stage.DownloadableFiles {
		downloadablefile.Unstage(stage)
	}

	for jpgimage := range stage.JpgImages {
		jpgimage.Unstage(stage)
	}

	for page := range stage.Pages {
		page.Unstage(stage)
	}

	for pngimage := range stage.PngImages {
		pngimage.Unstage(stage)
	}

	for section := range stage.Sections {
		section.Unstage(stage)
	}

	for svgimage := range stage.SvgImages {
		svgimage.Unstage(stage)
	}

	// end of insertion point for array nil
}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type Gongstruct interface{}

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
	GongGetUUID(stage *Stage) string
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
	case map[*DownloadableFile]any:
		return any(&stage.DownloadableFiles).(*Type)
	case map[*JpgImage]any:
		return any(&stage.JpgImages).(*Type)
	case map[*Page]any:
		return any(&stage.Pages).(*Type)
	case map[*PngImage]any:
		return any(&stage.PngImages).(*Type)
	case map[*Section]any:
		return any(&stage.Sections).(*Type)
	case map[*SvgImage]any:
		return any(&stage.SvgImages).(*Type)
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
	case *DownloadableFile:
		return any(stage.DownloadableFiles_mapString).(map[string]Type)
	case *JpgImage:
		return any(stage.JpgImages_mapString).(map[string]Type)
	case *Page:
		return any(stage.Pages_mapString).(map[string]Type)
	case *PngImage:
		return any(stage.PngImages_mapString).(map[string]Type)
	case *Section:
		return any(stage.Sections_mapString).(map[string]Type)
	case *SvgImage:
		return any(stage.SvgImages_mapString).(map[string]Type)
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
	case DownloadableFile:
		return any(&stage.DownloadableFiles).(*map[*Type]struct{})
	case JpgImage:
		return any(&stage.JpgImages).(*map[*Type]struct{})
	case Page:
		return any(&stage.Pages).(*map[*Type]struct{})
	case PngImage:
		return any(&stage.PngImages).(*map[*Type]struct{})
	case Section:
		return any(&stage.Sections).(*map[*Type]struct{})
	case SvgImage:
		return any(&stage.SvgImages).(*map[*Type]struct{})
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
	case *DownloadableFile:
		return any(&stage.DownloadableFiles).(*map[Type]struct{})
	case *JpgImage:
		return any(&stage.JpgImages).(*map[Type]struct{})
	case *Page:
		return any(&stage.Pages).(*map[Type]struct{})
	case *PngImage:
		return any(&stage.PngImages).(*map[Type]struct{})
	case *Section:
		return any(&stage.Sections).(*map[Type]struct{})
	case *SvgImage:
		return any(&stage.SvgImages).(*map[Type]struct{})
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
	case DownloadableFile:
		return any(&stage.DownloadableFiles_mapString).(*map[string]*Type)
	case JpgImage:
		return any(&stage.JpgImages_mapString).(*map[string]*Type)
	case Page:
		return any(&stage.Pages_mapString).(*map[string]*Type)
	case PngImage:
		return any(&stage.PngImages_mapString).(*map[string]*Type)
	case Section:
		return any(&stage.Sections_mapString).(*map[string]*Type)
	case SvgImage:
		return any(&stage.SvgImages_mapString).(*map[string]*Type)
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
			// field is initialized with an instance of Section with the name of the field
			Sections: []*Section{{Name: "Sections"}},
			// field is initialized with an instance of Page with the name of the field
			Pages: []*Page{{Name: "Pages"}},
			// field is initialized with an instance of Chapter with the name of the field
			SubChapters: []*Chapter{{Name: "SubChapters"}},
		}).(*Type)
	case Content:
		return any(&Content{
			// Initialisation of associations
			// field is initialized with an instance of Chapter with the name of the field
			Chapters: []*Chapter{{Name: "Chapters"}},
		}).(*Type)
	case DownloadableFile:
		return any(&DownloadableFile{
			// Initialisation of associations
		}).(*Type)
	case JpgImage:
		return any(&JpgImage{
			// Initialisation of associations
		}).(*Type)
	case Page:
		return any(&Page{
			// Initialisation of associations
			// field is initialized with an instance of Section with the name of the field
			Sections: []*Section{{Name: "Sections"}},
		}).(*Type)
	case PngImage:
		return any(&PngImage{
			// Initialisation of associations
		}).(*Type)
	case Section:
		return any(&Section{
			// Initialisation of associations
			// field is initialized with an instance of SvgImage with the name of the field
			SvgImage: &SvgImage{Name: "SvgImage"},
			// field is initialized with an instance of PngImage with the name of the field
			PngImage: &PngImage{Name: "PngImage"},
			// field is initialized with an instance of JpgImage with the name of the field
			JpgImage: &JpgImage{Name: "JpgImage"},
			// field is initialized with an instance of DownloadableFile with the name of the field
			DownloadableFile: &DownloadableFile{Name: "DownloadableFile"},
		}).(*Type)
	case SvgImage:
		return any(&SvgImage{
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
	// reverse maps of direct associations of DownloadableFile
	case DownloadableFile:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of JpgImage
	case JpgImage:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Page
	case Page:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of PngImage
	case PngImage:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Section
	case Section:
		switch fieldname {
		// insertion point for per direct association field
		case "SvgImage":
			res := make(map[*SvgImage][]*Section)
			for section := range stage.Sections {
				if section.SvgImage != nil {
					svgimage_ := section.SvgImage
					var sections []*Section
					_, ok := res[svgimage_]
					if ok {
						sections = res[svgimage_]
					} else {
						sections = make([]*Section, 0)
					}
					sections = append(sections, section)
					res[svgimage_] = sections
				}
			}
			return any(res).(map[*End][]*Start)
		case "PngImage":
			res := make(map[*PngImage][]*Section)
			for section := range stage.Sections {
				if section.PngImage != nil {
					pngimage_ := section.PngImage
					var sections []*Section
					_, ok := res[pngimage_]
					if ok {
						sections = res[pngimage_]
					} else {
						sections = make([]*Section, 0)
					}
					sections = append(sections, section)
					res[pngimage_] = sections
				}
			}
			return any(res).(map[*End][]*Start)
		case "JpgImage":
			res := make(map[*JpgImage][]*Section)
			for section := range stage.Sections {
				if section.JpgImage != nil {
					jpgimage_ := section.JpgImage
					var sections []*Section
					_, ok := res[jpgimage_]
					if ok {
						sections = res[jpgimage_]
					} else {
						sections = make([]*Section, 0)
					}
					sections = append(sections, section)
					res[jpgimage_] = sections
				}
			}
			return any(res).(map[*End][]*Start)
		case "DownloadableFile":
			res := make(map[*DownloadableFile][]*Section)
			for section := range stage.Sections {
				if section.DownloadableFile != nil {
					downloadablefile_ := section.DownloadableFile
					var sections []*Section
					_, ok := res[downloadablefile_]
					if ok {
						sections = res[downloadablefile_]
					} else {
						sections = make([]*Section, 0)
					}
					sections = append(sections, section)
					res[downloadablefile_] = sections
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of SvgImage
	case SvgImage:
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
		case "Sections":
			res := make(map[*Section][]*Chapter)
			for chapter := range stage.Chapters {
				for _, section_ := range chapter.Sections {
					res[section_] = append(res[section_], chapter)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Pages":
			res := make(map[*Page][]*Chapter)
			for chapter := range stage.Chapters {
				for _, page_ := range chapter.Pages {
					res[page_] = append(res[page_], chapter)
				}
			}
			return any(res).(map[*End][]*Start)
		case "SubChapters":
			res := make(map[*Chapter][]*Chapter)
			for chapter := range stage.Chapters {
				for _, chapter_ := range chapter.SubChapters {
					res[chapter_] = append(res[chapter_], chapter)
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
	// reverse maps of direct associations of DownloadableFile
	case DownloadableFile:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of JpgImage
	case JpgImage:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Page
	case Page:
		switch fieldname {
		// insertion point for per direct association field
		case "Sections":
			res := make(map[*Section][]*Page)
			for page := range stage.Pages {
				for _, section_ := range page.Sections {
					res[section_] = append(res[section_], page)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of PngImage
	case PngImage:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Section
	case Section:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SvgImage
	case SvgImage:
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
	case *DownloadableFile:
		res = "DownloadableFile"
	case *JpgImage:
		res = "JpgImage"
	case *Page:
		res = "Page"
	case *PngImage:
		res = "PngImage"
	case *Section:
		res = "Section"
	case *SvgImage:
		res = "SvgImage"
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
		rf.GongstructName = "Chapter"
		rf.Fieldname = "SubChapters"
		res = append(res, rf)
		rf.GongstructName = "Content"
		rf.Fieldname = "Chapters"
		res = append(res, rf)
	case *Content:
		var rf ReverseField
		_ = rf
	case *DownloadableFile:
		var rf ReverseField
		_ = rf
	case *JpgImage:
		var rf ReverseField
		_ = rf
	case *Page:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Chapter"
		rf.Fieldname = "Pages"
		res = append(res, rf)
	case *PngImage:
		var rf ReverseField
		_ = rf
	case *Section:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Chapter"
		rf.Fieldname = "Sections"
		res = append(res, rf)
		rf.GongstructName = "Page"
		rf.Fieldname = "Sections"
		res = append(res, rf)
	case *SvgImage:
		var rf ReverseField
		_ = rf
	}
	return
}

// insertion point for get fields header method
func (chapter *Chapter) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "MardownContent",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Sections",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Section",
		},
		{
			Name:                 "Pages",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Page",
		},
		{
			Name:                 "SubChapters",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Chapter",
		},
	}
	return
}

func (content *Content) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "MardownContent",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "ContentPath",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "OutputPath",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StaticPath",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "LogoSVGFile",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsBespokeLogoFileName",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "BespokeLogoFileName",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsBespokePageTileLogoFileName",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "BespokePageTileLogoFileName",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Target",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "Target",
		},
		{
			Name:                 "Chapters",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Chapter",
		},
		{
			Name:               "VersionInfo",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (downloadablefile *DownloadableFile) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Base64Content",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (jpgimage *JpgImage) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Base64Content",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (page *Page) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "MardownContent",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Sections",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Section",
		},
	}
	return
}

func (pngimage *PngImage) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Base64Content",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (section *Section) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "MardownContent",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsImage",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "SvgImage",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "SvgImage",
		},
		{
			Name:                 "PngImage",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "PngImage",
		},
		{
			Name:                 "JpgImage",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "JpgImage",
		},
		{
			Name:               "IsDownloadableFile",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "DownloadableFile",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "DownloadableFile",
		},
	}
	return
}

func (svgimage *SvgImage) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Content",
			GongFieldValueType: GongFieldValueTypeString,
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
	GongFieldValueTypeIntDuration     GongFieldValueType = "GongFieldValueTypeIntDuration"
	GongFieldValueTypeFloat           GongFieldValueType = "GongFieldValueTypeFloat"
	GongFieldValueTypeBool            GongFieldValueType = "GongFieldValueTypeBool"
	GongFieldValueTypeString          GongFieldValueType = "GongFieldValueTypeString"
	GongFieldValueTypeDate            GongFieldValueType = "GongFieldValueTypeDate"
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
	case "Sections":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range chapter.Sections {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Pages":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range chapter.Pages {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "SubChapters":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range chapter.SubChapters {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
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
	case "StaticPath":
		res.valueString = content.StaticPath
	case "LogoSVGFile":
		res.valueString = content.LogoSVGFile
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
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "VersionInfo":
		res.valueString = content.VersionInfo
	}
	return
}

func (downloadablefile *DownloadableFile) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = downloadablefile.Name
	case "Base64Content":
		res.valueString = downloadablefile.Base64Content
	}
	return
}

func (jpgimage *JpgImage) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = jpgimage.Name
	case "Base64Content":
		res.valueString = jpgimage.Base64Content
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
	case "Sections":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range page.Sections {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	}
	return
}

func (pngimage *PngImage) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = pngimage.Name
	case "Base64Content":
		res.valueString = pngimage.Base64Content
	}
	return
}

func (section *Section) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = section.Name
	case "MardownContent":
		res.valueString = section.MardownContent
	case "IsImage":
		res.valueString = fmt.Sprintf("%t", section.IsImage)
		res.valueBool = section.IsImage
		res.GongFieldValueType = GongFieldValueTypeBool
	case "SvgImage":
		res.GongFieldValueType = GongFieldValueTypePointer
		if section.SvgImage != nil {
			res.valueString = section.SvgImage.Name
			res.ids = section.SvgImage.GongGetUUID(stage)
		}
	case "PngImage":
		res.GongFieldValueType = GongFieldValueTypePointer
		if section.PngImage != nil {
			res.valueString = section.PngImage.Name
			res.ids = section.PngImage.GongGetUUID(stage)
		}
	case "JpgImage":
		res.GongFieldValueType = GongFieldValueTypePointer
		if section.JpgImage != nil {
			res.valueString = section.JpgImage.Name
			res.ids = section.JpgImage.GongGetUUID(stage)
		}
	case "IsDownloadableFile":
		res.valueString = fmt.Sprintf("%t", section.IsDownloadableFile)
		res.valueBool = section.IsDownloadableFile
		res.GongFieldValueType = GongFieldValueTypeBool
	case "DownloadableFile":
		res.GongFieldValueType = GongFieldValueTypePointer
		if section.DownloadableFile != nil {
			res.valueString = section.DownloadableFile.Name
			res.ids = section.DownloadableFile.GongGetUUID(stage)
		}
	}
	return
}

func (svgimage *SvgImage) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = svgimage.Name
	case "Content":
		res.valueString = svgimage.Content
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
	case "Sections":
		chapter.Sections = make([]*Section, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Sections {
					if stage.Section_stagedOrder[__instance__] == uint(id) {
						chapter.Sections = append(chapter.Sections, __instance__)
						break
					}
				}
			}
		}
	case "Pages":
		chapter.Pages = make([]*Page, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Pages {
					if stage.Page_stagedOrder[__instance__] == uint(id) {
						chapter.Pages = append(chapter.Pages, __instance__)
						break
					}
				}
			}
		}
	case "SubChapters":
		chapter.SubChapters = make([]*Chapter, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Chapters {
					if stage.Chapter_stagedOrder[__instance__] == uint(id) {
						chapter.SubChapters = append(chapter.SubChapters, __instance__)
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
	case "StaticPath":
		content.StaticPath = value.GetValueString()
	case "LogoSVGFile":
		content.LogoSVGFile = value.GetValueString()
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
					if stage.Chapter_stagedOrder[__instance__] == uint(id) {
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

func (downloadablefile *DownloadableFile) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		downloadablefile.Name = value.GetValueString()
	case "Base64Content":
		downloadablefile.Base64Content = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (jpgimage *JpgImage) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		jpgimage.Name = value.GetValueString()
	case "Base64Content":
		jpgimage.Base64Content = value.GetValueString()
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
	case "Sections":
		page.Sections = make([]*Section, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Sections {
					if stage.Section_stagedOrder[__instance__] == uint(id) {
						page.Sections = append(page.Sections, __instance__)
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

func (pngimage *PngImage) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		pngimage.Name = value.GetValueString()
	case "Base64Content":
		pngimage.Base64Content = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (section *Section) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		section.Name = value.GetValueString()
	case "MardownContent":
		section.MardownContent = value.GetValueString()
	case "IsImage":
		section.IsImage = value.GetValueBool()
	case "SvgImage":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			section.SvgImage = nil
			for __instance__ := range stage.SvgImages {
				if stage.SvgImage_stagedOrder[__instance__] == uint(id) {
					section.SvgImage = __instance__
					break
				}
			}
		}
	case "PngImage":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			section.PngImage = nil
			for __instance__ := range stage.PngImages {
				if stage.PngImage_stagedOrder[__instance__] == uint(id) {
					section.PngImage = __instance__
					break
				}
			}
		}
	case "JpgImage":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			section.JpgImage = nil
			for __instance__ := range stage.JpgImages {
				if stage.JpgImage_stagedOrder[__instance__] == uint(id) {
					section.JpgImage = __instance__
					break
				}
			}
		}
	case "IsDownloadableFile":
		section.IsDownloadableFile = value.GetValueBool()
	case "DownloadableFile":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			section.DownloadableFile = nil
			for __instance__ := range stage.DownloadableFiles {
				if stage.DownloadableFile_stagedOrder[__instance__] == uint(id) {
					section.DownloadableFile = __instance__
					break
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (svgimage *SvgImage) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		svgimage.Name = value.GetValueString()
	case "Content":
		svgimage.Content = value.GetValueString()
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

func (downloadablefile *DownloadableFile) GongGetGongstructName() string {
	return "DownloadableFile"
}

func (jpgimage *JpgImage) GongGetGongstructName() string {
	return "JpgImage"
}

func (page *Page) GongGetGongstructName() string {
	return "Page"
}

func (pngimage *PngImage) GongGetGongstructName() string {
	return "PngImage"
}

func (section *Section) GongGetGongstructName() string {
	return "Section"
}

func (svgimage *SvgImage) GongGetGongstructName() string {
	return "SvgImage"
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

	stage.DownloadableFiles_mapString = make(map[string]*DownloadableFile)
	for downloadablefile := range stage.DownloadableFiles {
		stage.DownloadableFiles_mapString[downloadablefile.Name] = downloadablefile
	}

	stage.JpgImages_mapString = make(map[string]*JpgImage)
	for jpgimage := range stage.JpgImages {
		stage.JpgImages_mapString[jpgimage.Name] = jpgimage
	}

	stage.Pages_mapString = make(map[string]*Page)
	for page := range stage.Pages {
		stage.Pages_mapString[page.Name] = page
	}

	stage.PngImages_mapString = make(map[string]*PngImage)
	for pngimage := range stage.PngImages {
		stage.PngImages_mapString[pngimage.Name] = pngimage
	}

	stage.Sections_mapString = make(map[string]*Section)
	for section := range stage.Sections {
		stage.Sections_mapString[section.Name] = section
	}

	stage.SvgImages_mapString = make(map[string]*SvgImage)
	for svgimage := range stage.SvgImages {
		stage.SvgImages_mapString[svgimage.Name] = svgimage
	}

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
