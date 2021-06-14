// generated from OrmFileSetupTemplate
package orm

import (
	"log"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

// genQuery return the name of the column
func genQuery( columnName string) string {
	return fmt.Sprintf("%s = ?", columnName)
}

// SetupModels connects to the sqlite database
func SetupModels(logMode bool, filepath string) *gorm.DB {

	db, err := gorm.Open(sqlite.Open(filepath), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	AutoMigrate(db)

	return db
}

// AutoMigrate migrates db with with orm Struct
func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate( // insertion point for reference to structs 
	  &GongBasicFieldDB{},
	  &GongEnumDB{},
	  &GongEnumValueDB{},
	  &GongStructDB{},
	  &GongTimeFieldDB{},
	  &ModelPkgDB{},
	  &PointerToGongStructFieldDB{},
	  &SliceOfPointerToGongStructFieldDB{},
	)

	if err != nil {
		msg := err.Error()
		panic("problem with migration " + msg + " on package github.com/fullstack-lang/gong/go")
	}
	log.Printf("Database Migration of package github.com/fullstack-lang/gong/go is OK")
}

func ResetDB(db *gorm.DB) { // insertion point for reference to structs 
	  db.Delete(&GongBasicFieldDB{})
	  db.Delete(&GongEnumDB{})
	  db.Delete(&GongEnumValueDB{})
	  db.Delete(&GongStructDB{})
	  db.Delete(&GongTimeFieldDB{})
	  db.Delete(&ModelPkgDB{})
	  db.Delete(&PointerToGongStructFieldDB{})
	  db.Delete(&SliceOfPointerToGongStructFieldDB{})
}