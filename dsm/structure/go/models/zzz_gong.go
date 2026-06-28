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

	structure_go "github.com/fullstack-lang/gong/dsm/structure/go"
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
	AllocatedResourceShapes                map[*AllocatedResourceShape]struct{}
	AllocatedResourceShapes_instance       map[*AllocatedResourceShape]*AllocatedResourceShape
	AllocatedResourceShapes_mapString      map[string]*AllocatedResourceShape
	AllocatedResourceShapeOrder            uint
	AllocatedResourceShape_stagedOrder     map[*AllocatedResourceShape]uint
	AllocatedResourceShape_orderStaged     map[uint]*AllocatedResourceShape
	AllocatedResourceShapes_reference      map[*AllocatedResourceShape]*AllocatedResourceShape
	AllocatedResourceShapes_referenceOrder map[*AllocatedResourceShape]uint

	// insertion point for slice of pointers maps
	OnAfterAllocatedResourceShapeCreateCallback OnAfterCreateInterface[AllocatedResourceShape]
	OnAfterAllocatedResourceShapeUpdateCallback OnAfterUpdateInterface[AllocatedResourceShape]
	OnAfterAllocatedResourceShapeDeleteCallback OnAfterDeleteInterface[AllocatedResourceShape]
	OnAfterAllocatedResourceShapeReadCallback   OnAfterReadInterface[AllocatedResourceShape]

	AllocatedSystemShapes                map[*AllocatedSystemShape]struct{}
	AllocatedSystemShapes_instance       map[*AllocatedSystemShape]*AllocatedSystemShape
	AllocatedSystemShapes_mapString      map[string]*AllocatedSystemShape
	AllocatedSystemShapeOrder            uint
	AllocatedSystemShape_stagedOrder     map[*AllocatedSystemShape]uint
	AllocatedSystemShape_orderStaged     map[uint]*AllocatedSystemShape
	AllocatedSystemShapes_reference      map[*AllocatedSystemShape]*AllocatedSystemShape
	AllocatedSystemShapes_referenceOrder map[*AllocatedSystemShape]uint

	// insertion point for slice of pointers maps
	OnAfterAllocatedSystemShapeCreateCallback OnAfterCreateInterface[AllocatedSystemShape]
	OnAfterAllocatedSystemShapeUpdateCallback OnAfterUpdateInterface[AllocatedSystemShape]
	OnAfterAllocatedSystemShapeDeleteCallback OnAfterDeleteInterface[AllocatedSystemShape]
	OnAfterAllocatedSystemShapeReadCallback   OnAfterReadInterface[AllocatedSystemShape]

	ControlFlows                map[*ControlFlow]struct{}
	ControlFlows_instance       map[*ControlFlow]*ControlFlow
	ControlFlows_mapString      map[string]*ControlFlow
	ControlFlowOrder            uint
	ControlFlow_stagedOrder     map[*ControlFlow]uint
	ControlFlow_orderStaged     map[uint]*ControlFlow
	ControlFlows_reference      map[*ControlFlow]*ControlFlow
	ControlFlows_referenceOrder map[*ControlFlow]uint

	// insertion point for slice of pointers maps
	OnAfterControlFlowCreateCallback OnAfterCreateInterface[ControlFlow]
	OnAfterControlFlowUpdateCallback OnAfterUpdateInterface[ControlFlow]
	OnAfterControlFlowDeleteCallback OnAfterDeleteInterface[ControlFlow]
	OnAfterControlFlowReadCallback   OnAfterReadInterface[ControlFlow]

	ControlFlowShapes                map[*ControlFlowShape]struct{}
	ControlFlowShapes_instance       map[*ControlFlowShape]*ControlFlowShape
	ControlFlowShapes_mapString      map[string]*ControlFlowShape
	ControlFlowShapeOrder            uint
	ControlFlowShape_stagedOrder     map[*ControlFlowShape]uint
	ControlFlowShape_orderStaged     map[uint]*ControlFlowShape
	ControlFlowShapes_reference      map[*ControlFlowShape]*ControlFlowShape
	ControlFlowShapes_referenceOrder map[*ControlFlowShape]uint

	// insertion point for slice of pointers maps
	OnAfterControlFlowShapeCreateCallback OnAfterCreateInterface[ControlFlowShape]
	OnAfterControlFlowShapeUpdateCallback OnAfterUpdateInterface[ControlFlowShape]
	OnAfterControlFlowShapeDeleteCallback OnAfterDeleteInterface[ControlFlowShape]
	OnAfterControlFlowShapeReadCallback   OnAfterReadInterface[ControlFlowShape]

	Datas                map[*Data]struct{}
	Datas_instance       map[*Data]*Data
	Datas_mapString      map[string]*Data
	DataOrder            uint
	Data_stagedOrder     map[*Data]uint
	Data_orderStaged     map[uint]*Data
	Datas_reference      map[*Data]*Data
	Datas_referenceOrder map[*Data]uint

	// insertion point for slice of pointers maps
	OnAfterDataCreateCallback OnAfterCreateInterface[Data]
	OnAfterDataUpdateCallback OnAfterUpdateInterface[Data]
	OnAfterDataDeleteCallback OnAfterDeleteInterface[Data]
	OnAfterDataReadCallback   OnAfterReadInterface[Data]

	DataFlows                map[*DataFlow]struct{}
	DataFlows_instance       map[*DataFlow]*DataFlow
	DataFlows_mapString      map[string]*DataFlow
	DataFlowOrder            uint
	DataFlow_stagedOrder     map[*DataFlow]uint
	DataFlow_orderStaged     map[uint]*DataFlow
	DataFlows_reference      map[*DataFlow]*DataFlow
	DataFlows_referenceOrder map[*DataFlow]uint

	// insertion point for slice of pointers maps
	DataFlow_Datas_reverseMap map[*Data]*DataFlow

	OnAfterDataFlowCreateCallback OnAfterCreateInterface[DataFlow]
	OnAfterDataFlowUpdateCallback OnAfterUpdateInterface[DataFlow]
	OnAfterDataFlowDeleteCallback OnAfterDeleteInterface[DataFlow]
	OnAfterDataFlowReadCallback   OnAfterReadInterface[DataFlow]

	DataFlowShapes                map[*DataFlowShape]struct{}
	DataFlowShapes_instance       map[*DataFlowShape]*DataFlowShape
	DataFlowShapes_mapString      map[string]*DataFlowShape
	DataFlowShapeOrder            uint
	DataFlowShape_stagedOrder     map[*DataFlowShape]uint
	DataFlowShape_orderStaged     map[uint]*DataFlowShape
	DataFlowShapes_reference      map[*DataFlowShape]*DataFlowShape
	DataFlowShapes_referenceOrder map[*DataFlowShape]uint

	// insertion point for slice of pointers maps
	OnAfterDataFlowShapeCreateCallback OnAfterCreateInterface[DataFlowShape]
	OnAfterDataFlowShapeUpdateCallback OnAfterUpdateInterface[DataFlowShape]
	OnAfterDataFlowShapeDeleteCallback OnAfterDeleteInterface[DataFlowShape]
	OnAfterDataFlowShapeReadCallback   OnAfterReadInterface[DataFlowShape]

	DataShapes                map[*DataShape]struct{}
	DataShapes_instance       map[*DataShape]*DataShape
	DataShapes_mapString      map[string]*DataShape
	DataShapeOrder            uint
	DataShape_stagedOrder     map[*DataShape]uint
	DataShape_orderStaged     map[uint]*DataShape
	DataShapes_reference      map[*DataShape]*DataShape
	DataShapes_referenceOrder map[*DataShape]uint

	// insertion point for slice of pointers maps
	OnAfterDataShapeCreateCallback OnAfterCreateInterface[DataShape]
	OnAfterDataShapeUpdateCallback OnAfterUpdateInterface[DataShape]
	OnAfterDataShapeDeleteCallback OnAfterDeleteInterface[DataShape]
	OnAfterDataShapeReadCallback   OnAfterReadInterface[DataShape]

	DiagramStructures                map[*DiagramStructure]struct{}
	DiagramStructures_instance       map[*DiagramStructure]*DiagramStructure
	DiagramStructures_mapString      map[string]*DiagramStructure
	DiagramStructureOrder            uint
	DiagramStructure_stagedOrder     map[*DiagramStructure]uint
	DiagramStructure_orderStaged     map[uint]*DiagramStructure
	DiagramStructures_reference      map[*DiagramStructure]*DiagramStructure
	DiagramStructures_referenceOrder map[*DiagramStructure]uint

	// insertion point for slice of pointers maps
	DiagramStructure_System_Shapes_reverseMap map[*SystemShape]*DiagramStructure

	DiagramStructure_SystemsWhoseNodeIsExpanded_reverseMap map[*System]*DiagramStructure

	DiagramStructure_Part_Shapes_reverseMap map[*PartShape]*DiagramStructure

	DiagramStructure_PartWhoseNodeIsExpanded_reverseMap map[*Part]*DiagramStructure

	DiagramStructure_ExternalPart_Shapes_reverseMap map[*ExternalPartShape]*DiagramStructure

	DiagramStructure_ExternalPartWhoseNodeIsExpanded_reverseMap map[*Part]*DiagramStructure

	DiagramStructure_ExternalPartsWhoseOutDataFlowsNodeIsExpanded_reverseMap map[*Part]*DiagramStructure

	DiagramStructure_ExternalPartsWhoseInDataFlowsNodeIsExpanded_reverseMap map[*Part]*DiagramStructure

	DiagramStructure_PortsWhoseNodeIsExpanded_reverseMap map[*Port]*DiagramStructure

	DiagramStructure_Port_Shapes_reverseMap map[*PortShape]*DiagramStructure

	DiagramStructure_ControlFlowsWhoseNodeIsExpanded_reverseMap map[*ControlFlow]*DiagramStructure

	DiagramStructure_ControlFlow_Shapes_reverseMap map[*ControlFlowShape]*DiagramStructure

	DiagramStructure_DataFlowsWhoseNodeIsExpanded_reverseMap map[*DataFlow]*DiagramStructure

	DiagramStructure_DataFlow_Shapes_reverseMap map[*DataFlowShape]*DiagramStructure

	DiagramStructure_DatasWhoseNodeIsExpanded_reverseMap map[*Data]*DiagramStructure

	DiagramStructure_Data_Shapes_reverseMap map[*DataShape]*DiagramStructure

	DiagramStructure_DataFlowsWhoseDataNodeIsExpanded_reverseMap map[*DataFlow]*DiagramStructure

	DiagramStructure_AllocatedResourcesWhoseNodeIsExpanded_reverseMap map[*Resource]*DiagramStructure

	DiagramStructure_AllocatedResourceShapes_reverseMap map[*AllocatedResourceShape]*DiagramStructure

	DiagramStructure_AllocatedSystemesWhoseNodeIsExpanded_reverseMap map[*System]*DiagramStructure

	DiagramStructure_AllocatedSystemShapes_reverseMap map[*AllocatedSystemShape]*DiagramStructure

	DiagramStructure_Note_Shapes_reverseMap map[*NoteShape]*DiagramStructure

	DiagramStructure_NotesWhoseNodeIsExpanded_reverseMap map[*Note]*DiagramStructure

	DiagramStructure_NotePortShapes_reverseMap map[*NotePortShape]*DiagramStructure

	OnAfterDiagramStructureCreateCallback OnAfterCreateInterface[DiagramStructure]
	OnAfterDiagramStructureUpdateCallback OnAfterUpdateInterface[DiagramStructure]
	OnAfterDiagramStructureDeleteCallback OnAfterDeleteInterface[DiagramStructure]
	OnAfterDiagramStructureReadCallback   OnAfterReadInterface[DiagramStructure]

	ExternalPartShapes                map[*ExternalPartShape]struct{}
	ExternalPartShapes_instance       map[*ExternalPartShape]*ExternalPartShape
	ExternalPartShapes_mapString      map[string]*ExternalPartShape
	ExternalPartShapeOrder            uint
	ExternalPartShape_stagedOrder     map[*ExternalPartShape]uint
	ExternalPartShape_orderStaged     map[uint]*ExternalPartShape
	ExternalPartShapes_reference      map[*ExternalPartShape]*ExternalPartShape
	ExternalPartShapes_referenceOrder map[*ExternalPartShape]uint

	// insertion point for slice of pointers maps
	OnAfterExternalPartShapeCreateCallback OnAfterCreateInterface[ExternalPartShape]
	OnAfterExternalPartShapeUpdateCallback OnAfterUpdateInterface[ExternalPartShape]
	OnAfterExternalPartShapeDeleteCallback OnAfterDeleteInterface[ExternalPartShape]
	OnAfterExternalPartShapeReadCallback   OnAfterReadInterface[ExternalPartShape]

	Librarys                map[*Library]struct{}
	Librarys_instance       map[*Library]*Library
	Librarys_mapString      map[string]*Library
	LibraryOrder            uint
	Library_stagedOrder     map[*Library]uint
	Library_orderStaged     map[uint]*Library
	Librarys_reference      map[*Library]*Library
	Librarys_referenceOrder map[*Library]uint

	// insertion point for slice of pointers maps
	Library_SubLibraries_reverseMap map[*Library]*Library

	Library_SubLibrariesWhoseNodeIsExpanded_reverseMap map[*Library]*Library

	Library_RootSystemes_reverseMap map[*System]*Library

	Library_SystemsWhoseNodeIsExpanded_reverseMap map[*System]*Library

	Library_RootDataFlows_reverseMap map[*DataFlow]*Library

	Library_DataFlowsWhoseNodeIsExpanded_reverseMap map[*DataFlow]*Library

	Library_RootDatas_reverseMap map[*Data]*Library

	Library_DatasWhoseNodeIsExpanded_reverseMap map[*Data]*Library

	Library_RootResources_reverseMap map[*Resource]*Library

	Library_ResourcesWhoseNodeIsExpanded_reverseMap map[*Resource]*Library

	Library_PartsWhoseNodeIsExpanded_reverseMap map[*Part]*Library

	Library_RootNotes_reverseMap map[*Note]*Library

	Library_NotesWhoseNodeIsExpanded_reverseMap map[*Note]*Library

	OnAfterLibraryCreateCallback OnAfterCreateInterface[Library]
	OnAfterLibraryUpdateCallback OnAfterUpdateInterface[Library]
	OnAfterLibraryDeleteCallback OnAfterDeleteInterface[Library]
	OnAfterLibraryReadCallback   OnAfterReadInterface[Library]

	Notes                map[*Note]struct{}
	Notes_instance       map[*Note]*Note
	Notes_mapString      map[string]*Note
	NoteOrder            uint
	Note_stagedOrder     map[*Note]uint
	Note_orderStaged     map[uint]*Note
	Notes_reference      map[*Note]*Note
	Notes_referenceOrder map[*Note]uint

	// insertion point for slice of pointers maps
	Note_Ports_reverseMap map[*Port]*Note

	OnAfterNoteCreateCallback OnAfterCreateInterface[Note]
	OnAfterNoteUpdateCallback OnAfterUpdateInterface[Note]
	OnAfterNoteDeleteCallback OnAfterDeleteInterface[Note]
	OnAfterNoteReadCallback   OnAfterReadInterface[Note]

	NotePortShapes                map[*NotePortShape]struct{}
	NotePortShapes_instance       map[*NotePortShape]*NotePortShape
	NotePortShapes_mapString      map[string]*NotePortShape
	NotePortShapeOrder            uint
	NotePortShape_stagedOrder     map[*NotePortShape]uint
	NotePortShape_orderStaged     map[uint]*NotePortShape
	NotePortShapes_reference      map[*NotePortShape]*NotePortShape
	NotePortShapes_referenceOrder map[*NotePortShape]uint

	// insertion point for slice of pointers maps
	OnAfterNotePortShapeCreateCallback OnAfterCreateInterface[NotePortShape]
	OnAfterNotePortShapeUpdateCallback OnAfterUpdateInterface[NotePortShape]
	OnAfterNotePortShapeDeleteCallback OnAfterDeleteInterface[NotePortShape]
	OnAfterNotePortShapeReadCallback   OnAfterReadInterface[NotePortShape]

	NoteShapes                map[*NoteShape]struct{}
	NoteShapes_instance       map[*NoteShape]*NoteShape
	NoteShapes_mapString      map[string]*NoteShape
	NoteShapeOrder            uint
	NoteShape_stagedOrder     map[*NoteShape]uint
	NoteShape_orderStaged     map[uint]*NoteShape
	NoteShapes_reference      map[*NoteShape]*NoteShape
	NoteShapes_referenceOrder map[*NoteShape]uint

	// insertion point for slice of pointers maps
	OnAfterNoteShapeCreateCallback OnAfterCreateInterface[NoteShape]
	OnAfterNoteShapeUpdateCallback OnAfterUpdateInterface[NoteShape]
	OnAfterNoteShapeDeleteCallback OnAfterDeleteInterface[NoteShape]
	OnAfterNoteShapeReadCallback   OnAfterReadInterface[NoteShape]

	Parts                map[*Part]struct{}
	Parts_instance       map[*Part]*Part
	Parts_mapString      map[string]*Part
	PartOrder            uint
	Part_stagedOrder     map[*Part]uint
	Part_orderStaged     map[uint]*Part
	Parts_reference      map[*Part]*Part
	Parts_referenceOrder map[*Part]uint

	// insertion point for slice of pointers maps
	Part_Ports_reverseMap map[*Port]*Part

	Part_ControlFlows_reverseMap map[*ControlFlow]*Part

	Part_PortWhoseOutControlFlowsNodeIsExpanded_reverseMap map[*Port]*Part

	Part_PortWhoseInControlFlowsNodeIsExpanded_reverseMap map[*Port]*Part

	Part_PortWhoseOutDataFlowsNodeIsExpanded_reverseMap map[*Port]*Part

	Part_PortWhoseInDataFlowsNodeIsExpanded_reverseMap map[*Port]*Part

	Part_PartAnchoredPath_reverseMap map[*PartAnchoredPath]*Part

	OnAfterPartCreateCallback OnAfterCreateInterface[Part]
	OnAfterPartUpdateCallback OnAfterUpdateInterface[Part]
	OnAfterPartDeleteCallback OnAfterDeleteInterface[Part]
	OnAfterPartReadCallback   OnAfterReadInterface[Part]

	PartAnchoredPaths                map[*PartAnchoredPath]struct{}
	PartAnchoredPaths_instance       map[*PartAnchoredPath]*PartAnchoredPath
	PartAnchoredPaths_mapString      map[string]*PartAnchoredPath
	PartAnchoredPathOrder            uint
	PartAnchoredPath_stagedOrder     map[*PartAnchoredPath]uint
	PartAnchoredPath_orderStaged     map[uint]*PartAnchoredPath
	PartAnchoredPaths_reference      map[*PartAnchoredPath]*PartAnchoredPath
	PartAnchoredPaths_referenceOrder map[*PartAnchoredPath]uint

	// insertion point for slice of pointers maps
	OnAfterPartAnchoredPathCreateCallback OnAfterCreateInterface[PartAnchoredPath]
	OnAfterPartAnchoredPathUpdateCallback OnAfterUpdateInterface[PartAnchoredPath]
	OnAfterPartAnchoredPathDeleteCallback OnAfterDeleteInterface[PartAnchoredPath]
	OnAfterPartAnchoredPathReadCallback   OnAfterReadInterface[PartAnchoredPath]

	PartShapes                map[*PartShape]struct{}
	PartShapes_instance       map[*PartShape]*PartShape
	PartShapes_mapString      map[string]*PartShape
	PartShapeOrder            uint
	PartShape_stagedOrder     map[*PartShape]uint
	PartShape_orderStaged     map[uint]*PartShape
	PartShapes_reference      map[*PartShape]*PartShape
	PartShapes_referenceOrder map[*PartShape]uint

	// insertion point for slice of pointers maps
	OnAfterPartShapeCreateCallback OnAfterCreateInterface[PartShape]
	OnAfterPartShapeUpdateCallback OnAfterUpdateInterface[PartShape]
	OnAfterPartShapeDeleteCallback OnAfterDeleteInterface[PartShape]
	OnAfterPartShapeReadCallback   OnAfterReadInterface[PartShape]

	Ports                map[*Port]struct{}
	Ports_instance       map[*Port]*Port
	Ports_mapString      map[string]*Port
	PortOrder            uint
	Port_stagedOrder     map[*Port]uint
	Port_orderStaged     map[uint]*Port
	Ports_reference      map[*Port]*Port
	Ports_referenceOrder map[*Port]uint

	// insertion point for slice of pointers maps
	OnAfterPortCreateCallback OnAfterCreateInterface[Port]
	OnAfterPortUpdateCallback OnAfterUpdateInterface[Port]
	OnAfterPortDeleteCallback OnAfterDeleteInterface[Port]
	OnAfterPortReadCallback   OnAfterReadInterface[Port]

	PortShapes                map[*PortShape]struct{}
	PortShapes_instance       map[*PortShape]*PortShape
	PortShapes_mapString      map[string]*PortShape
	PortShapeOrder            uint
	PortShape_stagedOrder     map[*PortShape]uint
	PortShape_orderStaged     map[uint]*PortShape
	PortShapes_reference      map[*PortShape]*PortShape
	PortShapes_referenceOrder map[*PortShape]uint

	// insertion point for slice of pointers maps
	OnAfterPortShapeCreateCallback OnAfterCreateInterface[PortShape]
	OnAfterPortShapeUpdateCallback OnAfterUpdateInterface[PortShape]
	OnAfterPortShapeDeleteCallback OnAfterDeleteInterface[PortShape]
	OnAfterPortShapeReadCallback   OnAfterReadInterface[PortShape]

	Resources                map[*Resource]struct{}
	Resources_instance       map[*Resource]*Resource
	Resources_mapString      map[string]*Resource
	ResourceOrder            uint
	Resource_stagedOrder     map[*Resource]uint
	Resource_orderStaged     map[uint]*Resource
	Resources_reference      map[*Resource]*Resource
	Resources_referenceOrder map[*Resource]uint

	// insertion point for slice of pointers maps
	OnAfterResourceCreateCallback OnAfterCreateInterface[Resource]
	OnAfterResourceUpdateCallback OnAfterUpdateInterface[Resource]
	OnAfterResourceDeleteCallback OnAfterDeleteInterface[Resource]
	OnAfterResourceReadCallback   OnAfterReadInterface[Resource]

	Systems                map[*System]struct{}
	Systems_instance       map[*System]*System
	Systems_mapString      map[string]*System
	SystemOrder            uint
	System_stagedOrder     map[*System]uint
	System_orderStaged     map[uint]*System
	Systems_reference      map[*System]*System
	Systems_referenceOrder map[*System]uint

	// insertion point for slice of pointers maps
	System_DiagramStructures_reverseMap map[*DiagramStructure]*System

	System_DiagramStructureWhoseNodeIsExpanded_reverseMap map[*DiagramStructure]*System

	System_SubSystemes_reverseMap map[*System]*System

	System_Parts_reverseMap map[*Part]*System

	System_PartWhoseNodeIsExpanded_reverseMap map[*Part]*System

	System_DataFlows_reverseMap map[*DataFlow]*System

	System_ExternalParts_reverseMap map[*Part]*System

	System_ExternalPartWhoseNodeIsExpanded_reverseMap map[*Part]*System

	OnAfterSystemCreateCallback OnAfterCreateInterface[System]
	OnAfterSystemUpdateCallback OnAfterUpdateInterface[System]
	OnAfterSystemDeleteCallback OnAfterDeleteInterface[System]
	OnAfterSystemReadCallback   OnAfterReadInterface[System]

	SystemShapes                map[*SystemShape]struct{}
	SystemShapes_instance       map[*SystemShape]*SystemShape
	SystemShapes_mapString      map[string]*SystemShape
	SystemShapeOrder            uint
	SystemShape_stagedOrder     map[*SystemShape]uint
	SystemShape_orderStaged     map[uint]*SystemShape
	SystemShapes_reference      map[*SystemShape]*SystemShape
	SystemShapes_referenceOrder map[*SystemShape]uint

	// insertion point for slice of pointers maps
	OnAfterSystemShapeCreateCallback OnAfterCreateInterface[SystemShape]
	OnAfterSystemShapeUpdateCallback OnAfterUpdateInterface[SystemShape]
	OnAfterSystemShapeDeleteCallback OnAfterDeleteInterface[SystemShape]
	OnAfterSystemShapeReadCallback   OnAfterReadInterface[SystemShape]

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
	stage.AllocatedResourceShapes_reference = make(map[*AllocatedResourceShape]*AllocatedResourceShape)
	stage.AllocatedResourceShapes_instance = make(map[*AllocatedResourceShape]*AllocatedResourceShape)
	stage.AllocatedResourceShapes_referenceOrder = make(map[*AllocatedResourceShape]uint)

	stage.AllocatedSystemShapes_reference = make(map[*AllocatedSystemShape]*AllocatedSystemShape)
	stage.AllocatedSystemShapes_instance = make(map[*AllocatedSystemShape]*AllocatedSystemShape)
	stage.AllocatedSystemShapes_referenceOrder = make(map[*AllocatedSystemShape]uint)

	stage.ControlFlows_reference = make(map[*ControlFlow]*ControlFlow)
	stage.ControlFlows_instance = make(map[*ControlFlow]*ControlFlow)
	stage.ControlFlows_referenceOrder = make(map[*ControlFlow]uint)

	stage.ControlFlowShapes_reference = make(map[*ControlFlowShape]*ControlFlowShape)
	stage.ControlFlowShapes_instance = make(map[*ControlFlowShape]*ControlFlowShape)
	stage.ControlFlowShapes_referenceOrder = make(map[*ControlFlowShape]uint)

	stage.Datas_reference = make(map[*Data]*Data)
	stage.Datas_instance = make(map[*Data]*Data)
	stage.Datas_referenceOrder = make(map[*Data]uint)

	stage.DataFlows_reference = make(map[*DataFlow]*DataFlow)
	stage.DataFlows_instance = make(map[*DataFlow]*DataFlow)
	stage.DataFlows_referenceOrder = make(map[*DataFlow]uint)

	stage.DataFlowShapes_reference = make(map[*DataFlowShape]*DataFlowShape)
	stage.DataFlowShapes_instance = make(map[*DataFlowShape]*DataFlowShape)
	stage.DataFlowShapes_referenceOrder = make(map[*DataFlowShape]uint)

	stage.DataShapes_reference = make(map[*DataShape]*DataShape)
	stage.DataShapes_instance = make(map[*DataShape]*DataShape)
	stage.DataShapes_referenceOrder = make(map[*DataShape]uint)

	stage.DiagramStructures_reference = make(map[*DiagramStructure]*DiagramStructure)
	stage.DiagramStructures_instance = make(map[*DiagramStructure]*DiagramStructure)
	stage.DiagramStructures_referenceOrder = make(map[*DiagramStructure]uint)

	stage.ExternalPartShapes_reference = make(map[*ExternalPartShape]*ExternalPartShape)
	stage.ExternalPartShapes_instance = make(map[*ExternalPartShape]*ExternalPartShape)
	stage.ExternalPartShapes_referenceOrder = make(map[*ExternalPartShape]uint)

	stage.Librarys_reference = make(map[*Library]*Library)
	stage.Librarys_instance = make(map[*Library]*Library)
	stage.Librarys_referenceOrder = make(map[*Library]uint)

	stage.Notes_reference = make(map[*Note]*Note)
	stage.Notes_instance = make(map[*Note]*Note)
	stage.Notes_referenceOrder = make(map[*Note]uint)

	stage.NotePortShapes_reference = make(map[*NotePortShape]*NotePortShape)
	stage.NotePortShapes_instance = make(map[*NotePortShape]*NotePortShape)
	stage.NotePortShapes_referenceOrder = make(map[*NotePortShape]uint)

	stage.NoteShapes_reference = make(map[*NoteShape]*NoteShape)
	stage.NoteShapes_instance = make(map[*NoteShape]*NoteShape)
	stage.NoteShapes_referenceOrder = make(map[*NoteShape]uint)

	stage.Parts_reference = make(map[*Part]*Part)
	stage.Parts_instance = make(map[*Part]*Part)
	stage.Parts_referenceOrder = make(map[*Part]uint)

	stage.PartAnchoredPaths_reference = make(map[*PartAnchoredPath]*PartAnchoredPath)
	stage.PartAnchoredPaths_instance = make(map[*PartAnchoredPath]*PartAnchoredPath)
	stage.PartAnchoredPaths_referenceOrder = make(map[*PartAnchoredPath]uint)

	stage.PartShapes_reference = make(map[*PartShape]*PartShape)
	stage.PartShapes_instance = make(map[*PartShape]*PartShape)
	stage.PartShapes_referenceOrder = make(map[*PartShape]uint)

	stage.Ports_reference = make(map[*Port]*Port)
	stage.Ports_instance = make(map[*Port]*Port)
	stage.Ports_referenceOrder = make(map[*Port]uint)

	stage.PortShapes_reference = make(map[*PortShape]*PortShape)
	stage.PortShapes_instance = make(map[*PortShape]*PortShape)
	stage.PortShapes_referenceOrder = make(map[*PortShape]uint)

	stage.Resources_reference = make(map[*Resource]*Resource)
	stage.Resources_instance = make(map[*Resource]*Resource)
	stage.Resources_referenceOrder = make(map[*Resource]uint)

	stage.Systems_reference = make(map[*System]*System)
	stage.Systems_instance = make(map[*System]*System)
	stage.Systems_referenceOrder = make(map[*System]uint)

	stage.SystemShapes_reference = make(map[*SystemShape]*SystemShape)
	stage.SystemShapes_instance = make(map[*SystemShape]*SystemShape)
	stage.SystemShapes_referenceOrder = make(map[*SystemShape]uint)

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
	var maxAllocatedResourceShapeOrder uint
	var foundAllocatedResourceShape bool
	for _, order := range stage.AllocatedResourceShape_stagedOrder {
		if !foundAllocatedResourceShape || order > maxAllocatedResourceShapeOrder {
			maxAllocatedResourceShapeOrder = order
			foundAllocatedResourceShape = true
		}
	}
	if foundAllocatedResourceShape {
		stage.AllocatedResourceShapeOrder = maxAllocatedResourceShapeOrder + 1
	} else {
		stage.AllocatedResourceShapeOrder = 0
	}

	var maxAllocatedSystemShapeOrder uint
	var foundAllocatedSystemShape bool
	for _, order := range stage.AllocatedSystemShape_stagedOrder {
		if !foundAllocatedSystemShape || order > maxAllocatedSystemShapeOrder {
			maxAllocatedSystemShapeOrder = order
			foundAllocatedSystemShape = true
		}
	}
	if foundAllocatedSystemShape {
		stage.AllocatedSystemShapeOrder = maxAllocatedSystemShapeOrder + 1
	} else {
		stage.AllocatedSystemShapeOrder = 0
	}

	var maxControlFlowOrder uint
	var foundControlFlow bool
	for _, order := range stage.ControlFlow_stagedOrder {
		if !foundControlFlow || order > maxControlFlowOrder {
			maxControlFlowOrder = order
			foundControlFlow = true
		}
	}
	if foundControlFlow {
		stage.ControlFlowOrder = maxControlFlowOrder + 1
	} else {
		stage.ControlFlowOrder = 0
	}

	var maxControlFlowShapeOrder uint
	var foundControlFlowShape bool
	for _, order := range stage.ControlFlowShape_stagedOrder {
		if !foundControlFlowShape || order > maxControlFlowShapeOrder {
			maxControlFlowShapeOrder = order
			foundControlFlowShape = true
		}
	}
	if foundControlFlowShape {
		stage.ControlFlowShapeOrder = maxControlFlowShapeOrder + 1
	} else {
		stage.ControlFlowShapeOrder = 0
	}

	var maxDataOrder uint
	var foundData bool
	for _, order := range stage.Data_stagedOrder {
		if !foundData || order > maxDataOrder {
			maxDataOrder = order
			foundData = true
		}
	}
	if foundData {
		stage.DataOrder = maxDataOrder + 1
	} else {
		stage.DataOrder = 0
	}

	var maxDataFlowOrder uint
	var foundDataFlow bool
	for _, order := range stage.DataFlow_stagedOrder {
		if !foundDataFlow || order > maxDataFlowOrder {
			maxDataFlowOrder = order
			foundDataFlow = true
		}
	}
	if foundDataFlow {
		stage.DataFlowOrder = maxDataFlowOrder + 1
	} else {
		stage.DataFlowOrder = 0
	}

	var maxDataFlowShapeOrder uint
	var foundDataFlowShape bool
	for _, order := range stage.DataFlowShape_stagedOrder {
		if !foundDataFlowShape || order > maxDataFlowShapeOrder {
			maxDataFlowShapeOrder = order
			foundDataFlowShape = true
		}
	}
	if foundDataFlowShape {
		stage.DataFlowShapeOrder = maxDataFlowShapeOrder + 1
	} else {
		stage.DataFlowShapeOrder = 0
	}

	var maxDataShapeOrder uint
	var foundDataShape bool
	for _, order := range stage.DataShape_stagedOrder {
		if !foundDataShape || order > maxDataShapeOrder {
			maxDataShapeOrder = order
			foundDataShape = true
		}
	}
	if foundDataShape {
		stage.DataShapeOrder = maxDataShapeOrder + 1
	} else {
		stage.DataShapeOrder = 0
	}

	var maxDiagramStructureOrder uint
	var foundDiagramStructure bool
	for _, order := range stage.DiagramStructure_stagedOrder {
		if !foundDiagramStructure || order > maxDiagramStructureOrder {
			maxDiagramStructureOrder = order
			foundDiagramStructure = true
		}
	}
	if foundDiagramStructure {
		stage.DiagramStructureOrder = maxDiagramStructureOrder + 1
	} else {
		stage.DiagramStructureOrder = 0
	}

	var maxExternalPartShapeOrder uint
	var foundExternalPartShape bool
	for _, order := range stage.ExternalPartShape_stagedOrder {
		if !foundExternalPartShape || order > maxExternalPartShapeOrder {
			maxExternalPartShapeOrder = order
			foundExternalPartShape = true
		}
	}
	if foundExternalPartShape {
		stage.ExternalPartShapeOrder = maxExternalPartShapeOrder + 1
	} else {
		stage.ExternalPartShapeOrder = 0
	}

	var maxLibraryOrder uint
	var foundLibrary bool
	for _, order := range stage.Library_stagedOrder {
		if !foundLibrary || order > maxLibraryOrder {
			maxLibraryOrder = order
			foundLibrary = true
		}
	}
	if foundLibrary {
		stage.LibraryOrder = maxLibraryOrder + 1
	} else {
		stage.LibraryOrder = 0
	}

	var maxNoteOrder uint
	var foundNote bool
	for _, order := range stage.Note_stagedOrder {
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

	var maxNotePortShapeOrder uint
	var foundNotePortShape bool
	for _, order := range stage.NotePortShape_stagedOrder {
		if !foundNotePortShape || order > maxNotePortShapeOrder {
			maxNotePortShapeOrder = order
			foundNotePortShape = true
		}
	}
	if foundNotePortShape {
		stage.NotePortShapeOrder = maxNotePortShapeOrder + 1
	} else {
		stage.NotePortShapeOrder = 0
	}

	var maxNoteShapeOrder uint
	var foundNoteShape bool
	for _, order := range stage.NoteShape_stagedOrder {
		if !foundNoteShape || order > maxNoteShapeOrder {
			maxNoteShapeOrder = order
			foundNoteShape = true
		}
	}
	if foundNoteShape {
		stage.NoteShapeOrder = maxNoteShapeOrder + 1
	} else {
		stage.NoteShapeOrder = 0
	}

	var maxPartOrder uint
	var foundPart bool
	for _, order := range stage.Part_stagedOrder {
		if !foundPart || order > maxPartOrder {
			maxPartOrder = order
			foundPart = true
		}
	}
	if foundPart {
		stage.PartOrder = maxPartOrder + 1
	} else {
		stage.PartOrder = 0
	}

	var maxPartAnchoredPathOrder uint
	var foundPartAnchoredPath bool
	for _, order := range stage.PartAnchoredPath_stagedOrder {
		if !foundPartAnchoredPath || order > maxPartAnchoredPathOrder {
			maxPartAnchoredPathOrder = order
			foundPartAnchoredPath = true
		}
	}
	if foundPartAnchoredPath {
		stage.PartAnchoredPathOrder = maxPartAnchoredPathOrder + 1
	} else {
		stage.PartAnchoredPathOrder = 0
	}

	var maxPartShapeOrder uint
	var foundPartShape bool
	for _, order := range stage.PartShape_stagedOrder {
		if !foundPartShape || order > maxPartShapeOrder {
			maxPartShapeOrder = order
			foundPartShape = true
		}
	}
	if foundPartShape {
		stage.PartShapeOrder = maxPartShapeOrder + 1
	} else {
		stage.PartShapeOrder = 0
	}

	var maxPortOrder uint
	var foundPort bool
	for _, order := range stage.Port_stagedOrder {
		if !foundPort || order > maxPortOrder {
			maxPortOrder = order
			foundPort = true
		}
	}
	if foundPort {
		stage.PortOrder = maxPortOrder + 1
	} else {
		stage.PortOrder = 0
	}

	var maxPortShapeOrder uint
	var foundPortShape bool
	for _, order := range stage.PortShape_stagedOrder {
		if !foundPortShape || order > maxPortShapeOrder {
			maxPortShapeOrder = order
			foundPortShape = true
		}
	}
	if foundPortShape {
		stage.PortShapeOrder = maxPortShapeOrder + 1
	} else {
		stage.PortShapeOrder = 0
	}

	var maxResourceOrder uint
	var foundResource bool
	for _, order := range stage.Resource_stagedOrder {
		if !foundResource || order > maxResourceOrder {
			maxResourceOrder = order
			foundResource = true
		}
	}
	if foundResource {
		stage.ResourceOrder = maxResourceOrder + 1
	} else {
		stage.ResourceOrder = 0
	}

	var maxSystemOrder uint
	var foundSystem bool
	for _, order := range stage.System_stagedOrder {
		if !foundSystem || order > maxSystemOrder {
			maxSystemOrder = order
			foundSystem = true
		}
	}
	if foundSystem {
		stage.SystemOrder = maxSystemOrder + 1
	} else {
		stage.SystemOrder = 0
	}

	var maxSystemShapeOrder uint
	var foundSystemShape bool
	for _, order := range stage.SystemShape_stagedOrder {
		if !foundSystemShape || order > maxSystemShapeOrder {
			maxSystemShapeOrder = order
			foundSystemShape = true
		}
	}
	if foundSystemShape {
		stage.SystemShapeOrder = maxSystemShapeOrder + 1
	} else {
		stage.SystemShapeOrder = 0
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
	case *AllocatedResourceShape:
		tmp := GetStructInstancesByOrder(stage.AllocatedResourceShapes, stage.AllocatedResourceShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *AllocatedResourceShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *AllocatedSystemShape:
		tmp := GetStructInstancesByOrder(stage.AllocatedSystemShapes, stage.AllocatedSystemShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *AllocatedSystemShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *ControlFlow:
		tmp := GetStructInstancesByOrder(stage.ControlFlows, stage.ControlFlow_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *ControlFlow implements.
			res = append(res, any(v).(T))
		}
		return res
	case *ControlFlowShape:
		tmp := GetStructInstancesByOrder(stage.ControlFlowShapes, stage.ControlFlowShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *ControlFlowShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Data:
		tmp := GetStructInstancesByOrder(stage.Datas, stage.Data_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Data implements.
			res = append(res, any(v).(T))
		}
		return res
	case *DataFlow:
		tmp := GetStructInstancesByOrder(stage.DataFlows, stage.DataFlow_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *DataFlow implements.
			res = append(res, any(v).(T))
		}
		return res
	case *DataFlowShape:
		tmp := GetStructInstancesByOrder(stage.DataFlowShapes, stage.DataFlowShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *DataFlowShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *DataShape:
		tmp := GetStructInstancesByOrder(stage.DataShapes, stage.DataShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *DataShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *DiagramStructure:
		tmp := GetStructInstancesByOrder(stage.DiagramStructures, stage.DiagramStructure_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *DiagramStructure implements.
			res = append(res, any(v).(T))
		}
		return res
	case *ExternalPartShape:
		tmp := GetStructInstancesByOrder(stage.ExternalPartShapes, stage.ExternalPartShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *ExternalPartShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Library:
		tmp := GetStructInstancesByOrder(stage.Librarys, stage.Library_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Library implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Note:
		tmp := GetStructInstancesByOrder(stage.Notes, stage.Note_stagedOrder)

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
	case *NotePortShape:
		tmp := GetStructInstancesByOrder(stage.NotePortShapes, stage.NotePortShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *NotePortShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *NoteShape:
		tmp := GetStructInstancesByOrder(stage.NoteShapes, stage.NoteShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *NoteShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Part:
		tmp := GetStructInstancesByOrder(stage.Parts, stage.Part_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Part implements.
			res = append(res, any(v).(T))
		}
		return res
	case *PartAnchoredPath:
		tmp := GetStructInstancesByOrder(stage.PartAnchoredPaths, stage.PartAnchoredPath_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *PartAnchoredPath implements.
			res = append(res, any(v).(T))
		}
		return res
	case *PartShape:
		tmp := GetStructInstancesByOrder(stage.PartShapes, stage.PartShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *PartShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Port:
		tmp := GetStructInstancesByOrder(stage.Ports, stage.Port_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Port implements.
			res = append(res, any(v).(T))
		}
		return res
	case *PortShape:
		tmp := GetStructInstancesByOrder(stage.PortShapes, stage.PortShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *PortShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Resource:
		tmp := GetStructInstancesByOrder(stage.Resources, stage.Resource_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Resource implements.
			res = append(res, any(v).(T))
		}
		return res
	case *System:
		tmp := GetStructInstancesByOrder(stage.Systems, stage.System_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *System implements.
			res = append(res, any(v).(T))
		}
		return res
	case *SystemShape:
		tmp := GetStructInstancesByOrder(stage.SystemShapes, stage.SystemShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *SystemShape implements.
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
	case "AllocatedResourceShape":
		res = GetNamedStructInstances(stage.AllocatedResourceShapes, stage.AllocatedResourceShape_stagedOrder)
	case "AllocatedSystemShape":
		res = GetNamedStructInstances(stage.AllocatedSystemShapes, stage.AllocatedSystemShape_stagedOrder)
	case "ControlFlow":
		res = GetNamedStructInstances(stage.ControlFlows, stage.ControlFlow_stagedOrder)
	case "ControlFlowShape":
		res = GetNamedStructInstances(stage.ControlFlowShapes, stage.ControlFlowShape_stagedOrder)
	case "Data":
		res = GetNamedStructInstances(stage.Datas, stage.Data_stagedOrder)
	case "DataFlow":
		res = GetNamedStructInstances(stage.DataFlows, stage.DataFlow_stagedOrder)
	case "DataFlowShape":
		res = GetNamedStructInstances(stage.DataFlowShapes, stage.DataFlowShape_stagedOrder)
	case "DataShape":
		res = GetNamedStructInstances(stage.DataShapes, stage.DataShape_stagedOrder)
	case "DiagramStructure":
		res = GetNamedStructInstances(stage.DiagramStructures, stage.DiagramStructure_stagedOrder)
	case "ExternalPartShape":
		res = GetNamedStructInstances(stage.ExternalPartShapes, stage.ExternalPartShape_stagedOrder)
	case "Library":
		res = GetNamedStructInstances(stage.Librarys, stage.Library_stagedOrder)
	case "Note":
		res = GetNamedStructInstances(stage.Notes, stage.Note_stagedOrder)
	case "NotePortShape":
		res = GetNamedStructInstances(stage.NotePortShapes, stage.NotePortShape_stagedOrder)
	case "NoteShape":
		res = GetNamedStructInstances(stage.NoteShapes, stage.NoteShape_stagedOrder)
	case "Part":
		res = GetNamedStructInstances(stage.Parts, stage.Part_stagedOrder)
	case "PartAnchoredPath":
		res = GetNamedStructInstances(stage.PartAnchoredPaths, stage.PartAnchoredPath_stagedOrder)
	case "PartShape":
		res = GetNamedStructInstances(stage.PartShapes, stage.PartShape_stagedOrder)
	case "Port":
		res = GetNamedStructInstances(stage.Ports, stage.Port_stagedOrder)
	case "PortShape":
		res = GetNamedStructInstances(stage.PortShapes, stage.PortShape_stagedOrder)
	case "Resource":
		res = GetNamedStructInstances(stage.Resources, stage.Resource_stagedOrder)
	case "System":
		res = GetNamedStructInstances(stage.Systems, stage.System_stagedOrder)
	case "SystemShape":
		res = GetNamedStructInstances(stage.SystemShapes, stage.SystemShape_stagedOrder)
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
	return "github.com/fullstack-lang/gong/dsm/structure/go/models"
}

func (stage *Stage) GetMap_GongStructName_InstancesNb() map[string]int {
	return stage.Map_GongStructName_InstancesNb
}

func (stage *Stage) GetModelsEmbededDir() embed.FS {
	return structure_go.GoModelsDir
}

func (stage *Stage) GetDigramsEmbededDir() embed.FS {
	return structure_go.GoDiagramsDir
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
	CommitAllocatedResourceShape(allocatedresourceshape *AllocatedResourceShape)
	CheckoutAllocatedResourceShape(allocatedresourceshape *AllocatedResourceShape)
	CommitAllocatedSystemShape(allocatedsystemshape *AllocatedSystemShape)
	CheckoutAllocatedSystemShape(allocatedsystemshape *AllocatedSystemShape)
	CommitControlFlow(controlflow *ControlFlow)
	CheckoutControlFlow(controlflow *ControlFlow)
	CommitControlFlowShape(controlflowshape *ControlFlowShape)
	CheckoutControlFlowShape(controlflowshape *ControlFlowShape)
	CommitData(data *Data)
	CheckoutData(data *Data)
	CommitDataFlow(dataflow *DataFlow)
	CheckoutDataFlow(dataflow *DataFlow)
	CommitDataFlowShape(dataflowshape *DataFlowShape)
	CheckoutDataFlowShape(dataflowshape *DataFlowShape)
	CommitDataShape(datashape *DataShape)
	CheckoutDataShape(datashape *DataShape)
	CommitDiagramStructure(diagramstructure *DiagramStructure)
	CheckoutDiagramStructure(diagramstructure *DiagramStructure)
	CommitExternalPartShape(externalpartshape *ExternalPartShape)
	CheckoutExternalPartShape(externalpartshape *ExternalPartShape)
	CommitLibrary(library *Library)
	CheckoutLibrary(library *Library)
	CommitNote(note *Note)
	CheckoutNote(note *Note)
	CommitNotePortShape(noteportshape *NotePortShape)
	CheckoutNotePortShape(noteportshape *NotePortShape)
	CommitNoteShape(noteshape *NoteShape)
	CheckoutNoteShape(noteshape *NoteShape)
	CommitPart(part *Part)
	CheckoutPart(part *Part)
	CommitPartAnchoredPath(partanchoredpath *PartAnchoredPath)
	CheckoutPartAnchoredPath(partanchoredpath *PartAnchoredPath)
	CommitPartShape(partshape *PartShape)
	CheckoutPartShape(partshape *PartShape)
	CommitPort(port *Port)
	CheckoutPort(port *Port)
	CommitPortShape(portshape *PortShape)
	CheckoutPortShape(portshape *PortShape)
	CommitResource(resource *Resource)
	CheckoutResource(resource *Resource)
	CommitSystem(system *System)
	CheckoutSystem(system *System)
	CommitSystemShape(systemshape *SystemShape)
	CheckoutSystemShape(systemshape *SystemShape)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(name string) (stage *Stage) {
	stage = &Stage{ // insertion point for array initiatialisation
		AllocatedResourceShapes:           make(map[*AllocatedResourceShape]struct{}),
		AllocatedResourceShapes_mapString: make(map[string]*AllocatedResourceShape),

		AllocatedSystemShapes:           make(map[*AllocatedSystemShape]struct{}),
		AllocatedSystemShapes_mapString: make(map[string]*AllocatedSystemShape),

		ControlFlows:           make(map[*ControlFlow]struct{}),
		ControlFlows_mapString: make(map[string]*ControlFlow),

		ControlFlowShapes:           make(map[*ControlFlowShape]struct{}),
		ControlFlowShapes_mapString: make(map[string]*ControlFlowShape),

		Datas:           make(map[*Data]struct{}),
		Datas_mapString: make(map[string]*Data),

		DataFlows:           make(map[*DataFlow]struct{}),
		DataFlows_mapString: make(map[string]*DataFlow),

		DataFlowShapes:           make(map[*DataFlowShape]struct{}),
		DataFlowShapes_mapString: make(map[string]*DataFlowShape),

		DataShapes:           make(map[*DataShape]struct{}),
		DataShapes_mapString: make(map[string]*DataShape),

		DiagramStructures:           make(map[*DiagramStructure]struct{}),
		DiagramStructures_mapString: make(map[string]*DiagramStructure),

		ExternalPartShapes:           make(map[*ExternalPartShape]struct{}),
		ExternalPartShapes_mapString: make(map[string]*ExternalPartShape),

		Librarys:           make(map[*Library]struct{}),
		Librarys_mapString: make(map[string]*Library),

		Notes:           make(map[*Note]struct{}),
		Notes_mapString: make(map[string]*Note),

		NotePortShapes:           make(map[*NotePortShape]struct{}),
		NotePortShapes_mapString: make(map[string]*NotePortShape),

		NoteShapes:           make(map[*NoteShape]struct{}),
		NoteShapes_mapString: make(map[string]*NoteShape),

		Parts:           make(map[*Part]struct{}),
		Parts_mapString: make(map[string]*Part),

		PartAnchoredPaths:           make(map[*PartAnchoredPath]struct{}),
		PartAnchoredPaths_mapString: make(map[string]*PartAnchoredPath),

		PartShapes:           make(map[*PartShape]struct{}),
		PartShapes_mapString: make(map[string]*PartShape),

		Ports:           make(map[*Port]struct{}),
		Ports_mapString: make(map[string]*Port),

		PortShapes:           make(map[*PortShape]struct{}),
		PortShapes_mapString: make(map[string]*PortShape),

		Resources:           make(map[*Resource]struct{}),
		Resources_mapString: make(map[string]*Resource),

		Systems:           make(map[*System]struct{}),
		Systems_mapString: make(map[string]*System),

		SystemShapes:           make(map[*SystemShape]struct{}),
		SystemShapes_mapString: make(map[string]*SystemShape),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		name: name,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		AllocatedResourceShape_stagedOrder: make(map[*AllocatedResourceShape]uint),
		AllocatedResourceShape_orderStaged: make(map[uint]*AllocatedResourceShape),
		AllocatedResourceShapes_reference:  make(map[*AllocatedResourceShape]*AllocatedResourceShape),

		AllocatedSystemShape_stagedOrder: make(map[*AllocatedSystemShape]uint),
		AllocatedSystemShape_orderStaged: make(map[uint]*AllocatedSystemShape),
		AllocatedSystemShapes_reference:  make(map[*AllocatedSystemShape]*AllocatedSystemShape),

		ControlFlow_stagedOrder: make(map[*ControlFlow]uint),
		ControlFlow_orderStaged: make(map[uint]*ControlFlow),
		ControlFlows_reference:  make(map[*ControlFlow]*ControlFlow),

		ControlFlowShape_stagedOrder: make(map[*ControlFlowShape]uint),
		ControlFlowShape_orderStaged: make(map[uint]*ControlFlowShape),
		ControlFlowShapes_reference:  make(map[*ControlFlowShape]*ControlFlowShape),

		Data_stagedOrder: make(map[*Data]uint),
		Data_orderStaged: make(map[uint]*Data),
		Datas_reference:  make(map[*Data]*Data),

		DataFlow_stagedOrder: make(map[*DataFlow]uint),
		DataFlow_orderStaged: make(map[uint]*DataFlow),
		DataFlows_reference:  make(map[*DataFlow]*DataFlow),

		DataFlowShape_stagedOrder: make(map[*DataFlowShape]uint),
		DataFlowShape_orderStaged: make(map[uint]*DataFlowShape),
		DataFlowShapes_reference:  make(map[*DataFlowShape]*DataFlowShape),

		DataShape_stagedOrder: make(map[*DataShape]uint),
		DataShape_orderStaged: make(map[uint]*DataShape),
		DataShapes_reference:  make(map[*DataShape]*DataShape),

		DiagramStructure_stagedOrder: make(map[*DiagramStructure]uint),
		DiagramStructure_orderStaged: make(map[uint]*DiagramStructure),
		DiagramStructures_reference:  make(map[*DiagramStructure]*DiagramStructure),

		ExternalPartShape_stagedOrder: make(map[*ExternalPartShape]uint),
		ExternalPartShape_orderStaged: make(map[uint]*ExternalPartShape),
		ExternalPartShapes_reference:  make(map[*ExternalPartShape]*ExternalPartShape),

		Library_stagedOrder: make(map[*Library]uint),
		Library_orderStaged: make(map[uint]*Library),
		Librarys_reference:  make(map[*Library]*Library),

		Note_stagedOrder: make(map[*Note]uint),
		Note_orderStaged: make(map[uint]*Note),
		Notes_reference:  make(map[*Note]*Note),

		NotePortShape_stagedOrder: make(map[*NotePortShape]uint),
		NotePortShape_orderStaged: make(map[uint]*NotePortShape),
		NotePortShapes_reference:  make(map[*NotePortShape]*NotePortShape),

		NoteShape_stagedOrder: make(map[*NoteShape]uint),
		NoteShape_orderStaged: make(map[uint]*NoteShape),
		NoteShapes_reference:  make(map[*NoteShape]*NoteShape),

		Part_stagedOrder: make(map[*Part]uint),
		Part_orderStaged: make(map[uint]*Part),
		Parts_reference:  make(map[*Part]*Part),

		PartAnchoredPath_stagedOrder: make(map[*PartAnchoredPath]uint),
		PartAnchoredPath_orderStaged: make(map[uint]*PartAnchoredPath),
		PartAnchoredPaths_reference:  make(map[*PartAnchoredPath]*PartAnchoredPath),

		PartShape_stagedOrder: make(map[*PartShape]uint),
		PartShape_orderStaged: make(map[uint]*PartShape),
		PartShapes_reference:  make(map[*PartShape]*PartShape),

		Port_stagedOrder: make(map[*Port]uint),
		Port_orderStaged: make(map[uint]*Port),
		Ports_reference:  make(map[*Port]*Port),

		PortShape_stagedOrder: make(map[*PortShape]uint),
		PortShape_orderStaged: make(map[uint]*PortShape),
		PortShapes_reference:  make(map[*PortShape]*PortShape),

		Resource_stagedOrder: make(map[*Resource]uint),
		Resource_orderStaged: make(map[uint]*Resource),
		Resources_reference:  make(map[*Resource]*Resource),

		System_stagedOrder: make(map[*System]uint),
		System_orderStaged: make(map[uint]*System),
		Systems_reference:  make(map[*System]*System),

		SystemShape_stagedOrder: make(map[*SystemShape]uint),
		SystemShape_orderStaged: make(map[uint]*SystemShape),
		SystemShapes_reference:  make(map[*SystemShape]*SystemShape),

		// end of insertion point
		GongUnmarshallers: map[string]ModelUnmarshaller{ // insertion point for unmarshallers
			"AllocatedResourceShape": &AllocatedResourceShapeUnmarshaller{},

			"AllocatedSystemShape": &AllocatedSystemShapeUnmarshaller{},

			"ControlFlow": &ControlFlowUnmarshaller{},

			"ControlFlowShape": &ControlFlowShapeUnmarshaller{},

			"Data": &DataUnmarshaller{},

			"DataFlow": &DataFlowUnmarshaller{},

			"DataFlowShape": &DataFlowShapeUnmarshaller{},

			"DataShape": &DataShapeUnmarshaller{},

			"DiagramStructure": &DiagramStructureUnmarshaller{},

			"ExternalPartShape": &ExternalPartShapeUnmarshaller{},

			"Library": &LibraryUnmarshaller{},

			"Note": &NoteUnmarshaller{},

			"NotePortShape": &NotePortShapeUnmarshaller{},

			"NoteShape": &NoteShapeUnmarshaller{},

			"Part": &PartUnmarshaller{},

			"PartAnchoredPath": &PartAnchoredPathUnmarshaller{},

			"PartShape": &PartShapeUnmarshaller{},

			"Port": &PortUnmarshaller{},

			"PortShape": &PortShapeUnmarshaller{},

			"Resource": &ResourceUnmarshaller{},

			"System": &SystemUnmarshaller{},

			"SystemShape": &SystemShapeUnmarshaller{},

			// end of insertion point
		},

		NamedStructs: []*NamedStruct{ // insertion point for order map initialisations
			{name: "AllocatedResourceShape"},
			{name: "AllocatedSystemShape"},
			{name: "ControlFlow"},
			{name: "ControlFlowShape"},
			{name: "Data"},
			{name: "DataFlow"},
			{name: "DataFlowShape"},
			{name: "DataShape"},
			{name: "DiagramStructure"},
			{name: "ExternalPartShape"},
			{name: "Library"},
			{name: "Note"},
			{name: "NotePortShape"},
			{name: "NoteShape"},
			{name: "Part"},
			{name: "PartAnchoredPath"},
			{name: "PartShape"},
			{name: "Port"},
			{name: "PortShape"},
			{name: "Resource"},
			{name: "System"},
			{name: "SystemShape"},
		}, // end of insertion point

		navigationMode: GongNavigationModeNormal,
	}

	return
}

func GetOrder[Type Gongstruct](stage *Stage, instance *Type) uint {
	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *AllocatedResourceShape:
		return stage.AllocatedResourceShape_stagedOrder[instance]
	case *AllocatedSystemShape:
		return stage.AllocatedSystemShape_stagedOrder[instance]
	case *ControlFlow:
		return stage.ControlFlow_stagedOrder[instance]
	case *ControlFlowShape:
		return stage.ControlFlowShape_stagedOrder[instance]
	case *Data:
		return stage.Data_stagedOrder[instance]
	case *DataFlow:
		return stage.DataFlow_stagedOrder[instance]
	case *DataFlowShape:
		return stage.DataFlowShape_stagedOrder[instance]
	case *DataShape:
		return stage.DataShape_stagedOrder[instance]
	case *DiagramStructure:
		return stage.DiagramStructure_stagedOrder[instance]
	case *ExternalPartShape:
		return stage.ExternalPartShape_stagedOrder[instance]
	case *Library:
		return stage.Library_stagedOrder[instance]
	case *Note:
		return stage.Note_stagedOrder[instance]
	case *NotePortShape:
		return stage.NotePortShape_stagedOrder[instance]
	case *NoteShape:
		return stage.NoteShape_stagedOrder[instance]
	case *Part:
		return stage.Part_stagedOrder[instance]
	case *PartAnchoredPath:
		return stage.PartAnchoredPath_stagedOrder[instance]
	case *PartShape:
		return stage.PartShape_stagedOrder[instance]
	case *Port:
		return stage.Port_stagedOrder[instance]
	case *PortShape:
		return stage.PortShape_stagedOrder[instance]
	case *Resource:
		return stage.Resource_stagedOrder[instance]
	case *System:
		return stage.System_stagedOrder[instance]
	case *SystemShape:
		return stage.SystemShape_stagedOrder[instance]
	default:
		return 0 // should not happen
	}
}

func GongGetInstanceFromOrder[Type PointerToGongstruct](stage *Stage, order uint) (res Type) {
	var t Type
	switch any(t).(type) {
	// insertion point for order map initialisations
	case *AllocatedResourceShape:
		return any(stage.AllocatedResourceShape_orderStaged[order]).(Type)
	case *AllocatedSystemShape:
		return any(stage.AllocatedSystemShape_orderStaged[order]).(Type)
	case *ControlFlow:
		return any(stage.ControlFlow_orderStaged[order]).(Type)
	case *ControlFlowShape:
		return any(stage.ControlFlowShape_orderStaged[order]).(Type)
	case *Data:
		return any(stage.Data_orderStaged[order]).(Type)
	case *DataFlow:
		return any(stage.DataFlow_orderStaged[order]).(Type)
	case *DataFlowShape:
		return any(stage.DataFlowShape_orderStaged[order]).(Type)
	case *DataShape:
		return any(stage.DataShape_orderStaged[order]).(Type)
	case *DiagramStructure:
		return any(stage.DiagramStructure_orderStaged[order]).(Type)
	case *ExternalPartShape:
		return any(stage.ExternalPartShape_orderStaged[order]).(Type)
	case *Library:
		return any(stage.Library_orderStaged[order]).(Type)
	case *Note:
		return any(stage.Note_orderStaged[order]).(Type)
	case *NotePortShape:
		return any(stage.NotePortShape_orderStaged[order]).(Type)
	case *NoteShape:
		return any(stage.NoteShape_orderStaged[order]).(Type)
	case *Part:
		return any(stage.Part_orderStaged[order]).(Type)
	case *PartAnchoredPath:
		return any(stage.PartAnchoredPath_orderStaged[order]).(Type)
	case *PartShape:
		return any(stage.PartShape_orderStaged[order]).(Type)
	case *Port:
		return any(stage.Port_orderStaged[order]).(Type)
	case *PortShape:
		return any(stage.PortShape_orderStaged[order]).(Type)
	case *Resource:
		return any(stage.Resource_orderStaged[order]).(Type)
	case *System:
		return any(stage.System_orderStaged[order]).(Type)
	case *SystemShape:
		return any(stage.SystemShape_orderStaged[order]).(Type)
	default:
		return // should not happen
	}
}

func GetOrderPointerGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) uint {
	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *AllocatedResourceShape:
		return stage.AllocatedResourceShape_stagedOrder[instance]
	case *AllocatedSystemShape:
		return stage.AllocatedSystemShape_stagedOrder[instance]
	case *ControlFlow:
		return stage.ControlFlow_stagedOrder[instance]
	case *ControlFlowShape:
		return stage.ControlFlowShape_stagedOrder[instance]
	case *Data:
		return stage.Data_stagedOrder[instance]
	case *DataFlow:
		return stage.DataFlow_stagedOrder[instance]
	case *DataFlowShape:
		return stage.DataFlowShape_stagedOrder[instance]
	case *DataShape:
		return stage.DataShape_stagedOrder[instance]
	case *DiagramStructure:
		return stage.DiagramStructure_stagedOrder[instance]
	case *ExternalPartShape:
		return stage.ExternalPartShape_stagedOrder[instance]
	case *Library:
		return stage.Library_stagedOrder[instance]
	case *Note:
		return stage.Note_stagedOrder[instance]
	case *NotePortShape:
		return stage.NotePortShape_stagedOrder[instance]
	case *NoteShape:
		return stage.NoteShape_stagedOrder[instance]
	case *Part:
		return stage.Part_stagedOrder[instance]
	case *PartAnchoredPath:
		return stage.PartAnchoredPath_stagedOrder[instance]
	case *PartShape:
		return stage.PartShape_stagedOrder[instance]
	case *Port:
		return stage.Port_stagedOrder[instance]
	case *PortShape:
		return stage.PortShape_stagedOrder[instance]
	case *Resource:
		return stage.Resource_stagedOrder[instance]
	case *System:
		return stage.System_stagedOrder[instance]
	case *SystemShape:
		return stage.SystemShape_stagedOrder[instance]
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
	stage.Map_GongStructName_InstancesNb["AllocatedResourceShape"] = len(stage.AllocatedResourceShapes)
	stage.Map_GongStructName_InstancesNb["AllocatedSystemShape"] = len(stage.AllocatedSystemShapes)
	stage.Map_GongStructName_InstancesNb["ControlFlow"] = len(stage.ControlFlows)
	stage.Map_GongStructName_InstancesNb["ControlFlowShape"] = len(stage.ControlFlowShapes)
	stage.Map_GongStructName_InstancesNb["Data"] = len(stage.Datas)
	stage.Map_GongStructName_InstancesNb["DataFlow"] = len(stage.DataFlows)
	stage.Map_GongStructName_InstancesNb["DataFlowShape"] = len(stage.DataFlowShapes)
	stage.Map_GongStructName_InstancesNb["DataShape"] = len(stage.DataShapes)
	stage.Map_GongStructName_InstancesNb["DiagramStructure"] = len(stage.DiagramStructures)
	stage.Map_GongStructName_InstancesNb["ExternalPartShape"] = len(stage.ExternalPartShapes)
	stage.Map_GongStructName_InstancesNb["Library"] = len(stage.Librarys)
	stage.Map_GongStructName_InstancesNb["Note"] = len(stage.Notes)
	stage.Map_GongStructName_InstancesNb["NotePortShape"] = len(stage.NotePortShapes)
	stage.Map_GongStructName_InstancesNb["NoteShape"] = len(stage.NoteShapes)
	stage.Map_GongStructName_InstancesNb["Part"] = len(stage.Parts)
	stage.Map_GongStructName_InstancesNb["PartAnchoredPath"] = len(stage.PartAnchoredPaths)
	stage.Map_GongStructName_InstancesNb["PartShape"] = len(stage.PartShapes)
	stage.Map_GongStructName_InstancesNb["Port"] = len(stage.Ports)
	stage.Map_GongStructName_InstancesNb["PortShape"] = len(stage.PortShapes)
	stage.Map_GongStructName_InstancesNb["Resource"] = len(stage.Resources)
	stage.Map_GongStructName_InstancesNb["System"] = len(stage.Systems)
	stage.Map_GongStructName_InstancesNb["SystemShape"] = len(stage.SystemShapes)
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
// Stage puts allocatedresourceshape to the model stage
func (allocatedresourceshape *AllocatedResourceShape) Stage(stage *Stage) *AllocatedResourceShape {
	if _, ok := stage.AllocatedResourceShapes[allocatedresourceshape]; !ok {
		stage.AllocatedResourceShapes[allocatedresourceshape] = struct{}{}
		stage.AllocatedResourceShape_stagedOrder[allocatedresourceshape] = stage.AllocatedResourceShapeOrder
		stage.AllocatedResourceShape_orderStaged[stage.AllocatedResourceShapeOrder] = allocatedresourceshape
		stage.AllocatedResourceShapeOrder++
	}
	stage.AllocatedResourceShapes_mapString[allocatedresourceshape.Name] = allocatedresourceshape

	return allocatedresourceshape
}

// StagePreserveOrder puts allocatedresourceshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AllocatedResourceShapeOrder
// - update stage.AllocatedResourceShapeOrder accordingly
func (allocatedresourceshape *AllocatedResourceShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.AllocatedResourceShapes[allocatedresourceshape]; !ok {
		stage.AllocatedResourceShapes[allocatedresourceshape] = struct{}{}

		if order > stage.AllocatedResourceShapeOrder {
			stage.AllocatedResourceShapeOrder = order
		}
		stage.AllocatedResourceShape_stagedOrder[allocatedresourceshape] = order
		stage.AllocatedResourceShape_orderStaged[order] = allocatedresourceshape
		stage.AllocatedResourceShapeOrder++
	}
	stage.AllocatedResourceShapes_mapString[allocatedresourceshape.Name] = allocatedresourceshape
}

// Unstage removes allocatedresourceshape off the model stage
func (allocatedresourceshape *AllocatedResourceShape) Unstage(stage *Stage) *AllocatedResourceShape {
	delete(stage.AllocatedResourceShapes, allocatedresourceshape)
	// issue1150
	// delete(stage.AllocatedResourceShape_stagedOrder, allocatedresourceshape)
	delete(stage.AllocatedResourceShapes_mapString, allocatedresourceshape.Name)

	return allocatedresourceshape
}

// UnstageVoid removes allocatedresourceshape off the model stage
func (allocatedresourceshape *AllocatedResourceShape) UnstageVoid(stage *Stage) {
	delete(stage.AllocatedResourceShapes, allocatedresourceshape)
	// issue1150
	// delete(stage.AllocatedResourceShape_stagedOrder, allocatedresourceshape)
	delete(stage.AllocatedResourceShapes_mapString, allocatedresourceshape.Name)
}

// commit allocatedresourceshape to the back repo (if it is already staged)
func (allocatedresourceshape *AllocatedResourceShape) Commit(stage *Stage) *AllocatedResourceShape {
	if _, ok := stage.AllocatedResourceShapes[allocatedresourceshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitAllocatedResourceShape(allocatedresourceshape)
		}
	}
	return allocatedresourceshape
}

func (allocatedresourceshape *AllocatedResourceShape) CommitVoid(stage *Stage) {
	allocatedresourceshape.Commit(stage)
}

func (allocatedresourceshape *AllocatedResourceShape) StageVoid(stage *Stage) {
	allocatedresourceshape.Stage(stage)
}

// Checkout allocatedresourceshape to the back repo (if it is already staged)
func (allocatedresourceshape *AllocatedResourceShape) Checkout(stage *Stage) *AllocatedResourceShape {
	if _, ok := stage.AllocatedResourceShapes[allocatedresourceshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutAllocatedResourceShape(allocatedresourceshape)
		}
	}
	return allocatedresourceshape
}

// for satisfaction of GongStruct interface
func (allocatedresourceshape *AllocatedResourceShape) GetName() (res string) {
	return allocatedresourceshape.Name
}

// for satisfaction of GongStruct interface
func (allocatedresourceshape *AllocatedResourceShape) SetName(name string) {
	allocatedresourceshape.Name = name
}

// Stage puts allocatedsystemshape to the model stage
func (allocatedsystemshape *AllocatedSystemShape) Stage(stage *Stage) *AllocatedSystemShape {
	if _, ok := stage.AllocatedSystemShapes[allocatedsystemshape]; !ok {
		stage.AllocatedSystemShapes[allocatedsystemshape] = struct{}{}
		stage.AllocatedSystemShape_stagedOrder[allocatedsystemshape] = stage.AllocatedSystemShapeOrder
		stage.AllocatedSystemShape_orderStaged[stage.AllocatedSystemShapeOrder] = allocatedsystemshape
		stage.AllocatedSystemShapeOrder++
	}
	stage.AllocatedSystemShapes_mapString[allocatedsystemshape.Name] = allocatedsystemshape

	return allocatedsystemshape
}

// StagePreserveOrder puts allocatedsystemshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AllocatedSystemShapeOrder
// - update stage.AllocatedSystemShapeOrder accordingly
func (allocatedsystemshape *AllocatedSystemShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.AllocatedSystemShapes[allocatedsystemshape]; !ok {
		stage.AllocatedSystemShapes[allocatedsystemshape] = struct{}{}

		if order > stage.AllocatedSystemShapeOrder {
			stage.AllocatedSystemShapeOrder = order
		}
		stage.AllocatedSystemShape_stagedOrder[allocatedsystemshape] = order
		stage.AllocatedSystemShape_orderStaged[order] = allocatedsystemshape
		stage.AllocatedSystemShapeOrder++
	}
	stage.AllocatedSystemShapes_mapString[allocatedsystemshape.Name] = allocatedsystemshape
}

// Unstage removes allocatedsystemshape off the model stage
func (allocatedsystemshape *AllocatedSystemShape) Unstage(stage *Stage) *AllocatedSystemShape {
	delete(stage.AllocatedSystemShapes, allocatedsystemshape)
	// issue1150
	// delete(stage.AllocatedSystemShape_stagedOrder, allocatedsystemshape)
	delete(stage.AllocatedSystemShapes_mapString, allocatedsystemshape.Name)

	return allocatedsystemshape
}

// UnstageVoid removes allocatedsystemshape off the model stage
func (allocatedsystemshape *AllocatedSystemShape) UnstageVoid(stage *Stage) {
	delete(stage.AllocatedSystemShapes, allocatedsystemshape)
	// issue1150
	// delete(stage.AllocatedSystemShape_stagedOrder, allocatedsystemshape)
	delete(stage.AllocatedSystemShapes_mapString, allocatedsystemshape.Name)
}

// commit allocatedsystemshape to the back repo (if it is already staged)
func (allocatedsystemshape *AllocatedSystemShape) Commit(stage *Stage) *AllocatedSystemShape {
	if _, ok := stage.AllocatedSystemShapes[allocatedsystemshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitAllocatedSystemShape(allocatedsystemshape)
		}
	}
	return allocatedsystemshape
}

func (allocatedsystemshape *AllocatedSystemShape) CommitVoid(stage *Stage) {
	allocatedsystemshape.Commit(stage)
}

func (allocatedsystemshape *AllocatedSystemShape) StageVoid(stage *Stage) {
	allocatedsystemshape.Stage(stage)
}

// Checkout allocatedsystemshape to the back repo (if it is already staged)
func (allocatedsystemshape *AllocatedSystemShape) Checkout(stage *Stage) *AllocatedSystemShape {
	if _, ok := stage.AllocatedSystemShapes[allocatedsystemshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutAllocatedSystemShape(allocatedsystemshape)
		}
	}
	return allocatedsystemshape
}

// for satisfaction of GongStruct interface
func (allocatedsystemshape *AllocatedSystemShape) GetName() (res string) {
	return allocatedsystemshape.Name
}

// for satisfaction of GongStruct interface
func (allocatedsystemshape *AllocatedSystemShape) SetName(name string) {
	allocatedsystemshape.Name = name
}

// Stage puts controlflow to the model stage
func (controlflow *ControlFlow) Stage(stage *Stage) *ControlFlow {
	if _, ok := stage.ControlFlows[controlflow]; !ok {
		stage.ControlFlows[controlflow] = struct{}{}
		stage.ControlFlow_stagedOrder[controlflow] = stage.ControlFlowOrder
		stage.ControlFlow_orderStaged[stage.ControlFlowOrder] = controlflow
		stage.ControlFlowOrder++
	}
	stage.ControlFlows_mapString[controlflow.Name] = controlflow

	return controlflow
}

// StagePreserveOrder puts controlflow to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ControlFlowOrder
// - update stage.ControlFlowOrder accordingly
func (controlflow *ControlFlow) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.ControlFlows[controlflow]; !ok {
		stage.ControlFlows[controlflow] = struct{}{}

		if order > stage.ControlFlowOrder {
			stage.ControlFlowOrder = order
		}
		stage.ControlFlow_stagedOrder[controlflow] = order
		stage.ControlFlow_orderStaged[order] = controlflow
		stage.ControlFlowOrder++
	}
	stage.ControlFlows_mapString[controlflow.Name] = controlflow
}

// Unstage removes controlflow off the model stage
func (controlflow *ControlFlow) Unstage(stage *Stage) *ControlFlow {
	delete(stage.ControlFlows, controlflow)
	// issue1150
	// delete(stage.ControlFlow_stagedOrder, controlflow)
	delete(stage.ControlFlows_mapString, controlflow.Name)

	return controlflow
}

// UnstageVoid removes controlflow off the model stage
func (controlflow *ControlFlow) UnstageVoid(stage *Stage) {
	delete(stage.ControlFlows, controlflow)
	// issue1150
	// delete(stage.ControlFlow_stagedOrder, controlflow)
	delete(stage.ControlFlows_mapString, controlflow.Name)
}

// commit controlflow to the back repo (if it is already staged)
func (controlflow *ControlFlow) Commit(stage *Stage) *ControlFlow {
	if _, ok := stage.ControlFlows[controlflow]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitControlFlow(controlflow)
		}
	}
	return controlflow
}

func (controlflow *ControlFlow) CommitVoid(stage *Stage) {
	controlflow.Commit(stage)
}

func (controlflow *ControlFlow) StageVoid(stage *Stage) {
	controlflow.Stage(stage)
}

// Checkout controlflow to the back repo (if it is already staged)
func (controlflow *ControlFlow) Checkout(stage *Stage) *ControlFlow {
	if _, ok := stage.ControlFlows[controlflow]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutControlFlow(controlflow)
		}
	}
	return controlflow
}

// for satisfaction of GongStruct interface
func (controlflow *ControlFlow) GetName() (res string) {
	return controlflow.Name
}

// for satisfaction of GongStruct interface
func (controlflow *ControlFlow) SetName(name string) {
	controlflow.Name = name
}

// Stage puts controlflowshape to the model stage
func (controlflowshape *ControlFlowShape) Stage(stage *Stage) *ControlFlowShape {
	if _, ok := stage.ControlFlowShapes[controlflowshape]; !ok {
		stage.ControlFlowShapes[controlflowshape] = struct{}{}
		stage.ControlFlowShape_stagedOrder[controlflowshape] = stage.ControlFlowShapeOrder
		stage.ControlFlowShape_orderStaged[stage.ControlFlowShapeOrder] = controlflowshape
		stage.ControlFlowShapeOrder++
	}
	stage.ControlFlowShapes_mapString[controlflowshape.Name] = controlflowshape

	return controlflowshape
}

// StagePreserveOrder puts controlflowshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ControlFlowShapeOrder
// - update stage.ControlFlowShapeOrder accordingly
func (controlflowshape *ControlFlowShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.ControlFlowShapes[controlflowshape]; !ok {
		stage.ControlFlowShapes[controlflowshape] = struct{}{}

		if order > stage.ControlFlowShapeOrder {
			stage.ControlFlowShapeOrder = order
		}
		stage.ControlFlowShape_stagedOrder[controlflowshape] = order
		stage.ControlFlowShape_orderStaged[order] = controlflowshape
		stage.ControlFlowShapeOrder++
	}
	stage.ControlFlowShapes_mapString[controlflowshape.Name] = controlflowshape
}

// Unstage removes controlflowshape off the model stage
func (controlflowshape *ControlFlowShape) Unstage(stage *Stage) *ControlFlowShape {
	delete(stage.ControlFlowShapes, controlflowshape)
	// issue1150
	// delete(stage.ControlFlowShape_stagedOrder, controlflowshape)
	delete(stage.ControlFlowShapes_mapString, controlflowshape.Name)

	return controlflowshape
}

// UnstageVoid removes controlflowshape off the model stage
func (controlflowshape *ControlFlowShape) UnstageVoid(stage *Stage) {
	delete(stage.ControlFlowShapes, controlflowshape)
	// issue1150
	// delete(stage.ControlFlowShape_stagedOrder, controlflowshape)
	delete(stage.ControlFlowShapes_mapString, controlflowshape.Name)
}

// commit controlflowshape to the back repo (if it is already staged)
func (controlflowshape *ControlFlowShape) Commit(stage *Stage) *ControlFlowShape {
	if _, ok := stage.ControlFlowShapes[controlflowshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitControlFlowShape(controlflowshape)
		}
	}
	return controlflowshape
}

func (controlflowshape *ControlFlowShape) CommitVoid(stage *Stage) {
	controlflowshape.Commit(stage)
}

func (controlflowshape *ControlFlowShape) StageVoid(stage *Stage) {
	controlflowshape.Stage(stage)
}

// Checkout controlflowshape to the back repo (if it is already staged)
func (controlflowshape *ControlFlowShape) Checkout(stage *Stage) *ControlFlowShape {
	if _, ok := stage.ControlFlowShapes[controlflowshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutControlFlowShape(controlflowshape)
		}
	}
	return controlflowshape
}

// for satisfaction of GongStruct interface
func (controlflowshape *ControlFlowShape) GetName() (res string) {
	return controlflowshape.Name
}

// for satisfaction of GongStruct interface
func (controlflowshape *ControlFlowShape) SetName(name string) {
	controlflowshape.Name = name
}

// Stage puts data to the model stage
func (data *Data) Stage(stage *Stage) *Data {
	if _, ok := stage.Datas[data]; !ok {
		stage.Datas[data] = struct{}{}
		stage.Data_stagedOrder[data] = stage.DataOrder
		stage.Data_orderStaged[stage.DataOrder] = data
		stage.DataOrder++
	}
	stage.Datas_mapString[data.Name] = data

	return data
}

// StagePreserveOrder puts data to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DataOrder
// - update stage.DataOrder accordingly
func (data *Data) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Datas[data]; !ok {
		stage.Datas[data] = struct{}{}

		if order > stage.DataOrder {
			stage.DataOrder = order
		}
		stage.Data_stagedOrder[data] = order
		stage.Data_orderStaged[order] = data
		stage.DataOrder++
	}
	stage.Datas_mapString[data.Name] = data
}

// Unstage removes data off the model stage
func (data *Data) Unstage(stage *Stage) *Data {
	delete(stage.Datas, data)
	// issue1150
	// delete(stage.Data_stagedOrder, data)
	delete(stage.Datas_mapString, data.Name)

	return data
}

// UnstageVoid removes data off the model stage
func (data *Data) UnstageVoid(stage *Stage) {
	delete(stage.Datas, data)
	// issue1150
	// delete(stage.Data_stagedOrder, data)
	delete(stage.Datas_mapString, data.Name)
}

// commit data to the back repo (if it is already staged)
func (data *Data) Commit(stage *Stage) *Data {
	if _, ok := stage.Datas[data]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitData(data)
		}
	}
	return data
}

