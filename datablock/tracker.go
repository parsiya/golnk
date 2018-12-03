package datablock

import (
	"encoding/binary"
	"encoding/hex"
	"io"
)

// TrackerDataBlock specifies data that can be used to resolve a link target
// if it's not found in its original location.
type TrackerDataBlock struct {
	// Length of the rest of the struct.
	Length uint32
	// Version must be 0x00000000.
	Version uint32
	// MachineID is the NetBIOS name of the last machine that the file was residing.
	MachineID [16]byte
	// Droid you are looking for?
	Droid [32]byte
	// DroidBirth?
	DroidBirth [32]byte
}

// Parse parses the TrackerDataBlock.
func (t *TrackerDataBlock) Parse(r io.Reader) error {
	// Read length.
	err := binary.Read(r, binary.LittleEndian, t.Length)
	if err != nil {
		return err
	}
	// Read Version.
	err = binary.Read(r, binary.LittleEndian, t.Version)
	if err != nil {
		return err
	}
	// Read MachineID.
	err = binary.Read(r, binary.LittleEndian, t.MachineID)
	if err != nil {
		return err
	}
	// Read Droid.
	err = binary.Read(r, binary.LittleEndian, t.Droid)
	if err != nil {
		return err
	}
	// Read DroidBirth.
	err = binary.Read(r, binary.LittleEndian, t.DroidBirth)
	if err != nil {
		return err
	}
	return nil
}

// TableString returns the fields in a [][]string that can be used with
// tablewriter.AppendBulk.
func (t TrackerDataBlock) TableString(r io.Reader) [][]string {
	var data [][]string
	data = append(data, []string{"Length2", uint32TableStr(t.Length)})
	data = append(data, []string{"Version", uint32TableStr(t.Version)})
	data = append(data, []string{"MachineID", hex.EncodeToString(t.MachineID[:])})
	data = append(data, []string{"Droid", hex.EncodeToString(t.Droid[:])})
	data = append(data, []string{"DroidBirth", hex.EncodeToString(t.DroidBirth[:])})

	return data
}
