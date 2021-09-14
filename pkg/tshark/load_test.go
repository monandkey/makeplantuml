package tshark

import (
	"testing"
	"local.packages/util"
	"local.packages/tshark"
)

var inputN = []map[string]string{
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

func TestNameResolution(t *testing.T) {
	const fileName string = "/home/makeplantuml/container/sample/testattach.pcapng"

	type Tests struct {
		name   string
		target string
		args   tshark.TsharkArgs
		want   error
	}

	tests := Tests {
		name: "Normal Case",
		args: tshark.TsharkArgs{
			Header: inputN,
		},
		want: nil,
	}

	util.PumlLocation.Path = "/home/makeplantuml/container/puml"

	t.Run(tests.name, func(t *testing.T) {
		cmd := tshark.TsharkArgs{}
		res := cmd.NameResolution(tests.args.Header, "../../profile/hosts")

		if res != tests.want {
			t.Errorf("The return value is not expected.\nres: %s\n", res)
		}
	})
}

func TestNameOrNfSelection(t *testing.T) {
	type required struct {
		name string
		nf   string
	}

	type Tests struct {
		name string
		args required
		want string
	}

	tests := []Tests{
		{
			name: "w/ name, w/o nf",
			args: required{
				name: "enb",
				nf:   "",
			},
			want: "enb",
		},
		{
			name: "w/ name, w/ nf",
			args: required{
				name: "enb",
				nf:   "enb01",
			},
			want: "enb01",
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			res := tshark.NameOrNfSelection(v.args.name, v.args.nf)
			if res != v.want {
				t.Errorf("The return value is not expected.\nres: %s\n", res)
			}
		})
	}
}
