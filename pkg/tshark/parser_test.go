package tshark

import (
	"testing"
	"local.packages/tshark"
)

var expOut = `
"1","2021-06-16 10:34:38.377646","172.16.10.10",,,,"46571","172.16.10.20",,,,"38412","NGAP/NAS-5GS","InitialUEMessage, Registration request",,"0x527eec87",,"124",
"2","2021-06-16 10:34:38.592668","10.244.166.179",,,"50718",,"10.103.54.119",,,"29510",,"HTTP2","HEADERS[3]: GET /nnrf-disc/v1/nf-instances?requester-nf-type=AMF&target-nf-type=AUSF",,,"0x0000f32b","180",
"3","2021-06-16 10:34:38.592685","10.244.166.179",,,"50718",,"10.244.166.129",,,"29510",,"HTTP2","HEADERS[3]: GET /nnrf-disc/v1/nf-instances?requester-nf-type=AMF&target-nf-type=AUSF",,,"0x000063c3","180",
"4","2021-06-16 10:34:38.938723","10.244.166.129",,,"29510",,"10.244.166.179",,,"50718",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x00006377","104",
"5","2021-06-16 10:34:38.938750","10.103.54.119",,,"29510",,"10.244.166.179",,,"50718",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x0000f2df","104",
"6","2021-06-16 10:34:38.950965","10.244.166.179",,,"57206",,"10.103.99.231",,,"29509",,"HTTP2","HEADERS[3]: POST /nausf-auth/v1/ue-authentications",,,"0x0000209e","182",
"7","2021-06-16 10:34:38.950970","10.244.166.179",,,"57206",,"10.244.166.157",,,"29509",,"HTTP2","HEADERS[3]: POST /nausf-auth/v1/ue-authentications",,,"0x000063e1","182",
"8","2021-06-16 10:34:39.151550","10.244.166.157",,,"45510",,"10.103.54.119",,,"29510",,"HTTP2","HEADERS[3]: GET /nnrf-disc/v1/nf-instances?requester-nf-type=AUSF&service-names=nudm-ueau&target-nf-type=UDM",,,"0x0000f327","198",
"9","2021-06-16 10:34:39.151555","10.244.166.157",,,"45510",,"10.244.166.129",,,"29510",,"HTTP2","HEADERS[3]: GET /nnrf-disc/v1/nf-instances?requester-nf-type=AUSF&service-names=nudm-ueau&target-nf-type=UDM",,,"0x000063bf","198",
"10","2021-06-16 10:34:39.165688","10.244.166.129",,,"29510",,"10.244.166.157",,,"45510",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x00006361","104",
`

func TestSetAddress(t *testing.T) {
	type address struct {
		v4    string
		v6    string
		lenv4 string
		lenv6 string
	}

	type Tests struct {
		name   string
		args   address
		want   string
	}

	tests := []Tests {
		{
			name: "v4 Only",
			args: address{
				v4:    "1.1.1.1",
				v6:    "",
				lenv4: "",
				lenv6: "",
			},
			want: "1.1.1.1",
		},
		{
			name: "v6 Only",
			args: address{
				v4:    "",
				v6:    "2001::0",
				lenv4: "",
				lenv6: "",
			},
			want: "2001::0",
		},
		{
			name: "v4v6 Dual",
			args: address{
				v4:    "1.1.1.1",
				v6:    "2001::0",
				lenv4: "100",
				lenv6: "120",
			},
			want: "1.1.1.1",
		},
		{
			name: "v4 Dual",
			args: address{
				v4:    "1.1.1.1,2.2.2.2",
				v6:    "",
				lenv4: "100,200",
				lenv6: "",
			},
			want: "1.1.1.1",
		},
		{
			name: "v6 Dual",
			args: address{
				v4:    "",
				v6:    "2001::0,2001::1",
				lenv4: "",
				lenv6: "200,100",
			},
			want: "2001::1",
		},
		{
			name: "v4v6 Dual Exp",
			args: address{
				v4:    "1.1.1.1,2.2.2.2",
				v6:    "2001::0,2001::1",
				lenv4: "50,150",
				lenv6: "200,100",
			},
			want: "1.1.1.1",
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			cmd := tshark.TsharkArgs{}
			res := cmd.SetAddress(v.args.v4, v.args.v6, v.args.lenv4, v.args.lenv6)
			if res != v.want {
				t.Errorf("The return value is not the expected value.\nres: %s\n", res)
			}
		})
	}
}

