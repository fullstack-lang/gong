
- [1. About Gong](#1-about-gong)
- [2. Gong basics](#2-gong-basics)
  - [2.1. The model](#21-the-model)
  - [2.2. A *gongstruct* is a named struct](#22-a-gongstruct-is-a-named-struct)
  - [2.3. A *gongenum* is a string/int enum](#23-a-gongenum-is-a-stringint-enum)
  - [2.4. *gongfield*s are fields within a gongstruct](#24-gongfields-are-fields-within-a-gongstruct)
  - [2.5. Gongnote](#25-gongnote)
- [3. A "hello world" stack in 5 minutes](#3-a-hello-world-stack-in-5-minutes)
  - [3.1. Prerequisite](#31-prerequisite)
  - [3.2. Generating the code with the `gongc` command](#32-generating-the-code-with-the-gongc-command)
  - [3.3. Injecting Data via REST](#33-injecting-data-via-rest)


# 1. About Gong

![gong logo](docs/images/gong%20logo.svg)

*With gong, a web application is a set of go/angular stacks.*

**Gong is a work in progress, uses it carefuly**

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


# 3. A "hello world" stack in 5 minutes

## 3.1. Prerequisite

- go 1.24 (see https://go.dev/doc/install)
- npm version >= 10 (see https://nodejs.org)
- angular version 19 (see https://angular.dev/)
- vscode (recommanded)

it is possible to generate a functionning stack in 5 minutes. 

## 3.2. Generating the code with the `gongc` command

> go install github.com/fullstack-lang/gong/go/cmd/gongc/gongc@main

In a terminal, below commands :

- `mkdir` creates a `helloworld` directory
- `mkdir` generates a sub directory `go/models`
- `echo` commands generates 2 go structs in this subdirectory
  - `Hello` which stores a way to say hello
  - `Country` which stores a country and an association to the way to say hello in this country 
- `gongc go/models` compiles the models
- `./helloworld` run the server

```bash
mkdir helloworld
cd helloworld
mkdir go
mkdir go/models
echo "package models
type Hello struct {
  Name string
  HelloType HelloType
}" > go/models/hello.go
echo "package models
type Country struct {
  Name string
  Hello *Hello
  AlternateHellos []*Hello
}" > go/models/country.go
echo "package models
type HelloType string
const (
  Casual HelloType = \"Casual\"
  Formal HelloType = \"Formal\"
)" > go/models/hello_type.go
gongc -skipStager=false go/models
cd go/cmd/helloworld
./helloworld -unmarshallFromCode=data/stage.go -marshallOnCommit=data/stage
cd ../../../..
rm -rf helloworld
```

Then, browse to [localhost:8080](http://localhost:8080) and add data manualy.

The option `-marshallOnCommit=stage` let a `stage.go` generated/updated as each save operation.

## 3.3. Injecting Data via REST

The backend of a gong application is a REST server. You can interact with the server via REST calls.

```
curl --request POST \
  --url http://localhost:8080/api/helloworld/go/v1/hellos \
  --header 'content-type: application/json' \
  --data '{"Name": "Bonjour"}'

curl --request POST \
  --url http://localhost:8080/api/helloworld/go/v1/hellos \
  --header 'content-type: application/json' \
  --data '{"Name": "Bonjorno"}'

curl --request POST \
  --url http://localhost:8080/api/helloworld/go/v1/countrys \
  --header 'content-type: application/json' \
  --data '{"Name": "France","HelloID":{"Int64":1,"Valid":true}}'

curl --request POST \
  --url http://localhost:8080/api/helloworld/go/v1/countrys \
  --header 'content-type: application/json' \
  --data '{"Name": "Italy","HelloID":{"Int64":2,"Valid":true}}'
  ```

