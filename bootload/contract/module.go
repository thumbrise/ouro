package contract

type Module interface {
	Name() string
	Bind(binder Binder) error
}
