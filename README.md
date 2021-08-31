- [Gong](#gong)
  - [About Gong](#about-gong)
  - [Prerequisite](#prerequisite)
    - [Go](#go)
    - [gcc](#gcc)
    - [go-swagger (optional)](#go-swagger-optional)
    - [Angular](#angular)
    - [Vscode (optional)](#vscode-optional)
  - [API description](#api-description)
- [Using gong](#using-gong)
  - [installing The gong compiler](#installing-the-gong-compiler)
  - [reusable stacks](#reusable-stacks)
  - [examples](#examples)

WORK IN PROGRESS

# Gong

![Example of a generated application with gong](docs/images/bookstore-client.png)
*Example of a generated application with gong*

## About Gong

Gong (go + ng) is a framework for rapid web application development (a.k.a. full stack development) based on go for the back-end and angular for the front-end.

The unit of development in gong is the **gong stack**. A gong stack can import other gong stacks. With gong, The front end and the back end of a gong stack are integrated as a whole.

The code in a gong stack is organized in a predefined directory structures. At the top are 2 directories:

- `go` for the go code
- `ng` for the angular code. 

By default, the main program `main.go` of the back-end of a gong stack provides the web server, the business logic and the database in one single binary. `main.go` is located in the root directory because it `embeds` the `ng` directory.

The data model and business logic of the web application is in the `go/models` directory. 

This repository (github.com/fullstack-lang/gong) is the home `gongc` (in go/gongc), a compiler that compiles the business logic written in `go` and generates code in `go` and `ng` directories.

This repository is also a gong stack that can be reused in other gong stacks. For instance, it is used in the gongdoc stack.

## Prerequisite

### Go

go version equal or above 1.16 is mandatory (cf. use of `embed` package). See https://golang.org for installation.

### gcc

A gong stack uses gorm for database access and sqlite as the default database. The sqlite driver requires cgo, which requires gcc.

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

See [gong back-end API](./gong-go-api.md) for API details.
See [gong back-end implementation](./gong-go-impl.md) for implementation details.

# Using gong

## installing The gong compiler

> cd go/gongc; go install

## reusable stacks 

https://github.com/fullstack-lang/gongdoc, a UML editor (based on jointjs) for documenting a gong model. gongdoc uses the gong stack.

https://github.com/fullstack-lang/gongsim, a stack for developping simulations

https://github.com/fullstack-lang/gongleaflet, a stack for developping application with leaflet carto components

https://github.com/fullstack-lang/gongsvg, a stack for developping application with svg graphical components

## examples

https://github.com/fullstack-lang/helloworld

https://github.com/fullstack-lang/bookstore

https://github.com/fullstack-lang/laundromat, An example that uses 3 stacks (gong, gongsim, gongdoc)

https://github.com/fullstack-lang/gongfly, An example that uses 4 stacks (gong, gongsim, gongdoc, gongleaflet)
