package main

import "github.com/spf13/viper"

type OuroBoros struct {
	viper *viper.Viper
	opts  *OuroBorosOptions
}

func Boros(vp *viper.Viper, opts *OuroBorosOptions) *OuroBoros {
	if vp == nil {
		panic("viper argument must be passed")
	}

	return &OuroBoros{
		viper: vp,
		opts:  opts,
	}
}

func (o *OuroBoros) SetConfigFile(file string) {}

type OuroBorosOptions struct {
	//nolint:unused // scaffold
	configTypes []interface{}
}
