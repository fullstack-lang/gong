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
	"sync"
	"time"
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

const (
	GongProbeTreeSidebarSuffix           = ":sidebar of the probe"
	GongProbeNavigationTreeSidebarSuffix = ":sidebar of the probe, navigation"
	GongProbeTableSuffix                 = ":table of the probe"
	GongProbeNotificationTableSuffix     = ":notification table of the probe"
	GongProbeFormSuffix                  = ":form of the probe"
	GongProbeSplitSuffix                 = ":probe of the probe"
	GongProbeLoadSuffix                  = ":load of the probe"
)

type GongMarshallingMode string

const (
	// the whole stage is generated at each marshall. This is the default
	GongMarshallingNormal GongMarshallingMode = "GongMarshallingNormal"

	// only the last commit is append to the marshall file
	GongMarshallingAppendCommit GongMarshallingMode = "GongMarshallingAppendCommit"
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

	OnAfterAllCreateCallback GongOnAfterCreateInterface[All]
	OnAfterAllUpdateCallback GongOnAfterUpdateInterface[All]
	OnAfterAllDeleteCallback GongOnAfterDeleteInterface[All]

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

	OnAfterAnnotationCreateCallback GongOnAfterCreateInterface[Annotation]
	OnAfterAnnotationUpdateCallback GongOnAfterUpdateInterface[Annotation]
	OnAfterAnnotationDeleteCallback GongOnAfterDeleteInterface[Annotation]

	Attributes                map[*Attribute]struct{}
	Attributes_instance       map[*Attribute]*Attribute
	Attributes_mapString      map[string]*Attribute
	AttributeOrder            uint
	Attribute_stagedOrder     map[*Attribute]uint
	Attribute_orderStaged     map[uint]*Attribute
	Attributes_reference      map[*Attribute]*Attribute
	Attributes_referenceOrder map[*Attribute]uint

	// insertion point for slice of pointers maps
	OnAfterAttributeCreateCallback GongOnAfterCreateInterface[Attribute]
	OnAfterAttributeUpdateCallback GongOnAfterUpdateInterface[Attribute]
	OnAfterAttributeDeleteCallback GongOnAfterDeleteInterface[Attribute]

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

	OnAfterAttributeGroupCreateCallback GongOnAfterCreateInterface[AttributeGroup]
	OnAfterAttributeGroupUpdateCallback GongOnAfterUpdateInterface[AttributeGroup]
	OnAfterAttributeGroupDeleteCallback GongOnAfterDeleteInterface[AttributeGroup]

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

	OnAfterChoiceCreateCallback GongOnAfterCreateInterface[Choice]
	OnAfterChoiceUpdateCallback GongOnAfterUpdateInterface[Choice]
	OnAfterChoiceDeleteCallback GongOnAfterDeleteInterface[Choice]

	ComplexContents                map[*ComplexContent]struct{}
	ComplexContents_instance       map[*ComplexContent]*ComplexContent
	ComplexContents_mapString      map[string]*ComplexContent
	ComplexContentOrder            uint
	ComplexContent_stagedOrder     map[*ComplexContent]uint
	ComplexContent_orderStaged     map[uint]*ComplexContent
	ComplexContents_reference      map[*ComplexContent]*ComplexContent
	ComplexContents_referenceOrder map[*ComplexContent]uint

	// insertion point for slice of pointers maps
	OnAfterComplexContentCreateCallback GongOnAfterCreateInterface[ComplexContent]
	OnAfterComplexContentUpdateCallback GongOnAfterUpdateInterface[ComplexContent]
	OnAfterComplexContentDeleteCallback GongOnAfterDeleteInterface[ComplexContent]

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

	OnAfterComplexTypeCreateCallback GongOnAfterCreateInterface[ComplexType]
	OnAfterComplexTypeUpdateCallback GongOnAfterUpdateInterface[ComplexType]
	OnAfterComplexTypeDeleteCallback GongOnAfterDeleteInterface[ComplexType]

	Documentations                map[*Documentation]struct{}
	Documentations_instance       map[*Documentation]*Documentation
	Documentations_mapString      map[string]*Documentation
	DocumentationOrder            uint
	Documentation_stagedOrder     map[*Documentation]uint
	Documentation_orderStaged     map[uint]*Documentation
	Documentations_reference      map[*Documentation]*Documentation
	Documentations_referenceOrder map[*Documentation]uint

	// insertion point for slice of pointers maps
	OnAfterDocumentationCreateCallback GongOnAfterCreateInterface[Documentation]
	OnAfterDocumentationUpdateCallback GongOnAfterUpdateInterface[Documentation]
	OnAfterDocumentationDeleteCallback GongOnAfterDeleteInterface[Documentation]

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

	OnAfterElementCreateCallback GongOnAfterCreateInterface[Element]
	OnAfterElementUpdateCallback GongOnAfterUpdateInterface[Element]
	OnAfterElementDeleteCallback GongOnAfterDeleteInterface[Element]

	Enumerations                map[*Enumeration]struct{}
	Enumerations_instance       map[*Enumeration]*Enumeration
	Enumerations_mapString      map[string]*Enumeration
	EnumerationOrder            uint
	Enumeration_stagedOrder     map[*Enumeration]uint
	Enumeration_orderStaged     map[uint]*Enumeration
	Enumerations_reference      map[*Enumeration]*Enumeration
	Enumerations_referenceOrder map[*Enumeration]uint

	// insertion point for slice of pointers maps
	OnAfterEnumerationCreateCallback GongOnAfterCreateInterface[Enumeration]
	OnAfterEnumerationUpdateCallback GongOnAfterUpdateInterface[Enumeration]
	OnAfterEnumerationDeleteCallback GongOnAfterDeleteInterface[Enumeration]

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

	OnAfterExtensionCreateCallback GongOnAfterCreateInterface[Extension]
	OnAfterExtensionUpdateCallback GongOnAfterUpdateInterface[Extension]
	OnAfterExtensionDeleteCallback GongOnAfterDeleteInterface[Extension]

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

	OnAfterGroupCreateCallback GongOnAfterCreateInterface[Group]
	OnAfterGroupUpdateCallback GongOnAfterUpdateInterface[Group]
	OnAfterGroupDeleteCallback GongOnAfterDeleteInterface[Group]

	Lengths                map[*Length]struct{}
	Lengths_instance       map[*Length]*Length
	Lengths_mapString      map[string]*Length
	LengthOrder            uint
	Length_stagedOrder     map[*Length]uint
	Length_orderStaged     map[uint]*Length
	Lengths_reference      map[*Length]*Length
	Lengths_referenceOrder map[*Length]uint

	// insertion point for slice of pointers maps
	OnAfterLengthCreateCallback GongOnAfterCreateInterface[Length]
	OnAfterLengthUpdateCallback GongOnAfterUpdateInterface[Length]
	OnAfterLengthDeleteCallback GongOnAfterDeleteInterface[Length]

	MaxInclusives                map[*MaxInclusive]struct{}
	MaxInclusives_instance       map[*MaxInclusive]*MaxInclusive
	MaxInclusives_mapString      map[string]*MaxInclusive
	MaxInclusiveOrder            uint
	MaxInclusive_stagedOrder     map[*MaxInclusive]uint
	MaxInclusive_orderStaged     map[uint]*MaxInclusive
	MaxInclusives_reference      map[*MaxInclusive]*MaxInclusive
	MaxInclusives_referenceOrder map[*MaxInclusive]uint

	// insertion point for slice of pointers maps
	OnAfterMaxInclusiveCreateCallback GongOnAfterCreateInterface[MaxInclusive]
	OnAfterMaxInclusiveUpdateCallback GongOnAfterUpdateInterface[MaxInclusive]
	OnAfterMaxInclusiveDeleteCallback GongOnAfterDeleteInterface[MaxInclusive]

	MaxLengths                map[*MaxLength]struct{}
	MaxLengths_instance       map[*MaxLength]*MaxLength
	MaxLengths_mapString      map[string]*MaxLength
	MaxLengthOrder            uint
	MaxLength_stagedOrder     map[*MaxLength]uint
	MaxLength_orderStaged     map[uint]*MaxLength
	MaxLengths_reference      map[*MaxLength]*MaxLength
	MaxLengths_referenceOrder map[*MaxLength]uint

	// insertion point for slice of pointers maps
	OnAfterMaxLengthCreateCallback GongOnAfterCreateInterface[MaxLength]
	OnAfterMaxLengthUpdateCallback GongOnAfterUpdateInterface[MaxLength]
	OnAfterMaxLengthDeleteCallback GongOnAfterDeleteInterface[MaxLength]

	MinInclusives                map[*MinInclusive]struct{}
	MinInclusives_instance       map[*MinInclusive]*MinInclusive
	MinInclusives_mapString      map[string]*MinInclusive
	MinInclusiveOrder            uint
	MinInclusive_stagedOrder     map[*MinInclusive]uint
	MinInclusive_orderStaged     map[uint]*MinInclusive
	MinInclusives_reference      map[*MinInclusive]*MinInclusive
	MinInclusives_referenceOrder map[*MinInclusive]uint

	// insertion point for slice of pointers maps
	OnAfterMinInclusiveCreateCallback GongOnAfterCreateInterface[MinInclusive]
	OnAfterMinInclusiveUpdateCallback GongOnAfterUpdateInterface[MinInclusive]
	OnAfterMinInclusiveDeleteCallback GongOnAfterDeleteInterface[MinInclusive]

	MinLengths                map[*MinLength]struct{}
	MinLengths_instance       map[*MinLength]*MinLength
	MinLengths_mapString      map[string]*MinLength
	MinLengthOrder            uint
	MinLength_stagedOrder     map[*MinLength]uint
	MinLength_orderStaged     map[uint]*MinLength
	MinLengths_reference      map[*MinLength]*MinLength
	MinLengths_referenceOrder map[*MinLength]uint

	// insertion point for slice of pointers maps
	OnAfterMinLengthCreateCallback GongOnAfterCreateInterface[MinLength]
	OnAfterMinLengthUpdateCallback GongOnAfterUpdateInterface[MinLength]
	OnAfterMinLengthDeleteCallback GongOnAfterDeleteInterface[MinLength]

	Patterns                map[*Pattern]struct{}
	Patterns_instance       map[*Pattern]*Pattern
	Patterns_mapString      map[string]*Pattern
	PatternOrder            uint
	Pattern_stagedOrder     map[*Pattern]uint
	Pattern_orderStaged     map[uint]*Pattern
	Patterns_reference      map[*Pattern]*Pattern
	Patterns_referenceOrder map[*Pattern]uint

	// insertion point for slice of pointers maps
	OnAfterPatternCreateCallback GongOnAfterCreateInterface[Pattern]
	OnAfterPatternUpdateCallback GongOnAfterUpdateInterface[Pattern]
	OnAfterPatternDeleteCallback GongOnAfterDeleteInterface[Pattern]

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

	OnAfterRestrictionCreateCallback GongOnAfterCreateInterface[Restriction]
	OnAfterRestrictionUpdateCallback GongOnAfterUpdateInterface[Restriction]
	OnAfterRestrictionDeleteCallback GongOnAfterDeleteInterface[Restriction]

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

	OnAfterSchemaCreateCallback GongOnAfterCreateInterface[Schema]
	OnAfterSchemaUpdateCallback GongOnAfterUpdateInterface[Schema]
	OnAfterSchemaDeleteCallback GongOnAfterDeleteInterface[Schema]

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

	OnAfterSequenceCreateCallback GongOnAfterCreateInterface[Sequence]
	OnAfterSequenceUpdateCallback GongOnAfterUpdateInterface[Sequence]
	OnAfterSequenceDeleteCallback GongOnAfterDeleteInterface[Sequence]

	SimpleContents                map[*SimpleContent]struct{}
	SimpleContents_instance       map[*SimpleContent]*SimpleContent
	SimpleContents_mapString      map[string]*SimpleContent
	SimpleContentOrder            uint
	SimpleContent_stagedOrder     map[*SimpleContent]uint
	SimpleContent_orderStaged     map[uint]*SimpleContent
	SimpleContents_reference      map[*SimpleContent]*SimpleContent
	SimpleContents_referenceOrder map[*SimpleContent]uint

	// insertion point for slice of pointers maps
	OnAfterSimpleContentCreateCallback GongOnAfterCreateInterface[SimpleContent]
	OnAfterSimpleContentUpdateCallback GongOnAfterUpdateInterface[SimpleContent]
	OnAfterSimpleContentDeleteCallback GongOnAfterDeleteInterface[SimpleContent]

	SimpleTypes                map[*SimpleType]struct{}
	SimpleTypes_instance       map[*SimpleType]*SimpleType
	SimpleTypes_mapString      map[string]*SimpleType
	SimpleTypeOrder            uint
	SimpleType_stagedOrder     map[*SimpleType]uint
	SimpleType_orderStaged     map[uint]*SimpleType
	SimpleTypes_reference      map[*SimpleType]*SimpleType
	SimpleTypes_referenceOrder map[*SimpleType]uint

	// insertion point for slice of pointers maps
	OnAfterSimpleTypeCreateCallback GongOnAfterCreateInterface[SimpleType]
	OnAfterSimpleTypeUpdateCallback GongOnAfterUpdateInterface[SimpleType]
	OnAfterSimpleTypeDeleteCallback GongOnAfterDeleteInterface[SimpleType]

	TotalDigits                map[*TotalDigit]struct{}
	TotalDigits_instance       map[*TotalDigit]*TotalDigit
	TotalDigits_mapString      map[string]*TotalDigit
	TotalDigitOrder            uint
	TotalDigit_stagedOrder     map[*TotalDigit]uint
	TotalDigit_orderStaged     map[uint]*TotalDigit
	TotalDigits_reference      map[*TotalDigit]*TotalDigit
	TotalDigits_referenceOrder map[*TotalDigit]uint

	// insertion point for slice of pointers maps
	OnAfterTotalDigitCreateCallback GongOnAfterCreateInterface[TotalDigit]
	OnAfterTotalDigitUpdateCallback GongOnAfterUpdateInterface[TotalDigit]
	OnAfterTotalDigitDeleteCallback GongOnAfterDeleteInterface[TotalDigit]

	Unions                map[*Union]struct{}
	Unions_instance       map[*Union]*Union
	Unions_mapString      map[string]*Union
	UnionOrder            uint
	Union_stagedOrder     map[*Union]uint
	Union_orderStaged     map[uint]*Union
	Unions_reference      map[*Union]*Union
	Unions_referenceOrder map[*Union]uint

	// insertion point for slice of pointers maps
	OnAfterUnionCreateCallback GongOnAfterCreateInterface[Union]
	OnAfterUnionUpdateCallback GongOnAfterUpdateInterface[Union]
	OnAfterUnionDeleteCallback GongOnAfterDeleteInterface[Union]

	WhiteSpaces                map[*WhiteSpace]struct{}
	WhiteSpaces_instance       map[*WhiteSpace]*WhiteSpace
	WhiteSpaces_mapString      map[string]*WhiteSpace
	WhiteSpaceOrder            uint
	WhiteSpace_stagedOrder     map[*WhiteSpace]uint
	WhiteSpace_orderStaged     map[uint]*WhiteSpace
	WhiteSpaces_reference      map[*WhiteSpace]*WhiteSpace
	WhiteSpaces_referenceOrder map[*WhiteSpace]uint

	// insertion point for slice of pointers maps
	OnAfterWhiteSpaceCreateCallback GongOnAfterCreateInterface[WhiteSpace]
	OnAfterWhiteSpaceUpdateCallback GongOnAfterUpdateInterface[WhiteSpace]
	OnAfterWhiteSpaceDeleteCallback GongOnAfterDeleteInterface[WhiteSpace]

	BackRepo GongBackRepoInterface

	// if set will be called before each commit to the back repo
	OnInitCommitCallback          GongOnInitCommitInterface
	OnInitCommitFromFrontCallback GongOnInitCommitInterface
	OnInitCommitFromBackCallback  GongOnInitCommitInterface

	// Private slices to hold the registered hooks
	beforeCommitHooks []func(stage *Stage)
	afterCommitHooks  []func(stage *Stage)

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

type GongStage = Stage

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
	err := stage.ParseAstString(commitToApply, true)
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
	err := stage.ParseAstString(commitToApply, true)
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
	__gong__clearReferences(&stage.Alls_reference, &stage.Alls_instance, &stage.Alls_referenceOrder)

	__gong__clearReferences(&stage.Annotations_reference, &stage.Annotations_instance, &stage.Annotations_referenceOrder)

	__gong__clearReferences(&stage.Attributes_reference, &stage.Attributes_instance, &stage.Attributes_referenceOrder)

	__gong__clearReferences(&stage.AttributeGroups_reference, &stage.AttributeGroups_instance, &stage.AttributeGroups_referenceOrder)

	__gong__clearReferences(&stage.Choices_reference, &stage.Choices_instance, &stage.Choices_referenceOrder)

	__gong__clearReferences(&stage.ComplexContents_reference, &stage.ComplexContents_instance, &stage.ComplexContents_referenceOrder)

	__gong__clearReferences(&stage.ComplexTypes_reference, &stage.ComplexTypes_instance, &stage.ComplexTypes_referenceOrder)

	__gong__clearReferences(&stage.Documentations_reference, &stage.Documentations_instance, &stage.Documentations_referenceOrder)

	__gong__clearReferences(&stage.Elements_reference, &stage.Elements_instance, &stage.Elements_referenceOrder)

	__gong__clearReferences(&stage.Enumerations_reference, &stage.Enumerations_instance, &stage.Enumerations_referenceOrder)

	__gong__clearReferences(&stage.Extensions_reference, &stage.Extensions_instance, &stage.Extensions_referenceOrder)

	__gong__clearReferences(&stage.Groups_reference, &stage.Groups_instance, &stage.Groups_referenceOrder)

	__gong__clearReferences(&stage.Lengths_reference, &stage.Lengths_instance, &stage.Lengths_referenceOrder)

	__gong__clearReferences(&stage.MaxInclusives_reference, &stage.MaxInclusives_instance, &stage.MaxInclusives_referenceOrder)

	__gong__clearReferences(&stage.MaxLengths_reference, &stage.MaxLengths_instance, &stage.MaxLengths_referenceOrder)

	__gong__clearReferences(&stage.MinInclusives_reference, &stage.MinInclusives_instance, &stage.MinInclusives_referenceOrder)

	__gong__clearReferences(&stage.MinLengths_reference, &stage.MinLengths_instance, &stage.MinLengths_referenceOrder)

	__gong__clearReferences(&stage.Patterns_reference, &stage.Patterns_instance, &stage.Patterns_referenceOrder)

	__gong__clearReferences(&stage.Restrictions_reference, &stage.Restrictions_instance, &stage.Restrictions_referenceOrder)

	__gong__clearReferences(&stage.Schemas_reference, &stage.Schemas_instance, &stage.Schemas_referenceOrder)

	__gong__clearReferences(&stage.Sequences_reference, &stage.Sequences_instance, &stage.Sequences_referenceOrder)

	__gong__clearReferences(&stage.SimpleContents_reference, &stage.SimpleContents_instance, &stage.SimpleContents_referenceOrder)

	__gong__clearReferences(&stage.SimpleTypes_reference, &stage.SimpleTypes_instance, &stage.SimpleTypes_referenceOrder)

	__gong__clearReferences(&stage.TotalDigits_reference, &stage.TotalDigits_instance, &stage.TotalDigits_referenceOrder)

	__gong__clearReferences(&stage.Unions_reference, &stage.Unions_instance, &stage.Unions_referenceOrder)

	__gong__clearReferences(&stage.WhiteSpaces_reference, &stage.WhiteSpaces_instance, &stage.WhiteSpaces_referenceOrder)

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
	stage.AllOrder = __gong__recomputeOrder(stage.All_stagedOrder)

	stage.AnnotationOrder = __gong__recomputeOrder(stage.Annotation_stagedOrder)

	stage.AttributeOrder = __gong__recomputeOrder(stage.Attribute_stagedOrder)

	stage.AttributeGroupOrder = __gong__recomputeOrder(stage.AttributeGroup_stagedOrder)

	stage.ChoiceOrder = __gong__recomputeOrder(stage.Choice_stagedOrder)

	stage.ComplexContentOrder = __gong__recomputeOrder(stage.ComplexContent_stagedOrder)

	stage.ComplexTypeOrder = __gong__recomputeOrder(stage.ComplexType_stagedOrder)

	stage.DocumentationOrder = __gong__recomputeOrder(stage.Documentation_stagedOrder)

	stage.ElementOrder = __gong__recomputeOrder(stage.Element_stagedOrder)

	stage.EnumerationOrder = __gong__recomputeOrder(stage.Enumeration_stagedOrder)

	stage.ExtensionOrder = __gong__recomputeOrder(stage.Extension_stagedOrder)

	stage.GroupOrder = __gong__recomputeOrder(stage.Group_stagedOrder)

	stage.LengthOrder = __gong__recomputeOrder(stage.Length_stagedOrder)

	stage.MaxInclusiveOrder = __gong__recomputeOrder(stage.MaxInclusive_stagedOrder)

	stage.MaxLengthOrder = __gong__recomputeOrder(stage.MaxLength_stagedOrder)

	stage.MinInclusiveOrder = __gong__recomputeOrder(stage.MinInclusive_stagedOrder)

	stage.MinLengthOrder = __gong__recomputeOrder(stage.MinLength_stagedOrder)

	stage.PatternOrder = __gong__recomputeOrder(stage.Pattern_stagedOrder)

	stage.RestrictionOrder = __gong__recomputeOrder(stage.Restriction_stagedOrder)

	stage.SchemaOrder = __gong__recomputeOrder(stage.Schema_stagedOrder)

	stage.SequenceOrder = __gong__recomputeOrder(stage.Sequence_stagedOrder)

	stage.SimpleContentOrder = __gong__recomputeOrder(stage.SimpleContent_stagedOrder)

	stage.SimpleTypeOrder = __gong__recomputeOrder(stage.SimpleType_stagedOrder)

	stage.TotalDigitOrder = __gong__recomputeOrder(stage.TotalDigit_stagedOrder)

	stage.UnionOrder = __gong__recomputeOrder(stage.Union_stagedOrder)

	stage.WhiteSpaceOrder = __gong__recomputeOrder(stage.WhiteSpace_stagedOrder)

	// end of insertion point for max order recomputation
}

