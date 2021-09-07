# Back Repo implementation

- [Back Repo implementation](#back-repo-implementation)
  - [dependency injection in the stage object](#dependency-injection-in-the-stage-object)
  - [The back repo object](#the-back-repo-object)

## dependency injection in the stage object

the `package models` is where the coding is done in the back-end (see the simulation [gongfly](https://github.com/fullstack-lang/gongfly) for instance).

Gong philosophy is that `package models` shall not be dependant on the back repository implementation. The dependency to the back is injected at the init of the stack back repo with the call to `orm.SetupModels()` in the `main.go` by default.

```go
    // setup GORM
    db := orm.SetupModels(false, "./test.db")

    ....

```

The generated code, one finds
```go

// SetupModels connects to the sqlite database
func SetupModels(logMode bool, filepath string) *gorm.DB {
    ...
	AutoMigrate(db)

	return db
}

...

// AutoMigrate migrates db with with orm Struct
func AutoMigrate(db *gorm.DB) {

    ...

	BackRepo.init(db)
}

// Init the BackRepoStruct inner variables and link to the database
func (backRepo *BackRepoStruct) init(db *gorm.DB) {
    ...

    // dependency injection 
    models.Stage.BackRepo = backRepo
}
```

The `models.Stage` singloton has a `models.Stage.BackRepo` that supports a generated interface.

```go
type StageStruct struct { // insertion point for definition of arrays registering instances
    ...
	
    BackRepo BackRepoInterface
}
```

In this generated interface, important function are the `Commit()` and `Checkout()` functions because they implement the major operations of the back repository.
```go
type BackRepoInterface interface {
	Commit(stage *StageStruct)
	Checkout(stage *StageStruct)
    ...
}
```

Indeed, when `Commit()` or `Checkout()` are called on the `models.Stage` singloton, the call to the back repo is performed via the interface.

```go
func (stage *StageStruct) Commit() {
	if stage.BackRepo != nil {
		stage.BackRepo.Commit(stage)
	}
}

func (stage *StageStruct) Checkout() {
	if stage.BackRepo != nil {
		stage.BackRepo.Checkout(stage)
	}
}
```

## The back repo object

Now, let's see in `orm/back_repo.go` how the back repo is implemented.


```go
var BackRepo BackRepoStruct

type BackRepoStruct struct {
	// insertion point for per struct back repo declarations
	BackRepoAclass BackRepoAclassStruct
}
```