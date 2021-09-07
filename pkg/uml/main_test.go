package uml

import (
	"fmt"
	"testing"
	"local.packages/cfg"
	"local.packages/uml"
	"local.packages/util"
	"local.packages/tshark"
)

func TestCreateTemplate(t *testing.T) {
	type Tests struct {
		name string
		args string
		want error
	}

	tests := []Tests{
		{
			name: "Normal Case [No Title]",
			args: "",
			want: nil,
		},
		{
			name: "Normal Case [Title]",
			args: "test",
			want: nil,
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			if err := uml.CreateTemplate(v.args); err != v.want {
				t.Errorf("err: %s\n", err)
			}
		})
	}
}

func TestWriteUml(t *testing.T) {
	const (
		fileName string = "tmp.puml"
		fileDir  string = "/home/makeplantuml/container/pkg/uml/puml/"
	)

	type Tests struct {
		name      string
		args      tshark.TsharkHeaders
		timestamp bool
		want      error
	}

	tests := []Tests{
		{
			name: "Normal Case [No Timestamp]",
			args: tshark.TsharkHeaders{
				tshark.TsharkHeader{
					Number:   "1",
					Time:     "2021-06-16 10:34:38.377646",
					SrcAddr:  "172.16.10.10",
					SrcPort:  "46571",
					DstAddr:  "172.16.10.20",
					DstPort:  "38412",
					Protocol: "NGAP/NAS-5GS",
					Message:  "InitialUEMessage, Registration request",
					Checksum: "0x527eec87",
				},
			},
			timestamp: false,
			want: nil,
		},
		{
			name: "Normal Case [No Timestamp]",
			args: tshark.TsharkHeaders{
				tshark.TsharkHeader{
					Number:   "2",
					Time:     "2021-06-16 10:34:38.377646",
					SrcAddr:  "10.244.166.179",
					SrcPort:  "50718",
					DstAddr:  "10.103.54.119",
					DstPort:  "29510",
					Protocol: "HTTP2",
					Message:  "GET /nnrf-disc/v1/nf-instances",
					Checksum: "0x0000f32b",
				},
			},
			timestamp: true,
			want: nil,
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			if err := uml.WriteUml(v.args, v.timestamp); err != v.want {
				t.Errorf("err: %s\n", err)
			
			} else {
				fmt.Println(util.FileRead(fileDir + fileName))
			}
		})
	}

	if err := util.FileRemove(fileDir); err != nil {
		t.Errorf("err: %s\n", err)
	}

	errorTests := Tests{
		name: "Error Case",
		args: tshark.TsharkHeaders{
			tshark.TsharkHeader{
				Number:   "3",
				Time:     "2021-06-16 10:34:38.592685",
				SrcAddr:  "10.244.166.179",
				SrcPort:  "50718",
				DstAddr:  "10.244.166.129",
				DstPort:  "29510",
				Protocol: "HTTP2",
				Message:  "200 OK",
				Checksum: "0x000063c3",
			},
		},
		timestamp: false,
		want: nil,
	}

	t.Run(errorTests.name, func(t *testing.T) {
		if err := uml.WriteUml(errorTests.args, errorTests.timestamp); err == errorTests.want {
			t.Errorf("The expectation is that it will fail because the target file does not exist.\n")
		}
	})
}

func TestRenderingUml(t *testing.T) {
	const (
		fileName  string = "tmp.puml"
		fileDir   string = "/home/makeplantuml/container/pkg/uml/puml/"
		resultDir string = "/home/makeplantuml/container/pkg/uml/result/"
	)

	type Tests struct {
		name string
		args tshark.TsharkHeaders
		want error
	}

	tests := Tests{
		name: "Normal Case",
		args: tshark.TsharkHeaders{
			tshark.TsharkHeader{
				Number:   "1",
				Time:     "2021-06-16 10:34:38.377646",
				SrcAddr:  "172.16.10.10",
				SrcPort:  "46571",
				DstAddr:  "172.16.10.20",
				DstPort:  "38412",
				Protocol: "NGAP/NAS-5GS",
				Message:  "InitialUEMessage, Registration request",
				Checksum: "0x527eec87",
			},
		},
		want: nil,
	}

	t.Run(tests.name, func(t *testing.T) {
		preparationRendering(t, tests.args)
		fmt.Println(util.FileRead(fileDir + fileName))

		if err := uml.RenderingUml(); err != tests.want {
			t.Errorf("Failed to render the puml file.\nerr: %s\n", err)
		}
	})

	if err := util.FileRemove(fileDir); err != nil {
		t.Errorf("err: %s\n", err)
	}

	tests = Tests{
		name: "Error Case",
		args: tshark.TsharkHeaders{
			tshark.TsharkHeader{
				Number:   "1",
				Time:     "2021-06-16 10:34:38.377646",
				SrcAddr:  "172.16.10.10",
				SrcPort:  "46571",
				DstAddr:  "172.16.10.20",
				DstPort:  "38412",
				Protocol: "NGAP/NAS-5GS",
				Message:  "InitialUEMessage, Registration request",
				Checksum: "0x527eec87",
			},
		},
		want: nil,
	}

	t.Run(tests.name, func(t *testing.T) {
		if err := uml.RenderingUml(); err == tests.want {
			t.Errorf("It is assumed that an error will be output because the target file does not exist.\nerr: %s\n", err)
		}
	})

	if err := util.FileRemove(resultDir); err != nil {
		t.Errorf("err: %s\n", err)
	}
}

func preparationRendering(t *testing.T, args tshark.TsharkHeaders) {
	cfg.CfgVal.Profile.Path.Plantuml = "/home/makeplantuml/container/docs/plantuml.jar"

	if err := uml.CreateTemplate(""); err != nil {
		t.Errorf("Failed to create template.\nerr: %s\n", err)
	}

	if err := uml.WriteUml(args, false); err != nil {
		t.Errorf("Writing to the puml file failed.\nerr: %s\n", err)
	}

	util.PumlLocation.Path = "/home/makeplantuml/container/pkg/uml/puml"
}
