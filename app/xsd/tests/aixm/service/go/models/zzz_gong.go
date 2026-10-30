// generated code - do not edit
package models

import (
	"cmp"
	"errors"
	"fmt"
	"log"
	"math"
	"slices"
	"sort"
	"strings"
	"time"

	gong_runtime "github.com/fullstack-lang/gong/pkg/runtime"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models/aixmfeatures"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models/aixmdatatypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models/aixmabstractgmlobjecttypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models/gml"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models/xlink"
)

// can be used for
//
//	days := __gong__abs(int(int(inferedInstance.ComputedDuration.Hours()) / 24))
func __gong__abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

var (
	_ = __gong__abs
	_ = strings.Clone("")
)

type GongMarshallingMode = gong_runtime.GongMarshallingMode

const (
	// the whole stage is generated at each marshall. This is the default
	GongMarshallingNormal = gong_runtime.GongMarshallingNormal

	// only the last commit is append to the marshall file
	GongMarshallingAppendCommit = gong_runtime.GongMarshallingAppendCommit
)

type gongStageNavigationMode = gong_runtime.StageNavigationMode

const (
	GongNavigationModeNormal     = gong_runtime.GongNavigationModeNormal
	GongNavigationModeNavigating = gong_runtime.GongNavigationModeNavigating
)

const (
	GongProbeTreeSidebarSuffix           = gong_runtime.GongProbeTreeSidebarSuffix
	GongProbeNavigationTreeSidebarSuffix = gong_runtime.GongProbeNavigationTreeSidebarSuffix
	GongProbeTableSuffix                 = gong_runtime.GongProbeTableSuffix
	GongProbeNotificationTableSuffix     = gong_runtime.GongProbeNotificationTableSuffix
	GongProbeFormSuffix                  = gong_runtime.GongProbeFormSuffix
	GongProbeSplitSuffix                 = gong_runtime.GongProbeSplitSuffix
	GongProbeLoadSuffix                  = gong_runtime.GongProbeLoadSuffix
)

func (stage *Stage) GetProbeTreeSidebarStageName() string {
	return stage.GetType() + ":" + stage.GetName() + GongProbeTreeSidebarSuffix
}

func (stage *Stage) GetProbeNavigationTreeSidebarStageName() string {
	return stage.GetType() + ":" + stage.GetName() + GongProbeNavigationTreeSidebarSuffix
}

func (stage *Stage) GetProbeFormStageName() string {
	return stage.GetType() + ":" + stage.GetName() + GongProbeFormSuffix
}

func (stage *Stage) GetProbeTableStageName() string {
	return stage.GetType() + ":" + stage.GetName() + GongProbeTableSuffix
}

func (stage *Stage) GetProbeNotificationTableStageName() string {
	return stage.GetType() + ":" + stage.GetName() + GongProbeNotificationTableSuffix
}

func (stage *Stage) GetProbeSplitStageName() string {
	return stage.GetType() + ":" + stage.GetName() + GongProbeSplitSuffix
}

func (stage *Stage) GetProbeLoadStageName() string {
	return stage.GetType() + ":" + stage.GetName() + GongProbeLoadSuffix
}

// errUnkownEnum is returns when a value cannot match enum values
var (
	errUnkownEnum = errors.New("unkown enum")
	_             = errUnkownEnum
)

// needed to avoid when fmt package is not needed by generated code
var _ = fmt.Sprintf

// idem for math package when not need by generated code
var _ = math.E

// swagger:ignore
type __void any

// needed for creating set of instances in the stage
var (
	__member __void
	_        = __member
)

// MetaPackageImport represents a package import needed by a meta/diagram file
type MetaPackageImport struct {
	Alias string
	Path  string
}

