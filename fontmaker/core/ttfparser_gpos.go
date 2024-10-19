package core

import "bytes"

// https://learn.microsoft.com/en-us/typography/opentype/spec/gsub#SS
func (t *TTFParser) ParseGSUB(fd *bytes.Reader) error {
	err := t.Seek(fd, "GSUB")
	if err == ErrTableNotFound {
		return nil
	} else if err != nil {
		return err
	}

	tableOffset, err := t.FTell(fd)
	if err != nil {
		return err
	}

	majorVersion, err := t.ReadUShort(fd)
	if err != nil {
		return err
	}
	minorVersion, err := t.ReadUShort(fd)
	if err != nil {
		return err
	}

	scriptListOffset, err := t.ReadUShort(fd)
	if err != nil {
		return err
	}
	featureListOffset, err := t.ReadUShort(fd)
	if err != nil {
		return err
	}
	lookupListOffset, err := t.ReadUShort(fd)
	if err != nil {
		return err
	}

	_ = majorVersion
	_ = minorVersion
	_ = scriptListOffset
	_ = featureListOffset
	_ = lookupListOffset

	t.parseLookupList(fd, tableOffset, lookupListOffset)

	return nil
}

// https://learn.microsoft.com/en-us/typography/opentype/otspec180/chapter2#lulTbl
func (t *TTFParser) parseLookupList(fd *bytes.Reader, parentOffset uint, lookupListOffset uint) error {
	offset := parentOffset + lookupListOffset
	_, err := fd.Seek(int64(offset), 0)
	if err != nil {
		return err
	}

	lookupCount, err := t.ReadUShort(fd)
	if err != nil {
		return err
	}
	if lookupCount == 0 {
		return nil
	}

	lookupOffsets := make([]uint, lookupCount)
	for i := 0; i < int(lookupCount); i++ {
		lookupOffsets[i], err = t.ReadUShort(fd)
		if err != nil {
			return err
		}
	}

	for i := 0; i < len(lookupOffsets); i++ {
		err = t.parseLookup(fd, uint(offset), lookupOffsets[i])
		if err != nil {
			return err
		}
	}

	return nil
}

