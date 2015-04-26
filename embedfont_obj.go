package gopdf

import (
	"bytes"
	"errors"
	"io/ioutil"
	"strconv"
	"strings"
)

type EmbedFontObj struct {
	buffer                 bytes.Buffer
	Data                   string
	zfontpath, ttffontpath string
	font                   IFont
}

var ERROR_NO_FONT_PATH = errors.New("no font path (zfontpath or ttffontpath)")

func (me *EmbedFontObj) Init(funcGetRoot func() *GoPdf) {
}

func (me *EmbedFontObj) GetFontBytes() ([]byte, error) {

	if strings.TrimSpace(me.zfontpath) != "" {
		b, err := ioutil.ReadFile(me.zfontpath)
		if err != nil {
			return nil, err
		}
		return b, nil
	} else if strings.TrimSpace(me.ttffontpath) != "" {
		/*var buff bytes.Buffer
		gzipwriter := zlib.NewWriter(&buff)

		fontbytes, err := ioutil.ReadFile(me.ttffontpath)
		if err != nil {
			return nil, err
		}
		_, err = gzipwriter.Write(fontbytes)
		if err != nil {
			return nil, err
		}
		defer gzipwriter.Close()
		return buff.Bytes(), nil*/
		b, err := ioutil.ReadFile("/home/oneplus/GOPATH/src/github.com/signintech/gopdf/res/fonts/Loma.z")
		if err != nil {
			return nil, err
		}
		return b, nil
	}

	return nil, ERROR_NO_FONT_PATH
}

func (me *EmbedFontObj) Build() {
	b, err := me.GetFontBytes()
	if err != nil {
		return
	}
	me.buffer.WriteString("<</Length " + strconv.Itoa(len(b)) + "\n")
	me.buffer.WriteString("/Filter /FlateDecode\n")
	me.buffer.WriteString("/Length1 " + strconv.Itoa(me.font.GetOriginalsize()) + "\n")
	me.buffer.WriteString(">>\n")
	me.buffer.WriteString("stream\n")
	me.buffer.Write(b)
	me.buffer.WriteString("\nendstream\n")
}

func (me *EmbedFontObj) GetType() string {
	return "EmbedFont"
}

func (me *EmbedFontObj) GetObjBuff() *bytes.Buffer {
	return &(me.buffer)
}

func (me *EmbedFontObj) SetFont(font IFont, zfontpath string) {
	me.font = font
	me.zfontpath = zfontpath
}

func (me *EmbedFontObj) SetFontTtf(font IFont, ttffontpath string) {
	me.font = font
	me.ttffontpath = ttffontpath
}
