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

	tone_go "github.com/fullstack-lang/gong/lib/tone/go"
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
	Freqencys                map[*Freqency]struct{}
	Freqencys_reference      map[*Freqency]*Freqency
	Freqencys_referenceOrder map[*Freqency]uint // diff Unstage needs the reference order
	Freqencys_mapString      map[string]*Freqency

	// insertion point for slice of pointers maps
	OnAfterFreqencyCreateCallback OnAfterCreateInterface[Freqency]
	OnAfterFreqencyUpdateCallback OnAfterUpdateInterface[Freqency]
	OnAfterFreqencyDeleteCallback OnAfterDeleteInterface[Freqency]
	OnAfterFreqencyReadCallback   OnAfterReadInterface[Freqency]

	Notes                map[*Note]struct{}
	Notes_reference      map[*Note]*Note
	Notes_referenceOrder map[*Note]uint // diff Unstage needs the reference order
	Notes_mapString      map[string]*Note

	// insertion point for slice of pointers maps
	Note_Frequencies_reverseMap map[*Freqency]*Note

	OnAfterNoteCreateCallback OnAfterCreateInterface[Note]
	OnAfterNoteUpdateCallback OnAfterUpdateInterface[Note]
	OnAfterNoteDeleteCallback OnAfterDeleteInterface[Note]
	OnAfterNoteReadCallback   OnAfterReadInterface[Note]

	Players                map[*Player]struct{}
	Players_reference      map[*Player]*Player
	Players_referenceOrder map[*Player]uint // diff Unstage needs the reference order
	Players_mapString      map[string]*Player

	// insertion point for slice of pointers maps
	OnAfterPlayerCreateCallback OnAfterCreateInterface[Player]
	OnAfterPlayerUpdateCallback OnAfterUpdateInterface[Player]
	OnAfterPlayerDeleteCallback OnAfterDeleteInterface[Player]
	OnAfterPlayerReadCallback   OnAfterReadInterface[Player]

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
	FreqencyOrder            uint
	FreqencyMap_Staged_Order map[*Freqency]uint

	NoteOrder            uint
	NoteMap_Staged_Order map[*Note]uint

	PlayerOrder            uint
	PlayerMap_Staged_Order map[*Player]uint

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
	var maxFreqencyOrder uint
	var foundFreqency bool
	for _, order := range stage.FreqencyMap_Staged_Order {
		if !foundFreqency || order > maxFreqencyOrder {
			maxFreqencyOrder = order
			foundFreqency = true
		}
	}
	if foundFreqency {
		stage.FreqencyOrder = maxFreqencyOrder + 1
	} else {
		stage.FreqencyOrder = 0
	}

	var maxNoteOrder uint
	var foundNote bool
	for _, order := range stage.NoteMap_Staged_Order {
		if !foundNote || order > maxNoteOrder {
			maxNoteOrder = order
			foundNote = true
		}
	}
	if foundNote {
		stage.NoteOrder = maxNoteOrder + 1
	} else {
		stage.NoteOrder = 0
	}

	var maxPlayerOrder uint
	var foundPlayer bool
	for _, order := range stage.PlayerMap_Staged_Order {
		if !foundPlayer || order > maxPlayerOrder {
			maxPlayerOrder = order
			foundPlayer = true
		}
	}
	if foundPlayer {
		stage.PlayerOrder = maxPlayerOrder + 1
	} else {
		stage.PlayerOrder = 0
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
	case *Freqency:
		tmp := GetStructInstancesByOrder(stage.Freqencys, stage.FreqencyMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Freqency implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Note:
		tmp := GetStructInstancesByOrder(stage.Notes, stage.NoteMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Note implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Player:
		tmp := GetStructInstancesByOrder(stage.Players, stage.PlayerMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Player implements.
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
	case "Freqency":
		res = GetNamedStructInstances(stage.Freqencys, stage.FreqencyMap_Staged_Order)
	case "Note":
		res = GetNamedStructInstances(stage.Notes, stage.NoteMap_Staged_Order)
	case "Player":
		res = GetNamedStructInstances(stage.Players, stage.PlayerMap_Staged_Order)
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
	return "github.com/fullstack-lang/gong/lib/tone/go/models"
}

func (stage *Stage) GetMap_GongStructName_InstancesNb() map[string]int {
	return stage.Map_GongStructName_InstancesNb
}

func (stage *Stage) GetModelsEmbededDir() embed.FS {
	return tone_go.GoModelsDir
}

func (stage *Stage) GetDigramsEmbededDir() embed.FS {
	return tone_go.GoDiagramsDir
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
	CommitFreqency(freqency *Freqency)
	CheckoutFreqency(freqency *Freqency)
	CommitNote(note *Note)
	CheckoutNote(note *Note)
	CommitPlayer(player *Player)
	CheckoutPlayer(player *Player)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(name string) (stage *Stage) {

	stage = &Stage{ // insertion point for array initiatialisation
		Freqencys:           make(map[*Freqency]struct{}),
		Freqencys_mapString: make(map[string]*Freqency),

		Notes:           make(map[*Note]struct{}),
		Notes_mapString: make(map[string]*Note),

		Players:           make(map[*Player]struct{}),
		Players_mapString: make(map[string]*Player),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		name: name,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		FreqencyMap_Staged_Order: make(map[*Freqency]uint),

		NoteMap_Staged_Order: make(map[*Note]uint),

		PlayerMap_Staged_Order: make(map[*Player]uint),

		// end of insertion point
		GongUnmarshallers: map[string]ModelUnmarshaller{ // insertion point for unmarshallers
			"Freqency": &FreqencyUnmarshaller{},

			"Note": &NoteUnmarshaller{},

			"Player": &PlayerUnmarshaller{},

			// end of insertion point
		},

		NamedStructs: []*NamedStruct{ // insertion point for order map initialisations
			{name: "Freqency"},
			{name: "Note"},
			{name: "Player"},
		}, // end of insertion point

		navigationMode: GongNavigationModeNormal,
	}

	return
}

func GetOrder[Type Gongstruct](stage *Stage, instance *Type) uint {

	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *Freqency:
		return stage.FreqencyMap_Staged_Order[instance]
	case *Note:
		return stage.NoteMap_Staged_Order[instance]
	case *Player:
		return stage.PlayerMap_Staged_Order[instance]
	default:
		return 0 // should not happen
	}
}

func GetOrderPointerGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) uint {

	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *Freqency:
		return stage.FreqencyMap_Staged_Order[instance]
	case *Note:
		return stage.NoteMap_Staged_Order[instance]
	case *Player:
		return stage.PlayerMap_Staged_Order[instance]
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
	stage.Map_GongStructName_InstancesNb["Freqency"] = len(stage.Freqencys)
	stage.Map_GongStructName_InstancesNb["Note"] = len(stage.Notes)
	stage.Map_GongStructName_InstancesNb["Player"] = len(stage.Players)
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
// Stage puts freqency to the model stage
func (freqency *Freqency) Stage(stage *Stage) *Freqency {

	if _, ok := stage.Freqencys[freqency]; !ok {
		stage.Freqencys[freqency] = struct{}{}
		stage.FreqencyMap_Staged_Order[freqency] = stage.FreqencyOrder
		stage.FreqencyOrder++
	}
	stage.Freqencys_mapString[freqency.Name] = freqency

	return freqency
}

// StagePreserveOrder puts freqency to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.FreqencyOrder
// - update stage.FreqencyOrder accordingly
func (freqency *Freqency) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.Freqencys[freqency]; !ok {
		stage.Freqencys[freqency] = struct{}{}

		if order > stage.FreqencyOrder {
			stage.FreqencyOrder = order
		}
		stage.FreqencyMap_Staged_Order[freqency] = order
		stage.FreqencyOrder++
	}
	stage.Freqencys_mapString[freqency.Name] = freqency
}

// Unstage removes freqency off the model stage
func (freqency *Freqency) Unstage(stage *Stage) *Freqency {
	delete(stage.Freqencys, freqency)
	delete(stage.FreqencyMap_Staged_Order, freqency)
	delete(stage.Freqencys_mapString, freqency.Name)

	return freqency
}

// UnstageVoid removes freqency off the model stage
func (freqency *Freqency) UnstageVoid(stage *Stage) {
	delete(stage.Freqencys, freqency)
	delete(stage.FreqencyMap_Staged_Order, freqency)
	delete(stage.Freqencys_mapString, freqency.Name)
}

// commit freqency to the back repo (if it is already staged)
func (freqency *Freqency) Commit(stage *Stage) *Freqency {
	if _, ok := stage.Freqencys[freqency]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitFreqency(freqency)
		}
	}
	return freqency
}

