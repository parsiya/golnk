package lnk

// File represents one lnk file.
type File struct {
	Data   []byte          // File content.
	Size   int             // File size.
	Header ShellLinkHeader // File header.
}

// ShellLinkHeader represents the lnk header.
type ShellLinkHeader struct {
	Header       uint32   // Header size: should be 0x4c.
	LinkCLSID    [16]byte // A class identifier, should be  00021401-0000-0000-C000-000000000046.
	LinkFlags    []string // File attributes about link target, originally a uint32.
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

// linkFlags defines what shell link structures are in the file.
var linkFlags = []string{
	"HasLinkTargetIDList",         // bit00 - ShellLinkHeader is followed by a LinkTargetIDList structure.
	"HasLinkInfo",                 // bit01 - LinkInfo in file.
	"HasName",                     // bit02 - NAME_String in file.
	"HasRelativePath",             // bit03 - RELATIVE_PATH in file.
	"HasWorkingDir",               // bit04 - WORKING_DIR in file.
	"HasArguments",                // bit05 - COMMAND_LINE_ARGUMENTS
	"HasIconLocation",             // bit06 - ICON_LOCATION
	"IsUnicode",                   // bit07 - Strings are in unicode
	"ForceNoLinkInfo",             // bit08 - LinkInfo is ignored
	"HasExpString",                // bit09 - The shell link is saved with an EnvironmentVariableDataBlock
	"RunInSeparateProcess",        // bit10 - Target runs in a 16-bit virtual machine
	"Unused1",                     // bit11 - ignore
	"HasDarwinID",                 // bit12 - The shell link is saved with a DarwinDataBlock
	"RunAsUser",                   // bit13 - The application is run as a different user when the target of the shell link is activated.
	"HasExpIcon",                  // bit14 - The shell link is saved with an IconEnvironmentDataBlock
	"NoPidlAlias",                 // bit15 - The file system location is represented in the shell namespace when the path to an item is parsed into an IDList.
	"Unused2",                     // bit16 - ignore
	"RunWithShimLayer",            // bit17 - The shell link is saved with a ShimDataBlock.
	"ForceNoLinkTrack",            // bit18 - The TrackerDataBlock is ignored.
	"EnableTargetMetadata",        // bit19 - The shell link attempts to collect target properties and store them in the PropertyStoreDataBlock (section 2.5.7) when the link target is set.
	"DisableLinkPathTracking",     // bit20 - The EnvironmentVariableDataBlock is ignored.
	"DisableKnownFolderTracking",  // bit21 - The SpecialFolderDataBlock (section 2.5.9) and the KnownFolderDataBlock (section 2.5.6) are ignored when loading the shell link. If this bit is set, these extra data blocks SHOULD NOT be saved when saving the shell link.
	"DisableKnownFolderAlias",     // bit22 - If the link has a KnownFolderDataBlock (section 2.5.6), the unaliased form of the known folder IDList SHOULD be used when translating the target IDList at the time that the link is loaded.
	"AllowLinkToLink",             // bit23 - Creating a link that references another link is enabled. Otherwise, specifying a link as the target IDList SHOULD NOT be allowed.
	"UnaliasOnSave",               // bit24 - When saving a link for which the target IDList is under a known folder, either the unaliased form of that known folder or the target IDList SHOULD be used.
	"PreferEnvironmentPath",       // bit25 - The target IDList SHOULD NOT be stored; instead, the path specified in the EnvironmentVariableDataBlock (section 2.5.4) SHOULD be used to refer to the target.
	"KeepLocalIDListForUNCTarget", // bit26 - When the target is a UNC name that refers to a location on a local machine, the local path IDList in the PropertyStoreDataBlock (section 2.5.7) SHOULD be stored, so it can be used when the link is loaded on the local machine.
}

// fileAttributesFlags represent target file attributes.
var fileAttributesFlags = []string{
	"FILE_ATTRIBUTE_READONLY",
	"FILE_ATTRIBUTE_HIDDEN",
	"FILE_ATTRIBUTE_SYSTEM",
	"Reserved1", // Must be zero.
	"FILE_ATTRIBUTE_DIRECTORY",
	"FILE_ATTRIBUTE_ARCHIVE",
	"Reserved2", // Must be zero.
	"FILE_ATTRIBUTE_NORMAL",
	"FILE_ATTRIBUTE_TEMPORARY",
	"FILE_ATTRIBUTE_SPARSE_FILE",
	"FILE_ATTRIBUTE_REPARSE_POINT",
	"FILE_ATTRIBUTE_COMPRESSED",
	"FILE_ATTRIBUTE_OFFLINE",
	"FILE_ATTRIBUTE_NOT_CONTENT_INDEXED",
	"FILE_ATTRIBUTE_ENCRYPTED",
}
