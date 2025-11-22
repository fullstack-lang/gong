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

	markdown_go "github.com/fullstack-lang/gong/lib/markdown/go"
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
	Contents           map[*Content]any
	Contents_mapString map[string]*Content

	// insertion point for slice of pointers maps
	OnAfterContentCreateCallback OnAfterCreateInterface[Content]
	OnAfterContentUpdateCallback OnAfterUpdateInterface[Content]
	OnAfterContentDeleteCallback OnAfterDeleteInterface[Content]
	OnAfterContentReadCallback   OnAfterReadInterface[Content]

	JpgImages           map[*JpgImage]any
	JpgImages_mapString map[string]*JpgImage

	// insertion point for slice of pointers maps
	OnAfterJpgImageCreateCallback OnAfterCreateInterface[JpgImage]
	OnAfterJpgImageUpdateCallback OnAfterUpdateInterface[JpgImage]
	OnAfterJpgImageDeleteCallback OnAfterDeleteInterface[JpgImage]
	OnAfterJpgImageReadCallback   OnAfterReadInterface[JpgImage]

	PngImages           map[*PngImage]any
	PngImages_mapString map[string]*PngImage

	// insertion point for slice of pointers maps
	OnAfterPngImageCreateCallback OnAfterCreateInterface[PngImage]
	OnAfterPngImageUpdateCallback OnAfterUpdateInterface[PngImage]
	OnAfterPngImageDeleteCallback OnAfterDeleteInterface[PngImage]
	OnAfterPngImageReadCallback   OnAfterReadInterface[PngImage]

	SvgImages           map[*SvgImage]any
	SvgImages_mapString map[string]*SvgImage

	// insertion point for slice of pointers maps
	OnAfterSvgImageCreateCallback OnAfterCreateInterface[SvgImage]
	OnAfterSvgImageUpdateCallback OnAfterUpdateInterface[SvgImage]
	OnAfterSvgImageDeleteCallback OnAfterDeleteInterface[SvgImage]
	OnAfterSvgImageReadCallback   OnAfterReadInterface[SvgImage]

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
	ContentOrder            uint
	ContentMap_Staged_Order map[*Content]uint

	JpgImageOrder            uint
	JpgImageMap_Staged_Order map[*JpgImage]uint

	PngImageOrder            uint
	PngImageMap_Staged_Order map[*PngImage]uint

	SvgImageOrder            uint
	SvgImageMap_Staged_Order map[*SvgImage]uint

	// end of insertion point

	NamedStructs []*NamedStruct
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
	case *Content:
		tmp := GetStructInstancesByOrder(stage.Contents, stage.ContentMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Content implements.
			res = append(res, any(v).(T))
		}
		return res
	case *JpgImage:
		tmp := GetStructInstancesByOrder(stage.JpgImages, stage.JpgImageMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *JpgImage implements.
			res = append(res, any(v).(T))
		}
		return res
	case *PngImage:
		tmp := GetStructInstancesByOrder(stage.PngImages, stage.PngImageMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *PngImage implements.
			res = append(res, any(v).(T))
		}
		return res
	case *SvgImage:
		tmp := GetStructInstancesByOrder(stage.SvgImages, stage.SvgImageMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *SvgImage implements.
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
	case "Content":
		res = GetNamedStructInstances(stage.Contents, stage.ContentMap_Staged_Order)
	case "JpgImage":
		res = GetNamedStructInstances(stage.JpgImages, stage.JpgImageMap_Staged_Order)
	case "PngImage":
		res = GetNamedStructInstances(stage.PngImages, stage.PngImageMap_Staged_Order)
	case "SvgImage":
		res = GetNamedStructInstances(stage.SvgImages, stage.SvgImageMap_Staged_Order)
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
	return "github.com/fullstack-lang/gong/lib/markdown/go/models"
}

func (stage *Stage) GetMap_GongStructName_InstancesNb() map[string]int {
	return stage.Map_GongStructName_InstancesNb
}

func (stage *Stage) GetModelsEmbededDir() embed.FS {
	return markdown_go.GoModelsDir
}

