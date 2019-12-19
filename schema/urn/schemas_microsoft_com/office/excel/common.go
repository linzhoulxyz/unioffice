// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package excel

import (
	"encoding/xml"
	"fmt"

	"github.com/linzhoulxyz/unioffice"
)

type ST_ObjectType byte

const (
	ST_ObjectTypeUnset    ST_ObjectType = 0
	ST_ObjectTypeButton   ST_ObjectType = 1
	ST_ObjectTypeCheckbox ST_ObjectType = 2
	ST_ObjectTypeDialog   ST_ObjectType = 3
	ST_ObjectTypeDrop     ST_ObjectType = 4
	ST_ObjectTypeEdit     ST_ObjectType = 5
	ST_ObjectTypeGBox     ST_ObjectType = 6
	ST_ObjectTypeLabel    ST_ObjectType = 7
	ST_ObjectTypeLineA    ST_ObjectType = 8
	ST_ObjectTypeList     ST_ObjectType = 9
	ST_ObjectTypeMovie    ST_ObjectType = 10
	ST_ObjectTypeNote     ST_ObjectType = 11
	ST_ObjectTypePict     ST_ObjectType = 12
	ST_ObjectTypeRadio    ST_ObjectType = 13
	ST_ObjectTypeRectA    ST_ObjectType = 14
	ST_ObjectTypeScroll   ST_ObjectType = 15
	ST_ObjectTypeSpin     ST_ObjectType = 16
	ST_ObjectTypeShape    ST_ObjectType = 17
	ST_ObjectTypeGroup    ST_ObjectType = 18
	ST_ObjectTypeRect     ST_ObjectType = 19
)

func (e ST_ObjectType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_ObjectTypeUnset:
		attr.Value = ""
	case ST_ObjectTypeButton:
		attr.Value = "Button"
	case ST_ObjectTypeCheckbox:
		attr.Value = "Checkbox"
	case ST_ObjectTypeDialog:
		attr.Value = "Dialog"
	case ST_ObjectTypeDrop:
		attr.Value = "Drop"
	case ST_ObjectTypeEdit:
		attr.Value = "Edit"
	case ST_ObjectTypeGBox:
		attr.Value = "GBox"
	case ST_ObjectTypeLabel:
		attr.Value = "Label"
	case ST_ObjectTypeLineA:
		attr.Value = "LineA"
	case ST_ObjectTypeList:
		attr.Value = "List"
	case ST_ObjectTypeMovie:
		attr.Value = "Movie"
	case ST_ObjectTypeNote:
		attr.Value = "Note"
	case ST_ObjectTypePict:
		attr.Value = "Pict"
	case ST_ObjectTypeRadio:
		attr.Value = "Radio"
	case ST_ObjectTypeRectA:
		attr.Value = "RectA"
	case ST_ObjectTypeScroll:
		attr.Value = "Scroll"
	case ST_ObjectTypeSpin:
		attr.Value = "Spin"
	case ST_ObjectTypeShape:
		attr.Value = "Shape"
	case ST_ObjectTypeGroup:
		attr.Value = "Group"
	case ST_ObjectTypeRect:
		attr.Value = "Rect"
	}
	return attr, nil
}

func (e *ST_ObjectType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "Button":
		*e = 1
	case "Checkbox":
		*e = 2
	case "Dialog":
		*e = 3
	case "Drop":
		*e = 4
	case "Edit":
		*e = 5
	case "GBox":
		*e = 6
	case "Label":
		*e = 7
	case "LineA":
		*e = 8
	case "List":
		*e = 9
	case "Movie":
		*e = 10
	case "Note":
		*e = 11
	case "Pict":
		*e = 12
	case "Radio":
		*e = 13
	case "RectA":
		*e = 14
	case "Scroll":
		*e = 15
	case "Spin":
		*e = 16
	case "Shape":
		*e = 17
	case "Group":
		*e = 18
	case "Rect":
		*e = 19
	}
	return nil
}

func (m ST_ObjectType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_ObjectType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		*m = 1
		return nil
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "Button":
			*m = 1
		case "Checkbox":
			*m = 2
		case "Dialog":
			*m = 3
		case "Drop":
			*m = 4
		case "Edit":
			*m = 5
		case "GBox":
			*m = 6
		case "Label":
			*m = 7
		case "LineA":
			*m = 8
		case "List":
			*m = 9
		case "Movie":
			*m = 10
		case "Note":
			*m = 11
		case "Pict":
			*m = 12
		case "Radio":
			*m = 13
		case "RectA":
			*m = 14
		case "Scroll":
			*m = 15
		case "Spin":
			*m = 16
		case "Shape":
			*m = 17
		case "Group":
			*m = 18
		case "Rect":
			*m = 19
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_ObjectType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "Button"
	case 2:
		return "Checkbox"
	case 3:
		return "Dialog"
	case 4:
		return "Drop"
	case 5:
		return "Edit"
	case 6:
		return "GBox"
	case 7:
		return "Label"
	case 8:
		return "LineA"
	case 9:
		return "List"
	case 10:
		return "Movie"
	case 11:
		return "Note"
	case 12:
		return "Pict"
	case 13:
		return "Radio"
	case 14:
		return "RectA"
	case 15:
		return "Scroll"
	case 16:
		return "Spin"
	case 17:
		return "Shape"
	case 18:
		return "Group"
	case 19:
		return "Rect"
	}
	return ""
}

func (m ST_ObjectType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_ObjectType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

// init registers constructor functions for dynamically creating elements based off the XML namespace and name
func init() {
	unioffice.RegisterConstructor("urn:schemas-microsoft-com:office:excel", "CT_ClientData", NewCT_ClientData)
	unioffice.RegisterConstructor("urn:schemas-microsoft-com:office:excel", "ClientData", NewClientData)
}
