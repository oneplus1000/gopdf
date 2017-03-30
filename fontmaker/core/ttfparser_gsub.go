package core

import (
	"bytes"
	"errors"
)

//ErrGSUBMajorVersionNotsubpport GSUB major version not subpport
var ErrGSUBMajorVersionNotsubpport = errors.New("GSUB major version not subpport")

//ErrGSUBMinorVersionNotsubpport GSUB minor version not subpport
var ErrGSUBMinorVersionNotsubpport = errors.New("GSUB minor version not subpport")

//ParseGSUB https://www.microsoft.com/typography/otspec/gsub.htm#EX1
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

	err = t.parseGsubScriptList(fd, uint(gsubOffset)+scriptListOffset)
	if err != nil {
		return err
	}

	err = t.parseGsubFeatureList(fd, uint(gsubOffset)+featureListOffset)
	if err != nil {
		return err
	}

	err = t.parseGsubLookupList(fd, uint(gsubOffset)+lookupListOffset)
	if err != nil {
		return err
	}

	return nil
}

func (t *TTFParser) parseGsubScriptList(fd *bytes.Reader, offset uint) error {
	return nil
}

func (t *TTFParser) parseGsubFeatureList(fd *bytes.Reader, offset uint) error {
	return nil
}

func (t *TTFParser) parseGsubLookupList(fd *bytes.Reader, offset uint) error {
	return nil
}
