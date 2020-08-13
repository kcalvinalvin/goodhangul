package hwp50

type wchar uint16

// enum for tag
// Either this is really bad engineering or these fuckers
// aren't documenting everything.
const (
	isPageFirstLine = iota
	isColumnFirstLine
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	isEmptySegment
	isFirstSegment
	isLastSegment
	hasAutoHypehn
	hasIndent
	hasParaHeadShape
	_
	_
	_
	_
	_
	_
	_
	_
	_
	// In the documentation, the reason for this bit is for
	// "Ease of implemenation". lmao
	badEngineeringCoveredWithAnExcuse
)

// BodyText is the main stream of bytes for an hwp file. Includes information
// about charts, images, etc.
type BodyText struct {
	ParaHeader              ParaHeader
	ParaChar                []wchar
	ParaCharShape           []paraCharShape
	ParaLineSeg             paraLineSeg
	ParaRangeTag            []byte
	CtrlHeader              [4]byte
	ListHeader              [6]byte
	PageDef                 [40]byte
	FootnoteShape           [30]byte
	PageBorderFill          [14]byte
	ShapeComponent          [4]byte
	Table                   []byte
	ShapeComponentLine      [20]byte
	ShapeComponentRectangle [9]byte
	ShapeComponentEllipse   [60]byte
	ShapeComponentArc       [25]byte
	ShapeComponentPolygon   []byte
	ShapeComponentCurve     []byte
	ShapeComponentOLE       [26]byte
	ShapeComponentPicture   []byte
	CtrlData                []byte
	EqEdit                  []byte
	ShapeComponentTextArt   []byte
	FormObject              []byte
	MemoShape               [22]byte
	MemoList                [4]byte
	ChartData               [2]byte
	VideaData               []byte
	ShapeComponentUnknown   [36]byte
}

type ParaHeader struct {
	// NChars is the number of chars in this paragraph
	NChars uint32

	ControlMask uint32

	ParaShapeID uint16

	ParaStyleID uint8

	SecSplitInfo uint8

	CharShapeInfo uint16

	RangeTagInfo uint16

	LineAlignInfo uint16

	SectionInsID uint32

	// Only applicable from v5.0.3.2
	TrackChange uint16
}

// paraCharShape determines the shape of a char
type paraCharShape struct {
	// pos is the count where the CharShape changes
	pos uint32
	// shapeID is the shape of the ParaChar
	shapeID uint32
}

type paraLineSeg struct {
	textStartPos   uint32
	lineVertPos    uint32
	lineHeight     int32
	textPosHeight  int32
	lenToBaseline  int32 // Measure from lineVertPos
	lineWidth      int32
	columnStartPos int32
	segmentWidth   int32
	tag            uint32
}

type tag struct {
}
