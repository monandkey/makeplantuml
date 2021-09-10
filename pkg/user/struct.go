package user

import (
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

type flags struct {
	i int
}

type Option func(*flags)
