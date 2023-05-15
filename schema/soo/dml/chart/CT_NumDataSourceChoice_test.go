// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart_test

import (
	"encoding/xml"
	"testing"

	"github.com/danangsigit/gooxml/schema/soo/dml/chart"
)

func TestCT_NumDataSourceChoiceConstructor(t *testing.T) {
	v := chart.NewCT_NumDataSourceChoice()
	if v == nil {
		t.Errorf("chart.NewCT_NumDataSourceChoice must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_NumDataSourceChoice should validate: %s", err)
	}
}

func TestCT_NumDataSourceChoiceMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_NumDataSourceChoice()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_NumDataSourceChoice()
	xml.Unmarshal(buf, v2)
}
