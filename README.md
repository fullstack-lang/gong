# 1. Try Domain Specific Modelling (DSM) in your browser

## System Engineering

- [Capture](https://fullstack-lang.github.io/gong/capture-app-portable.html) ([zip](https://fullstack-lang.github.io/gong/capture-app-portable.zip)): capture stakeholder needs, operational concerns, high-level requirements & analysis concepts
- [Structure](https://fullstack-lang.github.io/gong/structure-app-portable.html) ([zip](https://fullstack-lang.github.io/gong/structure-app-portable.zip)): model system architectures, hierarchical component breakdowns, ports, data & control flows
- [Scenario](https://fullstack-lang.github.io/gong/scenario-app-portable.html) ([zip](https://fullstack-lang.github.io/gong/scenario-app-portable.zip)): explore operational scenarios, actor state evolutions, timeline trajectories & parameter trade-offs
- [Process](https://fullstack-lang.github.io/gong/process-app-portable.html) ([zip](https://fullstack-lang.github.io/gong/process-app-portable.zip)): define operational processes, participant swimlanes, tasks, data flows & control flows
- [Statemachines](https://fullstack-lang.github.io/gong/statemachines-app-portable.html) ([zip](https://fullstack-lang.github.io/gong/statemachines-app-portable.zip)): specify discrete state machines, composite states, transitions, events & actions
- [Project](https://fullstack-lang.github.io/gong/project-app-portable.html) ([zip](https://fullstack-lang.github.io/gong/project-app-portable.zip)): define product breakdown structures (PBS), work breakdown structures (WBS), resource allocation & planning
- [Reqif](https://fullstack-lang.github.io/gong/reqif-app-portable.html) ([zip](https://fullstack-lang.github.io/gong/reqif-app-portable.zip)): view, edit and exchange requirements using the Requirements Interchange Format (ReqIF) standard

## Technical

- [SVG](https://fullstack-lang.github.io/gong/svg-app-portable.html) ([zip](https://fullstack-lang.github.io/gong/svg-app-portable.zip)): interactively render, edit and manipulate Scalable Vector Graphics (SVG) elements

## For fun

- [Barrgraph](https://fullstack-lang.github.io/gong/barrgraph-app-portable.html) ([zip](https://fullstack-lang.github.io/gong/barrgraph-app-portable.zip)): visualize and edit Alfred Barr's genealogical chart of modern art movements and artistic influences
- [Phylla](https://fullstack-lang.github.io/gong/phylla-app-portable.html) ([zip](https://fullstack-lang.github.io/gong/phylla-app-portable.zip)): explore botanical phyllotaxis spiral patterns, 2D/3D plant geometries and generative designs

# 2. Why Gong ?

Gong is a framework for lowering the effort for developping Domain Specific Modelling (DSM).

A DSM is an application that allows a users to edit data and diagrams.
The data is based on the abstract syntax of the Domain Specific Language (DSL). The diagrams are based on the concrete syntax of the DSL.

You can do DSM with General Purpose Modeling Languages (GPML) like UML or SysML that are standards with hundreds of elements in their abstract and concrete syntax. However, tailoring them to a specific domain often involves subsetting—ignoring the majority of the standard to focus on a narrow slice — and profiling (stereotypes) to bend generic concepts to specific needs.

Gong allows you to grow your DSM from a General Purpose Programming Language (GPPL), e.g. go. You start with a reference domain specific model (like Structure or Project). Then, you introduce abstract and concrete syntax elements as your understanding of the domain matures. 

Developing a robust DSM remains a complex task that requires familiarity with standard modeling patterns. Gong aims to democratize the development part of this process.

# 3. A "hello world" gong application

## 3.1. Prerequisite

- go 1.26 (see https://go.dev/doc/install)

## 3.2. Generating & running the code with the `gong` command

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
gong generate go/models
cd go/cmd/helloworld
./helloworld data/stage
cd ../../../..
rm -rf helloworld
```

Then, browse to [localhost:8080](http://localhost:8080) and add data manualy.

![helloworld.png](docs/images/helloworld.png)

# 4. Installing and compiling the gong repo

```bash
git clone https://github.com/fullstack-lang/gong
cd gong/go/cmd/gong
go install
cd ../../..
./scripts/run_gong_conditionally.sh
```

# 5. Status

Gong is a work in progress.