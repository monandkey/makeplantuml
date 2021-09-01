package makeplantuml

import (
	"os"
	"fmt"
	"regexp"
	"strings"
	"runtime"
	"os/exec"
)

type tsharkHeader struct {
	number   string
	time     string
	srcAddr  string
	srcPort  string
	dstAddr  string
	dstPort  string
	protocol string
	message  string
	checksum string
}

func setAddress(v4 string, v6 string) string {
	v4 = regexp.MustCompile("\"").ReplaceAllString(v4, "")
	v6 = regexp.MustCompile("\"").ReplaceAllString(v6, "")

	if v4 != "" && v6 == "" {
		return v4
	}

	if v4 == "" && v6 != "" {
		return v6
	}

	return ""
}

func setPortAndCheckSum(u string, t string, s string) string {
	u = regexp.MustCompile("\"").ReplaceAllString(u, "")
	t = regexp.MustCompile("\"").ReplaceAllString(t, "")
	s = regexp.MustCompile("\"").ReplaceAllString(s, "")

	if u != "" && t == "" && s == "" {
		return u
	}

	if u == "" && t != "" && s == "" {
		return t
	}

	if u == "" && t == "" && s != "" {
		return t
	}

	return ""
}

type tsharkHeaders []tsharkHeader

func (t tsharkHeaders) setHeader(out string) {
	for regexp.MustCompile(",,").Match([]byte(out)) {
		out = regexp.MustCompile(",{2}").ReplaceAllString(out, ",\"\",")
	}

	lines := strings.Split(out, "\n")

	for _, line := range lines {
		if regexp.MustCompile(",$").Match([]byte(line)) {
			line = regexp.MustCompile(",$").ReplaceAllString(line, ",\"\"")
		}

		column := regexp.MustCompile("\".*?\"").FindAllStringSubmatch(line, -1)
		if len(column) == 0 {
			continue
		}

		tNumber := regexp.MustCompile("\"").ReplaceAllString(column[0][0], "")
		tTime := regexp.MustCompile("\"").ReplaceAllString(column[1][0], "")
		tSrcAddr := setAddress(column[2][0], column[3][0])
		tSrcPort := setPortAndCheckSum(column[4][0], column[5][0], column[6][0])
		tDstAddr := setAddress(column[7][0], column[8][0])
		tDstPort := setPortAndCheckSum(column[9][0], column[10][0], column[11][0])
		tProtocol := regexp.MustCompile("\"").ReplaceAllString(column[12][0], "")
		tMessage := regexp.MustCompile("\"").ReplaceAllString(column[13][0], "")
		tChecksum := setPortAndCheckSum(column[14][0], column[15][0], column[16][0])

		th := tsharkHeader{
			number:   tNumber,
			time:     tTime,
			srcAddr:  tSrcAddr,
			srcPort:  tSrcPort,
			dstAddr:  tDstAddr,
			dstPort:  tDstPort,
			protocol: tProtocol,
			message:  tMessage,
			checksum: tChecksum,
		}
		t = append(t, th)
	}
}

func RunTshark() {
	var cmd string
	switch(runtime.GOOS) {
		case "windows":
			cmd = "C:/Program Files/Wireshark-3.2.3/tshark.exe"
		case "linux":
			cmd = "tshark"
		default:
			fmt.Println("Your OS not support.")
			return
	}

	out, err := exec.Command(cmd,
		"-r", "./sample/440110000001519.pcap",
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
		"-e", "udp.checksum",
		"-e", "sctp.checksum",
		"-e", "tcp.checksum",
	).Output()

	if err != nil {
		fmt.Println(os.Stderr, err)
		return
	}

	var fmtOut tsharkHeaders
	fmtOut.setHeader(string(out))
}
