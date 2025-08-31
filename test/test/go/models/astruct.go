package models

import (
	"log"
	"time"
)

// Astruct demonstrate basic gong features
type Astruct struct {

	// a "Name" field is necessary to generate a GongStruct
	Name string

	//gong:meta
	Field any

	// AnonymousStructField1 struct {
	// 	// Timevar time.Time
	// 	// Bollean bool

	// 	// Associationtob4 *Bstruct
	// 	TheName1 string
	// 	// TheName2        string
	// 	// SliceOfB4       []*Bstruct
	// 	// Duration2       time.Duration
	// }

	// AnonymousStructField2 struct {
	// 	TheName1 string
	// }

	// ONE-ZERO/ONE association is a supported type of gong
	Associationtob *Bstruct

	// ONE-MANY association is a supported type of gong
	// not that that in go, this is a MANY-MANY association. In gong, the association
	// is implemented in the destination of the association, therefore, a Bclass instance
	// is related to ZERO/ONE Aclass instance by Arrayofb field
	Anarrayofb []*Bstruct

	//gong:type this magic node means that the object points to a type
	// This means that when a branch is copied, the duplicate meachanism
	// will stop here and only duplicate the pointer, not the type
	Anotherassociationtob_2 *Bstruct

	// time.Time is a supported type of gong
	Date time.Time

	// gong:bespoketimeserializeformat "2006-01-02"
	Date2 time.Time

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

	// TextFieldBespokeSize is a field with a larger form field
	//gong:width 600 gong:height 300
	TextFieldBespokeSize string

	//gong:text, magic code to have the form editor have a text area instead of an input
	TextArea string
}

// presence of this function will let gong generates orchestrator code
func (astruct *Astruct) OnAfterUpdate(stage *Stage, stagedInstance, frontInstance *Astruct) {

	//
}

func (astruct *Astruct) OnAfterUpdateWithMouseEvent(stage *Stage, frontInstance *Astruct, mouseEvent *Gong__MouseEvent) {

	log.Println(astruct.Name, "Received a an update with a mouse event. Shift Key, ", mouseEvent.ShiftKey)
}
