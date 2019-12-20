// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package dml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_PositiveSize2D struct {
	CxAttr int64
	CyAttr int64
}

func NewCT_PositiveSize2D() *CT_PositiveSize2D {
	ret := &CT_PositiveSize2D{}
	ret.CxAttr = 0
	ret.CyAttr = 0
	return ret
}

func (m *CT_PositiveSize2D) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "cx"},
		Value: fmt.Sprintf("%v", m.CxAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "cy"},
		Value: fmt.Sprintf("%v", m.CyAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_PositiveSize2D) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CxAttr = 0
	m.CyAttr = 0
	for _, attr := range start.Attr {
		if attr.Name.Local == "cx" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.CxAttr = parsed
			continue
		}
		if attr.Name.Local == "cy" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.CyAttr = parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_PositiveSize2D: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_PositiveSize2D and its children
func (m *CT_PositiveSize2D) Validate() error {
	return m.ValidateWithPath("CT_PositiveSize2D")
}

// ValidateWithPath validates the CT_PositiveSize2D and its children, prefixing error messages with path
func (m *CT_PositiveSize2D) ValidateWithPath(path string) error {
	if m.CxAttr < 0 {
		return fmt.Errorf("%s/m.CxAttr must be >= 0 (have %v)", path, m.CxAttr)
	}
	if m.CxAttr > 27273042316900 {
		return fmt.Errorf("%s/m.CxAttr must be <= 27273042316900 (have %v)", path, m.CxAttr)
	}
	if m.CyAttr < 0 {
		return fmt.Errorf("%s/m.CyAttr must be >= 0 (have %v)", path, m.CyAttr)
	}
	if m.CyAttr > 27273042316900 {
		return fmt.Errorf("%s/m.CyAttr must be <= 27273042316900 (have %v)", path, m.CyAttr)
	}
	return nil
}

//calculate image width and height
//convention: https://docs.microsoft.com/en-us/dotnet/api/documentformat.openxml.drawing.wordprocessing.extent?view=openxml-2.8.1
type EmuSize struct {
	Width  string
	Height string
}

const (
	mm = 36000
	cm = 360000
	in = 914400
	pt = 12700
	pc = 152400
	pi = 152400
)

//return width and height by using pt
func (m *CT_PositiveSize2D) Size() EmuSize {
	return EmuSize{
		Width:  fmt.Sprintf("%spt", strconv.FormatInt(m.CxAttr/pt, 10)),
		Height: fmt.Sprintf("%spt", strconv.FormatInt(m.CyAttr/pt, 10))}
}
