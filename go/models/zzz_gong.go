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

	gong_go "github.com/fullstack-lang/gong/go"
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
	GongBasicFields                map[*GongBasicField]struct{}
	GongBasicFields_reference      map[*GongBasicField]*GongBasicField
	GongBasicFields_referenceOrder map[*GongBasicField]uint // diff Unstage needs the reference order
	GongBasicFields_mapString      map[string]*GongBasicField

	// insertion point for slice of pointers maps
	OnAfterGongBasicFieldCreateCallback OnAfterCreateInterface[GongBasicField]
	OnAfterGongBasicFieldUpdateCallback OnAfterUpdateInterface[GongBasicField]
	OnAfterGongBasicFieldDeleteCallback OnAfterDeleteInterface[GongBasicField]
	OnAfterGongBasicFieldReadCallback   OnAfterReadInterface[GongBasicField]

	GongEnums                map[*GongEnum]struct{}
	GongEnums_reference      map[*GongEnum]*GongEnum
	GongEnums_referenceOrder map[*GongEnum]uint // diff Unstage needs the reference order
	GongEnums_mapString      map[string]*GongEnum

	// insertion point for slice of pointers maps
	GongEnum_GongEnumValues_reverseMap map[*GongEnumValue]*GongEnum

	OnAfterGongEnumCreateCallback OnAfterCreateInterface[GongEnum]
	OnAfterGongEnumUpdateCallback OnAfterUpdateInterface[GongEnum]
	OnAfterGongEnumDeleteCallback OnAfterDeleteInterface[GongEnum]
	OnAfterGongEnumReadCallback   OnAfterReadInterface[GongEnum]

	GongEnumValues                map[*GongEnumValue]struct{}
	GongEnumValues_reference      map[*GongEnumValue]*GongEnumValue
	GongEnumValues_referenceOrder map[*GongEnumValue]uint // diff Unstage needs the reference order
	GongEnumValues_mapString      map[string]*GongEnumValue

	// insertion point for slice of pointers maps
	OnAfterGongEnumValueCreateCallback OnAfterCreateInterface[GongEnumValue]
	OnAfterGongEnumValueUpdateCallback OnAfterUpdateInterface[GongEnumValue]
	OnAfterGongEnumValueDeleteCallback OnAfterDeleteInterface[GongEnumValue]
	OnAfterGongEnumValueReadCallback   OnAfterReadInterface[GongEnumValue]

	GongLinks                map[*GongLink]struct{}
	GongLinks_reference      map[*GongLink]*GongLink
	GongLinks_referenceOrder map[*GongLink]uint // diff Unstage needs the reference order
	GongLinks_mapString      map[string]*GongLink

	// insertion point for slice of pointers maps
	OnAfterGongLinkCreateCallback OnAfterCreateInterface[GongLink]
	OnAfterGongLinkUpdateCallback OnAfterUpdateInterface[GongLink]
	OnAfterGongLinkDeleteCallback OnAfterDeleteInterface[GongLink]
	OnAfterGongLinkReadCallback   OnAfterReadInterface[GongLink]

	GongNotes                map[*GongNote]struct{}
	GongNotes_reference      map[*GongNote]*GongNote
	GongNotes_referenceOrder map[*GongNote]uint // diff Unstage needs the reference order
	GongNotes_mapString      map[string]*GongNote

	// insertion point for slice of pointers maps
	GongNote_Links_reverseMap map[*GongLink]*GongNote

	OnAfterGongNoteCreateCallback OnAfterCreateInterface[GongNote]
	OnAfterGongNoteUpdateCallback OnAfterUpdateInterface[GongNote]
	OnAfterGongNoteDeleteCallback OnAfterDeleteInterface[GongNote]
	OnAfterGongNoteReadCallback   OnAfterReadInterface[GongNote]

	GongStructs                map[*GongStruct]struct{}
	GongStructs_reference      map[*GongStruct]*GongStruct
	GongStructs_referenceOrder map[*GongStruct]uint // diff Unstage needs the reference order
	GongStructs_mapString      map[string]*GongStruct

	// insertion point for slice of pointers maps
	GongStruct_GongBasicFields_reverseMap map[*GongBasicField]*GongStruct

	GongStruct_GongTimeFields_reverseMap map[*GongTimeField]*GongStruct

	GongStruct_PointerToGongStructFields_reverseMap map[*PointerToGongStructField]*GongStruct

	GongStruct_SliceOfPointerToGongStructFields_reverseMap map[*SliceOfPointerToGongStructField]*GongStruct

	OnAfterGongStructCreateCallback OnAfterCreateInterface[GongStruct]
	OnAfterGongStructUpdateCallback OnAfterUpdateInterface[GongStruct]
	OnAfterGongStructDeleteCallback OnAfterDeleteInterface[GongStruct]
	OnAfterGongStructReadCallback   OnAfterReadInterface[GongStruct]

	GongTimeFields                map[*GongTimeField]struct{}
	GongTimeFields_reference      map[*GongTimeField]*GongTimeField
	GongTimeFields_referenceOrder map[*GongTimeField]uint // diff Unstage needs the reference order
	GongTimeFields_mapString      map[string]*GongTimeField

	// insertion point for slice of pointers maps
	OnAfterGongTimeFieldCreateCallback OnAfterCreateInterface[GongTimeField]
	OnAfterGongTimeFieldUpdateCallback OnAfterUpdateInterface[GongTimeField]
	OnAfterGongTimeFieldDeleteCallback OnAfterDeleteInterface[GongTimeField]
	OnAfterGongTimeFieldReadCallback   OnAfterReadInterface[GongTimeField]

	MetaReferences                map[*MetaReference]struct{}
	MetaReferences_reference      map[*MetaReference]*MetaReference
	MetaReferences_referenceOrder map[*MetaReference]uint // diff Unstage needs the reference order
	MetaReferences_mapString      map[string]*MetaReference

	// insertion point for slice of pointers maps
	OnAfterMetaReferenceCreateCallback OnAfterCreateInterface[MetaReference]
	OnAfterMetaReferenceUpdateCallback OnAfterUpdateInterface[MetaReference]
	OnAfterMetaReferenceDeleteCallback OnAfterDeleteInterface[MetaReference]
	OnAfterMetaReferenceReadCallback   OnAfterReadInterface[MetaReference]

	ModelPkgs                map[*ModelPkg]struct{}
	ModelPkgs_reference      map[*ModelPkg]*ModelPkg
	ModelPkgs_referenceOrder map[*ModelPkg]uint // diff Unstage needs the reference order
	ModelPkgs_mapString      map[string]*ModelPkg

	// insertion point for slice of pointers maps
	OnAfterModelPkgCreateCallback OnAfterCreateInterface[ModelPkg]
	OnAfterModelPkgUpdateCallback OnAfterUpdateInterface[ModelPkg]
	OnAfterModelPkgDeleteCallback OnAfterDeleteInterface[ModelPkg]
	OnAfterModelPkgReadCallback   OnAfterReadInterface[ModelPkg]

	PointerToGongStructFields                map[*PointerToGongStructField]struct{}
	PointerToGongStructFields_reference      map[*PointerToGongStructField]*PointerToGongStructField
	PointerToGongStructFields_referenceOrder map[*PointerToGongStructField]uint // diff Unstage needs the reference order
	PointerToGongStructFields_mapString      map[string]*PointerToGongStructField

	// insertion point for slice of pointers maps
	OnAfterPointerToGongStructFieldCreateCallback OnAfterCreateInterface[PointerToGongStructField]
	OnAfterPointerToGongStructFieldUpdateCallback OnAfterUpdateInterface[PointerToGongStructField]
	OnAfterPointerToGongStructFieldDeleteCallback OnAfterDeleteInterface[PointerToGongStructField]
	OnAfterPointerToGongStructFieldReadCallback   OnAfterReadInterface[PointerToGongStructField]

	SliceOfPointerToGongStructFields                map[*SliceOfPointerToGongStructField]struct{}
	SliceOfPointerToGongStructFields_reference      map[*SliceOfPointerToGongStructField]*SliceOfPointerToGongStructField
	SliceOfPointerToGongStructFields_referenceOrder map[*SliceOfPointerToGongStructField]uint // diff Unstage needs the reference order
	SliceOfPointerToGongStructFields_mapString      map[string]*SliceOfPointerToGongStructField

	// insertion point for slice of pointers maps
	OnAfterSliceOfPointerToGongStructFieldCreateCallback OnAfterCreateInterface[SliceOfPointerToGongStructField]
	OnAfterSliceOfPointerToGongStructFieldUpdateCallback OnAfterUpdateInterface[SliceOfPointerToGongStructField]
	OnAfterSliceOfPointerToGongStructFieldDeleteCallback OnAfterDeleteInterface[SliceOfPointerToGongStructField]
	OnAfterSliceOfPointerToGongStructFieldReadCallback   OnAfterReadInterface[SliceOfPointerToGongStructField]

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
	GongBasicFieldOrder            uint
	GongBasicFieldMap_Staged_Order map[*GongBasicField]uint

	GongEnumOrder            uint
	GongEnumMap_Staged_Order map[*GongEnum]uint

	GongEnumValueOrder            uint
	GongEnumValueMap_Staged_Order map[*GongEnumValue]uint

	GongLinkOrder            uint
	GongLinkMap_Staged_Order map[*GongLink]uint

	GongNoteOrder            uint
	GongNoteMap_Staged_Order map[*GongNote]uint

	GongStructOrder            uint
	GongStructMap_Staged_Order map[*GongStruct]uint

	GongTimeFieldOrder            uint
	GongTimeFieldMap_Staged_Order map[*GongTimeField]uint

	MetaReferenceOrder            uint
	MetaReferenceMap_Staged_Order map[*MetaReference]uint

	ModelPkgOrder            uint
	ModelPkgMap_Staged_Order map[*ModelPkg]uint

	PointerToGongStructFieldOrder            uint
	PointerToGongStructFieldMap_Staged_Order map[*PointerToGongStructField]uint

	SliceOfPointerToGongStructFieldOrder            uint
	SliceOfPointerToGongStructFieldMap_Staged_Order map[*SliceOfPointerToGongStructField]uint

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

	stage.ComputeReference()
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
	stage.ComputeReference()
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

	// recompute the next order for each struct
	// this is necessary because the order might have been incremented
	// during the commits that have been discarded
	// insertion point for max order recomputation 
	var maxGongBasicFieldOrder uint
	var foundGongBasicField bool
	for _, order := range stage.GongBasicFieldMap_Staged_Order {
		if !foundGongBasicField || order > maxGongBasicFieldOrder {
			maxGongBasicFieldOrder = order
			foundGongBasicField = true
		}
	}
	if foundGongBasicField {
		stage.GongBasicFieldOrder = maxGongBasicFieldOrder + 1
	} else {
		stage.GongBasicFieldOrder = 0
	}

	var maxGongEnumOrder uint
	var foundGongEnum bool
	for _, order := range stage.GongEnumMap_Staged_Order {
		if !foundGongEnum || order > maxGongEnumOrder {
			maxGongEnumOrder = order
			foundGongEnum = true
		}
	}
	if foundGongEnum {
		stage.GongEnumOrder = maxGongEnumOrder + 1
	} else {
		stage.GongEnumOrder = 0
	}

	var maxGongEnumValueOrder uint
	var foundGongEnumValue bool
	for _, order := range stage.GongEnumValueMap_Staged_Order {
		if !foundGongEnumValue || order > maxGongEnumValueOrder {
			maxGongEnumValueOrder = order
			foundGongEnumValue = true
		}
	}
	if foundGongEnumValue {
		stage.GongEnumValueOrder = maxGongEnumValueOrder + 1
	} else {
		stage.GongEnumValueOrder = 0
	}

	var maxGongLinkOrder uint
	var foundGongLink bool
	for _, order := range stage.GongLinkMap_Staged_Order {
		if !foundGongLink || order > maxGongLinkOrder {
			maxGongLinkOrder = order
			foundGongLink = true
		}
	}
	if foundGongLink {
		stage.GongLinkOrder = maxGongLinkOrder + 1
	} else {
		stage.GongLinkOrder = 0
	}

	var maxGongNoteOrder uint
	var foundGongNote bool
	for _, order := range stage.GongNoteMap_Staged_Order {
		if !foundGongNote || order > maxGongNoteOrder {
			maxGongNoteOrder = order
			foundGongNote = true
		}
	}
	if foundGongNote {
		stage.GongNoteOrder = maxGongNoteOrder + 1
	} else {
		stage.GongNoteOrder = 0
	}

	var maxGongStructOrder uint
	var foundGongStruct bool
	for _, order := range stage.GongStructMap_Staged_Order {
		if !foundGongStruct || order > maxGongStructOrder {
			maxGongStructOrder = order
			foundGongStruct = true
		}
	}
	if foundGongStruct {
		stage.GongStructOrder = maxGongStructOrder + 1
	} else {
		stage.GongStructOrder = 0
	}

	var maxGongTimeFieldOrder uint
	var foundGongTimeField bool
	for _, order := range stage.GongTimeFieldMap_Staged_Order {
		if !foundGongTimeField || order > maxGongTimeFieldOrder {
			maxGongTimeFieldOrder = order
			foundGongTimeField = true
		}
	}
	if foundGongTimeField {
		stage.GongTimeFieldOrder = maxGongTimeFieldOrder + 1
	} else {
		stage.GongTimeFieldOrder = 0
	}

	var maxMetaReferenceOrder uint
	var foundMetaReference bool
	for _, order := range stage.MetaReferenceMap_Staged_Order {
		if !foundMetaReference || order > maxMetaReferenceOrder {
			maxMetaReferenceOrder = order
			foundMetaReference = true
		}
	}
	if foundMetaReference {
		stage.MetaReferenceOrder = maxMetaReferenceOrder + 1
	} else {
		stage.MetaReferenceOrder = 0
	}

	var maxModelPkgOrder uint
	var foundModelPkg bool
	for _, order := range stage.ModelPkgMap_Staged_Order {
		if !foundModelPkg || order > maxModelPkgOrder {
			maxModelPkgOrder = order
			foundModelPkg = true
		}
	}
	if foundModelPkg {
		stage.ModelPkgOrder = maxModelPkgOrder + 1
	} else {
		stage.ModelPkgOrder = 0
	}

	var maxPointerToGongStructFieldOrder uint
	var foundPointerToGongStructField bool
	for _, order := range stage.PointerToGongStructFieldMap_Staged_Order {
		if !foundPointerToGongStructField || order > maxPointerToGongStructFieldOrder {
			maxPointerToGongStructFieldOrder = order
			foundPointerToGongStructField = true
		}
	}
	if foundPointerToGongStructField {
		stage.PointerToGongStructFieldOrder = maxPointerToGongStructFieldOrder + 1
	} else {
		stage.PointerToGongStructFieldOrder = 0
	}

	var maxSliceOfPointerToGongStructFieldOrder uint
	var foundSliceOfPointerToGongStructField bool
	for _, order := range stage.SliceOfPointerToGongStructFieldMap_Staged_Order {
		if !foundSliceOfPointerToGongStructField || order > maxSliceOfPointerToGongStructFieldOrder {
			maxSliceOfPointerToGongStructFieldOrder = order
			foundSliceOfPointerToGongStructField = true
		}
	}
	if foundSliceOfPointerToGongStructField {
		stage.SliceOfPointerToGongStructFieldOrder = maxSliceOfPointerToGongStructFieldOrder + 1
	} else {
		stage.SliceOfPointerToGongStructFieldOrder = 0
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
	case *GongBasicField:
		tmp := GetStructInstancesByOrder(stage.GongBasicFields, stage.GongBasicFieldMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *GongBasicField implements.
			res = append(res, any(v).(T))
		}
		return res
	case *GongEnum:
		tmp := GetStructInstancesByOrder(stage.GongEnums, stage.GongEnumMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *GongEnum implements.
			res = append(res, any(v).(T))
		}
		return res
	case *GongEnumValue:
		tmp := GetStructInstancesByOrder(stage.GongEnumValues, stage.GongEnumValueMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *GongEnumValue implements.
			res = append(res, any(v).(T))
		}
		return res
	case *GongLink:
		tmp := GetStructInstancesByOrder(stage.GongLinks, stage.GongLinkMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *GongLink implements.
			res = append(res, any(v).(T))
		}
		return res
	case *GongNote:
		tmp := GetStructInstancesByOrder(stage.GongNotes, stage.GongNoteMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *GongNote implements.
			res = append(res, any(v).(T))
		}
		return res
	case *GongStruct:
		tmp := GetStructInstancesByOrder(stage.GongStructs, stage.GongStructMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *GongStruct implements.
			res = append(res, any(v).(T))
		}
		return res
	case *GongTimeField:
		tmp := GetStructInstancesByOrder(stage.GongTimeFields, stage.GongTimeFieldMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *GongTimeField implements.
			res = append(res, any(v).(T))
		}
		return res
	case *MetaReference:
		tmp := GetStructInstancesByOrder(stage.MetaReferences, stage.MetaReferenceMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *MetaReference implements.
			res = append(res, any(v).(T))
		}
		return res
	case *ModelPkg:
		tmp := GetStructInstancesByOrder(stage.ModelPkgs, stage.ModelPkgMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *ModelPkg implements.
			res = append(res, any(v).(T))
		}
		return res
	case *PointerToGongStructField:
		tmp := GetStructInstancesByOrder(stage.PointerToGongStructFields, stage.PointerToGongStructFieldMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *PointerToGongStructField implements.
			res = append(res, any(v).(T))
		}
		return res
	case *SliceOfPointerToGongStructField:
		tmp := GetStructInstancesByOrder(stage.SliceOfPointerToGongStructFields, stage.SliceOfPointerToGongStructFieldMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *SliceOfPointerToGongStructField implements.
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
	case "GongBasicField":
		res = GetNamedStructInstances(stage.GongBasicFields, stage.GongBasicFieldMap_Staged_Order)
	case "GongEnum":
		res = GetNamedStructInstances(stage.GongEnums, stage.GongEnumMap_Staged_Order)
	case "GongEnumValue":
		res = GetNamedStructInstances(stage.GongEnumValues, stage.GongEnumValueMap_Staged_Order)
	case "GongLink":
		res = GetNamedStructInstances(stage.GongLinks, stage.GongLinkMap_Staged_Order)
	case "GongNote":
		res = GetNamedStructInstances(stage.GongNotes, stage.GongNoteMap_Staged_Order)
	case "GongStruct":
		res = GetNamedStructInstances(stage.GongStructs, stage.GongStructMap_Staged_Order)
	case "GongTimeField":
		res = GetNamedStructInstances(stage.GongTimeFields, stage.GongTimeFieldMap_Staged_Order)
	case "MetaReference":
		res = GetNamedStructInstances(stage.MetaReferences, stage.MetaReferenceMap_Staged_Order)
	case "ModelPkg":
		res = GetNamedStructInstances(stage.ModelPkgs, stage.ModelPkgMap_Staged_Order)
	case "PointerToGongStructField":
		res = GetNamedStructInstances(stage.PointerToGongStructFields, stage.PointerToGongStructFieldMap_Staged_Order)
	case "SliceOfPointerToGongStructField":
		res = GetNamedStructInstances(stage.SliceOfPointerToGongStructFields, stage.SliceOfPointerToGongStructFieldMap_Staged_Order)
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
	return "github.com/fullstack-lang/gong/go/models"
}

func (stage *Stage) GetMap_GongStructName_InstancesNb() map[string]int {
	return stage.Map_GongStructName_InstancesNb
}

func (stage *Stage) GetModelsEmbededDir() embed.FS {
	return gong_go.GoModelsDir
}

func (stage *Stage) GetDigramsEmbededDir() embed.FS {
	return gong_go.GoDiagramsDir
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
	CommitGongBasicField(gongbasicfield *GongBasicField)
	CheckoutGongBasicField(gongbasicfield *GongBasicField)
	CommitGongEnum(gongenum *GongEnum)
	CheckoutGongEnum(gongenum *GongEnum)
	CommitGongEnumValue(gongenumvalue *GongEnumValue)
	CheckoutGongEnumValue(gongenumvalue *GongEnumValue)
	CommitGongLink(gonglink *GongLink)
	CheckoutGongLink(gonglink *GongLink)
	CommitGongNote(gongnote *GongNote)
	CheckoutGongNote(gongnote *GongNote)
	CommitGongStruct(gongstruct *GongStruct)
	CheckoutGongStruct(gongstruct *GongStruct)
	CommitGongTimeField(gongtimefield *GongTimeField)
	CheckoutGongTimeField(gongtimefield *GongTimeField)
	CommitMetaReference(metareference *MetaReference)
	CheckoutMetaReference(metareference *MetaReference)
	CommitModelPkg(modelpkg *ModelPkg)
	CheckoutModelPkg(modelpkg *ModelPkg)
	CommitPointerToGongStructField(pointertogongstructfield *PointerToGongStructField)
	CheckoutPointerToGongStructField(pointertogongstructfield *PointerToGongStructField)
	CommitSliceOfPointerToGongStructField(sliceofpointertogongstructfield *SliceOfPointerToGongStructField)
	CheckoutSliceOfPointerToGongStructField(sliceofpointertogongstructfield *SliceOfPointerToGongStructField)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(name string) (stage *Stage) {

	stage = &Stage{ // insertion point for array initiatialisation
		GongBasicFields:           make(map[*GongBasicField]struct{}),
		GongBasicFields_mapString: make(map[string]*GongBasicField),

		GongEnums:           make(map[*GongEnum]struct{}),
		GongEnums_mapString: make(map[string]*GongEnum),

		GongEnumValues:           make(map[*GongEnumValue]struct{}),
		GongEnumValues_mapString: make(map[string]*GongEnumValue),

		GongLinks:           make(map[*GongLink]struct{}),
		GongLinks_mapString: make(map[string]*GongLink),

		GongNotes:           make(map[*GongNote]struct{}),
		GongNotes_mapString: make(map[string]*GongNote),

		GongStructs:           make(map[*GongStruct]struct{}),
		GongStructs_mapString: make(map[string]*GongStruct),

		GongTimeFields:           make(map[*GongTimeField]struct{}),
		GongTimeFields_mapString: make(map[string]*GongTimeField),

		MetaReferences:           make(map[*MetaReference]struct{}),
		MetaReferences_mapString: make(map[string]*MetaReference),

		ModelPkgs:           make(map[*ModelPkg]struct{}),
		ModelPkgs_mapString: make(map[string]*ModelPkg),

		PointerToGongStructFields:           make(map[*PointerToGongStructField]struct{}),
		PointerToGongStructFields_mapString: make(map[string]*PointerToGongStructField),

		SliceOfPointerToGongStructFields:           make(map[*SliceOfPointerToGongStructField]struct{}),
		SliceOfPointerToGongStructFields_mapString: make(map[string]*SliceOfPointerToGongStructField),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		name: name,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		GongBasicFieldMap_Staged_Order: make(map[*GongBasicField]uint),

		GongEnumMap_Staged_Order: make(map[*GongEnum]uint),

		GongEnumValueMap_Staged_Order: make(map[*GongEnumValue]uint),

		GongLinkMap_Staged_Order: make(map[*GongLink]uint),

		GongNoteMap_Staged_Order: make(map[*GongNote]uint),

		GongStructMap_Staged_Order: make(map[*GongStruct]uint),

		GongTimeFieldMap_Staged_Order: make(map[*GongTimeField]uint),

		MetaReferenceMap_Staged_Order: make(map[*MetaReference]uint),

		ModelPkgMap_Staged_Order: make(map[*ModelPkg]uint),

		PointerToGongStructFieldMap_Staged_Order: make(map[*PointerToGongStructField]uint),

		SliceOfPointerToGongStructFieldMap_Staged_Order: make(map[*SliceOfPointerToGongStructField]uint),

		// end of insertion point
		GongUnmarshallers: map[string]ModelUnmarshaller{ // insertion point for unmarshallers
			"GongBasicField": &GongBasicFieldUnmarshaller{},

			"GongEnum": &GongEnumUnmarshaller{},

			"GongEnumValue": &GongEnumValueUnmarshaller{},

			"GongLink": &GongLinkUnmarshaller{},

			"GongNote": &GongNoteUnmarshaller{},

			"GongStruct": &GongStructUnmarshaller{},

			"GongTimeField": &GongTimeFieldUnmarshaller{},

			"MetaReference": &MetaReferenceUnmarshaller{},

			"ModelPkg": &ModelPkgUnmarshaller{},

			"PointerToGongStructField": &PointerToGongStructFieldUnmarshaller{},

			"SliceOfPointerToGongStructField": &SliceOfPointerToGongStructFieldUnmarshaller{},

			// end of insertion point
		},

		NamedStructs: []*NamedStruct{ // insertion point for order map initialisations
			{name: "GongBasicField"},
			{name: "GongEnum"},
			{name: "GongEnumValue"},
			{name: "GongLink"},
			{name: "GongNote"},
			{name: "GongStruct"},
			{name: "GongTimeField"},
			{name: "MetaReference"},
			{name: "ModelPkg"},
			{name: "PointerToGongStructField"},
			{name: "SliceOfPointerToGongStructField"},
		}, // end of insertion point

		navigationMode: GongNavigationModeNormal,
	}

	return
}

func GetOrder[Type Gongstruct](stage *Stage, instance *Type) uint {

	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *GongBasicField:
		return stage.GongBasicFieldMap_Staged_Order[instance]
	case *GongEnum:
		return stage.GongEnumMap_Staged_Order[instance]
	case *GongEnumValue:
		return stage.GongEnumValueMap_Staged_Order[instance]
	case *GongLink:
		return stage.GongLinkMap_Staged_Order[instance]
	case *GongNote:
		return stage.GongNoteMap_Staged_Order[instance]
	case *GongStruct:
		return stage.GongStructMap_Staged_Order[instance]
	case *GongTimeField:
		return stage.GongTimeFieldMap_Staged_Order[instance]
	case *MetaReference:
		return stage.MetaReferenceMap_Staged_Order[instance]
	case *ModelPkg:
		return stage.ModelPkgMap_Staged_Order[instance]
	case *PointerToGongStructField:
		return stage.PointerToGongStructFieldMap_Staged_Order[instance]
	case *SliceOfPointerToGongStructField:
		return stage.SliceOfPointerToGongStructFieldMap_Staged_Order[instance]
	default:
		return 0 // should not happen
	}
}

func GetOrderPointerGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) uint {

	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *GongBasicField:
		return stage.GongBasicFieldMap_Staged_Order[instance]
	case *GongEnum:
		return stage.GongEnumMap_Staged_Order[instance]
	case *GongEnumValue:
		return stage.GongEnumValueMap_Staged_Order[instance]
	case *GongLink:
		return stage.GongLinkMap_Staged_Order[instance]
	case *GongNote:
		return stage.GongNoteMap_Staged_Order[instance]
	case *GongStruct:
		return stage.GongStructMap_Staged_Order[instance]
	case *GongTimeField:
		return stage.GongTimeFieldMap_Staged_Order[instance]
	case *MetaReference:
		return stage.MetaReferenceMap_Staged_Order[instance]
	case *ModelPkg:
		return stage.ModelPkgMap_Staged_Order[instance]
	case *PointerToGongStructField:
		return stage.PointerToGongStructFieldMap_Staged_Order[instance]
	case *SliceOfPointerToGongStructField:
		return stage.SliceOfPointerToGongStructFieldMap_Staged_Order[instance]
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
		stage.ComputeReference()
		if stage.GetProbeIF() != nil {
			stage.GetProbeIF().Refresh()
		}
	}
}

