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

package extract

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/dapperlabs/prototool/internal/reflect"
	ptesting "github.com/dapperlabs/prototool/internal/testing"
)

func TestOne(t *testing.T) {
	packageSet := requireGetPackageSet(t, "one")
	packageNameToPackage := packageSet.PackageNameToPackage()
	_, ok := packageNameToPackage["uber.proto.foo.v1"]
	require.True(t, ok)
	_, ok = packageNameToPackage["uber.proto.bar.v1"]
	require.True(t, ok)
}

func requireGetPackageSet(t *testing.T, subDirPath string) *PackageSet {
	packageSet, err := getPackageSet(subDirPath)
	require.NoError(t, err)
	return packageSet
}

func getPackageSet(subDirPath string) (*PackageSet, error) {
	fileDescriptorSets, err := ptesting.GetFileDescriptorSets(".", "testdata/"+subDirPath)
	if err != nil {
		return nil, err
	}
	reflectPackageSet, err := reflect.NewPackageSet(fileDescriptorSets.Unwrap()...)
	if err != nil {
		return nil, err
	}
	return NewPackageSet(reflectPackageSet)
}
