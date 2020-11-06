package wml

import (
	"encoding/xml"
	"fmt"
)

type CT_Choice struct {
	Drawing *CT_Drawing
}

func NewCT_Choice() *CT_Choice {
	ret := &CT_Choice{}
	return ret
}

func (m *CT_Choice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Choice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_Choice: %s", err)
		}

		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "drawing"}:
				m.Drawing = NewCT_Drawing()
				if err := d.DecodeElement(m.Drawing, &el); err != nil {
					return err
				}
			}
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_Choice and its children
func (m *CT_Choice) Validate() error {
	return m.ValidateWithPath("CT_Choice")
}

// ValidateWithPath validates the CT_Choice and its children, prefixing error messages with path
func (m *CT_Choice) ValidateWithPath(path string) error {
	return nil
}
