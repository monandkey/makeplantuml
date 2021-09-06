package tshark

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
