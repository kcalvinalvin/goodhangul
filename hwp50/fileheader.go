package hwp50

import (
	"encoding/binary"
	"fmt"
	"io"
)

// Signature is the first 32 bytes that the hwp50 file format is going
// to be.
//
// Signature는 hwp50 포맷의 첫 32 바이트를 뜻합니다.
var Signature = [32]byte{
	0x48, 0x57, 0x50, 0x20, 0x44, 0x6f, 0x63, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x20, 0x46, 0x69, 0x6c,
	0x65, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
}

// FileHeader is the 256 byte header for a hwp50 file. File properties are
// stored here.
//
// FileHeader는 hwp파일의 256 바이트 헤더 입니다. 파일 인식 정보가 저장되어
// 있습니다.
type FileHeader struct {
	// Sig represents that this file is an hwp file.
	//
	// Sig는 이 파일이 hwp 파일이라는 것을 뜻합니다.
	Sig [32]byte

	// Version is the next 4 bytes that hwp50 uses for file versioning
	// Determines if the file is compatible or not.
	//
	// Version은 Sig 다음의 파일의 버전을 나타내는 4바이트 입니다.
	// 파일이 호환 가능한지 여부를 나타냅니다.
	Version FileVersion

	// Fp is the next 4 bytes that uses bitflags for signaling
	// 18 properties.
	//
	// Fp 는 Version 다음의 파일의 속성을 알리는 18가지의 bitflag들
	// 입니다.
	Fp FirstProperty

	// Sp is the next 4 bytes after Fp that uses bitflags for signaling
	// 3 of these properties:
	// 1: License, CopyProtection, and if it's able to be copied without
	// modification.
	//
	// Sp 는 Fp다음의 3가지의 속성들을 알리는 bitflag입니다.
	// 라이센스, 복제 허용, 동일조건 하에 복제 허용을 위해 사용됩니다.
	Sp SecondProperty

	// EncryptVersion denotes the versioning of the encryption used
	// If the value is:
	// 0: none
	// 1: Hancom v2.5 and below
	// 2: Hancom v3.0 enchanced
	// 3: Hancom v3.0 Old
	// 4: Hancom v7.0 and newer
	//
	// EncryptVersion는 암호화 버전을 뜻합니다.
	// 값 의미:
	// 0: 없음
	// 1: Hancom v2.5 버전 이하
	// 2: Hancom v3.0 버전 enchanced
	// 3: Hancom v3.0 버전 old
	// 4: Hancom v7.0 버전 이후
	EncryptVersion uint32

	// KOGLCountry denotes which country the license is for
	// If the value is:
	// 6: A Korean KOGL License
	// 15: A US KOGL License
	//
	// KOGLCountry는 KOGL 라이센스가 어느 나라에 해당되는지를 뜻합니다.
	// 값 의미:
	// 6: 한국
	// 15: 미국
	KOGLCountry uint8

	//################################################
	//################################################
	// Bytes 207-256 are reserved.
	// 207-256바이트는 아직 사용하지 않습니다.
	//################################################
	//################################################
}