func (stage *Stage) ComputeInstancesNb() {
	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["GongBasicField"] = len(stage.GongBasicFields)
	stage.Map_GongStructName_InstancesNb["GongEnum"] = len(stage.GongEnums)
	stage.Map_GongStructName_InstancesNb["GongEnumValue"] = len(stage.GongEnumValues)
	stage.Map_GongStructName_InstancesNb["GongLink"] = len(stage.GongLinks)
	stage.Map_GongStructName_InstancesNb["GongNote"] = len(stage.GongNotes)
	stage.Map_GongStructName_InstancesNb["GongStruct"] = len(stage.GongStructs)
	stage.Map_GongStructName_InstancesNb["GongTimeField"] = len(stage.GongTimeFields)
	stage.Map_GongStructName_InstancesNb["MetaReference"] = len(stage.MetaReferences)
	stage.Map_GongStructName_InstancesNb["ModelPkg"] = len(stage.ModelPkgs)
	stage.Map_GongStructName_InstancesNb["PointerToGongStructField"] = len(stage.PointerToGongStructFields)
	stage.Map_GongStructName_InstancesNb["SliceOfPointerToGongStructField"] = len(stage.SliceOfPointerToGongStructFields)
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
// Stage puts gongbasicfield to the model stage
func (gongbasicfield *GongBasicField) Stage(stage *Stage) *GongBasicField {

	if _, ok := stage.GongBasicFields[gongbasicfield]; !ok {
		stage.GongBasicFields[gongbasicfield] = struct{}{}
		stage.GongBasicFieldMap_Staged_Order[gongbasicfield] = stage.GongBasicFieldOrder
		stage.GongBasicFieldOrder++
	}
	stage.GongBasicFields_mapString[gongbasicfield.Name] = gongbasicfield

	return gongbasicfield
}

// StagePreserveOrder puts gongbasicfield to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.GongBasicFieldOrder
// - update stage.GongBasicFieldOrder accordingly
func (gongbasicfield *GongBasicField) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.GongBasicFields[gongbasicfield]; !ok {
		stage.GongBasicFields[gongbasicfield] = struct{}{}

		if order > stage.GongBasicFieldOrder {
			stage.GongBasicFieldOrder = order
		}
		stage.GongBasicFieldMap_Staged_Order[gongbasicfield] = order
		stage.GongBasicFieldOrder++
	}
	stage.GongBasicFields_mapString[gongbasicfield.Name] = gongbasicfield
}

