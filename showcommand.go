package lnk

// ShowCommand valid values:
// 0x00000001 - SW_SHOWNORMAL - The application is open and its window is open in a normal fashion.
// 0x00000003 - SW_SHOWMAXIMIZED - The application is open, and keyboard focus is given to the application, but its window is not shown.
// 0x00000007 - SW_SHOWMINNOACTIVE - The application is open, but its window is not shown. It is not given the keyboard focus.
// All other values are SW_SHOWNORMAL.

// showCommand returns the string associated with ShowCommand uint32 field.
func showCommand(s uint32) string {
	switch s {
	case 0x03:
		return "SW_SHOWMAXIMIZED"
	case 0x07:
		return "SW_SHOWMINNOACTIVE"
	}
	// Anything other than these two (include 0x01) is "SW_SHOWNORMAL"
	return "SW_SHOWNORMAL"
}
