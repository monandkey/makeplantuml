package cmd

import (
	"github.com/spf13/cobra"
	"local.packages/user"
)

type config struct {
	java           string
	wireshark      string
	plantuml       string
	timestamp      bool
	nameResolution bool
	validation     bool
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
		validation:     false,
	}

	initCmd.Flags().StringVar(&initConfig.java, "java-path", initConfig.java, "Specify the location of java")
	initCmd.Flags().StringVar(&initConfig.wireshark, "wireshark-path", initConfig.wireshark, "Specify the location of Wireshark")
	initCmd.Flags().StringVar(&initConfig.plantuml, "plantuml-path", initConfig.plantuml, "Specify the location of PlantUML")
	initCmd.Flags().BoolVar(&initConfig.timestamp, "feature-timestamp", initConfig.timestamp, "Always add a timestamp")
	initCmd.Flags().BoolVar(&initConfig.nameResolution, "feature-name-resolution", initConfig.nameResolution, "Always resolve names")
	initCmd.Flags().BoolVar(&initConfig.validation, "validation-config", initConfig.validation, "Verify that the configuration settings are correct")

	initCmd.RunE = func(cmd *cobra.Command, args []string) error {
		var use user.ConfigUserMethod
		use = user.ConfigUserSelection()

		if initConfig.validation {
			use.Validate()
			return nil
		}

		str := map[string]string{
			"java":           initConfig.java,
			"wireshark":      initConfig.wireshark,
			"plantuml":       initConfig.plantuml,
		}

		isbool := map[string]bool{
			"timestamp":      initConfig.timestamp,
			"nameResolution": initConfig.nameResolution,
		}

		if err := use.SetArgs(str, isbool); err != nil {
			return err
		}

		if err := use.Writing(); err != nil {
			return err
		}
		return nil
	}
	rootCmd.AddCommand(initCmd)
}