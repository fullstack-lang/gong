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

	test_go "github.com/fullstack-lang/gong/test/test/go"
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

const ProbeTreeSidebarSuffix = ":sidebar of the probe"
const ProbeTableSuffix = ":table of the probe"
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

func (stage *Stage) GetProbeSplitStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeSplitSuffix
}

// errUnkownEnum is returns when a value cannot match enum values
var errUnkownEnum = errors.New("unkown enum")

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
	name               string
	commitId           uint // commitId is updated at each commit
	commitTimeStamp    time.Time
	contentWhenParsed  string
	commitIdWhenParsed uint
	generatesDiff      bool

	// insertion point for definition of arrays registering instances
	Astructs           map[*Astruct]any
	Astructs_mapString map[string]*Astruct

	// insertion point for slice of pointers maps
	Astruct_Anarrayofb_reverseMap map[*Bstruct]*Astruct

	Astruct_Dstruct4s_reverseMap map[*Dstruct]*Astruct

	Astruct_Anarrayofa_reverseMap map[*Astruct]*Astruct

	Astruct_Anotherarrayofb_reverseMap map[*Bstruct]*Astruct

	Astruct_AnarrayofbUse_reverseMap map[*AstructBstructUse]*Astruct

	Astruct_Anarrayofb2Use_reverseMap map[*AstructBstruct2Use]*Astruct

	OnAfterAstructCreateCallback OnAfterCreateInterface[Astruct]
	OnAfterAstructUpdateCallback OnAfterUpdateInterface[Astruct]
	OnAfterAstructDeleteCallback OnAfterDeleteInterface[Astruct]
	OnAfterAstructReadCallback   OnAfterReadInterface[Astruct]

	AstructBstruct2Uses           map[*AstructBstruct2Use]any
	AstructBstruct2Uses_mapString map[string]*AstructBstruct2Use

	// insertion point for slice of pointers maps
	OnAfterAstructBstruct2UseCreateCallback OnAfterCreateInterface[AstructBstruct2Use]
	OnAfterAstructBstruct2UseUpdateCallback OnAfterUpdateInterface[AstructBstruct2Use]
	OnAfterAstructBstruct2UseDeleteCallback OnAfterDeleteInterface[AstructBstruct2Use]
	OnAfterAstructBstruct2UseReadCallback   OnAfterReadInterface[AstructBstruct2Use]

	AstructBstructUses           map[*AstructBstructUse]any
	AstructBstructUses_mapString map[string]*AstructBstructUse

	// insertion point for slice of pointers maps
	OnAfterAstructBstructUseCreateCallback OnAfterCreateInterface[AstructBstructUse]
	OnAfterAstructBstructUseUpdateCallback OnAfterUpdateInterface[AstructBstructUse]
	OnAfterAstructBstructUseDeleteCallback OnAfterDeleteInterface[AstructBstructUse]
	OnAfterAstructBstructUseReadCallback   OnAfterReadInterface[AstructBstructUse]

	Bstructs           map[*Bstruct]any
	Bstructs_mapString map[string]*Bstruct

	// insertion point for slice of pointers maps
	OnAfterBstructCreateCallback OnAfterCreateInterface[Bstruct]
	OnAfterBstructUpdateCallback OnAfterUpdateInterface[Bstruct]
	OnAfterBstructDeleteCallback OnAfterDeleteInterface[Bstruct]
	OnAfterBstructReadCallback   OnAfterReadInterface[Bstruct]

	Dstructs           map[*Dstruct]any
	Dstructs_mapString map[string]*Dstruct

	// insertion point for slice of pointers maps
	Dstruct_Anarrayofb_reverseMap map[*Bstruct]*Dstruct

	Dstruct_Gstructs_reverseMap map[*Gstruct]*Dstruct

	OnAfterDstructCreateCallback OnAfterCreateInterface[Dstruct]
	OnAfterDstructUpdateCallback OnAfterUpdateInterface[Dstruct]
	OnAfterDstructDeleteCallback OnAfterDeleteInterface[Dstruct]
	OnAfterDstructReadCallback   OnAfterReadInterface[Dstruct]

	F0123456789012345678901234567890s           map[*F0123456789012345678901234567890]any
	F0123456789012345678901234567890s_mapString map[string]*F0123456789012345678901234567890

	// insertion point for slice of pointers maps
	OnAfterF0123456789012345678901234567890CreateCallback OnAfterCreateInterface[F0123456789012345678901234567890]
	OnAfterF0123456789012345678901234567890UpdateCallback OnAfterUpdateInterface[F0123456789012345678901234567890]
	OnAfterF0123456789012345678901234567890DeleteCallback OnAfterDeleteInterface[F0123456789012345678901234567890]
	OnAfterF0123456789012345678901234567890ReadCallback   OnAfterReadInterface[F0123456789012345678901234567890]

	Gstructs           map[*Gstruct]any
	Gstructs_mapString map[string]*Gstruct

	// insertion point for slice of pointers maps
	OnAfterGstructCreateCallback OnAfterCreateInterface[Gstruct]
	OnAfterGstructUpdateCallback OnAfterUpdateInterface[Gstruct]
	OnAfterGstructDeleteCallback OnAfterDeleteInterface[Gstruct]
	OnAfterGstructReadCallback   OnAfterReadInterface[Gstruct]

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
	AstructOrder            uint
	AstructMap_Staged_Order map[*Astruct]uint

	AstructBstruct2UseOrder            uint
	AstructBstruct2UseMap_Staged_Order map[*AstructBstruct2Use]uint

	AstructBstructUseOrder            uint
	AstructBstructUseMap_Staged_Order map[*AstructBstructUse]uint

	BstructOrder            uint
	BstructMap_Staged_Order map[*Bstruct]uint

	DstructOrder            uint
	DstructMap_Staged_Order map[*Dstruct]uint

	F0123456789012345678901234567890Order            uint
	F0123456789012345678901234567890Map_Staged_Order map[*F0123456789012345678901234567890]uint

	GstructOrder            uint
	GstructMap_Staged_Order map[*Gstruct]uint

	// end of insertion point

	NamedStructs []*NamedStruct

	// for the computation of the diff at each commit we need
	reference map[GongstructIF]GongstructIF
	modified  map[GongstructIF]struct{}
	new       map[GongstructIF]struct{}
	deleted   map[GongstructIF]struct{}
}

func (stage *Stage) GetCommitId() uint {
	return stage.commitId
}

func (stage *Stage) GetCommitTS() time.Time {
	return stage.commitTimeStamp
}

func (stage *Stage) SetGeneratesDiff(generatesDiff bool) {
	stage.generatesDiff = generatesDiff
}

// GetNamedStructs implements models.ProbebStage.
func (stage *Stage) GetNamedStructsNames() (res []string) {

	for _, namedStruct := range stage.NamedStructs {
		res = append(res, namedStruct.name)
	}

	return
}

func (stage *Stage) GetReference() map[GongstructIF]GongstructIF {
	return stage.reference
}

func (stage *Stage) GetModified() map[GongstructIF]struct{} {
	return stage.modified
}

func (stage *Stage) GetNew() map[GongstructIF]struct{} {
	return stage.new
}

