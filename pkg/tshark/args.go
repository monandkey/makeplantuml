package tshark

func (t TsharkArgs) GetArgs(kinds string, fileName string) []string {
	switch(kinds) {
		case "normal":
			return t.SetArgsNormal(fileName)

		case "handson":
			return t.SetArgsHandson(fileName)

		case "annotation":
			return t.SetArgsAnnotatio(fileName)
	}
	return t.Args
}

func (t TsharkArgs) SetArgsNormal(fileName string) []string {
	t.Args = []string{
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
	}
	return t.Args
}

func (t TsharkArgs) SetArgsHandson(fileName string) []string {
	t.Args = []string{
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
	return t.Args
}

func (t TsharkArgs) SetArgsAnnotatio(fileName string) []string {
	t.Args = []string{
		"-r", fileName,
		"-t", "ad",
		"-T", "fields",
		"-E", "separator=,",
		"-E", "quote=d",
		"-e", "frame.number",             //  0
		"-e", "_ws.col.Time",             //  1
		"-e", "ip.src",                   //  2
		"-e", "ipv6.src",                 //  3
		"-e", "udp.srcport",              //  4
		"-e", "tcp.srcport",              //  5
		"-e", "sctp.srcport",             //  6
		"-e", "ip.dst",                   //  7
		"-e", "ipv6.dst",                 //  8
		"-e", "udp.dstport",              //  9
		"-e", "tcp.dstport",              // 10
		"-e", "sctp.dstport",             // 11
		"-e", "_ws.col.Protocol",         // 12
		"-e", "_ws.col.Info",             // 13
		"-e", "ip.len",                   // 14
		"-e", "ipv6.plen",                // 15
		"-e", "nas_eps.emm.type_of_id",   // 16
        "-e", "nas_eps.emm.dcnr_cap",     // 17
		"-e", "nas_5gs.mm.type_id",       // 18
		"-e", "gtpv2.oi",                 // 19
		"-e", "gtpv2.si",                 // 20
		"-e", "gtpv2.dcnr",               // 21
		"-e", "gtpv2.cause",              // 22
		"-e", "pfcp.cause",               // 23
		"-e", "diameter.Result-Code",     // 24
		"-e", "diameter.CC-Request-Type", // 25
	}
	return t.Args
}