// DeserializeFileHeader reads the file header information from the hwp50 file
// NOTE: All variables are in little endian
//
// DeserializeFileHeader는 hwp50 파일에서 헤더를 읽습니다.
// NOTE: 모든 객체는 litten endian으로 저장되어 있습니다.
func (f *FileHeader) DeserializeFileHeader(r io.Reader) (err error) {
	// raw represents the hwp file's 256 byte header
	// raw는 256byte의 hwp 파일 헤더를 뜻합니다.
	raw := make([]byte, 256)

	_, err = r.Read(raw)
	if err != nil {
		return err
	}

	// Check if Signature matches.
	// 서명 확인
	var check [32]byte

	copy(check[:], raw[:32])

	if check != Signature {
		return fmt.Errorf("file corrupted or not a hwp file " +
			"이 파일은 손상되었거나 hwp 파일이 아닙니다")
	}

	// Set FileHeader Sig field
	// FileHeader Sig 필드 init
	copy(f.Sig[:], raw[:31])

	// Set FileHeader Version field
	// FileHeader Version 필드 init
	var versionBytes [4]byte

	copy(versionBytes[:], raw[31:31+4])

	// Set Version struct with 4 bytes of version (little endian)
	// 버전 struct
	err = f.Version.deserializeVersion(versionBytes)
	if err != nil {
		return err
	}

	fp := binary.LittleEndian.Uint32(raw[35 : 35+4])
	f.Fp = FirstProperty(fp) // typecast

	sp := binary.LittleEndian.Uint32(raw[39 : 39+4])
	f.Sp = SecondProperty(sp) // typecast

	f.EncryptVersion = binary.LittleEndian.Uint32(raw[44 : 44+4])

	f.KOGLCountry = raw[49]

	var (
		reserved [207]byte
		empty    [207]byte
	)

	copy(reserved[:], raw[50:])

	if reserved != empty {
		fmt.Println("WARNING:File corrupted or " +
			"new version of hwp available" +
			"주의: 파일이 손상되었거나 일부 지원하지 않는 새로운" +
			"버전의 hwp 파일입니다.")
		fmt.Printf("%x\n", reserved)
	}

	return nil
}

/*
 * Below represents the first 4 bytes that hwp50 uses for marking what
 * properties are available for the hwp file.
 * 및의 코드는 hwp50 형식의 파일의 속성을 나타내는 첫 4 바이트에 대한
 * 것입니다.
 *
 * Bits 0-17 are described below with methods.
 * 비트 0-17은 밑의 method에 나와 있습니다.
 *
 * Bits 18-31 are reserved.
 * 18-31비트는 아직 사용하지 않습니다.
 */

// FirstProperty represents the 4 byte bitflag that hwp50 uses for various settings
//
// FirstProperty는 hwp50 파일이 사용하는 4바이트 bitflag를 뜻합니다.
type FirstProperty uint32

// IsCompressed is the 0th of FirstProperty that denotes if the file
// is compressed
//
// IsCompressed는 파일이 압축되었는지를 나타내는 FirstProperty의 0번째 비트입니다
func (fp FirstProperty) IsCompressed() bool {
	return fp&(1<<0) == 1
}

// IsEncrypted is the 1st bit of FirstProperty that denotes if the
// file is encrypted
//
// IsEncrypted는 파일이 암호화 되었는지를 나타내는 FirstProperty의
// 1번째 비트입니다
func (fp FirstProperty) IsEncrypted() bool {
	return fp&(1<<1) == 1
}

// IsExported is the 2nd bit that denotes if the file is a file for distribution
// See directory distribution/
//
// IsExported는 파일이 배포용 문서인지를 나타내는 FirstProperty의 2번째
// 비트입니다.
// distribution/ 다이렉토리를 참고하시기 바랍니다.
func (fp FirstProperty) IsExported() bool {
	return fp&(1<<2) == 1
}

// HasScript is the 3rd bit that denotes if the file has scripts
//
// HasScript는 파일이 스크립트를 저장하는지를 나타내는 3번째 비트입니다
func (fp FirstProperty) HasScript() bool {
	return fp&(1<<3) == 1
}

// HasDRM is the 4th bit that denotes if the file is DRMed. Ew
//
// HasDRM은 파일이 DRM 걸려있는지를 나타내는 4번째 비트입니다.
func (fp FirstProperty) HasDRM() bool {
	return fp&(1<<4) == 1
}

// HasXMLTemplateStorage is the 5th bit that denotes if the file has
// XMLTemplate storage
//
// HasXMLTemplateStorage는 파일이 XMLTemplate storage가 있는지를
// 나타내는 5번째 비트입니다.
func (fp FirstProperty) HasXMLTemplateStorage() bool {
	return fp&(1<<5) == 1
}