func TestSetSetPortAndCheckSum(t *testing.T) {
	type port struct {
		udp  string
		tcp  string
		sctp string
	}

	type Tests struct {
		name   string
		args   port
		want   string
	}

	tests := []Tests{
		{
			name: "UDP Only",
			args: port{
				udp:  "2123",
				tcp:  "",
				sctp: "",
			},
			want: "2123",
		},
		{
			name: "TCP Only",
			args: port{
				udp:  "",
				tcp:  "3868",
				sctp: "",
			},
			want: "3868",
		},
		{
			name: "SCTP Only",
			args: port{
				udp:  "",
				tcp:  "",
				sctp: "36412",
			},
			want: "36412",
		},
		{
			name: "Unanticipated cases",
			args: port{
				udp:  "2123",
				tcp:  "3686",
				sctp: "36412",
			},
			want: "",
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			cmd := tshark.TsharkArgs{}
			res := cmd.SetPortAndCheckSum(v.args.udp, v.args.tcp, v.args.sctp)
			if res != v.want {
				t.Errorf("The return value is not the expected value.")
			}
		})
	}
}

func TestSetMessage(t *testing.T) {
	type required struct {
		message  string
		protocol string
	}

	type Tests struct {
		name   string
		args   required
		want   string
	}

	tests := []Tests{
		{
			name: "GTPv2",
			args: required{
				message:  "Create Session Request",
				protocol: "GTPv2",
			},
			want: "Create Session Request",
		},
		{
			name: "PFCP",
			args: required{
				message:  "PFCP Session Establishment Request",
				protocol: "PFCP",
			},
			want: "PFCP Session Establishment Request",
		},
		{
			name: "DIAMETER Request",
			args: required{
				message:  "cmd=Capabilities-Exchange Request(257) flags=R--- appl=Diameter Common Messages(0) h2h=7c8a72c3 e2e=f3d80eea",
				protocol: "DIAMETER",
			},
			want: "Capabilities-Exchange Request",
		},
		{
			name: "DIAMETER Answer",
			args: required{
				message:  "cmd=Capabilities-Exchange Answer(257) flags=---- appl=Diameter Common Messages(0) h2h=7c8a72c3 e2e=f3d80eea",
				protocol: "DIAMETER",
			},
			want: "Capabilities-Exchange Answer",
		},
		{
			name: "3GPP DIAMETER Request",
			args: required{
				message:  "cmd=3GPP-Authentication-Information Request(318) flags=RP-- appl=3GPP S6a/S6d(16777251) h2h=7c8a72c4 e2e=f3d80eeb | ",
				protocol: "DIAMETER",
			},
			want: "Authentication-Information Request",
		},
		{
			name: "3GPP DIAMETER Answer",
			args: required{
				message:  "SACK cmd=3GPP-Authentication-Information Answer(318) flags=-P-- appl=3GPP S6a/S6d(16777251) h2h=7c8a72c4 e2e=f3d80eeb | ",
				protocol: "DIAMETER",
			},
			want: "Authentication-Information Answer",
		},
		{
			name: "S1AP",
			args: required{
				message:  "S1SetupRequest",
				protocol: "S1AP",
			},
			want: "S1SetupRequest",
		},
		{
			name: "NGAP",
			args: required{
				message: "UplinkNASTransport",
				protocol: "NGAP/NAS-5GS",
			},
			want: "UplinkNASTransport",
		},
		{
			name: "HTTP2 Method GET",
			args: required{
				message:  "HEADERS[3]: GET /nnrf-disc/v1/nf-instances?requester-nf-type=SMF&target-nf-type=PCF",
				protocol: "HTTP2",
			},
			want: "GET /nnrf-disc/v1/nf-instances",
		},
		{
			name: "HTTP2 Method HEAD",
			args: required{
				message:  "",
				protocol: "HTTP2",
			},
			want: "",
		},
		{
			name: "HTTP2 Method POST",
			args: required{
				message:  "HEADERS[3]: POST /nsmf-pdusession/v1/sm-contexts/urn:uuid:d005b19a-e01e-4198-8a7c-1b68923ef52d/modify",
				protocol: "HTTP2",
			},
			want: "POST /nsmf-pdusession/v1/sm-contexts/urn:uuid:d005b19a-e01e-4198-8a7c-1b68923ef52d/modify",
		},
		{
			name: "HTTP2 Method PUT",
			args: required{
				message:  "HEADERS[3]: PUT /nudr-dr/v1/subscription-data/imsi-208930000000003/context-data/amf-3gpp-access",
				protocol: "HTTP2",
			},
			want: "PUT /nudr-dr/v1/subscription-data/imsi-208930000000003/context-data/amf-3gpp-access",
		},
		{
			name: "HTTP2 Method DELETE",
			args: required{
				message:  "HEADERS[3]: DELETE /npcf-am-policy-control/v1/policies/imsi-208930000000003-1",
				protocol: "HTTP2",
			},
			want: "DELETE /npcf-am-policy-control/v1/policies/imsi-208930000000003-1",
		},
		{
			name: "HTTP2 Method CONNECT",
			args: required{
				message:  "",
				protocol: "HTTP2",
			},
			want: "",
		},
		{
			name: "HTTP2 Method OPTIONS",
			args: required{
				message:  "",
				protocol: "HTTP2",
			},
			want: "",
		},
		{
			name: "HTTP2 Method TRACE",
			args: required{
				message:  "",
				protocol: "HTTP2",
			},
			want: "",
		},
		{
			name: "HTTP2 Method PATCH",
			args: required{
				message:  "HEADERS[3]: PATCH /nudr-dr/v1/subscription-data/imsi-208930000000003/authentication-data/authentication-subscription",
				protocol: "HTTP2",
			},
			want: "PATCH /nudr-dr/v1/subscription-data/imsi-208930000000003/authentication-data/authentication-subscription",
		},
		{
			name: "HTTP2 Method Response",
			args: required{
				message:  "HEADERS[3]: 200 OK",
				protocol: "HTTP2",
			},
			want: "200 OK",
		},
		{
			name: "HTTP2 PDU",
			args: required{
				message:  "DATA[3], JavaScript Object Notation (application/json), PDU session establishment accept (PDU session type IPv4 only allowed)",
				protocol: "HTTP2/JSON/NAS-5GS/NGAP",
			},
			want: "PDU session establishment accept (PDU session type IPv4 only allowed)",
		},
		{
			name: "HTTP2 Drop",
			args: required{
				message:  "SETTINGS[0], WINDOW_UPDATE[0]",
				protocol: "HTTP2",
			},
			want: "",
		},
		{
			name: "TCP Drop",
			args: required{
				message:  "8080 → 47970 [ACK] Seq=1 Ack=2 Win=27136 Len=0 TSval=1042549644 TSecr=1042549643",
				protocol: "TCP",
			},
			want: "",
		},
		{
			name: "SCTP Drop",
			args: required{
				message:  "HEARTBEAT",
				protocol: "SCTP",
			},
			want: "",
		},
		{
			name: "ICMP",
			args: required{
				message:  "Echo (ping) request  id=0x7348, seq=27032/39017, ttl=255 (no response found!)",
				protocol: "ICMP",
			},
			want: "Echo (ping) request  id=0x7348, seq=27032/39017, ttl=255 (no response found!)",
		},
		{
			name: "ICMPv6",
			args: required{
				message:  "Echo (ping) request  id=0x7348, seq=27032/39017, ttl=255 (no response found!)",
				protocol: "ICMPv6",
			},
			want: "Echo (ping) request  id=0x7348, seq=27032/39017, ttl=255 (no response found!)",
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			cmd := tshark.TsharkArgs{}
			res := cmd.SetMessage(v.args.message, v.args.protocol)
			if res != v.want {
				t.Errorf("The return value is not the expected value.\n %s", res)
			}
		})
	}
}