func (data *Data) CommitVoid(stage *Stage) {
	data.Commit(stage)
}

func (data *Data) StageVoid(stage *Stage) {
	data.Stage(stage)
}

// Checkout data to the back repo (if it is already staged)
func (data *Data) Checkout(stage *Stage) *Data {
	if _, ok := stage.Datas[data]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutData(data)
		}
	}
	return data
}

// for satisfaction of GongStruct interface
func (data *Data) GetName() (res string) {
	return data.Name
}

// for satisfaction of GongStruct interface
func (data *Data) SetName(name string) {
	data.Name = name
}

// Stage puts dataflow to the model stage
func (dataflow *DataFlow) Stage(stage *Stage) *DataFlow {
	if _, ok := stage.DataFlows[dataflow]; !ok {
		stage.DataFlows[dataflow] = struct{}{}
		stage.DataFlow_stagedOrder[dataflow] = stage.DataFlowOrder
		stage.DataFlow_orderStaged[stage.DataFlowOrder] = dataflow
		stage.DataFlowOrder++
	}
	stage.DataFlows_mapString[dataflow.Name] = dataflow

	return dataflow
}

// StagePreserveOrder puts dataflow to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DataFlowOrder
// - update stage.DataFlowOrder accordingly
func (dataflow *DataFlow) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.DataFlows[dataflow]; !ok {
		stage.DataFlows[dataflow] = struct{}{}

		if order > stage.DataFlowOrder {
			stage.DataFlowOrder = order
		}
		stage.DataFlow_stagedOrder[dataflow] = order
		stage.DataFlow_orderStaged[order] = dataflow
		stage.DataFlowOrder++
	}
	stage.DataFlows_mapString[dataflow.Name] = dataflow
}

