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
		&ClassdiagramDB{},
		&ClassshapeDB{},
		&FieldDB{},
		&GongdocCommandDB{},
		&GongdocStatusDB{},
		&LinkDB{},
		&PkgeltDB{},
		&PositionDB{},
		&StateDB{},
		&UmlscDB{},
		&VerticeDB{},
	)

	if _db.Error != nil {
		msg := _db.Error.Error()
		panic("problem with migration " + msg + " on package github.com/fullstack-lang/gong/stacks/gongdoc/go")
	}
	log.Printf("Migration OK")
}

func ResetDB(db *gorm.DB) { // insertion point for reference to structs
	db.Delete(&ClassdiagramDB{})
	db.Delete(&ClassshapeDB{})
	db.Delete(&FieldDB{})
	db.Delete(&GongdocCommandDB{})
	db.Delete(&GongdocStatusDB{})
	db.Delete(&LinkDB{})
	db.Delete(&PkgeltDB{})
	db.Delete(&PositionDB{})
	db.Delete(&StateDB{})
	db.Delete(&UmlscDB{})
	db.Delete(&VerticeDB{})
}