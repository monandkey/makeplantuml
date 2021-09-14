package uml

import (
	"os"
	"fmt"
	"local.packages/util"
)

func (u UmlArgs) WriteUml(kinds string, headers []map[string]string, tf bool) error {
	switch(kinds) {
		case "normal":
			if err := u.WriteUmlNormal(headers, tf); err != nil {
				return err
			}
		
		case "annotation":
			if err := u.WriteUmlAnnotatio(headers, tf); err != nil {
				return err
			}
	}
	return nil
}

func (u UmlArgs) WriteUmlNormal(headers []map[string]string, tf bool) error {
	file, err := os.OpenFile(util.PumlLocation.Path + "/tmp.puml", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	for i, header := range headers {
		if i == 0 {
			fmt.Fprintln(file, "")
		}

		str := "\"" + header["srcAddr"] + "\" -> \"" + header["dstAddr"] + "\" : ["+ header["number"] + "][" + header["protocol"] + "] " + header["message"]
		fmt.Fprintln(file, str)

		if tf {
			rnote := "rnote left: " + header["time"]
			fmt.Fprintln(file, rnote)
		}
	}
	fmt.Fprintln(file, "\n@enduml")
	return nil
}

func (u UmlArgs) WriteUmlAnnotatio(headers []map[string]string, tf bool) error {
	file, err := os.OpenFile(util.PumlLocation.Path + "/tmp.puml", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	for i, header := range headers {
		if i == 0 {
			fmt.Fprintln(file, "")
		}

		if header["annotation"] != "" {
			str := "\"" + header["srcAddr"] + "\" -> \"" + header["dstAddr"] + "\" : ["+ header["number"] + "][" + header["protocol"] + "] " + header["message"] + "\\n" + header["annotation"]
			fmt.Fprintln(file, str)

		} else {
			str := "\"" + header["srcAddr"] + "\" -> \"" + header["dstAddr"] + "\" : ["+ header["number"] + "][" + header["protocol"] + "] " + header["message"]
			fmt.Fprintln(file, str)
		} 

		// if header["annotation"] != "" {
		// 	file.Write(([]byte)("note right of " + header["srcAddr"] + "\n"))
		// 	file.Write(([]byte)(header["annotation"] + "\n"))
		// 	file.Write(([]byte)("end note\n"))
		// }

		if tf {
			rnote := "rnote left: " + header["time"]
			fmt.Fprintln(file, rnote)
		}
	}
	fmt.Fprintln(file, "\n@enduml")
	return nil
}
