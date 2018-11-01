package lnk

// VolumeID (section 2.3.1. of [SHLLINK]).
// Information about the volume that the link target was on when the link was created.
type VolumeID struct {
	// Size of VolumeID (including this?)
	Size uint32

	// Type of drive that target was stored on.
	DriveType uint32

	// Serial number of the volume.
	DriveSerialNumber uint32

	// Offset to the a null-terminated string that contains the volume label of
	// the drive that the link target is stored on.
	// If == 0x14, it must be ignored and VolumeLabelOffsetUnicode must be used.
	VolumeLabelOffset uint32

	// Offset to Unicode version of VolumeLabel.
	// Must not be present if VolumeLabelOffset is not 0x14.
	VolumeLabelOffsetUnicode uint32

	// VolumeLabel either ASCII-HEX or Unicode as defined by value of VolumeLabelOffset.
	Data string
}

// Different DriveTypes. The value of field is the index to this slice.
var driveType = []string{
	"DRIVE_UNKNOWN",
	"DRIVE_NO_ROOT_DIR",
	"DRIVE_REMOVABLE",
	"DRIVE_FIXED",
	"DRIVE_REMOTE",
	"DRIVE_CDROM",
	"DRIVE_RAMDISK",
}
