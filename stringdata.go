package lnk

import (
	"io"
)

// Optional StringData - Section 2.4.

// StringDataSection represents section 2.4 of the lnk.
// Fields are mostly optional and present if certain flags are set.
type StringDataSection struct {

	// On disk, all these fields have a uint16 size, followed by a string.
	// The string is not null-terminated.
	// The strings are unicode if IsUnicode is set in header.

	// NameString specifies a description of the shortcut that is displayed to
	// end users to identify the purpose of the shell link.
	// Present with HasName flag.
	NameString string

	// RelativePath specifies the location of the link target relative to the
	// file that contains the shell link.
	// Present with HasRelativePath flag.
	RelativePath string

	// WorkingDir specifies the file system path of the working directory to
	// be used when activating the link target.
	// Present with HasWorkingDir flag.
	WorkingDir string

	// CommandLineArguments stores the command-line arguments of the link target.
	// Present with HasArguments flag.
	CommandLineArguments string // String?

	// IconLocation specifies the location of the icon to be used.
	// Present with HasIconLocation flag.
	IconLocation string
}

// StringData parses the StringData portion of the lnk.
// flags is the ShellLinkHeader.LinkFlags.
func StringData(r io.Reader, linkFlags FlagMap) (st StringDataSection, err error) {

	// Read unicode strings if is unicode flag is set.
	isUnicode := linkFlags["IsUnicode"]

	// Read NameString if HasName flag is set.
	if linkFlags["HasName"] {
		name, _ := readStringData(r, isUnicode)
		// fmt.Println("Name", name)
		st.NameString = name
	}

	// Read NameString if HasName flag is set.
	if linkFlags["HasRelativePath"] {
		st.RelativePath, err = readStringData(r, isUnicode)
		if err != nil {
			return st, err
		}
	}

	// Read WorkingDir if HasWorkingDir flag is set.
	if linkFlags["HasWorkingDir"] {
		st.WorkingDir, err = readStringData(r, isUnicode)
		if err != nil {
			return st, err
		}
	}

	// Read CommandLineArguments if HasArguments flag is set.
	if linkFlags["HasArguments"] {
		st.CommandLineArguments, err = readStringData(r, isUnicode)
		if err != nil {
			return st, err
		}
	}

	// Read IconLocation if HasIconLocation flag is set.
	if linkFlags["HasIconLocation"] {
		st.IconLocation, err = readStringData(r, isUnicode)
		if err != nil {
			return st, err
		}
	}
	return st, err
}
