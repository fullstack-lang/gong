- [Gong](#gong)
  - [About Gong](#about-gong)
  - [Stack organization](#stack-organization)
  - [Prerequisite](#prerequisite)
    - [Go](#go)
    - [gcc](#gcc)
    - [go-swagger (optional)](#go-swagger-optional)
    - [Angular](#angular)
    - [Vscode (optional)](#vscode-optional)
  - [API description](#api-description)
- [Using gong](#using-gong)
  - [Running the gong test application](#running-the-gong-test-application)
  - [Testing the generation of the code](#testing-the-generation-of-the-code)
  - [Reusable stacks](#reusable-stacks)
  - [Examples](#examples)

# Gong

![gong logo](docs/images/gong%20logo.svg)

Gong is a work in progress.

## About Gong

Gong (go + ng) is a framework for rapid web application development (a.k.a. full stack development) based on go for the back-end and angular for the front-end.

The unit of development in gong is the **gong stack** (a "stack" in the rest of this document). A stack can import other stacks (both the front end and the back end of a stack are integrated as a whole).

Gong's stated goal is the rapid development of web applications for system engineering (see [paper](https://www.researchgate.net/publication/354237095_GONG_an_open_source_MBSE_toolset/references#fullTextFileContent) for details on this goal)

## Stack organization

Code is organized with a fixed directory structure. At the top are 2 directories:

- `go` for the go code, with the data model and business logic of the application in the `go/models` directory. 
- `ng` for the angular code.

By default, the main program `main.go` of the back-end of a stack provides the web server, the business logic and the database in one single binary. `main.go` is located in the root directory because it `embeds` the `ng` directory (thanks to go v1.16 `embeds` feature).

This repository (github.com/fullstack-lang/gong) is the home of `gongc` (in go/gongc), a compiler that compiles the business logic written in `go` and generates code in `go` and `ng` directories.

This repository is also the home of the `gong` stack whose data model is the description of the data model that is parsed by `gongc`. The gong stack an be reused in other stacks. (for instance, in gongdoc, an UML editor).

## Prerequisite

### Go

go version equal or above 1.16 is mandatory (cf. use of `embed` package). See https://golang.org for installation.

### gcc

A stack uses gorm for database access and sqlite as the default database. The sqlite driver requires cgo, which requires gcc.

### go-swagger (optional)

go-swagger is a go program is used after each `gongc` compilation to generate the project API in a `yml` file. *gongc* is robust to the absence of go-swagger but it is recommanded to use it if you need to document the API with yaml.

On mac/linux,

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

### Angular

Gong uses angular version >= 11 (see https://angular.io for installation)

### Vscode (optional)

Vscode is usefull & handy because the tasks definitions and debug configuration related to gong are provided in the repository.

## API description

See [gong back-end API](./docs/gong-go-api.md) for API details.
See [gong back-end implementation](./docs/gong-go-impl.md) for implementation details.

# Using gong

## Running the gong test application

the `test` directory contains a stack wit the generated code.

```
cd test/ng
npm i; ng build
cd ..
go run main.go
```

Then, browse to [localhost:8080](localhost:8080)

![test web application](docs/images/test.png)
*Example of a generated application with gong*

## Testing the generation of the code

Installing The gong compiler.

From the root directory.

> cd go/gongc; go install; cd ../..

Generating the code

> cd test; gongc go/models

## Reusable stacks

A gong application is a stack that can integrate other stacks. Below is a list of stacks that can be reused. 

https://github.com/fullstack-lang/gongdoc, a UML editor (based on jointjs) for documenting a gong model. gongdoc uses the gong stack.

https://github.com/fullstack-lang/gongsim, a stack for developping simulations

https://github.com/fullstack-lang/gongleaflet, a stack for developping application with leaflet carto components

https://github.com/fullstack-lang/gongsvg, a stack for developping application with svg graphical components

## Examples

https://github.com/fullstack-lang/helloworld is a recommanded starting point for understanding gong.

https://github.com/fullstack-lang/bookstore is a little more sophisticated example than helloworld.

https://github.com/fullstack-lang/laundromat, is a more sophisticated example. It is a simulation stack that reuses 3 other stacks (gong, gongsim, gongdoc)

https://github.com/fullstack-lang/gongfly, An airplane simulation that reuses 4 stacks (gong, gongsim, gongdoc, gongleaflet)
