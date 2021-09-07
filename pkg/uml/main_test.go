package uml

import (
	"fmt"
	"testing"
	"local.packages/uml"
	"local.packages/util"
)

func TestCreateTemplate(t *testing.T) {
	const (
		fileName string = "tmp.puml"
		fileDir string = "/home/makeplantuml/container/pkg/uml/puml/"
	)

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
				t.Errorf("Test")
			
			} else {
				fmt.Println(util.FileRead(fileDir + fileName))
				if err := util.FileRemove(fileDir); err != nil {
					fmt.Println(err)
				}
			}
		})
	}
}
