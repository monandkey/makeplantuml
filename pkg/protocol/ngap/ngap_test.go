package ngap

import (
	"testing"
	"local.packages/ngap"
)

func TestGetCause(t *testing.T) {
	type Tests struct{
		name string
		args int
		want string
	}

	tests := []Tests{
		{
			name: "Normal Case",
			args: 3,
			want: "Cause: Illegal UE\\n",
		},
		{
			name: "Error Case",
			args: 0,
			want: "",
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			res := ngap.GetCause(v.args)
			switch(v.name) {
				case "Normal Case":
					if res != v.want {
						t.Errorf("Failed to load the file.\n%s", res)
					}

				case "Error Case":
					if res != v.want {
						t.Errorf("Error handling is incorrect.\nres: %s want: %s", res, v.want)
					}		
			}
		})
	}
}

func TestGetTypeOfId(t *testing.T) {
	type Tests struct{
		name string
		args int
		want string
	}

	tests := []Tests{
		{
			name: "Normal Case",
			args: 1,
			want: "Attach type: SUCI\\n",
		},
		{
			name: "Error Case",
			args: 0,
			want: "",
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			res := ngap.GetTypeOfId(v.args)
			switch(v.name) {
				case "Normal Case":
					if res != v.want {
						t.Errorf("Failed to load the file.\n%s", res)
					}

				case "Error Case":
					if res != v.want {
						t.Errorf("Error handling is incorrect.\nres: %s want: %s", res, v.want)
					}		
			}
		})
	}
}
