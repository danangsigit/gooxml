// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package pml_test

import (
	"encoding/xml"
	"testing"

	"github.com/danangsigit/gooxml/schema/soo/pml"
)

func TestCT_NormalViewPropertiesConstructor(t *testing.T) {
	v := pml.NewCT_NormalViewProperties()
	if v == nil {
		t.Errorf("pml.NewCT_NormalViewProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_NormalViewProperties should validate: %s", err)
	}
}

func TestCT_NormalViewPropertiesMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_NormalViewProperties()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_NormalViewProperties()
	xml.Unmarshal(buf, v2)
}