func (freqency *Freqency) CommitVoid(stage *Stage) {
	freqency.Commit(stage)
}

func (freqency *Freqency) StageVoid(stage *Stage) {
	freqency.Stage(stage)
}

// Checkout freqency to the back repo (if it is already staged)
func (freqency *Freqency) Checkout(stage *Stage) *Freqency {
	if _, ok := stage.Freqencys[freqency]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutFreqency(freqency)
		}
	}
	return freqency
}

// for satisfaction of GongStruct interface
func (freqency *Freqency) GetName() (res string) {
	return freqency.Name
}

// for satisfaction of GongStruct interface
func (freqency *Freqency) SetName(name string) {
	freqency.Name = name
}

// Stage puts note to the model stage
func (note *Note) Stage(stage *Stage) *Note {

	if _, ok := stage.Notes[note]; !ok {
		stage.Notes[note] = struct{}{}
		stage.NoteMap_Staged_Order[note] = stage.NoteOrder
		stage.NoteOrder++
	}
	stage.Notes_mapString[note.Name] = note

	return note
}

// StagePreserveOrder puts note to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.NoteOrder
// - update stage.NoteOrder accordingly
func (note *Note) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.Notes[note]; !ok {
		stage.Notes[note] = struct{}{}

		if order > stage.NoteOrder {
			stage.NoteOrder = order
		}
		stage.NoteMap_Staged_Order[note] = order
		stage.NoteOrder++
	}
	stage.Notes_mapString[note.Name] = note
}

