package cmd

import (
	"os"
	"fmt"
	"strconv"
	"strings"
	"github.com/spf13/cobra"
	"local.packages/cfg"
)

type config struct {
	java           string
	wireshark      string
	plantuml       string
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
		plantuml:       "",
		timestamp:      false,
		nameResolution: false,
	}

	initCmd.Flags().StringVar(&initConfig.java, "java-path", initConfig.java, "Specify the location of java")
	initCmd.Flags().StringVar(&initConfig.wireshark, "wireshark-path", initConfig.wireshark, "Specify the location of Wireshark")
	initCmd.Flags().StringVar(&initConfig.plantuml, "plantuml-path", initConfig.plantuml, "Specify the location of PlantUML")
	initCmd.Flags().BoolVar(&initConfig.timestamp, "feature-timestamp", initConfig.timestamp, "Always add a timestamp")
	initCmd.Flags().BoolVar(&initConfig.nameResolution, "feature-name-resolution", initConfig.nameResolution, "Always resolve names")

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

		if initConfig.plantuml == "" {
			initConfig.plantuml = "default"
		}

		InitConfig := cfg.Config{
			Java:           initConfig.java,
			Wireshark:      initConfig.wireshark,
			Plantuml:       initConfig.plantuml,
			Timestamp:      initConfig.timestamp,
			NameResolution: initConfig.nameResolution,
		}

		if cfg.ExistInitConfig() {
			if err := cfg.InitializeConfig(InitConfig); err != nil {
				return err
			}

			fmt.Println("Create config file")

		} else {
			var a string
			fmt.Printf("overwrite ? y or n: ")
			fmt.Scan(&a)

			if a == "y" {
				if err := cfg.InitializeConfig(InitConfig); err != nil {
					return err
				}
				fmt.Println("Overwrite!!")
			}
		}
		return nil
	}
	rootCmd.AddCommand(initCmd)
}