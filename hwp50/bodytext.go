package hwp50

// BodyText is the main stream of bytes for an hwp file. Includes information
// about charts, images, etc.
type BodyText struct {
	ParaHeader              [22]byte
	ParaText                []byte
	ParaCharShape           []byte
	ParaLineSeg             []byte
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
