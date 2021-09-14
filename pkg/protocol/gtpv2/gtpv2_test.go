package gtpv2

import (
	"testing"
	"local.packages/gtpv2"
)

func TestGetResultCode(t *testing.T) {
	type Tests struct{
		name string
		args int
		want string
	}

	tests := []Tests{
		{
			name: "Normal Case",
			args: 16,
			want: "Cause: Request accepted\\n",
		},
		{
			name: "Error Case",
			args: 0,
			want: "",
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			res := gtpv2.GetCause(v.args)
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
			res := gtpv2.GetDcnr(v.args)
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

func TestGetOiIndication(t *testing.T) {
	type Tests struct{
		name string
		args int
		want string
	}

	tests := []Tests{
		{
			name: "Normal Case",
			args: 1,
			want: "OI: supported\\n",
		},
		{
			name: "Error Case",
			args: 0,
			want: "",
		},
	}
	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			res := gtpv2.GetOiIndication(v.args)
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

func TestGetSiIndication(t *testing.T) {
	type Tests struct{
		name string
		args int
		want string
	}

	tests := []Tests{
		{
			name: "Normal Case",
			args: 1,
			want: "SI: supported\\n",
		},
		{
			name: "Error Case",
			args: 0,
			want: "",
		},
	}
	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			res := gtpv2.GetSiIndication(v.args)
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
