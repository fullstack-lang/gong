// generated code - do not edit
package models

import (
	"cmp"
	"errors"
	"fmt"
	"math"
	"slices"
	"time"
)

func __Gong__Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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

// StageStruct enables storage of staged instances
// swagger:ignore
type StageStruct struct {
	name string

	// insertion point for definition of arrays registering instances
	Arrows           map[*Arrow]any
	Arrows_mapString map[string]*Arrow

	// insertion point for slice of pointers maps
	OnAfterArrowCreateCallback OnAfterCreateInterface[Arrow]
	OnAfterArrowUpdateCallback OnAfterUpdateInterface[Arrow]
	OnAfterArrowDeleteCallback OnAfterDeleteInterface[Arrow]
	OnAfterArrowReadCallback   OnAfterReadInterface[Arrow]

	Bars           map[*Bar]any
	Bars_mapString map[string]*Bar

	// insertion point for slice of pointers maps
	OnAfterBarCreateCallback OnAfterCreateInterface[Bar]
	OnAfterBarUpdateCallback OnAfterUpdateInterface[Bar]
	OnAfterBarDeleteCallback OnAfterDeleteInterface[Bar]
	OnAfterBarReadCallback   OnAfterReadInterface[Bar]

	Gantts           map[*Gantt]any
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

	Groups           map[*Group]any
	Groups_mapString map[string]*Group

	// insertion point for slice of pointers maps
	Group_GroupLanes_reverseMap map[*Lane]*Group

	OnAfterGroupCreateCallback OnAfterCreateInterface[Group]
	OnAfterGroupUpdateCallback OnAfterUpdateInterface[Group]
	OnAfterGroupDeleteCallback OnAfterDeleteInterface[Group]
	OnAfterGroupReadCallback   OnAfterReadInterface[Group]

	Lanes           map[*Lane]any
	Lanes_mapString map[string]*Lane

	// insertion point for slice of pointers maps
	Lane_Bars_reverseMap map[*Bar]*Lane

	OnAfterLaneCreateCallback OnAfterCreateInterface[Lane]
	OnAfterLaneUpdateCallback OnAfterUpdateInterface[Lane]
	OnAfterLaneDeleteCallback OnAfterDeleteInterface[Lane]
	OnAfterLaneReadCallback   OnAfterReadInterface[Lane]

	LaneUses           map[*LaneUse]any
	LaneUses_mapString map[string]*LaneUse

	// insertion point for slice of pointers maps
	OnAfterLaneUseCreateCallback OnAfterCreateInterface[LaneUse]
	OnAfterLaneUseUpdateCallback OnAfterUpdateInterface[LaneUse]
	OnAfterLaneUseDeleteCallback OnAfterDeleteInterface[LaneUse]
	OnAfterLaneUseReadCallback   OnAfterReadInterface[LaneUse]

	Milestones           map[*Milestone]any
	Milestones_mapString map[string]*Milestone

	// insertion point for slice of pointers maps
	Milestone_LanesToDisplayMilestoneUse_reverseMap map[*LaneUse]*Milestone

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
	Order            uint
	Map_Staged_Order map[any]uint
}

func (stage *StageStruct) GetType() string {
	return "github.com/fullstack-lang/gong/lib/gantt/go/models"
}

type GONG__Identifier struct {
	Ident string
	Type  GONG__ExpressionType
}

type OnInitCommitInterface interface {
	BeforeCommit(stage *StageStruct)
}

// OnAfterCreateInterface callback when an instance is updated from the front
type OnAfterCreateInterface[Type Gongstruct] interface {
	OnAfterCreate(stage *StageStruct,
		instance *Type)
}

// OnAfterReadInterface callback when an instance is updated from the front
type OnAfterReadInterface[Type Gongstruct] interface {
	OnAfterRead(stage *StageStruct,
		instance *Type)
}

// OnAfterUpdateInterface callback when an instance is updated from the front
type OnAfterUpdateInterface[Type Gongstruct] interface {
	OnAfterUpdate(stage *StageStruct, old, new *Type)
}

// OnAfterDeleteInterface callback when an instance is updated from the front
type OnAfterDeleteInterface[Type Gongstruct] interface {
	OnAfterDelete(stage *StageStruct,
		staged, front *Type)
}

