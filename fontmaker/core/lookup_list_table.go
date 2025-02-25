package core

type LookupListTable struct {
	LookupCount   uint16
	LookupOffsets []uint16
	LookupTables  []LookupTable
}

type LookupTable struct {
	LookupType       uint16
	LookupFlag       uint16
	SubTableCount    uint16
	SubTableOffset   []uint16
	MarkFilteringSet uint16
	SubTables        []LookupSubtable
}

type LookupSubtable interface {
	GetFormat() uint16
	GetBeginningOfSinglePosSubtable() int64
}

type LookupSubtableFormat1 struct {
	Format                       uint16
	BeginningOfSinglePosSubtable int64
	CoverageOffset               uint16
	ValueFormat                  uint16
	ValueRecord                  ValueRecord
}

func (l LookupSubtableFormat1) GetFormat() uint16 {
	return l.Format
}

func (l LookupSubtableFormat1) GetBeginningOfSinglePosSubtable() int64 {
	return l.BeginningOfSinglePosSubtable
}

type LookupSubtableFormat2 struct {
	Format uint16
}

type CoverageFormat1 struct {
	Format       uint16
	CoverageSize uint16
	GlyphArray   []uint16
}

/*
*
Lookup type 4 subtable: mark-to-base attachment positioning
Type	Name	Description
uint16	format	Format identifier — format = 1.
Offset16	markCoverageOffset	Offset to markCoverage table, from beginning of MarkBasePos subtable.
Offset16	baseCoverageOffset	Offset to baseCoverage table, from beginning of MarkBasePos subtable.
uint16	markClassCount	Number of classes defined for marks.
Offset16	markArrayOffset	Offset to MarkArray table, from beginning of MarkBasePos subtable.
Offset16	baseArrayOffset	Offset to BaseArray table, from beginning of MarkBasePos subtable.
*/
type LookupSubtableType4 struct {
	Format             uint16
	MarkCoverageOffset uint16
	BaseCoverageOffset uint16
	MarkClassCount     uint16
	MarkArrayOffset    uint16
	BaseArrayOffset    uint16
}
