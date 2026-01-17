// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour

var _ = __GONG_time_The_fool_doth_think_he_is_wise__

// insertion point
type Body_WOP struct {
	// insertion point

	Name string
}

func (from *Body) CopyBasicFields(to *Body) {
	// insertion point
	to.Name = from.Name
}

type Document_WOP struct {
	// insertion point

	Name string
}

func (from *Document) CopyBasicFields(to *Document) {
	// insertion point
	to.Name = from.Name
}

type Docx_WOP struct {
	// insertion point

	Name string
}

func (from *Docx) CopyBasicFields(to *Docx) {
	// insertion point
	to.Name = from.Name
}

type File_WOP struct {
	// insertion point

	Name string
}

func (from *File) CopyBasicFields(to *File) {
	// insertion point
	to.Name = from.Name
}

type Node_WOP struct {
	// insertion point

	Name string
}

func (from *Node) CopyBasicFields(to *Node) {
	// insertion point
	to.Name = from.Name
}

type Paragraph_WOP struct {
	// insertion point

	Name string

	Content string

	CollatedText string
}

func (from *Paragraph) CopyBasicFields(to *Paragraph) {
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

func (from *ParagraphProperties) CopyBasicFields(to *ParagraphProperties) {
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

func (from *ParagraphStyle) CopyBasicFields(to *ParagraphStyle) {
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

func (from *Rune) CopyBasicFields(to *Rune) {
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

func (from *RuneProperties) CopyBasicFields(to *RuneProperties) {
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

func (from *Table) CopyBasicFields(to *Table) {
	// insertion point
	to.Name = from.Name
	to.Content = from.Content
}

type TableColumn_WOP struct {
	// insertion point

	Name string

	Content string
}

func (from *TableColumn) CopyBasicFields(to *TableColumn) {
	// insertion point
	to.Name = from.Name
	to.Content = from.Content
}

type TableProperties_WOP struct {
	// insertion point

	Name string

	Content string
}

func (from *TableProperties) CopyBasicFields(to *TableProperties) {
	// insertion point
	to.Name = from.Name
	to.Content = from.Content
}

type TableRow_WOP struct {
	// insertion point

	Name string

	Content string
}

func (from *TableRow) CopyBasicFields(to *TableRow) {
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

func (from *TableStyle) CopyBasicFields(to *TableStyle) {
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

func (from *Text) CopyBasicFields(to *Text) {
	// insertion point
	to.Name = from.Name
	to.Content = from.Content
	to.PreserveWhiteSpace = from.PreserveWhiteSpace
}