// Unstage removes dataflow off the model stage
func (dataflow *DataFlow) Unstage(stage *Stage) *DataFlow {
	delete(stage.DataFlows, dataflow)
	// issue1150
	// delete(stage.DataFlow_stagedOrder, dataflow)
	delete(stage.DataFlows_mapString, dataflow.Name)

	return dataflow
}

// UnstageVoid removes dataflow off the model stage
func (dataflow *DataFlow) UnstageVoid(stage *Stage) {
	delete(stage.DataFlows, dataflow)
	// issue1150
	// delete(stage.DataFlow_stagedOrder, dataflow)
	delete(stage.DataFlows_mapString, dataflow.Name)
}

// commit dataflow to the back repo (if it is already staged)
func (dataflow *DataFlow) Commit(stage *Stage) *DataFlow {
	if _, ok := stage.DataFlows[dataflow]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDataFlow(dataflow)
		}
	}
	return dataflow
}

func (dataflow *DataFlow) CommitVoid(stage *Stage) {
	dataflow.Commit(stage)
}

func (dataflow *DataFlow) StageVoid(stage *Stage) {
	dataflow.Stage(stage)
}

// Checkout dataflow to the back repo (if it is already staged)
func (dataflow *DataFlow) Checkout(stage *Stage) *DataFlow {
	if _, ok := stage.DataFlows[dataflow]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutDataFlow(dataflow)
		}
	}
	return dataflow
}

// for satisfaction of GongStruct interface
func (dataflow *DataFlow) GetName() (res string) {
	return dataflow.Name
}

// for satisfaction of GongStruct interface
func (dataflow *DataFlow) SetName(name string) {
	dataflow.Name = name
}

// Stage puts dataflowshape to the model stage
func (dataflowshape *DataFlowShape) Stage(stage *Stage) *DataFlowShape {
	if _, ok := stage.DataFlowShapes[dataflowshape]; !ok {
		stage.DataFlowShapes[dataflowshape] = struct{}{}
		stage.DataFlowShape_stagedOrder[dataflowshape] = stage.DataFlowShapeOrder
		stage.DataFlowShape_orderStaged[stage.DataFlowShapeOrder] = dataflowshape
		stage.DataFlowShapeOrder++
	}
	stage.DataFlowShapes_mapString[dataflowshape.Name] = dataflowshape

	return dataflowshape
}

// StagePreserveOrder puts dataflowshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DataFlowShapeOrder
// - update stage.DataFlowShapeOrder accordingly
func (dataflowshape *DataFlowShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.DataFlowShapes[dataflowshape]; !ok {
		stage.DataFlowShapes[dataflowshape] = struct{}{}

		if order > stage.DataFlowShapeOrder {
			stage.DataFlowShapeOrder = order
		}
		stage.DataFlowShape_stagedOrder[dataflowshape] = order
		stage.DataFlowShape_orderStaged[order] = dataflowshape
		stage.DataFlowShapeOrder++
	}
	stage.DataFlowShapes_mapString[dataflowshape.Name] = dataflowshape
}

// Unstage removes dataflowshape off the model stage
func (dataflowshape *DataFlowShape) Unstage(stage *Stage) *DataFlowShape {
	delete(stage.DataFlowShapes, dataflowshape)
	// issue1150
	// delete(stage.DataFlowShape_stagedOrder, dataflowshape)
	delete(stage.DataFlowShapes_mapString, dataflowshape.Name)

	return dataflowshape
}

// UnstageVoid removes dataflowshape off the model stage
func (dataflowshape *DataFlowShape) UnstageVoid(stage *Stage) {
	delete(stage.DataFlowShapes, dataflowshape)
	// issue1150
	// delete(stage.DataFlowShape_stagedOrder, dataflowshape)
	delete(stage.DataFlowShapes_mapString, dataflowshape.Name)
}

// commit dataflowshape to the back repo (if it is already staged)
func (dataflowshape *DataFlowShape) Commit(stage *Stage) *DataFlowShape {
	if _, ok := stage.DataFlowShapes[dataflowshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDataFlowShape(dataflowshape)
		}
	}
	return dataflowshape
}

func (dataflowshape *DataFlowShape) CommitVoid(stage *Stage) {
	dataflowshape.Commit(stage)
}

func (dataflowshape *DataFlowShape) StageVoid(stage *Stage) {
	dataflowshape.Stage(stage)
}

// Checkout dataflowshape to the back repo (if it is already staged)
func (dataflowshape *DataFlowShape) Checkout(stage *Stage) *DataFlowShape {
	if _, ok := stage.DataFlowShapes[dataflowshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutDataFlowShape(dataflowshape)
		}
	}
	return dataflowshape
}

// for satisfaction of GongStruct interface
func (dataflowshape *DataFlowShape) GetName() (res string) {
	return dataflowshape.Name
}

// for satisfaction of GongStruct interface
func (dataflowshape *DataFlowShape) SetName(name string) {
	dataflowshape.Name = name
}

// Stage puts datashape to the model stage
func (datashape *DataShape) Stage(stage *Stage) *DataShape {
	if _, ok := stage.DataShapes[datashape]; !ok {
		stage.DataShapes[datashape] = struct{}{}
		stage.DataShape_stagedOrder[datashape] = stage.DataShapeOrder
		stage.DataShape_orderStaged[stage.DataShapeOrder] = datashape
		stage.DataShapeOrder++
	}
	stage.DataShapes_mapString[datashape.Name] = datashape

	return datashape
}

// StagePreserveOrder puts datashape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DataShapeOrder
// - update stage.DataShapeOrder accordingly
func (datashape *DataShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.DataShapes[datashape]; !ok {
		stage.DataShapes[datashape] = struct{}{}

		if order > stage.DataShapeOrder {
			stage.DataShapeOrder = order
		}
		stage.DataShape_stagedOrder[datashape] = order
		stage.DataShape_orderStaged[order] = datashape
		stage.DataShapeOrder++
	}
	stage.DataShapes_mapString[datashape.Name] = datashape
}

// Unstage removes datashape off the model stage
func (datashape *DataShape) Unstage(stage *Stage) *DataShape {
	delete(stage.DataShapes, datashape)
	// issue1150
	// delete(stage.DataShape_stagedOrder, datashape)
	delete(stage.DataShapes_mapString, datashape.Name)

	return datashape
}

// UnstageVoid removes datashape off the model stage
func (datashape *DataShape) UnstageVoid(stage *Stage) {
	delete(stage.DataShapes, datashape)
	// issue1150
	// delete(stage.DataShape_stagedOrder, datashape)
	delete(stage.DataShapes_mapString, datashape.Name)
}

// commit datashape to the back repo (if it is already staged)
func (datashape *DataShape) Commit(stage *Stage) *DataShape {
	if _, ok := stage.DataShapes[datashape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDataShape(datashape)
		}
	}
	return datashape
}

func (datashape *DataShape) CommitVoid(stage *Stage) {
	datashape.Commit(stage)
}

func (datashape *DataShape) StageVoid(stage *Stage) {
	datashape.Stage(stage)
}

// Checkout datashape to the back repo (if it is already staged)
func (datashape *DataShape) Checkout(stage *Stage) *DataShape {
	if _, ok := stage.DataShapes[datashape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutDataShape(datashape)
		}
	}
	return datashape
}

// for satisfaction of GongStruct interface
func (datashape *DataShape) GetName() (res string) {
	return datashape.Name
}

// for satisfaction of GongStruct interface
func (datashape *DataShape) SetName(name string) {
	datashape.Name = name
}

// Stage puts diagramstructure to the model stage
func (diagramstructure *DiagramStructure) Stage(stage *Stage) *DiagramStructure {
	if _, ok := stage.DiagramStructures[diagramstructure]; !ok {
		stage.DiagramStructures[diagramstructure] = struct{}{}
		stage.DiagramStructure_stagedOrder[diagramstructure] = stage.DiagramStructureOrder
		stage.DiagramStructure_orderStaged[stage.DiagramStructureOrder] = diagramstructure
		stage.DiagramStructureOrder++
	}
	stage.DiagramStructures_mapString[diagramstructure.Name] = diagramstructure

	return diagramstructure
}

// StagePreserveOrder puts diagramstructure to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DiagramStructureOrder
// - update stage.DiagramStructureOrder accordingly
func (diagramstructure *DiagramStructure) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.DiagramStructures[diagramstructure]; !ok {
		stage.DiagramStructures[diagramstructure] = struct{}{}

		if order > stage.DiagramStructureOrder {
			stage.DiagramStructureOrder = order
		}
		stage.DiagramStructure_stagedOrder[diagramstructure] = order
		stage.DiagramStructure_orderStaged[order] = diagramstructure
		stage.DiagramStructureOrder++
	}
	stage.DiagramStructures_mapString[diagramstructure.Name] = diagramstructure
}

// Unstage removes diagramstructure off the model stage
func (diagramstructure *DiagramStructure) Unstage(stage *Stage) *DiagramStructure {
	delete(stage.DiagramStructures, diagramstructure)
	// issue1150
	// delete(stage.DiagramStructure_stagedOrder, diagramstructure)
	delete(stage.DiagramStructures_mapString, diagramstructure.Name)

	return diagramstructure
}

// UnstageVoid removes diagramstructure off the model stage
func (diagramstructure *DiagramStructure) UnstageVoid(stage *Stage) {
	delete(stage.DiagramStructures, diagramstructure)
	// issue1150
	// delete(stage.DiagramStructure_stagedOrder, diagramstructure)
	delete(stage.DiagramStructures_mapString, diagramstructure.Name)
}

// commit diagramstructure to the back repo (if it is already staged)
func (diagramstructure *DiagramStructure) Commit(stage *Stage) *DiagramStructure {
	if _, ok := stage.DiagramStructures[diagramstructure]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDiagramStructure(diagramstructure)
		}
	}
	return diagramstructure
}

func (diagramstructure *DiagramStructure) CommitVoid(stage *Stage) {
	diagramstructure.Commit(stage)
}

func (diagramstructure *DiagramStructure) StageVoid(stage *Stage) {
	diagramstructure.Stage(stage)
}

// Checkout diagramstructure to the back repo (if it is already staged)
func (diagramstructure *DiagramStructure) Checkout(stage *Stage) *DiagramStructure {
	if _, ok := stage.DiagramStructures[diagramstructure]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutDiagramStructure(diagramstructure)
		}
	}
	return diagramstructure
}

// for satisfaction of GongStruct interface
func (diagramstructure *DiagramStructure) GetName() (res string) {
	return diagramstructure.Name
}

// for satisfaction of GongStruct interface
func (diagramstructure *DiagramStructure) SetName(name string) {
	diagramstructure.Name = name
}

// Stage puts externalpartshape to the model stage
func (externalpartshape *ExternalPartShape) Stage(stage *Stage) *ExternalPartShape {
	if _, ok := stage.ExternalPartShapes[externalpartshape]; !ok {
		stage.ExternalPartShapes[externalpartshape] = struct{}{}
		stage.ExternalPartShape_stagedOrder[externalpartshape] = stage.ExternalPartShapeOrder
		stage.ExternalPartShape_orderStaged[stage.ExternalPartShapeOrder] = externalpartshape
		stage.ExternalPartShapeOrder++
	}
	stage.ExternalPartShapes_mapString[externalpartshape.Name] = externalpartshape

	return externalpartshape
}

// StagePreserveOrder puts externalpartshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ExternalPartShapeOrder
// - update stage.ExternalPartShapeOrder accordingly
func (externalpartshape *ExternalPartShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.ExternalPartShapes[externalpartshape]; !ok {
		stage.ExternalPartShapes[externalpartshape] = struct{}{}

		if order > stage.ExternalPartShapeOrder {
			stage.ExternalPartShapeOrder = order
		}
		stage.ExternalPartShape_stagedOrder[externalpartshape] = order
		stage.ExternalPartShape_orderStaged[order] = externalpartshape
		stage.ExternalPartShapeOrder++
	}
	stage.ExternalPartShapes_mapString[externalpartshape.Name] = externalpartshape
}

// Unstage removes externalpartshape off the model stage
func (externalpartshape *ExternalPartShape) Unstage(stage *Stage) *ExternalPartShape {
	delete(stage.ExternalPartShapes, externalpartshape)
	// issue1150
	// delete(stage.ExternalPartShape_stagedOrder, externalpartshape)
	delete(stage.ExternalPartShapes_mapString, externalpartshape.Name)

	return externalpartshape
}

// UnstageVoid removes externalpartshape off the model stage
func (externalpartshape *ExternalPartShape) UnstageVoid(stage *Stage) {
	delete(stage.ExternalPartShapes, externalpartshape)
	// issue1150
	// delete(stage.ExternalPartShape_stagedOrder, externalpartshape)
	delete(stage.ExternalPartShapes_mapString, externalpartshape.Name)
}

// commit externalpartshape to the back repo (if it is already staged)
func (externalpartshape *ExternalPartShape) Commit(stage *Stage) *ExternalPartShape {
	if _, ok := stage.ExternalPartShapes[externalpartshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitExternalPartShape(externalpartshape)
		}
	}
	return externalpartshape
}

func (externalpartshape *ExternalPartShape) CommitVoid(stage *Stage) {
	externalpartshape.Commit(stage)
}

func (externalpartshape *ExternalPartShape) StageVoid(stage *Stage) {
	externalpartshape.Stage(stage)
}

// Checkout externalpartshape to the back repo (if it is already staged)
func (externalpartshape *ExternalPartShape) Checkout(stage *Stage) *ExternalPartShape {
	if _, ok := stage.ExternalPartShapes[externalpartshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutExternalPartShape(externalpartshape)
		}
	}
	return externalpartshape
}

// for satisfaction of GongStruct interface
func (externalpartshape *ExternalPartShape) GetName() (res string) {
	return externalpartshape.Name
}

// for satisfaction of GongStruct interface
func (externalpartshape *ExternalPartShape) SetName(name string) {
	externalpartshape.Name = name
}

// Stage puts library to the model stage
func (library *Library) Stage(stage *Stage) *Library {
	if _, ok := stage.Librarys[library]; !ok {
		stage.Librarys[library] = struct{}{}
		stage.Library_stagedOrder[library] = stage.LibraryOrder
		stage.Library_orderStaged[stage.LibraryOrder] = library
		stage.LibraryOrder++
	}
	stage.Librarys_mapString[library.Name] = library

	return library
}

// StagePreserveOrder puts library to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.LibraryOrder
// - update stage.LibraryOrder accordingly
func (library *Library) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Librarys[library]; !ok {
		stage.Librarys[library] = struct{}{}

		if order > stage.LibraryOrder {
			stage.LibraryOrder = order
		}
		stage.Library_stagedOrder[library] = order
		stage.Library_orderStaged[order] = library
		stage.LibraryOrder++
	}
	stage.Librarys_mapString[library.Name] = library
}

// Unstage removes library off the model stage
func (library *Library) Unstage(stage *Stage) *Library {
	delete(stage.Librarys, library)
	// issue1150
	// delete(stage.Library_stagedOrder, library)
	delete(stage.Librarys_mapString, library.Name)

	return library
}

// UnstageVoid removes library off the model stage
func (library *Library) UnstageVoid(stage *Stage) {
	delete(stage.Librarys, library)
	// issue1150
	// delete(stage.Library_stagedOrder, library)
	delete(stage.Librarys_mapString, library.Name)
}

// commit library to the back repo (if it is already staged)
func (library *Library) Commit(stage *Stage) *Library {
	if _, ok := stage.Librarys[library]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitLibrary(library)
		}
	}
	return library
}

func (library *Library) CommitVoid(stage *Stage) {
	library.Commit(stage)
}

func (library *Library) StageVoid(stage *Stage) {
	library.Stage(stage)
}

// Checkout library to the back repo (if it is already staged)
func (library *Library) Checkout(stage *Stage) *Library {
	if _, ok := stage.Librarys[library]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutLibrary(library)
		}
	}
	return library
}

// for satisfaction of GongStruct interface
func (library *Library) GetName() (res string) {
	return library.Name
}

// for satisfaction of GongStruct interface
func (library *Library) SetName(name string) {
	library.Name = name
}

// Stage puts note to the model stage
func (note *Note) Stage(stage *Stage) *Note {
	if _, ok := stage.Notes[note]; !ok {
		stage.Notes[note] = struct{}{}
		stage.Note_stagedOrder[note] = stage.NoteOrder
		stage.Note_orderStaged[stage.NoteOrder] = note
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
		stage.Note_stagedOrder[note] = order
		stage.Note_orderStaged[order] = note
		stage.NoteOrder++
	}
	stage.Notes_mapString[note.Name] = note
}

// Unstage removes note off the model stage
func (note *Note) Unstage(stage *Stage) *Note {
	delete(stage.Notes, note)
	// issue1150
	// delete(stage.Note_stagedOrder, note)
	delete(stage.Notes_mapString, note.Name)

	return note
}

// UnstageVoid removes note off the model stage
func (note *Note) UnstageVoid(stage *Stage) {
	delete(stage.Notes, note)
	// issue1150
	// delete(stage.Note_stagedOrder, note)
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

// Stage puts noteportshape to the model stage
func (noteportshape *NotePortShape) Stage(stage *Stage) *NotePortShape {
	if _, ok := stage.NotePortShapes[noteportshape]; !ok {
		stage.NotePortShapes[noteportshape] = struct{}{}
		stage.NotePortShape_stagedOrder[noteportshape] = stage.NotePortShapeOrder
		stage.NotePortShape_orderStaged[stage.NotePortShapeOrder] = noteportshape
		stage.NotePortShapeOrder++
	}
	stage.NotePortShapes_mapString[noteportshape.Name] = noteportshape

	return noteportshape
}

// StagePreserveOrder puts noteportshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.NotePortShapeOrder
// - update stage.NotePortShapeOrder accordingly
func (noteportshape *NotePortShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.NotePortShapes[noteportshape]; !ok {
		stage.NotePortShapes[noteportshape] = struct{}{}

		if order > stage.NotePortShapeOrder {
			stage.NotePortShapeOrder = order
		}
		stage.NotePortShape_stagedOrder[noteportshape] = order
		stage.NotePortShape_orderStaged[order] = noteportshape
		stage.NotePortShapeOrder++
	}
	stage.NotePortShapes_mapString[noteportshape.Name] = noteportshape
}

// Unstage removes noteportshape off the model stage
func (noteportshape *NotePortShape) Unstage(stage *Stage) *NotePortShape {
	delete(stage.NotePortShapes, noteportshape)
	// issue1150
	// delete(stage.NotePortShape_stagedOrder, noteportshape)
	delete(stage.NotePortShapes_mapString, noteportshape.Name)

	return noteportshape
}

// UnstageVoid removes noteportshape off the model stage
func (noteportshape *NotePortShape) UnstageVoid(stage *Stage) {
	delete(stage.NotePortShapes, noteportshape)
	// issue1150
	// delete(stage.NotePortShape_stagedOrder, noteportshape)
	delete(stage.NotePortShapes_mapString, noteportshape.Name)
}

// commit noteportshape to the back repo (if it is already staged)
func (noteportshape *NotePortShape) Commit(stage *Stage) *NotePortShape {
	if _, ok := stage.NotePortShapes[noteportshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitNotePortShape(noteportshape)
		}
	}
	return noteportshape
}

func (noteportshape *NotePortShape) CommitVoid(stage *Stage) {
	noteportshape.Commit(stage)
}

func (noteportshape *NotePortShape) StageVoid(stage *Stage) {
	noteportshape.Stage(stage)
}

// Checkout noteportshape to the back repo (if it is already staged)
func (noteportshape *NotePortShape) Checkout(stage *Stage) *NotePortShape {
	if _, ok := stage.NotePortShapes[noteportshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutNotePortShape(noteportshape)
		}
	}
	return noteportshape
}

