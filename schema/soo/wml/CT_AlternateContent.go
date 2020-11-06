package wml

import (
	"encoding/xml"
	"fmt"
	"github.com/linzhoulxyz/unioffice"
)

type CT_AlternateContent struct {
	Choice *CT_Choice
	Extra  []unioffice.Any
}

func NewCT_AlternateContent() *CT_AlternateContent {
	ret := &CT_AlternateContent{}
	return ret
}

func (m *CT_AlternateContent) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_AlternateContent) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_AlternateContent: %s", err)
		}

		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/markup-compatibility/2006", Local: "Choice"}:
				m.Choice = NewCT_Choice()
				if err := d.DecodeElement(m.Choice, &el); err != nil {
					return err
				}
			default:
				any := &unioffice.XSDAny{}
				if err := d.DecodeElement(any, &el); err != nil {
					return err
				}
				m.Extra = append(m.Extra, any)
			}
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_AlternateContent and its children
func (m *CT_AlternateContent) Validate() error {
	return m.ValidateWithPath("CT_AlternateContent")
}

// ValidateWithPath validates the CT_AlternateContent and its children, prefixing error messages with path
func (m *CT_AlternateContent) ValidateWithPath(path string) error {
	return nil
}
