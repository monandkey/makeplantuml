package util

import (
	"os"
	"testing"
	"local.packages/util"
)

func TestFileRead(t *testing.T) {
	const fileName string = "/home/makeplantuml/container/pkg/util/go.mod"

	type Tests struct{
		name string
		args string
		want error
	}

	tests := []Tests{
		{
			name: "Normal Case",
			args: fileName,
			want: nil,
		},
		{
			name: "Error Case",
			args: "./dummy.go",
			want: nil,
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			_, err := util.FileRead(v.args)
			switch(v.name) {
				case "Normal Case":
					if err != v.want {
						t.Errorf("Failed to load the file.")
					}

				case "Error Case":
					if err == v.want {
						t.Errorf("Error handling is incorrect.")
					}		
			}
		})	
	}
}

func TestFileRemove(t *testing.T) {
	const fileName string = "/home/makeplantuml/container/pkg/util/dummy.test"

	type Tests struct{
		name string
		args string
		want error
	}

	tests := Tests{
		name: "Normal Case",
		args: fileName,
		want: nil,
	}

	if err := fileCreate(fileName); err != nil {
		t.Errorf("Failed to create dummy file.")
	}

	t.Run(tests.name, func(t *testing.T) {
		if err := util.FileRemove(fileName); err != tests.want {
			t.Errorf("Failed to delete the file.")
		}
	})

	tests = Tests{
		name: "Error Case",
		args: fileName,
		want: nil,
	}

	t.Run(tests.name, func(t *testing.T) {
		if err := util.FileRemove(fileName); err == tests.want {
			t.Errorf("It is assumed that an error will be output because the target file does not exist.\n err: %s\n", err)
		}
	})
}

func fileCreate(fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	file.Write(([]byte)("dummy\n"))
	return nil
}