func (stage *Stage) GetDigramsEmbededDir() embed.FS {
	return markdown_go.GoDiagramsDir
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
	CommitContent(content *Content)
	CheckoutContent(content *Content)
	CommitJpgImage(jpgimage *JpgImage)
	CheckoutJpgImage(jpgimage *JpgImage)
	CommitPngImage(pngimage *PngImage)
	CheckoutPngImage(pngimage *PngImage)
	CommitSvgImage(svgimage *SvgImage)
	CheckoutSvgImage(svgimage *SvgImage)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(name string) (stage *Stage) {

	stage = &Stage{ // insertion point for array initiatialisation
		Contents:           make(map[*Content]any),
		Contents_mapString: make(map[string]*Content),

		JpgImages:           make(map[*JpgImage]any),
		JpgImages_mapString: make(map[string]*JpgImage),

		PngImages:           make(map[*PngImage]any),
		PngImages_mapString: make(map[string]*PngImage),

		SvgImages:           make(map[*SvgImage]any),
		SvgImages_mapString: make(map[string]*SvgImage),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		name: name,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		ContentMap_Staged_Order: make(map[*Content]uint),

		JpgImageMap_Staged_Order: make(map[*JpgImage]uint),

		PngImageMap_Staged_Order: make(map[*PngImage]uint),

		SvgImageMap_Staged_Order: make(map[*SvgImage]uint),

		// end of insertion point

		NamedStructs: []*NamedStruct{ // insertion point for order map initialisations
			{name: "Content"},
			{name: "JpgImage"},
			{name: "PngImage"},
			{name: "SvgImage"},
		}, // end of insertion point
	}

	return
}

func GetOrder[Type Gongstruct](stage *Stage, instance *Type) uint {

	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *Content:
		return stage.ContentMap_Staged_Order[instance]
	case *JpgImage:
		return stage.JpgImageMap_Staged_Order[instance]
	case *PngImage:
		return stage.PngImageMap_Staged_Order[instance]
	case *SvgImage:
		return stage.SvgImageMap_Staged_Order[instance]
	default:
		return 0 // should not happen
	}
}

func GetOrderPointerGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) uint {

	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *Content:
		return stage.ContentMap_Staged_Order[instance]
	case *JpgImage:
		return stage.JpgImageMap_Staged_Order[instance]
	case *PngImage:
		return stage.PngImageMap_Staged_Order[instance]
	case *SvgImage:
		return stage.SvgImageMap_Staged_Order[instance]
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
	if stage.OnInitCommitCallback != nil {
		stage.OnInitCommitCallback.BeforeCommit(stage)
	}
	if stage.OnInitCommitFromBackCallback != nil {
		stage.OnInitCommitFromBackCallback.BeforeCommit(stage)
	}

	stage.ComputeReverseMaps()
	stage.commitId++
	stage.commitTimeStamp = time.Now()

	if stage.BackRepo != nil {
		stage.BackRepo.Commit(stage)
	}

	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["Content"] = len(stage.Contents)
	stage.Map_GongStructName_InstancesNb["JpgImage"] = len(stage.JpgImages)
	stage.Map_GongStructName_InstancesNb["PngImage"] = len(stage.PngImages)
	stage.Map_GongStructName_InstancesNb["SvgImage"] = len(stage.SvgImages)

}

func (stage *Stage) Checkout() {
	if stage.BackRepo != nil {
		stage.BackRepo.Checkout(stage)
	}

	stage.ComputeReverseMaps()
	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["Content"] = len(stage.Contents)
	stage.Map_GongStructName_InstancesNb["JpgImage"] = len(stage.JpgImages)
	stage.Map_GongStructName_InstancesNb["PngImage"] = len(stage.PngImages)
	stage.Map_GongStructName_InstancesNb["SvgImage"] = len(stage.SvgImages)

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
// Stage puts content to the model stage
func (content *Content) Stage(stage *Stage) *Content {

	if _, ok := stage.Contents[content]; !ok {
		stage.Contents[content] = __member
		stage.ContentMap_Staged_Order[content] = stage.ContentOrder
		stage.ContentOrder++
	}
	stage.Contents_mapString[content.Name] = content

	return content
}

// Unstage removes content off the model stage
func (content *Content) Unstage(stage *Stage) *Content {
	delete(stage.Contents, content)
	delete(stage.Contents_mapString, content.Name)
	return content
}

// UnstageVoid removes content off the model stage
func (content *Content) UnstageVoid(stage *Stage) {
	delete(stage.Contents, content)
	delete(stage.Contents_mapString, content.Name)
}

// commit content to the back repo (if it is already staged)
func (content *Content) Commit(stage *Stage) *Content {
	if _, ok := stage.Contents[content]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitContent(content)
		}
	}
	return content
}

