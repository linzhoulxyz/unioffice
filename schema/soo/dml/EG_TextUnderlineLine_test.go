// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package dml_test

import (
	"encoding/xml"
	"testing"

	"github.com/linzhoulxyz/unioffice/schema/soo/dml"
)

func TestEG_TextUnderlineLineConstructor(t *testing.T) {
	v := dml.NewEG_TextUnderlineLine()
	if v == nil {
		t.Errorf("dml.NewEG_TextUnderlineLine must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.EG_TextUnderlineLine should validate: %s", err)
	}
}

func TestEG_TextUnderlineLineMarshalUnmarshal(t *testing.T) {
	v := dml.NewEG_TextUnderlineLine()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewEG_TextUnderlineLine()
	xml.Unmarshal(buf, v2)
}
