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

	xsd_go "github.com/fullstack-lang/gong/app/xsd/go"
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
	Alls                map[*All]struct{}
	Alls_instance       map[*All]*All
	Alls_mapString      map[string]*All
	AllOrder            uint
	All_stagedOrder     map[*All]uint
	All_orderStaged     map[uint]*All
	Alls_reference      map[*All]*All
	Alls_referenceOrder map[*All]uint

	// insertion point for slice of pointers maps
	All_Sequences_reverseMap map[*Sequence]*All

	All_Alls_reverseMap map[*All]*All

	All_Choices_reverseMap map[*Choice]*All

	All_Groups_reverseMap map[*Group]*All

	All_Elements_reverseMap map[*Element]*All

	OnAfterAllCreateCallback OnAfterCreateInterface[All]
	OnAfterAllUpdateCallback OnAfterUpdateInterface[All]
	OnAfterAllDeleteCallback OnAfterDeleteInterface[All]
	OnAfterAllReadCallback   OnAfterReadInterface[All]

	Annotations                map[*Annotation]struct{}
	Annotations_instance       map[*Annotation]*Annotation
	Annotations_mapString      map[string]*Annotation
	AnnotationOrder            uint
	Annotation_stagedOrder     map[*Annotation]uint
	Annotation_orderStaged     map[uint]*Annotation
	Annotations_reference      map[*Annotation]*Annotation
	Annotations_referenceOrder map[*Annotation]uint

	// insertion point for slice of pointers maps
	Annotation_Documentations_reverseMap map[*Documentation]*Annotation

	OnAfterAnnotationCreateCallback OnAfterCreateInterface[Annotation]
	OnAfterAnnotationUpdateCallback OnAfterUpdateInterface[Annotation]
	OnAfterAnnotationDeleteCallback OnAfterDeleteInterface[Annotation]
	OnAfterAnnotationReadCallback   OnAfterReadInterface[Annotation]

	Attributes                map[*Attribute]struct{}
	Attributes_instance       map[*Attribute]*Attribute
	Attributes_mapString      map[string]*Attribute
	AttributeOrder            uint
	Attribute_stagedOrder     map[*Attribute]uint
	Attribute_orderStaged     map[uint]*Attribute
	Attributes_reference      map[*Attribute]*Attribute
	Attributes_referenceOrder map[*Attribute]uint

	// insertion point for slice of pointers maps
	OnAfterAttributeCreateCallback OnAfterCreateInterface[Attribute]
	OnAfterAttributeUpdateCallback OnAfterUpdateInterface[Attribute]
	OnAfterAttributeDeleteCallback OnAfterDeleteInterface[Attribute]
	OnAfterAttributeReadCallback   OnAfterReadInterface[Attribute]

	AttributeGroups                map[*AttributeGroup]struct{}
	AttributeGroups_instance       map[*AttributeGroup]*AttributeGroup
	AttributeGroups_mapString      map[string]*AttributeGroup
	AttributeGroupOrder            uint
	AttributeGroup_stagedOrder     map[*AttributeGroup]uint
	AttributeGroup_orderStaged     map[uint]*AttributeGroup
	AttributeGroups_reference      map[*AttributeGroup]*AttributeGroup
	AttributeGroups_referenceOrder map[*AttributeGroup]uint

	// insertion point for slice of pointers maps
	AttributeGroup_AttributeGroups_reverseMap map[*AttributeGroup]*AttributeGroup

	AttributeGroup_Attributes_reverseMap map[*Attribute]*AttributeGroup

	OnAfterAttributeGroupCreateCallback OnAfterCreateInterface[AttributeGroup]
	OnAfterAttributeGroupUpdateCallback OnAfterUpdateInterface[AttributeGroup]
	OnAfterAttributeGroupDeleteCallback OnAfterDeleteInterface[AttributeGroup]
	OnAfterAttributeGroupReadCallback   OnAfterReadInterface[AttributeGroup]

	Choices                map[*Choice]struct{}
	Choices_instance       map[*Choice]*Choice
	Choices_mapString      map[string]*Choice
	ChoiceOrder            uint
	Choice_stagedOrder     map[*Choice]uint
	Choice_orderStaged     map[uint]*Choice
	Choices_reference      map[*Choice]*Choice
	Choices_referenceOrder map[*Choice]uint

	// insertion point for slice of pointers maps
	Choice_Sequences_reverseMap map[*Sequence]*Choice

	Choice_Alls_reverseMap map[*All]*Choice

	Choice_Choices_reverseMap map[*Choice]*Choice

	Choice_Groups_reverseMap map[*Group]*Choice

	Choice_Elements_reverseMap map[*Element]*Choice

	OnAfterChoiceCreateCallback OnAfterCreateInterface[Choice]
	OnAfterChoiceUpdateCallback OnAfterUpdateInterface[Choice]
	OnAfterChoiceDeleteCallback OnAfterDeleteInterface[Choice]
	OnAfterChoiceReadCallback   OnAfterReadInterface[Choice]

	ComplexContents                map[*ComplexContent]struct{}
	ComplexContents_instance       map[*ComplexContent]*ComplexContent
	ComplexContents_mapString      map[string]*ComplexContent
	ComplexContentOrder            uint
	ComplexContent_stagedOrder     map[*ComplexContent]uint
	ComplexContent_orderStaged     map[uint]*ComplexContent
	ComplexContents_reference      map[*ComplexContent]*ComplexContent
	ComplexContents_referenceOrder map[*ComplexContent]uint

	// insertion point for slice of pointers maps
	OnAfterComplexContentCreateCallback OnAfterCreateInterface[ComplexContent]
	OnAfterComplexContentUpdateCallback OnAfterUpdateInterface[ComplexContent]
	OnAfterComplexContentDeleteCallback OnAfterDeleteInterface[ComplexContent]
	OnAfterComplexContentReadCallback   OnAfterReadInterface[ComplexContent]

	ComplexTypes                map[*ComplexType]struct{}
	ComplexTypes_instance       map[*ComplexType]*ComplexType
	ComplexTypes_mapString      map[string]*ComplexType
	ComplexTypeOrder            uint
	ComplexType_stagedOrder     map[*ComplexType]uint
	ComplexType_orderStaged     map[uint]*ComplexType
	ComplexTypes_reference      map[*ComplexType]*ComplexType
	ComplexTypes_referenceOrder map[*ComplexType]uint

	// insertion point for slice of pointers maps
	ComplexType_Sequences_reverseMap map[*Sequence]*ComplexType

	ComplexType_Alls_reverseMap map[*All]*ComplexType

	ComplexType_Choices_reverseMap map[*Choice]*ComplexType

	ComplexType_Groups_reverseMap map[*Group]*ComplexType

	ComplexType_Elements_reverseMap map[*Element]*ComplexType

	ComplexType_Attributes_reverseMap map[*Attribute]*ComplexType

	ComplexType_AttributeGroups_reverseMap map[*AttributeGroup]*ComplexType

	OnAfterComplexTypeCreateCallback OnAfterCreateInterface[ComplexType]
	OnAfterComplexTypeUpdateCallback OnAfterUpdateInterface[ComplexType]
	OnAfterComplexTypeDeleteCallback OnAfterDeleteInterface[ComplexType]
	OnAfterComplexTypeReadCallback   OnAfterReadInterface[ComplexType]

	Documentations                map[*Documentation]struct{}
	Documentations_instance       map[*Documentation]*Documentation
	Documentations_mapString      map[string]*Documentation
	DocumentationOrder            uint
	Documentation_stagedOrder     map[*Documentation]uint
	Documentation_orderStaged     map[uint]*Documentation
	Documentations_reference      map[*Documentation]*Documentation
	Documentations_referenceOrder map[*Documentation]uint

	// insertion point for slice of pointers maps
	OnAfterDocumentationCreateCallback OnAfterCreateInterface[Documentation]
	OnAfterDocumentationUpdateCallback OnAfterUpdateInterface[Documentation]
	OnAfterDocumentationDeleteCallback OnAfterDeleteInterface[Documentation]
	OnAfterDocumentationReadCallback   OnAfterReadInterface[Documentation]

	Elements                map[*Element]struct{}
	Elements_instance       map[*Element]*Element
	Elements_mapString      map[string]*Element
	ElementOrder            uint
	Element_stagedOrder     map[*Element]uint
	Element_orderStaged     map[uint]*Element
	Elements_reference      map[*Element]*Element
	Elements_referenceOrder map[*Element]uint

	// insertion point for slice of pointers maps
	Element_Groups_reverseMap map[*Group]*Element

	OnAfterElementCreateCallback OnAfterCreateInterface[Element]
	OnAfterElementUpdateCallback OnAfterUpdateInterface[Element]
	OnAfterElementDeleteCallback OnAfterDeleteInterface[Element]
	OnAfterElementReadCallback   OnAfterReadInterface[Element]

	Enumerations                map[*Enumeration]struct{}
	Enumerations_instance       map[*Enumeration]*Enumeration
	Enumerations_mapString      map[string]*Enumeration
	EnumerationOrder            uint
	Enumeration_stagedOrder     map[*Enumeration]uint
	Enumeration_orderStaged     map[uint]*Enumeration
	Enumerations_reference      map[*Enumeration]*Enumeration
	Enumerations_referenceOrder map[*Enumeration]uint

	// insertion point for slice of pointers maps
	OnAfterEnumerationCreateCallback OnAfterCreateInterface[Enumeration]
	OnAfterEnumerationUpdateCallback OnAfterUpdateInterface[Enumeration]
	OnAfterEnumerationDeleteCallback OnAfterDeleteInterface[Enumeration]
	OnAfterEnumerationReadCallback   OnAfterReadInterface[Enumeration]

	Extensions                map[*Extension]struct{}
	Extensions_instance       map[*Extension]*Extension
	Extensions_mapString      map[string]*Extension
	ExtensionOrder            uint
	Extension_stagedOrder     map[*Extension]uint
	Extension_orderStaged     map[uint]*Extension
	Extensions_reference      map[*Extension]*Extension
	Extensions_referenceOrder map[*Extension]uint

	// insertion point for slice of pointers maps
	Extension_Sequences_reverseMap map[*Sequence]*Extension

	Extension_Alls_reverseMap map[*All]*Extension

	Extension_Choices_reverseMap map[*Choice]*Extension

	Extension_Groups_reverseMap map[*Group]*Extension

	Extension_Elements_reverseMap map[*Element]*Extension

	Extension_Attributes_reverseMap map[*Attribute]*Extension

	Extension_AttributeGroups_reverseMap map[*AttributeGroup]*Extension

	OnAfterExtensionCreateCallback OnAfterCreateInterface[Extension]
	OnAfterExtensionUpdateCallback OnAfterUpdateInterface[Extension]
	OnAfterExtensionDeleteCallback OnAfterDeleteInterface[Extension]
	OnAfterExtensionReadCallback   OnAfterReadInterface[Extension]

	Groups                map[*Group]struct{}
	Groups_instance       map[*Group]*Group
	Groups_mapString      map[string]*Group
	GroupOrder            uint
	Group_stagedOrder     map[*Group]uint
	Group_orderStaged     map[uint]*Group
	Groups_reference      map[*Group]*Group
	Groups_referenceOrder map[*Group]uint

	// insertion point for slice of pointers maps
	Group_Sequences_reverseMap map[*Sequence]*Group

	Group_Alls_reverseMap map[*All]*Group

	Group_Choices_reverseMap map[*Choice]*Group

	Group_Groups_reverseMap map[*Group]*Group

	Group_Elements_reverseMap map[*Element]*Group

	OnAfterGroupCreateCallback OnAfterCreateInterface[Group]
	OnAfterGroupUpdateCallback OnAfterUpdateInterface[Group]
	OnAfterGroupDeleteCallback OnAfterDeleteInterface[Group]
	OnAfterGroupReadCallback   OnAfterReadInterface[Group]

	Lengths                map[*Length]struct{}
	Lengths_instance       map[*Length]*Length
	Lengths_mapString      map[string]*Length
	LengthOrder            uint
	Length_stagedOrder     map[*Length]uint
	Length_orderStaged     map[uint]*Length
	Lengths_reference      map[*Length]*Length
	Lengths_referenceOrder map[*Length]uint

	// insertion point for slice of pointers maps
	OnAfterLengthCreateCallback OnAfterCreateInterface[Length]
	OnAfterLengthUpdateCallback OnAfterUpdateInterface[Length]
	OnAfterLengthDeleteCallback OnAfterDeleteInterface[Length]
	OnAfterLengthReadCallback   OnAfterReadInterface[Length]

	MaxInclusives                map[*MaxInclusive]struct{}
	MaxInclusives_instance       map[*MaxInclusive]*MaxInclusive
	MaxInclusives_mapString      map[string]*MaxInclusive
	MaxInclusiveOrder            uint
	MaxInclusive_stagedOrder     map[*MaxInclusive]uint
	MaxInclusive_orderStaged     map[uint]*MaxInclusive
	MaxInclusives_reference      map[*MaxInclusive]*MaxInclusive
	MaxInclusives_referenceOrder map[*MaxInclusive]uint

	// insertion point for slice of pointers maps
	OnAfterMaxInclusiveCreateCallback OnAfterCreateInterface[MaxInclusive]
	OnAfterMaxInclusiveUpdateCallback OnAfterUpdateInterface[MaxInclusive]
	OnAfterMaxInclusiveDeleteCallback OnAfterDeleteInterface[MaxInclusive]
	OnAfterMaxInclusiveReadCallback   OnAfterReadInterface[MaxInclusive]

	MaxLengths                map[*MaxLength]struct{}
	MaxLengths_instance       map[*MaxLength]*MaxLength
	MaxLengths_mapString      map[string]*MaxLength
	MaxLengthOrder            uint
	MaxLength_stagedOrder     map[*MaxLength]uint
	MaxLength_orderStaged     map[uint]*MaxLength
	MaxLengths_reference      map[*MaxLength]*MaxLength
	MaxLengths_referenceOrder map[*MaxLength]uint

	// insertion point for slice of pointers maps
	OnAfterMaxLengthCreateCallback OnAfterCreateInterface[MaxLength]
	OnAfterMaxLengthUpdateCallback OnAfterUpdateInterface[MaxLength]
	OnAfterMaxLengthDeleteCallback OnAfterDeleteInterface[MaxLength]
	OnAfterMaxLengthReadCallback   OnAfterReadInterface[MaxLength]

	MinInclusives                map[*MinInclusive]struct{}
	MinInclusives_instance       map[*MinInclusive]*MinInclusive
	MinInclusives_mapString      map[string]*MinInclusive
	MinInclusiveOrder            uint
	MinInclusive_stagedOrder     map[*MinInclusive]uint
	MinInclusive_orderStaged     map[uint]*MinInclusive
	MinInclusives_reference      map[*MinInclusive]*MinInclusive
	MinInclusives_referenceOrder map[*MinInclusive]uint

	// insertion point for slice of pointers maps
	OnAfterMinInclusiveCreateCallback OnAfterCreateInterface[MinInclusive]
	OnAfterMinInclusiveUpdateCallback OnAfterUpdateInterface[MinInclusive]
	OnAfterMinInclusiveDeleteCallback OnAfterDeleteInterface[MinInclusive]
	OnAfterMinInclusiveReadCallback   OnAfterReadInterface[MinInclusive]

	MinLengths                map[*MinLength]struct{}
	MinLengths_instance       map[*MinLength]*MinLength
	MinLengths_mapString      map[string]*MinLength
	MinLengthOrder            uint
	MinLength_stagedOrder     map[*MinLength]uint
	MinLength_orderStaged     map[uint]*MinLength
	MinLengths_reference      map[*MinLength]*MinLength
	MinLengths_referenceOrder map[*MinLength]uint

	// insertion point for slice of pointers maps
	OnAfterMinLengthCreateCallback OnAfterCreateInterface[MinLength]
	OnAfterMinLengthUpdateCallback OnAfterUpdateInterface[MinLength]
	OnAfterMinLengthDeleteCallback OnAfterDeleteInterface[MinLength]
	OnAfterMinLengthReadCallback   OnAfterReadInterface[MinLength]

	Patterns                map[*Pattern]struct{}
	Patterns_instance       map[*Pattern]*Pattern
	Patterns_mapString      map[string]*Pattern
	PatternOrder            uint
	Pattern_stagedOrder     map[*Pattern]uint
	Pattern_orderStaged     map[uint]*Pattern
	Patterns_reference      map[*Pattern]*Pattern
	Patterns_referenceOrder map[*Pattern]uint

	// insertion point for slice of pointers maps
	OnAfterPatternCreateCallback OnAfterCreateInterface[Pattern]
	OnAfterPatternUpdateCallback OnAfterUpdateInterface[Pattern]
	OnAfterPatternDeleteCallback OnAfterDeleteInterface[Pattern]
	OnAfterPatternReadCallback   OnAfterReadInterface[Pattern]

	Restrictions                map[*Restriction]struct{}
	Restrictions_instance       map[*Restriction]*Restriction
	Restrictions_mapString      map[string]*Restriction
	RestrictionOrder            uint
	Restriction_stagedOrder     map[*Restriction]uint
	Restriction_orderStaged     map[uint]*Restriction
	Restrictions_reference      map[*Restriction]*Restriction
	Restrictions_referenceOrder map[*Restriction]uint

	// insertion point for slice of pointers maps
	Restriction_Enumerations_reverseMap map[*Enumeration]*Restriction

	OnAfterRestrictionCreateCallback OnAfterCreateInterface[Restriction]
	OnAfterRestrictionUpdateCallback OnAfterUpdateInterface[Restriction]
	OnAfterRestrictionDeleteCallback OnAfterDeleteInterface[Restriction]
	OnAfterRestrictionReadCallback   OnAfterReadInterface[Restriction]

	Schemas                map[*Schema]struct{}
	Schemas_instance       map[*Schema]*Schema
	Schemas_mapString      map[string]*Schema
	SchemaOrder            uint
	Schema_stagedOrder     map[*Schema]uint
	Schema_orderStaged     map[uint]*Schema
	Schemas_reference      map[*Schema]*Schema
	Schemas_referenceOrder map[*Schema]uint

	// insertion point for slice of pointers maps
	Schema_Elements_reverseMap map[*Element]*Schema

	Schema_SimpleTypes_reverseMap map[*SimpleType]*Schema

	Schema_ComplexTypes_reverseMap map[*ComplexType]*Schema

	Schema_AttributeGroups_reverseMap map[*AttributeGroup]*Schema

	Schema_Groups_reverseMap map[*Group]*Schema

	OnAfterSchemaCreateCallback OnAfterCreateInterface[Schema]
	OnAfterSchemaUpdateCallback OnAfterUpdateInterface[Schema]
	OnAfterSchemaDeleteCallback OnAfterDeleteInterface[Schema]
	OnAfterSchemaReadCallback   OnAfterReadInterface[Schema]

	Sequences                map[*Sequence]struct{}
	Sequences_instance       map[*Sequence]*Sequence
	Sequences_mapString      map[string]*Sequence
	SequenceOrder            uint
	Sequence_stagedOrder     map[*Sequence]uint
	Sequence_orderStaged     map[uint]*Sequence
	Sequences_reference      map[*Sequence]*Sequence
	Sequences_referenceOrder map[*Sequence]uint

	// insertion point for slice of pointers maps
	Sequence_Sequences_reverseMap map[*Sequence]*Sequence

	Sequence_Alls_reverseMap map[*All]*Sequence

	Sequence_Choices_reverseMap map[*Choice]*Sequence

	Sequence_Groups_reverseMap map[*Group]*Sequence

	Sequence_Elements_reverseMap map[*Element]*Sequence

	OnAfterSequenceCreateCallback OnAfterCreateInterface[Sequence]
	OnAfterSequenceUpdateCallback OnAfterUpdateInterface[Sequence]
	OnAfterSequenceDeleteCallback OnAfterDeleteInterface[Sequence]
	OnAfterSequenceReadCallback   OnAfterReadInterface[Sequence]

	SimpleContents                map[*SimpleContent]struct{}
	SimpleContents_instance       map[*SimpleContent]*SimpleContent
	SimpleContents_mapString      map[string]*SimpleContent
	SimpleContentOrder            uint
	SimpleContent_stagedOrder     map[*SimpleContent]uint
	SimpleContent_orderStaged     map[uint]*SimpleContent
	SimpleContents_reference      map[*SimpleContent]*SimpleContent
	SimpleContents_referenceOrder map[*SimpleContent]uint

	// insertion point for slice of pointers maps
	OnAfterSimpleContentCreateCallback OnAfterCreateInterface[SimpleContent]
	OnAfterSimpleContentUpdateCallback OnAfterUpdateInterface[SimpleContent]
	OnAfterSimpleContentDeleteCallback OnAfterDeleteInterface[SimpleContent]
	OnAfterSimpleContentReadCallback   OnAfterReadInterface[SimpleContent]

	SimpleTypes                map[*SimpleType]struct{}
	SimpleTypes_instance       map[*SimpleType]*SimpleType
	SimpleTypes_mapString      map[string]*SimpleType
	SimpleTypeOrder            uint
	SimpleType_stagedOrder     map[*SimpleType]uint
	SimpleType_orderStaged     map[uint]*SimpleType
	SimpleTypes_reference      map[*SimpleType]*SimpleType
	SimpleTypes_referenceOrder map[*SimpleType]uint

	// insertion point for slice of pointers maps
	OnAfterSimpleTypeCreateCallback OnAfterCreateInterface[SimpleType]
	OnAfterSimpleTypeUpdateCallback OnAfterUpdateInterface[SimpleType]
	OnAfterSimpleTypeDeleteCallback OnAfterDeleteInterface[SimpleType]
	OnAfterSimpleTypeReadCallback   OnAfterReadInterface[SimpleType]

	TotalDigits                map[*TotalDigit]struct{}
	TotalDigits_instance       map[*TotalDigit]*TotalDigit
	TotalDigits_mapString      map[string]*TotalDigit
	TotalDigitOrder            uint
	TotalDigit_stagedOrder     map[*TotalDigit]uint
	TotalDigit_orderStaged     map[uint]*TotalDigit
	TotalDigits_reference      map[*TotalDigit]*TotalDigit
	TotalDigits_referenceOrder map[*TotalDigit]uint

	// insertion point for slice of pointers maps
	OnAfterTotalDigitCreateCallback OnAfterCreateInterface[TotalDigit]
	OnAfterTotalDigitUpdateCallback OnAfterUpdateInterface[TotalDigit]
	OnAfterTotalDigitDeleteCallback OnAfterDeleteInterface[TotalDigit]
	OnAfterTotalDigitReadCallback   OnAfterReadInterface[TotalDigit]

	Unions                map[*Union]struct{}
	Unions_instance       map[*Union]*Union
	Unions_mapString      map[string]*Union
	UnionOrder            uint
	Union_stagedOrder     map[*Union]uint
	Union_orderStaged     map[uint]*Union
	Unions_reference      map[*Union]*Union
	Unions_referenceOrder map[*Union]uint

	// insertion point for slice of pointers maps
	OnAfterUnionCreateCallback OnAfterCreateInterface[Union]
	OnAfterUnionUpdateCallback OnAfterUpdateInterface[Union]
	OnAfterUnionDeleteCallback OnAfterDeleteInterface[Union]
	OnAfterUnionReadCallback   OnAfterReadInterface[Union]

	WhiteSpaces                map[*WhiteSpace]struct{}
	WhiteSpaces_instance       map[*WhiteSpace]*WhiteSpace
	WhiteSpaces_mapString      map[string]*WhiteSpace
	WhiteSpaceOrder            uint
	WhiteSpace_stagedOrder     map[*WhiteSpace]uint
	WhiteSpace_orderStaged     map[uint]*WhiteSpace
	WhiteSpaces_reference      map[*WhiteSpace]*WhiteSpace
	WhiteSpaces_referenceOrder map[*WhiteSpace]uint

	// insertion point for slice of pointers maps
	OnAfterWhiteSpaceCreateCallback OnAfterCreateInterface[WhiteSpace]
	OnAfterWhiteSpaceUpdateCallback OnAfterUpdateInterface[WhiteSpace]
	OnAfterWhiteSpaceDeleteCallback OnAfterDeleteInterface[WhiteSpace]
	OnAfterWhiteSpaceReadCallback   OnAfterReadInterface[WhiteSpace]

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
	stage.Alls_reference = make(map[*All]*All)
	stage.Alls_instance = make(map[*All]*All)
	stage.Alls_referenceOrder = make(map[*All]uint)

	stage.Annotations_reference = make(map[*Annotation]*Annotation)
	stage.Annotations_instance = make(map[*Annotation]*Annotation)
	stage.Annotations_referenceOrder = make(map[*Annotation]uint)

	stage.Attributes_reference = make(map[*Attribute]*Attribute)
	stage.Attributes_instance = make(map[*Attribute]*Attribute)
	stage.Attributes_referenceOrder = make(map[*Attribute]uint)

	stage.AttributeGroups_reference = make(map[*AttributeGroup]*AttributeGroup)
	stage.AttributeGroups_instance = make(map[*AttributeGroup]*AttributeGroup)
	stage.AttributeGroups_referenceOrder = make(map[*AttributeGroup]uint)

	stage.Choices_reference = make(map[*Choice]*Choice)
	stage.Choices_instance = make(map[*Choice]*Choice)
	stage.Choices_referenceOrder = make(map[*Choice]uint)

	stage.ComplexContents_reference = make(map[*ComplexContent]*ComplexContent)
	stage.ComplexContents_instance = make(map[*ComplexContent]*ComplexContent)
	stage.ComplexContents_referenceOrder = make(map[*ComplexContent]uint)

	stage.ComplexTypes_reference = make(map[*ComplexType]*ComplexType)
	stage.ComplexTypes_instance = make(map[*ComplexType]*ComplexType)
	stage.ComplexTypes_referenceOrder = make(map[*ComplexType]uint)

	stage.Documentations_reference = make(map[*Documentation]*Documentation)
	stage.Documentations_instance = make(map[*Documentation]*Documentation)
	stage.Documentations_referenceOrder = make(map[*Documentation]uint)

	stage.Elements_reference = make(map[*Element]*Element)
	stage.Elements_instance = make(map[*Element]*Element)
	stage.Elements_referenceOrder = make(map[*Element]uint)

	stage.Enumerations_reference = make(map[*Enumeration]*Enumeration)
	stage.Enumerations_instance = make(map[*Enumeration]*Enumeration)
	stage.Enumerations_referenceOrder = make(map[*Enumeration]uint)

	stage.Extensions_reference = make(map[*Extension]*Extension)
	stage.Extensions_instance = make(map[*Extension]*Extension)
	stage.Extensions_referenceOrder = make(map[*Extension]uint)

	stage.Groups_reference = make(map[*Group]*Group)
	stage.Groups_instance = make(map[*Group]*Group)
	stage.Groups_referenceOrder = make(map[*Group]uint)

	stage.Lengths_reference = make(map[*Length]*Length)
	stage.Lengths_instance = make(map[*Length]*Length)
	stage.Lengths_referenceOrder = make(map[*Length]uint)

	stage.MaxInclusives_reference = make(map[*MaxInclusive]*MaxInclusive)
	stage.MaxInclusives_instance = make(map[*MaxInclusive]*MaxInclusive)
	stage.MaxInclusives_referenceOrder = make(map[*MaxInclusive]uint)

	stage.MaxLengths_reference = make(map[*MaxLength]*MaxLength)
	stage.MaxLengths_instance = make(map[*MaxLength]*MaxLength)
	stage.MaxLengths_referenceOrder = make(map[*MaxLength]uint)

	stage.MinInclusives_reference = make(map[*MinInclusive]*MinInclusive)
	stage.MinInclusives_instance = make(map[*MinInclusive]*MinInclusive)
	stage.MinInclusives_referenceOrder = make(map[*MinInclusive]uint)

	stage.MinLengths_reference = make(map[*MinLength]*MinLength)
	stage.MinLengths_instance = make(map[*MinLength]*MinLength)
	stage.MinLengths_referenceOrder = make(map[*MinLength]uint)

	stage.Patterns_reference = make(map[*Pattern]*Pattern)
	stage.Patterns_instance = make(map[*Pattern]*Pattern)
	stage.Patterns_referenceOrder = make(map[*Pattern]uint)

	stage.Restrictions_reference = make(map[*Restriction]*Restriction)
	stage.Restrictions_instance = make(map[*Restriction]*Restriction)
	stage.Restrictions_referenceOrder = make(map[*Restriction]uint)

	stage.Schemas_reference = make(map[*Schema]*Schema)
	stage.Schemas_instance = make(map[*Schema]*Schema)
	stage.Schemas_referenceOrder = make(map[*Schema]uint)

	stage.Sequences_reference = make(map[*Sequence]*Sequence)
	stage.Sequences_instance = make(map[*Sequence]*Sequence)
	stage.Sequences_referenceOrder = make(map[*Sequence]uint)

	stage.SimpleContents_reference = make(map[*SimpleContent]*SimpleContent)
	stage.SimpleContents_instance = make(map[*SimpleContent]*SimpleContent)
	stage.SimpleContents_referenceOrder = make(map[*SimpleContent]uint)

	stage.SimpleTypes_reference = make(map[*SimpleType]*SimpleType)
	stage.SimpleTypes_instance = make(map[*SimpleType]*SimpleType)
	stage.SimpleTypes_referenceOrder = make(map[*SimpleType]uint)

	stage.TotalDigits_reference = make(map[*TotalDigit]*TotalDigit)
	stage.TotalDigits_instance = make(map[*TotalDigit]*TotalDigit)
	stage.TotalDigits_referenceOrder = make(map[*TotalDigit]uint)

	stage.Unions_reference = make(map[*Union]*Union)
	stage.Unions_instance = make(map[*Union]*Union)
	stage.Unions_referenceOrder = make(map[*Union]uint)

	stage.WhiteSpaces_reference = make(map[*WhiteSpace]*WhiteSpace)
	stage.WhiteSpaces_instance = make(map[*WhiteSpace]*WhiteSpace)
	stage.WhiteSpaces_referenceOrder = make(map[*WhiteSpace]uint)

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
	var maxAllOrder uint
	var foundAll bool
	for _, order := range stage.All_stagedOrder {
		if !foundAll || order > maxAllOrder {
			maxAllOrder = order
			foundAll = true
		}
	}
	if foundAll {
		stage.AllOrder = maxAllOrder + 1
	} else {
		stage.AllOrder = 0
	}

	var maxAnnotationOrder uint
	var foundAnnotation bool
	for _, order := range stage.Annotation_stagedOrder {
		if !foundAnnotation || order > maxAnnotationOrder {
			maxAnnotationOrder = order
			foundAnnotation = true
		}
	}
	if foundAnnotation {
		stage.AnnotationOrder = maxAnnotationOrder + 1
	} else {
		stage.AnnotationOrder = 0
	}

	var maxAttributeOrder uint
	var foundAttribute bool
	for _, order := range stage.Attribute_stagedOrder {
		if !foundAttribute || order > maxAttributeOrder {
			maxAttributeOrder = order
			foundAttribute = true
		}
	}
	if foundAttribute {
		stage.AttributeOrder = maxAttributeOrder + 1
	} else {
		stage.AttributeOrder = 0
	}

	var maxAttributeGroupOrder uint
	var foundAttributeGroup bool
	for _, order := range stage.AttributeGroup_stagedOrder {
		if !foundAttributeGroup || order > maxAttributeGroupOrder {
			maxAttributeGroupOrder = order
			foundAttributeGroup = true
		}
	}
	if foundAttributeGroup {
		stage.AttributeGroupOrder = maxAttributeGroupOrder + 1
	} else {
		stage.AttributeGroupOrder = 0
	}

	var maxChoiceOrder uint
	var foundChoice bool
	for _, order := range stage.Choice_stagedOrder {
		if !foundChoice || order > maxChoiceOrder {
			maxChoiceOrder = order
			foundChoice = true
		}
	}
	if foundChoice {
		stage.ChoiceOrder = maxChoiceOrder + 1
	} else {
		stage.ChoiceOrder = 0
	}

	var maxComplexContentOrder uint
	var foundComplexContent bool
	for _, order := range stage.ComplexContent_stagedOrder {
		if !foundComplexContent || order > maxComplexContentOrder {
			maxComplexContentOrder = order
			foundComplexContent = true
		}
	}
	if foundComplexContent {
		stage.ComplexContentOrder = maxComplexContentOrder + 1
	} else {
		stage.ComplexContentOrder = 0
	}

	var maxComplexTypeOrder uint
	var foundComplexType bool
	for _, order := range stage.ComplexType_stagedOrder {
		if !foundComplexType || order > maxComplexTypeOrder {
			maxComplexTypeOrder = order
			foundComplexType = true
		}
	}
	if foundComplexType {
		stage.ComplexTypeOrder = maxComplexTypeOrder + 1
	} else {
		stage.ComplexTypeOrder = 0
	}

	var maxDocumentationOrder uint
	var foundDocumentation bool
	for _, order := range stage.Documentation_stagedOrder {
		if !foundDocumentation || order > maxDocumentationOrder {
			maxDocumentationOrder = order
			foundDocumentation = true
		}
	}
	if foundDocumentation {
		stage.DocumentationOrder = maxDocumentationOrder + 1
	} else {
		stage.DocumentationOrder = 0
	}

	var maxElementOrder uint
	var foundElement bool
	for _, order := range stage.Element_stagedOrder {
		if !foundElement || order > maxElementOrder {
			maxElementOrder = order
			foundElement = true
		}
	}
	if foundElement {
		stage.ElementOrder = maxElementOrder + 1
	} else {
		stage.ElementOrder = 0
	}

	var maxEnumerationOrder uint
	var foundEnumeration bool
	for _, order := range stage.Enumeration_stagedOrder {
		if !foundEnumeration || order > maxEnumerationOrder {
			maxEnumerationOrder = order
			foundEnumeration = true
		}
	}
	if foundEnumeration {
		stage.EnumerationOrder = maxEnumerationOrder + 1
	} else {
		stage.EnumerationOrder = 0
	}

	var maxExtensionOrder uint
	var foundExtension bool
	for _, order := range stage.Extension_stagedOrder {
		if !foundExtension || order > maxExtensionOrder {
			maxExtensionOrder = order
			foundExtension = true
		}
	}
	if foundExtension {
		stage.ExtensionOrder = maxExtensionOrder + 1
	} else {
		stage.ExtensionOrder = 0
	}

	var maxGroupOrder uint
	var foundGroup bool
	for _, order := range stage.Group_stagedOrder {
		if !foundGroup || order > maxGroupOrder {
			maxGroupOrder = order
			foundGroup = true
		}
	}
	if foundGroup {
		stage.GroupOrder = maxGroupOrder + 1
	} else {
		stage.GroupOrder = 0
	}

	var maxLengthOrder uint
	var foundLength bool
	for _, order := range stage.Length_stagedOrder {
		if !foundLength || order > maxLengthOrder {
			maxLengthOrder = order
			foundLength = true
		}
	}
	if foundLength {
		stage.LengthOrder = maxLengthOrder + 1
	} else {
		stage.LengthOrder = 0
	}

	var maxMaxInclusiveOrder uint
	var foundMaxInclusive bool
	for _, order := range stage.MaxInclusive_stagedOrder {
		if !foundMaxInclusive || order > maxMaxInclusiveOrder {
			maxMaxInclusiveOrder = order
			foundMaxInclusive = true
		}
	}
	if foundMaxInclusive {
		stage.MaxInclusiveOrder = maxMaxInclusiveOrder + 1
	} else {
		stage.MaxInclusiveOrder = 0
	}

	var maxMaxLengthOrder uint
	var foundMaxLength bool
	for _, order := range stage.MaxLength_stagedOrder {
		if !foundMaxLength || order > maxMaxLengthOrder {
			maxMaxLengthOrder = order
			foundMaxLength = true
		}
	}
	if foundMaxLength {
		stage.MaxLengthOrder = maxMaxLengthOrder + 1
	} else {
		stage.MaxLengthOrder = 0
	}

	var maxMinInclusiveOrder uint
	var foundMinInclusive bool
	for _, order := range stage.MinInclusive_stagedOrder {
		if !foundMinInclusive || order > maxMinInclusiveOrder {
			maxMinInclusiveOrder = order
			foundMinInclusive = true
		}
	}
	if foundMinInclusive {
		stage.MinInclusiveOrder = maxMinInclusiveOrder + 1
	} else {
		stage.MinInclusiveOrder = 0
	}

	var maxMinLengthOrder uint
	var foundMinLength bool
	for _, order := range stage.MinLength_stagedOrder {
		if !foundMinLength || order > maxMinLengthOrder {
			maxMinLengthOrder = order
			foundMinLength = true
		}
	}
	if foundMinLength {
		stage.MinLengthOrder = maxMinLengthOrder + 1
	} else {
		stage.MinLengthOrder = 0
	}

	var maxPatternOrder uint
	var foundPattern bool
	for _, order := range stage.Pattern_stagedOrder {
		if !foundPattern || order > maxPatternOrder {
			maxPatternOrder = order
			foundPattern = true
		}
	}
	if foundPattern {
		stage.PatternOrder = maxPatternOrder + 1
	} else {
		stage.PatternOrder = 0
	}

	var maxRestrictionOrder uint
	var foundRestriction bool
	for _, order := range stage.Restriction_stagedOrder {
		if !foundRestriction || order > maxRestrictionOrder {
			maxRestrictionOrder = order
			foundRestriction = true
		}
	}
	if foundRestriction {
		stage.RestrictionOrder = maxRestrictionOrder + 1
	} else {
		stage.RestrictionOrder = 0
	}

	var maxSchemaOrder uint
	var foundSchema bool
	for _, order := range stage.Schema_stagedOrder {
		if !foundSchema || order > maxSchemaOrder {
			maxSchemaOrder = order
			foundSchema = true
		}
	}
	if foundSchema {
		stage.SchemaOrder = maxSchemaOrder + 1
	} else {
		stage.SchemaOrder = 0
	}

	var maxSequenceOrder uint
	var foundSequence bool
	for _, order := range stage.Sequence_stagedOrder {
		if !foundSequence || order > maxSequenceOrder {
			maxSequenceOrder = order
			foundSequence = true
		}
	}
	if foundSequence {
		stage.SequenceOrder = maxSequenceOrder + 1
	} else {
		stage.SequenceOrder = 0
	}

	var maxSimpleContentOrder uint
	var foundSimpleContent bool
	for _, order := range stage.SimpleContent_stagedOrder {
		if !foundSimpleContent || order > maxSimpleContentOrder {
			maxSimpleContentOrder = order
			foundSimpleContent = true
		}
	}
	if foundSimpleContent {
		stage.SimpleContentOrder = maxSimpleContentOrder + 1
	} else {
		stage.SimpleContentOrder = 0
	}

	var maxSimpleTypeOrder uint
	var foundSimpleType bool
	for _, order := range stage.SimpleType_stagedOrder {
		if !foundSimpleType || order > maxSimpleTypeOrder {
			maxSimpleTypeOrder = order
			foundSimpleType = true
		}
	}
	if foundSimpleType {
		stage.SimpleTypeOrder = maxSimpleTypeOrder + 1
	} else {
		stage.SimpleTypeOrder = 0
	}

	var maxTotalDigitOrder uint
	var foundTotalDigit bool
	for _, order := range stage.TotalDigit_stagedOrder {
		if !foundTotalDigit || order > maxTotalDigitOrder {
			maxTotalDigitOrder = order
			foundTotalDigit = true
		}
	}
	if foundTotalDigit {
		stage.TotalDigitOrder = maxTotalDigitOrder + 1
	} else {
		stage.TotalDigitOrder = 0
	}

	var maxUnionOrder uint
	var foundUnion bool
	for _, order := range stage.Union_stagedOrder {
		if !foundUnion || order > maxUnionOrder {
			maxUnionOrder = order
			foundUnion = true
		}
	}
	if foundUnion {
		stage.UnionOrder = maxUnionOrder + 1
	} else {
		stage.UnionOrder = 0
	}

	var maxWhiteSpaceOrder uint
	var foundWhiteSpace bool
	for _, order := range stage.WhiteSpace_stagedOrder {
		if !foundWhiteSpace || order > maxWhiteSpaceOrder {
			maxWhiteSpaceOrder = order
			foundWhiteSpace = true
		}
	}
	if foundWhiteSpace {
		stage.WhiteSpaceOrder = maxWhiteSpaceOrder + 1
	} else {
		stage.WhiteSpaceOrder = 0
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
	case *All:
		tmp := GetStructInstancesByOrder(stage.Alls, stage.All_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *All implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Annotation:
		tmp := GetStructInstancesByOrder(stage.Annotations, stage.Annotation_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Annotation implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Attribute:
		tmp := GetStructInstancesByOrder(stage.Attributes, stage.Attribute_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Attribute implements.
			res = append(res, any(v).(T))
		}
		return res
	case *AttributeGroup:
		tmp := GetStructInstancesByOrder(stage.AttributeGroups, stage.AttributeGroup_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *AttributeGroup implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Choice:
		tmp := GetStructInstancesByOrder(stage.Choices, stage.Choice_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Choice implements.
			res = append(res, any(v).(T))
		}
		return res
	case *ComplexContent:
		tmp := GetStructInstancesByOrder(stage.ComplexContents, stage.ComplexContent_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *ComplexContent implements.
			res = append(res, any(v).(T))
		}
		return res
	case *ComplexType:
		tmp := GetStructInstancesByOrder(stage.ComplexTypes, stage.ComplexType_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *ComplexType implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Documentation:
		tmp := GetStructInstancesByOrder(stage.Documentations, stage.Documentation_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Documentation implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Element:
		tmp := GetStructInstancesByOrder(stage.Elements, stage.Element_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Element implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Enumeration:
		tmp := GetStructInstancesByOrder(stage.Enumerations, stage.Enumeration_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Enumeration implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Extension:
		tmp := GetStructInstancesByOrder(stage.Extensions, stage.Extension_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Extension implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Group:
		tmp := GetStructInstancesByOrder(stage.Groups, stage.Group_stagedOrder)

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
	case *Length:
		tmp := GetStructInstancesByOrder(stage.Lengths, stage.Length_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Length implements.
			res = append(res, any(v).(T))
		}
		return res
	case *MaxInclusive:
		tmp := GetStructInstancesByOrder(stage.MaxInclusives, stage.MaxInclusive_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *MaxInclusive implements.
			res = append(res, any(v).(T))
		}
		return res
	case *MaxLength:
		tmp := GetStructInstancesByOrder(stage.MaxLengths, stage.MaxLength_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *MaxLength implements.
			res = append(res, any(v).(T))
		}
		return res
	case *MinInclusive:
		tmp := GetStructInstancesByOrder(stage.MinInclusives, stage.MinInclusive_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *MinInclusive implements.
			res = append(res, any(v).(T))
		}
		return res
	case *MinLength:
		tmp := GetStructInstancesByOrder(stage.MinLengths, stage.MinLength_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *MinLength implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Pattern:
		tmp := GetStructInstancesByOrder(stage.Patterns, stage.Pattern_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Pattern implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Restriction:
		tmp := GetStructInstancesByOrder(stage.Restrictions, stage.Restriction_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Restriction implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Schema:
		tmp := GetStructInstancesByOrder(stage.Schemas, stage.Schema_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Schema implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Sequence:
		tmp := GetStructInstancesByOrder(stage.Sequences, stage.Sequence_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Sequence implements.
			res = append(res, any(v).(T))
		}
		return res
	case *SimpleContent:
		tmp := GetStructInstancesByOrder(stage.SimpleContents, stage.SimpleContent_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *SimpleContent implements.
			res = append(res, any(v).(T))
		}
		return res
	case *SimpleType:
		tmp := GetStructInstancesByOrder(stage.SimpleTypes, stage.SimpleType_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *SimpleType implements.
			res = append(res, any(v).(T))
		}
		return res
	case *TotalDigit:
		tmp := GetStructInstancesByOrder(stage.TotalDigits, stage.TotalDigit_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *TotalDigit implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Union:
		tmp := GetStructInstancesByOrder(stage.Unions, stage.Union_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Union implements.
			res = append(res, any(v).(T))
		}
		return res
	case *WhiteSpace:
		tmp := GetStructInstancesByOrder(stage.WhiteSpaces, stage.WhiteSpace_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *WhiteSpace implements.
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
	case "All":
		res = GetNamedStructInstances(stage.Alls, stage.All_stagedOrder)
	case "Annotation":
		res = GetNamedStructInstances(stage.Annotations, stage.Annotation_stagedOrder)
	case "Attribute":
		res = GetNamedStructInstances(stage.Attributes, stage.Attribute_stagedOrder)
	case "AttributeGroup":
		res = GetNamedStructInstances(stage.AttributeGroups, stage.AttributeGroup_stagedOrder)
	case "Choice":
		res = GetNamedStructInstances(stage.Choices, stage.Choice_stagedOrder)
	case "ComplexContent":
		res = GetNamedStructInstances(stage.ComplexContents, stage.ComplexContent_stagedOrder)
	case "ComplexType":
		res = GetNamedStructInstances(stage.ComplexTypes, stage.ComplexType_stagedOrder)
	case "Documentation":
		res = GetNamedStructInstances(stage.Documentations, stage.Documentation_stagedOrder)
	case "Element":
		res = GetNamedStructInstances(stage.Elements, stage.Element_stagedOrder)
	case "Enumeration":
		res = GetNamedStructInstances(stage.Enumerations, stage.Enumeration_stagedOrder)
	case "Extension":
		res = GetNamedStructInstances(stage.Extensions, stage.Extension_stagedOrder)
	case "Group":
		res = GetNamedStructInstances(stage.Groups, stage.Group_stagedOrder)
	case "Length":
		res = GetNamedStructInstances(stage.Lengths, stage.Length_stagedOrder)
	case "MaxInclusive":
		res = GetNamedStructInstances(stage.MaxInclusives, stage.MaxInclusive_stagedOrder)
	case "MaxLength":
		res = GetNamedStructInstances(stage.MaxLengths, stage.MaxLength_stagedOrder)
	case "MinInclusive":
		res = GetNamedStructInstances(stage.MinInclusives, stage.MinInclusive_stagedOrder)
	case "MinLength":
		res = GetNamedStructInstances(stage.MinLengths, stage.MinLength_stagedOrder)
	case "Pattern":
		res = GetNamedStructInstances(stage.Patterns, stage.Pattern_stagedOrder)
	case "Restriction":
		res = GetNamedStructInstances(stage.Restrictions, stage.Restriction_stagedOrder)
	case "Schema":
		res = GetNamedStructInstances(stage.Schemas, stage.Schema_stagedOrder)
	case "Sequence":
		res = GetNamedStructInstances(stage.Sequences, stage.Sequence_stagedOrder)
	case "SimpleContent":
		res = GetNamedStructInstances(stage.SimpleContents, stage.SimpleContent_stagedOrder)
	case "SimpleType":
		res = GetNamedStructInstances(stage.SimpleTypes, stage.SimpleType_stagedOrder)
	case "TotalDigit":
		res = GetNamedStructInstances(stage.TotalDigits, stage.TotalDigit_stagedOrder)
	case "Union":
		res = GetNamedStructInstances(stage.Unions, stage.Union_stagedOrder)
	case "WhiteSpace":
		res = GetNamedStructInstances(stage.WhiteSpaces, stage.WhiteSpace_stagedOrder)
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
	return "github.com/fullstack-lang/gong/app/xsd/go/models"
}

func (stage *Stage) GetMap_GongStructName_InstancesNb() map[string]int {
	return stage.Map_GongStructName_InstancesNb
}

func (stage *Stage) GetModelsEmbededDir() embed.FS {
	return xsd_go.GoModelsDir
}

func (stage *Stage) GetDigramsEmbededDir() embed.FS {
	return xsd_go.GoDiagramsDir
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
	CommitAll(all *All)
	CheckoutAll(all *All)
	CommitAnnotation(annotation *Annotation)
	CheckoutAnnotation(annotation *Annotation)
	CommitAttribute(attribute *Attribute)
	CheckoutAttribute(attribute *Attribute)
	CommitAttributeGroup(attributegroup *AttributeGroup)
	CheckoutAttributeGroup(attributegroup *AttributeGroup)
	CommitChoice(choice *Choice)
	CheckoutChoice(choice *Choice)
	CommitComplexContent(complexcontent *ComplexContent)
	CheckoutComplexContent(complexcontent *ComplexContent)
	CommitComplexType(complextype *ComplexType)
	CheckoutComplexType(complextype *ComplexType)
	CommitDocumentation(documentation *Documentation)
	CheckoutDocumentation(documentation *Documentation)
	CommitElement(element *Element)
	CheckoutElement(element *Element)
	CommitEnumeration(enumeration *Enumeration)
	CheckoutEnumeration(enumeration *Enumeration)
	CommitExtension(extension *Extension)
	CheckoutExtension(extension *Extension)
	CommitGroup(group *Group)
	CheckoutGroup(group *Group)
	CommitLength(length *Length)
	CheckoutLength(length *Length)
	CommitMaxInclusive(maxinclusive *MaxInclusive)
	CheckoutMaxInclusive(maxinclusive *MaxInclusive)
	CommitMaxLength(maxlength *MaxLength)
	CheckoutMaxLength(maxlength *MaxLength)
	CommitMinInclusive(mininclusive *MinInclusive)
	CheckoutMinInclusive(mininclusive *MinInclusive)
	CommitMinLength(minlength *MinLength)
	CheckoutMinLength(minlength *MinLength)
	CommitPattern(pattern *Pattern)
	CheckoutPattern(pattern *Pattern)
	CommitRestriction(restriction *Restriction)
	CheckoutRestriction(restriction *Restriction)
	CommitSchema(schema *Schema)
	CheckoutSchema(schema *Schema)
	CommitSequence(sequence *Sequence)
	CheckoutSequence(sequence *Sequence)
	CommitSimpleContent(simplecontent *SimpleContent)
	CheckoutSimpleContent(simplecontent *SimpleContent)
	CommitSimpleType(simpletype *SimpleType)
	CheckoutSimpleType(simpletype *SimpleType)
	CommitTotalDigit(totaldigit *TotalDigit)
	CheckoutTotalDigit(totaldigit *TotalDigit)
	CommitUnion(union *Union)
	CheckoutUnion(union *Union)
	CommitWhiteSpace(whitespace *WhiteSpace)
	CheckoutWhiteSpace(whitespace *WhiteSpace)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(name string) (stage *Stage) {
	stage = &Stage{ // insertion point for array initiatialisation
		Alls:           make(map[*All]struct{}),
		Alls_mapString: make(map[string]*All),

		Annotations:           make(map[*Annotation]struct{}),
		Annotations_mapString: make(map[string]*Annotation),

		Attributes:           make(map[*Attribute]struct{}),
		Attributes_mapString: make(map[string]*Attribute),

		AttributeGroups:           make(map[*AttributeGroup]struct{}),
		AttributeGroups_mapString: make(map[string]*AttributeGroup),

		Choices:           make(map[*Choice]struct{}),
		Choices_mapString: make(map[string]*Choice),

		ComplexContents:           make(map[*ComplexContent]struct{}),
		ComplexContents_mapString: make(map[string]*ComplexContent),

		ComplexTypes:           make(map[*ComplexType]struct{}),
		ComplexTypes_mapString: make(map[string]*ComplexType),

		Documentations:           make(map[*Documentation]struct{}),
		Documentations_mapString: make(map[string]*Documentation),

		Elements:           make(map[*Element]struct{}),
		Elements_mapString: make(map[string]*Element),

		Enumerations:           make(map[*Enumeration]struct{}),
		Enumerations_mapString: make(map[string]*Enumeration),

		Extensions:           make(map[*Extension]struct{}),
		Extensions_mapString: make(map[string]*Extension),

		Groups:           make(map[*Group]struct{}),
		Groups_mapString: make(map[string]*Group),

		Lengths:           make(map[*Length]struct{}),
		Lengths_mapString: make(map[string]*Length),

		MaxInclusives:           make(map[*MaxInclusive]struct{}),
		MaxInclusives_mapString: make(map[string]*MaxInclusive),

		MaxLengths:           make(map[*MaxLength]struct{}),
		MaxLengths_mapString: make(map[string]*MaxLength),

		MinInclusives:           make(map[*MinInclusive]struct{}),
		MinInclusives_mapString: make(map[string]*MinInclusive),

		MinLengths:           make(map[*MinLength]struct{}),
		MinLengths_mapString: make(map[string]*MinLength),

		Patterns:           make(map[*Pattern]struct{}),
		Patterns_mapString: make(map[string]*Pattern),

		Restrictions:           make(map[*Restriction]struct{}),
		Restrictions_mapString: make(map[string]*Restriction),

		Schemas:           make(map[*Schema]struct{}),
		Schemas_mapString: make(map[string]*Schema),

		Sequences:           make(map[*Sequence]struct{}),
		Sequences_mapString: make(map[string]*Sequence),

		SimpleContents:           make(map[*SimpleContent]struct{}),
		SimpleContents_mapString: make(map[string]*SimpleContent),

		SimpleTypes:           make(map[*SimpleType]struct{}),
		SimpleTypes_mapString: make(map[string]*SimpleType),

		TotalDigits:           make(map[*TotalDigit]struct{}),
		TotalDigits_mapString: make(map[string]*TotalDigit),

		Unions:           make(map[*Union]struct{}),
		Unions_mapString: make(map[string]*Union),

		WhiteSpaces:           make(map[*WhiteSpace]struct{}),
		WhiteSpaces_mapString: make(map[string]*WhiteSpace),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		name: name,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		All_stagedOrder: make(map[*All]uint),
		All_orderStaged: make(map[uint]*All),
		Alls_reference:  make(map[*All]*All),

		Annotation_stagedOrder: make(map[*Annotation]uint),
		Annotation_orderStaged: make(map[uint]*Annotation),
		Annotations_reference:  make(map[*Annotation]*Annotation),

		Attribute_stagedOrder: make(map[*Attribute]uint),
		Attribute_orderStaged: make(map[uint]*Attribute),
		Attributes_reference:  make(map[*Attribute]*Attribute),

		AttributeGroup_stagedOrder: make(map[*AttributeGroup]uint),
		AttributeGroup_orderStaged: make(map[uint]*AttributeGroup),
		AttributeGroups_reference:  make(map[*AttributeGroup]*AttributeGroup),

		Choice_stagedOrder: make(map[*Choice]uint),
		Choice_orderStaged: make(map[uint]*Choice),
		Choices_reference:  make(map[*Choice]*Choice),

		ComplexContent_stagedOrder: make(map[*ComplexContent]uint),
		ComplexContent_orderStaged: make(map[uint]*ComplexContent),
		ComplexContents_reference:  make(map[*ComplexContent]*ComplexContent),

		ComplexType_stagedOrder: make(map[*ComplexType]uint),
		ComplexType_orderStaged: make(map[uint]*ComplexType),
		ComplexTypes_reference:  make(map[*ComplexType]*ComplexType),

		Documentation_stagedOrder: make(map[*Documentation]uint),
		Documentation_orderStaged: make(map[uint]*Documentation),
		Documentations_reference:  make(map[*Documentation]*Documentation),

		Element_stagedOrder: make(map[*Element]uint),
		Element_orderStaged: make(map[uint]*Element),
		Elements_reference:  make(map[*Element]*Element),

		Enumeration_stagedOrder: make(map[*Enumeration]uint),
		Enumeration_orderStaged: make(map[uint]*Enumeration),
		Enumerations_reference:  make(map[*Enumeration]*Enumeration),

		Extension_stagedOrder: make(map[*Extension]uint),
		Extension_orderStaged: make(map[uint]*Extension),
		Extensions_reference:  make(map[*Extension]*Extension),

		Group_stagedOrder: make(map[*Group]uint),
		Group_orderStaged: make(map[uint]*Group),
		Groups_reference:  make(map[*Group]*Group),

		Length_stagedOrder: make(map[*Length]uint),
		Length_orderStaged: make(map[uint]*Length),
		Lengths_reference:  make(map[*Length]*Length),

		MaxInclusive_stagedOrder: make(map[*MaxInclusive]uint),
		MaxInclusive_orderStaged: make(map[uint]*MaxInclusive),
		MaxInclusives_reference:  make(map[*MaxInclusive]*MaxInclusive),

		MaxLength_stagedOrder: make(map[*MaxLength]uint),
		MaxLength_orderStaged: make(map[uint]*MaxLength),
		MaxLengths_reference:  make(map[*MaxLength]*MaxLength),

		MinInclusive_stagedOrder: make(map[*MinInclusive]uint),
		MinInclusive_orderStaged: make(map[uint]*MinInclusive),
		MinInclusives_reference:  make(map[*MinInclusive]*MinInclusive),

		MinLength_stagedOrder: make(map[*MinLength]uint),
		MinLength_orderStaged: make(map[uint]*MinLength),
		MinLengths_reference:  make(map[*MinLength]*MinLength),

		Pattern_stagedOrder: make(map[*Pattern]uint),
		Pattern_orderStaged: make(map[uint]*Pattern),
		Patterns_reference:  make(map[*Pattern]*Pattern),

		Restriction_stagedOrder: make(map[*Restriction]uint),
		Restriction_orderStaged: make(map[uint]*Restriction),
		Restrictions_reference:  make(map[*Restriction]*Restriction),

		Schema_stagedOrder: make(map[*Schema]uint),
		Schema_orderStaged: make(map[uint]*Schema),
		Schemas_reference:  make(map[*Schema]*Schema),

		Sequence_stagedOrder: make(map[*Sequence]uint),
		Sequence_orderStaged: make(map[uint]*Sequence),
		Sequences_reference:  make(map[*Sequence]*Sequence),

		SimpleContent_stagedOrder: make(map[*SimpleContent]uint),
		SimpleContent_orderStaged: make(map[uint]*SimpleContent),
		SimpleContents_reference:  make(map[*SimpleContent]*SimpleContent),

		SimpleType_stagedOrder: make(map[*SimpleType]uint),
		SimpleType_orderStaged: make(map[uint]*SimpleType),
		SimpleTypes_reference:  make(map[*SimpleType]*SimpleType),

		TotalDigit_stagedOrder: make(map[*TotalDigit]uint),
		TotalDigit_orderStaged: make(map[uint]*TotalDigit),
		TotalDigits_reference:  make(map[*TotalDigit]*TotalDigit),

		Union_stagedOrder: make(map[*Union]uint),
		Union_orderStaged: make(map[uint]*Union),
		Unions_reference:  make(map[*Union]*Union),

		WhiteSpace_stagedOrder: make(map[*WhiteSpace]uint),
		WhiteSpace_orderStaged: make(map[uint]*WhiteSpace),
		WhiteSpaces_reference:  make(map[*WhiteSpace]*WhiteSpace),

		// end of insertion point
		GongUnmarshallers: map[string]ModelUnmarshaller{ // insertion point for unmarshallers
			"All": &AllUnmarshaller{},

			"Annotation": &AnnotationUnmarshaller{},

			"Attribute": &AttributeUnmarshaller{},

			"AttributeGroup": &AttributeGroupUnmarshaller{},

			"Choice": &ChoiceUnmarshaller{},

			"ComplexContent": &ComplexContentUnmarshaller{},

			"ComplexType": &ComplexTypeUnmarshaller{},

			"Documentation": &DocumentationUnmarshaller{},

			"Element": &ElementUnmarshaller{},

			"Enumeration": &EnumerationUnmarshaller{},

			"Extension": &ExtensionUnmarshaller{},

			"Group": &GroupUnmarshaller{},

			"Length": &LengthUnmarshaller{},

			"MaxInclusive": &MaxInclusiveUnmarshaller{},

			"MaxLength": &MaxLengthUnmarshaller{},

			"MinInclusive": &MinInclusiveUnmarshaller{},

			"MinLength": &MinLengthUnmarshaller{},

			"Pattern": &PatternUnmarshaller{},

			"Restriction": &RestrictionUnmarshaller{},

			"Schema": &SchemaUnmarshaller{},

			"Sequence": &SequenceUnmarshaller{},

			"SimpleContent": &SimpleContentUnmarshaller{},

			"SimpleType": &SimpleTypeUnmarshaller{},

			"TotalDigit": &TotalDigitUnmarshaller{},

			"Union": &UnionUnmarshaller{},

			"WhiteSpace": &WhiteSpaceUnmarshaller{},

			// end of insertion point
		},

		NamedStructs: []*NamedStruct{ // insertion point for order map initialisations
			{name: "All"},
			{name: "Annotation"},
			{name: "Attribute"},
			{name: "AttributeGroup"},
			{name: "Choice"},
			{name: "ComplexContent"},
			{name: "ComplexType"},
			{name: "Documentation"},
			{name: "Element"},
			{name: "Enumeration"},
			{name: "Extension"},
			{name: "Group"},
			{name: "Length"},
			{name: "MaxInclusive"},
			{name: "MaxLength"},
			{name: "MinInclusive"},
			{name: "MinLength"},
			{name: "Pattern"},
			{name: "Restriction"},
			{name: "Schema"},
			{name: "Sequence"},
			{name: "SimpleContent"},
			{name: "SimpleType"},
			{name: "TotalDigit"},
			{name: "Union"},
			{name: "WhiteSpace"},
		}, // end of insertion point

		navigationMode: GongNavigationModeNormal,
	}

	return
}

func GetOrder[Type Gongstruct](stage *Stage, instance *Type) uint {
	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *All:
		return stage.All_stagedOrder[instance]
	case *Annotation:
		return stage.Annotation_stagedOrder[instance]
	case *Attribute:
		return stage.Attribute_stagedOrder[instance]
	case *AttributeGroup:
		return stage.AttributeGroup_stagedOrder[instance]
	case *Choice:
		return stage.Choice_stagedOrder[instance]
	case *ComplexContent:
		return stage.ComplexContent_stagedOrder[instance]
	case *ComplexType:
		return stage.ComplexType_stagedOrder[instance]
	case *Documentation:
		return stage.Documentation_stagedOrder[instance]
	case *Element:
		return stage.Element_stagedOrder[instance]
	case *Enumeration:
		return stage.Enumeration_stagedOrder[instance]
	case *Extension:
		return stage.Extension_stagedOrder[instance]
	case *Group:
		return stage.Group_stagedOrder[instance]
	case *Length:
		return stage.Length_stagedOrder[instance]
	case *MaxInclusive:
		return stage.MaxInclusive_stagedOrder[instance]
	case *MaxLength:
		return stage.MaxLength_stagedOrder[instance]
	case *MinInclusive:
		return stage.MinInclusive_stagedOrder[instance]
	case *MinLength:
		return stage.MinLength_stagedOrder[instance]
	case *Pattern:
		return stage.Pattern_stagedOrder[instance]
	case *Restriction:
		return stage.Restriction_stagedOrder[instance]
	case *Schema:
		return stage.Schema_stagedOrder[instance]
	case *Sequence:
		return stage.Sequence_stagedOrder[instance]
	case *SimpleContent:
		return stage.SimpleContent_stagedOrder[instance]
	case *SimpleType:
		return stage.SimpleType_stagedOrder[instance]
	case *TotalDigit:
		return stage.TotalDigit_stagedOrder[instance]
	case *Union:
		return stage.Union_stagedOrder[instance]
	case *WhiteSpace:
		return stage.WhiteSpace_stagedOrder[instance]
	default:
		return 0 // should not happen
	}
}

func GongGetInstanceFromOrder[Type PointerToGongstruct](stage *Stage, order uint) (res Type) {
	var t Type
	switch any(t).(type) {
	// insertion point for order map initialisations
	case *All:
		return any(stage.All_orderStaged[order]).(Type)
	case *Annotation:
		return any(stage.Annotation_orderStaged[order]).(Type)
	case *Attribute:
		return any(stage.Attribute_orderStaged[order]).(Type)
	case *AttributeGroup:
		return any(stage.AttributeGroup_orderStaged[order]).(Type)
	case *Choice:
		return any(stage.Choice_orderStaged[order]).(Type)
	case *ComplexContent:
		return any(stage.ComplexContent_orderStaged[order]).(Type)
	case *ComplexType:
		return any(stage.ComplexType_orderStaged[order]).(Type)
	case *Documentation:
		return any(stage.Documentation_orderStaged[order]).(Type)
	case *Element:
		return any(stage.Element_orderStaged[order]).(Type)
	case *Enumeration:
		return any(stage.Enumeration_orderStaged[order]).(Type)
	case *Extension:
		return any(stage.Extension_orderStaged[order]).(Type)
	case *Group:
		return any(stage.Group_orderStaged[order]).(Type)
	case *Length:
		return any(stage.Length_orderStaged[order]).(Type)
	case *MaxInclusive:
		return any(stage.MaxInclusive_orderStaged[order]).(Type)
	case *MaxLength:
		return any(stage.MaxLength_orderStaged[order]).(Type)
	case *MinInclusive:
		return any(stage.MinInclusive_orderStaged[order]).(Type)
	case *MinLength:
		return any(stage.MinLength_orderStaged[order]).(Type)
	case *Pattern:
		return any(stage.Pattern_orderStaged[order]).(Type)
	case *Restriction:
		return any(stage.Restriction_orderStaged[order]).(Type)
	case *Schema:
		return any(stage.Schema_orderStaged[order]).(Type)
	case *Sequence:
		return any(stage.Sequence_orderStaged[order]).(Type)
	case *SimpleContent:
		return any(stage.SimpleContent_orderStaged[order]).(Type)
	case *SimpleType:
		return any(stage.SimpleType_orderStaged[order]).(Type)
	case *TotalDigit:
		return any(stage.TotalDigit_orderStaged[order]).(Type)
	case *Union:
		return any(stage.Union_orderStaged[order]).(Type)
	case *WhiteSpace:
		return any(stage.WhiteSpace_orderStaged[order]).(Type)
	default:
		return // should not happen
	}
}

func GetOrderPointerGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) uint {
	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *All:
		return stage.All_stagedOrder[instance]
	case *Annotation:
		return stage.Annotation_stagedOrder[instance]
	case *Attribute:
		return stage.Attribute_stagedOrder[instance]
	case *AttributeGroup:
		return stage.AttributeGroup_stagedOrder[instance]
	case *Choice:
		return stage.Choice_stagedOrder[instance]
	case *ComplexContent:
		return stage.ComplexContent_stagedOrder[instance]
	case *ComplexType:
		return stage.ComplexType_stagedOrder[instance]
	case *Documentation:
		return stage.Documentation_stagedOrder[instance]
	case *Element:
		return stage.Element_stagedOrder[instance]
	case *Enumeration:
		return stage.Enumeration_stagedOrder[instance]
	case *Extension:
		return stage.Extension_stagedOrder[instance]
	case *Group:
		return stage.Group_stagedOrder[instance]
	case *Length:
		return stage.Length_stagedOrder[instance]
	case *MaxInclusive:
		return stage.MaxInclusive_stagedOrder[instance]
	case *MaxLength:
		return stage.MaxLength_stagedOrder[instance]
	case *MinInclusive:
		return stage.MinInclusive_stagedOrder[instance]
	case *MinLength:
		return stage.MinLength_stagedOrder[instance]
	case *Pattern:
		return stage.Pattern_stagedOrder[instance]
	case *Restriction:
		return stage.Restriction_stagedOrder[instance]
	case *Schema:
		return stage.Schema_stagedOrder[instance]
	case *Sequence:
		return stage.Sequence_stagedOrder[instance]
	case *SimpleContent:
		return stage.SimpleContent_stagedOrder[instance]
	case *SimpleType:
		return stage.SimpleType_stagedOrder[instance]
	case *TotalDigit:
		return stage.TotalDigit_stagedOrder[instance]
	case *Union:
		return stage.Union_stagedOrder[instance]
	case *WhiteSpace:
		return stage.WhiteSpace_stagedOrder[instance]
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
	stage.Map_GongStructName_InstancesNb["All"] = len(stage.Alls)
	stage.Map_GongStructName_InstancesNb["Annotation"] = len(stage.Annotations)
	stage.Map_GongStructName_InstancesNb["Attribute"] = len(stage.Attributes)
	stage.Map_GongStructName_InstancesNb["AttributeGroup"] = len(stage.AttributeGroups)
	stage.Map_GongStructName_InstancesNb["Choice"] = len(stage.Choices)
	stage.Map_GongStructName_InstancesNb["ComplexContent"] = len(stage.ComplexContents)
	stage.Map_GongStructName_InstancesNb["ComplexType"] = len(stage.ComplexTypes)
	stage.Map_GongStructName_InstancesNb["Documentation"] = len(stage.Documentations)
	stage.Map_GongStructName_InstancesNb["Element"] = len(stage.Elements)
	stage.Map_GongStructName_InstancesNb["Enumeration"] = len(stage.Enumerations)
	stage.Map_GongStructName_InstancesNb["Extension"] = len(stage.Extensions)
	stage.Map_GongStructName_InstancesNb["Group"] = len(stage.Groups)
	stage.Map_GongStructName_InstancesNb["Length"] = len(stage.Lengths)
	stage.Map_GongStructName_InstancesNb["MaxInclusive"] = len(stage.MaxInclusives)
	stage.Map_GongStructName_InstancesNb["MaxLength"] = len(stage.MaxLengths)
	stage.Map_GongStructName_InstancesNb["MinInclusive"] = len(stage.MinInclusives)
	stage.Map_GongStructName_InstancesNb["MinLength"] = len(stage.MinLengths)
	stage.Map_GongStructName_InstancesNb["Pattern"] = len(stage.Patterns)
	stage.Map_GongStructName_InstancesNb["Restriction"] = len(stage.Restrictions)
	stage.Map_GongStructName_InstancesNb["Schema"] = len(stage.Schemas)
	stage.Map_GongStructName_InstancesNb["Sequence"] = len(stage.Sequences)
	stage.Map_GongStructName_InstancesNb["SimpleContent"] = len(stage.SimpleContents)
	stage.Map_GongStructName_InstancesNb["SimpleType"] = len(stage.SimpleTypes)
	stage.Map_GongStructName_InstancesNb["TotalDigit"] = len(stage.TotalDigits)
	stage.Map_GongStructName_InstancesNb["Union"] = len(stage.Unions)
	stage.Map_GongStructName_InstancesNb["WhiteSpace"] = len(stage.WhiteSpaces)
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
// Stage puts all to the model stage
func (all *All) Stage(stage *Stage) *All {
	if _, ok := stage.Alls[all]; !ok {
		stage.Alls[all] = struct{}{}
		stage.All_stagedOrder[all] = stage.AllOrder
		stage.All_orderStaged[stage.AllOrder] = all
		stage.AllOrder++
	}
	stage.Alls_mapString[all.Name] = all

	return all
}

// StagePreserveOrder puts all to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AllOrder
// - update stage.AllOrder accordingly
func (all *All) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Alls[all]; !ok {
		stage.Alls[all] = struct{}{}

		if order > stage.AllOrder {
			stage.AllOrder = order
		}
		stage.All_stagedOrder[all] = order
		stage.All_orderStaged[order] = all
		stage.AllOrder++
	}
	stage.Alls_mapString[all.Name] = all
}

// Unstage removes all off the model stage
func (all *All) Unstage(stage *Stage) *All {
	delete(stage.Alls, all)
	// issue1150
	// delete(stage.All_stagedOrder, all)
	delete(stage.Alls_mapString, all.Name)

	return all
}

// UnstageVoid removes all off the model stage
func (all *All) UnstageVoid(stage *Stage) {
	delete(stage.Alls, all)
	// issue1150
	// delete(stage.All_stagedOrder, all)
	delete(stage.Alls_mapString, all.Name)
}

// commit all to the back repo (if it is already staged)
func (all *All) Commit(stage *Stage) *All {
	if _, ok := stage.Alls[all]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitAll(all)
		}
	}
	return all
}

func (all *All) CommitVoid(stage *Stage) {
	all.Commit(stage)
}

func (all *All) StageVoid(stage *Stage) {
	all.Stage(stage)
}

// Checkout all to the back repo (if it is already staged)
func (all *All) Checkout(stage *Stage) *All {
	if _, ok := stage.Alls[all]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutAll(all)
		}
	}
	return all
}

// for satisfaction of GongStruct interface
func (all *All) GetName() (res string) {
	return all.Name
}

// for satisfaction of GongStruct interface
func (all *All) SetName(name string) {
	all.Name = name
}

// Stage puts annotation to the model stage
func (annotation *Annotation) Stage(stage *Stage) *Annotation {
	if _, ok := stage.Annotations[annotation]; !ok {
		stage.Annotations[annotation] = struct{}{}
		stage.Annotation_stagedOrder[annotation] = stage.AnnotationOrder
		stage.Annotation_orderStaged[stage.AnnotationOrder] = annotation
		stage.AnnotationOrder++
	}
	stage.Annotations_mapString[annotation.Name] = annotation

	return annotation
}

// StagePreserveOrder puts annotation to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AnnotationOrder
// - update stage.AnnotationOrder accordingly
func (annotation *Annotation) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Annotations[annotation]; !ok {
		stage.Annotations[annotation] = struct{}{}

		if order > stage.AnnotationOrder {
			stage.AnnotationOrder = order
		}
		stage.Annotation_stagedOrder[annotation] = order
		stage.Annotation_orderStaged[order] = annotation
		stage.AnnotationOrder++
	}
	stage.Annotations_mapString[annotation.Name] = annotation
}

// Unstage removes annotation off the model stage
func (annotation *Annotation) Unstage(stage *Stage) *Annotation {
	delete(stage.Annotations, annotation)
	// issue1150
	// delete(stage.Annotation_stagedOrder, annotation)
	delete(stage.Annotations_mapString, annotation.Name)

	return annotation
}

// UnstageVoid removes annotation off the model stage
func (annotation *Annotation) UnstageVoid(stage *Stage) {
	delete(stage.Annotations, annotation)
	// issue1150
	// delete(stage.Annotation_stagedOrder, annotation)
	delete(stage.Annotations_mapString, annotation.Name)
}

// commit annotation to the back repo (if it is already staged)
func (annotation *Annotation) Commit(stage *Stage) *Annotation {
	if _, ok := stage.Annotations[annotation]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitAnnotation(annotation)
		}
	}
	return annotation
}

func (annotation *Annotation) CommitVoid(stage *Stage) {
	annotation.Commit(stage)
}

func (annotation *Annotation) StageVoid(stage *Stage) {
	annotation.Stage(stage)
}

// Checkout annotation to the back repo (if it is already staged)
func (annotation *Annotation) Checkout(stage *Stage) *Annotation {
	if _, ok := stage.Annotations[annotation]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutAnnotation(annotation)
		}
	}
	return annotation
}

// for satisfaction of GongStruct interface
func (annotation *Annotation) GetName() (res string) {
	return annotation.Name
}

// for satisfaction of GongStruct interface
func (annotation *Annotation) SetName(name string) {
	annotation.Name = name
}

// Stage puts attribute to the model stage
func (attribute *Attribute) Stage(stage *Stage) *Attribute {
	if _, ok := stage.Attributes[attribute]; !ok {
		stage.Attributes[attribute] = struct{}{}
		stage.Attribute_stagedOrder[attribute] = stage.AttributeOrder
		stage.Attribute_orderStaged[stage.AttributeOrder] = attribute
		stage.AttributeOrder++
	}
	stage.Attributes_mapString[attribute.Name] = attribute

	return attribute
}

// StagePreserveOrder puts attribute to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AttributeOrder
// - update stage.AttributeOrder accordingly
func (attribute *Attribute) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Attributes[attribute]; !ok {
		stage.Attributes[attribute] = struct{}{}

		if order > stage.AttributeOrder {
			stage.AttributeOrder = order
		}
		stage.Attribute_stagedOrder[attribute] = order
		stage.Attribute_orderStaged[order] = attribute
		stage.AttributeOrder++
	}
	stage.Attributes_mapString[attribute.Name] = attribute
}

// Unstage removes attribute off the model stage
func (attribute *Attribute) Unstage(stage *Stage) *Attribute {
	delete(stage.Attributes, attribute)
	// issue1150
	// delete(stage.Attribute_stagedOrder, attribute)
	delete(stage.Attributes_mapString, attribute.Name)

	return attribute
}

// UnstageVoid removes attribute off the model stage
func (attribute *Attribute) UnstageVoid(stage *Stage) {
	delete(stage.Attributes, attribute)
	// issue1150
	// delete(stage.Attribute_stagedOrder, attribute)
	delete(stage.Attributes_mapString, attribute.Name)
}

// commit attribute to the back repo (if it is already staged)
func (attribute *Attribute) Commit(stage *Stage) *Attribute {
	if _, ok := stage.Attributes[attribute]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitAttribute(attribute)
		}
	}
	return attribute
}

func (attribute *Attribute) CommitVoid(stage *Stage) {
	attribute.Commit(stage)
}

func (attribute *Attribute) StageVoid(stage *Stage) {
	attribute.Stage(stage)
}

// Checkout attribute to the back repo (if it is already staged)
func (attribute *Attribute) Checkout(stage *Stage) *Attribute {
	if _, ok := stage.Attributes[attribute]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutAttribute(attribute)
		}
	}
	return attribute
}

