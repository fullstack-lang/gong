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

	gantt_go "github.com/fullstack-lang/gong/lib/gantt/go"
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
// swagger:ignore
type Stage struct {
	name string

	// insertion point for definition of arrays registering instances
	Arrows           map[*Arrow]struct{}
	Arrows_reference map[*Arrow]*Arrow
	Arrows_mapString map[string]*Arrow

	// insertion point for slice of pointers maps
	OnAfterArrowCreateCallback OnAfterCreateInterface[Arrow]
	OnAfterArrowUpdateCallback OnAfterUpdateInterface[Arrow]
	OnAfterArrowDeleteCallback OnAfterDeleteInterface[Arrow]
	OnAfterArrowReadCallback   OnAfterReadInterface[Arrow]

	Bars           map[*Bar]struct{}
	Bars_reference map[*Bar]*Bar
	Bars_mapString map[string]*Bar

	// insertion point for slice of pointers maps
	OnAfterBarCreateCallback OnAfterCreateInterface[Bar]
	OnAfterBarUpdateCallback OnAfterUpdateInterface[Bar]
	OnAfterBarDeleteCallback OnAfterDeleteInterface[Bar]
	OnAfterBarReadCallback   OnAfterReadInterface[Bar]

	Gantts           map[*Gantt]struct{}
	Gantts_reference map[*Gantt]*Gantt
	Gantts_mapString map[string]*Gantt

	// insertion point for slice of pointers maps
	Gantt_Lanes_reverseMap map[*Lane]*Gantt

	Gantt_Milestones_reverseMap map[*Milestone]*Gantt

	Gantt_Groups_reverseMap map[*Group]*Gantt

	Gantt_Arrows_reverseMap map[*Arrow]*Gantt

	OnAfterGanttCreateCallback OnAfterCreateInterface[Gantt]
	OnAfterGanttUpdateCallback OnAfterUpdateInterface[Gantt]
	OnAfterGanttDeleteCallback OnAfterDeleteInterface[Gantt]
	OnAfterGanttReadCallback   OnAfterReadInterface[Gantt]

	Groups           map[*Group]struct{}
	Groups_reference map[*Group]*Group
	Groups_mapString map[string]*Group

	// insertion point for slice of pointers maps
	Group_GroupLanes_reverseMap map[*Lane]*Group

	OnAfterGroupCreateCallback OnAfterCreateInterface[Group]
	OnAfterGroupUpdateCallback OnAfterUpdateInterface[Group]
	OnAfterGroupDeleteCallback OnAfterDeleteInterface[Group]
	OnAfterGroupReadCallback   OnAfterReadInterface[Group]

	Lanes           map[*Lane]struct{}
	Lanes_reference map[*Lane]*Lane
	Lanes_mapString map[string]*Lane

	// insertion point for slice of pointers maps
	Lane_Bars_reverseMap map[*Bar]*Lane

	OnAfterLaneCreateCallback OnAfterCreateInterface[Lane]
	OnAfterLaneUpdateCallback OnAfterUpdateInterface[Lane]
	OnAfterLaneDeleteCallback OnAfterDeleteInterface[Lane]
	OnAfterLaneReadCallback   OnAfterReadInterface[Lane]

	LaneUses           map[*LaneUse]struct{}
	LaneUses_reference map[*LaneUse]*LaneUse
	LaneUses_mapString map[string]*LaneUse

	// insertion point for slice of pointers maps
	OnAfterLaneUseCreateCallback OnAfterCreateInterface[LaneUse]
	OnAfterLaneUseUpdateCallback OnAfterUpdateInterface[LaneUse]
	OnAfterLaneUseDeleteCallback OnAfterDeleteInterface[LaneUse]
	OnAfterLaneUseReadCallback   OnAfterReadInterface[LaneUse]

	Milestones           map[*Milestone]struct{}
	Milestones_reference map[*Milestone]*Milestone
	Milestones_mapString map[string]*Milestone

	// insertion point for slice of pointers maps
	Milestone_LanesToDisplay_reverseMap map[*Lane]*Milestone

	OnAfterMilestoneCreateCallback OnAfterCreateInterface[Milestone]
	OnAfterMilestoneUpdateCallback OnAfterUpdateInterface[Milestone]
	OnAfterMilestoneDeleteCallback OnAfterDeleteInterface[Milestone]
	OnAfterMilestoneReadCallback   OnAfterReadInterface[Milestone]

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
	ArrowOrder            uint
	ArrowMap_Staged_Order map[*Arrow]uint

	BarOrder            uint
	BarMap_Staged_Order map[*Bar]uint

	GanttOrder            uint
	GanttMap_Staged_Order map[*Gantt]uint

	GroupOrder            uint
	GroupMap_Staged_Order map[*Group]uint

	LaneOrder            uint
	LaneMap_Staged_Order map[*Lane]uint

	LaneUseOrder            uint
	LaneUseMap_Staged_Order map[*LaneUse]uint

	MilestoneOrder            uint
	MilestoneMap_Staged_Order map[*Milestone]uint

	// end of insertion point

	NamedStructs []*NamedStruct

	// probeIF is the interface to the probe that allows log
	// commit event to the probe
	probeIF ProbeIF
}

func (stage *Stage) SetProbeIF(probeIF ProbeIF) {
	stage.probeIF = probeIF
}

