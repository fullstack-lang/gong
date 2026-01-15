
- [1. Why Gong ?](#1-why-gong-)
- [2. A "hello world" Domain Specific Modeling Environment (DSME) application](#2-a-hello-world-domain-specific-modeling-environment-dsme-application)
  - [2.1. Prerequisite](#21-prerequisite)
  - [2.2. Generating \& running the code with the `gong` command](#22-generating--running-the-code-with-the-gong-command)
- [3. Installing and compiling the gong repo](#3-installing-and-compiling-the-gong-repo)
- [4. Status](#4-status)


# 1. Why Gong ?

Gong is a framework for lowering the effort for developping Domain Specific Modelling Environnement (DSME).

A DSME is an application that allows a users to edit data and diagrams.
The data is based on the abstract syntax of the Domain Specific Language (DSL). The diagrams are based on the concrete syntax of the DSL.

General Purpose Modeling Languages (GPML) like UML or SysML are standards with hundreds of elements in their abstract and concrete syntax. However, tailoring them to a specific domain often involves subsetting—ignoring the majority of the standard to focus on a narrow slice — and profiling (stereotypes) to bend generic concepts to specific needs.

Gong allows you to grow your DSML. You start with an empty metamodel or an existing small metamodel close to your need and progressively introduce abstract and concrete syntax elements only as your understanding of the domain matures. 

Developing a robust DSML remains a complex task that requires familiarity with standard metamodeling patterns. Gong aims to democratize the development part of this process.

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

Then, browse to [localhost:8080](http://localhost:8080) and add data manualy.

![helloworld.png](docs/images/helloworld.png)

# 3. Installing and compiling the gong repo

```bash
git clone https://github.com/fullstack-lang/gong
cd gong/go/cmd/gong
go install
cd ../../..
./scripts/run_gong_conditionally.sh
```

# 4. Status

Gong has been used in different settings. It is a work in progress that matures with each new DSME development.

For DSME examples see

- explore the "dsme" directory that contains small DSMEs to start from.
- https://github.com/fullstack-lang/gongreqif, a tool for analysing ReqIF files. The abstract syntax is the ReqIF model and the concrete syntax is the rendering of the specifications as markdown/HTML.