func (t *TTFParser) parseLookup(fd *bytes.Reader, parentOffset uint, lookupOffset uint) error {
	_, err := fd.Seek(int64(parentOffset+lookupOffset), 0)
	if err != nil {
		return err
	}
	lookupType, err := t.ReadUShort(fd)
	if err != nil {
		return err
	}

	lookupFlag, err := t.ReadUShort(fd)
	if err != nil {
		return err
	}
	subTableCount, err := t.ReadUShort(fd)
	if err != nil {
		return err
	}

	subtableOffsets := make([]uint, subTableCount)
	for i := 0; i < int(subTableCount); i++ {
		subtableOffsets[i], err = t.ReadUShort(fd)
		if err != nil {
			return err
		}
	}

	markFilteringSet, err := t.ReadUShort(fd)
	if err != nil {
		return err
	}

	_ = lookupType
	_ = lookupFlag
	_ = subTableCount
	_ = markFilteringSet

	for i := 0; i < len(subtableOffsets); i++ {
		switch lookupType {
		case 1:
			err = t.parseLookupType1Subtable(fd, parentOffset+lookupOffset, subtableOffsets[i])
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// https://learn.microsoft.com/en-us/typography/opentype/spec/gsub#lookup-type-1-subtable-single-substitution
func (t *TTFParser) parseLookupType1Subtable(fd *bytes.Reader, parentOffset, subtableOffset uint) error {
	_, err := fd.Seek(int64(parentOffset+subtableOffset), 0)
	if err != nil {
		return err
	}
	format, err := t.ReadUShort(fd)
	if err != nil {
		return err
	}

	switch format {
	case 1:
		err = t.parseSingleSubstitutionFormat1(fd, int64(parentOffset+subtableOffset))
		if err != nil {
			return err
		}
	case 2:
		err = t.parseSingleSubstitutionFormat2(fd, int64(parentOffset+subtableOffset))
		if err != nil {
			return err
		}
	}
	return nil
}

func (t *TTFParser) parseSingleSubstitutionFormat1(fd *bytes.Reader, parentOffset int64) error {
	coverageOffset, err := t.ReadUShort(fd)
	if err != nil {
		return err
	}
	deltaGlyphID, err := t.ReadShort(fd)
	if err != nil {
		return err
	}
	err = t.parseCoverageTable(fd, parentOffset+int64(coverageOffset))
	if err != nil {
		return err
	}
	_ = coverageOffset
	_ = deltaGlyphID
	return nil
}

func (t *TTFParser) parseSingleSubstitutionFormat2(fd *bytes.Reader, parentOffset int64) error {
	coverageOffset, err := t.ReadUShort(fd)
	if err != nil {
		return err
	}
	glyphCount, err := t.ReadUShort(fd)
	if err != nil {
		return err
	}

	substituteGlyphIDs := make([]uint, glyphCount)
	for i := 0; i < int(glyphCount); i++ {
		substituteGlyphID, err := t.ReadUShort(fd)
		if err != nil {
			return err
		}
		substituteGlyphIDs = append(substituteGlyphIDs, substituteGlyphID)
	}

	err = t.parseCoverageTable(fd, parentOffset+int64(coverageOffset))
	if err != nil {
		return err
	}

	_ = coverageOffset
	_ = glyphCount
	_ = substituteGlyphIDs
	return nil
}

func (t *TTFParser) parseCoverageTable(fd *bytes.Reader, offset int64) error {
	_, err := fd.Seek(offset, 0)
	if err != nil {
		return err
	}

	coverageFormat, err := t.ReadUShort(fd)
	if err != nil {
		return err
	}
	switch coverageFormat {
	case 1:
		return t.parseCoverageFormat1(fd)
	case 2:
		return t.parseCoverageFormat2(fd)
	}
	return nil
}

func (t *TTFParser) parseCoverageFormat1(fd *bytes.Reader) error {
	glyphCount, err := t.ReadUShort(fd)
	if err != nil {
		return err
	}
	glyphArray := make([]uint, glyphCount)
	for i := 0; i < int(glyphCount); i++ {
		glyphID, err := t.ReadUShort(fd)
		if err != nil {
			return err
		}
		glyphArray = append(glyphArray, glyphID)
	}
	_ = glyphArray
	return nil
}

func (t *TTFParser) parseCoverageFormat2(fd *bytes.Reader) error {
	rangeCount, err := t.ReadUShort(fd)
	if err != nil {
		return err
	}

	for i := 0; i < int(rangeCount); i++ {
		err = t.parseCoverageFormat2RangeRecord(fd)
		if err != nil {
			return err
		}
	}

	_ = rangeCount
	return nil
}

func (t *TTFParser) parseCoverageFormat2RangeRecord(fd *bytes.Reader) error {
	startGlyphID, err := t.ReadUShort(fd)
	if err != nil {
		return err
	}
	endGlyphID, err := t.ReadUShort(fd)
	if err != nil {
		return err
	}

	startCoverageIndex, err := t.ReadUShort(fd)
	if err != nil {
		return err
	}

	_ = startGlyphID
	_ = endGlyphID
	_ = startCoverageIndex

	return nil
}

/*

func (t *TTFParser) parseAdjustmentFormat1(fd *bytes.Reader) error {
	coverageOffset, err := t.ReadUShort(fd)
	if err != nil {
		return err
	}
	valueFormat, err := t.ReadUShort(fd)
	if err != nil {
		return err
	}
	valueRecord, err := t.parseValueRecord(fd)
	if err != nil {
		return err
	}
	_ = coverageOffset
	_ = valueFormat
	_ = valueRecord
	return nil
}

// parse ValueRecord
// https://learn.microsoft.com/en-us/typography/opentype/otspec181/gpos#valuerecord
func (t *TTFParser) parseValueRecord(fd *bytes.Reader) (ValueRecord, error) {
	vr := ValueRecord{}
	xplacement, err := t.ReadShort(fd)
	if err != nil {
		return vr, nil
	}
	yplacement, err := t.ReadShort(fd)
	if err != nil {
		return vr, nil
	}
	xadvance, err := t.ReadShort(fd)
	if err != nil {
		return vr, nil
	}
	yadvance, err := t.ReadShort(fd)
	if err != nil {
		return vr, nil
	}
	xplaDevice, err := t.ReadUShort(fd)
	if err != nil {
		return vr, nil
	}
	yplaDevice, err := t.ReadUShort(fd)
	if err != nil {
		return vr, nil
	}

	xadvDevice, err := t.ReadUShort(fd)
	if err != nil {
		return vr, nil
	}

	yadvDevice, err := t.ReadUShort(fd)
	if err != nil {
		return vr, nil
	}

	vr.XPlacement = xplacement
	vr.YPlacement = yplacement
	vr.XAdvance = xadvance
	vr.YAdvance = yadvance
	vr.XPlaDevice = xplaDevice
	vr.YPlaDevice = yplaDevice
	vr.XAdvDevice = xadvDevice
	vr.YAdvDevice = yadvDevice
	return vr, nil
}
*/
