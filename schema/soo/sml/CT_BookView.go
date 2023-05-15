// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/danangsigit/gooxml"
)

type CT_BookView struct {
	// Visibility
	VisibilityAttr ST_Visibility
	// Minimized
	MinimizedAttr *bool
	// Show Horizontal Scroll
	ShowHorizontalScrollAttr *bool
	// Show Vertical Scroll
	ShowVerticalScrollAttr *bool
	// Show Sheet Tabs
	ShowSheetTabsAttr *bool
	// Upper Left Corner (X Coordinate)
	XWindowAttr *int32
	// Upper Left Corner (Y Coordinate)
	YWindowAttr *int32
	// Window Width
	WindowWidthAttr *uint32
	// Window Height
	WindowHeightAttr *uint32
	// Sheet Tab Ratio
	TabRatioAttr *uint32
	// First Sheet
	FirstSheetAttr *uint32
	// Active Sheet Index
	ActiveTabAttr *uint32
	// AutoFilter Date Grouping
	AutoFilterDateGroupingAttr *bool
	ExtLst                     *CT_ExtensionList
}

func NewCT_BookView() *CT_BookView {
	ret := &CT_BookView{}
	return ret
}

func (m *CT_BookView) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.VisibilityAttr != ST_VisibilityUnset {
		attr, err := m.VisibilityAttr.MarshalXMLAttr(xml.Name{Local: "visibility"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.MinimizedAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "minimized"},
			Value: fmt.Sprintf("%d", b2i(*m.MinimizedAttr))})
	}
	if m.ShowHorizontalScrollAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showHorizontalScroll"},
			Value: fmt.Sprintf("%d", b2i(*m.ShowHorizontalScrollAttr))})
	}
	if m.ShowVerticalScrollAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showVerticalScroll"},
			Value: fmt.Sprintf("%d", b2i(*m.ShowVerticalScrollAttr))})
	}
	if m.ShowSheetTabsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showSheetTabs"},
			Value: fmt.Sprintf("%d", b2i(*m.ShowSheetTabsAttr))})
	}
	if m.XWindowAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xWindow"},
			Value: fmt.Sprintf("%v", *m.XWindowAttr)})
	}
	if m.YWindowAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "yWindow"},
			Value: fmt.Sprintf("%v", *m.YWindowAttr)})
	}
	if m.WindowWidthAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "windowWidth"},
			Value: fmt.Sprintf("%v", *m.WindowWidthAttr)})
	}
	if m.WindowHeightAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "windowHeight"},
			Value: fmt.Sprintf("%v", *m.WindowHeightAttr)})
	}
	if m.TabRatioAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "tabRatio"},
			Value: fmt.Sprintf("%v", *m.TabRatioAttr)})
	}
	if m.FirstSheetAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "firstSheet"},
			Value: fmt.Sprintf("%v", *m.FirstSheetAttr)})
	}
	if m.ActiveTabAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "activeTab"},
			Value: fmt.Sprintf("%v", *m.ActiveTabAttr)})
	}
	if m.AutoFilterDateGroupingAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "autoFilterDateGrouping"},
			Value: fmt.Sprintf("%d", b2i(*m.AutoFilterDateGroupingAttr))})
	}
	e.EncodeToken(start)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "ma:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_BookView) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "windowWidth" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.WindowWidthAttr = &pt
			continue
		}
		if attr.Name.Local == "visibility" {
			m.VisibilityAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "showHorizontalScroll" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowHorizontalScrollAttr = &parsed
			continue
		}
		if attr.Name.Local == "showVerticalScroll" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowVerticalScrollAttr = &parsed
			continue
		}
		if attr.Name.Local == "showSheetTabs" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowSheetTabsAttr = &parsed
			continue
		}
		if attr.Name.Local == "xWindow" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.XWindowAttr = &pt
			continue
		}
		if attr.Name.Local == "yWindow" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.YWindowAttr = &pt
			continue
		}
		if attr.Name.Local == "minimized" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.MinimizedAttr = &parsed
			continue
		}
		if attr.Name.Local == "windowHeight" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.WindowHeightAttr = &pt
			continue
		}
		if attr.Name.Local == "tabRatio" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.TabRatioAttr = &pt
			continue
		}
		if attr.Name.Local == "firstSheet" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.FirstSheetAttr = &pt
			continue
		}
		if attr.Name.Local == "activeTab" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.ActiveTabAttr = &pt
			continue
		}
		if attr.Name.Local == "autoFilterDateGrouping" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AutoFilterDateGroupingAttr = &parsed
			continue
		}
	}
lCT_BookView:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_BookView %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_BookView
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_BookView and its children
func (m *CT_BookView) Validate() error {
	return m.ValidateWithPath("CT_BookView")
}

// ValidateWithPath validates the CT_BookView and its children, prefixing error messages with path
func (m *CT_BookView) ValidateWithPath(path string) error {
	if err := m.VisibilityAttr.ValidateWithPath(path + "/VisibilityAttr"); err != nil {
		return err
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
