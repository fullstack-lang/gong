package models

// Book describes a book
// swagger:model Book
type Book struct {

	// The Name of the book
	Name string
	// The Author of the book
	Author string
	// The City of the publisheer
	City string
	// The Year of publication
	Year int
	// Price of the Book
	Price float64
	// Recommanded
	Recommanded bool

	Area *Area
}
