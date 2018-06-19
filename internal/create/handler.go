// Copyright (c) 2018 Uber Technologies, Inc.
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

package create

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/uber/prototool/internal/settings"
	"go.uber.org/zap"
)

var tmpl = template.Must(template.New("tmpl").Parse(`syntax = "proto3";

package {{.Pkg}};

option go_package = "{{.GoPkg}}";
option java_package = "{{.JavaPkg}}";`))

type tmplData struct {
	Pkg     string
	GoPkg   string
	JavaPkg string
}

type handler struct {
	logger         *zap.Logger
	configProvider settings.ConfigProvider
	pkg            string
}

func newHandler(options ...HandlerOption) *handler {
	handler := &handler{
		logger: zap.NewNop(),
	}
	for _, option := range options {
		option(handler)
	}
	handler.configProvider = settings.NewConfigProvider(
		settings.ConfigProviderWithLogger(handler.logger),
	)
	return handler
}

func (h *handler) Create(filePaths ...string) error {
	for _, filePath := range filePaths {
		if err := h.checkFilePath(filePath); err != nil {
			return err
		}
	}
	for _, filePath := range filePaths {
		if err := h.create(filePath); err != nil {
			return err
		}
	}
	return nil
}

func (h *handler) checkFilePath(filePath string) error {
	if filePath == "" {
		return errors.New("filePath empty")
	}
	dirPath := filepath.Dir(filePath)

	fileInfo, err := os.Stat(dirPath)
	if err != nil {
		return err
	}
	if !fileInfo.IsDir() {
		return fmt.Errorf("%q is not a directory somehow", dirPath)
	}
	if _, err := os.Stat(filePath); err == nil {
		return fmt.Errorf("%q already exists", filePath)
	}
	return nil
}

func (h *handler) create(filePath string) error {
	pkg, err := h.getPkg(filePath)
	if err != nil {
		return err
	}
	data, err := getData(
		&tmplData{
			Pkg:     pkg,
			GoPkg:   getGoPkg(pkg),
			JavaPkg: getJavaPkg(pkg),
		},
	)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filePath, data, 0644)
}

func (h *handler) getPkg(filePath string) (string, error) {
	absFilePath, err := filepath.Abs(filePath)
	if err != nil {
		return "", err
	}
	absDirPath := filepath.Dir(absFilePath)
	config, err := h.configProvider.GetForDir(absDirPath)
	if err != nil {
		return "", err
	}
	// no config file found
	if config.DirPath == "" {
		return DefaultPackage, nil
	}
	if config.DirPath == absDirPath {
		return DefaultPackage, nil
	}
	rel, err := filepath.Rel(config.DirPath, absDirPath)
	if err != nil {
		return "", err
	}
	// TODO
	if rel == "." {
		return "", fmt.Errorf("big problem")
	}
	return strings.Join(strings.Split(rel, string(os.PathSeparator)), "."), nil
}

func getGoPkg(pkg string) string {
	split := strings.Split(pkg, ".")
	return split[len(split)-1] + "pb"
}

func getJavaPkg(pkg string) string {
	return "com." + pkg + ".pb"
}

func getData(tmplData *tmplData) ([]byte, error) {
	buffer := bytes.NewBuffer(nil)
	if err := tmpl.Execute(buffer, tmplData); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}