func (stage *Stage) GetDeleted() map[GongstructIF]struct{} {
	return stage.deleted
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

func GetStructInstancesByOrderAuto[T PointerToGongstruct](stage *Stage) (res []T) {
	var t T
	switch any(t).(type) {
	// insertion point for case
	case *Astruct:
		tmp := GetStructInstancesByOrder(stage.Astructs, stage.AstructMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Astruct implements.
			res = append(res, any(v).(T))
		}
		return res
	case *AstructBstruct2Use:
		tmp := GetStructInstancesByOrder(stage.AstructBstruct2Uses, stage.AstructBstruct2UseMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *AstructBstruct2Use implements.
			res = append(res, any(v).(T))
		}
		return res
	case *AstructBstructUse:
		tmp := GetStructInstancesByOrder(stage.AstructBstructUses, stage.AstructBstructUseMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *AstructBstructUse implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Bstruct:
		tmp := GetStructInstancesByOrder(stage.Bstructs, stage.BstructMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Bstruct implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Dstruct:
		tmp := GetStructInstancesByOrder(stage.Dstructs, stage.DstructMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Dstruct implements.
			res = append(res, any(v).(T))
		}
		return res
	case *F0123456789012345678901234567890:
		tmp := GetStructInstancesByOrder(stage.F0123456789012345678901234567890s, stage.F0123456789012345678901234567890Map_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *F0123456789012345678901234567890 implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Gstruct:
		tmp := GetStructInstancesByOrder(stage.Gstructs, stage.GstructMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Gstruct implements.
			res = append(res, any(v).(T))
		}
		return res

	}
	return
}

func GetStructInstancesByOrder[T PointerToGongstruct](set map[T]any, order map[T]uint) (res []T) {

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
	case "Astruct":
		res = GetNamedStructInstances(stage.Astructs, stage.AstructMap_Staged_Order)
	case "AstructBstruct2Use":
		res = GetNamedStructInstances(stage.AstructBstruct2Uses, stage.AstructBstruct2UseMap_Staged_Order)
	case "AstructBstructUse":
		res = GetNamedStructInstances(stage.AstructBstructUses, stage.AstructBstructUseMap_Staged_Order)
	case "Bstruct":
		res = GetNamedStructInstances(stage.Bstructs, stage.BstructMap_Staged_Order)
	case "Dstruct":
		res = GetNamedStructInstances(stage.Dstructs, stage.DstructMap_Staged_Order)
	case "F0123456789012345678901234567890":
		res = GetNamedStructInstances(stage.F0123456789012345678901234567890s, stage.F0123456789012345678901234567890Map_Staged_Order)
	case "Gstruct":
		res = GetNamedStructInstances(stage.Gstructs, stage.GstructMap_Staged_Order)
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
	return "github.com/fullstack-lang/gong/test/test/go/models"
}

func (stage *Stage) GetMap_GongStructName_InstancesNb() map[string]int {
	return stage.Map_GongStructName_InstancesNb
}

func (stage *Stage) GetModelsEmbededDir() embed.FS {
	return test_go.GoModelsDir
}

func (stage *Stage) GetDigramsEmbededDir() embed.FS {
	return test_go.GoDiagramsDir
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
	CommitAstruct(astruct *Astruct)
	CheckoutAstruct(astruct *Astruct)
	CommitAstructBstruct2Use(astructbstruct2use *AstructBstruct2Use)
	CheckoutAstructBstruct2Use(astructbstruct2use *AstructBstruct2Use)
	CommitAstructBstructUse(astructbstructuse *AstructBstructUse)
	CheckoutAstructBstructUse(astructbstructuse *AstructBstructUse)
	CommitBstruct(bstruct *Bstruct)
	CheckoutBstruct(bstruct *Bstruct)
	CommitDstruct(dstruct *Dstruct)
	CheckoutDstruct(dstruct *Dstruct)
	CommitF0123456789012345678901234567890(f0123456789012345678901234567890 *F0123456789012345678901234567890)
	CheckoutF0123456789012345678901234567890(f0123456789012345678901234567890 *F0123456789012345678901234567890)
	CommitGstruct(gstruct *Gstruct)
	CheckoutGstruct(gstruct *Gstruct)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(name string) (stage *Stage) {

	stage = &Stage{ // insertion point for array initiatialisation
		Astructs:           make(map[*Astruct]any),
		Astructs_mapString: make(map[string]*Astruct),

		AstructBstruct2Uses:           make(map[*AstructBstruct2Use]any),
		AstructBstruct2Uses_mapString: make(map[string]*AstructBstruct2Use),

		AstructBstructUses:           make(map[*AstructBstructUse]any),
		AstructBstructUses_mapString: make(map[string]*AstructBstructUse),

		Bstructs:           make(map[*Bstruct]any),
		Bstructs_mapString: make(map[string]*Bstruct),

		Dstructs:           make(map[*Dstruct]any),
		Dstructs_mapString: make(map[string]*Dstruct),

		F0123456789012345678901234567890s:           make(map[*F0123456789012345678901234567890]any),
		F0123456789012345678901234567890s_mapString: make(map[string]*F0123456789012345678901234567890),

		Gstructs:           make(map[*Gstruct]any),
		Gstructs_mapString: make(map[string]*Gstruct),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		name: name,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		AstructMap_Staged_Order: make(map[*Astruct]uint),

		AstructBstruct2UseMap_Staged_Order: make(map[*AstructBstruct2Use]uint),

		AstructBstructUseMap_Staged_Order: make(map[*AstructBstructUse]uint),

		BstructMap_Staged_Order: make(map[*Bstruct]uint),

		DstructMap_Staged_Order: make(map[*Dstruct]uint),

		F0123456789012345678901234567890Map_Staged_Order: make(map[*F0123456789012345678901234567890]uint),

		GstructMap_Staged_Order: make(map[*Gstruct]uint),

		// end of insertion point

		NamedStructs: []*NamedStruct{ // insertion point for order map initialisations
			{name: "Astruct"},
			{name: "AstructBstruct2Use"},
			{name: "AstructBstructUse"},
			{name: "Bstruct"},
			{name: "Dstruct"},
			{name: "F0123456789012345678901234567890"},
			{name: "Gstruct"},
		}, // end of insertion point

		reference: make(map[GongstructIF]GongstructIF),
		new:       make(map[GongstructIF]struct{}),
		modified:  make(map[GongstructIF]struct{}),
		deleted:   make(map[GongstructIF]struct{}),
	}

	return
}

func GetOrder[Type Gongstruct](stage *Stage, instance *Type) uint {

	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *Astruct:
		return stage.AstructMap_Staged_Order[instance]
	case *AstructBstruct2Use:
		return stage.AstructBstruct2UseMap_Staged_Order[instance]
	case *AstructBstructUse:
		return stage.AstructBstructUseMap_Staged_Order[instance]
	case *Bstruct:
		return stage.BstructMap_Staged_Order[instance]
	case *Dstruct:
		return stage.DstructMap_Staged_Order[instance]
	case *F0123456789012345678901234567890:
		return stage.F0123456789012345678901234567890Map_Staged_Order[instance]
	case *Gstruct:
		return stage.GstructMap_Staged_Order[instance]
	default:
		return 0 // should not happen
	}
}

func GetOrderPointerGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) uint {

	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *Astruct:
		return stage.AstructMap_Staged_Order[instance]
	case *AstructBstruct2Use:
		return stage.AstructBstruct2UseMap_Staged_Order[instance]
	case *AstructBstructUse:
		return stage.AstructBstructUseMap_Staged_Order[instance]
	case *Bstruct:
		return stage.BstructMap_Staged_Order[instance]
	case *Dstruct:
		return stage.DstructMap_Staged_Order[instance]
	case *F0123456789012345678901234567890:
		return stage.F0123456789012345678901234567890Map_Staged_Order[instance]
	case *Gstruct:
		return stage.GstructMap_Staged_Order[instance]
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
	stage.commitId++
	stage.commitTimeStamp = time.Now()

	if stage.BackRepo != nil {
		stage.BackRepo.Commit(stage)
	}
	stage.ComputeInstancesNb()
	stage.ComputeReference()
}

func (stage *Stage) ComputeInstancesNb() {
	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["Astruct"] = len(stage.Astructs)
	stage.Map_GongStructName_InstancesNb["AstructBstruct2Use"] = len(stage.AstructBstruct2Uses)
	stage.Map_GongStructName_InstancesNb["AstructBstructUse"] = len(stage.AstructBstructUses)
	stage.Map_GongStructName_InstancesNb["Bstruct"] = len(stage.Bstructs)
	stage.Map_GongStructName_InstancesNb["Dstruct"] = len(stage.Dstructs)
	stage.Map_GongStructName_InstancesNb["F0123456789012345678901234567890"] = len(stage.F0123456789012345678901234567890s)
	stage.Map_GongStructName_InstancesNb["Gstruct"] = len(stage.Gstructs)
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
// Stage puts astruct to the model stage
func (astruct *Astruct) Stage(stage *Stage) *Astruct {

	if _, ok := stage.Astructs[astruct]; !ok {
		stage.Astructs[astruct] = __member
		stage.AstructMap_Staged_Order[astruct] = stage.AstructOrder
		stage.AstructOrder++
		stage.new[astruct] = struct{}{}
		delete(stage.deleted, astruct)
	} else {
		if _, ok := stage.new[astruct]; !ok {
			stage.modified[astruct] = struct{}{}
		}
	}
	stage.Astructs_mapString[astruct.Name] = astruct

	return astruct
}

// Unstage removes astruct off the model stage
func (astruct *Astruct) Unstage(stage *Stage) *Astruct {
	delete(stage.Astructs, astruct)
	delete(stage.Astructs_mapString, astruct.Name)

	if _, ok := stage.reference[astruct]; ok {
		stage.deleted[astruct] = struct{}{}
	} else {
		delete(stage.new, astruct)
	}
	return astruct
}

// UnstageVoid removes astruct off the model stage
func (astruct *Astruct) UnstageVoid(stage *Stage) {
	delete(stage.Astructs, astruct)
	delete(stage.Astructs_mapString, astruct.Name)
}

// commit astruct to the back repo (if it is already staged)
func (astruct *Astruct) Commit(stage *Stage) *Astruct {
	if _, ok := stage.Astructs[astruct]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitAstruct(astruct)
		}
	}
	return astruct
}

func (astruct *Astruct) CommitVoid(stage *Stage) {
	astruct.Commit(stage)
}

func (astruct *Astruct) StageVoid(stage *Stage) {
	astruct.Stage(stage)
}

// Checkout astruct to the back repo (if it is already staged)
func (astruct *Astruct) Checkout(stage *Stage) *Astruct {
	if _, ok := stage.Astructs[astruct]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutAstruct(astruct)
		}
	}
	return astruct
}

