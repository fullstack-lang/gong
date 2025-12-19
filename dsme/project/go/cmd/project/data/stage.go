package main

import (
	"time"

	"github.com/fullstack-lang/gong/dsme/project/go/models"
	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// _ point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.Stage) {

	// Declaration of instances to stage

	__Product__00000000_CI_1 := (&models.Product{}).Stage(stage)
	__Product__00000001_CI_2 := (&models.Product{}).Stage(stage)
	__Product__00000002_CI_1_1 := (&models.Product{}).Stage(stage)
	__Product__00000003_New_Product := (&models.Product{}).Stage(stage)
	__Product__00000004_New_Product := (&models.Product{}).Stage(stage)
	__Product__00000005_New_Product := (&models.Product{}).Stage(stage)

	__Project__00000000_Project_1 := (&models.Project{}).Stage(stage)

	__Root__00000000_Root := (&models.Root{}).Stage(stage)

	// Setup of values

	__Product__00000000_CI_1.Name = `CI 1`

	__Product__00000001_CI_2.Name = `CI 2`

	__Product__00000002_CI_1_1.Name = `CI 1.1`

	__Product__00000003_New_Product.Name = `New Product`

	__Product__00000004_New_Product.Name = `New Product`

	__Product__00000005_New_Product.Name = `New Product`

	__Project__00000000_Project_1.Name = `Project 1`

	__Root__00000000_Root.Name = `Root`

	// Setup of pointers
	// setup of Product instances pointers
	__Product__00000000_CI_1.SubProducts = append(__Product__00000000_CI_1.SubProducts, __Product__00000002_CI_1_1)
	__Product__00000000_CI_1.SubProducts = append(__Product__00000000_CI_1.SubProducts, __Product__00000004_New_Product)
	// setup of Project instances pointers
	__Project__00000000_Project_1.RootProducts = append(__Project__00000000_Project_1.RootProducts, __Product__00000000_CI_1)
	__Project__00000000_Project_1.RootProducts = append(__Project__00000000_Project_1.RootProducts, __Product__00000001_CI_2)
	__Project__00000000_Project_1.RootProducts = append(__Project__00000000_Project_1.RootProducts, __Product__00000003_New_Product)
	// setup of Root instances pointers
	__Root__00000000_Root.Projects = append(__Root__00000000_Root.Projects, __Project__00000000_Project_1)
}

