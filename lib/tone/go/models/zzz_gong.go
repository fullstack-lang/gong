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

	tone_go "github.com/fullstack-lang/gong/lib/tone/go"
)

func __Gong__Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

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
	Freqencys           map[*Freqency]any
	Freqencys_mapString map[string]*Freqency

	// insertion point for slice of pointers maps
	OnAfterFreqencyCreateCallback OnAfterCreateInterface[Freqency]
	OnAfterFreqencyUpdateCallback OnAfterUpdateInterface[Freqency]
	OnAfterFreqencyDeleteCallback OnAfterDeleteInterface[Freqency]
	OnAfterFreqencyReadCallback   OnAfterReadInterface[Freqency]

	Notes           map[*Note]any
	Notes_mapString map[string]*Note

	// insertion point for slice of pointers maps
	Note_Frequencies_reverseMap map[*Freqency]*Note

	OnAfterNoteCreateCallback OnAfterCreateInterface[Note]
	OnAfterNoteUpdateCallback OnAfterUpdateInterface[Note]
	OnAfterNoteDeleteCallback OnAfterDeleteInterface[Note]
	OnAfterNoteReadCallback   OnAfterReadInterface[Note]

	Players           map[*Player]any
	Players_mapString map[string]*Player

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
		Freqencys:           make(map[*Freqency]any),
		Freqencys_mapString: make(map[string]*Freqency),

		Notes:           make(map[*Note]any),
		Notes_mapString: make(map[string]*Note),

		Players:           make(map[*Player]any),
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

		NamedStructs: []*NamedStruct{ // insertion point for order map initialisations
			{name: "Freqency"},
			{name: "Note"},
			{name: "Player"},
		}, // end of insertion point
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
	stage.Map_GongStructName_InstancesNb["Freqency"] = len(stage.Freqencys)
	stage.Map_GongStructName_InstancesNb["Note"] = len(stage.Notes)
	stage.Map_GongStructName_InstancesNb["Player"] = len(stage.Players)

}

func (stage *Stage) Checkout() {
	if stage.BackRepo != nil {
		stage.BackRepo.Checkout(stage)
	}

	stage.ComputeReverseMaps()
	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["Freqency"] = len(stage.Freqencys)
	stage.Map_GongStructName_InstancesNb["Note"] = len(stage.Notes)
	stage.Map_GongStructName_InstancesNb["Player"] = len(stage.Players)

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
		stage.Freqencys[freqency] = __member
		stage.FreqencyMap_Staged_Order[freqency] = stage.FreqencyOrder
		stage.FreqencyOrder++
	}
	stage.Freqencys_mapString[freqency.Name] = freqency

	return freqency
}

// Unstage removes freqency off the model stage
func (freqency *Freqency) Unstage(stage *Stage) *Freqency {
	delete(stage.Freqencys, freqency)
	delete(stage.Freqencys_mapString, freqency.Name)
	return freqency
}