func (stage *Stage) SetDeltaMode(inDeltaMode bool) {
	stage.isInDeltaMode = inDeltaMode
}

func (stage *Stage) IsInDeltaMode() bool {
	return stage.isInDeltaMode
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

// GetInstancesByOrder is the Stage method returning a slice of generic pointers to gongstructs
// ordered by their order in the stage.
func (stage *Stage) GetInstancesByOrder[T GongstructPtr]() (res []T) {
	var t T
	switch any(t).(type) {
	// insertion point for case
	case *All:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Alls, stage.All_stagedOrder))
	case *Annotation:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Annotations, stage.Annotation_stagedOrder))
	case *Attribute:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Attributes, stage.Attribute_stagedOrder))
	case *AttributeGroup:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.AttributeGroups, stage.AttributeGroup_stagedOrder))
	case *Choice:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Choices, stage.Choice_stagedOrder))
	case *ComplexContent:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.ComplexContents, stage.ComplexContent_stagedOrder))
	case *ComplexType:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.ComplexTypes, stage.ComplexType_stagedOrder))
	case *Documentation:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Documentations, stage.Documentation_stagedOrder))
	case *Element:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Elements, stage.Element_stagedOrder))
	case *Enumeration:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Enumerations, stage.Enumeration_stagedOrder))
	case *Extension:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Extensions, stage.Extension_stagedOrder))
	case *Group:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Groups, stage.Group_stagedOrder))
	case *Length:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Lengths, stage.Length_stagedOrder))
	case *MaxInclusive:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.MaxInclusives, stage.MaxInclusive_stagedOrder))
	case *MaxLength:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.MaxLengths, stage.MaxLength_stagedOrder))
	case *MinInclusive:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.MinInclusives, stage.MinInclusive_stagedOrder))
	case *MinLength:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.MinLengths, stage.MinLength_stagedOrder))
	case *Pattern:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Patterns, stage.Pattern_stagedOrder))
	case *Restriction:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Restrictions, stage.Restriction_stagedOrder))
	case *Schema:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Schemas, stage.Schema_stagedOrder))
	case *Sequence:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Sequences, stage.Sequence_stagedOrder))
	case *SimpleContent:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.SimpleContents, stage.SimpleContent_stagedOrder))
	case *SimpleType:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.SimpleTypes, stage.SimpleType_stagedOrder))
	case *TotalDigit:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.TotalDigits, stage.TotalDigit_stagedOrder))
	case *Union:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Unions, stage.Union_stagedOrder))
	case *WhiteSpace:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.WhiteSpaces, stage.WhiteSpace_stagedOrder))

	}
	return
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
	return "github.com/fullstack-lang/gong/app/xsd/go/models"
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
		GongUnmarshallers: map[string]GongModelUnmarshaller{ // insertion point for unmarshallers
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

		navigationMode: GongNavigationModeNormal,
	}

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
	__gong__stage(stage.Alls, stage.All_stagedOrder, stage.All_orderStaged, &stage.AllOrder, stage.Alls_mapString, all, all.Name)
	return all
}