// for satisfaction of GongStruct interface
func (attribute *Attribute) GetName() (res string) {
	return attribute.Name
}

// for satisfaction of GongStruct interface
func (attribute *Attribute) SetName(name string) {
	attribute.Name = name
}

// Stage puts attributegroup to the model stage
func (attributegroup *AttributeGroup) Stage(stage *Stage) *AttributeGroup {
	if _, ok := stage.AttributeGroups[attributegroup]; !ok {
		stage.AttributeGroups[attributegroup] = struct{}{}
		stage.AttributeGroup_stagedOrder[attributegroup] = stage.AttributeGroupOrder
		stage.AttributeGroup_orderStaged[stage.AttributeGroupOrder] = attributegroup
		stage.AttributeGroupOrder++
	}
	stage.AttributeGroups_mapString[attributegroup.Name] = attributegroup

	return attributegroup
}

// StagePreserveOrder puts attributegroup to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AttributeGroupOrder
// - update stage.AttributeGroupOrder accordingly
func (attributegroup *AttributeGroup) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.AttributeGroups[attributegroup]; !ok {
		stage.AttributeGroups[attributegroup] = struct{}{}

		if order > stage.AttributeGroupOrder {
			stage.AttributeGroupOrder = order
		}
		stage.AttributeGroup_stagedOrder[attributegroup] = order
		stage.AttributeGroup_orderStaged[order] = attributegroup
		stage.AttributeGroupOrder++
	}
	stage.AttributeGroups_mapString[attributegroup.Name] = attributegroup
}