type BackRepoInterface interface {
	Commit(stage *StageStruct)
	Checkout(stage *StageStruct)
	Backup(stage *StageStruct, dirPath string)
	Restore(stage *StageStruct, dirPath string)
	BackupXL(stage *StageStruct, dirPath string)
	RestoreXL(stage *StageStruct, dirPath string)
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

func NewStage(name string) (stage *StageStruct) {

	stage = &StageStruct{ // insertion point for array initiatialisation
		Arrows:           make(map[*Arrow]any),
		Arrows_mapString: make(map[string]*Arrow),

		Bars:           make(map[*Bar]any),
		Bars_mapString: make(map[string]*Bar),

		Gantts:           make(map[*Gantt]any),
		Gantts_mapString: make(map[string]*Gantt),

		Groups:           make(map[*Group]any),
		Groups_mapString: make(map[string]*Group),

		Lanes:           make(map[*Lane]any),
		Lanes_mapString: make(map[string]*Lane),

		LaneUses:           make(map[*LaneUse]any),
		LaneUses_mapString: make(map[string]*LaneUse),

		Milestones:           make(map[*Milestone]any),
		Milestones_mapString: make(map[string]*Milestone),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		name: name,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		Map_Staged_Order: make(map[any]uint),
	}

	return
}

func (stage *StageStruct) GetName() string {
	return stage.name
}

func (stage *StageStruct) CommitWithSuspendedCallbacks() {

	tmp := stage.OnInitCommitFromBackCallback
	stage.OnInitCommitFromBackCallback = nil
	stage.Commit()
	stage.OnInitCommitFromBackCallback = tmp
}

func (stage *StageStruct) Commit() {
	stage.ComputeReverseMaps()

	if stage.BackRepo != nil {
		stage.BackRepo.Commit(stage)
	}

	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["Arrow"] = len(stage.Arrows)
	stage.Map_GongStructName_InstancesNb["Bar"] = len(stage.Bars)
	stage.Map_GongStructName_InstancesNb["Gantt"] = len(stage.Gantts)
	stage.Map_GongStructName_InstancesNb["Group"] = len(stage.Groups)
	stage.Map_GongStructName_InstancesNb["Lane"] = len(stage.Lanes)
	stage.Map_GongStructName_InstancesNb["LaneUse"] = len(stage.LaneUses)
	stage.Map_GongStructName_InstancesNb["Milestone"] = len(stage.Milestones)

}

func (stage *StageStruct) Checkout() {
	if stage.BackRepo != nil {
		stage.BackRepo.Checkout(stage)
	}

	stage.ComputeReverseMaps()
	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["Arrow"] = len(stage.Arrows)
	stage.Map_GongStructName_InstancesNb["Bar"] = len(stage.Bars)
	stage.Map_GongStructName_InstancesNb["Gantt"] = len(stage.Gantts)
	stage.Map_GongStructName_InstancesNb["Group"] = len(stage.Groups)
	stage.Map_GongStructName_InstancesNb["Lane"] = len(stage.Lanes)
	stage.Map_GongStructName_InstancesNb["LaneUse"] = len(stage.LaneUses)
	stage.Map_GongStructName_InstancesNb["Milestone"] = len(stage.Milestones)

}

// backup generates backup files in the dirPath
func (stage *StageStruct) Backup(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Backup(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *StageStruct) Restore(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Restore(stage, dirPath)
	}
}

// backup generates backup files in the dirPath
func (stage *StageStruct) BackupXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.BackupXL(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *StageStruct) RestoreXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.RestoreXL(stage, dirPath)
	}
}

// insertion point for cumulative sub template with model space calls
// Stage puts arrow to the model stage
func (arrow *Arrow) Stage(stage *StageStruct) *Arrow {

	if _, ok := stage.Arrows[arrow]; !ok {
		stage.Arrows[arrow] = __member
		stage.Map_Staged_Order[arrow] = stage.Order
		stage.Order++
	}
	stage.Arrows_mapString[arrow.Name] = arrow

	return arrow
}

// Unstage removes arrow off the model stage
func (arrow *Arrow) Unstage(stage *StageStruct) *Arrow {
	delete(stage.Arrows, arrow)
	delete(stage.Arrows_mapString, arrow.Name)
	return arrow
}

// UnstageVoid removes arrow off the model stage
func (arrow *Arrow) UnstageVoid(stage *StageStruct) {
	delete(stage.Arrows, arrow)
	delete(stage.Arrows_mapString, arrow.Name)
}

// commit arrow to the back repo (if it is already staged)
func (arrow *Arrow) Commit(stage *StageStruct) *Arrow {
	if _, ok := stage.Arrows[arrow]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitArrow(arrow)
		}
	}
	return arrow
}

func (arrow *Arrow) CommitVoid(stage *StageStruct) {
	arrow.Commit(stage)
}

// Checkout arrow to the back repo (if it is already staged)
func (arrow *Arrow) Checkout(stage *StageStruct) *Arrow {
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

// Stage puts bar to the model stage
func (bar *Bar) Stage(stage *StageStruct) *Bar {

	if _, ok := stage.Bars[bar]; !ok {
		stage.Bars[bar] = __member
		stage.Map_Staged_Order[bar] = stage.Order
		stage.Order++
	}
	stage.Bars_mapString[bar.Name] = bar

	return bar
}

// Unstage removes bar off the model stage
func (bar *Bar) Unstage(stage *StageStruct) *Bar {
	delete(stage.Bars, bar)
	delete(stage.Bars_mapString, bar.Name)
	return bar
}

// UnstageVoid removes bar off the model stage
func (bar *Bar) UnstageVoid(stage *StageStruct) {
	delete(stage.Bars, bar)
	delete(stage.Bars_mapString, bar.Name)
}

// commit bar to the back repo (if it is already staged)
func (bar *Bar) Commit(stage *StageStruct) *Bar {
	if _, ok := stage.Bars[bar]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitBar(bar)
		}
	}
	return bar
}

func (bar *Bar) CommitVoid(stage *StageStruct) {
	bar.Commit(stage)
}