func (stage *Stage) GetProbeIF() ProbeIF {
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
	case *Arrow:
		tmp := GetStructInstancesByOrder(stage.Arrows, stage.ArrowMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Arrow implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Bar:
		tmp := GetStructInstancesByOrder(stage.Bars, stage.BarMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Bar implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Gantt:
		tmp := GetStructInstancesByOrder(stage.Gantts, stage.GanttMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Gantt implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Group:
		tmp := GetStructInstancesByOrder(stage.Groups, stage.GroupMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Group implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Lane:
		tmp := GetStructInstancesByOrder(stage.Lanes, stage.LaneMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Lane implements.
			res = append(res, any(v).(T))
		}
		return res
	case *LaneUse:
		tmp := GetStructInstancesByOrder(stage.LaneUses, stage.LaneUseMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *LaneUse implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Milestone:
		tmp := GetStructInstancesByOrder(stage.Milestones, stage.MilestoneMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Milestone implements.
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
	case "Arrow":
		res = GetNamedStructInstances(stage.Arrows, stage.ArrowMap_Staged_Order)
	case "Bar":
		res = GetNamedStructInstances(stage.Bars, stage.BarMap_Staged_Order)
	case "Gantt":
		res = GetNamedStructInstances(stage.Gantts, stage.GanttMap_Staged_Order)
	case "Group":
		res = GetNamedStructInstances(stage.Groups, stage.GroupMap_Staged_Order)
	case "Lane":
		res = GetNamedStructInstances(stage.Lanes, stage.LaneMap_Staged_Order)
	case "LaneUse":
		res = GetNamedStructInstances(stage.LaneUses, stage.LaneUseMap_Staged_Order)
	case "Milestone":
		res = GetNamedStructInstances(stage.Milestones, stage.MilestoneMap_Staged_Order)
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
	return "github.com/fullstack-lang/gong/lib/gantt/go/models"
}

func (stage *Stage) GetMap_GongStructName_InstancesNb() map[string]int {
	return stage.Map_GongStructName_InstancesNb
}

func (stage *Stage) GetModelsEmbededDir() embed.FS {
	return gantt_go.GoModelsDir
}

func (stage *Stage) GetDigramsEmbededDir() embed.FS {
	return gantt_go.GoDiagramsDir
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
	CommitArrow(arrow *Arrow)
	CheckoutArrow(arrow *Arrow)
	CommitBar(bar *Bar)
	CheckoutBar(bar *Bar)
	CommitGantt(gantt *Gantt)
	CheckoutGantt(gantt *Gantt)
	CommitGroup(group *Group)
	CheckoutGroup(group *Group)
	CommitLane(lane *Lane)
	CheckoutLane(lane *Lane)
	CommitLaneUse(laneuse *LaneUse)
	CheckoutLaneUse(laneuse *LaneUse)
	CommitMilestone(milestone *Milestone)
	CheckoutMilestone(milestone *Milestone)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(name string) (stage *Stage) {

	stage = &Stage{ // insertion point for array initiatialisation
		Arrows:           make(map[*Arrow]struct{}),
		Arrows_mapString: make(map[string]*Arrow),

		Bars:           make(map[*Bar]struct{}),
		Bars_mapString: make(map[string]*Bar),

		Gantts:           make(map[*Gantt]struct{}),
		Gantts_mapString: make(map[string]*Gantt),

		Groups:           make(map[*Group]struct{}),
		Groups_mapString: make(map[string]*Group),

		Lanes:           make(map[*Lane]struct{}),
		Lanes_mapString: make(map[string]*Lane),

		LaneUses:           make(map[*LaneUse]struct{}),
		LaneUses_mapString: make(map[string]*LaneUse),

		Milestones:           make(map[*Milestone]struct{}),
		Milestones_mapString: make(map[string]*Milestone),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		name: name,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		ArrowMap_Staged_Order: make(map[*Arrow]uint),

		BarMap_Staged_Order: make(map[*Bar]uint),

		GanttMap_Staged_Order: make(map[*Gantt]uint),

		GroupMap_Staged_Order: make(map[*Group]uint),

		LaneMap_Staged_Order: make(map[*Lane]uint),

		LaneUseMap_Staged_Order: make(map[*LaneUse]uint),

		MilestoneMap_Staged_Order: make(map[*Milestone]uint),

		// end of insertion point

		NamedStructs: []*NamedStruct{ // insertion point for order map initialisations
			{name: "Arrow"},
			{name: "Bar"},
			{name: "Gantt"},
			{name: "Group"},
			{name: "Lane"},
			{name: "LaneUse"},
			{name: "Milestone"},
		}, // end of insertion point
	}

	return
}

func GetOrder[Type Gongstruct](stage *Stage, instance *Type) uint {

	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *Arrow:
		return stage.ArrowMap_Staged_Order[instance]
	case *Bar:
		return stage.BarMap_Staged_Order[instance]
	case *Gantt:
		return stage.GanttMap_Staged_Order[instance]
	case *Group:
		return stage.GroupMap_Staged_Order[instance]
	case *Lane:
		return stage.LaneMap_Staged_Order[instance]
	case *LaneUse:
		return stage.LaneUseMap_Staged_Order[instance]
	case *Milestone:
		return stage.MilestoneMap_Staged_Order[instance]
	default:
		return 0 // should not happen
	}
}

func GetOrderPointerGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) uint {

	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *Arrow:
		return stage.ArrowMap_Staged_Order[instance]
	case *Bar:
		return stage.BarMap_Staged_Order[instance]
	case *Gantt:
		return stage.GanttMap_Staged_Order[instance]
	case *Group:
		return stage.GroupMap_Staged_Order[instance]
	case *Lane:
		return stage.LaneMap_Staged_Order[instance]
	case *LaneUse:
		return stage.LaneUseMap_Staged_Order[instance]
	case *Milestone:
		return stage.MilestoneMap_Staged_Order[instance]
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
	stage.ComputeDifference()
	stage.ComputeReference()

	if stage.GetProbeIF() != nil {
		stage.GetProbeIF().AddNotification(time.Now(), "\t// Commit performed")
		stage.GetProbeIF().CommitNotificationTable()
	}
}

func (stage *Stage) ComputeInstancesNb() {
	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["Arrow"] = len(stage.Arrows)
	stage.Map_GongStructName_InstancesNb["Bar"] = len(stage.Bars)
	stage.Map_GongStructName_InstancesNb["Gantt"] = len(stage.Gantts)
	stage.Map_GongStructName_InstancesNb["Group"] = len(stage.Groups)
	stage.Map_GongStructName_InstancesNb["Lane"] = len(stage.Lanes)
	stage.Map_GongStructName_InstancesNb["LaneUse"] = len(stage.LaneUses)
	stage.Map_GongStructName_InstancesNb["Milestone"] = len(stage.Milestones)
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
// Stage puts arrow to the model stage
func (arrow *Arrow) Stage(stage *Stage) *Arrow {

	if _, ok := stage.Arrows[arrow]; !ok {
		stage.Arrows[arrow] = struct{}{}
		stage.ArrowMap_Staged_Order[arrow] = stage.ArrowOrder
		stage.ArrowOrder++
	}
	stage.Arrows_mapString[arrow.Name] = arrow

	return arrow
}

// StagePreserveOrder puts arrow to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ArrowOrder
// - update stage.ArrowOrder accordingly
func (arrow *Arrow) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.Arrows[arrow]; !ok {
		stage.Arrows[arrow] = struct{}{}

		if order > stage.ArrowOrder {
			stage.ArrowOrder = order
		}
		stage.ArrowMap_Staged_Order[arrow] = stage.ArrowOrder
		stage.ArrowOrder++
	}
	stage.Arrows_mapString[arrow.Name] = arrow
}

// Unstage removes arrow off the model stage
func (arrow *Arrow) Unstage(stage *Stage) *Arrow {
	delete(stage.Arrows, arrow)
	delete(stage.ArrowMap_Staged_Order, arrow)
	delete(stage.Arrows_mapString, arrow.Name)

	return arrow
}

// UnstageVoid removes arrow off the model stage
func (arrow *Arrow) UnstageVoid(stage *Stage) {
	delete(stage.Arrows, arrow)
	delete(stage.ArrowMap_Staged_Order, arrow)
	delete(stage.Arrows_mapString, arrow.Name)
}

// commit arrow to the back repo (if it is already staged)
func (arrow *Arrow) Commit(stage *Stage) *Arrow {
	if _, ok := stage.Arrows[arrow]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitArrow(arrow)
		}
	}
	return arrow
}

func (arrow *Arrow) CommitVoid(stage *Stage) {
	arrow.Commit(stage)
}

func (arrow *Arrow) StageVoid(stage *Stage) {
	arrow.Stage(stage)
}

// Checkout arrow to the back repo (if it is already staged)
func (arrow *Arrow) Checkout(stage *Stage) *Arrow {
	if _, ok := stage.Arrows[arrow]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutArrow(arrow)
		}
	}
	return arrow
}

// for satisfaction of GongStruct interface
func (arrow *Arrow) GetName() (res string) {
	return arrow.Name
}

// for satisfaction of GongStruct interface
func (arrow *Arrow) SetName(name string) {
	arrow.Name = name
}

// Stage puts bar to the model stage
func (bar *Bar) Stage(stage *Stage) *Bar {

	if _, ok := stage.Bars[bar]; !ok {
		stage.Bars[bar] = struct{}{}
		stage.BarMap_Staged_Order[bar] = stage.BarOrder
		stage.BarOrder++
	}
	stage.Bars_mapString[bar.Name] = bar

	return bar
}

// StagePreserveOrder puts bar to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.BarOrder
// - update stage.BarOrder accordingly
func (bar *Bar) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.Bars[bar]; !ok {
		stage.Bars[bar] = struct{}{}

		if order > stage.BarOrder {
			stage.BarOrder = order
		}
		stage.BarMap_Staged_Order[bar] = stage.BarOrder
		stage.BarOrder++
	}
	stage.Bars_mapString[bar.Name] = bar
}

