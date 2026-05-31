package viper

import (
	"context"
	"log/slog"
	"strings"

	"github.com/spf13/viper"
	"github.com/thumbrise/ouro/contract"
)

type Loader struct {
	viper *viper.Viper
}

func NewLoader(logger *slog.Logger) *Loader {
	vp := viper.NewWithOptions(
		viper.WithLogger(logger),
		viper.ExperimentalBindStruct(),
		viper.EnvKeyReplacer(strings.NewReplacer(".", "_")),
	)

	vp.SetTypeByDefaultValue(true)
	vp.AutomaticEnv()

	return &Loader{
		viper: vp,
	}
}

func (c Loader) Load(ctx context.Context, loadContext contract.LoadContext) error {
	var err error

	if loadContext.Key == "" {
		err = c.viper.Unmarshal(loadContext.Data)
	} else {
		err = c.viper.UnmarshalKey(loadContext.Key, loadContext.Data)
	}

	return err
}