// HasFileHistory is the 6th bit that denotes if the file history is included
//
// HasFileHistory는 파일이 이력을 저장했는지를 나타내는 6번째 비트입니다.
func (fp FirstProperty) HasFileHistory() bool {
	return fp&(1<<6) == 1
}

// HasDigitalSig is the 7th bit that denotes if the file has a digital signature
//
// HasDigitalSig는 전자 서명 정보가 있는지를 나타내는 7번째 비트입니다.
func (fp FirstProperty) HasDigitalSig() bool {
	return fp&(1<<7) == 1
}

// IsEncryptedWithKISAKey is the 8th bit that  denotes if the file is
// encrypted with a KISA key. A KISA key is a key that the Korean government
// authorized a private company for the safekeeping of all bank accounts in
// Korea. Every Korean person has one. English info at
// https://rootca.kisa.or.kr/kor/popup/foreigner_pop1_en.html
//
// IsEncryptedWithKISAKey는 공인인증서로 암호화 되었는지를 뜻하는 8번째
// 비트입니다.
func (fp FirstProperty) IsEncryptedWithKISAKey() bool {
	return fp&(1<<8) == 1
}

// HasSpareDigitalSig is the 9th bit that denotes if the file has a spare
// digital signature. Yeah idk what this means either.
//
// HasSpareDigitalSig는 파일이 전자 서명을 예비 저장하였는지를 뜻하는 9번째
// 비트입니다.
// 저도 뭔말인지 몰라요.
func (fp FirstProperty) HasSpareDigitalSig() bool {
	return fp&(1<<9) == 1
}

// HasKISADRM is the 10th bit that denotes if the file has a DRM with the
// aforementioned KISAKey.
//
// HasKISADRM은 공인인증서로 DRM 되었는지를 뜻하는 10번째 비트입니다.
func (fp FirstProperty) HasKISADRM() bool {
	return fp&(1<<10) == 1
}

// HasCCL is the 11th bit that denotes if the file has a CCL
// (Creative Commons License).
//
// HasCCL은 파일이 CCL(Creative Commons License)가 있는지를 뜻하는 11번째
// 비트입니다.
func (fp FirstProperty) HasCCL() bool {
	return fp&(1<<11) == 1
}

// IsMobileOptimized is the 12th bit that denotes if the file is mobile optimized
//
// IsMobileOptimized는 모바일 최적화가 되었는지를 뜻하는 13번째 비트입니다.
func (fp FirstProperty) IsMobileOptimized() bool {
	return fp&(1<<12) == 1
}

// IsPrivateInfoProtected is the 13th bit that denotes if the file is a
// private info protecting file. Idk what this means.
//
// IsPrivateInfoProtected는 파일이 개인정보 보안 문서인지를 뜻하는 14번째
// 비트입니다.
// 네. 저도 뭔말인지 몰라요.
func (fp FirstProperty) IsPrivateInfoProtected() bool {
	return fp&(1<<13) == 1
}

// IsModificationTracked is the 14th bit that denotes if the file tracks
// modification
//
// IsModificationTracked은 파일이 변경 추적을 하는지를 뜻하는 14번째 비트입니다
func (fp FirstProperty) IsModificationTracked() bool {
	return fp&(1<<14) == 1
}

// HasKOGLLicense is the 15th bit that denotes if the file has a KOGL license.
// KOGL license is a media license created by the Korean Government.
// There are 4 types of KOGL license.
// KOGL type 1: Must have source.
// KOGL type 2: Must have source and no commercial use.
// KOGL type 3: Must have source and no modifications.
// KOGL type 4: Must have source, no commercial use, no modifications
// Info at kogl.co.kr (In Korean).
//
// HasKOGLLicense는 KOGL 공공누리 저작권 문서가 있는지를 뜻하는 15번째
// 비트입니다. kogl.co.kr에서 한글로 라이센스에 대해서 보실 수 있습니다.
func (fp FirstProperty) HasKOGLLicense() bool {
	return fp&(1<<15) == 1
}