// Checkout bar to the back repo (if it is already staged)
func (bar *Bar) Checkout(stage *StageStruct) *Bar {
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

// Stage puts gantt to the model stage
func (gantt *Gantt) Stage(stage *StageStruct) *Gantt {

	if _, ok := stage.Gantts[gantt]; !ok {
		stage.Gantts[gantt] = __member
		stage.Map_Staged_Order[gantt] = stage.Order
		stage.Order++
	}
	stage.Gantts_mapString[gantt.Name] = gantt

	return gantt
}

// Unstage removes gantt off the model stage
func (gantt *Gantt) Unstage(stage *StageStruct) *Gantt {
	delete(stage.Gantts, gantt)
	delete(stage.Gantts_mapString, gantt.Name)
	return gantt
}

// UnstageVoid removes gantt off the model stage
func (gantt *Gantt) UnstageVoid(stage *StageStruct) {
	delete(stage.Gantts, gantt)
	delete(stage.Gantts_mapString, gantt.Name)
}

// commit gantt to the back repo (if it is already staged)
func (gantt *Gantt) Commit(stage *StageStruct) *Gantt {
	if _, ok := stage.Gantts[gantt]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitGantt(gantt)
		}
	}
	return gantt
}

func (gantt *Gantt) CommitVoid(stage *StageStruct) {
	gantt.Commit(stage)
}

// Checkout gantt to the back repo (if it is already staged)
func (gantt *Gantt) Checkout(stage *StageStruct) *Gantt {
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

// Stage puts group to the model stage
func (group *Group) Stage(stage *StageStruct) *Group {

	if _, ok := stage.Groups[group]; !ok {
		stage.Groups[group] = __member
		stage.Map_Staged_Order[group] = stage.Order
		stage.Order++
	}
	stage.Groups_mapString[group.Name] = group

	return group
}

// Unstage removes group off the model stage
func (group *Group) Unstage(stage *StageStruct) *Group {
	delete(stage.Groups, group)
	delete(stage.Groups_mapString, group.Name)
	return group
}

// UnstageVoid removes group off the model stage
func (group *Group) UnstageVoid(stage *StageStruct) {
	delete(stage.Groups, group)
	delete(stage.Groups_mapString, group.Name)
}

// commit group to the back repo (if it is already staged)
func (group *Group) Commit(stage *StageStruct) *Group {
	if _, ok := stage.Groups[group]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitGroup(group)
		}
	}
	return group
}

func (group *Group) CommitVoid(stage *StageStruct) {
	group.Commit(stage)
}

// Checkout group to the back repo (if it is already staged)
func (group *Group) Checkout(stage *StageStruct) *Group {
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

// Stage puts lane to the model stage
func (lane *Lane) Stage(stage *StageStruct) *Lane {

	if _, ok := stage.Lanes[lane]; !ok {
		stage.Lanes[lane] = __member
		stage.Map_Staged_Order[lane] = stage.Order
		stage.Order++
	}
	stage.Lanes_mapString[lane.Name] = lane

	return lane
}

// Unstage removes lane off the model stage
func (lane *Lane) Unstage(stage *StageStruct) *Lane {
	delete(stage.Lanes, lane)
	delete(stage.Lanes_mapString, lane.Name)
	return lane
}

// UnstageVoid removes lane off the model stage
func (lane *Lane) UnstageVoid(stage *StageStruct) {
	delete(stage.Lanes, lane)
	delete(stage.Lanes_mapString, lane.Name)
}

// commit lane to the back repo (if it is already staged)
func (lane *Lane) Commit(stage *StageStruct) *Lane {
	if _, ok := stage.Lanes[lane]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitLane(lane)
		}
	}
	return lane
}

func (lane *Lane) CommitVoid(stage *StageStruct) {
	lane.Commit(stage)
}

// Checkout lane to the back repo (if it is already staged)
func (lane *Lane) Checkout(stage *StageStruct) *Lane {
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

// Stage puts laneuse to the model stage
func (laneuse *LaneUse) Stage(stage *StageStruct) *LaneUse {

	if _, ok := stage.LaneUses[laneuse]; !ok {
		stage.LaneUses[laneuse] = __member
		stage.Map_Staged_Order[laneuse] = stage.Order
		stage.Order++
	}
	stage.LaneUses_mapString[laneuse.Name] = laneuse

	return laneuse
}

// Unstage removes laneuse off the model stage
func (laneuse *LaneUse) Unstage(stage *StageStruct) *LaneUse {
	delete(stage.LaneUses, laneuse)
	delete(stage.LaneUses_mapString, laneuse.Name)
	return laneuse
}

// UnstageVoid removes laneuse off the model stage
func (laneuse *LaneUse) UnstageVoid(stage *StageStruct) {
	delete(stage.LaneUses, laneuse)
	delete(stage.LaneUses_mapString, laneuse.Name)
}

// commit laneuse to the back repo (if it is already staged)
func (laneuse *LaneUse) Commit(stage *StageStruct) *LaneUse {
	if _, ok := stage.LaneUses[laneuse]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitLaneUse(laneuse)
		}
	}
	return laneuse
}

func (laneuse *LaneUse) CommitVoid(stage *StageStruct) {
	laneuse.Commit(stage)
}

