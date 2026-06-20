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

	books_go "github.com/fullstack-lang/gong/dsm/xsd/tests/books/go"
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
	BookTypes                map[*BookType]struct{}
	BookTypes_instance       map[*BookType]*BookType
	BookTypes_mapString      map[string]*BookType
	BookTypeOrder            uint
	BookType_stagedOrder     map[*BookType]uint
	BookType_orderStaged     map[uint]*BookType
	BookTypes_reference      map[*BookType]*BookType
	BookTypes_referenceOrder map[*BookType]uint

	// insertion point for slice of pointers maps
	BookType_Credit_reverseMap map[*Credit]*BookType

	OnAfterBookTypeCreateCallback OnAfterCreateInterface[BookType]
	OnAfterBookTypeUpdateCallback OnAfterUpdateInterface[BookType]
	OnAfterBookTypeDeleteCallback OnAfterDeleteInterface[BookType]
	OnAfterBookTypeReadCallback   OnAfterReadInterface[BookType]

	Bookss                map[*Books]struct{}
	Bookss_instance       map[*Books]*Books
	Bookss_mapString      map[string]*Books
	BooksOrder            uint
	Books_stagedOrder     map[*Books]uint
	Books_orderStaged     map[uint]*Books
	Bookss_reference      map[*Books]*Books
	Bookss_referenceOrder map[*Books]uint

	// insertion point for slice of pointers maps
	Books_Book_reverseMap map[*BookType]*Books

	OnAfterBooksCreateCallback OnAfterCreateInterface[Books]
	OnAfterBooksUpdateCallback OnAfterUpdateInterface[Books]
	OnAfterBooksDeleteCallback OnAfterDeleteInterface[Books]
	OnAfterBooksReadCallback   OnAfterReadInterface[Books]

	Credits                map[*Credit]struct{}
	Credits_instance       map[*Credit]*Credit
	Credits_mapString      map[string]*Credit
	CreditOrder            uint
	Credit_stagedOrder     map[*Credit]uint
	Credit_orderStaged     map[uint]*Credit
	Credits_reference      map[*Credit]*Credit
	Credits_referenceOrder map[*Credit]uint

	// insertion point for slice of pointers maps
	Credit_Link_reverseMap map[*Link]*Credit

	OnAfterCreditCreateCallback OnAfterCreateInterface[Credit]
	OnAfterCreditUpdateCallback OnAfterUpdateInterface[Credit]
	OnAfterCreditDeleteCallback OnAfterDeleteInterface[Credit]
	OnAfterCreditReadCallback   OnAfterReadInterface[Credit]

	Links                map[*Link]struct{}
	Links_instance       map[*Link]*Link
	Links_mapString      map[string]*Link
	LinkOrder            uint
	Link_stagedOrder     map[*Link]uint
	Link_orderStaged     map[uint]*Link
	Links_reference      map[*Link]*Link
	Links_referenceOrder map[*Link]uint

	// insertion point for slice of pointers maps
	OnAfterLinkCreateCallback OnAfterCreateInterface[Link]
	OnAfterLinkUpdateCallback OnAfterUpdateInterface[Link]
	OnAfterLinkDeleteCallback OnAfterDeleteInterface[Link]
	OnAfterLinkReadCallback   OnAfterReadInterface[Link]

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
	stage.BookTypes_reference = make(map[*BookType]*BookType)
	stage.BookTypes_instance = make(map[*BookType]*BookType)
	stage.BookTypes_referenceOrder = make(map[*BookType]uint)

	stage.Bookss_reference = make(map[*Books]*Books)
	stage.Bookss_instance = make(map[*Books]*Books)
	stage.Bookss_referenceOrder = make(map[*Books]uint)

	stage.Credits_reference = make(map[*Credit]*Credit)
	stage.Credits_instance = make(map[*Credit]*Credit)
	stage.Credits_referenceOrder = make(map[*Credit]uint)

	stage.Links_reference = make(map[*Link]*Link)
	stage.Links_instance = make(map[*Link]*Link)
	stage.Links_referenceOrder = make(map[*Link]uint)

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
	var maxBookTypeOrder uint
	var foundBookType bool
	for _, order := range stage.BookType_stagedOrder {
		if !foundBookType || order > maxBookTypeOrder {
			maxBookTypeOrder = order
			foundBookType = true
		}
	}
	if foundBookType {
		stage.BookTypeOrder = maxBookTypeOrder + 1
	} else {
		stage.BookTypeOrder = 0
	}

	var maxBooksOrder uint
	var foundBooks bool
	for _, order := range stage.Books_stagedOrder {
		if !foundBooks || order > maxBooksOrder {
			maxBooksOrder = order
			foundBooks = true
		}
	}
	if foundBooks {
		stage.BooksOrder = maxBooksOrder + 1
	} else {
		stage.BooksOrder = 0
	}

	var maxCreditOrder uint
	var foundCredit bool
	for _, order := range stage.Credit_stagedOrder {
		if !foundCredit || order > maxCreditOrder {
			maxCreditOrder = order
			foundCredit = true
		}
	}
	if foundCredit {
		stage.CreditOrder = maxCreditOrder + 1
	} else {
		stage.CreditOrder = 0
	}

	var maxLinkOrder uint
	var foundLink bool
	for _, order := range stage.Link_stagedOrder {
		if !foundLink || order > maxLinkOrder {
			maxLinkOrder = order
			foundLink = true
		}
	}
	if foundLink {
		stage.LinkOrder = maxLinkOrder + 1
	} else {
		stage.LinkOrder = 0
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
	case *BookType:
		tmp := GetStructInstancesByOrder(stage.BookTypes, stage.BookType_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *BookType implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Books:
		tmp := GetStructInstancesByOrder(stage.Bookss, stage.Books_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Books implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Credit:
		tmp := GetStructInstancesByOrder(stage.Credits, stage.Credit_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Credit implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Link:
		tmp := GetStructInstancesByOrder(stage.Links, stage.Link_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Link implements.
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
	case "BookType":
		res = GetNamedStructInstances(stage.BookTypes, stage.BookType_stagedOrder)
	case "Books":
		res = GetNamedStructInstances(stage.Bookss, stage.Books_stagedOrder)
	case "Credit":
		res = GetNamedStructInstances(stage.Credits, stage.Credit_stagedOrder)
	case "Link":
		res = GetNamedStructInstances(stage.Links, stage.Link_stagedOrder)
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
	return "github.com/fullstack-lang/gong/dsm/xsd/tests/books/go/models"
}

func (stage *Stage) GetMap_GongStructName_InstancesNb() map[string]int {
	return stage.Map_GongStructName_InstancesNb
}

func (stage *Stage) GetModelsEmbededDir() embed.FS {
	return books_go.GoModelsDir
}

func (stage *Stage) GetDigramsEmbededDir() embed.FS {
	return books_go.GoDiagramsDir
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
	CommitBookType(booktype *BookType)
	CheckoutBookType(booktype *BookType)
	CommitBooks(books *Books)
	CheckoutBooks(books *Books)
	CommitCredit(credit *Credit)
	CheckoutCredit(credit *Credit)
	CommitLink(link *Link)
	CheckoutLink(link *Link)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(name string) (stage *Stage) {
	stage = &Stage{ // insertion point for array initiatialisation
		BookTypes:           make(map[*BookType]struct{}),
		BookTypes_mapString: make(map[string]*BookType),

		Bookss:           make(map[*Books]struct{}),
		Bookss_mapString: make(map[string]*Books),

		Credits:           make(map[*Credit]struct{}),
		Credits_mapString: make(map[string]*Credit),

		Links:           make(map[*Link]struct{}),
		Links_mapString: make(map[string]*Link),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		name: name,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		BookType_stagedOrder: make(map[*BookType]uint),
		BookType_orderStaged: make(map[uint]*BookType),
		BookTypes_reference:  make(map[*BookType]*BookType),

		Books_stagedOrder: make(map[*Books]uint),
		Books_orderStaged: make(map[uint]*Books),
		Bookss_reference:  make(map[*Books]*Books),

		Credit_stagedOrder: make(map[*Credit]uint),
		Credit_orderStaged: make(map[uint]*Credit),
		Credits_reference:  make(map[*Credit]*Credit),

		Link_stagedOrder: make(map[*Link]uint),
		Link_orderStaged: make(map[uint]*Link),
		Links_reference:  make(map[*Link]*Link),

		// end of insertion point
		GongUnmarshallers: map[string]ModelUnmarshaller{ // insertion point for unmarshallers
			"BookType": &BookTypeUnmarshaller{},

			"Books": &BooksUnmarshaller{},

			"Credit": &CreditUnmarshaller{},

			"Link": &LinkUnmarshaller{},

			// end of insertion point
		},

		NamedStructs: []*NamedStruct{ // insertion point for order map initialisations
			{name: "BookType"},
			{name: "Books"},
			{name: "Credit"},
			{name: "Link"},
		}, // end of insertion point

		navigationMode: GongNavigationModeNormal,
	}

	return
}

func GetOrder[Type Gongstruct](stage *Stage, instance *Type) uint {
	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *BookType:
		return stage.BookType_stagedOrder[instance]
	case *Books:
		return stage.Books_stagedOrder[instance]
	case *Credit:
		return stage.Credit_stagedOrder[instance]
	case *Link:
		return stage.Link_stagedOrder[instance]
	default:
		return 0 // should not happen
	}
}

func GongGetInstanceFromOrder[Type PointerToGongstruct](stage *Stage, order uint) (res Type) {
	var t Type
	switch any(t).(type) {
	// insertion point for order map initialisations
	case *BookType:
		return any(stage.BookType_orderStaged[order]).(Type)
	case *Books:
		return any(stage.Books_orderStaged[order]).(Type)
	case *Credit:
		return any(stage.Credit_orderStaged[order]).(Type)
	case *Link:
		return any(stage.Link_orderStaged[order]).(Type)
	default:
		return // should not happen
	}
}

func GetOrderPointerGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) uint {
	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *BookType:
		return stage.BookType_stagedOrder[instance]
	case *Books:
		return stage.Books_stagedOrder[instance]
	case *Credit:
		return stage.Credit_stagedOrder[instance]
	case *Link:
		return stage.Link_stagedOrder[instance]
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
	stage.Map_GongStructName_InstancesNb["BookType"] = len(stage.BookTypes)
	stage.Map_GongStructName_InstancesNb["Books"] = len(stage.Bookss)
	stage.Map_GongStructName_InstancesNb["Credit"] = len(stage.Credits)
	stage.Map_GongStructName_InstancesNb["Link"] = len(stage.Links)
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
// Stage puts booktype to the model stage
func (booktype *BookType) Stage(stage *Stage) *BookType {
	if _, ok := stage.BookTypes[booktype]; !ok {
		stage.BookTypes[booktype] = struct{}{}
		stage.BookType_stagedOrder[booktype] = stage.BookTypeOrder
		stage.BookType_orderStaged[stage.BookTypeOrder] = booktype
		stage.BookTypeOrder++
	}
	stage.BookTypes_mapString[booktype.Name] = booktype

	return booktype
}

// StagePreserveOrder puts booktype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.BookTypeOrder
// - update stage.BookTypeOrder accordingly
func (booktype *BookType) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.BookTypes[booktype]; !ok {
		stage.BookTypes[booktype] = struct{}{}

		if order > stage.BookTypeOrder {
			stage.BookTypeOrder = order
		}
		stage.BookType_stagedOrder[booktype] = order
		stage.BookType_orderStaged[order] = booktype
		stage.BookTypeOrder++
	}
	stage.BookTypes_mapString[booktype.Name] = booktype
}

// Unstage removes booktype off the model stage
func (booktype *BookType) Unstage(stage *Stage) *BookType {
	delete(stage.BookTypes, booktype)
	// issue1150
	// delete(stage.BookType_stagedOrder, booktype)
	delete(stage.BookTypes_mapString, booktype.Name)

	return booktype
}

// UnstageVoid removes booktype off the model stage
func (booktype *BookType) UnstageVoid(stage *Stage) {
	delete(stage.BookTypes, booktype)
	// issue1150
	// delete(stage.BookType_stagedOrder, booktype)
	delete(stage.BookTypes_mapString, booktype.Name)
}

// commit booktype to the back repo (if it is already staged)
func (booktype *BookType) Commit(stage *Stage) *BookType {
	if _, ok := stage.BookTypes[booktype]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitBookType(booktype)
		}
	}
	return booktype
}

func (booktype *BookType) CommitVoid(stage *Stage) {
	booktype.Commit(stage)
}

func (booktype *BookType) StageVoid(stage *Stage) {
	booktype.Stage(stage)
}

// Checkout booktype to the back repo (if it is already staged)
func (booktype *BookType) Checkout(stage *Stage) *BookType {
	if _, ok := stage.BookTypes[booktype]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutBookType(booktype)
		}
	}
	return booktype
}

