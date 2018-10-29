package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"os"
)

// Lnk represents the structure of an lnk file.
// Almost everything is little-endian.
type Lnk struct {
	SHL ShellLinkHeader
}

// ShellLinkHeader represents the lnk header.
type ShellLinkHeader struct {
	Header       uint32   // Header size: should be 0x4c.
	LinkCLSID    [16]byte // A class identifier, should be  00021401-0000-0000-C000-000000000046.
	LinkFlags    uint32   // File attributes about link target. TODO: what format is this?
	CreationTime uint64   // Creation time of link target in UTC. TODO: what format is this? Could be zero.
	AccessTime   uint64   // Access time of link target. Could be zero.
	WriteTime    uint64   // Write time  of link target. Could be zero.
	FileSize     uint32   // Filesize of link target. If larger than capacity, it will have the LSB 32-bits of size.
	IconIndex    int32    // 32-bit signed integer, the index of an icon within a given icon location. TODO: is it just a number to create the icon of the lnk file based on the target?
	ShowCommand  uint32   // uint32 integer that is the expected windows state of the target after execution.
	// Valid values:
	// 0x00000001 - SW_SHOWNORMAL - The application is open and its window is open in a normal fashion.
	// 0x00000003 - SW_SHOWMAXIMIZED - The application is open, and keyboard focus is given to the application, but its window is not shown.
	// 0x00000007 - SW_SHOWMINNOACTIVE - The application is open, but its window is not shown. It is not given the keyboard focus.
	// All other values are SW_SHOWNORMAL.

	HotKey    uint32 // HotKeyFlags structure to launch the target.
	Reserved1 uint16 // Zero
	Reserved2 uint32 // Zero
	Reserved3 uint32 // Zero
}

// LinkFlags defines what shell link structures are in the file.
type LinkFlags struct {
	HasLinkTargetIDList         bool // bit00 - ShellLinkHeader is followd by a LinkTargetIDList structure.
	HasLinkInfo                 bool // bit01 - LinkInfo in file.
	HasName                     bool // bit02 - NAME_String in file.
	HasRelativePath             bool // bit03 - RELATIVE_PATH in file.
	HasWorkingDir               bool // bit04 - WORKING_DIR in file.
	HasArguments                bool // bit05 - COMMAND_LINE_ARGUMENTS
	HasIconLocation             bool // bit06 - ICON_LOCATION
	IsUnicode                   bool // bit07 - Strings are in unicode
	ForceNoLinkInfo             bool // bit08 - LinkInfo is ignored
	HasExpString                bool // bit09 - The shell link is saved with an EnvironmentVariableDataBlock
	RunInSeparateProcess        bool // bit10 - Target runs in a 16-bit virtual machine
	Unused1                     bool // bit11 - ignore
	HasDarwinID                 bool // bit12 - The shell link is saved with a DarwinDataBlock
	RunAsUser                   bool // bit13 - The application is run as a different user when the target of the shell link is activated.
	HasExpIcon                  bool // bit14 - The shell link is saved with an IconEnvironmentDataBlock
	NoPidlAlias                 bool // bit15 - The file system location is represented in the shell namespace when the path to an item is parsed into an IDList.
	Unused2                     bool // bit16 - ignore
	RunWithShimLayer            bool // bit17 - The shell link is saved with a ShimDataBlock.
	ForceNoLinkTrack            bool // bit18 - The TrackerDataBlock is ignored.
	EnableTargetMetadata        bool // bit19 - The shell link attempts to collect target properties and store them in the PropertyStoreDataBlock (section 2.5.7) when the link target is set.
	DisableLinkPathTracking     bool // bit20 - The EnvironmentVariableDataBlock is ignored.
	DisableKnownFolderTracking  bool // bit21 - The SpecialFolderDataBlock (section 2.5.9) and the KnownFolderDataBlock (section 2.5.6) are ignored when loading the shell link. If this bit is set, these extra data blocks SHOULD NOT be saved when saving the shell link.
	DisableKnownFolderAlias     bool // bit22 - If the link has a KnownFolderDataBlock (section 2.5.6), the unaliased form of the known folder IDList SHOULD be used when translating the target IDList at the time that the link is loaded.
	AllowLinkToLink             bool // bit23 - Creating a link that references another link is enabled. Otherwise, specifying a link as the target IDList SHOULD NOT be allowed.
	UnaliasOnSave               bool // bit24 - When saving a link for which the target IDList is under a known folder, either the unaliased form of that known folder or the target IDList SHOULD be used.
	PreferEnvironmentPath       bool // bit25 - The target IDList SHOULD NOT be stored; instead, the path specified in the EnvironmentVariableDataBlock (section 2.5.4) SHOULD be used to refer to the target.
	KeepLocalIDListForUNCTarget bool // bit26 - When the target is a UNC name that refers to a location on a local machine, the local path IDList in the PropertyStoreDataBlock (section 2.5.7) SHOULD be stored, so it can be used when the link is loaded on the local machine.

}

func main() {

	fi, err := os.Open("test.lnk")
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	// lnk files are small-ish, no reason not to read everything at once.
	lnkBytes, err := ioutil.ReadAll(fi)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Read %d bytes.\n", len(lnkBytes))

	buf := bytes.NewReader(lnkBytes)
	// First four bytes == header. Must be 0x4c in little-endian.
	var header uint32
	err = binary.Read(buf, binary.LittleEndian, &header)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%x\n", header)

	if header != 0x4c {
		fmt.Printf("Bad header - got %x - wanted %x", header, 0x4c)
	}
}
