package contract

type Binder interface {
	BindConfig(config interface{}) error
}