// for satisfaction of GongStruct interface
func (booktype *BookType) GetName() (res string) {
	return booktype.Name
}

// for satisfaction of GongStruct interface
func (booktype *BookType) SetName(name string) {
	booktype.Name = name
}

// Stage puts books to the model stage
func (books *Books) Stage(stage *Stage) *Books {
	if _, ok := stage.Bookss[books]; !ok {
		stage.Bookss[books] = struct{}{}
		stage.Books_stagedOrder[books] = stage.BooksOrder
		stage.Books_orderStaged[stage.BooksOrder] = books
		stage.BooksOrder++
	}
	stage.Bookss_mapString[books.Name] = books

	return books
}

// StagePreserveOrder puts books to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.BooksOrder
// - update stage.BooksOrder accordingly
func (books *Books) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Bookss[books]; !ok {
		stage.Bookss[books] = struct{}{}

		if order > stage.BooksOrder {
			stage.BooksOrder = order
		}
		stage.Books_stagedOrder[books] = order
		stage.Books_orderStaged[order] = books
		stage.BooksOrder++
	}
	stage.Bookss_mapString[books.Name] = books
}

// Unstage removes books off the model stage
func (books *Books) Unstage(stage *Stage) *Books {
	delete(stage.Bookss, books)
	// issue1150
	// delete(stage.Books_stagedOrder, books)
	delete(stage.Bookss_mapString, books.Name)

	return books
}

