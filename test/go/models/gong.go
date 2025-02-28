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
	path string

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

	Fstructs           map[*Fstruct]any
	Fstructs_mapString map[string]*Fstruct

	// insertion point for slice of pointers maps
	OnAfterFstructCreateCallback OnAfterCreateInterface[Fstruct]
	OnAfterFstructUpdateCallback OnAfterUpdateInterface[Fstruct]
	OnAfterFstructDeleteCallback OnAfterDeleteInterface[Fstruct]
	OnAfterFstructReadCallback   OnAfterReadInterface[Fstruct]

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
	Order            uint
	Map_Staged_Order map[any]uint
}

func (stage *StageStruct) GetType() string {
	return "github.com/fullstack-lang/gong/test/go/models"
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
	CommitFstruct(fstruct *Fstruct)
	CheckoutFstruct(fstruct *Fstruct)
	CommitGstruct(gstruct *Gstruct)
	CheckoutGstruct(gstruct *Gstruct)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(path string) (stage *StageStruct) {

	stage = &StageStruct{ // insertion point for array initiatialisation
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

		Fstructs:           make(map[*Fstruct]any),
		Fstructs_mapString: make(map[string]*Fstruct),

		Gstructs:           make(map[*Gstruct]any),
		Gstructs_mapString: make(map[string]*Gstruct),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		path: path,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		Map_Staged_Order: make(map[any]uint),
	}

	return
}

func (stage *StageStruct) GetPath() string {
	return stage.path
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
	stage.Map_GongStructName_InstancesNb["Astruct"] = len(stage.Astructs)
	stage.Map_GongStructName_InstancesNb["AstructBstruct2Use"] = len(stage.AstructBstruct2Uses)
	stage.Map_GongStructName_InstancesNb["AstructBstructUse"] = len(stage.AstructBstructUses)
	stage.Map_GongStructName_InstancesNb["Bstruct"] = len(stage.Bstructs)
	stage.Map_GongStructName_InstancesNb["Dstruct"] = len(stage.Dstructs)
	stage.Map_GongStructName_InstancesNb["Fstruct"] = len(stage.Fstructs)
	stage.Map_GongStructName_InstancesNb["Gstruct"] = len(stage.Gstructs)

}

func (stage *StageStruct) Checkout() {
	if stage.BackRepo != nil {
		stage.BackRepo.Checkout(stage)
	}

	stage.ComputeReverseMaps()
	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["Astruct"] = len(stage.Astructs)
	stage.Map_GongStructName_InstancesNb["AstructBstruct2Use"] = len(stage.AstructBstruct2Uses)
	stage.Map_GongStructName_InstancesNb["AstructBstructUse"] = len(stage.AstructBstructUses)
	stage.Map_GongStructName_InstancesNb["Bstruct"] = len(stage.Bstructs)
	stage.Map_GongStructName_InstancesNb["Dstruct"] = len(stage.Dstructs)
	stage.Map_GongStructName_InstancesNb["Fstruct"] = len(stage.Fstructs)
	stage.Map_GongStructName_InstancesNb["Gstruct"] = len(stage.Gstructs)

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
// Stage puts astruct to the model stage
func (astruct *Astruct) Stage(stage *StageStruct) *Astruct {

	if _, ok := stage.Astructs[astruct]; !ok {
		stage.Astructs[astruct] = __member
		stage.Map_Staged_Order[astruct] = stage.Order
		stage.Order++
	}
	stage.Astructs_mapString[astruct.Name] = astruct

	return astruct
}

// Unstage removes astruct off the model stage
func (astruct *Astruct) Unstage(stage *StageStruct) *Astruct {
	delete(stage.Astructs, astruct)
	delete(stage.Astructs_mapString, astruct.Name)
	return astruct
}

// UnstageVoid removes astruct off the model stage
func (astruct *Astruct) UnstageVoid(stage *StageStruct) {
	delete(stage.Astructs, astruct)
	delete(stage.Astructs_mapString, astruct.Name)
}

// commit astruct to the back repo (if it is already staged)
func (astruct *Astruct) Commit(stage *StageStruct) *Astruct {
	if _, ok := stage.Astructs[astruct]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitAstruct(astruct)
		}
	}
	return astruct
}

func (astruct *Astruct) CommitVoid(stage *StageStruct) {
	astruct.Commit(stage)
}

