package tshark

type TsharkArgs struct {
	Cmd    string
	Args   []string
	Out    []byte
	Header []map[string]string
}

type TsharkMethod interface {
	SetCmd()
	SetArgs(string)
	RunE() error
	Parse()
	NameResE(string) error
}
