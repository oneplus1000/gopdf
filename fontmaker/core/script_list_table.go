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
	return uintToString(s.ScriptTag)
}

type ScriptTable struct {
	DefaultLangSysOffset uint16
	LangSysCount         uint16
	DefaultLangSys       LangSysTable
	LangSysRecords       []LangSysRecord
}

type LangSysRecord struct {
	LangSysTag    uint
	langSysOffset uint16
	LangSys       LangSysTable
}

func (l LangSysRecord) LangSysTagString() string {
	return uintToString(l.LangSysTag)
}

type LangSysTable struct {
	LookupOrderOffset uint16
	ReqFeatureIndex   uint16
	FeatureIndex      []uint16
}

func uintToString(u uint) string {
	return string([]byte{
		byte(u >> 24),
		byte(u >> 16),
		byte(u >> 8),
		byte(u),
	})
}