// Stage enables storage of staged instances
type Stage struct {
	gong_runtime.StageCore

	// insertion point for definition of arrays registering instances
	AIXMBasicMessages                map[*AIXMBasicMessage]struct{}
	AIXMBasicMessages_instance       map[*AIXMBasicMessage]*AIXMBasicMessage
	AIXMBasicMessages_mapString      map[string]*AIXMBasicMessage
	AIXMBasicMessageOrder            uint
	AIXMBasicMessage_stagedOrder     map[*AIXMBasicMessage]uint
	AIXMBasicMessage_orderStaged     map[uint]*AIXMBasicMessage
	AIXMBasicMessages_reference      map[*AIXMBasicMessage]*AIXMBasicMessage
	AIXMBasicMessages_referenceOrder map[*AIXMBasicMessage]uint

	// insertion point for slice of pointers maps
	AIXMBasicMessage_HasMember_reverseMap map[*BasicMessageMemberAIXMPropertyType]*AIXMBasicMessage

	OnAfterAIXMBasicMessageCreateCallback GongOnAfterCreateInterface[AIXMBasicMessage]
	OnAfterAIXMBasicMessageUpdateCallback GongOnAfterUpdateInterface[AIXMBasicMessage]
	OnAfterAIXMBasicMessageDeleteCallback GongOnAfterDeleteInterface[AIXMBasicMessage]

	AIXMBasicMessageTypes                map[*AIXMBasicMessageType]struct{}
	AIXMBasicMessageTypes_instance       map[*AIXMBasicMessageType]*AIXMBasicMessageType
	AIXMBasicMessageTypes_mapString      map[string]*AIXMBasicMessageType
	AIXMBasicMessageTypeOrder            uint
	AIXMBasicMessageType_stagedOrder     map[*AIXMBasicMessageType]uint
	AIXMBasicMessageType_orderStaged     map[uint]*AIXMBasicMessageType
	AIXMBasicMessageTypes_reference      map[*AIXMBasicMessageType]*AIXMBasicMessageType
	AIXMBasicMessageTypes_referenceOrder map[*AIXMBasicMessageType]uint

	// insertion point for slice of pointers maps
	AIXMBasicMessageType_HasMember_reverseMap map[*BasicMessageMemberAIXMPropertyType]*AIXMBasicMessageType

	OnAfterAIXMBasicMessageTypeCreateCallback GongOnAfterCreateInterface[AIXMBasicMessageType]
	OnAfterAIXMBasicMessageTypeUpdateCallback GongOnAfterUpdateInterface[AIXMBasicMessageType]
	OnAfterAIXMBasicMessageTypeDeleteCallback GongOnAfterDeleteInterface[AIXMBasicMessageType]

	BasicMessageMemberAIXMPropertyTypes                map[*BasicMessageMemberAIXMPropertyType]struct{}
	BasicMessageMemberAIXMPropertyTypes_instance       map[*BasicMessageMemberAIXMPropertyType]*BasicMessageMemberAIXMPropertyType
	BasicMessageMemberAIXMPropertyTypes_mapString      map[string]*BasicMessageMemberAIXMPropertyType
	BasicMessageMemberAIXMPropertyTypeOrder            uint
	BasicMessageMemberAIXMPropertyType_stagedOrder     map[*BasicMessageMemberAIXMPropertyType]uint
	BasicMessageMemberAIXMPropertyType_orderStaged     map[uint]*BasicMessageMemberAIXMPropertyType
	BasicMessageMemberAIXMPropertyTypes_reference      map[*BasicMessageMemberAIXMPropertyType]*BasicMessageMemberAIXMPropertyType
	BasicMessageMemberAIXMPropertyTypes_referenceOrder map[*BasicMessageMemberAIXMPropertyType]uint

	// insertion point for slice of pointers maps
	OnAfterBasicMessageMemberAIXMPropertyTypeCreateCallback GongOnAfterCreateInterface[BasicMessageMemberAIXMPropertyType]
	OnAfterBasicMessageMemberAIXMPropertyTypeUpdateCallback GongOnAfterUpdateInterface[BasicMessageMemberAIXMPropertyType]
	OnAfterBasicMessageMemberAIXMPropertyTypeDeleteCallback GongOnAfterDeleteInterface[BasicMessageMemberAIXMPropertyType]

	BackRepo GongBackRepoInterface

	// if set will be called before each commit to the back repo
	OnInitCommitCallback          GongOnInitCommitInterface
	OnInitCommitFromFrontCallback GongOnInitCommitInterface
	OnInitCommitFromBackCallback  GongOnInitCommitInterface

	// store the number of instance per gongstruct
	Map_GongStructName_InstancesNb map[string]int

	// store meta package import
	MetaPackageImportPath  string
	MetaPackageImportAlias string
	MetaPackageImports     []*MetaPackageImport

	// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
	// map to enable docLink renaming when an identifier is renamed
	Map_DocLink_Renaming map[string]GONG__Identifier
	// the to be removed stops here

	// store the stage order of each instance in order to
	// preserve this order when serializing them
	// insertion point for order fields declaration
	// end of insertion point

	// GongUnmarshallers is the registry of all model unmarshallers
	GongUnmarshallers map[string]GongModelUnmarshaller

	// probeIF is the interface to the probe that allows log
	// commit event to the probe
	probeIF GongProbeIF
	AixmfeaturesStage *aixmfeatures.Stage
	AixmdatatypesStage *aixmdatatypes.Stage
	AixmabstractgmlobjecttypesStage *aixmabstractgmlobjecttypes.Stage
	GmlStage *gml.Stage
	XlinkStage *xlink.Stage
}

type GongStage = Stage

// RegisterBeforeCommit adds a hook that runs before the commit happens
func (s *Stage) RegisterBeforeCommit(hook func(stage *Stage)) {
	s.StageCore.RegisterBeforeCommitHook(func() { hook(s) })
}

// RegisterAfterCommit adds a hook that runs after the commit succeeds
func (s *Stage) RegisterAfterCommit(hook func(stage *Stage)) {
	s.StageCore.RegisterAfterCommitHook(func() { hook(s) })
}

// ApplyBackwardCommit applies the commit before the current one
func (stage *Stage) ApplyBackwardCommit() error {
	return stage.StageCore.ApplyBackwardCommit(stage.ParseAstString, stage.ComputeReferenceAndOrders)
}

// ApplyForwardCommit applies the commit after the current one
func (stage *Stage) ApplyForwardCommit() error {
	return stage.StageCore.ApplyForwardCommit(stage.ParseAstString, stage.ComputeReferenceAndOrders)
}

// ResetHard removes the more recent
// commitsBehind forward/backward Commits from the
// stage
func (stage *Stage) ResetHard() {
	stage.StageCore.ResetHardCore()

	stage.ComputeInstancesNb()
	if stage.OnInitCommitCallback != nil {
		stage.OnInitCommitCallback.BeforeCommit(stage)
	}
	if stage.OnInitCommitFromBackCallback != nil {
		stage.OnInitCommitFromBackCallback.BeforeCommit(stage)
	}

	// 1. Run all Before Commit hooks
	stage.StageCore.RunBeforeCommitHooks()

	// 2. Run all After Commit hooks
	stage.StageCore.RunAfterCommitHooks()
}

// Squash removes all commits and marshals the stage as a single commit
func (stage *Stage) Squash() {
	stage.StageCore.SquashCore()

	// insertion point for clear references
	__gong__clearReferences(&stage.AIXMBasicMessages_reference, &stage.AIXMBasicMessages_instance, &stage.AIXMBasicMessages_referenceOrder)

	__gong__clearReferences(&stage.AIXMBasicMessageTypes_reference, &stage.AIXMBasicMessageTypes_instance, &stage.AIXMBasicMessageTypes_referenceOrder)

	__gong__clearReferences(&stage.BasicMessageMemberAIXMPropertyTypes_reference, &stage.BasicMessageMemberAIXMPropertyTypes_instance, &stage.BasicMessageMemberAIXMPropertyTypes_referenceOrder)

	stage.ComputeInstancesNb()
	if stage.OnInitCommitCallback != nil {
		stage.OnInitCommitCallback.BeforeCommit(stage)
	}
	if stage.OnInitCommitFromBackCallback != nil {
		stage.OnInitCommitFromBackCallback.BeforeCommit(stage)
	}

	// 1. Run all Before Commit hooks
	stage.StageCore.RunBeforeCommitHooks()

	stage.StageCore.EndSquash()

	// 2. Run all After Commit hooks
	stage.StageCore.RunAfterCommitHooks()
}

