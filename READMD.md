# golnk - lnk Parser for Go

Reference: https://msdn.microsoft.com/en-us/library/dd871305.aspx
Version 5.0: https://winprotocoldoc.blob.core.windows.net/productionwindowsarchives/MS-SHLLINK/[MS-SHLLINK].pdf

# Structure

```
SHELL_LINK = SHELL_LINK_HEADER [LINKTARGET_IDLIST] [LINKINFO]
              [STRING_DATA] *EXTRA_DATA
```

## SHELL_LINK_HEADER
Contains identification information, timestamps, and flags that specify the presence of optional structures.

Little-endian, remember.

* HeaderSize (4 bytes): The size, in bytes, of this structure. This value MUST be 0x0000004C.
* LinkCLSID (16 bytes): A class identifier (CLSID). This value MUST be 00021401-0000-0000-C000-000000000046.
* LinkFlags (4 bytes): A LinkFlags structure (section 2.1.1) that specifies information about the shell link and the presence of optional portions of the structure.
* FileAttributes (4 bytes): A FileAttributesFlags structure (section 2.1.2) that specifies information about the link target.
* CreationTime (8 bytes): A FILETIME structure ([MS-DTYP] section 2.3.3) that specifies the creation time of the link target in UTC (Coordinated Universal Time). If the value is zero, there is no creation time set on the link target.
* AccessTime (8 bytes): A FILETIME structure ([MS-DTYP] section 2.3.3) that specifies the access time of the link target in UTC (Coordinated Universal Time). If the value is zero, there is no access time set on the link target.
* WriteTime (8 bytes): A FILETIME structure ([MS-DTYP] section 2.3.3) that specifies the write time of the link target in UTC (Coordinated Universal Time). If the value is zero, there is no write time set on the link target.
* FileSize (4 bytes): A 32-bit unsigned integer that specifies the size, in bytes, of the link target. If the link target file is larger than 0xFFFFFFFF, this value specifies the least significant 32 bits of the link target file size.
* IconIndex (4 bytes): A 32-bit signed integer that specifies the index of an icon within a given icon location.
* ShowCommand (4 bytes): A 32-bit unsigned integer that specifies the expected window state of an application launched by the link. This value SHOULD be one of the following.

"Unless otherwise specified, the value contained by size fields includes the size of size field itself."
