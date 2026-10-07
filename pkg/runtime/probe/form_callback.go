package probe

type FormCallbackIF interface {
	GetCreationMode() bool
	GetInstance() any
	GetGongstructName() string
	OnSave()
}