// UnstageVoid removes books off the model stage
func (books *Books) UnstageVoid(stage *Stage) {
	delete(stage.Bookss, books)
	// issue1150
	// delete(stage.Books_stagedOrder, books)
	delete(stage.Bookss_mapString, books.Name)
}

// commit books to the back repo (if it is already staged)
func (books *Books) Commit(stage *Stage) *Books {
	if _, ok := stage.Bookss[books]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitBooks(books)
		}
	}
	return books
}

func (books *Books) CommitVoid(stage *Stage) {
	books.Commit(stage)
}

func (books *Books) StageVoid(stage *Stage) {
	books.Stage(stage)
}

// Checkout books to the back repo (if it is already staged)
func (books *Books) Checkout(stage *Stage) *Books {
	if _, ok := stage.Bookss[books]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutBooks(books)
		}
	}
	return books
}

// for satisfaction of GongStruct interface
func (books *Books) GetName() (res string) {
	return books.Name
}

// for satisfaction of GongStruct interface
func (books *Books) SetName(name string) {
	books.Name = name
}

// Stage puts credit to the model stage
func (credit *Credit) Stage(stage *Stage) *Credit {
	if _, ok := stage.Credits[credit]; !ok {
		stage.Credits[credit] = struct{}{}
		stage.Credit_stagedOrder[credit] = stage.CreditOrder
		stage.Credit_orderStaged[stage.CreditOrder] = credit
		stage.CreditOrder++
	}
	stage.Credits_mapString[credit.Name] = credit

	return credit
}

