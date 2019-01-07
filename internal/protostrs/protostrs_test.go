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

package protostrs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGoPackage(t *testing.T) {
	assert.Equal(t, "", GoPackage(""))
	assert.Equal(t, "foopb", GoPackage("foo"))
	assert.Equal(t, "barpb", GoPackage("foo.bar"))
}

func TestGoPackageLastTwo(t *testing.T) {
	assert.Equal(t, "", GoPackageLastTwo(""))
	assert.Equal(t, "foo", GoPackageLastTwo("foo"))
	assert.Equal(t, "foobar", GoPackageLastTwo("foo.bar"))
	assert.Equal(t, "foobar", GoPackageLastTwo("first.foo.bar"))
}

func TestJavaOuterClassname(t *testing.T) {
	assert.Equal(t, "", JavaOuterClassname(""))
	assert.Equal(t, "FileProto", JavaOuterClassname("file.proto"))
	assert.Equal(t, "FileProto", JavaOuterClassname("file.txt"))
	assert.Equal(t, "FileProto", JavaOuterClassname("a/file.proto"))
	assert.Equal(t, "FileProto", JavaOuterClassname("a/b/file.proto"))
	assert.Equal(t, "FileOneProto", JavaOuterClassname("a/b/file_one.proto"))
	assert.Equal(t, "FileOneProto", JavaOuterClassname("a/b/file-one.proto"))
	assert.Equal(t, "FileOneProto", JavaOuterClassname("a/b/file one.proto"))
	assert.Equal(t, "FiLeOneTwoProto", JavaOuterClassname("a/b/fiLe_One_two.proto"))
	assert.Equal(t, "FileOneProto", JavaOuterClassname("a/b/file one.txt"))
}

func TestJavaPackage(t *testing.T) {
	assert.Equal(t, "", JavaPackage(""))
	assert.Equal(t, "com.foo", JavaPackage("foo"))
	assert.Equal(t, "com.foo.bar", JavaPackage("foo.bar"))
}

func TestMajorBetaVersion(t *testing.T) {
	testMajorBetaVersionValid(t, "foo.v1", 1, 0)
	testMajorBetaVersionValid(t, "foo.bar.v1", 1, 0)
	testMajorBetaVersionValid(t, "foo.bar.v18", 18, 0)
	testMajorBetaVersionValid(t, "foo.bar.v180", 180, 0)
	testMajorBetaVersionInvalid(t, "foo.v0")
	testMajorBetaVersionInvalid(t, "foo.barv1")
	testMajorBetaVersionInvalid(t, "barv1")
	testMajorBetaVersionInvalid(t, "v1")
	testMajorBetaVersionInvalid(t, "foo.barv")
	testMajorBetaVersionInvalid(t, "barv")
	testMajorBetaVersionInvalid(t, "v")
	testMajorBetaVersionInvalid(t, "foo.bar.v-1")
	testMajorBetaVersionValid(t, "foo.v1beta1", 1, 1)
	testMajorBetaVersionValid(t, "foo.bar.v1beta1", 1, 1)
	testMajorBetaVersionValid(t, "foo.bar.v18beta18", 18, 18)
	testMajorBetaVersionValid(t, "foo.bar.v180beta180", 180, 180)
	testMajorBetaVersionInvalid(t, "foo.v0beta0")
	testMajorBetaVersionInvalid(t, "foo.v1beta0")
	testMajorBetaVersionInvalid(t, "foo.v0beta1")
	testMajorBetaVersionInvalid(t, "foo.barv1beta1")
	testMajorBetaVersionInvalid(t, "barv1beta1")
	testMajorBetaVersionInvalid(t, "v1beta1")
	testMajorBetaVersionInvalid(t, "foo.barvbeta1")
	testMajorBetaVersionInvalid(t, "barvbeta1")
	testMajorBetaVersionInvalid(t, "vbeta1")
	testMajorBetaVersionInvalid(t, "foo.bar.v-1beta1")
	testMajorBetaVersionInvalid(t, "foo.v1beta")
	testMajorBetaVersionInvalid(t, "foo.beta1")
}

func testMajorBetaVersionValid(t *testing.T, packageName string, expectedMajorBetaVersion uint64, expectedBetaVersion uint64) {
	majorVersion, betaVersion, ok := MajorBetaVersion(packageName)
	assert.True(t, ok, packageName)
	assert.Equal(t, expectedMajorBetaVersion, majorVersion, packageName)
	assert.Equal(t, expectedBetaVersion, betaVersion, packageName)
}

func testMajorBetaVersionInvalid(t *testing.T, packageName string) {
	_, _, ok := MajorBetaVersion(packageName)
	assert.False(t, ok, packageName)
}
