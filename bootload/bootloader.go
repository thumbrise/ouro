package bootload

import (
	"github.com/thumbrise/ouro/bootload/contract"
)

type Bootloader struct {
	modules []contract.Module
	binder  Binder
}

func NewBootloader() *Bootloader {
	return &Bootloader{
		binder: Binder{},
	}
}

func (b *Bootloader) AddModules(modules ...contract.Module) {
	b.modules = append(b.modules, modules...)
}

func (b *Bootloader) Boot() {
	for _, module := range b.modules {
		err := b.bootModule(module)
		if err != nil {
			panic(err)
		}
	}
}

func (b *Bootloader) bootModule(m contract.Module) error {
	return m.Bind(b.binder)
}