// Checkout astruct to the back repo (if it is already staged)
func (astruct *Astruct) Checkout(stage *StageStruct) *Astruct {
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
func (astructbstruct2use *AstructBstruct2Use) Stage(stage *StageStruct) *AstructBstruct2Use {
	stage.AstructBstruct2Uses[astructbstruct2use] = __member
	stage.AstructBstruct2Uses_mapString[astructbstruct2use.Name] = astructbstruct2use

	return astructbstruct2use
}

// Unstage removes astructbstruct2use off the model stage
func (astructbstruct2use *AstructBstruct2Use) Unstage(stage *StageStruct) *AstructBstruct2Use {
	delete(stage.AstructBstruct2Uses, astructbstruct2use)
	delete(stage.AstructBstruct2Uses_mapString, astructbstruct2use.Name)
	return astructbstruct2use
}

// UnstageVoid removes astructbstruct2use off the model stage
func (astructbstruct2use *AstructBstruct2Use) UnstageVoid(stage *StageStruct) {
	delete(stage.AstructBstruct2Uses, astructbstruct2use)
	delete(stage.AstructBstruct2Uses_mapString, astructbstruct2use.Name)
}

// commit astructbstruct2use to the back repo (if it is already staged)
func (astructbstruct2use *AstructBstruct2Use) Commit(stage *StageStruct) *AstructBstruct2Use {
	if _, ok := stage.AstructBstruct2Uses[astructbstruct2use]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitAstructBstruct2Use(astructbstruct2use)
		}
	}
	return astructbstruct2use
}

func (astructbstruct2use *AstructBstruct2Use) CommitVoid(stage *StageStruct) {
	astructbstruct2use.Commit(stage)
}

// Checkout astructbstruct2use to the back repo (if it is already staged)
func (astructbstruct2use *AstructBstruct2Use) Checkout(stage *StageStruct) *AstructBstruct2Use {
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
func (astructbstructuse *AstructBstructUse) Stage(stage *StageStruct) *AstructBstructUse {
	stage.AstructBstructUses[astructbstructuse] = __member
	stage.AstructBstructUses_mapString[astructbstructuse.Name] = astructbstructuse

	return astructbstructuse
}

// Unstage removes astructbstructuse off the model stage
func (astructbstructuse *AstructBstructUse) Unstage(stage *StageStruct) *AstructBstructUse {
	delete(stage.AstructBstructUses, astructbstructuse)
	delete(stage.AstructBstructUses_mapString, astructbstructuse.Name)
	return astructbstructuse
}

// UnstageVoid removes astructbstructuse off the model stage
func (astructbstructuse *AstructBstructUse) UnstageVoid(stage *StageStruct) {
	delete(stage.AstructBstructUses, astructbstructuse)
	delete(stage.AstructBstructUses_mapString, astructbstructuse.Name)
}

// commit astructbstructuse to the back repo (if it is already staged)
func (astructbstructuse *AstructBstructUse) Commit(stage *StageStruct) *AstructBstructUse {
	if _, ok := stage.AstructBstructUses[astructbstructuse]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitAstructBstructUse(astructbstructuse)
		}
	}
	return astructbstructuse
}

func (astructbstructuse *AstructBstructUse) CommitVoid(stage *StageStruct) {
	astructbstructuse.Commit(stage)
}

// Checkout astructbstructuse to the back repo (if it is already staged)
func (astructbstructuse *AstructBstructUse) Checkout(stage *StageStruct) *AstructBstructUse {
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
func (bstruct *Bstruct) Stage(stage *StageStruct) *Bstruct {
	stage.Bstructs[bstruct] = __member
	stage.Bstructs_mapString[bstruct.Name] = bstruct

	return bstruct
}

// Unstage removes bstruct off the model stage
func (bstruct *Bstruct) Unstage(stage *StageStruct) *Bstruct {
	delete(stage.Bstructs, bstruct)
	delete(stage.Bstructs_mapString, bstruct.Name)
	return bstruct
}

// UnstageVoid removes bstruct off the model stage
func (bstruct *Bstruct) UnstageVoid(stage *StageStruct) {
	delete(stage.Bstructs, bstruct)
	delete(stage.Bstructs_mapString, bstruct.Name)
}

// commit bstruct to the back repo (if it is already staged)
func (bstruct *Bstruct) Commit(stage *StageStruct) *Bstruct {
	if _, ok := stage.Bstructs[bstruct]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitBstruct(bstruct)
		}
	}
	return bstruct
}

func (bstruct *Bstruct) CommitVoid(stage *StageStruct) {
	bstruct.Commit(stage)
}

