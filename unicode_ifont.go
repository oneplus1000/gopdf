package gopdf

//"fmt"
//iconv "github.com/djimenez/iconv-go"

type UnicodeIFont struct {
	family string
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
	return nil
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
