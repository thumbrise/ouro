package banana_shop

import (
	"github.com/thumbrise/ouro/bootload/contract"
)

type Module struct{}

func (m *Module) Name() string {
	return "Banana Shop"
}

func (m *Module) Bind(binder contract.Binder) error {
	return binder.BindConfig(Config{})
}
