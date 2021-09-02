package cmd

import (
	"os"
	"fmt"
	"strconv"
	"strings"
	"github.com/spf13/cobra"
	"local.packages/makeplantuml"
)

type config struct {
	java           string
	wireshark      string
	timestamp      bool
	nameResolution bool
}

func init() {
	initCmd := &cobra.Command{}
	initCmd.Use = "init"
	initCmd.Short = "Create the config file"

	initConfig := config{
		java:           "",
		wireshark:      "",
		timestamp:      false,
		nameResolution: false,
	}

	initCmd.Flags().StringVar(&initConfig.java, "java-path", initConfig.java, "")
	initCmd.Flags().StringVar(&initConfig.wireshark, "wireshark-path", initConfig.wireshark, "")
	initCmd.Flags().BoolVar(&initConfig.timestamp, "feature-timestamp", initConfig.timestamp, "")
	initCmd.Flags().BoolVar(&initConfig.nameResolution, "feature-name-resolution", initConfig.nameResolution, "")

	initCmd.RunE = func(cmd *cobra.Command, args []string) error {
		for i, v := range os.Args {
			if strings.Contains(v, "feature-timestamp") && os.Args[i+1] == "false" {
				initConfig.timestamp, _ = strconv.ParseBool(os.Args[i+1])
			}

			if strings.Contains(v, "feature-name-resolution") && os.Args[i+1] == "false" {
				initConfig.nameResolution, _ = strconv.ParseBool(os.Args[i+1])
			}
		}

		if initConfig.java == "" {
			initConfig.java = "default"
		}

		if initConfig.wireshark == "" {
			initConfig.wireshark = "default"
		}

		InitConfig := makeplantuml.Config{
			Java:           initConfig.java,
			Wireshark:      initConfig.wireshark,
			Timestamp:      initConfig.timestamp,
			NameResolution: initConfig.nameResolution,
		}

		if makeplantuml.ExistInitConfig() {
			makeplantuml.InitializeConfig(InitConfig)
			fmt.Println("Create config file")
			os.Exit(0)

		} else {
			var a string
			fmt.Printf("overwrite ? y or n: ")
			fmt.Scan(&a)

			if a == "y" {
				makeplantuml.InitializeConfig(InitConfig)
				fmt.Println("Overwrite!!")
			}
		}
		return nil
	}
	rootCmd.AddCommand(initCmd)
}