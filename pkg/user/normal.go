package user

import (
	"fmt"
	"regexp"
	"os/exec"
	"local.packages/tshark"
)

func (n NormalUser) new() tshark.TsharkMethod {
	return &NormalUser{}
}

func (n *NormalUser) SetCmd() {
	n.Cmd = tshark.GetTsharkCommand()
}

func (n *NormalUser) SetArgs(fileName string) {
	n.Args = []string{
		"-r", fileName,
		"-t", "ad",
		"-T", "fields",
		"-E", "separator=,",
		"-E", "quote=d",
		"-e", "frame.number",     //  0
		"-e", "_ws.col.Time",     //  1
		"-e", "ip.src",           //  2
		"-e", "ipv6.src",         //  3
		"-e", "udp.srcport",      //  4
		"-e", "tcp.srcport",      //  5
		"-e", "sctp.srcport",     //  6
		"-e", "ip.dst",           //  7
		"-e", "ipv6.dst",         //  8
		"-e", "udp.dstport",      //  9
		"-e", "tcp.dstport",      // 10
		"-e", "sctp.dstport",     // 11
		"-e", "_ws.col.Protocol", // 12
		"-e", "_ws.col.Info",     // 13
		"-e", "ip.len",           // 14
		"-e", "ipv6.plen",        // 15
	}
}

func (n *NormalUser) RunE() error {
	var err error
	n.Out, err = exec.Command(n.Cmd, n.Args...).Output()

	if err != nil {
		return err
	}
	return nil
}

func (n *NormalUser) Parse() {
	lines := tshark.ConvertOutputResultIntoArray(string(n.Out))

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
		n.Header = append(n.Header, []map[string]string{tmph}...)
	}
}

func (n *NormalUser) Display() {
	fmt.Println(n.Header)
	return
}

