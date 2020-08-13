package hwp50

// Hwp is a representation of the entire Hwp file
// This can be parsed into xml
type Hwp struct {
	DocInfo               DocInfo
	PrvText               []byte
	Scripts               []byte
	DefaultJScript        []byte
	BodyText              []BodyText
	PrvImage              []byte
	DocOptions            []byte
	FileHeader            FileHeader
	HwpSummaryInformation []byte
}

// TODO
func (hwp *Hwp) ToXML() {
}

// TODO
func (hwp *Hwp) ToPDF() {
}
