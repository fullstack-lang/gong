package models

import (
	"time"
)

// Astruct demonstrate basic gong features
type Astruct struct {

	// a "Name" field is necessary to generate a GongStruct
	Name string

	// ONE-ZERO/ONE association is a supported type of gong
	Associationtob *Bstruct

	// ONE-MANY association is a supported type of gong
	// not that that in go, this is a MANY-MANY association. In gong, the association
	// is implemented in the destination of the association, therefore, a Bclass instance
	// is related to ZERO/ONE Aclass instance by Arrayofb field
	Anarrayofb []*Bstruct

	Anotherassociationtob_2 *Bstruct

	// time.Time is a supported type of gong
	Date time.Time

	// bool is a supported type of gong
	Booleanfield bool

	// enums is a supported type of gong (if they are string)
	Aenum   AEnumType
	Aenum_2 AEnumType
	Benum   BEnumType
	CEnum   CEnumTypeInt

	// an embedded struct is supported (without field name)
	Cstruct
	// another embedded struct
	Estruct

	// float64 is a supported type of gong
	Floatfield float64

	// int is a supported type of gong
	Intfield int

	Anotherbooleanfield bool

	Duration1 time.Duration

	// gong supports different fields with the same destination struct, even itself
	Anarrayofa []*Astruct

	Anotherarrayofb []*Bstruct

	// MANY-MANY association, (because AclassBclassUse ends with "Use")
	AnarrayofbUse  []*AstructBstructUse
	Anarrayofb2Use []*AstructBstruct2Use

	// pointer to itself
	AnAstruct *Astruct

	// StructRef is a string that can be set by a gong:ident directive followed by
	// a doclink declaration such as [pkg.Name1]

	// FieldRef for a field ref [pkg.Name1]
	//gong:ident
	StructRef string

	// FieldRef for a field ref [pkg.Name1.Name2]
	//gong:ident
	FieldRef string

	// docLinkRemapping "pkg.IntEnum(0)"
	//gong:ident [pkg.Name1]
	EnumIntRef string

	// docLinkRemapping "pkg.StringEnum("")"
	//gong:ident [pkg.Name1]
	EnumStringRef string

	// docLinkRemapping "pkg.Name1"
	//gong:ident [pkg.Name1]
	EnumValue string

	// docLinkRemapping "pkg.Name1"
	//gong:ident [pkg.Name1]
	ConstIdentifierValue string

	//gong:width 80
	TextFieldBespokeWidth string

	//gong:text, magic code to have the form editor have a text area instead of an input
	TextArea string
}

// presence of this function will let gong generates orchestrator code
func (astruct *Astruct) OnAfterUpdate(stage *StageStruct, stagedInstance, frontInstance *Astruct) {

	//
}
