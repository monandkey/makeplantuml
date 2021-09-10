package user

func (h handsonUser) new() UserMethod {
	return &handsonUser{}
}

func (h *handsonUser) SetArgs(fileName string) {
	h.Args = []string{
		"-r", fileName,
		"-t", "ad",
		"-T", "fields",
		"-E", "separator=,",
		"-E", "quote=d",
		"-e", "frame.number",     //  0
		"-e", "_ws.col.Time",     //  1
		"-e", "ip.src",           //  2
		"-e", "ipv6.src",         //  3
		"-e", "udp.srcport",      //  4
		"-e", "tcp.srcport",      //  5
		"-e", "sctp.srcport",     //  6
		"-e", "ip.dst",           //  7
		"-e", "ipv6.dst",         //  8
		"-e", "udp.dstport",      //  9
		"-e", "tcp.dstport",      // 10
		"-e", "sctp.dstport",     // 11
		"-e", "_ws.col.Protocol", // 12
		"-e", "_ws.col.Info",     // 13
		"-e", "ip.len",           // 14
		"-e", "ipv6.plen",        // 15
		"-d", "tcp.port==29000-30000,http2",
	}
}
