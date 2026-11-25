package models

const (
	DataStoreDefaultWidth  = 140.0
	DataStoreDefaultHeight = 80.0
	DataStoreMinWidth      = 60.0
	DataStoreMinHeight     = 48.0
)

type DataStoreShape struct {
	Name string

	DataStore *DataStore

	RectShape
}

func (s *DataStoreShape) GetAbstractElement() AbstractType {
	if s.DataStore == nil {
		return nil
	}
	return s.DataStore
}

func (s *DataStoreShape) SetAbstractElement(abstractElement AbstractType) {
	s.DataStore = abstractElement.(*DataStore)
}

var _ ConcreteType = (*DataStoreShape)(nil)
var _ RectShapeInterface = (*DataStoreShape)(nil)