func (content *Content) CommitVoid(stage *Stage) {
	content.Commit(stage)
}

// Checkout content to the back repo (if it is already staged)
func (content *Content) Checkout(stage *Stage) *Content {
	if _, ok := stage.Contents[content]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutContent(content)
		}
	}
	return content
}

// for satisfaction of GongStruct interface
func (content *Content) GetName() (res string) {
	return content.Name
}

// Stage puts jpgimage to the model stage
func (jpgimage *JpgImage) Stage(stage *Stage) *JpgImage {

	if _, ok := stage.JpgImages[jpgimage]; !ok {
		stage.JpgImages[jpgimage] = __member
		stage.JpgImageMap_Staged_Order[jpgimage] = stage.JpgImageOrder
		stage.JpgImageOrder++
	}
	stage.JpgImages_mapString[jpgimage.Name] = jpgimage

	return jpgimage
}

// Unstage removes jpgimage off the model stage
func (jpgimage *JpgImage) Unstage(stage *Stage) *JpgImage {
	delete(stage.JpgImages, jpgimage)
	delete(stage.JpgImages_mapString, jpgimage.Name)
	return jpgimage
}

// UnstageVoid removes jpgimage off the model stage
func (jpgimage *JpgImage) UnstageVoid(stage *Stage) {
	delete(stage.JpgImages, jpgimage)
	delete(stage.JpgImages_mapString, jpgimage.Name)
}

// commit jpgimage to the back repo (if it is already staged)
func (jpgimage *JpgImage) Commit(stage *Stage) *JpgImage {
	if _, ok := stage.JpgImages[jpgimage]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitJpgImage(jpgimage)
		}
	}
	return jpgimage
}

func (jpgimage *JpgImage) CommitVoid(stage *Stage) {
	jpgimage.Commit(stage)
}

// Checkout jpgimage to the back repo (if it is already staged)
func (jpgimage *JpgImage) Checkout(stage *Stage) *JpgImage {
	if _, ok := stage.JpgImages[jpgimage]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutJpgImage(jpgimage)
		}
	}
	return jpgimage
}

// for satisfaction of GongStruct interface
func (jpgimage *JpgImage) GetName() (res string) {
	return jpgimage.Name
}

// Stage puts pngimage to the model stage
func (pngimage *PngImage) Stage(stage *Stage) *PngImage {

	if _, ok := stage.PngImages[pngimage]; !ok {
		stage.PngImages[pngimage] = __member
		stage.PngImageMap_Staged_Order[pngimage] = stage.PngImageOrder
		stage.PngImageOrder++
	}
	stage.PngImages_mapString[pngimage.Name] = pngimage

	return pngimage
}

// Unstage removes pngimage off the model stage
func (pngimage *PngImage) Unstage(stage *Stage) *PngImage {
	delete(stage.PngImages, pngimage)
	delete(stage.PngImages_mapString, pngimage.Name)
	return pngimage
}

// UnstageVoid removes pngimage off the model stage
func (pngimage *PngImage) UnstageVoid(stage *Stage) {
	delete(stage.PngImages, pngimage)
	delete(stage.PngImages_mapString, pngimage.Name)
}

// commit pngimage to the back repo (if it is already staged)
func (pngimage *PngImage) Commit(stage *Stage) *PngImage {
	if _, ok := stage.PngImages[pngimage]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitPngImage(pngimage)
		}
	}
	return pngimage
}

func (pngimage *PngImage) CommitVoid(stage *Stage) {
	pngimage.Commit(stage)
}

// Checkout pngimage to the back repo (if it is already staged)
func (pngimage *PngImage) Checkout(stage *Stage) *PngImage {
	if _, ok := stage.PngImages[pngimage]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutPngImage(pngimage)
		}
	}
	return pngimage
}

// for satisfaction of GongStruct interface
func (pngimage *PngImage) GetName() (res string) {
	return pngimage.Name
}

// Stage puts svgimage to the model stage
func (svgimage *SvgImage) Stage(stage *Stage) *SvgImage {

	if _, ok := stage.SvgImages[svgimage]; !ok {
		stage.SvgImages[svgimage] = __member
		stage.SvgImageMap_Staged_Order[svgimage] = stage.SvgImageOrder
		stage.SvgImageOrder++
	}
	stage.SvgImages_mapString[svgimage.Name] = svgimage

	return svgimage
}

