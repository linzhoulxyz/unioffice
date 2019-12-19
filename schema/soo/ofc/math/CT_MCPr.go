// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package math

import (
	"encoding/xml"

	"github.com/linzhoulxyz/unioffice"
)

type CT_MCPr struct {
	Count *CT_Integer255
	McJc  *CT_XAlign
}

func NewCT_MCPr() *CT_MCPr {
	ret := &CT_MCPr{}
	return ret
}

func (m *CT_MCPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Count != nil {
		secount := xml.StartElement{Name: xml.Name{Local: "m:count"}}
		e.EncodeElement(m.Count, secount)
	}
	if m.McJc != nil {
		semcJc := xml.StartElement{Name: xml.Name{Local: "m:mcJc"}}
		e.EncodeElement(m.McJc, semcJc)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_MCPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_MCPr:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "count"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "count"}:
				m.Count = NewCT_Integer255()
				if err := d.DecodeElement(m.Count, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "mcJc"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "mcJc"}:
				m.McJc = NewCT_XAlign()
				if err := d.DecodeElement(m.McJc, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_MCPr %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_MCPr
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_MCPr and its children
func (m *CT_MCPr) Validate() error {
	return m.ValidateWithPath("CT_MCPr")
}

// ValidateWithPath validates the CT_MCPr and its children, prefixing error messages with path
func (m *CT_MCPr) ValidateWithPath(path string) error {
	if m.Count != nil {
		if err := m.Count.ValidateWithPath(path + "/Count"); err != nil {
			return err
		}
	}
	if m.McJc != nil {
		if err := m.McJc.ValidateWithPath(path + "/McJc"); err != nil {
			return err
		}
	}
	return nil
}
