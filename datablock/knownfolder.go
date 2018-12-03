package datablock

import (
	"encoding/binary"
	"encoding/hex"
	"io"
)

// KnownFolderDataBlock - Section 2.5.6. Specifies the location of a known folder.
type KnownFolderDataBlock struct {
	// GUID of the target folder.
	KnownFolderID [16]byte
	// Offset of the target folder in the IDList.
	Offset uint32
}

// Parse parses the KnownFolderDataBlock. Parse gets the data minus the size and signature.
func (d *KnownFolderDataBlock) Parse(r io.Reader) error {
	// Read 16 bytes.
	err := binary.Read(r, binary.LittleEndian, d.KnownFolderID)
	if err != nil {
		return err
	}
	// Read offset and return.
	return binary.Read(r, binary.LittleEndian, d.Offset)
}

// TableString returns the fields in a [][]string that can be used with
// tablewriter.AppendBulk.
func (d KnownFolderDataBlock) TableString() [][]string {
	var data [][]string
	data = append(data, []string{"KnownFolderID", hex.EncodeToString(d.KnownFolderID[:])})
	data = append(data, []string{"Offset"}, uint32TableStr(d.Offset))

	return data
}
