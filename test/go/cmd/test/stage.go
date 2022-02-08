package main

import "github.com/fullstack-lang/gong/test/go/models"

var Stage models.StageStruct



func Unmarshall(stage *models.StageStruct) {

	// map of identifiers
	__Astruct__000000_A1 := new(models.Astruct).Stage()
	__Astruct__000001_A2 := new(models.Astruct).Stage()

	__Bstruct__000002_B2 := new(models.Bstruct).Stage()
	__Bstruct__000003_B3 := new(models.Bstruct).Stage()
	__Bstruct__000004_B1 := new(models.Bstruct).Stage()

	// initializers of values

	//Init Astruct A1
	__Astruct__000000_A1.Name = "A1"
	__Astruct__000000_A1.Aenum = "ENUM_VAL2"
	__Astruct__000000_A1.CFloatfield = 56.444000
	__Astruct__000000_A1.Intfield = 3

	//Init Astruct A2
	__Astruct__000001_A2.Name = "A2"
	__Astruct__000001_A2.Aenum = "ENUM_VAL2"
	__Astruct__000001_A2.CFloatfield = 0.100000
	__Astruct__000001_A2.Intfield = 0

	__Bstruct__000002_B2.Name = "B2"
	__Bstruct__000003_B3.Name = "B3"
	__Bstruct__000004_B1.Name = "B1"

	// initializers of pointers
	__Astruct__000000_A1.Associationtob = __Bstruct__000002_B2
	__Astruct__000000_A1.Anarrayofb = append(__Astruct__000000_A1.Anarrayofb, __Bstruct__000003_B3)
	__Astruct__000000_A1.Anarrayofb = append(__Astruct__000000_A1.Anarrayofb, __Bstruct__000002_B2)
	__Astruct__000000_A1.Anarrayofb = append(__Astruct__000000_A1.Anarrayofb, __Bstruct__000004_B1)

	return
}


