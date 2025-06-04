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

	load_go "github.com/fullstack-lang/gong/lib/load/go"
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
	name string

	// insertion point for definition of arrays registering instances
	FileToDownloads           map[*FileToDownload]any
	FileToDownloads_mapString map[string]*FileToDownload

	// insertion point for slice of pointers maps
	OnAfterFileToDownloadCreateCallback OnAfterCreateInterface[FileToDownload]
	OnAfterFileToDownloadUpdateCallback OnAfterUpdateInterface[FileToDownload]
	OnAfterFileToDownloadDeleteCallback OnAfterDeleteInterface[FileToDownload]
	OnAfterFileToDownloadReadCallback   OnAfterReadInterface[FileToDownload]

	FileToUploads           map[*FileToUpload]any
	FileToUploads_mapString map[string]*FileToUpload

	// insertion point for slice of pointers maps
	OnAfterFileToUploadCreateCallback OnAfterCreateInterface[FileToUpload]
	OnAfterFileToUploadUpdateCallback OnAfterUpdateInterface[FileToUpload]
	OnAfterFileToUploadDeleteCallback OnAfterDeleteInterface[FileToUpload]
	OnAfterFileToUploadReadCallback   OnAfterReadInterface[FileToUpload]

	Messages           map[*Message]any
	Messages_mapString map[string]*Message

	// insertion point for slice of pointers maps
	OnAfterMessageCreateCallback OnAfterCreateInterface[Message]
	OnAfterMessageUpdateCallback OnAfterUpdateInterface[Message]
	OnAfterMessageDeleteCallback OnAfterDeleteInterface[Message]
	OnAfterMessageReadCallback   OnAfterReadInterface[Message]

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
	FileToDownloadOrder            uint
	FileToDownloadMap_Staged_Order map[*FileToDownload]uint

	FileToUploadOrder            uint
	FileToUploadMap_Staged_Order map[*FileToUpload]uint

	MessageOrder            uint
	MessageMap_Staged_Order map[*Message]uint

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

	for _, instance := range orderedSet {
		res = append(res, instance)
	}

	return
}

func (stage *Stage) GetNamedStructNamesByOrder(namedStructName string) (res []string) {

	switch namedStructName {
	// insertion point for case
	case "FileToDownload":
		res = GetNamedStructInstances(stage.FileToDownloads, stage.FileToDownloadMap_Staged_Order)
	case "FileToUpload":
		res = GetNamedStructInstances(stage.FileToUploads, stage.FileToUploadMap_Staged_Order)
	case "Message":
		res = GetNamedStructInstances(stage.Messages, stage.MessageMap_Staged_Order)
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
	return "github.com/fullstack-lang/gong/lib/load/go/models"
}

func (stage *Stage) GetMap_GongStructName_InstancesNb() map[string]int {
	return stage.Map_GongStructName_InstancesNb
}

func (stage *Stage) GetModelsEmbededDir() embed.FS {
	return load_go.GoModelsDir
}

func (stage *Stage) GetDigramsEmbededDir() embed.FS {
	return load_go.GoDiagramsDir
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
	CommitFileToDownload(filetodownload *FileToDownload)
	CheckoutFileToDownload(filetodownload *FileToDownload)
	CommitFileToUpload(filetoupload *FileToUpload)
	CheckoutFileToUpload(filetoupload *FileToUpload)
	CommitMessage(message *Message)
	CheckoutMessage(message *Message)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(name string) (stage *Stage) {

	stage = &Stage{ // insertion point for array initiatialisation
		FileToDownloads:           make(map[*FileToDownload]any),
		FileToDownloads_mapString: make(map[string]*FileToDownload),

		FileToUploads:           make(map[*FileToUpload]any),
		FileToUploads_mapString: make(map[string]*FileToUpload),

		Messages:           make(map[*Message]any),
		Messages_mapString: make(map[string]*Message),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		name: name,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		FileToDownloadMap_Staged_Order: make(map[*FileToDownload]uint),

		FileToUploadMap_Staged_Order: make(map[*FileToUpload]uint),

		MessageMap_Staged_Order: make(map[*Message]uint),

		// end of insertion point

		NamedStructs: []*NamedStruct{ // insertion point for order map initialisations
			{name: "FileToDownload"},
			{name: "FileToUpload"},
			{name: "Message"},
		}, // end of insertion point
	}

	return
}

func GetOrder[Type Gongstruct](stage *Stage, instance *Type) uint {

	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *FileToDownload:
		return stage.FileToDownloadMap_Staged_Order[instance]
	case *FileToUpload:
		return stage.FileToUploadMap_Staged_Order[instance]
	case *Message:
		return stage.MessageMap_Staged_Order[instance]
	default:
		return 0 // should not happen
	}
}

func GetOrderPointerGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) uint {

	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *FileToDownload:
		return stage.FileToDownloadMap_Staged_Order[instance]
	case *FileToUpload:
		return stage.FileToUploadMap_Staged_Order[instance]
	case *Message:
		return stage.MessageMap_Staged_Order[instance]
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
	stage.Map_GongStructName_InstancesNb["FileToDownload"] = len(stage.FileToDownloads)
	stage.Map_GongStructName_InstancesNb["FileToUpload"] = len(stage.FileToUploads)
	stage.Map_GongStructName_InstancesNb["Message"] = len(stage.Messages)

}