// Unstage removes gongbasicfield off the model stage
func (gongbasicfield *GongBasicField) Unstage(stage *Stage) *GongBasicField {
	delete(stage.GongBasicFields, gongbasicfield)
	delete(stage.GongBasicFieldMap_Staged_Order, gongbasicfield)
	delete(stage.GongBasicFields_mapString, gongbasicfield.Name)

	return gongbasicfield
}

// UnstageVoid removes gongbasicfield off the model stage
func (gongbasicfield *GongBasicField) UnstageVoid(stage *Stage) {
	delete(stage.GongBasicFields, gongbasicfield)
	delete(stage.GongBasicFieldMap_Staged_Order, gongbasicfield)
	delete(stage.GongBasicFields_mapString, gongbasicfield.Name)
}

// commit gongbasicfield to the back repo (if it is already staged)
func (gongbasicfield *GongBasicField) Commit(stage *Stage) *GongBasicField {
	if _, ok := stage.GongBasicFields[gongbasicfield]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitGongBasicField(gongbasicfield)
		}
	}
	return gongbasicfield
}

func (gongbasicfield *GongBasicField) CommitVoid(stage *Stage) {
	gongbasicfield.Commit(stage)
}

func (gongbasicfield *GongBasicField) StageVoid(stage *Stage) {
	gongbasicfield.Stage(stage)
}

// Checkout gongbasicfield to the back repo (if it is already staged)
func (gongbasicfield *GongBasicField) Checkout(stage *Stage) *GongBasicField {
	if _, ok := stage.GongBasicFields[gongbasicfield]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutGongBasicField(gongbasicfield)
		}
	}
	return gongbasicfield
}

// for satisfaction of GongStruct interface
func (gongbasicfield *GongBasicField) GetName() (res string) {
	return gongbasicfield.Name
}

// for satisfaction of GongStruct interface
func (gongbasicfield *GongBasicField) SetName(name string) {
	gongbasicfield.Name = name
}

// Stage puts gongenum to the model stage
func (gongenum *GongEnum) Stage(stage *Stage) *GongEnum {

	if _, ok := stage.GongEnums[gongenum]; !ok {
		stage.GongEnums[gongenum] = struct{}{}
		stage.GongEnumMap_Staged_Order[gongenum] = stage.GongEnumOrder
		stage.GongEnumOrder++
	}
	stage.GongEnums_mapString[gongenum.Name] = gongenum

	return gongenum
}

// StagePreserveOrder puts gongenum to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.GongEnumOrder
// - update stage.GongEnumOrder accordingly
func (gongenum *GongEnum) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.GongEnums[gongenum]; !ok {
		stage.GongEnums[gongenum] = struct{}{}

		if order > stage.GongEnumOrder {
			stage.GongEnumOrder = order
		}
		stage.GongEnumMap_Staged_Order[gongenum] = order
		stage.GongEnumOrder++
	}
	stage.GongEnums_mapString[gongenum.Name] = gongenum
}

// Unstage removes gongenum off the model stage
func (gongenum *GongEnum) Unstage(stage *Stage) *GongEnum {
	delete(stage.GongEnums, gongenum)
	delete(stage.GongEnumMap_Staged_Order, gongenum)
	delete(stage.GongEnums_mapString, gongenum.Name)

	return gongenum
}

// UnstageVoid removes gongenum off the model stage
func (gongenum *GongEnum) UnstageVoid(stage *Stage) {
	delete(stage.GongEnums, gongenum)
	delete(stage.GongEnumMap_Staged_Order, gongenum)
	delete(stage.GongEnums_mapString, gongenum.Name)
}

// commit gongenum to the back repo (if it is already staged)
func (gongenum *GongEnum) Commit(stage *Stage) *GongEnum {
	if _, ok := stage.GongEnums[gongenum]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitGongEnum(gongenum)
		}
	}
	return gongenum
}

func (gongenum *GongEnum) CommitVoid(stage *Stage) {
	gongenum.Commit(stage)
}

func (gongenum *GongEnum) StageVoid(stage *Stage) {
	gongenum.Stage(stage)
}

// Checkout gongenum to the back repo (if it is already staged)
func (gongenum *GongEnum) Checkout(stage *Stage) *GongEnum {
	if _, ok := stage.GongEnums[gongenum]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutGongEnum(gongenum)
		}
	}
	return gongenum
}

// for satisfaction of GongStruct interface
func (gongenum *GongEnum) GetName() (res string) {
	return gongenum.Name
}

// for satisfaction of GongStruct interface
func (gongenum *GongEnum) SetName(name string) {
	gongenum.Name = name
}

// Stage puts gongenumvalue to the model stage
func (gongenumvalue *GongEnumValue) Stage(stage *Stage) *GongEnumValue {

	if _, ok := stage.GongEnumValues[gongenumvalue]; !ok {
		stage.GongEnumValues[gongenumvalue] = struct{}{}
		stage.GongEnumValueMap_Staged_Order[gongenumvalue] = stage.GongEnumValueOrder
		stage.GongEnumValueOrder++
	}
	stage.GongEnumValues_mapString[gongenumvalue.Name] = gongenumvalue

	return gongenumvalue
}

// StagePreserveOrder puts gongenumvalue to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.GongEnumValueOrder
// - update stage.GongEnumValueOrder accordingly
func (gongenumvalue *GongEnumValue) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.GongEnumValues[gongenumvalue]; !ok {
		stage.GongEnumValues[gongenumvalue] = struct{}{}

		if order > stage.GongEnumValueOrder {
			stage.GongEnumValueOrder = order
		}
		stage.GongEnumValueMap_Staged_Order[gongenumvalue] = order
		stage.GongEnumValueOrder++
	}
	stage.GongEnumValues_mapString[gongenumvalue.Name] = gongenumvalue
}

// Unstage removes gongenumvalue off the model stage
func (gongenumvalue *GongEnumValue) Unstage(stage *Stage) *GongEnumValue {
	delete(stage.GongEnumValues, gongenumvalue)
	delete(stage.GongEnumValueMap_Staged_Order, gongenumvalue)
	delete(stage.GongEnumValues_mapString, gongenumvalue.Name)

	return gongenumvalue
}

// UnstageVoid removes gongenumvalue off the model stage
func (gongenumvalue *GongEnumValue) UnstageVoid(stage *Stage) {
	delete(stage.GongEnumValues, gongenumvalue)
	delete(stage.GongEnumValueMap_Staged_Order, gongenumvalue)
	delete(stage.GongEnumValues_mapString, gongenumvalue.Name)
}

// commit gongenumvalue to the back repo (if it is already staged)
func (gongenumvalue *GongEnumValue) Commit(stage *Stage) *GongEnumValue {
	if _, ok := stage.GongEnumValues[gongenumvalue]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitGongEnumValue(gongenumvalue)
		}
	}
	return gongenumvalue
}

func (gongenumvalue *GongEnumValue) CommitVoid(stage *Stage) {
	gongenumvalue.Commit(stage)
}

func (gongenumvalue *GongEnumValue) StageVoid(stage *Stage) {
	gongenumvalue.Stage(stage)
}

// Checkout gongenumvalue to the back repo (if it is already staged)
func (gongenumvalue *GongEnumValue) Checkout(stage *Stage) *GongEnumValue {
	if _, ok := stage.GongEnumValues[gongenumvalue]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutGongEnumValue(gongenumvalue)
		}
	}
	return gongenumvalue
}

// for satisfaction of GongStruct interface
func (gongenumvalue *GongEnumValue) GetName() (res string) {
	return gongenumvalue.Name
}

// for satisfaction of GongStruct interface
func (gongenumvalue *GongEnumValue) SetName(name string) {
	gongenumvalue.Name = name
}

// Stage puts gonglink to the model stage
func (gonglink *GongLink) Stage(stage *Stage) *GongLink {

	if _, ok := stage.GongLinks[gonglink]; !ok {
		stage.GongLinks[gonglink] = struct{}{}
		stage.GongLinkMap_Staged_Order[gonglink] = stage.GongLinkOrder
		stage.GongLinkOrder++
	}
	stage.GongLinks_mapString[gonglink.Name] = gonglink

	return gonglink
}

// StagePreserveOrder puts gonglink to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.GongLinkOrder
// - update stage.GongLinkOrder accordingly
func (gonglink *GongLink) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.GongLinks[gonglink]; !ok {
		stage.GongLinks[gonglink] = struct{}{}

		if order > stage.GongLinkOrder {
			stage.GongLinkOrder = order
		}
		stage.GongLinkMap_Staged_Order[gonglink] = order
		stage.GongLinkOrder++
	}
	stage.GongLinks_mapString[gonglink.Name] = gonglink
}

// Unstage removes gonglink off the model stage
func (gonglink *GongLink) Unstage(stage *Stage) *GongLink {
	delete(stage.GongLinks, gonglink)
	delete(stage.GongLinkMap_Staged_Order, gonglink)
	delete(stage.GongLinks_mapString, gonglink.Name)

	return gonglink
}

// UnstageVoid removes gonglink off the model stage
func (gonglink *GongLink) UnstageVoid(stage *Stage) {
	delete(stage.GongLinks, gonglink)
	delete(stage.GongLinkMap_Staged_Order, gonglink)
	delete(stage.GongLinks_mapString, gonglink.Name)
}

// commit gonglink to the back repo (if it is already staged)
func (gonglink *GongLink) Commit(stage *Stage) *GongLink {
	if _, ok := stage.GongLinks[gonglink]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitGongLink(gonglink)
		}
	}
	return gonglink
}

