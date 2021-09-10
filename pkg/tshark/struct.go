package tshark

type TsharkHeaders []TsharkHeader

type TsharkHeader struct {
	Number   string
	Time     string
	SrcAddr  string
	SrcPort  string
	DstAddr  string
	DstPort  string
	Protocol string
	Message  string
	Checksum string
}

type DuplicatePacketeDropTshark struct {
	Number   string
	Time     string
	SrcAddr  string
	SrcPort  string
	DstAddr  string
	DstPort  string
	Protocol string
	Message  string
	Checksum string
	
}

type TsharkArgs struct {
	Cmd    string
	Args   []string
	Out    []byte
	Header []map[string]string
}

type HandsonTshark struct {
	TsharkArgs
	res TsharkHeaders
}

type TsharkMethod interface {
	SetCmd()
	SetArgs(string)
	RunE() error
	Parse()
	NameResE(string) error
}
