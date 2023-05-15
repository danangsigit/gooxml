// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package docPropsVTypes

import (
	"encoding/xml"
	"strconv"
	"time"

	"github.com/danangsigit/gooxml"
)

type Vector struct {
	CT_Vector
}

func NewVector() *Vector {
	ret := &Vector{}
	ret.CT_Vector = *NewCT_Vector()
	return ret
}

func (m *Vector) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return m.CT_Vector.MarshalXML(e, start)
}

func (m *Vector) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_Vector = *NewCT_Vector()
	for _, attr := range start.Attr {
		if attr.Name.Local == "size" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.SizeAttr = uint32(parsed)
			continue
		}
		if attr.Name.Local == "baseType" {
			m.BaseTypeAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
lVector:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "variant"}:
				tmp := NewVariant()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Variant = append(m.Variant, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "i1"}:
				var tmp int8
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.I1 = append(m.I1, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "i2"}:
				var tmp int16
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.I2 = append(m.I2, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "i4"}:
				var tmp int32
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.I4 = append(m.I4, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "i8"}:
				var tmp int64
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.I8 = append(m.I8, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "ui1"}:
				var tmp uint8
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.Ui1 = append(m.Ui1, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "ui2"}:
				var tmp uint16
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.Ui2 = append(m.Ui2, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "ui4"}:
				var tmp uint32
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.Ui4 = append(m.Ui4, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "ui8"}:
				var tmp uint64
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.Ui8 = append(m.Ui8, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "r4"}:
				var tmp float32
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.R4 = append(m.R4, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "r8"}:
				var tmp float64
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.R8 = append(m.R8, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "lpstr"}:
				var tmp string
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.Lpstr = append(m.Lpstr, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "lpwstr"}:
				var tmp string
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.Lpwstr = append(m.Lpwstr, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "bstr"}:
				var tmp string
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.Bstr = append(m.Bstr, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "date"}:
				var tmp time.Time
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.Date = append(m.Date, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "filetime"}:
				var tmp time.Time
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.Filetime = append(m.Filetime, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "bool"}:
				var tmp bool
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.Bool = append(m.Bool, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "cy"}:
				var tmp string
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.Cy = append(m.Cy, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "error"}:
				var tmp string
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.Error = append(m.Error, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "clsid"}:
				var tmp string
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.Clsid = append(m.Clsid, tmp)
			default:
				gooxml.Log("skipping unsupported element on Vector %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lVector
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the Vector and its children
func (m *Vector) Validate() error {
	return m.ValidateWithPath("Vector")
}

// ValidateWithPath validates the Vector and its children, prefixing error messages with path
func (m *Vector) ValidateWithPath(path string) error {
	if err := m.CT_Vector.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}