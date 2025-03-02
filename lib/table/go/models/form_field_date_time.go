package models

import "time"

type FormFieldDateTime struct {
	Name string

	Value time.Time
}

// GONGDOC(NoteOnFormFieldDateTime):
// [models.FormFieldDateTime]
// provides two form fields that will
// make the precision of go time.Time
const NoteOnFormFieldDateTime = ""
