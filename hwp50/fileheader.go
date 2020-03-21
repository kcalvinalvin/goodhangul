package hwp50

import (
	"fmt"
	"io"
)

const (
	// Signature is the first 32 bytes that the hwp50 file format is going
	// to be.
	// Signature는 첫 32 바이트의 hwp50 포맷을 뜻합니다.
	Signature = 0xd0, 0xcf, 0xd0, 0xcf, 0x11, 0xe0, 0xa1, 0xb1, 0x1a, 0xe1,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x3e, 0x00, 0x03, 0x00,
		0xfe, 0xff, 0x09, 0x00
)

// FirstProperty represents the 4 byte bitflag that hwp50 uses for various settings
// FirstProperty는 hwp50 파일이 사용하는 4바이트 bitflag를 뜻합니다.
type FirstProperty [4]byte

const (
	// IsCompressed denotes if the file is compressed
	// IsCompressed는 파일이 압축되었는지를 나타냅니다.
	IsCompressed FirstProperty = 1 << iota

	// IsEncrypted denotes if the file is encrypted
	// IsEncrypted는 파일이 암호화 되었는지를 나타냅니다.
	IsEncrypted

	// IsExported denotes if the file is an exported file.
	// See directory distribution/
	// IsExported는 파일이 배포용 문서인지를 나타냅니다.
	// distribution/ 다이렉토리를 참고하시기 바랍니다.
	IsExported

	// HasScript denotes if the file has scripts
	// HasScript는 파일이 스크립트를 저장하는지를 나타냅니다.
	HasScript

	// HasDRM denotes if the file is DRMed. Ew
	// HasDRM은 파일이 DRM 걸려있는지를 나타냅니다.
	HasDRM

	// HasXMLTemplateStorage denotes if the file has XMLTemplate storage
	// HasXMLTemplateStorage는 파일이 XMLTemplate storage가 있는지를
	// 나타냅니다.
	HasXMLTemplateStorage

	// HasFileHistory denotes if the file history is included
	// HasFileHistory는 파일이 이력을 저장했는지를 나타냅니다.
	HasFileHistory

	// HasDigitalSig denotes if the file has a digital signature
	// HasDigitalSig는 전자 서명 정보가 있는지를 나타냅니다.
	HasDigitalSig

	// IsEncryptedWithKISAKey denotes if the file is encrypted with a KISA key
	// A KISA key is a key that the Korean government authorized a private
	// company for the safekeeping of all bank accounts in Korea. Every
	// Korean person has one. English info at
	// https://rootca.kisa.or.kr/kor/popup/foreigner_pop1_en.html
	// IsEncryptedWithKISAKey는 공인인증서로 암호화 되었는지를 뜻합니다.
	IsEncryptedWithKISAKey

	// HasSpareDigitalSig denotes if the file has a spare digital signature.
	// Yeah idk what this means either.
	// HasSpareDigitalSig는 파일이 전자 서명을 예비 저장하였는지를 뜻합니다.
	// 저도 뭔말인지 몰라요.
	HasSpareDigitalSig

	// HasKISADRM denotes if the file has a DRM with the aforementioned
	// KISAKey.
	// HasKISADRM은 공인인증서로 DRM 되었는지를 뜻합니다.
	HasKISADRM

	// HasCCL denotes if the file has a CCL (Creative Commons License).
	// HasCCL은 파일이 CCL(Creative Commons License)가 있는지를 뜻합니다.
	HasCCL

	// IsMobileOptimized denotes if the file is mobile optimized
	// IsMobileOptimized는 모바일 최적화가 되었는지를 뜻합니다.
	IsMobileOptimized

	// IsPrivateInfoProtected denotes if the file is a private info
	// protecting file. Idk what this means.
	// IsPrivateInfoProtected는 파일이 개인정보 보안 문서인지를 뜻합니다.
	// 네. 저도 몰라요.
	IsPrivateInfoProtected

	// IsModificationTracked denote if the file tracks modification
	// IsModificationTracked는 파일이 변경 추적을 하는지를 뜻합니다.
	IsModificationTracked

	// HasKOGLLicense denotes if the file has a KOGL license.
	// KOGL license is a media license created by the Korean Government.
	// There are 4 types of KOGL license.
	// KOGL type 1: Must have source.
	// KOGL type 2: Must have source and no commercial use.
	// KOGL type 3: Must have source and no modifications.
	// KOGL type 4: Must have source, no commerical use, no modifications
	// Info at kogl.co.kr (In Korean).
	// HasKOGLLicense는 KOGL 공공누리 저작권 문서가 있는지를 뜻합니다.
	// kogl.co.kr에 한글로 라이센스에 대해서 보실 수 있습니다.
	HasKOGLLicense

	// HasVideoControls denotes if the file has video controls.
	// HasVideoControls는 파일이 비디오 컨트롤이 있는지를 뜻합니다.
	HasVideoControls

	// HasChapterControlField denotes if the file has chapter controls
	// HasChapterControlField는 차례 필드 컬트롤이 있는지를 뜻합니다.
	HasChapterControlField

	//################################################
	//################################################
	// Bits 18-31 are reserved.
	// 18-31비트는 아직 사용하지 않습니다.
	//################################################
	//################################################
)

