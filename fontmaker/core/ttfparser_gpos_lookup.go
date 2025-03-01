package core

import (
	"bytes"
)

// //https://learn.microsoft.com/en-us/typography/opentype/spec/gpos#lookup-type-4-subtable-mark-to-base-attachment-positioning
func (t *TTFParser) parseLookupTable(fd *bytes.Reader, lookupOffset int64) (LookupTable, error) {
	err := fdJumpTo(fd, lookupOffset)
	if err != nil {
		return LookupTable{}, err
	}

	lookupType, err := t.ReadUShortUint16(fd)
	if err != nil {
		return LookupTable{}, err
	}

	lookupFlag, err := t.ReadUShortUint16(fd)
	if err != nil {
		return LookupTable{}, err
	}

	subTableCount, err := t.ReadUShortUint16(fd)
	if err != nil {
		return LookupTable{}, err
	}

	subTableOffset := make([]uint16, subTableCount)
	for i := uint16(0); i < subTableCount; i++ {
		subTableOffset[i], err = t.ReadUShortUint16(fd)
		if err != nil {
			return LookupTable{}, err
		}
	}

	markFilteringSet := uint16(0)
	if lookupFlag == useMarkFilteringSet {
		markFilteringSet, err = t.ReadUShortUint16(fd)
		if err != nil {
			return LookupTable{}, err
		}
	}

	switch lookupType {
	case 4:
		{
			lk4, err := t.parseLookupSubTableType4(fd)
			if err != nil {
				return LookupTable{}, err
			}
			err = t.parseBaseArray(fd, lk4)
			if err != nil {
				return LookupTable{}, err
			}
		}
	}

	return LookupTable{
		LookupType:     lookupType,
		LookupFlag:     lookupFlag,
		SubTableCount:  subTableCount,
		SubTableOffset: subTableOffset,
		//SubTables:        subTables,
		MarkFilteringSet: markFilteringSet,
	}, nil
}

func (t *TTFParser) parseBaseArray(fd *bytes.Reader, lk LookupSubtableType4) error {
	baseCount, err := t.ReadUShortUint16(fd)
	if err != nil {
		return err
	}

	baseRecords := make([]BaseRecord, baseCount)
	for i := uint16(0); i < baseCount; i++ {
		baseRecords[i], err = t.parseBaseRecord(fd, lk.MarkClassCount)
		if err != nil {
			return err
		}
	}

	return nil
}

func (t *TTFParser) parseBaseRecord(fd *bytes.Reader, markClassCount uint16) (BaseRecord, error) {
	br := BaseRecord{}
	for i := uint16(0); i < markClassCount; i++ {
		markAnchorOffset, err := t.ReadUShortUint16(fd)
		if err != nil {
			return BaseRecord{}, err
		}
		br.BaseAnchorOffsets = append(br.BaseAnchorOffsets, markAnchorOffset)
	}
	return br, nil
}

func (t *TTFParser) parseLookupSubTableType4(fd *bytes.Reader) (LookupSubtableType4, error) {

	format, err := t.ReadUShortUint16(fd)
	if err != nil {
		return LookupSubtableType4{}, err
	}

	markCoverageOffset, err := t.ReadUShortUint16(fd)
	if err != nil {
		return LookupSubtableType4{}, err
	}

	baseCoverageOffset, err := t.ReadUShortUint16(fd)
	if err != nil {
		return LookupSubtableType4{}, err
	}

	markClassCount, err := t.ReadUShortUint16(fd)
	if err != nil {
		return LookupSubtableType4{}, err
	}

	markArrayOffset, err := t.ReadUShortUint16(fd)
	if err != nil {
		return LookupSubtableType4{}, err
	}

	baseArrayOffset, err := t.ReadUShortUint16(fd)
	if err != nil {
		return LookupSubtableType4{}, err
	}

	return LookupSubtableType4{
		Format:             format,
		MarkCoverageOffset: markCoverageOffset,
		BaseCoverageOffset: baseCoverageOffset,
		MarkClassCount:     markClassCount,
		MarkArrayOffset:    markArrayOffset,
		BaseArrayOffset:    baseArrayOffset,
	}, nil
}
