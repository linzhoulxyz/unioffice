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

func TestCT_OfficeArtExtensionConstructor(t *testing.T) {
	v := dml.NewCT_OfficeArtExtension()
	if v == nil {
		t.Errorf("dml.NewCT_OfficeArtExtension must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_OfficeArtExtension should validate: %s", err)
	}
}

func TestCT_OfficeArtExtensionMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_OfficeArtExtension()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_OfficeArtExtension()
	xml.Unmarshal(buf, v2)
}