// Unstage removes note off the model stage
func (note *Note) Unstage(stage *Stage) *Note {
	delete(stage.Notes, note)
	delete(stage.NoteMap_Staged_Order, note)
	delete(stage.Notes_mapString, note.Name)

	return note
}

// UnstageVoid removes note off the model stage
func (note *Note) UnstageVoid(stage *Stage) {
	delete(stage.Notes, note)
	delete(stage.NoteMap_Staged_Order, note)
	delete(stage.Notes_mapString, note.Name)
}

// commit note to the back repo (if it is already staged)
func (note *Note) Commit(stage *Stage) *Note {
	if _, ok := stage.Notes[note]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitNote(note)
		}
	}
	return note
}

func (note *Note) CommitVoid(stage *Stage) {
	note.Commit(stage)
}

func (note *Note) StageVoid(stage *Stage) {
	note.Stage(stage)
}

// Checkout note to the back repo (if it is already staged)
func (note *Note) Checkout(stage *Stage) *Note {
	if _, ok := stage.Notes[note]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutNote(note)
		}
	}
	return note
}

// for satisfaction of GongStruct interface
func (note *Note) GetName() (res string) {
	return note.Name
}

// for satisfaction of GongStruct interface
func (note *Note) SetName(name string) {
	note.Name = name
}

// Stage puts player to the model stage
func (player *Player) Stage(stage *Stage) *Player {

	if _, ok := stage.Players[player]; !ok {
		stage.Players[player] = struct{}{}
		stage.PlayerMap_Staged_Order[player] = stage.PlayerOrder
		stage.PlayerOrder++
	}
	stage.Players_mapString[player.Name] = player

	return player
}

// StagePreserveOrder puts player to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.PlayerOrder
// - update stage.PlayerOrder accordingly
func (player *Player) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.Players[player]; !ok {
		stage.Players[player] = struct{}{}

		if order > stage.PlayerOrder {
			stage.PlayerOrder = order
		}
		stage.PlayerMap_Staged_Order[player] = order
		stage.PlayerOrder++
	}
	stage.Players_mapString[player.Name] = player
}

// Unstage removes player off the model stage
func (player *Player) Unstage(stage *Stage) *Player {
	delete(stage.Players, player)
	delete(stage.PlayerMap_Staged_Order, player)
	delete(stage.Players_mapString, player.Name)

	return player
}

