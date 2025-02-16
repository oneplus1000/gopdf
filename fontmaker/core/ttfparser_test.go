package core

import (
	"path/filepath"
	"testing"
)

// TODO: เปลี่ยนให้ไป load จาก internal แทน (เพื่อให้สามารถทำการ test ได้ และไม่ติดเรื่องลิขสิทธิ์)
const fontTestDir = "/Users/oneplus/Code/Play/test_font"

func TestParseNotoSansThai(t *testing.T) {
	p := TTFParser{}
	err := p.Parse(filepath.Join(fontTestDir, "NotoSansThai-SemiBold.ttf"))
	if err != nil {
		t.Errorf("Error parsing font: %v", err)
	}
	if p.gpos == nil {
		t.Errorf("GPOS table not found")
	}

	//correct value
	var version uint = 0x00010000
	var scriptListCount int = 6
	var scriptTagtag0 = "DFLT"
	var scriptTagtag1 = "cyrl"

	if p.gpos.Version() != version {
		t.Errorf("GPOS table version not match")
	}
	if int(p.gpos.ScriptList.ScriptCount) != scriptListCount {
		t.Errorf("GPOS table script count not match")
	}

	if p.gpos.ScriptList.ScriptRecords[0].ScriptTagString() != scriptTagtag0 {
		x := p.gpos.ScriptList.ScriptRecords[0].ScriptTagString()
		t.Errorf("GPOS table script tag not match '%s'", x)
	}

	if p.gpos.ScriptList.ScriptRecords[0].ScriptTable.LangSysCount != 0 {
		t.Errorf("GPOS table script table lang sys count not match")
	}

	for i, scriptRecord := range p.gpos.ScriptList.ScriptRecords {
		if scriptRecord.ScriptTable.LangSysCount != 0 {
			t.Errorf("GPOS table script table lang sys count not match at %d %d", i, scriptRecord.ScriptTable.LangSysCount)
		}
		if scriptRecord.ScriptTable.DefaultLangSysOffset != 0 {
			if scriptRecord.ScriptTable.DefaultLangSys.LookupOrderOffset != 0 {
				t.Errorf("GPOS table script table default lang sys lookup order offset not match at %d %d", i, scriptRecord.ScriptTable.DefaultLangSys.LookupOrderOffset)
			}
			if scriptRecord.ScriptTable.DefaultLangSys.ReqFeatureIndex != 0xFFFF {
				t.Errorf("GPOS table script table default lang sys req feature index not match at %d %d", i, scriptRecord.ScriptTable.DefaultLangSys.ReqFeatureIndex)
			}
		}
	}

	if p.gpos.ScriptList.ScriptRecords[1].ScriptTagString() != scriptTagtag1 {
		x := p.gpos.ScriptList.ScriptRecords[1].ScriptTagString()
		t.Errorf("GPOS table script tag not match '%s'", x)
	}

}

func TestParseLomaBold(t *testing.T) {
	p := TTFParser{}
	err := p.Parse(filepath.Join(fontTestDir, "Loma-Bold.ttf"))
	if err != nil {
		t.Errorf("Error parsing font: %v", err)
	}
	if p.gpos == nil {
		t.Errorf("GPOS table not found")
	}

	if len(p.gpos.ScriptList.ScriptRecords) != 2 {
		t.Errorf("GPOS table script count not match")
	}

	for i, scriptRecord := range p.gpos.ScriptList.ScriptRecords {
		if i == 0 {
			if scriptRecord.ScriptTable.LangSysCount != 0 {
				t.Errorf("GPOS table script table lang sys count not match at %d %d", i, scriptRecord.ScriptTable.LangSysCount)
			}
		} else if i == 1 {
			if scriptRecord.ScriptTable.LangSysCount != 3 || len(scriptRecord.ScriptTable.LangSysRecords) != 3 {
				t.Errorf("GPOS table script table lang sys count not match at %d %d", i, scriptRecord.ScriptTable.LangSysCount)
			}
			for j, rc := range scriptRecord.ScriptTable.LangSysRecords {
				if j == 0 {
					if rc.LangSysTagString() != "KUY " {
						t.Errorf("GPOS table script table lang sys tag not match at %d %s", i, rc.LangSysTagString())
					}
					if len(rc.LangSys.FeatureIndex) != 2 {
						t.Errorf("GPOS table script table lang sys feature index not match at %d %d", i, len(rc.LangSys.FeatureIndex))
					}
				}
			}
		}
		if scriptRecord.ScriptTable.DefaultLangSysOffset != 0 {
			if scriptRecord.ScriptTable.DefaultLangSys.LookupOrderOffset != 0 {
				t.Errorf("GPOS table script table default lang sys lookup order offset not match at %d %d", i, scriptRecord.ScriptTable.DefaultLangSys.LookupOrderOffset)
			}
			if scriptRecord.ScriptTable.DefaultLangSys.ReqFeatureIndex != 0xFFFF {
				t.Errorf("GPOS table script table default lang sys req feature index not match at %d %d", i, scriptRecord.ScriptTable.DefaultLangSys.ReqFeatureIndex)
			}
		}
	}

}