// StagePreserveOrder puts credit to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CreditOrder
// - update stage.CreditOrder accordingly
func (credit *Credit) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Credits[credit]; !ok {
		stage.Credits[credit] = struct{}{}

		if order > stage.CreditOrder {
			stage.CreditOrder = order
		}
		stage.Credit_stagedOrder[credit] = order
		stage.Credit_orderStaged[order] = credit
		stage.CreditOrder++
	}
	stage.Credits_mapString[credit.Name] = credit
}

// Unstage removes credit off the model stage
func (credit *Credit) Unstage(stage *Stage) *Credit {
	delete(stage.Credits, credit)
	// issue1150
	// delete(stage.Credit_stagedOrder, credit)
	delete(stage.Credits_mapString, credit.Name)

	return credit
}

// UnstageVoid removes credit off the model stage
func (credit *Credit) UnstageVoid(stage *Stage) {
	delete(stage.Credits, credit)
	// issue1150
	// delete(stage.Credit_stagedOrder, credit)
	delete(stage.Credits_mapString, credit.Name)
}

// commit credit to the back repo (if it is already staged)
func (credit *Credit) Commit(stage *Stage) *Credit {
	if _, ok := stage.Credits[credit]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitCredit(credit)
		}
	}
	return credit
}

func (credit *Credit) CommitVoid(stage *Stage) {
	credit.Commit(stage)
}

