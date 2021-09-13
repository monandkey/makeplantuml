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
	uml.UmlArgs
}

type normalUser struct {
	baseUser
}

type handsonUser struct {
	baseUser
}

type toWritingUser struct {
	baseUser
}

type fromRenderingUser struct {
	baseUser
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