// Unstage removes attributegroup off the model stage
func (attributegroup *AttributeGroup) Unstage(stage *Stage) *AttributeGroup {
	delete(stage.AttributeGroups, attributegroup)
	// issue1150
	// delete(stage.AttributeGroup_stagedOrder, attributegroup)
	delete(stage.AttributeGroups_mapString, attributegroup.Name)

	return attributegroup
}

// UnstageVoid removes attributegroup off the model stage
func (attributegroup *AttributeGroup) UnstageVoid(stage *Stage) {
	delete(stage.AttributeGroups, attributegroup)
	// issue1150
	// delete(stage.AttributeGroup_stagedOrder, attributegroup)
	delete(stage.AttributeGroups_mapString, attributegroup.Name)
}

// commit attributegroup to the back repo (if it is already staged)
func (attributegroup *AttributeGroup) Commit(stage *Stage) *AttributeGroup {
	if _, ok := stage.AttributeGroups[attributegroup]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitAttributeGroup(attributegroup)
		}
	}
	return attributegroup
}

func (attributegroup *AttributeGroup) CommitVoid(stage *Stage) {
	attributegroup.Commit(stage)
}

func (attributegroup *AttributeGroup) StageVoid(stage *Stage) {
	attributegroup.Stage(stage)
}

// Checkout attributegroup to the back repo (if it is already staged)
func (attributegroup *AttributeGroup) Checkout(stage *Stage) *AttributeGroup {
	if _, ok := stage.AttributeGroups[attributegroup]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutAttributeGroup(attributegroup)
		}
	}
	return attributegroup
}

// for satisfaction of GongStruct interface
func (attributegroup *AttributeGroup) GetName() (res string) {
	return attributegroup.Name
}

// for satisfaction of GongStruct interface
func (attributegroup *AttributeGroup) SetName(name string) {
	attributegroup.Name = name
}

// Stage puts choice to the model stage
func (choice *Choice) Stage(stage *Stage) *Choice {
	if _, ok := stage.Choices[choice]; !ok {
		stage.Choices[choice] = struct{}{}
		stage.Choice_stagedOrder[choice] = stage.ChoiceOrder
		stage.Choice_orderStaged[stage.ChoiceOrder] = choice
		stage.ChoiceOrder++
	}
	stage.Choices_mapString[choice.Name] = choice

	return choice
}

// StagePreserveOrder puts choice to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ChoiceOrder
// - update stage.ChoiceOrder accordingly
func (choice *Choice) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Choices[choice]; !ok {
		stage.Choices[choice] = struct{}{}

		if order > stage.ChoiceOrder {
			stage.ChoiceOrder = order
		}
		stage.Choice_stagedOrder[choice] = order
		stage.Choice_orderStaged[order] = choice
		stage.ChoiceOrder++
	}
	stage.Choices_mapString[choice.Name] = choice
}

// Unstage removes choice off the model stage
func (choice *Choice) Unstage(stage *Stage) *Choice {
	delete(stage.Choices, choice)
	// issue1150
	// delete(stage.Choice_stagedOrder, choice)
	delete(stage.Choices_mapString, choice.Name)

	return choice
}

// UnstageVoid removes choice off the model stage
func (choice *Choice) UnstageVoid(stage *Stage) {
	delete(stage.Choices, choice)
	// issue1150
	// delete(stage.Choice_stagedOrder, choice)
	delete(stage.Choices_mapString, choice.Name)
}

// commit choice to the back repo (if it is already staged)
func (choice *Choice) Commit(stage *Stage) *Choice {
	if _, ok := stage.Choices[choice]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitChoice(choice)
		}
	}
	return choice
}

func (choice *Choice) CommitVoid(stage *Stage) {
	choice.Commit(stage)
}

func (choice *Choice) StageVoid(stage *Stage) {
	choice.Stage(stage)
}

// Checkout choice to the back repo (if it is already staged)
func (choice *Choice) Checkout(stage *Stage) *Choice {
	if _, ok := stage.Choices[choice]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutChoice(choice)
		}
	}
	return choice
}

// for satisfaction of GongStruct interface
func (choice *Choice) GetName() (res string) {
	return choice.Name
}

// for satisfaction of GongStruct interface
func (choice *Choice) SetName(name string) {
	choice.Name = name
}

// Stage puts complexcontent to the model stage
func (complexcontent *ComplexContent) Stage(stage *Stage) *ComplexContent {
	if _, ok := stage.ComplexContents[complexcontent]; !ok {
		stage.ComplexContents[complexcontent] = struct{}{}
		stage.ComplexContent_stagedOrder[complexcontent] = stage.ComplexContentOrder
		stage.ComplexContent_orderStaged[stage.ComplexContentOrder] = complexcontent
		stage.ComplexContentOrder++
	}
	stage.ComplexContents_mapString[complexcontent.Name] = complexcontent

	return complexcontent
}

// StagePreserveOrder puts complexcontent to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ComplexContentOrder
// - update stage.ComplexContentOrder accordingly
func (complexcontent *ComplexContent) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.ComplexContents[complexcontent]; !ok {
		stage.ComplexContents[complexcontent] = struct{}{}

		if order > stage.ComplexContentOrder {
			stage.ComplexContentOrder = order
		}
		stage.ComplexContent_stagedOrder[complexcontent] = order
		stage.ComplexContent_orderStaged[order] = complexcontent
		stage.ComplexContentOrder++
	}
	stage.ComplexContents_mapString[complexcontent.Name] = complexcontent
}

// Unstage removes complexcontent off the model stage
func (complexcontent *ComplexContent) Unstage(stage *Stage) *ComplexContent {
	delete(stage.ComplexContents, complexcontent)
	// issue1150
	// delete(stage.ComplexContent_stagedOrder, complexcontent)
	delete(stage.ComplexContents_mapString, complexcontent.Name)

	return complexcontent
}

// UnstageVoid removes complexcontent off the model stage
func (complexcontent *ComplexContent) UnstageVoid(stage *Stage) {
	delete(stage.ComplexContents, complexcontent)
	// issue1150
	// delete(stage.ComplexContent_stagedOrder, complexcontent)
	delete(stage.ComplexContents_mapString, complexcontent.Name)
}

// commit complexcontent to the back repo (if it is already staged)
func (complexcontent *ComplexContent) Commit(stage *Stage) *ComplexContent {
	if _, ok := stage.ComplexContents[complexcontent]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitComplexContent(complexcontent)
		}
	}
	return complexcontent
}

func (complexcontent *ComplexContent) CommitVoid(stage *Stage) {
	complexcontent.Commit(stage)
}

func (complexcontent *ComplexContent) StageVoid(stage *Stage) {
	complexcontent.Stage(stage)
}

// Checkout complexcontent to the back repo (if it is already staged)
func (complexcontent *ComplexContent) Checkout(stage *Stage) *ComplexContent {
	if _, ok := stage.ComplexContents[complexcontent]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutComplexContent(complexcontent)
		}
	}
	return complexcontent
}

// for satisfaction of GongStruct interface
func (complexcontent *ComplexContent) GetName() (res string) {
	return complexcontent.Name
}

// for satisfaction of GongStruct interface
func (complexcontent *ComplexContent) SetName(name string) {
	complexcontent.Name = name
}

// Stage puts complextype to the model stage
func (complextype *ComplexType) Stage(stage *Stage) *ComplexType {
	if _, ok := stage.ComplexTypes[complextype]; !ok {
		stage.ComplexTypes[complextype] = struct{}{}
		stage.ComplexType_stagedOrder[complextype] = stage.ComplexTypeOrder
		stage.ComplexType_orderStaged[stage.ComplexTypeOrder] = complextype
		stage.ComplexTypeOrder++
	}
	stage.ComplexTypes_mapString[complextype.Name] = complextype

	return complextype
}

// StagePreserveOrder puts complextype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ComplexTypeOrder
// - update stage.ComplexTypeOrder accordingly
func (complextype *ComplexType) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.ComplexTypes[complextype]; !ok {
		stage.ComplexTypes[complextype] = struct{}{}

		if order > stage.ComplexTypeOrder {
			stage.ComplexTypeOrder = order
		}
		stage.ComplexType_stagedOrder[complextype] = order
		stage.ComplexType_orderStaged[order] = complextype
		stage.ComplexTypeOrder++
	}
	stage.ComplexTypes_mapString[complextype.Name] = complextype
}

// Unstage removes complextype off the model stage
func (complextype *ComplexType) Unstage(stage *Stage) *ComplexType {
	delete(stage.ComplexTypes, complextype)
	// issue1150
	// delete(stage.ComplexType_stagedOrder, complextype)
	delete(stage.ComplexTypes_mapString, complextype.Name)

	return complextype
}

// UnstageVoid removes complextype off the model stage
func (complextype *ComplexType) UnstageVoid(stage *Stage) {
	delete(stage.ComplexTypes, complextype)
	// issue1150
	// delete(stage.ComplexType_stagedOrder, complextype)
	delete(stage.ComplexTypes_mapString, complextype.Name)
}

// commit complextype to the back repo (if it is already staged)
func (complextype *ComplexType) Commit(stage *Stage) *ComplexType {
	if _, ok := stage.ComplexTypes[complextype]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitComplexType(complextype)
		}
	}
	return complextype
}

func (complextype *ComplexType) CommitVoid(stage *Stage) {
	complextype.Commit(stage)
}

func (complextype *ComplexType) StageVoid(stage *Stage) {
	complextype.Stage(stage)
}

// Checkout complextype to the back repo (if it is already staged)
func (complextype *ComplexType) Checkout(stage *Stage) *ComplexType {
	if _, ok := stage.ComplexTypes[complextype]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutComplexType(complextype)
		}
	}
	return complextype
}

// for satisfaction of GongStruct interface
func (complextype *ComplexType) GetName() (res string) {
	return complextype.Name
}

// for satisfaction of GongStruct interface
func (complextype *ComplexType) SetName(name string) {
	complextype.Name = name
}

// Stage puts documentation to the model stage
func (documentation *Documentation) Stage(stage *Stage) *Documentation {
	if _, ok := stage.Documentations[documentation]; !ok {
		stage.Documentations[documentation] = struct{}{}
		stage.Documentation_stagedOrder[documentation] = stage.DocumentationOrder
		stage.Documentation_orderStaged[stage.DocumentationOrder] = documentation
		stage.DocumentationOrder++
	}
	stage.Documentations_mapString[documentation.Name] = documentation

	return documentation
}

// StagePreserveOrder puts documentation to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DocumentationOrder
// - update stage.DocumentationOrder accordingly
func (documentation *Documentation) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Documentations[documentation]; !ok {
		stage.Documentations[documentation] = struct{}{}

		if order > stage.DocumentationOrder {
			stage.DocumentationOrder = order
		}
		stage.Documentation_stagedOrder[documentation] = order
		stage.Documentation_orderStaged[order] = documentation
		stage.DocumentationOrder++
	}
	stage.Documentations_mapString[documentation.Name] = documentation
}

// Unstage removes documentation off the model stage
func (documentation *Documentation) Unstage(stage *Stage) *Documentation {
	delete(stage.Documentations, documentation)
	// issue1150
	// delete(stage.Documentation_stagedOrder, documentation)
	delete(stage.Documentations_mapString, documentation.Name)

	return documentation
}

// UnstageVoid removes documentation off the model stage
func (documentation *Documentation) UnstageVoid(stage *Stage) {
	delete(stage.Documentations, documentation)
	// issue1150
	// delete(stage.Documentation_stagedOrder, documentation)
	delete(stage.Documentations_mapString, documentation.Name)
}

// commit documentation to the back repo (if it is already staged)
func (documentation *Documentation) Commit(stage *Stage) *Documentation {
	if _, ok := stage.Documentations[documentation]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDocumentation(documentation)
		}
	}
	return documentation
}

func (documentation *Documentation) CommitVoid(stage *Stage) {
	documentation.Commit(stage)
}

func (documentation *Documentation) StageVoid(stage *Stage) {
	documentation.Stage(stage)
}

// Checkout documentation to the back repo (if it is already staged)
func (documentation *Documentation) Checkout(stage *Stage) *Documentation {
	if _, ok := stage.Documentations[documentation]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutDocumentation(documentation)
		}
	}
	return documentation
}

// for satisfaction of GongStruct interface
func (documentation *Documentation) GetName() (res string) {
	return documentation.Name
}

// for satisfaction of GongStruct interface
func (documentation *Documentation) SetName(name string) {
	documentation.Name = name
}

// Stage puts element to the model stage
func (element *Element) Stage(stage *Stage) *Element {
	if _, ok := stage.Elements[element]; !ok {
		stage.Elements[element] = struct{}{}
		stage.Element_stagedOrder[element] = stage.ElementOrder
		stage.Element_orderStaged[stage.ElementOrder] = element
		stage.ElementOrder++
	}
	stage.Elements_mapString[element.Name] = element

	return element
}

// StagePreserveOrder puts element to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ElementOrder
// - update stage.ElementOrder accordingly
func (element *Element) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Elements[element]; !ok {
		stage.Elements[element] = struct{}{}

		if order > stage.ElementOrder {
			stage.ElementOrder = order
		}
		stage.Element_stagedOrder[element] = order
		stage.Element_orderStaged[order] = element
		stage.ElementOrder++
	}
	stage.Elements_mapString[element.Name] = element
}

// Unstage removes element off the model stage
func (element *Element) Unstage(stage *Stage) *Element {
	delete(stage.Elements, element)
	// issue1150
	// delete(stage.Element_stagedOrder, element)
	delete(stage.Elements_mapString, element.Name)

	return element
}

// UnstageVoid removes element off the model stage
func (element *Element) UnstageVoid(stage *Stage) {
	delete(stage.Elements, element)
	// issue1150
	// delete(stage.Element_stagedOrder, element)
	delete(stage.Elements_mapString, element.Name)
}

// commit element to the back repo (if it is already staged)
func (element *Element) Commit(stage *Stage) *Element {
	if _, ok := stage.Elements[element]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitElement(element)
		}
	}
	return element
}

func (element *Element) CommitVoid(stage *Stage) {
	element.Commit(stage)
}

func (element *Element) StageVoid(stage *Stage) {
	element.Stage(stage)
}

// Checkout element to the back repo (if it is already staged)
func (element *Element) Checkout(stage *Stage) *Element {
	if _, ok := stage.Elements[element]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutElement(element)
		}
	}
	return element
}

// for satisfaction of GongStruct interface
func (element *Element) GetName() (res string) {
	return element.Name
}

// for satisfaction of GongStruct interface
func (element *Element) SetName(name string) {
	element.Name = name
}

// Stage puts enumeration to the model stage
func (enumeration *Enumeration) Stage(stage *Stage) *Enumeration {
	if _, ok := stage.Enumerations[enumeration]; !ok {
		stage.Enumerations[enumeration] = struct{}{}
		stage.Enumeration_stagedOrder[enumeration] = stage.EnumerationOrder
		stage.Enumeration_orderStaged[stage.EnumerationOrder] = enumeration
		stage.EnumerationOrder++
	}
	stage.Enumerations_mapString[enumeration.Name] = enumeration

	return enumeration
}

// StagePreserveOrder puts enumeration to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.EnumerationOrder
// - update stage.EnumerationOrder accordingly
func (enumeration *Enumeration) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Enumerations[enumeration]; !ok {
		stage.Enumerations[enumeration] = struct{}{}

		if order > stage.EnumerationOrder {
			stage.EnumerationOrder = order
		}
		stage.Enumeration_stagedOrder[enumeration] = order
		stage.Enumeration_orderStaged[order] = enumeration
		stage.EnumerationOrder++
	}
	stage.Enumerations_mapString[enumeration.Name] = enumeration
}

// Unstage removes enumeration off the model stage
func (enumeration *Enumeration) Unstage(stage *Stage) *Enumeration {
	delete(stage.Enumerations, enumeration)
	// issue1150
	// delete(stage.Enumeration_stagedOrder, enumeration)
	delete(stage.Enumerations_mapString, enumeration.Name)

	return enumeration
}

// UnstageVoid removes enumeration off the model stage
func (enumeration *Enumeration) UnstageVoid(stage *Stage) {
	delete(stage.Enumerations, enumeration)
	// issue1150
	// delete(stage.Enumeration_stagedOrder, enumeration)
	delete(stage.Enumerations_mapString, enumeration.Name)
}

// commit enumeration to the back repo (if it is already staged)
func (enumeration *Enumeration) Commit(stage *Stage) *Enumeration {
	if _, ok := stage.Enumerations[enumeration]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitEnumeration(enumeration)
		}
	}
	return enumeration
}

func (enumeration *Enumeration) CommitVoid(stage *Stage) {
	enumeration.Commit(stage)
}

func (enumeration *Enumeration) StageVoid(stage *Stage) {
	enumeration.Stage(stage)
}

// Checkout enumeration to the back repo (if it is already staged)
func (enumeration *Enumeration) Checkout(stage *Stage) *Enumeration {
	if _, ok := stage.Enumerations[enumeration]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutEnumeration(enumeration)
		}
	}
	return enumeration
}

// for satisfaction of GongStruct interface
func (enumeration *Enumeration) GetName() (res string) {
	return enumeration.Name
}

// for satisfaction of GongStruct interface
func (enumeration *Enumeration) SetName(name string) {
	enumeration.Name = name
}

// Stage puts extension to the model stage
func (extension *Extension) Stage(stage *Stage) *Extension {
	if _, ok := stage.Extensions[extension]; !ok {
		stage.Extensions[extension] = struct{}{}
		stage.Extension_stagedOrder[extension] = stage.ExtensionOrder
		stage.Extension_orderStaged[stage.ExtensionOrder] = extension
		stage.ExtensionOrder++
	}
	stage.Extensions_mapString[extension.Name] = extension

	return extension
}

// StagePreserveOrder puts extension to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ExtensionOrder
// - update stage.ExtensionOrder accordingly
func (extension *Extension) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Extensions[extension]; !ok {
		stage.Extensions[extension] = struct{}{}

		if order > stage.ExtensionOrder {
			stage.ExtensionOrder = order
		}
		stage.Extension_stagedOrder[extension] = order
		stage.Extension_orderStaged[order] = extension
		stage.ExtensionOrder++
	}
	stage.Extensions_mapString[extension.Name] = extension
}

// Unstage removes extension off the model stage
func (extension *Extension) Unstage(stage *Stage) *Extension {
	delete(stage.Extensions, extension)
	// issue1150
	// delete(stage.Extension_stagedOrder, extension)
	delete(stage.Extensions_mapString, extension.Name)

	return extension
}

// UnstageVoid removes extension off the model stage
func (extension *Extension) UnstageVoid(stage *Stage) {
	delete(stage.Extensions, extension)
	// issue1150
	// delete(stage.Extension_stagedOrder, extension)
	delete(stage.Extensions_mapString, extension.Name)
}

// commit extension to the back repo (if it is already staged)
func (extension *Extension) Commit(stage *Stage) *Extension {
	if _, ok := stage.Extensions[extension]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitExtension(extension)
		}
	}
	return extension
}

func (extension *Extension) CommitVoid(stage *Stage) {
	extension.Commit(stage)
}

func (extension *Extension) StageVoid(stage *Stage) {
	extension.Stage(stage)
}

// Checkout extension to the back repo (if it is already staged)
func (extension *Extension) Checkout(stage *Stage) *Extension {
	if _, ok := stage.Extensions[extension]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutExtension(extension)
		}
	}
	return extension
}

// for satisfaction of GongStruct interface
func (extension *Extension) GetName() (res string) {
	return extension.Name
}

// for satisfaction of GongStruct interface
func (extension *Extension) SetName(name string) {
	extension.Name = name
}

// Stage puts group to the model stage
func (group *Group) Stage(stage *Stage) *Group {
	if _, ok := stage.Groups[group]; !ok {
		stage.Groups[group] = struct{}{}
		stage.Group_stagedOrder[group] = stage.GroupOrder
		stage.Group_orderStaged[stage.GroupOrder] = group
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
		stage.Group_stagedOrder[group] = order
		stage.Group_orderStaged[order] = group
		stage.GroupOrder++
	}
	stage.Groups_mapString[group.Name] = group
}

// Unstage removes group off the model stage
func (group *Group) Unstage(stage *Stage) *Group {
	delete(stage.Groups, group)
	// issue1150
	// delete(stage.Group_stagedOrder, group)
	delete(stage.Groups_mapString, group.Name)

	return group
}

// UnstageVoid removes group off the model stage
func (group *Group) UnstageVoid(stage *Stage) {
	delete(stage.Groups, group)
	// issue1150
	// delete(stage.Group_stagedOrder, group)
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

// Stage puts length to the model stage
func (length *Length) Stage(stage *Stage) *Length {
	if _, ok := stage.Lengths[length]; !ok {
		stage.Lengths[length] = struct{}{}
		stage.Length_stagedOrder[length] = stage.LengthOrder
		stage.Length_orderStaged[stage.LengthOrder] = length
		stage.LengthOrder++
	}
	stage.Lengths_mapString[length.Name] = length

	return length
}

// StagePreserveOrder puts length to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.LengthOrder
// - update stage.LengthOrder accordingly
func (length *Length) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Lengths[length]; !ok {
		stage.Lengths[length] = struct{}{}

		if order > stage.LengthOrder {
			stage.LengthOrder = order
		}
		stage.Length_stagedOrder[length] = order
		stage.Length_orderStaged[order] = length
		stage.LengthOrder++
	}
	stage.Lengths_mapString[length.Name] = length
}

// Unstage removes length off the model stage
func (length *Length) Unstage(stage *Stage) *Length {
	delete(stage.Lengths, length)
	// issue1150
	// delete(stage.Length_stagedOrder, length)
	delete(stage.Lengths_mapString, length.Name)

	return length
}

// UnstageVoid removes length off the model stage
func (length *Length) UnstageVoid(stage *Stage) {
	delete(stage.Lengths, length)
	// issue1150
	// delete(stage.Length_stagedOrder, length)
	delete(stage.Lengths_mapString, length.Name)
}

// commit length to the back repo (if it is already staged)
func (length *Length) Commit(stage *Stage) *Length {
	if _, ok := stage.Lengths[length]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitLength(length)
		}
	}
	return length
}

func (length *Length) CommitVoid(stage *Stage) {
	length.Commit(stage)
}

func (length *Length) StageVoid(stage *Stage) {
	length.Stage(stage)
}

// Checkout length to the back repo (if it is already staged)
func (length *Length) Checkout(stage *Stage) *Length {
	if _, ok := stage.Lengths[length]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutLength(length)
		}
	}
	return length
}

// for satisfaction of GongStruct interface
func (length *Length) GetName() (res string) {
	return length.Name
}

// for satisfaction of GongStruct interface
func (length *Length) SetName(name string) {
	length.Name = name
}

// Stage puts maxinclusive to the model stage
func (maxinclusive *MaxInclusive) Stage(stage *Stage) *MaxInclusive {
	if _, ok := stage.MaxInclusives[maxinclusive]; !ok {
		stage.MaxInclusives[maxinclusive] = struct{}{}
		stage.MaxInclusive_stagedOrder[maxinclusive] = stage.MaxInclusiveOrder
		stage.MaxInclusive_orderStaged[stage.MaxInclusiveOrder] = maxinclusive
		stage.MaxInclusiveOrder++
	}
	stage.MaxInclusives_mapString[maxinclusive.Name] = maxinclusive

	return maxinclusive
}

