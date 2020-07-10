package hwp50

// DocInfo saves information about font, tab, styling, etc
//
// DocInfo는 본문의 글꼴, 글자 속성, 탭, 스타일등의 정보를 담고 있습니다.
type DocInfo struct {
	// DocumentProperites is the current hwp file's properties
	//
	// DocumentProperites는 현 hwp 파일의 속성을 나타냅니다.
	DocumentProperites DocumentProperites

	// IDMappings
	IDMappings IDMappings

	BinData []byte

	FaceName            []byte
	BorderFill          []byte
	CharShape           [72]byte
	TabDef              [14]byte
	Numbering           []byte
	Bullet              [10]byte
	ParagraphShape      [54]byte
	Style               []byte
	MemoShape           [22]byte
	TrackChangeAuthor   []byte
	FirstTrackChange    []byte
	DocData             []byte
	ForbiddenChar       []byte
	CompatibleDocument  [4]byte
	LayoutCompatibility [20]byte
	DistributeDocData   [256]byte
	SecondTrackChange   [1032]byte
}

// DocumentProperites is the property of the current hwp 5.0 file
//
// DocumentProperites는 현재 읽는 hwp 5.0 파일의 속성입니다.
type DocumentProperites struct {
	// SectionNum represents how many sections there are
	//
	// SectionNum은 구역 개수를 의미합니다.
	SectionNum uint16

	// PageStartNum refers to when the page starts.
	//
	// PageStartNum은 페이지가 언제 시작하는지를 뜻합니다.
	PageStartNum uint16

	// FootNoteStartNum refers to where the first footnote is.
	//
	// FootNoteStartNum은 각주 시작 번호를 뜻합니다.
	FootNoteStartNum uint16

	// EndNoteStartNum refers to where the first endnote is.
	//
	// EndNoteStartNum은 미주 시작 번호를 뜻합니다.
	EndNoteStartNum uint16

	// PictureStartNum refers to where the first picture is.
	//
	// PictureStartNum은 그림 시작 번호를 뜻합니다.
	PictureStartNum uint16

	// ChartStartNum refers to where the first chart is.
	//
	// ChartStartNum은 표 시작 번호를 뜻합니다.
	ChartStartNum uint16

	// EquationStartNum refers to where the first equation is.
	//
	// EquationStartNum은 수식 시작 번호를 뜻합니다.
	EquationStartNum uint16

	// ListID is the current list's ID within the current doc.
	//
	// ListID는 현 문서 내 리스트 ID입니다.
	ListID uint32

	// ParagraphID is the current paragraph's ID within the current doc.
	//
	// ParagraphID는 현 문서 내 문단 ID입니다.
	ParagraphID uint32

	// WordUnitLocationInParagraph refers to where the word is within the
	// Paragraph
	//
	// WordUnitLocationInParagraph는 현 문단 내 글자 단위 위치를
	// 의미합니다.
	WordUnitLocationInParagraph uint32
}

// IDMappings is a pointer to the elements
//
// IDMappings은 이
type IDMappings struct {
	BinData int32

	KoreanFont int32

	EnglishFont int32

	HanjaFont int32

	JapaneseFont int32

	EtcFont int32

	CharacterFont int32

	UserFont int32

	BorderAndBackground int32

	LetterShape int32

	TabDef int32

	ParagraphNum int32

	Style int32

	// Only relevant for hwp 5.0.2.1 and up
	MemoShape int32

	// Only relevant for hwp 5.0.3.2 and up
	TrackChanges int32

	// Only relevant for hwp 5.0.3.2 and up
	TrackChangeAuthor int32
}

// BinData stores data about pictures, OLE, etc.
type BinData struct {
	// Stores information on whehter BinData has compression,
	// images, and whether it was accssed before
	Property uint16

	// length of the absolute path
	// Only relevant when Type is "LINK"
	AbsLen uint16

	// location of the absolute path
	// ABSLen * 2 is the size
	// Only relevant when Type is "LINK"
	AbsLoc []byte

	// length of the relative path
	// Only relevant when Type is "LINK"
	RelLen uint16

	// location of the relative path
	// RelLen * 2 is the size
	// Only relevant when Type is "LINK"
	RelLoc []byte

	// Binary data's stored ID in BINDATASTORAGE
	// Only relevant when type is "EMBEDDING" or "STORAGE"
	BinDataID uint16

	// Only relevant when type is "EMBEDDING" or "STORAGE"
	BinDataNameLen uint16

	// Extension type
	// BinDataNameLen * 2 is the size
	// Only relevant when type is "EMBEDDING"
	// Available picture types: jpg, bmp, gif
	// Available OLE type: ole
	ExtType []byte
}

func (bd *BinData) HasExternalPicture() bool {
	// turn on bits 0~3
	mask := uint16(15)

	// take only the 1~3 bits
	value := mask & bd.Property

	return value == 0
}

