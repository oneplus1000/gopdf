package core

import (
	"bytes"
)

func (t *TTFParser) ParseGPOS(fd *bytes.Reader) error {
	t.gpos = nil //clear
	err := t.Seek(fd, "GPOS")
	if err == ErrTableNotFound {
		return nil
	} else if err != nil {
		return err
	}

	t.gpos = new(GPOSTable) //init
	err = t.parseGPOSHeader(fd)
	if err != nil {
		return err
	}
	err = t.parseGPOSScriptList(fd)
	if err != nil {
		return err
	}
	err = t.parseFeatureList(fd)
	if err != nil {
		return err
	}

	return nil

}

func (t *TTFParser) parseGPOSHeader(fd *bytes.Reader) error {
	majorVersion, err := t.ReadUShortUint16(fd)
	if err != nil {
		return err
	}
	t.gpos.MajorVersion = majorVersion

	minorVersion, err := t.ReadUShortUint16(fd)
	if err != nil {
		return err
	}
	t.gpos.MinorVersion = minorVersion

	scriptListOffset, err := t.ReadUShortUint16(fd)
	if err != nil {
		return err
	}
	t.gpos.ScriptListOffset = scriptListOffset

	featureListOffset, err := t.ReadUShortUint16(fd)
	if err != nil {
		return err
	}
	t.gpos.FeatureListOffset = featureListOffset

	lookupListOffset, err := t.ReadUShortUint16(fd)
	if err != nil {
		return err
	}
	t.gpos.LookupListOffset = lookupListOffset

	if majorVersion == 1 {
		featureVariationsOffset, err := t.ReadULong(fd)
		if err != nil {
			return err
		}
		t.gpos.FeatureVariationsOffset = featureVariationsOffset
	}
	return nil
}

func (t *TTFParser) parseGPOSScriptList(fd *bytes.Reader) error {
	err := t.Seek(fd, "GPOS")
	if err == ErrTableNotFound {
		return nil
	} else if err != nil {
		return err
	}
	gposOffset := fdCurrentOffset(fd) //save current offset

	scriptListOffset := t.gpos.ScriptListOffset
	err = t.Skip(fd, int(scriptListOffset)) //skip count
	if err != nil {
		return err
	}

	scriptCount, err := t.ReadUShortUint16(fd)
	if err != nil {
		return err
	}
	t.gpos.ScriptList.ScriptCount = scriptCount

	//check script records
	scriptRecords := make([]ScriptRecord, scriptCount)
	i := uint16(0)
	for i = 0; i < scriptCount; i++ {
		scriptTag, err := t.ReadULong(fd)
		if err != nil {
			return err
		}
		scriptOffet, err := t.ReadUShortUint16(fd)
		if err != nil {
			return err
		}
		scriptRecords[i] = ScriptRecord{
			ScriptTag:    scriptTag,
			ScriptOffset: scriptOffet,
		}
	}
	t.gpos.ScriptList.ScriptRecords = scriptRecords
	//parse script table
	beginOfScriptListTable := gposOffset + int64(scriptListOffset)
	for i, scriptRecord := range scriptRecords {
		scriptTable, err := t.parseScriptTable(fd, beginOfScriptListTable, scriptRecord)
		if err != nil {
			return err
		}
		scriptRecords[i].ScriptTable = scriptTable
	}

	return nil
}

func (t *TTFParser) parseFeatureList(fd *bytes.Reader) error {

	err := t.Seek(fd, "GPOS")
	if err == ErrTableNotFound {
		return nil
	} else if err != nil {
		return err
	}
	gposOffset := fdCurrentOffset(fd) //save current offset
	featureListOffset := t.gpos.FeatureListOffset
	err = t.Skip(fd, int(featureListOffset)) //skip count
	if err != nil {
		return err
	}

	featureCount, err := t.ReadUShortUint16(fd)
	if err != nil {
		return err
	}
	t.gpos.FeatureList.FeatureCount = featureCount

	//check feature records
	featureRecords := make([]FeatureRecord, featureCount)
	i := uint16(0)
	for i = 0; i < featureCount; i++ {
		featureTag, err := t.ReadULong(fd)
		if err != nil {
			return err
		}
		featureOffset, err := t.ReadUShortUint16(fd)
		if err != nil {
			return err
		}
		featureRecords[i] = FeatureRecord{
			FeatureTag:    featureTag,
			FeatureOffset: featureOffset,
		}
		t.gpos.FeatureList.FeatureRecords = featureRecords
	}

	//parse feature table
	beginningOfFeatureLis := gposOffset + int64(featureListOffset)
	for i = 0; i < featureCount; i++ {
		featureTable, err := t.parseFeatureTable(fd, beginningOfFeatureLis+int64(featureRecords[i].FeatureOffset))
		if err != nil {
			return err
		}
		featureRecords[i].FeatureTable = featureTable
	}

	return nil
}

