package lnk

import (
	"encoding/binary"
	"fmt"
	"io"
)

// ItemList structure.
// If Header has HasLinkTargetIDList, then header is immediately followed by
// one LinkTargetIDList structure.

// LinkTargetIDList contains information about the target of the link.
// Section 2.2 in= [MS-SHLLINK]
type LinkTargetIDList struct {
	// First two bytes is IDListSize.
	IDListSize uint16
	// Data containing IDLists and TerminalID.
	List IDList
}

// IDList represents a persisted item ID list.
type IDList struct {
	// ItemIDList contains the IDLists.
	ItemIDList []ItemID
	// TerminalID is 00 00.
	TerminalID uint16
}

// ItemID is an element from IDList.
// From [MS-SHLLINK]:
// "The data stored in a given ItemID is defined by the source that corresponds
// to the location in the target namespace of the preceding ItemIDs. This data
// uniquely identifies the items in that part of the namespace."
type ItemID struct {
	// Size of ItemID INCLUDING the size. Why?
	Size uint16
	// Data length is size-2 bytes.
	Data []byte
}

// LinkTarget returns a populated LinkTarget based on bytes passed. []byte
// should point to the start of the section. Normally this will be offset 0x4c
// of the lnk file.
func LinkTarget(r io.Reader) (li LinkTargetIDList, err error) {

	// Read the first two bytes to get the IDListSize.
	err = binary.Read(r, binary.LittleEndian, &li.IDListSize)
	if err != nil {
		return li, fmt.Errorf("lnk.LinkTarget: read IDListSize - %s", err.Error())
	}
	fmt.Println(li.IDListSize)

	// Instead of reading IDListSize bytes, we read uint16 which is length, if
	// this item is zero, we have reached TerminalID which is 00 00. If not, read
	// that many bytes. If the file format is wrong, we may bleed into the next
	// section, but then again the IDListSize might be wrong too.

	// listData := make([]byte, li.IDListSize-2)
	// err = binary.Read(r, binary.LittleEndian, &listData)
	// if err != nil {
	// 	return li, fmt.Errorf("lnk.LinkTarget: read IDList bytes - %s", err.Error())
	// }
	// fmt.Println(len(listData))
	// // Create an io.Reader from buffer.
	// buf := bytes.NewReader(listData[:])

	// 	// Populate TerminalID by reading the last two bytes.
	// 	err = binary.Read(r, binary.LittleEndian, &idList.TerminalID)
	// 	if err != nil {
	// 		return li, fmt.Errorf("lnk.LinkTarget: read IDList.TerminalID - %s", err.Error())
	// 	}
	// 	if idList.TerminalID != 0x00 {
	// 		return li,
	// 			fmt.Errorf("lnk.LinkTarget: TerminalID not zero - got %s", uint16Str(idList.TerminalID))
	// 	}

	var idList IDList

	// Start populating ItemIDs.
	var items []ItemID
	var itemSize uint16
	for {
		err = binary.Read(r, binary.LittleEndian, &itemSize)
		if err != nil {
			return li, fmt.Errorf("lnk.LinkTarget: read item size - %s", err.Error())
		}
		// Check if we have reach the TerminalID
		if itemSize == 0 {
			idList.TerminalID = itemSize
			fmt.Println("Reached TerminalID")
			break
		}
		// If not, read those many bytes-2.
		itemData := make([]byte, itemSize-2)
		err = binary.Read(r, binary.LittleEndian, &itemData)
		if err != nil {
			return li, fmt.Errorf("lnk.LinkTarget: read item data - %s", err.Error())
		}
		items = append(items, ItemID{Size: itemSize, Data: itemData})
	}

	fmt.Println(len(items))

	for _, it := range items {
		fmt.Println("Item Size:", it.Size)
		fmt.Println("Item Data:", string(it.Data))
		fmt.Println("--------------------")
	}

	return li, err
}
