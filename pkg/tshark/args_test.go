package tshark

import (
	"testing"
	"local.packages/tshark"
)

var expWantN = []string{
	"-r", "test.pcap",
	"-t", "ad",
	"-T", "fields",
	"-E", "separator=,",
	"-E", "quote=d",
	"-e", "frame.number",
	"-e", "_ws.col.Time",
	"-e", "ip.src",
	"-e", "ipv6.src",
	"-e", "udp.srcport",
	"-e", "tcp.srcport",
	"-e", "sctp.srcport",
	"-e", "ip.dst",
	"-e", "ipv6.dst",
	"-e", "udp.dstport",
	"-e", "tcp.dstport",
	"-e", "sctp.dstport",
	"-e", "_ws.col.Protocol",
	"-e", "_ws.col.Info",
	"-e", "ip.len",
	"-e", "ipv6.plen",
}

var expWantH = []string{
	"-r", "test.pcap",
	"-t", "ad",
	"-T", "fields",
	"-E", "separator=,",
	"-E", "quote=d",
	"-e", "frame.number",
	"-e", "_ws.col.Time",
	"-e", "ip.src",
	"-e", "ipv6.src",
	"-e", "udp.srcport",
	"-e", "tcp.srcport",
	"-e", "sctp.srcport",
	"-e", "ip.dst",
	"-e", "ipv6.dst",
	"-e", "udp.dstport",
	"-e", "tcp.dstport",
	"-e", "sctp.dstport",
	"-e", "_ws.col.Protocol",
	"-e", "_ws.col.Info",
	"-e", "ip.len",
	"-e", "ipv6.plen",
	"-d", "tcp.port==29000-30000,http2",
}

var expWantA = []string{
	"-r", "test.pcap",
	"-t", "ad",
	"-T", "fields",
	"-E", "separator=,",
	"-E", "quote=d",
	"-e", "frame.number",
	"-e", "_ws.col.Time",
	"-e", "ip.src",
	"-e", "ipv6.src",
	"-e", "udp.srcport",
	"-e", "tcp.srcport",
	"-e", "sctp.srcport",
	"-e", "ip.dst",
	"-e", "ipv6.dst",
	"-e", "udp.dstport",
	"-e", "tcp.dstport",
	"-e", "sctp.dstport",
	"-e", "_ws.col.Protocol",
	"-e", "_ws.col.Info",
	"-e", "ip.len",
	"-e", "ipv6.plen",
	"-e", "nas_eps.emm.type_of_id",
	"-e", "nas_eps.emm.dcnr_cap",
	"-e", "nas_5gs.mm.type_id",
	"-e", "gtpv2.oi",
	"-e", "gtpv2.si",
	"-e", "gtpv2.dcnr",
	"-e", "gtpv2.cause",
	"-e", "pfcp.cause",
	"-e", "diameter.Result-Code",
	"-e", "diameter.CC-Request-Type",
}

func TestGetArgs(t *testing.T) {
	type Tests struct {
		name string
		args string
		file string
		want []string
	}

	tests := []Tests{
		{
			name: "Normal Case",
			args: "normal",
			file: "test.pcap",
			want: expWantN,
		},
		{
			name: "Handson Case",
			args: "handson",
			file: "test.pcap",
			want: expWantH,
		},
		{
			name: "Annotation Case",
			args: "annotation",
			file: "test.pcap",
			want: expWantA,
		},
		{
			name: "Error Case",
			args: "error",
			file: "test.pcap",
			want: []string{},
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			cmd := tshark.TsharkArgs{}
			cmd.Args = cmd.GetArgs(v.args, v.file)

			for i, _ := range cmd.Args {
				if cmd.Args[i] != v.want[i] {
					t.Errorf("The result is not the expected value.")
				}
			}
		})
	}
}