// for satisfaction of GongStruct interface
func (astruct *Astruct) GetName() (res string) {
	return astruct.Name
}

// Stage puts astructbstruct2use to the model stage
func (astructbstruct2use *AstructBstruct2Use) Stage(stage *Stage) *AstructBstruct2Use {

	if _, ok := stage.AstructBstruct2Uses[astructbstruct2use]; !ok {
		stage.AstructBstruct2Uses[astructbstruct2use] = __member
		stage.AstructBstruct2UseMap_Staged_Order[astructbstruct2use] = stage.AstructBstruct2UseOrder
		stage.AstructBstruct2UseOrder++
		stage.new[astructbstruct2use] = struct{}{}
		delete(stage.deleted, astructbstruct2use)
	} else {
		if _, ok := stage.new[astructbstruct2use]; !ok {
			stage.modified[astructbstruct2use] = struct{}{}
		}
	}
	stage.AstructBstruct2Uses_mapString[astructbstruct2use.Name] = astructbstruct2use

	return astructbstruct2use
}

// Unstage removes astructbstruct2use off the model stage
func (astructbstruct2use *AstructBstruct2Use) Unstage(stage *Stage) *AstructBstruct2Use {
	delete(stage.AstructBstruct2Uses, astructbstruct2use)
	delete(stage.AstructBstruct2Uses_mapString, astructbstruct2use.Name)

	if _, ok := stage.reference[astructbstruct2use]; ok {
		stage.deleted[astructbstruct2use] = struct{}{}
	} else {
		delete(stage.new, astructbstruct2use)
	}
	return astructbstruct2use
}

// UnstageVoid removes astructbstruct2use off the model stage
func (astructbstruct2use *AstructBstruct2Use) UnstageVoid(stage *Stage) {
	delete(stage.AstructBstruct2Uses, astructbstruct2use)
	delete(stage.AstructBstruct2Uses_mapString, astructbstruct2use.Name)
}

// commit astructbstruct2use to the back repo (if it is already staged)
func (astructbstruct2use *AstructBstruct2Use) Commit(stage *Stage) *AstructBstruct2Use {
	if _, ok := stage.AstructBstruct2Uses[astructbstruct2use]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitAstructBstruct2Use(astructbstruct2use)
		}
	}
	return astructbstruct2use
}

func (astructbstruct2use *AstructBstruct2Use) CommitVoid(stage *Stage) {
	astructbstruct2use.Commit(stage)
}

func (astructbstruct2use *AstructBstruct2Use) StageVoid(stage *Stage) {
	astructbstruct2use.Stage(stage)
}

// Checkout astructbstruct2use to the back repo (if it is already staged)
func (astructbstruct2use *AstructBstruct2Use) Checkout(stage *Stage) *AstructBstruct2Use {
	if _, ok := stage.AstructBstruct2Uses[astructbstruct2use]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutAstructBstruct2Use(astructbstruct2use)
		}
	}
	return astructbstruct2use
}

// for satisfaction of GongStruct interface
func (astructbstruct2use *AstructBstruct2Use) GetName() (res string) {
	return astructbstruct2use.Name
}

// Stage puts astructbstructuse to the model stage
func (astructbstructuse *AstructBstructUse) Stage(stage *Stage) *AstructBstructUse {

	if _, ok := stage.AstructBstructUses[astructbstructuse]; !ok {
		stage.AstructBstructUses[astructbstructuse] = __member
		stage.AstructBstructUseMap_Staged_Order[astructbstructuse] = stage.AstructBstructUseOrder
		stage.AstructBstructUseOrder++
		stage.new[astructbstructuse] = struct{}{}
		delete(stage.deleted, astructbstructuse)
	} else {
		if _, ok := stage.new[astructbstructuse]; !ok {
			stage.modified[astructbstructuse] = struct{}{}
		}
	}
	stage.AstructBstructUses_mapString[astructbstructuse.Name] = astructbstructuse

	return astructbstructuse
}

// Unstage removes astructbstructuse off the model stage
func (astructbstructuse *AstructBstructUse) Unstage(stage *Stage) *AstructBstructUse {
	delete(stage.AstructBstructUses, astructbstructuse)
	delete(stage.AstructBstructUses_mapString, astructbstructuse.Name)

	if _, ok := stage.reference[astructbstructuse]; ok {
		stage.deleted[astructbstructuse] = struct{}{}
	} else {
		delete(stage.new, astructbstructuse)
	}
	return astructbstructuse
}

// UnstageVoid removes astructbstructuse off the model stage
func (astructbstructuse *AstructBstructUse) UnstageVoid(stage *Stage) {
	delete(stage.AstructBstructUses, astructbstructuse)
	delete(stage.AstructBstructUses_mapString, astructbstructuse.Name)
}

// commit astructbstructuse to the back repo (if it is already staged)
func (astructbstructuse *AstructBstructUse) Commit(stage *Stage) *AstructBstructUse {
	if _, ok := stage.AstructBstructUses[astructbstructuse]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitAstructBstructUse(astructbstructuse)
		}
	}
	return astructbstructuse
}

func (astructbstructuse *AstructBstructUse) CommitVoid(stage *Stage) {
	astructbstructuse.Commit(stage)
}

func (astructbstructuse *AstructBstructUse) StageVoid(stage *Stage) {
	astructbstructuse.Stage(stage)
}

// Checkout astructbstructuse to the back repo (if it is already staged)
func (astructbstructuse *AstructBstructUse) Checkout(stage *Stage) *AstructBstructUse {
	if _, ok := stage.AstructBstructUses[astructbstructuse]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutAstructBstructUse(astructbstructuse)
		}
	}
	return astructbstructuse
}

// for satisfaction of GongStruct interface
func (astructbstructuse *AstructBstructUse) GetName() (res string) {
	return astructbstructuse.Name
}

// Stage puts bstruct to the model stage
func (bstruct *Bstruct) Stage(stage *Stage) *Bstruct {

	if _, ok := stage.Bstructs[bstruct]; !ok {
		stage.Bstructs[bstruct] = __member
		stage.BstructMap_Staged_Order[bstruct] = stage.BstructOrder
		stage.BstructOrder++
		stage.new[bstruct] = struct{}{}
		delete(stage.deleted, bstruct)
	} else {
		if _, ok := stage.new[bstruct]; !ok {
			stage.modified[bstruct] = struct{}{}
		}
	}
	stage.Bstructs_mapString[bstruct.Name] = bstruct

	return bstruct
}

// Unstage removes bstruct off the model stage
func (bstruct *Bstruct) Unstage(stage *Stage) *Bstruct {
	delete(stage.Bstructs, bstruct)
	delete(stage.Bstructs_mapString, bstruct.Name)

	if _, ok := stage.reference[bstruct]; ok {
		stage.deleted[bstruct] = struct{}{}
	} else {
		delete(stage.new, bstruct)
	}
	return bstruct
}

// UnstageVoid removes bstruct off the model stage
func (bstruct *Bstruct) UnstageVoid(stage *Stage) {
	delete(stage.Bstructs, bstruct)
	delete(stage.Bstructs_mapString, bstruct.Name)
}

// commit bstruct to the back repo (if it is already staged)
func (bstruct *Bstruct) Commit(stage *Stage) *Bstruct {
	if _, ok := stage.Bstructs[bstruct]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitBstruct(bstruct)
		}
	}
	return bstruct
}

