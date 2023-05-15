// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wml

import (
	"encoding/xml"
	"fmt"

	"github.com/danangsigit/gooxml"
)

type CT_Tabs struct {
	// Custom Tab Stop
	Tab []*CT_TabStop
}

func NewCT_Tabs() *CT_Tabs {
	ret := &CT_Tabs{}
	return ret
}

func (m *CT_Tabs) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	setab := xml.StartElement{Name: xml.Name{Local: "w:tab"}}
	for _, c := range m.Tab {
		e.EncodeElement(c, setab)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Tabs) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Tabs:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "tab"}:
				tmp := NewCT_TabStop()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Tab = append(m.Tab, tmp)
			default:
				gooxml.Log("skipping unsupported element on CT_Tabs %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Tabs
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Tabs and its children
func (m *CT_Tabs) Validate() error {
	return m.ValidateWithPath("CT_Tabs")
}

// ValidateWithPath validates the CT_Tabs and its children, prefixing error messages with path
func (m *CT_Tabs) ValidateWithPath(path string) error {
	for i, v := range m.Tab {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Tab[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
