// generated code - do not edit
package dbgorm

import (
	"github.com/fullstack-lang/gong/lib/button/go/db"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var _ db.DBInterface = (*DBWrapper)(nil)

func NewDBWrapper(filename, table_prefix string, dst ...interface{}) *DBWrapper {
	dbWrapper := new(DBWrapper)

	// adjust naming strategy to the stack
	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: table_prefix, // table name prefix
		},
	}
	db, err := gorm.Open(sqlite.Open(filename), gormConfig)

	// since testsim is a multi threaded application. It is important to set up
	// only one open connexion at a time
	dbDB_inMemory, err := db.DB()
	if err != nil {
		panic("cannot access DB of db" + err.Error())
	}
	// it is mandatory to allow parallel access, otherwise, bizarre errors occurs
	dbDB_inMemory.SetMaxOpenConns(1)

	// adjust naming strategy to the stack
	db.Config.NamingStrategy = &schema.NamingStrategy{
		TablePrefix: table_prefix, // table name prefix
	}

	err = db.AutoMigrate( // insertion point for reference to structs
		dst...,
	)

	if err != nil {
		msg := err.Error()
		panic("problem with migration " + msg + " on package", table_prefix)
	}

	dbWrapper.db = db

	return dbWrapper
}

// DBWrapper wraps *gorm.DB and implements the *DBWrapper
type DBWrapper struct {
	db *gorm.DB
}

// Create inserts the value into the database and returns a new *DBWrapper and an error if any.
func (dbw *DBWrapper) Create(instanceDB any) (db.DBInterface, error) {
	tx := dbw.db.Create(instanceDB)
	return &DBWrapper{db: tx}, tx.Error
}

// Unscoped returns a new *DBWrapper that includes soft-deleted records.
func (dbw *DBWrapper) Unscoped() (db.DBInterface, error) {
	tx := dbw.db.Unscoped()
	return &DBWrapper{db: tx}, tx.Error
}

// Model specifies the model you would like to run db operations against.
func (dbw *DBWrapper) Model(instanceDB any) (db.DBInterface, error) {
	tx := dbw.db.Model(instanceDB)
	return &DBWrapper{db: tx}, tx.Error
}

// Delete deletes value from the database.
func (dbw *DBWrapper) Delete(instanceDB any) (db.DBInterface, error) {
	tx := dbw.db.Delete(instanceDB)
	return &DBWrapper{db: tx}, tx.Error
}

// Save updates value in the database.
func (dbw *DBWrapper) Save(instanceDB any) (db.DBInterface, error) {
	tx := dbw.db.Save(instanceDB)
	return &DBWrapper{db: tx}, tx.Error
}

// Updates updates multiple columns.
func (dbw *DBWrapper) Updates(instanceDB any) (db.DBInterface, error) {
	tx := dbw.db.Updates(instanceDB)
	return &DBWrapper{db: tx}, tx.Error
}

// Find finds records that match given conditions.
func (dbw *DBWrapper) Find(instanceDBs any) (db.DBInterface, error) {
	tx := dbw.db.Find(instanceDBs)
	return &DBWrapper{db: tx}, tx.Error
}

// First returns the first record that matches the conditions.
func (dbw *DBWrapper) First(instanceDB any, conds ...any) (db.DBInterface, error) {
	tx := dbw.db.First(instanceDB, conds...)
	return &DBWrapper{db: tx}, tx.Error
}
