//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by iocli

package main

import (
	autowire "github.com/alibaba/ioc-golang/autowire"
	"github.com/alibaba/ioc-golang/autowire/normal"
	singleton "github.com/alibaba/ioc-golang/autowire/singleton"
	util "github.com/alibaba/ioc-golang/autowire/util"
)

func init() {
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &app_{}
		},
	})
	singleton.RegisterStructDescriptor(&autowire.StructDescriptor{
		Alias: "appalias",
		Factory: func() interface{} {
			return &App{}
		},
	})
}

type app_ struct {
	Run_ func()
}

func (a *app_) Run() {
	a.Run_()
}
func GetApp() (*App, error) {
	i, err := singleton.GetImpl(util.GetSDIDByStructPtr(new(App)))
	if err != nil {
		return nil, err
	}
	impl := i.(*App)
	return impl, nil
}