func (gonglink *GongLink) CommitVoid(stage *Stage) {
	gonglink.Commit(stage)
}

func (gonglink *GongLink) StageVoid(stage *Stage) {
	gonglink.Stage(stage)
}

// Checkout gonglink to the back repo (if it is already staged)
func (gonglink *GongLink) Checkout(stage *Stage) *GongLink {
	if _, ok := stage.GongLinks[gonglink]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutGongLink(gonglink)
		}
	}
	return gonglink
}

// for satisfaction of GongStruct interface
func (gonglink *GongLink) GetName() (res string) {
	return gonglink.Name
}

// for satisfaction of GongStruct interface
func (gonglink *GongLink) SetName(name string) {
	gonglink.Name = name
}

// Stage puts gongnote to the model stage
func (gongnote *GongNote) Stage(stage *Stage) *GongNote {

	if _, ok := stage.GongNotes[gongnote]; !ok {
		stage.GongNotes[gongnote] = struct{}{}
		stage.GongNoteMap_Staged_Order[gongnote] = stage.GongNoteOrder
		stage.GongNoteOrder++
	}
	stage.GongNotes_mapString[gongnote.Name] = gongnote

	return gongnote
}

// StagePreserveOrder puts gongnote to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.GongNoteOrder
// - update stage.GongNoteOrder accordingly
func (gongnote *GongNote) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.GongNotes[gongnote]; !ok {
		stage.GongNotes[gongnote] = struct{}{}

		if order > stage.GongNoteOrder {
			stage.GongNoteOrder = order
		}
		stage.GongNoteMap_Staged_Order[gongnote] = order
		stage.GongNoteOrder++
	}
	stage.GongNotes_mapString[gongnote.Name] = gongnote
}

// Unstage removes gongnote off the model stage
func (gongnote *GongNote) Unstage(stage *Stage) *GongNote {
	delete(stage.GongNotes, gongnote)
	delete(stage.GongNoteMap_Staged_Order, gongnote)
	delete(stage.GongNotes_mapString, gongnote.Name)

	return gongnote
}

// UnstageVoid removes gongnote off the model stage
func (gongnote *GongNote) UnstageVoid(stage *Stage) {
	delete(stage.GongNotes, gongnote)
	delete(stage.GongNoteMap_Staged_Order, gongnote)
	delete(stage.GongNotes_mapString, gongnote.Name)
}

// commit gongnote to the back repo (if it is already staged)
func (gongnote *GongNote) Commit(stage *Stage) *GongNote {
	if _, ok := stage.GongNotes[gongnote]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitGongNote(gongnote)
		}
	}
	return gongnote
}

func (gongnote *GongNote) CommitVoid(stage *Stage) {
	gongnote.Commit(stage)
}

func (gongnote *GongNote) StageVoid(stage *Stage) {
	gongnote.Stage(stage)
}

// Checkout gongnote to the back repo (if it is already staged)
func (gongnote *GongNote) Checkout(stage *Stage) *GongNote {
	if _, ok := stage.GongNotes[gongnote]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutGongNote(gongnote)
		}
	}
	return gongnote
}

// for satisfaction of GongStruct interface
func (gongnote *GongNote) GetName() (res string) {
	return gongnote.Name
}

// for satisfaction of GongStruct interface
func (gongnote *GongNote) SetName(name string) {
	gongnote.Name = name
}

// Stage puts gongstruct to the model stage
func (gongstruct *GongStruct) Stage(stage *Stage) *GongStruct {

	if _, ok := stage.GongStructs[gongstruct]; !ok {
		stage.GongStructs[gongstruct] = struct{}{}
		stage.GongStructMap_Staged_Order[gongstruct] = stage.GongStructOrder
		stage.GongStructOrder++
	}
	stage.GongStructs_mapString[gongstruct.Name] = gongstruct

	return gongstruct
}

// StagePreserveOrder puts gongstruct to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.GongStructOrder
// - update stage.GongStructOrder accordingly
func (gongstruct *GongStruct) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.GongStructs[gongstruct]; !ok {
		stage.GongStructs[gongstruct] = struct{}{}

		if order > stage.GongStructOrder {
			stage.GongStructOrder = order
		}
		stage.GongStructMap_Staged_Order[gongstruct] = order
		stage.GongStructOrder++
	}
	stage.GongStructs_mapString[gongstruct.Name] = gongstruct
}

// Unstage removes gongstruct off the model stage
func (gongstruct *GongStruct) Unstage(stage *Stage) *GongStruct {
	delete(stage.GongStructs, gongstruct)
	delete(stage.GongStructMap_Staged_Order, gongstruct)
	delete(stage.GongStructs_mapString, gongstruct.Name)

	return gongstruct
}

// UnstageVoid removes gongstruct off the model stage
func (gongstruct *GongStruct) UnstageVoid(stage *Stage) {
	delete(stage.GongStructs, gongstruct)
	delete(stage.GongStructMap_Staged_Order, gongstruct)
	delete(stage.GongStructs_mapString, gongstruct.Name)
}

// commit gongstruct to the back repo (if it is already staged)
func (gongstruct *GongStruct) Commit(stage *Stage) *GongStruct {
	if _, ok := stage.GongStructs[gongstruct]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitGongStruct(gongstruct)
		}
	}
	return gongstruct
}

func (gongstruct *GongStruct) CommitVoid(stage *Stage) {
	gongstruct.Commit(stage)
}

func (gongstruct *GongStruct) StageVoid(stage *Stage) {
	gongstruct.Stage(stage)
}

// Checkout gongstruct to the back repo (if it is already staged)
func (gongstruct *GongStruct) Checkout(stage *Stage) *GongStruct {
	if _, ok := stage.GongStructs[gongstruct]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutGongStruct(gongstruct)
		}
	}
	return gongstruct
}

// for satisfaction of GongStruct interface
func (gongstruct *GongStruct) GetName() (res string) {
	return gongstruct.Name
}

// for satisfaction of GongStruct interface
func (gongstruct *GongStruct) SetName(name string) {
	gongstruct.Name = name
}

// Stage puts gongtimefield to the model stage
func (gongtimefield *GongTimeField) Stage(stage *Stage) *GongTimeField {

	if _, ok := stage.GongTimeFields[gongtimefield]; !ok {
		stage.GongTimeFields[gongtimefield] = struct{}{}
		stage.GongTimeFieldMap_Staged_Order[gongtimefield] = stage.GongTimeFieldOrder
		stage.GongTimeFieldOrder++
	}
	stage.GongTimeFields_mapString[gongtimefield.Name] = gongtimefield

	return gongtimefield
}

// StagePreserveOrder puts gongtimefield to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.GongTimeFieldOrder
// - update stage.GongTimeFieldOrder accordingly
func (gongtimefield *GongTimeField) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.GongTimeFields[gongtimefield]; !ok {
		stage.GongTimeFields[gongtimefield] = struct{}{}

		if order > stage.GongTimeFieldOrder {
			stage.GongTimeFieldOrder = order
		}
		stage.GongTimeFieldMap_Staged_Order[gongtimefield] = order
		stage.GongTimeFieldOrder++
	}
	stage.GongTimeFields_mapString[gongtimefield.Name] = gongtimefield
}

// Unstage removes gongtimefield off the model stage
func (gongtimefield *GongTimeField) Unstage(stage *Stage) *GongTimeField {
	delete(stage.GongTimeFields, gongtimefield)
	delete(stage.GongTimeFieldMap_Staged_Order, gongtimefield)
	delete(stage.GongTimeFields_mapString, gongtimefield.Name)

	return gongtimefield
}

// UnstageVoid removes gongtimefield off the model stage
func (gongtimefield *GongTimeField) UnstageVoid(stage *Stage) {
	delete(stage.GongTimeFields, gongtimefield)
	delete(stage.GongTimeFieldMap_Staged_Order, gongtimefield)
	delete(stage.GongTimeFields_mapString, gongtimefield.Name)
}

// commit gongtimefield to the back repo (if it is already staged)
func (gongtimefield *GongTimeField) Commit(stage *Stage) *GongTimeField {
	if _, ok := stage.GongTimeFields[gongtimefield]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitGongTimeField(gongtimefield)
		}
	}
	return gongtimefield
}

func (gongtimefield *GongTimeField) CommitVoid(stage *Stage) {
	gongtimefield.Commit(stage)
}

func (gongtimefield *GongTimeField) StageVoid(stage *Stage) {
	gongtimefield.Stage(stage)
}

// Checkout gongtimefield to the back repo (if it is already staged)
func (gongtimefield *GongTimeField) Checkout(stage *Stage) *GongTimeField {
	if _, ok := stage.GongTimeFields[gongtimefield]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutGongTimeField(gongtimefield)
		}
	}
	return gongtimefield
}

// for satisfaction of GongStruct interface
func (gongtimefield *GongTimeField) GetName() (res string) {
	return gongtimefield.Name
}

// for satisfaction of GongStruct interface
func (gongtimefield *GongTimeField) SetName(name string) {
	gongtimefield.Name = name
}

// Stage puts metareference to the model stage
func (metareference *MetaReference) Stage(stage *Stage) *MetaReference {

	if _, ok := stage.MetaReferences[metareference]; !ok {
		stage.MetaReferences[metareference] = struct{}{}
		stage.MetaReferenceMap_Staged_Order[metareference] = stage.MetaReferenceOrder
		stage.MetaReferenceOrder++
	}
	stage.MetaReferences_mapString[metareference.Name] = metareference

	return metareference
}

// StagePreserveOrder puts metareference to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MetaReferenceOrder
// - update stage.MetaReferenceOrder accordingly
func (metareference *MetaReference) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.MetaReferences[metareference]; !ok {
		stage.MetaReferences[metareference] = struct{}{}

		if order > stage.MetaReferenceOrder {
			stage.MetaReferenceOrder = order
		}
		stage.MetaReferenceMap_Staged_Order[metareference] = order
		stage.MetaReferenceOrder++
	}
	stage.MetaReferences_mapString[metareference.Name] = metareference
}

// Unstage removes metareference off the model stage
func (metareference *MetaReference) Unstage(stage *Stage) *MetaReference {
	delete(stage.MetaReferences, metareference)
	delete(stage.MetaReferenceMap_Staged_Order, metareference)
	delete(stage.MetaReferences_mapString, metareference.Name)

	return metareference
}

// UnstageVoid removes metareference off the model stage
func (metareference *MetaReference) UnstageVoid(stage *Stage) {
	delete(stage.MetaReferences, metareference)
	delete(stage.MetaReferenceMap_Staged_Order, metareference)
	delete(stage.MetaReferences_mapString, metareference.Name)
}

// commit metareference to the back repo (if it is already staged)
func (metareference *MetaReference) Commit(stage *Stage) *MetaReference {
	if _, ok := stage.MetaReferences[metareference]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitMetaReference(metareference)
		}
	}
	return metareference
}

func (metareference *MetaReference) CommitVoid(stage *Stage) {
	metareference.Commit(stage)
}

func (metareference *MetaReference) StageVoid(stage *Stage) {
	metareference.Stage(stage)
}

// Checkout metareference to the back repo (if it is already staged)
func (metareference *MetaReference) Checkout(stage *Stage) *MetaReference {
	if _, ok := stage.MetaReferences[metareference]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutMetaReference(metareference)
		}
	}
	return metareference
}

// for satisfaction of GongStruct interface
func (metareference *MetaReference) GetName() (res string) {
	return metareference.Name
}

// for satisfaction of GongStruct interface
func (metareference *MetaReference) SetName(name string) {
	metareference.Name = name
}

// Stage puts modelpkg to the model stage
func (modelpkg *ModelPkg) Stage(stage *Stage) *ModelPkg {

	if _, ok := stage.ModelPkgs[modelpkg]; !ok {
		stage.ModelPkgs[modelpkg] = struct{}{}
		stage.ModelPkgMap_Staged_Order[modelpkg] = stage.ModelPkgOrder
		stage.ModelPkgOrder++
	}
	stage.ModelPkgs_mapString[modelpkg.Name] = modelpkg

	return modelpkg
}

// StagePreserveOrder puts modelpkg to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ModelPkgOrder
// - update stage.ModelPkgOrder accordingly
func (modelpkg *ModelPkg) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.ModelPkgs[modelpkg]; !ok {
		stage.ModelPkgs[modelpkg] = struct{}{}

		if order > stage.ModelPkgOrder {
			stage.ModelPkgOrder = order
		}
		stage.ModelPkgMap_Staged_Order[modelpkg] = order
		stage.ModelPkgOrder++
	}
	stage.ModelPkgs_mapString[modelpkg.Name] = modelpkg
}

// Unstage removes modelpkg off the model stage
func (modelpkg *ModelPkg) Unstage(stage *Stage) *ModelPkg {
	delete(stage.ModelPkgs, modelpkg)
	delete(stage.ModelPkgMap_Staged_Order, modelpkg)
	delete(stage.ModelPkgs_mapString, modelpkg.Name)

	return modelpkg
}

