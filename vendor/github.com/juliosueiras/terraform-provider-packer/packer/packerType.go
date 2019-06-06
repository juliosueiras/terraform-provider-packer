package packer

type Packer struct {
	Builders []Builder `json:"buliders"`
}

type Builder struct {
	Type string `json:"type"`
	VMwareISOBuilder
	Communicator string `json:"communicator"`
}

type VMwareISOBuilder struct {
	IsoURL          string `json:"iso_url"`
	IsoChecksum     string `json:"iso_checksum"`
	IsoChecksumType string `json:"iso_checksum_type"`
	SSHUsername     string `json:"ssh_username"`
	SSHPassword     string `json:"ssh_password"`
	ShutdownCommand string `json:"shutdown_command"`
}
