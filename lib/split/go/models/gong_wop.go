// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour

// insertion point
type AsSplit_WOP struct {
	// insertion point
	Name string
	Direction Direction
}

func (from *AsSplit) CopyBasicFields(to *AsSplit) {
	// insertion point
	to.Name = from.Name
	to.Direction = from.Direction
}

type AsSplitArea_WOP struct {
	// insertion point
	Name string
	Size float64
	IsAny bool
}

func (from *AsSplitArea) CopyBasicFields(to *AsSplitArea) {
	// insertion point
	to.Name = from.Name
	to.Size = from.Size
	to.IsAny = from.IsAny
}

type Table_WOP struct {
	// insertion point
	Name string
	StackName string
	TableName string
}

func (from *Table) CopyBasicFields(to *Table) {
	// insertion point
	to.Name = from.Name
	to.StackName = from.StackName
	to.TableName = from.TableName
}

type Tree_WOP struct {
	// insertion point
	Name string
	StackName string
	TreeName string
}

func (from *Tree) CopyBasicFields(to *Tree) {
	// insertion point
	to.Name = from.Name
	to.StackName = from.StackName
	to.TreeName = from.TreeName
}

type View_WOP struct {
	// insertion point
	Name string
}

func (from *View) CopyBasicFields(to *View) {
	// insertion point
	to.Name = from.Name
}

