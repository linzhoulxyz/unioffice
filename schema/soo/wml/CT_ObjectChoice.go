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

type CT_ObjectChoice struct {
	Control     *CT_Control
	ObjectLink  *CT_ObjectLink
	ObjectEmbed *CT_ObjectEmbed
	Movie       *CT_Rel
}

func NewCT_ObjectChoice() *CT_ObjectChoice {
	ret := &CT_ObjectChoice{}
	return ret
}

func (m *CT_ObjectChoice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.Control != nil {
		secontrol := xml.StartElement{Name: xml.Name{Local: "w:control"}}
		e.EncodeElement(m.Control, secontrol)
	}
	if m.ObjectLink != nil {
		seobjectLink := xml.StartElement{Name: xml.Name{Local: "w:objectLink"}}
		e.EncodeElement(m.ObjectLink, seobjectLink)
	}
	if m.ObjectEmbed != nil {
		seobjectEmbed := xml.StartElement{Name: xml.Name{Local: "w:objectEmbed"}}
		e.EncodeElement(m.ObjectEmbed, seobjectEmbed)
	}
	if m.Movie != nil {
		semovie := xml.StartElement{Name: xml.Name{Local: "w:movie"}}
		e.EncodeElement(m.Movie, semovie)
	}
	return nil
}

func (m *CT_ObjectChoice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_ObjectChoice:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "control"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "control"}:
				m.Control = NewCT_Control()
				if err := d.DecodeElement(m.Control, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "objectLink"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "objectLink"}:
				m.ObjectLink = NewCT_ObjectLink()
				if err := d.DecodeElement(m.ObjectLink, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "objectEmbed"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "objectEmbed"}:
				m.ObjectEmbed = NewCT_ObjectEmbed()
				if err := d.DecodeElement(m.ObjectEmbed, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "movie"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "movie"}:
				m.Movie = NewCT_Rel()
				if err := d.DecodeElement(m.Movie, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_ObjectChoice %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ObjectChoice
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ObjectChoice and its children
func (m *CT_ObjectChoice) Validate() error {
	return m.ValidateWithPath("CT_ObjectChoice")
}

// ValidateWithPath validates the CT_ObjectChoice and its children, prefixing error messages with path
func (m *CT_ObjectChoice) ValidateWithPath(path string) error {
	if m.Control != nil {
		if err := m.Control.ValidateWithPath(path + "/Control"); err != nil {
			return err
		}
	}
	if m.ObjectLink != nil {
		if err := m.ObjectLink.ValidateWithPath(path + "/ObjectLink"); err != nil {
			return err
		}
	}
	if m.ObjectEmbed != nil {
		if err := m.ObjectEmbed.ValidateWithPath(path + "/ObjectEmbed"); err != nil {
			return err
		}
	}
	if m.Movie != nil {
		if err := m.Movie.ValidateWithPath(path + "/Movie"); err != nil {
			return err
		}
	}
	return nil
}