// Checkout bstruct to the back repo (if it is already staged)
func (bstruct *Bstruct) Checkout(stage *StageStruct) *Bstruct {
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
func (dstruct *Dstruct) Stage(stage *StageStruct) *Dstruct {
	stage.Dstructs[dstruct] = __member
	stage.Dstructs_mapString[dstruct.Name] = dstruct

	return dstruct
}

// Unstage removes dstruct off the model stage
func (dstruct *Dstruct) Unstage(stage *StageStruct) *Dstruct {
	delete(stage.Dstructs, dstruct)
	delete(stage.Dstructs_mapString, dstruct.Name)
	return dstruct
}

// UnstageVoid removes dstruct off the model stage
func (dstruct *Dstruct) UnstageVoid(stage *StageStruct) {
	delete(stage.Dstructs, dstruct)
	delete(stage.Dstructs_mapString, dstruct.Name)
}

// commit dstruct to the back repo (if it is already staged)
func (dstruct *Dstruct) Commit(stage *StageStruct) *Dstruct {
	if _, ok := stage.Dstructs[dstruct]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDstruct(dstruct)
		}
	}
	return dstruct
}

func (dstruct *Dstruct) CommitVoid(stage *StageStruct) {
	dstruct.Commit(stage)
}

// Checkout dstruct to the back repo (if it is already staged)
func (dstruct *Dstruct) Checkout(stage *StageStruct) *Dstruct {
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

// Stage puts fstruct to the model stage
func (fstruct *Fstruct) Stage(stage *StageStruct) *Fstruct {
	stage.Fstructs[fstruct] = __member
	stage.Fstructs_mapString[fstruct.Name] = fstruct

	return fstruct
}

// Unstage removes fstruct off the model stage
func (fstruct *Fstruct) Unstage(stage *StageStruct) *Fstruct {
	delete(stage.Fstructs, fstruct)
	delete(stage.Fstructs_mapString, fstruct.Name)
	return fstruct
}

// UnstageVoid removes fstruct off the model stage
func (fstruct *Fstruct) UnstageVoid(stage *StageStruct) {
	delete(stage.Fstructs, fstruct)
	delete(stage.Fstructs_mapString, fstruct.Name)
}

// commit fstruct to the back repo (if it is already staged)
func (fstruct *Fstruct) Commit(stage *StageStruct) *Fstruct {
	if _, ok := stage.Fstructs[fstruct]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitFstruct(fstruct)
		}
	}
	return fstruct
}

func (fstruct *Fstruct) CommitVoid(stage *StageStruct) {
	fstruct.Commit(stage)
}

// Checkout fstruct to the back repo (if it is already staged)
func (fstruct *Fstruct) Checkout(stage *StageStruct) *Fstruct {
	if _, ok := stage.Fstructs[fstruct]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutFstruct(fstruct)
		}
	}
	return fstruct
}

// for satisfaction of GongStruct interface
func (fstruct *Fstruct) GetName() (res string) {
	return fstruct.Name
}

// Stage puts gstruct to the model stage
func (gstruct *Gstruct) Stage(stage *StageStruct) *Gstruct {
	stage.Gstructs[gstruct] = __member
	stage.Gstructs_mapString[gstruct.Name] = gstruct

	return gstruct
}

// Unstage removes gstruct off the model stage
func (gstruct *Gstruct) Unstage(stage *StageStruct) *Gstruct {
	delete(stage.Gstructs, gstruct)
	delete(stage.Gstructs_mapString, gstruct.Name)
	return gstruct
}

// UnstageVoid removes gstruct off the model stage
func (gstruct *Gstruct) UnstageVoid(stage *StageStruct) {
	delete(stage.Gstructs, gstruct)
	delete(stage.Gstructs_mapString, gstruct.Name)
}

// commit gstruct to the back repo (if it is already staged)
func (gstruct *Gstruct) Commit(stage *StageStruct) *Gstruct {
	if _, ok := stage.Gstructs[gstruct]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitGstruct(gstruct)
		}
	}
	return gstruct
}

func (gstruct *Gstruct) CommitVoid(stage *StageStruct) {
	gstruct.Commit(stage)
}

// Checkout gstruct to the back repo (if it is already staged)
func (gstruct *Gstruct) Checkout(stage *StageStruct) *Gstruct {
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
	CreateORMFstruct(Fstruct *Fstruct)
	CreateORMGstruct(Gstruct *Gstruct)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMAstruct(Astruct *Astruct)
	DeleteORMAstructBstruct2Use(AstructBstruct2Use *AstructBstruct2Use)
	DeleteORMAstructBstructUse(AstructBstructUse *AstructBstructUse)
	DeleteORMBstruct(Bstruct *Bstruct)
	DeleteORMDstruct(Dstruct *Dstruct)
	DeleteORMFstruct(Fstruct *Fstruct)
	DeleteORMGstruct(Gstruct *Gstruct)
}

