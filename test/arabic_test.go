package test

import (
	"log"
	"testing"

	"github.com/signintech/gopdf"
)

//https://www.microsoft.com/typography/otspec/gsub.htm#EX1

func TestArabic(t *testing.T) {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 595.28, H: 841.89}}) //595.28, 841.89 = A4
	pdf.AddPage()
	/*err := pdf.AddTTFFont("HDZB_5", "./res/BNazanin.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}*/
	err := pdf.AddTTFFontWithOption("HDZB_5", "./res/NotoNaskhArabic-Regular.ttf", gopdf.TtfOption{
		UseKerning: true,
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
	pdf.Cell(nil, Reverse("احمد"))
	pdf.WritePdf("./out/test1.pdf")
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
