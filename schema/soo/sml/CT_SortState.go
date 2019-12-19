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

type CT_SortState struct {
	// Sort by Columns
	ColumnSortAttr *bool
	// Case Sensitive
	CaseSensitiveAttr *bool
	// Sort Method
	SortMethodAttr ST_SortMethod
	// Sort Range
	RefAttr string
	// Sort Condition
	SortCondition []*CT_SortCondition
	ExtLst        *CT_ExtensionList
}

func NewCT_SortState() *CT_SortState {
	ret := &CT_SortState{}
	return ret
}

func (m *CT_SortState) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.ColumnSortAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "columnSort"},
			Value: fmt.Sprintf("%d", b2i(*m.ColumnSortAttr))})
	}
	if m.CaseSensitiveAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "caseSensitive"},
			Value: fmt.Sprintf("%d", b2i(*m.CaseSensitiveAttr))})
	}
	if m.SortMethodAttr != ST_SortMethodUnset {
		attr, err := m.SortMethodAttr.MarshalXMLAttr(xml.Name{Local: "sortMethod"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ref"},
		Value: fmt.Sprintf("%v", m.RefAttr)})
	e.EncodeToken(start)
	if m.SortCondition != nil {
		sesortCondition := xml.StartElement{Name: xml.Name{Local: "ma:sortCondition"}}
		for _, c := range m.SortCondition {
			e.EncodeElement(c, sesortCondition)
		}
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "ma:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_SortState) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "columnSort" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ColumnSortAttr = &parsed
			continue
		}
		if attr.Name.Local == "caseSensitive" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.CaseSensitiveAttr = &parsed
			continue
		}
		if attr.Name.Local == "sortMethod" {
			m.SortMethodAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "ref" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RefAttr = parsed
			continue
		}
	}
lCT_SortState:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "sortCondition"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "sortCondition"}:
				tmp := NewCT_SortCondition()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.SortCondition = append(m.SortCondition, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_SortState %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_SortState
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_SortState and its children
func (m *CT_SortState) Validate() error {
	return m.ValidateWithPath("CT_SortState")
}

// ValidateWithPath validates the CT_SortState and its children, prefixing error messages with path
func (m *CT_SortState) ValidateWithPath(path string) error {
	if err := m.SortMethodAttr.ValidateWithPath(path + "/SortMethodAttr"); err != nil {
		return err
	}
	for i, v := range m.SortCondition {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/SortCondition[%d]", path, i)); err != nil {
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
