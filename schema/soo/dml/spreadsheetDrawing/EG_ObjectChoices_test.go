// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package spreadsheetDrawing_test

import (
	"encoding/xml"
	"testing"

	"github.com/linzhoulxyz/unioffice/schema/soo/dml/spreadsheetDrawing"
)

func TestEG_ObjectChoicesConstructor(t *testing.T) {
	v := spreadsheetDrawing.NewEG_ObjectChoices()
	if v == nil {
		t.Errorf("spreadsheetDrawing.NewEG_ObjectChoices must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed spreadsheetDrawing.EG_ObjectChoices should validate: %s", err)
	}
}

func TestEG_ObjectChoicesMarshalUnmarshal(t *testing.T) {
	v := spreadsheetDrawing.NewEG_ObjectChoices()
	buf, _ := xml.Marshal(v)
	v2 := spreadsheetDrawing.NewEG_ObjectChoices()
	xml.Unmarshal(buf, v2)
}
