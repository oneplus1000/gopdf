package core

import "testing"

func TestParseGPOS(t *testing.T) {
	ttfParser := TTFParser{}
	err := ttfParser.Parse("/Users/oneplus/Code/Play/eventsanook/assets/fonts/NotoSansThai-SemiBold.ttf")
	if err != nil {
		t.Error(err)
	}
}
