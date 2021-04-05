// generated from OrmFileSetupTemplate
package orm

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // justificiation for blank import : initialisaion of the sqlite driver
)

// genQuery return the name of the column
func genQuery(columnName string) string {
	return fmt.Sprintf("%s = ?", columnName)
}

// SetupModels connects to the sqlite database
func SetupModels(logMode bool, filepath string) *gorm.DB {
	db, err := gorm.Open("sqlite3", filepath)

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.LogMode(logMode)

	AutoMigrate(db)

	return db
}

// AutoMigrate migrates db with with orm Struct
func AutoMigrate(db *gorm.DB) {
	_db := db.AutoMigrate( // insertion point for reference to structs
		&GongBasicFieldDB{},
		&GongEnumDB{},
		&GongEnumValueDB{},
		&GongStructDB{},
		&ModelPkgDB{},
		&PointerToGongStructFieldDB{},
		&SliceOfPointerToGongStructFieldDB{},
	)

	if _db.Error != nil {
		msg := _db.Error.Error()
		panic("problem with migration " + msg + " on package github.com/fullstack-lang/gong/stacks/gong/go")
	}
	log.Printf("Migration OK")
}

func ResetDB(db *gorm.DB) { // insertion point for reference to structs
	db.Delete(&GongBasicFieldDB{})
	db.Delete(&GongEnumDB{})
	db.Delete(&GongEnumValueDB{})
	db.Delete(&GongStructDB{})
	db.Delete(&ModelPkgDB{})
	db.Delete(&PointerToGongStructFieldDB{})
	db.Delete(&SliceOfPointerToGongStructFieldDB{})
}