// StagePreserveOrder puts all to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AllOrder
// - update stage.AllOrder accordingly
func (all *All) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Alls, stage.All_stagedOrder, stage.All_orderStaged, &stage.AllOrder, stage.Alls_mapString, all, order, all.Name)
}

// Unstage removes all off the model stage
func (all *All) Unstage(stage *Stage) *All {
	__gong__unstage(stage.Alls, stage.Alls_mapString, all, all.Name)
	return all
}

// UnstageVoid removes all off the model stage
func (all *All) UnstageVoid(stage *Stage) {
	all.Unstage(stage)
}

func (all *All) StageVoid(stage *Stage) {
	all.Stage(stage)
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
	__gong__stage(stage.Annotations, stage.Annotation_stagedOrder, stage.Annotation_orderStaged, &stage.AnnotationOrder, stage.Annotations_mapString, annotation, annotation.Name)
	return annotation
}

// StagePreserveOrder puts annotation to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AnnotationOrder
// - update stage.AnnotationOrder accordingly
func (annotation *Annotation) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Annotations, stage.Annotation_stagedOrder, stage.Annotation_orderStaged, &stage.AnnotationOrder, stage.Annotations_mapString, annotation, order, annotation.Name)
}

// Unstage removes annotation off the model stage
func (annotation *Annotation) Unstage(stage *Stage) *Annotation {
	__gong__unstage(stage.Annotations, stage.Annotations_mapString, annotation, annotation.Name)
	return annotation
}

