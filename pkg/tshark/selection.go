package tshark

// func NormalMode() Option {
// 	return func(f *flags) {
// 		f.i = 1
// 	}
// }

// func HandonMode() Option {
// 	return func(f *flags) {
// 		f.i = 2
// 	}
// }

// func UseTsharkSelection(options ...Option) TsharkMethod {
// 	f := flags{i: 1}
// 	for _, option := range options {
// 		option(&f)
// 		if f.i == 1 {
// 			return NormalTshark{}.new()

// 		} else if f.i == 2 {
// 			return HandsonTshark{}.new()
// 		}
// 	}
// 	return NormalTshark{}.new()
// }
