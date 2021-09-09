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
	tshark string
	cmd    []string
	out    []byte
}

type NormalTshark struct {
	TsharkArgs
}

type HandsonTshark struct {
	TsharkArgs
	res TsharkHeaders
}

type TsharkMethod interface {
	SetTsharkCommand()
	CreateCommand()
	RunE() error
	SetHeader()
}

type flags struct {
	i int
}

type Option func(*flags)