// UnstageVoid removes modelpkg off the model stage
func (modelpkg *ModelPkg) UnstageVoid(stage *Stage) {
	delete(stage.ModelPkgs, modelpkg)
	delete(stage.ModelPkgMap_Staged_Order, modelpkg)
	delete(stage.ModelPkgs_mapString, modelpkg.Name)
}

// commit modelpkg to the back repo (if it is already staged)
func (modelpkg *ModelPkg) Commit(stage *Stage) *ModelPkg {
	if _, ok := stage.ModelPkgs[modelpkg]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitModelPkg(modelpkg)
		}
	}
	return modelpkg
}

func (modelpkg *ModelPkg) CommitVoid(stage *Stage) {
	modelpkg.Commit(stage)
}

func (modelpkg *ModelPkg) StageVoid(stage *Stage) {
	modelpkg.Stage(stage)
}

// Checkout modelpkg to the back repo (if it is already staged)
func (modelpkg *ModelPkg) Checkout(stage *Stage) *ModelPkg {
	if _, ok := stage.ModelPkgs[modelpkg]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutModelPkg(modelpkg)
		}
	}
	return modelpkg
}

// for satisfaction of GongStruct interface
func (modelpkg *ModelPkg) GetName() (res string) {
	return modelpkg.Name
}

// for satisfaction of GongStruct interface
func (modelpkg *ModelPkg) SetName(name string) {
	modelpkg.Name = name
}

// Stage puts pointertogongstructfield to the model stage
func (pointertogongstructfield *PointerToGongStructField) Stage(stage *Stage) *PointerToGongStructField {

	if _, ok := stage.PointerToGongStructFields[pointertogongstructfield]; !ok {
		stage.PointerToGongStructFields[pointertogongstructfield] = struct{}{}
		stage.PointerToGongStructFieldMap_Staged_Order[pointertogongstructfield] = stage.PointerToGongStructFieldOrder
		stage.PointerToGongStructFieldOrder++
	}
	stage.PointerToGongStructFields_mapString[pointertogongstructfield.Name] = pointertogongstructfield

	return pointertogongstructfield
}

// StagePreserveOrder puts pointertogongstructfield to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.PointerToGongStructFieldOrder
// - update stage.PointerToGongStructFieldOrder accordingly
func (pointertogongstructfield *PointerToGongStructField) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.PointerToGongStructFields[pointertogongstructfield]; !ok {
		stage.PointerToGongStructFields[pointertogongstructfield] = struct{}{}

		if order > stage.PointerToGongStructFieldOrder {
			stage.PointerToGongStructFieldOrder = order
		}
		stage.PointerToGongStructFieldMap_Staged_Order[pointertogongstructfield] = order
		stage.PointerToGongStructFieldOrder++
	}
	stage.PointerToGongStructFields_mapString[pointertogongstructfield.Name] = pointertogongstructfield
}

// Unstage removes pointertogongstructfield off the model stage
func (pointertogongstructfield *PointerToGongStructField) Unstage(stage *Stage) *PointerToGongStructField {
	delete(stage.PointerToGongStructFields, pointertogongstructfield)
	delete(stage.PointerToGongStructFieldMap_Staged_Order, pointertogongstructfield)
	delete(stage.PointerToGongStructFields_mapString, pointertogongstructfield.Name)

	return pointertogongstructfield
}

// UnstageVoid removes pointertogongstructfield off the model stage
func (pointertogongstructfield *PointerToGongStructField) UnstageVoid(stage *Stage) {
	delete(stage.PointerToGongStructFields, pointertogongstructfield)
	delete(stage.PointerToGongStructFieldMap_Staged_Order, pointertogongstructfield)
	delete(stage.PointerToGongStructFields_mapString, pointertogongstructfield.Name)
}

// commit pointertogongstructfield to the back repo (if it is already staged)
func (pointertogongstructfield *PointerToGongStructField) Commit(stage *Stage) *PointerToGongStructField {
	if _, ok := stage.PointerToGongStructFields[pointertogongstructfield]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitPointerToGongStructField(pointertogongstructfield)
		}
	}
	return pointertogongstructfield
}

func (pointertogongstructfield *PointerToGongStructField) CommitVoid(stage *Stage) {
	pointertogongstructfield.Commit(stage)
}

func (pointertogongstructfield *PointerToGongStructField) StageVoid(stage *Stage) {
	pointertogongstructfield.Stage(stage)
}

// Checkout pointertogongstructfield to the back repo (if it is already staged)
func (pointertogongstructfield *PointerToGongStructField) Checkout(stage *Stage) *PointerToGongStructField {
	if _, ok := stage.PointerToGongStructFields[pointertogongstructfield]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutPointerToGongStructField(pointertogongstructfield)
		}
	}
	return pointertogongstructfield
}

// for satisfaction of GongStruct interface
func (pointertogongstructfield *PointerToGongStructField) GetName() (res string) {
	return pointertogongstructfield.Name
}

// for satisfaction of GongStruct interface
func (pointertogongstructfield *PointerToGongStructField) SetName(name string) {
	pointertogongstructfield.Name = name
}

// Stage puts sliceofpointertogongstructfield to the model stage
func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) Stage(stage *Stage) *SliceOfPointerToGongStructField {

	if _, ok := stage.SliceOfPointerToGongStructFields[sliceofpointertogongstructfield]; !ok {
		stage.SliceOfPointerToGongStructFields[sliceofpointertogongstructfield] = struct{}{}
		stage.SliceOfPointerToGongStructFieldMap_Staged_Order[sliceofpointertogongstructfield] = stage.SliceOfPointerToGongStructFieldOrder
		stage.SliceOfPointerToGongStructFieldOrder++
	}
	stage.SliceOfPointerToGongStructFields_mapString[sliceofpointertogongstructfield.Name] = sliceofpointertogongstructfield

	return sliceofpointertogongstructfield
}

// StagePreserveOrder puts sliceofpointertogongstructfield to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.SliceOfPointerToGongStructFieldOrder
// - update stage.SliceOfPointerToGongStructFieldOrder accordingly
func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.SliceOfPointerToGongStructFields[sliceofpointertogongstructfield]; !ok {
		stage.SliceOfPointerToGongStructFields[sliceofpointertogongstructfield] = struct{}{}

		if order > stage.SliceOfPointerToGongStructFieldOrder {
			stage.SliceOfPointerToGongStructFieldOrder = order
		}
		stage.SliceOfPointerToGongStructFieldMap_Staged_Order[sliceofpointertogongstructfield] = order
		stage.SliceOfPointerToGongStructFieldOrder++
	}
	stage.SliceOfPointerToGongStructFields_mapString[sliceofpointertogongstructfield.Name] = sliceofpointertogongstructfield
}

// Unstage removes sliceofpointertogongstructfield off the model stage
func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) Unstage(stage *Stage) *SliceOfPointerToGongStructField {
	delete(stage.SliceOfPointerToGongStructFields, sliceofpointertogongstructfield)
	delete(stage.SliceOfPointerToGongStructFieldMap_Staged_Order, sliceofpointertogongstructfield)
	delete(stage.SliceOfPointerToGongStructFields_mapString, sliceofpointertogongstructfield.Name)

	return sliceofpointertogongstructfield
}

// UnstageVoid removes sliceofpointertogongstructfield off the model stage
func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) UnstageVoid(stage *Stage) {
	delete(stage.SliceOfPointerToGongStructFields, sliceofpointertogongstructfield)
	delete(stage.SliceOfPointerToGongStructFieldMap_Staged_Order, sliceofpointertogongstructfield)
	delete(stage.SliceOfPointerToGongStructFields_mapString, sliceofpointertogongstructfield.Name)
}

// commit sliceofpointertogongstructfield to the back repo (if it is already staged)
func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) Commit(stage *Stage) *SliceOfPointerToGongStructField {
	if _, ok := stage.SliceOfPointerToGongStructFields[sliceofpointertogongstructfield]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSliceOfPointerToGongStructField(sliceofpointertogongstructfield)
		}
	}
	return sliceofpointertogongstructfield
}

func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) CommitVoid(stage *Stage) {
	sliceofpointertogongstructfield.Commit(stage)
}

func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) StageVoid(stage *Stage) {
	sliceofpointertogongstructfield.Stage(stage)
}

// Checkout sliceofpointertogongstructfield to the back repo (if it is already staged)
func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) Checkout(stage *Stage) *SliceOfPointerToGongStructField {
	if _, ok := stage.SliceOfPointerToGongStructFields[sliceofpointertogongstructfield]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSliceOfPointerToGongStructField(sliceofpointertogongstructfield)
		}
	}
	return sliceofpointertogongstructfield
}

// for satisfaction of GongStruct interface
func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) GetName() (res string) {
	return sliceofpointertogongstructfield.Name
}

// for satisfaction of GongStruct interface
func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) SetName(name string) {
	sliceofpointertogongstructfield.Name = name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMGongBasicField(GongBasicField *GongBasicField)
	CreateORMGongEnum(GongEnum *GongEnum)
	CreateORMGongEnumValue(GongEnumValue *GongEnumValue)
	CreateORMGongLink(GongLink *GongLink)
	CreateORMGongNote(GongNote *GongNote)
	CreateORMGongStruct(GongStruct *GongStruct)
	CreateORMGongTimeField(GongTimeField *GongTimeField)
	CreateORMMetaReference(MetaReference *MetaReference)
	CreateORMModelPkg(ModelPkg *ModelPkg)
	CreateORMPointerToGongStructField(PointerToGongStructField *PointerToGongStructField)
	CreateORMSliceOfPointerToGongStructField(SliceOfPointerToGongStructField *SliceOfPointerToGongStructField)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMGongBasicField(GongBasicField *GongBasicField)
	DeleteORMGongEnum(GongEnum *GongEnum)
	DeleteORMGongEnumValue(GongEnumValue *GongEnumValue)
	DeleteORMGongLink(GongLink *GongLink)
	DeleteORMGongNote(GongNote *GongNote)
	DeleteORMGongStruct(GongStruct *GongStruct)
	DeleteORMGongTimeField(GongTimeField *GongTimeField)
	DeleteORMMetaReference(MetaReference *MetaReference)
	DeleteORMModelPkg(ModelPkg *ModelPkg)
	DeleteORMPointerToGongStructField(PointerToGongStructField *PointerToGongStructField)
	DeleteORMSliceOfPointerToGongStructField(SliceOfPointerToGongStructField *SliceOfPointerToGongStructField)
}

func (stage *Stage) Reset() { // insertion point for array reset
	stage.GongBasicFields = make(map[*GongBasicField]struct{})
	stage.GongBasicFields_mapString = make(map[string]*GongBasicField)
	stage.GongBasicFieldMap_Staged_Order = make(map[*GongBasicField]uint)
	stage.GongBasicFieldOrder = 0

	stage.GongEnums = make(map[*GongEnum]struct{})
	stage.GongEnums_mapString = make(map[string]*GongEnum)
	stage.GongEnumMap_Staged_Order = make(map[*GongEnum]uint)
	stage.GongEnumOrder = 0

	stage.GongEnumValues = make(map[*GongEnumValue]struct{})
	stage.GongEnumValues_mapString = make(map[string]*GongEnumValue)
	stage.GongEnumValueMap_Staged_Order = make(map[*GongEnumValue]uint)
	stage.GongEnumValueOrder = 0

	stage.GongLinks = make(map[*GongLink]struct{})
	stage.GongLinks_mapString = make(map[string]*GongLink)
	stage.GongLinkMap_Staged_Order = make(map[*GongLink]uint)
	stage.GongLinkOrder = 0

	stage.GongNotes = make(map[*GongNote]struct{})
	stage.GongNotes_mapString = make(map[string]*GongNote)
	stage.GongNoteMap_Staged_Order = make(map[*GongNote]uint)
	stage.GongNoteOrder = 0

	stage.GongStructs = make(map[*GongStruct]struct{})
	stage.GongStructs_mapString = make(map[string]*GongStruct)
	stage.GongStructMap_Staged_Order = make(map[*GongStruct]uint)
	stage.GongStructOrder = 0

	stage.GongTimeFields = make(map[*GongTimeField]struct{})
	stage.GongTimeFields_mapString = make(map[string]*GongTimeField)
	stage.GongTimeFieldMap_Staged_Order = make(map[*GongTimeField]uint)
	stage.GongTimeFieldOrder = 0

	stage.MetaReferences = make(map[*MetaReference]struct{})
	stage.MetaReferences_mapString = make(map[string]*MetaReference)
	stage.MetaReferenceMap_Staged_Order = make(map[*MetaReference]uint)
	stage.MetaReferenceOrder = 0

	stage.ModelPkgs = make(map[*ModelPkg]struct{})
	stage.ModelPkgs_mapString = make(map[string]*ModelPkg)
	stage.ModelPkgMap_Staged_Order = make(map[*ModelPkg]uint)
	stage.ModelPkgOrder = 0

	stage.PointerToGongStructFields = make(map[*PointerToGongStructField]struct{})
	stage.PointerToGongStructFields_mapString = make(map[string]*PointerToGongStructField)
	stage.PointerToGongStructFieldMap_Staged_Order = make(map[*PointerToGongStructField]uint)
	stage.PointerToGongStructFieldOrder = 0

	stage.SliceOfPointerToGongStructFields = make(map[*SliceOfPointerToGongStructField]struct{})
	stage.SliceOfPointerToGongStructFields_mapString = make(map[string]*SliceOfPointerToGongStructField)
	stage.SliceOfPointerToGongStructFieldMap_Staged_Order = make(map[*SliceOfPointerToGongStructField]uint)
	stage.SliceOfPointerToGongStructFieldOrder = 0

	if stage.GetProbeIF() != nil {
		stage.GetProbeIF().ResetNotifications()
	}
	if stage.IsInDeltaMode() {
		stage.ComputeReference()
	}
}

