package tshark

import (
	"testing"
	"local.packages/tshark"
)


func TestRunTshark(t *testing.T) {
	type Tests struct {
		name string
		args string
		want int
	}

	tests := Tests {
		name: "Normal Case",
		args: "/home/makeplantuml/container/sample/testattach.pcapng",
		want: 17,
	}

	t.Run(tests.name, func(t *testing.T) {
		got := tshark.RunTshark(tests.args)
		if len(got) == 0 {
			t.Errorf("The return value is zero.\nReturn Value: %d", len(got))

		} else if (len(got) % tests.want) == 0 {
			t.Errorf("The return value is not the expected value.\nValue: %d", len(got))
		}
	})

	tests = Tests {
			name: "Error Case",
			args: "sample/testattach.pcapng",
			want: 0,
	}

	t.Run(tests.name, func(t *testing.T) {
		got := tshark.RunTshark(tests.args)
		if len(got) != tests.want {
			t.Errorf("The expected value is zero.")
		}
	})
}

