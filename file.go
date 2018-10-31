package lnk

// File represents one lnk file.
type File struct {
	Data   []byte          // File content.
	Size   int             // File size.
	Header ShellLinkHeader // File header.
	// TODO: Add as you go.
}
