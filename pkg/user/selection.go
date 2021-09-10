package user

import (
	"local.packages/tshark"
)

func Normal() Option {
	return func(f *flags) {
		f.i = 1
	}
}

func Handon() Option {
	return func(f *flags) {
		f.i = 2
	}
}

func UseTsharkSelection(options ...Option) tshark.TsharkMethod {
	f := flags{i: 1}
	for _, option := range options {
		option(&f)
		if f.i == 1 {
			return NormalUser{}.new()

		} else if f.i == 2 {
			// return HandsonTshark{}.new()
		}
	}
	return NormalUser{}.new()
}