// UnstageVoid removes annotation off the model stage
func (annotation *Annotation) UnstageVoid(stage *Stage) {
	annotation.Unstage(stage)
}

func (annotation *Annotation) StageVoid(stage *Stage) {
	annotation.Stage(stage)
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
	__gong__stage(stage.Attributes, stage.Attribute_stagedOrder, stage.Attribute_orderStaged, &stage.AttributeOrder, stage.Attributes_mapString, attribute, attribute.Name)
	return attribute
}

// StagePreserveOrder puts attribute to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AttributeOrder
// - update stage.AttributeOrder accordingly
func (attribute *Attribute) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Attributes, stage.Attribute_stagedOrder, stage.Attribute_orderStaged, &stage.AttributeOrder, stage.Attributes_mapString, attribute, order, attribute.Name)
}

// Unstage removes attribute off the model stage
func (attribute *Attribute) Unstage(stage *Stage) *Attribute {
	__gong__unstage(stage.Attributes, stage.Attributes_mapString, attribute, attribute.Name)
	return attribute
}

// UnstageVoid removes attribute off the model stage
func (attribute *Attribute) UnstageVoid(stage *Stage) {
	attribute.Unstage(stage)
}

func (attribute *Attribute) StageVoid(stage *Stage) {
	attribute.Stage(stage)
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
	__gong__stage(stage.AttributeGroups, stage.AttributeGroup_stagedOrder, stage.AttributeGroup_orderStaged, &stage.AttributeGroupOrder, stage.AttributeGroups_mapString, attributegroup, attributegroup.Name)
	return attributegroup
}

// StagePreserveOrder puts attributegroup to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AttributeGroupOrder
// - update stage.AttributeGroupOrder accordingly
func (attributegroup *AttributeGroup) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AttributeGroups, stage.AttributeGroup_stagedOrder, stage.AttributeGroup_orderStaged, &stage.AttributeGroupOrder, stage.AttributeGroups_mapString, attributegroup, order, attributegroup.Name)
}

// Unstage removes attributegroup off the model stage
func (attributegroup *AttributeGroup) Unstage(stage *Stage) *AttributeGroup {
	__gong__unstage(stage.AttributeGroups, stage.AttributeGroups_mapString, attributegroup, attributegroup.Name)
	return attributegroup
}

// UnstageVoid removes attributegroup off the model stage
func (attributegroup *AttributeGroup) UnstageVoid(stage *Stage) {
	attributegroup.Unstage(stage)
}

func (attributegroup *AttributeGroup) StageVoid(stage *Stage) {
	attributegroup.Stage(stage)
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
	__gong__stage(stage.Choices, stage.Choice_stagedOrder, stage.Choice_orderStaged, &stage.ChoiceOrder, stage.Choices_mapString, choice, choice.Name)
	return choice
}

// StagePreserveOrder puts choice to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ChoiceOrder
// - update stage.ChoiceOrder accordingly
func (choice *Choice) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Choices, stage.Choice_stagedOrder, stage.Choice_orderStaged, &stage.ChoiceOrder, stage.Choices_mapString, choice, order, choice.Name)
}

// Unstage removes choice off the model stage
func (choice *Choice) Unstage(stage *Stage) *Choice {
	__gong__unstage(stage.Choices, stage.Choices_mapString, choice, choice.Name)
	return choice
}

// UnstageVoid removes choice off the model stage
func (choice *Choice) UnstageVoid(stage *Stage) {
	choice.Unstage(stage)
}

func (choice *Choice) StageVoid(stage *Stage) {
	choice.Stage(stage)
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
	__gong__stage(stage.ComplexContents, stage.ComplexContent_stagedOrder, stage.ComplexContent_orderStaged, &stage.ComplexContentOrder, stage.ComplexContents_mapString, complexcontent, complexcontent.Name)
	return complexcontent
}

// StagePreserveOrder puts complexcontent to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ComplexContentOrder
// - update stage.ComplexContentOrder accordingly
func (complexcontent *ComplexContent) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ComplexContents, stage.ComplexContent_stagedOrder, stage.ComplexContent_orderStaged, &stage.ComplexContentOrder, stage.ComplexContents_mapString, complexcontent, order, complexcontent.Name)
}

// Unstage removes complexcontent off the model stage
func (complexcontent *ComplexContent) Unstage(stage *Stage) *ComplexContent {
	__gong__unstage(stage.ComplexContents, stage.ComplexContents_mapString, complexcontent, complexcontent.Name)
	return complexcontent
}

// UnstageVoid removes complexcontent off the model stage
func (complexcontent *ComplexContent) UnstageVoid(stage *Stage) {
	complexcontent.Unstage(stage)
}

func (complexcontent *ComplexContent) StageVoid(stage *Stage) {
	complexcontent.Stage(stage)
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
	__gong__stage(stage.ComplexTypes, stage.ComplexType_stagedOrder, stage.ComplexType_orderStaged, &stage.ComplexTypeOrder, stage.ComplexTypes_mapString, complextype, complextype.Name)
	return complextype
}

// StagePreserveOrder puts complextype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ComplexTypeOrder
// - update stage.ComplexTypeOrder accordingly
func (complextype *ComplexType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ComplexTypes, stage.ComplexType_stagedOrder, stage.ComplexType_orderStaged, &stage.ComplexTypeOrder, stage.ComplexTypes_mapString, complextype, order, complextype.Name)
}

// Unstage removes complextype off the model stage
func (complextype *ComplexType) Unstage(stage *Stage) *ComplexType {
	__gong__unstage(stage.ComplexTypes, stage.ComplexTypes_mapString, complextype, complextype.Name)
	return complextype
}

// UnstageVoid removes complextype off the model stage
func (complextype *ComplexType) UnstageVoid(stage *Stage) {
	complextype.Unstage(stage)
}

func (complextype *ComplexType) StageVoid(stage *Stage) {
	complextype.Stage(stage)
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
	__gong__stage(stage.Documentations, stage.Documentation_stagedOrder, stage.Documentation_orderStaged, &stage.DocumentationOrder, stage.Documentations_mapString, documentation, documentation.Name)
	return documentation
}

// StagePreserveOrder puts documentation to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DocumentationOrder
// - update stage.DocumentationOrder accordingly
func (documentation *Documentation) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Documentations, stage.Documentation_stagedOrder, stage.Documentation_orderStaged, &stage.DocumentationOrder, stage.Documentations_mapString, documentation, order, documentation.Name)
}

// Unstage removes documentation off the model stage
func (documentation *Documentation) Unstage(stage *Stage) *Documentation {
	__gong__unstage(stage.Documentations, stage.Documentations_mapString, documentation, documentation.Name)
	return documentation
}

// UnstageVoid removes documentation off the model stage
func (documentation *Documentation) UnstageVoid(stage *Stage) {
	documentation.Unstage(stage)
}

func (documentation *Documentation) StageVoid(stage *Stage) {
	documentation.Stage(stage)
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
	__gong__stage(stage.Elements, stage.Element_stagedOrder, stage.Element_orderStaged, &stage.ElementOrder, stage.Elements_mapString, element, element.Name)
	return element
}

// StagePreserveOrder puts element to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ElementOrder
// - update stage.ElementOrder accordingly
func (element *Element) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Elements, stage.Element_stagedOrder, stage.Element_orderStaged, &stage.ElementOrder, stage.Elements_mapString, element, order, element.Name)
}

// Unstage removes element off the model stage
func (element *Element) Unstage(stage *Stage) *Element {
	__gong__unstage(stage.Elements, stage.Elements_mapString, element, element.Name)
	return element
}

// UnstageVoid removes element off the model stage
func (element *Element) UnstageVoid(stage *Stage) {
	element.Unstage(stage)
}

func (element *Element) StageVoid(stage *Stage) {
	element.Stage(stage)
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
	__gong__stage(stage.Enumerations, stage.Enumeration_stagedOrder, stage.Enumeration_orderStaged, &stage.EnumerationOrder, stage.Enumerations_mapString, enumeration, enumeration.Name)
	return enumeration
}

// StagePreserveOrder puts enumeration to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.EnumerationOrder
// - update stage.EnumerationOrder accordingly
func (enumeration *Enumeration) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Enumerations, stage.Enumeration_stagedOrder, stage.Enumeration_orderStaged, &stage.EnumerationOrder, stage.Enumerations_mapString, enumeration, order, enumeration.Name)
}