// recomputeOrders recomputes the next order for each struct
// this is necessary because the order might have been incremented
// during the commits that have been discarded
// insertion point for max order recomputation
func (stage *Stage) recomputeOrders() {
	// insertion point for max order recomputation
	stage.AIXMBasicMessageOrder = __gong__recomputeOrder(stage.AIXMBasicMessage_stagedOrder)

	stage.AIXMBasicMessageTypeOrder = __gong__recomputeOrder(stage.AIXMBasicMessageType_stagedOrder)

	stage.BasicMessageMemberAIXMPropertyTypeOrder = __gong__recomputeOrder(stage.BasicMessageMemberAIXMPropertyType_stagedOrder)

	// end of insertion point for max order recomputation
}

func (stage *Stage) SetProbeIF(probeIF GongProbeIF) {
	stage.probeIF = probeIF
}

func (stage *Stage) GetProbeIF() GongProbeIF {
	if stage.probeIF == nil {
		return nil
	}

	return stage.probeIF
}

// insertion point for stage ops
func (*AIXMBasicMessage) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.AIXMBasicMessages, stage.AIXMBasicMessage_stagedOrder)
}

func (*AIXMBasicMessage) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.AIXMBasicMessage_orderStaged[order]
}

func (*AIXMBasicMessage) GongGetInstancesMapByName(stage *Stage) any {
	return stage.AIXMBasicMessages_mapString
}

func (*AIXMBasicMessage) GongGetInstancesSet(stage *Stage) any {
	return &stage.AIXMBasicMessages
}

func (*AIXMBasicMessage) GongNewInstance() any {
	return new(AIXMBasicMessage)
}

func (*AIXMBasicMessageType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.AIXMBasicMessageTypes, stage.AIXMBasicMessageType_stagedOrder)
}

func (*AIXMBasicMessageType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.AIXMBasicMessageType_orderStaged[order]
}

func (*AIXMBasicMessageType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.AIXMBasicMessageTypes_mapString
}

func (*AIXMBasicMessageType) GongGetInstancesSet(stage *Stage) any {
	return &stage.AIXMBasicMessageTypes
}

func (*AIXMBasicMessageType) GongNewInstance() any {
	return new(AIXMBasicMessageType)
}

func (*BasicMessageMemberAIXMPropertyType) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.BasicMessageMemberAIXMPropertyTypes, stage.BasicMessageMemberAIXMPropertyType_stagedOrder)
}

func (*BasicMessageMemberAIXMPropertyType) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.BasicMessageMemberAIXMPropertyType_orderStaged[order]
}

func (*BasicMessageMemberAIXMPropertyType) GongGetInstancesMapByName(stage *Stage) any {
	return stage.BasicMessageMemberAIXMPropertyTypes_mapString
}

func (*BasicMessageMemberAIXMPropertyType) GongGetInstancesSet(stage *Stage) any {
	return &stage.BasicMessageMemberAIXMPropertyTypes
}

func (*BasicMessageMemberAIXMPropertyType) GongNewInstance() any {
	return new(BasicMessageMemberAIXMPropertyType)
}


// GetInstancesByOrder is the Stage method returning a slice of generic pointers to gongstructs
// ordered by their order in the stage.
func (stage *Stage) GetInstancesByOrder[T GongstructPtr]() (res []T) {
	if stage == nil {
		return nil
	}
	var t T
	return t.GongGetInstancesByOrder(stage).([]T)
}

func __gong__getStructInstancesByOrder[T GongstructPtr](set map[T]struct{}, order map[T]uint) (res []T) {
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
			log.Fatalf("getStructInstancesByOrder: pointer not found")
		}
		return i_order < j_order
	})

	res = append(res, orderedSet...)

	return
}

func __gong__castSlice[T any, S any](s []S) []T {
	res := make([]T, len(s))
	for i, v := range s {
		res[i] = any(v).(T)
	}
	return res
}

func __gong__stage[T comparable](
	instances map[T]struct{},
	stagedOrder map[T]uint,
	orderStaged map[uint]T,
	order *uint,
	mapString map[string]T,
	instance T,
	name string,
) {
	if _, ok := instances[instance]; !ok {
		instances[instance] = struct{}{}
		stagedOrder[instance] = *order
		orderStaged[*order] = instance
		*order++
	}
	mapString[name] = instance
}

func __gong__stagePreserveOrder[T comparable](
	instances map[T]struct{},
	stagedOrder map[T]uint,
	orderStaged map[uint]T,
	currentOrder *uint,
	mapString map[string]T,
	instance T,
	order uint,
	name string,
) {
	if _, ok := instances[instance]; !ok {
		instances[instance] = struct{}{}
		if order > *currentOrder {
			*currentOrder = order
		}
		stagedOrder[instance] = order
		orderStaged[order] = instance
		*currentOrder++
	}
	mapString[name] = instance
}

func __gong__unstage[T comparable](
	instances map[T]struct{},
	mapString map[string]T,
	instance T,
	name string,
) {
	delete(instances, instance)
	delete(mapString, name)
}

func __gong__recomputeOrder[T comparable](stagedOrder map[T]uint) uint {
	var maxOrder uint
	var found bool
	for _, order := range stagedOrder {
		if !found || order > maxOrder {
			maxOrder = order
			found = true
		}
	}
	if found {
		return maxOrder + 1
	}
	return 0
}

func __gong__rebuildMapString[T interface {
	comparable
	GetName() string
}](staged map[T]struct{}, mapString *map[string]T) {
	*mapString = make(map[string]T, len(staged))
	for instance := range staged {
		(*mapString)[instance.GetName()] = instance
	}
}

func __gong__clearReferences[T comparable](ref *map[T]T, inst *map[T]T, refOrder *map[T]uint) {
	*ref = make(map[T]T)
	*inst = make(map[T]T)
	*refOrder = make(map[T]uint)
}

func __gong__resetStageType[T comparable](staged *map[T]struct{}, mapString *map[string]T, stagedOrder *map[T]uint, order *uint) {
	*staged = make(map[T]struct{})
	*mapString = make(map[string]T)
	*stagedOrder = make(map[T]uint)
	*order = 0
}

func (stage *Stage) GetType() string {
	return "github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models"
}

type GONG__Identifier struct {
	Ident string
	Type  GONG__ExpressionType
}

type GongOnInitCommitInterface interface {
	BeforeCommit(stage *Stage)
}

