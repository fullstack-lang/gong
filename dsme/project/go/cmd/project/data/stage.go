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

	__Product__00000000_UX := (&models.Product{}).Stage(stage)
	__Product__00000001_Backend := (&models.Product{}).Stage(stage)
	__Product__00000002_CI_1_1 := (&models.Product{}).Stage(stage)
	__Product__00000003_New_Product := (&models.Product{}).Stage(stage)
	__Product__00000004_tree_products_PBS_ := (&models.Product{}).Stage(stage)
	__Product__00000005_New_Product := (&models.Product{}).Stage(stage)
	__Product__00000006_New_Product := (&models.Product{}).Stage(stage)

	__Project__00000000_Project_1 := (&models.Project{}).Stage(stage)

	__Root__00000000_Root := (&models.Root{}).Stage(stage)

	// Setup of values

	__Product__00000000_UX.Name = `UX`
	__Product__00000000_UX.IsExpanded = false

	__Product__00000001_Backend.Name = `Backend`
	__Product__00000001_Backend.IsExpanded = false

	__Product__00000002_CI_1_1.Name = `CI 1.1`
	__Product__00000002_CI_1_1.IsExpanded = false

	__Product__00000003_New_Product.Name = `New Product`
	__Product__00000003_New_Product.IsExpanded = false

	__Product__00000004_tree_products_PBS_.Name = `tree - products (PBS)`
	__Product__00000004_tree_products_PBS_.IsExpanded = false

	__Product__00000005_New_Product.Name = `New Product`
	__Product__00000005_New_Product.IsExpanded = false

	__Product__00000006_New_Product.Name = `New Product`
	__Product__00000006_New_Product.IsExpanded = false

	__Project__00000000_Project_1.Name = `Project 1`
	__Project__00000000_Project_1.IsExpanded = true

	__Root__00000000_Root.Name = `Root`

	// Setup of pointers
	// setup of Product instances pointers
	__Product__00000000_UX.SubProducts = append(__Product__00000000_UX.SubProducts, __Product__00000002_CI_1_1)
	__Product__00000000_UX.SubProducts = append(__Product__00000000_UX.SubProducts, __Product__00000004_tree_products_PBS_)
	__Product__00000001_Backend.SubProducts = append(__Product__00000001_Backend.SubProducts, __Product__00000006_New_Product)
	// setup of Project instances pointers
	__Project__00000000_Project_1.RootProducts = append(__Project__00000000_Project_1.RootProducts, __Product__00000000_UX)
	__Project__00000000_Project_1.RootProducts = append(__Project__00000000_Project_1.RootProducts, __Product__00000001_Backend)
	__Project__00000000_Project_1.RootProducts = append(__Project__00000000_Project_1.RootProducts, __Product__00000003_New_Product)
	// setup of Root instances pointers
	__Root__00000000_Root.Projects = append(__Root__00000000_Root.Projects, __Project__00000000_Project_1)
}

