package s1ap

import (
	"testing"
	"local.packages/s1ap"
)

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
			want: "Attach type: IMSI\\n",
		},
		{
			name: "Error Case",
			args: 0,
			want: "",
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			res := s1ap.GetTypeOfId(v.args)
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

func TestGetDcnr(t *testing.T) {
	type Tests struct{
		name string
		args int
		want string
	}

	tests := []Tests{
		{
			name: "Normal Case",
			args: 1,
			want: "DCNR: supported\\n",
		},
		{
			name: "Error Case",
			args: 0,
			want: "",
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			res := s1ap.GetDcnr(v.args)
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
