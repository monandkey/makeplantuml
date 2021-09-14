package tshark

import (
	"testing"
	"local.packages/tshark"
)

var inputN = `"1","2021-06-16 10:34:38.377646","172.16.10.10",,,,"46571","172.16.10.20",,,,"38412","NGAP/NAS-5GS","InitialUEMessage, Registration request",,"0x527eec87",,"124",`
var inputA = `"1","2021-06-16 10:34:38.377646","172.16.10.10",,,,"46571","172.16.10.20",,,,"38412","NGAP/NAS-5GS","InitialUEMessage, Registration request",,,,,,,,,,,,`

var expWantN = []map[string]string{
	{
		"number":   "1", 
		"time":     "2021-06-16 10:34:38.377646", 
		"srcAddr":  "172.16.10.10", 
		"srcPort":  "46571", 
		"dstAddr":  "172.16.10.20", 
		"dstPort":  "38412", 
		"protocol": "NGAP/NAS-5GS", 
		"message":  "InitialUEMessage, Registration request",
	},
}

var expWantA = []map[string]string{
	{
		"number":     "1", 
		"time":       "2021-06-16 10:34:38.377646", 
		"srcAddr":    "172.16.10.10", 
		"srcPort":    "46571", 
		"dstAddr":    "172.16.10.20", 
		"dstPort":    "38412", 
		"protocol":   "NGAP/NAS-5GS", 
		"message":    "InitialUEMessage, Registration request",
		"annotation": "",
	},
}

func TestParserOutput(t *testing.T) {
	type Tests struct {
		name string
		args string
		init string
		want []map[string]string
	}

	tests := []Tests{
		{
			name: "Normal Case",
			args: "normal",
			init: inputN,
			want: expWantN,
		},
		{
			name: "Annotation Case",
			args: "annotation",
			init: inputA,
			want: expWantA,
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			cmd := tshark.TsharkArgs{Out: []byte(v.init)}
			res := cmd.ParserOutput(v.args)
			if res[0]["number"] != v.want[0]["number"] {
				t.Errorf("The result is not the expected value.")
			}

			if res[0]["time"] != v.want[0]["time"] {
				t.Errorf("The result is not the expected value.")
			}

			if res[0]["srcAddr"] != v.want[0]["srcAddr"] {
				t.Errorf("The result is not the expected value.")
			}

			if res[0]["srcPort"] != v.want[0]["srcPort"] {
				t.Errorf("The result is not the expected value.")
			}

			if res[0]["dstAddr"] != v.want[0]["dstAddr"] {
				t.Errorf("The result is not the expected value.")
			}

			if res[0]["dstPort"] != v.want[0]["dstPort"] {
				t.Errorf("The result is not the expected value.")
			}

			if res[0]["protocol"] != v.want[0]["protocol"] {
				t.Errorf("The result is not the expected value.")
			}

			if res[0]["message"] != v.want[0]["message"] {
				t.Errorf("The result is not the expected value.")
			}

			if _, ok := v.want[0]["annotation"]; ok {
				if res[0]["annotation"] != v.want[0]["annotation"] {
					t.Errorf("The result is not the expected value.")
				}
			}
		})
	}
}