// Checkout laneuse to the back repo (if it is already staged)
func (laneuse *LaneUse) Checkout(stage *StageStruct) *LaneUse {
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

// Stage puts milestone to the model stage
func (milestone *Milestone) Stage(stage *StageStruct) *Milestone {

	if _, ok := stage.Milestones[milestone]; !ok {
		stage.Milestones[milestone] = __member
		stage.Map_Staged_Order[milestone] = stage.Order
		stage.Order++
	}
	stage.Milestones_mapString[milestone.Name] = milestone

	return milestone
}

// Unstage removes milestone off the model stage
func (milestone *Milestone) Unstage(stage *StageStruct) *Milestone {
	delete(stage.Milestones, milestone)
	delete(stage.Milestones_mapString, milestone.Name)
	return milestone
}

// UnstageVoid removes milestone off the model stage
func (milestone *Milestone) UnstageVoid(stage *StageStruct) {
	delete(stage.Milestones, milestone)
	delete(stage.Milestones_mapString, milestone.Name)
}

// commit milestone to the back repo (if it is already staged)
func (milestone *Milestone) Commit(stage *StageStruct) *Milestone {
	if _, ok := stage.Milestones[milestone]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitMilestone(milestone)
		}
	}
	return milestone
}

func (milestone *Milestone) CommitVoid(stage *StageStruct) {
	milestone.Commit(stage)
}

// Checkout milestone to the back repo (if it is already staged)
func (milestone *Milestone) Checkout(stage *StageStruct) *Milestone {
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

func (stage *StageStruct) Reset() { // insertion point for array reset
	stage.Arrows = make(map[*Arrow]any)
	stage.Arrows_mapString = make(map[string]*Arrow)

	stage.Bars = make(map[*Bar]any)
	stage.Bars_mapString = make(map[string]*Bar)

	stage.Gantts = make(map[*Gantt]any)
	stage.Gantts_mapString = make(map[string]*Gantt)

	stage.Groups = make(map[*Group]any)
	stage.Groups_mapString = make(map[string]*Group)

	stage.Lanes = make(map[*Lane]any)
	stage.Lanes_mapString = make(map[string]*Lane)

	stage.LaneUses = make(map[*LaneUse]any)
	stage.LaneUses_mapString = make(map[string]*LaneUse)

	stage.Milestones = make(map[*Milestone]any)
	stage.Milestones_mapString = make(map[string]*Milestone)

}

func (stage *StageStruct) Nil() { // insertion point for array nil
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

func (stage *StageStruct) Unstage() { // insertion point for array nil
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
type PointerToGongstruct interface {
	GetName() string
	CommitVoid(*StageStruct)
	UnstageVoid(stage *StageStruct)
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

func GetGongstrucsSorted[T PointerToGongstruct](stage *StageStruct) (sortedSlice []T) {

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
func GongGetSet[Type GongstructSet](stage *StageStruct) *Type {
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

// GongGetMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GongGetMap[Type GongstructMapString](stage *StageStruct) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case map[string]*Arrow:
		return any(&stage.Arrows_mapString).(*Type)
	case map[string]*Bar:
		return any(&stage.Bars_mapString).(*Type)
	case map[string]*Gantt:
		return any(&stage.Gantts_mapString).(*Type)
	case map[string]*Group:
		return any(&stage.Groups_mapString).(*Type)
	case map[string]*Lane:
		return any(&stage.Lanes_mapString).(*Type)
	case map[string]*LaneUse:
		return any(&stage.LaneUses_mapString).(*Type)
	case map[string]*Milestone:
		return any(&stage.Milestones_mapString).(*Type)
	default:
		return nil
	}
}

// GetGongstructInstancesSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSet[Type Gongstruct](stage *StageStruct) *map[*Type]any {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case Arrow:
		return any(&stage.Arrows).(*map[*Type]any)
	case Bar:
		return any(&stage.Bars).(*map[*Type]any)
	case Gantt:
		return any(&stage.Gantts).(*map[*Type]any)
	case Group:
		return any(&stage.Groups).(*map[*Type]any)
	case Lane:
		return any(&stage.Lanes).(*map[*Type]any)
	case LaneUse:
		return any(&stage.LaneUses).(*map[*Type]any)
	case Milestone:
		return any(&stage.Milestones).(*map[*Type]any)
	default:
		return nil
	}
}

// GetGongstructInstancesSetFromPointerType returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSetFromPointerType[Type PointerToGongstruct](stage *StageStruct) *map[Type]any {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case *Arrow:
		return any(&stage.Arrows).(*map[Type]any)
	case *Bar:
		return any(&stage.Bars).(*map[Type]any)
	case *Gantt:
		return any(&stage.Gantts).(*map[Type]any)
	case *Group:
		return any(&stage.Groups).(*map[Type]any)
	case *Lane:
		return any(&stage.Lanes).(*map[Type]any)
	case *LaneUse:
		return any(&stage.LaneUses).(*map[Type]any)
	case *Milestone:
		return any(&stage.Milestones).(*map[Type]any)
	default:
		return nil
	}
}

// GetGongstructInstancesMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GetGongstructInstancesMap[Type Gongstruct](stage *StageStruct) *map[string]*Type {
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
			// field is initialized with an instance of LaneUse with the name of the field
			LanesToDisplayMilestoneUse: []*LaneUse{{Name: "LanesToDisplayMilestoneUse"}},
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
func GetPointerReverseMap[Start, End Gongstruct](fieldname string, stage *StageStruct) map[*End][]*Start {

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
func GetSliceOfPointersReverseMap[Start, End Gongstruct](fieldname string, stage *StageStruct) map[*End]*Start {

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
			res := make(map[*Lane]*Gantt)
			for gantt := range stage.Gantts {
				for _, lane_ := range gantt.Lanes {
					res[lane_] = gantt
				}
			}
			return any(res).(map[*End]*Start)
		case "Milestones":
			res := make(map[*Milestone]*Gantt)
			for gantt := range stage.Gantts {
				for _, milestone_ := range gantt.Milestones {
					res[milestone_] = gantt
				}
			}
			return any(res).(map[*End]*Start)
		case "Groups":
			res := make(map[*Group]*Gantt)
			for gantt := range stage.Gantts {
				for _, group_ := range gantt.Groups {
					res[group_] = gantt
				}
			}
			return any(res).(map[*End]*Start)
		case "Arrows":
			res := make(map[*Arrow]*Gantt)
			for gantt := range stage.Gantts {
				for _, arrow_ := range gantt.Arrows {
					res[arrow_] = gantt
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of Group
	case Group:
		switch fieldname {
		// insertion point for per direct association field
		case "GroupLanes":
			res := make(map[*Lane]*Group)
			for group := range stage.Groups {
				for _, lane_ := range group.GroupLanes {
					res[lane_] = group
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of Lane
	case Lane:
		switch fieldname {
		// insertion point for per direct association field
		case "Bars":
			res := make(map[*Bar]*Lane)
			for lane := range stage.Lanes {
				for _, bar_ := range lane.Bars {
					res[bar_] = lane
				}
			}
			return any(res).(map[*End]*Start)
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
		case "LanesToDisplayMilestoneUse":
			res := make(map[*LaneUse]*Milestone)
			for milestone := range stage.Milestones {
				for _, laneuse_ := range milestone.LanesToDisplayMilestoneUse {
					res[laneuse_] = milestone
				}
			}
			return any(res).(map[*End]*Start)
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
	case Arrow:
		res = "Arrow"
	case Bar:
		res = "Bar"
	case Gantt:
		res = "Gantt"
	case Group:
		res = "Group"
	case Lane:
		res = "Lane"
	case LaneUse:
		res = "LaneUse"
	case Milestone:
		res = "Milestone"
	}
	return res
}

// GetPointerToGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetPointerToGongstructName[Type PointerToGongstruct]() (res string) {

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

// GetFields return the array of the fields
func GetFields[Type Gongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case Arrow:
		res = []string{"Name", "From", "To", "OptionnalColor", "OptionnalStroke"}
	case Bar:
		res = []string{"Name", "Start", "End", "ComputedDuration", "OptionnalColor", "OptionnalStroke", "FillOpacity", "StrokeWidth", "StrokeDashArray"}
	case Gantt:
		res = []string{"Name", "ComputedStart", "ComputedEnd", "ComputedDuration", "UseManualStartAndEndDates", "ManualStart", "ManualEnd", "LaneHeight", "RatioBarToLaneHeight", "YTopMargin", "XLeftText", "TextHeight", "XLeftLanes", "XRightMargin", "ArrowLengthToTheRightOfStartBar", "ArrowTipLenght", "TimeLine_Color", "TimeLine_FillOpacity", "TimeLine_Stroke", "TimeLine_StrokeWidth", "Group_Stroke", "Group_StrokeWidth", "Group_StrokeDashArray", "DateYOffset", "AlignOnStartEndOnYearStart", "Lanes", "Milestones", "Groups", "Arrows"}
	case Group:
		res = []string{"Name", "GroupLanes"}
	case Lane:
		res = []string{"Name", "Order", "Bars"}
	case LaneUse:
		res = []string{"Name", "Lane"}
	case Milestone:
		res = []string{"Name", "Date", "DisplayVerticalBar", "LanesToDisplayMilestoneUse"}
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
	case Arrow:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Gantt"
		rf.Fieldname = "Arrows"
		res = append(res, rf)
	case Bar:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Lane"
		rf.Fieldname = "Bars"
		res = append(res, rf)
	case Gantt:
		var rf ReverseField
		_ = rf
	case Group:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Gantt"
		rf.Fieldname = "Groups"
		res = append(res, rf)
	case Lane:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Gantt"
		rf.Fieldname = "Lanes"
		res = append(res, rf)
		rf.GongstructName = "Group"
		rf.Fieldname = "GroupLanes"
		res = append(res, rf)
	case LaneUse:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Milestone"
		rf.Fieldname = "LanesToDisplayMilestoneUse"
		res = append(res, rf)
	case Milestone:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Gantt"
		rf.Fieldname = "Milestones"
		res = append(res, rf)
	}
	return
}

// GetFieldsFromPointer return the array of the fields
func GetFieldsFromPointer[Type PointerToGongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *Arrow:
		res = []string{"Name", "From", "To", "OptionnalColor", "OptionnalStroke"}
	case *Bar:
		res = []string{"Name", "Start", "End", "ComputedDuration", "OptionnalColor", "OptionnalStroke", "FillOpacity", "StrokeWidth", "StrokeDashArray"}
	case *Gantt:
		res = []string{"Name", "ComputedStart", "ComputedEnd", "ComputedDuration", "UseManualStartAndEndDates", "ManualStart", "ManualEnd", "LaneHeight", "RatioBarToLaneHeight", "YTopMargin", "XLeftText", "TextHeight", "XLeftLanes", "XRightMargin", "ArrowLengthToTheRightOfStartBar", "ArrowTipLenght", "TimeLine_Color", "TimeLine_FillOpacity", "TimeLine_Stroke", "TimeLine_StrokeWidth", "Group_Stroke", "Group_StrokeWidth", "Group_StrokeDashArray", "DateYOffset", "AlignOnStartEndOnYearStart", "Lanes", "Milestones", "Groups", "Arrows"}
	case *Group:
		res = []string{"Name", "GroupLanes"}
	case *Lane:
		res = []string{"Name", "Order", "Bars"}
	case *LaneUse:
		res = []string{"Name", "Lane"}
	case *Milestone:
		res = []string{"Name", "Date", "DisplayVerticalBar", "LanesToDisplayMilestoneUse"}
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
	case *Arrow:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "From":
			if inferedInstance.From != nil {
				res.valueString = inferedInstance.From.Name
			}
		case "To":
			if inferedInstance.To != nil {
				res.valueString = inferedInstance.To.Name
			}
		case "OptionnalColor":
			res.valueString = inferedInstance.OptionnalColor
		case "OptionnalStroke":
			res.valueString = inferedInstance.OptionnalStroke
		}
	case *Bar:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Start":
			res.valueString = inferedInstance.Start.String()
		case "End":
			res.valueString = inferedInstance.End.String()
		case "ComputedDuration":
			if math.Abs(inferedInstance.ComputedDuration.Hours()) >= 24 {
				days := __Gong__Abs(int(int(inferedInstance.ComputedDuration.Hours()) / 24))
				months := int(days / 31)
				days = days - months*31

				remainingHours := int(inferedInstance.ComputedDuration.Hours()) % 24
				remainingMinutes := int(inferedInstance.ComputedDuration.Minutes()) % 60
				remainingSeconds := int(inferedInstance.ComputedDuration.Seconds()) % 60

				if inferedInstance.ComputedDuration.Hours() < 0 {
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
				res.valueString = fmt.Sprintf("%s\n", inferedInstance.ComputedDuration.String())
			}
		case "OptionnalColor":
			res.valueString = inferedInstance.OptionnalColor
		case "OptionnalStroke":
			res.valueString = inferedInstance.OptionnalStroke
		case "FillOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FillOpacity)
			res.valueFloat = inferedInstance.FillOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeWidth":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
			res.valueFloat = inferedInstance.StrokeWidth
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeDashArray":
			res.valueString = inferedInstance.StrokeDashArray
		}
	case *Gantt:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "ComputedStart":
			res.valueString = inferedInstance.ComputedStart.String()
		case "ComputedEnd":
			res.valueString = inferedInstance.ComputedEnd.String()
		case "ComputedDuration":
			if math.Abs(inferedInstance.ComputedDuration.Hours()) >= 24 {
				days := __Gong__Abs(int(int(inferedInstance.ComputedDuration.Hours()) / 24))
				months := int(days / 31)
				days = days - months*31

				remainingHours := int(inferedInstance.ComputedDuration.Hours()) % 24
				remainingMinutes := int(inferedInstance.ComputedDuration.Minutes()) % 60
				remainingSeconds := int(inferedInstance.ComputedDuration.Seconds()) % 60

				if inferedInstance.ComputedDuration.Hours() < 0 {
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
				res.valueString = fmt.Sprintf("%s\n", inferedInstance.ComputedDuration.String())
			}
		case "UseManualStartAndEndDates":
			res.valueString = fmt.Sprintf("%t", inferedInstance.UseManualStartAndEndDates)
			res.valueBool = inferedInstance.UseManualStartAndEndDates
			res.GongFieldValueType = GongFieldValueTypeBool
		case "ManualStart":
			res.valueString = inferedInstance.ManualStart.String()
		case "ManualEnd":
			res.valueString = inferedInstance.ManualEnd.String()
		case "LaneHeight":
			res.valueString = fmt.Sprintf("%f", inferedInstance.LaneHeight)
			res.valueFloat = inferedInstance.LaneHeight
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "RatioBarToLaneHeight":
			res.valueString = fmt.Sprintf("%f", inferedInstance.RatioBarToLaneHeight)
			res.valueFloat = inferedInstance.RatioBarToLaneHeight
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "YTopMargin":
			res.valueString = fmt.Sprintf("%f", inferedInstance.YTopMargin)
			res.valueFloat = inferedInstance.YTopMargin
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "XLeftText":
			res.valueString = fmt.Sprintf("%f", inferedInstance.XLeftText)
			res.valueFloat = inferedInstance.XLeftText
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "TextHeight":
			res.valueString = fmt.Sprintf("%f", inferedInstance.TextHeight)
			res.valueFloat = inferedInstance.TextHeight
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "XLeftLanes":
			res.valueString = fmt.Sprintf("%f", inferedInstance.XLeftLanes)
			res.valueFloat = inferedInstance.XLeftLanes
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "XRightMargin":
			res.valueString = fmt.Sprintf("%f", inferedInstance.XRightMargin)
			res.valueFloat = inferedInstance.XRightMargin
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "ArrowLengthToTheRightOfStartBar":
			res.valueString = fmt.Sprintf("%f", inferedInstance.ArrowLengthToTheRightOfStartBar)
			res.valueFloat = inferedInstance.ArrowLengthToTheRightOfStartBar
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "ArrowTipLenght":
			res.valueString = fmt.Sprintf("%f", inferedInstance.ArrowTipLenght)
			res.valueFloat = inferedInstance.ArrowTipLenght
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "TimeLine_Color":
			res.valueString = inferedInstance.TimeLine_Color
		case "TimeLine_FillOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.TimeLine_FillOpacity)
			res.valueFloat = inferedInstance.TimeLine_FillOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "TimeLine_Stroke":
			res.valueString = inferedInstance.TimeLine_Stroke
		case "TimeLine_StrokeWidth":
			res.valueString = fmt.Sprintf("%f", inferedInstance.TimeLine_StrokeWidth)
			res.valueFloat = inferedInstance.TimeLine_StrokeWidth
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Group_Stroke":
			res.valueString = inferedInstance.Group_Stroke
		case "Group_StrokeWidth":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Group_StrokeWidth)
			res.valueFloat = inferedInstance.Group_StrokeWidth
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Group_StrokeDashArray":
			res.valueString = inferedInstance.Group_StrokeDashArray
		case "DateYOffset":
			res.valueString = fmt.Sprintf("%f", inferedInstance.DateYOffset)
			res.valueFloat = inferedInstance.DateYOffset
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "AlignOnStartEndOnYearStart":
			res.valueString = fmt.Sprintf("%t", inferedInstance.AlignOnStartEndOnYearStart)
			res.valueBool = inferedInstance.AlignOnStartEndOnYearStart
			res.GongFieldValueType = GongFieldValueTypeBool
		case "Lanes":
			for idx, __instance__ := range inferedInstance.Lanes {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "Milestones":
			for idx, __instance__ := range inferedInstance.Milestones {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "Groups":
			for idx, __instance__ := range inferedInstance.Groups {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "Arrows":
			for idx, __instance__ := range inferedInstance.Arrows {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case *Group:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "GroupLanes":
			for idx, __instance__ := range inferedInstance.GroupLanes {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case *Lane:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Order":
			res.valueString = fmt.Sprintf("%d", inferedInstance.Order)
			res.valueInt = inferedInstance.Order
			res.GongFieldValueType = GongFieldValueTypeInt
		case "Bars":
			for idx, __instance__ := range inferedInstance.Bars {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case *LaneUse:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Lane":
			if inferedInstance.Lane != nil {
				res.valueString = inferedInstance.Lane.Name
			}
		}
	case *Milestone:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Date":
			res.valueString = inferedInstance.Date.String()
		case "DisplayVerticalBar":
			res.valueString = fmt.Sprintf("%t", inferedInstance.DisplayVerticalBar)
			res.valueBool = inferedInstance.DisplayVerticalBar
			res.GongFieldValueType = GongFieldValueTypeBool
		case "LanesToDisplayMilestoneUse":
			for idx, __instance__ := range inferedInstance.LanesToDisplayMilestoneUse {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	default:
		_ = inferedInstance
	}
	return
}

func GetFieldStringValue(instance any, fieldName string) (res GongFieldValue) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case Arrow:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "From":
			if inferedInstance.From != nil {
				res.valueString = inferedInstance.From.Name
			}
		case "To":
			if inferedInstance.To != nil {
				res.valueString = inferedInstance.To.Name
			}
		case "OptionnalColor":
			res.valueString = inferedInstance.OptionnalColor
		case "OptionnalStroke":
			res.valueString = inferedInstance.OptionnalStroke
		}
	case Bar:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Start":
			res.valueString = inferedInstance.Start.String()
		case "End":
			res.valueString = inferedInstance.End.String()
		case "ComputedDuration":
			if math.Abs(inferedInstance.ComputedDuration.Hours()) >= 24 {
				days := __Gong__Abs(int(int(inferedInstance.ComputedDuration.Hours()) / 24))
				months := int(days / 31)
				days = days - months*31

				remainingHours := int(inferedInstance.ComputedDuration.Hours()) % 24
				remainingMinutes := int(inferedInstance.ComputedDuration.Minutes()) % 60
				remainingSeconds := int(inferedInstance.ComputedDuration.Seconds()) % 60

				if inferedInstance.ComputedDuration.Hours() < 0 {
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
				res.valueString = fmt.Sprintf("%s\n", inferedInstance.ComputedDuration.String())
			}
		case "OptionnalColor":
			res.valueString = inferedInstance.OptionnalColor
		case "OptionnalStroke":
			res.valueString = inferedInstance.OptionnalStroke
		case "FillOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.FillOpacity)
			res.valueFloat = inferedInstance.FillOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeWidth":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StrokeWidth)
			res.valueFloat = inferedInstance.StrokeWidth
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeDashArray":
			res.valueString = inferedInstance.StrokeDashArray
		}
	case Gantt:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "ComputedStart":
			res.valueString = inferedInstance.ComputedStart.String()
		case "ComputedEnd":
			res.valueString = inferedInstance.ComputedEnd.String()
		case "ComputedDuration":
			if math.Abs(inferedInstance.ComputedDuration.Hours()) >= 24 {
				days := __Gong__Abs(int(int(inferedInstance.ComputedDuration.Hours()) / 24))
				months := int(days / 31)
				days = days - months*31

				remainingHours := int(inferedInstance.ComputedDuration.Hours()) % 24
				remainingMinutes := int(inferedInstance.ComputedDuration.Minutes()) % 60
				remainingSeconds := int(inferedInstance.ComputedDuration.Seconds()) % 60

				if inferedInstance.ComputedDuration.Hours() < 0 {
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
				res.valueString = fmt.Sprintf("%s\n", inferedInstance.ComputedDuration.String())
			}
		case "UseManualStartAndEndDates":
			res.valueString = fmt.Sprintf("%t", inferedInstance.UseManualStartAndEndDates)
			res.valueBool = inferedInstance.UseManualStartAndEndDates
			res.GongFieldValueType = GongFieldValueTypeBool
		case "ManualStart":
			res.valueString = inferedInstance.ManualStart.String()
		case "ManualEnd":
			res.valueString = inferedInstance.ManualEnd.String()
		case "LaneHeight":
			res.valueString = fmt.Sprintf("%f", inferedInstance.LaneHeight)
			res.valueFloat = inferedInstance.LaneHeight
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "RatioBarToLaneHeight":
			res.valueString = fmt.Sprintf("%f", inferedInstance.RatioBarToLaneHeight)
			res.valueFloat = inferedInstance.RatioBarToLaneHeight
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "YTopMargin":
			res.valueString = fmt.Sprintf("%f", inferedInstance.YTopMargin)
			res.valueFloat = inferedInstance.YTopMargin
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "XLeftText":
			res.valueString = fmt.Sprintf("%f", inferedInstance.XLeftText)
			res.valueFloat = inferedInstance.XLeftText
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "TextHeight":
			res.valueString = fmt.Sprintf("%f", inferedInstance.TextHeight)
			res.valueFloat = inferedInstance.TextHeight
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "XLeftLanes":
			res.valueString = fmt.Sprintf("%f", inferedInstance.XLeftLanes)
			res.valueFloat = inferedInstance.XLeftLanes
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "XRightMargin":
			res.valueString = fmt.Sprintf("%f", inferedInstance.XRightMargin)
			res.valueFloat = inferedInstance.XRightMargin
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "ArrowLengthToTheRightOfStartBar":
			res.valueString = fmt.Sprintf("%f", inferedInstance.ArrowLengthToTheRightOfStartBar)
			res.valueFloat = inferedInstance.ArrowLengthToTheRightOfStartBar
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "ArrowTipLenght":
			res.valueString = fmt.Sprintf("%f", inferedInstance.ArrowTipLenght)
			res.valueFloat = inferedInstance.ArrowTipLenght
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "TimeLine_Color":
			res.valueString = inferedInstance.TimeLine_Color
		case "TimeLine_FillOpacity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.TimeLine_FillOpacity)
			res.valueFloat = inferedInstance.TimeLine_FillOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "TimeLine_Stroke":
			res.valueString = inferedInstance.TimeLine_Stroke
		case "TimeLine_StrokeWidth":
			res.valueString = fmt.Sprintf("%f", inferedInstance.TimeLine_StrokeWidth)
			res.valueFloat = inferedInstance.TimeLine_StrokeWidth
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Group_Stroke":
			res.valueString = inferedInstance.Group_Stroke
		case "Group_StrokeWidth":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Group_StrokeWidth)
			res.valueFloat = inferedInstance.Group_StrokeWidth
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Group_StrokeDashArray":
			res.valueString = inferedInstance.Group_StrokeDashArray
		case "DateYOffset":
			res.valueString = fmt.Sprintf("%f", inferedInstance.DateYOffset)
			res.valueFloat = inferedInstance.DateYOffset
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "AlignOnStartEndOnYearStart":
			res.valueString = fmt.Sprintf("%t", inferedInstance.AlignOnStartEndOnYearStart)
			res.valueBool = inferedInstance.AlignOnStartEndOnYearStart
			res.GongFieldValueType = GongFieldValueTypeBool
		case "Lanes":
			for idx, __instance__ := range inferedInstance.Lanes {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "Milestones":
			for idx, __instance__ := range inferedInstance.Milestones {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "Groups":
			for idx, __instance__ := range inferedInstance.Groups {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "Arrows":
			for idx, __instance__ := range inferedInstance.Arrows {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case Group:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "GroupLanes":
			for idx, __instance__ := range inferedInstance.GroupLanes {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case Lane:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Order":
			res.valueString = fmt.Sprintf("%d", inferedInstance.Order)
			res.valueInt = inferedInstance.Order
			res.GongFieldValueType = GongFieldValueTypeInt
		case "Bars":
			for idx, __instance__ := range inferedInstance.Bars {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case LaneUse:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Lane":
			if inferedInstance.Lane != nil {
				res.valueString = inferedInstance.Lane.Name
			}
		}
	case Milestone:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Date":
			res.valueString = inferedInstance.Date.String()
		case "DisplayVerticalBar":
			res.valueString = fmt.Sprintf("%t", inferedInstance.DisplayVerticalBar)
			res.valueBool = inferedInstance.DisplayVerticalBar
			res.GongFieldValueType = GongFieldValueTypeBool
		case "LanesToDisplayMilestoneUse":
			for idx, __instance__ := range inferedInstance.LanesToDisplayMilestoneUse {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	default:
		_ = inferedInstance
	}
	return
}

// Last line of the template
