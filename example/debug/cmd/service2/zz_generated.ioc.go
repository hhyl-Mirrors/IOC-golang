//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by iocli

package service2

import (
	autowire "github.com/alibaba/ioc-golang/autowire"
	normal "github.com/alibaba/ioc-golang/autowire/normal"
	singleton "github.com/alibaba/ioc-golang/autowire/singleton"
	util "github.com/alibaba/ioc-golang/autowire/util"
)

func init() {
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &impl1_{}
		},
	})
	singleton.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &Impl1{}
		},
	})
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &impl2_{}
		},
	})
	singleton.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &Impl2{}
		},
	})
}

type impl1_ struct {
	Hello_ func(input string) string
}

func (i *impl1_) Hello(input string) string {
	return i.Hello_(input)
}

type impl2_ struct {
	Hello_ func(input string) string
}

func (i *impl2_) Hello(input string) string {
	return i.Hello_(input)
}
func GetImpl1() (*Impl1, error) {
	i, err := singleton.GetImpl(util.GetSDIDByStructPtr(new(Impl1)))
	if err != nil {
		return nil, err
	}
	impl := i.(*Impl1)
	return impl, nil
}
func GetImpl2() (*Impl2, error) {
	i, err := singleton.GetImpl(util.GetSDIDByStructPtr(new(Impl2)))
	if err != nil {
		return nil, err
	}
	impl := i.(*Impl2)
	return impl, nil
}
