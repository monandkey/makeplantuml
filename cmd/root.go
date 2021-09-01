package cmd

import (
	"os"
	"fmt"
	"github.com/spf13/cobra"
	"local.packages/makeplantuml"
)

type params struct {
	version            bool
	fileName           string
	timeStamp          bool
	pumlConvert        bool
	packetOptimization bool
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
		version:            false,
		fileName:           "",
		timeStamp:          false,
		pumlConvert:        false,
		packetOptimization: false,
	}

	rootCmd.Flags().BoolVarP(&params.version, "version", "v", params.version, "display version")
	rootCmd.Flags().StringVarP(&params.fileName, "filename", "f", params.fileName, "")
	rootCmd.Flags().BoolVarP(&params.timeStamp, "timestamp", "t", params.timeStamp, "")
	rootCmd.Flags().BoolVarP(&params.pumlConvert, "puml-convert", "p", params.pumlConvert, "")
	rootCmd.Flags().BoolVar(&params.packetOptimization, "packet-optimization",     params.packetOptimization, "")

	rootCmd.RunE = func(cmd *cobra.Command, args []string) error {
		if params.version {
			fmt.Println("version: 2.0.0")
			os.Exit(0)
		}

		if len(args) > 5 {
			return rootCmd.Help()
		}

		fmt.Println(makeplantuml.RunTshark())
		return nil
	}
}
