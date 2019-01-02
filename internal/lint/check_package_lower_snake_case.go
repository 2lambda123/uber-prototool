// Copyright (c) 2019 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package lint

import (
	"strings"

	"github.com/emicklei/proto"
	"github.com/uber/prototool/internal/strs"
	"github.com/uber/prototool/internal/text"
)

var packageLowerSnakeCaseLinter = NewLinter(
	"PACKAGE_LOWER_SNAKE_CASE",
	"Verifies that the package is lower_snake.case.",
	checkPackageLowerSnakeCase,
)

func checkPackageLowerSnakeCase(add func(*text.Failure), dirPath string, descriptors []*proto.Proto) error {
	return runVisitor(&packageLowerSnakeCaseVisitor{baseAddVisitor: newBaseAddVisitor(add)}, descriptors)
}

type packageLowerSnakeCaseVisitor struct {
	baseAddVisitor

	filename string
	pkg      *proto.Package
}

func (v *packageLowerSnakeCaseVisitor) OnStart(descriptor *proto.Proto) error {
	v.filename = descriptor.Filename
	v.pkg = nil
	return nil
}

func (v *packageLowerSnakeCaseVisitor) VisitPackage(pkg *proto.Package) {
	if v.pkg != nil {
		v.AddFailuref(pkg.Position, "multiple package declarations, first was %v", v.pkg)
		return
	}
	v.pkg = pkg
}

func (v *packageLowerSnakeCaseVisitor) Finally() error {
	if v.pkg != nil {
		if !strs.IsLowerSnakeCase(strings.Replace(v.pkg.Name, ".", "", -1)) {
			v.AddFailuref(v.pkg.Position, "Package should be lower_snake.case but was %q.", v.pkg.Name)
		}
	}
	return nil
}