// StagePreserveOrder puts maxinclusive to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MaxInclusiveOrder
// - update stage.MaxInclusiveOrder accordingly
func (maxinclusive *MaxInclusive) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.MaxInclusives[maxinclusive]; !ok {
		stage.MaxInclusives[maxinclusive] = struct{}{}

		if order > stage.MaxInclusiveOrder {
			stage.MaxInclusiveOrder = order
		}
		stage.MaxInclusive_stagedOrder[maxinclusive] = order
		stage.MaxInclusive_orderStaged[order] = maxinclusive
		stage.MaxInclusiveOrder++
	}
	stage.MaxInclusives_mapString[maxinclusive.Name] = maxinclusive
}

// Unstage removes maxinclusive off the model stage
func (maxinclusive *MaxInclusive) Unstage(stage *Stage) *MaxInclusive {
	delete(stage.MaxInclusives, maxinclusive)
	// issue1150
	// delete(stage.MaxInclusive_stagedOrder, maxinclusive)
	delete(stage.MaxInclusives_mapString, maxinclusive.Name)

	return maxinclusive
}

// UnstageVoid removes maxinclusive off the model stage
func (maxinclusive *MaxInclusive) UnstageVoid(stage *Stage) {
	delete(stage.MaxInclusives, maxinclusive)
	// issue1150
	// delete(stage.MaxInclusive_stagedOrder, maxinclusive)
	delete(stage.MaxInclusives_mapString, maxinclusive.Name)
}

// commit maxinclusive to the back repo (if it is already staged)
func (maxinclusive *MaxInclusive) Commit(stage *Stage) *MaxInclusive {
	if _, ok := stage.MaxInclusives[maxinclusive]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitMaxInclusive(maxinclusive)
		}
	}
	return maxinclusive
}

func (maxinclusive *MaxInclusive) CommitVoid(stage *Stage) {
	maxinclusive.Commit(stage)
}

func (maxinclusive *MaxInclusive) StageVoid(stage *Stage) {
	maxinclusive.Stage(stage)
}

// Checkout maxinclusive to the back repo (if it is already staged)
func (maxinclusive *MaxInclusive) Checkout(stage *Stage) *MaxInclusive {
	if _, ok := stage.MaxInclusives[maxinclusive]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutMaxInclusive(maxinclusive)
		}
	}
	return maxinclusive
}

// for satisfaction of GongStruct interface
func (maxinclusive *MaxInclusive) GetName() (res string) {
	return maxinclusive.Name
}

// for satisfaction of GongStruct interface
func (maxinclusive *MaxInclusive) SetName(name string) {
	maxinclusive.Name = name
}

// Stage puts maxlength to the model stage
func (maxlength *MaxLength) Stage(stage *Stage) *MaxLength {
	if _, ok := stage.MaxLengths[maxlength]; !ok {
		stage.MaxLengths[maxlength] = struct{}{}
		stage.MaxLength_stagedOrder[maxlength] = stage.MaxLengthOrder
		stage.MaxLength_orderStaged[stage.MaxLengthOrder] = maxlength
		stage.MaxLengthOrder++
	}
	stage.MaxLengths_mapString[maxlength.Name] = maxlength

	return maxlength
}

// StagePreserveOrder puts maxlength to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MaxLengthOrder
// - update stage.MaxLengthOrder accordingly
func (maxlength *MaxLength) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.MaxLengths[maxlength]; !ok {
		stage.MaxLengths[maxlength] = struct{}{}

		if order > stage.MaxLengthOrder {
			stage.MaxLengthOrder = order
		}
		stage.MaxLength_stagedOrder[maxlength] = order
		stage.MaxLength_orderStaged[order] = maxlength
		stage.MaxLengthOrder++
	}
	stage.MaxLengths_mapString[maxlength.Name] = maxlength
}

// Unstage removes maxlength off the model stage
func (maxlength *MaxLength) Unstage(stage *Stage) *MaxLength {
	delete(stage.MaxLengths, maxlength)
	// issue1150
	// delete(stage.MaxLength_stagedOrder, maxlength)
	delete(stage.MaxLengths_mapString, maxlength.Name)

	return maxlength
}

// UnstageVoid removes maxlength off the model stage
func (maxlength *MaxLength) UnstageVoid(stage *Stage) {
	delete(stage.MaxLengths, maxlength)
	// issue1150
	// delete(stage.MaxLength_stagedOrder, maxlength)
	delete(stage.MaxLengths_mapString, maxlength.Name)
}

// commit maxlength to the back repo (if it is already staged)
func (maxlength *MaxLength) Commit(stage *Stage) *MaxLength {
	if _, ok := stage.MaxLengths[maxlength]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitMaxLength(maxlength)
		}
	}
	return maxlength
}

func (maxlength *MaxLength) CommitVoid(stage *Stage) {
	maxlength.Commit(stage)
}

func (maxlength *MaxLength) StageVoid(stage *Stage) {
	maxlength.Stage(stage)
}

// Checkout maxlength to the back repo (if it is already staged)
func (maxlength *MaxLength) Checkout(stage *Stage) *MaxLength {
	if _, ok := stage.MaxLengths[maxlength]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutMaxLength(maxlength)
		}
	}
	return maxlength
}

// for satisfaction of GongStruct interface
func (maxlength *MaxLength) GetName() (res string) {
	return maxlength.Name
}

// for satisfaction of GongStruct interface
func (maxlength *MaxLength) SetName(name string) {
	maxlength.Name = name
}

// Stage puts mininclusive to the model stage
func (mininclusive *MinInclusive) Stage(stage *Stage) *MinInclusive {
	if _, ok := stage.MinInclusives[mininclusive]; !ok {
		stage.MinInclusives[mininclusive] = struct{}{}
		stage.MinInclusive_stagedOrder[mininclusive] = stage.MinInclusiveOrder
		stage.MinInclusive_orderStaged[stage.MinInclusiveOrder] = mininclusive
		stage.MinInclusiveOrder++
	}
	stage.MinInclusives_mapString[mininclusive.Name] = mininclusive

	return mininclusive
}

// StagePreserveOrder puts mininclusive to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MinInclusiveOrder
// - update stage.MinInclusiveOrder accordingly
func (mininclusive *MinInclusive) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.MinInclusives[mininclusive]; !ok {
		stage.MinInclusives[mininclusive] = struct{}{}

		if order > stage.MinInclusiveOrder {
			stage.MinInclusiveOrder = order
		}
		stage.MinInclusive_stagedOrder[mininclusive] = order
		stage.MinInclusive_orderStaged[order] = mininclusive
		stage.MinInclusiveOrder++
	}
	stage.MinInclusives_mapString[mininclusive.Name] = mininclusive
}

// Unstage removes mininclusive off the model stage
func (mininclusive *MinInclusive) Unstage(stage *Stage) *MinInclusive {
	delete(stage.MinInclusives, mininclusive)
	// issue1150
	// delete(stage.MinInclusive_stagedOrder, mininclusive)
	delete(stage.MinInclusives_mapString, mininclusive.Name)

	return mininclusive
}

// UnstageVoid removes mininclusive off the model stage
func (mininclusive *MinInclusive) UnstageVoid(stage *Stage) {
	delete(stage.MinInclusives, mininclusive)
	// issue1150
	// delete(stage.MinInclusive_stagedOrder, mininclusive)
	delete(stage.MinInclusives_mapString, mininclusive.Name)
}

// commit mininclusive to the back repo (if it is already staged)
func (mininclusive *MinInclusive) Commit(stage *Stage) *MinInclusive {
	if _, ok := stage.MinInclusives[mininclusive]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitMinInclusive(mininclusive)
		}
	}
	return mininclusive
}

func (mininclusive *MinInclusive) CommitVoid(stage *Stage) {
	mininclusive.Commit(stage)
}

func (mininclusive *MinInclusive) StageVoid(stage *Stage) {
	mininclusive.Stage(stage)
}

// Checkout mininclusive to the back repo (if it is already staged)
func (mininclusive *MinInclusive) Checkout(stage *Stage) *MinInclusive {
	if _, ok := stage.MinInclusives[mininclusive]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutMinInclusive(mininclusive)
		}
	}
	return mininclusive
}

// for satisfaction of GongStruct interface
func (mininclusive *MinInclusive) GetName() (res string) {
	return mininclusive.Name
}

// for satisfaction of GongStruct interface
func (mininclusive *MinInclusive) SetName(name string) {
	mininclusive.Name = name
}

// Stage puts minlength to the model stage
func (minlength *MinLength) Stage(stage *Stage) *MinLength {
	if _, ok := stage.MinLengths[minlength]; !ok {
		stage.MinLengths[minlength] = struct{}{}
		stage.MinLength_stagedOrder[minlength] = stage.MinLengthOrder
		stage.MinLength_orderStaged[stage.MinLengthOrder] = minlength
		stage.MinLengthOrder++
	}
	stage.MinLengths_mapString[minlength.Name] = minlength

	return minlength
}

// StagePreserveOrder puts minlength to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MinLengthOrder
// - update stage.MinLengthOrder accordingly
func (minlength *MinLength) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.MinLengths[minlength]; !ok {
		stage.MinLengths[minlength] = struct{}{}

		if order > stage.MinLengthOrder {
			stage.MinLengthOrder = order
		}
		stage.MinLength_stagedOrder[minlength] = order
		stage.MinLength_orderStaged[order] = minlength
		stage.MinLengthOrder++
	}
	stage.MinLengths_mapString[minlength.Name] = minlength
}

// Unstage removes minlength off the model stage
func (minlength *MinLength) Unstage(stage *Stage) *MinLength {
	delete(stage.MinLengths, minlength)
	// issue1150
	// delete(stage.MinLength_stagedOrder, minlength)
	delete(stage.MinLengths_mapString, minlength.Name)

	return minlength
}

// UnstageVoid removes minlength off the model stage
func (minlength *MinLength) UnstageVoid(stage *Stage) {
	delete(stage.MinLengths, minlength)
	// issue1150
	// delete(stage.MinLength_stagedOrder, minlength)
	delete(stage.MinLengths_mapString, minlength.Name)
}

// commit minlength to the back repo (if it is already staged)
func (minlength *MinLength) Commit(stage *Stage) *MinLength {
	if _, ok := stage.MinLengths[minlength]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitMinLength(minlength)
		}
	}
	return minlength
}

func (minlength *MinLength) CommitVoid(stage *Stage) {
	minlength.Commit(stage)
}

func (minlength *MinLength) StageVoid(stage *Stage) {
	minlength.Stage(stage)
}

// Checkout minlength to the back repo (if it is already staged)
func (minlength *MinLength) Checkout(stage *Stage) *MinLength {
	if _, ok := stage.MinLengths[minlength]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutMinLength(minlength)
		}
	}
	return minlength
}

// for satisfaction of GongStruct interface
func (minlength *MinLength) GetName() (res string) {
	return minlength.Name
}

// for satisfaction of GongStruct interface
func (minlength *MinLength) SetName(name string) {
	minlength.Name = name
}

// Stage puts pattern to the model stage
func (pattern *Pattern) Stage(stage *Stage) *Pattern {
	if _, ok := stage.Patterns[pattern]; !ok {
		stage.Patterns[pattern] = struct{}{}
		stage.Pattern_stagedOrder[pattern] = stage.PatternOrder
		stage.Pattern_orderStaged[stage.PatternOrder] = pattern
		stage.PatternOrder++
	}
	stage.Patterns_mapString[pattern.Name] = pattern

	return pattern
}

// StagePreserveOrder puts pattern to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.PatternOrder
// - update stage.PatternOrder accordingly
func (pattern *Pattern) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Patterns[pattern]; !ok {
		stage.Patterns[pattern] = struct{}{}

		if order > stage.PatternOrder {
			stage.PatternOrder = order
		}
		stage.Pattern_stagedOrder[pattern] = order
		stage.Pattern_orderStaged[order] = pattern
		stage.PatternOrder++
	}
	stage.Patterns_mapString[pattern.Name] = pattern
}

// Unstage removes pattern off the model stage
func (pattern *Pattern) Unstage(stage *Stage) *Pattern {
	delete(stage.Patterns, pattern)
	// issue1150
	// delete(stage.Pattern_stagedOrder, pattern)
	delete(stage.Patterns_mapString, pattern.Name)

	return pattern
}

// UnstageVoid removes pattern off the model stage
func (pattern *Pattern) UnstageVoid(stage *Stage) {
	delete(stage.Patterns, pattern)
	// issue1150
	// delete(stage.Pattern_stagedOrder, pattern)
	delete(stage.Patterns_mapString, pattern.Name)
}

// commit pattern to the back repo (if it is already staged)
func (pattern *Pattern) Commit(stage *Stage) *Pattern {
	if _, ok := stage.Patterns[pattern]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitPattern(pattern)
		}
	}
	return pattern
}

func (pattern *Pattern) CommitVoid(stage *Stage) {
	pattern.Commit(stage)
}

func (pattern *Pattern) StageVoid(stage *Stage) {
	pattern.Stage(stage)
}

// Checkout pattern to the back repo (if it is already staged)
func (pattern *Pattern) Checkout(stage *Stage) *Pattern {
	if _, ok := stage.Patterns[pattern]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutPattern(pattern)
		}
	}
	return pattern
}

// for satisfaction of GongStruct interface
func (pattern *Pattern) GetName() (res string) {
	return pattern.Name
}

// for satisfaction of GongStruct interface
func (pattern *Pattern) SetName(name string) {
	pattern.Name = name
}

// Stage puts restriction to the model stage
func (restriction *Restriction) Stage(stage *Stage) *Restriction {
	if _, ok := stage.Restrictions[restriction]; !ok {
		stage.Restrictions[restriction] = struct{}{}
		stage.Restriction_stagedOrder[restriction] = stage.RestrictionOrder
		stage.Restriction_orderStaged[stage.RestrictionOrder] = restriction
		stage.RestrictionOrder++
	}
	stage.Restrictions_mapString[restriction.Name] = restriction

	return restriction
}

// StagePreserveOrder puts restriction to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.RestrictionOrder
// - update stage.RestrictionOrder accordingly
func (restriction *Restriction) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Restrictions[restriction]; !ok {
		stage.Restrictions[restriction] = struct{}{}

		if order > stage.RestrictionOrder {
			stage.RestrictionOrder = order
		}
		stage.Restriction_stagedOrder[restriction] = order
		stage.Restriction_orderStaged[order] = restriction
		stage.RestrictionOrder++
	}
	stage.Restrictions_mapString[restriction.Name] = restriction
}

// Unstage removes restriction off the model stage
func (restriction *Restriction) Unstage(stage *Stage) *Restriction {
	delete(stage.Restrictions, restriction)
	// issue1150
	// delete(stage.Restriction_stagedOrder, restriction)
	delete(stage.Restrictions_mapString, restriction.Name)

	return restriction
}

// UnstageVoid removes restriction off the model stage
func (restriction *Restriction) UnstageVoid(stage *Stage) {
	delete(stage.Restrictions, restriction)
	// issue1150
	// delete(stage.Restriction_stagedOrder, restriction)
	delete(stage.Restrictions_mapString, restriction.Name)
}

// commit restriction to the back repo (if it is already staged)
func (restriction *Restriction) Commit(stage *Stage) *Restriction {
	if _, ok := stage.Restrictions[restriction]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitRestriction(restriction)
		}
	}
	return restriction
}

func (restriction *Restriction) CommitVoid(stage *Stage) {
	restriction.Commit(stage)
}

func (restriction *Restriction) StageVoid(stage *Stage) {
	restriction.Stage(stage)
}

// Checkout restriction to the back repo (if it is already staged)
func (restriction *Restriction) Checkout(stage *Stage) *Restriction {
	if _, ok := stage.Restrictions[restriction]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutRestriction(restriction)
		}
	}
	return restriction
}

// for satisfaction of GongStruct interface
func (restriction *Restriction) GetName() (res string) {
	return restriction.Name
}

// for satisfaction of GongStruct interface
func (restriction *Restriction) SetName(name string) {
	restriction.Name = name
}

// Stage puts schema to the model stage
func (schema *Schema) Stage(stage *Stage) *Schema {
	if _, ok := stage.Schemas[schema]; !ok {
		stage.Schemas[schema] = struct{}{}
		stage.Schema_stagedOrder[schema] = stage.SchemaOrder
		stage.Schema_orderStaged[stage.SchemaOrder] = schema
		stage.SchemaOrder++
	}
	stage.Schemas_mapString[schema.Name] = schema

	return schema
}

// StagePreserveOrder puts schema to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.SchemaOrder
// - update stage.SchemaOrder accordingly
func (schema *Schema) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Schemas[schema]; !ok {
		stage.Schemas[schema] = struct{}{}

		if order > stage.SchemaOrder {
			stage.SchemaOrder = order
		}
		stage.Schema_stagedOrder[schema] = order
		stage.Schema_orderStaged[order] = schema
		stage.SchemaOrder++
	}
	stage.Schemas_mapString[schema.Name] = schema
}

// Unstage removes schema off the model stage
func (schema *Schema) Unstage(stage *Stage) *Schema {
	delete(stage.Schemas, schema)
	// issue1150
	// delete(stage.Schema_stagedOrder, schema)
	delete(stage.Schemas_mapString, schema.Name)

	return schema
}

// UnstageVoid removes schema off the model stage
func (schema *Schema) UnstageVoid(stage *Stage) {
	delete(stage.Schemas, schema)
	// issue1150
	// delete(stage.Schema_stagedOrder, schema)
	delete(stage.Schemas_mapString, schema.Name)
}

// commit schema to the back repo (if it is already staged)
func (schema *Schema) Commit(stage *Stage) *Schema {
	if _, ok := stage.Schemas[schema]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSchema(schema)
		}
	}
	return schema
}

func (schema *Schema) CommitVoid(stage *Stage) {
	schema.Commit(stage)
}

func (schema *Schema) StageVoid(stage *Stage) {
	schema.Stage(stage)
}

// Checkout schema to the back repo (if it is already staged)
func (schema *Schema) Checkout(stage *Stage) *Schema {
	if _, ok := stage.Schemas[schema]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSchema(schema)
		}
	}
	return schema
}

// for satisfaction of GongStruct interface
func (schema *Schema) GetName() (res string) {
	return schema.Name
}

// for satisfaction of GongStruct interface
func (schema *Schema) SetName(name string) {
	schema.Name = name
}

// Stage puts sequence to the model stage
func (sequence *Sequence) Stage(stage *Stage) *Sequence {
	if _, ok := stage.Sequences[sequence]; !ok {
		stage.Sequences[sequence] = struct{}{}
		stage.Sequence_stagedOrder[sequence] = stage.SequenceOrder
		stage.Sequence_orderStaged[stage.SequenceOrder] = sequence
		stage.SequenceOrder++
	}
	stage.Sequences_mapString[sequence.Name] = sequence

	return sequence
}

// StagePreserveOrder puts sequence to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.SequenceOrder
// - update stage.SequenceOrder accordingly
func (sequence *Sequence) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Sequences[sequence]; !ok {
		stage.Sequences[sequence] = struct{}{}

		if order > stage.SequenceOrder {
			stage.SequenceOrder = order
		}
		stage.Sequence_stagedOrder[sequence] = order
		stage.Sequence_orderStaged[order] = sequence
		stage.SequenceOrder++
	}
	stage.Sequences_mapString[sequence.Name] = sequence
}

// Unstage removes sequence off the model stage
func (sequence *Sequence) Unstage(stage *Stage) *Sequence {
	delete(stage.Sequences, sequence)
	// issue1150
	// delete(stage.Sequence_stagedOrder, sequence)
	delete(stage.Sequences_mapString, sequence.Name)

	return sequence
}

// UnstageVoid removes sequence off the model stage
func (sequence *Sequence) UnstageVoid(stage *Stage) {
	delete(stage.Sequences, sequence)
	// issue1150
	// delete(stage.Sequence_stagedOrder, sequence)
	delete(stage.Sequences_mapString, sequence.Name)
}

// commit sequence to the back repo (if it is already staged)
func (sequence *Sequence) Commit(stage *Stage) *Sequence {
	if _, ok := stage.Sequences[sequence]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSequence(sequence)
		}
	}
	return sequence
}

func (sequence *Sequence) CommitVoid(stage *Stage) {
	sequence.Commit(stage)
}

func (sequence *Sequence) StageVoid(stage *Stage) {
	sequence.Stage(stage)
}

// Checkout sequence to the back repo (if it is already staged)
func (sequence *Sequence) Checkout(stage *Stage) *Sequence {
	if _, ok := stage.Sequences[sequence]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSequence(sequence)
		}
	}
	return sequence
}

// for satisfaction of GongStruct interface
func (sequence *Sequence) GetName() (res string) {
	return sequence.Name
}

// for satisfaction of GongStruct interface
func (sequence *Sequence) SetName(name string) {
	sequence.Name = name
}

// Stage puts simplecontent to the model stage
func (simplecontent *SimpleContent) Stage(stage *Stage) *SimpleContent {
	if _, ok := stage.SimpleContents[simplecontent]; !ok {
		stage.SimpleContents[simplecontent] = struct{}{}
		stage.SimpleContent_stagedOrder[simplecontent] = stage.SimpleContentOrder
		stage.SimpleContent_orderStaged[stage.SimpleContentOrder] = simplecontent
		stage.SimpleContentOrder++
	}
	stage.SimpleContents_mapString[simplecontent.Name] = simplecontent

	return simplecontent
}

// StagePreserveOrder puts simplecontent to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.SimpleContentOrder
// - update stage.SimpleContentOrder accordingly
func (simplecontent *SimpleContent) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.SimpleContents[simplecontent]; !ok {
		stage.SimpleContents[simplecontent] = struct{}{}

		if order > stage.SimpleContentOrder {
			stage.SimpleContentOrder = order
		}
		stage.SimpleContent_stagedOrder[simplecontent] = order
		stage.SimpleContent_orderStaged[order] = simplecontent
		stage.SimpleContentOrder++
	}
	stage.SimpleContents_mapString[simplecontent.Name] = simplecontent
}

// Unstage removes simplecontent off the model stage
func (simplecontent *SimpleContent) Unstage(stage *Stage) *SimpleContent {
	delete(stage.SimpleContents, simplecontent)
	// issue1150
	// delete(stage.SimpleContent_stagedOrder, simplecontent)
	delete(stage.SimpleContents_mapString, simplecontent.Name)

	return simplecontent
}

// UnstageVoid removes simplecontent off the model stage
func (simplecontent *SimpleContent) UnstageVoid(stage *Stage) {
	delete(stage.SimpleContents, simplecontent)
	// issue1150
	// delete(stage.SimpleContent_stagedOrder, simplecontent)
	delete(stage.SimpleContents_mapString, simplecontent.Name)
}

// commit simplecontent to the back repo (if it is already staged)
func (simplecontent *SimpleContent) Commit(stage *Stage) *SimpleContent {
	if _, ok := stage.SimpleContents[simplecontent]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSimpleContent(simplecontent)
		}
	}
	return simplecontent
}

func (simplecontent *SimpleContent) CommitVoid(stage *Stage) {
	simplecontent.Commit(stage)
}

func (simplecontent *SimpleContent) StageVoid(stage *Stage) {
	simplecontent.Stage(stage)
}

// Checkout simplecontent to the back repo (if it is already staged)
func (simplecontent *SimpleContent) Checkout(stage *Stage) *SimpleContent {
	if _, ok := stage.SimpleContents[simplecontent]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSimpleContent(simplecontent)
		}
	}
	return simplecontent
}

// for satisfaction of GongStruct interface
func (simplecontent *SimpleContent) GetName() (res string) {
	return simplecontent.Name
}

// for satisfaction of GongStruct interface
func (simplecontent *SimpleContent) SetName(name string) {
	simplecontent.Name = name
}

// Stage puts simpletype to the model stage
func (simpletype *SimpleType) Stage(stage *Stage) *SimpleType {
	if _, ok := stage.SimpleTypes[simpletype]; !ok {
		stage.SimpleTypes[simpletype] = struct{}{}
		stage.SimpleType_stagedOrder[simpletype] = stage.SimpleTypeOrder
		stage.SimpleType_orderStaged[stage.SimpleTypeOrder] = simpletype
		stage.SimpleTypeOrder++
	}
	stage.SimpleTypes_mapString[simpletype.Name] = simpletype

	return simpletype
}

// StagePreserveOrder puts simpletype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.SimpleTypeOrder
// - update stage.SimpleTypeOrder accordingly
func (simpletype *SimpleType) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.SimpleTypes[simpletype]; !ok {
		stage.SimpleTypes[simpletype] = struct{}{}

		if order > stage.SimpleTypeOrder {
			stage.SimpleTypeOrder = order
		}
		stage.SimpleType_stagedOrder[simpletype] = order
		stage.SimpleType_orderStaged[order] = simpletype
		stage.SimpleTypeOrder++
	}
	stage.SimpleTypes_mapString[simpletype.Name] = simpletype
}

// Unstage removes simpletype off the model stage
func (simpletype *SimpleType) Unstage(stage *Stage) *SimpleType {
	delete(stage.SimpleTypes, simpletype)
	// issue1150
	// delete(stage.SimpleType_stagedOrder, simpletype)
	delete(stage.SimpleTypes_mapString, simpletype.Name)

	return simpletype
}

// UnstageVoid removes simpletype off the model stage
func (simpletype *SimpleType) UnstageVoid(stage *Stage) {
	delete(stage.SimpleTypes, simpletype)
	// issue1150
	// delete(stage.SimpleType_stagedOrder, simpletype)
	delete(stage.SimpleTypes_mapString, simpletype.Name)
}

// commit simpletype to the back repo (if it is already staged)
func (simpletype *SimpleType) Commit(stage *Stage) *SimpleType {
	if _, ok := stage.SimpleTypes[simpletype]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSimpleType(simpletype)
		}
	}
	return simpletype
}

func (simpletype *SimpleType) CommitVoid(stage *Stage) {
	simpletype.Commit(stage)
}

func (simpletype *SimpleType) StageVoid(stage *Stage) {
	simpletype.Stage(stage)
}

// Checkout simpletype to the back repo (if it is already staged)
func (simpletype *SimpleType) Checkout(stage *Stage) *SimpleType {
	if _, ok := stage.SimpleTypes[simpletype]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSimpleType(simpletype)
		}
	}
	return simpletype
}

// for satisfaction of GongStruct interface
func (simpletype *SimpleType) GetName() (res string) {
	return simpletype.Name
}

// for satisfaction of GongStruct interface
func (simpletype *SimpleType) SetName(name string) {
	simpletype.Name = name
}

// Stage puts totaldigit to the model stage
func (totaldigit *TotalDigit) Stage(stage *Stage) *TotalDigit {
	if _, ok := stage.TotalDigits[totaldigit]; !ok {
		stage.TotalDigits[totaldigit] = struct{}{}
		stage.TotalDigit_stagedOrder[totaldigit] = stage.TotalDigitOrder
		stage.TotalDigit_orderStaged[stage.TotalDigitOrder] = totaldigit
		stage.TotalDigitOrder++
	}
	stage.TotalDigits_mapString[totaldigit.Name] = totaldigit

	return totaldigit
}

// StagePreserveOrder puts totaldigit to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TotalDigitOrder
// - update stage.TotalDigitOrder accordingly
func (totaldigit *TotalDigit) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.TotalDigits[totaldigit]; !ok {
		stage.TotalDigits[totaldigit] = struct{}{}

		if order > stage.TotalDigitOrder {
			stage.TotalDigitOrder = order
		}
		stage.TotalDigit_stagedOrder[totaldigit] = order
		stage.TotalDigit_orderStaged[order] = totaldigit
		stage.TotalDigitOrder++
	}
	stage.TotalDigits_mapString[totaldigit.Name] = totaldigit
}

// Unstage removes totaldigit off the model stage
func (totaldigit *TotalDigit) Unstage(stage *Stage) *TotalDigit {
	delete(stage.TotalDigits, totaldigit)
	// issue1150
	// delete(stage.TotalDigit_stagedOrder, totaldigit)
	delete(stage.TotalDigits_mapString, totaldigit.Name)

	return totaldigit
}

// UnstageVoid removes totaldigit off the model stage
func (totaldigit *TotalDigit) UnstageVoid(stage *Stage) {
	delete(stage.TotalDigits, totaldigit)
	// issue1150
	// delete(stage.TotalDigit_stagedOrder, totaldigit)
	delete(stage.TotalDigits_mapString, totaldigit.Name)
}

// commit totaldigit to the back repo (if it is already staged)
func (totaldigit *TotalDigit) Commit(stage *Stage) *TotalDigit {
	if _, ok := stage.TotalDigits[totaldigit]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTotalDigit(totaldigit)
		}
	}
	return totaldigit
}

func (totaldigit *TotalDigit) CommitVoid(stage *Stage) {
	totaldigit.Commit(stage)
}

func (totaldigit *TotalDigit) StageVoid(stage *Stage) {
	totaldigit.Stage(stage)
}

// Checkout totaldigit to the back repo (if it is already staged)
func (totaldigit *TotalDigit) Checkout(stage *Stage) *TotalDigit {
	if _, ok := stage.TotalDigits[totaldigit]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutTotalDigit(totaldigit)
		}
	}
	return totaldigit
}

// for satisfaction of GongStruct interface
func (totaldigit *TotalDigit) GetName() (res string) {
	return totaldigit.Name
}

// for satisfaction of GongStruct interface
func (totaldigit *TotalDigit) SetName(name string) {
	totaldigit.Name = name
}

// Stage puts union to the model stage
func (union *Union) Stage(stage *Stage) *Union {
	if _, ok := stage.Unions[union]; !ok {
		stage.Unions[union] = struct{}{}
		stage.Union_stagedOrder[union] = stage.UnionOrder
		stage.Union_orderStaged[stage.UnionOrder] = union
		stage.UnionOrder++
	}
	stage.Unions_mapString[union.Name] = union

	return union
}

// StagePreserveOrder puts union to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.UnionOrder
// - update stage.UnionOrder accordingly
func (union *Union) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Unions[union]; !ok {
		stage.Unions[union] = struct{}{}

		if order > stage.UnionOrder {
			stage.UnionOrder = order
		}
		stage.Union_stagedOrder[union] = order
		stage.Union_orderStaged[order] = union
		stage.UnionOrder++
	}
	stage.Unions_mapString[union.Name] = union
}

// Unstage removes union off the model stage
func (union *Union) Unstage(stage *Stage) *Union {
	delete(stage.Unions, union)
	// issue1150
	// delete(stage.Union_stagedOrder, union)
	delete(stage.Unions_mapString, union.Name)

	return union
}

// UnstageVoid removes union off the model stage
func (union *Union) UnstageVoid(stage *Stage) {
	delete(stage.Unions, union)
	// issue1150
	// delete(stage.Union_stagedOrder, union)
	delete(stage.Unions_mapString, union.Name)
}

// commit union to the back repo (if it is already staged)
func (union *Union) Commit(stage *Stage) *Union {
	if _, ok := stage.Unions[union]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitUnion(union)
		}
	}
	return union
}

func (union *Union) CommitVoid(stage *Stage) {
	union.Commit(stage)
}

func (union *Union) StageVoid(stage *Stage) {
	union.Stage(stage)
}

// Checkout union to the back repo (if it is already staged)
func (union *Union) Checkout(stage *Stage) *Union {
	if _, ok := stage.Unions[union]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutUnion(union)
		}
	}
	return union
}

// for satisfaction of GongStruct interface
func (union *Union) GetName() (res string) {
	return union.Name
}

// for satisfaction of GongStruct interface
func (union *Union) SetName(name string) {
	union.Name = name
}

// Stage puts whitespace to the model stage
func (whitespace *WhiteSpace) Stage(stage *Stage) *WhiteSpace {
	if _, ok := stage.WhiteSpaces[whitespace]; !ok {
		stage.WhiteSpaces[whitespace] = struct{}{}
		stage.WhiteSpace_stagedOrder[whitespace] = stage.WhiteSpaceOrder
		stage.WhiteSpace_orderStaged[stage.WhiteSpaceOrder] = whitespace
		stage.WhiteSpaceOrder++
	}
	stage.WhiteSpaces_mapString[whitespace.Name] = whitespace

	return whitespace
}

// StagePreserveOrder puts whitespace to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.WhiteSpaceOrder
// - update stage.WhiteSpaceOrder accordingly
func (whitespace *WhiteSpace) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.WhiteSpaces[whitespace]; !ok {
		stage.WhiteSpaces[whitespace] = struct{}{}

		if order > stage.WhiteSpaceOrder {
			stage.WhiteSpaceOrder = order
		}
		stage.WhiteSpace_stagedOrder[whitespace] = order
		stage.WhiteSpace_orderStaged[order] = whitespace
		stage.WhiteSpaceOrder++
	}
	stage.WhiteSpaces_mapString[whitespace.Name] = whitespace
}

// Unstage removes whitespace off the model stage
func (whitespace *WhiteSpace) Unstage(stage *Stage) *WhiteSpace {
	delete(stage.WhiteSpaces, whitespace)
	// issue1150
	// delete(stage.WhiteSpace_stagedOrder, whitespace)
	delete(stage.WhiteSpaces_mapString, whitespace.Name)

	return whitespace
}

// UnstageVoid removes whitespace off the model stage
func (whitespace *WhiteSpace) UnstageVoid(stage *Stage) {
	delete(stage.WhiteSpaces, whitespace)
	// issue1150
	// delete(stage.WhiteSpace_stagedOrder, whitespace)
	delete(stage.WhiteSpaces_mapString, whitespace.Name)
}

// commit whitespace to the back repo (if it is already staged)
func (whitespace *WhiteSpace) Commit(stage *Stage) *WhiteSpace {
	if _, ok := stage.WhiteSpaces[whitespace]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitWhiteSpace(whitespace)
		}
	}
	return whitespace
}

func (whitespace *WhiteSpace) CommitVoid(stage *Stage) {
	whitespace.Commit(stage)
}

func (whitespace *WhiteSpace) StageVoid(stage *Stage) {
	whitespace.Stage(stage)
}

// Checkout whitespace to the back repo (if it is already staged)
func (whitespace *WhiteSpace) Checkout(stage *Stage) *WhiteSpace {
	if _, ok := stage.WhiteSpaces[whitespace]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutWhiteSpace(whitespace)
		}
	}
	return whitespace
}

// for satisfaction of GongStruct interface
func (whitespace *WhiteSpace) GetName() (res string) {
	return whitespace.Name
}