// Unstage removes bar off the model stage
func (bar *Bar) Unstage(stage *Stage) *Bar {
	delete(stage.Bars, bar)
	delete(stage.BarMap_Staged_Order, bar)
	delete(stage.Bars_mapString, bar.Name)

	return bar
}

// UnstageVoid removes bar off the model stage
func (bar *Bar) UnstageVoid(stage *Stage) {
	delete(stage.Bars, bar)
	delete(stage.BarMap_Staged_Order, bar)
	delete(stage.Bars_mapString, bar.Name)
}

// commit bar to the back repo (if it is already staged)
func (bar *Bar) Commit(stage *Stage) *Bar {
	if _, ok := stage.Bars[bar]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitBar(bar)
		}
	}
	return bar
}

func (bar *Bar) CommitVoid(stage *Stage) {
	bar.Commit(stage)
}

func (bar *Bar) StageVoid(stage *Stage) {
	bar.Stage(stage)
}

// Checkout bar to the back repo (if it is already staged)
func (bar *Bar) Checkout(stage *Stage) *Bar {
	if _, ok := stage.Bars[bar]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutBar(bar)
		}
	}
	return bar
}

// for satisfaction of GongStruct interface
func (bar *Bar) GetName() (res string) {
	return bar.Name
}

// for satisfaction of GongStruct interface
func (bar *Bar) SetName(name string) {
	bar.Name = name
}

// Stage puts gantt to the model stage
func (gantt *Gantt) Stage(stage *Stage) *Gantt {

	if _, ok := stage.Gantts[gantt]; !ok {
		stage.Gantts[gantt] = struct{}{}
		stage.GanttMap_Staged_Order[gantt] = stage.GanttOrder
		stage.GanttOrder++
	}
	stage.Gantts_mapString[gantt.Name] = gantt

	return gantt
}

// StagePreserveOrder puts gantt to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.GanttOrder
// - update stage.GanttOrder accordingly
func (gantt *Gantt) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.Gantts[gantt]; !ok {
		stage.Gantts[gantt] = struct{}{}

		if order > stage.GanttOrder {
			stage.GanttOrder = order
		}
		stage.GanttMap_Staged_Order[gantt] = stage.GanttOrder
		stage.GanttOrder++
	}
	stage.Gantts_mapString[gantt.Name] = gantt
}

// Unstage removes gantt off the model stage
func (gantt *Gantt) Unstage(stage *Stage) *Gantt {
	delete(stage.Gantts, gantt)
	delete(stage.GanttMap_Staged_Order, gantt)
	delete(stage.Gantts_mapString, gantt.Name)

	return gantt
}

// UnstageVoid removes gantt off the model stage
func (gantt *Gantt) UnstageVoid(stage *Stage) {
	delete(stage.Gantts, gantt)
	delete(stage.GanttMap_Staged_Order, gantt)
	delete(stage.Gantts_mapString, gantt.Name)
}

// commit gantt to the back repo (if it is already staged)
func (gantt *Gantt) Commit(stage *Stage) *Gantt {
	if _, ok := stage.Gantts[gantt]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitGantt(gantt)
		}
	}
	return gantt
}

func (gantt *Gantt) CommitVoid(stage *Stage) {
	gantt.Commit(stage)
}

func (gantt *Gantt) StageVoid(stage *Stage) {
	gantt.Stage(stage)
}

// Checkout gantt to the back repo (if it is already staged)
func (gantt *Gantt) Checkout(stage *Stage) *Gantt {
	if _, ok := stage.Gantts[gantt]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutGantt(gantt)
		}
	}
	return gantt
}

// for satisfaction of GongStruct interface
func (gantt *Gantt) GetName() (res string) {
	return gantt.Name
}

// for satisfaction of GongStruct interface
func (gantt *Gantt) SetName(name string) {
	gantt.Name = name
}

// Stage puts group to the model stage
func (group *Group) Stage(stage *Stage) *Group {

	if _, ok := stage.Groups[group]; !ok {
		stage.Groups[group] = struct{}{}
		stage.GroupMap_Staged_Order[group] = stage.GroupOrder
		stage.GroupOrder++
	}
	stage.Groups_mapString[group.Name] = group

	return group
}

// StagePreserveOrder puts group to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.GroupOrder
// - update stage.GroupOrder accordingly
func (group *Group) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.Groups[group]; !ok {
		stage.Groups[group] = struct{}{}

		if order > stage.GroupOrder {
			stage.GroupOrder = order
		}
		stage.GroupMap_Staged_Order[group] = stage.GroupOrder
		stage.GroupOrder++
	}
	stage.Groups_mapString[group.Name] = group
}

// Unstage removes group off the model stage
func (group *Group) Unstage(stage *Stage) *Group {
	delete(stage.Groups, group)
	delete(stage.GroupMap_Staged_Order, group)
	delete(stage.Groups_mapString, group.Name)

	return group
}

// UnstageVoid removes group off the model stage
func (group *Group) UnstageVoid(stage *Stage) {
	delete(stage.Groups, group)
	delete(stage.GroupMap_Staged_Order, group)
	delete(stage.Groups_mapString, group.Name)
}

// commit group to the back repo (if it is already staged)
func (group *Group) Commit(stage *Stage) *Group {
	if _, ok := stage.Groups[group]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitGroup(group)
		}
	}
	return group
}

func (group *Group) CommitVoid(stage *Stage) {
	group.Commit(stage)
}

func (group *Group) StageVoid(stage *Stage) {
	group.Stage(stage)
}

// Checkout group to the back repo (if it is already staged)
func (group *Group) Checkout(stage *Stage) *Group {
	if _, ok := stage.Groups[group]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutGroup(group)
		}
	}
	return group
}

// for satisfaction of GongStruct interface
func (group *Group) GetName() (res string) {
	return group.Name
}

// for satisfaction of GongStruct interface
func (group *Group) SetName(name string) {
	group.Name = name
}

// Stage puts lane to the model stage
func (lane *Lane) Stage(stage *Stage) *Lane {

	if _, ok := stage.Lanes[lane]; !ok {
		stage.Lanes[lane] = struct{}{}
		stage.LaneMap_Staged_Order[lane] = stage.LaneOrder
		stage.LaneOrder++
	}
	stage.Lanes_mapString[lane.Name] = lane

	return lane
}

// StagePreserveOrder puts lane to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.LaneOrder
// - update stage.LaneOrder accordingly
func (lane *Lane) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.Lanes[lane]; !ok {
		stage.Lanes[lane] = struct{}{}

		if order > stage.LaneOrder {
			stage.LaneOrder = order
		}
		stage.LaneMap_Staged_Order[lane] = stage.LaneOrder
		stage.LaneOrder++
	}
	stage.Lanes_mapString[lane.Name] = lane
}

// Unstage removes lane off the model stage
func (lane *Lane) Unstage(stage *Stage) *Lane {
	delete(stage.Lanes, lane)
	delete(stage.LaneMap_Staged_Order, lane)
	delete(stage.Lanes_mapString, lane.Name)

	return lane
}

// UnstageVoid removes lane off the model stage
func (lane *Lane) UnstageVoid(stage *Stage) {
	delete(stage.Lanes, lane)
	delete(stage.LaneMap_Staged_Order, lane)
	delete(stage.Lanes_mapString, lane.Name)
}

// commit lane to the back repo (if it is already staged)
func (lane *Lane) Commit(stage *Stage) *Lane {
	if _, ok := stage.Lanes[lane]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitLane(lane)
		}
	}
	return lane
}

