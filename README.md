- [Gong](#gong)
  - [About Gong](#about-gong)
  - [Prerequisite](#prerequisite)
    - [Go](#go)
    - [gcc](#gcc)
    - [go-swagger (optional)](#go-swagger-optional)
    - [Angular](#angular)
    - [Vscode (optional)](#vscode-optional)
  - [Developping a gong stack](#developping-a-gong-stack)
- [Using gong](#using-gong)
  - [Running the gong test application](#running-the-gong-test-application)
  - [Testing the generation of the code](#testing-the-generation-of-the-code)
  - [Reusable stacks](#reusable-stacks)
  - [Examples](#examples)
  - [Generating a "hello world" stack in 5 minutes](#generating-a-hello-world-stack-in-5-minutes)

# Gong

![gong logo](docs/images/gong%20logo.svg)

With gong, a web application is a set of stacks. Each stack, based on go and angular, has its own data model.

*Gong is a work in progress. API of the framework is not yet stabilized/baselined*

## About Gong

Gong (go + ng) is a framework for rapid web application development (a.k.a. full stack development) based on go for the back-end and angular for the front-end. Gong idea is to leverages best in class components. The back-end leverages the [gin](https://github.com/gin-gonic/gin) web framework and [gorm](https://gorm.io/index.html), the fantastic ORM. The front-end leverages and [angular material](https://material.angular.io/) and [angular split](https://github.com/angular-split/angular-split).

The unit of development in gong is the **gong stack** (a "stack" in the rest of this document). A stack can import other stacks (both the front end and the back end of a stack are integrated as a whole). The granularity of a stack is similar to an angular components. There are available stacks for [jointjs](https://www.jointjs.com/) and [leaflet](https://leafletjs.com/).

Gong is similar in intent to [lorca](https://github.com/zserge/lorca), [wails](https://github.com/wailsapp/wails) and [fyne](https://github.com/fyne-io/fyne). However, the gong framework approach is different because it includes gongc, a go data model compiler to generate front-end and back-end code.

Also, gong's stated goal is narrower, it is the rapid development of web applications for system engineering (see [paper](https://www.researchgate.net/publication/354237095_GONG_an_open_source_MBSE_toolset/references#fullTextFileContent) for details on this goal)

## Prerequisite

### Go

go version equal or above 1.16 is mandatory (cf. use of `embed` package). See https://golang.org for installation.

### gcc

A stack uses gorm for database access and sqlite as the default database. The sqlite driver requires cgo, which requires gcc.

### go-swagger (optional)

[go-swagger](https://github.com/go-swagger/go-swagger) is a go program is used after each `gongc` compilation to generate the project API in a `yml` file. *gongc* is robust to the absence of go-swagger but it is recommanded to use it if you need to document the API with yaml.

### Angular

Gong uses angular version >= 11 (see https://angular.io for installation)

### Vscode (optional)

Vscode is usefull & handy because the tasks definitions and debug configuration related to gong are provided in the repository.

## Developping a gong stack

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

Then, browse to [localhost:8080](http://localhost:8080)

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

## Generating a "hello world" stack in 5 minutes

If prerequisite and gongc are installed, 
it is possible to generate a functionning stack in 5 minutes. 

In a terminal, below commands :

- `mkdir` creates a `helloworld` directory
- `go mod init` generates a go module `github.com/fullstack-lang/helloworld`
- `mkdir` generates a sub directory `go/models`
- `echo` commands generates 2 go structs in this subdirectory
  - `Hello` which stores a way to say hello
  - `Country` which stores a country and an association to the way to say hello in this country 
- `gongc go/models` compiles the models
- `go mod tidy` download needed go packages
- `go build` generates the single binary `./helloworld`
- `./helloworld` run the server

```bash
mkdir helloworld
cd helloworld
go mod init github.com/fullstack-lang/helloworld
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
}" > go/models/country.go
gongc go/models
go mod tidy
go build
./helloworld
```

Then, browse to [localhost:8080](http://localhost:8080)

All commands are executed fast (on an average computer) except `gongc go/models` which takes a few minutes. `gongc` can be long the first time it is executed for a stack because it perfoms `npm i` (installation of node packages) which requires the download of hundreds of megabytes.

If `gongc` is performed again, it will take a few seconds.

