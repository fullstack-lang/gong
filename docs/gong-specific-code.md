
Gong (go + ng) is a framework for rapid web application development (a.k.a. full stack development) based on go and angular. The go back-end uses [gin](https://github.com/gin-gonic/gin), [gorm](https://gorm.io/index.html) and sqlite (a pure go sqlite, no cgo needed). The angular front-end uses [angular material](https://material.angular.io/).

The unit of development in gong is the **gong stack** (a "stack" in the rest of this document). A stack can import other stacks (both the front end and the back end of a stack are integrated as a whole). The granularity of a stack is similar to an angular components.

Gong goal is that the developper facing complexity should be minimal (see [DHH's conceptual compression concept](https://m.signalvnoise.com/conceptual-compression-means-beginners-dont-need-to-know-sql-hallelujah/) or [Rob Pike's design of Go regarding complexity](https://go.dev/talks/2015/simplicity-is-complicated/).


# 2. Gong basics

There are different complexity levels for full stack development:

- level 1 : pure go code that uses gong front-end angular libs (should fit most needs).
- level 2 : go + angular code that combines gong libs and with *specific* angular code.

This chapter presents level 1 basics.

## 2.1. The model

The gong *element*s that forms a gong models are located in one package located in "go/models".

```go
// the package name shall be "models"
package models

// the gong models code
// ...
```

No particular import is necessary.

gong elements are *gongstruct*, *gongenum*, *gongfield* or *gongnote*.


## 2.2. A *gongstruct* is a named struct

```go
// Hello is a gongstruct
type Hello struct { // it is exported
  Name string // it has an exported "Name" field
}
```

## 2.3. A *gongenum* is a string/int enum


```go
// Gongenums can be of type tring or int.

type AEnumType string // It is exported

const (
    ENUM_VAL1 AEnumType = "ENUM_VAL1_NOT_THE_SAME"
    ENUM_VAL2 AEnumType = "ENUM_VAL2"
)

type CEnumTypeInt int

// values for EnumType
const (
    CENUM_VAL1 CEnumTypeInt = iota
    CENUM_VAL2
)
```

## 2.4. *gongfield*s are fields within a gongstruct

```go
// Hello is a gongstruct
type Hello struct { // it is exported
  Name string // it has an exported "Name" field

  // Those are gongfields. They are exported fields of type int, int64, float, float64 bool, string
  Floatfield float64
  Intfield int
  Booleanfield bool

  // Those are gongfields. They are exported fields of type time.Duration or time.Time
  Duration1 time.Duration
  Date time.Time

  // HelloP is a gongfield. It is an exported pointer to a gongstruct within the same package
  HelloP *Hello 

  // HellosP is a gongfield. It is an exported slice of pointers to a gongstruct within the same package
  HellosP []*Hello 

  // AEnum is a gongfield. It is an exported field of a gongenum type within the same pakage  
  AEnum AEnumType

  // CEnum is gongfield
  CEnum CEnumTypeInt
  
}
```

## 2.5. Gongnote

Gong provides easy, UML like, diagramming capability of the gong elements. You might want to
add notes to diagrams. Gongnotes are for this.

![example diagram](docs/images/example%20diagram.png)

```go
// GONGDOC(NoteOnOrganisation): Note on the organisation
//
// and [models.Bstruct],
// and [models.Astruct],
// and [models.Astruct],
// and [models.Astruct],
// and [models.Astruct.Associationtob],
// having the following const exported identifier allows for
// referencing the note from the UML note and allows for
// renaming
//
// # This is heading 1
//
// ## This is heading 1.1
//
//	-
const NoteOnOrganisation = ""
```