type OnInitCommitInterface = GongOnInitCommitInterface

// GongOnAfterCreateInterface callback when an instance is updated from the front
type GongOnAfterCreateInterface[Type Gongstruct] interface {
	OnAfterCreate(stage *Stage,
		instance *Type)
}

type OnAfterCreateInterface[Type Gongstruct] = GongOnAfterCreateInterface[Type]

// GongOnAfterUpdateInterface callback when an instance is updated from the front
type GongOnAfterUpdateInterface[Type Gongstruct] interface {
	OnAfterUpdate(stage *Stage, old, new *Type)
}

type OnAfterUpdateInterface[Type Gongstruct] = GongOnAfterUpdateInterface[Type]

// GongOnAfterDeleteInterface callback when an instance is updated from the front
type GongOnAfterDeleteInterface[Type Gongstruct] interface {
	OnAfterDelete(stage *Stage,
		staged, front *Type)
}

type OnAfterDeleteInterface[Type Gongstruct] = GongOnAfterDeleteInterface[Type]

type GongBackRepoInterface interface {
	Commit(stage *Stage)
	Checkout(stage *Stage)
	Backup(stage *Stage, dirPath string)
	Restore(stage *Stage, dirPath string)
	BackupXL(stage *Stage, dirPath string)
	RestoreXL(stage *Stage, dirPath string)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

type BackRepoInterface = GongBackRepoInterface

func NewStage(name string) (stage *Stage) {
	stage = &Stage{ // insertion point for array initiatialisation
		AIXMBasicMessages:           make(map[*AIXMBasicMessage]struct{}),
		AIXMBasicMessages_mapString: make(map[string]*AIXMBasicMessage),

		AIXMBasicMessageTypes:           make(map[*AIXMBasicMessageType]struct{}),
		AIXMBasicMessageTypes_mapString: make(map[string]*AIXMBasicMessageType),

		BasicMessageMemberAIXMPropertyTypes:           make(map[*BasicMessageMemberAIXMPropertyType]struct{}),
		BasicMessageMemberAIXMPropertyTypes_mapString: make(map[string]*BasicMessageMemberAIXMPropertyType),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		AIXMBasicMessage_stagedOrder: make(map[*AIXMBasicMessage]uint),
		AIXMBasicMessage_orderStaged: make(map[uint]*AIXMBasicMessage),
		AIXMBasicMessages_reference:  make(map[*AIXMBasicMessage]*AIXMBasicMessage),

		AIXMBasicMessageType_stagedOrder: make(map[*AIXMBasicMessageType]uint),
		AIXMBasicMessageType_orderStaged: make(map[uint]*AIXMBasicMessageType),
		AIXMBasicMessageTypes_reference:  make(map[*AIXMBasicMessageType]*AIXMBasicMessageType),

		BasicMessageMemberAIXMPropertyType_stagedOrder: make(map[*BasicMessageMemberAIXMPropertyType]uint),
		BasicMessageMemberAIXMPropertyType_orderStaged: make(map[uint]*BasicMessageMemberAIXMPropertyType),
		BasicMessageMemberAIXMPropertyTypes_reference:  make(map[*BasicMessageMemberAIXMPropertyType]*BasicMessageMemberAIXMPropertyType),

		// end of insertion point
		GongUnmarshallers: map[string]GongModelUnmarshaller{ // insertion point for unmarshallers
			"AIXMBasicMessage": &AIXMBasicMessageUnmarshaller{},

			"AIXMBasicMessageType": &AIXMBasicMessageTypeUnmarshaller{},

			"BasicMessageMemberAIXMPropertyType": &BasicMessageMemberAIXMPropertyTypeUnmarshaller{},

			// end of insertion point
		},
	}
	stage.StageCore.SetName(name)
	stage.StageCore.SetNavigationMode(GongNavigationModeNormal)
	subPath_AixmfeaturesStage := "aixmfeatures"
	if name != "" {
		subPath_AixmfeaturesStage = name + "_aixmfeatures"
	}
	stage.AixmfeaturesStage = aixmfeatures.NewStage(subPath_AixmfeaturesStage)
	subPath_AixmdatatypesStage := "aixmdatatypes"
	if name != "" {
		subPath_AixmdatatypesStage = name + "_aixmdatatypes"
	}
	stage.AixmdatatypesStage = aixmdatatypes.NewStage(subPath_AixmdatatypesStage)
	subPath_AixmabstractgmlobjecttypesStage := "aixmabstractgmlobjecttypes"
	if name != "" {
		subPath_AixmabstractgmlobjecttypesStage = name + "_aixmabstractgmlobjecttypes"
	}
	stage.AixmabstractgmlobjecttypesStage = aixmabstractgmlobjecttypes.NewStage(subPath_AixmabstractgmlobjecttypesStage)
	subPath_GmlStage := "gml"
	if name != "" {
		subPath_GmlStage = name + "_gml"
	}
	stage.GmlStage = gml.NewStage(subPath_GmlStage)
	subPath_XlinkStage := "xlink"
	if name != "" {
		subPath_XlinkStage = name + "_xlink"
	}
	stage.XlinkStage = xlink.NewStage(subPath_XlinkStage)

	return
}

// GetOrder is the Stage method returning the order of a gongstruct instance.
func (stage *Stage) GetOrder(instance GongstructIF) uint {
	if instance != nil {
		return instance.GongGetOrder(stage)
	}
	return 0
}

// GetInstanceFromOrder is the Stage method returning a gongstruct instance from its order.
func (stage *Stage) GetInstanceFromOrder[Type GongstructPtr](order uint) (res Type) {
	if stage == nil {
		return
	}
	var t Type
	val := t.GongGetInstanceFromOrder(stage, order)
	if val != nil {
		res = val.(Type)
	}
	return
}

func (stage *Stage) CommitWithSuspendedCallbacks() {
	tmp := stage.OnInitCommitFromBackCallback
	stage.OnInitCommitFromBackCallback = nil
	tmp2 := stage.StageCore.GetBeforeCommitHooks()
	stage.StageCore.ClearBeforeCommitHooks()
	tmp3 := stage.StageCore.GetAfterCommitHooks()
	stage.StageCore.ClearAfterCommitHooks()
	stage.Commit()
	stage.OnInitCommitFromBackCallback = tmp
	stage.StageCore.SetBeforeCommitHooks(tmp2)
	stage.StageCore.SetAfterCommitHooks(tmp3)
}

func (stage *Stage) Commit() {
	if stage.XlinkStage != nil {
		stage.XlinkStage.Commit()
	}
	if stage.GmlStage != nil {
		stage.GmlStage.Commit()
	}
	if stage.AixmabstractgmlobjecttypesStage != nil {
		stage.AixmabstractgmlobjecttypesStage.Commit()
	}
	if stage.AixmdatatypesStage != nil {
		stage.AixmdatatypesStage.Commit()
	}
	if stage.AixmfeaturesStage != nil {
		stage.AixmfeaturesStage.Commit()
	}
	stage.ComputeReverseMaps()

	if stage.OnInitCommitCallback != nil {
		stage.OnInitCommitCallback.BeforeCommit(stage)
	}
	if stage.OnInitCommitFromBackCallback != nil {
		stage.OnInitCommitFromBackCallback.BeforeCommit(stage)
	}

	// 1. Run all Before Commit hooks
	stage.StageCore.RunBeforeCommitHooks()

	if stage.BackRepo != nil {
		stage.BackRepo.Commit(stage)
	}
	stage.ComputeInstancesNb()

	// if a commit is applied when in navigation mode
	// this will reset the commits behind and swith the
	// naviagation
	if stage.IsInDeltaMode() && stage.GetNavigationMode() == GongNavigationModeNavigating && stage.GetCommitsBehind() > 0 {
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
	stage.StageCore.RunAfterCommitHooks()
}

func (stage *Stage) ComputeInstancesNb() {
	if stage.XlinkStage != nil {
		stage.XlinkStage.ComputeInstancesNb()
	}
	if stage.GmlStage != nil {
		stage.GmlStage.ComputeInstancesNb()
	}
	if stage.AixmabstractgmlobjecttypesStage != nil {
		stage.AixmabstractgmlobjecttypesStage.ComputeInstancesNb()
	}
	if stage.AixmdatatypesStage != nil {
		stage.AixmdatatypesStage.ComputeInstancesNb()
	}
	if stage.AixmfeaturesStage != nil {
		stage.AixmfeaturesStage.ComputeInstancesNb()
	}
	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["AIXMBasicMessage"] = len(stage.AIXMBasicMessages)
	stage.Map_GongStructName_InstancesNb["AIXMBasicMessageType"] = len(stage.AIXMBasicMessageTypes)
	stage.Map_GongStructName_InstancesNb["BasicMessageMemberAIXMPropertyType"] = len(stage.BasicMessageMemberAIXMPropertyTypes)
}

func (stage *Stage) Checkout() {
	if stage.XlinkStage != nil {
		stage.XlinkStage.Checkout()
	}
	if stage.GmlStage != nil {
		stage.GmlStage.Checkout()
	}
	if stage.AixmabstractgmlobjecttypesStage != nil {
		stage.AixmabstractgmlobjecttypesStage.Checkout()
	}
	if stage.AixmdatatypesStage != nil {
		stage.AixmdatatypesStage.Checkout()
	}
	if stage.AixmfeaturesStage != nil {
		stage.AixmfeaturesStage.Checkout()
	}
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
// Stage puts aixmbasicmessage to the model stage
func (aixmbasicmessage *AIXMBasicMessage) Stage(stage *Stage) *AIXMBasicMessage {
	__gong__stage(stage.AIXMBasicMessages, stage.AIXMBasicMessage_stagedOrder, stage.AIXMBasicMessage_orderStaged, &stage.AIXMBasicMessageOrder, stage.AIXMBasicMessages_mapString, aixmbasicmessage, aixmbasicmessage.Name)
	return aixmbasicmessage
}

// StagePreserveOrder puts aixmbasicmessage to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AIXMBasicMessageOrder
// - update stage.AIXMBasicMessageOrder accordingly
func (aixmbasicmessage *AIXMBasicMessage) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AIXMBasicMessages, stage.AIXMBasicMessage_stagedOrder, stage.AIXMBasicMessage_orderStaged, &stage.AIXMBasicMessageOrder, stage.AIXMBasicMessages_mapString, aixmbasicmessage, order, aixmbasicmessage.Name)
}

// Unstage removes aixmbasicmessage off the model stage
func (aixmbasicmessage *AIXMBasicMessage) Unstage(stage *Stage) *AIXMBasicMessage {
	__gong__unstage(stage.AIXMBasicMessages, stage.AIXMBasicMessages_mapString, aixmbasicmessage, aixmbasicmessage.Name)
	return aixmbasicmessage
}

// UnstageVoid removes aixmbasicmessage off the model stage
func (aixmbasicmessage *AIXMBasicMessage) UnstageVoid(stage *Stage) {
	aixmbasicmessage.Unstage(stage)
}

func (aixmbasicmessage *AIXMBasicMessage) StageVoid(stage *Stage) {
	aixmbasicmessage.Stage(stage)
}

// for satisfaction of GongStruct interface
func (aixmbasicmessage *AIXMBasicMessage) GetName() (res string) {
	return aixmbasicmessage.Name
}

// for satisfaction of GongStruct interface
func (aixmbasicmessage *AIXMBasicMessage) SetName(name string) {
	aixmbasicmessage.Name = name
}

// Stage puts aixmbasicmessagetype to the model stage
func (aixmbasicmessagetype *AIXMBasicMessageType) Stage(stage *Stage) *AIXMBasicMessageType {
	__gong__stage(stage.AIXMBasicMessageTypes, stage.AIXMBasicMessageType_stagedOrder, stage.AIXMBasicMessageType_orderStaged, &stage.AIXMBasicMessageTypeOrder, stage.AIXMBasicMessageTypes_mapString, aixmbasicmessagetype, aixmbasicmessagetype.Name)
	return aixmbasicmessagetype
}

// StagePreserveOrder puts aixmbasicmessagetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AIXMBasicMessageTypeOrder
// - update stage.AIXMBasicMessageTypeOrder accordingly
func (aixmbasicmessagetype *AIXMBasicMessageType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AIXMBasicMessageTypes, stage.AIXMBasicMessageType_stagedOrder, stage.AIXMBasicMessageType_orderStaged, &stage.AIXMBasicMessageTypeOrder, stage.AIXMBasicMessageTypes_mapString, aixmbasicmessagetype, order, aixmbasicmessagetype.Name)
}