func (credit *Credit) StageVoid(stage *Stage) {
	credit.Stage(stage)
}

// Checkout credit to the back repo (if it is already staged)
func (credit *Credit) Checkout(stage *Stage) *Credit {
	if _, ok := stage.Credits[credit]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutCredit(credit)
		}
	}
	return credit
}

// for satisfaction of GongStruct interface
func (credit *Credit) GetName() (res string) {
	return credit.Name
}

// for satisfaction of GongStruct interface
func (credit *Credit) SetName(name string) {
	credit.Name = name
}

// Stage puts link to the model stage
func (link *Link) Stage(stage *Stage) *Link {
	if _, ok := stage.Links[link]; !ok {
		stage.Links[link] = struct{}{}
		stage.Link_stagedOrder[link] = stage.LinkOrder
		stage.Link_orderStaged[stage.LinkOrder] = link
		stage.LinkOrder++
	}
	stage.Links_mapString[link.Name] = link

	return link
}

// StagePreserveOrder puts link to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.LinkOrder
// - update stage.LinkOrder accordingly
func (link *Link) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Links[link]; !ok {
		stage.Links[link] = struct{}{}

		if order > stage.LinkOrder {
			stage.LinkOrder = order
		}
		stage.Link_stagedOrder[link] = order
		stage.Link_orderStaged[order] = link
		stage.LinkOrder++
	}
	stage.Links_mapString[link.Name] = link
}

// Unstage removes link off the model stage
func (link *Link) Unstage(stage *Stage) *Link {
	delete(stage.Links, link)
	// issue1150
	// delete(stage.Link_stagedOrder, link)
	delete(stage.Links_mapString, link.Name)

	return link
}

// UnstageVoid removes link off the model stage
func (link *Link) UnstageVoid(stage *Stage) {
	delete(stage.Links, link)
	// issue1150
	// delete(stage.Link_stagedOrder, link)
	delete(stage.Links_mapString, link.Name)
}

// commit link to the back repo (if it is already staged)
func (link *Link) Commit(stage *Stage) *Link {
	if _, ok := stage.Links[link]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitLink(link)
		}
	}
	return link
}

func (link *Link) CommitVoid(stage *Stage) {
	link.Commit(stage)
}