func (lane *Lane) CommitVoid(stage *Stage) {
	lane.Commit(stage)
}

func (lane *Lane) StageVoid(stage *Stage) {
	lane.Stage(stage)
}

// Checkout lane to the back repo (if it is already staged)
func (lane *Lane) Checkout(stage *Stage) *Lane {
	if _, ok := stage.Lanes[lane]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutLane(lane)
		}
	}
	return lane
}

// for satisfaction of GongStruct interface
func (lane *Lane) GetName() (res string) {
	return lane.Name
}

// for satisfaction of GongStruct interface
func (lane *Lane) SetName(name string) {
	lane.Name = name
}

// Stage puts laneuse to the model stage
func (laneuse *LaneUse) Stage(stage *Stage) *LaneUse {

	if _, ok := stage.LaneUses[laneuse]; !ok {
		stage.LaneUses[laneuse] = struct{}{}
		stage.LaneUseMap_Staged_Order[laneuse] = stage.LaneUseOrder
		stage.LaneUseOrder++
	}
	stage.LaneUses_mapString[laneuse.Name] = laneuse

	return laneuse
}

// StagePreserveOrder puts laneuse to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.LaneUseOrder
// - update stage.LaneUseOrder accordingly
func (laneuse *LaneUse) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.LaneUses[laneuse]; !ok {
		stage.LaneUses[laneuse] = struct{}{}

		if order > stage.LaneUseOrder {
			stage.LaneUseOrder = order
		}
		stage.LaneUseMap_Staged_Order[laneuse] = stage.LaneUseOrder
		stage.LaneUseOrder++
	}
	stage.LaneUses_mapString[laneuse.Name] = laneuse
}

// Unstage removes laneuse off the model stage
func (laneuse *LaneUse) Unstage(stage *Stage) *LaneUse {
	delete(stage.LaneUses, laneuse)
	delete(stage.LaneUseMap_Staged_Order, laneuse)
	delete(stage.LaneUses_mapString, laneuse.Name)

	return laneuse
}

// UnstageVoid removes laneuse off the model stage
func (laneuse *LaneUse) UnstageVoid(stage *Stage) {
	delete(stage.LaneUses, laneuse)
	delete(stage.LaneUseMap_Staged_Order, laneuse)
	delete(stage.LaneUses_mapString, laneuse.Name)
}

// commit laneuse to the back repo (if it is already staged)
func (laneuse *LaneUse) Commit(stage *Stage) *LaneUse {
	if _, ok := stage.LaneUses[laneuse]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitLaneUse(laneuse)
		}
	}
	return laneuse
}

func (laneuse *LaneUse) CommitVoid(stage *Stage) {
	laneuse.Commit(stage)
}

func (laneuse *LaneUse) StageVoid(stage *Stage) {
	laneuse.Stage(stage)
}

// Checkout laneuse to the back repo (if it is already staged)
func (laneuse *LaneUse) Checkout(stage *Stage) *LaneUse {
	if _, ok := stage.LaneUses[laneuse]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutLaneUse(laneuse)
		}
	}
	return laneuse
}

// for satisfaction of GongStruct interface
func (laneuse *LaneUse) GetName() (res string) {
	return laneuse.Name
}

// for satisfaction of GongStruct interface
func (laneuse *LaneUse) SetName(name string) {
	laneuse.Name = name
}

// Stage puts milestone to the model stage
func (milestone *Milestone) Stage(stage *Stage) *Milestone {

	if _, ok := stage.Milestones[milestone]; !ok {
		stage.Milestones[milestone] = struct{}{}
		stage.MilestoneMap_Staged_Order[milestone] = stage.MilestoneOrder
		stage.MilestoneOrder++
	}
	stage.Milestones_mapString[milestone.Name] = milestone

	return milestone
}

// StagePreserveOrder puts milestone to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MilestoneOrder
// - update stage.MilestoneOrder accordingly
func (milestone *Milestone) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.Milestones[milestone]; !ok {
		stage.Milestones[milestone] = struct{}{}

		if order > stage.MilestoneOrder {
			stage.MilestoneOrder = order
		}
		stage.MilestoneMap_Staged_Order[milestone] = stage.MilestoneOrder
		stage.MilestoneOrder++
	}
	stage.Milestones_mapString[milestone.Name] = milestone
}

// Unstage removes milestone off the model stage
func (milestone *Milestone) Unstage(stage *Stage) *Milestone {
	delete(stage.Milestones, milestone)
	delete(stage.MilestoneMap_Staged_Order, milestone)
	delete(stage.Milestones_mapString, milestone.Name)

	return milestone
}

// UnstageVoid removes milestone off the model stage
func (milestone *Milestone) UnstageVoid(stage *Stage) {
	delete(stage.Milestones, milestone)
	delete(stage.MilestoneMap_Staged_Order, milestone)
	delete(stage.Milestones_mapString, milestone.Name)
}

// commit milestone to the back repo (if it is already staged)
func (milestone *Milestone) Commit(stage *Stage) *Milestone {
	if _, ok := stage.Milestones[milestone]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitMilestone(milestone)
		}
	}
	return milestone
}

func (milestone *Milestone) CommitVoid(stage *Stage) {
	milestone.Commit(stage)
}

func (milestone *Milestone) StageVoid(stage *Stage) {
	milestone.Stage(stage)
}

// Checkout milestone to the back repo (if it is already staged)
func (milestone *Milestone) Checkout(stage *Stage) *Milestone {
	if _, ok := stage.Milestones[milestone]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutMilestone(milestone)
		}
	}
	return milestone
}

// for satisfaction of GongStruct interface
func (milestone *Milestone) GetName() (res string) {
	return milestone.Name
}

// for satisfaction of GongStruct interface
func (milestone *Milestone) SetName(name string) {
	milestone.Name = name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMArrow(Arrow *Arrow)
	CreateORMBar(Bar *Bar)
	CreateORMGantt(Gantt *Gantt)
	CreateORMGroup(Group *Group)
	CreateORMLane(Lane *Lane)
	CreateORMLaneUse(LaneUse *LaneUse)
	CreateORMMilestone(Milestone *Milestone)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMArrow(Arrow *Arrow)
	DeleteORMBar(Bar *Bar)
	DeleteORMGantt(Gantt *Gantt)
	DeleteORMGroup(Group *Group)
	DeleteORMLane(Lane *Lane)
	DeleteORMLaneUse(LaneUse *LaneUse)
	DeleteORMMilestone(Milestone *Milestone)
}

func (stage *Stage) Reset() { // insertion point for array reset
	stage.Arrows = make(map[*Arrow]struct{})
	stage.Arrows_mapString = make(map[string]*Arrow)
	stage.ArrowMap_Staged_Order = make(map[*Arrow]uint)
	stage.ArrowOrder = 0

	stage.Bars = make(map[*Bar]struct{})
	stage.Bars_mapString = make(map[string]*Bar)
	stage.BarMap_Staged_Order = make(map[*Bar]uint)
	stage.BarOrder = 0

	stage.Gantts = make(map[*Gantt]struct{})
	stage.Gantts_mapString = make(map[string]*Gantt)
	stage.GanttMap_Staged_Order = make(map[*Gantt]uint)
	stage.GanttOrder = 0

	stage.Groups = make(map[*Group]struct{})
	stage.Groups_mapString = make(map[string]*Group)
	stage.GroupMap_Staged_Order = make(map[*Group]uint)
	stage.GroupOrder = 0

	stage.Lanes = make(map[*Lane]struct{})
	stage.Lanes_mapString = make(map[string]*Lane)
	stage.LaneMap_Staged_Order = make(map[*Lane]uint)
	stage.LaneOrder = 0

	stage.LaneUses = make(map[*LaneUse]struct{})
	stage.LaneUses_mapString = make(map[string]*LaneUse)
	stage.LaneUseMap_Staged_Order = make(map[*LaneUse]uint)
	stage.LaneUseOrder = 0

	stage.Milestones = make(map[*Milestone]struct{})
	stage.Milestones_mapString = make(map[string]*Milestone)
	stage.MilestoneMap_Staged_Order = make(map[*Milestone]uint)
	stage.MilestoneOrder = 0

	stage.ComputeReference()
}

