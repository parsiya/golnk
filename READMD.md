# lnk - lnk Parser for Go
lnk is a package for parsing Windows Shell Link (.lnk) files.

It's based on version 5.0 of the [MS-SHLLINK] document:

* Reference: https://msdn.microsoft.com/en-us/library/dd871305.aspx
* Version 5.0: https://winprotocoldoc.blob.core.windows.net/productionwindowsarchives/MS-SHLLINK/[MS-SHLLINK].pdf

## Shell Link Structure
Each file has at least one header (`SHELL_LINK_HEADER`) and one or more optional sections. The existence of these sections are defined by the `LinkFlags` uint32 in the header.

```
SHELL_LINK = SHELL_LINK_HEADER [LINKTARGET_IDLIST] [LINKINFO]
              [STRING_DATA] *EXTRA_DATA
```

Note: "Unless otherwise specified, the value contained by size fields includes the size of size field itself."

Currently lnk parses every section other than `EXTRA_DATA`. Different data blocks are identified and stored but it does not parse any of them other than identifying the type (via their signature) and storing the content. Data blocks are defined in section 2.5 of the specification.

## Usage
Pass a file name to `lnk.File` which returns a `LnkFile`:

``` go
type LnkFile struct {
	Header     ShellLinkHeader   // File header.
	IDList     LinkTargetIDList  // LinkTargetIDList.
	LinkInfo   LinkInfoSection   // LinkInfo.
	StringData StringDataSection // StringData.
	DataBlocks ExtraData         // ExtraData blocks.
}
```

Each section is a struct that is populated. See their fields in their respective source files. Path to the target file is usually in `LinkInfo.LocalBasePath`.

``` go
code
```





Create an `io.Reader` from the contents and then 