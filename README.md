- [Gong](#gong)
  - [About Gong](#about-gong)
  - [Stack configuration management](#stack-configuration-management)
  - [Gong is intended for system engineering tooling](#gong-is-intended-for-system-engineering-tooling)
  - [Prerequisite](#prerequisite)
    - [Go](#go)
    - [gcc](#gcc)
    - [go-swagger (optional)](#go-swagger-optional)
    - [npm](#npm)
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

Gong (go + ng) is a framework for rapid web application development (a.k.a. full stack development). It is based on two development langages: go for the back-end and angular for the front-end. Gong idea is to leverages best in class components. The back-end leverages [gin](https://github.com/gin-gonic/gin), a web framework and [gorm](https://gorm.io/index.html), an ORM. The front-end leverages [angular material](https://material.angular.io/), an library of angular components.

The unit of development in gong is the **gong stack** (a "stack" in the rest of this document). A stack can import other stacks (both the front end and the back end of a stack are integrated as a whole). The granularity of a stack is similar to an angular components. There are available stacks for [jointjs](https://www.jointjs.com/) and [leaflet](https://leafletjs.com/).

## Stack configuration management

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

## Gong is intended for system engineering tooling

Gong fullstack approach was inspired by [Ruby on Rails](https://rubyonrails.org/) and a more generaly the idea that complexity facing the programmer should be carefuly managed, as it is described in [conceptual compression concept](https://m.signalvnoise.com/conceptual-compression-means-beginners-dont-need-to-know-sql-hallelujah/) and [Rob Pike's design of Go regarding complexity](https://www.dotconferences.com/2015/11/rob-pike-simplicity-is-complicated).

Gong fullstack approach, with a backend in go, is similar in intent to [lorca](https://github.com/zserge/lorca), [wails](https://github.com/wailsapp/wails) and [fyne](https://github.com/fyne-io/fyne). However, the gong framework approach is different because it includes gongc, a go data model compiler to generate front-end and back-end code. In this sense, it is similar to [ent](https://github.com/ent/ent) which includes a ("shema as code") approach.

Also, gong's stated goal is narrower since it is the rapid development of web applications for system engineering (see [paper](https://www.researchgate.net/publication/354237095_GONG_an_open_source_MBSE_toolset/references#fullTextFileContent) for details on this goal)

## Prerequisite

### Go

go version equal or above 1.16 is mandatory (cf. use of `embed` package). See https://golang.org for installation.

### gcc

A stack uses gorm for database access and sqlite as the default database. The sqlite driver requires cgo, which requires gcc.

### go-swagger (optional)

[go-swagger](https://github.com/go-swagger/go-swagger) is a go program is used after each `gongc` compilation to generate the project API in a `yml` file. *gongc* is robust to the absence of go-swagger but it is recommanded to use it if you need to document the API with yaml.

### npm

Gong uses npm version >= 6.14 (see https://nodejs.org)

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

https://github.com/fullstack-lang/gongjointjs, a stack for developping application with jointjs interactive graphical component

## Examples

https://github.com/fullstack-lang/helloworld is a recommanded starting point for understanding gong.

https://github.com/fullstack-lang/bookstore is a little more sophisticated example than helloworld.

https://github.com/fullstack-lang/laundromat, is a more sophisticated example. It is a simulation stack that reuses 3 other stacks (gong, gongsim, gongdoc)

https://github.com/fullstack-lang/gongfly, an airplane simulation that reuses 4 stacks (gong, gongsim, gongdoc, gongleaflet)

https://github.com/fullstack-lang/gongproject, a project management application that reuses 3 stacks (gong, gongjointjs, gongdoc)


## Generating a "hello world" stack in 5 minutes

If prerequisite and gongc are installed, 
it is possible to generate a functionning stack in 5 minutes. 

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
}" > go/models/country.go
gongc go/models
./go/cmd/helloworld/helloworld
```

Then, browse to [localhost:8080](http://localhost:8080)

All commands are executed fast (on an average computer) except `gongc go/models` which takes a few minutes. `gongc` can be long the first time it is executed for a stack because it perfoms `npm i` (installation of node packages) which requires the download of hundreds of megabytes.

If `gongc` is performed again, it will take a few seconds.

