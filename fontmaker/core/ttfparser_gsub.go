package core

import (
	"bytes"
	"errors"
	"fmt"
)

//ErrGSUBMajorVersionNotsubpport GSUB major version not subpport
var ErrGSUBMajorVersionNotsubpport = errors.New("GSUB major version not subpport")

//ErrGSUBMinorVersionNotsubpport GSUB minor version not subpport
var ErrGSUBMinorVersionNotsubpport = errors.New("GSUB minor version not subpport")

//ParseGSUB gsub
//	- https://www.microsoft.com/typography/otspec/gsub.htm#EX1
//  - https://www.microsoft.com/typography/otspec/chapter2.htm
//support LookupType type 1,4
func (t *TTFParser) ParseGSUB(fd *bytes.Reader) error {
	err := t.Seek(fd, "GSUB")
	if err != nil {
		return err
	}

	gsubOffset := fd.Len()

	majorVersion, err := t.ReadUShort(fd)
	if err != nil {
		return err
	}

	if majorVersion != 1 {
		return ErrGSUBMajorVersionNotsubpport
	}

	minorVersion, err := t.ReadUShort(fd)
	if err != nil {
		return err
	}

	if minorVersion != 0 && minorVersion != 1 {
		return ErrGSUBMinorVersionNotsubpport
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

	err = t.parseGsubScriptList(fd, int64(gsubOffset+int(scriptListOffset)))
	if err != nil {
		return err
	}

	err = t.parseGsubFeatureList(fd, int64(gsubOffset+int(featureListOffset)))
	if err != nil {
		return err
	}

	err = t.parseGsubLookupList(fd, int64(gsubOffset+int(lookupListOffset)))
	if err != nil {
		return err
	}

	return nil
}

func (t *TTFParser) parseGsubScriptList(fd *bytes.Reader, offset int64) error {
	return nil
}

func (t *TTFParser) parseGsubFeatureList(fd *bytes.Reader, offset int64) error {
	return nil
}

func (t *TTFParser) parseGsubLookupList(fd *bytes.Reader, offset int64) error {
	var err error
	_, err = fd.Seek(offset, 0)
	if err != nil {
		return err
	}

	lookupCount, err := t.ReadUShort(fd)
	if err != nil {
		return err
	}
	fmt.Printf("-------------------------lookupCount=%d\n", lookupCount)

	i := uint(0)
	for i < lookupCount {
		lookupTableOffset, err := t.ReadUShort(fd)
		if err != nil {
			return err
		}
		fmt.Printf("lookupTableOffset=%d\n", lookupTableOffset)

		i++
	}

	return nil
}
