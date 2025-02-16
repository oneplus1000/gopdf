package core

type FeatureListTable struct {
	FeatureCount   uint16
	FeatureRecords []FeatureRecord
}

type FeatureRecord struct {
	FeatureTag    uint
	FeatureOffset uint16
	FeatureTable
}

func (f FeatureRecord) FeatureTagString() string {
	return uintToString(f.FeatureTag)
}

type FeatureTable struct {
	FeatureParamsOffset uint16
	LookupCount         uint16
	LookupListIndex     []uint16
}