// Unstage removes svgimage off the model stage
func (svgimage *SvgImage) Unstage(stage *Stage) *SvgImage {
	delete(stage.SvgImages, svgimage)
	delete(stage.SvgImages_mapString, svgimage.Name)
	return svgimage
}

// UnstageVoid removes svgimage off the model stage
func (svgimage *SvgImage) UnstageVoid(stage *Stage) {
	delete(stage.SvgImages, svgimage)
	delete(stage.SvgImages_mapString, svgimage.Name)
}

// commit svgimage to the back repo (if it is already staged)
func (svgimage *SvgImage) Commit(stage *Stage) *SvgImage {
	if _, ok := stage.SvgImages[svgimage]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSvgImage(svgimage)
		}
	}
	return svgimage
}

func (svgimage *SvgImage) CommitVoid(stage *Stage) {
	svgimage.Commit(stage)
}

// Checkout svgimage to the back repo (if it is already staged)
func (svgimage *SvgImage) Checkout(stage *Stage) *SvgImage {
	if _, ok := stage.SvgImages[svgimage]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSvgImage(svgimage)
		}
	}
	return svgimage
}

// for satisfaction of GongStruct interface
func (svgimage *SvgImage) GetName() (res string) {
	return svgimage.Name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMContent(Content *Content)
	CreateORMJpgImage(JpgImage *JpgImage)
	CreateORMPngImage(PngImage *PngImage)
	CreateORMSvgImage(SvgImage *SvgImage)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMContent(Content *Content)
	DeleteORMJpgImage(JpgImage *JpgImage)
	DeleteORMPngImage(PngImage *PngImage)
	DeleteORMSvgImage(SvgImage *SvgImage)
}

func (stage *Stage) Reset() { // insertion point for array reset
	stage.Contents = make(map[*Content]any)
	stage.Contents_mapString = make(map[string]*Content)
	stage.ContentMap_Staged_Order = make(map[*Content]uint)
	stage.ContentOrder = 0

	stage.JpgImages = make(map[*JpgImage]any)
	stage.JpgImages_mapString = make(map[string]*JpgImage)
	stage.JpgImageMap_Staged_Order = make(map[*JpgImage]uint)
	stage.JpgImageOrder = 0

	stage.PngImages = make(map[*PngImage]any)
	stage.PngImages_mapString = make(map[string]*PngImage)
	stage.PngImageMap_Staged_Order = make(map[*PngImage]uint)
	stage.PngImageOrder = 0

	stage.SvgImages = make(map[*SvgImage]any)
	stage.SvgImages_mapString = make(map[string]*SvgImage)
	stage.SvgImageMap_Staged_Order = make(map[*SvgImage]uint)
	stage.SvgImageOrder = 0

}

func (stage *Stage) Nil() { // insertion point for array nil
	stage.Contents = nil
	stage.Contents_mapString = nil

	stage.JpgImages = nil
	stage.JpgImages_mapString = nil

	stage.PngImages = nil
	stage.PngImages_mapString = nil

	stage.SvgImages = nil
	stage.SvgImages_mapString = nil

}

