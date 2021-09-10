package cmd

import (
	"os"
	"fmt"
	"github.com/spf13/cobra"
	"local.packages/user"
)

type params struct {
	version   bool
	fileName  string
	timeStamp bool
	title     string
	handson   bool
}

var rootCmd = &cobra.Command{}

func Execute() {
	err := rootCmd.Execute()
    if err != nil {
        os.Exit(0)
    }
}

func init() {
	rootCmd.Use = "makeplantuml"
	rootCmd.Short = "PCAP to PlantUML to PNG converter"

	params := params{
		version:   false,
		fileName:  "",
		timeStamp: false,
		title:     "",
		handson:   false,
	}

	rootCmd.Flags().BoolVarP(&params.version, "version", "v", params.version, "Display version.")
	rootCmd.Flags().StringVarP(&params.fileName, "filename", "f", params.fileName, "Target file name.")
	rootCmd.Flags().BoolVarP(&params.timeStamp, "timestamp", "t", params.timeStamp, "Print a timestamp")
	rootCmd.Flags().StringVar(&params.title, "puml-title", params.title, "Give PUML a title.")
	rootCmd.Flags().BoolVar(&params.handson, "handson-environment", params.handson, "For captures acquired in hands-on environment.")

	rootCmd.RunE = func(cmd *cobra.Command, args []string) error {
		if params.version {
			fmt.Println("version: 2.0.0")
			return nil
		}

		if params.fileName == "" {
			return rootCmd.Help()
		}

		var use user.UserMethod
		if params.handson {
			use = user.UseTsharkSelection(user.Handon())

		} else {
			use = user.UseTsharkSelection(user.Normal())
		}
		use.SetCmd()
		use.SetArgs(params.fileName)

		if err := use.RunE(); err != nil {
			return err
		}

		use.Parse()
		
		if err := use.NameResE("./profile/hosts"); err != nil {
			return err
		}

		if err := use.CreateE(params.title); err != nil {
			return err
		}

		if err := use.WritingE(params.timeStamp); err != nil {
			return err
		}

		if err := use.RenderingE(); err != nil {
			return err
		}
		return nil
	}
}