func TestSetAnnotation(t *testing.T) {
	type Tests struct {
		name   string
		args   [][]string
		want   string
	}

	tests := []Tests{
		{
			name: "S1AP Type of id",
			args: [][]string{
				{"1"},{""},{""},{""},{""},{""},{""},{""},{""},{""},{""},{""},
				{"S1AP"}, // 12
				{"0"}, // 13
				{"0"}, // 14
				{"0"}, // 15
				{"1"}, // 16
				{"0"}, // 17
				{"0"}, // 18
				{"0"}, // 19
				{"0"}, // 20
				{"0"}, // 21
				{"0"}, // 22
				{"0"}, // 23
				{"0"}, // 24
				{"0"}, // 25
			},
			want: "Attach type: IMSI",
		},
		{
			name: "S1AP DCNR",
			args: [][]string{
				{"1"},{""},{""},{""},{""},{""},{""},{""},{""},{""},{""},{""},
				{"S1AP"}, // 12
				{"0"}, // 13
				{"0"}, // 14
				{"0"}, // 15
				{"0"}, // 16
				{"1"}, // 17
				{"0"}, // 18
				{"0"}, // 19
				{"0"}, // 20
				{"0"}, // 21
				{"0"}, // 22
				{"0"}, // 23
				{"0"}, // 24
				{"0"}, // 25
			},
			want: "DCNR: supported",
		},
		{
			name: "S1AP Type of id and DCNR",
			args: [][]string{
				{"1"},{""},{""},{""},{""},{""},{""},{""},{""},{""},{""},{""},
				{"S1AP"}, // 12
				{"0"}, // 13
				{"0"}, // 14
				{"0"}, // 15
				{"1"}, // 16
				{"1"}, // 17
				{"0"}, // 18
				{"0"}, // 19
				{"0"}, // 20
				{"0"}, // 21
				{"0"}, // 22
				{"0"}, // 23
				{"0"}, // 24
				{"0"}, // 25
			},
			want: "Attach type: IMSI\\nDCNR: supported",
		},
		{
			name: "NGAP Type of id",
			args: [][]string{
				{"1"},{""},{""},{""},{""},{""},{""},{""},{""},{""},{""},{""},
				{"NGAP/NAS-5GS"}, // 12
				{"0"}, // 13
				{"0"}, // 14
				{"0"}, // 15
				{"0"}, // 16
				{"0"}, // 17
				{"1"}, // 18
				{"0"}, // 19
				{"0"}, // 20
				{"0"}, // 21
				{"0"}, // 22
				{"0"}, // 23
				{"0"}, // 24
				{"0"}, // 25
			},
			want: "Attach type: SUCI",
		},
		{
			name: "GTPv2 OI",
			args: [][]string{
				{"1"},{""},{""},{""},{""},{""},{""},{""},{""},{""},{""},{""},
				{"GTPv2"}, // 12
				{"0"}, // 13
				{"0"}, // 14
				{"0"}, // 15
				{"0"}, // 16
				{"0"}, // 17
				{"0"}, // 18
				{"1"}, // 19
				{"0"}, // 20
				{"0"}, // 21
				{"0"}, // 22
				{"0"}, // 23
				{"0"}, // 24
				{"0"}, // 25
			},
			want: "OI: supported",
		},
		{
			name: "GTPv2 SI",
			args: [][]string{
				{"1"},{""},{""},{""},{""},{""},{""},{""},{""},{""},{""},{""},
				{"GTPv2"}, // 12
				{"0"}, // 13
				{"0"}, // 14
				{"0"}, // 15
				{"0"}, // 16
				{"0"}, // 17
				{"0"}, // 18
				{"0"}, // 19
				{"1"}, // 20
				{"0"}, // 21
				{"0"}, // 22
				{"0"}, // 23
				{"0"}, // 24
				{"0"}, // 25
			},
			want: "SI: supported",
		},
		{
			name: "GTPv2 DCNR",
			args: [][]string{
				{"1"},{""},{""},{""},{""},{""},{""},{""},{""},{""},{""},{""},
				{"GTPv2"}, // 12
				{"0"}, // 13
				{"0"}, // 14
				{"0"}, // 15
				{"0"}, // 16
				{"0"}, // 17
				{"0"}, // 18
				{"0"}, // 19
				{"0"}, // 20
				{"1"}, // 21
				{"0"}, // 22
				{"0"}, // 23
				{"0"}, // 24
				{"0"}, // 25
			},
			want: "DCNR: supported",
		},
		{
			name: "GTPv2 Cause",
			args: [][]string{
				{"1"},{""},{""},{""},{""},{""},{""},{""},{""},{""},{""},{""},
				{"GTPv2"}, // 12
				{"0"}, // 13
				{"0"}, // 14
				{"0"}, // 15
				{"0"}, // 16
				{"0"}, // 17
				{"0"}, // 18
				{"0"}, // 19
				{"0"}, // 20
				{"0"}, // 21
				{"16"}, // 22
				{"0"}, // 23
				{"0"}, // 24
				{"0"}, // 25
			},
			want: "Cause: Request accepted",
		},
		{
			name: "GTPv2 OI and Cause",
			args: [][]string{
				{"1"},{""},{""},{""},{""},{""},{""},{""},{""},{""},{""},{""},
				{"GTPv2"}, // 12
				{"0"}, // 13
				{"0"}, // 14
				{"0"}, // 15
				{"0"}, // 16
				{"0"}, // 17
				{"0"}, // 18
				{"1"}, // 19
				{"0"}, // 20
				{"0"}, // 21
				{"16"}, // 22
				{"0"}, // 23
				{"0"}, // 24
				{"0"}, // 25
			},
			want: "OI: supported\\nCause: Request accepted",
		},
		{
			name: "PFCP Cause",
			args: [][]string{
				{"1"},{""},{""},{""},{""},{""},{""},{""},{""},{""},{""},{""},
				{"PFCP"}, // 12
				{"0"}, // 13
				{"0"}, // 14
				{"0"}, // 15
				{"0"}, // 16
				{"0"}, // 17
				{"0"}, // 18
				{"0"}, // 19
				{"0"}, // 20
				{"0"}, // 21
				{"0"}, // 22
				{"1"}, // 23
				{"0"}, // 24
				{"0"}, // 25
			},
			want: "Cause: Request accepted (success)",
		},
		{
			name: "Diameter Result Code",
			args: [][]string{
				{"1"},{""},{""},{""},{""},{""},{""},{""},{""},{""},{""},{""},
				{"DIAMETER"}, // 12
				{"0"}, // 13
				{"0"}, // 14
				{"0"}, // 15
				{"0"}, // 16
				{"0"}, // 17
				{"0"}, // 18
				{"0"}, // 19
				{"0"}, // 20
				{"0"}, // 21
				{"0"}, // 22
				{"0"}, // 23
				{"2001"}, // 24
				{"0"}, // 25
			},
			want: "Cause: DIAMETER_SUCCESS",
		},
		{
			name: "Diameter CC RequestType",
			args: [][]string{
				{"1"},{""},{""},{""},{""},{""},{""},{""},{""},{""},{""},{""},
				{"DIAMETER"}, // 12
				{"0"}, // 13
				{"0"}, // 14
				{"0"}, // 15
				{"0"}, // 16
				{"0"}, // 17
				{"0"}, // 18
				{"0"}, // 19
				{"0"}, // 20
				{"0"}, // 21
				{"0"}, // 22
				{"0"}, // 23
				{"0"}, // 24
				{"1"}, // 25
			},
			want: "CC-Request-Type: INITIAL_REQUEST",
		},
		{
			name: "Diameter Result Code and CC RequestType",
			args: [][]string{
				{"1"},{""},{""},{""},{""},{""},{""},{""},{""},{""},{""},{""},
				{"DIAMETER"}, // 12
				{"0"}, // 13
				{"0"}, // 14
				{"0"}, // 15
				{"0"}, // 16
				{"0"}, // 17
				{"0"}, // 18
				{"0"}, // 19
				{"0"}, // 20
				{"0"}, // 21
				{"0"}, // 22
				{"0"}, // 23
				{"2001"}, // 24
				{"1"}, // 25
			},
			want: "Cause: DIAMETER_SUCCESS\\nCC-Request-Type: INITIAL_REQUEST",
		},
		{
			name: "Empty Case",
			args: [][]string{
				{"1"},{""},{""},{""},{""},{""},{""},{""},{""},{""},{""},{""},
				{"GTP"}, // 12
				{"0"}, // 13
				{"0"}, // 14
				{"0"}, // 15
				{"0"}, // 16
				{"0"}, // 17
				{"0"}, // 18
				{"0"}, // 19
				{"0"}, // 20
				{"0"}, // 21
				{"0"}, // 22
				{"0"}, // 23
				{"0"}, // 24
				{"0"}, // 25
			},
			want: "",
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			cmd := tshark.TsharkArgs{}
			res := cmd.SetAnnotation(v.args)

			if res != v.want {
				t.Errorf("The return value is not the expected value.\n %s", res)
			}
		})
	}
}