// SecondProperty represents the 4 byte bitflag that hwp50 uses for various settings
// SecondProperty는 hwp50 파일이 사용하는 4바이트 bitflag를 뜻합니다.
type SecondProperty [4]byte

const (
	// LicenseInfo denotes if the file has a CCL or KOGL license.
	// LicenseInfo는 파일이 CCL 또는 KOGL 라이센스가 있는지를 뜻합니다.
	LicenseInfo SeondProperty = 1 << iota

	// IsCopyProtected denotes if the file cannot be copied.
	// IsCopyProtected는 파일이 복제 제한 되어있는지를 뜻합니다.
	IsCopyProtected

	// IsAllowedToCopyWithoutModification denotes if the file can be copied
	// wihout any modifications to the original file.
	// IsAllowedToCopyWithoutModification는 동일 조건 하에 복제가
	// 허용되는지를 뜻합니다.
	IsAllowedToCopyWithoutModification

	//################################################
	//################################################
	// Bits 3-31 are reserved.
	// 3-31비트는 아직 사용하지 않습니다.
	//################################################
	//################################################
)

// FileHeader is the 256 const bit header for a hwp50 file.
// FileHeader는 hwp파일의 const 256 비트 헤더 입니다.
var FileHeader struct {
	// Sig represents that this file is an hwp file.
	// Sig는 이 파일이 hwp 파일이라는 것을 뜻합니다.
	Sig [32]byte

	/* Version is a 4 byte representation of the hwp50 file versioning
	 * Version은 4 바이트 hwp50 파일 버전은 나타네는 형식입니다.

	 * First Byte: Biggest file structure change.
	 * Incompatible if this byte is different
	 * 첫 번째 바이트: 가장 큰 구조 바뀜을 뜻합니다. 다르면 호환 불가.

	 * Second Byte: Lesser file structure change.
	 * Incompatible if this byte is different
	 * 두 번째 바이트: 두 번째로 큰 구조 바뀜을 뜻합니다. 다르면 호환 불가.

	 * Third Byte: Additional features added but backwards compatible.
	 * 셋째 바이트: 추가 기능이 더해짐을 뜻합니다. 달라도 호환 가능.

	 * Fourth Byte: Addtiional info in Record. Backwards compatible.
	 * 네 번째 바이트: 추가 정보가 Record에 더해짐을 뜻합니다. 달라도 호환 가능.
	 */
	Version [4]byte

	Fp FirstProperty

	Sp SecondProperty

	// EncryptVersion denotes the versioning of the encryption used
	// If the value is:
	// 0: none
	// 1: Hancom v2.5 and below
	// 2: Hancom v3.0 enchanced
	// 3: Hancom v3.0 Old
	// 4: Hancom v7.0 and newer
	// EncryptVersion는 암호화 버전을 뜻합니다.
	EncryptVersion [4]byte

	// KOGLCountry denotes which country the license is for
	// If the value is:
	// 6: A Korean KOGL License
	// 15: A US KOGL License
	KOGLCountry byte

	//################################################
	//################################################
	// Bytes 207-256 are reserved.
	// 207-256바이트는 아직 사용하지 않습니다.
	//################################################
	//################################################
}

// DeserializeFileHeader reads the file header information from the hwp50 file
// DeserializeFileHeader는 hwp50 파일에서 헤더를 읽습니다.
func (f *FileHeader) DeserializeFileHeader(r io.Reader) (
	fileVer FileVersion, err error) {

	// raw represents the const 256 bit header
	// raw는 const 256bit의 hwp 파일 헤더를 뜻합니다.
	raw := make([]byte, 256)

	r.Read(raw[:])

	// Check if Signature matches.
	// 서명 확인
	if raw[:31] != Signature {
		return fmt.Errorf("File corrupted or not a hwp file" +
			"이 파일은 손상되었거나 hwp 파일이 아닙니다.")
	}

	f.Sig = raw[:31]
	f.Version = raw[31 : 31+4]
	f.Fp = raw[35 : 35+4]
	f.Sp = raw[39 : 39+4]
	f.EncryptVersion = raw[44 : 44+4]
	f.KOGLCountry = raw[49 : 49+1]

	if raw[50:] != 0 {
		fmt.Println("WARNING:File corrupted or" +
			"new version of hwp available" +
			"주의: 파일이 손상되었거나 일부 지원하지 않는 새로운" +
			"버전의 hwp 파일입니다.")
	}
	return
}
