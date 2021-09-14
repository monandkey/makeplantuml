package uml

import (
	"fmt"
	"testing"
	"local.packages/cfg"
	"local.packages/uml"
	"local.packages/util"
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
			cmd := uml.UmlArgs{}
			if err := cmd.CreateTemplate(v.args); err != v.want {
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
		args      []map[string]string
		mode      string
		timestamp bool
		want      error
	}

	tests := []Tests{
		{
			name: "Normal Case [No Timestamp]",
			args: []map[string]string{
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
			},
			mode: "normal",
			timestamp: false,
			want: nil,
		},
		{
			name: "Normal Case [No Timestamp]",
			args: []map[string]string{
				{
					"number":   "2",
					"time":     "2021-06-16 10:34:38.377646",
					"srcAddr":  "10.244.166.179",
					"srcPort":  "50718",
					"dstAddr":  "10.103.54.119",
					"dstPort":  "29510",
					"protocol": "HTTP2",
					"message":  "GET /nnrf-disc/v1/nf-instances",
				},
			},
			mode: "normal",
			timestamp: true,
			want: nil,
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			cmd := uml.UmlArgs{}
			if err := cmd.WriteUml(v.mode, v.args, v.timestamp); err != v.want {
				t.Errorf("err: %s\n", err)
			
			} else {
				res, err := util.FileRead(fileDir + fileName)
				if err != nil {
					t.Errorf("Failed to load the file.")
				}
				fmt.Println(res)
			}
		})
	}

	if err := util.FileRemove(fileDir); err != nil {
		t.Errorf("err: %s\n", err)
	}

	errorTests := Tests{
		name: "Error Case",
		args: []map[string]string{
			{
				"number":   "3",
				"time":     "2021-06-16 10:34:38.592685",
				"srcAddr":  "10.244.166.179",
				"srcPort":  "50718",
				"dstAddr":  "10.244.166.129",
				"dstPort":  "29510",
				"protocol": "HTTP2",
				"message":  "200 OK",
			},
		},
		mode: "normal",
		timestamp: false,
		want: nil,
	}

	t.Run(errorTests.name, func(t *testing.T) {
		tmp := util.PumlLocation.Path
		util.PumlLocation.Path = "/rooot"

		cmd := uml.UmlArgs{}
		if err := cmd.WriteUml(errorTests.mode, errorTests.args, errorTests.timestamp); err == errorTests.want {
			t.Errorf("The expectation is that it will fail because the target file does not exist.\nerr: %s", err)
		}
		util.PumlLocation.Path = tmp
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
		args []map[string]string
		file string
		want error
	}

	tests := Tests{
		name: "Normal Case",
		args: []map[string]string{
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
		},
		file: "/home/makeplantuml/container/puml/tmp.puml",
		want: nil,
	}

	t.Run(tests.name, func(t *testing.T) {
		preparationRendering(t, tests.args)
		res, err := util.FileRead(fileDir + fileName)
		if err != nil {
			t.Errorf("Failed to load the file.")
		}
		fmt.Println(res)

		cmd := uml.UmlArgs{}
		if err := cmd.RenderingUml(tests.file); err != tests.want {
			t.Errorf("Failed to render the puml file.\nerr: %s\n", err)
		}
	})

	if err := util.FileRemove(fileDir); err != nil {
		t.Errorf("err: %s\n", err)
	}

	tests = Tests{
		name: "Error Case",
		args: []map[string]string{
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
		},
		file: "/home/makeplantuml/container/puml/tmp",
		want: nil,
	}

	t.Run(tests.name, func(t *testing.T) {
		cmd := uml.UmlArgs{}
		if err := cmd.RenderingUml(tests.file); err == tests.want {
			t.Errorf("It is assumed that an error will be output because the target file does not exist.\nerr: %s\n", err)
		}
	})

	if err := util.FileRemove(resultDir); err != nil {
		t.Errorf("err: %s\n", err)
	}
}

func preparationRendering(t *testing.T, args []map[string]string) {
	cfg.Param.Profile.Path.Plantuml = cfg.PlantumlLongPath

	cmd := uml.UmlArgs{}
	if err := cmd.CreateTemplate(""); err != nil {
		t.Errorf("Failed to create template.\nerr: %s\n", err)
	}

	if err := cmd.WriteUml("normal", args, false); err != nil {
		t.Errorf("Writing to the puml file failed.\nerr: %s\n", err)
	}

	util.PumlLocation.Path = "/home/makeplantuml/container/pkg/uml/puml"
}
