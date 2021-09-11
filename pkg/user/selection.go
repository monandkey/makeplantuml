package user

func Config() Option {
	return func(f *flags) {
		f.i = 0
	}
}

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

func ToWriting() Option {
	return func(f *flags) {
		f.i = 3
	}
}

func UserSelection(options ...Option) UserMethod {
	f := flags{i: 1}
	for _, option := range options {
		option(&f)
		if f.i == 1 {
			return normalUser{}.new()

		} else if f.i == 2 {
			return handsonUser{}.new()

		} else if f.i == 3 {
			return toWritingUser{}.new()
		}
	}
	return normalUser{}.new()
}

func ConfigUserSelection(options ...Option) ConfigUserMethod {
	f := flags{i: 0}
	for _, option := range options {
		option(&f)
		if f.i == 0 {
			return configOpeUser{}.new()
		}
	}
	return configOpeUser{}.new()
}
