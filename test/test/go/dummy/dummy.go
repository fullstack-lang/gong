package dummy

type Dummy struct {
	Name string
}

type DummyTypeString string

// values for EnumType
const (
	A DummyTypeString = "A"
	B DummyTypeString = "B"
)

type DummyTypeInt int

// values for EnumType
const (
	ZERO DummyTypeInt = iota
	ONE
)

// GONGNOTE(ShortNote): this is an example of a short note
// It uses the DocLink convention for referencing Identifiers
// In this case [Dummy]
// the following declaration allows for the renaming of the note
// in gongdoc, the note will be referenced by this identifier
const ShortNote = ""

type Stage struct {
	Name string
}