func (bstruct *Bstruct) CommitVoid(stage *Stage) {
	bstruct.Commit(stage)
}

func (bstruct *Bstruct) StageVoid(stage *Stage) {
	bstruct.Stage(stage)
}

// Checkout bstruct to the back repo (if it is already staged)
func (bstruct *Bstruct) Checkout(stage *Stage) *Bstruct {
	if _, ok := stage.Bstructs[bstruct]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutBstruct(bstruct)
		}
	}
	return bstruct
}

// for satisfaction of GongStruct interface
func (bstruct *Bstruct) GetName() (res string) {
	return bstruct.Name
}

// Stage puts dstruct to the model stage
func (dstruct *Dstruct) Stage(stage *Stage) *Dstruct {

	if _, ok := stage.Dstructs[dstruct]; !ok {
		stage.Dstructs[dstruct] = __member
		stage.DstructMap_Staged_Order[dstruct] = stage.DstructOrder
		stage.DstructOrder++
		stage.new[dstruct] = struct{}{}
		delete(stage.deleted, dstruct)
	} else {
		if _, ok := stage.new[dstruct]; !ok {
			stage.modified[dstruct] = struct{}{}
		}
	}
	stage.Dstructs_mapString[dstruct.Name] = dstruct

	return dstruct
}

// Unstage removes dstruct off the model stage
func (dstruct *Dstruct) Unstage(stage *Stage) *Dstruct {
	delete(stage.Dstructs, dstruct)
	delete(stage.Dstructs_mapString, dstruct.Name)

	if _, ok := stage.reference[dstruct]; ok {
		stage.deleted[dstruct] = struct{}{}
	} else {
		delete(stage.new, dstruct)
	}
	return dstruct
}

// UnstageVoid removes dstruct off the model stage
func (dstruct *Dstruct) UnstageVoid(stage *Stage) {
	delete(stage.Dstructs, dstruct)
	delete(stage.Dstructs_mapString, dstruct.Name)
}

// commit dstruct to the back repo (if it is already staged)
func (dstruct *Dstruct) Commit(stage *Stage) *Dstruct {
	if _, ok := stage.Dstructs[dstruct]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDstruct(dstruct)
		}
	}
	return dstruct
}

func (dstruct *Dstruct) CommitVoid(stage *Stage) {
	dstruct.Commit(stage)
}

func (dstruct *Dstruct) StageVoid(stage *Stage) {
	dstruct.Stage(stage)
}

// Checkout dstruct to the back repo (if it is already staged)
func (dstruct *Dstruct) Checkout(stage *Stage) *Dstruct {
	if _, ok := stage.Dstructs[dstruct]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutDstruct(dstruct)
		}
	}
	return dstruct
}

// for satisfaction of GongStruct interface
func (dstruct *Dstruct) GetName() (res string) {
	return dstruct.Name
}

// Stage puts f0123456789012345678901234567890 to the model stage
func (f0123456789012345678901234567890 *F0123456789012345678901234567890) Stage(stage *Stage) *F0123456789012345678901234567890 {

	if _, ok := stage.F0123456789012345678901234567890s[f0123456789012345678901234567890]; !ok {
		stage.F0123456789012345678901234567890s[f0123456789012345678901234567890] = __member
		stage.F0123456789012345678901234567890Map_Staged_Order[f0123456789012345678901234567890] = stage.F0123456789012345678901234567890Order
		stage.F0123456789012345678901234567890Order++
		stage.new[f0123456789012345678901234567890] = struct{}{}
		delete(stage.deleted, f0123456789012345678901234567890)
	} else {
		if _, ok := stage.new[f0123456789012345678901234567890]; !ok {
			stage.modified[f0123456789012345678901234567890] = struct{}{}
		}
	}
	stage.F0123456789012345678901234567890s_mapString[f0123456789012345678901234567890.Name] = f0123456789012345678901234567890

	return f0123456789012345678901234567890
}

// Unstage removes f0123456789012345678901234567890 off the model stage
func (f0123456789012345678901234567890 *F0123456789012345678901234567890) Unstage(stage *Stage) *F0123456789012345678901234567890 {
	delete(stage.F0123456789012345678901234567890s, f0123456789012345678901234567890)
	delete(stage.F0123456789012345678901234567890s_mapString, f0123456789012345678901234567890.Name)

	if _, ok := stage.reference[f0123456789012345678901234567890]; ok {
		stage.deleted[f0123456789012345678901234567890] = struct{}{}
	} else {
		delete(stage.new, f0123456789012345678901234567890)
	}
	return f0123456789012345678901234567890
}

// UnstageVoid removes f0123456789012345678901234567890 off the model stage
func (f0123456789012345678901234567890 *F0123456789012345678901234567890) UnstageVoid(stage *Stage) {
	delete(stage.F0123456789012345678901234567890s, f0123456789012345678901234567890)
	delete(stage.F0123456789012345678901234567890s_mapString, f0123456789012345678901234567890.Name)
}

// commit f0123456789012345678901234567890 to the back repo (if it is already staged)
func (f0123456789012345678901234567890 *F0123456789012345678901234567890) Commit(stage *Stage) *F0123456789012345678901234567890 {
	if _, ok := stage.F0123456789012345678901234567890s[f0123456789012345678901234567890]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitF0123456789012345678901234567890(f0123456789012345678901234567890)
		}
	}
	return f0123456789012345678901234567890
}

func (f0123456789012345678901234567890 *F0123456789012345678901234567890) CommitVoid(stage *Stage) {
	f0123456789012345678901234567890.Commit(stage)
}

func (f0123456789012345678901234567890 *F0123456789012345678901234567890) StageVoid(stage *Stage) {
	f0123456789012345678901234567890.Stage(stage)
}

// Checkout f0123456789012345678901234567890 to the back repo (if it is already staged)
func (f0123456789012345678901234567890 *F0123456789012345678901234567890) Checkout(stage *Stage) *F0123456789012345678901234567890 {
	if _, ok := stage.F0123456789012345678901234567890s[f0123456789012345678901234567890]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutF0123456789012345678901234567890(f0123456789012345678901234567890)
		}
	}
	return f0123456789012345678901234567890
}

// for satisfaction of GongStruct interface
func (f0123456789012345678901234567890 *F0123456789012345678901234567890) GetName() (res string) {
	return f0123456789012345678901234567890.Name
}

// Stage puts gstruct to the model stage
func (gstruct *Gstruct) Stage(stage *Stage) *Gstruct {

	if _, ok := stage.Gstructs[gstruct]; !ok {
		stage.Gstructs[gstruct] = __member
		stage.GstructMap_Staged_Order[gstruct] = stage.GstructOrder
		stage.GstructOrder++
		stage.new[gstruct] = struct{}{}
		delete(stage.deleted, gstruct)
	} else {
		if _, ok := stage.new[gstruct]; !ok {
			stage.modified[gstruct] = struct{}{}
		}
	}
	stage.Gstructs_mapString[gstruct.Name] = gstruct

	return gstruct
}

// Unstage removes gstruct off the model stage
func (gstruct *Gstruct) Unstage(stage *Stage) *Gstruct {
	delete(stage.Gstructs, gstruct)
	delete(stage.Gstructs_mapString, gstruct.Name)

	if _, ok := stage.reference[gstruct]; ok {
		stage.deleted[gstruct] = struct{}{}
	} else {
		delete(stage.new, gstruct)
	}
	return gstruct
}

// UnstageVoid removes gstruct off the model stage
func (gstruct *Gstruct) UnstageVoid(stage *Stage) {
	delete(stage.Gstructs, gstruct)
	delete(stage.Gstructs_mapString, gstruct.Name)
}

// commit gstruct to the back repo (if it is already staged)
func (gstruct *Gstruct) Commit(stage *Stage) *Gstruct {
	if _, ok := stage.Gstructs[gstruct]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitGstruct(gstruct)
		}
	}
	return gstruct
}

func (gstruct *Gstruct) CommitVoid(stage *Stage) {
	gstruct.Commit(stage)
}

func (gstruct *Gstruct) StageVoid(stage *Stage) {
	gstruct.Stage(stage)
}

// Checkout gstruct to the back repo (if it is already staged)
func (gstruct *Gstruct) Checkout(stage *Stage) *Gstruct {
	if _, ok := stage.Gstructs[gstruct]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutGstruct(gstruct)
		}
	}
	return gstruct
}