// UnstageVoid removes player off the model stage
func (player *Player) UnstageVoid(stage *Stage) {
	delete(stage.Players, player)
	delete(stage.PlayerMap_Staged_Order, player)
	delete(stage.Players_mapString, player.Name)
}

// commit player to the back repo (if it is already staged)
func (player *Player) Commit(stage *Stage) *Player {
	if _, ok := stage.Players[player]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitPlayer(player)
		}
	}
	return player
}

func (player *Player) CommitVoid(stage *Stage) {
	player.Commit(stage)
}

func (player *Player) StageVoid(stage *Stage) {
	player.Stage(stage)
}

// Checkout player to the back repo (if it is already staged)
func (player *Player) Checkout(stage *Stage) *Player {
	if _, ok := stage.Players[player]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutPlayer(player)
		}
	}
	return player
}

// for satisfaction of GongStruct interface
func (player *Player) GetName() (res string) {
	return player.Name
}

// for satisfaction of GongStruct interface
func (player *Player) SetName(name string) {
	player.Name = name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMFreqency(Freqency *Freqency)
	CreateORMNote(Note *Note)
	CreateORMPlayer(Player *Player)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMFreqency(Freqency *Freqency)
	DeleteORMNote(Note *Note)
	DeleteORMPlayer(Player *Player)
}

func (stage *Stage) Reset() { // insertion point for array reset
	stage.Freqencys = make(map[*Freqency]struct{})
	stage.Freqencys_mapString = make(map[string]*Freqency)
	stage.FreqencyMap_Staged_Order = make(map[*Freqency]uint)
	stage.FreqencyOrder = 0

	stage.Notes = make(map[*Note]struct{})
	stage.Notes_mapString = make(map[string]*Note)
	stage.NoteMap_Staged_Order = make(map[*Note]uint)
	stage.NoteOrder = 0

	stage.Players = make(map[*Player]struct{})
	stage.Players_mapString = make(map[string]*Player)
	stage.PlayerMap_Staged_Order = make(map[*Player]uint)
	stage.PlayerOrder = 0

	if stage.GetProbeIF() != nil {
		stage.GetProbeIF().ResetNotifications()
	}
	if stage.IsInDeltaMode() {
		stage.ComputeReferenceAndOrders()
	}
}

func (stage *Stage) Nil() { // insertion point for array nil
	stage.Freqencys = nil
	stage.Freqencys_mapString = nil

	stage.Notes = nil
	stage.Notes_mapString = nil

	stage.Players = nil
	stage.Players_mapString = nil

}

