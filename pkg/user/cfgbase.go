package user

import (
	"os"
	"fmt"
	"errors"
	"strconv"
	"strings"
	"local.packages/cfg"
)

func (c *configBaseUser) SetArgs(params map[string]string, isbool map[string]bool) error {
	if len(os.Args) % 2 != 0 {
		return errors.New("The argument is missing.")
	}

	for i, v := range os.Args {
		if strings.Contains(v, "feature-timestamp") && os.Args[i+1] == "false" {
			c.Timestamp, _ = strconv.ParseBool(os.Args[i+1])
			break

		} else {
			c.Timestamp = isbool["timestamp"]
		}

		if strings.Contains(v, "feature-name-resolution") && os.Args[i+1] == "false" {
			c.NameResolution, _ = strconv.ParseBool(os.Args[i+1])
			break

		} else {
			c.NameResolution = isbool["nameResolution"]
		}
	}

	if params["java"] == "" {
		c.Java = "default"

	} else {
		c.Java = params["java"]
	}

	if params["wireshark"] == "" {
		c.Wireshark = "default"

	} else {
		c.Wireshark = params["wireshark"]
	}

	if params["plantuml"] == "" {
		c.Plantuml = "default"

	} else {
		c.Plantuml = params["plantuml"]
	}
	return nil
}

func (c configBaseUser) Writing() error {
	InitConfig := cfg.Config{
		Java:           c.Java,
		Wireshark:      c.Wireshark,
		Plantuml:       c.Plantuml,
		Timestamp:      c.Timestamp,
		NameResolution: c.NameResolution,
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

func (c *configBaseUser) Validate() error {
	if err := cfg.ValidationConfig(); err != nil {
		return err
	}
	return nil
}
