package cmd

import (
	"os"
	"fmt"
	"errors"
	"github.com/spf13/cobra"
	"local.packages/tshark"
	"local.packages/uml"
)

type params struct {
	version   bool
	fileName  string
	timeStamp bool
	title     string
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
		title:     "",
	}

	rootCmd.Flags().BoolVarP(&params.version, "version", "v", params.version, "display version")
	rootCmd.Flags().StringVarP(&params.fileName, "filename", "f", params.fileName, "")
	rootCmd.Flags().BoolVarP(&params.timeStamp, "timestamp", "t", params.timeStamp, "")
	rootCmd.Flags().StringVar(&params.title, "puml-title", params.title, "Give PUML a title.")

	rootCmd.RunE = func(cmd *cobra.Command, args []string) error {
		if params.version {
			fmt.Println("version: 2.0.0")
			os.Exit(0)
		}

		if params.fileName == "" {
			return rootCmd.Help()
		}

		t := tshark.RunTshark(params.fileName)
		if len(t) == 0 {
			return errors.New("The result of tshark execution is not the expected value.")
		}

		uml.CreateTemplate(params.title)
		tshark.NameResolution(t)
		uml.WriteUml(t, params.timeStamp)
		uml.RenderingUml()
		return nil
	}
}
