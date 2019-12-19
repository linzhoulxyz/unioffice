// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package content_types_test

import (
	"encoding/xml"
	"testing"

	"github.com/linzhoulxyz/unioffice/schema/soo/pkg/content_types"
)

func TestTypesConstructor(t *testing.T) {
	v := content_types.NewTypes()
	if v == nil {
		t.Errorf("content_types.NewTypes must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed content_types.Types should validate: %s", err)
	}
}

func TestTypesMarshalUnmarshal(t *testing.T) {
	v := content_types.NewTypes()
	buf, _ := xml.Marshal(v)
	v2 := content_types.NewTypes()
	xml.Unmarshal(buf, v2)
}
