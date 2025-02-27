- [1. Gong](#1-gong)
  - [1.1. About Gong](#11-about-gong)
  - [1.2. Gong aims low complexity](#12-gong-aims-low-complexity)
  - [1.3. Prerequisite](#13-prerequisite)
    - [1.3.1. Go](#131-go)
    - [1.3.2. go-swagger (optional)](#132-go-swagger-optional)
    - [1.3.3. npm](#133-npm)
    - [1.3.4. Angular](#134-angular)
    - [1.3.5. Vscode (optional)](#135-vscode-optional)
- [2. Gong is a Go sub language](#2-gong-is-a-go-sub-language)
  - [2.1. Gong specification](#21-gong-specification)
    - [2.1.1. package nomenclature](#211-package-nomenclature)
    - [2.1.2. Gongstruct](#212-gongstruct)
    - [2.1.3. Gongenum](#213-gongenum)
    - [2.1.4. Gongfields](#214-gongfields)
    - [2.1.5. Gongnote](#215-gongnote)
  - [2.2. Gong library](#22-gong-library)
- [3. Using gong](#3-using-gong)
  - [3.1. Running the gong test application](#31-running-the-gong-test-application)
  - [3.2. Testing the generation of the code](#32-testing-the-generation-of-the-code)
  - [3.3. Reusable stacks](#33-reusable-stacks)
  - [3.4. Examples](#34-examples)
- [4. Gong's Features](#4-gongs-features)
  - [4.1. Gong is a go sub langage for generating a full stack](#41-gong-is-a-go-sub-langage-for-generating-a-full-stack)
  - [4.2. Back end SQL and go code from gong code](#42-back-end-sql-and-go-code-from-gong-code)
  - [4.3. Controler go code from gong code](#43-controler-go-code-from-gong-code)
  - [4.4. Front end angular code from gong code](#44-front-end-angular-code-from-gong-code)
  - [4.5. Angular material code from gong code](#45-angular-material-code-from-gong-code)
  - [4.6. Front stage, back stage and repository programming model](#46-front-stage-back-stage-and-repository-programming-model)
  - [4.7. UML Code documentation as go code](#47-uml-code-documentation-as-go-code)
  - [4.8. Integrated Stack configuration management](#48-integrated-stack-configuration-management)
  - [4.9. Persistance as go code for enabling fast refactoring](#49-persistance-as-go-code-for-enabling-fast-refactoring)
  - [4.10. Further documentation](#410-further-documentation)
- [5. A "hello world" stack in 5 minutes](#5-a-hello-world-stack-in-5-minutes)
  - [5.1. Generating the code with the `gongc` command](#51-generating-the-code-with-the-gongc-command)
  - [5.2. Timing of the `gongc` command](#52-timing-of-the-gongc-command)
  - [5.3. Injecting Data via REST](#53-injecting-data-via-rest)

# 1. Gong

![gong logo](docs/images/gong%20logo.svg)

With gong, a web application is a set of stacks. Each stack is based on go and angular. Each stack has its own data model.

*Gong is a work in progress. API of the framework is not yet stabilized/baselined*

## 1.1. About Gong

Gong (go + ng) is a framework for rapid web application development (a.k.a. full stack development) based on go and angular. The go back-end uses [gin](https://github.com/gin-gonic/gin), [gorm](https://gorm.io/index.html) and sqlite (a pure go sqlite, no cgo needed). The angular front-end uses [angular material](https://material.angular.io/).

The unit of development in gong is the **gong stack** (a "stack" in the rest of this document). A stack can import other stacks (both the front end and the back end of a stack are integrated as a whole). The granularity of a stack is similar to an angular components.

Gong includes gongc, a go data model compiler to generate front-end and back-end code. Gongc compiles gong code, a go sub-langage, to go and angular.

## 1.2. Gong aims low complexity

Gong fullstack approach was inspired by the idea that complexity facing the programmer should be carefuly managed (for the idea, see [conceptual compression concept](https://m.signalvnoise.com/conceptual-compression-means-beginners-dont-need-to-know-sql-hallelujah/) or [Rob Pike's design of Go regarding complexity](https://go.dev/talks/2015/simplicity-is-complicated/).

## 1.3. Prerequisite

### 1.3.1. Go

go version equal or above 1.21. See https://golang.org for installation.

Gong uses sqlite3 in a cgo free configuration by default.

### 1.3.2. go-swagger (optional)

[go-swagger](https://github.com/go-swagger/go-swagger) is a go program that is used (as an option) after each `gongc` compilation to generate the project API in a `yml` file. *gongc* is robust to the absence of go-swagger but it is recommanded to use it if you need to document the back-end API with yaml.

### 1.3.3. npm

Gong uses npm version >= 8 (see https://nodejs.org)

### 1.3.4. Angular

Gong uses angular version 16 (see https://angular.io for installation)

### 1.3.5. Vscode (optional)

Vscode is usefull & handy because the tasks definitions and debug configuration related to gong are provided in the repository.

# 2. Gong is a Go sub language

Gong is a langage with a compiler and a library.

Gong is also a go sub language because it is not autonomous. It must be developped within a go envionment and
it needs the tools of the go toolchains. A gong program is a go program.

## 2.1. Gong specification

Gong specification is the go specification with additional constraints. If a go construct does not meet those contraints,
it is transparant to the gong compiler.

There are three objects in gong.

### 2.1.1. package nomenclature

A gong program is a go program (go > 1.20) developped with a go module. The gong code is developed in only
one package "go/models"

```go
package models
```

No particular inmport is necessary.

### 2.1.2. Gongstruct

```go
// Hello is a gongstruct
type Hello struct { // it is exported
  Name string // it has an exported "Name" field
}

// foo is not a gongstruct
type foo struct { // it is not exported
  Name string
}

// Bar is not a gongstruct
type Bar struct { // it is exported
  name string // it has no exported "Name" field
}

// Zong is not a gongstruct
//
// gong:ignore it has a "ignore" magic code
type Zong struct { // it is exported
  Name string // it has an exported "Name" field
}
```

### 2.1.3. Gongenum


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

### 2.1.4. Gongfields

Gongfields are fields within a gongstruct.

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

### 2.1.5. Gongnote

```go
// A gongnote is a string const with a comment
// that starts with the GONGDOC keyword
// It follows the "Note" defined in the [go doc](https://pkg.go.dev/go/doc) and 
// [go doc comment]([]https://pkg.go.dev/go/doc/comment) defined in the
// standard go library
//
// > A Note represents a marked comment starting with "MARKER(uid): note body".

// GONGDOC(NoteExample): Note example
// It can reference Gongstructs
// [models.Bstruct],
// [models.Astruct],
// Gongenums
// [models.AEnumType]
// Gongfields
// [models.Astruct.Associationtob], 
// having the following const exported identifier allows for
// referencing the note from the UML note and allows for
// renaming
//
// # This is heading 1
//
// ## This is heading 1.1
//
//	-
const NoteExample = ""
```


## 2.2. Gong library

The gongc compilers generates code within the "go/models" package and within other packages.

# 3. Using gong

## 3.1. Running the gong test application

the `test` directory contains a stack wit the generated code.

```
cd test/ng
npm i; ng build
cd ..
go run main.go
```

Then, browse to [localhost:8080](http://localhost:8080)

![test web application](docs/images/test.png)
*Example of a generated application with gong*

## 3.2. Testing the generation of the code

Installing The gong compiler.

From the root directory.

> cd go/gongc; go install; cd ../..

Generating the code

> cd test; gongc go/models

## 3.3. Reusable stacks

A gong application is a stack that can integrate other stacks. Below is a list of stacks that can be reused. 

https://github.com/fullstack-lang/gongdoc, a UML editor (based on jointjs) for documenting a gong model. gongdoc uses the gong stack.

https://github.com/fullstack-lang/gongsim, a stack for developping simulations

https://github.com/fullstack-lang/gongleaflet, a stack for developping application with leaflet carto components

https://github.com/fullstack-lang/gongsvg, a stack for developping application with svg graphical components

https://github.com/fullstack-lang/gongjointjs, a stack for developping application with jointjs interactive graphical component

## 3.4. Examples

https://github.com/fullstack-lang/helloworld is a recommanded starting point for understanding gong.

https://github.com/fullstack-lang/bookstore is a little more sophisticated example than helloworld.

https://github.com/fullstack-lang/laundromat, is a more sophisticated example. It is a simulation stack that reuses 3 other stacks (gong, gongsim, gongdoc)

https://github.com/fullstack-lang/gongfly, an airplane simulation that reuses 4 stacks (gong, gongsim, gongdoc, gongleaflet)

https://github.com/fullstack-lang/gongproject, a project management application that reuses 3 stacks (gong, gongjointjs, gongdoc)

# 4. Gong's Features


## 4.1. Gong is a go sub langage for generating a full stack

Gong is a sublangage of go (stereotyped go). It is comprised of one or many *gongstruct* and *gongenum*. The *gongc* compiler is based on the *go* compiler.

A gong project is a go module with a unique name (for instance `github.com/fullstack-lang/gongdoc`). *gongstruct* and *gongenum* have to be present in the `go/models` sub package.

A *gongstruct* is a go struct with a `Name` field.

```go
package models
type Hello struct {
  Name string
}
```

A *gongenum* is a go string with some const values. The string type must be commented with `// swagger:enum`

```go
package models

// swagger:enum AEnumType
type AEnumType string

// values for EnumType
const (
	ENUM_VAL1 AEnumType = "ENUM_VAL1"
	ENUM_VAL2 AEnumType = "ENUM_VAL2"
)
```

A *gongstruct* can have *gongfield*. A *gongfield* is an exported fields (starts with an Upperscase) with
type as `int`, `float64`, `bool`, `time.Time`, `time.Duration`, `pointer` to another `gongstruct` in 
the same package and `slice` of pointers
to another `gongstruct` in the same package (for ONE-MANY associations). 
MANY-MANY associations are also implemented (if the target of the ONE-MANY association ends with the `Use`).

```go
// Astruct demonstrate basic gong features
type Astruct struct {

	// a "Name" field is necessary to generate a GongStruct
	Name string

	// string a is a supported type of gong
	Description string

	// bool is a supported type of gong
	Booleanfield bool

	// enums is a supported type of gong (if they are string)
	Aenum AEnumType

	// an embedded struct is supported if it is not a gongstruct (i.e. it is without a field `Name`)
	Cstruct

	// float64 is a supported type of gong
	Floatfield float64

	// int is a supported type of gong
	Intfield int

	// time.Time is a supported type of gong
	Date time.Time

	// time.Duration is a supported type of gong
	Duration time.Duration

	// ONE-ZERO/ONE association is a supported type of gong
	Associationtob *Bstruct

	// ONE-MANY association is a supported type of gong
	Sliceofb []*Bstruct

	// MANY-MANY association, 
	// because AclassBclassUse ends with "Use", it will implement a ONE/ONE assocation to the Bstruct)
	AnarrayofbUse  []*AstructBstructUse
}
```


## 4.2. Back end SQL and go code from gong code

The [gorm](https://gorm.io/index.html) framework is a go API for ORM (Object Realtionship Management). This means you do not need to code SQL to configure/migrate the database. [gorm](https://gorm.io/index.html) acts as a [conceptual compression](https://m.signalvnoise.com/conceptual-compression-means-beginners-dont-need-to-know-sql-hallelujah/) 
since you do not need to know SQL code to program the database, you only need to master the go API.

However, *gorm* is still a concept you need to know to program the database in go. *gong* allows to ignore this conceptual layer
 since the *gorm* code is generated. Indeed, the *gongc* compiler compiles the models (the `go/models` package) 
 to generate the appropriate gorm code (in the `go/orm` package). Therefore, *gong* is another 
 [conceptual compression](https://m.signalvnoise.com/conceptual-compression-means-beginners-dont-need-to-know-sql-hallelujah/) 
 to avoid programming *gorm* code.

## 4.3. Controler go code from gong code

The [gin](https://github.com/gin-gonic/gin) framework is an API for programming the controllers that 
implement the REST web service of the application. [gin](https://github.com/gin-gonic/gin) acts as a 
[conceptual compression](https://m.signalvnoise.com/conceptual-compression-means-beginners-dont-need-to-know-sql-hallelujah/) 
since it provides a high level concept to program the REST api. 

However, you need to know to program [gin](https://github.com/gin-gonic/gin) for implementing the REST api. *gong* allows to ignore this conceptual layer since the *gin* code is generated by *gong*. The *gongc* compiles compiles the models (the `go/models` package) 
to generate the appropriate gin code (in the `go/controllers` package)

## 4.4. Front end angular code from gong code

Angular framework is an API for programming front-end of the application.

gongc compiles the models (the `go/models` package) to generate the appropriate angular code (in the `ng/projects/<name of the package>` angular workspace). This code provides an API to the front end code that
follows the data models of the the `go/models` package.

Each gongstruct/gongenum is compiled into an appropriate typescript code.

## 4.5. Angular material code from gong code

Angular Material is a set of front-end Angular components.

gongc compiles the models (the `go/models` package) to generate the appropriate angular material code, in the `ng/projects/<name of the package>` angular workspace :

- navigation tree between all gongstruct of the `go/models` package
- table for each gongstruct of the `go/models` package
- form for each gongstruct of the `go/models` package

## 4.6. Front stage, back stage and repository programming model 

Putting data to the database and retrieving data from the database (CRUD operations) is performed via an
API that is generated for each *gongstruct*.

The API follows loosely the [repository pattern](https://docs.microsoft.com/en-us/dotnet/architecture/microservices/microservice-ddd-cqrs-patterns/infrastructure-persistence-layer-design).

The generated API comprises `Stage()`, `Unstage()`, `Commit()`, `Delete()` for each *gongstruct*.


## 4.7. UML Code documentation as go code 

see [gongcode](https://github.com/fullstack-lang/gongdoc).

## 4.8. Integrated Stack configuration management

The configuration of *both* back-end and front-end code of a stack is a single configuration item.

This is done thank the go `module`. For the go code, it is the standard way of managing dependencies. For the angular/typescript/js code, it is done in four steps.

First, gong uses the go `embed` fearure that allows, by using four lines of go in a file stored in the angular workspace to directory,

```go
package ng
import "embed"
//go:embed projects
var Projects embed.FS
```

that the code in the angular workspace `projects` directory is stored into the `go module`.

The second step is to import the created `ng package` in the project that will use the stack. For instance, the following line makes avaible the `Project` directory to the project.

```go
	_ "github.com/fullstack-lang/gongjointjs/ng"
  ```

The third step is another go feature, the  `go mod vendor` command, that makes available the source code of all dependencies in a `vendor` directory simply by issuing the command. Then, the angular code is now in the directory `vendor/github.com/fullstack-lang/gongjointjs/ng`.

The four step is to define your front-end dependency by using the `tsconfig.json` file and point it the to import path into the `vendor` directory (instead of using the installation by `npm install` of the imported front code module). you are therefore assured that your back-end code and front-end code belong to the same configuration. (see the https://github.com/fullstack-lang/gongproject/blob/master/ng/tsconfig.json for an example of tsconfig.json configuration).

## 4.9. Persistance as go code for enabling fast refactoring

Gong's goal is to speed up development of full stack applications. Gong's goal is therefore to allow fast iterations of the database model and **content/database**.

An iteration of the data model can include an addition or removal of a concept (a go´struct') or the addition or removal of a field of a concept. In this case, the 'gorm' tool takes care of the content/database migration.

An iteration can also include a renaming of a field or the renaming of a struct. In this case, the code can be automatically changed by the use of the refactoring function of the ["go please" langage server](https://github.com/golang/tools/tree/master/gopls) for the backend and the [typescript language server](https://github.com/typescript-language-server/typescript-language-server) for the front end.

Wihout gong, if one needs to refactor the name of a gongstruct or the name of a field of a gongstruct, the content/database of the application must be refactored by hand.

For instance :
 - via a json file
 - via the sqlite table/column renaming
 - via [gorm](https://gorm.io/docs/migration.html)

With gong, data refactoring is automatic. Gong API provides a `Marshall()` function of the staged objects that generates an `Unmarshall()` function in go code (a persistance of the repository data as go code)

when refactoring the code, the generated go code is refactored. Therefore, no need to manualy refactor the data.

## 4.10. Further documentation

See [gong back-end API](./docs/gong-go-api.md) for API details.


See [gong back-end implementation](./docs/gong-go-impl.md) for implementation details.
# 5. A "hello world" stack in 5 minutes

If prerequisite and gongc are installed, 
it is possible to generate a functionning stack in 5 minutes. 

## 5.1. Generating the code with the `gongc` command

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
}" > go/models/hello.go
echo "package models
type Country struct {
Name string
Hello *Hello
AlternateHellos []*Hello
}" > go/models/country.go
gongc go/models
cd go/cmd/helloworld
./helloworld -unmarshallFromCode=stage.go -marshallOnCommit=stage 
```

Then, browse to [localhost:8080](http://localhost:8080) and add data manualy.

With the option `-marshallOnCommit=stage`, a `stage.go` file is generated as each save operation (along the default sqlite database `test.db`). When the application is restarted with the `--unmarshall=stage` , the data is injected from the `stage.go` file, not from the database.

## 5.2. Timing of the `gongc` command

`gongc go/models` takes a few minutes the first time it is executed. `gongc` can be long the first time it is executed for a stack because it perfoms `npm i` which can be long if it is the first time (3'37'' on a macbook pro with a 2,6 GHz 6-Core Intel Core i7). 

If `gongc` is performed again, it will take a few tens seconds (32'' on a macbook pro with a 2,6 GHz 6-Core Intel Core i7, 1'16'' on a Core I7 windows PC).

## 5.3. Injecting Data via REST

The backend of a gong application is a REST server (thanks to gin). You can interact with the server via REST calls.

For instance, if you start from an empty database, the following commands will inject proprer data.

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

Port ot ng 18
```bash
mv ../package.json ../package.json.tmp
find .. -type d -name "node_modules" -prune -exec rm -rf '{}' +
find .. -type d -name "dist" -prune -exec rm -rf '{}' +
find .. -type d -name ".angular" -prune -exec rm -rf '{}' +
find .. -name "package-lock.json" -prune -exec rm -rf '{}' +
npm i
ng update @angular/core@18 @angular/cli@18 --allow-dirty
ng update @angular/material@18 --allow-dirty
ng update angular-split@18 --allow-dirty
mv ../package.json.tmp ../package.json
find .. -type d -name "node_modules" -prune -exec rm -rf '{}' +
find .. -type d -name "dist" -prune -exec rm -rf '{}' +
find .. -type d -name ".angular" -prune -exec rm -rf '{}' +
find .. -name "package-lock.json" -prune -exec rm -rf '{}' +
npm i
ng build
```

same, when in a sub directory stack
```bash
mv ../../package.json ../../package.json.tmp
find ../.. -type d -name "node_modules" -prune -exec rm -rf '{}' +
find ../.. -type d -name "dist" -prune -exec rm -rf '{}' +
find ../.. -type d -name ".angular" -prune -exec rm -rf '{}' +
find ../.. -name "package-lock.json" -prune -exec rm -rf '{}' +
npm i
ng update @angular/core@18 @angular/cli@18 --allow-dirty
ng update @angular/material@18 --allow-dirty
ng update angular-split@18 --allow-dirty
mv ../../package.json.tmp ../../package.json
find ../.. -type d -name "node_modules" -prune -exec rm -rf '{}' +
find ../.. -type d -name "dist" -prune -exec rm -rf '{}' +
find ../.. -type d -name ".angular" -prune -exec rm -rf '{}' +
find ../.. -name "package-lock.json" -prune -exec rm -rf '{}' +
npm i
ng build
```