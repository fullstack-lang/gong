package models

const ModelsFileTemplate = `// generated code - do not edit
package models

import "encoding/xml"

// to avoid compilation error if no xml element is generated
var _ xml.Attr` +
	`

// insertion point for enums declaration{{` + string(rune(Level0AllGongenumsCode)) + `}}` +
	`

	// insertion point for gongstructs declarations{{` + string(rune(Level0AllGongstructsCode)) + `}}`

type Level0 int

const (
	Level0AllGongstructsCode Level0 = iota
	Level0AllGongenumsCode
	Level0Nb
)

var Level0Code map[Level0]string = // new line
map[Level0]string{
	Level0AllGongstructsCode: ``,
}

type Level1 int

const (
	Level1NamedStructCode Level1 = iota
	Level1UnNamedStructCode
	Level1NamedEnumCode
	Level1Nb
)

var Level1Code map[Level1]string = // new line
map[Level1]string{
	Level1NamedStructCode: `

// {{` + string(rune(Level2Structname)) +
		`}} Named source {{` + string(rune(Level2Source)) +
		`}}` + `{{` + string(rune(Level2Comment)) + `}}` +
		`
type {{` + string(rune(Level2Structname)) + `}} struct {
	Name string ` + "`" + "xml:\"-\"" + "`" + `

	// insertion point for fields{{` + string(rune(Level2Fields)) + `}}
}`,
	Level1UnNamedStructCode: `

// {{` + string(rune(Level2Structname)) +
		`}} UnNamed source {{` + string(rune(Level2Source)) +
		`}}
type {{` + string(rune(Level2Structname)) + `}} struct {

	// insertion point for fields{{` + string(rune(Level2Fields)) + `}}
}`,

	Level1NamedEnumCode: `` +

		`

// {{` + string(rune(Level2EnumComment)) + `}}` +
		`
type {{` + string(rune(Level2Enumname)) + `}} string

const ({{` + string(rune(Level2EnumValues)) + `}}
)`,
}

type Level2 int

const (
	Level2Structname Level2 = iota
	Level2Comment
	Level2Source
	Level2Fields

	Level2Enumname
	Level2EnumComment
	Level2EnumValues

	Level2Nb
)

var Level2Code map[Level2]string = // new line
map[Level2]string{
	Level2Structname: ``,
	Level2Comment:    ``,
	Level2Source:     ``,
	Level2Fields:     ``,
	Level2Enumname:   ``,
	Level2EnumValues: ``,
}

type Level3 int

const (
	Level3EnumValueIdentifier Level3 = iota
	Level3EnumValueString
	Level3Nb
)