// Unstage removes enumeration off the model stage
func (enumeration *Enumeration) Unstage(stage *Stage) *Enumeration {
	__gong__unstage(stage.Enumerations, stage.Enumerations_mapString, enumeration, enumeration.Name)
	return enumeration
}

// UnstageVoid removes enumeration off the model stage
func (enumeration *Enumeration) UnstageVoid(stage *Stage) {
	enumeration.Unstage(stage)
}

func (enumeration *Enumeration) StageVoid(stage *Stage) {
	enumeration.Stage(stage)
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
	__gong__stage(stage.Extensions, stage.Extension_stagedOrder, stage.Extension_orderStaged, &stage.ExtensionOrder, stage.Extensions_mapString, extension, extension.Name)
	return extension
}

// StagePreserveOrder puts extension to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ExtensionOrder
// - update stage.ExtensionOrder accordingly
func (extension *Extension) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Extensions, stage.Extension_stagedOrder, stage.Extension_orderStaged, &stage.ExtensionOrder, stage.Extensions_mapString, extension, order, extension.Name)
}

// Unstage removes extension off the model stage
func (extension *Extension) Unstage(stage *Stage) *Extension {
	__gong__unstage(stage.Extensions, stage.Extensions_mapString, extension, extension.Name)
	return extension
}

// UnstageVoid removes extension off the model stage
func (extension *Extension) UnstageVoid(stage *Stage) {
	extension.Unstage(stage)
}

func (extension *Extension) StageVoid(stage *Stage) {
	extension.Stage(stage)
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
	__gong__stage(stage.Groups, stage.Group_stagedOrder, stage.Group_orderStaged, &stage.GroupOrder, stage.Groups_mapString, group, group.Name)
	return group
}

// StagePreserveOrder puts group to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.GroupOrder
// - update stage.GroupOrder accordingly
func (group *Group) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Groups, stage.Group_stagedOrder, stage.Group_orderStaged, &stage.GroupOrder, stage.Groups_mapString, group, order, group.Name)
}

// Unstage removes group off the model stage
func (group *Group) Unstage(stage *Stage) *Group {
	__gong__unstage(stage.Groups, stage.Groups_mapString, group, group.Name)
	return group
}

// UnstageVoid removes group off the model stage
func (group *Group) UnstageVoid(stage *Stage) {
	group.Unstage(stage)
}

func (group *Group) StageVoid(stage *Stage) {
	group.Stage(stage)
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
	__gong__stage(stage.Lengths, stage.Length_stagedOrder, stage.Length_orderStaged, &stage.LengthOrder, stage.Lengths_mapString, length, length.Name)
	return length
}

// StagePreserveOrder puts length to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.LengthOrder
// - update stage.LengthOrder accordingly
func (length *Length) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Lengths, stage.Length_stagedOrder, stage.Length_orderStaged, &stage.LengthOrder, stage.Lengths_mapString, length, order, length.Name)
}

// Unstage removes length off the model stage
func (length *Length) Unstage(stage *Stage) *Length {
	__gong__unstage(stage.Lengths, stage.Lengths_mapString, length, length.Name)
	return length
}

// UnstageVoid removes length off the model stage
func (length *Length) UnstageVoid(stage *Stage) {
	length.Unstage(stage)
}

func (length *Length) StageVoid(stage *Stage) {
	length.Stage(stage)
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
	__gong__stage(stage.MaxInclusives, stage.MaxInclusive_stagedOrder, stage.MaxInclusive_orderStaged, &stage.MaxInclusiveOrder, stage.MaxInclusives_mapString, maxinclusive, maxinclusive.Name)
	return maxinclusive
}

// StagePreserveOrder puts maxinclusive to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MaxInclusiveOrder
// - update stage.MaxInclusiveOrder accordingly
func (maxinclusive *MaxInclusive) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MaxInclusives, stage.MaxInclusive_stagedOrder, stage.MaxInclusive_orderStaged, &stage.MaxInclusiveOrder, stage.MaxInclusives_mapString, maxinclusive, order, maxinclusive.Name)
}

// Unstage removes maxinclusive off the model stage
func (maxinclusive *MaxInclusive) Unstage(stage *Stage) *MaxInclusive {
	__gong__unstage(stage.MaxInclusives, stage.MaxInclusives_mapString, maxinclusive, maxinclusive.Name)
	return maxinclusive
}

// UnstageVoid removes maxinclusive off the model stage
func (maxinclusive *MaxInclusive) UnstageVoid(stage *Stage) {
	maxinclusive.Unstage(stage)
}

func (maxinclusive *MaxInclusive) StageVoid(stage *Stage) {
	maxinclusive.Stage(stage)
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
	__gong__stage(stage.MaxLengths, stage.MaxLength_stagedOrder, stage.MaxLength_orderStaged, &stage.MaxLengthOrder, stage.MaxLengths_mapString, maxlength, maxlength.Name)
	return maxlength
}

// StagePreserveOrder puts maxlength to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MaxLengthOrder
// - update stage.MaxLengthOrder accordingly
func (maxlength *MaxLength) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MaxLengths, stage.MaxLength_stagedOrder, stage.MaxLength_orderStaged, &stage.MaxLengthOrder, stage.MaxLengths_mapString, maxlength, order, maxlength.Name)
}

// Unstage removes maxlength off the model stage
func (maxlength *MaxLength) Unstage(stage *Stage) *MaxLength {
	__gong__unstage(stage.MaxLengths, stage.MaxLengths_mapString, maxlength, maxlength.Name)
	return maxlength
}

// UnstageVoid removes maxlength off the model stage
func (maxlength *MaxLength) UnstageVoid(stage *Stage) {
	maxlength.Unstage(stage)
}

func (maxlength *MaxLength) StageVoid(stage *Stage) {
	maxlength.Stage(stage)
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
	__gong__stage(stage.MinInclusives, stage.MinInclusive_stagedOrder, stage.MinInclusive_orderStaged, &stage.MinInclusiveOrder, stage.MinInclusives_mapString, mininclusive, mininclusive.Name)
	return mininclusive
}

// StagePreserveOrder puts mininclusive to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MinInclusiveOrder
// - update stage.MinInclusiveOrder accordingly
func (mininclusive *MinInclusive) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MinInclusives, stage.MinInclusive_stagedOrder, stage.MinInclusive_orderStaged, &stage.MinInclusiveOrder, stage.MinInclusives_mapString, mininclusive, order, mininclusive.Name)
}

// Unstage removes mininclusive off the model stage
func (mininclusive *MinInclusive) Unstage(stage *Stage) *MinInclusive {
	__gong__unstage(stage.MinInclusives, stage.MinInclusives_mapString, mininclusive, mininclusive.Name)
	return mininclusive
}

// UnstageVoid removes mininclusive off the model stage
func (mininclusive *MinInclusive) UnstageVoid(stage *Stage) {
	mininclusive.Unstage(stage)
}

func (mininclusive *MinInclusive) StageVoid(stage *Stage) {
	mininclusive.Stage(stage)
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
	__gong__stage(stage.MinLengths, stage.MinLength_stagedOrder, stage.MinLength_orderStaged, &stage.MinLengthOrder, stage.MinLengths_mapString, minlength, minlength.Name)
	return minlength
}

// StagePreserveOrder puts minlength to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MinLengthOrder
// - update stage.MinLengthOrder accordingly
func (minlength *MinLength) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MinLengths, stage.MinLength_stagedOrder, stage.MinLength_orderStaged, &stage.MinLengthOrder, stage.MinLengths_mapString, minlength, order, minlength.Name)
}

// Unstage removes minlength off the model stage
func (minlength *MinLength) Unstage(stage *Stage) *MinLength {
	__gong__unstage(stage.MinLengths, stage.MinLengths_mapString, minlength, minlength.Name)
	return minlength
}

// UnstageVoid removes minlength off the model stage
func (minlength *MinLength) UnstageVoid(stage *Stage) {
	minlength.Unstage(stage)
}

func (minlength *MinLength) StageVoid(stage *Stage) {
	minlength.Stage(stage)
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
	__gong__stage(stage.Patterns, stage.Pattern_stagedOrder, stage.Pattern_orderStaged, &stage.PatternOrder, stage.Patterns_mapString, pattern, pattern.Name)
	return pattern
}

// StagePreserveOrder puts pattern to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.PatternOrder
// - update stage.PatternOrder accordingly
func (pattern *Pattern) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Patterns, stage.Pattern_stagedOrder, stage.Pattern_orderStaged, &stage.PatternOrder, stage.Patterns_mapString, pattern, order, pattern.Name)
}

// Unstage removes pattern off the model stage
func (pattern *Pattern) Unstage(stage *Stage) *Pattern {
	__gong__unstage(stage.Patterns, stage.Patterns_mapString, pattern, pattern.Name)
	return pattern
}