// Unstage removes aixmbasicmessagetype off the model stage
func (aixmbasicmessagetype *AIXMBasicMessageType) Unstage(stage *Stage) *AIXMBasicMessageType {
	__gong__unstage(stage.AIXMBasicMessageTypes, stage.AIXMBasicMessageTypes_mapString, aixmbasicmessagetype, aixmbasicmessagetype.Name)
	return aixmbasicmessagetype
}

// UnstageVoid removes aixmbasicmessagetype off the model stage
func (aixmbasicmessagetype *AIXMBasicMessageType) UnstageVoid(stage *Stage) {
	aixmbasicmessagetype.Unstage(stage)
}

func (aixmbasicmessagetype *AIXMBasicMessageType) StageVoid(stage *Stage) {
	aixmbasicmessagetype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (aixmbasicmessagetype *AIXMBasicMessageType) GetName() (res string) {
	return aixmbasicmessagetype.Name
}

// for satisfaction of GongStruct interface
func (aixmbasicmessagetype *AIXMBasicMessageType) SetName(name string) {
	aixmbasicmessagetype.Name = name
}

// Stage puts basicmessagememberaixmpropertytype to the model stage
func (basicmessagememberaixmpropertytype *BasicMessageMemberAIXMPropertyType) Stage(stage *Stage) *BasicMessageMemberAIXMPropertyType {
	__gong__stage(stage.BasicMessageMemberAIXMPropertyTypes, stage.BasicMessageMemberAIXMPropertyType_stagedOrder, stage.BasicMessageMemberAIXMPropertyType_orderStaged, &stage.BasicMessageMemberAIXMPropertyTypeOrder, stage.BasicMessageMemberAIXMPropertyTypes_mapString, basicmessagememberaixmpropertytype, basicmessagememberaixmpropertytype.Name)
	return basicmessagememberaixmpropertytype
}

// StagePreserveOrder puts basicmessagememberaixmpropertytype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.BasicMessageMemberAIXMPropertyTypeOrder
// - update stage.BasicMessageMemberAIXMPropertyTypeOrder accordingly
func (basicmessagememberaixmpropertytype *BasicMessageMemberAIXMPropertyType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.BasicMessageMemberAIXMPropertyTypes, stage.BasicMessageMemberAIXMPropertyType_stagedOrder, stage.BasicMessageMemberAIXMPropertyType_orderStaged, &stage.BasicMessageMemberAIXMPropertyTypeOrder, stage.BasicMessageMemberAIXMPropertyTypes_mapString, basicmessagememberaixmpropertytype, order, basicmessagememberaixmpropertytype.Name)
}

