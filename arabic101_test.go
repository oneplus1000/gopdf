package gopdf

import (
	"log"
	"testing"
)

//https://www.microsoft.com/typography/otspec/gsub.htm#EX1

func TestArabic(t *testing.T) {
	pdf := GoPdf{}
	pdf.Start(Config{PageSize: Rect{W: 595.28, H: 841.89}}) //595.28, 841.89 = A4
	pdf.AddPage()
	/*err := pdf.AddTTFFont("HDZB_5", "./res/BNazanin.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}*/
	err := pdf.AddTTFFontWithOption("HDZB_5", "./test/res/NotoNaskhArabic-Regular.ttf", TtfOption{
		UseKerning: false,
	})
	if err != nil {
		log.Print(err.Error())
		return
	}

	err = pdf.SetFont("HDZB_5", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}
	pdf.Cell(nil, ArabicText("احمد"))
	pdf.WritePdf("./test/out/test1.pdf")
}
