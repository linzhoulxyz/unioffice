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

	"github.com/linzhoulxyz/unioffice"
)

type CT_TextListStyle struct {
	DefPPr  *CT_TextParagraphProperties
	Lvl1pPr *CT_TextParagraphProperties
	Lvl2pPr *CT_TextParagraphProperties
	Lvl3pPr *CT_TextParagraphProperties
	Lvl4pPr *CT_TextParagraphProperties
	Lvl5pPr *CT_TextParagraphProperties
	Lvl6pPr *CT_TextParagraphProperties
	Lvl7pPr *CT_TextParagraphProperties
	Lvl8pPr *CT_TextParagraphProperties
	Lvl9pPr *CT_TextParagraphProperties
	ExtLst  *CT_OfficeArtExtensionList
}

func NewCT_TextListStyle() *CT_TextListStyle {
	ret := &CT_TextListStyle{}
	return ret
}

func (m *CT_TextListStyle) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.DefPPr != nil {
		sedefPPr := xml.StartElement{Name: xml.Name{Local: "a:defPPr"}}
		e.EncodeElement(m.DefPPr, sedefPPr)
	}
	if m.Lvl1pPr != nil {
		selvl1pPr := xml.StartElement{Name: xml.Name{Local: "a:lvl1pPr"}}
		e.EncodeElement(m.Lvl1pPr, selvl1pPr)
	}
	if m.Lvl2pPr != nil {
		selvl2pPr := xml.StartElement{Name: xml.Name{Local: "a:lvl2pPr"}}
		e.EncodeElement(m.Lvl2pPr, selvl2pPr)
	}
	if m.Lvl3pPr != nil {
		selvl3pPr := xml.StartElement{Name: xml.Name{Local: "a:lvl3pPr"}}
		e.EncodeElement(m.Lvl3pPr, selvl3pPr)
	}
	if m.Lvl4pPr != nil {
		selvl4pPr := xml.StartElement{Name: xml.Name{Local: "a:lvl4pPr"}}
		e.EncodeElement(m.Lvl4pPr, selvl4pPr)
	}
	if m.Lvl5pPr != nil {
		selvl5pPr := xml.StartElement{Name: xml.Name{Local: "a:lvl5pPr"}}
		e.EncodeElement(m.Lvl5pPr, selvl5pPr)
	}
	if m.Lvl6pPr != nil {
		selvl6pPr := xml.StartElement{Name: xml.Name{Local: "a:lvl6pPr"}}
		e.EncodeElement(m.Lvl6pPr, selvl6pPr)
	}
	if m.Lvl7pPr != nil {
		selvl7pPr := xml.StartElement{Name: xml.Name{Local: "a:lvl7pPr"}}
		e.EncodeElement(m.Lvl7pPr, selvl7pPr)
	}
	if m.Lvl8pPr != nil {
		selvl8pPr := xml.StartElement{Name: xml.Name{Local: "a:lvl8pPr"}}
		e.EncodeElement(m.Lvl8pPr, selvl8pPr)
	}
	if m.Lvl9pPr != nil {
		selvl9pPr := xml.StartElement{Name: xml.Name{Local: "a:lvl9pPr"}}
		e.EncodeElement(m.Lvl9pPr, selvl9pPr)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "a:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TextListStyle) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_TextListStyle:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "defPPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "defPPr"}:
				m.DefPPr = NewCT_TextParagraphProperties()
				if err := d.DecodeElement(m.DefPPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "lvl1pPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "lvl1pPr"}:
				m.Lvl1pPr = NewCT_TextParagraphProperties()
				if err := d.DecodeElement(m.Lvl1pPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "lvl2pPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "lvl2pPr"}:
				m.Lvl2pPr = NewCT_TextParagraphProperties()
				if err := d.DecodeElement(m.Lvl2pPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "lvl3pPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "lvl3pPr"}:
				m.Lvl3pPr = NewCT_TextParagraphProperties()
				if err := d.DecodeElement(m.Lvl3pPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "lvl4pPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "lvl4pPr"}:
				m.Lvl4pPr = NewCT_TextParagraphProperties()
				if err := d.DecodeElement(m.Lvl4pPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "lvl5pPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "lvl5pPr"}:
				m.Lvl5pPr = NewCT_TextParagraphProperties()
				if err := d.DecodeElement(m.Lvl5pPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "lvl6pPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "lvl6pPr"}:
				m.Lvl6pPr = NewCT_TextParagraphProperties()
				if err := d.DecodeElement(m.Lvl6pPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "lvl7pPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "lvl7pPr"}:
				m.Lvl7pPr = NewCT_TextParagraphProperties()
				if err := d.DecodeElement(m.Lvl7pPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "lvl8pPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "lvl8pPr"}:
				m.Lvl8pPr = NewCT_TextParagraphProperties()
				if err := d.DecodeElement(m.Lvl8pPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "lvl9pPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "lvl9pPr"}:
				m.Lvl9pPr = NewCT_TextParagraphProperties()
				if err := d.DecodeElement(m.Lvl9pPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "extLst"}:
				m.ExtLst = NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_TextListStyle %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TextListStyle
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TextListStyle and its children
func (m *CT_TextListStyle) Validate() error {
	return m.ValidateWithPath("CT_TextListStyle")
}

// ValidateWithPath validates the CT_TextListStyle and its children, prefixing error messages with path
func (m *CT_TextListStyle) ValidateWithPath(path string) error {
	if m.DefPPr != nil {
		if err := m.DefPPr.ValidateWithPath(path + "/DefPPr"); err != nil {
			return err
		}
	}
	if m.Lvl1pPr != nil {
		if err := m.Lvl1pPr.ValidateWithPath(path + "/Lvl1pPr"); err != nil {
			return err
		}
	}
	if m.Lvl2pPr != nil {
		if err := m.Lvl2pPr.ValidateWithPath(path + "/Lvl2pPr"); err != nil {
			return err
		}
	}
	if m.Lvl3pPr != nil {
		if err := m.Lvl3pPr.ValidateWithPath(path + "/Lvl3pPr"); err != nil {
			return err
		}
	}
	if m.Lvl4pPr != nil {
		if err := m.Lvl4pPr.ValidateWithPath(path + "/Lvl4pPr"); err != nil {
			return err
		}
	}
	if m.Lvl5pPr != nil {
		if err := m.Lvl5pPr.ValidateWithPath(path + "/Lvl5pPr"); err != nil {
			return err
		}
	}
	if m.Lvl6pPr != nil {
		if err := m.Lvl6pPr.ValidateWithPath(path + "/Lvl6pPr"); err != nil {
			return err
		}
	}
	if m.Lvl7pPr != nil {
		if err := m.Lvl7pPr.ValidateWithPath(path + "/Lvl7pPr"); err != nil {
			return err
		}
	}
	if m.Lvl8pPr != nil {
		if err := m.Lvl8pPr.ValidateWithPath(path + "/Lvl8pPr"); err != nil {
			return err
		}
	}
	if m.Lvl9pPr != nil {
		if err := m.Lvl9pPr.ValidateWithPath(path + "/Lvl9pPr"); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
