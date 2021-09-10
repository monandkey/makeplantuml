package user

import (
	"local.packages/cfg"
	"local.packages/uml"
	"local.packages/tshark"
)

type UserMethod interface {
	tshark.TsharkMethod
	uml.UmlMethod
	Display()
}

type baseUser struct {
	tshark.TsharkArgs
}

type normalUser struct {
	baseUser
}

type handsonUser struct {
	normalUser
}


type ConfigUserMethod interface {
	cfg.ConfigMethod
}

type configBaseUser struct {
	cfg.Config
}

type configOpeUser struct {
	configBaseUser
}


type flags struct {
	i int
}

type Option func(*flags)
