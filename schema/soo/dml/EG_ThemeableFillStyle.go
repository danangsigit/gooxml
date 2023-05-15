// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package dml

import (
	"encoding/xml"

	"github.com/danangsigit/gooxml"
)

type EG_ThemeableFillStyle struct {
	Fill    *CT_FillProperties
	FillRef *CT_StyleMatrixReference
}

func NewEG_ThemeableFillStyle() *EG_ThemeableFillStyle {
	ret := &EG_ThemeableFillStyle{}
	return ret
}

func (m *EG_ThemeableFillStyle) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.Fill != nil {
		sefill := xml.StartElement{Name: xml.Name{Local: "a:fill"}}
		e.EncodeElement(m.Fill, sefill)
	}
	if m.FillRef != nil {
		sefillRef := xml.StartElement{Name: xml.Name{Local: "a:fillRef"}}
		e.EncodeElement(m.FillRef, sefillRef)
	}
	return nil
}

func (m *EG_ThemeableFillStyle) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_ThemeableFillStyle:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "fill"}:
				m.Fill = NewCT_FillProperties()
				if err := d.DecodeElement(m.Fill, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "fillRef"}:
				m.FillRef = NewCT_StyleMatrixReference()
				if err := d.DecodeElement(m.FillRef, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on EG_ThemeableFillStyle %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_ThemeableFillStyle
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the EG_ThemeableFillStyle and its children
func (m *EG_ThemeableFillStyle) Validate() error {
	return m.ValidateWithPath("EG_ThemeableFillStyle")
}

// ValidateWithPath validates the EG_ThemeableFillStyle and its children, prefixing error messages with path
func (m *EG_ThemeableFillStyle) ValidateWithPath(path string) error {
	if m.Fill != nil {
		if err := m.Fill.ValidateWithPath(path + "/Fill"); err != nil {
			return err
		}
	}
	if m.FillRef != nil {
		if err := m.FillRef.ValidateWithPath(path + "/FillRef"); err != nil {
			return err
		}
	}
	return nil
}