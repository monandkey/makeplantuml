package tshark

import (
	"os"
	"fmt"
	"runtime"
	"os/exec"
	"local.packages/cfg"
)

func RunTshark(fileName string) TsharkHeaders {
	var (
		fmtOut TsharkHeaders
		cmd string
	)

	if cfg.CfgVal.Profile.Path.Wireshark == "default" {
		switch(runtime.GOOS) {
			case "windows":
				cmd = "C:/Program Files/Wireshark/tshark.exe"
			case "linux":
				cmd = "tshark"
			default:
				fmt.Println("Your OS not support.")
				return fmtOut
		}

	} else {
		cmd = cfg.CfgVal.Profile.Path.Wireshark + "\\tshark.exe"
	}

	out, err := exec.Command(cmd,
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
		"-e", "udp.checksum",
		"-e", "sctp.checksum",
		"-e", "tcp.checksum",
		"-e", "ip.len",
		"-e", "ipv6.plen",
	).Output()

	if err != nil {
		fmt.Println(os.Stderr, err)
		return fmtOut
	}

	return fmtOut.SetHeader(string(out))
}