// UnstageVoid removes freqency off the model stage
func (freqency *Freqency) UnstageVoid(stage *Stage) {
	delete(stage.Freqencys, freqency)
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

// Stage puts note to the model stage
func (note *Note) Stage(stage *Stage) *Note {

	if _, ok := stage.Notes[note]; !ok {
		stage.Notes[note] = __member
		stage.NoteMap_Staged_Order[note] = stage.NoteOrder
		stage.NoteOrder++
	}
	stage.Notes_mapString[note.Name] = note

	return note
}

// Unstage removes note off the model stage
func (note *Note) Unstage(stage *Stage) *Note {
	delete(stage.Notes, note)
	delete(stage.Notes_mapString, note.Name)
	return note
}

// UnstageVoid removes note off the model stage
func (note *Note) UnstageVoid(stage *Stage) {
	delete(stage.Notes, note)
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

// Stage puts player to the model stage
func (player *Player) Stage(stage *Stage) *Player {

	if _, ok := stage.Players[player]; !ok {
		stage.Players[player] = __member
		stage.PlayerMap_Staged_Order[player] = stage.PlayerOrder
		stage.PlayerOrder++
	}
	stage.Players_mapString[player.Name] = player

	return player
}

// Unstage removes player off the model stage
func (player *Player) Unstage(stage *Stage) *Player {
	delete(stage.Players, player)
	delete(stage.Players_mapString, player.Name)
	return player
}

// UnstageVoid removes player off the model stage
func (player *Player) UnstageVoid(stage *Stage) {
	delete(stage.Players, player)
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
	stage.Freqencys = make(map[*Freqency]any)
	stage.Freqencys_mapString = make(map[string]*Freqency)
	stage.FreqencyMap_Staged_Order = make(map[*Freqency]uint)
	stage.FreqencyOrder = 0

	stage.Notes = make(map[*Note]any)
	stage.Notes_mapString = make(map[string]*Note)
	stage.NoteMap_Staged_Order = make(map[*Note]uint)
	stage.NoteOrder = 0

	stage.Players = make(map[*Player]any)
	stage.Players_mapString = make(map[string]*Player)
	stage.PlayerMap_Staged_Order = make(map[*Player]uint)
	stage.PlayerOrder = 0

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

// GongGetMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GongGetMap[Type GongstructMapString](stage *Stage) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case map[string]*Freqency:
		return any(&stage.Freqencys_mapString).(*Type)
	case map[string]*Note:
		return any(&stage.Notes_mapString).(*Type)
	case map[string]*Player:
		return any(&stage.Players_mapString).(*Type)
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
	case Freqency:
		return any(&stage.Freqencys).(*map[*Type]any)
	case Note:
		return any(&stage.Notes).(*map[*Type]any)
	case Player:
		return any(&stage.Players).(*map[*Type]any)
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
	case *Freqency:
		return any(&stage.Freqencys).(*map[Type]any)
	case *Note:
		return any(&stage.Notes).(*map[Type]any)
	case *Player:
		return any(&stage.Players).(*map[Type]any)
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
func GetSliceOfPointersReverseMap[Start, End Gongstruct](fieldname string, stage *Stage) map[*End]*Start {

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
			res := make(map[*Freqency]*Note)
			for note := range stage.Notes {
				for _, freqency_ := range note.Frequencies {
					res[freqency_] = note
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of Player
	case Player:
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
	case Freqency:
		res = "Freqency"
	case Note:
		res = "Note"
	case Player:
		res = "Player"
	}
	return res
}

// GetPointerToGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetPointerToGongstructName[Type PointerToGongstruct]() (res string) {

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

// GetFields return the array of the fields
func GetFields[Type Gongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case Freqency:
		res = []string{"Name"}
	case Note:
		res = []string{"Name", "Frequencies", "Start", "Duration", "Velocity", "Info"}
	case Player:
		res = []string{"Name", "Status"}
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
	case Freqency:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Note"
		rf.Fieldname = "Frequencies"
		res = append(res, rf)
	case Note:
		var rf ReverseField
		_ = rf
	case Player:
		var rf ReverseField
		_ = rf
	}
	return
}

// GetFieldsFromPointer return the array of the fields
func GetFieldsFromPointer[Type PointerToGongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *Freqency:
		res = []string{"Name"}
	case *Note:
		res = []string{"Name", "Frequencies", "Start", "Duration", "Velocity", "Info"}
	case *Player:
		res = []string{"Name", "Status"}
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
	case *Freqency:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case *Note:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Frequencies":
			for idx, __instance__ := range inferedInstance.Frequencies {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "Start":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Start)
			res.valueFloat = inferedInstance.Start
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Duration":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Duration)
			res.valueFloat = inferedInstance.Duration
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Velocity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Velocity)
			res.valueFloat = inferedInstance.Velocity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Info":
			res.valueString = inferedInstance.Info
		}
	case *Player:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Status":
			enum := inferedInstance.Status
			res.valueString = enum.ToCodeString()
		}
	default:
		_ = inferedInstance
	}
	return
}

func GetFieldStringValue(instance any, fieldName string) (res GongFieldValue) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case Freqency:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case Note:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Frequencies":
			for idx, __instance__ := range inferedInstance.Frequencies {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "Start":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Start)
			res.valueFloat = inferedInstance.Start
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Duration":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Duration)
			res.valueFloat = inferedInstance.Duration
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Velocity":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Velocity)
			res.valueFloat = inferedInstance.Velocity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Info":
			res.valueString = inferedInstance.Info
		}
	case Player:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Status":
			enum := inferedInstance.Status
			res.valueString = enum.ToCodeString()
		}
	default:
		_ = inferedInstance
	}
	return
}

// Last line of the template
