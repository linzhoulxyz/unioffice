// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/linzhoulxyz/unioffice"
)

type CT_MdxTuple struct {
	// Member Index Count
	CAttr *uint32
	// Server Formatting Culture Currency
	CtAttr *string
	// Server Formatting String Index
	SiAttr *uint32
	// Server Formatting Built-In Number Format Index
	FiAttr *uint32
	// Server Formatting Background Color
	BcAttr *string
	// Server Formatting Foreground Color
	FcAttr *string
	// Server Formatting Italic Font
	IAttr *bool
	// Server Formatting Underline Font
	UAttr *bool
	// Server Formatting Strikethrough Font
	StAttr *bool
	// Server Formatting Bold Font
	BAttr *bool
	// Member Unique Name Index
	N []*CT_MetadataStringIndex
}

func NewCT_MdxTuple() *CT_MdxTuple {
	ret := &CT_MdxTuple{}
	return ret
}

func (m *CT_MdxTuple) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.CAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "c"},
			Value: fmt.Sprintf("%v", *m.CAttr)})
	}
	if m.CtAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ct"},
			Value: fmt.Sprintf("%v", *m.CtAttr)})
	}
	if m.SiAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "si"},
			Value: fmt.Sprintf("%v", *m.SiAttr)})
	}
	if m.FiAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "fi"},
			Value: fmt.Sprintf("%v", *m.FiAttr)})
	}
	if m.BcAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "bc"},
			Value: fmt.Sprintf("%v", *m.BcAttr)})
	}
	if m.FcAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "fc"},
			Value: fmt.Sprintf("%v", *m.FcAttr)})
	}
	if m.IAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "i"},
			Value: fmt.Sprintf("%d", b2i(*m.IAttr))})
	}
	if m.UAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "u"},
			Value: fmt.Sprintf("%d", b2i(*m.UAttr))})
	}
	if m.StAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "st"},
			Value: fmt.Sprintf("%d", b2i(*m.StAttr))})
	}
	if m.BAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "b"},
			Value: fmt.Sprintf("%d", b2i(*m.BAttr))})
	}
	e.EncodeToken(start)
	if m.N != nil {
		sen := xml.StartElement{Name: xml.Name{Local: "ma:n"}}
		for _, c := range m.N {
			e.EncodeElement(c, sen)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_MdxTuple) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "c" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.CAttr = &pt
			continue
		}
		if attr.Name.Local == "ct" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CtAttr = &parsed
			continue
		}
		if attr.Name.Local == "si" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.SiAttr = &pt
			continue
		}
		if attr.Name.Local == "fi" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.FiAttr = &pt
			continue
		}
		if attr.Name.Local == "bc" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.BcAttr = &parsed
			continue
		}
		if attr.Name.Local == "fc" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.FcAttr = &parsed
			continue
		}
		if attr.Name.Local == "i" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.IAttr = &parsed
			continue
		}
		if attr.Name.Local == "u" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.UAttr = &parsed
			continue
		}
		if attr.Name.Local == "st" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.StAttr = &parsed
			continue
		}
		if attr.Name.Local == "b" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.BAttr = &parsed
			continue
		}
	}
lCT_MdxTuple:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "n"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "n"}:
				tmp := NewCT_MetadataStringIndex()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.N = append(m.N, tmp)
			default:
				unioffice.Log("skipping unsupported element on CT_MdxTuple %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_MdxTuple
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_MdxTuple and its children
func (m *CT_MdxTuple) Validate() error {
	return m.ValidateWithPath("CT_MdxTuple")
}

// ValidateWithPath validates the CT_MdxTuple and its children, prefixing error messages with path
func (m *CT_MdxTuple) ValidateWithPath(path string) error {
	for i, v := range m.N {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/N[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
