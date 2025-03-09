// generated code - do not edit
package db

type DBInterface interface {
	Create(instanceDB any) (DBInterface, error)
	Unscoped() (DBInterface, error)
	Model(instanceDB any) (DBInterface, error)
	Delete(instanceDB any) (DBInterface, error)
	Save(instanceDB any) (DBInterface, error)
	Updates(instanceDB any) (DBInterface, error)
	Find(instanceDBs any) (DBInterface, error)
	First(instanceDB any, conds ...any) (DBInterface, error)
}
