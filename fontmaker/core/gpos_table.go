package core

type GPOSTable struct {
	MajorVersion            uint16
	MinorVersion            uint16
	ScriptListOffset        uint16
	FeatureListOffset       uint16
	LookupListOffset        uint16
	FeatureVariationsOffset uint
	//table
	ScriptList  ScriptListTable
	FeatureList FeatureListTable
}

func (g GPOSTable) Version() uint {
	tmp := uint(g.MajorVersion)
	v := uint(tmp<<16) + uint(g.MinorVersion)
	return v
}