func (stage *StageStruct) Reset() { // insertion point for array reset
	stage.Astructs = make(map[*Astruct]any)
	stage.Astructs_mapString = make(map[string]*Astruct)

	stage.AstructBstruct2Uses = make(map[*AstructBstruct2Use]any)
	stage.AstructBstruct2Uses_mapString = make(map[string]*AstructBstruct2Use)

	stage.AstructBstructUses = make(map[*AstructBstructUse]any)
	stage.AstructBstructUses_mapString = make(map[string]*AstructBstructUse)

	stage.Bstructs = make(map[*Bstruct]any)
	stage.Bstructs_mapString = make(map[string]*Bstruct)

	stage.Dstructs = make(map[*Dstruct]any)
	stage.Dstructs_mapString = make(map[string]*Dstruct)

	stage.Fstructs = make(map[*Fstruct]any)
	stage.Fstructs_mapString = make(map[string]*Fstruct)

	stage.Gstructs = make(map[*Gstruct]any)
	stage.Gstructs_mapString = make(map[string]*Gstruct)

}

func (stage *StageStruct) Nil() { // insertion point for array nil
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

	stage.Fstructs = nil
	stage.Fstructs_mapString = nil

	stage.Gstructs = nil
	stage.Gstructs_mapString = nil

}

