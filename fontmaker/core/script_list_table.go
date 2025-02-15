package core

type ScriptListTable struct {
	ScriptCount   uint16
	ScriptRecords []ScriptRecord
}

type ScriptRecord struct {
	ScriptTag    uint
	ScriptOffset uint16
	ScriptTable
}

func (s ScriptRecord) ScriptTagString() string {
	return string([]byte{
		byte(s.ScriptTag >> 24),
		byte(s.ScriptTag >> 16),
		byte(s.ScriptTag >> 8),
		byte(s.ScriptTag),
	})
}

type ScriptTable struct {
	DefaultLangSysOffset uint16
	LangSysCount         uint16
	DefaultLangSys       LangSysTable
	LangSysRecords       []LangSysRecord
}

type LangSysRecord struct {
	LangSysTag uint
	LangSys    LangSysTable
}

type LangSysTable struct {
	LookupOrderOffset uint16
	ReqFeatureIndex   uint16
	FeatureIndex      []uint16
}
