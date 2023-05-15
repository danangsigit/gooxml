// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package math_test

import (
	"encoding/xml"
	"testing"

	"github.com/danangsigit/gooxml/schema/soo/ofc/math"
)

func TestCT_FuncPrConstructor(t *testing.T) {
	v := math.NewCT_FuncPr()
	if v == nil {
		t.Errorf("math.NewCT_FuncPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_FuncPr should validate: %s", err)
	}
}

func TestCT_FuncPrMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_FuncPr()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_FuncPr()
	xml.Unmarshal(buf, v2)
}