func (t *TTFParser) parseFeatureTable(fd *bytes.Reader, featureOffset int64) (FeatureTable, error) {
	err := fdJumpTo(fd, featureOffset)
	if err != nil {
		return FeatureTable{}, err
	}

	featureParamsOffset, err := t.ReadUShortUint16(fd)
	if err != nil {
		return FeatureTable{}, err
	}

	lookupCount, err := t.ReadUShortUint16(fd)
	if err != nil {
		return FeatureTable{}, err
	}

	lookupListIndex := make([]uint16, lookupCount)
	for i := uint16(0); i < lookupCount; i++ {
		lookupListIndex[i], err = t.ReadUShortUint16(fd)
		if err != nil {
			return FeatureTable{}, err
		}
	}

	return FeatureTable{
		FeatureParamsOffset: featureParamsOffset,
		LookupCount:         lookupCount,
		LookupListIndex:     lookupListIndex,
	}, nil
}

func (t *TTFParser) parseScriptTable(fd *bytes.Reader,
	scriptListTableOffset int64,
	record ScriptRecord,
) (ScriptTable, error) {
	beginningOfScriptTable := scriptListTableOffset + int64(record.ScriptOffset)
	err := fdJumpTo(fd, beginningOfScriptTable)
	if err != nil {
		return ScriptTable{}, err
	}

	scTable := ScriptTable{}

	defaultLangSysOffset, err := t.ReadUShortUint16(fd)
	if err != nil {
		return ScriptTable{}, err
	}
	scTable.DefaultLangSysOffset = defaultLangSysOffset

	langSysCount, err := t.ReadUShortUint16(fd)
	if err != nil {
		return ScriptTable{}, err
	}
	scTable.LangSysCount = langSysCount

	//read LangSysRecords
	//LangSysRecords https://learn.microsoft.com/en-us/typography/opentype/spec/chapter2 langSysRecords
	langSysRecords := make([]LangSysRecord, langSysCount)
	for i := uint16(0); i < langSysCount; i++ {
		langSysTag, err := t.ReadULong(fd)
		if err != nil {
			return ScriptTable{}, err
		}
		langSysOffset, err := t.ReadUShortUint16(fd)
		if err != nil {
			return ScriptTable{}, err
		}
		langSysRecords[i] = LangSysRecord{
			LangSysTag:    langSysTag,
			langSysOffset: langSysOffset,
			LangSys:       LangSysTable{},
		}
	}

	//read DefaultLangSys
	if defaultLangSysOffset != 0 {
		defaultLangSysTable, err := t.parseLangSysTable(fd, beginningOfScriptTable+int64(defaultLangSysOffset))
		if err != nil {
			return ScriptTable{}, err
		}
		scTable.DefaultLangSys = defaultLangSysTable
	}

	//read LangSys
	for i, langSysRecord := range langSysRecords {
		langSysTable, err := t.parseLangSysTable(fd, beginningOfScriptTable+int64(langSysRecord.langSysOffset))
		if err != nil {
			return ScriptTable{}, err
		}
		langSysRecords[i].LangSys = langSysTable
	}
	scTable.LangSysRecords = langSysRecords

	return scTable, nil
}

func (t *TTFParser) parseLangSysTable(fd *bytes.Reader, langSysOffset int64) (LangSysTable, error) {
	err := fdJumpTo(fd, langSysOffset)
	if err != nil {
		return LangSysTable{}, err
	}

	langSysTable := LangSysTable{}

	lookupOrder, err := t.ReadUShortUint16(fd)
	if err != nil {
		return LangSysTable{}, err
	}
	langSysTable.LookupOrderOffset = lookupOrder

	reqFeatureIndex, err := t.ReadUShortUint16(fd)
	if err != nil {
		return LangSysTable{}, err
	}
	langSysTable.ReqFeatureIndex = reqFeatureIndex

	featureIndexCount, err := t.ReadUShortUint16(fd)
	if err != nil {
		return LangSysTable{}, err
	}
	featureIndex := make([]uint16, featureIndexCount)
	for i := uint16(0); i < featureIndexCount; i++ {
		featureIndex[i], err = t.ReadUShortUint16(fd)
		if err != nil {
			return LangSysTable{}, err
		}
	}
	langSysTable.FeatureIndex = featureIndex

	return langSysTable, nil
}
