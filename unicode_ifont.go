package gopdf

import (
	"fmt"

	"github.com/signintech/gopdf/fontmaker/core"
)

//"fmt"
//iconv "github.com/djimenez/iconv-go"

type UnicodeIFont struct {
	family string
	desc   []FontDescItem
}

func (me *UnicodeIFont) Init() {

}

func (me *UnicodeIFont) GetType() string {
	return ""
}

func (me *UnicodeIFont) GetName() string {
	return me.family
}

func (me *UnicodeIFont) GetDesc() []FontDescItem {
	return me.desc
}
func (me *UnicodeIFont) GetUp() int {
	return 0
}

func (me *UnicodeIFont) GetUt() int {
	return 0
}
func (me *UnicodeIFont) GetCw() FontCw {
	var fcw FontCw
	return fcw
}

func (me *UnicodeIFont) GetEnc() string {
	return ""
}

func (me *UnicodeIFont) GetDiff() string {
	return ""
}

func (me *UnicodeIFont) GetOriginalsize() int {
	return 0
}

func (me *UnicodeIFont) SetFamily(family string) {
	me.family = family
}

func (me *UnicodeIFont) GetFamily() string {
	return ""
}

func (me *UnicodeIFont) SetTtf(ttffile string) error {
	var fontmaker core.FontMaker
	info, err := fontmaker.GetInfoFromTrueType(ttffile, nil)
	if err != nil {
		return err
	}

	// Ascent
	ascender, err := info.GetInt64("Ascender")
	if err != nil {
		return err
	}

	// Descent
	descender, err := info.GetInt64("Descender")
	if err != nil {
		return err
	}

	// CapHeight
	capHeight, err := info.GetInt64("CapHeight")
	if err == core.ERROR_NO_KEY_FOUND {
		capHeight = ascender
	} else if err != nil {
		return err
	}

	flags := 0
	isFixedPitch, err := info.GetBool("IsFixedPitch")
	if err != nil {
		return err
	}
	if isFixedPitch {
		flags += 1 << 0
	}
	flags += 1 << 5
	italicAngle, err := info.GetInt64("ItalicAngle")
	if italicAngle != 0 {
		flags += 1 << 6
	}

	fbb, err := info.GetInt64s("FontBBox")
	if err != nil {
		return err
	}

	// StemV
	stdVW, err := info.GetInt64("StdVW")
	issetStdVW := false
	if err == nil {
		issetStdVW = true
	} else if err == core.ERROR_NO_KEY_FOUND {
		issetStdVW = false
	} else {
		return err
	}

	bold, err := info.GetBool("Bold")
	if err != nil {
		return err
	}

	stemv := int64(0)
	if issetStdVW {
		stemv = stdVW
	} else if bold {
		stemv = 120
	} else {
		stemv = 70
	}

	missingWidth, err := info.GetInt64("MissingWidth")
	if err != nil {
		return err
	}

	desc := make([]FontDescItem, 8)
	desc[0] = FontDescItem{Key: "Ascent", Val: fmt.Sprintf("%d", ascender)}
	desc[1] = FontDescItem{Key: "Descent", Val: fmt.Sprintf("%d", descender)}
	desc[2] = FontDescItem{Key: "CapHeight", Val: fmt.Sprintf("%d", capHeight)}
	desc[3] = FontDescItem{Key: "Flags", Val: fmt.Sprintf("%d", flags)}
	desc[4] = FontDescItem{Key: "FontBBox", Val: fmt.Sprintf("[%d %d %d %d]", fbb[0], fbb[1], fbb[2], fbb[3])}
	desc[5] = FontDescItem{Key: "ItalicAngle", Val: fmt.Sprintf("%d", italicAngle)}
	desc[6] = FontDescItem{Key: "StemV", Val: fmt.Sprintf("%d", stemv)}
	desc[7] = FontDescItem{Key: "MissingWidth", Val: fmt.Sprintf("%d", missingWidth)}
	me.desc = desc

	return nil
}