// for satisfaction of GongStruct interface
func (whitespace *WhiteSpace) SetName(name string) {
	whitespace.Name = name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMAll(All *All)
	CreateORMAnnotation(Annotation *Annotation)
	CreateORMAttribute(Attribute *Attribute)
	CreateORMAttributeGroup(AttributeGroup *AttributeGroup)
	CreateORMChoice(Choice *Choice)
	CreateORMComplexContent(ComplexContent *ComplexContent)
	CreateORMComplexType(ComplexType *ComplexType)
	CreateORMDocumentation(Documentation *Documentation)
	CreateORMElement(Element *Element)
	CreateORMEnumeration(Enumeration *Enumeration)
	CreateORMExtension(Extension *Extension)
	CreateORMGroup(Group *Group)
	CreateORMLength(Length *Length)
	CreateORMMaxInclusive(MaxInclusive *MaxInclusive)
	CreateORMMaxLength(MaxLength *MaxLength)
	CreateORMMinInclusive(MinInclusive *MinInclusive)
	CreateORMMinLength(MinLength *MinLength)
	CreateORMPattern(Pattern *Pattern)
	CreateORMRestriction(Restriction *Restriction)
	CreateORMSchema(Schema *Schema)
	CreateORMSequence(Sequence *Sequence)
	CreateORMSimpleContent(SimpleContent *SimpleContent)
	CreateORMSimpleType(SimpleType *SimpleType)
	CreateORMTotalDigit(TotalDigit *TotalDigit)
	CreateORMUnion(Union *Union)
	CreateORMWhiteSpace(WhiteSpace *WhiteSpace)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMAll(All *All)
	DeleteORMAnnotation(Annotation *Annotation)
	DeleteORMAttribute(Attribute *Attribute)
	DeleteORMAttributeGroup(AttributeGroup *AttributeGroup)
	DeleteORMChoice(Choice *Choice)
	DeleteORMComplexContent(ComplexContent *ComplexContent)
	DeleteORMComplexType(ComplexType *ComplexType)
	DeleteORMDocumentation(Documentation *Documentation)
	DeleteORMElement(Element *Element)
	DeleteORMEnumeration(Enumeration *Enumeration)
	DeleteORMExtension(Extension *Extension)
	DeleteORMGroup(Group *Group)
	DeleteORMLength(Length *Length)
	DeleteORMMaxInclusive(MaxInclusive *MaxInclusive)
	DeleteORMMaxLength(MaxLength *MaxLength)
	DeleteORMMinInclusive(MinInclusive *MinInclusive)
	DeleteORMMinLength(MinLength *MinLength)
	DeleteORMPattern(Pattern *Pattern)
	DeleteORMRestriction(Restriction *Restriction)
	DeleteORMSchema(Schema *Schema)
	DeleteORMSequence(Sequence *Sequence)
	DeleteORMSimpleContent(SimpleContent *SimpleContent)
	DeleteORMSimpleType(SimpleType *SimpleType)
	DeleteORMTotalDigit(TotalDigit *TotalDigit)
	DeleteORMUnion(Union *Union)
	DeleteORMWhiteSpace(WhiteSpace *WhiteSpace)
}

func (stage *Stage) Reset() { // insertion point for array reset
	stage.Alls = make(map[*All]struct{})
	stage.Alls_mapString = make(map[string]*All)
	stage.All_stagedOrder = make(map[*All]uint)
	stage.AllOrder = 0

	stage.Annotations = make(map[*Annotation]struct{})
	stage.Annotations_mapString = make(map[string]*Annotation)
	stage.Annotation_stagedOrder = make(map[*Annotation]uint)
	stage.AnnotationOrder = 0

	stage.Attributes = make(map[*Attribute]struct{})
	stage.Attributes_mapString = make(map[string]*Attribute)
	stage.Attribute_stagedOrder = make(map[*Attribute]uint)
	stage.AttributeOrder = 0

	stage.AttributeGroups = make(map[*AttributeGroup]struct{})
	stage.AttributeGroups_mapString = make(map[string]*AttributeGroup)
	stage.AttributeGroup_stagedOrder = make(map[*AttributeGroup]uint)
	stage.AttributeGroupOrder = 0

	stage.Choices = make(map[*Choice]struct{})
	stage.Choices_mapString = make(map[string]*Choice)
	stage.Choice_stagedOrder = make(map[*Choice]uint)
	stage.ChoiceOrder = 0

	stage.ComplexContents = make(map[*ComplexContent]struct{})
	stage.ComplexContents_mapString = make(map[string]*ComplexContent)
	stage.ComplexContent_stagedOrder = make(map[*ComplexContent]uint)
	stage.ComplexContentOrder = 0

	stage.ComplexTypes = make(map[*ComplexType]struct{})
	stage.ComplexTypes_mapString = make(map[string]*ComplexType)
	stage.ComplexType_stagedOrder = make(map[*ComplexType]uint)
	stage.ComplexTypeOrder = 0

	stage.Documentations = make(map[*Documentation]struct{})
	stage.Documentations_mapString = make(map[string]*Documentation)
	stage.Documentation_stagedOrder = make(map[*Documentation]uint)
	stage.DocumentationOrder = 0

	stage.Elements = make(map[*Element]struct{})
	stage.Elements_mapString = make(map[string]*Element)
	stage.Element_stagedOrder = make(map[*Element]uint)
	stage.ElementOrder = 0

	stage.Enumerations = make(map[*Enumeration]struct{})
	stage.Enumerations_mapString = make(map[string]*Enumeration)
	stage.Enumeration_stagedOrder = make(map[*Enumeration]uint)
	stage.EnumerationOrder = 0

	stage.Extensions = make(map[*Extension]struct{})
	stage.Extensions_mapString = make(map[string]*Extension)
	stage.Extension_stagedOrder = make(map[*Extension]uint)
	stage.ExtensionOrder = 0

	stage.Groups = make(map[*Group]struct{})
	stage.Groups_mapString = make(map[string]*Group)
	stage.Group_stagedOrder = make(map[*Group]uint)
	stage.GroupOrder = 0

	stage.Lengths = make(map[*Length]struct{})
	stage.Lengths_mapString = make(map[string]*Length)
	stage.Length_stagedOrder = make(map[*Length]uint)
	stage.LengthOrder = 0

	stage.MaxInclusives = make(map[*MaxInclusive]struct{})
	stage.MaxInclusives_mapString = make(map[string]*MaxInclusive)
	stage.MaxInclusive_stagedOrder = make(map[*MaxInclusive]uint)
	stage.MaxInclusiveOrder = 0

	stage.MaxLengths = make(map[*MaxLength]struct{})
	stage.MaxLengths_mapString = make(map[string]*MaxLength)
	stage.MaxLength_stagedOrder = make(map[*MaxLength]uint)
	stage.MaxLengthOrder = 0

	stage.MinInclusives = make(map[*MinInclusive]struct{})
	stage.MinInclusives_mapString = make(map[string]*MinInclusive)
	stage.MinInclusive_stagedOrder = make(map[*MinInclusive]uint)
	stage.MinInclusiveOrder = 0

	stage.MinLengths = make(map[*MinLength]struct{})
	stage.MinLengths_mapString = make(map[string]*MinLength)
	stage.MinLength_stagedOrder = make(map[*MinLength]uint)
	stage.MinLengthOrder = 0

	stage.Patterns = make(map[*Pattern]struct{})
	stage.Patterns_mapString = make(map[string]*Pattern)
	stage.Pattern_stagedOrder = make(map[*Pattern]uint)
	stage.PatternOrder = 0

	stage.Restrictions = make(map[*Restriction]struct{})
	stage.Restrictions_mapString = make(map[string]*Restriction)
	stage.Restriction_stagedOrder = make(map[*Restriction]uint)
	stage.RestrictionOrder = 0

	stage.Schemas = make(map[*Schema]struct{})
	stage.Schemas_mapString = make(map[string]*Schema)
	stage.Schema_stagedOrder = make(map[*Schema]uint)
	stage.SchemaOrder = 0

	stage.Sequences = make(map[*Sequence]struct{})
	stage.Sequences_mapString = make(map[string]*Sequence)
	stage.Sequence_stagedOrder = make(map[*Sequence]uint)
	stage.SequenceOrder = 0

	stage.SimpleContents = make(map[*SimpleContent]struct{})
	stage.SimpleContents_mapString = make(map[string]*SimpleContent)
	stage.SimpleContent_stagedOrder = make(map[*SimpleContent]uint)
	stage.SimpleContentOrder = 0

	stage.SimpleTypes = make(map[*SimpleType]struct{})
	stage.SimpleTypes_mapString = make(map[string]*SimpleType)
	stage.SimpleType_stagedOrder = make(map[*SimpleType]uint)
	stage.SimpleTypeOrder = 0

	stage.TotalDigits = make(map[*TotalDigit]struct{})
	stage.TotalDigits_mapString = make(map[string]*TotalDigit)
	stage.TotalDigit_stagedOrder = make(map[*TotalDigit]uint)
	stage.TotalDigitOrder = 0

	stage.Unions = make(map[*Union]struct{})
	stage.Unions_mapString = make(map[string]*Union)
	stage.Union_stagedOrder = make(map[*Union]uint)
	stage.UnionOrder = 0

	stage.WhiteSpaces = make(map[*WhiteSpace]struct{})
	stage.WhiteSpaces_mapString = make(map[string]*WhiteSpace)
	stage.WhiteSpace_stagedOrder = make(map[*WhiteSpace]uint)
	stage.WhiteSpaceOrder = 0

	if stage.GetProbeIF() != nil {
		stage.GetProbeIF().ResetNotifications()
	}
	if stage.IsInDeltaMode() {
		stage.ComputeReferenceAndOrders()
	}
}

func (stage *Stage) Nil() { // insertion point for array nil
	stage.Alls = nil
	stage.Alls_mapString = nil

	stage.Annotations = nil
	stage.Annotations_mapString = nil

	stage.Attributes = nil
	stage.Attributes_mapString = nil

	stage.AttributeGroups = nil
	stage.AttributeGroups_mapString = nil

	stage.Choices = nil
	stage.Choices_mapString = nil

	stage.ComplexContents = nil
	stage.ComplexContents_mapString = nil

	stage.ComplexTypes = nil
	stage.ComplexTypes_mapString = nil

	stage.Documentations = nil
	stage.Documentations_mapString = nil

	stage.Elements = nil
	stage.Elements_mapString = nil

	stage.Enumerations = nil
	stage.Enumerations_mapString = nil

	stage.Extensions = nil
	stage.Extensions_mapString = nil

	stage.Groups = nil
	stage.Groups_mapString = nil

	stage.Lengths = nil
	stage.Lengths_mapString = nil

	stage.MaxInclusives = nil
	stage.MaxInclusives_mapString = nil

	stage.MaxLengths = nil
	stage.MaxLengths_mapString = nil

	stage.MinInclusives = nil
	stage.MinInclusives_mapString = nil

	stage.MinLengths = nil
	stage.MinLengths_mapString = nil

	stage.Patterns = nil
	stage.Patterns_mapString = nil

	stage.Restrictions = nil
	stage.Restrictions_mapString = nil

	stage.Schemas = nil
	stage.Schemas_mapString = nil

	stage.Sequences = nil
	stage.Sequences_mapString = nil

	stage.SimpleContents = nil
	stage.SimpleContents_mapString = nil

	stage.SimpleTypes = nil
	stage.SimpleTypes_mapString = nil

	stage.TotalDigits = nil
	stage.TotalDigits_mapString = nil

	stage.Unions = nil
	stage.Unions_mapString = nil

	stage.WhiteSpaces = nil
	stage.WhiteSpaces_mapString = nil

	// end of insertion point for array nil
}

