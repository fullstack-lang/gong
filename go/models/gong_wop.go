// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour

// insertion point
type GongBasicField_WOP struct {
	// insertion point
	Name string
	BasicKindName string
	DeclaredType string
	CompositeStructName string
	Index int
	IsDocLink bool
	IsTextArea bool
	IsBespokeWidth bool
	BespokeWidth int
	IsBespokeHeight bool
	BespokeHeight int
}

func (from *GongBasicField) CopyBasicFields(to *GongBasicField) {
	// insertion point
	to.Name = from.Name
	to.BasicKindName = from.BasicKindName
	to.DeclaredType = from.DeclaredType
	to.CompositeStructName = from.CompositeStructName
	to.Index = from.Index
	to.IsDocLink = from.IsDocLink
	to.IsTextArea = from.IsTextArea
	to.IsBespokeWidth = from.IsBespokeWidth
	to.BespokeWidth = from.BespokeWidth
	to.IsBespokeHeight = from.IsBespokeHeight
	to.BespokeHeight = from.BespokeHeight
}

type GongEnum_WOP struct {
	// insertion point
	Name string
	Type GongEnumType
}

func (from *GongEnum) CopyBasicFields(to *GongEnum) {
	// insertion point
	to.Name = from.Name
	to.Type = from.Type
}

type GongEnumValue_WOP struct {
	// insertion point
	Name string
	Value string
}

func (from *GongEnumValue) CopyBasicFields(to *GongEnumValue) {
	// insertion point
	to.Name = from.Name
	to.Value = from.Value
}

type GongLink_WOP struct {
	// insertion point
	Name string
	Recv string
	ImportPath string
}

func (from *GongLink) CopyBasicFields(to *GongLink) {
	// insertion point
	to.Name = from.Name
	to.Recv = from.Recv
	to.ImportPath = from.ImportPath
}

type GongNote_WOP struct {
	// insertion point
	Name string
	Body string
	BodyHTML string
}

func (from *GongNote) CopyBasicFields(to *GongNote) {
	// insertion point
	to.Name = from.Name
	to.Body = from.Body
	to.BodyHTML = from.BodyHTML
}

type GongStruct_WOP struct {
	// insertion point
	Name string
	HasOnAfterUpdateSignature bool
	IsIgnoredForFront bool
}

func (from *GongStruct) CopyBasicFields(to *GongStruct) {
	// insertion point
	to.Name = from.Name
	to.HasOnAfterUpdateSignature = from.HasOnAfterUpdateSignature
	to.IsIgnoredForFront = from.IsIgnoredForFront
}

type GongTimeField_WOP struct {
	// insertion point
	Name string
	Index int
	CompositeStructName string
}

func (from *GongTimeField) CopyBasicFields(to *GongTimeField) {
	// insertion point
	to.Name = from.Name
	to.Index = from.Index
	to.CompositeStructName = from.CompositeStructName
}

type Meta_WOP struct {
	// insertion point
	Name string
	Text string
}

func (from *Meta) CopyBasicFields(to *Meta) {
	// insertion point
	to.Name = from.Name
	to.Text = from.Text
}

type MetaReference_WOP struct {
	// insertion point
	Name string
}

func (from *MetaReference) CopyBasicFields(to *MetaReference) {
	// insertion point
	to.Name = from.Name
}

type ModelPkg_WOP struct {
	// insertion point
	Name string
	PkgPath string
	PathToGoSubDirectory string
	OrmPkgGenPath string
	DbOrmPkgGenPath string
	DbLiteOrmPkgGenPath string
	DbPkgGenPath string
	ControllersPkgGenPath string
	FullstackPkgGenPath string
	StackPkgGenPath string
	StaticPkgGenPath string
	ProbePkgGenPath string
	NgWorkspacePath string
	NgWorkspaceName string
	NgDataLibrarySourceCodeDirectory string
	NgSpecificLibrarySourceCodeDirectory string
	MaterialLibDatamodelTargetPath string
}

func (from *ModelPkg) CopyBasicFields(to *ModelPkg) {
	// insertion point
	to.Name = from.Name
	to.PkgPath = from.PkgPath
	to.PathToGoSubDirectory = from.PathToGoSubDirectory
	to.OrmPkgGenPath = from.OrmPkgGenPath
	to.DbOrmPkgGenPath = from.DbOrmPkgGenPath
	to.DbLiteOrmPkgGenPath = from.DbLiteOrmPkgGenPath
	to.DbPkgGenPath = from.DbPkgGenPath
	to.ControllersPkgGenPath = from.ControllersPkgGenPath
	to.FullstackPkgGenPath = from.FullstackPkgGenPath
	to.StackPkgGenPath = from.StackPkgGenPath
	to.StaticPkgGenPath = from.StaticPkgGenPath
	to.ProbePkgGenPath = from.ProbePkgGenPath
	to.NgWorkspacePath = from.NgWorkspacePath
	to.NgWorkspaceName = from.NgWorkspaceName
	to.NgDataLibrarySourceCodeDirectory = from.NgDataLibrarySourceCodeDirectory
	to.NgSpecificLibrarySourceCodeDirectory = from.NgSpecificLibrarySourceCodeDirectory
	to.MaterialLibDatamodelTargetPath = from.MaterialLibDatamodelTargetPath
}

type PointerToGongStructField_WOP struct {
	// insertion point
	Name string
	Index int
	CompositeStructName string
	IsType bool
}

func (from *PointerToGongStructField) CopyBasicFields(to *PointerToGongStructField) {
	// insertion point
	to.Name = from.Name
	to.Index = from.Index
	to.CompositeStructName = from.CompositeStructName
	to.IsType = from.IsType
}

type SliceOfPointerToGongStructField_WOP struct {
	// insertion point
	Name string
	Index int
	CompositeStructName string
}

func (from *SliceOfPointerToGongStructField) CopyBasicFields(to *SliceOfPointerToGongStructField) {
	// insertion point
	to.Name = from.Name
	to.Index = from.Index
	to.CompositeStructName = from.CompositeStructName
}

