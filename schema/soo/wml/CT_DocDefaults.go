// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package wml

import (
	"encoding/xml"

	"github.com/linzhoulxyz/unioffice"
)

type CT_DocDefaults struct {
	// Default Run Properties
	RPrDefault *CT_RPrDefault
	// Default Paragraph Properties
	PPrDefault *CT_PPrDefault
}

func NewCT_DocDefaults() *CT_DocDefaults {
	ret := &CT_DocDefaults{}
	return ret
}

func (m *CT_DocDefaults) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.RPrDefault != nil {
		serPrDefault := xml.StartElement{Name: xml.Name{Local: "w:rPrDefault"}}
		e.EncodeElement(m.RPrDefault, serPrDefault)
	}
	if m.PPrDefault != nil {
		sepPrDefault := xml.StartElement{Name: xml.Name{Local: "w:pPrDefault"}}
		e.EncodeElement(m.PPrDefault, sepPrDefault)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_DocDefaults) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_DocDefaults:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "rPrDefault"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "rPrDefault"}:
				m.RPrDefault = NewCT_RPrDefault()
				if err := d.DecodeElement(m.RPrDefault, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "pPrDefault"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "pPrDefault"}:
				m.PPrDefault = NewCT_PPrDefault()
				if err := d.DecodeElement(m.PPrDefault, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_DocDefaults %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_DocDefaults
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_DocDefaults and its children
func (m *CT_DocDefaults) Validate() error {
	return m.ValidateWithPath("CT_DocDefaults")
}

// ValidateWithPath validates the CT_DocDefaults and its children, prefixing error messages with path
func (m *CT_DocDefaults) ValidateWithPath(path string) error {
	if m.RPrDefault != nil {
		if err := m.RPrDefault.ValidateWithPath(path + "/RPrDefault"); err != nil {
			return err
		}
	}
	if m.PPrDefault != nil {
		if err := m.PPrDefault.ValidateWithPath(path + "/PPrDefault"); err != nil {
			return err
		}
	}
	return nil
}
