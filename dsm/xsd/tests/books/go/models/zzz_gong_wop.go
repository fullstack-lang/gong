// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour

var _ = __GONG_time_The_fool_doth_think_he_is_wise__

// insertion point
type A_books_WOP struct {
	// insertion point

	Name string
}

func (from *A_books) CopyBasicFields(to *A_books) {
	// insertion point
	to.Name = from.Name
}

type BookType_WOP struct {
	// insertion point

	Name string

	Edition string

	Isbn string

	Bestseller bool

	Title string

	Author string

	Year int

	Format string
}

func (from *BookType) CopyBasicFields(to *BookType) {
	// insertion point
	to.Name = from.Name
	to.Edition = from.Edition
	to.Isbn = from.Isbn
	to.Bestseller = from.Bestseller
	to.Title = from.Title
	to.Author = from.Author
	to.Year = from.Year
	to.Format = from.Format
}

type Books_WOP struct {
	// insertion point

	Name string

	Name string
}

func (from *Books) CopyBasicFields(to *Books) {
	// insertion point
	to.Name = from.Name
	to.Name = from.Name
}

type Credit_WOP struct {
	// insertion point

	Name string

	Page int

	Credit_type string

	Credit_words string

	Credit_symbol string
}

func (from *Credit) CopyBasicFields(to *Credit) {
	// insertion point
	to.Name = from.Name
	to.Page = from.Page
	to.Credit_type = from.Credit_type
	to.Credit_words = from.Credit_words
	to.Credit_symbol = from.Credit_symbol
}

type Link_WOP struct {
	// insertion point

	Name string

	NameXSD string

	EnclosedText string
}

func (from *Link) CopyBasicFields(to *Link) {
	// insertion point
	to.Name = from.Name
	to.NameXSD = from.NameXSD
	to.EnclosedText = from.EnclosedText
}

