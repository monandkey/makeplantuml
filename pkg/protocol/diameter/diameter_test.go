package diameter

import (
	"testing"
	"local.packages/diameter"
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
			args: 2001,
			want: "Cause: DIAMETER_SUCCESS\\n",
		},
		{
			name: "Error Case",
			args: 0,
			want: "",
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			res := diameter.GetResultCode(v.args)
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