func (link *Link) StageVoid(stage *Stage) {
	link.Stage(stage)
}

// Checkout link to the back repo (if it is already staged)
func (link *Link) Checkout(stage *Stage) *Link {
	if _, ok := stage.Links[link]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutLink(link)
		}
	}
	return link
}

// for satisfaction of GongStruct interface
func (link *Link) GetName() (res string) {
	return link.Name
}

// for satisfaction of GongStruct interface
func (link *Link) SetName(name string) {
	link.Name = name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMBookType(BookType *BookType)
	CreateORMBooks(Books *Books)
	CreateORMCredit(Credit *Credit)
	CreateORMLink(Link *Link)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMBookType(BookType *BookType)
	DeleteORMBooks(Books *Books)
	DeleteORMCredit(Credit *Credit)
	DeleteORMLink(Link *Link)
}

func (stage *Stage) Reset() { // insertion point for array reset
	stage.BookTypes = make(map[*BookType]struct{})
	stage.BookTypes_mapString = make(map[string]*BookType)
	stage.BookType_stagedOrder = make(map[*BookType]uint)
	stage.BookTypeOrder = 0

	stage.Bookss = make(map[*Books]struct{})
	stage.Bookss_mapString = make(map[string]*Books)
	stage.Books_stagedOrder = make(map[*Books]uint)
	stage.BooksOrder = 0

	stage.Credits = make(map[*Credit]struct{})
	stage.Credits_mapString = make(map[string]*Credit)
	stage.Credit_stagedOrder = make(map[*Credit]uint)
	stage.CreditOrder = 0

	stage.Links = make(map[*Link]struct{})
	stage.Links_mapString = make(map[string]*Link)
	stage.Link_stagedOrder = make(map[*Link]uint)
	stage.LinkOrder = 0

	if stage.GetProbeIF() != nil {
		stage.GetProbeIF().ResetNotifications()
	}
	if stage.IsInDeltaMode() {
		stage.ComputeReferenceAndOrders()
	}
}

func (stage *Stage) Nil() { // insertion point for array nil
	stage.BookTypes = nil
	stage.BookTypes_mapString = nil

	stage.Bookss = nil
	stage.Bookss_mapString = nil

	stage.Credits = nil
	stage.Credits_mapString = nil

	stage.Links = nil
	stage.Links_mapString = nil

	// end of insertion point for array nil
}