func (stage *StageStruct) Unstage() { // insertion point for array nil
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

	for fstruct := range stage.Fstructs {
		fstruct.Unstage(stage)
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
	case map[*Fstruct]any:
		return any(&stage.Fstructs).(*Type)
	case map[*Gstruct]any:
		return any(&stage.Gstructs).(*Type)
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
	case map[string]*Fstruct:
		return any(&stage.Fstructs_mapString).(*Type)
	case map[string]*Gstruct:
		return any(&stage.Gstructs_mapString).(*Type)
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
	case Fstruct:
		return any(&stage.Fstructs).(*map[*Type]any)
	case Gstruct:
		return any(&stage.Gstructs).(*map[*Type]any)
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
	case *Fstruct:
		return any(&stage.Fstructs).(*map[Type]any)
	case *Gstruct:
		return any(&stage.Gstructs).(*map[Type]any)
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
	case Fstruct:
		return any(&stage.Fstructs_mapString).(*map[string]*Type)
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
	case Fstruct:
		return any(&Fstruct{
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
func GetPointerReverseMap[Start, End Gongstruct](fieldname string, stage *StageStruct) map[*End][]*Start {

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
	// reverse maps of direct associations of Fstruct
	case Fstruct:
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
func GetSliceOfPointersReverseMap[Start, End Gongstruct](fieldname string, stage *StageStruct) map[*End]*Start {

	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of Astruct
	case Astruct:
		switch fieldname {
		// insertion point for per direct association field
		case "Anarrayofb":
			res := make(map[*Bstruct]*Astruct)
			for astruct := range stage.Astructs {
				for _, bstruct_ := range astruct.Anarrayofb {
					res[bstruct_] = astruct
				}
			}
			return any(res).(map[*End]*Start)
		case "Dstruct4s":
			res := make(map[*Dstruct]*Astruct)
			for astruct := range stage.Astructs {
				for _, dstruct_ := range astruct.Dstruct4s {
					res[dstruct_] = astruct
				}
			}
			return any(res).(map[*End]*Start)
		case "Anarrayofa":
			res := make(map[*Astruct]*Astruct)
			for astruct := range stage.Astructs {
				for _, astruct_ := range astruct.Anarrayofa {
					res[astruct_] = astruct
				}
			}
			return any(res).(map[*End]*Start)
		case "Anotherarrayofb":
			res := make(map[*Bstruct]*Astruct)
			for astruct := range stage.Astructs {
				for _, bstruct_ := range astruct.Anotherarrayofb {
					res[bstruct_] = astruct
				}
			}
			return any(res).(map[*End]*Start)
		case "AnarrayofbUse":
			res := make(map[*AstructBstructUse]*Astruct)
			for astruct := range stage.Astructs {
				for _, astructbstructuse_ := range astruct.AnarrayofbUse {
					res[astructbstructuse_] = astruct
				}
			}
			return any(res).(map[*End]*Start)
		case "Anarrayofb2Use":
			res := make(map[*AstructBstruct2Use]*Astruct)
			for astruct := range stage.Astructs {
				for _, astructbstruct2use_ := range astruct.Anarrayofb2Use {
					res[astructbstruct2use_] = astruct
				}
			}
			return any(res).(map[*End]*Start)
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
			res := make(map[*Bstruct]*Dstruct)
			for dstruct := range stage.Dstructs {
				for _, bstruct_ := range dstruct.Anarrayofb {
					res[bstruct_] = dstruct
				}
			}
			return any(res).(map[*End]*Start)
		case "Gstructs":
			res := make(map[*Gstruct]*Dstruct)
			for dstruct := range stage.Dstructs {
				for _, gstruct_ := range dstruct.Gstructs {
					res[gstruct_] = dstruct
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of Fstruct
	case Fstruct:
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

// GetGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetGongstructName[Type Gongstruct]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case Astruct:
		res = "Astruct"
	case AstructBstruct2Use:
		res = "AstructBstruct2Use"
	case AstructBstructUse:
		res = "AstructBstructUse"
	case Bstruct:
		res = "Bstruct"
	case Dstruct:
		res = "Dstruct"
	case Fstruct:
		res = "Fstruct"
	case Gstruct:
		res = "Gstruct"
	}
	return res
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
	case *Fstruct:
		res = "Fstruct"
	case *Gstruct:
		res = "Gstruct"
	}
	return res
}

// GetFields return the array of the fields
func GetFields[Type Gongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case Astruct:
		res = []string{"Name", "Associationtob", "Anarrayofb", "Anotherassociationtob_2", "Date", "Date2", "Booleanfield", "Aenum", "Aenum_2", "Benum", "CEnum", "CName", "CFloatfield", "Bstruct", "Bstruct2", "Dstruct", "Dstruct2", "Dstruct3", "Dstruct4", "Dstruct4s", "Floatfield", "Intfield", "Anotherbooleanfield", "Duration1", "Anarrayofa", "Anotherarrayofb", "AnarrayofbUse", "Anarrayofb2Use", "AnAstruct", "StructRef", "FieldRef", "EnumIntRef", "EnumStringRef", "EnumValue", "ConstIdentifierValue", "TextFieldBespokeSize", "TextArea"}
	case AstructBstruct2Use:
		res = []string{"Name", "Bstrcut2"}
	case AstructBstructUse:
		res = []string{"Name", "Bstruct2"}
	case Bstruct:
		res = []string{"Name", "Floatfield", "Floatfield2", "Intfield"}
	case Dstruct:
		res = []string{"Name", "Anarrayofb", "Gstruct", "Gstructs"}
	case Fstruct:
		res = []string{"Name", "Date"}
	case Gstruct:
		res = []string{"Name", "Floatfield", "Floatfield2", "Intfield"}
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
	case Astruct:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Astruct"
		rf.Fieldname = "Anarrayofa"
		res = append(res, rf)
	case AstructBstruct2Use:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Astruct"
		rf.Fieldname = "Anarrayofb2Use"
		res = append(res, rf)
	case AstructBstructUse:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Astruct"
		rf.Fieldname = "AnarrayofbUse"
		res = append(res, rf)
	case Bstruct:
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
	case Dstruct:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Astruct"
		rf.Fieldname = "Dstruct4s"
		res = append(res, rf)
	case Fstruct:
		var rf ReverseField
		_ = rf
	case Gstruct:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Dstruct"
		rf.Fieldname = "Gstructs"
		res = append(res, rf)
	}
	return
}

// GetFieldsFromPointer return the array of the fields
func GetFieldsFromPointer[Type PointerToGongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *Astruct:
		res = []string{"Name", "Associationtob", "Anarrayofb", "Anotherassociationtob_2", "Date", "Date2", "Booleanfield", "Aenum", "Aenum_2", "Benum", "CEnum", "CName", "CFloatfield", "Bstruct", "Bstruct2", "Dstruct", "Dstruct2", "Dstruct3", "Dstruct4", "Dstruct4s", "Floatfield", "Intfield", "Anotherbooleanfield", "Duration1", "Anarrayofa", "Anotherarrayofb", "AnarrayofbUse", "Anarrayofb2Use", "AnAstruct", "StructRef", "FieldRef", "EnumIntRef", "EnumStringRef", "EnumValue", "ConstIdentifierValue", "TextFieldBespokeSize", "TextArea"}
	case *AstructBstruct2Use:
		res = []string{"Name", "Bstrcut2"}
	case *AstructBstructUse:
		res = []string{"Name", "Bstruct2"}
	case *Bstruct:
		res = []string{"Name", "Floatfield", "Floatfield2", "Intfield"}
	case *Dstruct:
		res = []string{"Name", "Anarrayofb", "Gstruct", "Gstructs"}
	case *Fstruct:
		res = []string{"Name", "Date"}
	case *Gstruct:
		res = []string{"Name", "Floatfield", "Floatfield2", "Intfield"}
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
	case *Astruct:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Associationtob":
			if inferedInstance.Associationtob != nil {
				res.valueString = inferedInstance.Associationtob.Name
			}
		case "Anarrayofb":
			for idx, __instance__ := range inferedInstance.Anarrayofb {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "Anotherassociationtob_2":
			if inferedInstance.Anotherassociationtob_2 != nil {
				res.valueString = inferedInstance.Anotherassociationtob_2.Name
			}
		case "Date":
			res.valueString = inferedInstance.Date.String()
		case "Date2":
			res.valueString = inferedInstance.Date2.Format("2006-01-02")
		case "Booleanfield":
			res.valueString = fmt.Sprintf("%t", inferedInstance.Booleanfield)
			res.valueBool = inferedInstance.Booleanfield
			res.GongFieldValueType = GongFieldValueTypeBool
		case "Aenum":
			enum := inferedInstance.Aenum
			res.valueString = enum.ToCodeString()
		case "Aenum_2":
			enum := inferedInstance.Aenum_2
			res.valueString = enum.ToCodeString()
		case "Benum":
			enum := inferedInstance.Benum
			res.valueString = enum.ToCodeString()
		case "CEnum":
			enum := inferedInstance.CEnum
			res.valueString = enum.ToCodeString()
		case "CName":
			res.valueString = inferedInstance.CName
		case "CFloatfield":
			res.valueString = fmt.Sprintf("%f", inferedInstance.CFloatfield)
			res.valueFloat = inferedInstance.CFloatfield
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Bstruct":
			if inferedInstance.Bstruct != nil {
				res.valueString = inferedInstance.Bstruct.Name
			}
		case "Bstruct2":
			if inferedInstance.Bstruct2 != nil {
				res.valueString = inferedInstance.Bstruct2.Name
			}
		case "Dstruct":
			if inferedInstance.Dstruct != nil {
				res.valueString = inferedInstance.Dstruct.Name
			}
		case "Dstruct2":
			if inferedInstance.Dstruct2 != nil {
				res.valueString = inferedInstance.Dstruct2.Name
			}
		case "Dstruct3":
			if inferedInstance.Dstruct3 != nil {
				res.valueString = inferedInstance.Dstruct3.Name
			}
		case "Dstruct4":
			if inferedInstance.Dstruct4 != nil {
				res.valueString = inferedInstance.Dstruct4.Name
			}
		case "Dstruct4s":
			for idx, __instance__ := range inferedInstance.Dstruct4s {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "Floatfield":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Floatfield)
			res.valueFloat = inferedInstance.Floatfield
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Intfield":
			res.valueString = fmt.Sprintf("%d", inferedInstance.Intfield)
			res.valueInt = inferedInstance.Intfield
			res.GongFieldValueType = GongFieldValueTypeInt
		case "Anotherbooleanfield":
			res.valueString = fmt.Sprintf("%t", inferedInstance.Anotherbooleanfield)
			res.valueBool = inferedInstance.Anotherbooleanfield
			res.GongFieldValueType = GongFieldValueTypeBool
		case "Duration1":
			if math.Abs(inferedInstance.Duration1.Hours()) >= 24 {
				days := __Gong__Abs(int(int(inferedInstance.Duration1.Hours()) / 24))
				months := int(days / 31)
				days = days - months*31

				remainingHours := int(inferedInstance.Duration1.Hours()) % 24
				remainingMinutes := int(inferedInstance.Duration1.Minutes()) % 60
				remainingSeconds := int(inferedInstance.Duration1.Seconds()) % 60

				if inferedInstance.Duration1.Hours() < 0 {
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
				res.valueString = fmt.Sprintf("%s\n", inferedInstance.Duration1.String())
			}
		case "Anarrayofa":
			for idx, __instance__ := range inferedInstance.Anarrayofa {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "Anotherarrayofb":
			for idx, __instance__ := range inferedInstance.Anotherarrayofb {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "AnarrayofbUse":
			for idx, __instance__ := range inferedInstance.AnarrayofbUse {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "Anarrayofb2Use":
			for idx, __instance__ := range inferedInstance.Anarrayofb2Use {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "AnAstruct":
			if inferedInstance.AnAstruct != nil {
				res.valueString = inferedInstance.AnAstruct.Name
			}
		case "StructRef":
			res.valueString = inferedInstance.StructRef
		case "FieldRef":
			res.valueString = inferedInstance.FieldRef
		case "EnumIntRef":
			res.valueString = inferedInstance.EnumIntRef
		case "EnumStringRef":
			res.valueString = inferedInstance.EnumStringRef
		case "EnumValue":
			res.valueString = inferedInstance.EnumValue
		case "ConstIdentifierValue":
			res.valueString = inferedInstance.ConstIdentifierValue
		case "TextFieldBespokeSize":
			res.valueString = inferedInstance.TextFieldBespokeSize
		case "TextArea":
			res.valueString = inferedInstance.TextArea
		}
	case *AstructBstruct2Use:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Bstrcut2":
			if inferedInstance.Bstrcut2 != nil {
				res.valueString = inferedInstance.Bstrcut2.Name
			}
		}
	case *AstructBstructUse:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Bstruct2":
			if inferedInstance.Bstruct2 != nil {
				res.valueString = inferedInstance.Bstruct2.Name
			}
		}
	case *Bstruct:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Floatfield":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Floatfield)
			res.valueFloat = inferedInstance.Floatfield
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Floatfield2":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Floatfield2)
			res.valueFloat = inferedInstance.Floatfield2
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Intfield":
			res.valueString = fmt.Sprintf("%d", inferedInstance.Intfield)
			res.valueInt = inferedInstance.Intfield
			res.GongFieldValueType = GongFieldValueTypeInt
		}
	case *Dstruct:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Anarrayofb":
			for idx, __instance__ := range inferedInstance.Anarrayofb {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "Gstruct":
			if inferedInstance.Gstruct != nil {
				res.valueString = inferedInstance.Gstruct.Name
			}
		case "Gstructs":
			for idx, __instance__ := range inferedInstance.Gstructs {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case *Fstruct:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Date":
			res.valueString = inferedInstance.Date.String()
		}
	case *Gstruct:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Floatfield":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Floatfield)
			res.valueFloat = inferedInstance.Floatfield
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Floatfield2":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Floatfield2)
			res.valueFloat = inferedInstance.Floatfield2
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Intfield":
			res.valueString = fmt.Sprintf("%d", inferedInstance.Intfield)
			res.valueInt = inferedInstance.Intfield
			res.GongFieldValueType = GongFieldValueTypeInt
		}
	default:
		_ = inferedInstance
	}
	return
}

func GetFieldStringValue(instance any, fieldName string) (res GongFieldValue) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case Astruct:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Associationtob":
			if inferedInstance.Associationtob != nil {
				res.valueString = inferedInstance.Associationtob.Name
			}
		case "Anarrayofb":
			for idx, __instance__ := range inferedInstance.Anarrayofb {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "Anotherassociationtob_2":
			if inferedInstance.Anotherassociationtob_2 != nil {
				res.valueString = inferedInstance.Anotherassociationtob_2.Name
			}
		case "Date":
			res.valueString = inferedInstance.Date.String()
		case "Date2":
			res.valueString = inferedInstance.Date2.Format("2006-01-02")
		case "Booleanfield":
			res.valueString = fmt.Sprintf("%t", inferedInstance.Booleanfield)
			res.valueBool = inferedInstance.Booleanfield
			res.GongFieldValueType = GongFieldValueTypeBool
		case "Aenum":
			enum := inferedInstance.Aenum
			res.valueString = enum.ToCodeString()
		case "Aenum_2":
			enum := inferedInstance.Aenum_2
			res.valueString = enum.ToCodeString()
		case "Benum":
			enum := inferedInstance.Benum
			res.valueString = enum.ToCodeString()
		case "CEnum":
			enum := inferedInstance.CEnum
			res.valueString = enum.ToCodeString()
		case "CName":
			res.valueString = inferedInstance.CName
		case "CFloatfield":
			res.valueString = fmt.Sprintf("%f", inferedInstance.CFloatfield)
			res.valueFloat = inferedInstance.CFloatfield
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Bstruct":
			if inferedInstance.Bstruct != nil {
				res.valueString = inferedInstance.Bstruct.Name
			}
		case "Bstruct2":
			if inferedInstance.Bstruct2 != nil {
				res.valueString = inferedInstance.Bstruct2.Name
			}
		case "Dstruct":
			if inferedInstance.Dstruct != nil {
				res.valueString = inferedInstance.Dstruct.Name
			}
		case "Dstruct2":
			if inferedInstance.Dstruct2 != nil {
				res.valueString = inferedInstance.Dstruct2.Name
			}
		case "Dstruct3":
			if inferedInstance.Dstruct3 != nil {
				res.valueString = inferedInstance.Dstruct3.Name
			}
		case "Dstruct4":
			if inferedInstance.Dstruct4 != nil {
				res.valueString = inferedInstance.Dstruct4.Name
			}
		case "Dstruct4s":
			for idx, __instance__ := range inferedInstance.Dstruct4s {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "Floatfield":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Floatfield)
			res.valueFloat = inferedInstance.Floatfield
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Intfield":
			res.valueString = fmt.Sprintf("%d", inferedInstance.Intfield)
			res.valueInt = inferedInstance.Intfield
			res.GongFieldValueType = GongFieldValueTypeInt
		case "Anotherbooleanfield":
			res.valueString = fmt.Sprintf("%t", inferedInstance.Anotherbooleanfield)
			res.valueBool = inferedInstance.Anotherbooleanfield
			res.GongFieldValueType = GongFieldValueTypeBool
		case "Duration1":
			if math.Abs(inferedInstance.Duration1.Hours()) >= 24 {
				days := __Gong__Abs(int(int(inferedInstance.Duration1.Hours()) / 24))
				months := int(days / 31)
				days = days - months*31

				remainingHours := int(inferedInstance.Duration1.Hours()) % 24
				remainingMinutes := int(inferedInstance.Duration1.Minutes()) % 60
				remainingSeconds := int(inferedInstance.Duration1.Seconds()) % 60

				if inferedInstance.Duration1.Hours() < 0 {
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
				res.valueString = fmt.Sprintf("%s\n", inferedInstance.Duration1.String())
			}
		case "Anarrayofa":
			for idx, __instance__ := range inferedInstance.Anarrayofa {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "Anotherarrayofb":
			for idx, __instance__ := range inferedInstance.Anotherarrayofb {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "AnarrayofbUse":
			for idx, __instance__ := range inferedInstance.AnarrayofbUse {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "Anarrayofb2Use":
			for idx, __instance__ := range inferedInstance.Anarrayofb2Use {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "AnAstruct":
			if inferedInstance.AnAstruct != nil {
				res.valueString = inferedInstance.AnAstruct.Name
			}
		case "StructRef":
			res.valueString = inferedInstance.StructRef
		case "FieldRef":
			res.valueString = inferedInstance.FieldRef
		case "EnumIntRef":
			res.valueString = inferedInstance.EnumIntRef
		case "EnumStringRef":
			res.valueString = inferedInstance.EnumStringRef
		case "EnumValue":
			res.valueString = inferedInstance.EnumValue
		case "ConstIdentifierValue":
			res.valueString = inferedInstance.ConstIdentifierValue
		case "TextFieldBespokeSize":
			res.valueString = inferedInstance.TextFieldBespokeSize
		case "TextArea":
			res.valueString = inferedInstance.TextArea
		}
	case AstructBstruct2Use:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Bstrcut2":
			if inferedInstance.Bstrcut2 != nil {
				res.valueString = inferedInstance.Bstrcut2.Name
			}
		}
	case AstructBstructUse:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Bstruct2":
			if inferedInstance.Bstruct2 != nil {
				res.valueString = inferedInstance.Bstruct2.Name
			}
		}
	case Bstruct:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Floatfield":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Floatfield)
			res.valueFloat = inferedInstance.Floatfield
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Floatfield2":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Floatfield2)
			res.valueFloat = inferedInstance.Floatfield2
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Intfield":
			res.valueString = fmt.Sprintf("%d", inferedInstance.Intfield)
			res.valueInt = inferedInstance.Intfield
			res.GongFieldValueType = GongFieldValueTypeInt
		}
	case Dstruct:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Anarrayofb":
			for idx, __instance__ := range inferedInstance.Anarrayofb {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "Gstruct":
			if inferedInstance.Gstruct != nil {
				res.valueString = inferedInstance.Gstruct.Name
			}
		case "Gstructs":
			for idx, __instance__ := range inferedInstance.Gstructs {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case Fstruct:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Date":
			res.valueString = inferedInstance.Date.String()
		}
	case Gstruct:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Floatfield":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Floatfield)
			res.valueFloat = inferedInstance.Floatfield
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Floatfield2":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Floatfield2)
			res.valueFloat = inferedInstance.Floatfield2
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Intfield":
			res.valueString = fmt.Sprintf("%d", inferedInstance.Intfield)
			res.valueInt = inferedInstance.Intfield
			res.GongFieldValueType = GongFieldValueTypeInt
		}
	default:
		_ = inferedInstance
	}
	return
}

// Last line of the template
