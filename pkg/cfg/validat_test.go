package cfg

import (
	"testing"
	"local.packages/cfg"
)

func TestValidationConfig(t *testing.T) {
	type application struct {
		java      string
		wireshark string
		plantuml  string
	}

	type Tests struct {
		name string
		args application
		want error
	}

	tests := []Tests{
		{
			name: "Normal Case",
			args: application{
				java:      "default",
				wireshark: "default",
				plantuml:  cfg.PlantumlLongPath,
			},
			want: nil,
		},
		{
			name: "Error Case w/o Java",
			args: application{
				java:      "",
				wireshark: "default",
				plantuml:  cfg.PlantumlLongPath,
			},
			want: nil,
		},
		{
			name: "Error Case w/o Wireshark",
			args: application{
				java:      "default",
				wireshark: "",
				plantuml:  cfg.PlantumlLongPath,
			},
			want: nil,
		},
		{
			name: "Error Case w/o PlantUML",
			args: application{
				java:      "default",
				wireshark: "default",
				plantuml:  "",
			},
			want: nil,
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			cfg.Param.Profile.Path.Java = v.args.java
			cfg.Param.Profile.Path.Wireshark = v.args.wireshark
			cfg.Param.Profile.Path.Plantuml = v.args.plantuml

			switch(v.name) {
				case "Normal Case":
					if err := cfg.ValidationConfig(); err != nil {
						t.Errorf("%s", err)
					}
				default:
					if err := cfg.ValidationConfig(); err == nil {
						t.Errorf("%s", err)
					}				
			}
		})
	}
}