// for satisfaction of GongStruct interface
func (gstruct *Gstruct) GetName() (res string) {
	return gstruct.Name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMAstruct(Astruct *Astruct)
	CreateORMAstructBstruct2Use(AstructBstruct2Use *AstructBstruct2Use)
	CreateORMAstructBstructUse(AstructBstructUse *AstructBstructUse)
	CreateORMBstruct(Bstruct *Bstruct)
	CreateORMDstruct(Dstruct *Dstruct)
	CreateORMF0123456789012345678901234567890(F0123456789012345678901234567890 *F0123456789012345678901234567890)
	CreateORMGstruct(Gstruct *Gstruct)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMAstruct(Astruct *Astruct)
	DeleteORMAstructBstruct2Use(AstructBstruct2Use *AstructBstruct2Use)
	DeleteORMAstructBstructUse(AstructBstructUse *AstructBstructUse)
	DeleteORMBstruct(Bstruct *Bstruct)
	DeleteORMDstruct(Dstruct *Dstruct)
	DeleteORMF0123456789012345678901234567890(F0123456789012345678901234567890 *F0123456789012345678901234567890)
	DeleteORMGstruct(Gstruct *Gstruct)
}

func (stage *Stage) Reset() { // insertion point for array reset
	stage.Astructs = make(map[*Astruct]any)
	stage.Astructs_mapString = make(map[string]*Astruct)
	stage.AstructMap_Staged_Order = make(map[*Astruct]uint)
	stage.AstructOrder = 0

	stage.AstructBstruct2Uses = make(map[*AstructBstruct2Use]any)
	stage.AstructBstruct2Uses_mapString = make(map[string]*AstructBstruct2Use)
	stage.AstructBstruct2UseMap_Staged_Order = make(map[*AstructBstruct2Use]uint)
	stage.AstructBstruct2UseOrder = 0

	stage.AstructBstructUses = make(map[*AstructBstructUse]any)
	stage.AstructBstructUses_mapString = make(map[string]*AstructBstructUse)
	stage.AstructBstructUseMap_Staged_Order = make(map[*AstructBstructUse]uint)
	stage.AstructBstructUseOrder = 0

	stage.Bstructs = make(map[*Bstruct]any)
	stage.Bstructs_mapString = make(map[string]*Bstruct)
	stage.BstructMap_Staged_Order = make(map[*Bstruct]uint)
	stage.BstructOrder = 0

	stage.Dstructs = make(map[*Dstruct]any)
	stage.Dstructs_mapString = make(map[string]*Dstruct)
	stage.DstructMap_Staged_Order = make(map[*Dstruct]uint)
	stage.DstructOrder = 0

	stage.F0123456789012345678901234567890s = make(map[*F0123456789012345678901234567890]any)
	stage.F0123456789012345678901234567890s_mapString = make(map[string]*F0123456789012345678901234567890)
	stage.F0123456789012345678901234567890Map_Staged_Order = make(map[*F0123456789012345678901234567890]uint)
	stage.F0123456789012345678901234567890Order = 0

	stage.Gstructs = make(map[*Gstruct]any)
	stage.Gstructs_mapString = make(map[string]*Gstruct)
	stage.GstructMap_Staged_Order = make(map[*Gstruct]uint)
	stage.GstructOrder = 0

	stage.ComputeReference()
}

func (stage *Stage) Nil() { // insertion point for array nil
	stage.Astructs = nil
	stage.Astructs_mapString = nil

	stage.AstructBstruct2Uses = nil
	stage.AstructBstruct2Uses_mapString = nil

	stage.AstructBstructUses = nil
	stage.AstructBstructUses_mapString = nil

	stage.Bstructs = nil
	stage.Bstructs_mapString = nil

	stage.Dstructs = nil
	stage.Dstructs_mapString = nil

	stage.F0123456789012345678901234567890s = nil
	stage.F0123456789012345678901234567890s_mapString = nil

	stage.Gstructs = nil
	stage.Gstructs_mapString = nil

}