func (stage *Stage) Unstage() { // insertion point for array nil
	for booktype := range stage.BookTypes {
		booktype.Unstage(stage)
	}

	for books := range stage.Bookss {
		books.Unstage(stage)
	}

	for credit := range stage.Credits {
		credit.Unstage(stage)
	}

	for link := range stage.Links {
		link.Unstage(stage)
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
	case map[*BookType]any:
		return any(&stage.BookTypes).(*Type)
	case map[*Books]any:
		return any(&stage.Bookss).(*Type)
	case map[*Credit]any:
		return any(&stage.Credits).(*Type)
	case map[*Link]any:
		return any(&stage.Links).(*Type)
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
	case *BookType:
		return any(stage.BookTypes_mapString).(map[string]Type)
	case *Books:
		return any(stage.Bookss_mapString).(map[string]Type)
	case *Credit:
		return any(stage.Credits_mapString).(map[string]Type)
	case *Link:
		return any(stage.Links_mapString).(map[string]Type)
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
	case BookType:
		return any(&stage.BookTypes).(*map[*Type]struct{})
	case Books:
		return any(&stage.Bookss).(*map[*Type]struct{})
	case Credit:
		return any(&stage.Credits).(*map[*Type]struct{})
	case Link:
		return any(&stage.Links).(*map[*Type]struct{})
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
	case *BookType:
		return any(&stage.BookTypes).(*map[Type]struct{})
	case *Books:
		return any(&stage.Bookss).(*map[Type]struct{})
	case *Credit:
		return any(&stage.Credits).(*map[Type]struct{})
	case *Link:
		return any(&stage.Links).(*map[Type]struct{})
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
	case BookType:
		return any(&stage.BookTypes_mapString).(*map[string]*Type)
	case Books:
		return any(&stage.Bookss_mapString).(*map[string]*Type)
	case Credit:
		return any(&stage.Credits_mapString).(*map[string]*Type)
	case Link:
		return any(&stage.Links_mapString).(*map[string]*Type)
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
	case BookType:
		return any(&BookType{
			// Initialisation of associations
			// field is initialized with an instance of Credit with the name of the field
			Credit: []*Credit{{Name: "Credit"}},
		}).(*Type)
	case Books:
		return any(&Books{
			// Initialisation of associations
			// field is initialized with an instance of BookType with the name of the field
			Book: []*BookType{{Name: "Book"}},
		}).(*Type)
	case Credit:
		return any(&Credit{
			// Initialisation of associations
			// field is initialized with an instance of Link with the name of the field
			Link: []*Link{{Name: "Link"}},
		}).(*Type)
	case Link:
		return any(&Link{
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
	// reverse maps of direct associations of BookType
	case BookType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Books
	case Books:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Credit
	case Credit:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Link
	case Link:
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
	// reverse maps of direct associations of BookType
	case BookType:
		switch fieldname {
		// insertion point for per direct association field
		case "Credit":
			res := make(map[*Credit][]*BookType)
			for booktype := range stage.BookTypes {
				for _, credit_ := range booktype.Credit {
					res[credit_] = append(res[credit_], booktype)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Books
	case Books:
		switch fieldname {
		// insertion point for per direct association field
		case "Book":
			res := make(map[*BookType][]*Books)
			for books := range stage.Bookss {
				for _, booktype_ := range books.Book {
					res[booktype_] = append(res[booktype_], books)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Credit
	case Credit:
		switch fieldname {
		// insertion point for per direct association field
		case "Link":
			res := make(map[*Link][]*Credit)
			for credit := range stage.Credits {
				for _, link_ := range credit.Link {
					res[link_] = append(res[link_], credit)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Link
	case Link:
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
	case *BookType:
		res = "BookType"
	case *Books:
		res = "Books"
	case *Credit:
		res = "Credit"
	case *Link:
		res = "Link"
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
	case *BookType:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Books"
		rf.Fieldname = "Book"
		res = append(res, rf)
	case *Books:
		var rf ReverseField
		_ = rf
	case *Credit:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "BookType"
		rf.Fieldname = "Credit"
		res = append(res, rf)
	case *Link:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Credit"
		rf.Fieldname = "Link"
		res = append(res, rf)
	}
	return
}

// insertion point for get fields header method
func (booktype *BookType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Edition",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Isbn",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Bestseller",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "Title",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Author",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Year",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "Format",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Credit",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Credit",
		},
	}
	return
}

func (books *Books) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Book",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "BookType",
		},
	}
	return
}

func (credit *Credit) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Page",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "Credit_type",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Link",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Link",
		},
		{
			Name:               "Credit_words",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Credit_symbol",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (link *Link) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NameXSD",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EnclosedText",
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
func (booktype *BookType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = booktype.Name
	case "Edition":
		res.valueString = booktype.Edition
	case "Isbn":
		res.valueString = booktype.Isbn
	case "Bestseller":
		res.valueString = fmt.Sprintf("%t", booktype.Bestseller)
		res.valueBool = booktype.Bestseller
		res.GongFieldValueType = GongFieldValueTypeBool
	case "Title":
		res.valueString = booktype.Title
	case "Author":
		res.valueString = booktype.Author
	case "Year":
		res.valueString = fmt.Sprintf("%d", booktype.Year)
		res.valueInt = booktype.Year
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Format":
		res.valueString = booktype.Format
	case "Credit":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range booktype.Credit {
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

func (books *Books) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = books.Name
	case "Book":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range books.Book {
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

func (credit *Credit) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = credit.Name
	case "Page":
		res.valueString = fmt.Sprintf("%d", credit.Page)
		res.valueInt = credit.Page
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Credit_type":
		res.valueString = credit.Credit_type
	case "Link":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range credit.Link {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Credit_words":
		res.valueString = credit.Credit_words
	case "Credit_symbol":
		res.valueString = credit.Credit_symbol
	}
	return
}

func (link *Link) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = link.Name
	case "NameXSD":
		res.valueString = link.NameXSD
	case "EnclosedText":
		res.valueString = link.EnclosedText
	}
	return
}

func GetFieldStringValueFromPointer(instance GongstructIF, fieldName string, stage *Stage) (res GongFieldValue) {
	res = instance.GongGetFieldValue(fieldName, stage)
	return
}

// insertion point for generic set gongstruct field value
func (booktype *BookType) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		booktype.Name = value.GetValueString()
	case "Edition":
		booktype.Edition = value.GetValueString()
	case "Isbn":
		booktype.Isbn = value.GetValueString()
	case "Bestseller":
		booktype.Bestseller = value.GetValueBool()
	case "Title":
		booktype.Title = value.GetValueString()
	case "Author":
		booktype.Author = value.GetValueString()
	case "Year":
		booktype.Year = int(value.GetValueInt())
	case "Format":
		booktype.Format = value.GetValueString()
	case "Credit":
		booktype.Credit = make([]*Credit, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Credits {
					if stage.Credit_stagedOrder[__instance__] == uint(id) {
						booktype.Credit = append(booktype.Credit, __instance__)
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

func (books *Books) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		books.Name = value.GetValueString()
	case "Book":
		books.Book = make([]*BookType, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.BookTypes {
					if stage.BookType_stagedOrder[__instance__] == uint(id) {
						books.Book = append(books.Book, __instance__)
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

func (credit *Credit) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		credit.Name = value.GetValueString()
	case "Page":
		credit.Page = int(value.GetValueInt())
	case "Credit_type":
		credit.Credit_type = value.GetValueString()
	case "Link":
		credit.Link = make([]*Link, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Links {
					if stage.Link_stagedOrder[__instance__] == uint(id) {
						credit.Link = append(credit.Link, __instance__)
						break
					}
				}
			}
		}
	case "Credit_words":
		credit.Credit_words = value.GetValueString()
	case "Credit_symbol":
		credit.Credit_symbol = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (link *Link) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		link.Name = value.GetValueString()
	case "NameXSD":
		link.NameXSD = value.GetValueString()
	case "EnclosedText":
		link.EnclosedText = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func SetFieldStringValueFromPointer(instance GongstructIF, fieldName string, value GongFieldValue, stage *Stage) error {
	return instance.GongSetFieldValue(fieldName, value, stage)
}

// insertion point for generic get gongstruct name
func (booktype *BookType) GongGetGongstructName() string {
	return "BookType"
}

func (books *Books) GongGetGongstructName() string {
	return "Books"
}

func (credit *Credit) GongGetGongstructName() string {
	return "Credit"
}

func (link *Link) GongGetGongstructName() string {
	return "Link"
}

func GetGongstructNameFromPointer(instance GongstructIF) (res string) {
	res = instance.GongGetGongstructName()
	return
}

func (stage *Stage) ResetMapStrings() {
	// insertion point for generic get gongstruct name
	stage.BookTypes_mapString = make(map[string]*BookType)
	for booktype := range stage.BookTypes {
		stage.BookTypes_mapString[booktype.Name] = booktype
	}

	stage.Bookss_mapString = make(map[string]*Books)
	for books := range stage.Bookss {
		stage.Bookss_mapString[books.Name] = books
	}

	stage.Credits_mapString = make(map[string]*Credit)
	for credit := range stage.Credits {
		stage.Credits_mapString[credit.Name] = credit
	}

	stage.Links_mapString = make(map[string]*Link)
	for link := range stage.Links {
		stage.Links_mapString[link.Name] = link
	}

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
