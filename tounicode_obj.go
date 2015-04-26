package gopdf

import (
	"bytes"
	"strconv"
)

type ToUnicodeObj struct {
	font   IFont
	buffer bytes.Buffer
}

func (me *ToUnicodeObj) Init(func() *GoPdf) {

}

func (me *ToUnicodeObj) GetType() string {
	return ""
}
func (me *ToUnicodeObj) GetObjBuff() *bytes.Buffer {

	return &(me.buffer)
}

//สร้าง ข้อมูลใน pdf
func (me *ToUnicodeObj) Build() {
	b := me.content()
	me.buffer.WriteString("<</Length " + strconv.Itoa(len(b)) + "/Filter/FlateDecode>>\n")
	me.buffer.WriteString("stream\n")
	me.buffer.Write(b)
	me.buffer.WriteString("\nendstream\n")
}

func (me *ToUnicodeObj) SetFont(font IFont) {
	me.font = font
}

func (me *ToUnicodeObj) GetFont() IFont {
	return me.font
}

func (me *ToUnicodeObj) content() []byte {
	var buff bytes.Buffer
	buff.WriteString("/CIDInit/ProcSet findresource begin\n")
	buff.WriteString("12 dict begin\n")
	buff.WriteString("begincmap\n")
	buff.WriteString("/CIDSystemInfo<<\n")
	buff.WriteString("/Registry (Adobe)\n")
	buff.WriteString("/Ordering (UCS)\n")
	buff.WriteString("/Supplement 0\n")
	buff.WriteString(">> def\n")
	buff.WriteString("/CMapName/Adobe-Identity-UCS def\n")
	buff.WriteString("/CMapType 2 def\n")
	buff.WriteString("1 begincodespacerange\n")
	buff.WriteString("<00> <FF>\n")
	buff.WriteString("endcodespacerange\n")
	buff.WriteString("2 beginbfchar\n")
	buff.WriteString("<01> <0E01>\n")
	buff.WriteString("<02> <0E02>\n")
	buff.WriteString("<03> <0E04>\n")
	buff.WriteString("<04> <0E07>\n")
	buff.WriteString("endbfchar\n")
	buff.WriteString("endcmap\n")
	buff.WriteString("CMapName currentdict /CMap defineresource pop\n")
	buff.WriteString("end\n")
	buff.WriteString("end\n")
	return buff.Bytes()
}