func (stage *Stage) Unstage() { // insertion point for array nil
	for content := range stage.Contents {
		content.Unstage(stage)
	}

	for jpgimage := range stage.JpgImages {
		jpgimage.Unstage(stage)
	}

	for pngimage := range stage.PngImages {
		pngimage.Unstage(stage)
	}

	for svgimage := range stage.SvgImages {
		svgimage.Unstage(stage)
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
	case map[*Content]any:
		return any(&stage.Contents).(*Type)
	case map[*JpgImage]any:
		return any(&stage.JpgImages).(*Type)
	case map[*PngImage]any:
		return any(&stage.PngImages).(*Type)
	case map[*SvgImage]any:
		return any(&stage.SvgImages).(*Type)
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
	case map[string]*Content:
		return any(&stage.Contents_mapString).(*Type)
	case map[string]*JpgImage:
		return any(&stage.JpgImages_mapString).(*Type)
	case map[string]*PngImage:
		return any(&stage.PngImages_mapString).(*Type)
	case map[string]*SvgImage:
		return any(&stage.SvgImages_mapString).(*Type)
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
	case Content:
		return any(&stage.Contents).(*map[*Type]any)
	case JpgImage:
		return any(&stage.JpgImages).(*map[*Type]any)
	case PngImage:
		return any(&stage.PngImages).(*map[*Type]any)
	case SvgImage:
		return any(&stage.SvgImages).(*map[*Type]any)
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
	case *Content:
		return any(&stage.Contents).(*map[Type]any)
	case *JpgImage:
		return any(&stage.JpgImages).(*map[Type]any)
	case *PngImage:
		return any(&stage.PngImages).(*map[Type]any)
	case *SvgImage:
		return any(&stage.SvgImages).(*map[Type]any)
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
	case Content:
		return any(&stage.Contents_mapString).(*map[string]*Type)
	case JpgImage:
		return any(&stage.JpgImages_mapString).(*map[string]*Type)
	case PngImage:
		return any(&stage.PngImages_mapString).(*map[string]*Type)
	case SvgImage:
		return any(&stage.SvgImages_mapString).(*map[string]*Type)
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
	case Content:
		return any(&Content{
			// Initialisation of associations
		}).(*Type)
	case JpgImage:
		return any(&JpgImage{
			// Initialisation of associations
		}).(*Type)
	case PngImage:
		return any(&PngImage{
			// Initialisation of associations
		}).(*Type)
	case SvgImage:
		return any(&SvgImage{
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
	// reverse maps of direct associations of Content
	case Content:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of JpgImage
	case JpgImage:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of PngImage
	case PngImage:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SvgImage
	case SvgImage:
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
	// reverse maps of direct associations of Content
	case Content:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of JpgImage
	case JpgImage:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of PngImage
	case PngImage:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SvgImage
	case SvgImage:
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
	case *Content:
		res = "Content"
	case *JpgImage:
		res = "JpgImage"
	case *PngImage:
		res = "PngImage"
	case *SvgImage:
		res = "SvgImage"
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
	case *Content:
		var rf ReverseField
		_ = rf
	case *JpgImage:
		var rf ReverseField
		_ = rf
	case *PngImage:
		var rf ReverseField
		_ = rf
	case *SvgImage:
		var rf ReverseField
		_ = rf
	}
	return
}

// GetFieldsFromPointer return the array of the fields
func GetFieldsFromPointer[Type PointerToGongstruct]() (res []GongFieldHeader) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *Content:
		res = []GongFieldHeader{

			{
				Name:               "Name",
				GongFieldValueType: GongFieldValueTypeBasicKind,
			},
			{
				Name:               "Content",
				GongFieldValueType: GongFieldValueTypeBasicKind,
			},
		}
	case *JpgImage:
		res = []GongFieldHeader{

			{
				Name:               "Name",
				GongFieldValueType: GongFieldValueTypeBasicKind,
			},
			{
				Name:               "Base64Content",
				GongFieldValueType: GongFieldValueTypeBasicKind,
			},
		}
	case *PngImage:
		res = []GongFieldHeader{

			{
				Name:               "Name",
				GongFieldValueType: GongFieldValueTypeBasicKind,
			},
			{
				Name:               "Base64Content",
				GongFieldValueType: GongFieldValueTypeBasicKind,
			},
		}
	case *SvgImage:
		res = []GongFieldHeader{

			{
				Name:               "Name",
				GongFieldValueType: GongFieldValueTypeBasicKind,
			},
			{
				Name:               "Content",
				GongFieldValueType: GongFieldValueTypeBasicKind,
			},
		}
	}
	return
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

func GetFieldStringValueFromPointer(instance any, fieldName string, stage *Stage) (res GongFieldValue) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case *Content:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Content":
			res.valueString = inferedInstance.Content
		}
	case *JpgImage:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Base64Content":
			res.valueString = inferedInstance.Base64Content
		}
	case *PngImage:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Base64Content":
			res.valueString = inferedInstance.Base64Content
		}
	case *SvgImage:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Content":
			res.valueString = inferedInstance.Content
		}
	default:
		_ = inferedInstance
	}
	return
}

func GetFieldStringValue(instance any, fieldName string, stage *Stage) (res GongFieldValue) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case Content:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Content":
			res.valueString = inferedInstance.Content
		}
	case JpgImage:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Base64Content":
			res.valueString = inferedInstance.Base64Content
		}
	case PngImage:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Base64Content":
			res.valueString = inferedInstance.Base64Content
		}
	case SvgImage:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Content":
			res.valueString = inferedInstance.Content
		}
	default:
		_ = inferedInstance
	}
	return
}

// Last line of the template
