// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package diagram_test

import (
	"encoding/xml"
	"testing"

	"github.com/danangsigit/gooxml/schema/soo/dml/diagram"
)

func TestCT_ColorTransformHeaderConstructor(t *testing.T) {
	v := diagram.NewCT_ColorTransformHeader()
	if v == nil {
		t.Errorf("diagram.NewCT_ColorTransformHeader must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_ColorTransformHeader should validate: %s", err)
	}
}

func TestCT_ColorTransformHeaderMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_ColorTransformHeader()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_ColorTransformHeader()
	xml.Unmarshal(buf, v2)
}
