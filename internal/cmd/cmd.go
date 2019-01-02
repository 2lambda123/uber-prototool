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

// Package cmd contains the logic to setup Prototool with github.com/spf13/cobra.
//
// The packages internal/cmd/prototool, internal/gen/gen-prototool-bash-completion,
// internal/gen/gen-prototool-manpages and internal/gen/gen-prototool-zsh-completion
// are main packages that call into this package, and this package calls into
// internal/exec to execute the logic.
//
// This package also contains integration testing for Prototool.
package cmd

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"github.com/uber/prototool/internal/exec"
)

// when generating man pages, the current date is used
// this means every time we run make gen, a diff is created
// this gets extremely annoying and isn't very useful so we make it static here
// we could also not check in the man pages, but for now we have them checked in
var genManTime = time.Date(2018, time.January, 1, 0, 0, 0, 0, time.UTC)

// Do runs the command logic.
func Do(args []string, stdin io.Reader, stdout io.Writer, stderr io.Writer) int {
	return do(false, args, stdin, stdout, stderr)
}

func do(develMode bool, args []string, stdin io.Reader, stdout io.Writer, stderr io.Writer) int {
	return runRootCommand(develMode, args, stdin, stdout, stderr, (*cobra.Command).Execute)
}

// GenBashCompletion generates a bash completion file to the writer.
func GenBashCompletion(stdin io.Reader, stdout io.Writer, stderr io.Writer) int {
	return runRootCommandOutput(false, []string{}, stdin, stdout, stderr, (*cobra.Command).GenBashCompletion)
}

// GenZshCompletion generates a zsh completion file to the writer.
func GenZshCompletion(stdin io.Reader, stdout io.Writer, stderr io.Writer) int {
	return runRootCommandOutput(false, []string{}, stdin, stdout, stderr, (*cobra.Command).GenZshCompletion)
}

// GenManpages generates the manpages to the given directory.
func GenManpages(args []string, stdin io.Reader, stdout io.Writer, stderr io.Writer) int {
	return runRootCommand(false, args, stdin, stdout, stderr, func(cmd *cobra.Command) error {
		if len(args) != 1 {
			return fmt.Errorf("usage: %s dirPath", os.Args[0])
		}
		return doc.GenManTree(cmd, &doc.GenManHeader{
			Date: &genManTime,
			// Otherwise we get an annoying "Auto generated by spf13/cobra"
			// Maybe we want that, but I think it's better to just have this
			Source: "Prototool",
		}, args[0])
	})
}

// develMode turns on sub-commands and potentially flags that we do not expose during the build of the prototool binary
func runRootCommandOutput(develMode bool, args []string, stdin io.Reader, stdout io.Writer, stderr io.Writer, f func(*cobra.Command, io.Writer) error) int {
	return runRootCommand(develMode, args, stdin, stdout, stderr, func(cmd *cobra.Command) error { return f(cmd, stdout) })
}

// develMode turns on sub-commands and potentially flags that we do not expose during the build of the prototool binary
func runRootCommand(develMode bool, args []string, stdin io.Reader, stdout io.Writer, stderr io.Writer, f func(*cobra.Command) error) (exitCode int) {
	if err := checkOS(); err != nil {
		return printAndGetErrorExitCode(err, stdout)
	}
	if err := f(getRootCommand(&exitCode, develMode, args, stdin, stdout, stderr)); err != nil {
		return printAndGetErrorExitCode(err, stdout)
	}
	return exitCode
}

func getRootCommand(exitCodeAddr *int, develMode bool, args []string, stdin io.Reader, stdout io.Writer, stderr io.Writer) *cobra.Command {
	flags := &flags{}

	rootCmd := &cobra.Command{Use: "prototool"}
	rootCmd.AddCommand(allCmdTemplate.Build(exitCodeAddr, stdin, stdout, stderr, flags))
	rootCmd.AddCommand(compileCmdTemplate.Build(exitCodeAddr, stdin, stdout, stderr, flags))
	rootCmd.AddCommand(createCmdTemplate.Build(exitCodeAddr, stdin, stdout, stderr, flags))
	rootCmd.AddCommand(filesCmdTemplate.Build(exitCodeAddr, stdin, stdout, stderr, flags))
	rootCmd.AddCommand(formatCmdTemplate.Build(exitCodeAddr, stdin, stdout, stderr, flags))
	rootCmd.AddCommand(generateCmdTemplate.Build(exitCodeAddr, stdin, stdout, stderr, flags))
	rootCmd.AddCommand(grpcCmdTemplate.Build(exitCodeAddr, stdin, stdout, stderr, flags))
	configCmd := &cobra.Command{Use: "config"}
	configCmd.AddCommand(configInitCmdTemplate.Build(exitCodeAddr, stdin, stdout, stderr, flags))
	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(lintCmdTemplate.Build(exitCodeAddr, stdin, stdout, stderr, flags))
	rootCmd.AddCommand(versionCmdTemplate.Build(exitCodeAddr, stdin, stdout, stderr, flags))
	inspectCmd := &cobra.Command{Use: "inspect", Short: "Top-level command for inspection commands."}
	inspectCmd.AddCommand(inspectPackagesCmdTemplate.Build(exitCodeAddr, stdin, stdout, stderr, flags))
	inspectCmd.AddCommand(inspectPackageCmdTemplate.Build(exitCodeAddr, stdin, stdout, stderr, flags))
	inspectCmd.AddCommand(inspectPackageDepsCmdTemplate.Build(exitCodeAddr, stdin, stdout, stderr, flags))
	inspectCmd.AddCommand(inspectPackageImportersCmdTemplate.Build(exitCodeAddr, stdin, stdout, stderr, flags))
	rootCmd.AddCommand(inspectCmd)

	// flags bound to rootCmd are global flags
	flags.bindDebug(rootCmd.PersistentFlags())

	if develMode {
		rootCmd.AddCommand(cleanCmdTemplate.Build(exitCodeAddr, stdin, stdout, stderr, flags))
		rootCmd.AddCommand(downloadCmdTemplate.Build(exitCodeAddr, stdin, stdout, stderr, flags))

		// we may or may not want to expose these to users
		// but will not build them into the binary for v1.0
		flags.bindCachePath(rootCmd.PersistentFlags())
		flags.bindPrintFields(rootCmd.PersistentFlags())
	}

	rootCmd.SetArgs(args)
	rootCmd.SetOutput(stdout)

	return rootCmd
}

func checkOS() error {
	switch runtime.GOOS {
	case "darwin", "linux":
		return nil
	default:
		return fmt.Errorf("%s is not a supported operating system", runtime.GOOS)
	}
}

func printAndGetErrorExitCode(err error, stdout io.Writer) int {
	if errString := err.Error(); errString != "" {
		_, _ = fmt.Fprintln(stdout, errString)
	}
	if exitError, ok := err.(*exec.ExitError); ok {
		return exitError.Code
	}
	return 1
}
