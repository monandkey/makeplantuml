package cmd

import (
	"os"
	"fmt"
	"github.com/spf13/cobra"
	"local.packages/makeplantuml"
)

type params struct {
	version   bool
	fileName  string
	timeStamp bool
}

var rootCmd = &cobra.Command{}

func Execute() {
	err := rootCmd.Execute()
    if err != nil {
        os.Exit(0)
    }
}

func init() {
	rootCmd.Use = "pppc"
	rootCmd.Short = "PCAP to PlantUML to PNG converter"

	params := params{
		version:   false,
		fileName:  "",
		timeStamp: false,
	}

	rootCmd.Flags().BoolVarP(&params.version, "version", "v", params.version, "display version")
	rootCmd.Flags().StringVarP(&params.fileName, "filename", "f", params.fileName, "")
	rootCmd.Flags().BoolVarP(&params.timeStamp, "timestamp", "t", params.timeStamp, "")

	rootCmd.RunE = func(cmd *cobra.Command, args []string) error {
		if params.version {
			fmt.Println("version: 2.0.0")
			os.Exit(0)
		}

		if params.fileName == "" {
			return rootCmd.Help()
		}

		t := makeplantuml.RunTshark(params.fileName)
		makeplantuml.CreateTemplate()
		makeplantuml.NameResolution(t)
		makeplantuml.WriteUml(t)
		makeplantuml.RenderingUml()
		return nil
	}
}