// HasVideoControls is the 16th bit that denotes if the file has video controls.
//
// HasVideoControls는 파일이 비디오 컨트롤이 있는지를 뜻하는 16번째 비트입니다.
func (fp FirstProperty) HasVideoControls() bool {
	return fp&(1<<16) == 1
}

// HasChapterControlField is the 17th bit that denotes if the file has
// chapter controls
//
// HasChapterControlField는 차례 필드 컬트롤이 있는지를 뜻하는 17번째
// 비트입니다.
func (fp FirstProperty) HasChapterControlField() bool {
	return fp&(1<<17) == 1
}

/*
 * Below represents the second 4 bytes that hwp50 uses for marking what
 * properties are available for the hwp file.
 * 및의 코드는 hwp50 형식의 파일의 속성을 나타내는 첫 4 바이트에 대한
 * 것입니다.
 *
 * Bits 0-2 are described below with methods.
 * 비트 0-2은 밑의 method에 나와 있습니다.
 *
 * Bits 3-31 are reserved.
 * 3-31비트는 아직 사용하지 않습니다.
 */

// SecondProperty represents the 4 byte bitflag that hwp50 uses for various settings
//
// SecondProperty는 hwp50 파일이 사용하는 4바이트 bitflag를 뜻합니다.
type SecondProperty uint32

// HasLicenseInfo denotes if the file has a CCL or KOGL license.
//
// HasLicenseInfo는 파일이 CCL 또는 KOGL 라이센스가 있는지를 뜻합니다.
func (sp SecondProperty) HasLicenseInfo() bool {
	return sp&(1<<0) == 1
}

// IsCopyProtected denotes if the file cannot be copied.
//
// IsCopyProtected는 파일이 복제 제한 되어있는지를 뜻합니다.
func (sp SecondProperty) IsCopyProtected() bool {
	return sp&(1<<1) == 1
}

// IsAllowedToCopyWithoutModification denotes if the file can be copied
// wihout any modifications to the original file.
//
// IsAllowedToCopyWithoutModification는 동일 조건 하에 복제가
// 허용되는지를 뜻합니다.
func (sp SecondProperty) IsAllowedToCopyWithoutModification() bool {
	return sp&(1<<2) == 1
}

// FileVersion is a 4 byte representation of the hwp50 file versioning
//
// FileVersion은 4 바이트 hwp50 파일 버전은 나타네는 형식입니다.
type FileVersion struct {
	// First Byte: Biggest file structure change.
	// Incompatible if this byte is different
	//
	// 첫 번째 바이트: 가장 큰 구조 바뀜을 뜻합니다. 다르면 호환 불가.
	Major uint8

	// Second Byte: Lesser file structure change.
	// Incompatible if this byte is different
	//
	// 두 번째 바이트: 두 번째로 큰 구조 바뀜을 뜻합니다. 다르면 호환 불가
	Minor uint8

	// Third Byte: Additional features added but backwards compatible.
	//
	// 셋째 바이트: 추가 기능이 더해짐을 뜻합니다. 달라도 호환 가능.
	Micro uint8

	// Fourth Byte: Addtiional info in Record. Backwards compatible.
	//
	// 네 번째 바이트: 추가 정보가 Record에 더해짐을 뜻합니다. 달라도
	// 호환 가능.
	Extra uint8
}

// deserializeVersion deserializes version info from the given 4 bytes
// Argument b should be in little endian
//
// deserializeVersion은 4 바이트 array의 버전 정보를 deserialize 합니다.
// Argument b는 little endian 포맷으로 주어야 합니다.
func (fv *FileVersion) deserializeVersion(b [4]byte) error {
	fv.Major = b[2]
	fv.Minor = b[3]
	fv.Micro = b[0]
	fv.Extra = b[1]

	return nil
}
