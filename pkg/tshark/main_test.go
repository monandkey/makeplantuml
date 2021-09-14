package tshark

import (
	"testing"
	"local.packages/cfg"
	"local.packages/tshark"
)


func TestGetTsharkCommand(t *testing.T) {
	type Tests struct {
		name string
		want string
	}

	tests := Tests {
		name: "Normal Case",
		want: "tshark",
	}

	t.Run(tests.name, func(t *testing.T) {
		cmd := tshark.TsharkArgs{}
		res := cmd.GetTsharkCommand()
		if res != tests.want {
			t.Errorf("The result is not the expected value.")
		}
	})

	tests = Tests {
			name: "Not Default",
			want: "/usr/bin/tshark\\tshark.exe",
	}

	cfg.Param.Profile.Path.Wireshark = "/usr/bin/tshark"

	t.Run(tests.name, func(t *testing.T) {
		cmd := tshark.TsharkArgs{}
		res := cmd.GetTsharkCommand()
		if res != tests.want {
			t.Errorf("The result is not the expected value.")
		}
	})
}

func TestConvertOutputResultIntoArray(t *testing.T) {
	const out string = `"1","2021-06-16 10:34:38.377646","172.16.10.10",,,,"46571","172.16.10.20",,,,"38412","NGAP/NAS-5GS","InitialUEMessage, Registration request",,"0x527eec87",,"124",`
	const exe string = `"1","2021-06-16 10:34:38.377646","172.16.10.10","","","","46571","172.16.10.20","","","","38412","NGAP/NAS-5GS","InitialUEMessage, Registration request","","0x527eec87","","124",`

	type Tests struct {
		name string
		args string
		want string
	}

	tests := Tests{
		name: "Normal Case",
		args: out,
		want: exe,
	}

	t.Run(tests.name, func(t *testing.T) {
		cmd := tshark.TsharkArgs{}
		res := cmd.ConvertOutputResultIntoArray(tests.args)
		if res[0] != tests.want {
			t.Errorf("The replacement result is not the expected value..")
		}
	})
}

func TestCleanLineStrings(t *testing.T) {
	const out string = `"1","2021-06-16 10:34:38.377646","172.16.10.10","","","","46571","172.16.10.20","","","","38412","NGAP/NAS-5GS","InitialUEMessage, Registration request","","0x527eec87","","124",`
	const exe string = `"1","2021-06-16 10:34:38.377646","172.16.10.10","","","","46571","172.16.10.20","","","","38412","NGAP/NAS-5GS","InitialUEMessage, Registration request","","0x527eec87","","124",""`

	type Tests struct {
		name string
		args string
		want string
	}

	tests := Tests{
		name: "Normal Case",
		args: out,
		want: exe,
	}

	t.Run(tests.name, func(t *testing.T) {
		cmd := tshark.TsharkArgs{}
		res := cmd.CleanLineStrings(tests.args)
		if res != tests.want {
			t.Errorf("The replacement result is not the expected value.")
		}
	})
}