func (stage *Stage) Nil() { // insertion point for array nil
	stage.GongBasicFields = nil
	stage.GongBasicFields_mapString = nil

	stage.GongEnums = nil
	stage.GongEnums_mapString = nil

	stage.GongEnumValues = nil
	stage.GongEnumValues_mapString = nil

	stage.GongLinks = nil
	stage.GongLinks_mapString = nil

	stage.GongNotes = nil
	stage.GongNotes_mapString = nil

	stage.GongStructs = nil
	stage.GongStructs_mapString = nil

	stage.GongTimeFields = nil
	stage.GongTimeFields_mapString = nil

	stage.MetaReferences = nil
	stage.MetaReferences_mapString = nil

	stage.ModelPkgs = nil
	stage.ModelPkgs_mapString = nil

	stage.PointerToGongStructFields = nil
	stage.PointerToGongStructFields_mapString = nil

	stage.SliceOfPointerToGongStructFields = nil
	stage.SliceOfPointerToGongStructFields_mapString = nil

}

func (stage *Stage) Unstage() { // insertion point for array nil
	for gongbasicfield := range stage.GongBasicFields {
		gongbasicfield.Unstage(stage)
	}

	for gongenum := range stage.GongEnums {
		gongenum.Unstage(stage)
	}

	for gongenumvalue := range stage.GongEnumValues {
		gongenumvalue.Unstage(stage)
	}

	for gonglink := range stage.GongLinks {
		gonglink.Unstage(stage)
	}

	for gongnote := range stage.GongNotes {
		gongnote.Unstage(stage)
	}

	for gongstruct := range stage.GongStructs {
		gongstruct.Unstage(stage)
	}

	for gongtimefield := range stage.GongTimeFields {
		gongtimefield.Unstage(stage)
	}

	for metareference := range stage.MetaReferences {
		metareference.Unstage(stage)
	}

	for modelpkg := range stage.ModelPkgs {
		modelpkg.Unstage(stage)
	}

	for pointertogongstructfield := range stage.PointerToGongStructFields {
		pointertogongstructfield.Unstage(stage)
	}

	for sliceofpointertogongstructfield := range stage.SliceOfPointerToGongStructFields {
		sliceofpointertogongstructfield.Unstage(stage)
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
	case map[*GongBasicField]any:
		return any(&stage.GongBasicFields).(*Type)
	case map[*GongEnum]any:
		return any(&stage.GongEnums).(*Type)
	case map[*GongEnumValue]any:
		return any(&stage.GongEnumValues).(*Type)
	case map[*GongLink]any:
		return any(&stage.GongLinks).(*Type)
	case map[*GongNote]any:
		return any(&stage.GongNotes).(*Type)
	case map[*GongStruct]any:
		return any(&stage.GongStructs).(*Type)
	case map[*GongTimeField]any:
		return any(&stage.GongTimeFields).(*Type)
	case map[*MetaReference]any:
		return any(&stage.MetaReferences).(*Type)
	case map[*ModelPkg]any:
		return any(&stage.ModelPkgs).(*Type)
	case map[*PointerToGongStructField]any:
		return any(&stage.PointerToGongStructFields).(*Type)
	case map[*SliceOfPointerToGongStructField]any:
		return any(&stage.SliceOfPointerToGongStructFields).(*Type)
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
	case *GongBasicField:
		return any(stage.GongBasicFields_mapString).(map[string]Type)
	case *GongEnum:
		return any(stage.GongEnums_mapString).(map[string]Type)
	case *GongEnumValue:
		return any(stage.GongEnumValues_mapString).(map[string]Type)
	case *GongLink:
		return any(stage.GongLinks_mapString).(map[string]Type)
	case *GongNote:
		return any(stage.GongNotes_mapString).(map[string]Type)
	case *GongStruct:
		return any(stage.GongStructs_mapString).(map[string]Type)
	case *GongTimeField:
		return any(stage.GongTimeFields_mapString).(map[string]Type)
	case *MetaReference:
		return any(stage.MetaReferences_mapString).(map[string]Type)
	case *ModelPkg:
		return any(stage.ModelPkgs_mapString).(map[string]Type)
	case *PointerToGongStructField:
		return any(stage.PointerToGongStructFields_mapString).(map[string]Type)
	case *SliceOfPointerToGongStructField:
		return any(stage.SliceOfPointerToGongStructFields_mapString).(map[string]Type)
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
	case GongBasicField:
		return any(&stage.GongBasicFields).(*map[*Type]struct{})
	case GongEnum:
		return any(&stage.GongEnums).(*map[*Type]struct{})
	case GongEnumValue:
		return any(&stage.GongEnumValues).(*map[*Type]struct{})
	case GongLink:
		return any(&stage.GongLinks).(*map[*Type]struct{})
	case GongNote:
		return any(&stage.GongNotes).(*map[*Type]struct{})
	case GongStruct:
		return any(&stage.GongStructs).(*map[*Type]struct{})
	case GongTimeField:
		return any(&stage.GongTimeFields).(*map[*Type]struct{})
	case MetaReference:
		return any(&stage.MetaReferences).(*map[*Type]struct{})
	case ModelPkg:
		return any(&stage.ModelPkgs).(*map[*Type]struct{})
	case PointerToGongStructField:
		return any(&stage.PointerToGongStructFields).(*map[*Type]struct{})
	case SliceOfPointerToGongStructField:
		return any(&stage.SliceOfPointerToGongStructFields).(*map[*Type]struct{})
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
	case *GongBasicField:
		return any(&stage.GongBasicFields).(*map[Type]struct{})
	case *GongEnum:
		return any(&stage.GongEnums).(*map[Type]struct{})
	case *GongEnumValue:
		return any(&stage.GongEnumValues).(*map[Type]struct{})
	case *GongLink:
		return any(&stage.GongLinks).(*map[Type]struct{})
	case *GongNote:
		return any(&stage.GongNotes).(*map[Type]struct{})
	case *GongStruct:
		return any(&stage.GongStructs).(*map[Type]struct{})
	case *GongTimeField:
		return any(&stage.GongTimeFields).(*map[Type]struct{})
	case *MetaReference:
		return any(&stage.MetaReferences).(*map[Type]struct{})
	case *ModelPkg:
		return any(&stage.ModelPkgs).(*map[Type]struct{})
	case *PointerToGongStructField:
		return any(&stage.PointerToGongStructFields).(*map[Type]struct{})
	case *SliceOfPointerToGongStructField:
		return any(&stage.SliceOfPointerToGongStructFields).(*map[Type]struct{})
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
	case GongBasicField:
		return any(&stage.GongBasicFields_mapString).(*map[string]*Type)
	case GongEnum:
		return any(&stage.GongEnums_mapString).(*map[string]*Type)
	case GongEnumValue:
		return any(&stage.GongEnumValues_mapString).(*map[string]*Type)
	case GongLink:
		return any(&stage.GongLinks_mapString).(*map[string]*Type)
	case GongNote:
		return any(&stage.GongNotes_mapString).(*map[string]*Type)
	case GongStruct:
		return any(&stage.GongStructs_mapString).(*map[string]*Type)
	case GongTimeField:
		return any(&stage.GongTimeFields_mapString).(*map[string]*Type)
	case MetaReference:
		return any(&stage.MetaReferences_mapString).(*map[string]*Type)
	case ModelPkg:
		return any(&stage.ModelPkgs_mapString).(*map[string]*Type)
	case PointerToGongStructField:
		return any(&stage.PointerToGongStructFields_mapString).(*map[string]*Type)
	case SliceOfPointerToGongStructField:
		return any(&stage.SliceOfPointerToGongStructFields_mapString).(*map[string]*Type)
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
	case GongBasicField:
		return any(&GongBasicField{
			// Initialisation of associations
			// field is initialized with an instance of GongEnum with the name of the field
			GongEnum: &GongEnum{Name: "GongEnum"},
		}).(*Type)
	case GongEnum:
		return any(&GongEnum{
			// Initialisation of associations
			// field is initialized with an instance of GongEnumValue with the name of the field
			GongEnumValues: []*GongEnumValue{{Name: "GongEnumValues"}},
		}).(*Type)
	case GongEnumValue:
		return any(&GongEnumValue{
			// Initialisation of associations
		}).(*Type)
	case GongLink:
		return any(&GongLink{
			// Initialisation of associations
		}).(*Type)
	case GongNote:
		return any(&GongNote{
			// Initialisation of associations
			// field is initialized with an instance of GongLink with the name of the field
			Links: []*GongLink{{Name: "Links"}},
		}).(*Type)
	case GongStruct:
		return any(&GongStruct{
			// Initialisation of associations
			// field is initialized with an instance of GongBasicField with the name of the field
			GongBasicFields: []*GongBasicField{{Name: "GongBasicFields"}},
			// field is initialized with an instance of GongTimeField with the name of the field
			GongTimeFields: []*GongTimeField{{Name: "GongTimeFields"}},
			// field is initialized with an instance of PointerToGongStructField with the name of the field
			PointerToGongStructFields: []*PointerToGongStructField{{Name: "PointerToGongStructFields"}},
			// field is initialized with an instance of SliceOfPointerToGongStructField with the name of the field
			SliceOfPointerToGongStructFields: []*SliceOfPointerToGongStructField{{Name: "SliceOfPointerToGongStructFields"}},
		}).(*Type)
	case GongTimeField:
		return any(&GongTimeField{
			// Initialisation of associations
		}).(*Type)
	case MetaReference:
		return any(&MetaReference{
			// Initialisation of associations
		}).(*Type)
	case ModelPkg:
		return any(&ModelPkg{
			// Initialisation of associations
		}).(*Type)
	case PointerToGongStructField:
		return any(&PointerToGongStructField{
			// Initialisation of associations
			// field is initialized with an instance of GongStruct with the name of the field
			GongStruct: &GongStruct{Name: "GongStruct"},
		}).(*Type)
	case SliceOfPointerToGongStructField:
		return any(&SliceOfPointerToGongStructField{
			// Initialisation of associations
			// field is initialized with an instance of GongStruct with the name of the field
			GongStruct: &GongStruct{Name: "GongStruct"},
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
	// reverse maps of direct associations of GongBasicField
	case GongBasicField:
		switch fieldname {
		// insertion point for per direct association field
		case "GongEnum":
			res := make(map[*GongEnum][]*GongBasicField)
			for gongbasicfield := range stage.GongBasicFields {
				if gongbasicfield.GongEnum != nil {
					gongenum_ := gongbasicfield.GongEnum
					var gongbasicfields []*GongBasicField
					_, ok := res[gongenum_]
					if ok {
						gongbasicfields = res[gongenum_]
					} else {
						gongbasicfields = make([]*GongBasicField, 0)
					}
					gongbasicfields = append(gongbasicfields, gongbasicfield)
					res[gongenum_] = gongbasicfields
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of GongEnum
	case GongEnum:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GongEnumValue
	case GongEnumValue:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GongLink
	case GongLink:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GongNote
	case GongNote:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GongStruct
	case GongStruct:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GongTimeField
	case GongTimeField:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MetaReference
	case MetaReference:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ModelPkg
	case ModelPkg:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of PointerToGongStructField
	case PointerToGongStructField:
		switch fieldname {
		// insertion point for per direct association field
		case "GongStruct":
			res := make(map[*GongStruct][]*PointerToGongStructField)
			for pointertogongstructfield := range stage.PointerToGongStructFields {
				if pointertogongstructfield.GongStruct != nil {
					gongstruct_ := pointertogongstructfield.GongStruct
					var pointertogongstructfields []*PointerToGongStructField
					_, ok := res[gongstruct_]
					if ok {
						pointertogongstructfields = res[gongstruct_]
					} else {
						pointertogongstructfields = make([]*PointerToGongStructField, 0)
					}
					pointertogongstructfields = append(pointertogongstructfields, pointertogongstructfield)
					res[gongstruct_] = pointertogongstructfields
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of SliceOfPointerToGongStructField
	case SliceOfPointerToGongStructField:
		switch fieldname {
		// insertion point for per direct association field
		case "GongStruct":
			res := make(map[*GongStruct][]*SliceOfPointerToGongStructField)
			for sliceofpointertogongstructfield := range stage.SliceOfPointerToGongStructFields {
				if sliceofpointertogongstructfield.GongStruct != nil {
					gongstruct_ := sliceofpointertogongstructfield.GongStruct
					var sliceofpointertogongstructfields []*SliceOfPointerToGongStructField
					_, ok := res[gongstruct_]
					if ok {
						sliceofpointertogongstructfields = res[gongstruct_]
					} else {
						sliceofpointertogongstructfields = make([]*SliceOfPointerToGongStructField, 0)
					}
					sliceofpointertogongstructfields = append(sliceofpointertogongstructfields, sliceofpointertogongstructfield)
					res[gongstruct_] = sliceofpointertogongstructfields
				}
			}
			return any(res).(map[*End][]*Start)
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
	// reverse maps of direct associations of GongBasicField
	case GongBasicField:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GongEnum
	case GongEnum:
		switch fieldname {
		// insertion point for per direct association field
		case "GongEnumValues":
			res := make(map[*GongEnumValue][]*GongEnum)
			for gongenum := range stage.GongEnums {
				for _, gongenumvalue_ := range gongenum.GongEnumValues {
					res[gongenumvalue_] = append(res[gongenumvalue_], gongenum)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of GongEnumValue
	case GongEnumValue:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GongLink
	case GongLink:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GongNote
	case GongNote:
		switch fieldname {
		// insertion point for per direct association field
		case "Links":
			res := make(map[*GongLink][]*GongNote)
			for gongnote := range stage.GongNotes {
				for _, gonglink_ := range gongnote.Links {
					res[gonglink_] = append(res[gonglink_], gongnote)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of GongStruct
	case GongStruct:
		switch fieldname {
		// insertion point for per direct association field
		case "GongBasicFields":
			res := make(map[*GongBasicField][]*GongStruct)
			for gongstruct := range stage.GongStructs {
				for _, gongbasicfield_ := range gongstruct.GongBasicFields {
					res[gongbasicfield_] = append(res[gongbasicfield_], gongstruct)
				}
			}
			return any(res).(map[*End][]*Start)
		case "GongTimeFields":
			res := make(map[*GongTimeField][]*GongStruct)
			for gongstruct := range stage.GongStructs {
				for _, gongtimefield_ := range gongstruct.GongTimeFields {
					res[gongtimefield_] = append(res[gongtimefield_], gongstruct)
				}
			}
			return any(res).(map[*End][]*Start)
		case "PointerToGongStructFields":
			res := make(map[*PointerToGongStructField][]*GongStruct)
			for gongstruct := range stage.GongStructs {
				for _, pointertogongstructfield_ := range gongstruct.PointerToGongStructFields {
					res[pointertogongstructfield_] = append(res[pointertogongstructfield_], gongstruct)
				}
			}
			return any(res).(map[*End][]*Start)
		case "SliceOfPointerToGongStructFields":
			res := make(map[*SliceOfPointerToGongStructField][]*GongStruct)
			for gongstruct := range stage.GongStructs {
				for _, sliceofpointertogongstructfield_ := range gongstruct.SliceOfPointerToGongStructFields {
					res[sliceofpointertogongstructfield_] = append(res[sliceofpointertogongstructfield_], gongstruct)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of GongTimeField
	case GongTimeField:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MetaReference
	case MetaReference:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ModelPkg
	case ModelPkg:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of PointerToGongStructField
	case PointerToGongStructField:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SliceOfPointerToGongStructField
	case SliceOfPointerToGongStructField:
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
	case *GongBasicField:
		res = "GongBasicField"
	case *GongEnum:
		res = "GongEnum"
	case *GongEnumValue:
		res = "GongEnumValue"
	case *GongLink:
		res = "GongLink"
	case *GongNote:
		res = "GongNote"
	case *GongStruct:
		res = "GongStruct"
	case *GongTimeField:
		res = "GongTimeField"
	case *MetaReference:
		res = "MetaReference"
	case *ModelPkg:
		res = "ModelPkg"
	case *PointerToGongStructField:
		res = "PointerToGongStructField"
	case *SliceOfPointerToGongStructField:
		res = "SliceOfPointerToGongStructField"
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
	case *GongBasicField:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "GongStruct"
		rf.Fieldname = "GongBasicFields"
		res = append(res, rf)
	case *GongEnum:
		var rf ReverseField
		_ = rf
	case *GongEnumValue:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "GongEnum"
		rf.Fieldname = "GongEnumValues"
		res = append(res, rf)
	case *GongLink:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "GongNote"
		rf.Fieldname = "Links"
		res = append(res, rf)
	case *GongNote:
		var rf ReverseField
		_ = rf
	case *GongStruct:
		var rf ReverseField
		_ = rf
	case *GongTimeField:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "GongStruct"
		rf.Fieldname = "GongTimeFields"
		res = append(res, rf)
	case *MetaReference:
		var rf ReverseField
		_ = rf
	case *ModelPkg:
		var rf ReverseField
		_ = rf
	case *PointerToGongStructField:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "GongStruct"
		rf.Fieldname = "PointerToGongStructFields"
		res = append(res, rf)
	case *SliceOfPointerToGongStructField:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "GongStruct"
		rf.Fieldname = "SliceOfPointerToGongStructFields"
		res = append(res, rf)
	}
	return
}

// insertion point for get fields header method
func (gongbasicfield *GongBasicField) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "BasicKindName",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "GongEnum",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "GongEnum",
		},
		{
			Name:               "DeclaredType",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "CompositeStructName",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Index",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "IsTextArea",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "IsBespokeWidth",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "BespokeWidth",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "IsBespokeHeight",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "BespokeHeight",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (gongenum *GongEnum) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Type",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "GongEnumValues",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "GongEnumValue",
		},
	}
	return
}

func (gongenumvalue *GongEnumValue) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Value",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (gonglink *GongLink) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Recv",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "ImportPath",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (gongnote *GongNote) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Body",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "BodyHTML",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "Links",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "GongLink",
		},
	}
	return
}

func (gongstruct *GongStruct) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "GongBasicFields",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "GongBasicField",
		},
		{
			Name:                 "GongTimeFields",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "GongTimeField",
		},
		{
			Name:                 "PointerToGongStructFields",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "PointerToGongStructField",
		},
		{
			Name:                 "SliceOfPointerToGongStructFields",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "SliceOfPointerToGongStructField",
		},
		{
			Name:               "HasOnAfterUpdateSignature",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "IsIgnoredForFront",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (gongtimefield *GongTimeField) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Index",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "CompositeStructName",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "BespokeTimeFormat",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (metareference *MetaReference) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (modelpkg *ModelPkg) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "PkgPath",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "PathToGoSubDirectory",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "OrmPkgGenPath",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "DbOrmPkgGenPath",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "DbLiteOrmPkgGenPath",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "DbPkgGenPath",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "ControllersPkgGenPath",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "FullstackPkgGenPath",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StackPkgGenPath",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Level1StackPkgGenPath",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StaticPkgGenPath",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "ProbePkgGenPath",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "NgWorkspacePath",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "NgWorkspaceName",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "NgDataLibrarySourceCodeDirectory",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "NgSpecificLibrarySourceCodeDirectory",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "MaterialLibDatamodelTargetPath",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (pointertogongstructfield *PointerToGongStructField) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "GongStruct",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "GongStruct",
		},
		{
			Name:               "Index",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "CompositeStructName",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "IsType",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "GongStruct",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "GongStruct",
		},
		{
			Name:               "Index",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "CompositeStructName",
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
func (gongbasicfield *GongBasicField) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = gongbasicfield.Name
	case "BasicKindName":
		res.valueString = gongbasicfield.BasicKindName
	case "GongEnum":
		res.GongFieldValueType = GongFieldValueTypePointer
		if gongbasicfield.GongEnum != nil {
			res.valueString = gongbasicfield.GongEnum.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, gongbasicfield.GongEnum))
		}
	case "DeclaredType":
		res.valueString = gongbasicfield.DeclaredType
	case "CompositeStructName":
		res.valueString = gongbasicfield.CompositeStructName
	case "Index":
		res.valueString = fmt.Sprintf("%d", gongbasicfield.Index)
		res.valueInt = gongbasicfield.Index
		res.GongFieldValueType = GongFieldValueTypeInt
	case "IsTextArea":
		res.valueString = fmt.Sprintf("%t", gongbasicfield.IsTextArea)
		res.valueBool = gongbasicfield.IsTextArea
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsBespokeWidth":
		res.valueString = fmt.Sprintf("%t", gongbasicfield.IsBespokeWidth)
		res.valueBool = gongbasicfield.IsBespokeWidth
		res.GongFieldValueType = GongFieldValueTypeBool
	case "BespokeWidth":
		res.valueString = fmt.Sprintf("%d", gongbasicfield.BespokeWidth)
		res.valueInt = gongbasicfield.BespokeWidth
		res.GongFieldValueType = GongFieldValueTypeInt
	case "IsBespokeHeight":
		res.valueString = fmt.Sprintf("%t", gongbasicfield.IsBespokeHeight)
		res.valueBool = gongbasicfield.IsBespokeHeight
		res.GongFieldValueType = GongFieldValueTypeBool
	case "BespokeHeight":
		res.valueString = fmt.Sprintf("%d", gongbasicfield.BespokeHeight)
		res.valueInt = gongbasicfield.BespokeHeight
		res.GongFieldValueType = GongFieldValueTypeInt
	}
	return
}
func (gongenum *GongEnum) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = gongenum.Name
	case "Type":
		enum := gongenum.Type
		res.valueString = enum.ToCodeString()
	case "GongEnumValues":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range gongenum.GongEnumValues {
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
func (gongenumvalue *GongEnumValue) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = gongenumvalue.Name
	case "Value":
		res.valueString = gongenumvalue.Value
	}
	return
}
func (gonglink *GongLink) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = gonglink.Name
	case "Recv":
		res.valueString = gonglink.Recv
	case "ImportPath":
		res.valueString = gonglink.ImportPath
	}
	return
}
func (gongnote *GongNote) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = gongnote.Name
	case "Body":
		res.valueString = gongnote.Body
	case "BodyHTML":
		res.valueString = gongnote.BodyHTML
	case "Links":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range gongnote.Links {
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
func (gongstruct *GongStruct) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = gongstruct.Name
	case "GongBasicFields":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range gongstruct.GongBasicFields {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "GongTimeFields":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range gongstruct.GongTimeFields {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "PointerToGongStructFields":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range gongstruct.PointerToGongStructFields {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "SliceOfPointerToGongStructFields":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range gongstruct.SliceOfPointerToGongStructFields {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "HasOnAfterUpdateSignature":
		res.valueString = fmt.Sprintf("%t", gongstruct.HasOnAfterUpdateSignature)
		res.valueBool = gongstruct.HasOnAfterUpdateSignature
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsIgnoredForFront":
		res.valueString = fmt.Sprintf("%t", gongstruct.IsIgnoredForFront)
		res.valueBool = gongstruct.IsIgnoredForFront
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}
func (gongtimefield *GongTimeField) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = gongtimefield.Name
	case "Index":
		res.valueString = fmt.Sprintf("%d", gongtimefield.Index)
		res.valueInt = gongtimefield.Index
		res.GongFieldValueType = GongFieldValueTypeInt
	case "CompositeStructName":
		res.valueString = gongtimefield.CompositeStructName
	case "BespokeTimeFormat":
		res.valueString = gongtimefield.BespokeTimeFormat
	}
	return
}
func (metareference *MetaReference) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = metareference.Name
	}
	return
}
func (modelpkg *ModelPkg) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = modelpkg.Name
	case "PkgPath":
		res.valueString = modelpkg.PkgPath
	case "PathToGoSubDirectory":
		res.valueString = modelpkg.PathToGoSubDirectory
	case "OrmPkgGenPath":
		res.valueString = modelpkg.OrmPkgGenPath
	case "DbOrmPkgGenPath":
		res.valueString = modelpkg.DbOrmPkgGenPath
	case "DbLiteOrmPkgGenPath":
		res.valueString = modelpkg.DbLiteOrmPkgGenPath
	case "DbPkgGenPath":
		res.valueString = modelpkg.DbPkgGenPath
	case "ControllersPkgGenPath":
		res.valueString = modelpkg.ControllersPkgGenPath
	case "FullstackPkgGenPath":
		res.valueString = modelpkg.FullstackPkgGenPath
	case "StackPkgGenPath":
		res.valueString = modelpkg.StackPkgGenPath
	case "Level1StackPkgGenPath":
		res.valueString = modelpkg.Level1StackPkgGenPath
	case "StaticPkgGenPath":
		res.valueString = modelpkg.StaticPkgGenPath
	case "ProbePkgGenPath":
		res.valueString = modelpkg.ProbePkgGenPath
	case "NgWorkspacePath":
		res.valueString = modelpkg.NgWorkspacePath
	case "NgWorkspaceName":
		res.valueString = modelpkg.NgWorkspaceName
	case "NgDataLibrarySourceCodeDirectory":
		res.valueString = modelpkg.NgDataLibrarySourceCodeDirectory
	case "NgSpecificLibrarySourceCodeDirectory":
		res.valueString = modelpkg.NgSpecificLibrarySourceCodeDirectory
	case "MaterialLibDatamodelTargetPath":
		res.valueString = modelpkg.MaterialLibDatamodelTargetPath
	}
	return
}
func (pointertogongstructfield *PointerToGongStructField) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = pointertogongstructfield.Name
	case "GongStruct":
		res.GongFieldValueType = GongFieldValueTypePointer
		if pointertogongstructfield.GongStruct != nil {
			res.valueString = pointertogongstructfield.GongStruct.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, pointertogongstructfield.GongStruct))
		}
	case "Index":
		res.valueString = fmt.Sprintf("%d", pointertogongstructfield.Index)
		res.valueInt = pointertogongstructfield.Index
		res.GongFieldValueType = GongFieldValueTypeInt
	case "CompositeStructName":
		res.valueString = pointertogongstructfield.CompositeStructName
	case "IsType":
		res.valueString = fmt.Sprintf("%t", pointertogongstructfield.IsType)
		res.valueBool = pointertogongstructfield.IsType
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}
func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = sliceofpointertogongstructfield.Name
	case "GongStruct":
		res.GongFieldValueType = GongFieldValueTypePointer
		if sliceofpointertogongstructfield.GongStruct != nil {
			res.valueString = sliceofpointertogongstructfield.GongStruct.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, sliceofpointertogongstructfield.GongStruct))
		}
	case "Index":
		res.valueString = fmt.Sprintf("%d", sliceofpointertogongstructfield.Index)
		res.valueInt = sliceofpointertogongstructfield.Index
		res.GongFieldValueType = GongFieldValueTypeInt
	case "CompositeStructName":
		res.valueString = sliceofpointertogongstructfield.CompositeStructName
	}
	return
}
func GetFieldStringValueFromPointer(instance GongstructIF, fieldName string, stage *Stage) (res GongFieldValue) {

	res = instance.GongGetFieldValue(fieldName, stage)
	return
}

