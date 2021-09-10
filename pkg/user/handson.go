package user

import (
	"fmt"
	"regexp"
	"os/exec"
	"local.packages/uml"
	"local.packages/tshark"
)

func (h handsonUser) new() UserMethod {
	return &handsonUser{}
}

func (h *handsonUser) SetCmd() {
	h.Cmd = tshark.GetTsharkCommand()
}

func (h *handsonUser) SetArgs(fileName string) {
	h.Args = []string{
		"-r", fileName,
		"-t", "ad",
		"-T", "fields",
		"-E", "separator=,",
		"-E", "quote=d",
		"-e", "frame.number",
		"-e", "_ws.col.Time",
		"-e", "ip.src",
		"-e", "ipv6.src",
		"-e", "udp.srcport",
		"-e", "tcp.srcport",
		"-e", "sctp.srcport",
		"-e", "ip.dst",
		"-e", "ipv6.dst",
		"-e", "udp.dstport",
		"-e", "tcp.dstport",
		"-e", "sctp.dstport",
		"-e", "_ws.col.Protocol",
		"-e", "_ws.col.Info",
		"-e", "ip.len",
		"-e", "ipv6.plen",
		"-d", "tcp.port==29000-30000,http2",
	}
}

func (h *handsonUser) RunE() error {
	var err error
	h.Out, err = exec.Command(h.Cmd, h.Args...).Output()

	if err != nil {
		return err
	}

	return nil
}

func (h *handsonUser) Parse() {
	lines := tshark.ConvertOutputResultIntoArray(string(h.Out))

	for _, line := range lines {
		line = tshark.CleanLineStrings(line)

		column := regexp.MustCompile("\".*?\"").FindAllStringSubmatch(line, -1)
		if len(column) == 0 {
			continue
		}

		tmpNumber := regexp.MustCompile("\"").ReplaceAllString(column[0][0], "")
		tmpTime := regexp.MustCompile("\"").ReplaceAllString(column[1][0], "")
		tmpSrcAddr := tshark.SetAddress(column[2][0], column[3][0], column[14][0], column[15][0])
		tmpSrcPort := tshark.SetPortAndCheckSum(column[4][0], column[5][0], column[6][0])
		tmpDstAddr := tshark.SetAddress(column[7][0], column[8][0], column[14][0], column[15][0])
		tmpDstPort := tshark.SetPortAndCheckSum(column[9][0], column[10][0], column[11][0])
		tmpProtocol := regexp.MustCompile("\"").ReplaceAllString(column[12][0], "")
		tmpMessage := tshark.SetMessage(column[13][0], tmpProtocol)

		tmph := map[string]string{
			"number":   tmpNumber,
			"time":     tmpTime,
			"srcAddr":  tmpSrcAddr,
			"srcPort":  tmpSrcPort,
			"dstAddr":  tmpDstAddr,
			"dstPort":  tmpDstPort,
			"protocol": tmpProtocol,
			"message":  tmpMessage,
		}
		h.Header = append(h.Header, []map[string]string{tmph}...)
	}
}

func (n *handsonUser) NameResE(fileName string) error {
	if err := tshark.NameResolution(n.Header, fileName); err != nil {
		return err
	}
	return nil
}

func (n *handsonUser) Display() {
	fmt.Println(n.Header)
}

func (n *handsonUser) CreateE(title string) error {
	if err := uml.CreateTemplate(title); err != nil {
		return err
	}
	return nil
}

func (n *handsonUser) WritingE(timestamp bool) error {
	if err := uml.WriteUml(n.Header, timestamp); err != nil {
		return err
	}
	return nil
}

func (n *handsonUser) RenderingE() error {
	if err := uml.RenderingUml(); err != nil {
		return err
	}
	return nil
}