// Unstage removes basicmessagememberaixmpropertytype off the model stage
func (basicmessagememberaixmpropertytype *BasicMessageMemberAIXMPropertyType) Unstage(stage *Stage) *BasicMessageMemberAIXMPropertyType {
	__gong__unstage(stage.BasicMessageMemberAIXMPropertyTypes, stage.BasicMessageMemberAIXMPropertyTypes_mapString, basicmessagememberaixmpropertytype, basicmessagememberaixmpropertytype.Name)
	return basicmessagememberaixmpropertytype
}

// UnstageVoid removes basicmessagememberaixmpropertytype off the model stage
func (basicmessagememberaixmpropertytype *BasicMessageMemberAIXMPropertyType) UnstageVoid(stage *Stage) {
	basicmessagememberaixmpropertytype.Unstage(stage)
}

func (basicmessagememberaixmpropertytype *BasicMessageMemberAIXMPropertyType) StageVoid(stage *Stage) {
	basicmessagememberaixmpropertytype.Stage(stage)
}

// for satisfaction of GongStruct interface
func (basicmessagememberaixmpropertytype *BasicMessageMemberAIXMPropertyType) GetName() (res string) {
	return basicmessagememberaixmpropertytype.Name
}

// for satisfaction of GongStruct interface
func (basicmessagememberaixmpropertytype *BasicMessageMemberAIXMPropertyType) SetName(name string) {
	basicmessagememberaixmpropertytype.Name = name
}

func (stage *Stage) Reset() { // insertion point for array reset
	__gong__resetStageType(&stage.AIXMBasicMessages, &stage.AIXMBasicMessages_mapString, &stage.AIXMBasicMessage_stagedOrder, &stage.AIXMBasicMessageOrder)

	__gong__resetStageType(&stage.AIXMBasicMessageTypes, &stage.AIXMBasicMessageTypes_mapString, &stage.AIXMBasicMessageType_stagedOrder, &stage.AIXMBasicMessageTypeOrder)

	__gong__resetStageType(&stage.BasicMessageMemberAIXMPropertyTypes, &stage.BasicMessageMemberAIXMPropertyTypes_mapString, &stage.BasicMessageMemberAIXMPropertyType_stagedOrder, &stage.BasicMessageMemberAIXMPropertyTypeOrder)

	if stage.XlinkStage != nil {
		stage.XlinkStage.Reset()
	}
	if stage.GmlStage != nil {
		stage.GmlStage.Reset()
	}
	if stage.AixmabstractgmlobjecttypesStage != nil {
		stage.AixmabstractgmlobjecttypesStage.Reset()
	}
	if stage.AixmdatatypesStage != nil {
		stage.AixmdatatypesStage.Reset()
	}
	if stage.AixmfeaturesStage != nil {
		stage.AixmfeaturesStage.Reset()
	}
	if stage.GetProbeIF() != nil {
		stage.GetProbeIF().ResetNotifications()
	}
	if stage.IsInDeltaMode() {
		stage.ComputeReferenceAndOrders()
	}
}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type Gongstruct interface {
	GongGetAssociationName() any
}

type GongstructBasicField interface {
	int | float64 | bool | string | time.Time | time.Duration
}