func (stage *Stage) Checkout() {
	if stage.BackRepo != nil {
		stage.BackRepo.Checkout(stage)
	}

	stage.ComputeReverseMaps()
	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["FileToDownload"] = len(stage.FileToDownloads)
	stage.Map_GongStructName_InstancesNb["FileToUpload"] = len(stage.FileToUploads)
	stage.Map_GongStructName_InstancesNb["Message"] = len(stage.Messages)

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
// Stage puts filetodownload to the model stage
func (filetodownload *FileToDownload) Stage(stage *Stage) *FileToDownload {

	if _, ok := stage.FileToDownloads[filetodownload]; !ok {
		stage.FileToDownloads[filetodownload] = __member
		stage.FileToDownloadMap_Staged_Order[filetodownload] = stage.FileToDownloadOrder
		stage.FileToDownloadOrder++
	}
	stage.FileToDownloads_mapString[filetodownload.Name] = filetodownload

	return filetodownload
}

// Unstage removes filetodownload off the model stage
func (filetodownload *FileToDownload) Unstage(stage *Stage) *FileToDownload {
	delete(stage.FileToDownloads, filetodownload)
	delete(stage.FileToDownloads_mapString, filetodownload.Name)
	return filetodownload
}

// UnstageVoid removes filetodownload off the model stage
func (filetodownload *FileToDownload) UnstageVoid(stage *Stage) {
	delete(stage.FileToDownloads, filetodownload)
	delete(stage.FileToDownloads_mapString, filetodownload.Name)
}

// commit filetodownload to the back repo (if it is already staged)
func (filetodownload *FileToDownload) Commit(stage *Stage) *FileToDownload {
	if _, ok := stage.FileToDownloads[filetodownload]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitFileToDownload(filetodownload)
		}
	}
	return filetodownload
}

func (filetodownload *FileToDownload) CommitVoid(stage *Stage) {
	filetodownload.Commit(stage)
}

// Checkout filetodownload to the back repo (if it is already staged)
func (filetodownload *FileToDownload) Checkout(stage *Stage) *FileToDownload {
	if _, ok := stage.FileToDownloads[filetodownload]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutFileToDownload(filetodownload)
		}
	}
	return filetodownload
}

// for satisfaction of GongStruct interface
func (filetodownload *FileToDownload) GetName() (res string) {
	return filetodownload.Name
}

// Stage puts filetoupload to the model stage
func (filetoupload *FileToUpload) Stage(stage *Stage) *FileToUpload {

	if _, ok := stage.FileToUploads[filetoupload]; !ok {
		stage.FileToUploads[filetoupload] = __member
		stage.FileToUploadMap_Staged_Order[filetoupload] = stage.FileToUploadOrder
		stage.FileToUploadOrder++
	}
	stage.FileToUploads_mapString[filetoupload.Name] = filetoupload

	return filetoupload
}

