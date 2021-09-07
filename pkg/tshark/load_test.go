package tshark

import (
	"testing"
	"local.packages/tshark"
)

func TestNameResolution(t *testing.T) {
	fileName := "/home/makeplantuml/container/sample/testattach.pcapng"
	type Tests struct {
		name   string
		target string
		args   tshark.TsharkHeaders
		want   string
	}

	tests := Tests {
		name: "Normal Case",
		args: tshark.RunTshark(fileName),
		want: "",
	}

	t.Run(tests.name, func(t *testing.T) {
		tshark.NameResolution(tests.args, "../../profile/hosts")
		for _, v := range tests.args {
			if v.Number == tests.want {
				t.Errorf("Number: This field is empty.")
			}
			if v.Time == tests.want {
				t.Errorf("Time: This field is empty.")
			}
			if v.SrcAddr == tests.want {
				t.Errorf("SrcAddr: This field is empty.")
			}
			if v.SrcPort == tests.want {
				t.Errorf("SrcPort: This field is empty.")
			}
			if v.DstAddr == tests.want {
				t.Errorf("DstAddr: This field is empty.")
			}
			if v.DstPort == tests.want {
				t.Errorf("DstPort: This field is empty.")
			}
			if v.Protocol == tests.want {
				t.Errorf("Protocol: This field is empty.")
			}
			if v.Message == tests.want {
				t.Errorf("Message: This field is empty.")
			}
			if v.Checksum == tests.want {
				t.Errorf("Checksum: This field is empty.")
			}
		}
	})
}
