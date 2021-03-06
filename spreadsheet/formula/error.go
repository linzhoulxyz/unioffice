// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package formula

type Error struct {
	s string
}

func NewError(v string) Expression {
	return Error{v}
}

func (e Error) Eval(ctx Context, ev Evaluator) Result {
	return MakeErrorResult(e.s)
}

func (e Error) Reference(ctx Context, ev Evaluator) Reference {
	return ReferenceInvalid
}