type GongtructBasicField = GongstructBasicField

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type GongstructIF interface {
	GetName() string
	SetName(string)
	StageVoid(*Stage)
	UnstageVoid(stage *Stage)
	GongGetFieldHeaders() []GongFieldHeader
	GongGetFieldValue(fieldName string, stage *Stage) GongFieldValue
	GongGetGongstructName() string
	GongGetOrder(stage *Stage) uint
	GongGetReferenceIdentifier(stage *Stage) string
	GongGetIdentifier(stage *Stage) string
	GongCopy() GongstructIF
	GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) string
	GongGetUUID(stage *Stage) string
	GongAfterCreateFromFront(stage *Stage)
	GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF)
	GongAfterDeleteFromFront(stage *Stage, front GongstructIF)
	GongIsStaged(stage *Stage) bool
	GongStageBranch(stage *Stage)
	GongUnstageBranch(stage *Stage)

	GongGetInstancesByOrder(stage *Stage) any
	GongGetInstanceFromOrder(stage *Stage, order uint) any
	GongGetInstancesMapByName(stage *Stage) any
	GongGetInstancesSet(stage *Stage) any
	GongNewInstance() any
	GongGetReverseFields() []GongReverseField
}
type GongstructPtr interface {
	GongstructIF
	comparable
}

type PointerToGongstruct = GongstructPtr

func GongCompareGongstructByName[T GongstructPtr](a, b T) int {
	return cmp.Compare(a.GetName(), b.GetName())
}

func GongSortGongstructSetByName[T GongstructPtr](set map[T]struct{}) (sortedSlice []T) {
	for key := range set {
		sortedSlice = append(sortedSlice, key)
	}
	slices.SortFunc(sortedSlice, GongCompareGongstructByName)

	return
}

// GetInstancesSorted is the Stage method returning sorted instances of a gongstruct.
func (stage *Stage) GetInstancesSorted[T GongstructPtr]() (sortedSlice []T) {
	set := stage.GetInstancesSet[T]()
	sortedSlice = GongSortGongstructSetByName(*set)

	return
}

// GetInstancesMapByName is the Stage method returning a map of staged instances by their name.
func (stage *Stage) GetInstancesMapByName[Type GongstructIF]() map[string]Type {
	if stage == nil {
		return nil
	}
	var t Type
	if any(t) == nil {
		return nil
	}
	return t.GongGetInstancesMapByName(stage).(map[string]Type)
}

// GetInstancesSet is the Stage method returning the set of staged instances (pointer-type constraint).
func (stage *Stage) GetInstancesSet[Type GongstructPtr]() *map[Type]struct{} {
	if stage == nil {
		return nil
	}
	var t Type
	return t.GongGetInstancesSet(stage).(*map[Type]struct{})
}

// insertion point for instance with special fields
func (AIXMBasicMessage) GongGetAssociationName() any {
	return &AIXMBasicMessage{
			HasMember: []*BasicMessageMemberAIXMPropertyType{{Name: "HasMember"}},
	}
}

func (AIXMBasicMessageType) GongGetAssociationName() any {
	return &AIXMBasicMessageType{
			HasMember: []*BasicMessageMemberAIXMPropertyType{{Name: "HasMember"}},
	}
}

func (BasicMessageMemberAIXMPropertyType) GongGetAssociationName() any {
	return &BasicMessageMemberAIXMPropertyType{
	}
}


// GongGetAssociationName is a generic function that returns an instance of Type
// where each association is filled with an instance whose name is the name of the association
//
// This function can be handy for generating navigation function that are refactorable
func GongGetAssociationName[Type Gongstruct]() *Type {
	var t Type
	return t.GongGetAssociationName().(*Type)
}

// GetPointerReverseMap allows backtrack navigation of any Start.Fieldname
// associations (0..1) that is a pointer from one staged Gongstruct (type Start)
// instances to another (type End)
//
// The function provides a map with keys as instances of End and values to arrays of *Start
// the map is construed by iterating over all Start instances and populationg keys with End instances
// and values with slice of Start instances
// GetPointerReverseMap is the Stage method for backtrack navigation of pointer associations.
func (stage *Stage) GetPointerReverseMap[Start, End Gongstruct](fieldname string) map[*End][]*Start {
	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of AIXMBasicMessage
	case AIXMBasicMessage:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of AIXMBasicMessageType
	case AIXMBasicMessageType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of BasicMessageMemberAIXMPropertyType
	case BasicMessageMemberAIXMPropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	}
	return nil
}

// GetSliceOfPointersReverseMap is the Stage method for backtrack navigation of slice-of-pointers associations.
func (stage *Stage) GetSliceOfPointersReverseMap[Start, End Gongstruct](fieldname string) map[*End][]*Start {
	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of AIXMBasicMessage
	case AIXMBasicMessage:
		switch fieldname {
		// insertion point for per direct association field
		case "HasMember":
			res := make(map[*BasicMessageMemberAIXMPropertyType][]*AIXMBasicMessage)
			for aixmbasicmessage := range stage.AIXMBasicMessages {
				for _, basicmessagememberaixmpropertytype_ := range aixmbasicmessage.HasMember {
					res[basicmessagememberaixmpropertytype_] = append(res[basicmessagememberaixmpropertytype_], aixmbasicmessage)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of AIXMBasicMessageType
	case AIXMBasicMessageType:
		switch fieldname {
		// insertion point for per direct association field
		case "HasMember":
			res := make(map[*BasicMessageMemberAIXMPropertyType][]*AIXMBasicMessageType)
			for aixmbasicmessagetype := range stage.AIXMBasicMessageTypes {
				for _, basicmessagememberaixmpropertytype_ := range aixmbasicmessagetype.HasMember {
					res[basicmessagememberaixmpropertytype_] = append(res[basicmessagememberaixmpropertytype_], aixmbasicmessagetype)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of BasicMessageMemberAIXMPropertyType
	case BasicMessageMemberAIXMPropertyType:
		switch fieldname {
		// insertion point for per direct association field
		}
	}
	return nil
}

// GongNewInstance creates a new instance of the Gongstruct
func GongNewInstance[Type GongstructPtr]() (res Type) {
	var t Type
	return t.GongNewInstance().(Type)
}

func NewInstance[Type GongstructPtr]() (res Type) {
	return GongNewInstance[Type]()
}

func (stage *Stage) GongNewInstance[Type GongstructPtr]() (res Type) {
	res = GongNewInstance[Type]()
	var zero Type
	if res != zero {
		res.StageVoid(stage)
	}
	return res
}

func (stage *Stage) NewInstance[Type GongstructPtr]() (res Type) {
	return stage.GongNewInstance[Type]()
}

// GongGetPointerToGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GongGetPointerToGongstructName[Type GongstructIF]() (res string) {
	var t Type
	if any(t) == nil {
		return ""
	}
	return t.GongGetGongstructName()
}

func GetPointerToGongstructName[Type GongstructIF]() (res string) {
	return GongGetPointerToGongstructName[Type]()
}

type GongReverseField struct {
	GongstructName string
	Fieldname      string
}

type ReverseField = GongReverseField

// insertion point for generic get reverse fields
func (*AIXMBasicMessage) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*AIXMBasicMessageType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*BasicMessageMemberAIXMPropertyType) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "AIXMBasicMessage",
			Fieldname: "HasMember",
		},
		{
			GongstructName: "AIXMBasicMessageType",
			Fieldname: "HasMember",
		},
	}
}


