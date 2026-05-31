package main

import (
	"github.com/thumbrise/ouro/bootload"
	"github.com/thumbrise/ouro/bootloader_fixture/modules/banana_shop"
)

func main() {
	b := bootload.NewBootloader()
	b.AddModules(&banana_shop.Module{})
	b.Boot()
}
