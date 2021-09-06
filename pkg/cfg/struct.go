package cfg

type profile struct {
	Profile profiles `yaml:"profile`
}

type profiles struct {
	Path    path    `yaml:"path"`
	Feature feature `yaml:"feature"`
}

type path struct {
	Java      string `yaml:"java"`
	Wireshark string `yaml:"wireshark"`
	Plantuml  string `yaml:"plantuml"`
}

type feature struct {
	Timestamp      string `yaml:"timestamp"`
	NameResolution string `yaml:"nameResolution"`
}

type configPath struct {
	path string
}

type Config struct {
	Java           string
	Wireshark      string
	Plantuml       string
	Timestamp      bool
	NameResolution bool
}

type disassembledCharacter struct {
	homedir  string
	separate string
	filename string
}