// insertion point for generic set gongstruct field value
func (gongbasicfield *GongBasicField) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		gongbasicfield.Name = value.GetValueString()
	case "BasicKindName":
		gongbasicfield.BasicKindName = value.GetValueString()
	case "GongEnum":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			gongbasicfield.GongEnum = nil
			for __instance__ := range stage.GongEnums {
				if stage.GongEnumMap_Staged_Order[__instance__] == uint(id) {
					gongbasicfield.GongEnum = __instance__
					break
				}
			}
		}
	case "DeclaredType":
		gongbasicfield.DeclaredType = value.GetValueString()
	case "CompositeStructName":
		gongbasicfield.CompositeStructName = value.GetValueString()
	case "Index":
		gongbasicfield.Index = int(value.GetValueInt())
	case "IsTextArea":
		gongbasicfield.IsTextArea = value.GetValueBool()
	case "IsBespokeWidth":
		gongbasicfield.IsBespokeWidth = value.GetValueBool()
	case "BespokeWidth":
		gongbasicfield.BespokeWidth = int(value.GetValueInt())
	case "IsBespokeHeight":
		gongbasicfield.IsBespokeHeight = value.GetValueBool()
	case "BespokeHeight":
		gongbasicfield.BespokeHeight = int(value.GetValueInt())
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (gongenum *GongEnum) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		gongenum.Name = value.GetValueString()
	case "Type":
		gongenum.Type.FromCodeString(value.GetValueString())
	case "GongEnumValues":
		gongenum.GongEnumValues = make([]*GongEnumValue, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.GongEnumValues {
					if stage.GongEnumValueMap_Staged_Order[__instance__] == uint(id) {
						gongenum.GongEnumValues = append(gongenum.GongEnumValues, __instance__)
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

func (gongenumvalue *GongEnumValue) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		gongenumvalue.Name = value.GetValueString()
	case "Value":
		gongenumvalue.Value = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (gonglink *GongLink) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		gonglink.Name = value.GetValueString()
	case "Recv":
		gonglink.Recv = value.GetValueString()
	case "ImportPath":
		gonglink.ImportPath = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (gongnote *GongNote) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		gongnote.Name = value.GetValueString()
	case "Body":
		gongnote.Body = value.GetValueString()
	case "BodyHTML":
		gongnote.BodyHTML = value.GetValueString()
	case "Links":
		gongnote.Links = make([]*GongLink, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.GongLinks {
					if stage.GongLinkMap_Staged_Order[__instance__] == uint(id) {
						gongnote.Links = append(gongnote.Links, __instance__)
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

func (gongstruct *GongStruct) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		gongstruct.Name = value.GetValueString()
	case "GongBasicFields":
		gongstruct.GongBasicFields = make([]*GongBasicField, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.GongBasicFields {
					if stage.GongBasicFieldMap_Staged_Order[__instance__] == uint(id) {
						gongstruct.GongBasicFields = append(gongstruct.GongBasicFields, __instance__)
						break
					}
				}
			}
		}
	case "GongTimeFields":
		gongstruct.GongTimeFields = make([]*GongTimeField, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.GongTimeFields {
					if stage.GongTimeFieldMap_Staged_Order[__instance__] == uint(id) {
						gongstruct.GongTimeFields = append(gongstruct.GongTimeFields, __instance__)
						break
					}
				}
			}
		}
	case "PointerToGongStructFields":
		gongstruct.PointerToGongStructFields = make([]*PointerToGongStructField, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.PointerToGongStructFields {
					if stage.PointerToGongStructFieldMap_Staged_Order[__instance__] == uint(id) {
						gongstruct.PointerToGongStructFields = append(gongstruct.PointerToGongStructFields, __instance__)
						break
					}
				}
			}
		}
	case "SliceOfPointerToGongStructFields":
		gongstruct.SliceOfPointerToGongStructFields = make([]*SliceOfPointerToGongStructField, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.SliceOfPointerToGongStructFields {
					if stage.SliceOfPointerToGongStructFieldMap_Staged_Order[__instance__] == uint(id) {
						gongstruct.SliceOfPointerToGongStructFields = append(gongstruct.SliceOfPointerToGongStructFields, __instance__)
						break
					}
				}
			}
		}
	case "HasOnAfterUpdateSignature":
		gongstruct.HasOnAfterUpdateSignature = value.GetValueBool()
	case "IsIgnoredForFront":
		gongstruct.IsIgnoredForFront = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (gongtimefield *GongTimeField) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		gongtimefield.Name = value.GetValueString()
	case "Index":
		gongtimefield.Index = int(value.GetValueInt())
	case "CompositeStructName":
		gongtimefield.CompositeStructName = value.GetValueString()
	case "BespokeTimeFormat":
		gongtimefield.BespokeTimeFormat = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (metareference *MetaReference) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		metareference.Name = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (modelpkg *ModelPkg) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		modelpkg.Name = value.GetValueString()
	case "PkgPath":
		modelpkg.PkgPath = value.GetValueString()
	case "PathToGoSubDirectory":
		modelpkg.PathToGoSubDirectory = value.GetValueString()
	case "OrmPkgGenPath":
		modelpkg.OrmPkgGenPath = value.GetValueString()
	case "DbOrmPkgGenPath":
		modelpkg.DbOrmPkgGenPath = value.GetValueString()
	case "DbLiteOrmPkgGenPath":
		modelpkg.DbLiteOrmPkgGenPath = value.GetValueString()
	case "DbPkgGenPath":
		modelpkg.DbPkgGenPath = value.GetValueString()
	case "ControllersPkgGenPath":
		modelpkg.ControllersPkgGenPath = value.GetValueString()
	case "FullstackPkgGenPath":
		modelpkg.FullstackPkgGenPath = value.GetValueString()
	case "StackPkgGenPath":
		modelpkg.StackPkgGenPath = value.GetValueString()
	case "Level1StackPkgGenPath":
		modelpkg.Level1StackPkgGenPath = value.GetValueString()
	case "StaticPkgGenPath":
		modelpkg.StaticPkgGenPath = value.GetValueString()
	case "ProbePkgGenPath":
		modelpkg.ProbePkgGenPath = value.GetValueString()
	case "NgWorkspacePath":
		modelpkg.NgWorkspacePath = value.GetValueString()
	case "NgWorkspaceName":
		modelpkg.NgWorkspaceName = value.GetValueString()
	case "NgDataLibrarySourceCodeDirectory":
		modelpkg.NgDataLibrarySourceCodeDirectory = value.GetValueString()
	case "NgSpecificLibrarySourceCodeDirectory":
		modelpkg.NgSpecificLibrarySourceCodeDirectory = value.GetValueString()
	case "MaterialLibDatamodelTargetPath":
		modelpkg.MaterialLibDatamodelTargetPath = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (pointertogongstructfield *PointerToGongStructField) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		pointertogongstructfield.Name = value.GetValueString()
	case "GongStruct":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			pointertogongstructfield.GongStruct = nil
			for __instance__ := range stage.GongStructs {
				if stage.GongStructMap_Staged_Order[__instance__] == uint(id) {
					pointertogongstructfield.GongStruct = __instance__
					break
				}
			}
		}
	case "Index":
		pointertogongstructfield.Index = int(value.GetValueInt())
	case "CompositeStructName":
		pointertogongstructfield.CompositeStructName = value.GetValueString()
	case "IsType":
		pointertogongstructfield.IsType = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		sliceofpointertogongstructfield.Name = value.GetValueString()
	case "GongStruct":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			sliceofpointertogongstructfield.GongStruct = nil
			for __instance__ := range stage.GongStructs {
				if stage.GongStructMap_Staged_Order[__instance__] == uint(id) {
					sliceofpointertogongstructfield.GongStruct = __instance__
					break
				}
			}
		}
	case "Index":
		sliceofpointertogongstructfield.Index = int(value.GetValueInt())
	case "CompositeStructName":
		sliceofpointertogongstructfield.CompositeStructName = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func SetFieldStringValueFromPointer(instance GongstructIF, fieldName string, value GongFieldValue, stage *Stage) error {
	return instance.GongSetFieldValue(fieldName, value, stage)
}