// Unstage removes filetoupload off the model stage
func (filetoupload *FileToUpload) Unstage(stage *Stage) *FileToUpload {
	delete(stage.FileToUploads, filetoupload)
	delete(stage.FileToUploads_mapString, filetoupload.Name)
	return filetoupload
}

// UnstageVoid removes filetoupload off the model stage
func (filetoupload *FileToUpload) UnstageVoid(stage *Stage) {
	delete(stage.FileToUploads, filetoupload)
	delete(stage.FileToUploads_mapString, filetoupload.Name)
}

// commit filetoupload to the back repo (if it is already staged)
func (filetoupload *FileToUpload) Commit(stage *Stage) *FileToUpload {
	if _, ok := stage.FileToUploads[filetoupload]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitFileToUpload(filetoupload)
		}
	}
	return filetoupload
}

func (filetoupload *FileToUpload) CommitVoid(stage *Stage) {
	filetoupload.Commit(stage)
}

// Checkout filetoupload to the back repo (if it is already staged)
func (filetoupload *FileToUpload) Checkout(stage *Stage) *FileToUpload {
	if _, ok := stage.FileToUploads[filetoupload]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutFileToUpload(filetoupload)
		}
	}
	return filetoupload
}

// for satisfaction of GongStruct interface
func (filetoupload *FileToUpload) GetName() (res string) {
	return filetoupload.Name
}

// Stage puts message to the model stage
func (message *Message) Stage(stage *Stage) *Message {

	if _, ok := stage.Messages[message]; !ok {
		stage.Messages[message] = __member
		stage.MessageMap_Staged_Order[message] = stage.MessageOrder
		stage.MessageOrder++
	}
	stage.Messages_mapString[message.Name] = message

	return message
}

// Unstage removes message off the model stage
func (message *Message) Unstage(stage *Stage) *Message {
	delete(stage.Messages, message)
	delete(stage.Messages_mapString, message.Name)
	return message
}

// UnstageVoid removes message off the model stage
func (message *Message) UnstageVoid(stage *Stage) {
	delete(stage.Messages, message)
	delete(stage.Messages_mapString, message.Name)
}

// commit message to the back repo (if it is already staged)
func (message *Message) Commit(stage *Stage) *Message {
	if _, ok := stage.Messages[message]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitMessage(message)
		}
	}
	return message
}

func (message *Message) CommitVoid(stage *Stage) {
	message.Commit(stage)
}

// Checkout message to the back repo (if it is already staged)
func (message *Message) Checkout(stage *Stage) *Message {
	if _, ok := stage.Messages[message]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutMessage(message)
		}
	}
	return message
}

// for satisfaction of GongStruct interface
func (message *Message) GetName() (res string) {
	return message.Name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMFileToDownload(FileToDownload *FileToDownload)
	CreateORMFileToUpload(FileToUpload *FileToUpload)
	CreateORMMessage(Message *Message)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMFileToDownload(FileToDownload *FileToDownload)
	DeleteORMFileToUpload(FileToUpload *FileToUpload)
	DeleteORMMessage(Message *Message)
}

func (stage *Stage) Reset() { // insertion point for array reset
	stage.FileToDownloads = make(map[*FileToDownload]any)
	stage.FileToDownloads_mapString = make(map[string]*FileToDownload)
	stage.FileToDownloadMap_Staged_Order = make(map[*FileToDownload]uint)
	stage.FileToDownloadOrder = 0

	stage.FileToUploads = make(map[*FileToUpload]any)
	stage.FileToUploads_mapString = make(map[string]*FileToUpload)
	stage.FileToUploadMap_Staged_Order = make(map[*FileToUpload]uint)
	stage.FileToUploadOrder = 0

	stage.Messages = make(map[*Message]any)
	stage.Messages_mapString = make(map[string]*Message)
	stage.MessageMap_Staged_Order = make(map[*Message]uint)
	stage.MessageOrder = 0

}

func (stage *Stage) Nil() { // insertion point for array nil
	stage.FileToDownloads = nil
	stage.FileToDownloads_mapString = nil

	stage.FileToUploads = nil
	stage.FileToUploads_mapString = nil

	stage.Messages = nil
	stage.Messages_mapString = nil

}