func (stage *Stage) Unstage() { // insertion point for array nil
	for freqency := range stage.Freqencys {
		freqency.Unstage(stage)
	}

	for note := range stage.Notes {
		note.Unstage(stage)
	}

	for player := range stage.Players {
		player.Unstage(stage)
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
	case map[*Freqency]any:
		return any(&stage.Freqencys).(*Type)
	case map[*Note]any:
		return any(&stage.Notes).(*Type)
	case map[*Player]any:
		return any(&stage.Players).(*Type)
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
	case *Freqency:
		return any(stage.Freqencys_mapString).(map[string]Type)
	case *Note:
		return any(stage.Notes_mapString).(map[string]Type)
	case *Player:
		return any(stage.Players_mapString).(map[string]Type)
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
	case Freqency:
		return any(&stage.Freqencys).(*map[*Type]struct{})
	case Note:
		return any(&stage.Notes).(*map[*Type]struct{})
	case Player:
		return any(&stage.Players).(*map[*Type]struct{})
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
	case *Freqency:
		return any(&stage.Freqencys).(*map[Type]struct{})
	case *Note:
		return any(&stage.Notes).(*map[Type]struct{})
	case *Player:
		return any(&stage.Players).(*map[Type]struct{})
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
	case Freqency:
		return any(&stage.Freqencys_mapString).(*map[string]*Type)
	case Note:
		return any(&stage.Notes_mapString).(*map[string]*Type)
	case Player:
		return any(&stage.Players_mapString).(*map[string]*Type)
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
	case Freqency:
		return any(&Freqency{
			// Initialisation of associations
		}).(*Type)
	case Note:
		return any(&Note{
			// Initialisation of associations
			// field is initialized with an instance of Freqency with the name of the field
			Frequencies: []*Freqency{{Name: "Frequencies"}},
		}).(*Type)
	case Player:
		return any(&Player{
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
	// reverse maps of direct associations of Freqency
	case Freqency:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Note
	case Note:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Player
	case Player:
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
	// reverse maps of direct associations of Freqency
	case Freqency:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Note
	case Note:
		switch fieldname {
		// insertion point for per direct association field
		case "Frequencies":
			res := make(map[*Freqency][]*Note)
			for note := range stage.Notes {
				for _, freqency_ := range note.Frequencies {
					res[freqency_] = append(res[freqency_], note)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Player
	case Player:
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
	case *Freqency:
		res = "Freqency"
	case *Note:
		res = "Note"
	case *Player:
		res = "Player"
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
	case *Freqency:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Note"
		rf.Fieldname = "Frequencies"
		res = append(res, rf)
	case *Note:
		var rf ReverseField
		_ = rf
	case *Player:
		var rf ReverseField
		_ = rf
	}
	return
}

// insertion point for get fields header method
func (freqency *Freqency) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (note *Note) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "Frequencies",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Freqency",
		},
		{
			Name:               "Start",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Duration",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Velocity",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Info",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (player *Player) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Status",
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
func (freqency *Freqency) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = freqency.Name
	}
	return
}
func (note *Note) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = note.Name
	case "Frequencies":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range note.Frequencies {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "Start":
		res.valueString = fmt.Sprintf("%f", note.Start)
		res.valueFloat = note.Start
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Duration":
		res.valueString = fmt.Sprintf("%f", note.Duration)
		res.valueFloat = note.Duration
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Velocity":
		res.valueString = fmt.Sprintf("%f", note.Velocity)
		res.valueFloat = note.Velocity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Info":
		res.valueString = note.Info
	}
	return
}
func (player *Player) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = player.Name
	case "Status":
		enum := player.Status
		res.valueString = enum.ToCodeString()
	}
	return
}
func GetFieldStringValueFromPointer(instance GongstructIF, fieldName string, stage *Stage) (res GongFieldValue) {

	res = instance.GongGetFieldValue(fieldName, stage)
	return
}

// insertion point for generic set gongstruct field value
func (freqency *Freqency) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		freqency.Name = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (note *Note) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		note.Name = value.GetValueString()
	case "Frequencies":
		note.Frequencies = make([]*Freqency, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Freqencys {
					if stage.FreqencyMap_Staged_Order[__instance__] == uint(id) {
						note.Frequencies = append(note.Frequencies, __instance__)
						break
					}
				}
			}
		}
	case "Start":
		note.Start = value.GetValueFloat()
	case "Duration":
		note.Duration = value.GetValueFloat()
	case "Velocity":
		note.Velocity = value.GetValueFloat()
	case "Info":
		note.Info = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (player *Player) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		player.Name = value.GetValueString()
	case "Status":
		player.Status.FromCodeString(value.GetValueString())
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func SetFieldStringValueFromPointer(instance GongstructIF, fieldName string, value GongFieldValue, stage *Stage) error {
	return instance.GongSetFieldValue(fieldName, value, stage)
}

// insertion point for generic get gongstruct name
func (freqency *Freqency) GongGetGongstructName() string {
	return "Freqency"
}

func (note *Note) GongGetGongstructName() string {
	return "Note"
}

func (player *Player) GongGetGongstructName() string {
	return "Player"
}

func GetGongstructNameFromPointer(instance GongstructIF) (res string) {
	res = instance.GongGetGongstructName()
	return
}

func (stage *Stage) ResetMapStrings() {

	// insertion point for generic get gongstruct name
	stage.Freqencys_mapString = make(map[string]*Freqency)
	for freqency := range stage.Freqencys {
		stage.Freqencys_mapString[freqency.Name] = freqency
	}

	stage.Notes_mapString = make(map[string]*Note)
	for note := range stage.Notes {
		stage.Notes_mapString[note.Name] = note
	}

	stage.Players_mapString = make(map[string]*Player)
	for player := range stage.Players {
		stage.Players_mapString[player.Name] = player
	}

}

// Last line of the template
