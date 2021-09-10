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

type normalUser struct {
	tshark.TsharkArgs
}

type flags struct {
	i int
}

type Option func(*flags)
