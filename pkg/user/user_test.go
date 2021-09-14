package user

import (
	"testing"
	"local.packages/cfg"
	"local.packages/util"
	"local.packages/user"
)

func TestBaseUser(t *testing.T) {
	var (
		fileName  = "/home/makeplantuml/container/sample/testattach.pcapng"
		hostFile  = "/home/makeplantuml/container/profile/hosts"
		title     = "Test"
		timestamp = false
		pumlFile  = "/home/makeplantuml/container/pkg/user/puml/tmp.puml"	
	)

	cfg.Param.Profile.Path.Plantuml = cfg.PlantumlLongPath

	type Tests struct {
		name string
		args int
		want error
	}

	tests := []Tests{
		{
			name: "Base User",
			args: 0,
			want: nil,	
		},
		{
			name: "Handson User",
			args: 1,
			want: nil,	
		},
		{
			name: "ToWriting User",
			args: 2,
			want: nil,	
		},
		{
			name: "FromRendering User",
			args: 3,
			want: nil,	
		},
		{
			name: "Annotatio User",
			args: 4,
			want: nil,	
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			var use user.UserMethod
			if v.args == 0 {
				use = user.UserSelection(user.Normal())
			}

			if v.args == 1 {
				use = user.UserSelection(user.Handon())
			}

			if v.args == 2 {
				use = user.UserSelection(user.ToWriting())
			}

			if v.args == 3 {
				use = user.UserSelection(user.FromRendering())
			}

			if v.args == 4 {
				use = user.UserSelection(user.Annotatio())
			}

			use.SetCmd()
			use.SetArgs(fileName)

			if err := use.RunE(); err != v.want {
				t.Errorf("There is an error in RunE.\nerr: %s", err)
			}

			use.Parse()

			if err := use.CreateE(title); err != v.want {
				t.Errorf("There is an error in CreateE.\nerr: %s", err)
			}

			if err := use.NameResE(hostFile); err != v.want {
				t.Errorf("There is an error in NameResE.\nerr: %s", err)
			}

			if err := use.WritingE(timestamp); err != v.want {
				t.Errorf("There is an error in WritingE.\nerr: %s", err)
			}

			if err := use.RenderingE(pumlFile); err != v.want {
				t.Errorf("There is an error in RenderingE.\nerr: %s", err)
			}	
		})
	}

	util.FileRemove("/home/makeplantuml/container/pkg/user/puml")
	util.FileRemove("/home/makeplantuml/container/pkg/user/result")
}
