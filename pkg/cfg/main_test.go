package cfg

import (
	"fmt"
	"reflect"
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

func TestInitializeConfig(t *testing.T) {
	type Tests struct {
		name string
		args cfg.Config
		want error
	}

	tests := Tests{
		name: "Normal Case",
		args: cfg.Config{
			Java:           "default",
			Wireshark:      "default",
			Plantuml:       "default",
			Timestamp:      false,
			NameResolution: false,
		},
		want: nil,
	}

	t.Run(tests.name, func(t *testing.T) {
		if err := cfg.InitializeConfig(tests.args); err != tests.want {
			t.Errorf("err: %s\n", err)
		}
	})
}

func TestConfigLoad(t *testing.T) {
	type Config struct {
		Java           string
		Wireshark      string
		Plantuml       string
		Timestamp      string
		NameResolution string
	}

	type Tests struct {
		name string
		args Config
		want string
	}

	tests := []Tests{
		{
			name: "Normal Case",
			args: Config{
				Java:           "default",
				Wireshark:      "default",
				Plantuml:       "default",
				Timestamp:      "false",
				NameResolution: "false",
			},
			want: "",
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			if reflect.TypeOf(cfg.CfgVal.Profile.Path.Java) != reflect.TypeOf(v.args.Java) {
				t.Errorf("The type string is the expected value.")
			}

			if reflect.TypeOf(cfg.CfgVal.Profile.Path.Wireshark) != reflect.TypeOf(v.args.Wireshark) {
				t.Errorf("The type string is the expected value.")
			}

			if reflect.TypeOf(cfg.CfgVal.Profile.Path.Plantuml) != reflect.TypeOf(v.args.Plantuml) {
				t.Errorf("The type string is the expected value.")
			}

			if reflect.TypeOf(cfg.CfgVal.Profile.Feature.Timestamp) != reflect.TypeOf(v.args.Timestamp) {
				t.Errorf("The type bool is the expected value.")
			}

			if reflect.TypeOf(cfg.CfgVal.Profile.Feature.NameResolution) != reflect.TypeOf(v.args.NameResolution) {
				t.Errorf("The type bool is the expected value.")
			}
		})
	}
}