func (bd *BinData) HasEmbeddedPicture() bool {
	// turn on bits 0~3
	mask := uint16(15)

	// take only the 1~3 bits
	value := mask & bd.Property

	return value == 1
}

func (bd *BinData) HasStorage() bool {
	// turn on bits 0~3
	mask := uint16(15)

	// take only the 1~3 bits
	value := mask & bd.Property

	return value == 2
}

func (bd *BinData) IsDefultStorageMode() bool {
	// turn on bits 4~5
	mask := uint16(3 << 4)

	// take only the 1~3 bits
	value := mask & bd.Property

	return value == 0
}

func (bd *BinData) AlwaysCompress() bool {
	// turn on bits 4~5
	mask := uint16(3 << 4)

	// take only the 1~3 bits
	value := mask & bd.Property

	return value == 1
}

func (bd *BinData) NeverCompress() bool {
	// turn on bits 4~5
	mask := uint16(3 << 4)

	// take only the 1~3 bits
	value := mask & bd.Property

	return value == 2
}

func (bd *BinData) NeverAccessed() bool {
	// turn on bits 8~9
	mask := uint16(3 << 8)

	// take only the 1~3 bits
	value := mask & bd.Property

	return value == 0
}

func (bd *BinData) AccessSucessful() bool {
	// turn on bits 8~9
	mask := uint16(3 << 8)

	// take only the 1~3 bits
	value := mask & bd.Property

	return value == 1
}

func (bd *BinData) AccessFailed() bool {
	// turn on bits 8~9
	mask := uint16(3 << 8)

	// take only the 1~3 bits
	value := mask & bd.Property

	return value == 2
}

func (bd *BinData) AccessFailedAndErrorIgnored() bool {
	// turn on bits 8~9
	mask := uint16(3 << 8)

	// take only the 1~3 bits
	value := mask & bd.Property

	return value == 3
}

func (bd *BinData) GetAbsLocSize() uint16 {
	return bd.AbsLen * 2
}

func (bd *BinData) GetRelLocSize() uint16 {
	return bd.RelLen * 2
}

func (bd *BinData) GetExtTypeSize() uint16 {
	return bd.BinDataNameLen * 2
}

// FaceName stores information about the font used
type FaceName struct {
	//
	Property byte

	// Length of the FName
	FNameLen uint16

	// Name of the font used
	FName []byte

	// Length of the substitute FName when FName is not
	// available
	SubFNameLen uint16

	// Name of the substitute font
	// Used when FName is not available
	SubFName []byte

	// Property of the face used
	FaceType [10]byte

	// Default face name length
	DefaultFNameLen uint16

	// Name of the default font
	DefaultFName []byte
}

func (fn *FaceName) GetFaceNameLen() uint16 {
	return fn.FNameLen
}

func (fn *FaceName) GetSubFNameLen() uint16 {
	return fn.SubFNameLen
}

func (fn *FaceName) GetDefaultFNameLen() uint16 {
	return fn.DefaultFNameLen
}

type BorderFill struct {
	Property uint16

	BorderType [4]uint8

	BorderThickness [4]uint8

	BorderColor [4]uint32

	DiagonalLineType uint8

	DiagonalLineThickness uint8

	DiagonalLineColor uint32

	FillInfo []byte
}

func (bf *BorderFill) Has3DEffect() bool {
	// turn on the 0th bit
	mask := uint16(1)

	value := mask & bf.Property

	return value == 1
}

func (bf *BorderFill) HasShadows() bool {
	// turn on the 1st bit
	mask := uint16(2)
	value := mask & bf.Property
	return value == 1
}

func (bf *BorderFill) GetSlashShape() uint16 {
	// turn on bits 2~4
	mask := uint16(7 << 2)
	value := mask & bf.Property
	return value
}

func (bf *BorderFill) GetBackSlashShape() uint16 {
	// turn on bits 5~7
	mask := uint16(7 << 5)
	value := mask & bf.Property
	return value
}
func (bf *BorderFill) IsSlashBent() uint16 {
	// turn on bits 8~9
	mask := uint16(3 << 9)
	value := mask & bf.Property
	return value
}
func (bf *BorderFill) IsBackSlashBent() bool {
	// turn on bit 10
	mask := uint16(1 << 10)
	value := mask & bf.Property
	return value == 1
}
func (bf *BorderFill) IsSlash180flipped() bool {
	mask := uint16(1 << 11)
	value := mask & bf.Property
	return value == 1
}

func (bf *BorderFill) IsBackSlash180flipped() bool {
	mask := uint16(1 << 12)
	value := mask & bf.Property
	return value == 1
}

func (bf *BorderFill) HasCenterLine() bool {
	mask := uint16(1 << 13)
	value := mask & bf.Property
	return value == 1
}
