package lnk

import (
	"encoding/hex"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

var sampleFixtures = []string{
	"remote.directory.xp.test",
	"remote.file.xp.test",
	"test-orig.lnk",
	"test.lnk",
	"test.lnk.bak",
	"vbox-svr-win10.lnk",
	"Visual Studio Code.lnk",
	"Windows Store.lnk",
}

var remoteFixtures = []string{
	"remote.directory.xp.test",
	"remote.file.xp.test",
}

func samplePath(name string) string {
	return filepath.Join("test", name)
}

func readSampleFixture(t *testing.T, name string) LnkFile {
	t.Helper()

	fixturePath := samplePath(name)
	parsed, err := File(fixturePath)
	if err != nil {
		t.Fatalf("File(%q) error = %v", fixturePath, err)
	}
	return parsed
}

func TestFileParsesSampleFixtures(t *testing.T) {
	for _, fixture := range sampleFixtures {
		t.Run(fixture, func(t *testing.T) {
			parsed := readSampleFixture(t, fixture)
			if parsed.Header.Magic == 0 {
				t.Fatalf("parsed header magic is zero for %q", fixture)
			}
		})
	}
}

func TestReadMatchesFileForSampleFixtures(t *testing.T) {
	for _, fixture := range sampleFixtures {
		t.Run(fixture, func(t *testing.T) {
			fixturePath := samplePath(fixture)

			fi, err := os.Open(fixturePath)
			if err != nil {
				t.Fatalf("os.Open(%q) error = %v", fixturePath, err)
			}
			defer fi.Close()

			stat, err := fi.Stat()
			if err != nil {
				t.Fatalf("Stat(%q) error = %v", fixturePath, err)
			}

			fromReader, err := Read(fi, uint64(stat.Size()))
			if err != nil {
				t.Fatalf("Read(%q) error = %v", fixturePath, err)
			}

			fromFile := readSampleFixture(t, fixture)
			if !reflect.DeepEqual(fromReader, fromFile) {
				t.Fatalf("Read(%q) and File(%q) returned different results", fixturePath, fixturePath)
			}
		})
	}
}

func TestSampleHeadersLookValid(t *testing.T) {
	for _, fixture := range sampleFixtures {
		t.Run(fixture, func(t *testing.T) {
			parsed := readSampleFixture(t, fixture)

			if parsed.Header.Magic != headerSize {
				t.Fatalf("Header.Magic = %#x, want %#x", parsed.Header.Magic, headerSize)
			}
			if got := hex.EncodeToString(parsed.Header.LinkCLSID[:]); got != classID {
				t.Fatalf("Header.LinkCLSID = %q, want %q", got, classID)
			}
			if len(parsed.Header.Raw) != headerSize {
				t.Fatalf("len(Header.Raw) = %d, want %d", len(parsed.Header.Raw), headerSize)
			}
			if parsed.Header.Reserved1 != 0 || parsed.Header.Reserved2 != 0 || parsed.Header.Reserved3 != 0 {
				t.Fatalf("reserved header fields were not zero: %#v %#v %#v", parsed.Header.Reserved1, parsed.Header.Reserved2, parsed.Header.Reserved3)
			}
		})
	}
}

func TestSampleSectionsTrackHeaderFlags(t *testing.T) {
	for _, fixture := range sampleFixtures {
		t.Run(fixture, func(t *testing.T) {
			parsed := readSampleFixture(t, fixture)

			if parsed.Header.LinkFlags["HasLinkTargetIDList"] {
				if parsed.IDList.IDListSize == 0 {
					t.Fatalf("HasLinkTargetIDList was set but IDListSize was zero")
				}
			} else if parsed.IDList.IDListSize != 0 {
				t.Fatalf("HasLinkTargetIDList was not set but IDListSize = %d", parsed.IDList.IDListSize)
			}

			if parsed.Header.LinkFlags["HasLinkInfo"] {
				if parsed.LinkInfo.Size == 0 {
					t.Fatalf("HasLinkInfo was set but LinkInfo.Size was zero")
				}
				if len(parsed.LinkInfo.Raw) == 0 {
					t.Fatalf("HasLinkInfo was set but LinkInfo.Raw was empty")
				}
			} else if parsed.LinkInfo.Size != 0 {
				t.Fatalf("HasLinkInfo was not set but LinkInfo.Size = %d", parsed.LinkInfo.Size)
			}

			if parsed.LinkInfo.CommonNetworkRelativeLinkOffset != 0 && parsed.LinkInfo.NetworkRelativeLink.Size == 0 {
				t.Fatalf("CommonNetworkRelativeLinkOffset was set but parsed network link size was zero")
			}

			if parsed.DataBlocks.TerminalBlock >= 0x04 {
				t.Fatalf("TerminalBlock = %#x, want a value < 0x04", parsed.DataBlocks.TerminalBlock)
			}

			for index, block := range parsed.DataBlocks.Blocks {
				if block.Size < 8 {
					t.Fatalf("block %d size = %d, want >= 8", index, block.Size)
				}
				if block.Type == "" {
					t.Fatalf("block %d type was empty", index)
				}
			}
		})
	}
}

func TestRemoteFixturesContainNetworkLinkInfo(t *testing.T) {
	for _, fixture := range remoteFixtures {
		t.Run(fixture, func(t *testing.T) {
			parsed := readSampleFixture(t, fixture)

			if !parsed.Header.LinkFlags["HasLinkInfo"] {
				t.Fatalf("HasLinkInfo was false for remote fixture %q", fixture)
			}
			if parsed.LinkInfo.CommonNetworkRelativeLinkOffset == 0 {
				t.Fatalf("CommonNetworkRelativeLinkOffset was zero for remote fixture %q", fixture)
			}
			if parsed.LinkInfo.NetworkRelativeLink.Size == 0 {
				t.Fatalf("NetworkRelativeLink.Size was zero for remote fixture %q", fixture)
			}
			if parsed.LinkInfo.CommonPathSuffix == "" {
				t.Fatalf("CommonPathSuffix was empty for remote fixture %q", fixture)
			}
		})
	}
}

func TestSampleStringersProduceStructuredOutput(t *testing.T) {
	for _, fixture := range []string{"test.lnk", "remote.directory.xp.test", "Visual Studio Code.lnk"} {
		t.Run(fixture, func(t *testing.T) {
			parsed := readSampleFixture(t, fixture)

			headerOut := strings.ToLower(strings.ReplaceAll(parsed.Header.String(), " ", ""))
			if !strings.Contains(headerOut, "magic") || !strings.Contains(headerOut, "value") {
				t.Fatalf("Header.String() output did not include expected table content for %q", fixture)
			}
			if parsed.LinkInfo.Size != 0 {
				linkInfoOut := strings.ToLower(strings.ReplaceAll(parsed.LinkInfo.String(), " ", ""))
				if !strings.Contains(linkInfoOut, "headersize") || !strings.Contains(linkInfoOut, "value") {
					t.Fatalf("LinkInfo.String() output did not include expected table content for %q", fixture)
				}
			}
			if parsed.LinkInfo.CommonNetworkRelativeLinkOffset != 0 {
				networkOut := strings.ToLower(strings.ReplaceAll(parsed.LinkInfo.NetworkRelativeLink.String(), " ", ""))
				if !strings.Contains(networkOut, "networkprovidertype") || !strings.Contains(networkOut, "flags") {
					t.Fatalf("CommonNetworkRelativeLink.String() output did not include expected table content for %q", fixture)
				}
			}
		})
	}
}

func TestHeaderRejectsTooSmallMaxSize(t *testing.T) {
	fi, err := os.Open(samplePath("test.lnk"))
	if err != nil {
		t.Fatalf("os.Open(test fixture) error = %v", err)
	}
	defer fi.Close()

	_, err = Header(fi, 8)
	if err == nil {
		t.Fatal("Header() error = nil, want an error for too-small maxSize")
	}
	if !strings.Contains(err.Error(), "invalid computed size") {
		t.Fatalf("Header() error = %q, want an invalid computed size error", err.Error())
	}
}