// for satisfaction of GongStruct interface
func (noteportshape *NotePortShape) GetName() (res string) {
	return noteportshape.Name
}

// for satisfaction of GongStruct interface
func (noteportshape *NotePortShape) SetName(name string) {
	noteportshape.Name = name
}

// Stage puts noteshape to the model stage
func (noteshape *NoteShape) Stage(stage *Stage) *NoteShape {
	if _, ok := stage.NoteShapes[noteshape]; !ok {
		stage.NoteShapes[noteshape] = struct{}{}
		stage.NoteShape_stagedOrder[noteshape] = stage.NoteShapeOrder
		stage.NoteShape_orderStaged[stage.NoteShapeOrder] = noteshape
		stage.NoteShapeOrder++
	}
	stage.NoteShapes_mapString[noteshape.Name] = noteshape

	return noteshape
}

// StagePreserveOrder puts noteshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.NoteShapeOrder
// - update stage.NoteShapeOrder accordingly
func (noteshape *NoteShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.NoteShapes[noteshape]; !ok {
		stage.NoteShapes[noteshape] = struct{}{}

		if order > stage.NoteShapeOrder {
			stage.NoteShapeOrder = order
		}
		stage.NoteShape_stagedOrder[noteshape] = order
		stage.NoteShape_orderStaged[order] = noteshape
		stage.NoteShapeOrder++
	}
	stage.NoteShapes_mapString[noteshape.Name] = noteshape
}

// Unstage removes noteshape off the model stage
func (noteshape *NoteShape) Unstage(stage *Stage) *NoteShape {
	delete(stage.NoteShapes, noteshape)
	// issue1150
	// delete(stage.NoteShape_stagedOrder, noteshape)
	delete(stage.NoteShapes_mapString, noteshape.Name)

	return noteshape
}

// UnstageVoid removes noteshape off the model stage
func (noteshape *NoteShape) UnstageVoid(stage *Stage) {
	delete(stage.NoteShapes, noteshape)
	// issue1150
	// delete(stage.NoteShape_stagedOrder, noteshape)
	delete(stage.NoteShapes_mapString, noteshape.Name)
}

// commit noteshape to the back repo (if it is already staged)
func (noteshape *NoteShape) Commit(stage *Stage) *NoteShape {
	if _, ok := stage.NoteShapes[noteshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitNoteShape(noteshape)
		}
	}
	return noteshape
}

func (noteshape *NoteShape) CommitVoid(stage *Stage) {
	noteshape.Commit(stage)
}

func (noteshape *NoteShape) StageVoid(stage *Stage) {
	noteshape.Stage(stage)
}

// Checkout noteshape to the back repo (if it is already staged)
func (noteshape *NoteShape) Checkout(stage *Stage) *NoteShape {
	if _, ok := stage.NoteShapes[noteshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutNoteShape(noteshape)
		}
	}
	return noteshape
}

// for satisfaction of GongStruct interface
func (noteshape *NoteShape) GetName() (res string) {
	return noteshape.Name
}

// for satisfaction of GongStruct interface
func (noteshape *NoteShape) SetName(name string) {
	noteshape.Name = name
}

// Stage puts part to the model stage
func (part *Part) Stage(stage *Stage) *Part {
	if _, ok := stage.Parts[part]; !ok {
		stage.Parts[part] = struct{}{}
		stage.Part_stagedOrder[part] = stage.PartOrder
		stage.Part_orderStaged[stage.PartOrder] = part
		stage.PartOrder++
	}
	stage.Parts_mapString[part.Name] = part

	return part
}

// StagePreserveOrder puts part to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.PartOrder
// - update stage.PartOrder accordingly
func (part *Part) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Parts[part]; !ok {
		stage.Parts[part] = struct{}{}

		if order > stage.PartOrder {
			stage.PartOrder = order
		}
		stage.Part_stagedOrder[part] = order
		stage.Part_orderStaged[order] = part
		stage.PartOrder++
	}
	stage.Parts_mapString[part.Name] = part
}

// Unstage removes part off the model stage
func (part *Part) Unstage(stage *Stage) *Part {
	delete(stage.Parts, part)
	// issue1150
	// delete(stage.Part_stagedOrder, part)
	delete(stage.Parts_mapString, part.Name)

	return part
}

// UnstageVoid removes part off the model stage
func (part *Part) UnstageVoid(stage *Stage) {
	delete(stage.Parts, part)
	// issue1150
	// delete(stage.Part_stagedOrder, part)
	delete(stage.Parts_mapString, part.Name)
}

// commit part to the back repo (if it is already staged)
func (part *Part) Commit(stage *Stage) *Part {
	if _, ok := stage.Parts[part]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitPart(part)
		}
	}
	return part
}

func (part *Part) CommitVoid(stage *Stage) {
	part.Commit(stage)
}

func (part *Part) StageVoid(stage *Stage) {
	part.Stage(stage)
}

// Checkout part to the back repo (if it is already staged)
func (part *Part) Checkout(stage *Stage) *Part {
	if _, ok := stage.Parts[part]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutPart(part)
		}
	}
	return part
}

// for satisfaction of GongStruct interface
func (part *Part) GetName() (res string) {
	return part.Name
}

// for satisfaction of GongStruct interface
func (part *Part) SetName(name string) {
	part.Name = name
}

// Stage puts partanchoredpath to the model stage
func (partanchoredpath *PartAnchoredPath) Stage(stage *Stage) *PartAnchoredPath {
	if _, ok := stage.PartAnchoredPaths[partanchoredpath]; !ok {
		stage.PartAnchoredPaths[partanchoredpath] = struct{}{}
		stage.PartAnchoredPath_stagedOrder[partanchoredpath] = stage.PartAnchoredPathOrder
		stage.PartAnchoredPath_orderStaged[stage.PartAnchoredPathOrder] = partanchoredpath
		stage.PartAnchoredPathOrder++
	}
	stage.PartAnchoredPaths_mapString[partanchoredpath.Name] = partanchoredpath

	return partanchoredpath
}

// StagePreserveOrder puts partanchoredpath to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.PartAnchoredPathOrder
// - update stage.PartAnchoredPathOrder accordingly
func (partanchoredpath *PartAnchoredPath) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.PartAnchoredPaths[partanchoredpath]; !ok {
		stage.PartAnchoredPaths[partanchoredpath] = struct{}{}

		if order > stage.PartAnchoredPathOrder {
			stage.PartAnchoredPathOrder = order
		}
		stage.PartAnchoredPath_stagedOrder[partanchoredpath] = order
		stage.PartAnchoredPath_orderStaged[order] = partanchoredpath
		stage.PartAnchoredPathOrder++
	}
	stage.PartAnchoredPaths_mapString[partanchoredpath.Name] = partanchoredpath
}

// Unstage removes partanchoredpath off the model stage
func (partanchoredpath *PartAnchoredPath) Unstage(stage *Stage) *PartAnchoredPath {
	delete(stage.PartAnchoredPaths, partanchoredpath)
	// issue1150
	// delete(stage.PartAnchoredPath_stagedOrder, partanchoredpath)
	delete(stage.PartAnchoredPaths_mapString, partanchoredpath.Name)

	return partanchoredpath
}

// UnstageVoid removes partanchoredpath off the model stage
func (partanchoredpath *PartAnchoredPath) UnstageVoid(stage *Stage) {
	delete(stage.PartAnchoredPaths, partanchoredpath)
	// issue1150
	// delete(stage.PartAnchoredPath_stagedOrder, partanchoredpath)
	delete(stage.PartAnchoredPaths_mapString, partanchoredpath.Name)
}

// commit partanchoredpath to the back repo (if it is already staged)
func (partanchoredpath *PartAnchoredPath) Commit(stage *Stage) *PartAnchoredPath {
	if _, ok := stage.PartAnchoredPaths[partanchoredpath]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitPartAnchoredPath(partanchoredpath)
		}
	}
	return partanchoredpath
}

func (partanchoredpath *PartAnchoredPath) CommitVoid(stage *Stage) {
	partanchoredpath.Commit(stage)
}

func (partanchoredpath *PartAnchoredPath) StageVoid(stage *Stage) {
	partanchoredpath.Stage(stage)
}

// Checkout partanchoredpath to the back repo (if it is already staged)
func (partanchoredpath *PartAnchoredPath) Checkout(stage *Stage) *PartAnchoredPath {
	if _, ok := stage.PartAnchoredPaths[partanchoredpath]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutPartAnchoredPath(partanchoredpath)
		}
	}
	return partanchoredpath
}

// for satisfaction of GongStruct interface
func (partanchoredpath *PartAnchoredPath) GetName() (res string) {
	return partanchoredpath.Name
}

// for satisfaction of GongStruct interface
func (partanchoredpath *PartAnchoredPath) SetName(name string) {
	partanchoredpath.Name = name
}

// Stage puts partshape to the model stage
func (partshape *PartShape) Stage(stage *Stage) *PartShape {
	if _, ok := stage.PartShapes[partshape]; !ok {
		stage.PartShapes[partshape] = struct{}{}
		stage.PartShape_stagedOrder[partshape] = stage.PartShapeOrder
		stage.PartShape_orderStaged[stage.PartShapeOrder] = partshape
		stage.PartShapeOrder++
	}
	stage.PartShapes_mapString[partshape.Name] = partshape

	return partshape
}

// StagePreserveOrder puts partshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.PartShapeOrder
// - update stage.PartShapeOrder accordingly
func (partshape *PartShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.PartShapes[partshape]; !ok {
		stage.PartShapes[partshape] = struct{}{}

		if order > stage.PartShapeOrder {
			stage.PartShapeOrder = order
		}
		stage.PartShape_stagedOrder[partshape] = order
		stage.PartShape_orderStaged[order] = partshape
		stage.PartShapeOrder++
	}
	stage.PartShapes_mapString[partshape.Name] = partshape
}

// Unstage removes partshape off the model stage
func (partshape *PartShape) Unstage(stage *Stage) *PartShape {
	delete(stage.PartShapes, partshape)
	// issue1150
	// delete(stage.PartShape_stagedOrder, partshape)
	delete(stage.PartShapes_mapString, partshape.Name)

	return partshape
}

// UnstageVoid removes partshape off the model stage
func (partshape *PartShape) UnstageVoid(stage *Stage) {
	delete(stage.PartShapes, partshape)
	// issue1150
	// delete(stage.PartShape_stagedOrder, partshape)
	delete(stage.PartShapes_mapString, partshape.Name)
}

// commit partshape to the back repo (if it is already staged)
func (partshape *PartShape) Commit(stage *Stage) *PartShape {
	if _, ok := stage.PartShapes[partshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitPartShape(partshape)
		}
	}
	return partshape
}

func (partshape *PartShape) CommitVoid(stage *Stage) {
	partshape.Commit(stage)
}

func (partshape *PartShape) StageVoid(stage *Stage) {
	partshape.Stage(stage)
}

// Checkout partshape to the back repo (if it is already staged)
func (partshape *PartShape) Checkout(stage *Stage) *PartShape {
	if _, ok := stage.PartShapes[partshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutPartShape(partshape)
		}
	}
	return partshape
}

// for satisfaction of GongStruct interface
func (partshape *PartShape) GetName() (res string) {
	return partshape.Name
}

// for satisfaction of GongStruct interface
func (partshape *PartShape) SetName(name string) {
	partshape.Name = name
}

// Stage puts port to the model stage
func (port *Port) Stage(stage *Stage) *Port {
	if _, ok := stage.Ports[port]; !ok {
		stage.Ports[port] = struct{}{}
		stage.Port_stagedOrder[port] = stage.PortOrder
		stage.Port_orderStaged[stage.PortOrder] = port
		stage.PortOrder++
	}
	stage.Ports_mapString[port.Name] = port

	return port
}

// StagePreserveOrder puts port to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.PortOrder
// - update stage.PortOrder accordingly
func (port *Port) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Ports[port]; !ok {
		stage.Ports[port] = struct{}{}

		if order > stage.PortOrder {
			stage.PortOrder = order
		}
		stage.Port_stagedOrder[port] = order
		stage.Port_orderStaged[order] = port
		stage.PortOrder++
	}
	stage.Ports_mapString[port.Name] = port
}

// Unstage removes port off the model stage
func (port *Port) Unstage(stage *Stage) *Port {
	delete(stage.Ports, port)
	// issue1150
	// delete(stage.Port_stagedOrder, port)
	delete(stage.Ports_mapString, port.Name)

	return port
}

// UnstageVoid removes port off the model stage
func (port *Port) UnstageVoid(stage *Stage) {
	delete(stage.Ports, port)
	// issue1150
	// delete(stage.Port_stagedOrder, port)
	delete(stage.Ports_mapString, port.Name)
}

// commit port to the back repo (if it is already staged)
func (port *Port) Commit(stage *Stage) *Port {
	if _, ok := stage.Ports[port]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitPort(port)
		}
	}
	return port
}

func (port *Port) CommitVoid(stage *Stage) {
	port.Commit(stage)
}

func (port *Port) StageVoid(stage *Stage) {
	port.Stage(stage)
}

// Checkout port to the back repo (if it is already staged)
func (port *Port) Checkout(stage *Stage) *Port {
	if _, ok := stage.Ports[port]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutPort(port)
		}
	}
	return port
}

// for satisfaction of GongStruct interface
func (port *Port) GetName() (res string) {
	return port.Name
}

// for satisfaction of GongStruct interface
func (port *Port) SetName(name string) {
	port.Name = name
}

// Stage puts portshape to the model stage
func (portshape *PortShape) Stage(stage *Stage) *PortShape {
	if _, ok := stage.PortShapes[portshape]; !ok {
		stage.PortShapes[portshape] = struct{}{}
		stage.PortShape_stagedOrder[portshape] = stage.PortShapeOrder
		stage.PortShape_orderStaged[stage.PortShapeOrder] = portshape
		stage.PortShapeOrder++
	}
	stage.PortShapes_mapString[portshape.Name] = portshape

	return portshape
}

// StagePreserveOrder puts portshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.PortShapeOrder
// - update stage.PortShapeOrder accordingly
func (portshape *PortShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.PortShapes[portshape]; !ok {
		stage.PortShapes[portshape] = struct{}{}

		if order > stage.PortShapeOrder {
			stage.PortShapeOrder = order
		}
		stage.PortShape_stagedOrder[portshape] = order
		stage.PortShape_orderStaged[order] = portshape
		stage.PortShapeOrder++
	}
	stage.PortShapes_mapString[portshape.Name] = portshape
}

// Unstage removes portshape off the model stage
func (portshape *PortShape) Unstage(stage *Stage) *PortShape {
	delete(stage.PortShapes, portshape)
	// issue1150
	// delete(stage.PortShape_stagedOrder, portshape)
	delete(stage.PortShapes_mapString, portshape.Name)

	return portshape
}

// UnstageVoid removes portshape off the model stage
func (portshape *PortShape) UnstageVoid(stage *Stage) {
	delete(stage.PortShapes, portshape)
	// issue1150
	// delete(stage.PortShape_stagedOrder, portshape)
	delete(stage.PortShapes_mapString, portshape.Name)
}

// commit portshape to the back repo (if it is already staged)
func (portshape *PortShape) Commit(stage *Stage) *PortShape {
	if _, ok := stage.PortShapes[portshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitPortShape(portshape)
		}
	}
	return portshape
}

func (portshape *PortShape) CommitVoid(stage *Stage) {
	portshape.Commit(stage)
}

func (portshape *PortShape) StageVoid(stage *Stage) {
	portshape.Stage(stage)
}

// Checkout portshape to the back repo (if it is already staged)
func (portshape *PortShape) Checkout(stage *Stage) *PortShape {
	if _, ok := stage.PortShapes[portshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutPortShape(portshape)
		}
	}
	return portshape
}

// for satisfaction of GongStruct interface
func (portshape *PortShape) GetName() (res string) {
	return portshape.Name
}

// for satisfaction of GongStruct interface
func (portshape *PortShape) SetName(name string) {
	portshape.Name = name
}

// Stage puts resource to the model stage
func (resource *Resource) Stage(stage *Stage) *Resource {
	if _, ok := stage.Resources[resource]; !ok {
		stage.Resources[resource] = struct{}{}
		stage.Resource_stagedOrder[resource] = stage.ResourceOrder
		stage.Resource_orderStaged[stage.ResourceOrder] = resource
		stage.ResourceOrder++
	}
	stage.Resources_mapString[resource.Name] = resource

	return resource
}

// StagePreserveOrder puts resource to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ResourceOrder
// - update stage.ResourceOrder accordingly
func (resource *Resource) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Resources[resource]; !ok {
		stage.Resources[resource] = struct{}{}

		if order > stage.ResourceOrder {
			stage.ResourceOrder = order
		}
		stage.Resource_stagedOrder[resource] = order
		stage.Resource_orderStaged[order] = resource
		stage.ResourceOrder++
	}
	stage.Resources_mapString[resource.Name] = resource
}

// Unstage removes resource off the model stage
func (resource *Resource) Unstage(stage *Stage) *Resource {
	delete(stage.Resources, resource)
	// issue1150
	// delete(stage.Resource_stagedOrder, resource)
	delete(stage.Resources_mapString, resource.Name)

	return resource
}

// UnstageVoid removes resource off the model stage
func (resource *Resource) UnstageVoid(stage *Stage) {
	delete(stage.Resources, resource)
	// issue1150
	// delete(stage.Resource_stagedOrder, resource)
	delete(stage.Resources_mapString, resource.Name)
}

// commit resource to the back repo (if it is already staged)
func (resource *Resource) Commit(stage *Stage) *Resource {
	if _, ok := stage.Resources[resource]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitResource(resource)
		}
	}
	return resource
}

func (resource *Resource) CommitVoid(stage *Stage) {
	resource.Commit(stage)
}

func (resource *Resource) StageVoid(stage *Stage) {
	resource.Stage(stage)
}

// Checkout resource to the back repo (if it is already staged)
func (resource *Resource) Checkout(stage *Stage) *Resource {
	if _, ok := stage.Resources[resource]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutResource(resource)
		}
	}
	return resource
}

// for satisfaction of GongStruct interface
func (resource *Resource) GetName() (res string) {
	return resource.Name
}

// for satisfaction of GongStruct interface
func (resource *Resource) SetName(name string) {
	resource.Name = name
}

// Stage puts system to the model stage
func (system *System) Stage(stage *Stage) *System {
	if _, ok := stage.Systems[system]; !ok {
		stage.Systems[system] = struct{}{}
		stage.System_stagedOrder[system] = stage.SystemOrder
		stage.System_orderStaged[stage.SystemOrder] = system
		stage.SystemOrder++
	}
	stage.Systems_mapString[system.Name] = system

	return system
}

// StagePreserveOrder puts system to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.SystemOrder
// - update stage.SystemOrder accordingly
func (system *System) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Systems[system]; !ok {
		stage.Systems[system] = struct{}{}

		if order > stage.SystemOrder {
			stage.SystemOrder = order
		}
		stage.System_stagedOrder[system] = order
		stage.System_orderStaged[order] = system
		stage.SystemOrder++
	}
	stage.Systems_mapString[system.Name] = system
}

// Unstage removes system off the model stage
func (system *System) Unstage(stage *Stage) *System {
	delete(stage.Systems, system)
	// issue1150
	// delete(stage.System_stagedOrder, system)
	delete(stage.Systems_mapString, system.Name)

	return system
}

// UnstageVoid removes system off the model stage
func (system *System) UnstageVoid(stage *Stage) {
	delete(stage.Systems, system)
	// issue1150
	// delete(stage.System_stagedOrder, system)
	delete(stage.Systems_mapString, system.Name)
}

// commit system to the back repo (if it is already staged)
func (system *System) Commit(stage *Stage) *System {
	if _, ok := stage.Systems[system]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSystem(system)
		}
	}
	return system
}

func (system *System) CommitVoid(stage *Stage) {
	system.Commit(stage)
}

func (system *System) StageVoid(stage *Stage) {
	system.Stage(stage)
}

// Checkout system to the back repo (if it is already staged)
func (system *System) Checkout(stage *Stage) *System {
	if _, ok := stage.Systems[system]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSystem(system)
		}
	}
	return system
}

// for satisfaction of GongStruct interface
func (system *System) GetName() (res string) {
	return system.Name
}

// for satisfaction of GongStruct interface
func (system *System) SetName(name string) {
	system.Name = name
}

// Stage puts systemshape to the model stage
func (systemshape *SystemShape) Stage(stage *Stage) *SystemShape {
	if _, ok := stage.SystemShapes[systemshape]; !ok {
		stage.SystemShapes[systemshape] = struct{}{}
		stage.SystemShape_stagedOrder[systemshape] = stage.SystemShapeOrder
		stage.SystemShape_orderStaged[stage.SystemShapeOrder] = systemshape
		stage.SystemShapeOrder++
	}
	stage.SystemShapes_mapString[systemshape.Name] = systemshape

	return systemshape
}

// StagePreserveOrder puts systemshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.SystemShapeOrder
// - update stage.SystemShapeOrder accordingly
func (systemshape *SystemShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.SystemShapes[systemshape]; !ok {
		stage.SystemShapes[systemshape] = struct{}{}

		if order > stage.SystemShapeOrder {
			stage.SystemShapeOrder = order
		}
		stage.SystemShape_stagedOrder[systemshape] = order
		stage.SystemShape_orderStaged[order] = systemshape
		stage.SystemShapeOrder++
	}
	stage.SystemShapes_mapString[systemshape.Name] = systemshape
}

// Unstage removes systemshape off the model stage
func (systemshape *SystemShape) Unstage(stage *Stage) *SystemShape {
	delete(stage.SystemShapes, systemshape)
	// issue1150
	// delete(stage.SystemShape_stagedOrder, systemshape)
	delete(stage.SystemShapes_mapString, systemshape.Name)

	return systemshape
}

// UnstageVoid removes systemshape off the model stage
func (systemshape *SystemShape) UnstageVoid(stage *Stage) {
	delete(stage.SystemShapes, systemshape)
	// issue1150
	// delete(stage.SystemShape_stagedOrder, systemshape)
	delete(stage.SystemShapes_mapString, systemshape.Name)
}

// commit systemshape to the back repo (if it is already staged)
func (systemshape *SystemShape) Commit(stage *Stage) *SystemShape {
	if _, ok := stage.SystemShapes[systemshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSystemShape(systemshape)
		}
	}
	return systemshape
}

func (systemshape *SystemShape) CommitVoid(stage *Stage) {
	systemshape.Commit(stage)
}

func (systemshape *SystemShape) StageVoid(stage *Stage) {
	systemshape.Stage(stage)
}

// Checkout systemshape to the back repo (if it is already staged)
func (systemshape *SystemShape) Checkout(stage *Stage) *SystemShape {
	if _, ok := stage.SystemShapes[systemshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSystemShape(systemshape)
		}
	}
	return systemshape
}

// for satisfaction of GongStruct interface
func (systemshape *SystemShape) GetName() (res string) {
	return systemshape.Name
}

// for satisfaction of GongStruct interface
func (systemshape *SystemShape) SetName(name string) {
	systemshape.Name = name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMAllocatedResourceShape(AllocatedResourceShape *AllocatedResourceShape)
	CreateORMAllocatedSystemShape(AllocatedSystemShape *AllocatedSystemShape)
	CreateORMControlFlow(ControlFlow *ControlFlow)
	CreateORMControlFlowShape(ControlFlowShape *ControlFlowShape)
	CreateORMData(Data *Data)
	CreateORMDataFlow(DataFlow *DataFlow)
	CreateORMDataFlowShape(DataFlowShape *DataFlowShape)
	CreateORMDataShape(DataShape *DataShape)
	CreateORMDiagramStructure(DiagramStructure *DiagramStructure)
	CreateORMExternalPartShape(ExternalPartShape *ExternalPartShape)
	CreateORMLibrary(Library *Library)
	CreateORMNote(Note *Note)
	CreateORMNotePortShape(NotePortShape *NotePortShape)
	CreateORMNoteShape(NoteShape *NoteShape)
	CreateORMPart(Part *Part)
	CreateORMPartAnchoredPath(PartAnchoredPath *PartAnchoredPath)
	CreateORMPartShape(PartShape *PartShape)
	CreateORMPort(Port *Port)
	CreateORMPortShape(PortShape *PortShape)
	CreateORMResource(Resource *Resource)
	CreateORMSystem(System *System)
	CreateORMSystemShape(SystemShape *SystemShape)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMAllocatedResourceShape(AllocatedResourceShape *AllocatedResourceShape)
	DeleteORMAllocatedSystemShape(AllocatedSystemShape *AllocatedSystemShape)
	DeleteORMControlFlow(ControlFlow *ControlFlow)
	DeleteORMControlFlowShape(ControlFlowShape *ControlFlowShape)
	DeleteORMData(Data *Data)
	DeleteORMDataFlow(DataFlow *DataFlow)
	DeleteORMDataFlowShape(DataFlowShape *DataFlowShape)
	DeleteORMDataShape(DataShape *DataShape)
	DeleteORMDiagramStructure(DiagramStructure *DiagramStructure)
	DeleteORMExternalPartShape(ExternalPartShape *ExternalPartShape)
	DeleteORMLibrary(Library *Library)
	DeleteORMNote(Note *Note)
	DeleteORMNotePortShape(NotePortShape *NotePortShape)
	DeleteORMNoteShape(NoteShape *NoteShape)
	DeleteORMPart(Part *Part)
	DeleteORMPartAnchoredPath(PartAnchoredPath *PartAnchoredPath)
	DeleteORMPartShape(PartShape *PartShape)
	DeleteORMPort(Port *Port)
	DeleteORMPortShape(PortShape *PortShape)
	DeleteORMResource(Resource *Resource)
	DeleteORMSystem(System *System)
	DeleteORMSystemShape(SystemShape *SystemShape)
}

func (stage *Stage) Reset() { // insertion point for array reset
	stage.AllocatedResourceShapes = make(map[*AllocatedResourceShape]struct{})
	stage.AllocatedResourceShapes_mapString = make(map[string]*AllocatedResourceShape)
	stage.AllocatedResourceShape_stagedOrder = make(map[*AllocatedResourceShape]uint)
	stage.AllocatedResourceShapeOrder = 0

	stage.AllocatedSystemShapes = make(map[*AllocatedSystemShape]struct{})
	stage.AllocatedSystemShapes_mapString = make(map[string]*AllocatedSystemShape)
	stage.AllocatedSystemShape_stagedOrder = make(map[*AllocatedSystemShape]uint)
	stage.AllocatedSystemShapeOrder = 0

	stage.ControlFlows = make(map[*ControlFlow]struct{})
	stage.ControlFlows_mapString = make(map[string]*ControlFlow)
	stage.ControlFlow_stagedOrder = make(map[*ControlFlow]uint)
	stage.ControlFlowOrder = 0

	stage.ControlFlowShapes = make(map[*ControlFlowShape]struct{})
	stage.ControlFlowShapes_mapString = make(map[string]*ControlFlowShape)
	stage.ControlFlowShape_stagedOrder = make(map[*ControlFlowShape]uint)
	stage.ControlFlowShapeOrder = 0

	stage.Datas = make(map[*Data]struct{})
	stage.Datas_mapString = make(map[string]*Data)
	stage.Data_stagedOrder = make(map[*Data]uint)
	stage.DataOrder = 0

	stage.DataFlows = make(map[*DataFlow]struct{})
	stage.DataFlows_mapString = make(map[string]*DataFlow)
	stage.DataFlow_stagedOrder = make(map[*DataFlow]uint)
	stage.DataFlowOrder = 0

	stage.DataFlowShapes = make(map[*DataFlowShape]struct{})
	stage.DataFlowShapes_mapString = make(map[string]*DataFlowShape)
	stage.DataFlowShape_stagedOrder = make(map[*DataFlowShape]uint)
	stage.DataFlowShapeOrder = 0

	stage.DataShapes = make(map[*DataShape]struct{})
	stage.DataShapes_mapString = make(map[string]*DataShape)
	stage.DataShape_stagedOrder = make(map[*DataShape]uint)
	stage.DataShapeOrder = 0

	stage.DiagramStructures = make(map[*DiagramStructure]struct{})
	stage.DiagramStructures_mapString = make(map[string]*DiagramStructure)
	stage.DiagramStructure_stagedOrder = make(map[*DiagramStructure]uint)
	stage.DiagramStructureOrder = 0

	stage.ExternalPartShapes = make(map[*ExternalPartShape]struct{})
	stage.ExternalPartShapes_mapString = make(map[string]*ExternalPartShape)
	stage.ExternalPartShape_stagedOrder = make(map[*ExternalPartShape]uint)
	stage.ExternalPartShapeOrder = 0

	stage.Librarys = make(map[*Library]struct{})
	stage.Librarys_mapString = make(map[string]*Library)
	stage.Library_stagedOrder = make(map[*Library]uint)
	stage.LibraryOrder = 0

	stage.Notes = make(map[*Note]struct{})
	stage.Notes_mapString = make(map[string]*Note)
	stage.Note_stagedOrder = make(map[*Note]uint)
	stage.NoteOrder = 0

	stage.NotePortShapes = make(map[*NotePortShape]struct{})
	stage.NotePortShapes_mapString = make(map[string]*NotePortShape)
	stage.NotePortShape_stagedOrder = make(map[*NotePortShape]uint)
	stage.NotePortShapeOrder = 0

	stage.NoteShapes = make(map[*NoteShape]struct{})
	stage.NoteShapes_mapString = make(map[string]*NoteShape)
	stage.NoteShape_stagedOrder = make(map[*NoteShape]uint)
	stage.NoteShapeOrder = 0

	stage.Parts = make(map[*Part]struct{})
	stage.Parts_mapString = make(map[string]*Part)
	stage.Part_stagedOrder = make(map[*Part]uint)
	stage.PartOrder = 0

	stage.PartAnchoredPaths = make(map[*PartAnchoredPath]struct{})
	stage.PartAnchoredPaths_mapString = make(map[string]*PartAnchoredPath)
	stage.PartAnchoredPath_stagedOrder = make(map[*PartAnchoredPath]uint)
	stage.PartAnchoredPathOrder = 0

	stage.PartShapes = make(map[*PartShape]struct{})
	stage.PartShapes_mapString = make(map[string]*PartShape)
	stage.PartShape_stagedOrder = make(map[*PartShape]uint)
	stage.PartShapeOrder = 0

	stage.Ports = make(map[*Port]struct{})
	stage.Ports_mapString = make(map[string]*Port)
	stage.Port_stagedOrder = make(map[*Port]uint)
	stage.PortOrder = 0

	stage.PortShapes = make(map[*PortShape]struct{})
	stage.PortShapes_mapString = make(map[string]*PortShape)
	stage.PortShape_stagedOrder = make(map[*PortShape]uint)
	stage.PortShapeOrder = 0

	stage.Resources = make(map[*Resource]struct{})
	stage.Resources_mapString = make(map[string]*Resource)
	stage.Resource_stagedOrder = make(map[*Resource]uint)
	stage.ResourceOrder = 0

	stage.Systems = make(map[*System]struct{})
	stage.Systems_mapString = make(map[string]*System)
	stage.System_stagedOrder = make(map[*System]uint)
	stage.SystemOrder = 0

	stage.SystemShapes = make(map[*SystemShape]struct{})
	stage.SystemShapes_mapString = make(map[string]*SystemShape)
	stage.SystemShape_stagedOrder = make(map[*SystemShape]uint)
	stage.SystemShapeOrder = 0

	if stage.GetProbeIF() != nil {
		stage.GetProbeIF().ResetNotifications()
	}
	if stage.IsInDeltaMode() {
		stage.ComputeReferenceAndOrders()
	}
}

func (stage *Stage) Nil() { // insertion point for array nil
	stage.AllocatedResourceShapes = nil
	stage.AllocatedResourceShapes_mapString = nil

	stage.AllocatedSystemShapes = nil
	stage.AllocatedSystemShapes_mapString = nil

	stage.ControlFlows = nil
	stage.ControlFlows_mapString = nil

	stage.ControlFlowShapes = nil
	stage.ControlFlowShapes_mapString = nil

	stage.Datas = nil
	stage.Datas_mapString = nil

	stage.DataFlows = nil
	stage.DataFlows_mapString = nil

	stage.DataFlowShapes = nil
	stage.DataFlowShapes_mapString = nil

	stage.DataShapes = nil
	stage.DataShapes_mapString = nil

	stage.DiagramStructures = nil
	stage.DiagramStructures_mapString = nil

	stage.ExternalPartShapes = nil
	stage.ExternalPartShapes_mapString = nil

	stage.Librarys = nil
	stage.Librarys_mapString = nil

	stage.Notes = nil
	stage.Notes_mapString = nil

	stage.NotePortShapes = nil
	stage.NotePortShapes_mapString = nil

	stage.NoteShapes = nil
	stage.NoteShapes_mapString = nil

	stage.Parts = nil
	stage.Parts_mapString = nil

	stage.PartAnchoredPaths = nil
	stage.PartAnchoredPaths_mapString = nil

	stage.PartShapes = nil
	stage.PartShapes_mapString = nil

	stage.Ports = nil
	stage.Ports_mapString = nil

	stage.PortShapes = nil
	stage.PortShapes_mapString = nil

	stage.Resources = nil
	stage.Resources_mapString = nil

	stage.Systems = nil
	stage.Systems_mapString = nil

	stage.SystemShapes = nil
	stage.SystemShapes_mapString = nil

	// end of insertion point for array nil
}