func (stage *Stage) Unstage() { // insertion point for array nil
	for all := range stage.Alls {
		all.Unstage(stage)
	}

	for annotation := range stage.Annotations {
		annotation.Unstage(stage)
	}

	for attribute := range stage.Attributes {
		attribute.Unstage(stage)
	}

	for attributegroup := range stage.AttributeGroups {
		attributegroup.Unstage(stage)
	}

	for choice := range stage.Choices {
		choice.Unstage(stage)
	}

	for complexcontent := range stage.ComplexContents {
		complexcontent.Unstage(stage)
	}

	for complextype := range stage.ComplexTypes {
		complextype.Unstage(stage)
	}

	for documentation := range stage.Documentations {
		documentation.Unstage(stage)
	}

	for element := range stage.Elements {
		element.Unstage(stage)
	}

	for enumeration := range stage.Enumerations {
		enumeration.Unstage(stage)
	}

	for extension := range stage.Extensions {
		extension.Unstage(stage)
	}

	for group := range stage.Groups {
		group.Unstage(stage)
	}

	for length := range stage.Lengths {
		length.Unstage(stage)
	}

	for maxinclusive := range stage.MaxInclusives {
		maxinclusive.Unstage(stage)
	}

	for maxlength := range stage.MaxLengths {
		maxlength.Unstage(stage)
	}

	for mininclusive := range stage.MinInclusives {
		mininclusive.Unstage(stage)
	}

	for minlength := range stage.MinLengths {
		minlength.Unstage(stage)
	}

	for pattern := range stage.Patterns {
		pattern.Unstage(stage)
	}

	for restriction := range stage.Restrictions {
		restriction.Unstage(stage)
	}

	for schema := range stage.Schemas {
		schema.Unstage(stage)
	}

	for sequence := range stage.Sequences {
		sequence.Unstage(stage)
	}

	for simplecontent := range stage.SimpleContents {
		simplecontent.Unstage(stage)
	}

	for simpletype := range stage.SimpleTypes {
		simpletype.Unstage(stage)
	}

	for totaldigit := range stage.TotalDigits {
		totaldigit.Unstage(stage)
	}

	for union := range stage.Unions {
		union.Unstage(stage)
	}

	for whitespace := range stage.WhiteSpaces {
		whitespace.Unstage(stage)
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
	case map[*All]any:
		return any(&stage.Alls).(*Type)
	case map[*Annotation]any:
		return any(&stage.Annotations).(*Type)
	case map[*Attribute]any:
		return any(&stage.Attributes).(*Type)
	case map[*AttributeGroup]any:
		return any(&stage.AttributeGroups).(*Type)
	case map[*Choice]any:
		return any(&stage.Choices).(*Type)
	case map[*ComplexContent]any:
		return any(&stage.ComplexContents).(*Type)
	case map[*ComplexType]any:
		return any(&stage.ComplexTypes).(*Type)
	case map[*Documentation]any:
		return any(&stage.Documentations).(*Type)
	case map[*Element]any:
		return any(&stage.Elements).(*Type)
	case map[*Enumeration]any:
		return any(&stage.Enumerations).(*Type)
	case map[*Extension]any:
		return any(&stage.Extensions).(*Type)
	case map[*Group]any:
		return any(&stage.Groups).(*Type)
	case map[*Length]any:
		return any(&stage.Lengths).(*Type)
	case map[*MaxInclusive]any:
		return any(&stage.MaxInclusives).(*Type)
	case map[*MaxLength]any:
		return any(&stage.MaxLengths).(*Type)
	case map[*MinInclusive]any:
		return any(&stage.MinInclusives).(*Type)
	case map[*MinLength]any:
		return any(&stage.MinLengths).(*Type)
	case map[*Pattern]any:
		return any(&stage.Patterns).(*Type)
	case map[*Restriction]any:
		return any(&stage.Restrictions).(*Type)
	case map[*Schema]any:
		return any(&stage.Schemas).(*Type)
	case map[*Sequence]any:
		return any(&stage.Sequences).(*Type)
	case map[*SimpleContent]any:
		return any(&stage.SimpleContents).(*Type)
	case map[*SimpleType]any:
		return any(&stage.SimpleTypes).(*Type)
	case map[*TotalDigit]any:
		return any(&stage.TotalDigits).(*Type)
	case map[*Union]any:
		return any(&stage.Unions).(*Type)
	case map[*WhiteSpace]any:
		return any(&stage.WhiteSpaces).(*Type)
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
	case *All:
		return any(stage.Alls_mapString).(map[string]Type)
	case *Annotation:
		return any(stage.Annotations_mapString).(map[string]Type)
	case *Attribute:
		return any(stage.Attributes_mapString).(map[string]Type)
	case *AttributeGroup:
		return any(stage.AttributeGroups_mapString).(map[string]Type)
	case *Choice:
		return any(stage.Choices_mapString).(map[string]Type)
	case *ComplexContent:
		return any(stage.ComplexContents_mapString).(map[string]Type)
	case *ComplexType:
		return any(stage.ComplexTypes_mapString).(map[string]Type)
	case *Documentation:
		return any(stage.Documentations_mapString).(map[string]Type)
	case *Element:
		return any(stage.Elements_mapString).(map[string]Type)
	case *Enumeration:
		return any(stage.Enumerations_mapString).(map[string]Type)
	case *Extension:
		return any(stage.Extensions_mapString).(map[string]Type)
	case *Group:
		return any(stage.Groups_mapString).(map[string]Type)
	case *Length:
		return any(stage.Lengths_mapString).(map[string]Type)
	case *MaxInclusive:
		return any(stage.MaxInclusives_mapString).(map[string]Type)
	case *MaxLength:
		return any(stage.MaxLengths_mapString).(map[string]Type)
	case *MinInclusive:
		return any(stage.MinInclusives_mapString).(map[string]Type)
	case *MinLength:
		return any(stage.MinLengths_mapString).(map[string]Type)
	case *Pattern:
		return any(stage.Patterns_mapString).(map[string]Type)
	case *Restriction:
		return any(stage.Restrictions_mapString).(map[string]Type)
	case *Schema:
		return any(stage.Schemas_mapString).(map[string]Type)
	case *Sequence:
		return any(stage.Sequences_mapString).(map[string]Type)
	case *SimpleContent:
		return any(stage.SimpleContents_mapString).(map[string]Type)
	case *SimpleType:
		return any(stage.SimpleTypes_mapString).(map[string]Type)
	case *TotalDigit:
		return any(stage.TotalDigits_mapString).(map[string]Type)
	case *Union:
		return any(stage.Unions_mapString).(map[string]Type)
	case *WhiteSpace:
		return any(stage.WhiteSpaces_mapString).(map[string]Type)
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
	case All:
		return any(&stage.Alls).(*map[*Type]struct{})
	case Annotation:
		return any(&stage.Annotations).(*map[*Type]struct{})
	case Attribute:
		return any(&stage.Attributes).(*map[*Type]struct{})
	case AttributeGroup:
		return any(&stage.AttributeGroups).(*map[*Type]struct{})
	case Choice:
		return any(&stage.Choices).(*map[*Type]struct{})
	case ComplexContent:
		return any(&stage.ComplexContents).(*map[*Type]struct{})
	case ComplexType:
		return any(&stage.ComplexTypes).(*map[*Type]struct{})
	case Documentation:
		return any(&stage.Documentations).(*map[*Type]struct{})
	case Element:
		return any(&stage.Elements).(*map[*Type]struct{})
	case Enumeration:
		return any(&stage.Enumerations).(*map[*Type]struct{})
	case Extension:
		return any(&stage.Extensions).(*map[*Type]struct{})
	case Group:
		return any(&stage.Groups).(*map[*Type]struct{})
	case Length:
		return any(&stage.Lengths).(*map[*Type]struct{})
	case MaxInclusive:
		return any(&stage.MaxInclusives).(*map[*Type]struct{})
	case MaxLength:
		return any(&stage.MaxLengths).(*map[*Type]struct{})
	case MinInclusive:
		return any(&stage.MinInclusives).(*map[*Type]struct{})
	case MinLength:
		return any(&stage.MinLengths).(*map[*Type]struct{})
	case Pattern:
		return any(&stage.Patterns).(*map[*Type]struct{})
	case Restriction:
		return any(&stage.Restrictions).(*map[*Type]struct{})
	case Schema:
		return any(&stage.Schemas).(*map[*Type]struct{})
	case Sequence:
		return any(&stage.Sequences).(*map[*Type]struct{})
	case SimpleContent:
		return any(&stage.SimpleContents).(*map[*Type]struct{})
	case SimpleType:
		return any(&stage.SimpleTypes).(*map[*Type]struct{})
	case TotalDigit:
		return any(&stage.TotalDigits).(*map[*Type]struct{})
	case Union:
		return any(&stage.Unions).(*map[*Type]struct{})
	case WhiteSpace:
		return any(&stage.WhiteSpaces).(*map[*Type]struct{})
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
	case *All:
		return any(&stage.Alls).(*map[Type]struct{})
	case *Annotation:
		return any(&stage.Annotations).(*map[Type]struct{})
	case *Attribute:
		return any(&stage.Attributes).(*map[Type]struct{})
	case *AttributeGroup:
		return any(&stage.AttributeGroups).(*map[Type]struct{})
	case *Choice:
		return any(&stage.Choices).(*map[Type]struct{})
	case *ComplexContent:
		return any(&stage.ComplexContents).(*map[Type]struct{})
	case *ComplexType:
		return any(&stage.ComplexTypes).(*map[Type]struct{})
	case *Documentation:
		return any(&stage.Documentations).(*map[Type]struct{})
	case *Element:
		return any(&stage.Elements).(*map[Type]struct{})
	case *Enumeration:
		return any(&stage.Enumerations).(*map[Type]struct{})
	case *Extension:
		return any(&stage.Extensions).(*map[Type]struct{})
	case *Group:
		return any(&stage.Groups).(*map[Type]struct{})
	case *Length:
		return any(&stage.Lengths).(*map[Type]struct{})
	case *MaxInclusive:
		return any(&stage.MaxInclusives).(*map[Type]struct{})
	case *MaxLength:
		return any(&stage.MaxLengths).(*map[Type]struct{})
	case *MinInclusive:
		return any(&stage.MinInclusives).(*map[Type]struct{})
	case *MinLength:
		return any(&stage.MinLengths).(*map[Type]struct{})
	case *Pattern:
		return any(&stage.Patterns).(*map[Type]struct{})
	case *Restriction:
		return any(&stage.Restrictions).(*map[Type]struct{})
	case *Schema:
		return any(&stage.Schemas).(*map[Type]struct{})
	case *Sequence:
		return any(&stage.Sequences).(*map[Type]struct{})
	case *SimpleContent:
		return any(&stage.SimpleContents).(*map[Type]struct{})
	case *SimpleType:
		return any(&stage.SimpleTypes).(*map[Type]struct{})
	case *TotalDigit:
		return any(&stage.TotalDigits).(*map[Type]struct{})
	case *Union:
		return any(&stage.Unions).(*map[Type]struct{})
	case *WhiteSpace:
		return any(&stage.WhiteSpaces).(*map[Type]struct{})
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
	case All:
		return any(&stage.Alls_mapString).(*map[string]*Type)
	case Annotation:
		return any(&stage.Annotations_mapString).(*map[string]*Type)
	case Attribute:
		return any(&stage.Attributes_mapString).(*map[string]*Type)
	case AttributeGroup:
		return any(&stage.AttributeGroups_mapString).(*map[string]*Type)
	case Choice:
		return any(&stage.Choices_mapString).(*map[string]*Type)
	case ComplexContent:
		return any(&stage.ComplexContents_mapString).(*map[string]*Type)
	case ComplexType:
		return any(&stage.ComplexTypes_mapString).(*map[string]*Type)
	case Documentation:
		return any(&stage.Documentations_mapString).(*map[string]*Type)
	case Element:
		return any(&stage.Elements_mapString).(*map[string]*Type)
	case Enumeration:
		return any(&stage.Enumerations_mapString).(*map[string]*Type)
	case Extension:
		return any(&stage.Extensions_mapString).(*map[string]*Type)
	case Group:
		return any(&stage.Groups_mapString).(*map[string]*Type)
	case Length:
		return any(&stage.Lengths_mapString).(*map[string]*Type)
	case MaxInclusive:
		return any(&stage.MaxInclusives_mapString).(*map[string]*Type)
	case MaxLength:
		return any(&stage.MaxLengths_mapString).(*map[string]*Type)
	case MinInclusive:
		return any(&stage.MinInclusives_mapString).(*map[string]*Type)
	case MinLength:
		return any(&stage.MinLengths_mapString).(*map[string]*Type)
	case Pattern:
		return any(&stage.Patterns_mapString).(*map[string]*Type)
	case Restriction:
		return any(&stage.Restrictions_mapString).(*map[string]*Type)
	case Schema:
		return any(&stage.Schemas_mapString).(*map[string]*Type)
	case Sequence:
		return any(&stage.Sequences_mapString).(*map[string]*Type)
	case SimpleContent:
		return any(&stage.SimpleContents_mapString).(*map[string]*Type)
	case SimpleType:
		return any(&stage.SimpleTypes_mapString).(*map[string]*Type)
	case TotalDigit:
		return any(&stage.TotalDigits_mapString).(*map[string]*Type)
	case Union:
		return any(&stage.Unions_mapString).(*map[string]*Type)
	case WhiteSpace:
		return any(&stage.WhiteSpaces_mapString).(*map[string]*Type)
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
	case All:
		return any(&All{
			// Initialisation of associations
			// field is initialized with Annotated problem with composites

		}).(*Type)
	case Annotation:
		return any(&Annotation{
			// Initialisation of associations
			// field is initialized with an instance of Documentation with the name of the field
			Documentations: []*Documentation{{Name: "Documentations"}},
		}).(*Type)
	case Attribute:
		return any(&Attribute{
			// Initialisation of associations
			// field is initialized with Annotated problem with composites

		}).(*Type)
	case AttributeGroup:
		return any(&AttributeGroup{
			// Initialisation of associations
			// field is initialized with an instance of AttributeGroup with the name of the field
			AttributeGroups: []*AttributeGroup{{Name: "AttributeGroups"}},
			// field is initialized with an instance of Attribute with the name of the field
			Attributes: []*Attribute{{Name: "Attributes"}},
			// field is initialized with Annotated problem with composites

		}).(*Type)
	case Choice:
		return any(&Choice{
			// Initialisation of associations
			// field is initialized with Annotated problem with composites

		}).(*Type)
	case ComplexContent:
		return any(&ComplexContent{
			// Initialisation of associations
		}).(*Type)
	case ComplexType:
		return any(&ComplexType{
			// Initialisation of associations
			// field is initialized with an instance of Element with the name of the field
			OuterElement: &Element{Name: "OuterElement"},
			// field is initialized with an instance of Extension with the name of the field
			Extension: &Extension{Name: "Extension"},
			// field is initialized with an instance of SimpleContent with the name of the field
			SimpleContent: &SimpleContent{Name: "SimpleContent"},
			// field is initialized with an instance of ComplexContent with the name of the field
			ComplexContent: &ComplexContent{Name: "ComplexContent"},
			// field is initialized with an instance of Attribute with the name of the field
			Attributes: []*Attribute{{Name: "Attributes"}},
			// field is initialized with an instance of AttributeGroup with the name of the field
			AttributeGroups: []*AttributeGroup{{Name: "AttributeGroups"}},
			// field is initialized with Annotated problem with composites

		}).(*Type)
	case Documentation:
		return any(&Documentation{
			// Initialisation of associations
		}).(*Type)
	case Element:
		return any(&Element{
			// Initialisation of associations
			// field is initialized with an instance of SimpleType with the name of the field
			SimpleType: &SimpleType{Name: "SimpleType"},
			// field is initialized with an instance of ComplexType with the name of the field
			ComplexType: &ComplexType{Name: "ComplexType"},
			// field is initialized with an instance of Group with the name of the field
			Groups: []*Group{{Name: "Groups"}},
			// field is initialized with Annotated problem with composites

		}).(*Type)
	case Enumeration:
		return any(&Enumeration{
			// Initialisation of associations
			// field is initialized with Annotated problem with composites

		}).(*Type)
	case Extension:
		return any(&Extension{
			// Initialisation of associations
			// field is initialized with an instance of Attribute with the name of the field
			Attributes: []*Attribute{{Name: "Attributes"}},
			// field is initialized with an instance of AttributeGroup with the name of the field
			AttributeGroups: []*AttributeGroup{{Name: "AttributeGroups"}},
		}).(*Type)
	case Group:
		return any(&Group{
			// Initialisation of associations
			// field is initialized with an instance of Element with the name of the field
			OuterElement: &Element{Name: "OuterElement"},
			// field is initialized with Annotated problem with composites

		}).(*Type)
	case Length:
		return any(&Length{
			// Initialisation of associations
			// field is initialized with Annotated problem with composites

		}).(*Type)
	case MaxInclusive:
		return any(&MaxInclusive{
			// Initialisation of associations
			// field is initialized with Annotated problem with composites

		}).(*Type)
	case MaxLength:
		return any(&MaxLength{
			// Initialisation of associations
			// field is initialized with Annotated problem with composites

		}).(*Type)
	case MinInclusive:
		return any(&MinInclusive{
			// Initialisation of associations
			// field is initialized with Annotated problem with composites

		}).(*Type)
	case MinLength:
		return any(&MinLength{
			// Initialisation of associations
			// field is initialized with Annotated problem with composites

		}).(*Type)
	case Pattern:
		return any(&Pattern{
			// Initialisation of associations
			// field is initialized with Annotated problem with composites

		}).(*Type)
	case Restriction:
		return any(&Restriction{
			// Initialisation of associations
			// field is initialized with an instance of Enumeration with the name of the field
			Enumerations: []*Enumeration{{Name: "Enumerations"}},
			// field is initialized with an instance of MinInclusive with the name of the field
			MinInclusive: &MinInclusive{Name: "MinInclusive"},
			// field is initialized with an instance of MaxInclusive with the name of the field
			MaxInclusive: &MaxInclusive{Name: "MaxInclusive"},
			// field is initialized with an instance of Pattern with the name of the field
			Pattern: &Pattern{Name: "Pattern"},
			// field is initialized with an instance of WhiteSpace with the name of the field
			WhiteSpace: &WhiteSpace{Name: "WhiteSpace"},
			// field is initialized with an instance of MinLength with the name of the field
			MinLength: &MinLength{Name: "MinLength"},
			// field is initialized with an instance of MaxLength with the name of the field
			MaxLength: &MaxLength{Name: "MaxLength"},
			// field is initialized with an instance of Length with the name of the field
			Length: &Length{Name: "Length"},
			// field is initialized with an instance of TotalDigit with the name of the field
			TotalDigit: &TotalDigit{Name: "TotalDigit"},
			// field is initialized with Annotated problem with composites

		}).(*Type)
	case Schema:
		return any(&Schema{
			// Initialisation of associations
			// field is initialized with an instance of Element with the name of the field
			Elements: []*Element{{Name: "Elements"}},
			// field is initialized with an instance of SimpleType with the name of the field
			SimpleTypes: []*SimpleType{{Name: "SimpleTypes"}},
			// field is initialized with an instance of ComplexType with the name of the field
			ComplexTypes: []*ComplexType{{Name: "ComplexTypes"}},
			// field is initialized with an instance of AttributeGroup with the name of the field
			AttributeGroups: []*AttributeGroup{{Name: "AttributeGroups"}},
			// field is initialized with an instance of Group with the name of the field
			Groups: []*Group{{Name: "Groups"}},
			// field is initialized with Annotated problem with composites

		}).(*Type)
	case Sequence:
		return any(&Sequence{
			// Initialisation of associations
			// field is initialized with Annotated problem with composites

		}).(*Type)
	case SimpleContent:
		return any(&SimpleContent{
			// Initialisation of associations
			// field is initialized with an instance of Extension with the name of the field
			Extension: &Extension{Name: "Extension"},
			// field is initialized with an instance of Restriction with the name of the field
			Restriction: &Restriction{Name: "Restriction"},
		}).(*Type)
	case SimpleType:
		return any(&SimpleType{
			// Initialisation of associations
			// field is initialized with an instance of Restriction with the name of the field
			Restriction: &Restriction{Name: "Restriction"},
			// field is initialized with an instance of Union with the name of the field
			Union: &Union{Name: "Union"},
			// field is initialized with Annotated problem with composites

		}).(*Type)
	case TotalDigit:
		return any(&TotalDigit{
			// Initialisation of associations
			// field is initialized with Annotated problem with composites

		}).(*Type)
	case Union:
		return any(&Union{
			// Initialisation of associations
			// field is initialized with Annotated problem with composites

		}).(*Type)
	case WhiteSpace:
		return any(&WhiteSpace{
			// Initialisation of associations
			// field is initialized with Annotated problem with composites

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
	// reverse maps of direct associations of All
	case All:
		switch fieldname {
		// insertion point for per direct association field
		case "Annotation":
			res := make(map[*Annotation][]*All)
			for all := range stage.Alls {
				if all.Annotation != nil {
					annotation_ := all.Annotation
					var alls []*All
					_, ok := res[annotation_]
					if ok {
						alls = res[annotation_]
					} else {
						alls = make([]*All, 0)
					}
					alls = append(alls, all)
					res[annotation_] = alls
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Annotation
	case Annotation:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Attribute
	case Attribute:
		switch fieldname {
		// insertion point for per direct association field
		case "Annotation":
			res := make(map[*Annotation][]*Attribute)
			for attribute := range stage.Attributes {
				if attribute.Annotation != nil {
					annotation_ := attribute.Annotation
					var attributes []*Attribute
					_, ok := res[annotation_]
					if ok {
						attributes = res[annotation_]
					} else {
						attributes = make([]*Attribute, 0)
					}
					attributes = append(attributes, attribute)
					res[annotation_] = attributes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of AttributeGroup
	case AttributeGroup:
		switch fieldname {
		// insertion point for per direct association field
		case "Annotation":
			res := make(map[*Annotation][]*AttributeGroup)
			for attributegroup := range stage.AttributeGroups {
				if attributegroup.Annotation != nil {
					annotation_ := attributegroup.Annotation
					var attributegroups []*AttributeGroup
					_, ok := res[annotation_]
					if ok {
						attributegroups = res[annotation_]
					} else {
						attributegroups = make([]*AttributeGroup, 0)
					}
					attributegroups = append(attributegroups, attributegroup)
					res[annotation_] = attributegroups
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Choice
	case Choice:
		switch fieldname {
		// insertion point for per direct association field
		case "Annotation":
			res := make(map[*Annotation][]*Choice)
			for choice := range stage.Choices {
				if choice.Annotation != nil {
					annotation_ := choice.Annotation
					var choices []*Choice
					_, ok := res[annotation_]
					if ok {
						choices = res[annotation_]
					} else {
						choices = make([]*Choice, 0)
					}
					choices = append(choices, choice)
					res[annotation_] = choices
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ComplexContent
	case ComplexContent:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ComplexType
	case ComplexType:
		switch fieldname {
		// insertion point for per direct association field
		case "OuterElement":
			res := make(map[*Element][]*ComplexType)
			for complextype := range stage.ComplexTypes {
				if complextype.OuterElement != nil {
					element_ := complextype.OuterElement
					var complextypes []*ComplexType
					_, ok := res[element_]
					if ok {
						complextypes = res[element_]
					} else {
						complextypes = make([]*ComplexType, 0)
					}
					complextypes = append(complextypes, complextype)
					res[element_] = complextypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "Annotation":
			res := make(map[*Annotation][]*ComplexType)
			for complextype := range stage.ComplexTypes {
				if complextype.Annotation != nil {
					annotation_ := complextype.Annotation
					var complextypes []*ComplexType
					_, ok := res[annotation_]
					if ok {
						complextypes = res[annotation_]
					} else {
						complextypes = make([]*ComplexType, 0)
					}
					complextypes = append(complextypes, complextype)
					res[annotation_] = complextypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "Extension":
			res := make(map[*Extension][]*ComplexType)
			for complextype := range stage.ComplexTypes {
				if complextype.Extension != nil {
					extension_ := complextype.Extension
					var complextypes []*ComplexType
					_, ok := res[extension_]
					if ok {
						complextypes = res[extension_]
					} else {
						complextypes = make([]*ComplexType, 0)
					}
					complextypes = append(complextypes, complextype)
					res[extension_] = complextypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "SimpleContent":
			res := make(map[*SimpleContent][]*ComplexType)
			for complextype := range stage.ComplexTypes {
				if complextype.SimpleContent != nil {
					simplecontent_ := complextype.SimpleContent
					var complextypes []*ComplexType
					_, ok := res[simplecontent_]
					if ok {
						complextypes = res[simplecontent_]
					} else {
						complextypes = make([]*ComplexType, 0)
					}
					complextypes = append(complextypes, complextype)
					res[simplecontent_] = complextypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "ComplexContent":
			res := make(map[*ComplexContent][]*ComplexType)
			for complextype := range stage.ComplexTypes {
				if complextype.ComplexContent != nil {
					complexcontent_ := complextype.ComplexContent
					var complextypes []*ComplexType
					_, ok := res[complexcontent_]
					if ok {
						complextypes = res[complexcontent_]
					} else {
						complextypes = make([]*ComplexType, 0)
					}
					complextypes = append(complextypes, complextype)
					res[complexcontent_] = complextypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Documentation
	case Documentation:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Element
	case Element:
		switch fieldname {
		// insertion point for per direct association field
		case "Annotation":
			res := make(map[*Annotation][]*Element)
			for element := range stage.Elements {
				if element.Annotation != nil {
					annotation_ := element.Annotation
					var elements []*Element
					_, ok := res[annotation_]
					if ok {
						elements = res[annotation_]
					} else {
						elements = make([]*Element, 0)
					}
					elements = append(elements, element)
					res[annotation_] = elements
				}
			}
			return any(res).(map[*End][]*Start)
		case "SimpleType":
			res := make(map[*SimpleType][]*Element)
			for element := range stage.Elements {
				if element.SimpleType != nil {
					simpletype_ := element.SimpleType
					var elements []*Element
					_, ok := res[simpletype_]
					if ok {
						elements = res[simpletype_]
					} else {
						elements = make([]*Element, 0)
					}
					elements = append(elements, element)
					res[simpletype_] = elements
				}
			}
			return any(res).(map[*End][]*Start)
		case "ComplexType":
			res := make(map[*ComplexType][]*Element)
			for element := range stage.Elements {
				if element.ComplexType != nil {
					complextype_ := element.ComplexType
					var elements []*Element
					_, ok := res[complextype_]
					if ok {
						elements = res[complextype_]
					} else {
						elements = make([]*Element, 0)
					}
					elements = append(elements, element)
					res[complextype_] = elements
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Enumeration
	case Enumeration:
		switch fieldname {
		// insertion point for per direct association field
		case "Annotation":
			res := make(map[*Annotation][]*Enumeration)
			for enumeration := range stage.Enumerations {
				if enumeration.Annotation != nil {
					annotation_ := enumeration.Annotation
					var enumerations []*Enumeration
					_, ok := res[annotation_]
					if ok {
						enumerations = res[annotation_]
					} else {
						enumerations = make([]*Enumeration, 0)
					}
					enumerations = append(enumerations, enumeration)
					res[annotation_] = enumerations
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Extension
	case Extension:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Group
	case Group:
		switch fieldname {
		// insertion point for per direct association field
		case "Annotation":
			res := make(map[*Annotation][]*Group)
			for group := range stage.Groups {
				if group.Annotation != nil {
					annotation_ := group.Annotation
					var groups []*Group
					_, ok := res[annotation_]
					if ok {
						groups = res[annotation_]
					} else {
						groups = make([]*Group, 0)
					}
					groups = append(groups, group)
					res[annotation_] = groups
				}
			}
			return any(res).(map[*End][]*Start)
		case "OuterElement":
			res := make(map[*Element][]*Group)
			for group := range stage.Groups {
				if group.OuterElement != nil {
					element_ := group.OuterElement
					var groups []*Group
					_, ok := res[element_]
					if ok {
						groups = res[element_]
					} else {
						groups = make([]*Group, 0)
					}
					groups = append(groups, group)
					res[element_] = groups
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Length
	case Length:
		switch fieldname {
		// insertion point for per direct association field
		case "Annotation":
			res := make(map[*Annotation][]*Length)
			for length := range stage.Lengths {
				if length.Annotation != nil {
					annotation_ := length.Annotation
					var lengths []*Length
					_, ok := res[annotation_]
					if ok {
						lengths = res[annotation_]
					} else {
						lengths = make([]*Length, 0)
					}
					lengths = append(lengths, length)
					res[annotation_] = lengths
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of MaxInclusive
	case MaxInclusive:
		switch fieldname {
		// insertion point for per direct association field
		case "Annotation":
			res := make(map[*Annotation][]*MaxInclusive)
			for maxinclusive := range stage.MaxInclusives {
				if maxinclusive.Annotation != nil {
					annotation_ := maxinclusive.Annotation
					var maxinclusives []*MaxInclusive
					_, ok := res[annotation_]
					if ok {
						maxinclusives = res[annotation_]
					} else {
						maxinclusives = make([]*MaxInclusive, 0)
					}
					maxinclusives = append(maxinclusives, maxinclusive)
					res[annotation_] = maxinclusives
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of MaxLength
	case MaxLength:
		switch fieldname {
		// insertion point for per direct association field
		case "Annotation":
			res := make(map[*Annotation][]*MaxLength)
			for maxlength := range stage.MaxLengths {
				if maxlength.Annotation != nil {
					annotation_ := maxlength.Annotation
					var maxlengths []*MaxLength
					_, ok := res[annotation_]
					if ok {
						maxlengths = res[annotation_]
					} else {
						maxlengths = make([]*MaxLength, 0)
					}
					maxlengths = append(maxlengths, maxlength)
					res[annotation_] = maxlengths
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of MinInclusive
	case MinInclusive:
		switch fieldname {
		// insertion point for per direct association field
		case "Annotation":
			res := make(map[*Annotation][]*MinInclusive)
			for mininclusive := range stage.MinInclusives {
				if mininclusive.Annotation != nil {
					annotation_ := mininclusive.Annotation
					var mininclusives []*MinInclusive
					_, ok := res[annotation_]
					if ok {
						mininclusives = res[annotation_]
					} else {
						mininclusives = make([]*MinInclusive, 0)
					}
					mininclusives = append(mininclusives, mininclusive)
					res[annotation_] = mininclusives
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of MinLength
	case MinLength:
		switch fieldname {
		// insertion point for per direct association field
		case "Annotation":
			res := make(map[*Annotation][]*MinLength)
			for minlength := range stage.MinLengths {
				if minlength.Annotation != nil {
					annotation_ := minlength.Annotation
					var minlengths []*MinLength
					_, ok := res[annotation_]
					if ok {
						minlengths = res[annotation_]
					} else {
						minlengths = make([]*MinLength, 0)
					}
					minlengths = append(minlengths, minlength)
					res[annotation_] = minlengths
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Pattern
	case Pattern:
		switch fieldname {
		// insertion point for per direct association field
		case "Annotation":
			res := make(map[*Annotation][]*Pattern)
			for pattern := range stage.Patterns {
				if pattern.Annotation != nil {
					annotation_ := pattern.Annotation
					var patterns []*Pattern
					_, ok := res[annotation_]
					if ok {
						patterns = res[annotation_]
					} else {
						patterns = make([]*Pattern, 0)
					}
					patterns = append(patterns, pattern)
					res[annotation_] = patterns
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Restriction
	case Restriction:
		switch fieldname {
		// insertion point for per direct association field
		case "Annotation":
			res := make(map[*Annotation][]*Restriction)
			for restriction := range stage.Restrictions {
				if restriction.Annotation != nil {
					annotation_ := restriction.Annotation
					var restrictions []*Restriction
					_, ok := res[annotation_]
					if ok {
						restrictions = res[annotation_]
					} else {
						restrictions = make([]*Restriction, 0)
					}
					restrictions = append(restrictions, restriction)
					res[annotation_] = restrictions
				}
			}
			return any(res).(map[*End][]*Start)
		case "MinInclusive":
			res := make(map[*MinInclusive][]*Restriction)
			for restriction := range stage.Restrictions {
				if restriction.MinInclusive != nil {
					mininclusive_ := restriction.MinInclusive
					var restrictions []*Restriction
					_, ok := res[mininclusive_]
					if ok {
						restrictions = res[mininclusive_]
					} else {
						restrictions = make([]*Restriction, 0)
					}
					restrictions = append(restrictions, restriction)
					res[mininclusive_] = restrictions
				}
			}
			return any(res).(map[*End][]*Start)
		case "MaxInclusive":
			res := make(map[*MaxInclusive][]*Restriction)
			for restriction := range stage.Restrictions {
				if restriction.MaxInclusive != nil {
					maxinclusive_ := restriction.MaxInclusive
					var restrictions []*Restriction
					_, ok := res[maxinclusive_]
					if ok {
						restrictions = res[maxinclusive_]
					} else {
						restrictions = make([]*Restriction, 0)
					}
					restrictions = append(restrictions, restriction)
					res[maxinclusive_] = restrictions
				}
			}
			return any(res).(map[*End][]*Start)
		case "Pattern":
			res := make(map[*Pattern][]*Restriction)
			for restriction := range stage.Restrictions {
				if restriction.Pattern != nil {
					pattern_ := restriction.Pattern
					var restrictions []*Restriction
					_, ok := res[pattern_]
					if ok {
						restrictions = res[pattern_]
					} else {
						restrictions = make([]*Restriction, 0)
					}
					restrictions = append(restrictions, restriction)
					res[pattern_] = restrictions
				}
			}
			return any(res).(map[*End][]*Start)
		case "WhiteSpace":
			res := make(map[*WhiteSpace][]*Restriction)
			for restriction := range stage.Restrictions {
				if restriction.WhiteSpace != nil {
					whitespace_ := restriction.WhiteSpace
					var restrictions []*Restriction
					_, ok := res[whitespace_]
					if ok {
						restrictions = res[whitespace_]
					} else {
						restrictions = make([]*Restriction, 0)
					}
					restrictions = append(restrictions, restriction)
					res[whitespace_] = restrictions
				}
			}
			return any(res).(map[*End][]*Start)
		case "MinLength":
			res := make(map[*MinLength][]*Restriction)
			for restriction := range stage.Restrictions {
				if restriction.MinLength != nil {
					minlength_ := restriction.MinLength
					var restrictions []*Restriction
					_, ok := res[minlength_]
					if ok {
						restrictions = res[minlength_]
					} else {
						restrictions = make([]*Restriction, 0)
					}
					restrictions = append(restrictions, restriction)
					res[minlength_] = restrictions
				}
			}
			return any(res).(map[*End][]*Start)
		case "MaxLength":
			res := make(map[*MaxLength][]*Restriction)
			for restriction := range stage.Restrictions {
				if restriction.MaxLength != nil {
					maxlength_ := restriction.MaxLength
					var restrictions []*Restriction
					_, ok := res[maxlength_]
					if ok {
						restrictions = res[maxlength_]
					} else {
						restrictions = make([]*Restriction, 0)
					}
					restrictions = append(restrictions, restriction)
					res[maxlength_] = restrictions
				}
			}
			return any(res).(map[*End][]*Start)
		case "Length":
			res := make(map[*Length][]*Restriction)
			for restriction := range stage.Restrictions {
				if restriction.Length != nil {
					length_ := restriction.Length
					var restrictions []*Restriction
					_, ok := res[length_]
					if ok {
						restrictions = res[length_]
					} else {
						restrictions = make([]*Restriction, 0)
					}
					restrictions = append(restrictions, restriction)
					res[length_] = restrictions
				}
			}
			return any(res).(map[*End][]*Start)
		case "TotalDigit":
			res := make(map[*TotalDigit][]*Restriction)
			for restriction := range stage.Restrictions {
				if restriction.TotalDigit != nil {
					totaldigit_ := restriction.TotalDigit
					var restrictions []*Restriction
					_, ok := res[totaldigit_]
					if ok {
						restrictions = res[totaldigit_]
					} else {
						restrictions = make([]*Restriction, 0)
					}
					restrictions = append(restrictions, restriction)
					res[totaldigit_] = restrictions
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Schema
	case Schema:
		switch fieldname {
		// insertion point for per direct association field
		case "Annotation":
			res := make(map[*Annotation][]*Schema)
			for schema := range stage.Schemas {
				if schema.Annotation != nil {
					annotation_ := schema.Annotation
					var schemas []*Schema
					_, ok := res[annotation_]
					if ok {
						schemas = res[annotation_]
					} else {
						schemas = make([]*Schema, 0)
					}
					schemas = append(schemas, schema)
					res[annotation_] = schemas
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Sequence
	case Sequence:
		switch fieldname {
		// insertion point for per direct association field
		case "Annotation":
			res := make(map[*Annotation][]*Sequence)
			for sequence := range stage.Sequences {
				if sequence.Annotation != nil {
					annotation_ := sequence.Annotation
					var sequences []*Sequence
					_, ok := res[annotation_]
					if ok {
						sequences = res[annotation_]
					} else {
						sequences = make([]*Sequence, 0)
					}
					sequences = append(sequences, sequence)
					res[annotation_] = sequences
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of SimpleContent
	case SimpleContent:
		switch fieldname {
		// insertion point for per direct association field
		case "Extension":
			res := make(map[*Extension][]*SimpleContent)
			for simplecontent := range stage.SimpleContents {
				if simplecontent.Extension != nil {
					extension_ := simplecontent.Extension
					var simplecontents []*SimpleContent
					_, ok := res[extension_]
					if ok {
						simplecontents = res[extension_]
					} else {
						simplecontents = make([]*SimpleContent, 0)
					}
					simplecontents = append(simplecontents, simplecontent)
					res[extension_] = simplecontents
				}
			}
			return any(res).(map[*End][]*Start)
		case "Restriction":
			res := make(map[*Restriction][]*SimpleContent)
			for simplecontent := range stage.SimpleContents {
				if simplecontent.Restriction != nil {
					restriction_ := simplecontent.Restriction
					var simplecontents []*SimpleContent
					_, ok := res[restriction_]
					if ok {
						simplecontents = res[restriction_]
					} else {
						simplecontents = make([]*SimpleContent, 0)
					}
					simplecontents = append(simplecontents, simplecontent)
					res[restriction_] = simplecontents
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of SimpleType
	case SimpleType:
		switch fieldname {
		// insertion point for per direct association field
		case "Annotation":
			res := make(map[*Annotation][]*SimpleType)
			for simpletype := range stage.SimpleTypes {
				if simpletype.Annotation != nil {
					annotation_ := simpletype.Annotation
					var simpletypes []*SimpleType
					_, ok := res[annotation_]
					if ok {
						simpletypes = res[annotation_]
					} else {
						simpletypes = make([]*SimpleType, 0)
					}
					simpletypes = append(simpletypes, simpletype)
					res[annotation_] = simpletypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "Restriction":
			res := make(map[*Restriction][]*SimpleType)
			for simpletype := range stage.SimpleTypes {
				if simpletype.Restriction != nil {
					restriction_ := simpletype.Restriction
					var simpletypes []*SimpleType
					_, ok := res[restriction_]
					if ok {
						simpletypes = res[restriction_]
					} else {
						simpletypes = make([]*SimpleType, 0)
					}
					simpletypes = append(simpletypes, simpletype)
					res[restriction_] = simpletypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "Union":
			res := make(map[*Union][]*SimpleType)
			for simpletype := range stage.SimpleTypes {
				if simpletype.Union != nil {
					union_ := simpletype.Union
					var simpletypes []*SimpleType
					_, ok := res[union_]
					if ok {
						simpletypes = res[union_]
					} else {
						simpletypes = make([]*SimpleType, 0)
					}
					simpletypes = append(simpletypes, simpletype)
					res[union_] = simpletypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of TotalDigit
	case TotalDigit:
		switch fieldname {
		// insertion point for per direct association field
		case "Annotation":
			res := make(map[*Annotation][]*TotalDigit)
			for totaldigit := range stage.TotalDigits {
				if totaldigit.Annotation != nil {
					annotation_ := totaldigit.Annotation
					var totaldigits []*TotalDigit
					_, ok := res[annotation_]
					if ok {
						totaldigits = res[annotation_]
					} else {
						totaldigits = make([]*TotalDigit, 0)
					}
					totaldigits = append(totaldigits, totaldigit)
					res[annotation_] = totaldigits
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Union
	case Union:
		switch fieldname {
		// insertion point for per direct association field
		case "Annotation":
			res := make(map[*Annotation][]*Union)
			for union := range stage.Unions {
				if union.Annotation != nil {
					annotation_ := union.Annotation
					var unions []*Union
					_, ok := res[annotation_]
					if ok {
						unions = res[annotation_]
					} else {
						unions = make([]*Union, 0)
					}
					unions = append(unions, union)
					res[annotation_] = unions
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of WhiteSpace
	case WhiteSpace:
		switch fieldname {
		// insertion point for per direct association field
		case "Annotation":
			res := make(map[*Annotation][]*WhiteSpace)
			for whitespace := range stage.WhiteSpaces {
				if whitespace.Annotation != nil {
					annotation_ := whitespace.Annotation
					var whitespaces []*WhiteSpace
					_, ok := res[annotation_]
					if ok {
						whitespaces = res[annotation_]
					} else {
						whitespaces = make([]*WhiteSpace, 0)
					}
					whitespaces = append(whitespaces, whitespace)
					res[annotation_] = whitespaces
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
	// reverse maps of direct associations of All
	case All:
		switch fieldname {
		// insertion point for per direct association field
		case "Sequences":
			res := make(map[*Sequence][]*All)
			for all := range stage.Alls {
				for _, sequence_ := range all.Sequences {
					res[sequence_] = append(res[sequence_], all)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Alls":
			res := make(map[*All][]*All)
			for all := range stage.Alls {
				for _, all_ := range all.Alls {
					res[all_] = append(res[all_], all)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Choices":
			res := make(map[*Choice][]*All)
			for all := range stage.Alls {
				for _, choice_ := range all.Choices {
					res[choice_] = append(res[choice_], all)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Groups":
			res := make(map[*Group][]*All)
			for all := range stage.Alls {
				for _, group_ := range all.Groups {
					res[group_] = append(res[group_], all)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Elements":
			res := make(map[*Element][]*All)
			for all := range stage.Alls {
				for _, element_ := range all.Elements {
					res[element_] = append(res[element_], all)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Annotation
	case Annotation:
		switch fieldname {
		// insertion point for per direct association field
		case "Documentations":
			res := make(map[*Documentation][]*Annotation)
			for annotation := range stage.Annotations {
				for _, documentation_ := range annotation.Documentations {
					res[documentation_] = append(res[documentation_], annotation)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Attribute
	case Attribute:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of AttributeGroup
	case AttributeGroup:
		switch fieldname {
		// insertion point for per direct association field
		case "AttributeGroups":
			res := make(map[*AttributeGroup][]*AttributeGroup)
			for attributegroup := range stage.AttributeGroups {
				for _, attributegroup_ := range attributegroup.AttributeGroups {
					res[attributegroup_] = append(res[attributegroup_], attributegroup)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Attributes":
			res := make(map[*Attribute][]*AttributeGroup)
			for attributegroup := range stage.AttributeGroups {
				for _, attribute_ := range attributegroup.Attributes {
					res[attribute_] = append(res[attribute_], attributegroup)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Choice
	case Choice:
		switch fieldname {
		// insertion point for per direct association field
		case "Sequences":
			res := make(map[*Sequence][]*Choice)
			for choice := range stage.Choices {
				for _, sequence_ := range choice.Sequences {
					res[sequence_] = append(res[sequence_], choice)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Alls":
			res := make(map[*All][]*Choice)
			for choice := range stage.Choices {
				for _, all_ := range choice.Alls {
					res[all_] = append(res[all_], choice)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Choices":
			res := make(map[*Choice][]*Choice)
			for choice := range stage.Choices {
				for _, choice_ := range choice.Choices {
					res[choice_] = append(res[choice_], choice)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Groups":
			res := make(map[*Group][]*Choice)
			for choice := range stage.Choices {
				for _, group_ := range choice.Groups {
					res[group_] = append(res[group_], choice)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Elements":
			res := make(map[*Element][]*Choice)
			for choice := range stage.Choices {
				for _, element_ := range choice.Elements {
					res[element_] = append(res[element_], choice)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ComplexContent
	case ComplexContent:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ComplexType
	case ComplexType:
		switch fieldname {
		// insertion point for per direct association field
		case "Sequences":
			res := make(map[*Sequence][]*ComplexType)
			for complextype := range stage.ComplexTypes {
				for _, sequence_ := range complextype.Sequences {
					res[sequence_] = append(res[sequence_], complextype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Alls":
			res := make(map[*All][]*ComplexType)
			for complextype := range stage.ComplexTypes {
				for _, all_ := range complextype.Alls {
					res[all_] = append(res[all_], complextype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Choices":
			res := make(map[*Choice][]*ComplexType)
			for complextype := range stage.ComplexTypes {
				for _, choice_ := range complextype.Choices {
					res[choice_] = append(res[choice_], complextype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Groups":
			res := make(map[*Group][]*ComplexType)
			for complextype := range stage.ComplexTypes {
				for _, group_ := range complextype.Groups {
					res[group_] = append(res[group_], complextype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Elements":
			res := make(map[*Element][]*ComplexType)
			for complextype := range stage.ComplexTypes {
				for _, element_ := range complextype.Elements {
					res[element_] = append(res[element_], complextype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Attributes":
			res := make(map[*Attribute][]*ComplexType)
			for complextype := range stage.ComplexTypes {
				for _, attribute_ := range complextype.Attributes {
					res[attribute_] = append(res[attribute_], complextype)
				}
			}
			return any(res).(map[*End][]*Start)
		case "AttributeGroups":
			res := make(map[*AttributeGroup][]*ComplexType)
			for complextype := range stage.ComplexTypes {
				for _, attributegroup_ := range complextype.AttributeGroups {
					res[attributegroup_] = append(res[attributegroup_], complextype)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Documentation
	case Documentation:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Element
	case Element:
		switch fieldname {
		// insertion point for per direct association field
		case "Groups":
			res := make(map[*Group][]*Element)
			for element := range stage.Elements {
				for _, group_ := range element.Groups {
					res[group_] = append(res[group_], element)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Enumeration
	case Enumeration:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Extension
	case Extension:
		switch fieldname {
		// insertion point for per direct association field
		case "Sequences":
			res := make(map[*Sequence][]*Extension)
			for extension := range stage.Extensions {
				for _, sequence_ := range extension.Sequences {
					res[sequence_] = append(res[sequence_], extension)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Alls":
			res := make(map[*All][]*Extension)
			for extension := range stage.Extensions {
				for _, all_ := range extension.Alls {
					res[all_] = append(res[all_], extension)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Choices":
			res := make(map[*Choice][]*Extension)
			for extension := range stage.Extensions {
				for _, choice_ := range extension.Choices {
					res[choice_] = append(res[choice_], extension)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Groups":
			res := make(map[*Group][]*Extension)
			for extension := range stage.Extensions {
				for _, group_ := range extension.Groups {
					res[group_] = append(res[group_], extension)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Elements":
			res := make(map[*Element][]*Extension)
			for extension := range stage.Extensions {
				for _, element_ := range extension.Elements {
					res[element_] = append(res[element_], extension)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Attributes":
			res := make(map[*Attribute][]*Extension)
			for extension := range stage.Extensions {
				for _, attribute_ := range extension.Attributes {
					res[attribute_] = append(res[attribute_], extension)
				}
			}
			return any(res).(map[*End][]*Start)
		case "AttributeGroups":
			res := make(map[*AttributeGroup][]*Extension)
			for extension := range stage.Extensions {
				for _, attributegroup_ := range extension.AttributeGroups {
					res[attributegroup_] = append(res[attributegroup_], extension)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Group
	case Group:
		switch fieldname {
		// insertion point for per direct association field
		case "Sequences":
			res := make(map[*Sequence][]*Group)
			for group := range stage.Groups {
				for _, sequence_ := range group.Sequences {
					res[sequence_] = append(res[sequence_], group)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Alls":
			res := make(map[*All][]*Group)
			for group := range stage.Groups {
				for _, all_ := range group.Alls {
					res[all_] = append(res[all_], group)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Choices":
			res := make(map[*Choice][]*Group)
			for group := range stage.Groups {
				for _, choice_ := range group.Choices {
					res[choice_] = append(res[choice_], group)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Groups":
			res := make(map[*Group][]*Group)
			for group := range stage.Groups {
				for _, group_ := range group.Groups {
					res[group_] = append(res[group_], group)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Elements":
			res := make(map[*Element][]*Group)
			for group := range stage.Groups {
				for _, element_ := range group.Elements {
					res[element_] = append(res[element_], group)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Length
	case Length:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MaxInclusive
	case MaxInclusive:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MaxLength
	case MaxLength:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MinInclusive
	case MinInclusive:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MinLength
	case MinLength:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Pattern
	case Pattern:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Restriction
	case Restriction:
		switch fieldname {
		// insertion point for per direct association field
		case "Enumerations":
			res := make(map[*Enumeration][]*Restriction)
			for restriction := range stage.Restrictions {
				for _, enumeration_ := range restriction.Enumerations {
					res[enumeration_] = append(res[enumeration_], restriction)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Schema
	case Schema:
		switch fieldname {
		// insertion point for per direct association field
		case "Elements":
			res := make(map[*Element][]*Schema)
			for schema := range stage.Schemas {
				for _, element_ := range schema.Elements {
					res[element_] = append(res[element_], schema)
				}
			}
			return any(res).(map[*End][]*Start)
		case "SimpleTypes":
			res := make(map[*SimpleType][]*Schema)
			for schema := range stage.Schemas {
				for _, simpletype_ := range schema.SimpleTypes {
					res[simpletype_] = append(res[simpletype_], schema)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ComplexTypes":
			res := make(map[*ComplexType][]*Schema)
			for schema := range stage.Schemas {
				for _, complextype_ := range schema.ComplexTypes {
					res[complextype_] = append(res[complextype_], schema)
				}
			}
			return any(res).(map[*End][]*Start)
		case "AttributeGroups":
			res := make(map[*AttributeGroup][]*Schema)
			for schema := range stage.Schemas {
				for _, attributegroup_ := range schema.AttributeGroups {
					res[attributegroup_] = append(res[attributegroup_], schema)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Groups":
			res := make(map[*Group][]*Schema)
			for schema := range stage.Schemas {
				for _, group_ := range schema.Groups {
					res[group_] = append(res[group_], schema)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Sequence
	case Sequence:
		switch fieldname {
		// insertion point for per direct association field
		case "Sequences":
			res := make(map[*Sequence][]*Sequence)
			for sequence := range stage.Sequences {
				for _, sequence_ := range sequence.Sequences {
					res[sequence_] = append(res[sequence_], sequence)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Alls":
			res := make(map[*All][]*Sequence)
			for sequence := range stage.Sequences {
				for _, all_ := range sequence.Alls {
					res[all_] = append(res[all_], sequence)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Choices":
			res := make(map[*Choice][]*Sequence)
			for sequence := range stage.Sequences {
				for _, choice_ := range sequence.Choices {
					res[choice_] = append(res[choice_], sequence)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Groups":
			res := make(map[*Group][]*Sequence)
			for sequence := range stage.Sequences {
				for _, group_ := range sequence.Groups {
					res[group_] = append(res[group_], sequence)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Elements":
			res := make(map[*Element][]*Sequence)
			for sequence := range stage.Sequences {
				for _, element_ := range sequence.Elements {
					res[element_] = append(res[element_], sequence)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of SimpleContent
	case SimpleContent:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SimpleType
	case SimpleType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TotalDigit
	case TotalDigit:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Union
	case Union:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of WhiteSpace
	case WhiteSpace:
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
	case *All:
		res = "All"
	case *Annotation:
		res = "Annotation"
	case *Attribute:
		res = "Attribute"
	case *AttributeGroup:
		res = "AttributeGroup"
	case *Choice:
		res = "Choice"
	case *ComplexContent:
		res = "ComplexContent"
	case *ComplexType:
		res = "ComplexType"
	case *Documentation:
		res = "Documentation"
	case *Element:
		res = "Element"
	case *Enumeration:
		res = "Enumeration"
	case *Extension:
		res = "Extension"
	case *Group:
		res = "Group"
	case *Length:
		res = "Length"
	case *MaxInclusive:
		res = "MaxInclusive"
	case *MaxLength:
		res = "MaxLength"
	case *MinInclusive:
		res = "MinInclusive"
	case *MinLength:
		res = "MinLength"
	case *Pattern:
		res = "Pattern"
	case *Restriction:
		res = "Restriction"
	case *Schema:
		res = "Schema"
	case *Sequence:
		res = "Sequence"
	case *SimpleContent:
		res = "SimpleContent"
	case *SimpleType:
		res = "SimpleType"
	case *TotalDigit:
		res = "TotalDigit"
	case *Union:
		res = "Union"
	case *WhiteSpace:
		res = "WhiteSpace"
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
	case *All:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "All"
		rf.Fieldname = "Alls"
		res = append(res, rf)
		rf.GongstructName = "Choice"
		rf.Fieldname = "Alls"
		res = append(res, rf)
		rf.GongstructName = "ComplexType"
		rf.Fieldname = "Alls"
		res = append(res, rf)
		rf.GongstructName = "Extension"
		rf.Fieldname = "Alls"
		res = append(res, rf)
		rf.GongstructName = "Group"
		rf.Fieldname = "Alls"
		res = append(res, rf)
		rf.GongstructName = "Sequence"
		rf.Fieldname = "Alls"
		res = append(res, rf)
	case *Annotation:
		var rf ReverseField
		_ = rf
	case *Attribute:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "AttributeGroup"
		rf.Fieldname = "Attributes"
		res = append(res, rf)
		rf.GongstructName = "ComplexType"
		rf.Fieldname = "Attributes"
		res = append(res, rf)
		rf.GongstructName = "Extension"
		rf.Fieldname = "Attributes"
		res = append(res, rf)
	case *AttributeGroup:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "AttributeGroup"
		rf.Fieldname = "AttributeGroups"
		res = append(res, rf)
		rf.GongstructName = "ComplexType"
		rf.Fieldname = "AttributeGroups"
		res = append(res, rf)
		rf.GongstructName = "Extension"
		rf.Fieldname = "AttributeGroups"
		res = append(res, rf)
		rf.GongstructName = "Schema"
		rf.Fieldname = "AttributeGroups"
		res = append(res, rf)
	case *Choice:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "All"
		rf.Fieldname = "Choices"
		res = append(res, rf)
		rf.GongstructName = "Choice"
		rf.Fieldname = "Choices"
		res = append(res, rf)
		rf.GongstructName = "ComplexType"
		rf.Fieldname = "Choices"
		res = append(res, rf)
		rf.GongstructName = "Extension"
		rf.Fieldname = "Choices"
		res = append(res, rf)
		rf.GongstructName = "Group"
		rf.Fieldname = "Choices"
		res = append(res, rf)
		rf.GongstructName = "Sequence"
		rf.Fieldname = "Choices"
		res = append(res, rf)
	case *ComplexContent:
		var rf ReverseField
		_ = rf
	case *ComplexType:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Schema"
		rf.Fieldname = "ComplexTypes"
		res = append(res, rf)
	case *Documentation:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Annotation"
		rf.Fieldname = "Documentations"
		res = append(res, rf)
	case *Element:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "All"
		rf.Fieldname = "Elements"
		res = append(res, rf)
		rf.GongstructName = "Choice"
		rf.Fieldname = "Elements"
		res = append(res, rf)
		rf.GongstructName = "ComplexType"
		rf.Fieldname = "Elements"
		res = append(res, rf)
		rf.GongstructName = "Extension"
		rf.Fieldname = "Elements"
		res = append(res, rf)
		rf.GongstructName = "Group"
		rf.Fieldname = "Elements"
		res = append(res, rf)
		rf.GongstructName = "Schema"
		rf.Fieldname = "Elements"
		res = append(res, rf)
		rf.GongstructName = "Sequence"
		rf.Fieldname = "Elements"
		res = append(res, rf)
	case *Enumeration:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Restriction"
		rf.Fieldname = "Enumerations"
		res = append(res, rf)
	case *Extension:
		var rf ReverseField
		_ = rf
	case *Group:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "All"
		rf.Fieldname = "Groups"
		res = append(res, rf)
		rf.GongstructName = "Choice"
		rf.Fieldname = "Groups"
		res = append(res, rf)
		rf.GongstructName = "ComplexType"
		rf.Fieldname = "Groups"
		res = append(res, rf)
		rf.GongstructName = "Element"
		rf.Fieldname = "Groups"
		res = append(res, rf)
		rf.GongstructName = "Extension"
		rf.Fieldname = "Groups"
		res = append(res, rf)
		rf.GongstructName = "Group"
		rf.Fieldname = "Groups"
		res = append(res, rf)
		rf.GongstructName = "Schema"
		rf.Fieldname = "Groups"
		res = append(res, rf)
		rf.GongstructName = "Sequence"
		rf.Fieldname = "Groups"
		res = append(res, rf)
	case *Length:
		var rf ReverseField
		_ = rf
	case *MaxInclusive:
		var rf ReverseField
		_ = rf
	case *MaxLength:
		var rf ReverseField
		_ = rf
	case *MinInclusive:
		var rf ReverseField
		_ = rf
	case *MinLength:
		var rf ReverseField
		_ = rf
	case *Pattern:
		var rf ReverseField
		_ = rf
	case *Restriction:
		var rf ReverseField
		_ = rf
	case *Schema:
		var rf ReverseField
		_ = rf
	case *Sequence:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "All"
		rf.Fieldname = "Sequences"
		res = append(res, rf)
		rf.GongstructName = "Choice"
		rf.Fieldname = "Sequences"
		res = append(res, rf)
		rf.GongstructName = "ComplexType"
		rf.Fieldname = "Sequences"
		res = append(res, rf)
		rf.GongstructName = "Extension"
		rf.Fieldname = "Sequences"
		res = append(res, rf)
		rf.GongstructName = "Group"
		rf.Fieldname = "Sequences"
		res = append(res, rf)
		rf.GongstructName = "Sequence"
		rf.Fieldname = "Sequences"
		res = append(res, rf)
	case *SimpleContent:
		var rf ReverseField
		_ = rf
	case *SimpleType:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Schema"
		rf.Fieldname = "SimpleTypes"
		res = append(res, rf)
	case *TotalDigit:
		var rf ReverseField
		_ = rf
	case *Union:
		var rf ReverseField
		_ = rf
	case *WhiteSpace:
		var rf ReverseField
		_ = rf
	}
	return
}

// insertion point for get fields header method
func (all *All) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Annotation",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Annotation",
		},
		{
			Name:               "OuterElementName",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Sequences",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Sequence",
		},
		{
			Name:                 "Alls",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "All",
		},
		{
			Name:                 "Choices",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Choice",
		},
		{
			Name:                 "Groups",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Group",
		},
		{
			Name:                 "Elements",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Element",
		},
		{
			Name:               "Order",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "Depth",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "MinOccurs",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "MaxOccurs",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (annotation *Annotation) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Documentations",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Documentation",
		},
	}
	return
}

func (attribute *Attribute) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "Type",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Annotation",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Annotation",
		},
		{
			Name:               "HasNameConflict",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "GoIdentifier",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Default",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Use",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Form",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Fixed",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Ref",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "TargetNamespace",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "SimpleType",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IDXSD",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (attributegroup *AttributeGroup) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "Annotation",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Annotation",
		},
		{
			Name:               "HasNameConflict",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "GoIdentifier",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "AttributeGroups",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "AttributeGroup",
		},
		{
			Name:               "Ref",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Attributes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Attribute",
		},
		{
			Name:               "Order",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "Depth",
			GongFieldValueType: GongFieldValueTypeInt,
		},
	}
	return
}

func (choice *Choice) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Annotation",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Annotation",
		},
		{
			Name:               "OuterElementName",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Sequences",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Sequence",
		},
		{
			Name:                 "Alls",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "All",
		},
		{
			Name:                 "Choices",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Choice",
		},
		{
			Name:                 "Groups",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Group",
		},
		{
			Name:                 "Elements",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Element",
		},
		{
			Name:               "Order",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "Depth",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "MinOccurs",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "MaxOccurs",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsDuplicatedInXSD",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (complexcontent *ComplexContent) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (complextype *ComplexType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "HasNameConflict",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "GoIdentifier",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsAnonymous",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "OuterElement",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Element",
		},
		{
			Name:                 "Annotation",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Annotation",
		},
		{
			Name:               "NameXSD",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "OuterElementName",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Sequences",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Sequence",
		},
		{
			Name:                 "Alls",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "All",
		},
		{
			Name:                 "Choices",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Choice",
		},
		{
			Name:                 "Groups",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Group",
		},
		{
			Name:                 "Elements",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Element",
		},
		{
			Name:               "Order",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "Depth",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "MinOccurs",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "MaxOccurs",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Extension",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Extension",
		},
		{
			Name:                 "SimpleContent",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "SimpleContent",
		},
		{
			Name:                 "ComplexContent",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ComplexContent",
		},
		{
			Name:                 "Attributes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Attribute",
		},
		{
			Name:                 "AttributeGroups",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "AttributeGroup",
		},
		{
			Name:               "IsDuplicatedInXSD",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (documentation *Documentation) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Text",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Source",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Lang",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (element *Element) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Order",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "Depth",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "HasNameConflict",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "GoIdentifier",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Annotation",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Annotation",
		},
		{
			Name:               "NameXSD",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Type",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "MinOccurs",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "MaxOccurs",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Default",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Fixed",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Nillable",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Ref",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Abstract",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Form",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Block",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Final",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "SimpleType",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "SimpleType",
		},
		{
			Name:                 "ComplexType",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ComplexType",
		},
		{
			Name:                 "Groups",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Group",
		},
		{
			Name:               "IsDuplicatedInXSD",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (enumeration *Enumeration) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Annotation",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Annotation",
		},
		{
			Name:               "Value",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (extension *Extension) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "OuterElementName",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Sequences",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Sequence",
		},
		{
			Name:                 "Alls",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "All",
		},
		{
			Name:                 "Choices",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Choice",
		},
		{
			Name:                 "Groups",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Group",
		},
		{
			Name:                 "Elements",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Element",
		},
		{
			Name:               "Order",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "Depth",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "MinOccurs",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "MaxOccurs",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Base",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Ref",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Attributes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Attribute",
		},
		{
			Name:                 "AttributeGroups",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "AttributeGroup",
		},
	}
	return
}

func (group *Group) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Annotation",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Annotation",
		},
		{
			Name:               "NameXSD",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Ref",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsAnonymous",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "OuterElement",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Element",
		},
		{
			Name:               "HasNameConflict",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "GoIdentifier",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "OuterElementName",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Sequences",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Sequence",
		},
		{
			Name:                 "Alls",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "All",
		},
		{
			Name:                 "Choices",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Choice",
		},
		{
			Name:                 "Groups",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Group",
		},
		{
			Name:                 "Elements",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Element",
		},
		{
			Name:               "Order",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "Depth",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "MinOccurs",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "MaxOccurs",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (length *Length) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Annotation",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Annotation",
		},
		{
			Name:               "Value",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (maxinclusive *MaxInclusive) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Annotation",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Annotation",
		},
		{
			Name:               "Value",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (maxlength *MaxLength) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Annotation",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Annotation",
		},
		{
			Name:               "Value",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (mininclusive *MinInclusive) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Annotation",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Annotation",
		},
		{
			Name:               "Value",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (minlength *MinLength) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Annotation",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Annotation",
		},
		{
			Name:               "Value",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (pattern *Pattern) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Annotation",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Annotation",
		},
		{
			Name:               "Value",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (restriction *Restriction) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Annotation",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Annotation",
		},
		{
			Name:               "Base",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Enumerations",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Enumeration",
		},
		{
			Name:                 "MinInclusive",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "MinInclusive",
		},
		{
			Name:                 "MaxInclusive",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "MaxInclusive",
		},
		{
			Name:                 "Pattern",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Pattern",
		},
		{
			Name:                 "WhiteSpace",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "WhiteSpace",
		},
		{
			Name:                 "MinLength",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "MinLength",
		},
		{
			Name:                 "MaxLength",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "MaxLength",
		},
		{
			Name:                 "Length",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Length",
		},
		{
			Name:                 "TotalDigit",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "TotalDigit",
		},
	}
	return
}

func (schema *Schema) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Xs",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Annotation",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Annotation",
		},
		{
			Name:                 "Elements",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Element",
		},
		{
			Name:                 "SimpleTypes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "SimpleType",
		},
		{
			Name:                 "ComplexTypes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ComplexType",
		},
		{
			Name:                 "AttributeGroups",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "AttributeGroup",
		},
		{
			Name:                 "Groups",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Group",
		},
		{
			Name:               "Order",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "Depth",
			GongFieldValueType: GongFieldValueTypeInt,
		},
	}
	return
}

func (sequence *Sequence) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Annotation",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Annotation",
		},
		{
			Name:               "OuterElementName",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Sequences",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Sequence",
		},
		{
			Name:                 "Alls",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "All",
		},
		{
			Name:                 "Choices",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Choice",
		},
		{
			Name:                 "Groups",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Group",
		},
		{
			Name:                 "Elements",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Element",
		},
		{
			Name:               "Order",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "Depth",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "MinOccurs",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "MaxOccurs",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (simplecontent *SimpleContent) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Extension",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Extension",
		},
		{
			Name:                 "Restriction",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Restriction",
		},
	}
	return
}

func (simpletype *SimpleType) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Annotation",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Annotation",
		},
		{
			Name:               "NameXSD",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Restriction",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Restriction",
		},
		{
			Name:                 "Union",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Union",
		},
		{
			Name:               "Order",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "Depth",
			GongFieldValueType: GongFieldValueTypeInt,
		},
	}
	return
}

func (totaldigit *TotalDigit) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Annotation",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Annotation",
		},
		{
			Name:               "Value",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (union *Union) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Annotation",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Annotation",
		},
		{
			Name:               "MemberTypes",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (whitespace *WhiteSpace) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Annotation",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Annotation",
		},
		{
			Name:               "Value",
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
func (all *All) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = all.Name
	case "Annotation":
		res.GongFieldValueType = GongFieldValueTypePointer
		if all.Annotation != nil {
			res.valueString = all.Annotation.Name
			res.ids = all.Annotation.GongGetUUID(stage)
		}
	case "OuterElementName":
		res.valueString = all.OuterElementName
	case "Sequences":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range all.Sequences {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Alls":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range all.Alls {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Choices":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range all.Choices {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Groups":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range all.Groups {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Elements":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range all.Elements {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Order":
		res.valueString = fmt.Sprintf("%d", all.Order)
		res.valueInt = all.Order
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Depth":
		res.valueString = fmt.Sprintf("%d", all.Depth)
		res.valueInt = all.Depth
		res.GongFieldValueType = GongFieldValueTypeInt
	case "MinOccurs":
		res.valueString = all.MinOccurs
	case "MaxOccurs":
		res.valueString = all.MaxOccurs
	}
	return
}

func (annotation *Annotation) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = annotation.Name
	case "Documentations":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range annotation.Documentations {
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

func (attribute *Attribute) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = attribute.Name
	case "NameXSD":
		res.valueString = attribute.NameXSD
	case "Type":
		res.valueString = attribute.Type
	case "Annotation":
		res.GongFieldValueType = GongFieldValueTypePointer
		if attribute.Annotation != nil {
			res.valueString = attribute.Annotation.Name
			res.ids = attribute.Annotation.GongGetUUID(stage)
		}
	case "HasNameConflict":
		res.valueString = fmt.Sprintf("%t", attribute.HasNameConflict)
		res.valueBool = attribute.HasNameConflict
		res.GongFieldValueType = GongFieldValueTypeBool
	case "GoIdentifier":
		res.valueString = attribute.GoIdentifier
	case "Default":
		res.valueString = attribute.Default
	case "Use":
		res.valueString = attribute.Use
	case "Form":
		res.valueString = attribute.Form
	case "Fixed":
		res.valueString = attribute.Fixed
	case "Ref":
		res.valueString = attribute.Ref
	case "TargetNamespace":
		res.valueString = attribute.TargetNamespace
	case "SimpleType":
		res.valueString = attribute.SimpleType
	case "IDXSD":
		res.valueString = attribute.IDXSD
	}
	return
}

func (attributegroup *AttributeGroup) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = attributegroup.Name
	case "NameXSD":
		res.valueString = attributegroup.NameXSD
	case "Annotation":
		res.GongFieldValueType = GongFieldValueTypePointer
		if attributegroup.Annotation != nil {
			res.valueString = attributegroup.Annotation.Name
			res.ids = attributegroup.Annotation.GongGetUUID(stage)
		}
	case "HasNameConflict":
		res.valueString = fmt.Sprintf("%t", attributegroup.HasNameConflict)
		res.valueBool = attributegroup.HasNameConflict
		res.GongFieldValueType = GongFieldValueTypeBool
	case "GoIdentifier":
		res.valueString = attributegroup.GoIdentifier
	case "AttributeGroups":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range attributegroup.AttributeGroups {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Ref":
		res.valueString = attributegroup.Ref
	case "Attributes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range attributegroup.Attributes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Order":
		res.valueString = fmt.Sprintf("%d", attributegroup.Order)
		res.valueInt = attributegroup.Order
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Depth":
		res.valueString = fmt.Sprintf("%d", attributegroup.Depth)
		res.valueInt = attributegroup.Depth
		res.GongFieldValueType = GongFieldValueTypeInt
	}
	return
}

func (choice *Choice) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = choice.Name
	case "Annotation":
		res.GongFieldValueType = GongFieldValueTypePointer
		if choice.Annotation != nil {
			res.valueString = choice.Annotation.Name
			res.ids = choice.Annotation.GongGetUUID(stage)
		}
	case "OuterElementName":
		res.valueString = choice.OuterElementName
	case "Sequences":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range choice.Sequences {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Alls":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range choice.Alls {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Choices":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range choice.Choices {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Groups":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range choice.Groups {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Elements":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range choice.Elements {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Order":
		res.valueString = fmt.Sprintf("%d", choice.Order)
		res.valueInt = choice.Order
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Depth":
		res.valueString = fmt.Sprintf("%d", choice.Depth)
		res.valueInt = choice.Depth
		res.GongFieldValueType = GongFieldValueTypeInt
	case "MinOccurs":
		res.valueString = choice.MinOccurs
	case "MaxOccurs":
		res.valueString = choice.MaxOccurs
	case "IsDuplicatedInXSD":
		res.valueString = fmt.Sprintf("%t", choice.IsDuplicatedInXSD)
		res.valueBool = choice.IsDuplicatedInXSD
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (complexcontent *ComplexContent) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = complexcontent.Name
	}
	return
}

func (complextype *ComplexType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = complextype.Name
	case "HasNameConflict":
		res.valueString = fmt.Sprintf("%t", complextype.HasNameConflict)
		res.valueBool = complextype.HasNameConflict
		res.GongFieldValueType = GongFieldValueTypeBool
	case "GoIdentifier":
		res.valueString = complextype.GoIdentifier
	case "IsAnonymous":
		res.valueString = fmt.Sprintf("%t", complextype.IsAnonymous)
		res.valueBool = complextype.IsAnonymous
		res.GongFieldValueType = GongFieldValueTypeBool
	case "OuterElement":
		res.GongFieldValueType = GongFieldValueTypePointer
		if complextype.OuterElement != nil {
			res.valueString = complextype.OuterElement.Name
			res.ids = complextype.OuterElement.GongGetUUID(stage)
		}
	case "Annotation":
		res.GongFieldValueType = GongFieldValueTypePointer
		if complextype.Annotation != nil {
			res.valueString = complextype.Annotation.Name
			res.ids = complextype.Annotation.GongGetUUID(stage)
		}
	case "NameXSD":
		res.valueString = complextype.NameXSD
	case "OuterElementName":
		res.valueString = complextype.OuterElementName
	case "Sequences":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range complextype.Sequences {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Alls":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range complextype.Alls {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Choices":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range complextype.Choices {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Groups":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range complextype.Groups {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Elements":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range complextype.Elements {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Order":
		res.valueString = fmt.Sprintf("%d", complextype.Order)
		res.valueInt = complextype.Order
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Depth":
		res.valueString = fmt.Sprintf("%d", complextype.Depth)
		res.valueInt = complextype.Depth
		res.GongFieldValueType = GongFieldValueTypeInt
	case "MinOccurs":
		res.valueString = complextype.MinOccurs
	case "MaxOccurs":
		res.valueString = complextype.MaxOccurs
	case "Extension":
		res.GongFieldValueType = GongFieldValueTypePointer
		if complextype.Extension != nil {
			res.valueString = complextype.Extension.Name
			res.ids = complextype.Extension.GongGetUUID(stage)
		}
	case "SimpleContent":
		res.GongFieldValueType = GongFieldValueTypePointer
		if complextype.SimpleContent != nil {
			res.valueString = complextype.SimpleContent.Name
			res.ids = complextype.SimpleContent.GongGetUUID(stage)
		}
	case "ComplexContent":
		res.GongFieldValueType = GongFieldValueTypePointer
		if complextype.ComplexContent != nil {
			res.valueString = complextype.ComplexContent.Name
			res.ids = complextype.ComplexContent.GongGetUUID(stage)
		}
	case "Attributes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range complextype.Attributes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "AttributeGroups":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range complextype.AttributeGroups {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsDuplicatedInXSD":
		res.valueString = fmt.Sprintf("%t", complextype.IsDuplicatedInXSD)
		res.valueBool = complextype.IsDuplicatedInXSD
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (documentation *Documentation) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = documentation.Name
	case "Text":
		res.valueString = documentation.Text
	case "Source":
		res.valueString = documentation.Source
	case "Lang":
		res.valueString = documentation.Lang
	}
	return
}

func (element *Element) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = element.Name
	case "Order":
		res.valueString = fmt.Sprintf("%d", element.Order)
		res.valueInt = element.Order
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Depth":
		res.valueString = fmt.Sprintf("%d", element.Depth)
		res.valueInt = element.Depth
		res.GongFieldValueType = GongFieldValueTypeInt
	case "HasNameConflict":
		res.valueString = fmt.Sprintf("%t", element.HasNameConflict)
		res.valueBool = element.HasNameConflict
		res.GongFieldValueType = GongFieldValueTypeBool
	case "GoIdentifier":
		res.valueString = element.GoIdentifier
	case "Annotation":
		res.GongFieldValueType = GongFieldValueTypePointer
		if element.Annotation != nil {
			res.valueString = element.Annotation.Name
			res.ids = element.Annotation.GongGetUUID(stage)
		}
	case "NameXSD":
		res.valueString = element.NameXSD
	case "Type":
		res.valueString = element.Type
	case "MinOccurs":
		res.valueString = element.MinOccurs
	case "MaxOccurs":
		res.valueString = element.MaxOccurs
	case "Default":
		res.valueString = element.Default
	case "Fixed":
		res.valueString = element.Fixed
	case "Nillable":
		res.valueString = element.Nillable
	case "Ref":
		res.valueString = element.Ref
	case "Abstract":
		res.valueString = element.Abstract
	case "Form":
		res.valueString = element.Form
	case "Block":
		res.valueString = element.Block
	case "Final":
		res.valueString = element.Final
	case "SimpleType":
		res.GongFieldValueType = GongFieldValueTypePointer
		if element.SimpleType != nil {
			res.valueString = element.SimpleType.Name
			res.ids = element.SimpleType.GongGetUUID(stage)
		}
	case "ComplexType":
		res.GongFieldValueType = GongFieldValueTypePointer
		if element.ComplexType != nil {
			res.valueString = element.ComplexType.Name
			res.ids = element.ComplexType.GongGetUUID(stage)
		}
	case "Groups":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range element.Groups {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsDuplicatedInXSD":
		res.valueString = fmt.Sprintf("%t", element.IsDuplicatedInXSD)
		res.valueBool = element.IsDuplicatedInXSD
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (enumeration *Enumeration) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = enumeration.Name
	case "Annotation":
		res.GongFieldValueType = GongFieldValueTypePointer
		if enumeration.Annotation != nil {
			res.valueString = enumeration.Annotation.Name
			res.ids = enumeration.Annotation.GongGetUUID(stage)
		}
	case "Value":
		res.valueString = enumeration.Value
	}
	return
}

func (extension *Extension) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = extension.Name
	case "OuterElementName":
		res.valueString = extension.OuterElementName
	case "Sequences":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range extension.Sequences {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Alls":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range extension.Alls {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Choices":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range extension.Choices {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Groups":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range extension.Groups {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Elements":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range extension.Elements {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Order":
		res.valueString = fmt.Sprintf("%d", extension.Order)
		res.valueInt = extension.Order
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Depth":
		res.valueString = fmt.Sprintf("%d", extension.Depth)
		res.valueInt = extension.Depth
		res.GongFieldValueType = GongFieldValueTypeInt
	case "MinOccurs":
		res.valueString = extension.MinOccurs
	case "MaxOccurs":
		res.valueString = extension.MaxOccurs
	case "Base":
		res.valueString = extension.Base
	case "Ref":
		res.valueString = extension.Ref
	case "Attributes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range extension.Attributes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "AttributeGroups":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range extension.AttributeGroups {
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

func (group *Group) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = group.Name
	case "Annotation":
		res.GongFieldValueType = GongFieldValueTypePointer
		if group.Annotation != nil {
			res.valueString = group.Annotation.Name
			res.ids = group.Annotation.GongGetUUID(stage)
		}
	case "NameXSD":
		res.valueString = group.NameXSD
	case "Ref":
		res.valueString = group.Ref
	case "IsAnonymous":
		res.valueString = fmt.Sprintf("%t", group.IsAnonymous)
		res.valueBool = group.IsAnonymous
		res.GongFieldValueType = GongFieldValueTypeBool
	case "OuterElement":
		res.GongFieldValueType = GongFieldValueTypePointer
		if group.OuterElement != nil {
			res.valueString = group.OuterElement.Name
			res.ids = group.OuterElement.GongGetUUID(stage)
		}
	case "HasNameConflict":
		res.valueString = fmt.Sprintf("%t", group.HasNameConflict)
		res.valueBool = group.HasNameConflict
		res.GongFieldValueType = GongFieldValueTypeBool
	case "GoIdentifier":
		res.valueString = group.GoIdentifier
	case "OuterElementName":
		res.valueString = group.OuterElementName
	case "Sequences":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range group.Sequences {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Alls":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range group.Alls {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Choices":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range group.Choices {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Groups":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range group.Groups {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Elements":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range group.Elements {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Order":
		res.valueString = fmt.Sprintf("%d", group.Order)
		res.valueInt = group.Order
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Depth":
		res.valueString = fmt.Sprintf("%d", group.Depth)
		res.valueInt = group.Depth
		res.GongFieldValueType = GongFieldValueTypeInt
	case "MinOccurs":
		res.valueString = group.MinOccurs
	case "MaxOccurs":
		res.valueString = group.MaxOccurs
	}
	return
}

func (length *Length) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = length.Name
	case "Annotation":
		res.GongFieldValueType = GongFieldValueTypePointer
		if length.Annotation != nil {
			res.valueString = length.Annotation.Name
			res.ids = length.Annotation.GongGetUUID(stage)
		}
	case "Value":
		res.valueString = length.Value
	}
	return
}

func (maxinclusive *MaxInclusive) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = maxinclusive.Name
	case "Annotation":
		res.GongFieldValueType = GongFieldValueTypePointer
		if maxinclusive.Annotation != nil {
			res.valueString = maxinclusive.Annotation.Name
			res.ids = maxinclusive.Annotation.GongGetUUID(stage)
		}
	case "Value":
		res.valueString = maxinclusive.Value
	}
	return
}

func (maxlength *MaxLength) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = maxlength.Name
	case "Annotation":
		res.GongFieldValueType = GongFieldValueTypePointer
		if maxlength.Annotation != nil {
			res.valueString = maxlength.Annotation.Name
			res.ids = maxlength.Annotation.GongGetUUID(stage)
		}
	case "Value":
		res.valueString = maxlength.Value
	}
	return
}

func (mininclusive *MinInclusive) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = mininclusive.Name
	case "Annotation":
		res.GongFieldValueType = GongFieldValueTypePointer
		if mininclusive.Annotation != nil {
			res.valueString = mininclusive.Annotation.Name
			res.ids = mininclusive.Annotation.GongGetUUID(stage)
		}
	case "Value":
		res.valueString = mininclusive.Value
	}
	return
}

func (minlength *MinLength) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = minlength.Name
	case "Annotation":
		res.GongFieldValueType = GongFieldValueTypePointer
		if minlength.Annotation != nil {
			res.valueString = minlength.Annotation.Name
			res.ids = minlength.Annotation.GongGetUUID(stage)
		}
	case "Value":
		res.valueString = minlength.Value
	}
	return
}

func (pattern *Pattern) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = pattern.Name
	case "Annotation":
		res.GongFieldValueType = GongFieldValueTypePointer
		if pattern.Annotation != nil {
			res.valueString = pattern.Annotation.Name
			res.ids = pattern.Annotation.GongGetUUID(stage)
		}
	case "Value":
		res.valueString = pattern.Value
	}
	return
}

func (restriction *Restriction) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = restriction.Name
	case "Annotation":
		res.GongFieldValueType = GongFieldValueTypePointer
		if restriction.Annotation != nil {
			res.valueString = restriction.Annotation.Name
			res.ids = restriction.Annotation.GongGetUUID(stage)
		}
	case "Base":
		res.valueString = restriction.Base
	case "Enumerations":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range restriction.Enumerations {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "MinInclusive":
		res.GongFieldValueType = GongFieldValueTypePointer
		if restriction.MinInclusive != nil {
			res.valueString = restriction.MinInclusive.Name
			res.ids = restriction.MinInclusive.GongGetUUID(stage)
		}
	case "MaxInclusive":
		res.GongFieldValueType = GongFieldValueTypePointer
		if restriction.MaxInclusive != nil {
			res.valueString = restriction.MaxInclusive.Name
			res.ids = restriction.MaxInclusive.GongGetUUID(stage)
		}
	case "Pattern":
		res.GongFieldValueType = GongFieldValueTypePointer
		if restriction.Pattern != nil {
			res.valueString = restriction.Pattern.Name
			res.ids = restriction.Pattern.GongGetUUID(stage)
		}
	case "WhiteSpace":
		res.GongFieldValueType = GongFieldValueTypePointer
		if restriction.WhiteSpace != nil {
			res.valueString = restriction.WhiteSpace.Name
			res.ids = restriction.WhiteSpace.GongGetUUID(stage)
		}
	case "MinLength":
		res.GongFieldValueType = GongFieldValueTypePointer
		if restriction.MinLength != nil {
			res.valueString = restriction.MinLength.Name
			res.ids = restriction.MinLength.GongGetUUID(stage)
		}
	case "MaxLength":
		res.GongFieldValueType = GongFieldValueTypePointer
		if restriction.MaxLength != nil {
			res.valueString = restriction.MaxLength.Name
			res.ids = restriction.MaxLength.GongGetUUID(stage)
		}
	case "Length":
		res.GongFieldValueType = GongFieldValueTypePointer
		if restriction.Length != nil {
			res.valueString = restriction.Length.Name
			res.ids = restriction.Length.GongGetUUID(stage)
		}
	case "TotalDigit":
		res.GongFieldValueType = GongFieldValueTypePointer
		if restriction.TotalDigit != nil {
			res.valueString = restriction.TotalDigit.Name
			res.ids = restriction.TotalDigit.GongGetUUID(stage)
		}
	}
	return
}

func (schema *Schema) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = schema.Name
	case "Xs":
		res.valueString = schema.Xs
	case "Annotation":
		res.GongFieldValueType = GongFieldValueTypePointer
		if schema.Annotation != nil {
			res.valueString = schema.Annotation.Name
			res.ids = schema.Annotation.GongGetUUID(stage)
		}
	case "Elements":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range schema.Elements {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "SimpleTypes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range schema.SimpleTypes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "ComplexTypes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range schema.ComplexTypes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "AttributeGroups":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range schema.AttributeGroups {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Groups":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range schema.Groups {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Order":
		res.valueString = fmt.Sprintf("%d", schema.Order)
		res.valueInt = schema.Order
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Depth":
		res.valueString = fmt.Sprintf("%d", schema.Depth)
		res.valueInt = schema.Depth
		res.GongFieldValueType = GongFieldValueTypeInt
	}
	return
}

func (sequence *Sequence) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = sequence.Name
	case "Annotation":
		res.GongFieldValueType = GongFieldValueTypePointer
		if sequence.Annotation != nil {
			res.valueString = sequence.Annotation.Name
			res.ids = sequence.Annotation.GongGetUUID(stage)
		}
	case "OuterElementName":
		res.valueString = sequence.OuterElementName
	case "Sequences":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range sequence.Sequences {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Alls":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range sequence.Alls {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Choices":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range sequence.Choices {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Groups":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range sequence.Groups {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Elements":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range sequence.Elements {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Order":
		res.valueString = fmt.Sprintf("%d", sequence.Order)
		res.valueInt = sequence.Order
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Depth":
		res.valueString = fmt.Sprintf("%d", sequence.Depth)
		res.valueInt = sequence.Depth
		res.GongFieldValueType = GongFieldValueTypeInt
	case "MinOccurs":
		res.valueString = sequence.MinOccurs
	case "MaxOccurs":
		res.valueString = sequence.MaxOccurs
	}
	return
}

func (simplecontent *SimpleContent) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = simplecontent.Name
	case "Extension":
		res.GongFieldValueType = GongFieldValueTypePointer
		if simplecontent.Extension != nil {
			res.valueString = simplecontent.Extension.Name
			res.ids = simplecontent.Extension.GongGetUUID(stage)
		}
	case "Restriction":
		res.GongFieldValueType = GongFieldValueTypePointer
		if simplecontent.Restriction != nil {
			res.valueString = simplecontent.Restriction.Name
			res.ids = simplecontent.Restriction.GongGetUUID(stage)
		}
	}
	return
}

func (simpletype *SimpleType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = simpletype.Name
	case "Annotation":
		res.GongFieldValueType = GongFieldValueTypePointer
		if simpletype.Annotation != nil {
			res.valueString = simpletype.Annotation.Name
			res.ids = simpletype.Annotation.GongGetUUID(stage)
		}
	case "NameXSD":
		res.valueString = simpletype.NameXSD
	case "Restriction":
		res.GongFieldValueType = GongFieldValueTypePointer
		if simpletype.Restriction != nil {
			res.valueString = simpletype.Restriction.Name
			res.ids = simpletype.Restriction.GongGetUUID(stage)
		}
	case "Union":
		res.GongFieldValueType = GongFieldValueTypePointer
		if simpletype.Union != nil {
			res.valueString = simpletype.Union.Name
			res.ids = simpletype.Union.GongGetUUID(stage)
		}
	case "Order":
		res.valueString = fmt.Sprintf("%d", simpletype.Order)
		res.valueInt = simpletype.Order
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Depth":
		res.valueString = fmt.Sprintf("%d", simpletype.Depth)
		res.valueInt = simpletype.Depth
		res.GongFieldValueType = GongFieldValueTypeInt
	}
	return
}

func (totaldigit *TotalDigit) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = totaldigit.Name
	case "Annotation":
		res.GongFieldValueType = GongFieldValueTypePointer
		if totaldigit.Annotation != nil {
			res.valueString = totaldigit.Annotation.Name
			res.ids = totaldigit.Annotation.GongGetUUID(stage)
		}
	case "Value":
		res.valueString = totaldigit.Value
	}
	return
}

func (union *Union) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = union.Name
	case "Annotation":
		res.GongFieldValueType = GongFieldValueTypePointer
		if union.Annotation != nil {
			res.valueString = union.Annotation.Name
			res.ids = union.Annotation.GongGetUUID(stage)
		}
	case "MemberTypes":
		res.valueString = union.MemberTypes
	}
	return
}

func (whitespace *WhiteSpace) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = whitespace.Name
	case "Annotation":
		res.GongFieldValueType = GongFieldValueTypePointer
		if whitespace.Annotation != nil {
			res.valueString = whitespace.Annotation.Name
			res.ids = whitespace.Annotation.GongGetUUID(stage)
		}
	case "Value":
		res.valueString = whitespace.Value
	}
	return
}

func GetFieldStringValueFromPointer(instance GongstructIF, fieldName string, stage *Stage) (res GongFieldValue) {
	res = instance.GongGetFieldValue(fieldName, stage)
	return
}

// insertion point for generic set gongstruct field value
func (all *All) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		all.Name = value.GetValueString()
	case "Annotation":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			all.Annotation = nil
			for __instance__ := range stage.Annotations {
				if stage.Annotation_stagedOrder[__instance__] == uint(id) {
					all.Annotation = __instance__
					break
				}
			}
		}
	case "OuterElementName":
		all.OuterElementName = value.GetValueString()
	case "Sequences":
		all.Sequences = make([]*Sequence, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Sequences {
					if stage.Sequence_stagedOrder[__instance__] == uint(id) {
						all.Sequences = append(all.Sequences, __instance__)
						break
					}
				}
			}
		}
	case "Alls":
		all.Alls = make([]*All, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Alls {
					if stage.All_stagedOrder[__instance__] == uint(id) {
						all.Alls = append(all.Alls, __instance__)
						break
					}
				}
			}
		}
	case "Choices":
		all.Choices = make([]*Choice, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Choices {
					if stage.Choice_stagedOrder[__instance__] == uint(id) {
						all.Choices = append(all.Choices, __instance__)
						break
					}
				}
			}
		}
	case "Groups":
		all.Groups = make([]*Group, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Groups {
					if stage.Group_stagedOrder[__instance__] == uint(id) {
						all.Groups = append(all.Groups, __instance__)
						break
					}
				}
			}
		}
	case "Elements":
		all.Elements = make([]*Element, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Elements {
					if stage.Element_stagedOrder[__instance__] == uint(id) {
						all.Elements = append(all.Elements, __instance__)
						break
					}
				}
			}
		}
	case "Order":
		all.Order = int(value.GetValueInt())
	case "Depth":
		all.Depth = int(value.GetValueInt())
	case "MinOccurs":
		all.MinOccurs = value.GetValueString()
	case "MaxOccurs":
		all.MaxOccurs = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (annotation *Annotation) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		annotation.Name = value.GetValueString()
	case "Documentations":
		annotation.Documentations = make([]*Documentation, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Documentations {
					if stage.Documentation_stagedOrder[__instance__] == uint(id) {
						annotation.Documentations = append(annotation.Documentations, __instance__)
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

func (attribute *Attribute) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		attribute.Name = value.GetValueString()
	case "NameXSD":
		attribute.NameXSD = value.GetValueString()
	case "Type":
		attribute.Type = value.GetValueString()
	case "Annotation":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			attribute.Annotation = nil
			for __instance__ := range stage.Annotations {
				if stage.Annotation_stagedOrder[__instance__] == uint(id) {
					attribute.Annotation = __instance__
					break
				}
			}
		}
	case "HasNameConflict":
		attribute.HasNameConflict = value.GetValueBool()
	case "GoIdentifier":
		attribute.GoIdentifier = value.GetValueString()
	case "Default":
		attribute.Default = value.GetValueString()
	case "Use":
		attribute.Use = value.GetValueString()
	case "Form":
		attribute.Form = value.GetValueString()
	case "Fixed":
		attribute.Fixed = value.GetValueString()
	case "Ref":
		attribute.Ref = value.GetValueString()
	case "TargetNamespace":
		attribute.TargetNamespace = value.GetValueString()
	case "SimpleType":
		attribute.SimpleType = value.GetValueString()
	case "IDXSD":
		attribute.IDXSD = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (attributegroup *AttributeGroup) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		attributegroup.Name = value.GetValueString()
	case "NameXSD":
		attributegroup.NameXSD = value.GetValueString()
	case "Annotation":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			attributegroup.Annotation = nil
			for __instance__ := range stage.Annotations {
				if stage.Annotation_stagedOrder[__instance__] == uint(id) {
					attributegroup.Annotation = __instance__
					break
				}
			}
		}
	case "HasNameConflict":
		attributegroup.HasNameConflict = value.GetValueBool()
	case "GoIdentifier":
		attributegroup.GoIdentifier = value.GetValueString()
	case "AttributeGroups":
		attributegroup.AttributeGroups = make([]*AttributeGroup, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.AttributeGroups {
					if stage.AttributeGroup_stagedOrder[__instance__] == uint(id) {
						attributegroup.AttributeGroups = append(attributegroup.AttributeGroups, __instance__)
						break
					}
				}
			}
		}
	case "Ref":
		attributegroup.Ref = value.GetValueString()
	case "Attributes":
		attributegroup.Attributes = make([]*Attribute, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Attributes {
					if stage.Attribute_stagedOrder[__instance__] == uint(id) {
						attributegroup.Attributes = append(attributegroup.Attributes, __instance__)
						break
					}
				}
			}
		}
	case "Order":
		attributegroup.Order = int(value.GetValueInt())
	case "Depth":
		attributegroup.Depth = int(value.GetValueInt())
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (choice *Choice) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		choice.Name = value.GetValueString()
	case "Annotation":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			choice.Annotation = nil
			for __instance__ := range stage.Annotations {
				if stage.Annotation_stagedOrder[__instance__] == uint(id) {
					choice.Annotation = __instance__
					break
				}
			}
		}
	case "OuterElementName":
		choice.OuterElementName = value.GetValueString()
	case "Sequences":
		choice.Sequences = make([]*Sequence, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Sequences {
					if stage.Sequence_stagedOrder[__instance__] == uint(id) {
						choice.Sequences = append(choice.Sequences, __instance__)
						break
					}
				}
			}
		}
	case "Alls":
		choice.Alls = make([]*All, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Alls {
					if stage.All_stagedOrder[__instance__] == uint(id) {
						choice.Alls = append(choice.Alls, __instance__)
						break
					}
				}
			}
		}
	case "Choices":
		choice.Choices = make([]*Choice, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Choices {
					if stage.Choice_stagedOrder[__instance__] == uint(id) {
						choice.Choices = append(choice.Choices, __instance__)
						break
					}
				}
			}
		}
	case "Groups":
		choice.Groups = make([]*Group, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Groups {
					if stage.Group_stagedOrder[__instance__] == uint(id) {
						choice.Groups = append(choice.Groups, __instance__)
						break
					}
				}
			}
		}
	case "Elements":
		choice.Elements = make([]*Element, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Elements {
					if stage.Element_stagedOrder[__instance__] == uint(id) {
						choice.Elements = append(choice.Elements, __instance__)
						break
					}
				}
			}
		}
	case "Order":
		choice.Order = int(value.GetValueInt())
	case "Depth":
		choice.Depth = int(value.GetValueInt())
	case "MinOccurs":
		choice.MinOccurs = value.GetValueString()
	case "MaxOccurs":
		choice.MaxOccurs = value.GetValueString()
	case "IsDuplicatedInXSD":
		choice.IsDuplicatedInXSD = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (complexcontent *ComplexContent) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		complexcontent.Name = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (complextype *ComplexType) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		complextype.Name = value.GetValueString()
	case "HasNameConflict":
		complextype.HasNameConflict = value.GetValueBool()
	case "GoIdentifier":
		complextype.GoIdentifier = value.GetValueString()
	case "IsAnonymous":
		complextype.IsAnonymous = value.GetValueBool()
	case "OuterElement":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			complextype.OuterElement = nil
			for __instance__ := range stage.Elements {
				if stage.Element_stagedOrder[__instance__] == uint(id) {
					complextype.OuterElement = __instance__
					break
				}
			}
		}
	case "Annotation":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			complextype.Annotation = nil
			for __instance__ := range stage.Annotations {
				if stage.Annotation_stagedOrder[__instance__] == uint(id) {
					complextype.Annotation = __instance__
					break
				}
			}
		}
	case "NameXSD":
		complextype.NameXSD = value.GetValueString()
	case "OuterElementName":
		complextype.OuterElementName = value.GetValueString()
	case "Sequences":
		complextype.Sequences = make([]*Sequence, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Sequences {
					if stage.Sequence_stagedOrder[__instance__] == uint(id) {
						complextype.Sequences = append(complextype.Sequences, __instance__)
						break
					}
				}
			}
		}
	case "Alls":
		complextype.Alls = make([]*All, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Alls {
					if stage.All_stagedOrder[__instance__] == uint(id) {
						complextype.Alls = append(complextype.Alls, __instance__)
						break
					}
				}
			}
		}
	case "Choices":
		complextype.Choices = make([]*Choice, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Choices {
					if stage.Choice_stagedOrder[__instance__] == uint(id) {
						complextype.Choices = append(complextype.Choices, __instance__)
						break
					}
				}
			}
		}
	case "Groups":
		complextype.Groups = make([]*Group, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Groups {
					if stage.Group_stagedOrder[__instance__] == uint(id) {
						complextype.Groups = append(complextype.Groups, __instance__)
						break
					}
				}
			}
		}
	case "Elements":
		complextype.Elements = make([]*Element, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Elements {
					if stage.Element_stagedOrder[__instance__] == uint(id) {
						complextype.Elements = append(complextype.Elements, __instance__)
						break
					}
				}
			}
		}
	case "Order":
		complextype.Order = int(value.GetValueInt())
	case "Depth":
		complextype.Depth = int(value.GetValueInt())
	case "MinOccurs":
		complextype.MinOccurs = value.GetValueString()
	case "MaxOccurs":
		complextype.MaxOccurs = value.GetValueString()
	case "Extension":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			complextype.Extension = nil
			for __instance__ := range stage.Extensions {
				if stage.Extension_stagedOrder[__instance__] == uint(id) {
					complextype.Extension = __instance__
					break
				}
			}
		}
	case "SimpleContent":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			complextype.SimpleContent = nil
			for __instance__ := range stage.SimpleContents {
				if stage.SimpleContent_stagedOrder[__instance__] == uint(id) {
					complextype.SimpleContent = __instance__
					break
				}
			}
		}
	case "ComplexContent":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			complextype.ComplexContent = nil
			for __instance__ := range stage.ComplexContents {
				if stage.ComplexContent_stagedOrder[__instance__] == uint(id) {
					complextype.ComplexContent = __instance__
					break
				}
			}
		}
	case "Attributes":
		complextype.Attributes = make([]*Attribute, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Attributes {
					if stage.Attribute_stagedOrder[__instance__] == uint(id) {
						complextype.Attributes = append(complextype.Attributes, __instance__)
						break
					}
				}
			}
		}
	case "AttributeGroups":
		complextype.AttributeGroups = make([]*AttributeGroup, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.AttributeGroups {
					if stage.AttributeGroup_stagedOrder[__instance__] == uint(id) {
						complextype.AttributeGroups = append(complextype.AttributeGroups, __instance__)
						break
					}
				}
			}
		}
	case "IsDuplicatedInXSD":
		complextype.IsDuplicatedInXSD = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (documentation *Documentation) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		documentation.Name = value.GetValueString()
	case "Text":
		documentation.Text = value.GetValueString()
	case "Source":
		documentation.Source = value.GetValueString()
	case "Lang":
		documentation.Lang = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (element *Element) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		element.Name = value.GetValueString()
	case "Order":
		element.Order = int(value.GetValueInt())
	case "Depth":
		element.Depth = int(value.GetValueInt())
	case "HasNameConflict":
		element.HasNameConflict = value.GetValueBool()
	case "GoIdentifier":
		element.GoIdentifier = value.GetValueString()
	case "Annotation":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			element.Annotation = nil
			for __instance__ := range stage.Annotations {
				if stage.Annotation_stagedOrder[__instance__] == uint(id) {
					element.Annotation = __instance__
					break
				}
			}
		}
	case "NameXSD":
		element.NameXSD = value.GetValueString()
	case "Type":
		element.Type = value.GetValueString()
	case "MinOccurs":
		element.MinOccurs = value.GetValueString()
	case "MaxOccurs":
		element.MaxOccurs = value.GetValueString()
	case "Default":
		element.Default = value.GetValueString()
	case "Fixed":
		element.Fixed = value.GetValueString()
	case "Nillable":
		element.Nillable = value.GetValueString()
	case "Ref":
		element.Ref = value.GetValueString()
	case "Abstract":
		element.Abstract = value.GetValueString()
	case "Form":
		element.Form = value.GetValueString()
	case "Block":
		element.Block = value.GetValueString()
	case "Final":
		element.Final = value.GetValueString()
	case "SimpleType":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			element.SimpleType = nil
			for __instance__ := range stage.SimpleTypes {
				if stage.SimpleType_stagedOrder[__instance__] == uint(id) {
					element.SimpleType = __instance__
					break
				}
			}
		}
	case "ComplexType":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			element.ComplexType = nil
			for __instance__ := range stage.ComplexTypes {
				if stage.ComplexType_stagedOrder[__instance__] == uint(id) {
					element.ComplexType = __instance__
					break
				}
			}
		}
	case "Groups":
		element.Groups = make([]*Group, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Groups {
					if stage.Group_stagedOrder[__instance__] == uint(id) {
						element.Groups = append(element.Groups, __instance__)
						break
					}
				}
			}
		}
	case "IsDuplicatedInXSD":
		element.IsDuplicatedInXSD = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (enumeration *Enumeration) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		enumeration.Name = value.GetValueString()
	case "Annotation":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			enumeration.Annotation = nil
			for __instance__ := range stage.Annotations {
				if stage.Annotation_stagedOrder[__instance__] == uint(id) {
					enumeration.Annotation = __instance__
					break
				}
			}
		}
	case "Value":
		enumeration.Value = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (extension *Extension) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		extension.Name = value.GetValueString()
	case "OuterElementName":
		extension.OuterElementName = value.GetValueString()
	case "Sequences":
		extension.Sequences = make([]*Sequence, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Sequences {
					if stage.Sequence_stagedOrder[__instance__] == uint(id) {
						extension.Sequences = append(extension.Sequences, __instance__)
						break
					}
				}
			}
		}
	case "Alls":
		extension.Alls = make([]*All, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Alls {
					if stage.All_stagedOrder[__instance__] == uint(id) {
						extension.Alls = append(extension.Alls, __instance__)
						break
					}
				}
			}
		}
	case "Choices":
		extension.Choices = make([]*Choice, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Choices {
					if stage.Choice_stagedOrder[__instance__] == uint(id) {
						extension.Choices = append(extension.Choices, __instance__)
						break
					}
				}
			}
		}
	case "Groups":
		extension.Groups = make([]*Group, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Groups {
					if stage.Group_stagedOrder[__instance__] == uint(id) {
						extension.Groups = append(extension.Groups, __instance__)
						break
					}
				}
			}
		}
	case "Elements":
		extension.Elements = make([]*Element, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Elements {
					if stage.Element_stagedOrder[__instance__] == uint(id) {
						extension.Elements = append(extension.Elements, __instance__)
						break
					}
				}
			}
		}
	case "Order":
		extension.Order = int(value.GetValueInt())
	case "Depth":
		extension.Depth = int(value.GetValueInt())
	case "MinOccurs":
		extension.MinOccurs = value.GetValueString()
	case "MaxOccurs":
		extension.MaxOccurs = value.GetValueString()
	case "Base":
		extension.Base = value.GetValueString()
	case "Ref":
		extension.Ref = value.GetValueString()
	case "Attributes":
		extension.Attributes = make([]*Attribute, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Attributes {
					if stage.Attribute_stagedOrder[__instance__] == uint(id) {
						extension.Attributes = append(extension.Attributes, __instance__)
						break
					}
				}
			}
		}
	case "AttributeGroups":
		extension.AttributeGroups = make([]*AttributeGroup, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.AttributeGroups {
					if stage.AttributeGroup_stagedOrder[__instance__] == uint(id) {
						extension.AttributeGroups = append(extension.AttributeGroups, __instance__)
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
	case "Annotation":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			group.Annotation = nil
			for __instance__ := range stage.Annotations {
				if stage.Annotation_stagedOrder[__instance__] == uint(id) {
					group.Annotation = __instance__
					break
				}
			}
		}
	case "NameXSD":
		group.NameXSD = value.GetValueString()
	case "Ref":
		group.Ref = value.GetValueString()
	case "IsAnonymous":
		group.IsAnonymous = value.GetValueBool()
	case "OuterElement":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			group.OuterElement = nil
			for __instance__ := range stage.Elements {
				if stage.Element_stagedOrder[__instance__] == uint(id) {
					group.OuterElement = __instance__
					break
				}
			}
		}
	case "HasNameConflict":
		group.HasNameConflict = value.GetValueBool()
	case "GoIdentifier":
		group.GoIdentifier = value.GetValueString()
	case "OuterElementName":
		group.OuterElementName = value.GetValueString()
	case "Sequences":
		group.Sequences = make([]*Sequence, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Sequences {
					if stage.Sequence_stagedOrder[__instance__] == uint(id) {
						group.Sequences = append(group.Sequences, __instance__)
						break
					}
				}
			}
		}
	case "Alls":
		group.Alls = make([]*All, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Alls {
					if stage.All_stagedOrder[__instance__] == uint(id) {
						group.Alls = append(group.Alls, __instance__)
						break
					}
				}
			}
		}
	case "Choices":
		group.Choices = make([]*Choice, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Choices {
					if stage.Choice_stagedOrder[__instance__] == uint(id) {
						group.Choices = append(group.Choices, __instance__)
						break
					}
				}
			}
		}
	case "Groups":
		group.Groups = make([]*Group, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Groups {
					if stage.Group_stagedOrder[__instance__] == uint(id) {
						group.Groups = append(group.Groups, __instance__)
						break
					}
				}
			}
		}
	case "Elements":
		group.Elements = make([]*Element, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Elements {
					if stage.Element_stagedOrder[__instance__] == uint(id) {
						group.Elements = append(group.Elements, __instance__)
						break
					}
				}
			}
		}
	case "Order":
		group.Order = int(value.GetValueInt())
	case "Depth":
		group.Depth = int(value.GetValueInt())
	case "MinOccurs":
		group.MinOccurs = value.GetValueString()
	case "MaxOccurs":
		group.MaxOccurs = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (length *Length) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		length.Name = value.GetValueString()
	case "Annotation":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			length.Annotation = nil
			for __instance__ := range stage.Annotations {
				if stage.Annotation_stagedOrder[__instance__] == uint(id) {
					length.Annotation = __instance__
					break
				}
			}
		}
	case "Value":
		length.Value = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (maxinclusive *MaxInclusive) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		maxinclusive.Name = value.GetValueString()
	case "Annotation":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			maxinclusive.Annotation = nil
			for __instance__ := range stage.Annotations {
				if stage.Annotation_stagedOrder[__instance__] == uint(id) {
					maxinclusive.Annotation = __instance__
					break
				}
			}
		}
	case "Value":
		maxinclusive.Value = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (maxlength *MaxLength) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		maxlength.Name = value.GetValueString()
	case "Annotation":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			maxlength.Annotation = nil
			for __instance__ := range stage.Annotations {
				if stage.Annotation_stagedOrder[__instance__] == uint(id) {
					maxlength.Annotation = __instance__
					break
				}
			}
		}
	case "Value":
		maxlength.Value = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (mininclusive *MinInclusive) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		mininclusive.Name = value.GetValueString()
	case "Annotation":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			mininclusive.Annotation = nil
			for __instance__ := range stage.Annotations {
				if stage.Annotation_stagedOrder[__instance__] == uint(id) {
					mininclusive.Annotation = __instance__
					break
				}
			}
		}
	case "Value":
		mininclusive.Value = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (minlength *MinLength) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		minlength.Name = value.GetValueString()
	case "Annotation":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			minlength.Annotation = nil
			for __instance__ := range stage.Annotations {
				if stage.Annotation_stagedOrder[__instance__] == uint(id) {
					minlength.Annotation = __instance__
					break
				}
			}
		}
	case "Value":
		minlength.Value = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (pattern *Pattern) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		pattern.Name = value.GetValueString()
	case "Annotation":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			pattern.Annotation = nil
			for __instance__ := range stage.Annotations {
				if stage.Annotation_stagedOrder[__instance__] == uint(id) {
					pattern.Annotation = __instance__
					break
				}
			}
		}
	case "Value":
		pattern.Value = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (restriction *Restriction) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		restriction.Name = value.GetValueString()
	case "Annotation":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			restriction.Annotation = nil
			for __instance__ := range stage.Annotations {
				if stage.Annotation_stagedOrder[__instance__] == uint(id) {
					restriction.Annotation = __instance__
					break
				}
			}
		}
	case "Base":
		restriction.Base = value.GetValueString()
	case "Enumerations":
		restriction.Enumerations = make([]*Enumeration, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Enumerations {
					if stage.Enumeration_stagedOrder[__instance__] == uint(id) {
						restriction.Enumerations = append(restriction.Enumerations, __instance__)
						break
					}
				}
			}
		}
	case "MinInclusive":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			restriction.MinInclusive = nil
			for __instance__ := range stage.MinInclusives {
				if stage.MinInclusive_stagedOrder[__instance__] == uint(id) {
					restriction.MinInclusive = __instance__
					break
				}
			}
		}
	case "MaxInclusive":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			restriction.MaxInclusive = nil
			for __instance__ := range stage.MaxInclusives {
				if stage.MaxInclusive_stagedOrder[__instance__] == uint(id) {
					restriction.MaxInclusive = __instance__
					break
				}
			}
		}
	case "Pattern":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			restriction.Pattern = nil
			for __instance__ := range stage.Patterns {
				if stage.Pattern_stagedOrder[__instance__] == uint(id) {
					restriction.Pattern = __instance__
					break
				}
			}
		}
	case "WhiteSpace":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			restriction.WhiteSpace = nil
			for __instance__ := range stage.WhiteSpaces {
				if stage.WhiteSpace_stagedOrder[__instance__] == uint(id) {
					restriction.WhiteSpace = __instance__
					break
				}
			}
		}
	case "MinLength":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			restriction.MinLength = nil
			for __instance__ := range stage.MinLengths {
				if stage.MinLength_stagedOrder[__instance__] == uint(id) {
					restriction.MinLength = __instance__
					break
				}
			}
		}
	case "MaxLength":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			restriction.MaxLength = nil
			for __instance__ := range stage.MaxLengths {
				if stage.MaxLength_stagedOrder[__instance__] == uint(id) {
					restriction.MaxLength = __instance__
					break
				}
			}
		}
	case "Length":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			restriction.Length = nil
			for __instance__ := range stage.Lengths {
				if stage.Length_stagedOrder[__instance__] == uint(id) {
					restriction.Length = __instance__
					break
				}
			}
		}
	case "TotalDigit":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			restriction.TotalDigit = nil
			for __instance__ := range stage.TotalDigits {
				if stage.TotalDigit_stagedOrder[__instance__] == uint(id) {
					restriction.TotalDigit = __instance__
					break
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (schema *Schema) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		schema.Name = value.GetValueString()
	case "Xs":
		schema.Xs = value.GetValueString()
	case "Annotation":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			schema.Annotation = nil
			for __instance__ := range stage.Annotations {
				if stage.Annotation_stagedOrder[__instance__] == uint(id) {
					schema.Annotation = __instance__
					break
				}
			}
		}
	case "Elements":
		schema.Elements = make([]*Element, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Elements {
					if stage.Element_stagedOrder[__instance__] == uint(id) {
						schema.Elements = append(schema.Elements, __instance__)
						break
					}
				}
			}
		}
	case "SimpleTypes":
		schema.SimpleTypes = make([]*SimpleType, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.SimpleTypes {
					if stage.SimpleType_stagedOrder[__instance__] == uint(id) {
						schema.SimpleTypes = append(schema.SimpleTypes, __instance__)
						break
					}
				}
			}
		}
	case "ComplexTypes":
		schema.ComplexTypes = make([]*ComplexType, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.ComplexTypes {
					if stage.ComplexType_stagedOrder[__instance__] == uint(id) {
						schema.ComplexTypes = append(schema.ComplexTypes, __instance__)
						break
					}
				}
			}
		}
	case "AttributeGroups":
		schema.AttributeGroups = make([]*AttributeGroup, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.AttributeGroups {
					if stage.AttributeGroup_stagedOrder[__instance__] == uint(id) {
						schema.AttributeGroups = append(schema.AttributeGroups, __instance__)
						break
					}
				}
			}
		}
	case "Groups":
		schema.Groups = make([]*Group, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Groups {
					if stage.Group_stagedOrder[__instance__] == uint(id) {
						schema.Groups = append(schema.Groups, __instance__)
						break
					}
				}
			}
		}
	case "Order":
		schema.Order = int(value.GetValueInt())
	case "Depth":
		schema.Depth = int(value.GetValueInt())
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (sequence *Sequence) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		sequence.Name = value.GetValueString()
	case "Annotation":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			sequence.Annotation = nil
			for __instance__ := range stage.Annotations {
				if stage.Annotation_stagedOrder[__instance__] == uint(id) {
					sequence.Annotation = __instance__
					break
				}
			}
		}
	case "OuterElementName":
		sequence.OuterElementName = value.GetValueString()
	case "Sequences":
		sequence.Sequences = make([]*Sequence, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Sequences {
					if stage.Sequence_stagedOrder[__instance__] == uint(id) {
						sequence.Sequences = append(sequence.Sequences, __instance__)
						break
					}
				}
			}
		}
	case "Alls":
		sequence.Alls = make([]*All, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Alls {
					if stage.All_stagedOrder[__instance__] == uint(id) {
						sequence.Alls = append(sequence.Alls, __instance__)
						break
					}
				}
			}
		}
	case "Choices":
		sequence.Choices = make([]*Choice, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Choices {
					if stage.Choice_stagedOrder[__instance__] == uint(id) {
						sequence.Choices = append(sequence.Choices, __instance__)
						break
					}
				}
			}
		}
	case "Groups":
		sequence.Groups = make([]*Group, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Groups {
					if stage.Group_stagedOrder[__instance__] == uint(id) {
						sequence.Groups = append(sequence.Groups, __instance__)
						break
					}
				}
			}
		}
	case "Elements":
		sequence.Elements = make([]*Element, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Elements {
					if stage.Element_stagedOrder[__instance__] == uint(id) {
						sequence.Elements = append(sequence.Elements, __instance__)
						break
					}
				}
			}
		}
	case "Order":
		sequence.Order = int(value.GetValueInt())
	case "Depth":
		sequence.Depth = int(value.GetValueInt())
	case "MinOccurs":
		sequence.MinOccurs = value.GetValueString()
	case "MaxOccurs":
		sequence.MaxOccurs = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (simplecontent *SimpleContent) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		simplecontent.Name = value.GetValueString()
	case "Extension":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			simplecontent.Extension = nil
			for __instance__ := range stage.Extensions {
				if stage.Extension_stagedOrder[__instance__] == uint(id) {
					simplecontent.Extension = __instance__
					break
				}
			}
		}
	case "Restriction":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			simplecontent.Restriction = nil
			for __instance__ := range stage.Restrictions {
				if stage.Restriction_stagedOrder[__instance__] == uint(id) {
					simplecontent.Restriction = __instance__
					break
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (simpletype *SimpleType) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		simpletype.Name = value.GetValueString()
	case "Annotation":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			simpletype.Annotation = nil
			for __instance__ := range stage.Annotations {
				if stage.Annotation_stagedOrder[__instance__] == uint(id) {
					simpletype.Annotation = __instance__
					break
				}
			}
		}
	case "NameXSD":
		simpletype.NameXSD = value.GetValueString()
	case "Restriction":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			simpletype.Restriction = nil
			for __instance__ := range stage.Restrictions {
				if stage.Restriction_stagedOrder[__instance__] == uint(id) {
					simpletype.Restriction = __instance__
					break
				}
			}
		}
	case "Union":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			simpletype.Union = nil
			for __instance__ := range stage.Unions {
				if stage.Union_stagedOrder[__instance__] == uint(id) {
					simpletype.Union = __instance__
					break
				}
			}
		}
	case "Order":
		simpletype.Order = int(value.GetValueInt())
	case "Depth":
		simpletype.Depth = int(value.GetValueInt())
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (totaldigit *TotalDigit) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		totaldigit.Name = value.GetValueString()
	case "Annotation":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			totaldigit.Annotation = nil
			for __instance__ := range stage.Annotations {
				if stage.Annotation_stagedOrder[__instance__] == uint(id) {
					totaldigit.Annotation = __instance__
					break
				}
			}
		}
	case "Value":
		totaldigit.Value = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (union *Union) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		union.Name = value.GetValueString()
	case "Annotation":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			union.Annotation = nil
			for __instance__ := range stage.Annotations {
				if stage.Annotation_stagedOrder[__instance__] == uint(id) {
					union.Annotation = __instance__
					break
				}
			}
		}
	case "MemberTypes":
		union.MemberTypes = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (whitespace *WhiteSpace) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		whitespace.Name = value.GetValueString()
	case "Annotation":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			whitespace.Annotation = nil
			for __instance__ := range stage.Annotations {
				if stage.Annotation_stagedOrder[__instance__] == uint(id) {
					whitespace.Annotation = __instance__
					break
				}
			}
		}
	case "Value":
		whitespace.Value = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func SetFieldStringValueFromPointer(instance GongstructIF, fieldName string, value GongFieldValue, stage *Stage) error {
	return instance.GongSetFieldValue(fieldName, value, stage)
}

// insertion point for generic get gongstruct name
func (all *All) GongGetGongstructName() string {
	return "All"
}

func (annotation *Annotation) GongGetGongstructName() string {
	return "Annotation"
}

func (attribute *Attribute) GongGetGongstructName() string {
	return "Attribute"
}

func (attributegroup *AttributeGroup) GongGetGongstructName() string {
	return "AttributeGroup"
}

func (choice *Choice) GongGetGongstructName() string {
	return "Choice"
}

func (complexcontent *ComplexContent) GongGetGongstructName() string {
	return "ComplexContent"
}

func (complextype *ComplexType) GongGetGongstructName() string {
	return "ComplexType"
}

func (documentation *Documentation) GongGetGongstructName() string {
	return "Documentation"
}

func (element *Element) GongGetGongstructName() string {
	return "Element"
}

func (enumeration *Enumeration) GongGetGongstructName() string {
	return "Enumeration"
}

func (extension *Extension) GongGetGongstructName() string {
	return "Extension"
}

func (group *Group) GongGetGongstructName() string {
	return "Group"
}

func (length *Length) GongGetGongstructName() string {
	return "Length"
}

func (maxinclusive *MaxInclusive) GongGetGongstructName() string {
	return "MaxInclusive"
}

func (maxlength *MaxLength) GongGetGongstructName() string {
	return "MaxLength"
}

func (mininclusive *MinInclusive) GongGetGongstructName() string {
	return "MinInclusive"
}

func (minlength *MinLength) GongGetGongstructName() string {
	return "MinLength"
}

func (pattern *Pattern) GongGetGongstructName() string {
	return "Pattern"
}

func (restriction *Restriction) GongGetGongstructName() string {
	return "Restriction"
}

func (schema *Schema) GongGetGongstructName() string {
	return "Schema"
}

func (sequence *Sequence) GongGetGongstructName() string {
	return "Sequence"
}

func (simplecontent *SimpleContent) GongGetGongstructName() string {
	return "SimpleContent"
}

func (simpletype *SimpleType) GongGetGongstructName() string {
	return "SimpleType"
}

func (totaldigit *TotalDigit) GongGetGongstructName() string {
	return "TotalDigit"
}

func (union *Union) GongGetGongstructName() string {
	return "Union"
}

func (whitespace *WhiteSpace) GongGetGongstructName() string {
	return "WhiteSpace"
}

func GetGongstructNameFromPointer(instance GongstructIF) (res string) {
	res = instance.GongGetGongstructName()
	return
}

func (stage *Stage) ResetMapStrings() {
	// insertion point for generic get gongstruct name
	stage.Alls_mapString = make(map[string]*All)
	for all := range stage.Alls {
		stage.Alls_mapString[all.Name] = all
	}

	stage.Annotations_mapString = make(map[string]*Annotation)
	for annotation := range stage.Annotations {
		stage.Annotations_mapString[annotation.Name] = annotation
	}

	stage.Attributes_mapString = make(map[string]*Attribute)
	for attribute := range stage.Attributes {
		stage.Attributes_mapString[attribute.Name] = attribute
	}

	stage.AttributeGroups_mapString = make(map[string]*AttributeGroup)
	for attributegroup := range stage.AttributeGroups {
		stage.AttributeGroups_mapString[attributegroup.Name] = attributegroup
	}

	stage.Choices_mapString = make(map[string]*Choice)
	for choice := range stage.Choices {
		stage.Choices_mapString[choice.Name] = choice
	}

	stage.ComplexContents_mapString = make(map[string]*ComplexContent)
	for complexcontent := range stage.ComplexContents {
		stage.ComplexContents_mapString[complexcontent.Name] = complexcontent
	}

	stage.ComplexTypes_mapString = make(map[string]*ComplexType)
	for complextype := range stage.ComplexTypes {
		stage.ComplexTypes_mapString[complextype.Name] = complextype
	}

	stage.Documentations_mapString = make(map[string]*Documentation)
	for documentation := range stage.Documentations {
		stage.Documentations_mapString[documentation.Name] = documentation
	}

	stage.Elements_mapString = make(map[string]*Element)
	for element := range stage.Elements {
		stage.Elements_mapString[element.Name] = element
	}

	stage.Enumerations_mapString = make(map[string]*Enumeration)
	for enumeration := range stage.Enumerations {
		stage.Enumerations_mapString[enumeration.Name] = enumeration
	}

	stage.Extensions_mapString = make(map[string]*Extension)
	for extension := range stage.Extensions {
		stage.Extensions_mapString[extension.Name] = extension
	}

	stage.Groups_mapString = make(map[string]*Group)
	for group := range stage.Groups {
		stage.Groups_mapString[group.Name] = group
	}

	stage.Lengths_mapString = make(map[string]*Length)
	for length := range stage.Lengths {
		stage.Lengths_mapString[length.Name] = length
	}

	stage.MaxInclusives_mapString = make(map[string]*MaxInclusive)
	for maxinclusive := range stage.MaxInclusives {
		stage.MaxInclusives_mapString[maxinclusive.Name] = maxinclusive
	}

	stage.MaxLengths_mapString = make(map[string]*MaxLength)
	for maxlength := range stage.MaxLengths {
		stage.MaxLengths_mapString[maxlength.Name] = maxlength
	}

	stage.MinInclusives_mapString = make(map[string]*MinInclusive)
	for mininclusive := range stage.MinInclusives {
		stage.MinInclusives_mapString[mininclusive.Name] = mininclusive
	}

	stage.MinLengths_mapString = make(map[string]*MinLength)
	for minlength := range stage.MinLengths {
		stage.MinLengths_mapString[minlength.Name] = minlength
	}

	stage.Patterns_mapString = make(map[string]*Pattern)
	for pattern := range stage.Patterns {
		stage.Patterns_mapString[pattern.Name] = pattern
	}

	stage.Restrictions_mapString = make(map[string]*Restriction)
	for restriction := range stage.Restrictions {
		stage.Restrictions_mapString[restriction.Name] = restriction
	}

	stage.Schemas_mapString = make(map[string]*Schema)
	for schema := range stage.Schemas {
		stage.Schemas_mapString[schema.Name] = schema
	}

	stage.Sequences_mapString = make(map[string]*Sequence)
	for sequence := range stage.Sequences {
		stage.Sequences_mapString[sequence.Name] = sequence
	}

	stage.SimpleContents_mapString = make(map[string]*SimpleContent)
	for simplecontent := range stage.SimpleContents {
		stage.SimpleContents_mapString[simplecontent.Name] = simplecontent
	}

	stage.SimpleTypes_mapString = make(map[string]*SimpleType)
	for simpletype := range stage.SimpleTypes {
		stage.SimpleTypes_mapString[simpletype.Name] = simpletype
	}

	stage.TotalDigits_mapString = make(map[string]*TotalDigit)
	for totaldigit := range stage.TotalDigits {
		stage.TotalDigits_mapString[totaldigit.Name] = totaldigit
	}

	stage.Unions_mapString = make(map[string]*Union)
	for union := range stage.Unions {
		stage.Unions_mapString[union.Name] = union
	}

	stage.WhiteSpaces_mapString = make(map[string]*WhiteSpace)
	for whitespace := range stage.WhiteSpaces {
		stage.WhiteSpaces_mapString[whitespace.Name] = whitespace
	}

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
