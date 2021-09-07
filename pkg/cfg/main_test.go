package cfg

import (
	"fmt"
	"testing"
	"os/exec"
	"local.packages/cfg"
)

func TestExistInitConfig(t *testing.T) {
	type Tests struct {
		name string
		want bool
	}

	tests := []Tests{
		{
			name: "Normal Case",
			want: true,
		},
		{
			name: "Error Case",
			want: false,
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			if v.name == "Error Case" {
				cmd, err := exec.Command(
					"mv",
					"/root/.makeplantuml.yml",
					"/root/.makeplantuml.yaml",
					).CombinedOutput()

				if err != nil {
					fmt.Printf("cmd: %s", cmd)
					fmt.Printf("err: %s\n", err)
				}
			}

			res := cfg.ExistInitConfig()
			if res == v.want {
				t.Errorf("The result is not the expected behavior.")
			}

			if v.name == "Error Case" {
				cmd, err := exec.Command(
					"mv",
					"/root/.makeplantuml.yaml",
					"/root/.makeplantuml.yml",
					).CombinedOutput()

				if err != nil {
					fmt.Printf("cmd: %s", cmd)
					fmt.Printf("err: %s\n", err)
				}
			}
		})
	}
}
