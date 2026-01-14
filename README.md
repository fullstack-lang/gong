
- [1. About Gong](#1-about-gong)
- [2. A "hello world" Domain Specific Modeling Environment (DSME) application](#2-a-hello-world-domain-specific-modeling-environment-dsme-application)
  - [2.1. Prerequisite](#21-prerequisite)
  - [2.2. Generating \& running the code with the `gong` command](#22-generating--running-the-code-with-the-gong-command)
- [3. Installing and compiling the gong repo](#3-installing-and-compiling-the-gong-repo)


# 1. About Gong

Gong is a framework for developping Domain Specific Modelling Environnement (DSME).

A DSME is an application that allows a users to edit data and diagrams.
The data is based on the abstract syntax of the Domain Specific Language (DSL). The diagrams are based on the concrete syntax of the DSL.

General Purpose Modelling Language (GPML) like UML or SysML have also abstract/concrete syntax. 


**Gong is a work in progress, use it carefuly**

# 2. A "hello world" Domain Specific Modeling Environment (DSME) application

## 2.1. Prerequisite

- go 1.24 (see https://go.dev/doc/install)

## 2.2. Generating & running the code with the `gong` command

In a terminal, execute the following commands:


```bash
go install github.com/fullstack-lang/gong/go/cmd/gong@main
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
gong generate --level1 go/models
cd go/cmd/helloworld
./helloworld -unmarshallFromCode=data/stage.go -marshallOnCommit=data/stage
cd ../../../..
rm -rf helloworld
```

In a terminal, below commands :

- `mkdir` creates a `helloworld` directory
- `mkdir` generates a sub directory `go/models`
- `echo` commands generates 2 go structs in this subdirectory
  - `Hello` which stores a way to say hello
  - `Country` which stores a country and an association to the way to say hello in this country 
- `gong --level1 go/models` compiles the models
- `./helloworld` run the server

Then, browse to [localhost:8080](http://localhost:8080) and add data manualy.

The option `-marshallOnCommit=stage` let a `stage.go` generated/updated as each save operation.

# 3. Installing and compiling the gong repo

```bash
git clone https://github.com/fullstack-lang/gong
cd gong/go/cmd/gong
go install
cd ../../..
./scripts/run_gong_conditionally.sh
```