func GongGetReverseFields[Type GongstructIF]() (res []GongReverseField) {
	var t Type
	if any(t) == nil {
		return nil
	}
	return t.GongGetReverseFields()
}

func GetReverseFields[Type GongstructIF]() (res []GongReverseField) {
	return GongGetReverseFields[Type]()
}

// insertion point for get fields header method
func (aixmbasicmessage *AIXMBasicMessage) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Id",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Description",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "SequenceNumber",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "AggregationType",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "HasMember",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "BasicMessageMemberAIXMPropertyType",
		},
	}
	return
}

func (aixmbasicmessagetype *AIXMBasicMessageType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Id",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Description",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "SequenceNumber",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "AggregationType",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "HasMember",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "BasicMessageMemberAIXMPropertyType",
		},
	}
	return
}

func (basicmessagememberaixmpropertytype *BasicMessageMemberAIXMPropertyType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NilReason",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Type_XSD",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Href",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Role",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Arcrole",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Title",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Show",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Actuate",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

// GongGetFieldsFromPointer return the array of the fields
func GongGetFieldsFromPointer[Type GongstructPtr]() (res []GongFieldHeader) {
	var ret Type
	return ret.GongGetFieldHeaders()
}

func GetFieldsFromPointer[Type GongstructPtr]() (res []GongFieldHeader) {
	return GongGetFieldsFromPointer[Type]()
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
func (aixmbasicmessage *AIXMBasicMessage) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = aixmbasicmessage.Name
	case "Id":
		res.valueString = aixmbasicmessage.Id
	case "Description":
		res.valueString = aixmbasicmessage.Description
	case "SequenceNumber":
		res.valueString = fmt.Sprintf("%d", aixmbasicmessage.SequenceNumber)
		res.valueInt = aixmbasicmessage.SequenceNumber
		res.GongFieldValueType = GongFieldValueTypeInt
	case "AggregationType":
		res.valueString = aixmbasicmessage.AggregationType
	case "HasMember":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range aixmbasicmessage.HasMember {
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

func (aixmbasicmessagetype *AIXMBasicMessageType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = aixmbasicmessagetype.Name
	case "Id":
		res.valueString = aixmbasicmessagetype.Id
	case "Description":
		res.valueString = aixmbasicmessagetype.Description
	case "SequenceNumber":
		res.valueString = fmt.Sprintf("%d", aixmbasicmessagetype.SequenceNumber)
		res.valueInt = aixmbasicmessagetype.SequenceNumber
		res.GongFieldValueType = GongFieldValueTypeInt
	case "AggregationType":
		res.valueString = aixmbasicmessagetype.AggregationType
	case "HasMember":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range aixmbasicmessagetype.HasMember {
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

func (basicmessagememberaixmpropertytype *BasicMessageMemberAIXMPropertyType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = basicmessagememberaixmpropertytype.Name
	case "NilReason":
		res.valueString = basicmessagememberaixmpropertytype.NilReason
	case "Type_XSD":
		res.valueString = basicmessagememberaixmpropertytype.Type_XSD
	case "Href":
		res.valueString = basicmessagememberaixmpropertytype.Href
	case "Role":
		res.valueString = basicmessagememberaixmpropertytype.Role
	case "Arcrole":
		res.valueString = basicmessagememberaixmpropertytype.Arcrole
	case "Title":
		res.valueString = basicmessagememberaixmpropertytype.Title
	case "Show":
		res.valueString = basicmessagememberaixmpropertytype.Show
	case "Actuate":
		res.valueString = basicmessagememberaixmpropertytype.Actuate
	}
	return
}

func (stage *Stage) GetFieldStringValueFromPointer(instance GongstructIF, fieldName string) (res GongFieldValue) {
	res = instance.GongGetFieldValue(fieldName, stage)
	return
}

func GetFieldStringValueFromPointer(instance GongstructIF, fieldName string, stage *Stage) (res GongFieldValue) {
	return stage.GetFieldStringValueFromPointer(instance, fieldName)
}

// insertion point for generic get gongstruct name
func (aixmbasicmessage *AIXMBasicMessage) GongGetGongstructName() string {
	return "AIXMBasicMessage"
}

func (aixmbasicmessagetype *AIXMBasicMessageType) GongGetGongstructName() string {
	return "AIXMBasicMessageType"
}

func (basicmessagememberaixmpropertytype *BasicMessageMemberAIXMPropertyType) GongGetGongstructName() string {
	return "BasicMessageMemberAIXMPropertyType"
}

func GongGetGongstructNameFromPointer(instance GongstructIF) (res string) {
	res = instance.GongGetGongstructName()
	return
}

func GetGongstructNameFromPointer(instance GongstructIF) (res string) {
	return GongGetGongstructNameFromPointer(instance)
}

func (stage *Stage) ResetMapStrings() {
	// insertion point for generic get gongstruct name
	__gong__rebuildMapString(stage.AIXMBasicMessages, &stage.AIXMBasicMessages_mapString)

	__gong__rebuildMapString(stage.AIXMBasicMessageTypes, &stage.AIXMBasicMessageTypes_mapString)

	__gong__rebuildMapString(stage.BasicMessageMemberAIXMPropertyTypes, &stage.BasicMessageMemberAIXMPropertyTypes_mapString)

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
