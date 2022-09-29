// generated from OrmFileSetupTemplate
package orm

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// genQuery return the name of the column
func genQuery(columnName string) string {
	return fmt.Sprintf("%s = ?", columnName)
}

// SetupModels connects to the sqlite database
func SetupModels(logMode bool, filepath string) *gorm.DB {
	// adjust naming strategy to the stack
	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "github_com_fullstack_lang_gongdoc_go_", // table name prefix
		},
	}
	db, err := gorm.Open(sqlite.Open(filepath), gormConfig)

	if err != nil {
		panic("Failed to connect to database!")
	}

	AutoMigrate(db)

	return db
}

// AutoMigrate migrates db with with orm Struct
func AutoMigrate(db *gorm.DB) {
	// adjust naming strategy to the stack
	db.Config.NamingStrategy = &schema.NamingStrategy{
		TablePrefix: "github_com_fullstack_lang_gongdoc_go_", // table name prefix
	}

	err := db.AutoMigrate( // insertion point for reference to structs
		&ClassdiagramDB{},
		&ClassshapeDB{},
		&FieldDB{},
		&GongStructDB{},
		&GongdocCommandDB{},
		&GongdocStatusDB{},
		&LinkDB{},
		&NoteDB{},
		&PkgeltDB{},
		&PositionDB{},
		&UmlStateDB{},
		&UmlscDB{},
		&VerticeDB{},
	)

	if err != nil {
		msg := err.Error()
		panic("problem with migration " + msg + " on package github.com/fullstack-lang/gongdoc/go")
	}
	// log.Printf("Database Migration of package github.com/fullstack-lang/gongdoc/go is OK")

	BackRepo.init(db)
}

func ResetDB(db *gorm.DB) { // insertion point for reference to structs
	db.Delete(&ClassdiagramDB{})
	db.Delete(&ClassshapeDB{})
	db.Delete(&FieldDB{})
	db.Delete(&GongStructDB{})
	db.Delete(&GongdocCommandDB{})
	db.Delete(&GongdocStatusDB{})
	db.Delete(&LinkDB{})
	db.Delete(&NoteDB{})
	db.Delete(&PkgeltDB{})
	db.Delete(&PositionDB{})
	db.Delete(&UmlStateDB{})
	db.Delete(&UmlscDB{})
	db.Delete(&VerticeDB{})
}