// UnstageVoid removes pattern off the model stage
func (pattern *Pattern) UnstageVoid(stage *Stage) {
	pattern.Unstage(stage)
}

func (pattern *Pattern) StageVoid(stage *Stage) {
	pattern.Stage(stage)
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
	__gong__stage(stage.Restrictions, stage.Restriction_stagedOrder, stage.Restriction_orderStaged, &stage.RestrictionOrder, stage.Restrictions_mapString, restriction, restriction.Name)
	return restriction
}

// StagePreserveOrder puts restriction to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.RestrictionOrder
// - update stage.RestrictionOrder accordingly
func (restriction *Restriction) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Restrictions, stage.Restriction_stagedOrder, stage.Restriction_orderStaged, &stage.RestrictionOrder, stage.Restrictions_mapString, restriction, order, restriction.Name)
}

// Unstage removes restriction off the model stage
func (restriction *Restriction) Unstage(stage *Stage) *Restriction {
	__gong__unstage(stage.Restrictions, stage.Restrictions_mapString, restriction, restriction.Name)
	return restriction
}

// UnstageVoid removes restriction off the model stage
func (restriction *Restriction) UnstageVoid(stage *Stage) {
	restriction.Unstage(stage)
}

func (restriction *Restriction) StageVoid(stage *Stage) {
	restriction.Stage(stage)
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
	__gong__stage(stage.Schemas, stage.Schema_stagedOrder, stage.Schema_orderStaged, &stage.SchemaOrder, stage.Schemas_mapString, schema, schema.Name)
	return schema
}

// StagePreserveOrder puts schema to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.SchemaOrder
// - update stage.SchemaOrder accordingly
func (schema *Schema) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Schemas, stage.Schema_stagedOrder, stage.Schema_orderStaged, &stage.SchemaOrder, stage.Schemas_mapString, schema, order, schema.Name)
}

// Unstage removes schema off the model stage
func (schema *Schema) Unstage(stage *Stage) *Schema {
	__gong__unstage(stage.Schemas, stage.Schemas_mapString, schema, schema.Name)
	return schema
}

// UnstageVoid removes schema off the model stage
func (schema *Schema) UnstageVoid(stage *Stage) {
	schema.Unstage(stage)
}

func (schema *Schema) StageVoid(stage *Stage) {
	schema.Stage(stage)
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
	__gong__stage(stage.Sequences, stage.Sequence_stagedOrder, stage.Sequence_orderStaged, &stage.SequenceOrder, stage.Sequences_mapString, sequence, sequence.Name)
	return sequence
}

// StagePreserveOrder puts sequence to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.SequenceOrder
// - update stage.SequenceOrder accordingly
func (sequence *Sequence) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Sequences, stage.Sequence_stagedOrder, stage.Sequence_orderStaged, &stage.SequenceOrder, stage.Sequences_mapString, sequence, order, sequence.Name)
}

// Unstage removes sequence off the model stage
func (sequence *Sequence) Unstage(stage *Stage) *Sequence {
	__gong__unstage(stage.Sequences, stage.Sequences_mapString, sequence, sequence.Name)
	return sequence
}

// UnstageVoid removes sequence off the model stage
func (sequence *Sequence) UnstageVoid(stage *Stage) {
	sequence.Unstage(stage)
}

func (sequence *Sequence) StageVoid(stage *Stage) {
	sequence.Stage(stage)
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
	__gong__stage(stage.SimpleContents, stage.SimpleContent_stagedOrder, stage.SimpleContent_orderStaged, &stage.SimpleContentOrder, stage.SimpleContents_mapString, simplecontent, simplecontent.Name)
	return simplecontent
}

// StagePreserveOrder puts simplecontent to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.SimpleContentOrder
// - update stage.SimpleContentOrder accordingly
func (simplecontent *SimpleContent) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.SimpleContents, stage.SimpleContent_stagedOrder, stage.SimpleContent_orderStaged, &stage.SimpleContentOrder, stage.SimpleContents_mapString, simplecontent, order, simplecontent.Name)
}

// Unstage removes simplecontent off the model stage
func (simplecontent *SimpleContent) Unstage(stage *Stage) *SimpleContent {
	__gong__unstage(stage.SimpleContents, stage.SimpleContents_mapString, simplecontent, simplecontent.Name)
	return simplecontent
}

// UnstageVoid removes simplecontent off the model stage
func (simplecontent *SimpleContent) UnstageVoid(stage *Stage) {
	simplecontent.Unstage(stage)
}

func (simplecontent *SimpleContent) StageVoid(stage *Stage) {
	simplecontent.Stage(stage)
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
	__gong__stage(stage.SimpleTypes, stage.SimpleType_stagedOrder, stage.SimpleType_orderStaged, &stage.SimpleTypeOrder, stage.SimpleTypes_mapString, simpletype, simpletype.Name)
	return simpletype
}

// StagePreserveOrder puts simpletype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.SimpleTypeOrder
// - update stage.SimpleTypeOrder accordingly
func (simpletype *SimpleType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.SimpleTypes, stage.SimpleType_stagedOrder, stage.SimpleType_orderStaged, &stage.SimpleTypeOrder, stage.SimpleTypes_mapString, simpletype, order, simpletype.Name)
}

// Unstage removes simpletype off the model stage
func (simpletype *SimpleType) Unstage(stage *Stage) *SimpleType {
	__gong__unstage(stage.SimpleTypes, stage.SimpleTypes_mapString, simpletype, simpletype.Name)
	return simpletype
}

// UnstageVoid removes simpletype off the model stage
func (simpletype *SimpleType) UnstageVoid(stage *Stage) {
	simpletype.Unstage(stage)
}

func (simpletype *SimpleType) StageVoid(stage *Stage) {
	simpletype.Stage(stage)
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
	__gong__stage(stage.TotalDigits, stage.TotalDigit_stagedOrder, stage.TotalDigit_orderStaged, &stage.TotalDigitOrder, stage.TotalDigits_mapString, totaldigit, totaldigit.Name)
	return totaldigit
}

// StagePreserveOrder puts totaldigit to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TotalDigitOrder
// - update stage.TotalDigitOrder accordingly
func (totaldigit *TotalDigit) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.TotalDigits, stage.TotalDigit_stagedOrder, stage.TotalDigit_orderStaged, &stage.TotalDigitOrder, stage.TotalDigits_mapString, totaldigit, order, totaldigit.Name)
}

// Unstage removes totaldigit off the model stage
func (totaldigit *TotalDigit) Unstage(stage *Stage) *TotalDigit {
	__gong__unstage(stage.TotalDigits, stage.TotalDigits_mapString, totaldigit, totaldigit.Name)
	return totaldigit
}

// UnstageVoid removes totaldigit off the model stage
func (totaldigit *TotalDigit) UnstageVoid(stage *Stage) {
	totaldigit.Unstage(stage)
}

func (totaldigit *TotalDigit) StageVoid(stage *Stage) {
	totaldigit.Stage(stage)
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
	__gong__stage(stage.Unions, stage.Union_stagedOrder, stage.Union_orderStaged, &stage.UnionOrder, stage.Unions_mapString, union, union.Name)
	return union
}

// StagePreserveOrder puts union to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.UnionOrder
// - update stage.UnionOrder accordingly
func (union *Union) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Unions, stage.Union_stagedOrder, stage.Union_orderStaged, &stage.UnionOrder, stage.Unions_mapString, union, order, union.Name)
}

// Unstage removes union off the model stage
func (union *Union) Unstage(stage *Stage) *Union {
	__gong__unstage(stage.Unions, stage.Unions_mapString, union, union.Name)
	return union
}

// UnstageVoid removes union off the model stage
func (union *Union) UnstageVoid(stage *Stage) {
	union.Unstage(stage)
}

func (union *Union) StageVoid(stage *Stage) {
	union.Stage(stage)
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
	__gong__stage(stage.WhiteSpaces, stage.WhiteSpace_stagedOrder, stage.WhiteSpace_orderStaged, &stage.WhiteSpaceOrder, stage.WhiteSpaces_mapString, whitespace, whitespace.Name)
	return whitespace
}

// StagePreserveOrder puts whitespace to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.WhiteSpaceOrder
// - update stage.WhiteSpaceOrder accordingly
func (whitespace *WhiteSpace) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.WhiteSpaces, stage.WhiteSpace_stagedOrder, stage.WhiteSpace_orderStaged, &stage.WhiteSpaceOrder, stage.WhiteSpaces_mapString, whitespace, order, whitespace.Name)
}

// Unstage removes whitespace off the model stage
func (whitespace *WhiteSpace) Unstage(stage *Stage) *WhiteSpace {
	__gong__unstage(stage.WhiteSpaces, stage.WhiteSpaces_mapString, whitespace, whitespace.Name)
	return whitespace
}

// UnstageVoid removes whitespace off the model stage
func (whitespace *WhiteSpace) UnstageVoid(stage *Stage) {
	whitespace.Unstage(stage)
}

