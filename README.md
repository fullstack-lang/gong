- [Gong](#gong)
  - [About Gong](#about-gong)
  - [Prerequisite](#prerequisite)
    - [Go](#go)
    - [gcc](#gcc)
    - [go-swagger (optional, for generating Postman files)](#go-swagger-optional-for-generating-postman-files)
    - [Angular](#angular)
    - [Vscode](#vscode)
  - [Description of a gong application](#description-of-a-gong-application)
- [Using gong](#using-gong)
  - [installing The gong compiler](#installing-the-gong-compiler)
  - [reusable stacks](#reusable-stacks)
  - [examples](#examples)

WORK IN PROGRESS

# Gong

![Example of a generated application with gong](docs/images/bookstore-client.png)
*Example of a generated application with gong*

## About Gong

Gong (go + ng) is a framework for rapid web application development (full stack) based on go for the back-end and angular for the front-end.

The code in a gong stack is organized in a predefined directory structures. At the top are 2 directories `go` for the go code and `ng` for the angular code. The data model of the web application is in the `go/models` directory. 

github.com/fullstack-lang/gong is the repository for the code generation tool.

It is mainly a compilers, `gongc` (in go/gongc), a compiler that compiles the business logic written in `go`, extracts some elements in a `gong` language and generates code in `go` and `ng` directories.

## Prerequisite

### Go

go version equal or above 1.16 is mandatory (cf. use of `embed` package). See https://golang.org for installation.

### gcc

a gong stack uses gorm for the database and sqlite as the default database. Unfortunatly, the sqlite driver requires cgo, which requires gcc.

### go-swagger (optional, for generating Postman files)

swagger is a go program is used after each `gongc` compilation to generate the project API in a `yml` file. gongc is robust to the absence of go-swagger but it is recommanded to use it

It is not mandatory to install it.

On mac/linux, One way is

```bash
dir=$(mktemp -d) 
git clone https://github.com/go-swagger/go-swagger "$dir" 
cd "$dir"
go install ./cmd/swagger
```

on windows with powershell, creates and go into `go-swagger`
```bash
git clone https://github.com/go-swagger/go-swagger
go install ./go-swagger/cmd/swagger
```

setting up the path for `swagger` on mac/linux
```sh
export PATH=$PATH:$(go env GOPATH)/bin:$HOME/go/bin
```

### Angular

The `ng` command is used by different gong programs. Gong uses ng version >= 11 (see https://angular.io for installation)

### Vscode

Vscode is usefull & handy because the tasks definitionq and debug configuration related to gong are provided in the repository.

(note: No makefile is provided).

## Description of a gong application

A go package (for instance `<path>/go/models` ) written following the `gong` langage constraints can be compiled by the `gongc` compiler into a stack of integrated components:
- a set of `go` packages for the backend
- an `angular` library for the front end. 

This stack can be packaged into a reusable `gong` library to be used in another full stack developpemnent (a *bookstore* example is provided in the repository).

- a `go/orm` package, leveraging gorm, the fantastic go ORM, for the persistance into GORM supported database (sqlite3 in memory, sqlite3 file, postgres, ...)
- a `go/controllers` package, leveraging the gin framework, an HTTP web framework written in Go (Golang)
- a `go/controllers/<path>.yaml` open api 2.0 interface definition (thks to go-swagger), it provides a RESTful interface for  developing and consuming an API of the gong package

- a `ng/projects/<path>` angular service library for accessing gong object with some an angular material library with commonly used material components: table, editor, presentation, splitter presentation, arborescence presentation

if a gong variable data is created on the backend, a constraint is to register all instances on a store.

See [gong back-end API](./gong-go-api.md.md) for API details.
See [gong back-end implementation](./gong-go-impl.md.md) for implementation details.

# Using gong

## installing The gong compiler

> cd go/gongc; go install

## reusable stacks 

https://github.com/fullstack-lang/gongdoc, a UML editor for documenting a gong model

https://github.com/fullstack-lang/gongsim, a stack for developping simulations

https://github.com/fullstack-lang/gongleaflet, a stack for developping application with leaflet carto components

https://github.com/fullstack-lang/gongsvg, a stack for developping application with svg graphical components

## examples

https://github.com/fullstack-lang/helloworld

https://github.com/fullstack-lang/bookstore

https://github.com/fullstack-lang/laundromat, An example that uses 3 stacks (gong, gongsim, gongdoc)
