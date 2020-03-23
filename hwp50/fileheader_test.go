package hwp50

import (
	"fmt"
	"os"
	"testing"
)

// TestDeserializeFileHeader seeks to the hwp header from a .hwp file and
// reads it.
// TestDeserializeFileHeader는 .hwp 파일 헤더를 파일에서 찾아서 읽습니다.
func TestDeserializeFileHeader(t *testing.T) {
	f, err := os.Open("testdata")
	if err != nil {
		t.Fatal(err)
	}

	var compare [32]byte
	var offset int

	// Bit of a hack. Seeks for the hwp header as .hwp file uses the OLE
	// file structure. Records the offset
	// OLE 파일 구조를 무시하고 hwp파일 Signature만 찾아서 offset을
	// 기록합니다.
	for {
		// read 32 bytes to the compare var
		// 파일에서 32바이트씩 읽기
		// i returns how many bytes read
		// i는 얼마나 읽혔는지를 나타냅니다
		i, err := f.Read(compare[:])
		if err != nil {
			t.Fatal(err)
		}
		// If Signature found, then break
		// Signature 찾으면 break
		if compare == Signature {
			break
		}
		// add onto the offset last
		// 마지막으로 offset에 값 추가하기
		offset += i
	}

	// Seek back to the start of the hwp file header
	// hwp Signature를 찾은 곳으로 seek
	f.Seek(int64(offset), 0)

	// init header
	// 헤더 init
	var h FileHeader
	err = h.DeserializeFileHeader(f)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("FileVersion", h.Version.Major, h.Version.Minor,
		h.Version.Micro, h.Version.Extra)

	fmt.Println("")
	fmt.Println("IsCompressed?:", h.Fp.IsCompressed())
	fmt.Println("IsEncrypted?:", h.Fp.IsEncrypted())
	fmt.Println("IsExported?:", h.Fp.IsExported())
	fmt.Println("HasScript?:", h.Fp.HasScript())
	fmt.Println("HasDRM?:", h.Fp.HasDRM())
	fmt.Println("HasXMLTemplateStorage?:", h.Fp.HasXMLTemplateStorage())
	fmt.Println("HasFileHistory?:", h.Fp.HasFileHistory())
	fmt.Println("HasDigitalSig?:", h.Fp.HasDigitalSig())
	fmt.Println("IsEncryptedWithKISAKey?:", h.Fp.IsEncryptedWithKISAKey())
	fmt.Println("HasSpareDigitalSig?:", h.Fp.HasSpareDigitalSig())
	fmt.Println("HasKISADRM?:", h.Fp.HasKISADRM())
	fmt.Println("HasCCL?:", h.Fp.HasCCL())
	fmt.Println("IsMobileOptimized?:", h.Fp.IsMobileOptimized())
	fmt.Println("IsPrivateInfoProtected?:", h.Fp.IsPrivateInfoProtected())
	fmt.Println("IsModificationTracked?:", h.Fp.IsModificationTracked())
	fmt.Println("HasKOGLLicense?:", h.Fp.HasKOGLLicense())
	fmt.Println("HasVideoControls?:", h.Fp.HasVideoControls())
	fmt.Println("HasChapterControlField?:", h.Fp.HasChapterControlField())

	fmt.Println("")
	fmt.Println("HasLicenseInfo?:", h.Sp.HasLicenseInfo())
	fmt.Println("IsCopyProtected?:", h.Sp.IsCopyProtected())
	fmt.Println("IsAllowedToCopyWithoutModification?:",
		h.Sp.IsAllowedToCopyWithoutModification())

	fmt.Println("")
	fmt.Println("EncryptVersion", h.EncryptVersion)

	fmt.Println("")
	fmt.Println("KOGLCountry", h.KOGLCountry)
}
