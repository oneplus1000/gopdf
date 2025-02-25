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
			_, err = t.parseLookupSubTableType4(fd)
			if err != nil {
				return LookupTable{}, err
			}
		}
	}

	/*
		subTables := make([]LookupSubtable, subTableCount)
		for i, offset := range subTableOffset {
			err := fdJumpTo(fd, lookupOffset+int64(offset))
			if err != nil {
				return LookupTable{}, err
			}
			format, err := t.ReadUShortUint16(fd)
			if err != nil {
				return LookupTable{}, err
			}
			switch format {
			case 1:
				{
					subtableFormat1, err := t.parseLookupSubtableFormat1(fd)
					if err != nil {
						return LookupTable{}, err
					}
					subtableFormat1.BeginningOfSinglePosSubtable = lookupOffset + int64(offset)
					subTables[i] = subtableFormat1
				}
			}
		}

		for _, subTable := range subTables {
			if subTable.GetFormat() == 1 {
				subTableFormat1 := subTable.(LookupSubtableFormat1)
				err := fdJumpTo(fd, subTable.GetBeginningOfSinglePosSubtable()+int64(subTableFormat1.CoverageOffset))
				if err != nil {
					return LookupTable{}, err
				}
				coverageFormat, err := t.ReadUShortUint16(fd)
				if err != nil {
					return LookupTable{}, err
				}
				if coverageFormat == 1 {
					//parse coverage format 1
					_, err = t.parseCoverageFormat1(fd)
					if err != nil {
						return LookupTable{}, err
					}
				} else if coverageFormat == 2 {
					//parse coverage format 2
					_, err := t.ReadUShortUint16(fd)
					if err != nil {
						return LookupTable{}, err
					}
					_, err = t.ReadUShortUint16(fd)
					if err != nil {
						return LookupTable{}, err
					}
				}
			}
		}*/

	return LookupTable{
		LookupType:     lookupType,
		LookupFlag:     lookupFlag,
		SubTableCount:  subTableCount,
		SubTableOffset: subTableOffset,
		//SubTables:        subTables,
		MarkFilteringSet: markFilteringSet,
	}, nil
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