func (whitespace *WhiteSpace) StageVoid(stage *Stage) {
	whitespace.Stage(stage)
}

// for satisfaction of GongStruct interface
func (whitespace *WhiteSpace) GetName() (res string) {
	return whitespace.Name
}

// for satisfaction of GongStruct interface
func (whitespace *WhiteSpace) SetName(name string) {
	whitespace.Name = name
}

func (stage *Stage) Reset() { // insertion point for array reset
	__gong__resetStageType(&stage.Alls, &stage.Alls_mapString, &stage.All_stagedOrder, &stage.AllOrder)

	__gong__resetStageType(&stage.Annotations, &stage.Annotations_mapString, &stage.Annotation_stagedOrder, &stage.AnnotationOrder)

	__gong__resetStageType(&stage.Attributes, &stage.Attributes_mapString, &stage.Attribute_stagedOrder, &stage.AttributeOrder)

	__gong__resetStageType(&stage.AttributeGroups, &stage.AttributeGroups_mapString, &stage.AttributeGroup_stagedOrder, &stage.AttributeGroupOrder)

	__gong__resetStageType(&stage.Choices, &stage.Choices_mapString, &stage.Choice_stagedOrder, &stage.ChoiceOrder)

	__gong__resetStageType(&stage.ComplexContents, &stage.ComplexContents_mapString, &stage.ComplexContent_stagedOrder, &stage.ComplexContentOrder)

	__gong__resetStageType(&stage.ComplexTypes, &stage.ComplexTypes_mapString, &stage.ComplexType_stagedOrder, &stage.ComplexTypeOrder)

	__gong__resetStageType(&stage.Documentations, &stage.Documentations_mapString, &stage.Documentation_stagedOrder, &stage.DocumentationOrder)

	__gong__resetStageType(&stage.Elements, &stage.Elements_mapString, &stage.Element_stagedOrder, &stage.ElementOrder)

	__gong__resetStageType(&stage.Enumerations, &stage.Enumerations_mapString, &stage.Enumeration_stagedOrder, &stage.EnumerationOrder)

	__gong__resetStageType(&stage.Extensions, &stage.Extensions_mapString, &stage.Extension_stagedOrder, &stage.ExtensionOrder)

	__gong__resetStageType(&stage.Groups, &stage.Groups_mapString, &stage.Group_stagedOrder, &stage.GroupOrder)

	__gong__resetStageType(&stage.Lengths, &stage.Lengths_mapString, &stage.Length_stagedOrder, &stage.LengthOrder)

	__gong__resetStageType(&stage.MaxInclusives, &stage.MaxInclusives_mapString, &stage.MaxInclusive_stagedOrder, &stage.MaxInclusiveOrder)

	__gong__resetStageType(&stage.MaxLengths, &stage.MaxLengths_mapString, &stage.MaxLength_stagedOrder, &stage.MaxLengthOrder)

	__gong__resetStageType(&stage.MinInclusives, &stage.MinInclusives_mapString, &stage.MinInclusive_stagedOrder, &stage.MinInclusiveOrder)

	__gong__resetStageType(&stage.MinLengths, &stage.MinLengths_mapString, &stage.MinLength_stagedOrder, &stage.MinLengthOrder)

	__gong__resetStageType(&stage.Patterns, &stage.Patterns_mapString, &stage.Pattern_stagedOrder, &stage.PatternOrder)

	__gong__resetStageType(&stage.Restrictions, &stage.Restrictions_mapString, &stage.Restriction_stagedOrder, &stage.RestrictionOrder)

	__gong__resetStageType(&stage.Schemas, &stage.Schemas_mapString, &stage.Schema_stagedOrder, &stage.SchemaOrder)

	__gong__resetStageType(&stage.Sequences, &stage.Sequences_mapString, &stage.Sequence_stagedOrder, &stage.SequenceOrder)

	__gong__resetStageType(&stage.SimpleContents, &stage.SimpleContents_mapString, &stage.SimpleContent_stagedOrder, &stage.SimpleContentOrder)

	__gong__resetStageType(&stage.SimpleTypes, &stage.SimpleTypes_mapString, &stage.SimpleType_stagedOrder, &stage.SimpleTypeOrder)

	__gong__resetStageType(&stage.TotalDigits, &stage.TotalDigits_mapString, &stage.TotalDigit_stagedOrder, &stage.TotalDigitOrder)

	__gong__resetStageType(&stage.Unions, &stage.Unions_mapString, &stage.Union_stagedOrder, &stage.UnionOrder)

	__gong__resetStageType(&stage.WhiteSpaces, &stage.WhiteSpaces_mapString, &stage.WhiteSpace_stagedOrder, &stage.WhiteSpaceOrder)

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
type Gongstruct any

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

// GetInstancesSet is the Stage method returning the set of staged instances (pointer-type constraint).
func (stage *Stage) GetInstancesSet[Type GongstructPtr]() *map[Type]struct{} {
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

// GongGetAssociationName is a generic function that returns an instance of Type
// where each association is filled with an instance whose name is the name of the association
//
// This function can be handy for generating navigation function that are refactorable
func GongGetAssociationName[Type Gongstruct]() *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for instance with special fields
	case All:
		return any(&All{
			Annotation: &Annotation{Name: "Annotation"},
			Sequences: []*Sequence{{Name: "Sequences"}},
			Alls: []*All{{Name: "Alls"}},
			Choices: []*Choice{{Name: "Choices"}},
			Groups: []*Group{{Name: "Groups"}},
			Elements: []*Element{{Name: "Elements"}},
		}).(*Type)
	case Annotation:
		return any(&Annotation{
			Documentations: []*Documentation{{Name: "Documentations"}},
		}).(*Type)
	case Attribute:
		return any(&Attribute{
			Annotation: &Annotation{Name: "Annotation"},
		}).(*Type)
	case AttributeGroup:
		return any(&AttributeGroup{
			Annotation: &Annotation{Name: "Annotation"},
			AttributeGroups: []*AttributeGroup{{Name: "AttributeGroups"}},
			Attributes: []*Attribute{{Name: "Attributes"}},
		}).(*Type)
	case Choice:
		return any(&Choice{
			Annotation: &Annotation{Name: "Annotation"},
			Sequences: []*Sequence{{Name: "Sequences"}},
			Alls: []*All{{Name: "Alls"}},
			Choices: []*Choice{{Name: "Choices"}},
			Groups: []*Group{{Name: "Groups"}},
			Elements: []*Element{{Name: "Elements"}},
		}).(*Type)
	case ComplexType:
		return any(&ComplexType{
			OuterElement: &Element{Name: "OuterElement"},
			Annotation: &Annotation{Name: "Annotation"},
			Sequences: []*Sequence{{Name: "Sequences"}},
			Alls: []*All{{Name: "Alls"}},
			Choices: []*Choice{{Name: "Choices"}},
			Groups: []*Group{{Name: "Groups"}},
			Elements: []*Element{{Name: "Elements"}},
			Extension: &Extension{Name: "Extension"},
			SimpleContent: &SimpleContent{Name: "SimpleContent"},
			ComplexContent: &ComplexContent{Name: "ComplexContent"},
			Attributes: []*Attribute{{Name: "Attributes"}},
			AttributeGroups: []*AttributeGroup{{Name: "AttributeGroups"}},
		}).(*Type)
	case Element:
		return any(&Element{
			Annotation: &Annotation{Name: "Annotation"},
			SimpleType: &SimpleType{Name: "SimpleType"},
			ComplexType: &ComplexType{Name: "ComplexType"},
			Groups: []*Group{{Name: "Groups"}},
		}).(*Type)
	case Enumeration:
		return any(&Enumeration{
			Annotation: &Annotation{Name: "Annotation"},
		}).(*Type)
	case Extension:
		return any(&Extension{
			Sequences: []*Sequence{{Name: "Sequences"}},
			Alls: []*All{{Name: "Alls"}},
			Choices: []*Choice{{Name: "Choices"}},
			Groups: []*Group{{Name: "Groups"}},
			Elements: []*Element{{Name: "Elements"}},
			Attributes: []*Attribute{{Name: "Attributes"}},
			AttributeGroups: []*AttributeGroup{{Name: "AttributeGroups"}},
		}).(*Type)
	case Group:
		return any(&Group{
			Annotation: &Annotation{Name: "Annotation"},
			OuterElement: &Element{Name: "OuterElement"},
			Sequences: []*Sequence{{Name: "Sequences"}},
			Alls: []*All{{Name: "Alls"}},
			Choices: []*Choice{{Name: "Choices"}},
			Groups: []*Group{{Name: "Groups"}},
			Elements: []*Element{{Name: "Elements"}},
		}).(*Type)
	case Length:
		return any(&Length{
			Annotation: &Annotation{Name: "Annotation"},
		}).(*Type)
	case MaxInclusive:
		return any(&MaxInclusive{
			Annotation: &Annotation{Name: "Annotation"},
		}).(*Type)
	case MaxLength:
		return any(&MaxLength{
			Annotation: &Annotation{Name: "Annotation"},
		}).(*Type)
	case MinInclusive:
		return any(&MinInclusive{
			Annotation: &Annotation{Name: "Annotation"},
		}).(*Type)
	case MinLength:
		return any(&MinLength{
			Annotation: &Annotation{Name: "Annotation"},
		}).(*Type)
	case Pattern:
		return any(&Pattern{
			Annotation: &Annotation{Name: "Annotation"},
		}).(*Type)
	case Restriction:
		return any(&Restriction{
			Annotation: &Annotation{Name: "Annotation"},
			Enumerations: []*Enumeration{{Name: "Enumerations"}},
			MinInclusive: &MinInclusive{Name: "MinInclusive"},
			MaxInclusive: &MaxInclusive{Name: "MaxInclusive"},
			Pattern: &Pattern{Name: "Pattern"},
			WhiteSpace: &WhiteSpace{Name: "WhiteSpace"},
			MinLength: &MinLength{Name: "MinLength"},
			MaxLength: &MaxLength{Name: "MaxLength"},
			Length: &Length{Name: "Length"},
			TotalDigit: &TotalDigit{Name: "TotalDigit"},
		}).(*Type)
	case Schema:
		return any(&Schema{
			Annotation: &Annotation{Name: "Annotation"},
			Elements: []*Element{{Name: "Elements"}},
			SimpleTypes: []*SimpleType{{Name: "SimpleTypes"}},
			ComplexTypes: []*ComplexType{{Name: "ComplexTypes"}},
			AttributeGroups: []*AttributeGroup{{Name: "AttributeGroups"}},
			Groups: []*Group{{Name: "Groups"}},
		}).(*Type)
	case Sequence:
		return any(&Sequence{
			Annotation: &Annotation{Name: "Annotation"},
			Sequences: []*Sequence{{Name: "Sequences"}},
			Alls: []*All{{Name: "Alls"}},
			Choices: []*Choice{{Name: "Choices"}},
			Groups: []*Group{{Name: "Groups"}},
			Elements: []*Element{{Name: "Elements"}},
		}).(*Type)
	case SimpleContent:
		return any(&SimpleContent{
			Extension: &Extension{Name: "Extension"},
			Restriction: &Restriction{Name: "Restriction"},
		}).(*Type)
	case SimpleType:
		return any(&SimpleType{
			Annotation: &Annotation{Name: "Annotation"},
			Restriction: &Restriction{Name: "Restriction"},
			Union: &Union{Name: "Union"},
		}).(*Type)
	case TotalDigit:
		return any(&TotalDigit{
			Annotation: &Annotation{Name: "Annotation"},
		}).(*Type)
	case Union:
		return any(&Union{
			Annotation: &Annotation{Name: "Annotation"},
		}).(*Type)
	case WhiteSpace:
		return any(&WhiteSpace{
			Annotation: &Annotation{Name: "Annotation"},
		}).(*Type)
	default:
		return &ret
	}
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

// GetSliceOfPointersReverseMap is the Stage method for backtrack navigation of slice-of-pointers associations.
func (stage *Stage) GetSliceOfPointersReverseMap[Start, End Gongstruct](fieldname string) map[*End][]*Start {
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

// GongNewInstance creates a new instance of the Gongstruct
func GongNewInstance[Type GongstructPtr]() (res Type) {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic new instance
	case *All:
		res = any(new(All)).(Type)
	case *Annotation:
		res = any(new(Annotation)).(Type)
	case *Attribute:
		res = any(new(Attribute)).(Type)
	case *AttributeGroup:
		res = any(new(AttributeGroup)).(Type)
	case *Choice:
		res = any(new(Choice)).(Type)
	case *ComplexContent:
		res = any(new(ComplexContent)).(Type)
	case *ComplexType:
		res = any(new(ComplexType)).(Type)
	case *Documentation:
		res = any(new(Documentation)).(Type)
	case *Element:
		res = any(new(Element)).(Type)
	case *Enumeration:
		res = any(new(Enumeration)).(Type)
	case *Extension:
		res = any(new(Extension)).(Type)
	case *Group:
		res = any(new(Group)).(Type)
	case *Length:
		res = any(new(Length)).(Type)
	case *MaxInclusive:
		res = any(new(MaxInclusive)).(Type)
	case *MaxLength:
		res = any(new(MaxLength)).(Type)
	case *MinInclusive:
		res = any(new(MinInclusive)).(Type)
	case *MinLength:
		res = any(new(MinLength)).(Type)
	case *Pattern:
		res = any(new(Pattern)).(Type)
	case *Restriction:
		res = any(new(Restriction)).(Type)
	case *Schema:
		res = any(new(Schema)).(Type)
	case *Sequence:
		res = any(new(Sequence)).(Type)
	case *SimpleContent:
		res = any(new(SimpleContent)).(Type)
	case *SimpleType:
		res = any(new(SimpleType)).(Type)
	case *TotalDigit:
		res = any(new(TotalDigit)).(Type)
	case *Union:
		res = any(new(Union)).(Type)
	case *WhiteSpace:
		res = any(new(WhiteSpace)).(Type)
	}
	return res
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

func GetPointerToGongstructName[Type GongstructIF]() (res string) {
	return GongGetPointerToGongstructName[Type]()
}

type GongReverseField struct {
	GongstructName string
	Fieldname      string
}

type ReverseField = GongReverseField

func GongGetReverseFields[Type GongstructIF]() (res []GongReverseField) {
	res = make([]GongReverseField, 0)

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

func GetReverseFields[Type GongstructIF]() (res []GongReverseField) {
	return GongGetReverseFields[Type]()
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

func (stage *Stage) GetFieldStringValueFromPointer(instance GongstructIF, fieldName string) (res GongFieldValue) {
	res = instance.GongGetFieldValue(fieldName, stage)
	return
}

func GetFieldStringValueFromPointer(instance GongstructIF, fieldName string, stage *Stage) (res GongFieldValue) {
	return stage.GetFieldStringValueFromPointer(instance, fieldName)
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

func GongGetGongstructNameFromPointer(instance GongstructIF) (res string) {
	res = instance.GongGetGongstructName()
	return
}

func GetGongstructNameFromPointer(instance GongstructIF) (res string) {
	return GongGetGongstructNameFromPointer(instance)
}

func (stage *Stage) ResetMapStrings() {
	// insertion point for generic get gongstruct name
	__gong__rebuildMapString(stage.Alls, &stage.Alls_mapString)

	__gong__rebuildMapString(stage.Annotations, &stage.Annotations_mapString)

	__gong__rebuildMapString(stage.Attributes, &stage.Attributes_mapString)

	__gong__rebuildMapString(stage.AttributeGroups, &stage.AttributeGroups_mapString)

	__gong__rebuildMapString(stage.Choices, &stage.Choices_mapString)

	__gong__rebuildMapString(stage.ComplexContents, &stage.ComplexContents_mapString)

	__gong__rebuildMapString(stage.ComplexTypes, &stage.ComplexTypes_mapString)

	__gong__rebuildMapString(stage.Documentations, &stage.Documentations_mapString)

	__gong__rebuildMapString(stage.Elements, &stage.Elements_mapString)

	__gong__rebuildMapString(stage.Enumerations, &stage.Enumerations_mapString)

	__gong__rebuildMapString(stage.Extensions, &stage.Extensions_mapString)

	__gong__rebuildMapString(stage.Groups, &stage.Groups_mapString)

	__gong__rebuildMapString(stage.Lengths, &stage.Lengths_mapString)

	__gong__rebuildMapString(stage.MaxInclusives, &stage.MaxInclusives_mapString)

	__gong__rebuildMapString(stage.MaxLengths, &stage.MaxLengths_mapString)

	__gong__rebuildMapString(stage.MinInclusives, &stage.MinInclusives_mapString)

	__gong__rebuildMapString(stage.MinLengths, &stage.MinLengths_mapString)

	__gong__rebuildMapString(stage.Patterns, &stage.Patterns_mapString)

	__gong__rebuildMapString(stage.Restrictions, &stage.Restrictions_mapString)

	__gong__rebuildMapString(stage.Schemas, &stage.Schemas_mapString)

	__gong__rebuildMapString(stage.Sequences, &stage.Sequences_mapString)

	__gong__rebuildMapString(stage.SimpleContents, &stage.SimpleContents_mapString)

	__gong__rebuildMapString(stage.SimpleTypes, &stage.SimpleTypes_mapString)

	__gong__rebuildMapString(stage.TotalDigits, &stage.TotalDigits_mapString)

	__gong__rebuildMapString(stage.Unions, &stage.Unions_mapString)

	__gong__rebuildMapString(stage.WhiteSpaces, &stage.WhiteSpaces_mapString)

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
