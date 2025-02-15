package core

import (
	"path/filepath"
	"testing"
)

const fontRootPath = "/Users/oneplus/Code/Play/test_font"

func TestParseGPOSNoto(t *testing.T) {
	ttfParser := TTFParser{}
	err := ttfParser.Parse(filepath.Join(fontRootPath, "NotoSansThai-SemiBold.ttf"))
	if err != nil {
		t.Error(err)
	}
}

func TestParseGPOSLoma(t *testing.T) {
	ttfParser := TTFParser{}
	err := ttfParser.Parse(filepath.Join(fontRootPath, "Loma-Bold.ttf"))
	if err != nil {
		t.Error(err)
	}
}