// insertion point for generic get gongstruct name
func (gongbasicfield *GongBasicField) GongGetGongstructName() string {
	return "GongBasicField"
}

func (gongenum *GongEnum) GongGetGongstructName() string {
	return "GongEnum"
}

func (gongenumvalue *GongEnumValue) GongGetGongstructName() string {
	return "GongEnumValue"
}

func (gonglink *GongLink) GongGetGongstructName() string {
	return "GongLink"
}

func (gongnote *GongNote) GongGetGongstructName() string {
	return "GongNote"
}

func (gongstruct *GongStruct) GongGetGongstructName() string {
	return "GongStruct"
}

func (gongtimefield *GongTimeField) GongGetGongstructName() string {
	return "GongTimeField"
}

func (metareference *MetaReference) GongGetGongstructName() string {
	return "MetaReference"
}

func (modelpkg *ModelPkg) GongGetGongstructName() string {
	return "ModelPkg"
}

func (pointertogongstructfield *PointerToGongStructField) GongGetGongstructName() string {
	return "PointerToGongStructField"
}

func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) GongGetGongstructName() string {
	return "SliceOfPointerToGongStructField"
}

func GetGongstructNameFromPointer(instance GongstructIF) (res string) {
	res = instance.GongGetGongstructName()
	return
}

func (stage *Stage) ResetMapStrings() {

	// insertion point for generic get gongstruct name
	stage.GongBasicFields_mapString = make(map[string]*GongBasicField)
	for gongbasicfield := range stage.GongBasicFields {
		stage.GongBasicFields_mapString[gongbasicfield.Name] = gongbasicfield
	}

	stage.GongEnums_mapString = make(map[string]*GongEnum)
	for gongenum := range stage.GongEnums {
		stage.GongEnums_mapString[gongenum.Name] = gongenum
	}

	stage.GongEnumValues_mapString = make(map[string]*GongEnumValue)
	for gongenumvalue := range stage.GongEnumValues {
		stage.GongEnumValues_mapString[gongenumvalue.Name] = gongenumvalue
	}

	stage.GongLinks_mapString = make(map[string]*GongLink)
	for gonglink := range stage.GongLinks {
		stage.GongLinks_mapString[gonglink.Name] = gonglink
	}

	stage.GongNotes_mapString = make(map[string]*GongNote)
	for gongnote := range stage.GongNotes {
		stage.GongNotes_mapString[gongnote.Name] = gongnote
	}

	stage.GongStructs_mapString = make(map[string]*GongStruct)
	for gongstruct := range stage.GongStructs {
		stage.GongStructs_mapString[gongstruct.Name] = gongstruct
	}

	stage.GongTimeFields_mapString = make(map[string]*GongTimeField)
	for gongtimefield := range stage.GongTimeFields {
		stage.GongTimeFields_mapString[gongtimefield.Name] = gongtimefield
	}

	stage.MetaReferences_mapString = make(map[string]*MetaReference)
	for metareference := range stage.MetaReferences {
		stage.MetaReferences_mapString[metareference.Name] = metareference
	}

	stage.ModelPkgs_mapString = make(map[string]*ModelPkg)
	for modelpkg := range stage.ModelPkgs {
		stage.ModelPkgs_mapString[modelpkg.Name] = modelpkg
	}

	stage.PointerToGongStructFields_mapString = make(map[string]*PointerToGongStructField)
	for pointertogongstructfield := range stage.PointerToGongStructFields {
		stage.PointerToGongStructFields_mapString[pointertogongstructfield.Name] = pointertogongstructfield
	}

	stage.SliceOfPointerToGongStructFields_mapString = make(map[string]*SliceOfPointerToGongStructField)
	for sliceofpointertogongstructfield := range stage.SliceOfPointerToGongStructFields {
		stage.SliceOfPointerToGongStructFields_mapString[sliceofpointertogongstructfield.Name] = sliceofpointertogongstructfield
	}

}

// Last line of the template
