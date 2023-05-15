// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wml_test

import (
	"encoding/xml"
	"testing"

	"github.com/danangsigit/gooxml/schema/soo/wml"
)

func TestCT_TcPrChangeConstructor(t *testing.T) {
	v := wml.NewCT_TcPrChange()
	if v == nil {
		t.Errorf("wml.NewCT_TcPrChange must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_TcPrChange should validate: %s", err)
	}
}

func TestCT_TcPrChangeMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_TcPrChange()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_TcPrChange()
	xml.Unmarshal(buf, v2)
}