func (stage *Stage) Unstage() { // insertion point for array nil
	for allocatedresourceshape := range stage.AllocatedResourceShapes {
		allocatedresourceshape.Unstage(stage)
	}

	for allocatedsystemshape := range stage.AllocatedSystemShapes {
		allocatedsystemshape.Unstage(stage)
	}

	for controlflow := range stage.ControlFlows {
		controlflow.Unstage(stage)
	}

	for controlflowshape := range stage.ControlFlowShapes {
		controlflowshape.Unstage(stage)
	}

	for data := range stage.Datas {
		data.Unstage(stage)
	}

	for dataflow := range stage.DataFlows {
		dataflow.Unstage(stage)
	}

	for dataflowshape := range stage.DataFlowShapes {
		dataflowshape.Unstage(stage)
	}

	for datashape := range stage.DataShapes {
		datashape.Unstage(stage)
	}

	for diagramstructure := range stage.DiagramStructures {
		diagramstructure.Unstage(stage)
	}

	for externalpartshape := range stage.ExternalPartShapes {
		externalpartshape.Unstage(stage)
	}

	for library := range stage.Librarys {
		library.Unstage(stage)
	}

	for note := range stage.Notes {
		note.Unstage(stage)
	}

	for noteportshape := range stage.NotePortShapes {
		noteportshape.Unstage(stage)
	}

	for noteshape := range stage.NoteShapes {
		noteshape.Unstage(stage)
	}

	for part := range stage.Parts {
		part.Unstage(stage)
	}

	for partanchoredpath := range stage.PartAnchoredPaths {
		partanchoredpath.Unstage(stage)
	}

	for partshape := range stage.PartShapes {
		partshape.Unstage(stage)
	}

	for port := range stage.Ports {
		port.Unstage(stage)
	}

	for portshape := range stage.PortShapes {
		portshape.Unstage(stage)
	}

	for resource := range stage.Resources {
		resource.Unstage(stage)
	}

	for system := range stage.Systems {
		system.Unstage(stage)
	}

	for systemshape := range stage.SystemShapes {
		systemshape.Unstage(stage)
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
	case map[*AllocatedResourceShape]any:
		return any(&stage.AllocatedResourceShapes).(*Type)
	case map[*AllocatedSystemShape]any:
		return any(&stage.AllocatedSystemShapes).(*Type)
	case map[*ControlFlow]any:
		return any(&stage.ControlFlows).(*Type)
	case map[*ControlFlowShape]any:
		return any(&stage.ControlFlowShapes).(*Type)
	case map[*Data]any:
		return any(&stage.Datas).(*Type)
	case map[*DataFlow]any:
		return any(&stage.DataFlows).(*Type)
	case map[*DataFlowShape]any:
		return any(&stage.DataFlowShapes).(*Type)
	case map[*DataShape]any:
		return any(&stage.DataShapes).(*Type)
	case map[*DiagramStructure]any:
		return any(&stage.DiagramStructures).(*Type)
	case map[*ExternalPartShape]any:
		return any(&stage.ExternalPartShapes).(*Type)
	case map[*Library]any:
		return any(&stage.Librarys).(*Type)
	case map[*Note]any:
		return any(&stage.Notes).(*Type)
	case map[*NotePortShape]any:
		return any(&stage.NotePortShapes).(*Type)
	case map[*NoteShape]any:
		return any(&stage.NoteShapes).(*Type)
	case map[*Part]any:
		return any(&stage.Parts).(*Type)
	case map[*PartAnchoredPath]any:
		return any(&stage.PartAnchoredPaths).(*Type)
	case map[*PartShape]any:
		return any(&stage.PartShapes).(*Type)
	case map[*Port]any:
		return any(&stage.Ports).(*Type)
	case map[*PortShape]any:
		return any(&stage.PortShapes).(*Type)
	case map[*Resource]any:
		return any(&stage.Resources).(*Type)
	case map[*System]any:
		return any(&stage.Systems).(*Type)
	case map[*SystemShape]any:
		return any(&stage.SystemShapes).(*Type)
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
	case *AllocatedResourceShape:
		return any(stage.AllocatedResourceShapes_mapString).(map[string]Type)
	case *AllocatedSystemShape:
		return any(stage.AllocatedSystemShapes_mapString).(map[string]Type)
	case *ControlFlow:
		return any(stage.ControlFlows_mapString).(map[string]Type)
	case *ControlFlowShape:
		return any(stage.ControlFlowShapes_mapString).(map[string]Type)
	case *Data:
		return any(stage.Datas_mapString).(map[string]Type)
	case *DataFlow:
		return any(stage.DataFlows_mapString).(map[string]Type)
	case *DataFlowShape:
		return any(stage.DataFlowShapes_mapString).(map[string]Type)
	case *DataShape:
		return any(stage.DataShapes_mapString).(map[string]Type)
	case *DiagramStructure:
		return any(stage.DiagramStructures_mapString).(map[string]Type)
	case *ExternalPartShape:
		return any(stage.ExternalPartShapes_mapString).(map[string]Type)
	case *Library:
		return any(stage.Librarys_mapString).(map[string]Type)
	case *Note:
		return any(stage.Notes_mapString).(map[string]Type)
	case *NotePortShape:
		return any(stage.NotePortShapes_mapString).(map[string]Type)
	case *NoteShape:
		return any(stage.NoteShapes_mapString).(map[string]Type)
	case *Part:
		return any(stage.Parts_mapString).(map[string]Type)
	case *PartAnchoredPath:
		return any(stage.PartAnchoredPaths_mapString).(map[string]Type)
	case *PartShape:
		return any(stage.PartShapes_mapString).(map[string]Type)
	case *Port:
		return any(stage.Ports_mapString).(map[string]Type)
	case *PortShape:
		return any(stage.PortShapes_mapString).(map[string]Type)
	case *Resource:
		return any(stage.Resources_mapString).(map[string]Type)
	case *System:
		return any(stage.Systems_mapString).(map[string]Type)
	case *SystemShape:
		return any(stage.SystemShapes_mapString).(map[string]Type)
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
	case AllocatedResourceShape:
		return any(&stage.AllocatedResourceShapes).(*map[*Type]struct{})
	case AllocatedSystemShape:
		return any(&stage.AllocatedSystemShapes).(*map[*Type]struct{})
	case ControlFlow:
		return any(&stage.ControlFlows).(*map[*Type]struct{})
	case ControlFlowShape:
		return any(&stage.ControlFlowShapes).(*map[*Type]struct{})
	case Data:
		return any(&stage.Datas).(*map[*Type]struct{})
	case DataFlow:
		return any(&stage.DataFlows).(*map[*Type]struct{})
	case DataFlowShape:
		return any(&stage.DataFlowShapes).(*map[*Type]struct{})
	case DataShape:
		return any(&stage.DataShapes).(*map[*Type]struct{})
	case DiagramStructure:
		return any(&stage.DiagramStructures).(*map[*Type]struct{})
	case ExternalPartShape:
		return any(&stage.ExternalPartShapes).(*map[*Type]struct{})
	case Library:
		return any(&stage.Librarys).(*map[*Type]struct{})
	case Note:
		return any(&stage.Notes).(*map[*Type]struct{})
	case NotePortShape:
		return any(&stage.NotePortShapes).(*map[*Type]struct{})
	case NoteShape:
		return any(&stage.NoteShapes).(*map[*Type]struct{})
	case Part:
		return any(&stage.Parts).(*map[*Type]struct{})
	case PartAnchoredPath:
		return any(&stage.PartAnchoredPaths).(*map[*Type]struct{})
	case PartShape:
		return any(&stage.PartShapes).(*map[*Type]struct{})
	case Port:
		return any(&stage.Ports).(*map[*Type]struct{})
	case PortShape:
		return any(&stage.PortShapes).(*map[*Type]struct{})
	case Resource:
		return any(&stage.Resources).(*map[*Type]struct{})
	case System:
		return any(&stage.Systems).(*map[*Type]struct{})
	case SystemShape:
		return any(&stage.SystemShapes).(*map[*Type]struct{})
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
	case *AllocatedResourceShape:
		return any(&stage.AllocatedResourceShapes).(*map[Type]struct{})
	case *AllocatedSystemShape:
		return any(&stage.AllocatedSystemShapes).(*map[Type]struct{})
	case *ControlFlow:
		return any(&stage.ControlFlows).(*map[Type]struct{})
	case *ControlFlowShape:
		return any(&stage.ControlFlowShapes).(*map[Type]struct{})
	case *Data:
		return any(&stage.Datas).(*map[Type]struct{})
	case *DataFlow:
		return any(&stage.DataFlows).(*map[Type]struct{})
	case *DataFlowShape:
		return any(&stage.DataFlowShapes).(*map[Type]struct{})
	case *DataShape:
		return any(&stage.DataShapes).(*map[Type]struct{})
	case *DiagramStructure:
		return any(&stage.DiagramStructures).(*map[Type]struct{})
	case *ExternalPartShape:
		return any(&stage.ExternalPartShapes).(*map[Type]struct{})
	case *Library:
		return any(&stage.Librarys).(*map[Type]struct{})
	case *Note:
		return any(&stage.Notes).(*map[Type]struct{})
	case *NotePortShape:
		return any(&stage.NotePortShapes).(*map[Type]struct{})
	case *NoteShape:
		return any(&stage.NoteShapes).(*map[Type]struct{})
	case *Part:
		return any(&stage.Parts).(*map[Type]struct{})
	case *PartAnchoredPath:
		return any(&stage.PartAnchoredPaths).(*map[Type]struct{})
	case *PartShape:
		return any(&stage.PartShapes).(*map[Type]struct{})
	case *Port:
		return any(&stage.Ports).(*map[Type]struct{})
	case *PortShape:
		return any(&stage.PortShapes).(*map[Type]struct{})
	case *Resource:
		return any(&stage.Resources).(*map[Type]struct{})
	case *System:
		return any(&stage.Systems).(*map[Type]struct{})
	case *SystemShape:
		return any(&stage.SystemShapes).(*map[Type]struct{})
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
	case AllocatedResourceShape:
		return any(&stage.AllocatedResourceShapes_mapString).(*map[string]*Type)
	case AllocatedSystemShape:
		return any(&stage.AllocatedSystemShapes_mapString).(*map[string]*Type)
	case ControlFlow:
		return any(&stage.ControlFlows_mapString).(*map[string]*Type)
	case ControlFlowShape:
		return any(&stage.ControlFlowShapes_mapString).(*map[string]*Type)
	case Data:
		return any(&stage.Datas_mapString).(*map[string]*Type)
	case DataFlow:
		return any(&stage.DataFlows_mapString).(*map[string]*Type)
	case DataFlowShape:
		return any(&stage.DataFlowShapes_mapString).(*map[string]*Type)
	case DataShape:
		return any(&stage.DataShapes_mapString).(*map[string]*Type)
	case DiagramStructure:
		return any(&stage.DiagramStructures_mapString).(*map[string]*Type)
	case ExternalPartShape:
		return any(&stage.ExternalPartShapes_mapString).(*map[string]*Type)
	case Library:
		return any(&stage.Librarys_mapString).(*map[string]*Type)
	case Note:
		return any(&stage.Notes_mapString).(*map[string]*Type)
	case NotePortShape:
		return any(&stage.NotePortShapes_mapString).(*map[string]*Type)
	case NoteShape:
		return any(&stage.NoteShapes_mapString).(*map[string]*Type)
	case Part:
		return any(&stage.Parts_mapString).(*map[string]*Type)
	case PartAnchoredPath:
		return any(&stage.PartAnchoredPaths_mapString).(*map[string]*Type)
	case PartShape:
		return any(&stage.PartShapes_mapString).(*map[string]*Type)
	case Port:
		return any(&stage.Ports_mapString).(*map[string]*Type)
	case PortShape:
		return any(&stage.PortShapes_mapString).(*map[string]*Type)
	case Resource:
		return any(&stage.Resources_mapString).(*map[string]*Type)
	case System:
		return any(&stage.Systems_mapString).(*map[string]*Type)
	case SystemShape:
		return any(&stage.SystemShapes_mapString).(*map[string]*Type)
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
	case AllocatedResourceShape:
		return any(&AllocatedResourceShape{
			// Initialisation of associations
			// field is initialized with an instance of Part with the name of the field
			Part: &Part{Name: "Part"},
			// field is initialized with an instance of Resource with the name of the field
			Resource: &Resource{Name: "Resource"},
		}).(*Type)
	case AllocatedSystemShape:
		return any(&AllocatedSystemShape{
			// Initialisation of associations
			// field is initialized with an instance of Part with the name of the field
			Part: &Part{Name: "Part"},
			// field is initialized with an instance of System with the name of the field
			System: &System{Name: "System"},
		}).(*Type)
	case ControlFlow:
		return any(&ControlFlow{
			// Initialisation of associations
			// field is initialized with an instance of Port with the name of the field
			Start: &Port{Name: "Start"},
			// field is initialized with an instance of Port with the name of the field
			End: &Port{Name: "End"},
		}).(*Type)
	case ControlFlowShape:
		return any(&ControlFlowShape{
			// Initialisation of associations
			// field is initialized with an instance of ControlFlow with the name of the field
			ControlFlow: &ControlFlow{Name: "ControlFlow"},
		}).(*Type)
	case Data:
		return any(&Data{
			// Initialisation of associations
		}).(*Type)
	case DataFlow:
		return any(&DataFlow{
			// Initialisation of associations
			// field is initialized with an instance of Port with the name of the field
			StartPort: &Port{Name: "StartPort"},
			// field is initialized with an instance of Port with the name of the field
			EndPort: &Port{Name: "EndPort"},
			// field is initialized with an instance of Part with the name of the field
			StartExternalPart: &Part{Name: "StartExternalPart"},
			// field is initialized with an instance of Part with the name of the field
			EndExternalPart: &Part{Name: "EndExternalPart"},
			// field is initialized with an instance of Data with the name of the field
			Datas: []*Data{{Name: "Datas"}},
		}).(*Type)
	case DataFlowShape:
		return any(&DataFlowShape{
			// Initialisation of associations
			// field is initialized with an instance of DataFlow with the name of the field
			DataFlow: &DataFlow{Name: "DataFlow"},
		}).(*Type)
	case DataShape:
		return any(&DataShape{
			// Initialisation of associations
			// field is initialized with an instance of Data with the name of the field
			Data: &Data{Name: "Data"},
			// field is initialized with an instance of DataFlow with the name of the field
			DataFlow: &DataFlow{Name: "DataFlow"},
		}).(*Type)
	case DiagramStructure:
		return any(&DiagramStructure{
			// Initialisation of associations
			// field is initialized with an instance of SystemShape with the name of the field
			System_Shapes: []*SystemShape{{Name: "System_Shapes"}},
			// field is initialized with an instance of System with the name of the field
			SystemsWhoseNodeIsExpanded: []*System{{Name: "SystemsWhoseNodeIsExpanded"}},
			// field is initialized with an instance of PartShape with the name of the field
			Part_Shapes: []*PartShape{{Name: "Part_Shapes"}},
			// field is initialized with an instance of Part with the name of the field
			PartWhoseNodeIsExpanded: []*Part{{Name: "PartWhoseNodeIsExpanded"}},
			// field is initialized with an instance of ExternalPartShape with the name of the field
			ExternalPart_Shapes: []*ExternalPartShape{{Name: "ExternalPart_Shapes"}},
			// field is initialized with an instance of Part with the name of the field
			ExternalPartWhoseNodeIsExpanded: []*Part{{Name: "ExternalPartWhoseNodeIsExpanded"}},
			// field is initialized with an instance of Part with the name of the field
			ExternalPartsWhoseOutDataFlowsNodeIsExpanded: []*Part{{Name: "ExternalPartsWhoseOutDataFlowsNodeIsExpanded"}},
			// field is initialized with an instance of Part with the name of the field
			ExternalPartsWhoseInDataFlowsNodeIsExpanded: []*Part{{Name: "ExternalPartsWhoseInDataFlowsNodeIsExpanded"}},
			// field is initialized with an instance of Port with the name of the field
			PortsWhoseNodeIsExpanded: []*Port{{Name: "PortsWhoseNodeIsExpanded"}},
			// field is initialized with an instance of PortShape with the name of the field
			Port_Shapes: []*PortShape{{Name: "Port_Shapes"}},
			// field is initialized with an instance of ControlFlow with the name of the field
			ControlFlowsWhoseNodeIsExpanded: []*ControlFlow{{Name: "ControlFlowsWhoseNodeIsExpanded"}},
			// field is initialized with an instance of ControlFlowShape with the name of the field
			ControlFlow_Shapes: []*ControlFlowShape{{Name: "ControlFlow_Shapes"}},
			// field is initialized with an instance of DataFlow with the name of the field
			DataFlowsWhoseNodeIsExpanded: []*DataFlow{{Name: "DataFlowsWhoseNodeIsExpanded"}},
			// field is initialized with an instance of DataFlowShape with the name of the field
			DataFlow_Shapes: []*DataFlowShape{{Name: "DataFlow_Shapes"}},
			// field is initialized with an instance of Data with the name of the field
			DatasWhoseNodeIsExpanded: []*Data{{Name: "DatasWhoseNodeIsExpanded"}},
			// field is initialized with an instance of DataShape with the name of the field
			Data_Shapes: []*DataShape{{Name: "Data_Shapes"}},
			// field is initialized with an instance of DataFlow with the name of the field
			DataFlowsWhoseDataNodeIsExpanded: []*DataFlow{{Name: "DataFlowsWhoseDataNodeIsExpanded"}},
			// field is initialized with an instance of Resource with the name of the field
			AllocatedResourcesWhoseNodeIsExpanded: []*Resource{{Name: "AllocatedResourcesWhoseNodeIsExpanded"}},
			// field is initialized with an instance of AllocatedResourceShape with the name of the field
			AllocatedResourceShapes: []*AllocatedResourceShape{{Name: "AllocatedResourceShapes"}},
			// field is initialized with an instance of System with the name of the field
			AllocatedSystemesWhoseNodeIsExpanded: []*System{{Name: "AllocatedSystemesWhoseNodeIsExpanded"}},
			// field is initialized with an instance of AllocatedSystemShape with the name of the field
			AllocatedSystemShapes: []*AllocatedSystemShape{{Name: "AllocatedSystemShapes"}},
			// field is initialized with an instance of NoteShape with the name of the field
			Note_Shapes: []*NoteShape{{Name: "Note_Shapes"}},
			// field is initialized with an instance of Note with the name of the field
			NotesWhoseNodeIsExpanded: []*Note{{Name: "NotesWhoseNodeIsExpanded"}},
			// field is initialized with an instance of NotePortShape with the name of the field
			NotePortShapes: []*NotePortShape{{Name: "NotePortShapes"}},
		}).(*Type)
	case ExternalPartShape:
		return any(&ExternalPartShape{
			// Initialisation of associations
			// field is initialized with an instance of Part with the name of the field
			Part: &Part{Name: "Part"},
		}).(*Type)
	case Library:
		return any(&Library{
			// Initialisation of associations
			// field is initialized with an instance of Library with the name of the field
			SubLibraries: []*Library{{Name: "SubLibraries"}},
			// field is initialized with an instance of Library with the name of the field
			SubLibrariesWhoseNodeIsExpanded: []*Library{{Name: "SubLibrariesWhoseNodeIsExpanded"}},
			// field is initialized with an instance of System with the name of the field
			RootSystemes: []*System{{Name: "RootSystemes"}},
			// field is initialized with an instance of System with the name of the field
			SystemsWhoseNodeIsExpanded: []*System{{Name: "SystemsWhoseNodeIsExpanded"}},
			// field is initialized with an instance of DataFlow with the name of the field
			RootDataFlows: []*DataFlow{{Name: "RootDataFlows"}},
			// field is initialized with an instance of DataFlow with the name of the field
			DataFlowsWhoseNodeIsExpanded: []*DataFlow{{Name: "DataFlowsWhoseNodeIsExpanded"}},
			// field is initialized with an instance of Data with the name of the field
			RootDatas: []*Data{{Name: "RootDatas"}},
			// field is initialized with an instance of Data with the name of the field
			DatasWhoseNodeIsExpanded: []*Data{{Name: "DatasWhoseNodeIsExpanded"}},
			// field is initialized with an instance of Resource with the name of the field
			RootResources: []*Resource{{Name: "RootResources"}},
			// field is initialized with an instance of Resource with the name of the field
			ResourcesWhoseNodeIsExpanded: []*Resource{{Name: "ResourcesWhoseNodeIsExpanded"}},
			// field is initialized with an instance of Part with the name of the field
			PartsWhoseNodeIsExpanded: []*Part{{Name: "PartsWhoseNodeIsExpanded"}},
			// field is initialized with an instance of Note with the name of the field
			RootNotes: []*Note{{Name: "RootNotes"}},
			// field is initialized with an instance of Note with the name of the field
			NotesWhoseNodeIsExpanded: []*Note{{Name: "NotesWhoseNodeIsExpanded"}},
		}).(*Type)
	case Note:
		return any(&Note{
			// Initialisation of associations
			// field is initialized with an instance of Port with the name of the field
			Ports: []*Port{{Name: "Ports"}},
		}).(*Type)
	case NotePortShape:
		return any(&NotePortShape{
			// Initialisation of associations
			// field is initialized with an instance of Note with the name of the field
			Note: &Note{Name: "Note"},
			// field is initialized with an instance of Port with the name of the field
			Port: &Port{Name: "Port"},
		}).(*Type)
	case NoteShape:
		return any(&NoteShape{
			// Initialisation of associations
			// field is initialized with an instance of Note with the name of the field
			Note: &Note{Name: "Note"},
		}).(*Type)
	case Part:
		return any(&Part{
			// Initialisation of associations
			// field is initialized with an instance of Port with the name of the field
			Ports: []*Port{{Name: "Ports"}},
			// field is initialized with an instance of ControlFlow with the name of the field
			ControlFlows: []*ControlFlow{{Name: "ControlFlows"}},
			// field is initialized with an instance of Port with the name of the field
			PortWhoseOutControlFlowsNodeIsExpanded: []*Port{{Name: "PortWhoseOutControlFlowsNodeIsExpanded"}},
			// field is initialized with an instance of Port with the name of the field
			PortWhoseInControlFlowsNodeIsExpanded: []*Port{{Name: "PortWhoseInControlFlowsNodeIsExpanded"}},
			// field is initialized with an instance of Port with the name of the field
			PortWhoseOutDataFlowsNodeIsExpanded: []*Port{{Name: "PortWhoseOutDataFlowsNodeIsExpanded"}},
			// field is initialized with an instance of Port with the name of the field
			PortWhoseInDataFlowsNodeIsExpanded: []*Port{{Name: "PortWhoseInDataFlowsNodeIsExpanded"}},
			// field is initialized with an instance of PartAnchoredPath with the name of the field
			PartAnchoredPath: []*PartAnchoredPath{{Name: "PartAnchoredPath"}},
		}).(*Type)
	case PartAnchoredPath:
		return any(&PartAnchoredPath{
			// Initialisation of associations
		}).(*Type)
	case PartShape:
		return any(&PartShape{
			// Initialisation of associations
			// field is initialized with an instance of Part with the name of the field
			Part: &Part{Name: "Part"},
		}).(*Type)
	case Port:
		return any(&Port{
			// Initialisation of associations
			// field is initialized with an instance of System with the name of the field
			Type: &System{Name: "Type"},
		}).(*Type)
	case PortShape:
		return any(&PortShape{
			// Initialisation of associations
			// field is initialized with an instance of Port with the name of the field
			Port: &Port{Name: "Port"},
		}).(*Type)
	case Resource:
		return any(&Resource{
			// Initialisation of associations
		}).(*Type)
	case System:
		return any(&System{
			// Initialisation of associations
			// field is initialized with an instance of DiagramStructure with the name of the field
			DiagramStructures: []*DiagramStructure{{Name: "DiagramStructures"}},
			// field is initialized with an instance of DiagramStructure with the name of the field
			DiagramStructureWhoseNodeIsExpanded: []*DiagramStructure{{Name: "DiagramStructureWhoseNodeIsExpanded"}},
			// field is initialized with an instance of System with the name of the field
			SubSystemes: []*System{{Name: "SubSystemes"}},
			// field is initialized with an instance of Part with the name of the field
			Parts: []*Part{{Name: "Parts"}},
			// field is initialized with an instance of Part with the name of the field
			PartWhoseNodeIsExpanded: []*Part{{Name: "PartWhoseNodeIsExpanded"}},
			// field is initialized with an instance of DataFlow with the name of the field
			DataFlows: []*DataFlow{{Name: "DataFlows"}},
			// field is initialized with an instance of Part with the name of the field
			ExternalParts: []*Part{{Name: "ExternalParts"}},
			// field is initialized with an instance of Part with the name of the field
			ExternalPartWhoseNodeIsExpanded: []*Part{{Name: "ExternalPartWhoseNodeIsExpanded"}},
		}).(*Type)
	case SystemShape:
		return any(&SystemShape{
			// Initialisation of associations
			// field is initialized with an instance of System with the name of the field
			System: &System{Name: "System"},
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
	// reverse maps of direct associations of AllocatedResourceShape
	case AllocatedResourceShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Part":
			res := make(map[*Part][]*AllocatedResourceShape)
			for allocatedresourceshape := range stage.AllocatedResourceShapes {
				if allocatedresourceshape.Part != nil {
					part_ := allocatedresourceshape.Part
					var allocatedresourceshapes []*AllocatedResourceShape
					_, ok := res[part_]
					if ok {
						allocatedresourceshapes = res[part_]
					} else {
						allocatedresourceshapes = make([]*AllocatedResourceShape, 0)
					}
					allocatedresourceshapes = append(allocatedresourceshapes, allocatedresourceshape)
					res[part_] = allocatedresourceshapes
				}
			}
			return any(res).(map[*End][]*Start)
		case "Resource":
			res := make(map[*Resource][]*AllocatedResourceShape)
			for allocatedresourceshape := range stage.AllocatedResourceShapes {
				if allocatedresourceshape.Resource != nil {
					resource_ := allocatedresourceshape.Resource
					var allocatedresourceshapes []*AllocatedResourceShape
					_, ok := res[resource_]
					if ok {
						allocatedresourceshapes = res[resource_]
					} else {
						allocatedresourceshapes = make([]*AllocatedResourceShape, 0)
					}
					allocatedresourceshapes = append(allocatedresourceshapes, allocatedresourceshape)
					res[resource_] = allocatedresourceshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of AllocatedSystemShape
	case AllocatedSystemShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Part":
			res := make(map[*Part][]*AllocatedSystemShape)
			for allocatedsystemshape := range stage.AllocatedSystemShapes {
				if allocatedsystemshape.Part != nil {
					part_ := allocatedsystemshape.Part
					var allocatedsystemshapes []*AllocatedSystemShape
					_, ok := res[part_]
					if ok {
						allocatedsystemshapes = res[part_]
					} else {
						allocatedsystemshapes = make([]*AllocatedSystemShape, 0)
					}
					allocatedsystemshapes = append(allocatedsystemshapes, allocatedsystemshape)
					res[part_] = allocatedsystemshapes
				}
			}
			return any(res).(map[*End][]*Start)
		case "System":
			res := make(map[*System][]*AllocatedSystemShape)
			for allocatedsystemshape := range stage.AllocatedSystemShapes {
				if allocatedsystemshape.System != nil {
					system_ := allocatedsystemshape.System
					var allocatedsystemshapes []*AllocatedSystemShape
					_, ok := res[system_]
					if ok {
						allocatedsystemshapes = res[system_]
					} else {
						allocatedsystemshapes = make([]*AllocatedSystemShape, 0)
					}
					allocatedsystemshapes = append(allocatedsystemshapes, allocatedsystemshape)
					res[system_] = allocatedsystemshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ControlFlow
	case ControlFlow:
		switch fieldname {
		// insertion point for per direct association field
		case "Start":
			res := make(map[*Port][]*ControlFlow)
			for controlflow := range stage.ControlFlows {
				if controlflow.Start != nil {
					port_ := controlflow.Start
					var controlflows []*ControlFlow
					_, ok := res[port_]
					if ok {
						controlflows = res[port_]
					} else {
						controlflows = make([]*ControlFlow, 0)
					}
					controlflows = append(controlflows, controlflow)
					res[port_] = controlflows
				}
			}
			return any(res).(map[*End][]*Start)
		case "End":
			res := make(map[*Port][]*ControlFlow)
			for controlflow := range stage.ControlFlows {
				if controlflow.End != nil {
					port_ := controlflow.End
					var controlflows []*ControlFlow
					_, ok := res[port_]
					if ok {
						controlflows = res[port_]
					} else {
						controlflows = make([]*ControlFlow, 0)
					}
					controlflows = append(controlflows, controlflow)
					res[port_] = controlflows
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ControlFlowShape
	case ControlFlowShape:
		switch fieldname {
		// insertion point for per direct association field
		case "ControlFlow":
			res := make(map[*ControlFlow][]*ControlFlowShape)
			for controlflowshape := range stage.ControlFlowShapes {
				if controlflowshape.ControlFlow != nil {
					controlflow_ := controlflowshape.ControlFlow
					var controlflowshapes []*ControlFlowShape
					_, ok := res[controlflow_]
					if ok {
						controlflowshapes = res[controlflow_]
					} else {
						controlflowshapes = make([]*ControlFlowShape, 0)
					}
					controlflowshapes = append(controlflowshapes, controlflowshape)
					res[controlflow_] = controlflowshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Data
	case Data:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DataFlow
	case DataFlow:
		switch fieldname {
		// insertion point for per direct association field
		case "StartPort":
			res := make(map[*Port][]*DataFlow)
			for dataflow := range stage.DataFlows {
				if dataflow.StartPort != nil {
					port_ := dataflow.StartPort
					var dataflows []*DataFlow
					_, ok := res[port_]
					if ok {
						dataflows = res[port_]
					} else {
						dataflows = make([]*DataFlow, 0)
					}
					dataflows = append(dataflows, dataflow)
					res[port_] = dataflows
				}
			}
			return any(res).(map[*End][]*Start)
		case "EndPort":
			res := make(map[*Port][]*DataFlow)
			for dataflow := range stage.DataFlows {
				if dataflow.EndPort != nil {
					port_ := dataflow.EndPort
					var dataflows []*DataFlow
					_, ok := res[port_]
					if ok {
						dataflows = res[port_]
					} else {
						dataflows = make([]*DataFlow, 0)
					}
					dataflows = append(dataflows, dataflow)
					res[port_] = dataflows
				}
			}
			return any(res).(map[*End][]*Start)
		case "StartExternalPart":
			res := make(map[*Part][]*DataFlow)
			for dataflow := range stage.DataFlows {
				if dataflow.StartExternalPart != nil {
					part_ := dataflow.StartExternalPart
					var dataflows []*DataFlow
					_, ok := res[part_]
					if ok {
						dataflows = res[part_]
					} else {
						dataflows = make([]*DataFlow, 0)
					}
					dataflows = append(dataflows, dataflow)
					res[part_] = dataflows
				}
			}
			return any(res).(map[*End][]*Start)
		case "EndExternalPart":
			res := make(map[*Part][]*DataFlow)
			for dataflow := range stage.DataFlows {
				if dataflow.EndExternalPart != nil {
					part_ := dataflow.EndExternalPart
					var dataflows []*DataFlow
					_, ok := res[part_]
					if ok {
						dataflows = res[part_]
					} else {
						dataflows = make([]*DataFlow, 0)
					}
					dataflows = append(dataflows, dataflow)
					res[part_] = dataflows
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of DataFlowShape
	case DataFlowShape:
		switch fieldname {
		// insertion point for per direct association field
		case "DataFlow":
			res := make(map[*DataFlow][]*DataFlowShape)
			for dataflowshape := range stage.DataFlowShapes {
				if dataflowshape.DataFlow != nil {
					dataflow_ := dataflowshape.DataFlow
					var dataflowshapes []*DataFlowShape
					_, ok := res[dataflow_]
					if ok {
						dataflowshapes = res[dataflow_]
					} else {
						dataflowshapes = make([]*DataFlowShape, 0)
					}
					dataflowshapes = append(dataflowshapes, dataflowshape)
					res[dataflow_] = dataflowshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of DataShape
	case DataShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Data":
			res := make(map[*Data][]*DataShape)
			for datashape := range stage.DataShapes {
				if datashape.Data != nil {
					data_ := datashape.Data
					var datashapes []*DataShape
					_, ok := res[data_]
					if ok {
						datashapes = res[data_]
					} else {
						datashapes = make([]*DataShape, 0)
					}
					datashapes = append(datashapes, datashape)
					res[data_] = datashapes
				}
			}
			return any(res).(map[*End][]*Start)
		case "DataFlow":
			res := make(map[*DataFlow][]*DataShape)
			for datashape := range stage.DataShapes {
				if datashape.DataFlow != nil {
					dataflow_ := datashape.DataFlow
					var datashapes []*DataShape
					_, ok := res[dataflow_]
					if ok {
						datashapes = res[dataflow_]
					} else {
						datashapes = make([]*DataShape, 0)
					}
					datashapes = append(datashapes, datashape)
					res[dataflow_] = datashapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of DiagramStructure
	case DiagramStructure:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ExternalPartShape
	case ExternalPartShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Part":
			res := make(map[*Part][]*ExternalPartShape)
			for externalpartshape := range stage.ExternalPartShapes {
				if externalpartshape.Part != nil {
					part_ := externalpartshape.Part
					var externalpartshapes []*ExternalPartShape
					_, ok := res[part_]
					if ok {
						externalpartshapes = res[part_]
					} else {
						externalpartshapes = make([]*ExternalPartShape, 0)
					}
					externalpartshapes = append(externalpartshapes, externalpartshape)
					res[part_] = externalpartshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Library
	case Library:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Note
	case Note:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of NotePortShape
	case NotePortShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Note":
			res := make(map[*Note][]*NotePortShape)
			for noteportshape := range stage.NotePortShapes {
				if noteportshape.Note != nil {
					note_ := noteportshape.Note
					var noteportshapes []*NotePortShape
					_, ok := res[note_]
					if ok {
						noteportshapes = res[note_]
					} else {
						noteportshapes = make([]*NotePortShape, 0)
					}
					noteportshapes = append(noteportshapes, noteportshape)
					res[note_] = noteportshapes
				}
			}
			return any(res).(map[*End][]*Start)
		case "Port":
			res := make(map[*Port][]*NotePortShape)
			for noteportshape := range stage.NotePortShapes {
				if noteportshape.Port != nil {
					port_ := noteportshape.Port
					var noteportshapes []*NotePortShape
					_, ok := res[port_]
					if ok {
						noteportshapes = res[port_]
					} else {
						noteportshapes = make([]*NotePortShape, 0)
					}
					noteportshapes = append(noteportshapes, noteportshape)
					res[port_] = noteportshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of NoteShape
	case NoteShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Note":
			res := make(map[*Note][]*NoteShape)
			for noteshape := range stage.NoteShapes {
				if noteshape.Note != nil {
					note_ := noteshape.Note
					var noteshapes []*NoteShape
					_, ok := res[note_]
					if ok {
						noteshapes = res[note_]
					} else {
						noteshapes = make([]*NoteShape, 0)
					}
					noteshapes = append(noteshapes, noteshape)
					res[note_] = noteshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Part
	case Part:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of PartAnchoredPath
	case PartAnchoredPath:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of PartShape
	case PartShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Part":
			res := make(map[*Part][]*PartShape)
			for partshape := range stage.PartShapes {
				if partshape.Part != nil {
					part_ := partshape.Part
					var partshapes []*PartShape
					_, ok := res[part_]
					if ok {
						partshapes = res[part_]
					} else {
						partshapes = make([]*PartShape, 0)
					}
					partshapes = append(partshapes, partshape)
					res[part_] = partshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Port
	case Port:
		switch fieldname {
		// insertion point for per direct association field
		case "Type":
			res := make(map[*System][]*Port)
			for port := range stage.Ports {
				if port.Type != nil {
					system_ := port.Type
					var ports []*Port
					_, ok := res[system_]
					if ok {
						ports = res[system_]
					} else {
						ports = make([]*Port, 0)
					}
					ports = append(ports, port)
					res[system_] = ports
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of PortShape
	case PortShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Port":
			res := make(map[*Port][]*PortShape)
			for portshape := range stage.PortShapes {
				if portshape.Port != nil {
					port_ := portshape.Port
					var portshapes []*PortShape
					_, ok := res[port_]
					if ok {
						portshapes = res[port_]
					} else {
						portshapes = make([]*PortShape, 0)
					}
					portshapes = append(portshapes, portshape)
					res[port_] = portshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Resource
	case Resource:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of System
	case System:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SystemShape
	case SystemShape:
		switch fieldname {
		// insertion point for per direct association field
		case "System":
			res := make(map[*System][]*SystemShape)
			for systemshape := range stage.SystemShapes {
				if systemshape.System != nil {
					system_ := systemshape.System
					var systemshapes []*SystemShape
					_, ok := res[system_]
					if ok {
						systemshapes = res[system_]
					} else {
						systemshapes = make([]*SystemShape, 0)
					}
					systemshapes = append(systemshapes, systemshape)
					res[system_] = systemshapes
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
	// reverse maps of direct associations of AllocatedResourceShape
	case AllocatedResourceShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of AllocatedSystemShape
	case AllocatedSystemShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ControlFlow
	case ControlFlow:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ControlFlowShape
	case ControlFlowShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Data
	case Data:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DataFlow
	case DataFlow:
		switch fieldname {
		// insertion point for per direct association field
		case "Datas":
			res := make(map[*Data][]*DataFlow)
			for dataflow := range stage.DataFlows {
				for _, data_ := range dataflow.Datas {
					res[data_] = append(res[data_], dataflow)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of DataFlowShape
	case DataFlowShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DataShape
	case DataShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DiagramStructure
	case DiagramStructure:
		switch fieldname {
		// insertion point for per direct association field
		case "System_Shapes":
			res := make(map[*SystemShape][]*DiagramStructure)
			for diagramstructure := range stage.DiagramStructures {
				for _, systemshape_ := range diagramstructure.System_Shapes {
					res[systemshape_] = append(res[systemshape_], diagramstructure)
				}
			}
			return any(res).(map[*End][]*Start)
		case "SystemsWhoseNodeIsExpanded":
			res := make(map[*System][]*DiagramStructure)
			for diagramstructure := range stage.DiagramStructures {
				for _, system_ := range diagramstructure.SystemsWhoseNodeIsExpanded {
					res[system_] = append(res[system_], diagramstructure)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Part_Shapes":
			res := make(map[*PartShape][]*DiagramStructure)
			for diagramstructure := range stage.DiagramStructures {
				for _, partshape_ := range diagramstructure.Part_Shapes {
					res[partshape_] = append(res[partshape_], diagramstructure)
				}
			}
			return any(res).(map[*End][]*Start)
		case "PartWhoseNodeIsExpanded":
			res := make(map[*Part][]*DiagramStructure)
			for diagramstructure := range stage.DiagramStructures {
				for _, part_ := range diagramstructure.PartWhoseNodeIsExpanded {
					res[part_] = append(res[part_], diagramstructure)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ExternalPart_Shapes":
			res := make(map[*ExternalPartShape][]*DiagramStructure)
			for diagramstructure := range stage.DiagramStructures {
				for _, externalpartshape_ := range diagramstructure.ExternalPart_Shapes {
					res[externalpartshape_] = append(res[externalpartshape_], diagramstructure)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ExternalPartWhoseNodeIsExpanded":
			res := make(map[*Part][]*DiagramStructure)
			for diagramstructure := range stage.DiagramStructures {
				for _, part_ := range diagramstructure.ExternalPartWhoseNodeIsExpanded {
					res[part_] = append(res[part_], diagramstructure)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ExternalPartsWhoseOutDataFlowsNodeIsExpanded":
			res := make(map[*Part][]*DiagramStructure)
			for diagramstructure := range stage.DiagramStructures {
				for _, part_ := range diagramstructure.ExternalPartsWhoseOutDataFlowsNodeIsExpanded {
					res[part_] = append(res[part_], diagramstructure)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ExternalPartsWhoseInDataFlowsNodeIsExpanded":
			res := make(map[*Part][]*DiagramStructure)
			for diagramstructure := range stage.DiagramStructures {
				for _, part_ := range diagramstructure.ExternalPartsWhoseInDataFlowsNodeIsExpanded {
					res[part_] = append(res[part_], diagramstructure)
				}
			}
			return any(res).(map[*End][]*Start)
		case "PortsWhoseNodeIsExpanded":
			res := make(map[*Port][]*DiagramStructure)
			for diagramstructure := range stage.DiagramStructures {
				for _, port_ := range diagramstructure.PortsWhoseNodeIsExpanded {
					res[port_] = append(res[port_], diagramstructure)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Port_Shapes":
			res := make(map[*PortShape][]*DiagramStructure)
			for diagramstructure := range stage.DiagramStructures {
				for _, portshape_ := range diagramstructure.Port_Shapes {
					res[portshape_] = append(res[portshape_], diagramstructure)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ControlFlowsWhoseNodeIsExpanded":
			res := make(map[*ControlFlow][]*DiagramStructure)
			for diagramstructure := range stage.DiagramStructures {
				for _, controlflow_ := range diagramstructure.ControlFlowsWhoseNodeIsExpanded {
					res[controlflow_] = append(res[controlflow_], diagramstructure)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ControlFlow_Shapes":
			res := make(map[*ControlFlowShape][]*DiagramStructure)
			for diagramstructure := range stage.DiagramStructures {
				for _, controlflowshape_ := range diagramstructure.ControlFlow_Shapes {
					res[controlflowshape_] = append(res[controlflowshape_], diagramstructure)
				}
			}
			return any(res).(map[*End][]*Start)
		case "DataFlowsWhoseNodeIsExpanded":
			res := make(map[*DataFlow][]*DiagramStructure)
			for diagramstructure := range stage.DiagramStructures {
				for _, dataflow_ := range diagramstructure.DataFlowsWhoseNodeIsExpanded {
					res[dataflow_] = append(res[dataflow_], diagramstructure)
				}
			}
			return any(res).(map[*End][]*Start)
		case "DataFlow_Shapes":
			res := make(map[*DataFlowShape][]*DiagramStructure)
			for diagramstructure := range stage.DiagramStructures {
				for _, dataflowshape_ := range diagramstructure.DataFlow_Shapes {
					res[dataflowshape_] = append(res[dataflowshape_], diagramstructure)
				}
			}
			return any(res).(map[*End][]*Start)
		case "DatasWhoseNodeIsExpanded":
			res := make(map[*Data][]*DiagramStructure)
			for diagramstructure := range stage.DiagramStructures {
				for _, data_ := range diagramstructure.DatasWhoseNodeIsExpanded {
					res[data_] = append(res[data_], diagramstructure)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Data_Shapes":
			res := make(map[*DataShape][]*DiagramStructure)
			for diagramstructure := range stage.DiagramStructures {
				for _, datashape_ := range diagramstructure.Data_Shapes {
					res[datashape_] = append(res[datashape_], diagramstructure)
				}
			}
			return any(res).(map[*End][]*Start)
		case "DataFlowsWhoseDataNodeIsExpanded":
			res := make(map[*DataFlow][]*DiagramStructure)
			for diagramstructure := range stage.DiagramStructures {
				for _, dataflow_ := range diagramstructure.DataFlowsWhoseDataNodeIsExpanded {
					res[dataflow_] = append(res[dataflow_], diagramstructure)
				}
			}
			return any(res).(map[*End][]*Start)
		case "AllocatedResourcesWhoseNodeIsExpanded":
			res := make(map[*Resource][]*DiagramStructure)
			for diagramstructure := range stage.DiagramStructures {
				for _, resource_ := range diagramstructure.AllocatedResourcesWhoseNodeIsExpanded {
					res[resource_] = append(res[resource_], diagramstructure)
				}
			}
			return any(res).(map[*End][]*Start)
		case "AllocatedResourceShapes":
			res := make(map[*AllocatedResourceShape][]*DiagramStructure)
			for diagramstructure := range stage.DiagramStructures {
				for _, allocatedresourceshape_ := range diagramstructure.AllocatedResourceShapes {
					res[allocatedresourceshape_] = append(res[allocatedresourceshape_], diagramstructure)
				}
			}
			return any(res).(map[*End][]*Start)
		case "AllocatedSystemesWhoseNodeIsExpanded":
			res := make(map[*System][]*DiagramStructure)
			for diagramstructure := range stage.DiagramStructures {
				for _, system_ := range diagramstructure.AllocatedSystemesWhoseNodeIsExpanded {
					res[system_] = append(res[system_], diagramstructure)
				}
			}
			return any(res).(map[*End][]*Start)
		case "AllocatedSystemShapes":
			res := make(map[*AllocatedSystemShape][]*DiagramStructure)
			for diagramstructure := range stage.DiagramStructures {
				for _, allocatedsystemshape_ := range diagramstructure.AllocatedSystemShapes {
					res[allocatedsystemshape_] = append(res[allocatedsystemshape_], diagramstructure)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Note_Shapes":
			res := make(map[*NoteShape][]*DiagramStructure)
			for diagramstructure := range stage.DiagramStructures {
				for _, noteshape_ := range diagramstructure.Note_Shapes {
					res[noteshape_] = append(res[noteshape_], diagramstructure)
				}
			}
			return any(res).(map[*End][]*Start)
		case "NotesWhoseNodeIsExpanded":
			res := make(map[*Note][]*DiagramStructure)
			for diagramstructure := range stage.DiagramStructures {
				for _, note_ := range diagramstructure.NotesWhoseNodeIsExpanded {
					res[note_] = append(res[note_], diagramstructure)
				}
			}
			return any(res).(map[*End][]*Start)
		case "NotePortShapes":
			res := make(map[*NotePortShape][]*DiagramStructure)
			for diagramstructure := range stage.DiagramStructures {
				for _, noteportshape_ := range diagramstructure.NotePortShapes {
					res[noteportshape_] = append(res[noteportshape_], diagramstructure)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ExternalPartShape
	case ExternalPartShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Library
	case Library:
		switch fieldname {
		// insertion point for per direct association field
		case "SubLibraries":
			res := make(map[*Library][]*Library)
			for library := range stage.Librarys {
				for _, library_ := range library.SubLibraries {
					res[library_] = append(res[library_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		case "SubLibrariesWhoseNodeIsExpanded":
			res := make(map[*Library][]*Library)
			for library := range stage.Librarys {
				for _, library_ := range library.SubLibrariesWhoseNodeIsExpanded {
					res[library_] = append(res[library_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		case "RootSystemes":
			res := make(map[*System][]*Library)
			for library := range stage.Librarys {
				for _, system_ := range library.RootSystemes {
					res[system_] = append(res[system_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		case "SystemsWhoseNodeIsExpanded":
			res := make(map[*System][]*Library)
			for library := range stage.Librarys {
				for _, system_ := range library.SystemsWhoseNodeIsExpanded {
					res[system_] = append(res[system_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		case "RootDataFlows":
			res := make(map[*DataFlow][]*Library)
			for library := range stage.Librarys {
				for _, dataflow_ := range library.RootDataFlows {
					res[dataflow_] = append(res[dataflow_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		case "DataFlowsWhoseNodeIsExpanded":
			res := make(map[*DataFlow][]*Library)
			for library := range stage.Librarys {
				for _, dataflow_ := range library.DataFlowsWhoseNodeIsExpanded {
					res[dataflow_] = append(res[dataflow_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		case "RootDatas":
			res := make(map[*Data][]*Library)
			for library := range stage.Librarys {
				for _, data_ := range library.RootDatas {
					res[data_] = append(res[data_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		case "DatasWhoseNodeIsExpanded":
			res := make(map[*Data][]*Library)
			for library := range stage.Librarys {
				for _, data_ := range library.DatasWhoseNodeIsExpanded {
					res[data_] = append(res[data_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		case "RootResources":
			res := make(map[*Resource][]*Library)
			for library := range stage.Librarys {
				for _, resource_ := range library.RootResources {
					res[resource_] = append(res[resource_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ResourcesWhoseNodeIsExpanded":
			res := make(map[*Resource][]*Library)
			for library := range stage.Librarys {
				for _, resource_ := range library.ResourcesWhoseNodeIsExpanded {
					res[resource_] = append(res[resource_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		case "PartsWhoseNodeIsExpanded":
			res := make(map[*Part][]*Library)
			for library := range stage.Librarys {
				for _, part_ := range library.PartsWhoseNodeIsExpanded {
					res[part_] = append(res[part_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		case "RootNotes":
			res := make(map[*Note][]*Library)
			for library := range stage.Librarys {
				for _, note_ := range library.RootNotes {
					res[note_] = append(res[note_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		case "NotesWhoseNodeIsExpanded":
			res := make(map[*Note][]*Library)
			for library := range stage.Librarys {
				for _, note_ := range library.NotesWhoseNodeIsExpanded {
					res[note_] = append(res[note_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Note
	case Note:
		switch fieldname {
		// insertion point for per direct association field
		case "Ports":
			res := make(map[*Port][]*Note)
			for note := range stage.Notes {
				for _, port_ := range note.Ports {
					res[port_] = append(res[port_], note)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of NotePortShape
	case NotePortShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of NoteShape
	case NoteShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Part
	case Part:
		switch fieldname {
		// insertion point for per direct association field
		case "Ports":
			res := make(map[*Port][]*Part)
			for part := range stage.Parts {
				for _, port_ := range part.Ports {
					res[port_] = append(res[port_], part)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ControlFlows":
			res := make(map[*ControlFlow][]*Part)
			for part := range stage.Parts {
				for _, controlflow_ := range part.ControlFlows {
					res[controlflow_] = append(res[controlflow_], part)
				}
			}
			return any(res).(map[*End][]*Start)
		case "PortWhoseOutControlFlowsNodeIsExpanded":
			res := make(map[*Port][]*Part)
			for part := range stage.Parts {
				for _, port_ := range part.PortWhoseOutControlFlowsNodeIsExpanded {
					res[port_] = append(res[port_], part)
				}
			}
			return any(res).(map[*End][]*Start)
		case "PortWhoseInControlFlowsNodeIsExpanded":
			res := make(map[*Port][]*Part)
			for part := range stage.Parts {
				for _, port_ := range part.PortWhoseInControlFlowsNodeIsExpanded {
					res[port_] = append(res[port_], part)
				}
			}
			return any(res).(map[*End][]*Start)
		case "PortWhoseOutDataFlowsNodeIsExpanded":
			res := make(map[*Port][]*Part)
			for part := range stage.Parts {
				for _, port_ := range part.PortWhoseOutDataFlowsNodeIsExpanded {
					res[port_] = append(res[port_], part)
				}
			}
			return any(res).(map[*End][]*Start)
		case "PortWhoseInDataFlowsNodeIsExpanded":
			res := make(map[*Port][]*Part)
			for part := range stage.Parts {
				for _, port_ := range part.PortWhoseInDataFlowsNodeIsExpanded {
					res[port_] = append(res[port_], part)
				}
			}
			return any(res).(map[*End][]*Start)
		case "PartAnchoredPath":
			res := make(map[*PartAnchoredPath][]*Part)
			for part := range stage.Parts {
				for _, partanchoredpath_ := range part.PartAnchoredPath {
					res[partanchoredpath_] = append(res[partanchoredpath_], part)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of PartAnchoredPath
	case PartAnchoredPath:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of PartShape
	case PartShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Port
	case Port:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of PortShape
	case PortShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Resource
	case Resource:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of System
	case System:
		switch fieldname {
		// insertion point for per direct association field
		case "DiagramStructures":
			res := make(map[*DiagramStructure][]*System)
			for system := range stage.Systems {
				for _, diagramstructure_ := range system.DiagramStructures {
					res[diagramstructure_] = append(res[diagramstructure_], system)
				}
			}
			return any(res).(map[*End][]*Start)
		case "DiagramStructureWhoseNodeIsExpanded":
			res := make(map[*DiagramStructure][]*System)
			for system := range stage.Systems {
				for _, diagramstructure_ := range system.DiagramStructureWhoseNodeIsExpanded {
					res[diagramstructure_] = append(res[diagramstructure_], system)
				}
			}
			return any(res).(map[*End][]*Start)
		case "SubSystemes":
			res := make(map[*System][]*System)
			for system := range stage.Systems {
				for _, system_ := range system.SubSystemes {
					res[system_] = append(res[system_], system)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Parts":
			res := make(map[*Part][]*System)
			for system := range stage.Systems {
				for _, part_ := range system.Parts {
					res[part_] = append(res[part_], system)
				}
			}
			return any(res).(map[*End][]*Start)
		case "PartWhoseNodeIsExpanded":
			res := make(map[*Part][]*System)
			for system := range stage.Systems {
				for _, part_ := range system.PartWhoseNodeIsExpanded {
					res[part_] = append(res[part_], system)
				}
			}
			return any(res).(map[*End][]*Start)
		case "DataFlows":
			res := make(map[*DataFlow][]*System)
			for system := range stage.Systems {
				for _, dataflow_ := range system.DataFlows {
					res[dataflow_] = append(res[dataflow_], system)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ExternalParts":
			res := make(map[*Part][]*System)
			for system := range stage.Systems {
				for _, part_ := range system.ExternalParts {
					res[part_] = append(res[part_], system)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ExternalPartWhoseNodeIsExpanded":
			res := make(map[*Part][]*System)
			for system := range stage.Systems {
				for _, part_ := range system.ExternalPartWhoseNodeIsExpanded {
					res[part_] = append(res[part_], system)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of SystemShape
	case SystemShape:
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
	case *AllocatedResourceShape:
		res = "AllocatedResourceShape"
	case *AllocatedSystemShape:
		res = "AllocatedSystemShape"
	case *ControlFlow:
		res = "ControlFlow"
	case *ControlFlowShape:
		res = "ControlFlowShape"
	case *Data:
		res = "Data"
	case *DataFlow:
		res = "DataFlow"
	case *DataFlowShape:
		res = "DataFlowShape"
	case *DataShape:
		res = "DataShape"
	case *DiagramStructure:
		res = "DiagramStructure"
	case *ExternalPartShape:
		res = "ExternalPartShape"
	case *Library:
		res = "Library"
	case *Note:
		res = "Note"
	case *NotePortShape:
		res = "NotePortShape"
	case *NoteShape:
		res = "NoteShape"
	case *Part:
		res = "Part"
	case *PartAnchoredPath:
		res = "PartAnchoredPath"
	case *PartShape:
		res = "PartShape"
	case *Port:
		res = "Port"
	case *PortShape:
		res = "PortShape"
	case *Resource:
		res = "Resource"
	case *System:
		res = "System"
	case *SystemShape:
		res = "SystemShape"
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
	case *AllocatedResourceShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramStructure"
		rf.Fieldname = "AllocatedResourceShapes"
		res = append(res, rf)
	case *AllocatedSystemShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramStructure"
		rf.Fieldname = "AllocatedSystemShapes"
		res = append(res, rf)
	case *ControlFlow:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramStructure"
		rf.Fieldname = "ControlFlowsWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Part"
		rf.Fieldname = "ControlFlows"
		res = append(res, rf)
	case *ControlFlowShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramStructure"
		rf.Fieldname = "ControlFlow_Shapes"
		res = append(res, rf)
	case *Data:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DataFlow"
		rf.Fieldname = "Datas"
		res = append(res, rf)
		rf.GongstructName = "DiagramStructure"
		rf.Fieldname = "DatasWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "RootDatas"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "DatasWhoseNodeIsExpanded"
		res = append(res, rf)
	case *DataFlow:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramStructure"
		rf.Fieldname = "DataFlowsWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "DiagramStructure"
		rf.Fieldname = "DataFlowsWhoseDataNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "RootDataFlows"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "DataFlowsWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "System"
		rf.Fieldname = "DataFlows"
		res = append(res, rf)
	case *DataFlowShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramStructure"
		rf.Fieldname = "DataFlow_Shapes"
		res = append(res, rf)
	case *DataShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramStructure"
		rf.Fieldname = "Data_Shapes"
		res = append(res, rf)
	case *DiagramStructure:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "System"
		rf.Fieldname = "DiagramStructures"
		res = append(res, rf)
		rf.GongstructName = "System"
		rf.Fieldname = "DiagramStructureWhoseNodeIsExpanded"
		res = append(res, rf)
	case *ExternalPartShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramStructure"
		rf.Fieldname = "ExternalPart_Shapes"
		res = append(res, rf)
	case *Library:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Library"
		rf.Fieldname = "SubLibraries"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "SubLibrariesWhoseNodeIsExpanded"
		res = append(res, rf)
	case *Note:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramStructure"
		rf.Fieldname = "NotesWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "RootNotes"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "NotesWhoseNodeIsExpanded"
		res = append(res, rf)
	case *NotePortShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramStructure"
		rf.Fieldname = "NotePortShapes"
		res = append(res, rf)
	case *NoteShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramStructure"
		rf.Fieldname = "Note_Shapes"
		res = append(res, rf)
	case *Part:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramStructure"
		rf.Fieldname = "PartWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "DiagramStructure"
		rf.Fieldname = "ExternalPartWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "DiagramStructure"
		rf.Fieldname = "ExternalPartsWhoseOutDataFlowsNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "DiagramStructure"
		rf.Fieldname = "ExternalPartsWhoseInDataFlowsNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "PartsWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "System"
		rf.Fieldname = "Parts"
		res = append(res, rf)
		rf.GongstructName = "System"
		rf.Fieldname = "PartWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "System"
		rf.Fieldname = "ExternalParts"
		res = append(res, rf)
		rf.GongstructName = "System"
		rf.Fieldname = "ExternalPartWhoseNodeIsExpanded"
		res = append(res, rf)
	case *PartAnchoredPath:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Part"
		rf.Fieldname = "PartAnchoredPath"
		res = append(res, rf)
	case *PartShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramStructure"
		rf.Fieldname = "Part_Shapes"
		res = append(res, rf)
	case *Port:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramStructure"
		rf.Fieldname = "PortsWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Note"
		rf.Fieldname = "Ports"
		res = append(res, rf)
		rf.GongstructName = "Part"
		rf.Fieldname = "Ports"
		res = append(res, rf)
		rf.GongstructName = "Part"
		rf.Fieldname = "PortWhoseOutControlFlowsNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Part"
		rf.Fieldname = "PortWhoseInControlFlowsNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Part"
		rf.Fieldname = "PortWhoseOutDataFlowsNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Part"
		rf.Fieldname = "PortWhoseInDataFlowsNodeIsExpanded"
		res = append(res, rf)
	case *PortShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramStructure"
		rf.Fieldname = "Port_Shapes"
		res = append(res, rf)
	case *Resource:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramStructure"
		rf.Fieldname = "AllocatedResourcesWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "RootResources"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "ResourcesWhoseNodeIsExpanded"
		res = append(res, rf)
	case *System:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramStructure"
		rf.Fieldname = "SystemsWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "DiagramStructure"
		rf.Fieldname = "AllocatedSystemesWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "RootSystemes"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "SystemsWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "System"
		rf.Fieldname = "SubSystemes"
		res = append(res, rf)
	case *SystemShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramStructure"
		rf.Fieldname = "System_Shapes"
		res = append(res, rf)
	}
	return
}

// insertion point for get fields header method
func (allocatedresourceshape *AllocatedResourceShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Part",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Part",
		},
		{
			Name:                 "Resource",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Resource",
		},
	}
	return
}

func (allocatedsystemshape *AllocatedSystemShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Part",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Part",
		},
		{
			Name:                 "System",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "System",
		},
	}
	return
}

func (controlflow *ControlFlow) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Description",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "ComputedPrefix",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "LayoutDirection",
			GongFieldValueType:   GongFieldValueTypeInt,
			TargetGongstructName: "LayoutDirection",
		},
		{
			Name:                 "Start",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Port",
		},
		{
			Name:                 "End",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Port",
		},
	}
	return
}

func (controlflowshape *ControlFlowShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "ControlFlow",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ControlFlow",
		},
		{
			Name:               "StartRatio",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "EndRatio",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:                 "StartOrientation",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "OrientationType",
		},
		{
			Name:                 "EndOrientation",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "OrientationType",
		},
		{
			Name:               "CornerOffsetRatio",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "IsHidden",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (data *Data) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Acronym",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Description",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "ComputedPrefix",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "LayoutDirection",
			GongFieldValueType:   GongFieldValueTypeInt,
			TargetGongstructName: "LayoutDirection",
		},
		{
			Name:               "SVG_Path",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "InverseAppliedScaling",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
	}
	return
}

func (dataflow *DataFlow) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "StartPort",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Port",
		},
		{
			Name:                 "EndPort",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Port",
		},
		{
			Name:                 "StartExternalPart",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Part",
		},
		{
			Name:                 "EndExternalPart",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Part",
		},
		{
			Name:                 "Datas",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Data",
		},
		{
			Name:               "Description",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Type",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "DataFlowType",
		},
		{
			Name:                 "Direction",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "DataFlowDirection",
		},
		{
			Name:               "IsDatasNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "ComputedPrefix",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "LayoutDirection",
			GongFieldValueType:   GongFieldValueTypeInt,
			TargetGongstructName: "LayoutDirection",
		},
	}
	return
}

func (dataflowshape *DataFlowShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "DataFlow",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "DataFlow",
		},
		{
			Name:               "StartRatio",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "EndRatio",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:                 "StartOrientation",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "OrientationType",
		},
		{
			Name:                 "EndOrientation",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "OrientationType",
		},
		{
			Name:               "CornerOffsetRatio",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "IsHidden",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (datashape *DataShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Data",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Data",
		},
		{
			Name:                 "DataFlow",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "DataFlow",
		},
	}
	return
}

func (diagramstructure *DiagramStructure) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Description",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "ComputedPrefix",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "LayoutDirection",
			GongFieldValueType:   GongFieldValueTypeInt,
			TargetGongstructName: "LayoutDirection",
		},
		{
			Name:               "IsChecked",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsEditable_",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsShowPrefix",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "DefaultBoxWidth",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "DefaultBoxHeigth",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Width",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Height",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:                 "System_Shapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "SystemShape",
		},
		{
			Name:               "IsSystemsNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "SystemsWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "System",
		},
		{
			Name:                 "Part_Shapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "PartShape",
		},
		{
			Name:               "IsPartsNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "PartWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Part",
		},
		{
			Name:                 "ExternalPart_Shapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ExternalPartShape",
		},
		{
			Name:               "IsExternalPartsNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "ExternalPartWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Part",
		},
		{
			Name:                 "ExternalPartsWhoseOutDataFlowsNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Part",
		},
		{
			Name:                 "ExternalPartsWhoseInDataFlowsNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Part",
		},
		{
			Name:                 "PortsWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Port",
		},
		{
			Name:                 "Port_Shapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "PortShape",
		},
		{
			Name:                 "ControlFlowsWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ControlFlow",
		},
		{
			Name:                 "ControlFlow_Shapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ControlFlowShape",
		},
		{
			Name:                 "DataFlowsWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DataFlow",
		},
		{
			Name:                 "DataFlow_Shapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DataFlowShape",
		},
		{
			Name:                 "DatasWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Data",
		},
		{
			Name:                 "Data_Shapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DataShape",
		},
		{
			Name:                 "DataFlowsWhoseDataNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DataFlow",
		},
		{
			Name:                 "AllocatedResourcesWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Resource",
		},
		{
			Name:                 "AllocatedResourceShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "AllocatedResourceShape",
		},
		{
			Name:                 "AllocatedSystemesWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "System",
		},
		{
			Name:                 "AllocatedSystemShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "AllocatedSystemShape",
		},
		{
			Name:                 "Note_Shapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "NoteShape",
		},
		{
			Name:                 "NotesWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Note",
		},
		{
			Name:               "IsNotesNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "NotePortShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "NotePortShape",
		},
	}
	return
}

func (externalpartshape *ExternalPartShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Part",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Part",
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "X",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Y",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Width",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Height",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "IsHidden",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "TailHeigth",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
	}
	return
}

func (library *Library) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Description",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "ComputedPrefix",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "LayoutDirection",
			GongFieldValueType:   GongFieldValueTypeInt,
			TargetGongstructName: "LayoutDirection",
		},
		{
			Name:               "IsRootLibrary",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "SubLibraries",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Library",
		},
		{
			Name:               "IsSubLibrariesNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "SubLibrariesWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Library",
		},
		{
			Name:               "NbPixPerCharacter",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "LogoSVGFile",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "RootSystemes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "System",
		},
		{
			Name:               "IsSystemesNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "SystemsWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "System",
		},
		{
			Name:                 "RootDataFlows",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DataFlow",
		},
		{
			Name:               "IsDataFlowsNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "DataFlowsWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DataFlow",
		},
		{
			Name:                 "RootDatas",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Data",
		},
		{
			Name:               "IsDatasNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "DatasWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Data",
		},
		{
			Name:                 "RootResources",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Resource",
		},
		{
			Name:               "IsResourcesNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "ResourcesWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Resource",
		},
		{
			Name:                 "PartsWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Part",
		},
		{
			Name:                 "RootNotes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Note",
		},
		{
			Name:               "IsNotesNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "NotesWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Note",
		},
		{
			Name:               "IsExpandedTmp",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (note *Note) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Description",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "ComputedPrefix",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "LayoutDirection",
			GongFieldValueType:   GongFieldValueTypeInt,
			TargetGongstructName: "LayoutDirection",
		},
		{
			Name:               "IsPortsNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "Ports",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Port",
		},
	}
	return
}

func (noteportshape *NotePortShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Note",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Note",
		},
		{
			Name:                 "Port",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Port",
		},
		{
			Name:               "StartRatio",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "EndRatio",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:                 "StartOrientation",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "OrientationType",
		},
		{
			Name:                 "EndOrientation",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "OrientationType",
		},
		{
			Name:               "CornerOffsetRatio",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "IsHidden",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (noteshape *NoteShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Note",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Note",
		},
		{
			Name:               "X",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Y",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Width",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Height",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "IsHidden",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (part *Part) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Description",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "ComputedPrefix",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "LayoutDirection",
			GongFieldValueType:   GongFieldValueTypeInt,
			TargetGongstructName: "LayoutDirection",
		},
		{
			Name:               "IsPortsNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "Ports",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Port",
		},
		{
			Name:               "IsControlFlowsNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "ControlFlows",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ControlFlow",
		},
		{
			Name:                 "PortWhoseOutControlFlowsNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Port",
		},
		{
			Name:                 "PortWhoseInControlFlowsNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Port",
		},
		{
			Name:               "IsDataFlowsNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "PortWhoseOutDataFlowsNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Port",
		},
		{
			Name:                 "PortWhoseInDataFlowsNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Port",
		},
		{
			Name:                 "PartAnchoredPath",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "PartAnchoredPath",
		},
	}
	return
}

func (partanchoredpath *PartAnchoredPath) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Definition",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "X_Offset",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Y_Offset",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:                 "RectAnchorType",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "RectAnchorType",
		},
		{
			Name:               "ScalePropotionnally",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "AppliedScaling",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "FillOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Stroke",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StrokeOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeWidth",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StrokeDashArray",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "StrokeDashArrayWhenSelected",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Transform",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (partshape *PartShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Part",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Part",
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "X",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Y",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Width",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Height",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "IsHidden",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (port *Port) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Description",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "ComputedPrefix",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "LayoutDirection",
			GongFieldValueType:   GongFieldValueTypeInt,
			TargetGongstructName: "LayoutDirection",
		},
		{
			Name:               "IsStartPort",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsEndPort",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "Type",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "System",
		},
		{
			Name:               "IsPortNameNotSystemName",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (portshape *PortShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Port",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Port",
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "X",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Y",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Width",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Height",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "IsHidden",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (resource *Resource) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Acronym",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Description",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "ComputedPrefix",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "LayoutDirection",
			GongFieldValueType:   GongFieldValueTypeInt,
			TargetGongstructName: "LayoutDirection",
		},
		{
			Name:               "SVG_Path",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "InverseAppliedScaling",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
	}
	return
}

func (system *System) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Description",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "ComputedPrefix",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "LayoutDirection",
			GongFieldValueType:   GongFieldValueTypeInt,
			TargetGongstructName: "LayoutDirection",
		},
		{
			Name:               "SVG_Path",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "InverseAppliedScaling",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:                 "DiagramStructures",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DiagramStructure",
		},
		{
			Name:                 "DiagramStructureWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DiagramStructure",
		},
		{
			Name:               "IsSubSystemNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "SubSystemes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "System",
		},
		{
			Name:                 "Parts",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Part",
		},
		{
			Name:                 "PartWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Part",
		},
		{
			Name:                 "DataFlows",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DataFlow",
		},
		{
			Name:               "IsDataFlowsNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "ExternalParts",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Part",
		},
		{
			Name:                 "ExternalPartWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Part",
		},
	}
	return
}

func (systemshape *SystemShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "System",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "System",
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "X",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Y",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Width",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Height",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "IsHidden",
			GongFieldValueType: GongFieldValueTypeBool,
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
func (allocatedresourceshape *AllocatedResourceShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = allocatedresourceshape.Name
	case "Part":
		res.GongFieldValueType = GongFieldValueTypePointer
		if allocatedresourceshape.Part != nil {
			res.valueString = allocatedresourceshape.Part.Name
			res.ids = allocatedresourceshape.Part.GongGetUUID(stage)
		}
	case "Resource":
		res.GongFieldValueType = GongFieldValueTypePointer
		if allocatedresourceshape.Resource != nil {
			res.valueString = allocatedresourceshape.Resource.Name
			res.ids = allocatedresourceshape.Resource.GongGetUUID(stage)
		}
	}
	return
}

func (allocatedsystemshape *AllocatedSystemShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = allocatedsystemshape.Name
	case "Part":
		res.GongFieldValueType = GongFieldValueTypePointer
		if allocatedsystemshape.Part != nil {
			res.valueString = allocatedsystemshape.Part.Name
			res.ids = allocatedsystemshape.Part.GongGetUUID(stage)
		}
	case "System":
		res.GongFieldValueType = GongFieldValueTypePointer
		if allocatedsystemshape.System != nil {
			res.valueString = allocatedsystemshape.System.Name
			res.ids = allocatedsystemshape.System.GongGetUUID(stage)
		}
	}
	return
}

func (controlflow *ControlFlow) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = controlflow.Name
	case "Description":
		res.valueString = controlflow.Description
	case "ComputedPrefix":
		res.valueString = controlflow.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", controlflow.IsExpanded)
		res.valueBool = controlflow.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "LayoutDirection":
		enum := controlflow.LayoutDirection
		res.valueString = enum.ToCodeString()
	case "Start":
		res.GongFieldValueType = GongFieldValueTypePointer
		if controlflow.Start != nil {
			res.valueString = controlflow.Start.Name
			res.ids = controlflow.Start.GongGetUUID(stage)
		}
	case "End":
		res.GongFieldValueType = GongFieldValueTypePointer
		if controlflow.End != nil {
			res.valueString = controlflow.End.Name
			res.ids = controlflow.End.GongGetUUID(stage)
		}
	}
	return
}

func (controlflowshape *ControlFlowShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = controlflowshape.Name
	case "ControlFlow":
		res.GongFieldValueType = GongFieldValueTypePointer
		if controlflowshape.ControlFlow != nil {
			res.valueString = controlflowshape.ControlFlow.Name
			res.ids = controlflowshape.ControlFlow.GongGetUUID(stage)
		}
	case "StartRatio":
		res.valueString = fmt.Sprintf("%f", controlflowshape.StartRatio)
		res.valueFloat = controlflowshape.StartRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndRatio":
		res.valueString = fmt.Sprintf("%f", controlflowshape.EndRatio)
		res.valueFloat = controlflowshape.EndRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartOrientation":
		enum := controlflowshape.StartOrientation
		res.valueString = enum.ToCodeString()
	case "EndOrientation":
		enum := controlflowshape.EndOrientation
		res.valueString = enum.ToCodeString()
	case "CornerOffsetRatio":
		res.valueString = fmt.Sprintf("%f", controlflowshape.CornerOffsetRatio)
		res.valueFloat = controlflowshape.CornerOffsetRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", controlflowshape.IsHidden)
		res.valueBool = controlflowshape.IsHidden
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (data *Data) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = data.Name
	case "Acronym":
		res.valueString = data.Acronym
	case "Description":
		res.valueString = data.Description
	case "ComputedPrefix":
		res.valueString = data.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", data.IsExpanded)
		res.valueBool = data.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "LayoutDirection":
		enum := data.LayoutDirection
		res.valueString = enum.ToCodeString()
	case "SVG_Path":
		res.valueString = data.SVG_Path
	case "InverseAppliedScaling":
		res.valueString = fmt.Sprintf("%f", data.InverseAppliedScaling)
		res.valueFloat = data.InverseAppliedScaling
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}

func (dataflow *DataFlow) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = dataflow.Name
	case "StartPort":
		res.GongFieldValueType = GongFieldValueTypePointer
		if dataflow.StartPort != nil {
			res.valueString = dataflow.StartPort.Name
			res.ids = dataflow.StartPort.GongGetUUID(stage)
		}
	case "EndPort":
		res.GongFieldValueType = GongFieldValueTypePointer
		if dataflow.EndPort != nil {
			res.valueString = dataflow.EndPort.Name
			res.ids = dataflow.EndPort.GongGetUUID(stage)
		}
	case "StartExternalPart":
		res.GongFieldValueType = GongFieldValueTypePointer
		if dataflow.StartExternalPart != nil {
			res.valueString = dataflow.StartExternalPart.Name
			res.ids = dataflow.StartExternalPart.GongGetUUID(stage)
		}
	case "EndExternalPart":
		res.GongFieldValueType = GongFieldValueTypePointer
		if dataflow.EndExternalPart != nil {
			res.valueString = dataflow.EndExternalPart.Name
			res.ids = dataflow.EndExternalPart.GongGetUUID(stage)
		}
	case "Datas":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range dataflow.Datas {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Description":
		res.valueString = dataflow.Description
	case "Type":
		enum := dataflow.Type
		res.valueString = enum.ToCodeString()
	case "Direction":
		enum := dataflow.Direction
		res.valueString = enum.ToCodeString()
	case "IsDatasNodeExpanded":
		res.valueString = fmt.Sprintf("%t", dataflow.IsDatasNodeExpanded)
		res.valueBool = dataflow.IsDatasNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ComputedPrefix":
		res.valueString = dataflow.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", dataflow.IsExpanded)
		res.valueBool = dataflow.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "LayoutDirection":
		enum := dataflow.LayoutDirection
		res.valueString = enum.ToCodeString()
	}
	return
}

func (dataflowshape *DataFlowShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = dataflowshape.Name
	case "DataFlow":
		res.GongFieldValueType = GongFieldValueTypePointer
		if dataflowshape.DataFlow != nil {
			res.valueString = dataflowshape.DataFlow.Name
			res.ids = dataflowshape.DataFlow.GongGetUUID(stage)
		}
	case "StartRatio":
		res.valueString = fmt.Sprintf("%f", dataflowshape.StartRatio)
		res.valueFloat = dataflowshape.StartRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndRatio":
		res.valueString = fmt.Sprintf("%f", dataflowshape.EndRatio)
		res.valueFloat = dataflowshape.EndRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartOrientation":
		enum := dataflowshape.StartOrientation
		res.valueString = enum.ToCodeString()
	case "EndOrientation":
		enum := dataflowshape.EndOrientation
		res.valueString = enum.ToCodeString()
	case "CornerOffsetRatio":
		res.valueString = fmt.Sprintf("%f", dataflowshape.CornerOffsetRatio)
		res.valueFloat = dataflowshape.CornerOffsetRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", dataflowshape.IsHidden)
		res.valueBool = dataflowshape.IsHidden
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (datashape *DataShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = datashape.Name
	case "Data":
		res.GongFieldValueType = GongFieldValueTypePointer
		if datashape.Data != nil {
			res.valueString = datashape.Data.Name
			res.ids = datashape.Data.GongGetUUID(stage)
		}
	case "DataFlow":
		res.GongFieldValueType = GongFieldValueTypePointer
		if datashape.DataFlow != nil {
			res.valueString = datashape.DataFlow.Name
			res.ids = datashape.DataFlow.GongGetUUID(stage)
		}
	}
	return
}

func (diagramstructure *DiagramStructure) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = diagramstructure.Name
	case "Description":
		res.valueString = diagramstructure.Description
	case "ComputedPrefix":
		res.valueString = diagramstructure.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", diagramstructure.IsExpanded)
		res.valueBool = diagramstructure.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "LayoutDirection":
		enum := diagramstructure.LayoutDirection
		res.valueString = enum.ToCodeString()
	case "IsChecked":
		res.valueString = fmt.Sprintf("%t", diagramstructure.IsChecked)
		res.valueBool = diagramstructure.IsChecked
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsEditable_":
		res.valueString = fmt.Sprintf("%t", diagramstructure.IsEditable_)
		res.valueBool = diagramstructure.IsEditable_
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsShowPrefix":
		res.valueString = fmt.Sprintf("%t", diagramstructure.IsShowPrefix)
		res.valueBool = diagramstructure.IsShowPrefix
		res.GongFieldValueType = GongFieldValueTypeBool
	case "DefaultBoxWidth":
		res.valueString = fmt.Sprintf("%f", diagramstructure.DefaultBoxWidth)
		res.valueFloat = diagramstructure.DefaultBoxWidth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "DefaultBoxHeigth":
		res.valueString = fmt.Sprintf("%f", diagramstructure.DefaultBoxHeigth)
		res.valueFloat = diagramstructure.DefaultBoxHeigth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Width":
		res.valueString = fmt.Sprintf("%f", diagramstructure.Width)
		res.valueFloat = diagramstructure.Width
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Height":
		res.valueString = fmt.Sprintf("%f", diagramstructure.Height)
		res.valueFloat = diagramstructure.Height
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "System_Shapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramstructure.System_Shapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsSystemsNodeExpanded":
		res.valueString = fmt.Sprintf("%t", diagramstructure.IsSystemsNodeExpanded)
		res.valueBool = diagramstructure.IsSystemsNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "SystemsWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramstructure.SystemsWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Part_Shapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramstructure.Part_Shapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsPartsNodeExpanded":
		res.valueString = fmt.Sprintf("%t", diagramstructure.IsPartsNodeExpanded)
		res.valueBool = diagramstructure.IsPartsNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "PartWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramstructure.PartWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "ExternalPart_Shapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramstructure.ExternalPart_Shapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsExternalPartsNodeExpanded":
		res.valueString = fmt.Sprintf("%t", diagramstructure.IsExternalPartsNodeExpanded)
		res.valueBool = diagramstructure.IsExternalPartsNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ExternalPartWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramstructure.ExternalPartWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "ExternalPartsWhoseOutDataFlowsNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramstructure.ExternalPartsWhoseOutDataFlowsNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "ExternalPartsWhoseInDataFlowsNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramstructure.ExternalPartsWhoseInDataFlowsNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "PortsWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramstructure.PortsWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Port_Shapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramstructure.Port_Shapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "ControlFlowsWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramstructure.ControlFlowsWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "ControlFlow_Shapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramstructure.ControlFlow_Shapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "DataFlowsWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramstructure.DataFlowsWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "DataFlow_Shapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramstructure.DataFlow_Shapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "DatasWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramstructure.DatasWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Data_Shapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramstructure.Data_Shapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "DataFlowsWhoseDataNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramstructure.DataFlowsWhoseDataNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "AllocatedResourcesWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramstructure.AllocatedResourcesWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "AllocatedResourceShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramstructure.AllocatedResourceShapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "AllocatedSystemesWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramstructure.AllocatedSystemesWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "AllocatedSystemShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramstructure.AllocatedSystemShapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Note_Shapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramstructure.Note_Shapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "NotesWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramstructure.NotesWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsNotesNodeExpanded":
		res.valueString = fmt.Sprintf("%t", diagramstructure.IsNotesNodeExpanded)
		res.valueBool = diagramstructure.IsNotesNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "NotePortShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramstructure.NotePortShapes {
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

func (externalpartshape *ExternalPartShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = externalpartshape.Name
	case "Part":
		res.GongFieldValueType = GongFieldValueTypePointer
		if externalpartshape.Part != nil {
			res.valueString = externalpartshape.Part.Name
			res.ids = externalpartshape.Part.GongGetUUID(stage)
		}
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", externalpartshape.IsExpanded)
		res.valueBool = externalpartshape.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "X":
		res.valueString = fmt.Sprintf("%f", externalpartshape.X)
		res.valueFloat = externalpartshape.X
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y":
		res.valueString = fmt.Sprintf("%f", externalpartshape.Y)
		res.valueFloat = externalpartshape.Y
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Width":
		res.valueString = fmt.Sprintf("%f", externalpartshape.Width)
		res.valueFloat = externalpartshape.Width
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Height":
		res.valueString = fmt.Sprintf("%f", externalpartshape.Height)
		res.valueFloat = externalpartshape.Height
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", externalpartshape.IsHidden)
		res.valueBool = externalpartshape.IsHidden
		res.GongFieldValueType = GongFieldValueTypeBool
	case "TailHeigth":
		res.valueString = fmt.Sprintf("%f", externalpartshape.TailHeigth)
		res.valueFloat = externalpartshape.TailHeigth
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}

func (library *Library) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = library.Name
	case "Description":
		res.valueString = library.Description
	case "ComputedPrefix":
		res.valueString = library.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", library.IsExpanded)
		res.valueBool = library.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "LayoutDirection":
		enum := library.LayoutDirection
		res.valueString = enum.ToCodeString()
	case "IsRootLibrary":
		res.valueString = fmt.Sprintf("%t", library.IsRootLibrary)
		res.valueBool = library.IsRootLibrary
		res.GongFieldValueType = GongFieldValueTypeBool
	case "SubLibraries":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range library.SubLibraries {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsSubLibrariesNodeExpanded":
		res.valueString = fmt.Sprintf("%t", library.IsSubLibrariesNodeExpanded)
		res.valueBool = library.IsSubLibrariesNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "SubLibrariesWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range library.SubLibrariesWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "NbPixPerCharacter":
		res.valueString = fmt.Sprintf("%f", library.NbPixPerCharacter)
		res.valueFloat = library.NbPixPerCharacter
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "LogoSVGFile":
		res.valueString = library.LogoSVGFile
	case "RootSystemes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range library.RootSystemes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsSystemesNodeExpanded":
		res.valueString = fmt.Sprintf("%t", library.IsSystemesNodeExpanded)
		res.valueBool = library.IsSystemesNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "SystemsWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range library.SystemsWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "RootDataFlows":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range library.RootDataFlows {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsDataFlowsNodeExpanded":
		res.valueString = fmt.Sprintf("%t", library.IsDataFlowsNodeExpanded)
		res.valueBool = library.IsDataFlowsNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "DataFlowsWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range library.DataFlowsWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "RootDatas":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range library.RootDatas {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsDatasNodeExpanded":
		res.valueString = fmt.Sprintf("%t", library.IsDatasNodeExpanded)
		res.valueBool = library.IsDatasNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "DatasWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range library.DatasWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "RootResources":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range library.RootResources {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsResourcesNodeExpanded":
		res.valueString = fmt.Sprintf("%t", library.IsResourcesNodeExpanded)
		res.valueBool = library.IsResourcesNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ResourcesWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range library.ResourcesWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "PartsWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range library.PartsWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "RootNotes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range library.RootNotes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsNotesNodeExpanded":
		res.valueString = fmt.Sprintf("%t", library.IsNotesNodeExpanded)
		res.valueBool = library.IsNotesNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "NotesWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range library.NotesWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsExpandedTmp":
		res.valueString = fmt.Sprintf("%t", library.IsExpandedTmp)
		res.valueBool = library.IsExpandedTmp
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (note *Note) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = note.Name
	case "Description":
		res.valueString = note.Description
	case "ComputedPrefix":
		res.valueString = note.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", note.IsExpanded)
		res.valueBool = note.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "LayoutDirection":
		enum := note.LayoutDirection
		res.valueString = enum.ToCodeString()
	case "IsPortsNodeExpanded":
		res.valueString = fmt.Sprintf("%t", note.IsPortsNodeExpanded)
		res.valueBool = note.IsPortsNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "Ports":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range note.Ports {
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

func (noteportshape *NotePortShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = noteportshape.Name
	case "Note":
		res.GongFieldValueType = GongFieldValueTypePointer
		if noteportshape.Note != nil {
			res.valueString = noteportshape.Note.Name
			res.ids = noteportshape.Note.GongGetUUID(stage)
		}
	case "Port":
		res.GongFieldValueType = GongFieldValueTypePointer
		if noteportshape.Port != nil {
			res.valueString = noteportshape.Port.Name
			res.ids = noteportshape.Port.GongGetUUID(stage)
		}
	case "StartRatio":
		res.valueString = fmt.Sprintf("%f", noteportshape.StartRatio)
		res.valueFloat = noteportshape.StartRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndRatio":
		res.valueString = fmt.Sprintf("%f", noteportshape.EndRatio)
		res.valueFloat = noteportshape.EndRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartOrientation":
		enum := noteportshape.StartOrientation
		res.valueString = enum.ToCodeString()
	case "EndOrientation":
		enum := noteportshape.EndOrientation
		res.valueString = enum.ToCodeString()
	case "CornerOffsetRatio":
		res.valueString = fmt.Sprintf("%f", noteportshape.CornerOffsetRatio)
		res.valueFloat = noteportshape.CornerOffsetRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", noteportshape.IsHidden)
		res.valueBool = noteportshape.IsHidden
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (noteshape *NoteShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = noteshape.Name
	case "Note":
		res.GongFieldValueType = GongFieldValueTypePointer
		if noteshape.Note != nil {
			res.valueString = noteshape.Note.Name
			res.ids = noteshape.Note.GongGetUUID(stage)
		}
	case "X":
		res.valueString = fmt.Sprintf("%f", noteshape.X)
		res.valueFloat = noteshape.X
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y":
		res.valueString = fmt.Sprintf("%f", noteshape.Y)
		res.valueFloat = noteshape.Y
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Width":
		res.valueString = fmt.Sprintf("%f", noteshape.Width)
		res.valueFloat = noteshape.Width
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Height":
		res.valueString = fmt.Sprintf("%f", noteshape.Height)
		res.valueFloat = noteshape.Height
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", noteshape.IsHidden)
		res.valueBool = noteshape.IsHidden
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (part *Part) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = part.Name
	case "Description":
		res.valueString = part.Description
	case "ComputedPrefix":
		res.valueString = part.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", part.IsExpanded)
		res.valueBool = part.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "LayoutDirection":
		enum := part.LayoutDirection
		res.valueString = enum.ToCodeString()
	case "IsPortsNodeExpanded":
		res.valueString = fmt.Sprintf("%t", part.IsPortsNodeExpanded)
		res.valueBool = part.IsPortsNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "Ports":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range part.Ports {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsControlFlowsNodeExpanded":
		res.valueString = fmt.Sprintf("%t", part.IsControlFlowsNodeExpanded)
		res.valueBool = part.IsControlFlowsNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ControlFlows":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range part.ControlFlows {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "PortWhoseOutControlFlowsNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range part.PortWhoseOutControlFlowsNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "PortWhoseInControlFlowsNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range part.PortWhoseInControlFlowsNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsDataFlowsNodeExpanded":
		res.valueString = fmt.Sprintf("%t", part.IsDataFlowsNodeExpanded)
		res.valueBool = part.IsDataFlowsNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "PortWhoseOutDataFlowsNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range part.PortWhoseOutDataFlowsNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "PortWhoseInDataFlowsNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range part.PortWhoseInDataFlowsNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "PartAnchoredPath":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range part.PartAnchoredPath {
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

func (partanchoredpath *PartAnchoredPath) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = partanchoredpath.Name
	case "Definition":
		res.valueString = partanchoredpath.Definition
	case "X_Offset":
		res.valueString = fmt.Sprintf("%f", partanchoredpath.X_Offset)
		res.valueFloat = partanchoredpath.X_Offset
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y_Offset":
		res.valueString = fmt.Sprintf("%f", partanchoredpath.Y_Offset)
		res.valueFloat = partanchoredpath.Y_Offset
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "RectAnchorType":
		enum := partanchoredpath.RectAnchorType
		res.valueString = enum.ToCodeString()
	case "ScalePropotionnally":
		res.valueString = fmt.Sprintf("%t", partanchoredpath.ScalePropotionnally)
		res.valueBool = partanchoredpath.ScalePropotionnally
		res.GongFieldValueType = GongFieldValueTypeBool
	case "AppliedScaling":
		res.valueString = fmt.Sprintf("%f", partanchoredpath.AppliedScaling)
		res.valueFloat = partanchoredpath.AppliedScaling
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Color":
		res.valueString = partanchoredpath.Color
	case "FillOpacity":
		res.valueString = fmt.Sprintf("%f", partanchoredpath.FillOpacity)
		res.valueFloat = partanchoredpath.FillOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Stroke":
		res.valueString = partanchoredpath.Stroke
	case "StrokeOpacity":
		res.valueString = fmt.Sprintf("%f", partanchoredpath.StrokeOpacity)
		res.valueFloat = partanchoredpath.StrokeOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeWidth":
		res.valueString = fmt.Sprintf("%f", partanchoredpath.StrokeWidth)
		res.valueFloat = partanchoredpath.StrokeWidth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StrokeDashArray":
		res.valueString = partanchoredpath.StrokeDashArray
	case "StrokeDashArrayWhenSelected":
		res.valueString = partanchoredpath.StrokeDashArrayWhenSelected
	case "Transform":
		res.valueString = partanchoredpath.Transform
	}
	return
}

func (partshape *PartShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = partshape.Name
	case "Part":
		res.GongFieldValueType = GongFieldValueTypePointer
		if partshape.Part != nil {
			res.valueString = partshape.Part.Name
			res.ids = partshape.Part.GongGetUUID(stage)
		}
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", partshape.IsExpanded)
		res.valueBool = partshape.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "X":
		res.valueString = fmt.Sprintf("%f", partshape.X)
		res.valueFloat = partshape.X
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y":
		res.valueString = fmt.Sprintf("%f", partshape.Y)
		res.valueFloat = partshape.Y
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Width":
		res.valueString = fmt.Sprintf("%f", partshape.Width)
		res.valueFloat = partshape.Width
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Height":
		res.valueString = fmt.Sprintf("%f", partshape.Height)
		res.valueFloat = partshape.Height
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", partshape.IsHidden)
		res.valueBool = partshape.IsHidden
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (port *Port) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = port.Name
	case "Description":
		res.valueString = port.Description
	case "ComputedPrefix":
		res.valueString = port.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", port.IsExpanded)
		res.valueBool = port.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "LayoutDirection":
		enum := port.LayoutDirection
		res.valueString = enum.ToCodeString()
	case "IsStartPort":
		res.valueString = fmt.Sprintf("%t", port.IsStartPort)
		res.valueBool = port.IsStartPort
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsEndPort":
		res.valueString = fmt.Sprintf("%t", port.IsEndPort)
		res.valueBool = port.IsEndPort
		res.GongFieldValueType = GongFieldValueTypeBool
	case "Type":
		res.GongFieldValueType = GongFieldValueTypePointer
		if port.Type != nil {
			res.valueString = port.Type.Name
			res.ids = port.Type.GongGetUUID(stage)
		}
	case "IsPortNameNotSystemName":
		res.valueString = fmt.Sprintf("%t", port.IsPortNameNotSystemName)
		res.valueBool = port.IsPortNameNotSystemName
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (portshape *PortShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = portshape.Name
	case "Port":
		res.GongFieldValueType = GongFieldValueTypePointer
		if portshape.Port != nil {
			res.valueString = portshape.Port.Name
			res.ids = portshape.Port.GongGetUUID(stage)
		}
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", portshape.IsExpanded)
		res.valueBool = portshape.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "X":
		res.valueString = fmt.Sprintf("%f", portshape.X)
		res.valueFloat = portshape.X
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y":
		res.valueString = fmt.Sprintf("%f", portshape.Y)
		res.valueFloat = portshape.Y
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Width":
		res.valueString = fmt.Sprintf("%f", portshape.Width)
		res.valueFloat = portshape.Width
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Height":
		res.valueString = fmt.Sprintf("%f", portshape.Height)
		res.valueFloat = portshape.Height
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", portshape.IsHidden)
		res.valueBool = portshape.IsHidden
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (resource *Resource) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = resource.Name
	case "Acronym":
		res.valueString = resource.Acronym
	case "Description":
		res.valueString = resource.Description
	case "ComputedPrefix":
		res.valueString = resource.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", resource.IsExpanded)
		res.valueBool = resource.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "LayoutDirection":
		enum := resource.LayoutDirection
		res.valueString = enum.ToCodeString()
	case "SVG_Path":
		res.valueString = resource.SVG_Path
	case "InverseAppliedScaling":
		res.valueString = fmt.Sprintf("%f", resource.InverseAppliedScaling)
		res.valueFloat = resource.InverseAppliedScaling
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}

func (system *System) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = system.Name
	case "Description":
		res.valueString = system.Description
	case "ComputedPrefix":
		res.valueString = system.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", system.IsExpanded)
		res.valueBool = system.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "LayoutDirection":
		enum := system.LayoutDirection
		res.valueString = enum.ToCodeString()
	case "SVG_Path":
		res.valueString = system.SVG_Path
	case "InverseAppliedScaling":
		res.valueString = fmt.Sprintf("%f", system.InverseAppliedScaling)
		res.valueFloat = system.InverseAppliedScaling
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "DiagramStructures":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range system.DiagramStructures {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "DiagramStructureWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range system.DiagramStructureWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsSubSystemNodeExpanded":
		res.valueString = fmt.Sprintf("%t", system.IsSubSystemNodeExpanded)
		res.valueBool = system.IsSubSystemNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "SubSystemes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range system.SubSystemes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Parts":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range system.Parts {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "PartWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range system.PartWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "DataFlows":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range system.DataFlows {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsDataFlowsNodeExpanded":
		res.valueString = fmt.Sprintf("%t", system.IsDataFlowsNodeExpanded)
		res.valueBool = system.IsDataFlowsNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ExternalParts":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range system.ExternalParts {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "ExternalPartWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range system.ExternalPartWhoseNodeIsExpanded {
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

func (systemshape *SystemShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = systemshape.Name
	case "System":
		res.GongFieldValueType = GongFieldValueTypePointer
		if systemshape.System != nil {
			res.valueString = systemshape.System.Name
			res.ids = systemshape.System.GongGetUUID(stage)
		}
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", systemshape.IsExpanded)
		res.valueBool = systemshape.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "X":
		res.valueString = fmt.Sprintf("%f", systemshape.X)
		res.valueFloat = systemshape.X
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y":
		res.valueString = fmt.Sprintf("%f", systemshape.Y)
		res.valueFloat = systemshape.Y
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Width":
		res.valueString = fmt.Sprintf("%f", systemshape.Width)
		res.valueFloat = systemshape.Width
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Height":
		res.valueString = fmt.Sprintf("%f", systemshape.Height)
		res.valueFloat = systemshape.Height
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", systemshape.IsHidden)
		res.valueBool = systemshape.IsHidden
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func GetFieldStringValueFromPointer(instance GongstructIF, fieldName string, stage *Stage) (res GongFieldValue) {
	res = instance.GongGetFieldValue(fieldName, stage)
	return
}

// insertion point for generic set gongstruct field value
func (allocatedresourceshape *AllocatedResourceShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		allocatedresourceshape.Name = value.GetValueString()
	case "Part":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			allocatedresourceshape.Part = nil
			for __instance__ := range stage.Parts {
				if stage.Part_stagedOrder[__instance__] == uint(id) {
					allocatedresourceshape.Part = __instance__
					break
				}
			}
		}
	case "Resource":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			allocatedresourceshape.Resource = nil
			for __instance__ := range stage.Resources {
				if stage.Resource_stagedOrder[__instance__] == uint(id) {
					allocatedresourceshape.Resource = __instance__
					break
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (allocatedsystemshape *AllocatedSystemShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		allocatedsystemshape.Name = value.GetValueString()
	case "Part":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			allocatedsystemshape.Part = nil
			for __instance__ := range stage.Parts {
				if stage.Part_stagedOrder[__instance__] == uint(id) {
					allocatedsystemshape.Part = __instance__
					break
				}
			}
		}
	case "System":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			allocatedsystemshape.System = nil
			for __instance__ := range stage.Systems {
				if stage.System_stagedOrder[__instance__] == uint(id) {
					allocatedsystemshape.System = __instance__
					break
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (controlflow *ControlFlow) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		controlflow.Name = value.GetValueString()
	case "Description":
		controlflow.Description = value.GetValueString()
	case "ComputedPrefix":
		controlflow.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		controlflow.IsExpanded = value.GetValueBool()
	case "LayoutDirection":
		controlflow.LayoutDirection.FromCodeString(value.GetValueString())
	case "Start":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			controlflow.Start = nil
			for __instance__ := range stage.Ports {
				if stage.Port_stagedOrder[__instance__] == uint(id) {
					controlflow.Start = __instance__
					break
				}
			}
		}
	case "End":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			controlflow.End = nil
			for __instance__ := range stage.Ports {
				if stage.Port_stagedOrder[__instance__] == uint(id) {
					controlflow.End = __instance__
					break
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (controlflowshape *ControlFlowShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		controlflowshape.Name = value.GetValueString()
	case "ControlFlow":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			controlflowshape.ControlFlow = nil
			for __instance__ := range stage.ControlFlows {
				if stage.ControlFlow_stagedOrder[__instance__] == uint(id) {
					controlflowshape.ControlFlow = __instance__
					break
				}
			}
		}
	case "StartRatio":
		controlflowshape.StartRatio = value.GetValueFloat()
	case "EndRatio":
		controlflowshape.EndRatio = value.GetValueFloat()
	case "StartOrientation":
		controlflowshape.StartOrientation.FromCodeString(value.GetValueString())
	case "EndOrientation":
		controlflowshape.EndOrientation.FromCodeString(value.GetValueString())
	case "CornerOffsetRatio":
		controlflowshape.CornerOffsetRatio = value.GetValueFloat()
	case "IsHidden":
		controlflowshape.IsHidden = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (data *Data) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		data.Name = value.GetValueString()
	case "Acronym":
		data.Acronym = value.GetValueString()
	case "Description":
		data.Description = value.GetValueString()
	case "ComputedPrefix":
		data.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		data.IsExpanded = value.GetValueBool()
	case "LayoutDirection":
		data.LayoutDirection.FromCodeString(value.GetValueString())
	case "SVG_Path":
		data.SVG_Path = value.GetValueString()
	case "InverseAppliedScaling":
		data.InverseAppliedScaling = value.GetValueFloat()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (dataflow *DataFlow) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		dataflow.Name = value.GetValueString()
	case "StartPort":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			dataflow.StartPort = nil
			for __instance__ := range stage.Ports {
				if stage.Port_stagedOrder[__instance__] == uint(id) {
					dataflow.StartPort = __instance__
					break
				}
			}
		}
	case "EndPort":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			dataflow.EndPort = nil
			for __instance__ := range stage.Ports {
				if stage.Port_stagedOrder[__instance__] == uint(id) {
					dataflow.EndPort = __instance__
					break
				}
			}
		}
	case "StartExternalPart":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			dataflow.StartExternalPart = nil
			for __instance__ := range stage.Parts {
				if stage.Part_stagedOrder[__instance__] == uint(id) {
					dataflow.StartExternalPart = __instance__
					break
				}
			}
		}
	case "EndExternalPart":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			dataflow.EndExternalPart = nil
			for __instance__ := range stage.Parts {
				if stage.Part_stagedOrder[__instance__] == uint(id) {
					dataflow.EndExternalPart = __instance__
					break
				}
			}
		}
	case "Datas":
		dataflow.Datas = make([]*Data, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Datas {
					if stage.Data_stagedOrder[__instance__] == uint(id) {
						dataflow.Datas = append(dataflow.Datas, __instance__)
						break
					}
				}
			}
		}
	case "Description":
		dataflow.Description = value.GetValueString()
	case "Type":
		dataflow.Type.FromCodeString(value.GetValueString())
	case "Direction":
		dataflow.Direction.FromCodeString(value.GetValueString())
	case "IsDatasNodeExpanded":
		dataflow.IsDatasNodeExpanded = value.GetValueBool()
	case "ComputedPrefix":
		dataflow.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		dataflow.IsExpanded = value.GetValueBool()
	case "LayoutDirection":
		dataflow.LayoutDirection.FromCodeString(value.GetValueString())
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (dataflowshape *DataFlowShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		dataflowshape.Name = value.GetValueString()
	case "DataFlow":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			dataflowshape.DataFlow = nil
			for __instance__ := range stage.DataFlows {
				if stage.DataFlow_stagedOrder[__instance__] == uint(id) {
					dataflowshape.DataFlow = __instance__
					break
				}
			}
		}
	case "StartRatio":
		dataflowshape.StartRatio = value.GetValueFloat()
	case "EndRatio":
		dataflowshape.EndRatio = value.GetValueFloat()
	case "StartOrientation":
		dataflowshape.StartOrientation.FromCodeString(value.GetValueString())
	case "EndOrientation":
		dataflowshape.EndOrientation.FromCodeString(value.GetValueString())
	case "CornerOffsetRatio":
		dataflowshape.CornerOffsetRatio = value.GetValueFloat()
	case "IsHidden":
		dataflowshape.IsHidden = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (datashape *DataShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		datashape.Name = value.GetValueString()
	case "Data":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			datashape.Data = nil
			for __instance__ := range stage.Datas {
				if stage.Data_stagedOrder[__instance__] == uint(id) {
					datashape.Data = __instance__
					break
				}
			}
		}
	case "DataFlow":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			datashape.DataFlow = nil
			for __instance__ := range stage.DataFlows {
				if stage.DataFlow_stagedOrder[__instance__] == uint(id) {
					datashape.DataFlow = __instance__
					break
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (diagramstructure *DiagramStructure) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		diagramstructure.Name = value.GetValueString()
	case "Description":
		diagramstructure.Description = value.GetValueString()
	case "ComputedPrefix":
		diagramstructure.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		diagramstructure.IsExpanded = value.GetValueBool()
	case "LayoutDirection":
		diagramstructure.LayoutDirection.FromCodeString(value.GetValueString())
	case "IsChecked":
		diagramstructure.IsChecked = value.GetValueBool()
	case "IsEditable_":
		diagramstructure.IsEditable_ = value.GetValueBool()
	case "IsShowPrefix":
		diagramstructure.IsShowPrefix = value.GetValueBool()
	case "DefaultBoxWidth":
		diagramstructure.DefaultBoxWidth = value.GetValueFloat()
	case "DefaultBoxHeigth":
		diagramstructure.DefaultBoxHeigth = value.GetValueFloat()
	case "Width":
		diagramstructure.Width = value.GetValueFloat()
	case "Height":
		diagramstructure.Height = value.GetValueFloat()
	case "System_Shapes":
		diagramstructure.System_Shapes = make([]*SystemShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.SystemShapes {
					if stage.SystemShape_stagedOrder[__instance__] == uint(id) {
						diagramstructure.System_Shapes = append(diagramstructure.System_Shapes, __instance__)
						break
					}
				}
			}
		}
	case "IsSystemsNodeExpanded":
		diagramstructure.IsSystemsNodeExpanded = value.GetValueBool()
	case "SystemsWhoseNodeIsExpanded":
		diagramstructure.SystemsWhoseNodeIsExpanded = make([]*System, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Systems {
					if stage.System_stagedOrder[__instance__] == uint(id) {
						diagramstructure.SystemsWhoseNodeIsExpanded = append(diagramstructure.SystemsWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "Part_Shapes":
		diagramstructure.Part_Shapes = make([]*PartShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.PartShapes {
					if stage.PartShape_stagedOrder[__instance__] == uint(id) {
						diagramstructure.Part_Shapes = append(diagramstructure.Part_Shapes, __instance__)
						break
					}
				}
			}
		}
	case "IsPartsNodeExpanded":
		diagramstructure.IsPartsNodeExpanded = value.GetValueBool()
	case "PartWhoseNodeIsExpanded":
		diagramstructure.PartWhoseNodeIsExpanded = make([]*Part, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Parts {
					if stage.Part_stagedOrder[__instance__] == uint(id) {
						diagramstructure.PartWhoseNodeIsExpanded = append(diagramstructure.PartWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "ExternalPart_Shapes":
		diagramstructure.ExternalPart_Shapes = make([]*ExternalPartShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.ExternalPartShapes {
					if stage.ExternalPartShape_stagedOrder[__instance__] == uint(id) {
						diagramstructure.ExternalPart_Shapes = append(diagramstructure.ExternalPart_Shapes, __instance__)
						break
					}
				}
			}
		}
	case "IsExternalPartsNodeExpanded":
		diagramstructure.IsExternalPartsNodeExpanded = value.GetValueBool()
	case "ExternalPartWhoseNodeIsExpanded":
		diagramstructure.ExternalPartWhoseNodeIsExpanded = make([]*Part, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Parts {
					if stage.Part_stagedOrder[__instance__] == uint(id) {
						diagramstructure.ExternalPartWhoseNodeIsExpanded = append(diagramstructure.ExternalPartWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "ExternalPartsWhoseOutDataFlowsNodeIsExpanded":
		diagramstructure.ExternalPartsWhoseOutDataFlowsNodeIsExpanded = make([]*Part, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Parts {
					if stage.Part_stagedOrder[__instance__] == uint(id) {
						diagramstructure.ExternalPartsWhoseOutDataFlowsNodeIsExpanded = append(diagramstructure.ExternalPartsWhoseOutDataFlowsNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "ExternalPartsWhoseInDataFlowsNodeIsExpanded":
		diagramstructure.ExternalPartsWhoseInDataFlowsNodeIsExpanded = make([]*Part, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Parts {
					if stage.Part_stagedOrder[__instance__] == uint(id) {
						diagramstructure.ExternalPartsWhoseInDataFlowsNodeIsExpanded = append(diagramstructure.ExternalPartsWhoseInDataFlowsNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "PortsWhoseNodeIsExpanded":
		diagramstructure.PortsWhoseNodeIsExpanded = make([]*Port, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Ports {
					if stage.Port_stagedOrder[__instance__] == uint(id) {
						diagramstructure.PortsWhoseNodeIsExpanded = append(diagramstructure.PortsWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "Port_Shapes":
		diagramstructure.Port_Shapes = make([]*PortShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.PortShapes {
					if stage.PortShape_stagedOrder[__instance__] == uint(id) {
						diagramstructure.Port_Shapes = append(diagramstructure.Port_Shapes, __instance__)
						break
					}
				}
			}
		}
	case "ControlFlowsWhoseNodeIsExpanded":
		diagramstructure.ControlFlowsWhoseNodeIsExpanded = make([]*ControlFlow, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.ControlFlows {
					if stage.ControlFlow_stagedOrder[__instance__] == uint(id) {
						diagramstructure.ControlFlowsWhoseNodeIsExpanded = append(diagramstructure.ControlFlowsWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "ControlFlow_Shapes":
		diagramstructure.ControlFlow_Shapes = make([]*ControlFlowShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.ControlFlowShapes {
					if stage.ControlFlowShape_stagedOrder[__instance__] == uint(id) {
						diagramstructure.ControlFlow_Shapes = append(diagramstructure.ControlFlow_Shapes, __instance__)
						break
					}
				}
			}
		}
	case "DataFlowsWhoseNodeIsExpanded":
		diagramstructure.DataFlowsWhoseNodeIsExpanded = make([]*DataFlow, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.DataFlows {
					if stage.DataFlow_stagedOrder[__instance__] == uint(id) {
						diagramstructure.DataFlowsWhoseNodeIsExpanded = append(diagramstructure.DataFlowsWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "DataFlow_Shapes":
		diagramstructure.DataFlow_Shapes = make([]*DataFlowShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.DataFlowShapes {
					if stage.DataFlowShape_stagedOrder[__instance__] == uint(id) {
						diagramstructure.DataFlow_Shapes = append(diagramstructure.DataFlow_Shapes, __instance__)
						break
					}
				}
			}
		}
	case "DatasWhoseNodeIsExpanded":
		diagramstructure.DatasWhoseNodeIsExpanded = make([]*Data, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Datas {
					if stage.Data_stagedOrder[__instance__] == uint(id) {
						diagramstructure.DatasWhoseNodeIsExpanded = append(diagramstructure.DatasWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "Data_Shapes":
		diagramstructure.Data_Shapes = make([]*DataShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.DataShapes {
					if stage.DataShape_stagedOrder[__instance__] == uint(id) {
						diagramstructure.Data_Shapes = append(diagramstructure.Data_Shapes, __instance__)
						break
					}
				}
			}
		}
	case "DataFlowsWhoseDataNodeIsExpanded":
		diagramstructure.DataFlowsWhoseDataNodeIsExpanded = make([]*DataFlow, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.DataFlows {
					if stage.DataFlow_stagedOrder[__instance__] == uint(id) {
						diagramstructure.DataFlowsWhoseDataNodeIsExpanded = append(diagramstructure.DataFlowsWhoseDataNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "AllocatedResourcesWhoseNodeIsExpanded":
		diagramstructure.AllocatedResourcesWhoseNodeIsExpanded = make([]*Resource, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Resources {
					if stage.Resource_stagedOrder[__instance__] == uint(id) {
						diagramstructure.AllocatedResourcesWhoseNodeIsExpanded = append(diagramstructure.AllocatedResourcesWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "AllocatedResourceShapes":
		diagramstructure.AllocatedResourceShapes = make([]*AllocatedResourceShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.AllocatedResourceShapes {
					if stage.AllocatedResourceShape_stagedOrder[__instance__] == uint(id) {
						diagramstructure.AllocatedResourceShapes = append(diagramstructure.AllocatedResourceShapes, __instance__)
						break
					}
				}
			}
		}
	case "AllocatedSystemesWhoseNodeIsExpanded":
		diagramstructure.AllocatedSystemesWhoseNodeIsExpanded = make([]*System, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Systems {
					if stage.System_stagedOrder[__instance__] == uint(id) {
						diagramstructure.AllocatedSystemesWhoseNodeIsExpanded = append(diagramstructure.AllocatedSystemesWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "AllocatedSystemShapes":
		diagramstructure.AllocatedSystemShapes = make([]*AllocatedSystemShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.AllocatedSystemShapes {
					if stage.AllocatedSystemShape_stagedOrder[__instance__] == uint(id) {
						diagramstructure.AllocatedSystemShapes = append(diagramstructure.AllocatedSystemShapes, __instance__)
						break
					}
				}
			}
		}
	case "Note_Shapes":
		diagramstructure.Note_Shapes = make([]*NoteShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.NoteShapes {
					if stage.NoteShape_stagedOrder[__instance__] == uint(id) {
						diagramstructure.Note_Shapes = append(diagramstructure.Note_Shapes, __instance__)
						break
					}
				}
			}
		}
	case "NotesWhoseNodeIsExpanded":
		diagramstructure.NotesWhoseNodeIsExpanded = make([]*Note, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Notes {
					if stage.Note_stagedOrder[__instance__] == uint(id) {
						diagramstructure.NotesWhoseNodeIsExpanded = append(diagramstructure.NotesWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "IsNotesNodeExpanded":
		diagramstructure.IsNotesNodeExpanded = value.GetValueBool()
	case "NotePortShapes":
		diagramstructure.NotePortShapes = make([]*NotePortShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.NotePortShapes {
					if stage.NotePortShape_stagedOrder[__instance__] == uint(id) {
						diagramstructure.NotePortShapes = append(diagramstructure.NotePortShapes, __instance__)
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

func (externalpartshape *ExternalPartShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		externalpartshape.Name = value.GetValueString()
	case "Part":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			externalpartshape.Part = nil
			for __instance__ := range stage.Parts {
				if stage.Part_stagedOrder[__instance__] == uint(id) {
					externalpartshape.Part = __instance__
					break
				}
			}
		}
	case "IsExpanded":
		externalpartshape.IsExpanded = value.GetValueBool()
	case "X":
		externalpartshape.X = value.GetValueFloat()
	case "Y":
		externalpartshape.Y = value.GetValueFloat()
	case "Width":
		externalpartshape.Width = value.GetValueFloat()
	case "Height":
		externalpartshape.Height = value.GetValueFloat()
	case "IsHidden":
		externalpartshape.IsHidden = value.GetValueBool()
	case "TailHeigth":
		externalpartshape.TailHeigth = value.GetValueFloat()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (library *Library) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		library.Name = value.GetValueString()
	case "Description":
		library.Description = value.GetValueString()
	case "ComputedPrefix":
		library.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		library.IsExpanded = value.GetValueBool()
	case "LayoutDirection":
		library.LayoutDirection.FromCodeString(value.GetValueString())
	case "IsRootLibrary":
		library.IsRootLibrary = value.GetValueBool()
	case "SubLibraries":
		library.SubLibraries = make([]*Library, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Librarys {
					if stage.Library_stagedOrder[__instance__] == uint(id) {
						library.SubLibraries = append(library.SubLibraries, __instance__)
						break
					}
				}
			}
		}
	case "IsSubLibrariesNodeExpanded":
		library.IsSubLibrariesNodeExpanded = value.GetValueBool()
	case "SubLibrariesWhoseNodeIsExpanded":
		library.SubLibrariesWhoseNodeIsExpanded = make([]*Library, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Librarys {
					if stage.Library_stagedOrder[__instance__] == uint(id) {
						library.SubLibrariesWhoseNodeIsExpanded = append(library.SubLibrariesWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "NbPixPerCharacter":
		library.NbPixPerCharacter = value.GetValueFloat()
	case "LogoSVGFile":
		library.LogoSVGFile = value.GetValueString()
	case "RootSystemes":
		library.RootSystemes = make([]*System, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Systems {
					if stage.System_stagedOrder[__instance__] == uint(id) {
						library.RootSystemes = append(library.RootSystemes, __instance__)
						break
					}
				}
			}
		}
	case "IsSystemesNodeExpanded":
		library.IsSystemesNodeExpanded = value.GetValueBool()
	case "SystemsWhoseNodeIsExpanded":
		library.SystemsWhoseNodeIsExpanded = make([]*System, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Systems {
					if stage.System_stagedOrder[__instance__] == uint(id) {
						library.SystemsWhoseNodeIsExpanded = append(library.SystemsWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "RootDataFlows":
		library.RootDataFlows = make([]*DataFlow, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.DataFlows {
					if stage.DataFlow_stagedOrder[__instance__] == uint(id) {
						library.RootDataFlows = append(library.RootDataFlows, __instance__)
						break
					}
				}
			}
		}
	case "IsDataFlowsNodeExpanded":
		library.IsDataFlowsNodeExpanded = value.GetValueBool()
	case "DataFlowsWhoseNodeIsExpanded":
		library.DataFlowsWhoseNodeIsExpanded = make([]*DataFlow, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.DataFlows {
					if stage.DataFlow_stagedOrder[__instance__] == uint(id) {
						library.DataFlowsWhoseNodeIsExpanded = append(library.DataFlowsWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "RootDatas":
		library.RootDatas = make([]*Data, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Datas {
					if stage.Data_stagedOrder[__instance__] == uint(id) {
						library.RootDatas = append(library.RootDatas, __instance__)
						break
					}
				}
			}
		}
	case "IsDatasNodeExpanded":
		library.IsDatasNodeExpanded = value.GetValueBool()
	case "DatasWhoseNodeIsExpanded":
		library.DatasWhoseNodeIsExpanded = make([]*Data, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Datas {
					if stage.Data_stagedOrder[__instance__] == uint(id) {
						library.DatasWhoseNodeIsExpanded = append(library.DatasWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "RootResources":
		library.RootResources = make([]*Resource, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Resources {
					if stage.Resource_stagedOrder[__instance__] == uint(id) {
						library.RootResources = append(library.RootResources, __instance__)
						break
					}
				}
			}
		}
	case "IsResourcesNodeExpanded":
		library.IsResourcesNodeExpanded = value.GetValueBool()
	case "ResourcesWhoseNodeIsExpanded":
		library.ResourcesWhoseNodeIsExpanded = make([]*Resource, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Resources {
					if stage.Resource_stagedOrder[__instance__] == uint(id) {
						library.ResourcesWhoseNodeIsExpanded = append(library.ResourcesWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "PartsWhoseNodeIsExpanded":
		library.PartsWhoseNodeIsExpanded = make([]*Part, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Parts {
					if stage.Part_stagedOrder[__instance__] == uint(id) {
						library.PartsWhoseNodeIsExpanded = append(library.PartsWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "RootNotes":
		library.RootNotes = make([]*Note, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Notes {
					if stage.Note_stagedOrder[__instance__] == uint(id) {
						library.RootNotes = append(library.RootNotes, __instance__)
						break
					}
				}
			}
		}
	case "IsNotesNodeExpanded":
		library.IsNotesNodeExpanded = value.GetValueBool()
	case "NotesWhoseNodeIsExpanded":
		library.NotesWhoseNodeIsExpanded = make([]*Note, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Notes {
					if stage.Note_stagedOrder[__instance__] == uint(id) {
						library.NotesWhoseNodeIsExpanded = append(library.NotesWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "IsExpandedTmp":
		library.IsExpandedTmp = value.GetValueBool()
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
	case "Description":
		note.Description = value.GetValueString()
	case "ComputedPrefix":
		note.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		note.IsExpanded = value.GetValueBool()
	case "LayoutDirection":
		note.LayoutDirection.FromCodeString(value.GetValueString())
	case "IsPortsNodeExpanded":
		note.IsPortsNodeExpanded = value.GetValueBool()
	case "Ports":
		note.Ports = make([]*Port, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Ports {
					if stage.Port_stagedOrder[__instance__] == uint(id) {
						note.Ports = append(note.Ports, __instance__)
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

func (noteportshape *NotePortShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		noteportshape.Name = value.GetValueString()
	case "Note":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			noteportshape.Note = nil
			for __instance__ := range stage.Notes {
				if stage.Note_stagedOrder[__instance__] == uint(id) {
					noteportshape.Note = __instance__
					break
				}
			}
		}
	case "Port":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			noteportshape.Port = nil
			for __instance__ := range stage.Ports {
				if stage.Port_stagedOrder[__instance__] == uint(id) {
					noteportshape.Port = __instance__
					break
				}
			}
		}
	case "StartRatio":
		noteportshape.StartRatio = value.GetValueFloat()
	case "EndRatio":
		noteportshape.EndRatio = value.GetValueFloat()
	case "StartOrientation":
		noteportshape.StartOrientation.FromCodeString(value.GetValueString())
	case "EndOrientation":
		noteportshape.EndOrientation.FromCodeString(value.GetValueString())
	case "CornerOffsetRatio":
		noteportshape.CornerOffsetRatio = value.GetValueFloat()
	case "IsHidden":
		noteportshape.IsHidden = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (noteshape *NoteShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		noteshape.Name = value.GetValueString()
	case "Note":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			noteshape.Note = nil
			for __instance__ := range stage.Notes {
				if stage.Note_stagedOrder[__instance__] == uint(id) {
					noteshape.Note = __instance__
					break
				}
			}
		}
	case "X":
		noteshape.X = value.GetValueFloat()
	case "Y":
		noteshape.Y = value.GetValueFloat()
	case "Width":
		noteshape.Width = value.GetValueFloat()
	case "Height":
		noteshape.Height = value.GetValueFloat()
	case "IsHidden":
		noteshape.IsHidden = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (part *Part) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		part.Name = value.GetValueString()
	case "Description":
		part.Description = value.GetValueString()
	case "ComputedPrefix":
		part.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		part.IsExpanded = value.GetValueBool()
	case "LayoutDirection":
		part.LayoutDirection.FromCodeString(value.GetValueString())
	case "IsPortsNodeExpanded":
		part.IsPortsNodeExpanded = value.GetValueBool()
	case "Ports":
		part.Ports = make([]*Port, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Ports {
					if stage.Port_stagedOrder[__instance__] == uint(id) {
						part.Ports = append(part.Ports, __instance__)
						break
					}
				}
			}
		}
	case "IsControlFlowsNodeExpanded":
		part.IsControlFlowsNodeExpanded = value.GetValueBool()
	case "ControlFlows":
		part.ControlFlows = make([]*ControlFlow, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.ControlFlows {
					if stage.ControlFlow_stagedOrder[__instance__] == uint(id) {
						part.ControlFlows = append(part.ControlFlows, __instance__)
						break
					}
				}
			}
		}
	case "PortWhoseOutControlFlowsNodeIsExpanded":
		part.PortWhoseOutControlFlowsNodeIsExpanded = make([]*Port, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Ports {
					if stage.Port_stagedOrder[__instance__] == uint(id) {
						part.PortWhoseOutControlFlowsNodeIsExpanded = append(part.PortWhoseOutControlFlowsNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "PortWhoseInControlFlowsNodeIsExpanded":
		part.PortWhoseInControlFlowsNodeIsExpanded = make([]*Port, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Ports {
					if stage.Port_stagedOrder[__instance__] == uint(id) {
						part.PortWhoseInControlFlowsNodeIsExpanded = append(part.PortWhoseInControlFlowsNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "IsDataFlowsNodeExpanded":
		part.IsDataFlowsNodeExpanded = value.GetValueBool()
	case "PortWhoseOutDataFlowsNodeIsExpanded":
		part.PortWhoseOutDataFlowsNodeIsExpanded = make([]*Port, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Ports {
					if stage.Port_stagedOrder[__instance__] == uint(id) {
						part.PortWhoseOutDataFlowsNodeIsExpanded = append(part.PortWhoseOutDataFlowsNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "PortWhoseInDataFlowsNodeIsExpanded":
		part.PortWhoseInDataFlowsNodeIsExpanded = make([]*Port, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Ports {
					if stage.Port_stagedOrder[__instance__] == uint(id) {
						part.PortWhoseInDataFlowsNodeIsExpanded = append(part.PortWhoseInDataFlowsNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "PartAnchoredPath":
		part.PartAnchoredPath = make([]*PartAnchoredPath, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.PartAnchoredPaths {
					if stage.PartAnchoredPath_stagedOrder[__instance__] == uint(id) {
						part.PartAnchoredPath = append(part.PartAnchoredPath, __instance__)
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

func (partanchoredpath *PartAnchoredPath) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		partanchoredpath.Name = value.GetValueString()
	case "Definition":
		partanchoredpath.Definition = value.GetValueString()
	case "X_Offset":
		partanchoredpath.X_Offset = value.GetValueFloat()
	case "Y_Offset":
		partanchoredpath.Y_Offset = value.GetValueFloat()
	case "RectAnchorType":
		partanchoredpath.RectAnchorType.FromCodeString(value.GetValueString())
	case "ScalePropotionnally":
		partanchoredpath.ScalePropotionnally = value.GetValueBool()
	case "AppliedScaling":
		partanchoredpath.AppliedScaling = value.GetValueFloat()
	case "Color":
		partanchoredpath.Color = value.GetValueString()
	case "FillOpacity":
		partanchoredpath.FillOpacity = value.GetValueFloat()
	case "Stroke":
		partanchoredpath.Stroke = value.GetValueString()
	case "StrokeOpacity":
		partanchoredpath.StrokeOpacity = value.GetValueFloat()
	case "StrokeWidth":
		partanchoredpath.StrokeWidth = value.GetValueFloat()
	case "StrokeDashArray":
		partanchoredpath.StrokeDashArray = value.GetValueString()
	case "StrokeDashArrayWhenSelected":
		partanchoredpath.StrokeDashArrayWhenSelected = value.GetValueString()
	case "Transform":
		partanchoredpath.Transform = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (partshape *PartShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		partshape.Name = value.GetValueString()
	case "Part":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			partshape.Part = nil
			for __instance__ := range stage.Parts {
				if stage.Part_stagedOrder[__instance__] == uint(id) {
					partshape.Part = __instance__
					break
				}
			}
		}
	case "IsExpanded":
		partshape.IsExpanded = value.GetValueBool()
	case "X":
		partshape.X = value.GetValueFloat()
	case "Y":
		partshape.Y = value.GetValueFloat()
	case "Width":
		partshape.Width = value.GetValueFloat()
	case "Height":
		partshape.Height = value.GetValueFloat()
	case "IsHidden":
		partshape.IsHidden = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (port *Port) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		port.Name = value.GetValueString()
	case "Description":
		port.Description = value.GetValueString()
	case "ComputedPrefix":
		port.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		port.IsExpanded = value.GetValueBool()
	case "LayoutDirection":
		port.LayoutDirection.FromCodeString(value.GetValueString())
	case "IsStartPort":
		port.IsStartPort = value.GetValueBool()
	case "IsEndPort":
		port.IsEndPort = value.GetValueBool()
	case "Type":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			port.Type = nil
			for __instance__ := range stage.Systems {
				if stage.System_stagedOrder[__instance__] == uint(id) {
					port.Type = __instance__
					break
				}
			}
		}
	case "IsPortNameNotSystemName":
		port.IsPortNameNotSystemName = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (portshape *PortShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		portshape.Name = value.GetValueString()
	case "Port":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			portshape.Port = nil
			for __instance__ := range stage.Ports {
				if stage.Port_stagedOrder[__instance__] == uint(id) {
					portshape.Port = __instance__
					break
				}
			}
		}
	case "IsExpanded":
		portshape.IsExpanded = value.GetValueBool()
	case "X":
		portshape.X = value.GetValueFloat()
	case "Y":
		portshape.Y = value.GetValueFloat()
	case "Width":
		portshape.Width = value.GetValueFloat()
	case "Height":
		portshape.Height = value.GetValueFloat()
	case "IsHidden":
		portshape.IsHidden = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (resource *Resource) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		resource.Name = value.GetValueString()
	case "Acronym":
		resource.Acronym = value.GetValueString()
	case "Description":
		resource.Description = value.GetValueString()
	case "ComputedPrefix":
		resource.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		resource.IsExpanded = value.GetValueBool()
	case "LayoutDirection":
		resource.LayoutDirection.FromCodeString(value.GetValueString())
	case "SVG_Path":
		resource.SVG_Path = value.GetValueString()
	case "InverseAppliedScaling":
		resource.InverseAppliedScaling = value.GetValueFloat()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (system *System) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		system.Name = value.GetValueString()
	case "Description":
		system.Description = value.GetValueString()
	case "ComputedPrefix":
		system.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		system.IsExpanded = value.GetValueBool()
	case "LayoutDirection":
		system.LayoutDirection.FromCodeString(value.GetValueString())
	case "SVG_Path":
		system.SVG_Path = value.GetValueString()
	case "InverseAppliedScaling":
		system.InverseAppliedScaling = value.GetValueFloat()
	case "DiagramStructures":
		system.DiagramStructures = make([]*DiagramStructure, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.DiagramStructures {
					if stage.DiagramStructure_stagedOrder[__instance__] == uint(id) {
						system.DiagramStructures = append(system.DiagramStructures, __instance__)
						break
					}
				}
			}
		}
	case "DiagramStructureWhoseNodeIsExpanded":
		system.DiagramStructureWhoseNodeIsExpanded = make([]*DiagramStructure, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.DiagramStructures {
					if stage.DiagramStructure_stagedOrder[__instance__] == uint(id) {
						system.DiagramStructureWhoseNodeIsExpanded = append(system.DiagramStructureWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "IsSubSystemNodeExpanded":
		system.IsSubSystemNodeExpanded = value.GetValueBool()
	case "SubSystemes":
		system.SubSystemes = make([]*System, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Systems {
					if stage.System_stagedOrder[__instance__] == uint(id) {
						system.SubSystemes = append(system.SubSystemes, __instance__)
						break
					}
				}
			}
		}
	case "Parts":
		system.Parts = make([]*Part, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Parts {
					if stage.Part_stagedOrder[__instance__] == uint(id) {
						system.Parts = append(system.Parts, __instance__)
						break
					}
				}
			}
		}
	case "PartWhoseNodeIsExpanded":
		system.PartWhoseNodeIsExpanded = make([]*Part, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Parts {
					if stage.Part_stagedOrder[__instance__] == uint(id) {
						system.PartWhoseNodeIsExpanded = append(system.PartWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "DataFlows":
		system.DataFlows = make([]*DataFlow, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.DataFlows {
					if stage.DataFlow_stagedOrder[__instance__] == uint(id) {
						system.DataFlows = append(system.DataFlows, __instance__)
						break
					}
				}
			}
		}
	case "IsDataFlowsNodeExpanded":
		system.IsDataFlowsNodeExpanded = value.GetValueBool()
	case "ExternalParts":
		system.ExternalParts = make([]*Part, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Parts {
					if stage.Part_stagedOrder[__instance__] == uint(id) {
						system.ExternalParts = append(system.ExternalParts, __instance__)
						break
					}
				}
			}
		}
	case "ExternalPartWhoseNodeIsExpanded":
		system.ExternalPartWhoseNodeIsExpanded = make([]*Part, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Parts {
					if stage.Part_stagedOrder[__instance__] == uint(id) {
						system.ExternalPartWhoseNodeIsExpanded = append(system.ExternalPartWhoseNodeIsExpanded, __instance__)
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

func (systemshape *SystemShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		systemshape.Name = value.GetValueString()
	case "System":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			systemshape.System = nil
			for __instance__ := range stage.Systems {
				if stage.System_stagedOrder[__instance__] == uint(id) {
					systemshape.System = __instance__
					break
				}
			}
		}
	case "IsExpanded":
		systemshape.IsExpanded = value.GetValueBool()
	case "X":
		systemshape.X = value.GetValueFloat()
	case "Y":
		systemshape.Y = value.GetValueFloat()
	case "Width":
		systemshape.Width = value.GetValueFloat()
	case "Height":
		systemshape.Height = value.GetValueFloat()
	case "IsHidden":
		systemshape.IsHidden = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func SetFieldStringValueFromPointer(instance GongstructIF, fieldName string, value GongFieldValue, stage *Stage) error {
	return instance.GongSetFieldValue(fieldName, value, stage)
}

// insertion point for generic get gongstruct name
func (allocatedresourceshape *AllocatedResourceShape) GongGetGongstructName() string {
	return "AllocatedResourceShape"
}

func (allocatedsystemshape *AllocatedSystemShape) GongGetGongstructName() string {
	return "AllocatedSystemShape"
}

func (controlflow *ControlFlow) GongGetGongstructName() string {
	return "ControlFlow"
}

func (controlflowshape *ControlFlowShape) GongGetGongstructName() string {
	return "ControlFlowShape"
}

func (data *Data) GongGetGongstructName() string {
	return "Data"
}

func (dataflow *DataFlow) GongGetGongstructName() string {
	return "DataFlow"
}

func (dataflowshape *DataFlowShape) GongGetGongstructName() string {
	return "DataFlowShape"
}

func (datashape *DataShape) GongGetGongstructName() string {
	return "DataShape"
}

func (diagramstructure *DiagramStructure) GongGetGongstructName() string {
	return "DiagramStructure"
}

func (externalpartshape *ExternalPartShape) GongGetGongstructName() string {
	return "ExternalPartShape"
}

func (library *Library) GongGetGongstructName() string {
	return "Library"
}

func (note *Note) GongGetGongstructName() string {
	return "Note"
}

func (noteportshape *NotePortShape) GongGetGongstructName() string {
	return "NotePortShape"
}

func (noteshape *NoteShape) GongGetGongstructName() string {
	return "NoteShape"
}

func (part *Part) GongGetGongstructName() string {
	return "Part"
}

func (partanchoredpath *PartAnchoredPath) GongGetGongstructName() string {
	return "PartAnchoredPath"
}

func (partshape *PartShape) GongGetGongstructName() string {
	return "PartShape"
}

func (port *Port) GongGetGongstructName() string {
	return "Port"
}

func (portshape *PortShape) GongGetGongstructName() string {
	return "PortShape"
}

func (resource *Resource) GongGetGongstructName() string {
	return "Resource"
}

func (system *System) GongGetGongstructName() string {
	return "System"
}

func (systemshape *SystemShape) GongGetGongstructName() string {
	return "SystemShape"
}

func GetGongstructNameFromPointer(instance GongstructIF) (res string) {
	res = instance.GongGetGongstructName()
	return
}

func (stage *Stage) ResetMapStrings() {
	// insertion point for generic get gongstruct name
	stage.AllocatedResourceShapes_mapString = make(map[string]*AllocatedResourceShape)
	for allocatedresourceshape := range stage.AllocatedResourceShapes {
		stage.AllocatedResourceShapes_mapString[allocatedresourceshape.Name] = allocatedresourceshape
	}

	stage.AllocatedSystemShapes_mapString = make(map[string]*AllocatedSystemShape)
	for allocatedsystemshape := range stage.AllocatedSystemShapes {
		stage.AllocatedSystemShapes_mapString[allocatedsystemshape.Name] = allocatedsystemshape
	}

	stage.ControlFlows_mapString = make(map[string]*ControlFlow)
	for controlflow := range stage.ControlFlows {
		stage.ControlFlows_mapString[controlflow.Name] = controlflow
	}

	stage.ControlFlowShapes_mapString = make(map[string]*ControlFlowShape)
	for controlflowshape := range stage.ControlFlowShapes {
		stage.ControlFlowShapes_mapString[controlflowshape.Name] = controlflowshape
	}

	stage.Datas_mapString = make(map[string]*Data)
	for data := range stage.Datas {
		stage.Datas_mapString[data.Name] = data
	}

	stage.DataFlows_mapString = make(map[string]*DataFlow)
	for dataflow := range stage.DataFlows {
		stage.DataFlows_mapString[dataflow.Name] = dataflow
	}

	stage.DataFlowShapes_mapString = make(map[string]*DataFlowShape)
	for dataflowshape := range stage.DataFlowShapes {
		stage.DataFlowShapes_mapString[dataflowshape.Name] = dataflowshape
	}

	stage.DataShapes_mapString = make(map[string]*DataShape)
	for datashape := range stage.DataShapes {
		stage.DataShapes_mapString[datashape.Name] = datashape
	}

	stage.DiagramStructures_mapString = make(map[string]*DiagramStructure)
	for diagramstructure := range stage.DiagramStructures {
		stage.DiagramStructures_mapString[diagramstructure.Name] = diagramstructure
	}

	stage.ExternalPartShapes_mapString = make(map[string]*ExternalPartShape)
	for externalpartshape := range stage.ExternalPartShapes {
		stage.ExternalPartShapes_mapString[externalpartshape.Name] = externalpartshape
	}

	stage.Librarys_mapString = make(map[string]*Library)
	for library := range stage.Librarys {
		stage.Librarys_mapString[library.Name] = library
	}

	stage.Notes_mapString = make(map[string]*Note)
	for note := range stage.Notes {
		stage.Notes_mapString[note.Name] = note
	}

	stage.NotePortShapes_mapString = make(map[string]*NotePortShape)
	for noteportshape := range stage.NotePortShapes {
		stage.NotePortShapes_mapString[noteportshape.Name] = noteportshape
	}

	stage.NoteShapes_mapString = make(map[string]*NoteShape)
	for noteshape := range stage.NoteShapes {
		stage.NoteShapes_mapString[noteshape.Name] = noteshape
	}

	stage.Parts_mapString = make(map[string]*Part)
	for part := range stage.Parts {
		stage.Parts_mapString[part.Name] = part
	}

	stage.PartAnchoredPaths_mapString = make(map[string]*PartAnchoredPath)
	for partanchoredpath := range stage.PartAnchoredPaths {
		stage.PartAnchoredPaths_mapString[partanchoredpath.Name] = partanchoredpath
	}

	stage.PartShapes_mapString = make(map[string]*PartShape)
	for partshape := range stage.PartShapes {
		stage.PartShapes_mapString[partshape.Name] = partshape
	}

	stage.Ports_mapString = make(map[string]*Port)
	for port := range stage.Ports {
		stage.Ports_mapString[port.Name] = port
	}

	stage.PortShapes_mapString = make(map[string]*PortShape)
	for portshape := range stage.PortShapes {
		stage.PortShapes_mapString[portshape.Name] = portshape
	}

	stage.Resources_mapString = make(map[string]*Resource)
	for resource := range stage.Resources {
		stage.Resources_mapString[resource.Name] = resource
	}

	stage.Systems_mapString = make(map[string]*System)
	for system := range stage.Systems {
		stage.Systems_mapString[system.Name] = system
	}

	stage.SystemShapes_mapString = make(map[string]*SystemShape)
	for systemshape := range stage.SystemShapes {
		stage.SystemShapes_mapString[systemshape.Name] = systemshape
	}

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
