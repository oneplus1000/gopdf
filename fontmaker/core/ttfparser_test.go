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

	//script list
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

	//feature list
	if len(p.gpos.FeatureList.FeatureRecords) != 3 {
		t.Errorf("GPOS table feature count not match")
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

	//feature list
	if p.gpos.FeatureList.FeatureCount != 3 || len(p.gpos.FeatureList.FeatureRecords) != 3 {
		t.Errorf("GPOS table feature count not match")
	}

	for i, featureRecord := range p.gpos.FeatureList.FeatureRecords {
		if i == 0 {
			if featureRecord.FeatureTagString() != "kern" {
				t.Errorf("GPOS table feature tag not match at %d %s", i, featureRecord.FeatureTagString())
			}
			if featureRecord.FeatureTable.LookupCount != 1 || len(featureRecord.FeatureTable.LookupListIndex) != 1 {
				t.Errorf("GPOS table feature lookup count not match at %d %d", i, featureRecord.FeatureTable.LookupCount)
			}
		} else if i == 1 {
			if featureRecord.FeatureTagString() != "mark" {
				t.Errorf("GPOS table feature tag not match at %d %s", i, featureRecord.FeatureTagString())
			}
			if featureRecord.FeatureTable.LookupCount != 2 || len(featureRecord.FeatureTable.LookupListIndex) != 2 {
				t.Errorf("GPOS table feature lookup count not match at %d %d", i, featureRecord.FeatureTable.LookupCount)
			}
		} else if i == 2 {
			if featureRecord.FeatureTagString() != "mkmk" {
				t.Errorf("GPOS table feature tag not match at %d %s", i, featureRecord.FeatureTagString())
			}
			if featureRecord.FeatureTable.LookupCount != 1 || len(featureRecord.FeatureTable.LookupListIndex) != 1 {
				t.Errorf("GPOS table feature lookup count not match at %d %d", i, featureRecord.FeatureTable.LookupCount)
			}
		}
	}

	//lookup list
	if p.gpos.LookupList.LookupCount != 4 || len(p.gpos.LookupList.LookupOffsets) != 4 || len(p.gpos.LookupList.LookupTables) != 4 {
		t.Errorf("GPOS table lookup count not match")
	}

	for i, lookupTable := range p.gpos.LookupList.LookupTables {
		if i == 0 {
			if lookupTable.LookupType != 4 {
				t.Errorf("GPOS table lookup type not match at %d %d", i, lookupTable.LookupType)
			}
			if lookupTable.LookupFlag != 0 {
				t.Errorf("GPOS table lookup flag not match at %d %d", i, lookupTable.LookupFlag)
			}
			if lookupTable.SubTableCount != 1 || len(lookupTable.SubTableOffset) != 1 || len(lookupTable.SubTables) != 1 {
				t.Errorf("GPOS table lookup sub table count not match at %d %d", i, lookupTable.SubTableCount)
			}
			for _, sub := range lookupTable.SubTables {

				if sub.GetFormat() != 1 {
					t.Errorf("GPOS table lookup sub table format not match at %d %d", i, sub.GetFormat())
				}
				//if sub.(*LookupSubtableFormat1).GlyphCount != 3 || len(sub.(*LookupSubtableFormat1).GlyphIdArray) != 3 {
				//	t.Errorf("GPOS table lookup sub table glyph count not match at %d %d", i, sub.(*LookupSubtableFormat1).GlyphCount)
				//}
			}
		}
	}

}