func (stage *Stage) Unstage() { // insertion point for array nil
	for filetodownload := range stage.FileToDownloads {
		filetodownload.Unstage(stage)
	}

	for filetoupload := range stage.FileToUploads {
		filetoupload.Unstage(stage)
	}

	for message := range stage.Messages {
		message.Unstage(stage)
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
	case map[*FileToDownload]any:
		return any(&stage.FileToDownloads).(*Type)
	case map[*FileToUpload]any:
		return any(&stage.FileToUploads).(*Type)
	case map[*Message]any:
		return any(&stage.Messages).(*Type)
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
	case map[string]*FileToDownload:
		return any(&stage.FileToDownloads_mapString).(*Type)
	case map[string]*FileToUpload:
		return any(&stage.FileToUploads_mapString).(*Type)
	case map[string]*Message:
		return any(&stage.Messages_mapString).(*Type)
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
	case FileToDownload:
		return any(&stage.FileToDownloads).(*map[*Type]any)
	case FileToUpload:
		return any(&stage.FileToUploads).(*map[*Type]any)
	case Message:
		return any(&stage.Messages).(*map[*Type]any)
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
	case *FileToDownload:
		return any(&stage.FileToDownloads).(*map[Type]any)
	case *FileToUpload:
		return any(&stage.FileToUploads).(*map[Type]any)
	case *Message:
		return any(&stage.Messages).(*map[Type]any)
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
	case FileToDownload:
		return any(&stage.FileToDownloads_mapString).(*map[string]*Type)
	case FileToUpload:
		return any(&stage.FileToUploads_mapString).(*map[string]*Type)
	case Message:
		return any(&stage.Messages_mapString).(*map[string]*Type)
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
	case FileToDownload:
		return any(&FileToDownload{
			// Initialisation of associations
		}).(*Type)
	case FileToUpload:
		return any(&FileToUpload{
			// Initialisation of associations
		}).(*Type)
	case Message:
		return any(&Message{
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
	// reverse maps of direct associations of FileToDownload
	case FileToDownload:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FileToUpload
	case FileToUpload:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Message
	case Message:
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
	// reverse maps of direct associations of FileToDownload
	case FileToDownload:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FileToUpload
	case FileToUpload:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Message
	case Message:
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
	case FileToDownload:
		res = "FileToDownload"
	case FileToUpload:
		res = "FileToUpload"
	case Message:
		res = "Message"
	}
	return res
}

// GetPointerToGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetPointerToGongstructName[Type PointerToGongstruct]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *FileToDownload:
		res = "FileToDownload"
	case *FileToUpload:
		res = "FileToUpload"
	case *Message:
		res = "Message"
	}
	return res
}

// GetFields return the array of the fields
func GetFields[Type Gongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case FileToDownload:
		res = []string{"Name", "Content"}
	case FileToUpload:
		res = []string{"Name", "Base64EncodedContent"}
	case Message:
		res = []string{"Name"}
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
	case FileToDownload:
		var rf ReverseField
		_ = rf
	case FileToUpload:
		var rf ReverseField
		_ = rf
	case Message:
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
	case *FileToDownload:
		res = []string{"Name", "Content"}
	case *FileToUpload:
		res = []string{"Name", "Base64EncodedContent"}
	case *Message:
		res = []string{"Name"}
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
	case *FileToDownload:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Content":
			res.valueString = inferedInstance.Content
		}
	case *FileToUpload:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Base64EncodedContent":
			res.valueString = inferedInstance.Base64EncodedContent
		}
	case *Message:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	default:
		_ = inferedInstance
	}
	return
}

func GetFieldStringValue(instance any, fieldName string) (res GongFieldValue) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case FileToDownload:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Content":
			res.valueString = inferedInstance.Content
		}
	case FileToUpload:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Base64EncodedContent":
			res.valueString = inferedInstance.Base64EncodedContent
		}
	case Message:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	default:
		_ = inferedInstance
	}
	return
}

// Last line of the template