func (stage *Stage) Unstage() { // insertion point for array nil
	for astruct := range stage.Astructs {
		astruct.Unstage(stage)
	}

	for astructbstruct2use := range stage.AstructBstruct2Uses {
		astructbstruct2use.Unstage(stage)
	}

	for astructbstructuse := range stage.AstructBstructUses {
		astructbstructuse.Unstage(stage)
	}

	for bstruct := range stage.Bstructs {
		bstruct.Unstage(stage)
	}

	for dstruct := range stage.Dstructs {
		dstruct.Unstage(stage)
	}

	for f0123456789012345678901234567890 := range stage.F0123456789012345678901234567890s {
		f0123456789012345678901234567890.Unstage(stage)
	}

	for gstruct := range stage.Gstructs {
		gstruct.Unstage(stage)
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
	CommitVoid(*Stage)
	StageVoid(*Stage)
	UnstageVoid(stage *Stage)
	GongGetFieldHeaders() []GongFieldHeader
	GongClean(stage *Stage)
	GongGetFieldValue(fieldName string, stage *Stage) GongFieldValue
	GongCopy() GongstructIF
}
type PointerToGongstruct interface {
	GongstructIF
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
	case map[*Astruct]any:
		return any(&stage.Astructs).(*Type)
	case map[*AstructBstruct2Use]any:
		return any(&stage.AstructBstruct2Uses).(*Type)
	case map[*AstructBstructUse]any:
		return any(&stage.AstructBstructUses).(*Type)
	case map[*Bstruct]any:
		return any(&stage.Bstructs).(*Type)
	case map[*Dstruct]any:
		return any(&stage.Dstructs).(*Type)
	case map[*F0123456789012345678901234567890]any:
		return any(&stage.F0123456789012345678901234567890s).(*Type)
	case map[*Gstruct]any:
		return any(&stage.Gstructs).(*Type)
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
	case map[string]*Astruct:
		return any(&stage.Astructs_mapString).(*Type)
	case map[string]*AstructBstruct2Use:
		return any(&stage.AstructBstruct2Uses_mapString).(*Type)
	case map[string]*AstructBstructUse:
		return any(&stage.AstructBstructUses_mapString).(*Type)
	case map[string]*Bstruct:
		return any(&stage.Bstructs_mapString).(*Type)
	case map[string]*Dstruct:
		return any(&stage.Dstructs_mapString).(*Type)
	case map[string]*F0123456789012345678901234567890:
		return any(&stage.F0123456789012345678901234567890s_mapString).(*Type)
	case map[string]*Gstruct:
		return any(&stage.Gstructs_mapString).(*Type)
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
	case Astruct:
		return any(&stage.Astructs).(*map[*Type]any)
	case AstructBstruct2Use:
		return any(&stage.AstructBstruct2Uses).(*map[*Type]any)
	case AstructBstructUse:
		return any(&stage.AstructBstructUses).(*map[*Type]any)
	case Bstruct:
		return any(&stage.Bstructs).(*map[*Type]any)
	case Dstruct:
		return any(&stage.Dstructs).(*map[*Type]any)
	case F0123456789012345678901234567890:
		return any(&stage.F0123456789012345678901234567890s).(*map[*Type]any)
	case Gstruct:
		return any(&stage.Gstructs).(*map[*Type]any)
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
	case *Astruct:
		return any(&stage.Astructs).(*map[Type]any)
	case *AstructBstruct2Use:
		return any(&stage.AstructBstruct2Uses).(*map[Type]any)
	case *AstructBstructUse:
		return any(&stage.AstructBstructUses).(*map[Type]any)
	case *Bstruct:
		return any(&stage.Bstructs).(*map[Type]any)
	case *Dstruct:
		return any(&stage.Dstructs).(*map[Type]any)
	case *F0123456789012345678901234567890:
		return any(&stage.F0123456789012345678901234567890s).(*map[Type]any)
	case *Gstruct:
		return any(&stage.Gstructs).(*map[Type]any)
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
	case Astruct:
		return any(&stage.Astructs_mapString).(*map[string]*Type)
	case AstructBstruct2Use:
		return any(&stage.AstructBstruct2Uses_mapString).(*map[string]*Type)
	case AstructBstructUse:
		return any(&stage.AstructBstructUses_mapString).(*map[string]*Type)
	case Bstruct:
		return any(&stage.Bstructs_mapString).(*map[string]*Type)
	case Dstruct:
		return any(&stage.Dstructs_mapString).(*map[string]*Type)
	case F0123456789012345678901234567890:
		return any(&stage.F0123456789012345678901234567890s_mapString).(*map[string]*Type)
	case Gstruct:
		return any(&stage.Gstructs_mapString).(*map[string]*Type)
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
	case Astruct:
		return any(&Astruct{
			// Initialisation of associations
			// field is initialized with an instance of Bstruct with the name of the field
			Associationtob: &Bstruct{Name: "Associationtob"},
			// field is initialized with an instance of Bstruct with the name of the field
			Anarrayofb: []*Bstruct{{Name: "Anarrayofb"}},
			// field is initialized with an instance of Bstruct with the name of the field
			Anotherassociationtob_2: &Bstruct{Name: "Anotherassociationtob_2"},
			// field is initialized with an instance of Astruct with the name of the field
			Anarrayofa: []*Astruct{{Name: "Anarrayofa"}},
			// field is initialized with an instance of Bstruct with the name of the field
			Anotherarrayofb: []*Bstruct{{Name: "Anotherarrayofb"}},
			// field is initialized with an instance of AstructBstructUse with the name of the field
			AnarrayofbUse: []*AstructBstructUse{{Name: "AnarrayofbUse"}},
			// field is initialized with an instance of AstructBstruct2Use with the name of the field
			Anarrayofb2Use: []*AstructBstruct2Use{{Name: "Anarrayofb2Use"}},
			// field is initialized with an instance of Astruct with the name of the field
			AnAstruct: &Astruct{Name: "AnAstruct"},
			// field is initialized with Cstruct problem with composites

			// field is initialized with Estruct problem with composites

		}).(*Type)
	case AstructBstruct2Use:
		return any(&AstructBstruct2Use{
			// Initialisation of associations
			// field is initialized with an instance of Bstruct with the name of the field
			Bstrcut2: &Bstruct{Name: "Bstrcut2"},
		}).(*Type)
	case AstructBstructUse:
		return any(&AstructBstructUse{
			// Initialisation of associations
			// field is initialized with an instance of Bstruct with the name of the field
			Bstruct2: &Bstruct{Name: "Bstruct2"},
		}).(*Type)
	case Bstruct:
		return any(&Bstruct{
			// Initialisation of associations
		}).(*Type)
	case Dstruct:
		return any(&Dstruct{
			// Initialisation of associations
			// field is initialized with an instance of Bstruct with the name of the field
			Anarrayofb: []*Bstruct{{Name: "Anarrayofb"}},
			// field is initialized with an instance of Gstruct with the name of the field
			Gstruct: &Gstruct{Bstruct: Bstruct{Name: "Gstruct"}},
			// field is initialized with an instance of Gstruct with the name of the field
			Gstructs: []*Gstruct{{Bstruct: Bstruct{Name: "Gstructs"}}},
		}).(*Type)
	case F0123456789012345678901234567890:
		return any(&F0123456789012345678901234567890{
			// Initialisation of associations
		}).(*Type)
	case Gstruct:
		return any(&Gstruct{
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
	// reverse maps of direct associations of Astruct
	case Astruct:
		switch fieldname {
		// insertion point for per direct association field
		case "Associationtob":
			res := make(map[*Bstruct][]*Astruct)
			for astruct := range stage.Astructs {
				if astruct.Associationtob != nil {
					bstruct_ := astruct.Associationtob
					var astructs []*Astruct
					_, ok := res[bstruct_]
					if ok {
						astructs = res[bstruct_]
					} else {
						astructs = make([]*Astruct, 0)
					}
					astructs = append(astructs, astruct)
					res[bstruct_] = astructs
				}
			}
			return any(res).(map[*End][]*Start)
		case "Anotherassociationtob_2":
			res := make(map[*Bstruct][]*Astruct)
			for astruct := range stage.Astructs {
				if astruct.Anotherassociationtob_2 != nil {
					bstruct_ := astruct.Anotherassociationtob_2
					var astructs []*Astruct
					_, ok := res[bstruct_]
					if ok {
						astructs = res[bstruct_]
					} else {
						astructs = make([]*Astruct, 0)
					}
					astructs = append(astructs, astruct)
					res[bstruct_] = astructs
				}
			}
			return any(res).(map[*End][]*Start)
		case "Bstruct":
			res := make(map[*Bstruct][]*Astruct)
			for astruct := range stage.Astructs {
				if astruct.Bstruct != nil {
					bstruct_ := astruct.Bstruct
					var astructs []*Astruct
					_, ok := res[bstruct_]
					if ok {
						astructs = res[bstruct_]
					} else {
						astructs = make([]*Astruct, 0)
					}
					astructs = append(astructs, astruct)
					res[bstruct_] = astructs
				}
			}
			return any(res).(map[*End][]*Start)
		case "Bstruct2":
			res := make(map[*Bstruct][]*Astruct)
			for astruct := range stage.Astructs {
				if astruct.Bstruct2 != nil {
					bstruct_ := astruct.Bstruct2
					var astructs []*Astruct
					_, ok := res[bstruct_]
					if ok {
						astructs = res[bstruct_]
					} else {
						astructs = make([]*Astruct, 0)
					}
					astructs = append(astructs, astruct)
					res[bstruct_] = astructs
				}
			}
			return any(res).(map[*End][]*Start)
		case "Dstruct":
			res := make(map[*Dstruct][]*Astruct)
			for astruct := range stage.Astructs {
				if astruct.Dstruct != nil {
					dstruct_ := astruct.Dstruct
					var astructs []*Astruct
					_, ok := res[dstruct_]
					if ok {
						astructs = res[dstruct_]
					} else {
						astructs = make([]*Astruct, 0)
					}
					astructs = append(astructs, astruct)
					res[dstruct_] = astructs
				}
			}
			return any(res).(map[*End][]*Start)
		case "Dstruct2":
			res := make(map[*Dstruct][]*Astruct)
			for astruct := range stage.Astructs {
				if astruct.Dstruct2 != nil {
					dstruct_ := astruct.Dstruct2
					var astructs []*Astruct
					_, ok := res[dstruct_]
					if ok {
						astructs = res[dstruct_]
					} else {
						astructs = make([]*Astruct, 0)
					}
					astructs = append(astructs, astruct)
					res[dstruct_] = astructs
				}
			}
			return any(res).(map[*End][]*Start)
		case "Dstruct3":
			res := make(map[*Dstruct][]*Astruct)
			for astruct := range stage.Astructs {
				if astruct.Dstruct3 != nil {
					dstruct_ := astruct.Dstruct3
					var astructs []*Astruct
					_, ok := res[dstruct_]
					if ok {
						astructs = res[dstruct_]
					} else {
						astructs = make([]*Astruct, 0)
					}
					astructs = append(astructs, astruct)
					res[dstruct_] = astructs
				}
			}
			return any(res).(map[*End][]*Start)
		case "Dstruct4":
			res := make(map[*Dstruct][]*Astruct)
			for astruct := range stage.Astructs {
				if astruct.Dstruct4 != nil {
					dstruct_ := astruct.Dstruct4
					var astructs []*Astruct
					_, ok := res[dstruct_]
					if ok {
						astructs = res[dstruct_]
					} else {
						astructs = make([]*Astruct, 0)
					}
					astructs = append(astructs, astruct)
					res[dstruct_] = astructs
				}
			}
			return any(res).(map[*End][]*Start)
		case "AnAstruct":
			res := make(map[*Astruct][]*Astruct)
			for astruct := range stage.Astructs {
				if astruct.AnAstruct != nil {
					astruct_ := astruct.AnAstruct
					var astructs []*Astruct
					_, ok := res[astruct_]
					if ok {
						astructs = res[astruct_]
					} else {
						astructs = make([]*Astruct, 0)
					}
					astructs = append(astructs, astruct)
					res[astruct_] = astructs
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of AstructBstruct2Use
	case AstructBstruct2Use:
		switch fieldname {
		// insertion point for per direct association field
		case "Bstrcut2":
			res := make(map[*Bstruct][]*AstructBstruct2Use)
			for astructbstruct2use := range stage.AstructBstruct2Uses {
				if astructbstruct2use.Bstrcut2 != nil {
					bstruct_ := astructbstruct2use.Bstrcut2
					var astructbstruct2uses []*AstructBstruct2Use
					_, ok := res[bstruct_]
					if ok {
						astructbstruct2uses = res[bstruct_]
					} else {
						astructbstruct2uses = make([]*AstructBstruct2Use, 0)
					}
					astructbstruct2uses = append(astructbstruct2uses, astructbstruct2use)
					res[bstruct_] = astructbstruct2uses
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of AstructBstructUse
	case AstructBstructUse:
		switch fieldname {
		// insertion point for per direct association field
		case "Bstruct2":
			res := make(map[*Bstruct][]*AstructBstructUse)
			for astructbstructuse := range stage.AstructBstructUses {
				if astructbstructuse.Bstruct2 != nil {
					bstruct_ := astructbstructuse.Bstruct2
					var astructbstructuses []*AstructBstructUse
					_, ok := res[bstruct_]
					if ok {
						astructbstructuses = res[bstruct_]
					} else {
						astructbstructuses = make([]*AstructBstructUse, 0)
					}
					astructbstructuses = append(astructbstructuses, astructbstructuse)
					res[bstruct_] = astructbstructuses
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Bstruct
	case Bstruct:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Dstruct
	case Dstruct:
		switch fieldname {
		// insertion point for per direct association field
		case "Gstruct":
			res := make(map[*Gstruct][]*Dstruct)
			for dstruct := range stage.Dstructs {
				if dstruct.Gstruct != nil {
					gstruct_ := dstruct.Gstruct
					var dstructs []*Dstruct
					_, ok := res[gstruct_]
					if ok {
						dstructs = res[gstruct_]
					} else {
						dstructs = make([]*Dstruct, 0)
					}
					dstructs = append(dstructs, dstruct)
					res[gstruct_] = dstructs
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of F0123456789012345678901234567890
	case F0123456789012345678901234567890:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Gstruct
	case Gstruct:
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
	// reverse maps of direct associations of Astruct
	case Astruct:
		switch fieldname {
		// insertion point for per direct association field
		case "Anarrayofb":
			res := make(map[*Bstruct][]*Astruct)
			for astruct := range stage.Astructs {
				for _, bstruct_ := range astruct.Anarrayofb {
					res[bstruct_] = append(res[bstruct_], astruct)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Dstruct4s":
			res := make(map[*Dstruct][]*Astruct)
			for astruct := range stage.Astructs {
				for _, dstruct_ := range astruct.Dstruct4s {
					res[dstruct_] = append(res[dstruct_], astruct)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Anarrayofa":
			res := make(map[*Astruct][]*Astruct)
			for astruct := range stage.Astructs {
				for _, astruct_ := range astruct.Anarrayofa {
					res[astruct_] = append(res[astruct_], astruct)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Anotherarrayofb":
			res := make(map[*Bstruct][]*Astruct)
			for astruct := range stage.Astructs {
				for _, bstruct_ := range astruct.Anotherarrayofb {
					res[bstruct_] = append(res[bstruct_], astruct)
				}
			}
			return any(res).(map[*End][]*Start)
		case "AnarrayofbUse":
			res := make(map[*AstructBstructUse][]*Astruct)
			for astruct := range stage.Astructs {
				for _, astructbstructuse_ := range astruct.AnarrayofbUse {
					res[astructbstructuse_] = append(res[astructbstructuse_], astruct)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Anarrayofb2Use":
			res := make(map[*AstructBstruct2Use][]*Astruct)
			for astruct := range stage.Astructs {
				for _, astructbstruct2use_ := range astruct.Anarrayofb2Use {
					res[astructbstruct2use_] = append(res[astructbstruct2use_], astruct)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of AstructBstruct2Use
	case AstructBstruct2Use:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of AstructBstructUse
	case AstructBstructUse:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Bstruct
	case Bstruct:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Dstruct
	case Dstruct:
		switch fieldname {
		// insertion point for per direct association field
		case "Anarrayofb":
			res := make(map[*Bstruct][]*Dstruct)
			for dstruct := range stage.Dstructs {
				for _, bstruct_ := range dstruct.Anarrayofb {
					res[bstruct_] = append(res[bstruct_], dstruct)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Gstructs":
			res := make(map[*Gstruct][]*Dstruct)
			for dstruct := range stage.Dstructs {
				for _, gstruct_ := range dstruct.Gstructs {
					res[gstruct_] = append(res[gstruct_], dstruct)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of F0123456789012345678901234567890
	case F0123456789012345678901234567890:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Gstruct
	case Gstruct:
		switch fieldname {
		// insertion point for per direct association field
		}
	}
	return nil
}

// GetPointerToGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetPointerToGongstructName[Type PointerToGongstruct]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *Astruct:
		res = "Astruct"
	case *AstructBstruct2Use:
		res = "AstructBstruct2Use"
	case *AstructBstructUse:
		res = "AstructBstructUse"
	case *Bstruct:
		res = "Bstruct"
	case *Dstruct:
		res = "Dstruct"
	case *F0123456789012345678901234567890:
		res = "F0123456789012345678901234567890"
	case *Gstruct:
		res = "Gstruct"
	}
	return res
}

type ReverseField struct {
	GongstructName string
	Fieldname      string
}

func GetReverseFields[Type PointerToGongstruct]() (res []ReverseField) {

	res = make([]ReverseField, 0)

	var ret Type

	switch any(ret).(type) {

	// insertion point for generic get gongstruct name
	case *Astruct:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Astruct"
		rf.Fieldname = "Anarrayofa"
		res = append(res, rf)
	case *AstructBstruct2Use:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Astruct"
		rf.Fieldname = "Anarrayofb2Use"
		res = append(res, rf)
	case *AstructBstructUse:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Astruct"
		rf.Fieldname = "AnarrayofbUse"
		res = append(res, rf)
	case *Bstruct:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Astruct"
		rf.Fieldname = "Anarrayofb"
		res = append(res, rf)
		rf.GongstructName = "Astruct"
		rf.Fieldname = "Anotherarrayofb"
		res = append(res, rf)
		rf.GongstructName = "Dstruct"
		rf.Fieldname = "Anarrayofb"
		res = append(res, rf)
	case *Dstruct:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Astruct"
		rf.Fieldname = "Dstruct4s"
		res = append(res, rf)
	case *F0123456789012345678901234567890:
		var rf ReverseField
		_ = rf
	case *Gstruct:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Dstruct"
		rf.Fieldname = "Gstructs"
		res = append(res, rf)
	}
	return
}

// insertion point for get fields header method
func (astruct *Astruct) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Field",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Associationtob",
			GongFieldValueType: GongFieldValueTypePointer,
		},
		{
			Name:               "Anarrayofb",
			GongFieldValueType: GongFieldValueTypeSliceOfPointers,
		},
		{
			Name:               "Anotherassociationtob_2",
			GongFieldValueType: GongFieldValueTypePointer,
		},
		{
			Name:               "Date",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Date2",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Booleanfield",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Aenum",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Aenum_2",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Benum",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "CEnum",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "CName",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "CFloatfield",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Bstruct",
			GongFieldValueType: GongFieldValueTypePointer,
		},
		{
			Name:               "Bstruct2",
			GongFieldValueType: GongFieldValueTypePointer,
		},
		{
			Name:               "Dstruct",
			GongFieldValueType: GongFieldValueTypePointer,
		},
		{
			Name:               "Dstruct2",
			GongFieldValueType: GongFieldValueTypePointer,
		},
		{
			Name:               "Dstruct3",
			GongFieldValueType: GongFieldValueTypePointer,
		},
		{
			Name:               "Dstruct4",
			GongFieldValueType: GongFieldValueTypePointer,
		},
		{
			Name:               "Dstruct4s",
			GongFieldValueType: GongFieldValueTypeSliceOfPointers,
		},
		{
			Name:               "Floatfield",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Intfield",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Anotherbooleanfield",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Duration1",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Anarrayofa",
			GongFieldValueType: GongFieldValueTypeSliceOfPointers,
		},
		{
			Name:               "Anotherarrayofb",
			GongFieldValueType: GongFieldValueTypeSliceOfPointers,
		},
		{
			Name:               "AnarrayofbUse",
			GongFieldValueType: GongFieldValueTypeSliceOfPointers,
		},
		{
			Name:               "Anarrayofb2Use",
			GongFieldValueType: GongFieldValueTypeSliceOfPointers,
		},
		{
			Name:               "AnAstruct",
			GongFieldValueType: GongFieldValueTypePointer,
		},
		{
			Name:               "TextFieldBespokeSize",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "TextArea",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (astructbstruct2use *AstructBstruct2Use) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Bstrcut2",
			GongFieldValueType: GongFieldValueTypePointer,
		},
	}
	return
}

func (astructbstructuse *AstructBstructUse) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Bstruct2",
			GongFieldValueType: GongFieldValueTypePointer,
		},
	}
	return
}

func (bstruct *Bstruct) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Floatfield",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Floatfield2",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Intfield",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (dstruct *Dstruct) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Anarrayofb",
			GongFieldValueType: GongFieldValueTypeSliceOfPointers,
		},
		{
			Name:               "Gstruct",
			GongFieldValueType: GongFieldValueTypePointer,
		},
		{
			Name:               "Gstructs",
			GongFieldValueType: GongFieldValueTypeSliceOfPointers,
		},
	}
	return
}

func (f0123456789012345678901234567890 *F0123456789012345678901234567890) GongGetFieldHeaders() (res []GongFieldHeader) {
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
	}
	return
}

func (gstruct *Gstruct) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Floatfield",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Floatfield2",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Intfield",
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
	GongFieldValueType
	Name string
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
func (astruct *Astruct) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = astruct.Name
	case "Associationtob":
		res.GongFieldValueType = GongFieldValueTypePointer
		if astruct.Associationtob != nil {
			res.valueString = astruct.Associationtob.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, astruct.Associationtob))
		}
	case "Anarrayofb":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range astruct.Anarrayofb {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "Anotherassociationtob_2":
		res.GongFieldValueType = GongFieldValueTypePointer
		if astruct.Anotherassociationtob_2 != nil {
			res.valueString = astruct.Anotherassociationtob_2.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, astruct.Anotherassociationtob_2))
		}
	case "Date":
		res.valueString = astruct.Date.String()
	case "Date2":
		res.valueString = astruct.Date2.Format("2006-01-02")
	case "Booleanfield":
		res.valueString = fmt.Sprintf("%t", astruct.Booleanfield)
		res.valueBool = astruct.Booleanfield
		res.GongFieldValueType = GongFieldValueTypeBool
	case "Aenum":
		enum := astruct.Aenum
		res.valueString = enum.ToCodeString()
	case "Aenum_2":
		enum := astruct.Aenum_2
		res.valueString = enum.ToCodeString()
	case "Benum":
		enum := astruct.Benum
		res.valueString = enum.ToCodeString()
	case "CEnum":
		enum := astruct.CEnum
		res.valueString = enum.ToCodeString()
	case "CName":
		res.valueString = astruct.CName
	case "CFloatfield":
		res.valueString = fmt.Sprintf("%f", astruct.CFloatfield)
		res.valueFloat = astruct.CFloatfield
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Bstruct":
		res.GongFieldValueType = GongFieldValueTypePointer
		if astruct.Bstruct != nil {
			res.valueString = astruct.Bstruct.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, astruct.Bstruct))
		}
	case "Bstruct2":
		res.GongFieldValueType = GongFieldValueTypePointer
		if astruct.Bstruct2 != nil {
			res.valueString = astruct.Bstruct2.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, astruct.Bstruct2))
		}
	case "Dstruct":
		res.GongFieldValueType = GongFieldValueTypePointer
		if astruct.Dstruct != nil {
			res.valueString = astruct.Dstruct.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, astruct.Dstruct))
		}
	case "Dstruct2":
		res.GongFieldValueType = GongFieldValueTypePointer
		if astruct.Dstruct2 != nil {
			res.valueString = astruct.Dstruct2.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, astruct.Dstruct2))
		}
	case "Dstruct3":
		res.GongFieldValueType = GongFieldValueTypePointer
		if astruct.Dstruct3 != nil {
			res.valueString = astruct.Dstruct3.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, astruct.Dstruct3))
		}
	case "Dstruct4":
		res.GongFieldValueType = GongFieldValueTypePointer
		if astruct.Dstruct4 != nil {
			res.valueString = astruct.Dstruct4.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, astruct.Dstruct4))
		}
	case "Dstruct4s":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range astruct.Dstruct4s {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "Floatfield":
		res.valueString = fmt.Sprintf("%f", astruct.Floatfield)
		res.valueFloat = astruct.Floatfield
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Intfield":
		res.valueString = fmt.Sprintf("%d", astruct.Intfield)
		res.valueInt = astruct.Intfield
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Anotherbooleanfield":
		res.valueString = fmt.Sprintf("%t", astruct.Anotherbooleanfield)
		res.valueBool = astruct.Anotherbooleanfield
		res.GongFieldValueType = GongFieldValueTypeBool
	case "Duration1":
		if math.Abs(astruct.Duration1.Hours()) >= 24 {
			days := __Gong__Abs(int(int(astruct.Duration1.Hours()) / 24))
			months := int(days / 31)
			days = days - months*31

			remainingHours := int(astruct.Duration1.Hours()) % 24
			remainingMinutes := int(astruct.Duration1.Minutes()) % 60
			remainingSeconds := int(astruct.Duration1.Seconds()) % 60

			if astruct.Duration1.Hours() < 0 {
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
			res.valueString = fmt.Sprintf("%s\n", astruct.Duration1.String())
		}
	case "Anarrayofa":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range astruct.Anarrayofa {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "Anotherarrayofb":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range astruct.Anotherarrayofb {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "AnarrayofbUse":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range astruct.AnarrayofbUse {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "Anarrayofb2Use":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range astruct.Anarrayofb2Use {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "AnAstruct":
		res.GongFieldValueType = GongFieldValueTypePointer
		if astruct.AnAstruct != nil {
			res.valueString = astruct.AnAstruct.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, astruct.AnAstruct))
		}
	case "TextFieldBespokeSize":
		res.valueString = astruct.TextFieldBespokeSize
	case "TextArea":
		res.valueString = astruct.TextArea
	}
	return
}
func (astructbstruct2use *AstructBstruct2Use) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = astructbstruct2use.Name
	case "Bstrcut2":
		res.GongFieldValueType = GongFieldValueTypePointer
		if astructbstruct2use.Bstrcut2 != nil {
			res.valueString = astructbstruct2use.Bstrcut2.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, astructbstruct2use.Bstrcut2))
		}
	}
	return
}
func (astructbstructuse *AstructBstructUse) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = astructbstructuse.Name
	case "Bstruct2":
		res.GongFieldValueType = GongFieldValueTypePointer
		if astructbstructuse.Bstruct2 != nil {
			res.valueString = astructbstructuse.Bstruct2.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, astructbstructuse.Bstruct2))
		}
	}
	return
}
func (bstruct *Bstruct) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = bstruct.Name
	case "Floatfield":
		res.valueString = fmt.Sprintf("%f", bstruct.Floatfield)
		res.valueFloat = bstruct.Floatfield
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Floatfield2":
		res.valueString = fmt.Sprintf("%f", bstruct.Floatfield2)
		res.valueFloat = bstruct.Floatfield2
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Intfield":
		res.valueString = fmt.Sprintf("%d", bstruct.Intfield)
		res.valueInt = bstruct.Intfield
		res.GongFieldValueType = GongFieldValueTypeInt
	}
	return
}
func (dstruct *Dstruct) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = dstruct.Name
	case "Anarrayofb":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range dstruct.Anarrayofb {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "Gstruct":
		res.GongFieldValueType = GongFieldValueTypePointer
		if dstruct.Gstruct != nil {
			res.valueString = dstruct.Gstruct.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, dstruct.Gstruct))
		}
	case "Gstructs":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range dstruct.Gstructs {
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
func (f0123456789012345678901234567890 *F0123456789012345678901234567890) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = f0123456789012345678901234567890.Name
	case "Date":
		res.valueString = f0123456789012345678901234567890.Date.String()
	}
	return
}
func (gstruct *Gstruct) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = gstruct.Name
	case "Floatfield":
		res.valueString = fmt.Sprintf("%f", gstruct.Floatfield)
		res.valueFloat = gstruct.Floatfield
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Floatfield2":
		res.valueString = fmt.Sprintf("%f", gstruct.Floatfield2)
		res.valueFloat = gstruct.Floatfield2
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Intfield":
		res.valueString = fmt.Sprintf("%d", gstruct.Intfield)
		res.valueInt = gstruct.Intfield
		res.GongFieldValueType = GongFieldValueTypeInt
	}
	return
}

func GetFieldStringValueFromPointer(instance GongstructIF, fieldName string, stage *Stage) (res GongFieldValue) {

	res = instance.GongGetFieldValue(fieldName, stage)
	return
}

// Last line of the template
