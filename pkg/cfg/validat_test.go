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
		want string
	}

	tests := []Tests{
		{
			name: "Normal Case",
			args: application{
				java:      "default",
				wireshark: "default",
				plantuml:  cfg.PlantumlLongPath,
			},
			want: "",
		},
		{
			name: "w/o Java",
			args: application{
				java:      "",
				wireshark: "default",
				plantuml:  cfg.PlantumlLongPath,
			},
			want: "",
		},
		{
			name: "w/o Wireshark",
			args: application{
				java:      "default",
				wireshark: "",
				plantuml:  cfg.PlantumlLongPath,
			},
			want: "",
		},
		{
			name: "w/o PlantUML",
			args: application{
				java:      "default",
				wireshark: "default",
				plantuml:  "",
			},
			want: "",
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			cfg.Param.Profile.Path.Java = v.args.java
			cfg.Param.Profile.Path.Wireshark = v.args.wireshark
			cfg.Param.Profile.Path.Plantuml = v.args.plantuml
			cfg.ValidationConfig()
		})
	}
}