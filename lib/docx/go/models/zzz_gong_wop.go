// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var _ = time.Hour

// insertion point
type Body_WOP struct {
	// insertion point

	Name string
}

func (from *Body) GongCopyBasicFields(to *Body) {
	// insertion point
	to.Name = from.Name
}

type Document_WOP struct {
	// insertion point

	Name string
}

func (from *Document) GongCopyBasicFields(to *Document) {
	// insertion point
	to.Name = from.Name
}

type Docx_WOP struct {
	// insertion point

	Name string
}

func (from *Docx) GongCopyBasicFields(to *Docx) {
	// insertion point
	to.Name = from.Name
}

type File_WOP struct {
	// insertion point

	Name string
}

func (from *File) GongCopyBasicFields(to *File) {
	// insertion point
	*to = *from
}

type Node_WOP struct {
	// insertion point

	Name string
}

func (from *Node) GongCopyBasicFields(to *Node) {
	// insertion point
	to.Name = from.Name
}

type Paragraph_WOP struct {
	// insertion point

	Name string

	Content string

	CollatedText string
}

func (from *Paragraph) GongCopyBasicFields(to *Paragraph) {
	// insertion point
	to.Name = from.Name
	to.Content = from.Content
	to.CollatedText = from.CollatedText
}

type ParagraphProperties_WOP struct {
	// insertion point

	Name string

	Content string
}

func (from *ParagraphProperties) GongCopyBasicFields(to *ParagraphProperties) {
	// insertion point
	to.Name = from.Name
	to.Content = from.Content
}

type ParagraphStyle_WOP struct {
	// insertion point

	Name string

	Content string

	ValAttr string
}

func (from *ParagraphStyle) GongCopyBasicFields(to *ParagraphStyle) {
	// insertion point
	to.Name = from.Name
	to.Content = from.Content
	to.ValAttr = from.ValAttr
}

type Rune_WOP struct {
	// insertion point

	Name string

	Content string
}

func (from *Rune) GongCopyBasicFields(to *Rune) {
	// insertion point
	to.Name = from.Name
	to.Content = from.Content
}

type RuneProperties_WOP struct {
	// insertion point

	Name string

	IsBold bool

	IsStrike bool

	IsItalic bool

	Content string
}

func (from *RuneProperties) GongCopyBasicFields(to *RuneProperties) {
	// insertion point
	to.Name = from.Name
	to.IsBold = from.IsBold
	to.IsStrike = from.IsStrike
	to.IsItalic = from.IsItalic
	to.Content = from.Content
}

type Table_WOP struct {
	// insertion point

	Name string

	Content string
}

func (from *Table) GongCopyBasicFields(to *Table) {
	// insertion point
	to.Name = from.Name
	to.Content = from.Content
}

type TableColumn_WOP struct {
	// insertion point

	Name string

	Content string
}

func (from *TableColumn) GongCopyBasicFields(to *TableColumn) {
	// insertion point
	to.Name = from.Name
	to.Content = from.Content
}

type TableProperties_WOP struct {
	// insertion point

	Name string

	Content string
}

func (from *TableProperties) GongCopyBasicFields(to *TableProperties) {
	// insertion point
	to.Name = from.Name
	to.Content = from.Content
}

type TableRow_WOP struct {
	// insertion point

	Name string

	Content string
}

func (from *TableRow) GongCopyBasicFields(to *TableRow) {
	// insertion point
	to.Name = from.Name
	to.Content = from.Content
}

type TableStyle_WOP struct {
	// insertion point

	Name string

	Content string

	Val string
}

func (from *TableStyle) GongCopyBasicFields(to *TableStyle) {
	// insertion point
	to.Name = from.Name
	to.Content = from.Content
	to.Val = from.Val
}

type Text_WOP struct {
	// insertion point

	Name string

	Content string

	PreserveWhiteSpace bool
}

func (from *Text) GongCopyBasicFields(to *Text) {
	// insertion point
	to.Name = from.Name
	to.Content = from.Content
	to.PreserveWhiteSpace = from.PreserveWhiteSpace
}

// end of insertion point
