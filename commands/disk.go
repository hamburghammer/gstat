package commands

import (
	"encoding/json"

	"github.com/hamburghammer/gstat/args"
	goDisk "github.com/shirou/gopsutil/disk"
)

// Disk the struct for the DI to compose the disk space reader
type Disk struct {
	ReadDiskStats func(string) (*goDisk.UsageStat, error)
}

// NewDisk is a ctor for the Disk struct
func NewDisk() Disk {
	return Disk{ReadDiskStats: goDisk.Usage}
}

// Exec gets the disk space value for the root partition and maps it to the executiondata struct
func (d Disk) Exec(args args.Arguments) ([]byte, error) {
	if !args.Disk {
		return []byte{}, nil
	}

	usage, err := d.ReadDiskStats("/")
	if err != nil {
		return []byte{}, err
	}

	data := struct {
		Disk Memory `json:"disk"`
	}{Disk: Memory{Used: bytesToMegaByte(usage.Used), Total: bytesToMegaByte(usage.Total)}}
	return json.Marshal(data)
}
