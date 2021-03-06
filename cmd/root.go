package cmd

import (
	"os"
	"github.com/spf13/cobra"
	"local.packages/user"
)

type params struct {
	version       bool
	fileName      string
	timeStamp     bool
	title         string
	handson       bool
	toWriting     bool
	fromRendering bool
	pumlFile      string
	annotation    bool
}

var rootCmd = &cobra.Command{}

const rootExample = `makeplantuml --version | --help |
             --filename [filename] | --puml-title [title] | --timestamp | [--handson-environment] | [--inserting-annotations] |
             --filename [filename] --create-puml | --puml-title [title] | --timestamp | [--handson-environment] | [--inserting-annotations] |
             --rendering-puml | [--pumlfile [pumlfile]]`

func Execute() {
	err := rootCmd.Execute()
    if err != nil {
        os.Exit(0)
    }
}

func init() {
	rootCmd.Use = "makeplantuml"
	rootCmd.Short = "PCAP to PlantUML to PNG converter"
	rootCmd.Example = rootExample
	rootCmd.Version = "2.0.0"

	params := params{
		version:       false,
		fileName:      "",
		timeStamp:     false,
		title:         "",
		handson:       false,
		toWriting:     false,
		fromRendering: false,
		pumlFile:      "./puml/tmp.puml",
		annotation:     false,
	}

	rootCmd.Flags().BoolVarP(&params.version, "version", "v", params.version, "Display version.")
	rootCmd.Flags().StringVarP(&params.fileName, "filename", "f", params.fileName, "Target PCAP file name.")
	rootCmd.Flags().BoolVarP(&params.timeStamp, "timestamp", "t", params.timeStamp, "Print a timestamp")
	rootCmd.Flags().StringVar(&params.title, "puml-title", params.title, "Give PUML a title.")
	rootCmd.Flags().BoolVar(&params.handson, "handson-environment", params.handson, "For captures acquired in hands-on environment.")
	rootCmd.Flags().BoolVar(&params.toWriting, "create-puml", params.toWriting, "Creating a PUML file.")
	rootCmd.Flags().BoolVar(&params.fromRendering, "rendering-puml", params.fromRendering, "Renders the specified PUML file.")
	rootCmd.Flags().StringVar(&params.pumlFile, "puml-file", params.pumlFile, "Specify the rendered PUML file.")
	rootCmd.Flags().BoolVar(&params.annotation, "inserting-annotations", params.annotation, "Adding annotations to messages.")

	rootCmd.RunE = func(cmd *cobra.Command, args []string) error {
		if params.fileName == "" && !(params.fromRendering) {
			return rootCmd.Help()
		}

		var use user.UserMethod
		if params.handson {
			use = user.UserSelection(user.Handon())
		
		} else if params.toWriting {
			use = user.UserSelection(user.ToWriting())
		
		} else if params.fromRendering {
			use = user.UserSelection(user.FromRendering())

		} else if params.annotation {
			use = user.UserSelection(user.Annotatio())

		} else {
			use = user.UserSelection(user.Normal())
		}

		use.SetCmd()
		use.SetArgs(params.fileName)

		if err := use.RunE(); err != nil {
			return err
		}

		use.Parse()

		if err := use.CreateE(params.title); err != nil {
			return err
		}

		if err := use.NameResE("./profile/hosts"); err != nil {
			return err
		}

		if err := use.WritingE(params.timeStamp); err != nil {
			return err
		}

		if err := use.RenderingE(params.pumlFile); err != nil {
			return err
		}
		return nil
	}
}