func (stage *Stage) Nil() { // insertion point for array nil
	stage.Arrows = nil
	stage.Arrows_mapString = nil

	stage.Bars = nil
	stage.Bars_mapString = nil

	stage.Gantts = nil
	stage.Gantts_mapString = nil

	stage.Groups = nil
	stage.Groups_mapString = nil

	stage.Lanes = nil
	stage.Lanes_mapString = nil

	stage.LaneUses = nil
	stage.LaneUses_mapString = nil

	stage.Milestones = nil
	stage.Milestones_mapString = nil

}

func (stage *Stage) Unstage() { // insertion point for array nil
	for arrow := range stage.Arrows {
		arrow.Unstage(stage)
	}

	for bar := range stage.Bars {
		bar.Unstage(stage)
	}

	for gantt := range stage.Gantts {
		gantt.Unstage(stage)
	}

	for group := range stage.Groups {
		group.Unstage(stage)
	}

	for lane := range stage.Lanes {
		lane.Unstage(stage)
	}

	for laneuse := range stage.LaneUses {
		laneuse.Unstage(stage)
	}

	for milestone := range stage.Milestones {
		milestone.Unstage(stage)
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
	GongClean(stage *Stage)
	GongGetFieldValue(fieldName string, stage *Stage) GongFieldValue
	GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error
	GongGetGongstructName() string
	GongGetOrder(stage *Stage) uint
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
	case map[*Arrow]any:
		return any(&stage.Arrows).(*Type)
	case map[*Bar]any:
		return any(&stage.Bars).(*Type)
	case map[*Gantt]any:
		return any(&stage.Gantts).(*Type)
	case map[*Group]any:
		return any(&stage.Groups).(*Type)
	case map[*Lane]any:
		return any(&stage.Lanes).(*Type)
	case map[*LaneUse]any:
		return any(&stage.LaneUses).(*Type)
	case map[*Milestone]any:
		return any(&stage.Milestones).(*Type)
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
	case *Arrow:
		return any(stage.Arrows_mapString).(map[string]Type)
	case *Bar:
		return any(stage.Bars_mapString).(map[string]Type)
	case *Gantt:
		return any(stage.Gantts_mapString).(map[string]Type)
	case *Group:
		return any(stage.Groups_mapString).(map[string]Type)
	case *Lane:
		return any(stage.Lanes_mapString).(map[string]Type)
	case *LaneUse:
		return any(stage.LaneUses_mapString).(map[string]Type)
	case *Milestone:
		return any(stage.Milestones_mapString).(map[string]Type)
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
	case Arrow:
		return any(&stage.Arrows).(*map[*Type]struct{})
	case Bar:
		return any(&stage.Bars).(*map[*Type]struct{})
	case Gantt:
		return any(&stage.Gantts).(*map[*Type]struct{})
	case Group:
		return any(&stage.Groups).(*map[*Type]struct{})
	case Lane:
		return any(&stage.Lanes).(*map[*Type]struct{})
	case LaneUse:
		return any(&stage.LaneUses).(*map[*Type]struct{})
	case Milestone:
		return any(&stage.Milestones).(*map[*Type]struct{})
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
	case *Arrow:
		return any(&stage.Arrows).(*map[Type]struct{})
	case *Bar:
		return any(&stage.Bars).(*map[Type]struct{})
	case *Gantt:
		return any(&stage.Gantts).(*map[Type]struct{})
	case *Group:
		return any(&stage.Groups).(*map[Type]struct{})
	case *Lane:
		return any(&stage.Lanes).(*map[Type]struct{})
	case *LaneUse:
		return any(&stage.LaneUses).(*map[Type]struct{})
	case *Milestone:
		return any(&stage.Milestones).(*map[Type]struct{})
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
	case Arrow:
		return any(&stage.Arrows_mapString).(*map[string]*Type)
	case Bar:
		return any(&stage.Bars_mapString).(*map[string]*Type)
	case Gantt:
		return any(&stage.Gantts_mapString).(*map[string]*Type)
	case Group:
		return any(&stage.Groups_mapString).(*map[string]*Type)
	case Lane:
		return any(&stage.Lanes_mapString).(*map[string]*Type)
	case LaneUse:
		return any(&stage.LaneUses_mapString).(*map[string]*Type)
	case Milestone:
		return any(&stage.Milestones_mapString).(*map[string]*Type)
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
	case Arrow:
		return any(&Arrow{
			// Initialisation of associations
			// field is initialized with an instance of Bar with the name of the field
			From: &Bar{Name: "From"},
			// field is initialized with an instance of Bar with the name of the field
			To: &Bar{Name: "To"},
		}).(*Type)
	case Bar:
		return any(&Bar{
			// Initialisation of associations
		}).(*Type)
	case Gantt:
		return any(&Gantt{
			// Initialisation of associations
			// field is initialized with an instance of Lane with the name of the field
			Lanes: []*Lane{{Name: "Lanes"}},
			// field is initialized with an instance of Milestone with the name of the field
			Milestones: []*Milestone{{Name: "Milestones"}},
			// field is initialized with an instance of Group with the name of the field
			Groups: []*Group{{Name: "Groups"}},
			// field is initialized with an instance of Arrow with the name of the field
			Arrows: []*Arrow{{Name: "Arrows"}},
		}).(*Type)
	case Group:
		return any(&Group{
			// Initialisation of associations
			// field is initialized with an instance of Lane with the name of the field
			GroupLanes: []*Lane{{Name: "GroupLanes"}},
		}).(*Type)
	case Lane:
		return any(&Lane{
			// Initialisation of associations
			// field is initialized with an instance of Bar with the name of the field
			Bars: []*Bar{{Name: "Bars"}},
		}).(*Type)
	case LaneUse:
		return any(&LaneUse{
			// Initialisation of associations
			// field is initialized with an instance of Lane with the name of the field
			Lane: &Lane{Name: "Lane"},
		}).(*Type)
	case Milestone:
		return any(&Milestone{
			// Initialisation of associations
			// field is initialized with an instance of Lane with the name of the field
			LanesToDisplay: []*Lane{{Name: "LanesToDisplay"}},
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
	// reverse maps of direct associations of Arrow
	case Arrow:
		switch fieldname {
		// insertion point for per direct association field
		case "From":
			res := make(map[*Bar][]*Arrow)
			for arrow := range stage.Arrows {
				if arrow.From != nil {
					bar_ := arrow.From
					var arrows []*Arrow
					_, ok := res[bar_]
					if ok {
						arrows = res[bar_]
					} else {
						arrows = make([]*Arrow, 0)
					}
					arrows = append(arrows, arrow)
					res[bar_] = arrows
				}
			}
			return any(res).(map[*End][]*Start)
		case "To":
			res := make(map[*Bar][]*Arrow)
			for arrow := range stage.Arrows {
				if arrow.To != nil {
					bar_ := arrow.To
					var arrows []*Arrow
					_, ok := res[bar_]
					if ok {
						arrows = res[bar_]
					} else {
						arrows = make([]*Arrow, 0)
					}
					arrows = append(arrows, arrow)
					res[bar_] = arrows
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Bar
	case Bar:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Gantt
	case Gantt:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Group
	case Group:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Lane
	case Lane:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of LaneUse
	case LaneUse:
		switch fieldname {
		// insertion point for per direct association field
		case "Lane":
			res := make(map[*Lane][]*LaneUse)
			for laneuse := range stage.LaneUses {
				if laneuse.Lane != nil {
					lane_ := laneuse.Lane
					var laneuses []*LaneUse
					_, ok := res[lane_]
					if ok {
						laneuses = res[lane_]
					} else {
						laneuses = make([]*LaneUse, 0)
					}
					laneuses = append(laneuses, laneuse)
					res[lane_] = laneuses
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Milestone
	case Milestone:
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
	// reverse maps of direct associations of Arrow
	case Arrow:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Bar
	case Bar:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Gantt
	case Gantt:
		switch fieldname {
		// insertion point for per direct association field
		case "Lanes":
			res := make(map[*Lane][]*Gantt)
			for gantt := range stage.Gantts {
				for _, lane_ := range gantt.Lanes {
					res[lane_] = append(res[lane_], gantt)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Milestones":
			res := make(map[*Milestone][]*Gantt)
			for gantt := range stage.Gantts {
				for _, milestone_ := range gantt.Milestones {
					res[milestone_] = append(res[milestone_], gantt)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Groups":
			res := make(map[*Group][]*Gantt)
			for gantt := range stage.Gantts {
				for _, group_ := range gantt.Groups {
					res[group_] = append(res[group_], gantt)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Arrows":
			res := make(map[*Arrow][]*Gantt)
			for gantt := range stage.Gantts {
				for _, arrow_ := range gantt.Arrows {
					res[arrow_] = append(res[arrow_], gantt)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Group
	case Group:
		switch fieldname {
		// insertion point for per direct association field
		case "GroupLanes":
			res := make(map[*Lane][]*Group)
			for group := range stage.Groups {
				for _, lane_ := range group.GroupLanes {
					res[lane_] = append(res[lane_], group)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Lane
	case Lane:
		switch fieldname {
		// insertion point for per direct association field
		case "Bars":
			res := make(map[*Bar][]*Lane)
			for lane := range stage.Lanes {
				for _, bar_ := range lane.Bars {
					res[bar_] = append(res[bar_], lane)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of LaneUse
	case LaneUse:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Milestone
	case Milestone:
		switch fieldname {
		// insertion point for per direct association field
		case "LanesToDisplay":
			res := make(map[*Lane][]*Milestone)
			for milestone := range stage.Milestones {
				for _, lane_ := range milestone.LanesToDisplay {
					res[lane_] = append(res[lane_], milestone)
				}
			}
			return any(res).(map[*End][]*Start)
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
	case *Arrow:
		res = "Arrow"
	case *Bar:
		res = "Bar"
	case *Gantt:
		res = "Gantt"
	case *Group:
		res = "Group"
	case *Lane:
		res = "Lane"
	case *LaneUse:
		res = "LaneUse"
	case *Milestone:
		res = "Milestone"
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
	case *Arrow:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Gantt"
		rf.Fieldname = "Arrows"
		res = append(res, rf)
	case *Bar:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Lane"
		rf.Fieldname = "Bars"
		res = append(res, rf)
	case *Gantt:
		var rf ReverseField
		_ = rf
	case *Group:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Gantt"
		rf.Fieldname = "Groups"
		res = append(res, rf)
	case *Lane:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Gantt"
		rf.Fieldname = "Lanes"
		res = append(res, rf)
		rf.GongstructName = "Group"
		rf.Fieldname = "GroupLanes"
		res = append(res, rf)
		rf.GongstructName = "Milestone"
		rf.Fieldname = "LanesToDisplay"
		res = append(res, rf)
	case *LaneUse:
		var rf ReverseField
		_ = rf
	case *Milestone:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Gantt"
		rf.Fieldname = "Milestones"
		res = append(res, rf)
	}
	return
}

// insertion point for get fields header method
func (arrow *Arrow) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "From",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Bar",
		},
		{
			Name:                 "To",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Bar",
		},
		{
			Name:               "OptionnalColor",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "OptionnalStroke",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (bar *Bar) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Start",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "End",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "ComputedDuration",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "OptionnalColor",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "OptionnalStroke",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "FillOpacity",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StrokeWidth",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StrokeDashArray",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (gantt *Gantt) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "ComputedStart",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "ComputedEnd",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "ComputedDuration",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "UseManualStartAndEndDates",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "ManualStart",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "ManualEnd",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "LaneHeight",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "RatioBarToLaneHeight",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "YTopMargin",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "XLeftText",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "TextHeight",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "XLeftLanes",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "XRightMargin",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "ArrowLengthToTheRightOfStartBar",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "ArrowTipLenght",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "TimeLine_Color",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "TimeLine_FillOpacity",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "TimeLine_Stroke",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "TimeLine_StrokeWidth",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Group_Stroke",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Group_StrokeWidth",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Group_StrokeDashArray",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "DateYOffset",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "AlignOnStartEndOnYearStart",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "Lanes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Lane",
		},
		{
			Name:                 "Milestones",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Milestone",
		},
		{
			Name:                 "Groups",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Group",
		},
		{
			Name:                 "Arrows",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Arrow",
		},
	}
	return
}

func (group *Group) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "GroupLanes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Lane",
		},
	}
	return
}

func (lane *Lane) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Order",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "Bars",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Bar",
		},
	}
	return
}

func (laneuse *LaneUse) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "Lane",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Lane",
		},
	}
	return
}

func (milestone *Milestone) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Date",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "DisplayVerticalBar",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "LanesToDisplay",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Lane",
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
func (arrow *Arrow) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = arrow.Name
	case "From":
		res.GongFieldValueType = GongFieldValueTypePointer
		if arrow.From != nil {
			res.valueString = arrow.From.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, arrow.From))
		}
	case "To":
		res.GongFieldValueType = GongFieldValueTypePointer
		if arrow.To != nil {
			res.valueString = arrow.To.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, arrow.To))
		}
	case "OptionnalColor":
		res.valueString = arrow.OptionnalColor
	case "OptionnalStroke":
		res.valueString = arrow.OptionnalStroke
	}
	return
}
func (bar *Bar) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = bar.Name
	case "Start":
		res.valueString = bar.Start.String()
	case "End":
		res.valueString = bar.End.String()
	case "ComputedDuration":
		if math.Abs(bar.ComputedDuration.Hours()) >= 24 {
			days := __Gong__Abs(int(int(bar.ComputedDuration.Hours()) / 24))
			months := int(days / 31)
			days = days - months*31

			remainingHours := int(bar.ComputedDuration.Hours()) % 24
			remainingMinutes := int(bar.ComputedDuration.Minutes()) % 60
			remainingSeconds := int(bar.ComputedDuration.Seconds()) % 60

			if bar.ComputedDuration.Hours() < 0 {
				res.valueString = "- "
			}

			if months > 0 {
				if months > 1 {
					res.valueString = res.valueString + fmt.Sprintf("%d months", months)
				} else {
					res.valueString = res.valueString + fmt.Sprintf("%d month", months)
				}
			}
			if days > 0 {
				if months != 0 {
					res.valueString = res.valueString + ", "
				}
				if days > 1 {
					res.valueString = res.valueString + fmt.Sprintf("%d days", days)
				} else {
					res.valueString = res.valueString + fmt.Sprintf("%d day", days)
				}

			}
			if remainingHours != 0 || remainingMinutes != 0 || remainingSeconds != 0 {
				if days != 0 || (days == 0 && months != 0) {
					res.valueString = res.valueString + ", "
				}
				res.valueString = res.valueString + fmt.Sprintf("%d hours, %d minutes, %d seconds\n", remainingHours, remainingMinutes, remainingSeconds)
			}
		} else {
			res.valueString = fmt.Sprintf("%s\n", bar.ComputedDuration.String())
		}
	case "OptionnalColor":
		res.valueString = bar.OptionnalColor
	case "OptionnalStroke":
		res.valueString = bar.OptionnalStroke
	case "FillOpacity":
		res.valueString = fmt.Sprintf("%f", bar.FillOpacity)
		res.valueFloat = bar.FillOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeWidth":
		res.valueString = fmt.Sprintf("%f", bar.StrokeWidth)
		res.valueFloat = bar.StrokeWidth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeDashArray":
		res.valueString = bar.StrokeDashArray
	}
	return
}
func (gantt *Gantt) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = gantt.Name
	case "ComputedStart":
		res.valueString = gantt.ComputedStart.String()
	case "ComputedEnd":
		res.valueString = gantt.ComputedEnd.String()
	case "ComputedDuration":
		if math.Abs(gantt.ComputedDuration.Hours()) >= 24 {
			days := __Gong__Abs(int(int(gantt.ComputedDuration.Hours()) / 24))
			months := int(days / 31)
			days = days - months*31

			remainingHours := int(gantt.ComputedDuration.Hours()) % 24
			remainingMinutes := int(gantt.ComputedDuration.Minutes()) % 60
			remainingSeconds := int(gantt.ComputedDuration.Seconds()) % 60

			if gantt.ComputedDuration.Hours() < 0 {
				res.valueString = "- "
			}

			if months > 0 {
				if months > 1 {
					res.valueString = res.valueString + fmt.Sprintf("%d months", months)
				} else {
					res.valueString = res.valueString + fmt.Sprintf("%d month", months)
				}
			}
			if days > 0 {
				if months != 0 {
					res.valueString = res.valueString + ", "
				}
				if days > 1 {
					res.valueString = res.valueString + fmt.Sprintf("%d days", days)
				} else {
					res.valueString = res.valueString + fmt.Sprintf("%d day", days)
				}

			}
			if remainingHours != 0 || remainingMinutes != 0 || remainingSeconds != 0 {
				if days != 0 || (days == 0 && months != 0) {
					res.valueString = res.valueString + ", "
				}
				res.valueString = res.valueString + fmt.Sprintf("%d hours, %d minutes, %d seconds\n", remainingHours, remainingMinutes, remainingSeconds)
			}
		} else {
			res.valueString = fmt.Sprintf("%s\n", gantt.ComputedDuration.String())
		}
	case "UseManualStartAndEndDates":
		res.valueString = fmt.Sprintf("%t", gantt.UseManualStartAndEndDates)
		res.valueBool = gantt.UseManualStartAndEndDates
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ManualStart":
		res.valueString = gantt.ManualStart.String()
	case "ManualEnd":
		res.valueString = gantt.ManualEnd.String()
	case "LaneHeight":
		res.valueString = fmt.Sprintf("%f", gantt.LaneHeight)
		res.valueFloat = gantt.LaneHeight
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "RatioBarToLaneHeight":
		res.valueString = fmt.Sprintf("%f", gantt.RatioBarToLaneHeight)
		res.valueFloat = gantt.RatioBarToLaneHeight
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "YTopMargin":
		res.valueString = fmt.Sprintf("%f", gantt.YTopMargin)
		res.valueFloat = gantt.YTopMargin
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "XLeftText":
		res.valueString = fmt.Sprintf("%f", gantt.XLeftText)
		res.valueFloat = gantt.XLeftText
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "TextHeight":
		res.valueString = fmt.Sprintf("%f", gantt.TextHeight)
		res.valueFloat = gantt.TextHeight
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "XLeftLanes":
		res.valueString = fmt.Sprintf("%f", gantt.XLeftLanes)
		res.valueFloat = gantt.XLeftLanes
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "XRightMargin":
		res.valueString = fmt.Sprintf("%f", gantt.XRightMargin)
		res.valueFloat = gantt.XRightMargin
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "ArrowLengthToTheRightOfStartBar":
		res.valueString = fmt.Sprintf("%f", gantt.ArrowLengthToTheRightOfStartBar)
		res.valueFloat = gantt.ArrowLengthToTheRightOfStartBar
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "ArrowTipLenght":
		res.valueString = fmt.Sprintf("%f", gantt.ArrowTipLenght)
		res.valueFloat = gantt.ArrowTipLenght
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "TimeLine_Color":
		res.valueString = gantt.TimeLine_Color
	case "TimeLine_FillOpacity":
		res.valueString = fmt.Sprintf("%f", gantt.TimeLine_FillOpacity)
		res.valueFloat = gantt.TimeLine_FillOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "TimeLine_Stroke":
		res.valueString = gantt.TimeLine_Stroke
	case "TimeLine_StrokeWidth":
		res.valueString = fmt.Sprintf("%f", gantt.TimeLine_StrokeWidth)
		res.valueFloat = gantt.TimeLine_StrokeWidth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Group_Stroke":
		res.valueString = gantt.Group_Stroke
	case "Group_StrokeWidth":
		res.valueString = fmt.Sprintf("%f", gantt.Group_StrokeWidth)
		res.valueFloat = gantt.Group_StrokeWidth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Group_StrokeDashArray":
		res.valueString = gantt.Group_StrokeDashArray
	case "DateYOffset":
		res.valueString = fmt.Sprintf("%f", gantt.DateYOffset)
		res.valueFloat = gantt.DateYOffset
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "AlignOnStartEndOnYearStart":
		res.valueString = fmt.Sprintf("%t", gantt.AlignOnStartEndOnYearStart)
		res.valueBool = gantt.AlignOnStartEndOnYearStart
		res.GongFieldValueType = GongFieldValueTypeBool
	case "Lanes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range gantt.Lanes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "Milestones":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range gantt.Milestones {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "Groups":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range gantt.Groups {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "Arrows":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range gantt.Arrows {
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
func (group *Group) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = group.Name
	case "GroupLanes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range group.GroupLanes {
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
func (lane *Lane) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = lane.Name
	case "Order":
		res.valueString = fmt.Sprintf("%d", lane.Order)
		res.valueInt = lane.Order
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Bars":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range lane.Bars {
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
func (laneuse *LaneUse) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = laneuse.Name
	case "Lane":
		res.GongFieldValueType = GongFieldValueTypePointer
		if laneuse.Lane != nil {
			res.valueString = laneuse.Lane.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, laneuse.Lane))
		}
	}
	return
}
func (milestone *Milestone) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = milestone.Name
	case "Date":
		res.valueString = milestone.Date.String()
	case "DisplayVerticalBar":
		res.valueString = fmt.Sprintf("%t", milestone.DisplayVerticalBar)
		res.valueBool = milestone.DisplayVerticalBar
		res.GongFieldValueType = GongFieldValueTypeBool
	case "LanesToDisplay":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range milestone.LanesToDisplay {
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
func GetFieldStringValueFromPointer(instance GongstructIF, fieldName string, stage *Stage) (res GongFieldValue) {

	res = instance.GongGetFieldValue(fieldName, stage)
	return
}

// insertion point for generic set gongstruct field value
func (arrow *Arrow) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		arrow.Name = value.GetValueString()
	case "From":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			arrow.From = nil
			for __instance__ := range stage.Bars {
				if stage.BarMap_Staged_Order[__instance__] == uint(id) {
					arrow.From = __instance__
					break
				}
			}
		}
	case "To":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			arrow.To = nil
			for __instance__ := range stage.Bars {
				if stage.BarMap_Staged_Order[__instance__] == uint(id) {
					arrow.To = __instance__
					break
				}
			}
		}
	case "OptionnalColor":
		arrow.OptionnalColor = value.GetValueString()
	case "OptionnalStroke":
		arrow.OptionnalStroke = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (bar *Bar) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		bar.Name = value.GetValueString()
	case "OptionnalColor":
		bar.OptionnalColor = value.GetValueString()
	case "OptionnalStroke":
		bar.OptionnalStroke = value.GetValueString()
	case "FillOpacity":
		bar.FillOpacity = value.GetValueFloat()
	case "StrokeWidth":
		bar.StrokeWidth = value.GetValueFloat()
	case "StrokeDashArray":
		bar.StrokeDashArray = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (gantt *Gantt) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		gantt.Name = value.GetValueString()
	case "UseManualStartAndEndDates":
		gantt.UseManualStartAndEndDates = value.GetValueBool()
	case "LaneHeight":
		gantt.LaneHeight = value.GetValueFloat()
	case "RatioBarToLaneHeight":
		gantt.RatioBarToLaneHeight = value.GetValueFloat()
	case "YTopMargin":
		gantt.YTopMargin = value.GetValueFloat()
	case "XLeftText":
		gantt.XLeftText = value.GetValueFloat()
	case "TextHeight":
		gantt.TextHeight = value.GetValueFloat()
	case "XLeftLanes":
		gantt.XLeftLanes = value.GetValueFloat()
	case "XRightMargin":
		gantt.XRightMargin = value.GetValueFloat()
	case "ArrowLengthToTheRightOfStartBar":
		gantt.ArrowLengthToTheRightOfStartBar = value.GetValueFloat()
	case "ArrowTipLenght":
		gantt.ArrowTipLenght = value.GetValueFloat()
	case "TimeLine_Color":
		gantt.TimeLine_Color = value.GetValueString()
	case "TimeLine_FillOpacity":
		gantt.TimeLine_FillOpacity = value.GetValueFloat()
	case "TimeLine_Stroke":
		gantt.TimeLine_Stroke = value.GetValueString()
	case "TimeLine_StrokeWidth":
		gantt.TimeLine_StrokeWidth = value.GetValueFloat()
	case "Group_Stroke":
		gantt.Group_Stroke = value.GetValueString()
	case "Group_StrokeWidth":
		gantt.Group_StrokeWidth = value.GetValueFloat()
	case "Group_StrokeDashArray":
		gantt.Group_StrokeDashArray = value.GetValueString()
	case "DateYOffset":
		gantt.DateYOffset = value.GetValueFloat()
	case "AlignOnStartEndOnYearStart":
		gantt.AlignOnStartEndOnYearStart = value.GetValueBool()
	case "Lanes":
		gantt.Lanes = make([]*Lane, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Lanes {
					if stage.LaneMap_Staged_Order[__instance__] == uint(id) {
						gantt.Lanes = append(gantt.Lanes, __instance__)
						break
					}
				}
			}
		}
	case "Milestones":
		gantt.Milestones = make([]*Milestone, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Milestones {
					if stage.MilestoneMap_Staged_Order[__instance__] == uint(id) {
						gantt.Milestones = append(gantt.Milestones, __instance__)
						break
					}
				}
			}
		}
	case "Groups":
		gantt.Groups = make([]*Group, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Groups {
					if stage.GroupMap_Staged_Order[__instance__] == uint(id) {
						gantt.Groups = append(gantt.Groups, __instance__)
						break
					}
				}
			}
		}
	case "Arrows":
		gantt.Arrows = make([]*Arrow, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Arrows {
					if stage.ArrowMap_Staged_Order[__instance__] == uint(id) {
						gantt.Arrows = append(gantt.Arrows, __instance__)
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

func (group *Group) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		group.Name = value.GetValueString()
	case "GroupLanes":
		group.GroupLanes = make([]*Lane, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Lanes {
					if stage.LaneMap_Staged_Order[__instance__] == uint(id) {
						group.GroupLanes = append(group.GroupLanes, __instance__)
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

func (lane *Lane) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		lane.Name = value.GetValueString()
	case "Order":
		lane.Order = int(value.GetValueInt())
	case "Bars":
		lane.Bars = make([]*Bar, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Bars {
					if stage.BarMap_Staged_Order[__instance__] == uint(id) {
						lane.Bars = append(lane.Bars, __instance__)
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

func (laneuse *LaneUse) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		laneuse.Name = value.GetValueString()
	case "Lane":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			laneuse.Lane = nil
			for __instance__ := range stage.Lanes {
				if stage.LaneMap_Staged_Order[__instance__] == uint(id) {
					laneuse.Lane = __instance__
					break
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (milestone *Milestone) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		milestone.Name = value.GetValueString()
	case "DisplayVerticalBar":
		milestone.DisplayVerticalBar = value.GetValueBool()
	case "LanesToDisplay":
		milestone.LanesToDisplay = make([]*Lane, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Lanes {
					if stage.LaneMap_Staged_Order[__instance__] == uint(id) {
						milestone.LanesToDisplay = append(milestone.LanesToDisplay, __instance__)
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

func SetFieldStringValueFromPointer(instance GongstructIF, fieldName string, value GongFieldValue, stage *Stage) error {
	return instance.GongSetFieldValue(fieldName, value, stage)
}

// insertion point for generic get gongstruct name
func (arrow *Arrow) GongGetGongstructName() string {
	return "Arrow"
}

func (bar *Bar) GongGetGongstructName() string {
	return "Bar"
}

func (gantt *Gantt) GongGetGongstructName() string {
	return "Gantt"
}

func (group *Group) GongGetGongstructName() string {
	return "Group"
}

func (lane *Lane) GongGetGongstructName() string {
	return "Lane"
}

func (laneuse *LaneUse) GongGetGongstructName() string {
	return "LaneUse"
}

func (milestone *Milestone) GongGetGongstructName() string {
	return "Milestone"
}

func GetGongstructNameFromPointer(instance GongstructIF) (res string) {
	res = instance.GongGetGongstructName()
	return
}

func (stage *Stage) ResetMapStrings() {

	// insertion point for generic get gongstruct name
	stage.Arrows_mapString = make(map[string]*Arrow)
	for arrow := range stage.Arrows {
		stage.Arrows_mapString[arrow.Name] = arrow
	}

	stage.Bars_mapString = make(map[string]*Bar)
	for bar := range stage.Bars {
		stage.Bars_mapString[bar.Name] = bar
	}

	stage.Gantts_mapString = make(map[string]*Gantt)
	for gantt := range stage.Gantts {
		stage.Gantts_mapString[gantt.Name] = gantt
	}

	stage.Groups_mapString = make(map[string]*Group)
	for group := range stage.Groups {
		stage.Groups_mapString[group.Name] = group
	}

	stage.Lanes_mapString = make(map[string]*Lane)
	for lane := range stage.Lanes {
		stage.Lanes_mapString[lane.Name] = lane
	}

	stage.LaneUses_mapString = make(map[string]*LaneUse)
	for laneuse := range stage.LaneUses {
		stage.LaneUses_mapString[laneuse.Name] = laneuse
	}

	stage.Milestones_mapString = make(map[string]*Milestone)
	for milestone := range stage.Milestones {
		stage.Milestones_mapString[milestone.Name] = milestone
	}

}

// Last line of the template
