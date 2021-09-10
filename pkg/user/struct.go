package user

import (
	"local.packages/tshark"
)
type NormalUser struct {
	tshark.TsharkArgs
}

type flags struct {
	i int
}

type Option func(*flags)
